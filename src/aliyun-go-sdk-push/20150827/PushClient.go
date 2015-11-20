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

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-push/20150827/BatchGetDevicesInfoResponse"

	"aliyun-go-sdk-push/20150827/RevertRpcResponse"

	"aliyun-go-sdk-push/20150827/PushByteMessageResponse"
)

type PushClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *PushClient {
	p := &PushClient{}
	p.Version = "2015-08-27"
	p.ProductName = "Push"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *PushClient) BatchGetDevicesInfo(batchGetDevicesInfoRequest *BatchGetDevicesInfoRequest) (push_BatchGetDevicesInfoResponse.BatchGetDevicesInfoResponse, error) {
	var resObj push_BatchGetDevicesInfoResponse.BatchGetDevicesInfoResponse

	if batchGetDevicesInfoRequest == nil {
		batchGetDevicesInfoRequest = new(BatchGetDevicesInfoRequest)
	}
	err := c.DoAction(batchGetDevicesInfoRequest, &resObj)
	return resObj, err
}

func (c *PushClient) RevertRpc(revertRpcRequest *RevertRpcRequest) (push_RevertRpcResponse.RevertRpcResponse, error) {
	var resObj push_RevertRpcResponse.RevertRpcResponse

	if revertRpcRequest == nil {
		revertRpcRequest = new(RevertRpcRequest)
	}
	err := c.DoAction(revertRpcRequest, &resObj)
	return resObj, err
}

func (c *PushClient) PushByteMessage(pushByteMessageRequest *PushByteMessageRequest) (push_PushByteMessageResponse.PushByteMessageResponse, error) {
	var resObj push_PushByteMessageResponse.PushByteMessageResponse

	if pushByteMessageRequest == nil {
		pushByteMessageRequest = new(PushByteMessageRequest)
	}
	err := c.DoAction(pushByteMessageRequest, &resObj)
	return resObj, err
}
