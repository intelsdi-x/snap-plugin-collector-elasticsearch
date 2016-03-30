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
)

const (
	// Name of plugin
	name = "elasticsearch"
	// Version of plugin
	version = 1
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
func (p *Elasticsearch) CollectMetrics(mts []plugin.PluginMetricType) ([]plugin.PluginMetricType, error) {
	metrics := []plugin.PluginMetricType{}

	// flags to hit ES server only once
	// during one round of the collection
	hasNodeMetric := false
	hasClusterMetric := false

	var nodeStatsMap map[string]map[string]plugin.PluginMetricType
	var clusterStatsMap map[string]plugin.PluginMetricType
	var err error
	for _, m := range mts {
		switch m.Namespace()[2] {
		case "node":
			if !hasNodeMetric {
				nodeStatsMap, err = getNodeMetrics(mts[0])
				handleErr(err)
				hasNodeMetric = true
			}

			if m.Namespace()[3] == "*" {
				for i, node := range nodeStatsMap {
					m.Namespace()[3] = i
					metrics = append(metrics, node[strings.Join(m.Namespace(), "/")])
				}
			} else {
				for _, x := range nodeStatsMap {
					if value, ok := x[strings.Join(m.Namespace(), "/")]; ok {
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
			metrics = append(metrics, clusterStatsMap[strings.Join(m.Namespace(), "/")])
		default:
			// filter out the invalid metrics
			log.Println(invalidMetricType, m.Namespace())
		}

	}

	return metrics, nil
}

// GetMetricTypes returns the metric types exposed by Elasticsearch
func (p *Elasticsearch) GetMetricTypes(pct plugin.PluginConfigType) ([]plugin.PluginMetricType, error) {
	mtsType, err := getMetrics(pct)
	handleErr(err)

	return mtsType, nil
}

// GetConfigPolicy returns a ConfigPolicy
func (p *Elasticsearch) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	c := cpolicy.New()
	return c, nil
}

func getNodeMetrics(pmt interface{}) (map[string]map[string]plugin.PluginMetricType, error) {
	host := getServer(pmt)

	esNodeMetrics := NewESNodeMetric(host, timeout)
	mMap, err := esNodeMetrics.GetNodeData()
	if err != nil {
		return nil, err
	}

	return mMap, nil
}

func getClusterMetrics(pmt interface{}) (map[string]plugin.PluginMetricType, error) {
	host := getServer(pmt)

	esClusterMetrics := NewESClusterMetric(host, timeout)
	metrics, err := esClusterMetrics.GetClusterData()
	if err != nil {
		return nil, err
	}
	return metrics, nil
}

func getMetrics(pct plugin.PluginConfigType) ([]plugin.PluginMetricType, error) {
	mMap, err := getNodeMetrics(pct)
	handleErr(err)

	mts := []plugin.PluginMetricType{}
	for _, n := range mMap {
		for ns, _ := range n {
			namespace := strings.Split(ns, "/")
			namespace[3] = "*"
			mts = append(mts, plugin.PluginMetricType{Namespace_: namespace})
		}
		break
	}

	metrics, err := getClusterMetrics(pct)
	handleErr(err)

	for _, m := range metrics {
		mts = append(mts, plugin.PluginMetricType{Namespace_: m.Namespace()})
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
	items, err := config.GetConfigItems(cfg, []string{esHost, esPort})
	if err != nil {
		log.Fatal(err.Error())
	}

	server := items[esHost].(string)
	port := items[esPort].(int)

	return server + ":" + strconv.Itoa(port)
}
