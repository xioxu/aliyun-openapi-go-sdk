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

package ace

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-ace/20150713/GetMonitorDataResponse"

	"aliyun-go-sdk-ace/20150713/DescribeAppLogsResponse"
)

type AceClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *AceClient {
	p := &AceClient{}
	p.Version = "2015-07-13"
	p.ProductName = "Ace"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *AceClient) GetMonitorData(getMonitorDataRequest *GetMonitorDataRequest) (ace_GetMonitorDataResponse.GetMonitorDataResponse, error) {
	var resObj ace_GetMonitorDataResponse.GetMonitorDataResponse

	if getMonitorDataRequest == nil {
		getMonitorDataRequest = new(GetMonitorDataRequest)
	}
	err := c.DoAction(getMonitorDataRequest, &resObj)
	return resObj, err
}

func (c *AceClient) DescribeAppLogs(describeAppLogsRequest *DescribeAppLogsRequest) (ace_DescribeAppLogsResponse.DescribeAppLogsResponse, error) {
	var resObj ace_DescribeAppLogsResponse.DescribeAppLogsResponse

	if describeAppLogsRequest == nil {
		describeAppLogsRequest = new(DescribeAppLogsRequest)
	}
	err := c.DoAction(describeAppLogsRequest, &resObj)
	return resObj, err
}
