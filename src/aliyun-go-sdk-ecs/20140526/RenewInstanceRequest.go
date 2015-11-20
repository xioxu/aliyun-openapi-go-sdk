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

type RenewInstanceRequest struct {
	OwnerId                     int64
	ResourceOwnerAccount        string
	ResourceOwnerId             int64
	OwnerAccount                string
	InstanceId                  string
	InstanceType                string
	InternetMaxBandwidthOut     int32
	InternetChargeType          string
	Period                      int32
	RebootTime                  string
	CovertDiskPortable_1_DiskId string
	CovertDiskPortable_2_DiskId string
	CovertDiskPortable_3_DiskId string
	CovertDiskPortable_4_DiskId string
}

func (r *RenewInstanceRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 14)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId", r.InstanceId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceType", r.InstanceType)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "InternetMaxBandwidthOut", r.InternetMaxBandwidthOut)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InternetChargeType", r.InternetChargeType)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "Period", r.Period)
	core.AddToMapIfNotDefaultValueStr(queryVals, "RebootTime", r.RebootTime)
	core.AddToMapIfNotDefaultValueStr(queryVals, "CovertDiskPortable.1.DiskId", r.CovertDiskPortable_1_DiskId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "CovertDiskPortable.2.DiskId", r.CovertDiskPortable_2_DiskId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "CovertDiskPortable.3.DiskId", r.CovertDiskPortable_3_DiskId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "CovertDiskPortable.4.DiskId", r.CovertDiskPortable_4_DiskId)

	return queryVals
}

func (r *RenewInstanceRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *RenewInstanceRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *RenewInstanceRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *RenewInstanceRequest) GetPath() string {
	return ""
}

func (r *RenewInstanceRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *RenewInstanceRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *RenewInstanceRequest) GetActionName() string {
	return "RenewInstance"
}
