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

type GetMetricsRequest struct {
	ProjectName      string
	MetricStreamName string
	MetricName       string
}

func (r *GetMetricsRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 3)

	core.AddToMapIfNotDefaultValueStr(queryVals, "ProjectName", r.ProjectName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "MetricStreamName", r.MetricStreamName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "MetricName", r.MetricName)

	return queryVals
}

func (r *GetMetricsRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *GetMetricsRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *GetMetricsRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *GetMetricsRequest) GetPath() string {
	return ""
}

func (r *GetMetricsRequest) GetRequestMethod() string {
	return "GET"
}

func (r *GetMetricsRequest) GetProtocol() string {
	return "HTTP"
}

func (r *GetMetricsRequest) GetActionName() string {
	return "GetMetrics"
}
