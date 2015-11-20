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

type CreateDBInstanceforFirstPayRequest struct {
	Uid               int64
	Bid               string
	UidLoginEmail     string
	BidLoginEmail     string
	Engine            string
	EngineVersion     string
	DBInstanceClass   string
	DBInstanceStorage int32
	DBInstanceNetType string
	CharacterSetName  string
	DBInstanceRemarks string
	ClientToken       string
	OwnerAccount      string
}

func (r *CreateDBInstanceforFirstPayRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 13)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "uid", r.Uid)
	core.AddToMapIfNotDefaultValueStr(queryVals, "bid", r.Bid)
	core.AddToMapIfNotDefaultValueStr(queryVals, "uidLoginEmail", r.UidLoginEmail)
	core.AddToMapIfNotDefaultValueStr(queryVals, "bidLoginEmail", r.BidLoginEmail)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Engine", r.Engine)
	core.AddToMapIfNotDefaultValueStr(queryVals, "EngineVersion", r.EngineVersion)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DBInstanceClass", r.DBInstanceClass)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "DBInstanceStorage", r.DBInstanceStorage)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DBInstanceNetType", r.DBInstanceNetType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "CharacterSetName", r.CharacterSetName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DBInstanceRemarks", r.DBInstanceRemarks)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ClientToken", r.ClientToken)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *CreateDBInstanceforFirstPayRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *CreateDBInstanceforFirstPayRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *CreateDBInstanceforFirstPayRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *CreateDBInstanceforFirstPayRequest) GetPath() string {
	return ""
}

func (r *CreateDBInstanceforFirstPayRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *CreateDBInstanceforFirstPayRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *CreateDBInstanceforFirstPayRequest) GetActionName() string {
	return "CreateDBInstanceforFirstPay"
}
