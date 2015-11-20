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

type DescribeScalingActivitiesRequest struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	PageNumber           int32
	PageSize             int32
	ScalingGroupId       string
	StatusCode           string
	ScalingActivityId_1  string
	ScalingActivityId_2  string
	ScalingActivityId_3  string
	ScalingActivityId_4  string
	ScalingActivityId_5  string
	ScalingActivityId_6  string
	ScalingActivityId_7  string
	ScalingActivityId_8  string
	ScalingActivityId_9  string
	ScalingActivityId_10 string
	ScalingActivityId_11 string
	ScalingActivityId_12 string
	ScalingActivityId_13 string
	ScalingActivityId_14 string
	ScalingActivityId_15 string
	ScalingActivityId_16 string
	ScalingActivityId_17 string
	ScalingActivityId_18 string
	ScalingActivityId_19 string
	ScalingActivityId_20 string
	OwnerAccount         string
}

func (r *DescribeScalingActivitiesRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 28)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageNumber", r.PageNumber)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageSize", r.PageSize)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupId", r.ScalingGroupId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StatusCode", r.StatusCode)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingActivityId.1", r.ScalingActivityId_1)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingActivityId.2", r.ScalingActivityId_2)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingActivityId.3", r.ScalingActivityId_3)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingActivityId.4", r.ScalingActivityId_4)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingActivityId.5", r.ScalingActivityId_5)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingActivityId.6", r.ScalingActivityId_6)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingActivityId.7", r.ScalingActivityId_7)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingActivityId.8", r.ScalingActivityId_8)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingActivityId.9", r.ScalingActivityId_9)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingActivityId.10", r.ScalingActivityId_10)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingActivityId.11", r.ScalingActivityId_11)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingActivityId.12", r.ScalingActivityId_12)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingActivityId.13", r.ScalingActivityId_13)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingActivityId.14", r.ScalingActivityId_14)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingActivityId.15", r.ScalingActivityId_15)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingActivityId.16", r.ScalingActivityId_16)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingActivityId.17", r.ScalingActivityId_17)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingActivityId.18", r.ScalingActivityId_18)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingActivityId.19", r.ScalingActivityId_19)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingActivityId.20", r.ScalingActivityId_20)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *DescribeScalingActivitiesRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *DescribeScalingActivitiesRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *DescribeScalingActivitiesRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *DescribeScalingActivitiesRequest) GetPath() string {
	return ""
}

func (r *DescribeScalingActivitiesRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *DescribeScalingActivitiesRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *DescribeScalingActivitiesRequest) GetActionName() string {
	return "DescribeScalingActivities"
}
