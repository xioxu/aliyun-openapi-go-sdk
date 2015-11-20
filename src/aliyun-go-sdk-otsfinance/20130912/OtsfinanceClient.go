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

package otsfinance

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-otsfinance/20130912/UpdateUserResponse"

	"aliyun-go-sdk-otsfinance/20130912/UpdateInstanceResponse"

	"aliyun-go-sdk-otsfinance/20130912/ListInstanceResponse"

	"aliyun-go-sdk-otsfinance/20130912/InsertUserResponse"

	"aliyun-go-sdk-otsfinance/20130912/InsertInstanceResponse"

	"aliyun-go-sdk-otsfinance/20130912/GetUserResponse"

	"aliyun-go-sdk-otsfinance/20130912/GetInstanceResponse"

	"aliyun-go-sdk-otsfinance/20130912/DeleteUserResponse"

	"aliyun-go-sdk-otsfinance/20130912/DeleteInstanceResponse"
)

type OtsfinanceClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *OtsfinanceClient {
	p := &OtsfinanceClient{}
	p.Version = "2013-09-12"
	p.ProductName = "OtsFinance"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *OtsfinanceClient) UpdateUser(updateUserRequest *UpdateUserRequest) (otsfinance_UpdateUserResponse.UpdateUserResponse, error) {
	var resObj otsfinance_UpdateUserResponse.UpdateUserResponse

	if updateUserRequest == nil {
		updateUserRequest = new(UpdateUserRequest)
	}
	err := c.DoAction(updateUserRequest, &resObj)
	return resObj, err
}

func (c *OtsfinanceClient) UpdateInstance(updateInstanceRequest *UpdateInstanceRequest) (otsfinance_UpdateInstanceResponse.UpdateInstanceResponse, error) {
	var resObj otsfinance_UpdateInstanceResponse.UpdateInstanceResponse

	if updateInstanceRequest == nil {
		updateInstanceRequest = new(UpdateInstanceRequest)
	}
	err := c.DoAction(updateInstanceRequest, &resObj)
	return resObj, err
}

func (c *OtsfinanceClient) ListInstance(listInstanceRequest *ListInstanceRequest) (otsfinance_ListInstanceResponse.ListInstanceResponse, error) {
	var resObj otsfinance_ListInstanceResponse.ListInstanceResponse

	if listInstanceRequest == nil {
		listInstanceRequest = new(ListInstanceRequest)
	}
	err := c.DoAction(listInstanceRequest, &resObj)
	return resObj, err
}

func (c *OtsfinanceClient) InsertUser(insertUserRequest *InsertUserRequest) (otsfinance_InsertUserResponse.InsertUserResponse, error) {
	var resObj otsfinance_InsertUserResponse.InsertUserResponse

	if insertUserRequest == nil {
		insertUserRequest = new(InsertUserRequest)
	}
	err := c.DoAction(insertUserRequest, &resObj)
	return resObj, err
}

func (c *OtsfinanceClient) InsertInstance(insertInstanceRequest *InsertInstanceRequest) (otsfinance_InsertInstanceResponse.InsertInstanceResponse, error) {
	var resObj otsfinance_InsertInstanceResponse.InsertInstanceResponse

	if insertInstanceRequest == nil {
		insertInstanceRequest = new(InsertInstanceRequest)
	}
	err := c.DoAction(insertInstanceRequest, &resObj)
	return resObj, err
}

func (c *OtsfinanceClient) GetUser(getUserRequest *GetUserRequest) (otsfinance_GetUserResponse.GetUserResponse, error) {
	var resObj otsfinance_GetUserResponse.GetUserResponse

	if getUserRequest == nil {
		getUserRequest = new(GetUserRequest)
	}
	err := c.DoAction(getUserRequest, &resObj)
	return resObj, err
}

func (c *OtsfinanceClient) GetInstance(getInstanceRequest *GetInstanceRequest) (otsfinance_GetInstanceResponse.GetInstanceResponse, error) {
	var resObj otsfinance_GetInstanceResponse.GetInstanceResponse

	if getInstanceRequest == nil {
		getInstanceRequest = new(GetInstanceRequest)
	}
	err := c.DoAction(getInstanceRequest, &resObj)
	return resObj, err
}

func (c *OtsfinanceClient) DeleteUser(deleteUserRequest *DeleteUserRequest) (otsfinance_DeleteUserResponse.DeleteUserResponse, error) {
	var resObj otsfinance_DeleteUserResponse.DeleteUserResponse

	if deleteUserRequest == nil {
		deleteUserRequest = new(DeleteUserRequest)
	}
	err := c.DoAction(deleteUserRequest, &resObj)
	return resObj, err
}

func (c *OtsfinanceClient) DeleteInstance(deleteInstanceRequest *DeleteInstanceRequest) (otsfinance_DeleteInstanceResponse.DeleteInstanceResponse, error) {
	var resObj otsfinance_DeleteInstanceResponse.DeleteInstanceResponse

	if deleteInstanceRequest == nil {
		deleteInstanceRequest = new(DeleteInstanceRequest)
	}
	err := c.DoAction(deleteInstanceRequest, &resObj)
	return resObj, err
}
