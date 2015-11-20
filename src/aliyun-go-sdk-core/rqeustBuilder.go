/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package core

import (
	"aliyun-go-sdk-core/parameterHelper"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	RequestMethodSeperator = "|"
)

//HttpRequestBuilder is an interface definition for build http request
type HttpRequestBuilder interface {
	Buid(*BaseClient, BaseRequest) *http.Request
	Sign(*BaseClient, BaseRequest, *http.Request)
}

//Returns the first valid http request method from reqmethods
//Simple reqmethods: GET、GET|POST、PUT...
func GetRequestMethod(reqmethods string) string {
	reqMtdArr := strings.Split(reqmethods, RequestMethodSeperator)
	retVal := strings.ToUpper(reqMtdArr[0])

	if retVal == "GET" || retVal == "POST" || retVal == "DELETE" || retVal == "PUT" || retVal == "MOVE" {
		return retVal
	} else {
		panic(fmt.Sprintf("Invalid http request method:%s", retVal))
	}
}

func getRequestUrl(b *BaseClient, reqObj BaseRequest) string {
	protocol := strings.ToUpper(reqObj.GetProtocol())

	if protocol == "" {
		panic("Wrong accessing protoco:" + protocol)
	}

	if protocol != "HTTP" && protocol != "HTTPS" {
		protocol = strings.Split(protocol, "|")[0]
	}

	return protocol + "://" + concatUrl(GetProductUrl(b), reqObj.GetPath())
}

func GetProductUrl(h *BaseClient) string {
	return h.Profile.GetServiceUrl(h.ProductName)
}

//Returns the common accesing arguments described in
//https://docs.aliyun.com/#/pub/ecs/open-api/requestmethod&commonparams
func getCommonRequestArguments(b *BaseClient, request BaseRequest) map[string]string {
	var commonArgs = map[string]string{
		"Format":           b.Profile.FormatType,
		"Version":          b.Version,
		"AccessKeyId":      b.Profile.AccessKeyId,
		"SignatureMethod":  b.Profile.Signer.GetSignatureMethod(),
		"Timestamp":        parameterHelper.GetISO8601Time(),
		"SignatureVersion": b.Profile.Signer.GetVersion(),
		"SignatureNonce":   parameterHelper.GetSignatureNonce(),
		"Action":           request.GetActionName(),
		"RegionId":         b.Profile.RegionId,
	}

	return commonArgs
}

func concatUrl(baseUrl string, subPath string) string {
	if subPath == "" || baseUrl == "" {
		return baseUrl
	}

	if strings.HasSuffix(baseUrl, "/") {
		if strings.HasPrefix(subPath, "/") {
			return baseUrl + subPath[1:]
		} else {
			return baseUrl + subPath
		}
	} else {
		if strings.HasPrefix(subPath, "/") {
			return baseUrl + subPath
		} else {
			return baseUrl + "/" + subPath
		}

	}
}

func mergeMap(to map[string]string, source map[string]string) {
	for k, v := range source {
		to[k] = v
	}
}

func mergeMapToUrlValues(to url.Values, source map[string]string) {
	for k, v := range source {
		to.Add(k, v)
	}
}

func mergeUrlValues(to url.Values, source url.Values) {
	for k, v := range source {
		for _, v1 := range v {
			to.Add(k, v1)
		}
	}
}

func getAcceptContentType(requestFormat string, requestMtd string) string {
	if strings.ToLower(requestMtd) == "put" {
		return "application/octet-stream"
	}

	if strings.ToLower(requestFormat) == "xml" {
		return "application/xml"
	} else if strings.ToLower(requestFormat) == "json" {
		return "application/json"
	} else {
		return "application/octet-stream"
	}
}

//Add sdk version infomation to request head
func addSdkVersion(req *http.Request) {
	req.Header.Del("x-sdk-client")
	req.Header.Add("x-sdk-client", "Golang/1.0.0")
}

func buildRequest(queryArgs, bodyArgs map[string]string, reqMtd, accessUrl string) *http.Request {
	var req *http.Request
	if reqMtd == "GET" {

		mergeMap(queryArgs, bodyArgs)

		req, _ = http.NewRequest(reqMtd, accessUrl, nil)

	} else {
		req, _ = http.NewRequest(reqMtd, accessUrl, strings.NewReader(parameterHelper.ConvertMapToUrlValues(bodyArgs).Encode()))
	}

	currQryVals := req.URL.Query()
	mergeUrlValues(currQryVals, parameterHelper.ConvertMapToUrlValues(queryArgs))
	req.URL.RawQuery = currQryVals.Encode()

	return req
}
