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

var htmlContent = `
<div class="nav-b-block layout-relative clearfix" layer-type="layer-wrap">
	<h1 class="logo"><a href="//sports.sina.com.cn/" target="_blank">新浪体育</a></h1>
	<ul class="links clearfix">
		<li><a href="//www.sina.com.cn/" suda-uatrack="key=ty0526&value=blk_sports_nav_www">新浪首页</a></li>
		<li><a href="//news.sina.com.cn/" suda-uatrack="key=ty0526&value=blk_sports_nav_news">新闻</a></li>
		<li><a href="//sports.sina.com.cn/" suda-uatrack="key=ty0526&value=blk_sports_nav_sports">体育</a></li>
		<li><a href="//finance.sina.com.cn/" suda-uatrack="key=ty0526&value=blk_sports_nav_finance">财经</a></li>
		<li><a href="//ent.sina.com.cn/" suda-uatrack="key=ty0526&value=blk_sports_nav_ent">娱乐</a></li>
		<li><a href="//tech.sina.com.cn/" suda-uatrack="key=ty0526&value=blk_sports_nav_tech">科技</a></li>
		<li><a href="//blog.sina.com.cn/" suda-uatrack="key=ty0526&value=blk_sports_nav_blog">博客</a></li>
		<li><a href="//photo.sina.com.cn/" suda-uatrack="key=ty0526&value=blk_sports_nav_pic">图片</a></li>
		<li><a href="http://zhuanlan.sina.com.cn/" suda-uatrack="key=ty0526&value=blk_sports_nav_zl">专栏</a></li>
		<li class="more"><a href="javascript:;" layer-type="layer-nav" layer-data="eventType=mouseenter" class="">...</a></li>
	</ul>
	<div class="more-layer" layer-type="layer-cont" layer-data="eventType=mouseenter" suda-uatrack="key=ty0526&value=blk_sports_nav_more" style="display:none;margin-left:102px;">
		<i class="ico ico-arrow" layer-type="arrow" style="margin-left:-102px;"></i>
		<ul class="clearfix">
			<li><a href="//auto.sina.com.cn/">汽车</a></li>
			<li><a href="//edu.sina.com.cn/">教育</a></li>
			<li><a href="//fashion.sina.com.cn/">时尚</a></li>
			<li><a href="//eladies.sina.com.cn/">女性</a></li>
			<li><a href="//astro.sina.com.cn/">星座</a></li>
			<li><a href="//med.sina.com/">医药</a></li>
			<li><a href="//www.leju.com/#source=pc_sina_dbdh2&source_ext=pc_sina">房产</a></li>
			<li><a href="//history.sina.com.cn/">历史</a></li>
			<li><a href="//video.sina.com.cn/">视频</a></li>
			<li><a href="//collection.sina.com.cn/">收藏</a></li>
			<li><a href="//baby.sina.com.cn/">育儿</a></li>
			<li><a href="//book.sina.com.cn/">读书</a></li>
			<li><a href="//fo.sina.com.cn/">佛学</a></li>
			<li><a href="//games.sina.com.cn/">游戏</a></li>
			<li><a href="//travel.sina.com.cn/">旅游</a></li>
			<li><a href="//mail.sina.com.cn/">邮箱</a></li>
			<li><a href="//news.sina.com.cn/guide/">导航</a></li>
		</ul>
	</div>
</div>
`

func TestCssSelector_QueryAll(t *testing.T) {
	sel := NewCssSelector(strings.NewReader(htmlContent))
	nodes := sel.QueryAll("div.more-layer ul li a")
	for _, node := range nodes {
		if node.FindAttr("href") == "" {
			t.Error("href is empty")
		}
	}
}

func TestCssSelector_Query(t *testing.T) {
	sel := NewCssSelector(strings.NewReader(htmlContent))
	node := sel.Query("div.more-layer ul li a")
	if node.FindAttr("href") == "" {
		t.Error("href is empty")
	}
}
