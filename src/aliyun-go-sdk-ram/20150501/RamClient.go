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

	"aliyun-go-sdk-ram/20150501/UpdateRoleResponse"

	"aliyun-go-sdk-ram/20150501/SetSecurityPreferenceResponse"

	"aliyun-go-sdk-ram/20150501/ListRolesResponse"

	"aliyun-go-sdk-ram/20150501/ListPoliciesForRoleResponse"

	"aliyun-go-sdk-ram/20150501/GetSecurityPreferenceResponse"

	"aliyun-go-sdk-ram/20150501/GetRoleResponse"

	"aliyun-go-sdk-ram/20150501/GetAccountSummaryResponse"

	"aliyun-go-sdk-ram/20150501/DetachPolicyFromRoleResponse"

	"aliyun-go-sdk-ram/20150501/DeleteRoleResponse"

	"aliyun-go-sdk-ram/20150501/CreateRoleResponse"

	"aliyun-go-sdk-ram/20150501/AttachPolicyToRoleResponse"

	"aliyun-go-sdk-ram/20150501/UnbindMFADeviceResponse"

	"aliyun-go-sdk-ram/20150501/ListVirtualMFADevicesResponse"

	"aliyun-go-sdk-ram/20150501/GetUserMFAInfoResponse"

	"aliyun-go-sdk-ram/20150501/DeleteVirtualMFADeviceResponse"

	"aliyun-go-sdk-ram/20150501/CreateVirtualMFADeviceResponse"

	"aliyun-go-sdk-ram/20150501/BindMFADeviceResponse"

	"aliyun-go-sdk-ram/20150501/UpdateLoginProfileResponse"

	"aliyun-go-sdk-ram/20150501/GetLoginProfileResponse"

	"aliyun-go-sdk-ram/20150501/DeleteLoginProfileResponse"

	"aliyun-go-sdk-ram/20150501/CreateLoginProfileResponse"

	"aliyun-go-sdk-ram/20150501/UpdateUserResponse"

	"aliyun-go-sdk-ram/20150501/ListUsersResponse"

	"aliyun-go-sdk-ram/20150501/GetUserResponse"

	"aliyun-go-sdk-ram/20150501/DeleteUserResponse"

	"aliyun-go-sdk-ram/20150501/CreateUserResponse"

	"aliyun-go-sdk-ram/20150501/UpdateAccessKeyResponse"

	"aliyun-go-sdk-ram/20150501/ListAccessKeysResponse"

	"aliyun-go-sdk-ram/20150501/DeleteAccessKeyResponse"

	"aliyun-go-sdk-ram/20150501/CreateAccessKeyResponse"

	"aliyun-go-sdk-ram/20150501/GetServiceStatusResponse"

	"aliyun-go-sdk-ram/20150501/DeactivateServiceResponse"

	"aliyun-go-sdk-ram/20150501/ActivateServiceResponse"

	"aliyun-go-sdk-ram/20150501/SetPasswordPolicyResponse"

	"aliyun-go-sdk-ram/20150501/SetAccountAliasResponse"

	"aliyun-go-sdk-ram/20150501/GetPasswordPolicyResponse"

	"aliyun-go-sdk-ram/20150501/GetAccountAliasResponse"

	"aliyun-go-sdk-ram/20150501/ClearAccountAliasResponse"

	"aliyun-go-sdk-ram/20150501/SetDefaultPolicyVersionResponse"

	"aliyun-go-sdk-ram/20150501/ListPolicyVersionsResponse"

	"aliyun-go-sdk-ram/20150501/GetPolicyVersionResponse"

	"aliyun-go-sdk-ram/20150501/DeletePolicyVersionResponse"

	"aliyun-go-sdk-ram/20150501/CreatePolicyVersionResponse"

	"aliyun-go-sdk-ram/20150501/ListPoliciesForUserResponse"

	"aliyun-go-sdk-ram/20150501/ListPoliciesForGroupResponse"

	"aliyun-go-sdk-ram/20150501/ListEntitiesForPolicyResponse"

	"aliyun-go-sdk-ram/20150501/DetachPolicyFromUserResponse"

	"aliyun-go-sdk-ram/20150501/DetachPolicyFromGroupResponse"

	"aliyun-go-sdk-ram/20150501/AttachPolicyToUserResponse"

	"aliyun-go-sdk-ram/20150501/AttachPolicyToGroupResponse"

	"aliyun-go-sdk-ram/20150501/ListPoliciesResponse"

	"aliyun-go-sdk-ram/20150501/GetPolicyResponse"

	"aliyun-go-sdk-ram/20150501/DeletePolicyResponse"

	"aliyun-go-sdk-ram/20150501/CreatePolicyResponse"

	"aliyun-go-sdk-ram/20150501/RemoveUserFromGroupResponse"

	"aliyun-go-sdk-ram/20150501/ListUsersForGroupResponse"

	"aliyun-go-sdk-ram/20150501/ListGroupsForUserResponse"

	"aliyun-go-sdk-ram/20150501/AddUserToGroupResponse"

	"aliyun-go-sdk-ram/20150501/UpdateGroupResponse"

	"aliyun-go-sdk-ram/20150501/ListGroupsResponse"

	"aliyun-go-sdk-ram/20150501/GetGroupResponse"

	"aliyun-go-sdk-ram/20150501/DeleteGroupResponse"

	"aliyun-go-sdk-ram/20150501/CreateGroupResponse"
)

type RamClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *RamClient {
	p := &RamClient{}
	p.Version = "2015-05-01"
	p.ProductName = "Ram"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *RamClient) UpdateRole(updateRoleRequest *UpdateRoleRequest) (ram_UpdateRoleResponse.UpdateRoleResponse, error) {
	var resObj ram_UpdateRoleResponse.UpdateRoleResponse

	if updateRoleRequest == nil {
		updateRoleRequest = new(UpdateRoleRequest)
	}
	err := c.DoAction(updateRoleRequest, &resObj)
	return resObj, err
}

func (c *RamClient) SetSecurityPreference(setSecurityPreferenceRequest *SetSecurityPreferenceRequest) (ram_SetSecurityPreferenceResponse.SetSecurityPreferenceResponse, error) {
	var resObj ram_SetSecurityPreferenceResponse.SetSecurityPreferenceResponse

	if setSecurityPreferenceRequest == nil {
		setSecurityPreferenceRequest = new(SetSecurityPreferenceRequest)
	}
	err := c.DoAction(setSecurityPreferenceRequest, &resObj)
	return resObj, err
}

func (c *RamClient) ListRoles(listRolesRequest *ListRolesRequest) (ram_ListRolesResponse.ListRolesResponse, error) {
	var resObj ram_ListRolesResponse.ListRolesResponse

	if listRolesRequest == nil {
		listRolesRequest = new(ListRolesRequest)
	}
	err := c.DoAction(listRolesRequest, &resObj)
	return resObj, err
}

func (c *RamClient) ListPoliciesForRole(listPoliciesForRoleRequest *ListPoliciesForRoleRequest) (ram_ListPoliciesForRoleResponse.ListPoliciesForRoleResponse, error) {
	var resObj ram_ListPoliciesForRoleResponse.ListPoliciesForRoleResponse

	if listPoliciesForRoleRequest == nil {
		listPoliciesForRoleRequest = new(ListPoliciesForRoleRequest)
	}
	err := c.DoAction(listPoliciesForRoleRequest, &resObj)
	return resObj, err
}

func (c *RamClient) GetSecurityPreference(getSecurityPreferenceRequest *GetSecurityPreferenceRequest) (ram_GetSecurityPreferenceResponse.GetSecurityPreferenceResponse, error) {
	var resObj ram_GetSecurityPreferenceResponse.GetSecurityPreferenceResponse

	if getSecurityPreferenceRequest == nil {
		getSecurityPreferenceRequest = new(GetSecurityPreferenceRequest)
	}
	err := c.DoAction(getSecurityPreferenceRequest, &resObj)
	return resObj, err
}

func (c *RamClient) GetRole(getRoleRequest *GetRoleRequest) (ram_GetRoleResponse.GetRoleResponse, error) {
	var resObj ram_GetRoleResponse.GetRoleResponse

	if getRoleRequest == nil {
		getRoleRequest = new(GetRoleRequest)
	}
	err := c.DoAction(getRoleRequest, &resObj)
	return resObj, err
}

func (c *RamClient) GetAccountSummary(getAccountSummaryRequest *GetAccountSummaryRequest) (ram_GetAccountSummaryResponse.GetAccountSummaryResponse, error) {
	var resObj ram_GetAccountSummaryResponse.GetAccountSummaryResponse

	if getAccountSummaryRequest == nil {
		getAccountSummaryRequest = new(GetAccountSummaryRequest)
	}
	err := c.DoAction(getAccountSummaryRequest, &resObj)
	return resObj, err
}

func (c *RamClient) DetachPolicyFromRole(detachPolicyFromRoleRequest *DetachPolicyFromRoleRequest) (ram_DetachPolicyFromRoleResponse.DetachPolicyFromRoleResponse, error) {
	var resObj ram_DetachPolicyFromRoleResponse.DetachPolicyFromRoleResponse

	if detachPolicyFromRoleRequest == nil {
		detachPolicyFromRoleRequest = new(DetachPolicyFromRoleRequest)
	}
	err := c.DoAction(detachPolicyFromRoleRequest, &resObj)
	return resObj, err
}

func (c *RamClient) DeleteRole(deleteRoleRequest *DeleteRoleRequest) (ram_DeleteRoleResponse.DeleteRoleResponse, error) {
	var resObj ram_DeleteRoleResponse.DeleteRoleResponse

	if deleteRoleRequest == nil {
		deleteRoleRequest = new(DeleteRoleRequest)
	}
	err := c.DoAction(deleteRoleRequest, &resObj)
	return resObj, err
}

func (c *RamClient) CreateRole(createRoleRequest *CreateRoleRequest) (ram_CreateRoleResponse.CreateRoleResponse, error) {
	var resObj ram_CreateRoleResponse.CreateRoleResponse

	if createRoleRequest == nil {
		createRoleRequest = new(CreateRoleRequest)
	}
	err := c.DoAction(createRoleRequest, &resObj)
	return resObj, err
}

func (c *RamClient) AttachPolicyToRole(attachPolicyToRoleRequest *AttachPolicyToRoleRequest) (ram_AttachPolicyToRoleResponse.AttachPolicyToRoleResponse, error) {
	var resObj ram_AttachPolicyToRoleResponse.AttachPolicyToRoleResponse

	if attachPolicyToRoleRequest == nil {
		attachPolicyToRoleRequest = new(AttachPolicyToRoleRequest)
	}
	err := c.DoAction(attachPolicyToRoleRequest, &resObj)
	return resObj, err
}

func (c *RamClient) UnbindMFADevice(unbindMFADeviceRequest *UnbindMFADeviceRequest) (ram_UnbindMFADeviceResponse.UnbindMFADeviceResponse, error) {
	var resObj ram_UnbindMFADeviceResponse.UnbindMFADeviceResponse

	if unbindMFADeviceRequest == nil {
		unbindMFADeviceRequest = new(UnbindMFADeviceRequest)
	}
	err := c.DoAction(unbindMFADeviceRequest, &resObj)
	return resObj, err
}

func (c *RamClient) ListVirtualMFADevices(listVirtualMFADevicesRequest *ListVirtualMFADevicesRequest) (ram_ListVirtualMFADevicesResponse.ListVirtualMFADevicesResponse, error) {
	var resObj ram_ListVirtualMFADevicesResponse.ListVirtualMFADevicesResponse

	if listVirtualMFADevicesRequest == nil {
		listVirtualMFADevicesRequest = new(ListVirtualMFADevicesRequest)
	}
	err := c.DoAction(listVirtualMFADevicesRequest, &resObj)
	return resObj, err
}

func (c *RamClient) GetUserMFAInfo(getUserMFAInfoRequest *GetUserMFAInfoRequest) (ram_GetUserMFAInfoResponse.GetUserMFAInfoResponse, error) {
	var resObj ram_GetUserMFAInfoResponse.GetUserMFAInfoResponse

	if getUserMFAInfoRequest == nil {
		getUserMFAInfoRequest = new(GetUserMFAInfoRequest)
	}
	err := c.DoAction(getUserMFAInfoRequest, &resObj)
	return resObj, err
}

func (c *RamClient) DeleteVirtualMFADevice(deleteVirtualMFADeviceRequest *DeleteVirtualMFADeviceRequest) (ram_DeleteVirtualMFADeviceResponse.DeleteVirtualMFADeviceResponse, error) {
	var resObj ram_DeleteVirtualMFADeviceResponse.DeleteVirtualMFADeviceResponse

	if deleteVirtualMFADeviceRequest == nil {
		deleteVirtualMFADeviceRequest = new(DeleteVirtualMFADeviceRequest)
	}
	err := c.DoAction(deleteVirtualMFADeviceRequest, &resObj)
	return resObj, err
}

func (c *RamClient) CreateVirtualMFADevice(createVirtualMFADeviceRequest *CreateVirtualMFADeviceRequest) (ram_CreateVirtualMFADeviceResponse.CreateVirtualMFADeviceResponse, error) {
	var resObj ram_CreateVirtualMFADeviceResponse.CreateVirtualMFADeviceResponse

	if createVirtualMFADeviceRequest == nil {
		createVirtualMFADeviceRequest = new(CreateVirtualMFADeviceRequest)
	}
	err := c.DoAction(createVirtualMFADeviceRequest, &resObj)
	return resObj, err
}

func (c *RamClient) BindMFADevice(bindMFADeviceRequest *BindMFADeviceRequest) (ram_BindMFADeviceResponse.BindMFADeviceResponse, error) {
	var resObj ram_BindMFADeviceResponse.BindMFADeviceResponse

	if bindMFADeviceRequest == nil {
		bindMFADeviceRequest = new(BindMFADeviceRequest)
	}
	err := c.DoAction(bindMFADeviceRequest, &resObj)
	return resObj, err
}

func (c *RamClient) UpdateLoginProfile(updateLoginProfileRequest *UpdateLoginProfileRequest) (ram_UpdateLoginProfileResponse.UpdateLoginProfileResponse, error) {
	var resObj ram_UpdateLoginProfileResponse.UpdateLoginProfileResponse

	if updateLoginProfileRequest == nil {
		updateLoginProfileRequest = new(UpdateLoginProfileRequest)
	}
	err := c.DoAction(updateLoginProfileRequest, &resObj)
	return resObj, err
}

func (c *RamClient) GetLoginProfile(getLoginProfileRequest *GetLoginProfileRequest) (ram_GetLoginProfileResponse.GetLoginProfileResponse, error) {
	var resObj ram_GetLoginProfileResponse.GetLoginProfileResponse

	if getLoginProfileRequest == nil {
		getLoginProfileRequest = new(GetLoginProfileRequest)
	}
	err := c.DoAction(getLoginProfileRequest, &resObj)
	return resObj, err
}

func (c *RamClient) DeleteLoginProfile(deleteLoginProfileRequest *DeleteLoginProfileRequest) (ram_DeleteLoginProfileResponse.DeleteLoginProfileResponse, error) {
	var resObj ram_DeleteLoginProfileResponse.DeleteLoginProfileResponse

	if deleteLoginProfileRequest == nil {
		deleteLoginProfileRequest = new(DeleteLoginProfileRequest)
	}
	err := c.DoAction(deleteLoginProfileRequest, &resObj)
	return resObj, err
}

func (c *RamClient) CreateLoginProfile(createLoginProfileRequest *CreateLoginProfileRequest) (ram_CreateLoginProfileResponse.CreateLoginProfileResponse, error) {
	var resObj ram_CreateLoginProfileResponse.CreateLoginProfileResponse

	if createLoginProfileRequest == nil {
		createLoginProfileRequest = new(CreateLoginProfileRequest)
	}
	err := c.DoAction(createLoginProfileRequest, &resObj)
	return resObj, err
}

func (c *RamClient) UpdateUser(updateUserRequest *UpdateUserRequest) (ram_UpdateUserResponse.UpdateUserResponse, error) {
	var resObj ram_UpdateUserResponse.UpdateUserResponse

	if updateUserRequest == nil {
		updateUserRequest = new(UpdateUserRequest)
	}
	err := c.DoAction(updateUserRequest, &resObj)
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

func (c *RamClient) GetUser(getUserRequest *GetUserRequest) (ram_GetUserResponse.GetUserResponse, error) {
	var resObj ram_GetUserResponse.GetUserResponse

	if getUserRequest == nil {
		getUserRequest = new(GetUserRequest)
	}
	err := c.DoAction(getUserRequest, &resObj)
	return resObj, err
}

func (c *RamClient) DeleteUser(deleteUserRequest *DeleteUserRequest) (ram_DeleteUserResponse.DeleteUserResponse, error) {
	var resObj ram_DeleteUserResponse.DeleteUserResponse

	if deleteUserRequest == nil {
		deleteUserRequest = new(DeleteUserRequest)
	}
	err := c.DoAction(deleteUserRequest, &resObj)
	return resObj, err
}

func (c *RamClient) CreateUser(createUserRequest *CreateUserRequest) (ram_CreateUserResponse.CreateUserResponse, error) {
	var resObj ram_CreateUserResponse.CreateUserResponse

	if createUserRequest == nil {
		createUserRequest = new(CreateUserRequest)
	}
	err := c.DoAction(createUserRequest, &resObj)
	return resObj, err
}

func (c *RamClient) UpdateAccessKey(updateAccessKeyRequest *UpdateAccessKeyRequest) (ram_UpdateAccessKeyResponse.UpdateAccessKeyResponse, error) {
	var resObj ram_UpdateAccessKeyResponse.UpdateAccessKeyResponse

	if updateAccessKeyRequest == nil {
		updateAccessKeyRequest = new(UpdateAccessKeyRequest)
	}
	err := c.DoAction(updateAccessKeyRequest, &resObj)
	return resObj, err
}

func (c *RamClient) ListAccessKeys(listAccessKeysRequest *ListAccessKeysRequest) (ram_ListAccessKeysResponse.ListAccessKeysResponse, error) {
	var resObj ram_ListAccessKeysResponse.ListAccessKeysResponse

	if listAccessKeysRequest == nil {
		listAccessKeysRequest = new(ListAccessKeysRequest)
	}
	err := c.DoAction(listAccessKeysRequest, &resObj)
	return resObj, err
}

func (c *RamClient) DeleteAccessKey(deleteAccessKeyRequest *DeleteAccessKeyRequest) (ram_DeleteAccessKeyResponse.DeleteAccessKeyResponse, error) {
	var resObj ram_DeleteAccessKeyResponse.DeleteAccessKeyResponse

	if deleteAccessKeyRequest == nil {
		deleteAccessKeyRequest = new(DeleteAccessKeyRequest)
	}
	err := c.DoAction(deleteAccessKeyRequest, &resObj)
	return resObj, err
}

func (c *RamClient) CreateAccessKey(createAccessKeyRequest *CreateAccessKeyRequest) (ram_CreateAccessKeyResponse.CreateAccessKeyResponse, error) {
	var resObj ram_CreateAccessKeyResponse.CreateAccessKeyResponse

	if createAccessKeyRequest == nil {
		createAccessKeyRequest = new(CreateAccessKeyRequest)
	}
	err := c.DoAction(createAccessKeyRequest, &resObj)
	return resObj, err
}

func (c *RamClient) GetServiceStatus(getServiceStatusRequest *GetServiceStatusRequest) (ram_GetServiceStatusResponse.GetServiceStatusResponse, error) {
	var resObj ram_GetServiceStatusResponse.GetServiceStatusResponse

	if getServiceStatusRequest == nil {
		getServiceStatusRequest = new(GetServiceStatusRequest)
	}
	err := c.DoAction(getServiceStatusRequest, &resObj)
	return resObj, err
}

func (c *RamClient) DeactivateService(deactivateServiceRequest *DeactivateServiceRequest) (ram_DeactivateServiceResponse.DeactivateServiceResponse, error) {
	var resObj ram_DeactivateServiceResponse.DeactivateServiceResponse

	if deactivateServiceRequest == nil {
		deactivateServiceRequest = new(DeactivateServiceRequest)
	}
	err := c.DoAction(deactivateServiceRequest, &resObj)
	return resObj, err
}

func (c *RamClient) ActivateService(activateServiceRequest *ActivateServiceRequest) (ram_ActivateServiceResponse.ActivateServiceResponse, error) {
	var resObj ram_ActivateServiceResponse.ActivateServiceResponse

	if activateServiceRequest == nil {
		activateServiceRequest = new(ActivateServiceRequest)
	}
	err := c.DoAction(activateServiceRequest, &resObj)
	return resObj, err
}

func (c *RamClient) SetPasswordPolicy(setPasswordPolicyRequest *SetPasswordPolicyRequest) (ram_SetPasswordPolicyResponse.SetPasswordPolicyResponse, error) {
	var resObj ram_SetPasswordPolicyResponse.SetPasswordPolicyResponse

	if setPasswordPolicyRequest == nil {
		setPasswordPolicyRequest = new(SetPasswordPolicyRequest)
	}
	err := c.DoAction(setPasswordPolicyRequest, &resObj)
	return resObj, err
}

func (c *RamClient) SetAccountAlias(setAccountAliasRequest *SetAccountAliasRequest) (ram_SetAccountAliasResponse.SetAccountAliasResponse, error) {
	var resObj ram_SetAccountAliasResponse.SetAccountAliasResponse

	if setAccountAliasRequest == nil {
		setAccountAliasRequest = new(SetAccountAliasRequest)
	}
	err := c.DoAction(setAccountAliasRequest, &resObj)
	return resObj, err
}

func (c *RamClient) GetPasswordPolicy(getPasswordPolicyRequest *GetPasswordPolicyRequest) (ram_GetPasswordPolicyResponse.GetPasswordPolicyResponse, error) {
	var resObj ram_GetPasswordPolicyResponse.GetPasswordPolicyResponse

	if getPasswordPolicyRequest == nil {
		getPasswordPolicyRequest = new(GetPasswordPolicyRequest)
	}
	err := c.DoAction(getPasswordPolicyRequest, &resObj)
	return resObj, err
}

func (c *RamClient) GetAccountAlias(getAccountAliasRequest *GetAccountAliasRequest) (ram_GetAccountAliasResponse.GetAccountAliasResponse, error) {
	var resObj ram_GetAccountAliasResponse.GetAccountAliasResponse

	if getAccountAliasRequest == nil {
		getAccountAliasRequest = new(GetAccountAliasRequest)
	}
	err := c.DoAction(getAccountAliasRequest, &resObj)
	return resObj, err
}

func (c *RamClient) ClearAccountAlias(clearAccountAliasRequest *ClearAccountAliasRequest) (ram_ClearAccountAliasResponse.ClearAccountAliasResponse, error) {
	var resObj ram_ClearAccountAliasResponse.ClearAccountAliasResponse

	if clearAccountAliasRequest == nil {
		clearAccountAliasRequest = new(ClearAccountAliasRequest)
	}
	err := c.DoAction(clearAccountAliasRequest, &resObj)
	return resObj, err
}

func (c *RamClient) SetDefaultPolicyVersion(setDefaultPolicyVersionRequest *SetDefaultPolicyVersionRequest) (ram_SetDefaultPolicyVersionResponse.SetDefaultPolicyVersionResponse, error) {
	var resObj ram_SetDefaultPolicyVersionResponse.SetDefaultPolicyVersionResponse

	if setDefaultPolicyVersionRequest == nil {
		setDefaultPolicyVersionRequest = new(SetDefaultPolicyVersionRequest)
	}
	err := c.DoAction(setDefaultPolicyVersionRequest, &resObj)
	return resObj, err
}

func (c *RamClient) ListPolicyVersions(listPolicyVersionsRequest *ListPolicyVersionsRequest) (ram_ListPolicyVersionsResponse.ListPolicyVersionsResponse, error) {
	var resObj ram_ListPolicyVersionsResponse.ListPolicyVersionsResponse

	if listPolicyVersionsRequest == nil {
		listPolicyVersionsRequest = new(ListPolicyVersionsRequest)
	}
	err := c.DoAction(listPolicyVersionsRequest, &resObj)
	return resObj, err
}

func (c *RamClient) GetPolicyVersion(getPolicyVersionRequest *GetPolicyVersionRequest) (ram_GetPolicyVersionResponse.GetPolicyVersionResponse, error) {
	var resObj ram_GetPolicyVersionResponse.GetPolicyVersionResponse

	if getPolicyVersionRequest == nil {
		getPolicyVersionRequest = new(GetPolicyVersionRequest)
	}
	err := c.DoAction(getPolicyVersionRequest, &resObj)
	return resObj, err
}

func (c *RamClient) DeletePolicyVersion(deletePolicyVersionRequest *DeletePolicyVersionRequest) (ram_DeletePolicyVersionResponse.DeletePolicyVersionResponse, error) {
	var resObj ram_DeletePolicyVersionResponse.DeletePolicyVersionResponse

	if deletePolicyVersionRequest == nil {
		deletePolicyVersionRequest = new(DeletePolicyVersionRequest)
	}
	err := c.DoAction(deletePolicyVersionRequest, &resObj)
	return resObj, err
}

func (c *RamClient) CreatePolicyVersion(createPolicyVersionRequest *CreatePolicyVersionRequest) (ram_CreatePolicyVersionResponse.CreatePolicyVersionResponse, error) {
	var resObj ram_CreatePolicyVersionResponse.CreatePolicyVersionResponse

	if createPolicyVersionRequest == nil {
		createPolicyVersionRequest = new(CreatePolicyVersionRequest)
	}
	err := c.DoAction(createPolicyVersionRequest, &resObj)
	return resObj, err
}

func (c *RamClient) ListPoliciesForUser(listPoliciesForUserRequest *ListPoliciesForUserRequest) (ram_ListPoliciesForUserResponse.ListPoliciesForUserResponse, error) {
	var resObj ram_ListPoliciesForUserResponse.ListPoliciesForUserResponse

	if listPoliciesForUserRequest == nil {
		listPoliciesForUserRequest = new(ListPoliciesForUserRequest)
	}
	err := c.DoAction(listPoliciesForUserRequest, &resObj)
	return resObj, err
}

func (c *RamClient) ListPoliciesForGroup(listPoliciesForGroupRequest *ListPoliciesForGroupRequest) (ram_ListPoliciesForGroupResponse.ListPoliciesForGroupResponse, error) {
	var resObj ram_ListPoliciesForGroupResponse.ListPoliciesForGroupResponse

	if listPoliciesForGroupRequest == nil {
		listPoliciesForGroupRequest = new(ListPoliciesForGroupRequest)
	}
	err := c.DoAction(listPoliciesForGroupRequest, &resObj)
	return resObj, err
}

func (c *RamClient) ListEntitiesForPolicy(listEntitiesForPolicyRequest *ListEntitiesForPolicyRequest) (ram_ListEntitiesForPolicyResponse.ListEntitiesForPolicyResponse, error) {
	var resObj ram_ListEntitiesForPolicyResponse.ListEntitiesForPolicyResponse

	if listEntitiesForPolicyRequest == nil {
		listEntitiesForPolicyRequest = new(ListEntitiesForPolicyRequest)
	}
	err := c.DoAction(listEntitiesForPolicyRequest, &resObj)
	return resObj, err
}

func (c *RamClient) DetachPolicyFromUser(detachPolicyFromUserRequest *DetachPolicyFromUserRequest) (ram_DetachPolicyFromUserResponse.DetachPolicyFromUserResponse, error) {
	var resObj ram_DetachPolicyFromUserResponse.DetachPolicyFromUserResponse

	if detachPolicyFromUserRequest == nil {
		detachPolicyFromUserRequest = new(DetachPolicyFromUserRequest)
	}
	err := c.DoAction(detachPolicyFromUserRequest, &resObj)
	return resObj, err
}

func (c *RamClient) DetachPolicyFromGroup(detachPolicyFromGroupRequest *DetachPolicyFromGroupRequest) (ram_DetachPolicyFromGroupResponse.DetachPolicyFromGroupResponse, error) {
	var resObj ram_DetachPolicyFromGroupResponse.DetachPolicyFromGroupResponse

	if detachPolicyFromGroupRequest == nil {
		detachPolicyFromGroupRequest = new(DetachPolicyFromGroupRequest)
	}
	err := c.DoAction(detachPolicyFromGroupRequest, &resObj)
	return resObj, err
}

func (c *RamClient) AttachPolicyToUser(attachPolicyToUserRequest *AttachPolicyToUserRequest) (ram_AttachPolicyToUserResponse.AttachPolicyToUserResponse, error) {
	var resObj ram_AttachPolicyToUserResponse.AttachPolicyToUserResponse

	if attachPolicyToUserRequest == nil {
		attachPolicyToUserRequest = new(AttachPolicyToUserRequest)
	}
	err := c.DoAction(attachPolicyToUserRequest, &resObj)
	return resObj, err
}

func (c *RamClient) AttachPolicyToGroup(attachPolicyToGroupRequest *AttachPolicyToGroupRequest) (ram_AttachPolicyToGroupResponse.AttachPolicyToGroupResponse, error) {
	var resObj ram_AttachPolicyToGroupResponse.AttachPolicyToGroupResponse

	if attachPolicyToGroupRequest == nil {
		attachPolicyToGroupRequest = new(AttachPolicyToGroupRequest)
	}
	err := c.DoAction(attachPolicyToGroupRequest, &resObj)
	return resObj, err
}

func (c *RamClient) ListPolicies(listPoliciesRequest *ListPoliciesRequest) (ram_ListPoliciesResponse.ListPoliciesResponse, error) {
	var resObj ram_ListPoliciesResponse.ListPoliciesResponse

	if listPoliciesRequest == nil {
		listPoliciesRequest = new(ListPoliciesRequest)
	}
	err := c.DoAction(listPoliciesRequest, &resObj)
	return resObj, err
}

func (c *RamClient) GetPolicy(getPolicyRequest *GetPolicyRequest) (ram_GetPolicyResponse.GetPolicyResponse, error) {
	var resObj ram_GetPolicyResponse.GetPolicyResponse

	if getPolicyRequest == nil {
		getPolicyRequest = new(GetPolicyRequest)
	}
	err := c.DoAction(getPolicyRequest, &resObj)
	return resObj, err
}

func (c *RamClient) DeletePolicy(deletePolicyRequest *DeletePolicyRequest) (ram_DeletePolicyResponse.DeletePolicyResponse, error) {
	var resObj ram_DeletePolicyResponse.DeletePolicyResponse

	if deletePolicyRequest == nil {
		deletePolicyRequest = new(DeletePolicyRequest)
	}
	err := c.DoAction(deletePolicyRequest, &resObj)
	return resObj, err
}

func (c *RamClient) CreatePolicy(createPolicyRequest *CreatePolicyRequest) (ram_CreatePolicyResponse.CreatePolicyResponse, error) {
	var resObj ram_CreatePolicyResponse.CreatePolicyResponse

	if createPolicyRequest == nil {
		createPolicyRequest = new(CreatePolicyRequest)
	}
	err := c.DoAction(createPolicyRequest, &resObj)
	return resObj, err
}

func (c *RamClient) RemoveUserFromGroup(removeUserFromGroupRequest *RemoveUserFromGroupRequest) (ram_RemoveUserFromGroupResponse.RemoveUserFromGroupResponse, error) {
	var resObj ram_RemoveUserFromGroupResponse.RemoveUserFromGroupResponse

	if removeUserFromGroupRequest == nil {
		removeUserFromGroupRequest = new(RemoveUserFromGroupRequest)
	}
	err := c.DoAction(removeUserFromGroupRequest, &resObj)
	return resObj, err
}

func (c *RamClient) ListUsersForGroup(listUsersForGroupRequest *ListUsersForGroupRequest) (ram_ListUsersForGroupResponse.ListUsersForGroupResponse, error) {
	var resObj ram_ListUsersForGroupResponse.ListUsersForGroupResponse

	if listUsersForGroupRequest == nil {
		listUsersForGroupRequest = new(ListUsersForGroupRequest)
	}
	err := c.DoAction(listUsersForGroupRequest, &resObj)
	return resObj, err
}

func (c *RamClient) ListGroupsForUser(listGroupsForUserRequest *ListGroupsForUserRequest) (ram_ListGroupsForUserResponse.ListGroupsForUserResponse, error) {
	var resObj ram_ListGroupsForUserResponse.ListGroupsForUserResponse

	if listGroupsForUserRequest == nil {
		listGroupsForUserRequest = new(ListGroupsForUserRequest)
	}
	err := c.DoAction(listGroupsForUserRequest, &resObj)
	return resObj, err
}

func (c *RamClient) AddUserToGroup(addUserToGroupRequest *AddUserToGroupRequest) (ram_AddUserToGroupResponse.AddUserToGroupResponse, error) {
	var resObj ram_AddUserToGroupResponse.AddUserToGroupResponse

	if addUserToGroupRequest == nil {
		addUserToGroupRequest = new(AddUserToGroupRequest)
	}
	err := c.DoAction(addUserToGroupRequest, &resObj)
	return resObj, err
}

func (c *RamClient) UpdateGroup(updateGroupRequest *UpdateGroupRequest) (ram_UpdateGroupResponse.UpdateGroupResponse, error) {
	var resObj ram_UpdateGroupResponse.UpdateGroupResponse

	if updateGroupRequest == nil {
		updateGroupRequest = new(UpdateGroupRequest)
	}
	err := c.DoAction(updateGroupRequest, &resObj)
	return resObj, err
}

func (c *RamClient) ListGroups(listGroupsRequest *ListGroupsRequest) (ram_ListGroupsResponse.ListGroupsResponse, error) {
	var resObj ram_ListGroupsResponse.ListGroupsResponse

	if listGroupsRequest == nil {
		listGroupsRequest = new(ListGroupsRequest)
	}
	err := c.DoAction(listGroupsRequest, &resObj)
	return resObj, err
}

func (c *RamClient) GetGroup(getGroupRequest *GetGroupRequest) (ram_GetGroupResponse.GetGroupResponse, error) {
	var resObj ram_GetGroupResponse.GetGroupResponse

	if getGroupRequest == nil {
		getGroupRequest = new(GetGroupRequest)
	}
	err := c.DoAction(getGroupRequest, &resObj)
	return resObj, err
}

func (c *RamClient) DeleteGroup(deleteGroupRequest *DeleteGroupRequest) (ram_DeleteGroupResponse.DeleteGroupResponse, error) {
	var resObj ram_DeleteGroupResponse.DeleteGroupResponse

	if deleteGroupRequest == nil {
		deleteGroupRequest = new(DeleteGroupRequest)
	}
	err := c.DoAction(deleteGroupRequest, &resObj)
	return resObj, err
}

func (c *RamClient) CreateGroup(createGroupRequest *CreateGroupRequest) (ram_CreateGroupResponse.CreateGroupResponse, error) {
	var resObj ram_CreateGroupResponse.CreateGroupResponse

	if createGroupRequest == nil {
		createGroupRequest = new(CreateGroupRequest)
	}
	err := c.DoAction(createGroupRequest, &resObj)
	return resObj, err
}
