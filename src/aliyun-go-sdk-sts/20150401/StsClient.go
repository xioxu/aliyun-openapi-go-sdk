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

package sts

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-sts/20150401/AssumeRoleWithServiceIdentityResponse"

	"aliyun-go-sdk-sts/20150401/AssumeRoleResponse"
)

type StsClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *StsClient {
	p := &StsClient{}
	p.Version = "2015-04-01"
	p.ProductName = "Sts"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *StsClient) AssumeRoleWithServiceIdentity(assumeRoleWithServiceIdentityRequest *AssumeRoleWithServiceIdentityRequest) (sts_AssumeRoleWithServiceIdentityResponse.AssumeRoleWithServiceIdentityResponse, error) {
	var resObj sts_AssumeRoleWithServiceIdentityResponse.AssumeRoleWithServiceIdentityResponse

	if assumeRoleWithServiceIdentityRequest == nil {
		assumeRoleWithServiceIdentityRequest = new(AssumeRoleWithServiceIdentityRequest)
	}
	err := c.DoAction(assumeRoleWithServiceIdentityRequest, &resObj)
	return resObj, err
}

func (c *StsClient) AssumeRole(assumeRoleRequest *AssumeRoleRequest) (sts_AssumeRoleResponse.AssumeRoleResponse, error) {
	var resObj sts_AssumeRoleResponse.AssumeRoleResponse

	if assumeRoleRequest == nil {
		assumeRoleRequest = new(AssumeRoleRequest)
	}
	err := c.DoAction(assumeRoleRequest, &resObj)
	return resObj, err
}
