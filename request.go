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
	"log"
	"net/http"
)

const (
	GET  = "GET"
	POST = "POST"
)

type (
	HttpFailFunc func(errReq *HttpRequest, errors []error) error

	Request interface {
		//URL returns the URL of the request
		URL() string

		//SetOnFail sets the callback function to be called when the request fails
		SetOnFail(failFunc HttpFailFunc)
	}

	//HttpRequest is a wrapper around http.Request, user send this to scheduler for fetching the web page
	HttpRequest struct {
		*http.Request
		respCb        HttpResponseFunc
		failCb        HttpFailFunc
		successRespCh chan *HttpResponse
		failReqCh     chan *HttpRequest
		//downloader will try download the page for maxRetry times, add all errors to this field
		retryTimes int
		errs       []error
	}
)

func NewHttpRequest(method string, url string, respFunc HttpResponseFunc) *HttpRequest {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	httpReq := &HttpRequest{
		successRespCh: make(chan *HttpResponse),
		respCb:        respFunc,
		failCb:        nil,
		Request:       request,
		errs:          make([]error, 0),
		retryTimes:    0,
	}
	return httpReq
}

func (br *HttpRequest) SetOnFail(failFunc HttpFailFunc) {
	br.failCb = failFunc
}

func (br *HttpRequest) URL() string {
	return br.Request.URL.String()
}

func (br *HttpRequest) RetryTimes() int {
	return br.retryTimes
}
