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

package main

import (
	"github.com/x-debug/beetlecrawl"
	"github.com/x-debug/beetlecrawl/examples/spiders"
	"log"
)

func main() {
	appConf, err := beetlecrawl.LoadConfig("./examples/examples.yaml")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Downloader.MaxRetry: %d\n", appConf.Downloader.MaxRetry)

	scheduler := beetlecrawl.NewYieldScheduler(appConf)
	scheduler.AddSpider(&spiders.SinaSpider{})
	scheduler.AddSpider(&spiders.BlogSpider{})
	scheduler.AddSpider(&spiders.ErrorSpider{})
	log.Fatal(scheduler.Run())
}
