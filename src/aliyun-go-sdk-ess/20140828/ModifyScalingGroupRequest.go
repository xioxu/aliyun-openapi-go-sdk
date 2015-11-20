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

type ModifyScalingGroupRequest struct {
	OwnerId                      int64
	ResourceOwnerAccount         string
	ResourceOwnerId              int64
	ScalingGroupId               string
	ScalingGroupName             string
	MinSize                      int32
	MaxSize                      int32
	DefaultCooldown              int32
	RemovalPolicy_1              string
	RemovalPolicy_2              string
	ActiveScalingConfigurationId string
	OwnerAccount                 string
}

func (r *ModifyScalingGroupRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 12)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupId", r.ScalingGroupId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupName", r.ScalingGroupName)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "MinSize", r.MinSize)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "MaxSize", r.MaxSize)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "DefaultCooldown", r.DefaultCooldown)
	core.AddToMapIfNotDefaultValueStr(queryVals, "RemovalPolicy.1", r.RemovalPolicy_1)
	core.AddToMapIfNotDefaultValueStr(queryVals, "RemovalPolicy.2", r.RemovalPolicy_2)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ActiveScalingConfigurationId", r.ActiveScalingConfigurationId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *ModifyScalingGroupRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *ModifyScalingGroupRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *ModifyScalingGroupRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *ModifyScalingGroupRequest) GetPath() string {
	return ""
}

func (r *ModifyScalingGroupRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *ModifyScalingGroupRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *ModifyScalingGroupRequest) GetActionName() string {
	return "ModifyScalingGroup"
}
