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

type DescribeScalingConfigurationsRequest struct {
	OwnerId                     int64
	ResourceOwnerAccount        string
	ResourceOwnerId             int64
	PageNumber                  int32
	PageSize                    int32
	ScalingGroupId              string
	ScalingConfigurationId_1    string
	ScalingConfigurationId_2    string
	ScalingConfigurationId_3    string
	ScalingConfigurationId_4    string
	ScalingConfigurationId_5    string
	ScalingConfigurationId_6    string
	ScalingConfigurationId_7    string
	ScalingConfigurationId_8    string
	ScalingConfigurationId_9    string
	ScalingConfigurationId_10   string
	ScalingConfigurationName_1  string
	ScalingConfigurationName_2  string
	ScalingConfigurationName_3  string
	ScalingConfigurationName_4  string
	ScalingConfigurationName_5  string
	ScalingConfigurationName_6  string
	ScalingConfigurationName_7  string
	ScalingConfigurationName_8  string
	ScalingConfigurationName_9  string
	ScalingConfigurationName_10 string
	OwnerAccount                string
}

func (r *DescribeScalingConfigurationsRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 27)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageNumber", r.PageNumber)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageSize", r.PageSize)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupId", r.ScalingGroupId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingConfigurationId.1", r.ScalingConfigurationId_1)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingConfigurationId.2", r.ScalingConfigurationId_2)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingConfigurationId.3", r.ScalingConfigurationId_3)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingConfigurationId.4", r.ScalingConfigurationId_4)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingConfigurationId.5", r.ScalingConfigurationId_5)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingConfigurationId.6", r.ScalingConfigurationId_6)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingConfigurationId.7", r.ScalingConfigurationId_7)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingConfigurationId.8", r.ScalingConfigurationId_8)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingConfigurationId.9", r.ScalingConfigurationId_9)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingConfigurationId.10", r.ScalingConfigurationId_10)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingConfigurationName.1", r.ScalingConfigurationName_1)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingConfigurationName.2", r.ScalingConfigurationName_2)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingConfigurationName.3", r.ScalingConfigurationName_3)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingConfigurationName.4", r.ScalingConfigurationName_4)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingConfigurationName.5", r.ScalingConfigurationName_5)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingConfigurationName.6", r.ScalingConfigurationName_6)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingConfigurationName.7", r.ScalingConfigurationName_7)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingConfigurationName.8", r.ScalingConfigurationName_8)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingConfigurationName.9", r.ScalingConfigurationName_9)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingConfigurationName.10", r.ScalingConfigurationName_10)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *DescribeScalingConfigurationsRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *DescribeScalingConfigurationsRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *DescribeScalingConfigurationsRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *DescribeScalingConfigurationsRequest) GetPath() string {
	return ""
}

func (r *DescribeScalingConfigurationsRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *DescribeScalingConfigurationsRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *DescribeScalingConfigurationsRequest) GetActionName() string {
	return "DescribeScalingConfigurations"
}
