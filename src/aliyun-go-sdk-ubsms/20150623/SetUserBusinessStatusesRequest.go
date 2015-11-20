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

package ubsms

import "aliyun-go-sdk-core"

type SetUserBusinessStatusesRequest struct {
	Uid           string
	ServiceCode   string
	StatusKey1    string
	StatusValue1  string
	StatusKey2    string
	StatusValue2  string
	StatusKey3    string
	StatusValue3  string
	StatusKey4    string
	StatusValue4  string
	StatusKey5    string
	StatusValue5  string
	StatusKey6    string
	StatusValue6  string
	StatusKey7    string
	StatusValue7  string
	StatusKey8    string
	StatusValue8  string
	StatusKey9    string
	StatusValue9  string
	StatusKey10   string
	StatusValue10 string
	Password      string
}

func (r *SetUserBusinessStatusesRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 23)

	core.AddToMapIfNotDefaultValueStr(queryVals, "Uid", r.Uid)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ServiceCode", r.ServiceCode)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StatusKey1", r.StatusKey1)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StatusValue1", r.StatusValue1)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StatusKey2", r.StatusKey2)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StatusValue2", r.StatusValue2)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StatusKey3", r.StatusKey3)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StatusValue3", r.StatusValue3)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StatusKey4", r.StatusKey4)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StatusValue4", r.StatusValue4)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StatusKey5", r.StatusKey5)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StatusValue5", r.StatusValue5)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StatusKey6", r.StatusKey6)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StatusValue6", r.StatusValue6)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StatusKey7", r.StatusKey7)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StatusValue7", r.StatusValue7)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StatusKey8", r.StatusKey8)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StatusValue8", r.StatusValue8)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StatusKey9", r.StatusKey9)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StatusValue9", r.StatusValue9)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StatusKey10", r.StatusKey10)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StatusValue10", r.StatusValue10)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Password", r.Password)

	return queryVals
}

func (r *SetUserBusinessStatusesRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *SetUserBusinessStatusesRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *SetUserBusinessStatusesRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *SetUserBusinessStatusesRequest) GetPath() string {
	return ""
}

func (r *SetUserBusinessStatusesRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *SetUserBusinessStatusesRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *SetUserBusinessStatusesRequest) GetActionName() string {
	return "SetUserBusinessStatuses"
}
