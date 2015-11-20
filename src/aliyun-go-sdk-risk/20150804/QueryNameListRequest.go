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

package risk

import "aliyun-go-sdk-core"

type QueryNameListRequest struct {
	Type      string
	DataType  string
	DataValue string
	QueryLike string
	Extend    string
}

func (r *QueryNameListRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 5)

	core.AddToMapIfNotDefaultValueStr(queryVals, "Type", r.Type)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataType", r.DataType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataValue", r.DataValue)
	core.AddToMapIfNotDefaultValueStr(queryVals, "QueryLike", r.QueryLike)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Extend", r.Extend)

	return queryVals
}

func (r *QueryNameListRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *QueryNameListRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *QueryNameListRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *QueryNameListRequest) GetPath() string {
	return ""
}

func (r *QueryNameListRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *QueryNameListRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *QueryNameListRequest) GetActionName() string {
	return "QueryNameList"
}
