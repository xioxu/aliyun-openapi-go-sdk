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

type WriteUsscRequest struct {
	AppKey       string
	SignTime     string
	Sign         string
	MteeCode     string
	CodeType     string
	IdType       string
	UserId       string
	ChannelType  string
	VerifyResult string
	UmidToken    string
	Collina      string
	Ip           string
	Extend       string
}

func (r *WriteUsscRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 13)

	core.AddToMapIfNotDefaultValueStr(queryVals, "AppKey", r.AppKey)
	core.AddToMapIfNotDefaultValueStr(queryVals, "SignTime", r.SignTime)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Sign", r.Sign)
	core.AddToMapIfNotDefaultValueStr(queryVals, "MteeCode", r.MteeCode)
	core.AddToMapIfNotDefaultValueStr(queryVals, "CodeType", r.CodeType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "IdType", r.IdType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "UserId", r.UserId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ChannelType", r.ChannelType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "VerifyResult", r.VerifyResult)
	core.AddToMapIfNotDefaultValueStr(queryVals, "UmidToken", r.UmidToken)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Collina", r.Collina)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Ip", r.Ip)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Extend", r.Extend)

	return queryVals
}

func (r *WriteUsscRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *WriteUsscRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *WriteUsscRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *WriteUsscRequest) GetPath() string {
	return ""
}

func (r *WriteUsscRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *WriteUsscRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *WriteUsscRequest) GetActionName() string {
	return "WriteUssc"
}
