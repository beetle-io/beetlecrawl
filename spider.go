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

type Spider interface {
	Init() error
	Name() string
	YieldHttp(req Request) error
	SetScheduler(scheduler Scheduler)
}

type BaseSpider struct {
	sched Scheduler
}

func (bs *BaseSpider) Init() error {
	return nil
}

func (bs *BaseSpider) Name() string {
	return "base_spider"
}

func (bs *BaseSpider) YieldHttp(req Request) error {
	return bs.sched.EmitHttpReq(req.(*HttpRequest))
}

func (bs *BaseSpider) SetScheduler(scheduler Scheduler) {
	bs.sched = scheduler
}
