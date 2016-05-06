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
	"os"
	"reflect"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/core"
)

const (
	clusterEndPoint = "/_cluster/health"
	clusterNsPrefix = "intel/elasticsearch/cluster"
	HOST            = "host"
	FROM            = "from"
)

var (
	namespaces = []plugin.MetricType{}
)

type cluster struct {
	ClusterName                 string  `json:"cluster_name"`
	Status                      string  `json:"status"`
	TimeOut                     bool    `json:"timed_out"`
	NumberOfNodes               uint    `json:"number_of_nodes"`
	NumberOfDataNodes           uint    `json:"number_of_data_nodes"`
	ActivePrimaryShards         uint    `json:"active_primary_shards"`
	ActiveShards                uint    `json:"active_shards"`
	RelocatingShards            uint    `json:"relocating_shards"`
	InitializingShards          uint    `json:"initializing_shards"`
	UnassignedShards            uint    `json:"unassigned_shards"`
	DelayedUnassignedShards     uint    `json:"delayed_unassigned_shards"`
	NumberOfPendingTasks        uint    `json:"number_of_pending_tasks"`
	NumberOfInFlightFetch       uint    `json:"number_of_in_flight_fetch"`
	TaskMaxWaitingInQueueMillis int64   `json:"task_max_waiting_in_queue_millis"`
	ActiveShardsPercentAsNumber float32 `json:"active_shards_percent_as_number"`
}

// NewESClusterMetric returns a new instance of ESMetric
func NewESClusterMetric(url string, timeout time.Duration) ESMetric {
	return ESMetric{
		client: NewClient(url, clusterEndPoint, timeout),
	}
}

// GetClusterData collects the ES cluster metrics. Otherwise,
// an error is returned.
func (esm *ESMetric) GetClusterData() (map[string]plugin.MetricType, error) {
	mts := map[string]plugin.MetricType{}
	host, _ := os.Hostname()

	resp, err := esm.client.httpClient.Get(esm.client.GetUrl())
	if err != nil {
		esLog.WithFields(log.Fields{
			"_block": "get-cluster-data-http",
			"error":  err,
		}).Error("get ES cluster data error")
		return nil, err
	}
	defer resp.Body.Close()

	var v *cluster
	err = json.NewDecoder(resp.Body).Decode(&v)
	if err != nil {
		esLog.WithFields(log.Fields{
			"_block": "get-cluster-data-decoding",
			"error":  err,
		}).Error("decoding ES cluster data error")
		return nil, err
	}

	elem := reflect.ValueOf(v).Elem()
	for i := 0; i < elem.NumField(); i++ {
		ns := strings.Split(clusterNsPrefix, "/")
		ns = append(ns, elem.Type().Field(i).Tag.Get("json"))

		mts[strings.Join(ns, "/")] = plugin.MetricType{
			Namespace_: core.NewNamespace(ns...),
			Data_:      elem.Field(i).Interface(),
			Timestamp_: time.Now(),
			Tags_:      map[string]string{FROM: host},
		}
	}
	return mts, nil
}
