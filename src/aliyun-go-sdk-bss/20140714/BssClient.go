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

package bss

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-bss/20140714/SetResourceBusinessStatusResponse"

	"aliyun-go-sdk-bss/20140714/DescribeCouponListResponse"

	"aliyun-go-sdk-bss/20140714/DescribeCouponDetailResponse"

	"aliyun-go-sdk-bss/20140714/DescribeCashDetailResponse"
)

type BssClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *BssClient {
	p := &BssClient{}
	p.Version = "2014-07-14"
	p.ProductName = "Bss"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *BssClient) SetResourceBusinessStatus(setResourceBusinessStatusRequest *SetResourceBusinessStatusRequest) (bss_SetResourceBusinessStatusResponse.SetResourceBusinessStatusResponse, error) {
	var resObj bss_SetResourceBusinessStatusResponse.SetResourceBusinessStatusResponse

	if setResourceBusinessStatusRequest == nil {
		setResourceBusinessStatusRequest = new(SetResourceBusinessStatusRequest)
	}
	err := c.DoAction(setResourceBusinessStatusRequest, &resObj)
	return resObj, err
}

func (c *BssClient) DescribeCouponList(describeCouponListRequest *DescribeCouponListRequest) (bss_DescribeCouponListResponse.DescribeCouponListResponse, error) {
	var resObj bss_DescribeCouponListResponse.DescribeCouponListResponse

	if describeCouponListRequest == nil {
		describeCouponListRequest = new(DescribeCouponListRequest)
	}
	err := c.DoAction(describeCouponListRequest, &resObj)
	return resObj, err
}

func (c *BssClient) DescribeCouponDetail(describeCouponDetailRequest *DescribeCouponDetailRequest) (bss_DescribeCouponDetailResponse.DescribeCouponDetailResponse, error) {
	var resObj bss_DescribeCouponDetailResponse.DescribeCouponDetailResponse

	if describeCouponDetailRequest == nil {
		describeCouponDetailRequest = new(DescribeCouponDetailRequest)
	}
	err := c.DoAction(describeCouponDetailRequest, &resObj)
	return resObj, err
}

func (c *BssClient) DescribeCashDetail(describeCashDetailRequest *DescribeCashDetailRequest) (bss_DescribeCashDetailResponse.DescribeCashDetailResponse, error) {
	var resObj bss_DescribeCashDetailResponse.DescribeCashDetailResponse

	if describeCashDetailRequest == nil {
		describeCashDetailRequest = new(DescribeCashDetailRequest)
	}
	err := c.DoAction(describeCashDetailRequest, &resObj)
	return resObj, err
}
