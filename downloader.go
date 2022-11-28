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
	"sync"
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
		conf         *DownloaderConfig
		failReqs     []*HttpRequest
		failLock     sync.Mutex
	}
)

func newHttpDownloader(conf *DownloaderConfig) *httpDownloader {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	httpDownloader := &httpDownloader{
		downloadReqs: make(chan *HttpRequest, defaultRequestQueueSize),
		closeCh:      make(chan struct{}),
		httpClient:   &http.Client{Transport: tr},
		conf:         conf,
		failReqs:     make([]*HttpRequest, 0),
	}
	return httpDownloader
}

func (hd *httpDownloader) Init() error {
	go hd.run()
	return nil
}

func (hd *httpDownloader) run() {
	for req := range hd.downloadReqs {
		go func(req *HttpRequest) {
			resp, err := hd.httpClient.Do(req.Request)
			if err != nil {
				req.errs = append(req.errs, err)
				if req.retryTimes < hd.conf.MaxRetry {
					req.retryTimes += 1
					_ = hd.DownloadAsync(req)
				} else {
					hd.failLock.Lock()
					hd.failReqs = append(hd.failReqs, req)
					hd.failLock.Unlock()
					req.failReqCh <- req
				}
				return
			}
			req.successRespCh <- newHttpResponse(resp, req)
		}(req)
	}
}

func (hd *httpDownloader) DownloadAsync(req Request) error {
	hd.downloadReqs <- req.(*HttpRequest)
	return nil
}
