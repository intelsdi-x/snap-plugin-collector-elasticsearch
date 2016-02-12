<!--
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
-->

# snap collector plugin - Elasticsearch

This plugin collects Elasticsearch cluster and nodes statistics using snap telemetry engine.

The intention for this plugin is to collect metrics for Elasticsearch nodes and cluster health.

This plugin is used in the [snap framework] (http://github.com/intelsdi-x/snap).


1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Operating systems](#openrating-systems)
  * [Installation](#installation)
  * [Configuration and Usage](#configuration-and-usage)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license)
6. [Acknowledgements](#acknowledgements)

## Getting Started

In order to use this plugin you need the Elasticsearch cluster that you can collect metrics from.

### System Requirements

* [snap](http://github.com/intelsdi-x/snap)
* Elasticsearch node/cluster
* [golang 1.5+](https://golang.org/dl/)

### Operating systems
All OSs currently supported by snap:
* Linux/amd64
* Darwin/amd64

### Installation

#### Install Elasticsearch:
To install sysstat package from the official repositories simply use:
- For Darwin/amd64: `brew install elasticsearch`
- For Docker installation: 

##### Get and run the latest elasticsearch
```
$ docker pull elasticsearch
Using default tag: latest
latest: Pulling from library/elasticsearch
77e39ee82117: Pull complete 
...
Digest: sha256:c8ca4802e3ee5d6165da4c2158eee14d4dc7f4f39e1edbc62a003d659be35f3c
Status: Downloaded newer image for elasticsearch:latest

$ docker run -d --name snap-elasticsearch -p 9200:9200 -p 9300:9300 elasticsearch
2878aeba854c7f65c52f86b8ccf6f5ebf221cb63ff021d519ffd4572065fd183
```
##### Check Elasticsearch cluster health
```
$ http://DOCKERHOST:9200/_cluster/health?pretty
{
  "cluster_name" : "elasticsearch",
  "status" : "green",
  "timed_out" : false,
  "number_of_nodes" : 1,
  "number_of_data_nodes" : 1,
  "active_primary_shards" : 0,
  "active_shards" : 0,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 0,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 100.0
}
```
#### To build the plugin binary:
Get the source by running a `go get` to fetch the code:
```
$ go get github.com/intelsdi-x/snap-plugin-collector-elasticsearch
```

Build the plugin by running make within the cloned repo:
```
$ cd $GOPATH/src/github.com/intelsdi-x/snap-plugin-collector-elasticsearch && make
```
This builds the plugin in `/build/rootfs/`

#### Builds
You can also download prebuilt binaries for OS X and Linux (64-bit) at the [releases](https://github.com/intelsdi-x/snap-plugin-collector-elasticsearch/releases) page

### Configuration and Usage
* Set up the [snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)
* Ensure `$SNAP_PATH` is exported  
`export SNAP_PATH=$GOPATH/src/github.com/intelsdi-x/snap/build`
* Ensure `SNAP_ES_HOST` is exported

## Documentation

To learn more about this plugin:

* [snap elasticsearch examples](#examples)

### Collected Metrics
This plugin has the ability to gather the following metrics:

**Elastic Node Statistics**

This collector supports all node metrics for Elasticsearch 2.1.1. 

Metric namespace prefix: /intel/elasticsearch/node/{id}

Namespace |
------------ |
/intel/elastcisearch/node/{id}/host|
/intel/elastcisearch/node/{id}/timestamp|
/intel/elastcisearch/node/{id}/name|
/intel/elastcisearch/node/{id}/indices/docs/count|
/intel/elastcisearch/node/{id}/indices/docs/deleted|
/intel/elastcisearch/node/{id}/indices/store/size_in_bytes|
/intel/elastcisearch/node/{id}/indices/store/throttle_time_in_millis|
/intel/elastcisearch/node/{id}/indices/indexing/index_total|
/intel/elastcisearch/node/{id}/indices/indexing/index_time_in_millis|
/intel/elastcisearch/node/{id}/indices/indexing/index_current|
/intel/elastcisearch/node/{id}/indices/indexing/index_failed|
/intel/elastcisearch/node/{id}/indices/indexing/delete_total|
/intel/elastcisearch/node/{id}/indices/indexing/delete_time_in_millis|
/intel/elastcisearch/node/{id}/indices/indexing/delete_current|
/intel/elastcisearch/node/{id}/indices/indexing/noop_update_total|
/intel/elastcisearch/node/{id}/indices/indexing/is_throttled|
/intel/elastcisearch/node/{id}/indices/indexing/throttle_time_in_millis|
/intel/elastcisearch/node/{id}/indices/get/total|
/intel/elastcisearch/node/{id}/indices/get/timeInMillis|
/intel/elastcisearch/node/{id}/indices/get/exists_total|
/intel/elastcisearch/node/{id}/indices/get/exists_time_in_millis|
/intel/elastcisearch/node/{id}/indices/get/missing_total|
/intel/elastcisearch/node/{id}/indices/get/missing_time_in_millis|
/intel/elastcisearch/node/{id}/indices/get/current|
/intel/elastcisearch/node/{id}/indices/search/open_contexts|
/intel/elastcisearch/node/{id}/indices/search/query_total|
/intel/elastcisearch/node/{id}/indices/search/query_time_in_millis|
/intel/elastcisearch/node/{id}/indices/search/query_current|
/intel/elastcisearch/node/{id}/indices/search/fetch_total|
/intel/elastcisearch/node/{id}/indices/search/fetch_time_in_millis|
/intel/elastcisearch/node/{id}/indices/search/fetch_current|
/intel/elastcisearch/node/{id}/indices/search/scroll_total|
/intel/elastcisearch/node/{id}/indices/search/scroll_time_in_millis|
/intel/elastcisearch/node/{id}/indices/search/scroll_current|
/intel/elastcisearch/node/{id}/indices/merges/current|
/intel/elastcisearch/node/{id}/indices/merges/current_docs|
/intel/elastcisearch/node/{id}/indices/merges/current_size_in_bytes|
/intel/elastcisearch/node/{id}/indices/merges/total|
/intel/elastcisearch/node/{id}/indices/merges/total_time_in_millis|
/intel/elastcisearch/node/{id}/indices/merges/total_docs|
/intel/elastcisearch/node/{id}/indices/merges/total_size_in_bytes|
/intel/elastcisearch/node/{id}/indices/merges/total_stopped_time_in_millis|
/intel/elastcisearch/node/{id}/indices/merges/total_throttled_time_in_millis|
/intel/elastcisearch/node/{id}/indices/merges/total_auto_throttle_in_bytes|
/intel/elastcisearch/node/{id}/indices/refresh/total|
/intel/elastcisearch/node/{id}/indices/refresh/total_time_in_millis|
/intel/elastcisearch/node/{id}/indices/flush/total|
/intel/elastcisearch/node/{id}/indices/flush/total_time_in_millis|
/intel/elastcisearch/node/{id}/indices/warmer/current|
/intel/elastcisearch/node/{id}/indices/warmer/total|
/intel/elastcisearch/node/{id}/indices/warmer/total_time_in_millis|
/intel/elastcisearch/node/{id}/indices/query_cache/memory_size_in_bytes|
/intel/elastcisearch/node/{id}/indices/query_cache/total_count|
/intel/elastcisearch/node/{id}/indices/query_cache/hit_count|
/intel/elastcisearch/node/{id}/indices/query_cache/miss_count|
/intel/elastcisearch/node/{id}/indices/query_cache/cache_size|
/intel/elastcisearch/node/{id}/indices/query_cache/cache_count|
/intel/elastcisearch/node/{id}/indices/query_cache/evictions|
/intel/elastcisearch/node/{id}/indices/fielddata/memory_size_in_bytes|
/intel/elastcisearch/node/{id}/indices/fielddata/evictions|
/intel/elastcisearch/node/{id}/indices/percolate/total|
/intel/elastcisearch/node/{id}/indices/percolate/time_in_millis|
/intel/elastcisearch/node/{id}/indices/percolate/current|
/intel/elastcisearch/node/{id}/indices/percolate/memory_size_in_bytes|
/intel/elastcisearch/node/{id}/indices/percolate/memory_size|
/intel/elastcisearch/node/{id}/indices/percolate/queries|
/intel/elastcisearch/node/{id}/indices/completion/size_in_bytes|
/intel/elastcisearch/node/{id}/indices/segments/count|
/intel/elastcisearch/node/{id}/indices/segments/memory_in_bytes|
/intel/elastcisearch/node/{id}/indices/segments/terms_memory_in_bytes|
/intel/elastcisearch/node/{id}/indices/segments/stored_fields_memory_in_bytes|
/intel/elastcisearch/node/{id}/indices/segments/term_vectors_memory_in_bytes|
/intel/elastcisearch/node/{id}/indices/segments/norms_memory_in_bytes|
/intel/elastcisearch/node/{id}/indices/segments/doc_values_memory_in_bytes|
/intel/elastcisearch/node/{id}/indices/segments/index_writer_memory_in_bytes|
/intel/elastcisearch/node/{id}/indices/segments/index_writer_max_memory_in_bytes|
/intel/elastcisearch/node/{id}/indices/segments/version_map_memory_in_bytes|
/intel/elastcisearch/node/{id}/indices/segments/fixed_bit_set_memory_in_bytes|
/intel/elastcisearch/node/{id}/indices/translog/operations|
/intel/elastcisearch/node/{id}/indices/translog/size_in_bytes|
/intel/elastcisearch/node/{id}/indices/suggest/total|
/intel/elastcisearch/node/{id}/indices/suggest/time_in_millis|
/intel/elastcisearch/node/{id}/indices/suggest/current|
/intel/elastcisearch/node/{id}/indices/request_cache/memory_size_in_bytes|
/intel/elastcisearch/node/{id}/indices/request_cache/evictions|
/intel/elastcisearch/node/{id}/indices/request_cache|
/intel/elastcisearch/node/{id}/indices/request_cache/miss_count|
/intel/elastcisearch/node/{id}/indices/recovery/current_as_source|
/intel/elastcisearch/node/{id}/indices/recovery/current_as_target|
/intel/elastcisearch/node/{id}/indices/recovery/throttle_time_in_millis|
/intel/elastcisearch/node/{id}/os/timestamp|
/intel/elastcisearch/node/{id}/os/load_average|
/intel/elastcisearch/node/{id}/os/mem/total_in_bytes|
/intel/elastcisearch/node/{id}/os/mem/free_in_bytes|
/intel/elastcisearch/node/{id}/os/mem/used_in_bytes|
/intel/elastcisearch/node/{id}/os/mem/free_percent|
/intel/elastcisearch/node/{id}/os/mem/used_percent|
/intel/elastcisearch/node/{id}/os/swap/total_in_bytes|
/intel/elastcisearch/node/{id}/os/swap/free_in_bytes|
/intel/elastcisearch/node/{id}/os/swap/used_in_bytes|
/intel/elastcisearch/node/{id}/process/timestamp|
/intel/elastcisearch/node/{id}/process/open_file_descriptors|
/intel/elastcisearch/node/{id}/process/max_file_descriptors|
/intel/elastcisearch/node/{id}/process/cpu/percent|
/intel/elastcisearch/node/{id}/process/cpu/total_in_millis|
/intel/elastcisearch/node/{id}/process/mem/total_virtual_in_bytes|
/intel/elastcisearch/node/{id}/jvm/timestamp|
/intel/elastcisearch/node/{id}/jvm/uptime_in_millis|
/intel/elastcisearch/node/{id}/jvm/mem/heap_used_in_bytes|
/intel/elastcisearch/node/{id}/jvm/mem/heap_used_percent|
/intel/elastcisearch/node/{id}/jvm/mem/heap_committed_in_bytes|
/intel/elastcisearch/node/{id}/jvm/mem/heap_max_in_bytes|
/intel/elastcisearch/node/{id}/jvm/mem/non_heap_used_in_bytes|
/intel/elastcisearch/node/{id}/jvm/mem/non_heap_committed_in_bytes|
/intel/elastcisearch/node/{id}/jvm/mem/pools/young/used_in_bytes|
/intel/elastcisearch/node/{id}/jvm/mem/pools/young/max_in_bytes|
/intel/elastcisearch/node/{id}/jvm/mem/pools/young/peak_used_in_bytes|
/intel/elastcisearch/node/{id}/jvm/mem/pools/young/peak_max_in_bytes|
/intel/elastcisearch/node/{id}/jvm/mem/pools/Survivor/used_in_bytes|
/intel/elastcisearch/node/{id}/jvm/mem/pools/Survivor/max_in_bytes|
/intel/elastcisearch/node/{id}/jvm/mem/pools/Survivor/peak_used_in_bytes|
/intel/elastcisearch/node/{id}/jvm/mem/pools/Survivor/peak_max_in_bytes|
/intel/elastcisearch/node/{id}/jvm/mem/pools/old/used_in_bytes|
/intel/elastcisearch/node/{id}/jvm/mem/pools/old/max_in_bytes|
/intel/elastcisearch/node/{id}/jvm/mem/pools/old/peak_used_in_bytes|
/intel/elastcisearch/node/{id}/jvm/mem/pools/old/peak_max_in_bytes|
/intel/elastcisearch/node/{id}/jvm/threads/count|
/intel/elastcisearch/node/{id}/jvm/threads/peak_count|
/intel/elastcisearch/node/{id}/jvm/gc/collectors/young/collection_count|
/intel/elastcisearch/node/{id}/jvm/gc/collectors/young/collection_time_in_millis|
/intel/elastcisearch/node/{id}/jvm/gc/collectors/old/collection_count|
/intel/elastcisearch/node/{id}/jvm/gc/collectors/old/collection_time_in_millis|
/intel/elastcisearch/node/{id}/jvm/buffer_pools/direct/count|
/intel/elastcisearch/node/{id}/jvm/buffer_pools/direct/used_in_bytes|
/intel/elastcisearch/node/{id}/jvm/buffer_pools/direct/total_capacity_in_bytes|
/intel/elastcisearch/node/{id}/jvm/buffer_pools/mapped/count|
/intel/elastcisearch/node/{id}/jvm/buffer_pools/mapped/used_in_bytes|
/intel/elastcisearch/node/{id}/jvm/buffer_pools/mapped/total_capacity_in_bytes|
/intel/elastcisearch/node/{id}/jvm/classes/current_loaded_count|
/intel/elastcisearch/node/{id}/jvm/classes/total_loaded_count|
/intel/elastcisearch/node/{id}/jvm/classes/total_unloaded_count|
/intel/elastcisearch/node/{id}/thread_pool/bulk/threads|
/intel/elastcisearch/node/{id}/thread_pool/bulk/queue|
/intel/elastcisearch/node/{id}/thread_pool/bulk/active|
/intel/elastcisearch/node/{id}/thread_pool/bulk/rejected|
/intel/elastcisearch/node/{id}/thread_pool/bulk/largest|
/intel/elastcisearch/node/{id}/thread_pool/bulk/completed|
/intel/elastcisearch/node/{id}/thread_pool/fetch_shard_started/threads|
/intel/elastcisearch/node/{id}/thread_pool/fetch_shard_started/queue|
/intel/elastcisearch/node/{id}/thread_pool/fetch_shard_started/active|
/intel/elastcisearch/node/{id}/thread_pool/fetch_shard_started/rejected|
/intel/elastcisearch/node/{id}/thread_pool/fetch_shard_started/largest|
/intel/elastcisearch/node/{id}/thread_pool/fetch_shard_started/completed|
/intel/elastcisearch/node/{id}/thread_pool/fetch_shard_store/threads|
/intel/elastcisearch/node/{id}/thread_pool/fetch_shard_store/queue|
/intel/elastcisearch/node/{id}/thread_pool/fetch_shard_store/active|
/intel/elastcisearch/node/{id}/thread_pool/fetch_shard_store/rejected|
/intel/elastcisearch/node/{id}/thread_pool/fetch_shard_store/largest|
/intel/elastcisearch/node/{id}/thread_pool/fetch_shard_store/completed|
/intel/elastcisearch/node/{id}/thread_pool/flush/threads|
/intel/elastcisearch/node/{id}/thread_pool/flush/queue|
/intel/elastcisearch/node/{id}/thread_pool/flush/active|
/intel/elastcisearch/node/{id}/thread_pool/flush/rejected|
/intel/elastcisearch/node/{id}/thread_pool/flush/largest|
/intel/elastcisearch/node/{id}/thread_pool/flush/completed|
/intel/elastcisearch/node/{id}/thread_pool/force_merge/threads|
/intel/elastcisearch/node/{id}/thread_pool/force_merge/queue|
/intel/elastcisearch/node/{id}/thread_pool/force_merge/active|
/intel/elastcisearch/node/{id}/thread_pool/force_merge/rejected|
/intel/elastcisearch/node/{id}/thread_pool/force_merge/largest|
/intel/elastcisearch/node/{id}/thread_pool/force_merge/completed|
/intel/elastcisearch/node/{id}/thread_pool/generic/threads|
/intel/elastcisearch/node/{id}/thread_pool/generic/queue|
/intel/elastcisearch/node/{id}/thread_pool/generic/active|
/intel/elastcisearch/node/{id}/thread_pool/generic/rejected|
/intel/elastcisearch/node/{id}/thread_pool/generic/largest|
/intel/elastcisearch/node/{id}/thread_pool/generic/completed|
/intel/elastcisearch/node/{id}/thread_pool/get/threads|
/intel/elastcisearch/node/{id}/thread_pool/get/queue|
/intel/elastcisearch/node/{id}/thread_pool/get/active|
/intel/elastcisearch/node/{id}/thread_pool/get/rejected|
/intel/elastcisearch/node/{id}/thread_pool/get/largest|
/intel/elastcisearch/node/{id}/thread_pool/get/completed|
/intel/elastcisearch/node/{id}/thread_pool/index/threads|
/intel/elastcisearch/node/{id}/thread_pool/index/queue|
/intel/elastcisearch/node/{id}/thread_pool/index/active|
/intel/elastcisearch/node/{id}/thread_pool/index/rejected|
/intel/elastcisearch/node/{id}/thread_pool/index/largest|
/intel/elastcisearch/node/{id}/thread_pool/index/completed|
/intel/elastcisearch/node/{id}/thread_pool/listener/threads|
/intel/elastcisearch/node/{id}/thread_pool/listener/queue|
/intel/elastcisearch/node/{id}/thread_pool/listener/active|
/intel/elastcisearch/node/{id}/thread_pool/listener/rejected|
/intel/elastcisearch/node/{id}/thread_pool/listener/largest|
/intel/elastcisearch/node/{id}/thread_pool/listener/completed|
/intel/elastcisearch/node/{id}/thread_pool/management/threads|
/intel/elastcisearch/node/{id}/thread_pool/management/queue|
/intel/elastcisearch/node/{id}/thread_pool/management/active|
/intel/elastcisearch/node/{id}/thread_pool/management/rejected|
/intel/elastcisearch/node/{id}/thread_pool/management/largest|
/intel/elastcisearch/node/{id}/thread_pool/management/completed|
/intel/elastcisearch/node/{id}/thread_pool/percolate/threads|
/intel/elastcisearch/node/{id}/thread_pool/percolate/queue|
/intel/elastcisearch/node/{id}/thread_pool/percolate/active|
/intel/elastcisearch/node/{id}/thread_pool/percolate/rejected|
/intel/elastcisearch/node/{id}/thread_pool/percolate/largest|
/intel/elastcisearch/node/{id}/thread_pool/percolate/completed|
/intel/elastcisearch/node/{id}/thread_pool/refresh/threads|
/intel/elastcisearch/node/{id}/thread_pool/refresh/queue|
/intel/elastcisearch/node/{id}/thread_pool/refresh/active|
/intel/elastcisearch/node/{id}/thread_pool/refresh/rejected|
/intel/elastcisearch/node/{id}/thread_pool/refresh/largest|
/intel/elastcisearch/node/{id}/thread_pool/refresh/completed|
/intel/elastcisearch/node/{id}/thread_pool/search/threads|
/intel/elastcisearch/node/{id}/thread_pool/search/queue|
/intel/elastcisearch/node/{id}/thread_pool/search/active|
/intel/elastcisearch/node/{id}/thread_pool/search/rejected|
/intel/elastcisearch/node/{id}/thread_pool/search/largest|
/intel/elastcisearch/node/{id}/thread_pool/search/completed|
/intel/elastcisearch/node/{id}/thread_pool/snapshot/threads|
/intel/elastcisearch/node/{id}/thread_pool/snapshot/queue|
/intel/elastcisearch/node/{id}/thread_pool/snapshot/active|
/intel/elastcisearch/node/{id}/thread_pool/snapshot/rejected|
/intel/elastcisearch/node/{id}/thread_pool/snapshot/largest|
/intel/elastcisearch/node/{id}/thread_pool/snapshot/completed|
/intel/elastcisearch/node/{id}/thread_pool/suggest/threads|
/intel/elastcisearch/node/{id}/thread_pool/suggest/queue|
/intel/elastcisearch/node/{id}/thread_pool/suggest/active|
/intel/elastcisearch/node/{id}/thread_pool/suggest/rejected|
/intel/elastcisearch/node/{id}/thread_pool/suggest/largest|
/intel/elastcisearch/node/{id}/thread_pool/suggest/completed|
/intel/elastcisearch/node/{id}/thread_pool/warmer/threads|
/intel/elastcisearch/node/{id}/thread_pool/warmer/queue|
/intel/elastcisearch/node/{id}/thread_pool/warmer/active|
/intel/elastcisearch/node/{id}/thread_pool/warmer/rejected|
/intel/elastcisearch/node/{id}/thread_pool/warmer/largest|
/intel/elastcisearch/node/{id}/thread_pool/warmer/completed|
/intel/elastcisearch/node/{id}/fs/timestamp|
/intel/elastcisearch/node/{id}/fs/total/total_in_bytes|
/intel/elastcisearch/node/{id}/fs/total/free_in_bytes|
/intel/elastcisearch/node/{id}/fs/total/available_in_bytes|
/intel/elastcisearch/node/{id}/fs/data/transport/server_open|
/intel/elastcisearch/node/{id}/fs/data/transport/rx_count|
/intel/elastcisearch/node/{id}/fs/data/transport/rx_size_in_bytes|
/intel/elastcisearch/node/{id}/fs/data/transport/tx_count|
/intel/elastcisearch/node/{id}/fs/data/transport/tx_size_in_bytes|
/intel/elastcisearch/node/{id}/http/current_open|
/intel/elastcisearch/node/{id}/http/total_opened|
/intel/elastcisearch/node/{id}/breakers/request/limit_size_in_bytes|
/intel/elastcisearch/node/{id}/breakers/request/limit_size|
/intel/elastcisearch/node/{id}/breakers/request/estimated_size_in_bytes|
/intel/elastcisearch/node/{id}/breakers/request/estimated_size|
/intel/elastcisearch/node/{id}/breakers/request/overhead|
/intel/elastcisearch/node/{id}/breakers/request/tripped|
/intel/elastcisearch/node/{id}/breakers/fielddata/limit_size_in_bytes|
/intel/elastcisearch/node/{id}/breakers/fielddata/limit_size|
/intel/elastcisearch/node/{id}/breakers/fielddata/estimated_size_in_bytes|
/intel/elastcisearch/node/{id}/breakers/fielddata/estimated_size|
/intel/elastcisearch/node/{id}/breakers/fielddata/overhead|
/intel/elastcisearch/node/{id}/breakers/fielddata/tripped|
/intel/elastcisearch/node/{id}/breakers/parent/limit_size_in_bytes|
/intel/elastcisearch/node/{id}/breakers/parent/limit_size|
/intel/elastcisearch/node/{id}/breakers/parent/estimated_size_in_bytes|
/intel/elastcisearch/node/{id}/breakers/parent/estimated_size|
/intel/elastcisearch/node/{id}/breakers/parent/overhead|
/intel/elastcisearch/node/{id}/breakers/parent/tripped|
/intel/elastcisearch/node/{id}/script/compilations|
/intel/elastcisearch/node/{id}/script/cache_evictions|

**Elasticsearch Cluster Statistics**

This collector supports all cluster metrics for Elasticsearch 2.1.1. 

Metric namespace prefix: /intel/elasticsearch/cluster

Namespace |
------------ |
/intel/elasticsearch/cluster/cluster_name|
/intel/elasticsearch/cluster/status|
/intel/elasticsearch/cluster/timed_out|
/intel/elasticsearch/cluster/number_of_nodes|
/intel/elasticsearch/cluster/number_of_data_nodes|
/intel/elasticsearch/cluster/active_primary_shards|
/intel/elasticsearch/cluster/active_shards|
/intel/elasticsearch/cluster/relocating_shards|
/intel/elasticsearch/cluster/initializing_shards|
/intel/elasticsearch/cluster/unassigned_shards|
/intel/elasticsearch/cluster/delayed_unassigned_shards|
/intel/elasticsearch/cluster/number_of_pending_tasks|
/intel/elasticsearch/cluster/number_of_in_flight_fetch|
/intel/elasticsearch/cluster/task_max_waiting_in_queue_millis|
/intel/elasticsearch/cluster/active_shards_percent_as_number|

### Examples
Example running snap-plugin-collector-elasticsearch, passthru processor, and writing data to a file.

![Dockerized example](https://media2.giphy.com/avatars/snapsnap/ubJywMcap0zU.gif)

In one terminal window, open the snap daemon (in this case with logging set to 1 and trust disabled):
```
$ $SNAP_PATH/bin/snapd -l 1 -t 0
```
In another terminal window:
Load snap-plugin-collector-elasticsearch
```
$ $SNAP_PATH/bin/snapctl plugin load $SNAP_PATH/plugin/snap-plugin-collector-elasticsearch
Plugin loaded
Name: elasticsearch
Version: 1
Type: collector
Signed: false
Loaded Time: Sat, 13 Feb 2016 17:05:47 PST
```
See available metrics for your system (this is just part of the list)
```
$SNAP_PATH/bin/snapctl metric list                                
NAMESPACE 												 VERSIONS
/intel/elasticsearch/cluster/active_primary_shards 							 1
/intel/elasticsearch/cluster/active_shards 								 1
/intel/elasticsearch/cluster/active_shards_percent_as_number 						 1
/intel/elasticsearch/cluster/cluster_name 								 1
/intel/elasticsearch/cluster/delayed_unassigned_shards 							 1
/intel/elasticsearch/cluster/initializing_shards 							 1
/intel/elasticsearch/cluster/number_of_data_nodes 							 1
/intel/elasticsearch/cluster/number_of_in_flight_fetch 							 1
/intel/elasticsearch/cluster/number_of_nodes 								 1
/intel/elasticsearch/cluster/number_of_pending_tasks 							 1
/intel/elasticsearch/cluster/relocating_shards 								 1
/intel/elasticsearch/cluster/status 									 1
/intel/elasticsearch/cluster/task_max_waiting_in_queue_millis 						 1
/intel/elasticsearch/cluster/timed_out 									 1
/intel/elasticsearch/cluster/unassigned_shards 								 1
/intel/elasticsearch/node/wmya7Qp9S7OWtKugsX55IQ/breakers/fielddata/estimated_size 			 1
/intel/elasticsearch/node/wmya7Qp9S7OWtKugsX55IQ/breakers/fielddata/estimated_size_in_bytes 		 1
/intel/elasticsearch/node/wmya7Qp9S7OWtKugsX55IQ/breakers/fielddata/limit_size 				 1
/intel/elasticsearch/node/wmya7Qp9S7OWtKugsX55IQ/breakers/fielddata/limit_size_in_bytes 		 1
/intel/elasticsearch/node/wmya7Qp9S7OWtKugsX55IQ/breakers/fielddata/overhead 				 1
/intel/elasticsearch/node/wmya7Qp9S7OWtKugsX55IQ/breakers/fielddata/tripped 				 1
/intel/elasticsearch/node/wmya7Qp9S7OWtKugsX55IQ/breakers/parent/estimated_size 			 1
/intel/elasticsearch/node/wmya7Qp9S7OWtKugsX55IQ/breakers/parent/estimated_size_in_bytes 		 1
```

Load passthru plugin for processing:
```
$SNAP_PATH/bin/snapctl plugin load $SNAP_PATH/plugin/snap-processor-passthru
Plugin loaded
Name: passthru
Version: 1
Type: processor
Signed: false
Loaded Time: Sat, 13 Feb 2016 17:06:03 PST
```

Load file plugin for publishing:
```
$SNAP_PATH/bin/snapctl plugin load $SNAP_PATH/plugin/snap-publisher-file  
Plugin loaded
Name: file
Version: 3
Type: publisher
Signed: false
Loaded Time: Sat, 13 Feb 2016 17:06:17 PST
```

Create a task manifest file (e.g. `elasticsearch-file.json`. replace node id):    
```json
{
    "version": 1,
    "schedule": {
        "type": "simple",
        "interval": "1s"
    },
    "workflow": {
        "collect": {
            "metrics": {
                "/intel/elasticsearch/cluster/unassigned_shards": {},
                "/intel/elasticsearch/cluster/active_shards": {},
                "/intel/elasticsearch/node/wmya7Qp9S7OWtKugsX55IQ/indices/docs/count": {},
                "/intel/elasticsearch/node/wmya7Qp9S7OWtKugsX55IQ/indices/merges/current_size_in_bytes": {},
                "/intel/elasticsearch/node/wmya7Qp9S7OWtKugsX55IQ/indices/search/open_contexts": {},
                "/intel/elasticsearch/node/wmya7Qp9S7OWtKugsX55IQ/jvm/mem/heap_used_in_bytes": {},
                "/intel/elasticsearch/node/wmya7Qp9S7OWtKugsX55IQ/os/load_average": {},
                "/intel/elasticsearch/node/wmya7Qp9S7OWtKugsX55IQ/thread_pool/fetch_shard_started/completed": {},
                "/intel/elasticsearch/node/wmya7Qp9S7OWtKugsX55IQ/thread_pool/force_merge/threads": {},
                "/intel/elasticsearch/node/wmya7Qp9S7OWtKugsX55IQ/transport/tx_size_in_bytes": {}
            },
            "config": {
                "/intel/mock": {
                    "password": "secret",
                    "user": "root"
                }
            },
            "process": [
                {
                    "plugin_name": "passthru",
                    "process": null,
                    "publish": [
                        {                         
                            "plugin_name": "file",
                            "config": {
                                "file": "/tmp/published_elasticsearch"
                            }
                        }
                    ],
                    "config": null
                }
            ],
            "publish": null
        }
    }
}
```

Create task:
```
$SNAP_PATH/bin/snapctl task create -t ../../task/elaticsearch-task.json
Using task manifest to create task
Task created
ID: 5aadafc8-a7a1-427c-892c-87e680235563
Name: Task-5aadafc8-a7a1-427c-892c-87e680235563
State: Running
```

See file output (this is just part of the file):
```
48090-05-02 00:29:45 -0800 PST|[intel elasticsearch node wmya7Qp9S7OWtKugsX55IQ indices fielddata memory_size_in_bytes]|0|127.0.0.1
48090-05-02 00:29:45 -0800 PST|[intel elasticsearch node wmya7Qp9S7OWtKugsX55IQ indices segments norms_memory_in_bytes]|61632|127.0.0.1
48090-05-02 00:29:45 -0800 PST|[intel elasticsearch node wmya7Qp9S7OWtKugsX55IQ indices segments index_writer_max_memory_in_bytes]|17920000|127.0.0.1
48090-05-02 00:29:45 -0800 PST|[intel elasticsearch node wmya7Qp9S7OWtKugsX55IQ indices request_cache miss_count]|0|127.0.0.1
48090-05-02 00:29:45 -0800 PST|[intel elasticsearch node wmya7Qp9S7OWtKugsX55IQ thread_pool bulk largest]|0|127.0.0.1
48090-05-02 00:29:45 -0800 PST|[intel elasticsearch node wmya7Qp9S7OWtKugsX55IQ thread_pool refresh completed]|0|127.0.0.1
48090-05-02 00:29:45 -0800 PST|[intel elasticsearch node wmya7Qp9S7OWtKugsX55IQ thread_pool warmer queue]|0|127.0.0.1
48090-05-02 00:29:45 -0800 PST|[intel elasticsearch node wmya7Qp9S7OWtKugsX55IQ indices indexing index_total]|0|127.0.0.1
48090-05-02 00:29:45 -0800 PST|[intel elasticsearch node wmya7Qp9S7OWtKugsX55IQ os mem total_in_bytes]|17179869184|127.0.0.1
48090-05-02 00:29:45 -0800 PST|[intel elasticsearch node wmya7Qp9S7OWtKugsX55IQ jvm timestamp]|1455417187986|127.0.0.1
48090-05-02 00:29:45 -0800 PST|[intel elasticsearch node wmya7Qp9S7OWtKugsX55IQ thread_pool flush largest]|0|127.0.0.1
48090-05-02 00:29:45 -0800 PST|[intel elasticsearch node wmya7Qp9S7OWtKugsX55IQ indices indexing delete_total]|0|127.0.0.1
48090-05-02 00:29:45 -0800 PST|[intel elasticsearch node wmya7Qp9S7OWtKugsX55IQ jvm gc collectors young collection_time_in_millis]|4465|127.0.0.1
48090-05-02 00:29:45 -0800 PST|[intel elasticsearch node wmya7Qp9S7OWtKugsX55IQ thread_pool bulk active]|0|127.0.0.1
48090-05-02 00:29:45 -0800 PST|[intel elasticsearch node wmya7Qp9S7OWtKugsX55IQ indices flush total]|0|127.0.0.1
48090-05-02 00:29:45 -0800 PST|[intel elasticsearch node wmya7Qp9S7OWtKugsX55IQ transport tx_size_in_bytes]|2592|127.0.0.1
48090-05-02 00:29:45 -0800 PST|[intel elasticsearch node wmya7Qp9S7OWtKugsX55IQ breakers fielddata estimated_size_in_bytes]|0|127.0.0.1
48090-05-02 00:29:45 -0800 PST|[intel elasticsearch node wmya7Qp9S7OWtKugsX55IQ indices docs deleted]|0|127.0.0.1
48090-05-02 00:29:45 -0800 PST|[intel elasticsearch node wmya7Qp9S7OWtKugsX55IQ indices merges total_docs]|0|127.0.0.1
48090-05-02 00:29:45 -0800 PST|[intel elasticsearch node wmya7Qp9S7OWtKugsX55IQ os mem free_in_bytes]|25284608|127.0.0.1
48090-05-02 00:29:45 -0800 PST|[intel elasticsearch node wmya7Qp9S7OWtKugsX55IQ jvm classes total_unloaded_count]|0|127.0.0.1
48090-05-02 00:29:45 -0800 PST|[intel elasticsearch node wmya7Qp9S7OWtKugsX55IQ thread_pool warmer threads]|1|127.0.0.1
48090-05-02 00:29:45 -0800 PST|[intel elasticsearch node wmya7Qp9S7OWtKugsX55IQ breakers request estimated_size]|0b|127.0.0.1
48090-05-02 00:29:45 -0800 PST|[intel elasticsearch node wmya7Qp9S7OWtKugsX55IQ indices warmer total_time_in_millis]|22|127.0.0.1
48090-05-02 00:29:45 -0800 PST|[intel elasticsearch node wmya7Qp9S7OWtKugsX55IQ indices request_cache memory_size_in_bytes]|0|127.0.0.1
48090-05-02 00:29:45 -0800 PST|[intel elasticsearch node wmya7Qp9S7OWtKugsX55IQ jvm mem pools old used_in_bytes]|24052744|127.0.0.1
48090-05-02 00:29:45 -0800 PST|[intel elasticsearch node wmya7Qp9S7OWtKugsX55IQ thread_pool listener completed]|987|127.0.0.1
```

### Roadmap
This plugin is still in active development. As we launch this plugin, we have a few items in mind for the next few releases:
- [ ] Additional error handling

If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-elasticsearch/issues) 
and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-collector-elasticsearch/pulls).

## Community Support
This repository is one of **many** plugins in the **snap**, a powerful telemetry agent framework. See the full project at 
http://github.com/intelsdi-x/snap. To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support).


## Contributing
We love contributions!

There is more than one way to give back, from examples to blogs to code updates.

## License

[snap](http://github.com/intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).


## Acknowledgements

* Author: [@candysmurf](https://github.com/candysmurf/)
