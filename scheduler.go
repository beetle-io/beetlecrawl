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
	"log"
)

const (
	CountOfDownloader   = 10
	CountOfUserRequests = 1000
)

type (
	Scheduler interface {
		EmitHttpReq(req *HttpRequest) error
		AddSpider(spider Spider) error
		Spiders() []Spider
		Run() error
	}

	//YieldScheduler is the schedule center of main loop, it will schedule the
	//request to downloader on the local
	YieldScheduler struct {
		spiders []Spider

		httpDownloaders []*httpDownloader
		userHttpReqs    chan *HttpRequest
		//TODO atomic
		countOfUserHttpReqs int
		downloadHttpReqs    chan *HttpResponse
	}
)

func NewYieldScheduler() *YieldScheduler {
	downloaders := make([]*httpDownloader, 0)
	for i := 0; i < CountOfDownloader; i++ {
		//TODO export to user
		downloader := newHttpDownloader()
		if err := downloader.Init(); err != nil {
			log.Printf("init downloader error, %v", err)
			continue
		}
		downloaders = append(downloaders, downloader)
	}

	return &YieldScheduler{
		httpDownloaders:  downloaders,
		spiders:          make([]Spider, 0),
		userHttpReqs:     make(chan *HttpRequest, CountOfUserRequests),
		downloadHttpReqs: make(chan *HttpResponse),
	}
}

func (ys *YieldScheduler) EmitHttpReq(req *HttpRequest) error {
	ys.userHttpReqs <- req
	ys.countOfUserHttpReqs = ys.countOfUserHttpReqs + 1
	return nil
}

func (ys *YieldScheduler) AddSpider(spider Spider) error {
	spider.SetScheduler(ys)
	ys.spiders = append(ys.spiders, spider)
	return nil
}

func (ys *YieldScheduler) Spiders() []Spider {
	return ys.spiders
}

func (ys *YieldScheduler) Run() error {
	for _, spider := range ys.spiders {
		if err := spider.Init(); err != nil {
			return err
		}
	}

	for {
		select {
		case req := <-ys.userHttpReqs:
			log.Printf("get user req %v", req)
			req.respCh = ys.downloadHttpReqs
			//TODO downloader select
			_ = ys.httpDownloaders[0].DownloadAsync(req)
		case resp := <-ys.downloadHttpReqs:
			go func() {
				//TODO defer the panic on the spider

				if err := resp.req.respCb(resp); err != nil {
					log.Printf("Spider callback error, %v", err)
				}

				defer resp.Body.Close()
			}()
		}
	}
}
