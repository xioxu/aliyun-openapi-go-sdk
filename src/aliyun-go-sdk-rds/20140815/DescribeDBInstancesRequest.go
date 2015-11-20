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

package rds

import "aliyun-go-sdk-core"

type DescribeDBInstancesRequest struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	ClientToken          string
	ProxyId              string
	Engine               string
	DBInstanceStatus     string
	SearchKey            string
	DBInstanceId         string
	DBInstanceType       string
	PageSize             int32
	PageNumber           int32
	InstanceNetworkType  string
	ConnectionMode       string
	OwnerAccount         string
}

func (r *DescribeDBInstancesRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 15)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ClientToken", r.ClientToken)
	core.AddToMapIfNotDefaultValueStr(queryVals, "proxyId", r.ProxyId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Engine", r.Engine)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DBInstanceStatus", r.DBInstanceStatus)
	core.AddToMapIfNotDefaultValueStr(queryVals, "SearchKey", r.SearchKey)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DBInstanceId", r.DBInstanceId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DBInstanceType", r.DBInstanceType)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageSize", r.PageSize)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageNumber", r.PageNumber)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceNetworkType", r.InstanceNetworkType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ConnectionMode", r.ConnectionMode)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *DescribeDBInstancesRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *DescribeDBInstancesRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *DescribeDBInstancesRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *DescribeDBInstancesRequest) GetPath() string {
	return ""
}

func (r *DescribeDBInstancesRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *DescribeDBInstancesRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *DescribeDBInstancesRequest) GetActionName() string {
	return "DescribeDBInstances"
}
