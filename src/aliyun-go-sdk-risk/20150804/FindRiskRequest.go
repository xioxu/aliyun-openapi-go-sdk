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

package risk

import "aliyun-go-sdk-core"

type FindRiskRequest struct {
	MteeCode  string
	CodeType  string
	IdType    string
	UserId    string
	Collina   string
	UmidToken string
	Ip        string
	Email     string
	Phone     string
	Umid      string
	Extend    string
}

func (r *FindRiskRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 11)

	core.AddToMapIfNotDefaultValueStr(queryVals, "MteeCode", r.MteeCode)
	core.AddToMapIfNotDefaultValueStr(queryVals, "CodeType", r.CodeType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "IdType", r.IdType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "UserId", r.UserId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Collina", r.Collina)
	core.AddToMapIfNotDefaultValueStr(queryVals, "UmidToken", r.UmidToken)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Ip", r.Ip)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Email", r.Email)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Phone", r.Phone)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Umid", r.Umid)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Extend", r.Extend)

	return queryVals
}

func (r *FindRiskRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *FindRiskRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *FindRiskRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *FindRiskRequest) GetPath() string {
	return ""
}

func (r *FindRiskRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *FindRiskRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *FindRiskRequest) GetActionName() string {
	return "FindRisk"
}
