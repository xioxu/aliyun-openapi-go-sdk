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

type DescribeDisksRequest struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	ZoneId               string
	DiskIds              string
	InstanceId           string
	DiskType             string
	Category             string
	Status               string
	SnapshotId           string
	Portable             bool
	DeleteWithInstance   bool
	DeleteAutoSnapshot   bool
	PageNumber           int32
	PageSize             int32
	OwnerAccount         string
	DiskName             string
	EnableAutoSnapshot   bool
	DiskChargeType       string
	LockReason           string
	Filter_1_Key         string
	Filter_2_Key         string
	Filter_1_Value       string
	Filter_2_Value       string
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

func (r *DescribeDisksRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 34)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ZoneId", r.ZoneId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DiskIds", r.DiskIds)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId", r.InstanceId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DiskType", r.DiskType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Category", r.Category)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Status", r.Status)
	core.AddToMapIfNotDefaultValueStr(queryVals, "SnapshotId", r.SnapshotId)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "Portable", r.Portable)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "DeleteWithInstance", r.DeleteWithInstance)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "DeleteAutoSnapshot", r.DeleteAutoSnapshot)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageNumber", r.PageNumber)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageSize", r.PageSize)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DiskName", r.DiskName)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "EnableAutoSnapshot", r.EnableAutoSnapshot)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DiskChargeType", r.DiskChargeType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "LockReason", r.LockReason)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Filter.1.Key", r.Filter_1_Key)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Filter.2.Key", r.Filter_2_Key)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Filter.1.Value", r.Filter_1_Value)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Filter.2.Value", r.Filter_2_Value)
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

func (r *DescribeDisksRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *DescribeDisksRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *DescribeDisksRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *DescribeDisksRequest) GetPath() string {
	return ""
}

func (r *DescribeDisksRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *DescribeDisksRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *DescribeDisksRequest) GetActionName() string {
	return "DescribeDisks"
}
