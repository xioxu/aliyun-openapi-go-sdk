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

package drds

import "aliyun-go-sdk-core"

type DescribeDrdsInstancesRequest struct {
	Type string
}

func (r *DescribeDrdsInstancesRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 1)

	core.AddToMapIfNotDefaultValueStr(queryVals, "Type", r.Type)

	return queryVals
}

func (r *DescribeDrdsInstancesRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *DescribeDrdsInstancesRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *DescribeDrdsInstancesRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *DescribeDrdsInstancesRequest) GetPath() string {
	return ""
}

func (r *DescribeDrdsInstancesRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *DescribeDrdsInstancesRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *DescribeDrdsInstancesRequest) GetActionName() string {
	return "DescribeDrdsInstances"
}
