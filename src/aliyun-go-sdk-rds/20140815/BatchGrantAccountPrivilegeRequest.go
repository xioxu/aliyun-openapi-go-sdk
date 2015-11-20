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

type BatchGrantAccountPrivilegeRequest struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	DBInstanceId         string
	AccountName          string
	DBName               string
	AccountPrivilege     string
	OwnerAccount         string
}

func (r *BatchGrantAccountPrivilegeRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 8)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DBInstanceId", r.DBInstanceId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "AccountName", r.AccountName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DBName", r.DBName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "AccountPrivilege", r.AccountPrivilege)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *BatchGrantAccountPrivilegeRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *BatchGrantAccountPrivilegeRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *BatchGrantAccountPrivilegeRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *BatchGrantAccountPrivilegeRequest) GetPath() string {
	return ""
}

func (r *BatchGrantAccountPrivilegeRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *BatchGrantAccountPrivilegeRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *BatchGrantAccountPrivilegeRequest) GetActionName() string {
	return "BatchGrantAccountPrivilege"
}
