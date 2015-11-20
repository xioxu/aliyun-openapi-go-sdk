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

type CreateJobRequest struct {
	JobName    string
	JobType    string
	Parameter  string
	JobFailAct int32
}

func (r *CreateJobRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 4)

	core.AddToMapIfNotDefaultValueStr(queryVals, "JobName", r.JobName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "JobType", r.JobType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Parameter", r.Parameter)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "JobFailAct", r.JobFailAct)

	return queryVals
}

func (r *CreateJobRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *CreateJobRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *CreateJobRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *CreateJobRequest) GetPath() string {
	return ""
}

func (r *CreateJobRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *CreateJobRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *CreateJobRequest) GetActionName() string {
	return "CreateJob"
}
