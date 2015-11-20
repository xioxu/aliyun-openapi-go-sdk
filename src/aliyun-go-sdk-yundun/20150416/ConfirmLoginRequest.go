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

type ConfirmLoginRequest struct {
	InstanceId string
	SourceIp   string
	Time       string
}

func (r *ConfirmLoginRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 3)

	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId", r.InstanceId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "SourceIp", r.SourceIp)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Time", r.Time)

	return queryVals
}

func (r *ConfirmLoginRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *ConfirmLoginRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *ConfirmLoginRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *ConfirmLoginRequest) GetPath() string {
	return ""
}

func (r *ConfirmLoginRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *ConfirmLoginRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *ConfirmLoginRequest) GetActionName() string {
	return "ConfirmLogin"
}
