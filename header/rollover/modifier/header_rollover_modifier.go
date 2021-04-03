// Copyright 2015 Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package modifier

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/devopsfaith/krakend/logging"
	"github.com/google/martian/parse"
)

type AppendModifier struct {
	Name               string   `json:"name"`
	Values             []string `json:"values"`
	PrintSelectedValue bool     `json:"printSelectedValue"`
}

// ModifyRequest appends the header at name with value to the request.
func (m *AppendModifier) ModifyRequest(req *http.Request) error {
	if m.PrintSelectedValue {
		logger, _ := logging.NewLogger("DEBUG", os.Stdout, "")
		logger.Debug("selected header value: " + m.Values[0])
	}

	return RequestHeader(req).Set(m.Name, m.Values[0])
}

type Header struct {
	h http.Header

	host func() string
	cl   func() int64
	te   func() []string

	setHost func(string)
	setCL   func(int64)
	setTE   func([]string)
}

func RequestHeader(req *http.Request) *Header {
	return &Header{
		h:       req.Header,
		host:    func() string { return req.Host },
		cl:      func() int64 { return req.ContentLength },
		te:      func() []string { return req.TransferEncoding },
		setHost: func(host string) { req.Host = host },
		setCL:   func(cl int64) { req.ContentLength = cl },
		setTE:   func(te []string) { req.TransferEncoding = te },
	}
}

func (h *Header) Set(name, value string) error {
	switch http.CanonicalHeaderKey(name) {
	case "Host":
		h.setHost(value)
	case "Content-Length":
		cl, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}

		h.setCL(cl)
	case "Transfer-Encoding":
		h.setTE([]string{value})
	default:
		h.h.Set(name, value)
	}

	return nil
}

// AppendModifierFromJSON takes a JSON message as a byte slice and returns
// an appendModifier and an error.
//
// Example JSON configuration message:
// {
//  "scope": ["request", "result"],
//  "name": "X-Martian",
//  "value": "true"
// }
func FromJSON(b []byte) (*parse.Result, error) {
	msg := &AppendModifier{}
	if err := json.Unmarshal(b, msg); err != nil {
		return nil, err
	}

	return parse.NewResult(msg, []parse.ModifierType{parse.Request})
}
