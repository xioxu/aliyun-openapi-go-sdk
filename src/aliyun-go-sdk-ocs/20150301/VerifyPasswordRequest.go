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

package ocs

import "aliyun-go-sdk-core"

type VerifyPasswordRequest struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	OwnerAccount         string
	InstanceId           string
	Password             string
}

func (r *VerifyPasswordRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 6)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId", r.InstanceId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Password", r.Password)

	return queryVals
}

func (r *VerifyPasswordRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *VerifyPasswordRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *VerifyPasswordRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *VerifyPasswordRequest) GetPath() string {
	return ""
}

func (r *VerifyPasswordRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *VerifyPasswordRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *VerifyPasswordRequest) GetActionName() string {
	return "VerifyPassword"
}
