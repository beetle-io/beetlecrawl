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
	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
	"io"
)

//CssSelector is the css selector
type (
	CssSelector struct {
		*html.Node
	}

	CssSelectors []*CssSelector
)

func NewCssSelector(reader io.Reader) *CssSelector {
	node, err := html.Parse(reader)
	if err != nil {
		return nil
	}
	return newCssSelNode(node)
}

func newCssSelNode(htmlNode *html.Node) *CssSelector {
	return &CssSelector{Node: htmlNode}
}

//Query find the nodes by css expression
func (sn *CssSelector) Query(cssExpr string) CssSelectors {
	var selector cascadia.Sel
	var err error
	if selector, err = cascadia.Parse(cssExpr); err != nil {
		return []*CssSelector{}
	}

	htmlNodes := cascadia.QueryAll(sn.Node, selector)
	nodes := make([]*CssSelector, 0)
	for _, node := range htmlNodes {
		nodes = append(nodes, newCssSelNode(node))
	}
	return nodes
}

func (sn *CssSelector) QuerySingle(cssExpr string) *CssSelector {
	var selector cascadia.Sel
	var err error
	if selector, err = cascadia.Parse(cssExpr); err != nil {
		return &CssSelector{}
	}

	htmlNode := cascadia.Query(sn.Node, selector)
	return newCssSelNode(htmlNode)
}

//FindAttr find the attribute value of the node
func (sn *CssSelector) FindAttr(name string) string {
	for _, attr := range sn.Attr {
		if attr.Key == name {
			return attr.Val
		}
	}
	return ""
}

func (sns CssSelectors) First() *CssSelector {
	if len(sns) == 0 {
		return nil
	}
	return sns[0]
}

func (sns CssSelectors) Last() *CssSelector {
	if len(sns) == 0 {
		return nil
	}
	return sns[len(sns)-1]
}
