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
	"github.com/x-debug/beetlecrawl"
	"log"
)

type SinaSpider struct {
	beetlecrawl.BaseSpider
}

func (ss *SinaSpider) Name() string {
	return "sina_spider"
}

func (ss *SinaSpider) Init() error {
	log.Printf("init %s\n", ss.Name())
	return ss.YieldHttp(beetlecrawl.NewHttpRequest(beetlecrawl.GET, "https://sports.sina.com.cn/", ss.ParseList))
}

func (ss *SinaSpider) ParseList(resp *beetlecrawl.HttpResponse) (err error) {
	log.Printf("parse list %s, http status code %d\n", ss.Name(), resp.StatusCode)
	hrefs := resp.Css().QueryAll("div.more-layer ul li a")
	for _, href := range hrefs {
		log.Printf("href: %s\n", href.FindAttr("href"))
	}
	return err
}
