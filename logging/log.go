//
// Copyright (c) 2018 Juniper Networks, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logging

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Function doesn't return error, because it is just for logging.
// If conversion to json returns error we want to log variable as raw
func VariableToJSON(variable interface{}) string {
	jsonOutput, err := json.Marshal(variable)
	if err != nil {
		log.Debugln("Converting to JSON error:", err)
		return fmt.Sprintf("Cannot convert request to JSON. Raw output: %s", variable)
	}
	return string(jsonOutput)
}

func HTTPMessageLogger(param interface{}) string {

	if param == nil {
		return "Empty request/response. Some error occured."
	}

	var body io.ReadCloser

	switch param.(type) {
	default:
		log.Debugln("HTTPMessageLogger run with wrong parameter (not request or response)")
		return "HTTPMessageLogger run with wrong parameter (not request or response)"
	case *http.Request:
		body = (param.(*http.Request)).Body
	case *http.Response:
		body = (param.(*http.Response)).Body
	}

	var buf []byte
	var err error

	if body == nil {
		buf = []byte("Body is empty.")
	} else {
		buf, err = ioutil.ReadAll(body)
		if err != nil {
			log.Debugln("Cannot read request/response body.", err)
			buf = []byte("")
		}
	}

	switch param.(type) {
	case *http.Request:
		request := param.(*http.Request)
		request.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
		return fmt.Sprintf("[%s]=>[%s] { Header: %s, Body: %s }", request.Method, request.URL.String(), VariableToJSON(request.Header), buf)
	case *http.Response:
		response := param.(*http.Response)
		response.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
		return fmt.Sprintf("{ Status: %s, Header: %s, Body: %s }", response.Status, VariableToJSON(response.Header), buf)
	}

	// NEVER REACHED
	return ""
}
