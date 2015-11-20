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

type RevertRpcRequest struct {
	AppId      int64
	DeviceId   string
	RpcContent string
	TimeOut    int32
}

func (r *RevertRpcRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 4)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "AppId", r.AppId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DeviceId", r.DeviceId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "RpcContent", r.RpcContent)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "TimeOut", r.TimeOut)

	return queryVals
}

func (r *RevertRpcRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *RevertRpcRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *RevertRpcRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *RevertRpcRequest) GetPath() string {
	return ""
}

func (r *RevertRpcRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *RevertRpcRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *RevertRpcRequest) GetActionName() string {
	return "RevertRpc"
}
