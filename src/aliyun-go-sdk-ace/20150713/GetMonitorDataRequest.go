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

package ace

import "aliyun-go-sdk-core"

type GetMonitorDataRequest struct {
	AppId     int64
	Item      string
	StartTime string
	EndTime   string
	CurPage   int32
	PageSize  int32
}

func (r *GetMonitorDataRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 6)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "AppId", r.AppId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Item", r.Item)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StartTime", r.StartTime)
	core.AddToMapIfNotDefaultValueStr(queryVals, "EndTime", r.EndTime)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "CurPage", r.CurPage)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageSize", r.PageSize)

	return queryVals
}

func (r *GetMonitorDataRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *GetMonitorDataRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *GetMonitorDataRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *GetMonitorDataRequest) GetPath() string {
	return ""
}

func (r *GetMonitorDataRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *GetMonitorDataRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *GetMonitorDataRequest) GetActionName() string {
	return "GetMonitorData"
}
