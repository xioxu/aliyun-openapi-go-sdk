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

package emr

import "aliyun-go-sdk-core"

type ListExecutePlanExecuteRecordsRequest struct {
	ExecutePlanId int64
	IsDesc        bool
	Status        string
	PageNumber    int32
	PageSize      int32
}

func (r *ListExecutePlanExecuteRecordsRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 5)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "ExecutePlanId", r.ExecutePlanId)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "IsDesc", r.IsDesc)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Status", r.Status)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageNumber", r.PageNumber)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageSize", r.PageSize)

	return queryVals
}

func (r *ListExecutePlanExecuteRecordsRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *ListExecutePlanExecuteRecordsRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *ListExecutePlanExecuteRecordsRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *ListExecutePlanExecuteRecordsRequest) GetPath() string {
	return ""
}

func (r *ListExecutePlanExecuteRecordsRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *ListExecutePlanExecuteRecordsRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *ListExecutePlanExecuteRecordsRequest) GetActionName() string {
	return "ListExecutePlanExecuteRecords"
}
