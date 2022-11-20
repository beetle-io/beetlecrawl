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
)

func NewCssSelector(reader io.Reader) *CssSelector {
	return &CssSelector{rawReader: reader}
}

func (cs *CssSelector) QueryAll(cssExpr string) []*SelNode {
	node, err := html.Parse(cs.rawReader)
	if err != nil {
		return []*SelNode{}
	}

	var selector cascadia.Sel
	if selector, err = cascadia.Parse(cssExpr); err != nil {
		return []*SelNode{}
	}
	return cs.toSelNodes(cascadia.QueryAll(node, selector))
}

func (cs *CssSelector) Query(cssExpr string) *SelNode {
	node, err := html.Parse(cs.rawReader)
	if err != nil {
		return &SelNode{node}
	}

	var selector cascadia.Sel
	if selector, err = cascadia.Parse(cssExpr); err != nil {
		return &SelNode{node}
	}
	return &SelNode{cascadia.Query(node, selector)}
}

func (cs *CssSelector) toSelNodes(nodes []*html.Node) []*SelNode {
	selNodes := make([]*SelNode, 0)
	for _, node := range nodes {
		selNodes = append(selNodes, &SelNode{node})
	}
	return selNodes
}

func (sn *SelNode) FindAttr(name string) string {
	for _, attr := range sn.Attr {
		if attr.Key == name {
			return attr.Val
		}
	}
	return ""
}
