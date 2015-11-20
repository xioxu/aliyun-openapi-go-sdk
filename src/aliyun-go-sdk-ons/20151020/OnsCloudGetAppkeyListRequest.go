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

type OnsCloudGetAppkeyListRequest struct {
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int64
	IsvId        int64
}

func (r *OnsCloudGetAppkeyListRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 4)

	core.AddToMapIfNotDefaultValueStr(queryVals, "OnsRegionId", r.OnsRegionId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OnsPlatform", r.OnsPlatform)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "PreventCache", r.PreventCache)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "IsvId", r.IsvId)

	return queryVals
}

func (r *OnsCloudGetAppkeyListRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *OnsCloudGetAppkeyListRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *OnsCloudGetAppkeyListRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *OnsCloudGetAppkeyListRequest) GetPath() string {
	return ""
}

func (r *OnsCloudGetAppkeyListRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *OnsCloudGetAppkeyListRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *OnsCloudGetAppkeyListRequest) GetActionName() string {
	return "OnsCloudGetAppkeyList"
}
