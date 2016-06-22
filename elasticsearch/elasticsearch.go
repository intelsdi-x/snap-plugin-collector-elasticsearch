/*
http://www.apache.org/licenses/LICENSE-2.0.txt

Copyright 2016 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package elasticsearch

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/intelsdi-x/snap-plugin-utilities/config"
	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core"
)

const (
	// Name of plugin
	name = "elasticsearch"
	// Version of plugin
	version = 2
	// Type of plugin
	pluginType = plugin.CollectorPluginType

	// Timeout duration for HTTP connection
	timeout = 5 * time.Second

	esHost = "server"
	esPort = "port"

	invalidMetricType = "Invalid metric type found"
)

// Meta returns the snap plug.PluginMeta type
func Meta() *plugin.PluginMeta {
	return plugin.NewPluginMeta(name, version, pluginType, []string{plugin.SnapGOBContentType}, []string{plugin.SnapGOBContentType})
}

//  NewElasticsearchCollector returns a new instance of Elasticsearch struct
func NewElasticsearchCollector() *Elasticsearch {
	return &Elasticsearch{}
}

type Elasticsearch struct {
}

// CollectMetrics returns metrics from Elasticsearch
func (p *Elasticsearch) CollectMetrics(mts []plugin.MetricType) ([]plugin.MetricType, error) {
	metrics := []plugin.MetricType{}

	// flags to hit ES server only once
	// during one round of the collection
	hasNodeMetric := false
	hasClusterMetric := false

	var nodeStatsMap map[string]map[string]plugin.MetricType
	var clusterStatsMap map[string]plugin.MetricType
	var err error
	for _, m := range mts {
		switch m.Namespace()[2].Value {
		case "node":
			if !hasNodeMetric {
				nodeStatsMap, err = getNodeMetrics(mts[0])
				handleErr(err)
				hasNodeMetric = true
			}

			if m.Namespace()[3].Value == "*" {
				for i, node := range nodeStatsMap {
					m.Namespace()[3].Value = i
					metrics = append(metrics, node[strings.TrimLeft(m.Namespace().String(), "/")])
				}
			} else {
				for _, x := range nodeStatsMap {
					if value, ok := x[strings.TrimLeft(m.Namespace().String(), "/")]; ok {
						metrics = append(metrics, value)
						break
					}
				}
			}
		case "cluster":
			if !hasClusterMetric {
				clusterStatsMap, err = getClusterMetrics(mts[0])
				handleErr(err)
				hasClusterMetric = true
			}
			metrics = append(metrics, clusterStatsMap[strings.TrimLeft(m.Namespace().String(), "/")])
		default:
			// filter out the invalid metrics
			log.Println(invalidMetricType, m.Namespace())
		}

	}
	return metrics, nil
}

// GetMetricTypes returns the metric types exposed by Elasticsearch
func (p *Elasticsearch) GetMetricTypes(pct plugin.ConfigType) ([]plugin.MetricType, error) {
	mtsType, err := getMetrics(pct)
	handleErr(err)

	return mtsType, nil
}

// GetConfigPolicy returns a ConfigPolicy
func (p *Elasticsearch) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	c := cpolicy.New()
	return c, nil
}

func getNodeMetrics(pmt interface{}) (map[string]map[string]plugin.MetricType, error) {
	host := getServer(pmt)

	esNodeMetrics := NewESNodeMetric(host, timeout)
	mMap, err := esNodeMetrics.GetNodeData()
	if err != nil {
		return nil, err
	}

	return mMap, nil
}

func getClusterMetrics(pmt interface{}) (map[string]plugin.MetricType, error) {
	host := getServer(pmt)

	esClusterMetrics := NewESClusterMetric(host, timeout)
	metrics, err := esClusterMetrics.GetClusterData()
	if err != nil {
		return nil, err
	}
	return metrics, nil
}

func getMetrics(pct plugin.ConfigType) ([]plugin.MetricType, error) {
	mMap, err := getNodeMetrics(pct)
	handleErr(err)

	mts := []plugin.MetricType{}
	for _, n := range mMap {
		for ns, _ := range n {
			namespace := strings.Split(ns, "/")
			namespace[3] = "*"
			mts = append(mts, plugin.MetricType{Namespace_: core.NewNamespace(namespace...)})
		}
		break
	}

	metrics, err := getClusterMetrics(pct)
	handleErr(err)

	for _, m := range metrics {
		mts = append(mts, plugin.MetricType{Namespace_: m.Namespace()})
	}

	return mts, nil
}

func joinNamespace(ns []string) string {
	return "/" + strings.Join(ns, "/")
}

func handleErr(e error) {
	if e != nil {
		log.Fatal(e.Error())
	}
}

func getServer(cfg interface{}) string {
	items, err := config.GetConfigItems(cfg, esHost, esPort)
	if err != nil {
		log.Fatal(err.Error())
	}

	server := items[esHost].(string)
	port := items[esPort].(int)

	return server + ":" + strconv.Itoa(port)
}
