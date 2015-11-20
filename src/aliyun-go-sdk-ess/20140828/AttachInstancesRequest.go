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

type AttachInstancesRequest struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	ScalingGroupId       string
	InstanceId_1         string
	InstanceId_2         string
	InstanceId_3         string
	InstanceId_4         string
	InstanceId_5         string
	InstanceId_6         string
	InstanceId_7         string
	InstanceId_8         string
	InstanceId_9         string
	InstanceId_10        string
	InstanceId_11        string
	InstanceId_12        string
	InstanceId_13        string
	InstanceId_14        string
	InstanceId_15        string
	InstanceId_16        string
	InstanceId_17        string
	InstanceId_18        string
	InstanceId_19        string
	InstanceId_20        string
	OwnerAccount         string
}

func (r *AttachInstancesRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 25)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupId", r.ScalingGroupId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId.1", r.InstanceId_1)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId.2", r.InstanceId_2)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId.3", r.InstanceId_3)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId.4", r.InstanceId_4)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId.5", r.InstanceId_5)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId.6", r.InstanceId_6)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId.7", r.InstanceId_7)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId.8", r.InstanceId_8)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId.9", r.InstanceId_9)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId.10", r.InstanceId_10)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId.11", r.InstanceId_11)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId.12", r.InstanceId_12)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId.13", r.InstanceId_13)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId.14", r.InstanceId_14)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId.15", r.InstanceId_15)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId.16", r.InstanceId_16)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId.17", r.InstanceId_17)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId.18", r.InstanceId_18)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId.19", r.InstanceId_19)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId.20", r.InstanceId_20)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *AttachInstancesRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *AttachInstancesRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *AttachInstancesRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *AttachInstancesRequest) GetPath() string {
	return ""
}

func (r *AttachInstancesRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *AttachInstancesRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *AttachInstancesRequest) GetActionName() string {
	return "AttachInstances"
}
