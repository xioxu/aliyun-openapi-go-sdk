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

	"aliyun-go-sdk-ossadmin/20140326/StopOssResponse"

	"aliyun-go-sdk-ossadmin/20140326/RestartOssResponse"
)

type OssadminClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *OssadminClient {
	p := &OssadminClient{}
	p.Version = "2014-03-26"
	p.ProductName = "OssAdmin"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *OssadminClient) StopOss(stopOssRequest *StopOssRequest) (ossadmin_StopOssResponse.StopOssResponse, error) {
	var resObj ossadmin_StopOssResponse.StopOssResponse

	if stopOssRequest == nil {
		stopOssRequest = new(StopOssRequest)
	}
	err := c.DoAction(stopOssRequest, &resObj)
	return resObj, err
}

func (c *OssadminClient) RestartOss(restartOssRequest *RestartOssRequest) (ossadmin_RestartOssResponse.RestartOssResponse, error) {
	var resObj ossadmin_RestartOssResponse.RestartOssResponse

	if restartOssRequest == nil {
		restartOssRequest = new(RestartOssRequest)
	}
	err := c.DoAction(restartOssRequest, &resObj)
	return resObj, err
}
