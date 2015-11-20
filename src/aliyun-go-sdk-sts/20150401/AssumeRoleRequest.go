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

package sts

import "aliyun-go-sdk-core"

type AssumeRoleRequest struct {
	DurationSeconds int64
	Policy          string
	RoleArn         string
	RoleSessionName string
}

func (r *AssumeRoleRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 4)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "DurationSeconds", r.DurationSeconds)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Policy", r.Policy)
	core.AddToMapIfNotDefaultValueStr(queryVals, "RoleArn", r.RoleArn)
	core.AddToMapIfNotDefaultValueStr(queryVals, "RoleSessionName", r.RoleSessionName)

	return queryVals
}

func (r *AssumeRoleRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *AssumeRoleRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *AssumeRoleRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *AssumeRoleRequest) GetPath() string {
	return ""
}

func (r *AssumeRoleRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *AssumeRoleRequest) GetProtocol() string {
	return "HTTPS"
}

func (r *AssumeRoleRequest) GetActionName() string {
	return "AssumeRole"
}
