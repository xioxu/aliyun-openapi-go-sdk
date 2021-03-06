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

type ModifyScalingRuleRequest struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	ScalingRuleId        string
	ScalingRuleName      string
	Cooldown             int32
	AdjustmentType       string
	AdjustmentValue      int32
	OwnerAccount         string
}

func (r *ModifyScalingRuleRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 9)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleId", r.ScalingRuleId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingRuleName", r.ScalingRuleName)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "Cooldown", r.Cooldown)
	core.AddToMapIfNotDefaultValueStr(queryVals, "AdjustmentType", r.AdjustmentType)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "AdjustmentValue", r.AdjustmentValue)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *ModifyScalingRuleRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *ModifyScalingRuleRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *ModifyScalingRuleRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *ModifyScalingRuleRequest) GetPath() string {
	return ""
}

func (r *ModifyScalingRuleRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *ModifyScalingRuleRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *ModifyScalingRuleRequest) GetActionName() string {
	return "ModifyScalingRule"
}
