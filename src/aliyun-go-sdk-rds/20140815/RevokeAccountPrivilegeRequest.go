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

package rds

import "aliyun-go-sdk-core"

type RevokeAccountPrivilegeRequest struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	DBInstanceId         string
	AccountName          string
	DBName               string
	OwnerAccount         string
}

func (r *RevokeAccountPrivilegeRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 7)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DBInstanceId", r.DBInstanceId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "AccountName", r.AccountName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DBName", r.DBName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *RevokeAccountPrivilegeRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *RevokeAccountPrivilegeRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *RevokeAccountPrivilegeRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *RevokeAccountPrivilegeRequest) GetPath() string {
	return ""
}

func (r *RevokeAccountPrivilegeRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *RevokeAccountPrivilegeRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *RevokeAccountPrivilegeRequest) GetActionName() string {
	return "RevokeAccountPrivilege"
}
