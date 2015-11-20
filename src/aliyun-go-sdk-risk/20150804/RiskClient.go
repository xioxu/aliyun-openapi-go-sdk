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

package risk

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-risk/20150804/FindRiskResponse"

	"aliyun-go-sdk-risk/20150804/WriteUsscResponse"

	"aliyun-go-sdk-risk/20150804/ValidateVerifyCodeResponse"

	"aliyun-go-sdk-risk/20150804/SendVerifyCodeResponse"

	"aliyun-go-sdk-risk/20150804/QueryNameListResponse"
)

type RiskClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *RiskClient {
	p := &RiskClient{}
	p.Version = "2015-08-04"
	p.ProductName = "Risk"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *RiskClient) FindRisk(findRiskRequest *FindRiskRequest) (risk_FindRiskResponse.FindRiskResponse, error) {
	var resObj risk_FindRiskResponse.FindRiskResponse

	if findRiskRequest == nil {
		findRiskRequest = new(FindRiskRequest)
	}
	err := c.DoAction(findRiskRequest, &resObj)
	return resObj, err
}

func (c *RiskClient) WriteUssc(writeUsscRequest *WriteUsscRequest) (risk_WriteUsscResponse.WriteUsscResponse, error) {
	var resObj risk_WriteUsscResponse.WriteUsscResponse

	if writeUsscRequest == nil {
		writeUsscRequest = new(WriteUsscRequest)
	}
	err := c.DoAction(writeUsscRequest, &resObj)
	return resObj, err
}

func (c *RiskClient) ValidateVerifyCode(validateVerifyCodeRequest *ValidateVerifyCodeRequest) (risk_ValidateVerifyCodeResponse.ValidateVerifyCodeResponse, error) {
	var resObj risk_ValidateVerifyCodeResponse.ValidateVerifyCodeResponse

	if validateVerifyCodeRequest == nil {
		validateVerifyCodeRequest = new(ValidateVerifyCodeRequest)
	}
	err := c.DoAction(validateVerifyCodeRequest, &resObj)
	return resObj, err
}

func (c *RiskClient) SendVerifyCode(sendVerifyCodeRequest *SendVerifyCodeRequest) (risk_SendVerifyCodeResponse.SendVerifyCodeResponse, error) {
	var resObj risk_SendVerifyCodeResponse.SendVerifyCodeResponse

	if sendVerifyCodeRequest == nil {
		sendVerifyCodeRequest = new(SendVerifyCodeRequest)
	}
	err := c.DoAction(sendVerifyCodeRequest, &resObj)
	return resObj, err
}

func (c *RiskClient) QueryNameList(queryNameListRequest *QueryNameListRequest) (risk_QueryNameListResponse.QueryNameListResponse, error) {
	var resObj risk_QueryNameListResponse.QueryNameListResponse

	if queryNameListRequest == nil {
		queryNameListRequest = new(QueryNameListRequest)
	}
	err := c.DoAction(queryNameListRequest, &resObj)
	return resObj, err
}
