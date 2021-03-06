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

package emr

import "aliyun-go-sdk-core"

type ModifyExecutePlanRequest struct {
	ClusterId     int64
	ExecutePlanId int64
	Name          string
	Strategy      int32
	TimeInterval  int32
	StartTime     string
	TimeUnit      string
	JobId         string
}

func (r *ModifyExecutePlanRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 8)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "ClusterId", r.ClusterId)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ExecutePlanId", r.ExecutePlanId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Name", r.Name)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "Strategy", r.Strategy)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "TimeInterval", r.TimeInterval)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StartTime", r.StartTime)
	core.AddToMapIfNotDefaultValueStr(queryVals, "TimeUnit", r.TimeUnit)
	core.AddToMapIfNotDefaultValueStr(queryVals, "JobId", r.JobId)

	return queryVals
}

func (r *ModifyExecutePlanRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *ModifyExecutePlanRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *ModifyExecutePlanRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *ModifyExecutePlanRequest) GetPath() string {
	return ""
}

func (r *ModifyExecutePlanRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *ModifyExecutePlanRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *ModifyExecutePlanRequest) GetActionName() string {
	return "ModifyExecutePlan"
}
