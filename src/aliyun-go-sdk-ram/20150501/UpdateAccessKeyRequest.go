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

package ram

import "aliyun-go-sdk-core"

type UpdateAccessKeyRequest struct {
	UserName        string
	UserAccessKeyId string
	Status          string
}

func (r *UpdateAccessKeyRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 3)

	core.AddToMapIfNotDefaultValueStr(queryVals, "UserName", r.UserName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "UserAccessKeyId", r.UserAccessKeyId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Status", r.Status)

	return queryVals
}

func (r *UpdateAccessKeyRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *UpdateAccessKeyRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *UpdateAccessKeyRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *UpdateAccessKeyRequest) GetPath() string {
	return ""
}

func (r *UpdateAccessKeyRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *UpdateAccessKeyRequest) GetProtocol() string {
	return "HTTPS"
}

func (r *UpdateAccessKeyRequest) GetActionName() string {
	return "UpdateAccessKey"
}
