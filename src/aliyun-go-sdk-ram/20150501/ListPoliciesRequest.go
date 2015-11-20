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

type ListPoliciesRequest struct {
	PolicyType string
	Marker     string
	MaxItems   int32
}

func (r *ListPoliciesRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 3)

	core.AddToMapIfNotDefaultValueStr(queryVals, "PolicyType", r.PolicyType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Marker", r.Marker)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "MaxItems", r.MaxItems)

	return queryVals
}

func (r *ListPoliciesRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *ListPoliciesRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *ListPoliciesRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *ListPoliciesRequest) GetPath() string {
	return ""
}

func (r *ListPoliciesRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *ListPoliciesRequest) GetProtocol() string {
	return "HTTPS"
}

func (r *ListPoliciesRequest) GetActionName() string {
	return "ListPolicies"
}
