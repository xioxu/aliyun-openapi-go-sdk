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

package otsshihua

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-otsshihua/20151026/ListInstanceResponse"

	"aliyun-go-sdk-otsshihua/20151026/InsertInstanceResponse"

	"aliyun-go-sdk-otsshihua/20151026/GetInstanceResponse"

	"aliyun-go-sdk-otsshihua/20151026/DeleteInstanceResponse"
)

type OtsshihuaClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *OtsshihuaClient {
	p := &OtsshihuaClient{}
	p.Version = "2015-10-26"
	p.ProductName = "OtsShihua"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *OtsshihuaClient) ListInstance(listInstanceRequest *ListInstanceRequest) (otsshihua_ListInstanceResponse.ListInstanceResponse, error) {
	var resObj otsshihua_ListInstanceResponse.ListInstanceResponse

	if listInstanceRequest == nil {
		listInstanceRequest = new(ListInstanceRequest)
	}
	err := c.DoAction(listInstanceRequest, &resObj)
	return resObj, err
}

func (c *OtsshihuaClient) InsertInstance(insertInstanceRequest *InsertInstanceRequest) (otsshihua_InsertInstanceResponse.InsertInstanceResponse, error) {
	var resObj otsshihua_InsertInstanceResponse.InsertInstanceResponse

	if insertInstanceRequest == nil {
		insertInstanceRequest = new(InsertInstanceRequest)
	}
	err := c.DoAction(insertInstanceRequest, &resObj)
	return resObj, err
}

func (c *OtsshihuaClient) GetInstance(getInstanceRequest *GetInstanceRequest) (otsshihua_GetInstanceResponse.GetInstanceResponse, error) {
	var resObj otsshihua_GetInstanceResponse.GetInstanceResponse

	if getInstanceRequest == nil {
		getInstanceRequest = new(GetInstanceRequest)
	}
	err := c.DoAction(getInstanceRequest, &resObj)
	return resObj, err
}

func (c *OtsshihuaClient) DeleteInstance(deleteInstanceRequest *DeleteInstanceRequest) (otsshihua_DeleteInstanceResponse.DeleteInstanceResponse, error) {
	var resObj otsshihua_DeleteInstanceResponse.DeleteInstanceResponse

	if deleteInstanceRequest == nil {
		deleteInstanceRequest = new(DeleteInstanceRequest)
	}
	err := c.DoAction(deleteInstanceRequest, &resObj)
	return resObj, err
}
