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

package ess

import "aliyun-go-sdk-core"

type DescribeScalingGroupsRequest struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	PageNumber           int32
	PageSize             int32
	ScalingGroupId_1     string
	ScalingGroupId_2     string
	ScalingGroupId_3     string
	ScalingGroupId_4     string
	ScalingGroupId_5     string
	ScalingGroupId_6     string
	ScalingGroupId_7     string
	ScalingGroupId_8     string
	ScalingGroupId_9     string
	ScalingGroupId_10    string
	ScalingGroupId_12    string
	ScalingGroupId_13    string
	ScalingGroupId_14    string
	ScalingGroupId_15    string
	ScalingGroupId_16    string
	ScalingGroupId_17    string
	ScalingGroupId_18    string
	ScalingGroupId_19    string
	ScalingGroupId_20    string
	ScalingGroupName_1   string
	ScalingGroupName_2   string
	ScalingGroupName_3   string
	ScalingGroupName_4   string
	ScalingGroupName_5   string
	ScalingGroupName_6   string
	ScalingGroupName_7   string
	ScalingGroupName_8   string
	ScalingGroupName_9   string
	ScalingGroupName_10  string
	ScalingGroupName_11  string
	ScalingGroupName_12  string
	ScalingGroupName_13  string
	ScalingGroupName_14  string
	ScalingGroupName_15  string
	ScalingGroupName_16  string
	ScalingGroupName_17  string
	ScalingGroupName_18  string
	ScalingGroupName_19  string
	ScalingGroupName_20  string
	OwnerAccount         string
}

func (r *DescribeScalingGroupsRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 45)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageNumber", r.PageNumber)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageSize", r.PageSize)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupId.1", r.ScalingGroupId_1)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupId.2", r.ScalingGroupId_2)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupId.3", r.ScalingGroupId_3)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupId.4", r.ScalingGroupId_4)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupId.5", r.ScalingGroupId_5)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupId.6", r.ScalingGroupId_6)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupId.7", r.ScalingGroupId_7)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupId.8", r.ScalingGroupId_8)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupId.9", r.ScalingGroupId_9)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupId.10", r.ScalingGroupId_10)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupId.12", r.ScalingGroupId_12)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupId.13", r.ScalingGroupId_13)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupId.14", r.ScalingGroupId_14)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupId.15", r.ScalingGroupId_15)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupId.16", r.ScalingGroupId_16)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupId.17", r.ScalingGroupId_17)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupId.18", r.ScalingGroupId_18)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupId.19", r.ScalingGroupId_19)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupId.20", r.ScalingGroupId_20)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupName.1", r.ScalingGroupName_1)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupName.2", r.ScalingGroupName_2)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupName.3", r.ScalingGroupName_3)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupName.4", r.ScalingGroupName_4)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupName.5", r.ScalingGroupName_5)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupName.6", r.ScalingGroupName_6)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupName.7", r.ScalingGroupName_7)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupName.8", r.ScalingGroupName_8)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupName.9", r.ScalingGroupName_9)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupName.10", r.ScalingGroupName_10)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupName.11", r.ScalingGroupName_11)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupName.12", r.ScalingGroupName_12)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupName.13", r.ScalingGroupName_13)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupName.14", r.ScalingGroupName_14)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupName.15", r.ScalingGroupName_15)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupName.16", r.ScalingGroupName_16)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupName.17", r.ScalingGroupName_17)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupName.18", r.ScalingGroupName_18)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupName.19", r.ScalingGroupName_19)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupName.20", r.ScalingGroupName_20)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *DescribeScalingGroupsRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *DescribeScalingGroupsRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *DescribeScalingGroupsRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *DescribeScalingGroupsRequest) GetPath() string {
	return ""
}

func (r *DescribeScalingGroupsRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *DescribeScalingGroupsRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *DescribeScalingGroupsRequest) GetActionName() string {
	return "DescribeScalingGroups"
}
