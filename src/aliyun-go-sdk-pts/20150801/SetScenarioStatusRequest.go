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

package pts

import "aliyun-go-sdk-core"

type SetScenarioStatusRequest struct {
	Wskey      string
	ScenarioId int32
	Status     int32
	NodeIp     string
}

func (r *SetScenarioStatusRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 4)

	core.AddToMapIfNotDefaultValueStr(queryVals, "Wskey", r.Wskey)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "ScenarioId", r.ScenarioId)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "Status", r.Status)
	core.AddToMapIfNotDefaultValueStr(queryVals, "NodeIp", r.NodeIp)

	return queryVals
}

func (r *SetScenarioStatusRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *SetScenarioStatusRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *SetScenarioStatusRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *SetScenarioStatusRequest) GetPath() string {
	return ""
}

func (r *SetScenarioStatusRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *SetScenarioStatusRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *SetScenarioStatusRequest) GetActionName() string {
	return "SetScenarioStatus"
}
