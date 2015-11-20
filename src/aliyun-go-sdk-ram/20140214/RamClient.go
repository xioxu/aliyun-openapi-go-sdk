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

package ram

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-ram/20140214/RemoveUserResponse"

	"aliyun-go-sdk-ram/20140214/PutUserPolicyResponse"

	"aliyun-go-sdk-ram/20140214/ListUsersResponse"

	"aliyun-go-sdk-ram/20140214/ListUserPoliciesResponse"

	"aliyun-go-sdk-ram/20140214/GetUserPolicyResponse"

	"aliyun-go-sdk-ram/20140214/GetUserResponse"

	"aliyun-go-sdk-ram/20140214/DeleteUserPolicyResponse"

	"aliyun-go-sdk-ram/20140214/AddUserResponse"
)

type RamClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *RamClient {
	p := &RamClient{}
	p.Version = "2014-02-14"
	p.ProductName = "Ram"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *RamClient) RemoveUser(removeUserRequest *RemoveUserRequest) (ram_RemoveUserResponse.RemoveUserResponse, error) {
	var resObj ram_RemoveUserResponse.RemoveUserResponse

	if removeUserRequest == nil {
		removeUserRequest = new(RemoveUserRequest)
	}
	err := c.DoAction(removeUserRequest, &resObj)
	return resObj, err
}

func (c *RamClient) PutUserPolicy(putUserPolicyRequest *PutUserPolicyRequest) (ram_PutUserPolicyResponse.PutUserPolicyResponse, error) {
	var resObj ram_PutUserPolicyResponse.PutUserPolicyResponse

	if putUserPolicyRequest == nil {
		putUserPolicyRequest = new(PutUserPolicyRequest)
	}
	err := c.DoAction(putUserPolicyRequest, &resObj)
	return resObj, err
}

func (c *RamClient) ListUsers(listUsersRequest *ListUsersRequest) (ram_ListUsersResponse.ListUsersResponse, error) {
	var resObj ram_ListUsersResponse.ListUsersResponse

	if listUsersRequest == nil {
		listUsersRequest = new(ListUsersRequest)
	}
	err := c.DoAction(listUsersRequest, &resObj)
	return resObj, err
}

func (c *RamClient) ListUserPolicies(listUserPoliciesRequest *ListUserPoliciesRequest) (ram_ListUserPoliciesResponse.ListUserPoliciesResponse, error) {
	var resObj ram_ListUserPoliciesResponse.ListUserPoliciesResponse

	if listUserPoliciesRequest == nil {
		listUserPoliciesRequest = new(ListUserPoliciesRequest)
	}
	err := c.DoAction(listUserPoliciesRequest, &resObj)
	return resObj, err
}

func (c *RamClient) GetUserPolicy(getUserPolicyRequest *GetUserPolicyRequest) (ram_GetUserPolicyResponse.GetUserPolicyResponse, error) {
	var resObj ram_GetUserPolicyResponse.GetUserPolicyResponse

	if getUserPolicyRequest == nil {
		getUserPolicyRequest = new(GetUserPolicyRequest)
	}
	err := c.DoAction(getUserPolicyRequest, &resObj)
	return resObj, err
}

func (c *RamClient) GetUser(getUserRequest *GetUserRequest) (ram_GetUserResponse.GetUserResponse, error) {
	var resObj ram_GetUserResponse.GetUserResponse

	if getUserRequest == nil {
		getUserRequest = new(GetUserRequest)
	}
	err := c.DoAction(getUserRequest, &resObj)
	return resObj, err
}

func (c *RamClient) DeleteUserPolicy(deleteUserPolicyRequest *DeleteUserPolicyRequest) (ram_DeleteUserPolicyResponse.DeleteUserPolicyResponse, error) {
	var resObj ram_DeleteUserPolicyResponse.DeleteUserPolicyResponse

	if deleteUserPolicyRequest == nil {
		deleteUserPolicyRequest = new(DeleteUserPolicyRequest)
	}
	err := c.DoAction(deleteUserPolicyRequest, &resObj)
	return resObj, err
}

func (c *RamClient) AddUser(addUserRequest *AddUserRequest) (ram_AddUserResponse.AddUserResponse, error) {
	var resObj ram_AddUserResponse.AddUserResponse

	if addUserRequest == nil {
		addUserRequest = new(AddUserRequest)
	}
	err := c.DoAction(addUserRequest, &resObj)
	return resObj, err
}
