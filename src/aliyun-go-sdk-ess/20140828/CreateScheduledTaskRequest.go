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

package ess

import "aliyun-go-sdk-core"

type CreateScheduledTaskRequest struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	ScheduledTaskName    string
	Description          string
	ScheduledAction      string
	RecurrenceEndTime    string
	LaunchTime           string
	RecurrenceType       string
	RecurrenceValue      string
	TaskEnabled          bool
	LaunchExpirationTime int32
	OwnerAccount         string
}

func (r *CreateScheduledTaskRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 13)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskName", r.ScheduledTaskName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Description", r.Description)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledAction", r.ScheduledAction)
	core.AddToMapIfNotDefaultValueStr(queryVals, "RecurrenceEndTime", r.RecurrenceEndTime)
	core.AddToMapIfNotDefaultValueStr(queryVals, "LaunchTime", r.LaunchTime)
	core.AddToMapIfNotDefaultValueStr(queryVals, "RecurrenceType", r.RecurrenceType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "RecurrenceValue", r.RecurrenceValue)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "TaskEnabled", r.TaskEnabled)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "LaunchExpirationTime", r.LaunchExpirationTime)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *CreateScheduledTaskRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *CreateScheduledTaskRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *CreateScheduledTaskRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *CreateScheduledTaskRequest) GetPath() string {
	return ""
}

func (r *CreateScheduledTaskRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *CreateScheduledTaskRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *CreateScheduledTaskRequest) GetActionName() string {
	return "CreateScheduledTask"
}
