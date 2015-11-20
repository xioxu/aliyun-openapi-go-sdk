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

package ram

import "aliyun-go-sdk-core"

type SetDefaultPolicyVersionRequest struct {
	PolicyName string
	VersionId  string
}

func (r *SetDefaultPolicyVersionRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 2)

	core.AddToMapIfNotDefaultValueStr(queryVals, "PolicyName", r.PolicyName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "VersionId", r.VersionId)

	return queryVals
}

func (r *SetDefaultPolicyVersionRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *SetDefaultPolicyVersionRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *SetDefaultPolicyVersionRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *SetDefaultPolicyVersionRequest) GetPath() string {
	return ""
}

func (r *SetDefaultPolicyVersionRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *SetDefaultPolicyVersionRequest) GetProtocol() string {
	return "HTTPS"
}

func (r *SetDefaultPolicyVersionRequest) GetActionName() string {
	return "SetDefaultPolicyVersion"
}
