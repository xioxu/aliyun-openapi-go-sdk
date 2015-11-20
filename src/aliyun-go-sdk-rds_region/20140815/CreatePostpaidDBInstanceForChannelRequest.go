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

package rds_region

import "aliyun-go-sdk-core"

type CreatePostpaidDBInstanceForChannelRequest struct {
	OwnerId               int64
	ResourceOwnerAccount  string
	ResourceOwnerId       int64
	ClientToken           string
	Engine                string
	EngineVersion         string
	DBInstanceClass       string
	DBInstanceStorage     int32
	DBInstanceNetType     string
	DBInstanceDescription string
	SecurityIPList        string
	OwnerAccount          string
}

func (r *CreatePostpaidDBInstanceForChannelRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 12)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ClientToken", r.ClientToken)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Engine", r.Engine)
	core.AddToMapIfNotDefaultValueStr(queryVals, "EngineVersion", r.EngineVersion)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DBInstanceClass", r.DBInstanceClass)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "DBInstanceStorage", r.DBInstanceStorage)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DBInstanceNetType", r.DBInstanceNetType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DBInstanceDescription", r.DBInstanceDescription)
	core.AddToMapIfNotDefaultValueStr(queryVals, "SecurityIPList", r.SecurityIPList)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *CreatePostpaidDBInstanceForChannelRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *CreatePostpaidDBInstanceForChannelRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *CreatePostpaidDBInstanceForChannelRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *CreatePostpaidDBInstanceForChannelRequest) GetPath() string {
	return ""
}

func (r *CreatePostpaidDBInstanceForChannelRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *CreatePostpaidDBInstanceForChannelRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *CreatePostpaidDBInstanceForChannelRequest) GetActionName() string {
	return "CreatePostpaidDBInstanceForChannel"
}
