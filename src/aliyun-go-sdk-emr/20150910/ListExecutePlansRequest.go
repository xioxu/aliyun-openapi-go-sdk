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

type ListExecutePlansRequest struct {
	Strategy   int32
	Status     string
	IsDesc     bool
	PageNumber int32
	PageSize   int32
}

func (r *ListExecutePlansRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 5)

	core.AddToMapIfNotDefaultValueInt32(queryVals, "Strategy", r.Strategy)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Status", r.Status)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "IsDesc", r.IsDesc)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageNumber", r.PageNumber)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageSize", r.PageSize)

	return queryVals
}

func (r *ListExecutePlansRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *ListExecutePlansRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *ListExecutePlansRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *ListExecutePlansRequest) GetPath() string {
	return ""
}

func (r *ListExecutePlansRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *ListExecutePlansRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *ListExecutePlansRequest) GetActionName() string {
	return "ListExecutePlans"
}
