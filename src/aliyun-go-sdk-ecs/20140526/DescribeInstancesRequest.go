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

type DescribeInstancesRequest struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	VpcId                string
	VSwitchId            string
	ZoneId               string
	InstanceNetworkType  string
	SecurityGroupId      string
	InstanceIds          string
	PageNumber           int32
	PageSize             int32
	InnerIpAddresses     string
	PrivateIpAddresses   string
	PublicIpAddresses    string
	OwnerAccount         string
	InstanceChargeType   string
	InternetChargeType   string
	InstanceName         string
	ImageId              string
	Status               string
	LockReason           string
	Filter_1_Key         string
	Filter_2_Key         string
	Filter_3_Key         string
	Filter_4_Key         string
	Filter_1_Value       string
	Filter_2_Value       string
	Filter_3_Value       string
	Filter_4_Value       string
	DeviceAvailable      bool
	IoOptimized          bool
	Tag_1_Key            string
	Tag_2_Key            string
	Tag_3_Key            string
	Tag_4_Key            string
	Tag_5_Key            string
	Tag_1_Value          string
	Tag_2_Value          string
	Tag_3_Value          string
	Tag_4_Value          string
	Tag_5_Value          string
}

func (r *DescribeInstancesRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 41)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "VpcId", r.VpcId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "VSwitchId", r.VSwitchId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ZoneId", r.ZoneId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceNetworkType", r.InstanceNetworkType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "SecurityGroupId", r.SecurityGroupId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceIds", r.InstanceIds)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageNumber", r.PageNumber)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageSize", r.PageSize)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InnerIpAddresses", r.InnerIpAddresses)
	core.AddToMapIfNotDefaultValueStr(queryVals, "PrivateIpAddresses", r.PrivateIpAddresses)
	core.AddToMapIfNotDefaultValueStr(queryVals, "PublicIpAddresses", r.PublicIpAddresses)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceChargeType", r.InstanceChargeType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InternetChargeType", r.InternetChargeType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceName", r.InstanceName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ImageId", r.ImageId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Status", r.Status)
	core.AddToMapIfNotDefaultValueStr(queryVals, "LockReason", r.LockReason)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Filter.1.Key", r.Filter_1_Key)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Filter.2.Key", r.Filter_2_Key)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Filter.3.Key", r.Filter_3_Key)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Filter.4.Key", r.Filter_4_Key)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Filter.1.Value", r.Filter_1_Value)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Filter.2.Value", r.Filter_2_Value)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Filter.3.Value", r.Filter_3_Value)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Filter.4.Value", r.Filter_4_Value)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "DeviceAvailable", r.DeviceAvailable)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "IoOptimized", r.IoOptimized)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Tag.1.Key", r.Tag_1_Key)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Tag.2.Key", r.Tag_2_Key)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Tag.3.Key", r.Tag_3_Key)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Tag.4.Key", r.Tag_4_Key)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Tag.5.Key", r.Tag_5_Key)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Tag.1.Value", r.Tag_1_Value)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Tag.2.Value", r.Tag_2_Value)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Tag.3.Value", r.Tag_3_Value)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Tag.4.Value", r.Tag_4_Value)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Tag.5.Value", r.Tag_5_Value)

	return queryVals
}

func (r *DescribeInstancesRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *DescribeInstancesRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *DescribeInstancesRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *DescribeInstancesRequest) GetPath() string {
	return ""
}

func (r *DescribeInstancesRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *DescribeInstancesRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *DescribeInstancesRequest) GetActionName() string {
	return "DescribeInstances"
}
