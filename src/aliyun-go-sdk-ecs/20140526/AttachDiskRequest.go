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

type AttachDiskRequest struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	InstanceId           string
	DiskId               string
	Device               string
	DeleteWithInstance   bool
	OwnerAccount         string
}

func (r *AttachDiskRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 8)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId", r.InstanceId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DiskId", r.DiskId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Device", r.Device)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "DeleteWithInstance", r.DeleteWithInstance)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *AttachDiskRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *AttachDiskRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *AttachDiskRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *AttachDiskRequest) GetPath() string {
	return ""
}

func (r *AttachDiskRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *AttachDiskRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *AttachDiskRequest) GetActionName() string {
	return "AttachDisk"
}
