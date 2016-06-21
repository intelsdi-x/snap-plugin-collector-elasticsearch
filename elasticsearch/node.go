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

// Package elasticsearch ports Elasticsearch metrics into snap metrics
package elasticsearch

type node struct {
	Host       string      `json:"host"`
	Timestamp  int64       `json:"timestamp"`
	Name       string      `json:"name"`
	Indices    *indices    `json:"indices"`
	OS         *nodeOs     `json:"os"`
	Process    *process    `json:"process"`
	JVM        *jvm        `json:"jvm"`
	ThreadPool *threadPool `json:"thread_pool"`
	FS         *fs         `json:"fs"`
	Transport  *transport  `json:"transport"`
	HTTP       *nodehttp   `json:"http"`
	Breakers   *breakers   `json:"breakers"`
	Script     *script     `json:"script"`
}

type docs struct {
	Count   int64 `json:"count"`
	Deleted int64 `json:"deleted"`
}

type store struct {
	SizeInBytes          int64 `json:"size_in_bytes"`
	ThrottleTimeInMillis int64 `json:"throttle_time_in_millis"`
}

type indexing struct {
	IndexTotal           int64 `json:"index_total"`
	IndexTimeInMillis    int64 `json:"index_time_in_millis"`
	IndexCurrent         int64 `json:"index_current"`
	IndexFailed          int64 `json:"index_failed"`
	DeleteTotal          int64 `json:"delete_total"`
	DeleteTimeInMillis   int64 `json:"delete_time_in_millis"`
	DeleteCurrent        int64 `json:"delete_current"`
	NoopUpdateTotal      int32 `json:"noop_update_total"`
	IsThrottled          bool  `json:"is_throttled"`
	ThrottleTimeInMillis int64 `json:"throttle_time_in_millis"`
}

type get struct {
	Total               int64 `json:"total"`
	TimeInMillis        int64 `json:"timeInMillis"`
	ExistsTotal         int64 `json:"exists_total"`
	ExistsTimeInMillis  int64 `json:"exists_time_in_millis"`
	MissingTotal        int64 `json:"missing_total"`
	MissingTimeInMillis int64 `json:"missing_time_in_millis"`
	Current             int64 `json:"current"`
}

type search struct {
	OpenContext        int64 `json:"open_contexts"`
	QueryTotal         int64 `json:"query_total"`
	QueryTimeInMillis  int64 `json:"query_time_in_millis"`
	QueryCurrent       int64 `json:"query_current"`
	FetchTotal         int64 `json:"fetch_total"`
	FetchTimeInMillis  int64 `json:"fetch_time_in_millis"`
	FetchCurrent       int64 `json:"fetch_current"`
	ScrollTotal        int64 `json:"scroll_total"`
	ScrollTimeInMillis int64 `json:"scroll_time_in_millis"`
	ScrollCurrent      int64 `json:"scroll_current"`
}

type merges struct {
	Current                    int64 `json:"current"`
	CurrentDocs                int64 `json:"current_docs"`
	CurrentSizeInBytes         int64 `json:"current_size_in_bytes"`
	Total                      int64 `json:"total"`
	TotalTimeInMillis          int64 `json:"total_time_in_millis"`
	TotalDocs                  int64 `json:"total_docs"`
	TotalSizeInBytes           int64 `json:"total_size_in_bytes"`
	TotalStoppedTimeInMillis   int64 `json:"total_stopped_time_in_millis"`
	TotalThrottledTimeInMillis int64 `json:"total_throttled_time_in_millis"`
	TotalAutoThrottleInBytes   int64 `json:"total_auto_throttle_in_bytes"`
}

type refresh struct {
	Total             int64 `json:"total"`
	TotalTimeInMillis int64 `json:"total_time_in_millis"`
}

type flush struct {
	Total             int64 `json:"total"`
	TotalTimeInMillis int64 `json:"total_time_in_millis"`
}

type warmer struct {
	Current           int64 `json:"current"`
	Total             int64 `json:"total"`
	TotalTimeInMillis int64 `json:"total_time_in_millis"`
}

type queryCache struct {
	MemorySizeInBytes int64 `json:"memory_size_in_bytes"`
	TotalCount        int64 `json:"total_count"`
	HitCount          int64 `json:"hit_count"`
	MissCount         int64 `json:"miss_count"`
	CacheSize         int64 `json:"cache_size"`
	CacheCount        int64 `json:"cache_count"`
	Eviction          int64 `json:"evictions"`
}

type fielddata struct {
	MemorySizeInBytes int64 `json:"memory_size_in_bytes"`
	Evictions         int64 `json:"evictions"`
}

type percolate struct {
	Total             int64  `json:"total"`
	TimeInMillis      int64  `json:"time_in_millis"`
	Current           int64  `json:"current"`
	MemorySizeInBytes int64  `json:"memory_size_in_bytes"`
	MemorySize        string `json:"memory_size"`
	Queries           int64  `json:"queries"`
}

type completion struct {
	SizeInBytes int64 `json:"size_in_bytes"`
}
type segments struct {
	Count                       int64 `json:"count"`
	MemoryInBytes               int64 `json:"memory_in_bytes"`
	TermsMemoryInBytes          int64 `json:"terms_memory_in_bytes"`
	StoredFieldsMemoryInBytes   int64 `json:"stored_fields_memory_in_bytes"`
	TermVectorsMemoryInBytes    int64 `json:"term_vectors_memory_in_bytes"`
	NormsMemoryInBytes          int64 `json:"norms_memory_in_bytes"`
	DocValuesMemoryInBytes      int64 `json:"doc_values_memory_in_bytes"`
	IndexWriterMemoryInBytes    int64 `json:"index_writer_memory_in_bytes"`
	IndexWriterMaxMemoryInBytes int64 `json:"index_writer_max_memory_in_bytes"`
	VersionMapMemoryInBytes     int64 `json:"version_map_memory_in_bytes"`
	FixedBitSetMemoryInBytes    int64 `json:"fixed_bit_set_memory_in_bytes"`
}

type translog struct {
	Operations  int64 `json:"operations"`
	SizeInBytes int64 `json:"size_in_bytes"`
}

type suggest struct {
	Total        int64 `json:"total"`
	TimeInMillis int64 `json:"time_in_millis"`
	Current      int64 `json:"current"`
}

type requestCache struct {
	MemorySizeInBytes int64 `json:"memory_size_in_bytes"`
	Eviction          int64 `json:"evictions"`
	HitCount          int64 `josn:"hit_count"`
	MissCount         int64 `json:"miss_count"`
}

type recovery struct {
	CurrentAsSource      int64 `json:"current_as_source"`
	CurrentAsTarget      int64 `json:"current_as_target"`
	ThrottleTimeInMillis int64 `json:"throttle_time_in_millis"`
}

type indices struct {
	Docs         *docs         `json:"docs"`
	Store        *store        `json:"store"`
	Indexing     *indexing     `json:"indexing"`
	Get          *get          `json:"get"`
	Search       *search       `json:"search"`
	Merges       *merges       `json:"merges"`
	Refresh      *refresh      `json:"refresh"`
	Flush        *flush        `json:"flush"`
	Warmer       *warmer       `json:"warmer"`
	QueryCache   *queryCache   `json:"query_cache"`
	FieldData    *fielddata    `json:"fielddata"`
	Percolate    *percolate    `json:"percolate"`
	Completion   *completion   `json:"completion"`
	Segments     *segments     `json:"segments"`
	Translog     *translog     `json:"translog"`
	Suggest      *suggest      `json:"suggest"`
	RequestCache *requestCache `json:"request_cache"`
	Recovery     *recovery     `json:"recovery"`
}

type nodeOs struct {
	Timestamp   int64   `json:"timestamp"`
	LoadAverage float64 `json:"load_average"`
	Mem         *osmem  `json:"mem"`
	Swap        *swap   `json:"swap"`
}

type osmem struct {
	TotalInBytes   int64 `json:"total_in_bytes"`
	FreeInBytes    int64 `json:"free_in_bytes"`
	UsedInBytes    int64 `json:"used_in_bytes"`
	FreePercentage int   `json:"free_percent"`
	UsedPercent    int   `json:"used_percent"`
}

type swap struct {
	TotalInBytes int64 `json:"total_in_bytes"`
	FreeInBytes  int64 `json:"free_in_bytes"`
	UsedInBytes  int64 `json:"used_in_bytes"`
}

type process struct {
	Timestamp    int64 `json:"timestamp"`
	OpenFileDesc int32 `json:"open_file_descriptors"`
	MaxFileDesc  int32 `json:"max_file_descriptors"`
	Cpu          *cpu  `json:"cpu"`
	Mem          *pmem `json:"mem"`
}

type cpu struct {
	Percent       int   `json:"percent"`
	TotalInMillis int64 `json:"total_in_millis"`
}

type pmem struct {
	TotalVirtualInBytes int64 `json:"total_virtual_in_bytes"`
}

type jvm struct {
	Timestamp      int64        `json:"timestamp"`
	UptimeInMillis int64        `json:"uptime_in_millis"`
	Mem            *jvmMem      `json:"mem"`
	Threads        *thread      `json:"threads"`
	Gc             *gc          `json:"gc"`
	BufferPools    *bufferPools `json:"buffer_pools"`
	Classes        *classes     `json:"classes"`
}

type jvmMem struct {
	HeapUsedInBytes         int64  `json:"heap_used_in_bytes"`
	HeapUsedPercent         int    `json:"heap_used_percent"`
	HeapCommittedInBytes    int64  `json:"heap_committed_in_bytes"`
	HeapMaxInBytes          int64  `json:"heap_max_in_bytes"`
	NonHeapUsedInBytes      int64  `json:"non_heap_used_in_bytes"`
	NonHeapCommittedInBytes int64  `json:"non_heap_committed_in_bytes"`
	Pools                   *pools `json:"pools"`
}

type pools struct {
	Young    *young    `json:"young"`
	Survivor *survivor `json:"Survivor"`
	Old      *old      `json:"old"`
}

type young struct {
	UsedInBytes     int64 `json:"used_in_bytes"`
	MaxInBytes      int64 `json:"max_in_bytes"`
	PeakUsedInBytes int64 `json:"peak_used_in_bytes"`
	PeakMaxInBytes  int64 `json:"peak_max_in_bytes"`
}

type survivor struct {
	UsedInBytes     int64 `json:"used_in_bytes"`
	MaxInBytes      int64 `json:"max_in_bytes"`
	PeakUsedIbBytes int64 `json:"peak_used_in_bytes"`
	PeakMaxInBytes  int64 `json:"peak_max_in_bytes"`
}

type old struct {
	UsedInBytes     int64 `json:"used_in_bytes"`
	MaxInBytes      int   `json:"max_in_bytes"`
	PeakUsedInBytes int64 `json:"peak_used_in_bytes"`
	PeakMaxInBytes  int64 `json:"peak_max_in_bytes"`
}

type thread struct {
	Count     int `json:"count"`
	PeakCount int `json:"peak_count"`
}

type gc struct {
	Collectors *collectors `json:"collectors"`
}

type collectors struct {
	Young *cyoung `json:"young"`
	Old   *cold   `json:"old"`
}

type cyoung struct {
	CollectorCount        int   `json:"collection_count"`
	CollectorTimeInMillis int64 `json:"collection_time_in_millis"`
}

type cold struct {
	CollectorCount        int `json:"collection_count"`
	CollectorTimeInMillis int `json:"collection_time_in_millis"`
}

type bufferPools struct {
	Direct *direct `json:"direct"`
	Mapped *mapped `json:"mapped"`
}

type direct struct {
	Count                int   `json:"count"`
	UsedInBytes          int64 `json:"used_in_bytes"`
	TotalCapacityInBytes int64 `json:"total_capacity_in_bytes"`
}

type mapped struct {
	Count                int   `json:"count"`
	UsedInBytes          int64 `json:"used_in_bytes"`
	TotalCapacityInBytes int64 `json:"total_capacity_in_bytes"`
}

type classes struct {
	CurrentLoadedCount int32 `json:"current_loaded_count"`
	TotalLoadedCount   int32 `json:"total_loaded_count"`
	TotalUnloadedCount int32 `json:"total_unloaded_count"`
}

type threadPool struct {
	Bulk              *threadType `json:"bulk"`
	FetchShardStarted *threadType `json:"fetch_shard_started"`
	FetchShardStore   *threadType `json:"fetch_shard_store"`
	Flush             *threadType `json:"flush"`
	ForceMerge        *threadType `json:"force_merge"`
	Generic           *threadType `json:"generic"`
	Get               *threadType `json:"get"`
	Index             *threadType `json:"index"`
	Listener          *threadType `json:"listener"`
	Management        *threadType `json:"management"`
	Percolate         *threadType `json:"percolate"`
	Refresh           *threadType `json:"refresh"`
	Search            *threadType `json:"search"`
	Snapshot          *threadType `json:"snapshot"`
	Suggest           *threadType `json:"suggest"`
	Warmer            *threadType `json:"warmer"`
}

type threadType struct {
	Threads   uint `json:"threads"`
	Queue     uint `json:"queue"`
	Active    uint `json:"active"`
	Rejected  uint `json:"rejected"`
	Largest   uint `json:"largest"`
	Completed uint `json:"completed"`
}

type fs struct {
	Timestamp int64         `json:"timestamp"`
	Total     *total        `json:"total"`
	Data      []interface{} `json:"data"`
}

type total struct {
	TotalInBytes     int64 `json:"total_in_bytes"`
	FreeInBytes      int64 `json:"free_in_bytes"`
	AvailableInBytes int64 `json:"available_in_bytes"`
}

type data struct {
	Path             string `json:"path"`
	Mount            string `json:"mount"`
	Type             string `json:"type"`
	TotalInBytes     int64  `json:"total_in_bytes"`
	FreeInBytes      int64  `json:"free_in_bytes"`
	AvailableInBytes int64  `json:"available_in_bytes"`
}

type transport struct {
	ServerOpen    int32 `json:"server_open"`
	RxCount       int64 `json:"rx_count"`
	RxSizeInBytes int64 `json:"rx_size_in_bytes"`
	TxCount       int64 `json:"tx_count"`
	TxSizeInBytes int64 `json:"tx_size_in_bytes"`
}

type nodehttp struct {
	CurrentOpen int64 `json:"current_open"`
	TotalOpened int64 `json:"total_opened"`
}

type breakers struct {
	Request   *request    `json:"request"`
	Fielddata *bfielddata `json:"fielddata"`
	Parent    *parent     `json:"parent"`
}

type request struct {
	LimitSizeInBytes     int64   `json:"limit_size_in_bytes"`
	LimitSize            string  `json:"limit_size"`
	EstimatedSizeInBytes int64   `json:"estimated_size_in_bytes"`
	EstimatedSize        string  `json:"estimated_size"`
	Overhead             float32 `json:"overhead"`
	Tripped              int32   `json:"tripped"`
}

type bfielddata struct {
	LimitedSizeInBytes   int64   `json:"limit_size_in_bytes"`
	LimitSize            string  `json:"limit_size"`
	EstimatedSizeInBytes int64   `json:"estimated_size_in_bytes"`
	EstimatedSize        string  `json:"estimated_size"`
	Overhead             float32 `json:"overhead"`
	Tripped              int32   `json:"tripped"`
}

type parent struct {
	LimitSizeInBytes     int64   `json:"limit_size_in_bytes"`
	LimitSize            string  `json:"limit_size"`
	EstimatedSizeInBytes int64   `json:"estimated_size_in_bytes"`
	EstimatedSize        string  `json:"estimated_size"`
	Overhead             float32 `json:"overhead"`
	Tripped              int32   `json:"tripped"`
}

type script struct {
	Compilations   int64 `json:"compilations"`
	CacheEvictions int64 `json:"cache_evictions"`
}
