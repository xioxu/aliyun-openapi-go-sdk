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

type BatchQueryMetricRequest struct {
	Project    string
	Metric     string
	Period     string
	StartTime  string
	EndTime    string
	Dimensions string
	Extend     string
	Filter     string
}

func (r *BatchQueryMetricRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 8)

	core.AddToMapIfNotDefaultValueStr(queryVals, "Project", r.Project)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Metric", r.Metric)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Period", r.Period)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StartTime", r.StartTime)
	core.AddToMapIfNotDefaultValueStr(queryVals, "EndTime", r.EndTime)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Dimensions", r.Dimensions)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Extend", r.Extend)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Filter", r.Filter)

	return queryVals
}

func (r *BatchQueryMetricRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *BatchQueryMetricRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *BatchQueryMetricRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *BatchQueryMetricRequest) GetPath() string {
	return ""
}

func (r *BatchQueryMetricRequest) GetRequestMethod() string {
	return "GET"
}

func (r *BatchQueryMetricRequest) GetProtocol() string {
	return "HTTP"
}

func (r *BatchQueryMetricRequest) GetActionName() string {
	return "BatchQueryMetric"
}
