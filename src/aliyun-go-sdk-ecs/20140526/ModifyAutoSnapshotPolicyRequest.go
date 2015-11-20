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

type ModifyAutoSnapshotPolicyRequest struct {
	OwnerId                           int64
	ResourceOwnerAccount              string
	ResourceOwnerId                   int64
	SystemDiskPolicyEnabled           bool
	SystemDiskPolicyTimePeriod        int32
	SystemDiskPolicyRetentionDays     int32
	SystemDiskPolicyRetentionLastWeek bool
	DataDiskPolicyEnabled             bool
	DataDiskPolicyTimePeriod          int32
	DataDiskPolicyRetentionDays       int32
	DataDiskPolicyRetentionLastWeek   bool
	OwnerAccount                      string
}

func (r *ModifyAutoSnapshotPolicyRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 12)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "SystemDiskPolicyEnabled", r.SystemDiskPolicyEnabled)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "SystemDiskPolicyTimePeriod", r.SystemDiskPolicyTimePeriod)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "SystemDiskPolicyRetentionDays", r.SystemDiskPolicyRetentionDays)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "SystemDiskPolicyRetentionLastWeek", r.SystemDiskPolicyRetentionLastWeek)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "DataDiskPolicyEnabled", r.DataDiskPolicyEnabled)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "DataDiskPolicyTimePeriod", r.DataDiskPolicyTimePeriod)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "DataDiskPolicyRetentionDays", r.DataDiskPolicyRetentionDays)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "DataDiskPolicyRetentionLastWeek", r.DataDiskPolicyRetentionLastWeek)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *ModifyAutoSnapshotPolicyRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *ModifyAutoSnapshotPolicyRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *ModifyAutoSnapshotPolicyRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *ModifyAutoSnapshotPolicyRequest) GetPath() string {
	return ""
}

func (r *ModifyAutoSnapshotPolicyRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *ModifyAutoSnapshotPolicyRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *ModifyAutoSnapshotPolicyRequest) GetActionName() string {
	return "ModifyAutoSnapshotPolicy"
}
