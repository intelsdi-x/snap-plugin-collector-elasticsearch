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
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/core"
	log "github.com/sirupsen/logrus"
)

var (
	mtsMap     = map[string]map[string]plugin.MetricType{}
	treeMap    = map[string]int{}
	fieldCount int

	esLog        = log.WithField("_module", "es-node-client")
	nodeNsPrefix = "intel/elasticsearch/node/%s"
	nodeEndPoint = "/_nodes/stats"
)

type esData struct {
	ClusterName string           `json:"cluster_name"`
	Nodes       map[string]*node `json:"nodes"`
}

// ESMetric defines the metric needed data
type ESMetric struct {
	client    *HTTPClient
	host      string
	id        string
	timestamp int64
	dataMap   map[string]interface{}
}

// NewESNodeMetric returns a new instance of ESMetric
func NewESNodeMetric(url string, timeout time.Duration) ESMetric {
	return ESMetric{
		client: NewClient(url, nodeEndPoint, timeout),
	}
}

// GetNodeData returns an array of elasticsearch node metrics.
// All metrics of all nodes within the same cluster will be returned.
func (esm *ESMetric) GetNodeData() (map[string]map[string]plugin.MetricType, error) {
	resp, err := esm.client.httpClient.Get(esm.client.GetUrl())
	if err != nil {
		esLog.WithFields(log.Fields{
			"_block": "get-node-data-http",
			"error":  err,
		}).Error("get ES data error")
		return nil, err
	}
	defer resp.Body.Close()

	var v esData
	err = json.NewDecoder(resp.Body).Decode(&v)
	if err != nil {
		esLog.WithFields(log.Fields{
			"_block": "get-node-data-decoding",
			"error":  err,
		}).Error("decoding response error")
		return nil, err
	}

	nodes := v.Nodes
	nsStack := NewStack(0)

	// loops through different nodes in the same elasticsearch cluster
	for id, node := range nodes {
		// builds namespace based on the node id
		prefix := fmt.Sprintf(nodeNsPrefix, id)

		// pushes the node namespace onto the stack
		nsStack.Clear()
		nsStack.Push(prefix)

		// stores a map entry of the node namespace along with the number of fields
		treeMap[nodeNsPrefix] = reflect.TypeOf(node).Elem().NumField()

		esm.id = id
		esm.host = node.Host
		esm.timestamp = node.Timestamp
		esm.dataMap = map[string]interface{}{}

		// parses the data from the node
		err = esm.parseData(reflect.ValueOf(node), nsStack)
		esm.setESNodeMetrics()
	}
	return mtsMap, err
}

// parseData parses elasticsearch metric data into snap complaint metric type
func (esm *ESMetric) parseData(obj reflect.Value, nsStack *stack) error {
	switch obj.Kind() {
	case
		reflect.Int, reflect.Int32, reflect.Int64,
		reflect.Bool, reflect.String, reflect.Float32,
		reflect.Float64, reflect.Uint:

		str := strings.Join(nsStack.All(), "/")
		esm.dataMap[str] = obj.Interface()

		cleanStack(nsStack)
	case reflect.Ptr:
		if !obj.IsNil() {
			typ := obj.Elem()

			for i := 0; i < typ.NumField(); i++ {
				subty := reflect.Indirect(obj).Type().Field(i).Type.Kind()
				tag := reflect.Indirect(obj).Type().Field(i).Tag.Get("json")
				ty := reflect.Indirect(obj).Type().Field(i).Type
				if ty.Kind() == reflect.Ptr {
					ty = ty.Elem()
				}

				if subty == reflect.Ptr {
					fieldCount = ty.NumField()
				} else {
					fieldCount = 1
				}
				nsStack.Push(tag)
				treeMap[tag] = fieldCount

				err := esm.parseData(typ.Field(i), nsStack)
				if err != nil {
					esLog.WithFields(log.Fields{
						"_block": "parse-data-recursively",
						"error":  err,
					}).Error("parse data error")
					return err
				}
			}
		}
	case reflect.Slice:
		str := strings.Join(nsStack.All(), "/")
		slices := obj.Interface().([]interface{})
		for _, v := range slices {
			mv := v.(map[string]interface{})
			for k, s := range mv {
				str = strings.Join(nsStack.All(), "/") + "/" + k
				esm.dataMap[str] = s
			}
		}
		cleanStack(nsStack)
	default:
		esLog.WithFields(log.Fields{
			"_block": "parse-data-default",
			"error":  obj.Kind(),
		}).Error("unsupported data type")
	}
	return nil
}

func (esm *ESMetric) setESNodeMetrics() {
	mts := map[string]plugin.MetricType{}
	for n, m := range esm.dataMap {
		dpt := plugin.MetricType{
			Namespace_: core.NewNamespace(strings.Split(n, "/")...),
			Data_:      m,
			Tags_:      map[string]string{HOST: esm.host},
			Timestamp_: time.Now(),
		}
		mts[n] = dpt
	}
	mtsMap[esm.id] = mts
}

func cleanStack(nsStack *stack) error {
	leaf, err := nsStack.Pop()
	if err != nil {
		return err
	}

	// minus count for each pop
	if treeMap[leaf] > 0 {
		treeMap[leaf]--

		node, err := nsStack.Peek()
		if err != nil {
			return err
		}
		// minus parent node count
		treeMap[node]--

		// pop out a node when the count is 0
		// and continue checking its parent in
		// the while loop like fashion
		for treeMap[node] == 0 {
			nsStack.Pop()
			node, err = nsStack.Peek()
			treeMap[node]--
		}
	}
	return nil
}
