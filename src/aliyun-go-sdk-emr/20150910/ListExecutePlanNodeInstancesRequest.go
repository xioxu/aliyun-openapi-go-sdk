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

type ListExecutePlanNodeInstancesRequest struct {
	ExecutePlanWorkNodeId int64
}

func (r *ListExecutePlanNodeInstancesRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 1)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "ExecutePlanWorkNodeId", r.ExecutePlanWorkNodeId)

	return queryVals
}

func (r *ListExecutePlanNodeInstancesRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *ListExecutePlanNodeInstancesRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *ListExecutePlanNodeInstancesRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *ListExecutePlanNodeInstancesRequest) GetPath() string {
	return ""
}

func (r *ListExecutePlanNodeInstancesRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *ListExecutePlanNodeInstancesRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *ListExecutePlanNodeInstancesRequest) GetActionName() string {
	return "ListExecutePlanNodeInstances"
}
