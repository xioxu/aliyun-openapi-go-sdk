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

type AddTagsRequest struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	ResourceType         string
	ResourceId           string
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

func (r *AddTagsRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 15)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceType", r.ResourceType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceId", r.ResourceId)
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

func (r *AddTagsRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *AddTagsRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *AddTagsRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *AddTagsRequest) GetPath() string {
	return ""
}

func (r *AddTagsRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *AddTagsRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *AddTagsRequest) GetActionName() string {
	return "AddTags"
}
