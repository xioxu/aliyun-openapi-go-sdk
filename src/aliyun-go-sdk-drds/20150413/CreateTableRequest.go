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

type CreateTableRequest struct {
	DrdsInstanceId     string
	DbName             string
	DdlSql             string
	ShardType          string
	ShardKey           string
	AllowFullTableScan string
}

func (r *CreateTableRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 6)

	core.AddToMapIfNotDefaultValueStr(queryVals, "DrdsInstanceId", r.DrdsInstanceId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DbName", r.DbName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DdlSql", r.DdlSql)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ShardType", r.ShardType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ShardKey", r.ShardKey)
	core.AddToMapIfNotDefaultValueStr(queryVals, "AllowFullTableScan", r.AllowFullTableScan)

	return queryVals
}

func (r *CreateTableRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *CreateTableRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *CreateTableRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *CreateTableRequest) GetPath() string {
	return ""
}

func (r *CreateTableRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *CreateTableRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *CreateTableRequest) GetActionName() string {
	return "CreateTable"
}
