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

import "aliyun-go-sdk-core"

type DescribeCouponListRequest struct {
	Status            string
	StartDeliveryTime string
	EndDeliveryTime   string
	PageSize          int32
	PageNum           int32
}

func (r *DescribeCouponListRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 5)

	core.AddToMapIfNotDefaultValueStr(queryVals, "Status", r.Status)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StartDeliveryTime", r.StartDeliveryTime)
	core.AddToMapIfNotDefaultValueStr(queryVals, "EndDeliveryTime", r.EndDeliveryTime)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageSize", r.PageSize)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageNum", r.PageNum)

	return queryVals
}

func (r *DescribeCouponListRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *DescribeCouponListRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *DescribeCouponListRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *DescribeCouponListRequest) GetPath() string {
	return ""
}

func (r *DescribeCouponListRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *DescribeCouponListRequest) GetProtocol() string {
	return "HTTPS"
}

func (r *DescribeCouponListRequest) GetActionName() string {
	return "DescribeCouponList"
}
