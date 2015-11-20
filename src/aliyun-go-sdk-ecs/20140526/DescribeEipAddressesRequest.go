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

package ecs

import "aliyun-go-sdk-core"

type DescribeEipAddressesRequest struct {
	OwnerId                int64
	ResourceOwnerAccount   string
	ResourceOwnerId        int64
	Status                 string
	EipAddress             string
	AllocationId           string
	PageNumber             int32
	PageSize               int32
	OwnerAccount           string
	Filter_1_Key           string
	Filter_2_Key           string
	Filter_1_Value         string
	Filter_2_Value         string
	LockReason             string
	AssociatedInstanceType string
	AssociatedInstanceId   string
}

func (r *DescribeEipAddressesRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 16)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Status", r.Status)
	core.AddToMapIfNotDefaultValueStr(queryVals, "EipAddress", r.EipAddress)
	core.AddToMapIfNotDefaultValueStr(queryVals, "AllocationId", r.AllocationId)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageNumber", r.PageNumber)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageSize", r.PageSize)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Filter.1.Key", r.Filter_1_Key)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Filter.2.Key", r.Filter_2_Key)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Filter.1.Value", r.Filter_1_Value)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Filter.2.Value", r.Filter_2_Value)
	core.AddToMapIfNotDefaultValueStr(queryVals, "LockReason", r.LockReason)
	core.AddToMapIfNotDefaultValueStr(queryVals, "AssociatedInstanceType", r.AssociatedInstanceType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "AssociatedInstanceId", r.AssociatedInstanceId)

	return queryVals
}

func (r *DescribeEipAddressesRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *DescribeEipAddressesRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *DescribeEipAddressesRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *DescribeEipAddressesRequest) GetPath() string {
	return ""
}

func (r *DescribeEipAddressesRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *DescribeEipAddressesRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *DescribeEipAddressesRequest) GetActionName() string {
	return "DescribeEipAddresses"
}
