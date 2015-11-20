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

package rds

import "aliyun-go-sdk-core"

type RemoveTagsFromResourceRequest struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	ClientToken          string
	ProxyId              string
	DBInstanceId         string
	Tag_1_key            string
	Tag_2_key            string
	Tag_3_key            string
	Tag_4_key            string
	Tag_5_key            string
	Tag_1_value          string
	Tag_2_value          string
	Tag_3_value          string
	Tag_4_value          string
	Tag_5_value          string
	OwnerAccount         string
}

func (r *RemoveTagsFromResourceRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 17)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ClientToken", r.ClientToken)
	core.AddToMapIfNotDefaultValueStr(queryVals, "proxyId", r.ProxyId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DBInstanceId", r.DBInstanceId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Tag.1.key", r.Tag_1_key)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Tag.2.key", r.Tag_2_key)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Tag.3.key", r.Tag_3_key)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Tag.4.key", r.Tag_4_key)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Tag.5.key", r.Tag_5_key)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Tag.1.value", r.Tag_1_value)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Tag.2.value", r.Tag_2_value)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Tag.3.value", r.Tag_3_value)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Tag.4.value", r.Tag_4_value)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Tag.5.value", r.Tag_5_value)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *RemoveTagsFromResourceRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *RemoveTagsFromResourceRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *RemoveTagsFromResourceRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *RemoveTagsFromResourceRequest) GetPath() string {
	return ""
}

func (r *RemoveTagsFromResourceRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *RemoveTagsFromResourceRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *RemoveTagsFromResourceRequest) GetActionName() string {
	return "RemoveTagsFromResource"
}
