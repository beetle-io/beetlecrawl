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

package selector

import (
	"strings"
	"testing"
)

var xpathHtmlContent = `
<div class="news-list-a">
	<ul>
		<li class="item">
			<a href="https://sports.sina.com.cn/basketball/nba/2022-11-24/doc-imqqsmrp7346307.shtml" target="_blank">贝弗利本赛季投篮命中率仅26.6% 全联盟最低</a>
		</li>
		<li class="item">
			<a href="https://sports.sina.com.cn/basketball/nba/2022-11-24/doc-imqmmthc5769865.shtml" target="_blank">科比大女儿被陌生男子骚扰 法院已签署限制令</a>
		</li>
		<li class="item">
			<a class="hot-news" href="https://sports.sina.com.cn/basketball/nba/2022-11-21/doc-imqmmthc5409983.shtml" target="_blank">威少连续4场助攻上双 大胜马刺手指却意外受伤</a>
		</li>
	</ul>
</div>
`

func TestXPathSelector_Query(t *testing.T) {
	liNodes := NewXPathSelector(strings.NewReader(xpathHtmlContent)).Query("//li")
	for _, node := range liNodes {
		if "_blank" != node.QuerySingle("//a/@target").FirstChild.Data {
			t.Error("xpath selector query failed")
		}
	}
}

func TestXPathSelector_QuerySingle(t *testing.T) {
	firstNode := NewXPathSelector(strings.NewReader(xpathHtmlContent)).QuerySingle("//div[@class='news-list-a']/ul/li[1]/a/@href")
	if firstNode.FindAttr("href") == "" {
		t.Error("xpath selector query single failed")
	}
}

func TestXPathSelector_QueryClass(t *testing.T) {
	xpathSelector := NewXPathSelector(strings.NewReader(xpathHtmlContent))
	aNodes := xpathSelector.Query("//li[@class='item']/a")
	for _, node := range aNodes {
		if node.FirstChild.Data == "" {
			t.Error("xpath selector query class failed")
		}
	}
}

func TestXPathSelector_QueryText(t *testing.T) {
	xpathSelector := NewXPathSelector(strings.NewReader(xpathHtmlContent))
	secondNode := xpathSelector.QuerySingle("//div[@class='news-list-a']").QuerySingle("//li[2]").QuerySingle("//a")
	if secondNode.FirstChild.Data != "科比大女儿被陌生男子骚扰 法院已签署限制令" {
		t.Error("xpath selector query text failed")
	}
}
