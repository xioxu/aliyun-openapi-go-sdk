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

type DeleteCNameWafRequest struct {
	InstanceId   string
	Domain       string
	CnameId      int32
	InstanceType string
}

func (r *DeleteCNameWafRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 4)

	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId", r.InstanceId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Domain", r.Domain)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "CnameId", r.CnameId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceType", r.InstanceType)

	return queryVals
}

func (r *DeleteCNameWafRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *DeleteCNameWafRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *DeleteCNameWafRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *DeleteCNameWafRequest) GetPath() string {
	return ""
}

func (r *DeleteCNameWafRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *DeleteCNameWafRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *DeleteCNameWafRequest) GetActionName() string {
	return "DeleteCNameWaf"
}
