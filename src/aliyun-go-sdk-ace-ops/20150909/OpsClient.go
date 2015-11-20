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

package ops

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-ace-ops/20150909/ReportResponse"

	"aliyun-go-sdk-ace-ops/20150909/QueryResponse"
)

type OpsClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *OpsClient {
	p := &OpsClient{}
	p.Version = "2015-09-09"
	p.ProductName = "Ace-ops"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *OpsClient) Report(reportRequest *ReportRequest) (ops_ReportResponse.ReportResponse, error) {
	var resObj ops_ReportResponse.ReportResponse

	if reportRequest == nil {
		reportRequest = new(ReportRequest)
	}
	err := c.DoAction(reportRequest, &resObj)
	return resObj, err
}

func (c *OpsClient) Query(queryRequest *QueryRequest) (ops_QueryResponse.QueryResponse, error) {
	var resObj ops_QueryResponse.QueryResponse

	if queryRequest == nil {
		queryRequest = new(QueryRequest)
	}
	err := c.DoAction(queryRequest, &resObj)
	return resObj, err
}
