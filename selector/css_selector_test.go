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

var cssHtmlContent = `
	<div class="more-layer">
		<i class="ico ico-arrow"></i>
		<ul class="clearfix">
			<li><a href="//auto.sina.com.cn/">汽车</a></li>
			<li class="edu"><a href="//edu.sina.com.cn/">教育</a></li>
			<li><a href="//fashion.sina.com.cn/">时尚</a></li>
		</ul>
	</div>
`

func TestCssSelector_Query(t *testing.T) {
	sel := NewCssSelector(strings.NewReader(cssHtmlContent))
	nodes := sel.Query("div.more-layer ul li a")
	for _, node := range nodes {
		if node.FindAttr("href") != "//auto.sina.com.cn/" &&
			node.FindAttr("href") != "//edu.sina.com.cn/" &&
			node.FindAttr("href") != "//fashion.sina.com.cn/" {
			t.Error("href error")
		}
	}
}

func TestCssSelector_QuerySingle(t *testing.T) {
	sel := NewCssSelector(strings.NewReader(cssHtmlContent))
	node := sel.Query("div.more-layer ul li a").First()
	if node == nil {
		t.Error("first link is empty")
	}
}

func TestCssSelector_QueryClass(t *testing.T) {
	sel := NewCssSelector(strings.NewReader(cssHtmlContent))
	eduNode := sel.Query("div.more-layer ul li.edu").First()
	if eduNode == nil {
		t.Error("edu link is empty")
	}

	eduHref := eduNode.QuerySingle("a").FindAttr("href")
	if eduHref != "//edu.sina.com.cn/" {
		t.Error("edu link href error")
	}
}

func TestCssSelector_QueryText(t *testing.T) {
	sel := NewCssSelector(strings.NewReader(cssHtmlContent))
	eduNode := sel.Query("div.more-layer ul li.edu").First()
	if eduNode == nil {
		t.Error("edu link is empty")
	}

	eduText := eduNode.QuerySingle("a").FirstChild.Data
	if eduText != "教育" {
		t.Errorf("edu link text error %s", eduText)
	}
}
