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

type CreateVSwitchRequest struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	ZoneId               string
	CidrBlock            string
	VpcId                string
	VSwitchName          string
	Description          string
	ClientToken          string
	OwnerAccount         string
}

func (r *CreateVSwitchRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 10)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ZoneId", r.ZoneId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "CidrBlock", r.CidrBlock)
	core.AddToMapIfNotDefaultValueStr(queryVals, "VpcId", r.VpcId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "VSwitchName", r.VSwitchName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Description", r.Description)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ClientToken", r.ClientToken)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *CreateVSwitchRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *CreateVSwitchRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *CreateVSwitchRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *CreateVSwitchRequest) GetPath() string {
	return ""
}

func (r *CreateVSwitchRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *CreateVSwitchRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *CreateVSwitchRequest) GetActionName() string {
	return "CreateVSwitch"
}
