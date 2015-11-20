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

package pts

import "aliyun-go-sdk-core"

type ReportLogSampleRequest struct {
	Wskey      string
	ScenarioId int32
	LogSample  string
}

func (r *ReportLogSampleRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 3)

	core.AddToMapIfNotDefaultValueStr(queryVals, "Wskey", r.Wskey)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "ScenarioId", r.ScenarioId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "LogSample", r.LogSample)

	return queryVals
}

func (r *ReportLogSampleRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *ReportLogSampleRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *ReportLogSampleRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *ReportLogSampleRequest) GetPath() string {
	return ""
}

func (r *ReportLogSampleRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *ReportLogSampleRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *ReportLogSampleRequest) GetActionName() string {
	return "ReportLogSample"
}
