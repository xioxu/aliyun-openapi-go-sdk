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

type CheckAutoSnapshotPolicyRequest struct {
	OwnerId                           int64
	ResourceOwnerAccount              string
	ResourceOwnerId                   int64
	OwnerAccount                      string
	SystemDiskPolicyEnabled           bool
	SystemDiskPolicyTimePeriod        int32
	SystemDiskPolicyRetentionDays     int32
	SystemDiskPolicyRetentionLastWeek bool
	DataDiskPolicyEnabled             bool
	DataDiskPolicyTimePeriod          int32
	DataDiskPolicyRetentionDays       int32
	DataDiskPolicyRetentionLastWeek   bool
}

func (r *CheckAutoSnapshotPolicyRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 12)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "SystemDiskPolicyEnabled", r.SystemDiskPolicyEnabled)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "SystemDiskPolicyTimePeriod", r.SystemDiskPolicyTimePeriod)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "SystemDiskPolicyRetentionDays", r.SystemDiskPolicyRetentionDays)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "SystemDiskPolicyRetentionLastWeek", r.SystemDiskPolicyRetentionLastWeek)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "DataDiskPolicyEnabled", r.DataDiskPolicyEnabled)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "DataDiskPolicyTimePeriod", r.DataDiskPolicyTimePeriod)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "DataDiskPolicyRetentionDays", r.DataDiskPolicyRetentionDays)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "DataDiskPolicyRetentionLastWeek", r.DataDiskPolicyRetentionLastWeek)

	return queryVals
}

func (r *CheckAutoSnapshotPolicyRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *CheckAutoSnapshotPolicyRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *CheckAutoSnapshotPolicyRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *CheckAutoSnapshotPolicyRequest) GetPath() string {
	return ""
}

func (r *CheckAutoSnapshotPolicyRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *CheckAutoSnapshotPolicyRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *CheckAutoSnapshotPolicyRequest) GetActionName() string {
	return "CheckAutoSnapshotPolicy"
}
