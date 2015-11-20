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

package ops

import "aliyun-go-sdk-core"

type ReportRequest struct {
	Group       string
	Name        string
	Ip          string
	TypeAlias   string
	Softversion string
	Config      string
}

func (r *ReportRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 6)

	core.AddToMapIfNotDefaultValueStr(queryVals, "group", r.Group)
	core.AddToMapIfNotDefaultValueStr(queryVals, "name", r.Name)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ip", r.Ip)
	core.AddToMapIfNotDefaultValueStr(queryVals, "type", r.TypeAlias)
	core.AddToMapIfNotDefaultValueStr(queryVals, "softversion", r.Softversion)
	core.AddToMapIfNotDefaultValueStr(queryVals, "config", r.Config)

	return queryVals
}

func (r *ReportRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *ReportRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *ReportRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *ReportRequest) GetPath() string {
	return ""
}

func (r *ReportRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *ReportRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *ReportRequest) GetActionName() string {
	return "report"
}
