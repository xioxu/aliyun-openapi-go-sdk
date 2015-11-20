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

type CreateInstanceRequest struct {
	OwnerId                       int64
	ResourceOwnerAccount          string
	ResourceOwnerId               int64
	ImageId                       string
	InstanceType                  string
	SecurityGroupId               string
	InstanceName                  string
	InternetChargeType            string
	InternetMaxBandwidthIn        int32
	InternetMaxBandwidthOut       int32
	HostName                      string
	Password                      string
	ZoneId                        string
	ClusterId                     string
	ClientToken                   string
	VlanId                        string
	InnerIpAddress                string
	SystemDisk_Category           string
	SystemDisk_DiskName           string
	SystemDisk_Description        string
	DataDisk_1_Size               int32
	DataDisk_1_Category           string
	DataDisk_1_SnapshotId         string
	DataDisk_1_DiskName           string
	DataDisk_1_Description        string
	DataDisk_1_Device             string
	DataDisk_1_DeleteWithInstance bool
	DataDisk_2_Size               int32
	DataDisk_2_Category           string
	DataDisk_2_SnapshotId         string
	DataDisk_2_DiskName           string
	DataDisk_2_Description        string
	DataDisk_2_Device             string
	DataDisk_2_DeleteWithInstance bool
	DataDisk_3_Size               int32
	DataDisk_3_Category           string
	DataDisk_3_SnapshotId         string
	DataDisk_3_DiskName           string
	DataDisk_3_Description        string
	DataDisk_3_Device             string
	DataDisk_3_DeleteWithInstance bool
	DataDisk_4_Size               int32
	DataDisk_4_Category           string
	DataDisk_4_SnapshotId         string
	DataDisk_4_DiskName           string
	DataDisk_4_Description        string
	DataDisk_4_Device             string
	DataDisk_4_DeleteWithInstance bool
	NodeControllerId              string
	Description                   string
	VSwitchId                     string
	PrivateIpAddress              string
	IoOptimized                   string
	OwnerAccount                  string
	UseAdditionalService          bool
	InstanceChargeType            string
	Period                        int32
}

func (r *CreateInstanceRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 57)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ImageId", r.ImageId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceType", r.InstanceType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "SecurityGroupId", r.SecurityGroupId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceName", r.InstanceName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InternetChargeType", r.InternetChargeType)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "InternetMaxBandwidthIn", r.InternetMaxBandwidthIn)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "InternetMaxBandwidthOut", r.InternetMaxBandwidthOut)
	core.AddToMapIfNotDefaultValueStr(queryVals, "HostName", r.HostName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Password", r.Password)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ZoneId", r.ZoneId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ClusterId", r.ClusterId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ClientToken", r.ClientToken)
	core.AddToMapIfNotDefaultValueStr(queryVals, "VlanId", r.VlanId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InnerIpAddress", r.InnerIpAddress)
	core.AddToMapIfNotDefaultValueStr(queryVals, "SystemDisk.Category", r.SystemDisk_Category)
	core.AddToMapIfNotDefaultValueStr(queryVals, "SystemDisk.DiskName", r.SystemDisk_DiskName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "SystemDisk.Description", r.SystemDisk_Description)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "DataDisk.1.Size", r.DataDisk_1_Size)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.1.Category", r.DataDisk_1_Category)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.1.SnapshotId", r.DataDisk_1_SnapshotId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.1.DiskName", r.DataDisk_1_DiskName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.1.Description", r.DataDisk_1_Description)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.1.Device", r.DataDisk_1_Device)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "DataDisk.1.DeleteWithInstance", r.DataDisk_1_DeleteWithInstance)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "DataDisk.2.Size", r.DataDisk_2_Size)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.2.Category", r.DataDisk_2_Category)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.2.SnapshotId", r.DataDisk_2_SnapshotId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.2.DiskName", r.DataDisk_2_DiskName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.2.Description", r.DataDisk_2_Description)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.2.Device", r.DataDisk_2_Device)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "DataDisk.2.DeleteWithInstance", r.DataDisk_2_DeleteWithInstance)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "DataDisk.3.Size", r.DataDisk_3_Size)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.3.Category", r.DataDisk_3_Category)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.3.SnapshotId", r.DataDisk_3_SnapshotId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.3.DiskName", r.DataDisk_3_DiskName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.3.Description", r.DataDisk_3_Description)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.3.Device", r.DataDisk_3_Device)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "DataDisk.3.DeleteWithInstance", r.DataDisk_3_DeleteWithInstance)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "DataDisk.4.Size", r.DataDisk_4_Size)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.4.Category", r.DataDisk_4_Category)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.4.SnapshotId", r.DataDisk_4_SnapshotId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.4.DiskName", r.DataDisk_4_DiskName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.4.Description", r.DataDisk_4_Description)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataDisk.4.Device", r.DataDisk_4_Device)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "DataDisk.4.DeleteWithInstance", r.DataDisk_4_DeleteWithInstance)
	core.AddToMapIfNotDefaultValueStr(queryVals, "NodeControllerId", r.NodeControllerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Description", r.Description)
	core.AddToMapIfNotDefaultValueStr(queryVals, "VSwitchId", r.VSwitchId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "PrivateIpAddress", r.PrivateIpAddress)
	core.AddToMapIfNotDefaultValueStr(queryVals, "IoOptimized", r.IoOptimized)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "UseAdditionalService", r.UseAdditionalService)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceChargeType", r.InstanceChargeType)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "Period", r.Period)

	return queryVals
}

func (r *CreateInstanceRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *CreateInstanceRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *CreateInstanceRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *CreateInstanceRequest) GetPath() string {
	return ""
}

func (r *CreateInstanceRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *CreateInstanceRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *CreateInstanceRequest) GetActionName() string {
	return "CreateInstance"
}
