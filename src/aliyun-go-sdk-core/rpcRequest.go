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
	"net/http"
	"net/url"
	"strings"
)

type RpcHttpRequestBuilder struct {
}

func (r *RpcHttpRequestBuilder) Buid(b *BaseClient, reqObject BaseRequest) *http.Request {
	queryArgs := reqObject.GetQuery()
	bodyArgs := reqObject.GetBody()

	commonArgs := getCommonRequestArguments(b, reqObject)
	mergeMap(queryArgs, commonArgs)

	reqMtd := GetRequestMethod(reqObject.GetRequestMethod())

	accessUrl := getRequestUrl(b, reqObject)

	req := buildRequest(queryArgs, bodyArgs, reqMtd, accessUrl)

	addSdkVersion(req)

	return req
}

func (*RpcHttpRequestBuilder) Sign(b *BaseClient, reqObject BaseRequest, req *http.Request) {
	queryArgs := req.URL.Query()
	bodyArgs := reqObject.GetBody()

	tobeSignArgs := make(url.Values)

	reqMtd := GetRequestMethod(reqObject.GetRequestMethod())

	//If reqMtd is 'GET', the queryArgs should already contains body arguments
	if reqMtd != "GET" {
		mergeMapToUrlValues(tobeSignArgs, bodyArgs)
	}

	mergeUrlValues(tobeSignArgs, queryArgs)

	//Refer to https://docs.aliyun.com/#/pub/ecs/open-api/requestmethod&signature
	toSignatureStr := strings.Replace(tobeSignArgs.Encode(), "+", "%20", -1)
	toSignatureStr = strings.Replace(toSignatureStr, "*", "%2A", -1)
	toSignatureStr = strings.Replace(toSignatureStr, "%7E", "~", -1)

	//The official document has a mistake which said we should replace = & with the following code, but actully we do not need it
	//toSignatureStr = strings.Replace(toSignatureStr, "=", "%3D", -1)
	//toSignatureStr = strings.Replace(toSignatureStr, "&", "%26", -1)

	toSignatureStr = strings.ToUpper(GetRequestMethod(reqObject.GetRequestMethod())) + "&%2F&" + url.QueryEscape(toSignatureStr)

	signatureStr := b.Profile.Signer.MakeSignature(b.Profile.AccessSecret+"&", toSignatureStr)

	queryArgs.Add("Signature", signatureStr)

	req.URL.RawQuery = queryArgs.Encode()
}
