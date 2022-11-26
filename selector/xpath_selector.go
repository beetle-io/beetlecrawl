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
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"io"
)

type (
	XPathSelector struct {
		*html.Node
	}

	XPathSelectors []*XPathSelector
)

func NewXPathSelector(reader io.Reader) *XPathSelector {
	node, err := html.Parse(reader)
	if err != nil {
		return nil
	}
	return newXpathSelNode(node)
}

func newXpathSelNode(htmlNode *html.Node) *XPathSelector {
	return &XPathSelector{Node: htmlNode}
}

func (xs *XPathSelector) Query(xpathExpr string) XPathSelectors {
	htmlNodes := htmlquery.Find(xs.Node, xpathExpr)
	nodes := make([]*XPathSelector, 0)
	for _, node := range htmlNodes {
		nodes = append(nodes, newXpathSelNode(node))
	}
	return nodes
}

func (xs *XPathSelector) QuerySingle(xpathExpr string) *XPathSelector {
	node := htmlquery.FindOne(xs.Node, xpathExpr)
	return newXpathSelNode(node)
}

func (xs *XPathSelector) FindAttr(name string) string {
	return htmlquery.SelectAttr(xs.Node, name)
}

func (xls XPathSelectors) First() *XPathSelector {
	if len(xls) > 0 {
		return xls[0]
	}
	return nil
}

func (xls XPathSelectors) Last() *XPathSelector {
	if len(xls) > 0 {
		return xls[len(xls)-1]
	}
	return nil
}
