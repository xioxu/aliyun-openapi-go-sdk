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

package cms

import "aliyun-go-sdk-core"

type ListSqlMetricsRequest struct {
	ProjectName string
	MetricName  string
	Page        int64
	PageSize    int64
}

func (r *ListSqlMetricsRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 4)

	core.AddToMapIfNotDefaultValueStr(queryVals, "ProjectName", r.ProjectName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "MetricName", r.MetricName)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "Page", r.Page)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "PageSize", r.PageSize)

	return queryVals
}

func (r *ListSqlMetricsRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *ListSqlMetricsRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *ListSqlMetricsRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *ListSqlMetricsRequest) GetPath() string {
	return ""
}

func (r *ListSqlMetricsRequest) GetRequestMethod() string {
	return "GET"
}

func (r *ListSqlMetricsRequest) GetProtocol() string {
	return "HTTP"
}

func (r *ListSqlMetricsRequest) GetActionName() string {
	return "ListSqlMetrics"
}
