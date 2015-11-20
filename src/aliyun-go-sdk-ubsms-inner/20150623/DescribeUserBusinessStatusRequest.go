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

package inner

import "aliyun-go-sdk-core"

type DescribeUserBusinessStatusRequest struct {
	CallerBid   string
	Uid         string
	ServiceCode string
	Password    string
}

func (r *DescribeUserBusinessStatusRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 4)

	core.AddToMapIfNotDefaultValueStr(queryVals, "callerBid", r.CallerBid)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Uid", r.Uid)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ServiceCode", r.ServiceCode)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Password", r.Password)

	return queryVals
}

func (r *DescribeUserBusinessStatusRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *DescribeUserBusinessStatusRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *DescribeUserBusinessStatusRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *DescribeUserBusinessStatusRequest) GetPath() string {
	return ""
}

func (r *DescribeUserBusinessStatusRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *DescribeUserBusinessStatusRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *DescribeUserBusinessStatusRequest) GetActionName() string {
	return "DescribeUserBusinessStatus"
}
