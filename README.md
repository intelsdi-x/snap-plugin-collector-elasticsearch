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

This plugin collects Elasticsearch cluster and nodes statistics using snap telemetry framework.

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
* [golang 1.6+](https://golang.org/dl/)

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
* Ensure [snap global configuration](./examples/cfg/snap-config-sample.json) is defined 

## Documentation

To learn more about this plugin:

* [snap elasticsearch examples](#examples)

### Collected Metrics
This plugin can gather Elasticsearch node and cluster level statistics. 
The node level statistics are similar for each node except that each node has a different node id. To show the statistics for all nodes inside a cluster,  using a wildcard * to represent 
node id in the task manifest. Otherwise, a particular node id may be specified.

* [Node Metrics](METRICS_NODE.md)
* [Cluster Metrics](METRICS_CLUSTER.md)

In the node level, this plugin collects metrics listed the next catalog. 

| Metric Name| Description |
| :------------ | :------------- |
|indices| Indices stats about size, document count, indexing and deletion times, search times, field cache size, merges and flushes|
|os| Operating system stats, load average, mem, swap|
|process| Process statistics, memory consumption, cpu usage, open file descriptors|
|jvm| JVM stats, memory pool information, garbage collection, buffer pools, number of loaded/unloaded classes|
|thread_pool| Statistics about each thread pool, including current size, queue and rejected tasks|
|fs| File system information, data path, free disk space, read/write stats|
|http| HTTP connection information|
|breaks| Statistics about the field data circuit breaker|
|script| Computing the grades stats based on a script|

### Examples
Example running snap-plugin-collector-elasticsearch, passthru processor, and writing data to a file.

![Dockerized example](http://media.giphy.com/media/3osxY87TeMy7jGrbDW/giphy.gif)

In one terminal window, open the snap daemon (in this case with logging set to 1 and trust disabled):
```
$ $SNAP_PATH/bin/snapd -l 1 -t 0 --config <path to snap global config>
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
AMESPACE 									 VERSIONS
/intel/elasticsearch/cluster/active_primary_shards 				 1
/intel/elasticsearch/cluster/active_shards 					 1
/intel/elasticsearch/cluster/active_shards_percent_as_number 			 1
/intel/elasticsearch/cluster/cluster_name 					 1
/intel/elasticsearch/cluster/delayed_unassigned_shards 				 1
/intel/elasticsearch/cluster/number_of_data_nodes 							 1
/intel/elasticsearch/cluster/number_of_in_flight_fetch 							 1
/intel/elasticsearch/cluster/number_of_nodes 								 1
/intel/elasticsearch/cluster/number_of_pending_tasks 							 1
/intel/elasticsearch/cluster/relocating_shards 								 1
/intel/elasticsearch/cluster/status 									 1
/intel/elasticsearch/cluster/task_max_waiting_in_queue_millis 						 1
/intel/elasticsearch/cluster/timed_out 									 1
/intel/elasticsearch/cluster/unassigned_shards 								 1
/intel/elasticsearch/node/*/thread_pool/warmer/largest 				 1
/intel/elasticsearch/node/*/thread_pool/warmer/queue 				 1
/intel/elasticsearch/node/*/thread_pool/warmer/rejected 			 1
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

Create a task manifest file (e.g. `elasticsearch-task.json`). 
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
                "/intel/elasticsearch/node/*/indices/docs/count": {},
                "/intel/elasticsearch/node/*/indices/merges/current_size_in_bytes": {},
                "/intel/elasticsearch/node/*/indices/search/open_contexts": {},
                "/intel/elasticsearch/node/*/jvm/mem/heap_used_in_bytes": {},
                "/intel/elasticsearch/node/*/os/load_average": {},
                "/intel/elasticsearch/node/*/thread_pool/fetch_shard_started/completed": {},
                "/intel/elasticsearch/node/*/thread_pool/force_merge/threads": {},
                "/intel/elasticsearch/node/*/transport/tx_size_in_bytes": {}
            },
            "config": {},
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
48178-05-24 14:55:10 -0800 PST|[intel elasticsearch node F2fP3bedSsqK8S_v440enA os mem total_in_bytes]|1044631552|172.17.0.6
48178-05-24 14:55:10 -0800 PST|[intel elasticsearch node F2fP3bedSsqK8S_v440enA timestamp]|1458196124110|172.17.0.6
2016-03-16 23:28:53.281902987 -0700 PDT|[intel elasticsearch cluster status]|green|egu-mac01.lan
48178-05-24 15:11:51 -0800 PST|[intel elasticsearch node F2fP3bedSsqK8S_v440enA host]|172.17.0.6|172.17.0.6
48178-05-24 15:11:51 -0800 PST|[intel elasticsearch node F2fP3bedSsqK8S_v440enA indices merges total_throttled_time_in_millis]|0|172.17.0.6
48178-05-24 15:11:51 -0800 PST|[intel elasticsearch node F2fP3bedSsqK8S_v440enA jvm classes total_loaded_count]|6636|172.17.0.6
48178-05-24 15:11:51 -0800 PST|[intel elasticsearch node F2fP3bedSsqK8S_v440enA os mem total_in_bytes]|1044631552|172.17.0.6
48178-05-24 15:11:51 -0800 PST|[intel elasticsearch node F2fP3bedSsqK8S_v440enA timestamp]|1458196125111|172.17.0.6
2016-03-16 23:28:54.282481719 -0700 PDT|[intel elasticsearch cluster status]|green|egu-mac01.lan
48178-05-24 15:28:31 -0800 PST|[intel elasticsearch node F2fP3bedSsqK8S_v440enA host]|172.17.0.6|172.17.0.6
48178-05-24 15:28:31 -0800 PST|[intel elasticsearch node F2fP3bedSsqK8S_v440enA indices merges total_throttled_time_in_millis]|0|172.17.0.6
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
* Author: [@geauxvirtual](https://github.com/geauxvirtual)
* Author: [@jcooklin](https://github.com/jcooklin)
