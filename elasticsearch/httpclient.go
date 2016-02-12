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
	"net/http"
	"net/url"
	"time"
)

// HTTPClient defines the client for HTTP communication
type HTTPClient struct {
	url        string
	httpClient *http.Client
	endPoint   string
}

// NewClient returns a new instance of HTTPClient
func NewClient(url, endpoint string, timeout time.Duration) *HTTPClient {
	return &HTTPClient{
		url:        url,
		httpClient: &http.Client{Timeout: timeout},
		endPoint:   endpoint,
	}
}

// GetUrl returns the URL of a HTTPClient
func (hc *HTTPClient) GetUrl() string {
	u := url.URL{
		Scheme: "http",
		Host:   hc.url,
		Path:   hc.endPoint,
	}
	return u.String()
}
