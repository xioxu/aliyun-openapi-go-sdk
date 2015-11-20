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

package ubsms

import "aliyun-go-sdk-core"

type NotifyUserBusinessCommandRequest struct {
	Uid         string
	ServiceCode string
	Cmd         string
	Region      string
	InstanceId  string
	ClientToken string
	Password    string
}

func (r *NotifyUserBusinessCommandRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 7)

	core.AddToMapIfNotDefaultValueStr(queryVals, "Uid", r.Uid)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ServiceCode", r.ServiceCode)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Cmd", r.Cmd)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Region", r.Region)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceId", r.InstanceId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ClientToken", r.ClientToken)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Password", r.Password)

	return queryVals
}

func (r *NotifyUserBusinessCommandRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *NotifyUserBusinessCommandRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *NotifyUserBusinessCommandRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *NotifyUserBusinessCommandRequest) GetPath() string {
	return ""
}

func (r *NotifyUserBusinessCommandRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *NotifyUserBusinessCommandRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *NotifyUserBusinessCommandRequest) GetActionName() string {
	return "NotifyUserBusinessCommand"
}
