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
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

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

	// Timeout duration
	timeout = 5 * time.Second

	noHostErr = "SNAP_ES_HOST enviromental variable must be set"
	esHost    = "SNAP_ES_HOST"
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
	metrics := make([]plugin.PluginMetricType, len(mts))
	esMts, err := getMetrics()
	if err != nil {
		return nil, err
	}
	metrics = append(metrics, esMts...)

	return metrics, nil
}

// GetMetricTypes returns the metric types exposed by Elasticsearch
func (p *Elasticsearch) GetMetricTypes(_ plugin.PluginConfigType) ([]plugin.PluginMetricType, error) {
	// gather all metrics to populate namespaces
	_, err := getMetrics()
	if err != nil {
		return nil, err
	}

	// aggregate all metric namespaces for all nodes
	nodeMetricType := getESNodeMetricTypes()

	// gather cluster metric namespaces
	esClusterMetricType := getESClusterMetricType()
	nodeMetricType = append(nodeMetricType, esClusterMetricType...)

	return nodeMetricType, nil
}

// GetConfigPolicy returns a ConfigPolicy
func (p *Elasticsearch) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	c := cpolicy.New()
	return c, nil
}

func getMetrics() ([]plugin.PluginMetricType, error) {
	host := getServer()

	esNodeMetrics := NewESNodeMetric(host, timeout)
	mts, err := esNodeMetrics.GetNodeData()
	if err != nil {
		return nil, err
	}

	esClusterMetrics := NewESClusterMetric(host, timeout)
	metrics, err := esClusterMetrics.GetClusterData()
	if err != nil {
		return nil, err
	}

	mts = append(mts, metrics...)

	return mts, nil
}

func joinNamespace(ns []string) string {
	return "/" + strings.Join(ns, "/")
}

func prettyPrint(mts []plugin.PluginMetricType) error {
	var out bytes.Buffer
	mtsb, _, _ := plugin.MarshalPluginMetricTypes(plugin.SnapJSONContentType, mts)
	if err := json.Indent(&out, mtsb, "", "  "); err != nil {
		return err
	}
	fmt.Println(out.String())
	return nil
}

func handleErr(e error) {
	if e != nil {
		panic(e)
	}
}

func getServer() string {
	host := os.Getenv(esHost)
	if host == "" {
		panic(noHostErr)
	}
	return host + ":9200"
}
