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

package yundun

import "aliyun-go-sdk-core"

type ListInstanceInfosRequest struct {
	JstOwnerId   int64
	PageNumber   int32
	PageSize     int32
	Region       string
	EventType    string
	InstanceName string
	InstanceType string
	InstanceIds  string
}

func (r *ListInstanceInfosRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 8)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "JstOwnerId", r.JstOwnerId)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageNumber", r.PageNumber)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageSize", r.PageSize)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Region", r.Region)
	core.AddToMapIfNotDefaultValueStr(queryVals, "EventType", r.EventType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceName", r.InstanceName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceType", r.InstanceType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceIds", r.InstanceIds)

	return queryVals
}

func (r *ListInstanceInfosRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *ListInstanceInfosRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *ListInstanceInfosRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *ListInstanceInfosRequest) GetPath() string {
	return ""
}

func (r *ListInstanceInfosRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *ListInstanceInfosRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *ListInstanceInfosRequest) GetActionName() string {
	return "ListInstanceInfos"
}
