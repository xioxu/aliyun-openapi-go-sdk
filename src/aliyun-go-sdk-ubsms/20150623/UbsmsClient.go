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

package ubsms

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-ubsms/20150623/DescribeBusinessStatusResponse"

	"aliyun-go-sdk-ubsms/20150623/SetUserBusinessStatusesResponse"

	"aliyun-go-sdk-ubsms/20150623/NotifyUserBusinessCommandResponse"

	"aliyun-go-sdk-ubsms/20150623/SetUserBusinessStatusResponse"
)

type UbsmsClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *UbsmsClient {
	p := &UbsmsClient{}
	p.Version = "2015-06-23"
	p.ProductName = "Ubsms"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *UbsmsClient) DescribeBusinessStatus(describeBusinessStatusRequest *DescribeBusinessStatusRequest) (ubsms_DescribeBusinessStatusResponse.DescribeBusinessStatusResponse, error) {
	var resObj ubsms_DescribeBusinessStatusResponse.DescribeBusinessStatusResponse

	if describeBusinessStatusRequest == nil {
		describeBusinessStatusRequest = new(DescribeBusinessStatusRequest)
	}
	err := c.DoAction(describeBusinessStatusRequest, &resObj)
	return resObj, err
}

func (c *UbsmsClient) SetUserBusinessStatuses(setUserBusinessStatusesRequest *SetUserBusinessStatusesRequest) (ubsms_SetUserBusinessStatusesResponse.SetUserBusinessStatusesResponse, error) {
	var resObj ubsms_SetUserBusinessStatusesResponse.SetUserBusinessStatusesResponse

	if setUserBusinessStatusesRequest == nil {
		setUserBusinessStatusesRequest = new(SetUserBusinessStatusesRequest)
	}
	err := c.DoAction(setUserBusinessStatusesRequest, &resObj)
	return resObj, err
}

func (c *UbsmsClient) NotifyUserBusinessCommand(notifyUserBusinessCommandRequest *NotifyUserBusinessCommandRequest) (ubsms_NotifyUserBusinessCommandResponse.NotifyUserBusinessCommandResponse, error) {
	var resObj ubsms_NotifyUserBusinessCommandResponse.NotifyUserBusinessCommandResponse

	if notifyUserBusinessCommandRequest == nil {
		notifyUserBusinessCommandRequest = new(NotifyUserBusinessCommandRequest)
	}
	err := c.DoAction(notifyUserBusinessCommandRequest, &resObj)
	return resObj, err
}

func (c *UbsmsClient) SetUserBusinessStatus(setUserBusinessStatusRequest *SetUserBusinessStatusRequest) (ubsms_SetUserBusinessStatusResponse.SetUserBusinessStatusResponse, error) {
	var resObj ubsms_SetUserBusinessStatusResponse.SetUserBusinessStatusResponse

	if setUserBusinessStatusRequest == nil {
		setUserBusinessStatusRequest = new(SetUserBusinessStatusRequest)
	}
	err := c.DoAction(setUserBusinessStatusRequest, &resObj)
	return resObj, err
}
