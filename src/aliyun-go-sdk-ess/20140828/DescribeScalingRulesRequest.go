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

type DescribeScalingRulesRequest struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	PageNumber           int32
	PageSize             int32
	ScalingGroupId       string
	ScalingRuleId_1      string
	ScalingRuleId_2      string
	ScalingRuleId_3      string
	ScalingRuleId_4      string
	ScalingRuleId_5      string
	ScalingRuleId_6      string
	ScalingRuleId_7      string
	ScalingRuleId_8      string
	ScalingRuleId_9      string
	ScalingRuleId_10     string
	ScalingRuleName_1    string
	ScalingRuleName_2    string
	ScalingRuleName_3    string
	ScalingRuleName_4    string
	ScalingRuleName_5    string
	ScalingRuleName_6    string
	ScalingRuleName_7    string
	ScalingRuleName_8    string
	ScalingRuleName_9    string
	ScalingRuleName_10   string
	ScalingRuleAri_1     string
	ScalingRuleAri_2     string
	ScalingRuleAri_3     string
	ScalingRuleAri_4     string
	ScalingRuleAri_5     string
	ScalingRuleAri_6     string
	ScalingRuleAri_7     string
	ScalingRuleAri_8     string
	ScalingRuleAri_9     string
	ScalingRuleAri_10    string
	OwnerAccount         string
}

func (r *DescribeScalingRulesRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 37)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageNumber", r.PageNumber)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageSize", r.PageSize)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupId", r.ScalingGroupId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleId.1", r.ScalingRuleId_1)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleId.2", r.ScalingRuleId_2)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleId.3", r.ScalingRuleId_3)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleId.4", r.ScalingRuleId_4)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleId.5", r.ScalingRuleId_5)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleId.6", r.ScalingRuleId_6)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleId.7", r.ScalingRuleId_7)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleId.8", r.ScalingRuleId_8)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleId.9", r.ScalingRuleId_9)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleId.10", r.ScalingRuleId_10)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleName.1", r.ScalingRuleName_1)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleName.2", r.ScalingRuleName_2)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleName.3", r.ScalingRuleName_3)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleName.4", r.ScalingRuleName_4)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleName.5", r.ScalingRuleName_5)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleName.6", r.ScalingRuleName_6)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleName.7", r.ScalingRuleName_7)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleName.8", r.ScalingRuleName_8)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleName.9", r.ScalingRuleName_9)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleName.10", r.ScalingRuleName_10)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleAri.1", r.ScalingRuleAri_1)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleAri.2", r.ScalingRuleAri_2)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleAri.3", r.ScalingRuleAri_3)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleAri.4", r.ScalingRuleAri_4)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleAri.5", r.ScalingRuleAri_5)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleAri.6", r.ScalingRuleAri_6)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleAri.7", r.ScalingRuleAri_7)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleAri.8", r.ScalingRuleAri_8)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleAri.9", r.ScalingRuleAri_9)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleAri.10", r.ScalingRuleAri_10)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *DescribeScalingRulesRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *DescribeScalingRulesRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *DescribeScalingRulesRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *DescribeScalingRulesRequest) GetPath() string {
	return ""
}

func (r *DescribeScalingRulesRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *DescribeScalingRulesRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *DescribeScalingRulesRequest) GetActionName() string {
	return "DescribeScalingRules"
}
