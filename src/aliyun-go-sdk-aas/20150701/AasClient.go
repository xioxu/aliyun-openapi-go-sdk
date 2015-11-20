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

package aas

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-aas/20150701/DeleteAccessKeyForAccountResponse"

	"aliyun-go-sdk-aas/20150701/CreateAliyunAccountResponse"

	"aliyun-go-sdk-aas/20150701/CreateAccessKeyForAccountResponse"

	"aliyun-go-sdk-aas/20150701/UpdateAccessKeyStatusForAccountResponse"

	"aliyun-go-sdk-aas/20150701/ListAccessKeysForAccountResponse"

	"aliyun-go-sdk-aas/20150701/GetBasicInfoForAccountResponse"
)

type AasClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *AasClient {
	p := &AasClient{}
	p.Version = "2015-07-01"
	p.ProductName = "Aas"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *AasClient) DeleteAccessKeyForAccount(deleteAccessKeyForAccountRequest *DeleteAccessKeyForAccountRequest) (aas_DeleteAccessKeyForAccountResponse.DeleteAccessKeyForAccountResponse, error) {
	var resObj aas_DeleteAccessKeyForAccountResponse.DeleteAccessKeyForAccountResponse

	if deleteAccessKeyForAccountRequest == nil {
		deleteAccessKeyForAccountRequest = new(DeleteAccessKeyForAccountRequest)
	}
	err := c.DoAction(deleteAccessKeyForAccountRequest, &resObj)
	return resObj, err
}

func (c *AasClient) CreateAliyunAccount(createAliyunAccountRequest *CreateAliyunAccountRequest) (aas_CreateAliyunAccountResponse.CreateAliyunAccountResponse, error) {
	var resObj aas_CreateAliyunAccountResponse.CreateAliyunAccountResponse

	if createAliyunAccountRequest == nil {
		createAliyunAccountRequest = new(CreateAliyunAccountRequest)
	}
	err := c.DoAction(createAliyunAccountRequest, &resObj)
	return resObj, err
}

func (c *AasClient) CreateAccessKeyForAccount(createAccessKeyForAccountRequest *CreateAccessKeyForAccountRequest) (aas_CreateAccessKeyForAccountResponse.CreateAccessKeyForAccountResponse, error) {
	var resObj aas_CreateAccessKeyForAccountResponse.CreateAccessKeyForAccountResponse

	if createAccessKeyForAccountRequest == nil {
		createAccessKeyForAccountRequest = new(CreateAccessKeyForAccountRequest)
	}
	err := c.DoAction(createAccessKeyForAccountRequest, &resObj)
	return resObj, err
}

func (c *AasClient) UpdateAccessKeyStatusForAccount(updateAccessKeyStatusForAccountRequest *UpdateAccessKeyStatusForAccountRequest) (aas_UpdateAccessKeyStatusForAccountResponse.UpdateAccessKeyStatusForAccountResponse, error) {
	var resObj aas_UpdateAccessKeyStatusForAccountResponse.UpdateAccessKeyStatusForAccountResponse

	if updateAccessKeyStatusForAccountRequest == nil {
		updateAccessKeyStatusForAccountRequest = new(UpdateAccessKeyStatusForAccountRequest)
	}
	err := c.DoAction(updateAccessKeyStatusForAccountRequest, &resObj)
	return resObj, err
}

func (c *AasClient) ListAccessKeysForAccount(listAccessKeysForAccountRequest *ListAccessKeysForAccountRequest) (aas_ListAccessKeysForAccountResponse.ListAccessKeysForAccountResponse, error) {
	var resObj aas_ListAccessKeysForAccountResponse.ListAccessKeysForAccountResponse

	if listAccessKeysForAccountRequest == nil {
		listAccessKeysForAccountRequest = new(ListAccessKeysForAccountRequest)
	}
	err := c.DoAction(listAccessKeysForAccountRequest, &resObj)
	return resObj, err
}

func (c *AasClient) GetBasicInfoForAccount(getBasicInfoForAccountRequest *GetBasicInfoForAccountRequest) (aas_GetBasicInfoForAccountResponse.GetBasicInfoForAccountResponse, error) {
	var resObj aas_GetBasicInfoForAccountResponse.GetBasicInfoForAccountResponse

	if getBasicInfoForAccountRequest == nil {
		getBasicInfoForAccountRequest = new(GetBasicInfoForAccountRequest)
	}
	err := c.DoAction(getBasicInfoForAccountRequest, &resObj)
	return resObj, err
}
