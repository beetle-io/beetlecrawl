// Copyright 2022 beetlecrawl Project Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package beetlecrawl

import (
	"net/http"
	"time"
)

const defaultRequestQueueSize = 1000

type (
	Downloader interface {
		Init() error
		DownloadAsync(req Request) error
	}

	httpDownloader struct {
		downloadReqs chan *HttpRequest
		closeCh      chan struct{}
		httpClient   *http.Client
	}
)

func newHttpDownloader() *httpDownloader {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	httpDownloader := &httpDownloader{
		downloadReqs: make(chan *HttpRequest, defaultRequestQueueSize),
		closeCh:      make(chan struct{}),
		httpClient:   &http.Client{Transport: tr},
	}
	return httpDownloader
}

func (hd *httpDownloader) Init() error {
	go hd.run()
	return nil
}

func (hd *httpDownloader) run() {
	for req := range hd.downloadReqs {
		resp, err := hd.httpClient.Do(req.Request)
		if err != nil {
			continue
		}

		req.respCh <- newHttpResponse(resp, req)
	}
}

func (hd *httpDownloader) DownloadAsync(req Request) error {
	hd.downloadReqs <- req.(*HttpRequest)
	return nil
}
