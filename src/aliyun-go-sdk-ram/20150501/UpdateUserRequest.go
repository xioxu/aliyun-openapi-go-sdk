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

type UpdateUserRequest struct {
	UserName       string
	NewUserName    string
	NewDisplayName string
	NewMobilePhone string
	NewEmail       string
	NewComments    string
}

func (r *UpdateUserRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 6)

	core.AddToMapIfNotDefaultValueStr(queryVals, "UserName", r.UserName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "NewUserName", r.NewUserName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "NewDisplayName", r.NewDisplayName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "NewMobilePhone", r.NewMobilePhone)
	core.AddToMapIfNotDefaultValueStr(queryVals, "NewEmail", r.NewEmail)
	core.AddToMapIfNotDefaultValueStr(queryVals, "NewComments", r.NewComments)

	return queryVals
}

func (r *UpdateUserRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *UpdateUserRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *UpdateUserRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *UpdateUserRequest) GetPath() string {
	return ""
}

func (r *UpdateUserRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *UpdateUserRequest) GetProtocol() string {
	return "HTTPS"
}

func (r *UpdateUserRequest) GetActionName() string {
	return "UpdateUser"
}
