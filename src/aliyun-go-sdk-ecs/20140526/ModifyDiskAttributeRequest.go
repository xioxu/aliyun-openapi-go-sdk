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

type ModifyDiskAttributeRequest struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	DiskId               string
	DiskName             string
	Description          string
	DeleteWithInstance   bool
	DeleteAutoSnapshot   bool
	EnableAutoSnapshot   bool
	OwnerAccount         string
}

func (r *ModifyDiskAttributeRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 10)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DiskId", r.DiskId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DiskName", r.DiskName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Description", r.Description)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "DeleteWithInstance", r.DeleteWithInstance)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "DeleteAutoSnapshot", r.DeleteAutoSnapshot)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "EnableAutoSnapshot", r.EnableAutoSnapshot)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *ModifyDiskAttributeRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *ModifyDiskAttributeRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *ModifyDiskAttributeRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *ModifyDiskAttributeRequest) GetPath() string {
	return ""
}

func (r *ModifyDiskAttributeRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *ModifyDiskAttributeRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *ModifyDiskAttributeRequest) GetActionName() string {
	return "ModifyDiskAttribute"
}
