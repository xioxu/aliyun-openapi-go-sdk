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

type CreateReadOnlyDBInstanceRequest struct {
	OwnerId               int64
	ResourceOwnerAccount  string
	ResourceOwnerId       int64
	ClientToken           string
	ZoneId                string
	DBInstanceId          string
	DBInstanceClass       string
	DBInstanceStorage     int32
	EngineVersion         string
	PayType               string
	DBInstanceDescription string
	InstanceNetworkType   string
	VPCId                 string
	VSwitchId             string
	PrivateIpAddress      string
	OwnerAccount          string
}

func (r *CreateReadOnlyDBInstanceRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 16)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ClientToken", r.ClientToken)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ZoneId", r.ZoneId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DBInstanceId", r.DBInstanceId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DBInstanceClass", r.DBInstanceClass)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "DBInstanceStorage", r.DBInstanceStorage)
	core.AddToMapIfNotDefaultValueStr(queryVals, "EngineVersion", r.EngineVersion)
	core.AddToMapIfNotDefaultValueStr(queryVals, "PayType", r.PayType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DBInstanceDescription", r.DBInstanceDescription)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InstanceNetworkType", r.InstanceNetworkType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "VPCId", r.VPCId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "VSwitchId", r.VSwitchId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "PrivateIpAddress", r.PrivateIpAddress)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *CreateReadOnlyDBInstanceRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *CreateReadOnlyDBInstanceRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *CreateReadOnlyDBInstanceRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *CreateReadOnlyDBInstanceRequest) GetPath() string {
	return ""
}

func (r *CreateReadOnlyDBInstanceRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *CreateReadOnlyDBInstanceRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *CreateReadOnlyDBInstanceRequest) GetActionName() string {
	return "CreateReadOnlyDBInstance"
}
