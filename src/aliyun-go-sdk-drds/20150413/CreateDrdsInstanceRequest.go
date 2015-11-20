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

package drds

import "aliyun-go-sdk-core"

type CreateDrdsInstanceRequest struct {
	Description   string
	ZoneId        string
	Type          string
	Quantity      int32
	Specification string
	PayType       string
	VpcId         string
	VswitchId     string
}

func (r *CreateDrdsInstanceRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 8)

	core.AddToMapIfNotDefaultValueStr(queryVals, "Description", r.Description)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ZoneId", r.ZoneId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Type", r.Type)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "Quantity", r.Quantity)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Specification", r.Specification)
	core.AddToMapIfNotDefaultValueStr(queryVals, "PayType", r.PayType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "VpcId", r.VpcId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "VswitchId", r.VswitchId)

	return queryVals
}

func (r *CreateDrdsInstanceRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *CreateDrdsInstanceRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *CreateDrdsInstanceRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *CreateDrdsInstanceRequest) GetPath() string {
	return ""
}

func (r *CreateDrdsInstanceRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *CreateDrdsInstanceRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *CreateDrdsInstanceRequest) GetActionName() string {
	return "CreateDrdsInstance"
}
