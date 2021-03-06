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

package ons

import "aliyun-go-sdk-core"

type OnsTrendClusterOutputTpsRequest struct {
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int64
	Cluster      string
	BeginTime    int64
	EndTime      int64
	Period       int64
}

func (r *OnsTrendClusterOutputTpsRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 7)

	core.AddToMapIfNotDefaultValueStr(queryVals, "OnsRegionId", r.OnsRegionId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OnsPlatform", r.OnsPlatform)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "PreventCache", r.PreventCache)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Cluster", r.Cluster)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "BeginTime", r.BeginTime)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "EndTime", r.EndTime)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "Period", r.Period)

	return queryVals
}

func (r *OnsTrendClusterOutputTpsRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *OnsTrendClusterOutputTpsRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *OnsTrendClusterOutputTpsRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *OnsTrendClusterOutputTpsRequest) GetPath() string {
	return ""
}

func (r *OnsTrendClusterOutputTpsRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *OnsTrendClusterOutputTpsRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *OnsTrendClusterOutputTpsRequest) GetActionName() string {
	return "OnsTrendClusterOutputTps"
}
