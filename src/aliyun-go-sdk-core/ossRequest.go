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
	"sort"
	"strings"
)

type OssHttpRequestBuilder struct {
}

func (r *OssHttpRequestBuilder) Buid(b *BaseClient, reqObject BaseRequest) *http.Request {
	queryArgs := reqObject.GetQuery()
	//commonArgs := getCommonRequestArguments(b, reqObject)
	//mergeMap(queryArgs, commonArgs)

	bodyArgs := reqObject.GetBody()

	reqMtd := GetRequestMethod(reqObject.GetRequestMethod())

	accessUrl := getRequestUrl(b, reqObject)

	req := buildRequest(queryArgs, bodyArgs, reqMtd, accessUrl)

	addHeader(req, "Date", parameterHelper.GetRFC2616Date())

	//oss only provides xml format
	//addHeader(req, "Accept", getAcceptContentType(b.Profile.FormatType, reqMtd))

	if len(bodyArgs) > 0 {
		bodyUrlVals := parameterHelper.ConvertMapToUrlValues(bodyArgs)
		content := bodyUrlVals.Encode()

		contentMd5 := parameterHelper.Md5Str([]byte(content))
		req.Header.Add("Content-MD5", contentMd5)
		req.Header.Add("Content-Length", fmt.Sprintf("%d", len([]byte(content))))
		req.Header.Add("Content-Type", getAcceptContentType(b.Profile.FormatType, reqMtd))
	}

	addSdkVersion(req)
	return req
}

func (*OssHttpRequestBuilder) Sign(b *BaseClient, reqObject BaseRequest, req *http.Request) {

	strArr := make([]string, 0)
	strArr = append(strArr, GetRequestMethod(reqObject.GetRequestMethod()))

	strArr = appendToArray(strArr, req, "Content-MD5")
	strArr = appendToArray(strArr, req, "Content-Type")
	strArr = appendToArray(strArr, req, "Date")

	acsKeys := make([]string, 0)

	sort.Strings(acsKeys)

	strArr = append(strArr, req.URL.RequestURI())
	tobeSign := strings.Join(strArr, HEADER_SEPARATOR)

	signatureStr := b.Profile.Signer.MakeSignature(b.Profile.AccessSecret, tobeSign)

	signVal := "OSS " + b.Profile.AccessKeyId + ":" + signatureStr
	req.Header.Add("Authorization", signVal)

}
