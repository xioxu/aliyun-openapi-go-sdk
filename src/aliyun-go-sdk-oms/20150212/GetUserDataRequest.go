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

package oms

import "aliyun-go-sdk-core"

type GetUserDataRequest struct {
	OwnerId      int64
	OwnerAccount string
	ProductName  string
	DataType     string
	StartTime    string
	EndTime      string
	NextToken    string
	MaxResult    int32
}

func (r *GetUserDataRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 8)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ProductName", r.ProductName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DataType", r.DataType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StartTime", r.StartTime)
	core.AddToMapIfNotDefaultValueStr(queryVals, "EndTime", r.EndTime)
	core.AddToMapIfNotDefaultValueStr(queryVals, "NextToken", r.NextToken)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "MaxResult", r.MaxResult)

	return queryVals
}

func (r *GetUserDataRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *GetUserDataRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *GetUserDataRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *GetUserDataRequest) GetPath() string {
	return ""
}

func (r *GetUserDataRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *GetUserDataRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *GetUserDataRequest) GetActionName() string {
	return "GetUserData"
}
