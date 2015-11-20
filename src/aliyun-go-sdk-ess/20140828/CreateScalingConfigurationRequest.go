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

type CreateScalingConfigurationRequest struct {
	OwnerId                       int64
	ResourceOwnerAccount          string
	ResourceOwnerId               int64
	ScalingGroupId                string
	ImageId                       string
	InstanceType                  string
	SecurityGroupId               string
	InternetChargeType            string
	InternetMaxBandwidthIn        int32
	InternetMaxBandwidthOut       int32
	SystemDisk_Category           string
	ScalingConfigurationName      string
	DataDisk_1_Size               int32
	DataDisk_2_Size               int32
	DataDisk_3_Size               int32
	DataDisk_4_Size               int32
	DataDisk_1_Category           string
	DataDisk_2_Category           string
	DataDisk_3_Category           string
	DataDisk_4_Category           string
	DataDisk_1_SnapshotId         string
	DataDisk_2_SnapshotId         string
	DataDisk_3_SnapshotId         string
	DataDisk_4_SnapshotId         string
	DataDisk_1_Device             string
	DataDisk_2_Device             string
	DataDisk_3_Device             string
	DataDisk_4_Device             string
	DataDisk_1_DeleteWithInstance string
	DataDisk_2_DeleteWithInstance string
	DataDisk_3_DeleteWithInstance string
	DataDisk_4_DeleteWithInstance string
	OwnerAccount                  string
}

func (r *CreateScalingConfigurationRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 33)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingGroupId", r.ScalingGroupId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ImageId", r.ImageId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceType", r.InstanceType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "SecurityGroupId", r.SecurityGroupId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InternetChargeType", r.InternetChargeType)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "InternetMaxBandwidthIn", r.InternetMaxBandwidthIn)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "InternetMaxBandwidthOut", r.InternetMaxBandwidthOut)
	core.AddToMapIfNotDefaultValueStr(queryVals, "SystemDisk.Category", r.SystemDisk_Category)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScalingConfigurationName", r.ScalingConfigurationName)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "DataDisk.1.Size", r.DataDisk_1_Size)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "DataDisk.2.Size", r.DataDisk_2_Size)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "DataDisk.3.Size", r.DataDisk_3_Size)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "DataDisk.4.Size", r.DataDisk_4_Size)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.1.Category", r.DataDisk_1_Category)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.2.Category", r.DataDisk_2_Category)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.3.Category", r.DataDisk_3_Category)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.4.Category", r.DataDisk_4_Category)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.1.SnapshotId", r.DataDisk_1_SnapshotId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.2.SnapshotId", r.DataDisk_2_SnapshotId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.3.SnapshotId", r.DataDisk_3_SnapshotId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.4.SnapshotId", r.DataDisk_4_SnapshotId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.1.Device", r.DataDisk_1_Device)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.2.Device", r.DataDisk_2_Device)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.3.Device", r.DataDisk_3_Device)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.4.Device", r.DataDisk_4_Device)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.1.DeleteWithInstance", r.DataDisk_1_DeleteWithInstance)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.2.DeleteWithInstance", r.DataDisk_2_DeleteWithInstance)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.3.DeleteWithInstance", r.DataDisk_3_DeleteWithInstance)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.4.DeleteWithInstance", r.DataDisk_4_DeleteWithInstance)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *CreateScalingConfigurationRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *CreateScalingConfigurationRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *CreateScalingConfigurationRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *CreateScalingConfigurationRequest) GetPath() string {
	return ""
}

func (r *CreateScalingConfigurationRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *CreateScalingConfigurationRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *CreateScalingConfigurationRequest) GetActionName() string {
	return "CreateScalingConfiguration"
}
