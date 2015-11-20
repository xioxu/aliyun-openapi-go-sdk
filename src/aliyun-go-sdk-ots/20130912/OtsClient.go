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

package ots

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-ots/20130912/GetInstanceResponse"

	"aliyun-go-sdk-ots/20130912/DeleteUserResponse"

	"aliyun-go-sdk-ots/20130912/DeleteInstanceResponse"

	"aliyun-go-sdk-ots/20130912/UpdateUserResponse"

	"aliyun-go-sdk-ots/20130912/UpdateInstanceResponse"

	"aliyun-go-sdk-ots/20130912/ListInstanceResponse"

	"aliyun-go-sdk-ots/20130912/InsertUserResponse"

	"aliyun-go-sdk-ots/20130912/InsertInstanceResponse"

	"aliyun-go-sdk-ots/20130912/GetUserResponse"
)

type OtsClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *OtsClient {
	p := &OtsClient{}
	p.Version = "2013-09-12"
	p.ProductName = "Ots"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *OtsClient) GetInstance(getInstanceRequest *GetInstanceRequest) (ots_GetInstanceResponse.GetInstanceResponse, error) {
	var resObj ots_GetInstanceResponse.GetInstanceResponse

	if getInstanceRequest == nil {
		getInstanceRequest = new(GetInstanceRequest)
	}
	err := c.DoAction(getInstanceRequest, &resObj)
	return resObj, err
}

func (c *OtsClient) DeleteUser(deleteUserRequest *DeleteUserRequest) (ots_DeleteUserResponse.DeleteUserResponse, error) {
	var resObj ots_DeleteUserResponse.DeleteUserResponse

	if deleteUserRequest == nil {
		deleteUserRequest = new(DeleteUserRequest)
	}
	err := c.DoAction(deleteUserRequest, &resObj)
	return resObj, err
}

func (c *OtsClient) DeleteInstance(deleteInstanceRequest *DeleteInstanceRequest) (ots_DeleteInstanceResponse.DeleteInstanceResponse, error) {
	var resObj ots_DeleteInstanceResponse.DeleteInstanceResponse

	if deleteInstanceRequest == nil {
		deleteInstanceRequest = new(DeleteInstanceRequest)
	}
	err := c.DoAction(deleteInstanceRequest, &resObj)
	return resObj, err
}

func (c *OtsClient) UpdateUser(updateUserRequest *UpdateUserRequest) (ots_UpdateUserResponse.UpdateUserResponse, error) {
	var resObj ots_UpdateUserResponse.UpdateUserResponse

	if updateUserRequest == nil {
		updateUserRequest = new(UpdateUserRequest)
	}
	err := c.DoAction(updateUserRequest, &resObj)
	return resObj, err
}

func (c *OtsClient) UpdateInstance(updateInstanceRequest *UpdateInstanceRequest) (ots_UpdateInstanceResponse.UpdateInstanceResponse, error) {
	var resObj ots_UpdateInstanceResponse.UpdateInstanceResponse

	if updateInstanceRequest == nil {
		updateInstanceRequest = new(UpdateInstanceRequest)
	}
	err := c.DoAction(updateInstanceRequest, &resObj)
	return resObj, err
}

func (c *OtsClient) ListInstance(listInstanceRequest *ListInstanceRequest) (ots_ListInstanceResponse.ListInstanceResponse, error) {
	var resObj ots_ListInstanceResponse.ListInstanceResponse

	if listInstanceRequest == nil {
		listInstanceRequest = new(ListInstanceRequest)
	}
	err := c.DoAction(listInstanceRequest, &resObj)
	return resObj, err
}

func (c *OtsClient) InsertUser(insertUserRequest *InsertUserRequest) (ots_InsertUserResponse.InsertUserResponse, error) {
	var resObj ots_InsertUserResponse.InsertUserResponse

	if insertUserRequest == nil {
		insertUserRequest = new(InsertUserRequest)
	}
	err := c.DoAction(insertUserRequest, &resObj)
	return resObj, err
}

func (c *OtsClient) InsertInstance(insertInstanceRequest *InsertInstanceRequest) (ots_InsertInstanceResponse.InsertInstanceResponse, error) {
	var resObj ots_InsertInstanceResponse.InsertInstanceResponse

	if insertInstanceRequest == nil {
		insertInstanceRequest = new(InsertInstanceRequest)
	}
	err := c.DoAction(insertInstanceRequest, &resObj)
	return resObj, err
}

func (c *OtsClient) GetUser(getUserRequest *GetUserRequest) (ots_GetUserResponse.GetUserResponse, error) {
	var resObj ots_GetUserResponse.GetUserResponse

	if getUserRequest == nil {
		getUserRequest = new(GetUserRequest)
	}
	err := c.DoAction(getUserRequest, &resObj)
	return resObj, err
}
