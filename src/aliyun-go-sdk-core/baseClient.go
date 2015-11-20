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
	"aliyun-go-sdk-core/objectBuilder"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

type BaseClient struct {
	ProductName         string
	Version             string
	cookieJar           *cookiejar.Jar
	client              *http.Client
	Profile             *AccessProfile
	ApiStyle            string
	disableGlobalCookie bool
	HttpRequestBuilder  HttpRequestBuilder
}

//This method is used to do http request call, you do not need to call it directlly
//All the internal service call  dependens here
func (h *BaseClient) DoAction(reqObj BaseRequest, resObj BaseResponse) error {

	if h.cookieJar == nil {
		h.cookieJar = new(cookiejar.Jar)
	}

	requestCookieJar := h.cookieJar

	if h.disableGlobalCookie {
		requestCookieJar = nil
	}

	if h.client == nil {
		h.client = &http.Client{}
	}

	h.client.Jar = requestCookieJar

	reqBuilder := h.HttpRequestBuilder
	req := reqBuilder.Buid(h, reqObj)
	reqBuilder.Sign(h, reqObj, req)

	res, err2 := h.client.Do(req)
	CheckError(err2, fmt.Sprintf("Http request failed: %s", req.URL.String()))

	defer res.Body.Close()
	body, err3 := ioutil.ReadAll(res.Body)

	if res.StatusCode == 200 {
		CheckError(err3, fmt.Sprintf("Read reponse body failed: %s", req.URL.String()))

		var unmarshaledObj interface{}
		json.Unmarshal([]byte(body), &unmarshaledObj)

		reconstructedMap := make(map[string]interface{})
		builder.FlatJsonObjToMap("", unmarshaledObj.(map[string]interface{}), reconstructedMap)
		resObj.BuildProperties("", reconstructedMap)

		//err := builder.UnmarshalJson(responseBody, resObj)
		return nil
	} else {
		return errors.New(string(string(body)))
	}
}
