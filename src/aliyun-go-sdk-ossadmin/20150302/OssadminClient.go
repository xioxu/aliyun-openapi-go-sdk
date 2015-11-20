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

package ossadmin

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-ossadmin/20150302/StopOssInstanceResponse"

	"aliyun-go-sdk-ossadmin/20150302/RestartOssInstanceResponse"

	"aliyun-go-sdk-ossadmin/20150302/ReleaseOssInstanceResponse"

	"aliyun-go-sdk-ossadmin/20150302/CreateOssInstanceResponse"
)

type OssadminClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *OssadminClient {
	p := &OssadminClient{}
	p.Version = "2015-03-02"
	p.ProductName = "OssAdmin"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *OssadminClient) StopOssInstance(stopOssInstanceRequest *StopOssInstanceRequest) (ossadmin_StopOssInstanceResponse.StopOssInstanceResponse, error) {
	var resObj ossadmin_StopOssInstanceResponse.StopOssInstanceResponse

	if stopOssInstanceRequest == nil {
		stopOssInstanceRequest = new(StopOssInstanceRequest)
	}
	err := c.DoAction(stopOssInstanceRequest, &resObj)
	return resObj, err
}

func (c *OssadminClient) RestartOssInstance(restartOssInstanceRequest *RestartOssInstanceRequest) (ossadmin_RestartOssInstanceResponse.RestartOssInstanceResponse, error) {
	var resObj ossadmin_RestartOssInstanceResponse.RestartOssInstanceResponse

	if restartOssInstanceRequest == nil {
		restartOssInstanceRequest = new(RestartOssInstanceRequest)
	}
	err := c.DoAction(restartOssInstanceRequest, &resObj)
	return resObj, err
}

func (c *OssadminClient) ReleaseOssInstance(releaseOssInstanceRequest *ReleaseOssInstanceRequest) (ossadmin_ReleaseOssInstanceResponse.ReleaseOssInstanceResponse, error) {
	var resObj ossadmin_ReleaseOssInstanceResponse.ReleaseOssInstanceResponse

	if releaseOssInstanceRequest == nil {
		releaseOssInstanceRequest = new(ReleaseOssInstanceRequest)
	}
	err := c.DoAction(releaseOssInstanceRequest, &resObj)
	return resObj, err
}

func (c *OssadminClient) CreateOssInstance(createOssInstanceRequest *CreateOssInstanceRequest) (ossadmin_CreateOssInstanceResponse.CreateOssInstanceResponse, error) {
	var resObj ossadmin_CreateOssInstanceResponse.CreateOssInstanceResponse

	if createOssInstanceRequest == nil {
		createOssInstanceRequest = new(CreateOssInstanceRequest)
	}
	err := c.DoAction(createOssInstanceRequest, &resObj)
	return resObj, err
}
