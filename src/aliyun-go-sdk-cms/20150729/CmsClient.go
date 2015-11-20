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

package cms

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-cms/20150729/PutMetricDatumResponse"

	"aliyun-go-sdk-cms/20150729/DescribeMetricDatumResponse"
)

type CmsClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *CmsClient {
	p := &CmsClient{}
	p.Version = "2015-07-29"
	p.ProductName = "Cms"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *CmsClient) PutMetricDatum(putMetricDatumRequest *PutMetricDatumRequest) (cms_PutMetricDatumResponse.PutMetricDatumResponse, error) {
	var resObj cms_PutMetricDatumResponse.PutMetricDatumResponse

	if putMetricDatumRequest == nil {
		putMetricDatumRequest = new(PutMetricDatumRequest)
	}
	err := c.DoAction(putMetricDatumRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) DescribeMetricDatum(describeMetricDatumRequest *DescribeMetricDatumRequest) (cms_DescribeMetricDatumResponse.DescribeMetricDatumResponse, error) {
	var resObj cms_DescribeMetricDatumResponse.DescribeMetricDatumResponse

	if describeMetricDatumRequest == nil {
		describeMetricDatumRequest = new(DescribeMetricDatumRequest)
	}
	err := c.DoAction(describeMetricDatumRequest, &resObj)
	return resObj, err
}
