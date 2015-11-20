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

type QueryListMetricRequest struct {
	Project    string
	Metric     string
	Period     string
	StartTime  string
	EndTime    string
	Dimensions string
	Page       string
	Length     string
	Extend     string
	Filter     string
}

func (r *QueryListMetricRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 10)

	core.AddToMapIfNotDefaultValueStr(queryVals, "Project", r.Project)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Metric", r.Metric)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Period", r.Period)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StartTime", r.StartTime)
	core.AddToMapIfNotDefaultValueStr(queryVals, "EndTime", r.EndTime)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Dimensions", r.Dimensions)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Page", r.Page)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Length", r.Length)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Extend", r.Extend)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Filter", r.Filter)

	return queryVals
}

func (r *QueryListMetricRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *QueryListMetricRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *QueryListMetricRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *QueryListMetricRequest) GetPath() string {
	return ""
}

func (r *QueryListMetricRequest) GetRequestMethod() string {
	return "GET"
}

func (r *QueryListMetricRequest) GetProtocol() string {
	return "HTTP"
}

func (r *QueryListMetricRequest) GetActionName() string {
	return "QueryListMetric"
}
