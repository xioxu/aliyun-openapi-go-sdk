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

package oms

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-oms/20150212/GetProductDefineResponse"

	"aliyun-go-sdk-oms/20150212/GetUserDataResponse"
)

type OmsClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *OmsClient {
	p := &OmsClient{}
	p.Version = "2015-02-12"
	p.ProductName = "Oms"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *OmsClient) GetProductDefine(getProductDefineRequest *GetProductDefineRequest) (oms_GetProductDefineResponse.GetProductDefineResponse, error) {
	var resObj oms_GetProductDefineResponse.GetProductDefineResponse

	if getProductDefineRequest == nil {
		getProductDefineRequest = new(GetProductDefineRequest)
	}
	err := c.DoAction(getProductDefineRequest, &resObj)
	return resObj, err
}

func (c *OmsClient) GetUserData(getUserDataRequest *GetUserDataRequest) (oms_GetUserDataResponse.GetUserDataResponse, error) {
	var resObj oms_GetUserDataResponse.GetUserDataResponse

	if getUserDataRequest == nil {
		getUserDataRequest = new(GetUserDataRequest)
	}
	err := c.DoAction(getUserDataRequest, &resObj)
	return resObj, err
}
