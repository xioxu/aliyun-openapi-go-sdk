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

type GetMetricStatisticsRequest struct {
	Namespace  string
	MetricName string
	StartTime  string
	EndTime    string
	Interval   string
	Dimensions string
	NextToken  int32
	Length     int32
}

func (r *GetMetricStatisticsRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 8)

	core.AddToMapIfNotDefaultValueStr(queryVals, "Namespace", r.Namespace)
	core.AddToMapIfNotDefaultValueStr(queryVals, "MetricName", r.MetricName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StartTime", r.StartTime)
	core.AddToMapIfNotDefaultValueStr(queryVals, "EndTime", r.EndTime)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Interval", r.Interval)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Dimensions", r.Dimensions)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "NextToken", r.NextToken)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "Length", r.Length)

	return queryVals
}

func (r *GetMetricStatisticsRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *GetMetricStatisticsRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *GetMetricStatisticsRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *GetMetricStatisticsRequest) GetPath() string {
	return ""
}

func (r *GetMetricStatisticsRequest) GetRequestMethod() string {
	return "GET"
}

func (r *GetMetricStatisticsRequest) GetProtocol() string {
	return "HTTP"
}

func (r *GetMetricStatisticsRequest) GetActionName() string {
	return "GetMetricStatistics"
}
