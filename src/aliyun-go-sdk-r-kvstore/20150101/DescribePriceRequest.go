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

package kvstore

import "aliyun-go-sdk-core"

type DescribePriceRequest struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	OwnerAccount         string
	Capacity             int64
	OrderType            string
	ZoneId               string
	ChargeType           string
	Period               int64
	Quantity             int64
	InstanceId           string
}

func (r *DescribePriceRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 11)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "Capacity", r.Capacity)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OrderType", r.OrderType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ZoneId", r.ZoneId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ChargeType", r.ChargeType)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "Period", r.Period)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "Quantity", r.Quantity)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId", r.InstanceId)

	return queryVals
}

func (r *DescribePriceRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *DescribePriceRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *DescribePriceRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *DescribePriceRequest) GetPath() string {
	return ""
}

func (r *DescribePriceRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *DescribePriceRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *DescribePriceRequest) GetActionName() string {
	return "DescribePrice"
}
