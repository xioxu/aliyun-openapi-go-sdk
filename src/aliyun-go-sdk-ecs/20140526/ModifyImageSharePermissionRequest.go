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

type ModifyImageSharePermissionRequest struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	ImageId              string
	AddAccount_1         string
	AddAccount_2         string
	AddAccount_3         string
	AddAccount_4         string
	AddAccount_5         string
	AddAccount_6         string
	AddAccount_7         string
	AddAccount_8         string
	AddAccount_9         string
	AddAccount_10        string
	RemoveAccount_1      string
	RemoveAccount_2      string
	RemoveAccount_3      string
	RemoveAccount_4      string
	RemoveAccount_5      string
	RemoveAccount_6      string
	RemoveAccount_7      string
	RemoveAccount_8      string
	RemoveAccount_9      string
	RemoveAccount_10     string
	OwnerAccount         string
}

func (r *ModifyImageSharePermissionRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 25)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ImageId", r.ImageId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "AddAccount.1", r.AddAccount_1)
	core.AddToMapIfNotDefaultValueStr(queryVals, "AddAccount.2", r.AddAccount_2)
	core.AddToMapIfNotDefaultValueStr(queryVals, "AddAccount.3", r.AddAccount_3)
	core.AddToMapIfNotDefaultValueStr(queryVals, "AddAccount.4", r.AddAccount_4)
	core.AddToMapIfNotDefaultValueStr(queryVals, "AddAccount.5", r.AddAccount_5)
	core.AddToMapIfNotDefaultValueStr(queryVals, "AddAccount.6", r.AddAccount_6)
	core.AddToMapIfNotDefaultValueStr(queryVals, "AddAccount.7", r.AddAccount_7)
	core.AddToMapIfNotDefaultValueStr(queryVals, "AddAccount.8", r.AddAccount_8)
	core.AddToMapIfNotDefaultValueStr(queryVals, "AddAccount.9", r.AddAccount_9)
	core.AddToMapIfNotDefaultValueStr(queryVals, "AddAccount.10", r.AddAccount_10)
	core.AddToMapIfNotDefaultValueStr(queryVals, "RemoveAccount.1", r.RemoveAccount_1)
	core.AddToMapIfNotDefaultValueStr(queryVals, "RemoveAccount.2", r.RemoveAccount_2)
	core.AddToMapIfNotDefaultValueStr(queryVals, "RemoveAccount.3", r.RemoveAccount_3)
	core.AddToMapIfNotDefaultValueStr(queryVals, "RemoveAccount.4", r.RemoveAccount_4)
	core.AddToMapIfNotDefaultValueStr(queryVals, "RemoveAccount.5", r.RemoveAccount_5)
	core.AddToMapIfNotDefaultValueStr(queryVals, "RemoveAccount.6", r.RemoveAccount_6)
	core.AddToMapIfNotDefaultValueStr(queryVals, "RemoveAccount.7", r.RemoveAccount_7)
	core.AddToMapIfNotDefaultValueStr(queryVals, "RemoveAccount.8", r.RemoveAccount_8)
	core.AddToMapIfNotDefaultValueStr(queryVals, "RemoveAccount.9", r.RemoveAccount_9)
	core.AddToMapIfNotDefaultValueStr(queryVals, "RemoveAccount.10", r.RemoveAccount_10)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *ModifyImageSharePermissionRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *ModifyImageSharePermissionRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *ModifyImageSharePermissionRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *ModifyImageSharePermissionRequest) GetPath() string {
	return ""
}

func (r *ModifyImageSharePermissionRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *ModifyImageSharePermissionRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *ModifyImageSharePermissionRequest) GetActionName() string {
	return "ModifyImageSharePermission"
}
