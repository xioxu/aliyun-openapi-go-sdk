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

package yundun

import "aliyun-go-sdk-core"

type AddCNameWafRequest struct {
	InstanceId   string
	InstanceType string
	Domain       string
}

func (r *AddCNameWafRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 3)

	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId", r.InstanceId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceType", r.InstanceType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Domain", r.Domain)

	return queryVals
}

func (r *AddCNameWafRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *AddCNameWafRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *AddCNameWafRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *AddCNameWafRequest) GetPath() string {
	return ""
}

func (r *AddCNameWafRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *AddCNameWafRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *AddCNameWafRequest) GetActionName() string {
	return "AddCNameWaf"
}
