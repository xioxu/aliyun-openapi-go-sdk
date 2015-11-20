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

package ots

import "aliyun-go-sdk-core"

type InsertInstanceRequest struct {
	InstanceName  string
	ClusterName   string
	WriteCapacity int32
	ReadCapacity  int32
	EntityQuota   int32
	Description   string
}

func (r *InsertInstanceRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 6)

	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceName", r.InstanceName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ClusterName", r.ClusterName)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "WriteCapacity", r.WriteCapacity)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "ReadCapacity", r.ReadCapacity)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "EntityQuota", r.EntityQuota)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Description", r.Description)

	return queryVals
}

func (r *InsertInstanceRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *InsertInstanceRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *InsertInstanceRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *InsertInstanceRequest) GetPath() string {
	return ""
}

func (r *InsertInstanceRequest) GetRequestMethod() string {
	return "POST"
}

func (r *InsertInstanceRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *InsertInstanceRequest) GetActionName() string {
	return "InsertInstance"
}
