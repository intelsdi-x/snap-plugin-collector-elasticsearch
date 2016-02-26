//
// +build integration

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
	"os"
	"testing"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/core/ctypes"
	. "github.com/smartystreets/goconvey/convey"
)

func TestESCollectMetrics(t *testing.T) {
	config := make(map[string]ctypes.ConfigValue)

	Convey("Elasticsearch collector", t, func() {
		config["host"] = ctypes.ConfigValueStr{Value: os.Getenv("SNAP_ES_HOST")}
		config["port"] = ctypes.ConfigValueInt{Value: 9200}

		p := NewElasticsearchCollector()
		Convey("collect metrics", func() {
			mts := []plugin.PluginMetricType{
				plugin.PluginMetricType{
					Namespace_: []string{
						"intel", "elasticsearch", "node", "wmya7Qp9S7OWtKugsX55IQ",
						"thread_pool", "force_merge", "completed"},
				},
				plugin.PluginMetricType{
					Namespace_: []string{
						"intel", "elasticsearch", "node", "wmya7Qp9S7OWtKugsX55IQ",
						"indices", "docs", "count"},
				},
				plugin.PluginMetricType{
					Namespace_: []string{
						"intel", "elasticsearch", "node", "wmya7Qp9S7OWtKugsX55IQ",
						"jvm", "mem", "heap_max_in_bytes"},
				},
				plugin.PluginMetricType{
					Namespace_: []string{
						"intel", "elasticsearch", "node", "wmya7Qp9S7OWtKugsX55IQ",
						"os", "mem", "free_percent"},
				},
			}
			metrics, err := p.CollectMetrics(mts)
			So(err, ShouldBeNil)
			So(metrics, ShouldNotBeNil)
		})
	})
}
