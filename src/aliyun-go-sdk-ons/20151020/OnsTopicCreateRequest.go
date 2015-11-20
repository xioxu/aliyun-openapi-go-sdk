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

type OnsTopicCreateRequest struct {
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int64
	Topic        string
	Cluster      string
	QueueNum     int32
	Order        bool
	Qps          int64
	Status       int32
	Remark       string
	Appkey       string
	AppName      string
}

func (r *OnsTopicCreateRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 12)

	core.AddToMapIfNotDefaultValueStr(queryVals, "OnsRegionId", r.OnsRegionId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OnsPlatform", r.OnsPlatform)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "PreventCache", r.PreventCache)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Topic", r.Topic)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Cluster", r.Cluster)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "QueueNum", r.QueueNum)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "Order", r.Order)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "Qps", r.Qps)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "Status", r.Status)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Remark", r.Remark)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Appkey", r.Appkey)
	core.AddToMapIfNotDefaultValueStr(queryVals, "AppName", r.AppName)

	return queryVals
}

func (r *OnsTopicCreateRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *OnsTopicCreateRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *OnsTopicCreateRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *OnsTopicCreateRequest) GetPath() string {
	return ""
}

func (r *OnsTopicCreateRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *OnsTopicCreateRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *OnsTopicCreateRequest) GetActionName() string {
	return "OnsTopicCreate"
}
