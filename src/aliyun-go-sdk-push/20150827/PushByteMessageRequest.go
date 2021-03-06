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

package push

import "aliyun-go-sdk-core"

type PushByteMessageRequest struct {
	AppId       int64
	SendType    int32
	Accounts    string
	DeviceIds   string
	PushContent string
}

func (r *PushByteMessageRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 5)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "AppId", r.AppId)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "SendType", r.SendType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Accounts", r.Accounts)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DeviceIds", r.DeviceIds)
	core.AddToMapIfNotDefaultValueStr(queryVals, "PushContent", r.PushContent)

	return queryVals
}

func (r *PushByteMessageRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *PushByteMessageRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *PushByteMessageRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *PushByteMessageRequest) GetPath() string {
	return ""
}

func (r *PushByteMessageRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *PushByteMessageRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *PushByteMessageRequest) GetActionName() string {
	return "PushByteMessage"
}
