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

type OnsConsumerAccumulateRequest struct {
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int64
	ConsumerId   string
	Detail       bool
}

func (r *OnsConsumerAccumulateRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 5)

	core.AddToMapIfNotDefaultValueStr(queryVals, "OnsRegionId", r.OnsRegionId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OnsPlatform", r.OnsPlatform)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "PreventCache", r.PreventCache)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ConsumerId", r.ConsumerId)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "Detail", r.Detail)

	return queryVals
}

func (r *OnsConsumerAccumulateRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *OnsConsumerAccumulateRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *OnsConsumerAccumulateRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *OnsConsumerAccumulateRequest) GetPath() string {
	return ""
}

func (r *OnsConsumerAccumulateRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *OnsConsumerAccumulateRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *OnsConsumerAccumulateRequest) GetActionName() string {
	return "OnsConsumerAccumulate"
}
