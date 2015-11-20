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

type ModifyImageShareGroupPermissionRequest struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	ImageId              string
	AddGroup_1           string
	RemoveGroup_1        string
	OwnerAccount         string
}

func (r *ModifyImageShareGroupPermissionRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 7)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ImageId", r.ImageId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "AddGroup.1", r.AddGroup_1)
	core.AddToMapIfNotDefaultValueStr(queryVals, "RemoveGroup.1", r.RemoveGroup_1)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *ModifyImageShareGroupPermissionRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *ModifyImageShareGroupPermissionRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *ModifyImageShareGroupPermissionRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *ModifyImageShareGroupPermissionRequest) GetPath() string {
	return ""
}

func (r *ModifyImageShareGroupPermissionRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *ModifyImageShareGroupPermissionRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *ModifyImageShareGroupPermissionRequest) GetActionName() string {
	return "ModifyImageShareGroupPermission"
}
