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

type ModifyJobRequest struct {
	Id           int64
	Name         string
	Type         string
	EnvParameter string
	FailAct      int32
}

func (r *ModifyJobRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 5)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "Id", r.Id)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Name", r.Name)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Type", r.Type)
	core.AddToMapIfNotDefaultValueStr(queryVals, "EnvParameter", r.EnvParameter)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "FailAct", r.FailAct)

	return queryVals
}

func (r *ModifyJobRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *ModifyJobRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *ModifyJobRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *ModifyJobRequest) GetPath() string {
	return ""
}

func (r *ModifyJobRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *ModifyJobRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *ModifyJobRequest) GetActionName() string {
	return "ModifyJob"
}
