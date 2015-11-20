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

type DescribeScheduledTasksRequest struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	PageNumber           int32
	PageSize             int32
	ScheduledAction_1    string
	ScheduledAction_2    string
	ScheduledAction_3    string
	ScheduledAction_4    string
	ScheduledAction_5    string
	ScheduledAction_6    string
	ScheduledAction_7    string
	ScheduledAction_8    string
	ScheduledAction_9    string
	ScheduledAction_10   string
	ScheduledAction_11   string
	ScheduledAction_12   string
	ScheduledAction_13   string
	ScheduledAction_14   string
	ScheduledAction_15   string
	ScheduledAction_16   string
	ScheduledAction_17   string
	ScheduledAction_18   string
	ScheduledAction_19   string
	ScheduledAction_20   string
	ScheduledTaskId_1    string
	ScheduledTaskId_2    string
	ScheduledTaskId_3    string
	ScheduledTaskId_4    string
	ScheduledTaskId_5    string
	ScheduledTaskId_6    string
	ScheduledTaskId_7    string
	ScheduledTaskId_8    string
	ScheduledTaskId_9    string
	ScheduledTaskId_10   string
	ScheduledTaskId_11   string
	ScheduledTaskId_12   string
	ScheduledTaskId_13   string
	ScheduledTaskId_14   string
	ScheduledTaskId_15   string
	ScheduledTaskId_16   string
	ScheduledTaskId_17   string
	ScheduledTaskId_18   string
	ScheduledTaskId_19   string
	ScheduledTaskId_20   string
	ScheduledTaskName_1  string
	ScheduledTaskName_2  string
	ScheduledTaskName_3  string
	ScheduledTaskName_4  string
	ScheduledTaskName_5  string
	ScheduledTaskName_6  string
	ScheduledTaskName_7  string
	ScheduledTaskName_8  string
	ScheduledTaskName_9  string
	ScheduledTaskName_10 string
	ScheduledTaskName_11 string
	ScheduledTaskName_12 string
	ScheduledTaskName_13 string
	ScheduledTaskName_14 string
	ScheduledTaskName_15 string
	ScheduledTaskName_16 string
	ScheduledTaskName_17 string
	ScheduledTaskName_18 string
	ScheduledTaskName_19 string
	ScheduledTaskName_20 string
	OwnerAccount         string
}

func (r *DescribeScheduledTasksRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 66)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageNumber", r.PageNumber)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageSize", r.PageSize)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledAction.1", r.ScheduledAction_1)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledAction.2", r.ScheduledAction_2)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledAction.3", r.ScheduledAction_3)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledAction.4", r.ScheduledAction_4)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledAction.5", r.ScheduledAction_5)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledAction.6", r.ScheduledAction_6)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledAction.7", r.ScheduledAction_7)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledAction.8", r.ScheduledAction_8)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledAction.9", r.ScheduledAction_9)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledAction.10", r.ScheduledAction_10)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledAction.11", r.ScheduledAction_11)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledAction.12", r.ScheduledAction_12)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledAction.13", r.ScheduledAction_13)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledAction.14", r.ScheduledAction_14)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledAction.15", r.ScheduledAction_15)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledAction.16", r.ScheduledAction_16)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledAction.17", r.ScheduledAction_17)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledAction.18", r.ScheduledAction_18)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledAction.19", r.ScheduledAction_19)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledAction.20", r.ScheduledAction_20)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskId.1", r.ScheduledTaskId_1)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskId.2", r.ScheduledTaskId_2)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskId.3", r.ScheduledTaskId_3)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskId.4", r.ScheduledTaskId_4)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskId.5", r.ScheduledTaskId_5)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskId.6", r.ScheduledTaskId_6)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskId.7", r.ScheduledTaskId_7)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskId.8", r.ScheduledTaskId_8)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskId.9", r.ScheduledTaskId_9)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskId.10", r.ScheduledTaskId_10)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskId.11", r.ScheduledTaskId_11)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskId.12", r.ScheduledTaskId_12)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskId.13", r.ScheduledTaskId_13)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskId.14", r.ScheduledTaskId_14)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskId.15", r.ScheduledTaskId_15)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskId.16", r.ScheduledTaskId_16)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskId.17", r.ScheduledTaskId_17)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskId.18", r.ScheduledTaskId_18)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskId.19", r.ScheduledTaskId_19)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskId.20", r.ScheduledTaskId_20)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskName.1", r.ScheduledTaskName_1)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskName.2", r.ScheduledTaskName_2)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskName.3", r.ScheduledTaskName_3)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskName.4", r.ScheduledTaskName_4)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskName.5", r.ScheduledTaskName_5)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskName.6", r.ScheduledTaskName_6)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskName.7", r.ScheduledTaskName_7)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskName.8", r.ScheduledTaskName_8)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskName.9", r.ScheduledTaskName_9)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskName.10", r.ScheduledTaskName_10)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskName.11", r.ScheduledTaskName_11)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskName.12", r.ScheduledTaskName_12)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskName.13", r.ScheduledTaskName_13)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskName.14", r.ScheduledTaskName_14)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskName.15", r.ScheduledTaskName_15)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskName.16", r.ScheduledTaskName_16)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskName.17", r.ScheduledTaskName_17)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskName.18", r.ScheduledTaskName_18)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskName.19", r.ScheduledTaskName_19)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ScheduledTaskName.20", r.ScheduledTaskName_20)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *DescribeScheduledTasksRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *DescribeScheduledTasksRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *DescribeScheduledTasksRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *DescribeScheduledTasksRequest) GetPath() string {
	return ""
}

func (r *DescribeScheduledTasksRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *DescribeScheduledTasksRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *DescribeScheduledTasksRequest) GetActionName() string {
	return "DescribeScheduledTasks"
}
