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

package spiders

import (
	"errors"
	"github.com/x-debug/beetlecrawl"
	"log"
)

var ErrNotHandle = errors.New("http not handle")

type ErrorSpider struct {
	beetlecrawl.BaseSpider
}

func (es *ErrorSpider) Name() string {
	return "error_spider"
}

func (es *ErrorSpider) Init() error {
	es.SetupError(es.OnError)
	//The url is not exist, the spider MUST handle the error
	es.YieldHttp(beetlecrawl.NewHttpRequest(beetlecrawl.GET, "https://chenxf1.org/archives/detail.html", es.ParseDetail))
	return nil
}

func (es *ErrorSpider) ParseDetail(resp *beetlecrawl.HttpResponse) error {
	//The response can't arrive here because of the http error
	return nil
}

func (es *ErrorSpider) OnError(errReq *beetlecrawl.HttpRequest, errs []error) error {
	log.Printf("Url of error: %s, retry times of error: %d\n", errReq.URL(), errReq.RetryTimes())
	if len(errs) == 0 {
		return ErrNotHandle
	}
	return nil
}
