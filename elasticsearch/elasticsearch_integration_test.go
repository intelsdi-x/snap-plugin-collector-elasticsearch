//
// +build medium

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
	"testing"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/core"
	"github.com/intelsdi-x/snap/core/cdata"
	"github.com/intelsdi-x/snap/core/ctypes"
	. "github.com/smartystreets/goconvey/convey"
)

func TestESCollectMetrics(t *testing.T) {
	cfg := setupCfg("192.168.99.100", 9200)

	Convey("Elasticsearch collector", t, func() {
		p := NewElasticsearchCollector()
		p.GetMetricTypes(cfg)

		Convey("collect metrics", func() {
			mts := []plugin.MetricType{
				plugin.MetricType{
					Namespace_: core.NewNamespace(
						"intel", "elasticsearch", "node", "JWSiB4YtQF64iKavAkg_fQ",
						"jvm", "buffer_pools", "direct", "total_capacity_in_bytes"),
					Config_: cfg.ConfigDataNode,
				},
				plugin.MetricType{
					Namespace_: core.NewNamespace(
						"intel", "elasticsearch", "node", "JWSiB4YtQF64iKavAkg_fQ",
						"process", "cpu", "total_in_millis"),
					Config_: cfg.ConfigDataNode,
				},
				plugin.MetricType{
					Namespace_: core.NewNamespace(
						"intel", "elasticsearch", "node", "JWSiB4YtQF64iKavAkg_fQ",
						"jvm", "mem", "heap_max_in_bytes"),
					Config_: cfg.ConfigDataNode,
				},
				plugin.MetricType{
					Namespace_: core.NewNamespace(
						"intel", "elasticsearch", "node", "*",
						"os", "mem", "free_percent"),
					Config_: cfg.ConfigDataNode,
				},
				plugin.MetricType{
					Namespace_: core.NewNamespace(
						"intel", "elasticsearch", "node", "*",
						"thread_pool", "management", "completed"),
					Config_: cfg.ConfigDataNode,
				},
			}
			metrics, err := p.CollectMetrics(mts)
			So(err, ShouldBeNil)
			So(metrics, ShouldNotBeEmpty)
		})

		Convey("collect cluster metrics", func() {
			mts := []plugin.MetricType{
				plugin.MetricType{
					Namespace_: core.NewNamespace(
						"intel", "elasticsearch", "cluster",
						"status"),
					Config_: cfg.ConfigDataNode,
				},
				plugin.MetricType{
					Namespace_: core.NewNamespace(
						"intel", "elasticsearch", "cluster",
						"number_of_nodes"),
					Config_: cfg.ConfigDataNode,
				},
			}
			metrics, err := p.CollectMetrics(mts)
			So(err, ShouldBeNil)
			So(metrics, ShouldNotBeEmpty)
			So(metrics[0].Data(), ShouldNotBeNil)
			So(metrics[1].Data(), ShouldBeGreaterThan, 0)
		})
	})
}

func setupCfg(server string, port int) plugin.ConfigType {
	node := cdata.NewNode()
	node.AddItem("server", ctypes.ConfigValueStr{Value: server})
	node.AddItem("port", ctypes.ConfigValueInt{Value: port})
	return plugin.ConfigType{ConfigDataNode: node}
}
