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

import (
	"github.com/x-debug/beetlecrawl/selector"
	"io"
	"net/http"
)

type (
	HttpResponseFunc func(resp *HttpResponse) error

	Response interface {
		Request() Request
	}

	//HttpResponse is the http response
	HttpResponse struct {
		*http.Response
		req    *HttpRequest
		cssSel *selector.CssSelector
	}
)

func (br *HttpResponse) GetRequest() *HttpRequest {
	return br.req
}

func (br *HttpResponse) BodyBytes() ([]byte, error) {
	return io.ReadAll(br.Body)
}

func (br *HttpResponse) BodyString() (string, error) {
	bodyBytes, err := br.BodyBytes()
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}

func (br *HttpResponse) Css() *selector.CssSelector {
	if br.cssSel == nil {
		br.cssSel = selector.NewCssSelector(br.Body)
	}
	return br.cssSel
}

func newHttpResponse(nativeResp *http.Response, req *HttpRequest) *HttpResponse {
	return &HttpResponse{nativeResp, req, nil}
}
