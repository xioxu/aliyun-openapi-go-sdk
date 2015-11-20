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

type CreateRoleRequest struct {
	RoleName                 string
	Description              string
	AssumeRolePolicyDocument string
}

func (r *CreateRoleRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 3)

	core.AddToMapIfNotDefaultValueStr(queryVals, "RoleName", r.RoleName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Description", r.Description)
	core.AddToMapIfNotDefaultValueStr(queryVals, "AssumeRolePolicyDocument", r.AssumeRolePolicyDocument)

	return queryVals
}

func (r *CreateRoleRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *CreateRoleRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *CreateRoleRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *CreateRoleRequest) GetPath() string {
	return ""
}

func (r *CreateRoleRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *CreateRoleRequest) GetProtocol() string {
	return "HTTPS"
}

func (r *CreateRoleRequest) GetActionName() string {
	return "CreateRole"
}
