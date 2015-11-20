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

type DescribeImagesRequest struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	Status               string
	ImageId              string
	ShowExpired          bool
	SnapshotId           string
	ImageName            string
	ImageOwnerAlias      string
	PageNumber           int32
	PageSize             int32
	OwnerAccount         string
	Filter_1_Key         string
	Filter_2_Key         string
	Filter_1_Value       string
	Filter_2_Value       string
	Usage                string
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

func (r *DescribeImagesRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 27)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Status", r.Status)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ImageId", r.ImageId)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "ShowExpired", r.ShowExpired)
	core.AddToMapIfNotDefaultValueStr(queryVals, "SnapshotId", r.SnapshotId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ImageName", r.ImageName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ImageOwnerAlias", r.ImageOwnerAlias)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageNumber", r.PageNumber)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageSize", r.PageSize)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Filter.1.Key", r.Filter_1_Key)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Filter.2.Key", r.Filter_2_Key)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Filter.1.Value", r.Filter_1_Value)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Filter.2.Value", r.Filter_2_Value)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Usage", r.Usage)
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

func (r *DescribeImagesRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *DescribeImagesRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *DescribeImagesRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *DescribeImagesRequest) GetPath() string {
	return ""
}

func (r *DescribeImagesRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *DescribeImagesRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *DescribeImagesRequest) GetActionName() string {
	return "DescribeImages"
}
