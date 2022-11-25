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
		rawReader io.Reader
	}

	SelNode struct {
		*html.Node
	}

	SelNodes []*SelNode
)

func NewCssSelector(reader io.Reader) *CssSelector {
	return &CssSelector{rawReader: reader}
}

func (cs *CssSelector) Query(cssExpr string) SelNodes {
	node, err := html.Parse(cs.rawReader)
	if err != nil {
		return []*SelNode{}
	}

	return newSelNode(node).Query(cssExpr)
}

func newSelNode(htmlNode *html.Node) *SelNode {
	return &SelNode{Node: htmlNode}
}

//Query find the nodes by css expression
func (sn *SelNode) Query(cssExpr string) SelNodes {
	var selector cascadia.Sel
	var err error
	if selector, err = cascadia.Parse(cssExpr); err != nil {
		return []*SelNode{}
	}

	htmlNodes := cascadia.QueryAll(sn.Node, selector)
	nodes := make([]*SelNode, 0)
	for _, node := range htmlNodes {
		nodes = append(nodes, newSelNode(node))
	}
	return nodes
}

func (sn *SelNode) QuerySingle(cssExpr string) *SelNode {
	var selector cascadia.Sel
	var err error
	if selector, err = cascadia.Parse(cssExpr); err != nil {
		return &SelNode{}
	}

	htmlNode := cascadia.Query(sn.Node, selector)
	return newSelNode(htmlNode)
}

//FindAttr find the attribute value of the node
func (sn *SelNode) FindAttr(name string) string {
	for _, attr := range sn.Attr {
		if attr.Key == name {
			return attr.Val
		}
	}
	return ""
}

func (sns SelNodes) First() *SelNode {
	if len(sns) == 0 {
		return nil
	}
	return sns[0]
}

func (sns SelNodes) Last() *SelNode {
	if len(sns) == 0 {
		return nil
	}
	return sns[len(sns)-1]
}
