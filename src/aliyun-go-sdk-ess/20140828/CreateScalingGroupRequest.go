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

type CreateScalingGroupRequest struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	ScalingGroupName     string
	MinSize              int32
	MaxSize              int32
	DefaultCooldown      int32
	LoadBalancerId       string
	DBInstanceId_1       string
	DBInstanceId_2       string
	DBInstanceId_3       string
	RemovalPolicy_1      string
	RemovalPolicy_2      string
	OwnerAccount         string
}

func (r *CreateScalingGroupRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 14)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupName", r.ScalingGroupName)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "MinSize", r.MinSize)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "MaxSize", r.MaxSize)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "DefaultCooldown", r.DefaultCooldown)
	core.AddToMapIfNotDefaultValueStr(queryVals, "LoadBalancerId", r.LoadBalancerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DBInstanceId.1", r.DBInstanceId_1)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DBInstanceId.2", r.DBInstanceId_2)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DBInstanceId.3", r.DBInstanceId_3)
	core.AddToMapIfNotDefaultValueStr(queryVals, "RemovalPolicy.1", r.RemovalPolicy_1)
	core.AddToMapIfNotDefaultValueStr(queryVals, "RemovalPolicy.2", r.RemovalPolicy_2)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *CreateScalingGroupRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *CreateScalingGroupRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *CreateScalingGroupRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *CreateScalingGroupRequest) GetPath() string {
	return ""
}

func (r *CreateScalingGroupRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *CreateScalingGroupRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *CreateScalingGroupRequest) GetActionName() string {
	return "CreateScalingGroup"
}
