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

type UpdateSQLMetricsRequest struct {
	ProjectName string
	MetricName  string
	Sql         string
	IsPublic    int32
}

func (r *UpdateSQLMetricsRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 3)

	core.AddToMapIfNotDefaultValueStr(queryVals, "ProjectName", r.ProjectName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "MetricName", r.MetricName)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "IsPublic", r.IsPublic)

	return queryVals
}

func (r *UpdateSQLMetricsRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 1)

	core.AddToMapIfNotDefaultValueStr(bodyVals, "Sql", r.Sql)

	return bodyVals
}

func (r *UpdateSQLMetricsRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *UpdateSQLMetricsRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *UpdateSQLMetricsRequest) GetPath() string {
	return ""
}

func (r *UpdateSQLMetricsRequest) GetRequestMethod() string {
	return "POST"
}

func (r *UpdateSQLMetricsRequest) GetProtocol() string {
	return "HTTP"
}

func (r *UpdateSQLMetricsRequest) GetActionName() string {
	return "updateSQLMetrics"
}
