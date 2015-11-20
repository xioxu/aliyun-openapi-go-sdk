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

package inner

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-ubsms-inner/20150623/DescribeBusinessStatusResponse"

	"aliyun-go-sdk-ubsms-inner/20150623/SetUserSecurityStatusResponse"

	"aliyun-go-sdk-ubsms-inner/20150623/SetUserBusinessStatusesResponse"

	"aliyun-go-sdk-ubsms-inner/20150623/NotifyUserBusinessCommandResponse"

	"aliyun-go-sdk-ubsms-inner/20150623/DescribeUserBusinessStatusResponse"
)

type InnerClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *InnerClient {
	p := &InnerClient{}
	p.Version = "2015-06-23"
	p.ProductName = "Ubsms-inner"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *InnerClient) DescribeBusinessStatus(describeBusinessStatusRequest *DescribeBusinessStatusRequest) (inner_DescribeBusinessStatusResponse.DescribeBusinessStatusResponse, error) {
	var resObj inner_DescribeBusinessStatusResponse.DescribeBusinessStatusResponse

	if describeBusinessStatusRequest == nil {
		describeBusinessStatusRequest = new(DescribeBusinessStatusRequest)
	}
	err := c.DoAction(describeBusinessStatusRequest, &resObj)
	return resObj, err
}

func (c *InnerClient) SetUserSecurityStatus(setUserSecurityStatusRequest *SetUserSecurityStatusRequest) (inner_SetUserSecurityStatusResponse.SetUserSecurityStatusResponse, error) {
	var resObj inner_SetUserSecurityStatusResponse.SetUserSecurityStatusResponse

	if setUserSecurityStatusRequest == nil {
		setUserSecurityStatusRequest = new(SetUserSecurityStatusRequest)
	}
	err := c.DoAction(setUserSecurityStatusRequest, &resObj)
	return resObj, err
}

func (c *InnerClient) SetUserBusinessStatuses(setUserBusinessStatusesRequest *SetUserBusinessStatusesRequest) (inner_SetUserBusinessStatusesResponse.SetUserBusinessStatusesResponse, error) {
	var resObj inner_SetUserBusinessStatusesResponse.SetUserBusinessStatusesResponse

	if setUserBusinessStatusesRequest == nil {
		setUserBusinessStatusesRequest = new(SetUserBusinessStatusesRequest)
	}
	err := c.DoAction(setUserBusinessStatusesRequest, &resObj)
	return resObj, err
}

func (c *InnerClient) NotifyUserBusinessCommand(notifyUserBusinessCommandRequest *NotifyUserBusinessCommandRequest) (inner_NotifyUserBusinessCommandResponse.NotifyUserBusinessCommandResponse, error) {
	var resObj inner_NotifyUserBusinessCommandResponse.NotifyUserBusinessCommandResponse

	if notifyUserBusinessCommandRequest == nil {
		notifyUserBusinessCommandRequest = new(NotifyUserBusinessCommandRequest)
	}
	err := c.DoAction(notifyUserBusinessCommandRequest, &resObj)
	return resObj, err
}

func (c *InnerClient) DescribeUserBusinessStatus(describeUserBusinessStatusRequest *DescribeUserBusinessStatusRequest) (inner_DescribeUserBusinessStatusResponse.DescribeUserBusinessStatusResponse, error) {
	var resObj inner_DescribeUserBusinessStatusResponse.DescribeUserBusinessStatusResponse

	if describeUserBusinessStatusRequest == nil {
		describeUserBusinessStatusRequest = new(DescribeUserBusinessStatusRequest)
	}
	err := c.DoAction(describeUserBusinessStatusRequest, &resObj)
	return resObj, err
}
