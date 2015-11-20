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

package ecs

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-ecs/20140526/UnassociateHaVipResponse"

	"aliyun-go-sdk-ecs/20140526/ModifyHaVipAttributeResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeHaVipsResponse"

	"aliyun-go-sdk-ecs/20140526/DeleteHaVipResponse"

	"aliyun-go-sdk-ecs/20140526/CreateHaVipResponse"

	"aliyun-go-sdk-ecs/20140526/AssociateHaVipResponse"

	"aliyun-go-sdk-ecs/20140526/RenewInstanceResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeTagKeysResponse"

	"aliyun-go-sdk-ecs/20140526/RemoveTagsResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeTagsResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeResourceByTagsResponse"

	"aliyun-go-sdk-ecs/20140526/AddTagsResponse"

	"aliyun-go-sdk-ecs/20140526/UnassociateEipAddressResponse"

	"aliyun-go-sdk-ecs/20140526/StopInstanceResponse"

	"aliyun-go-sdk-ecs/20140526/StartInstanceResponse"

	"aliyun-go-sdk-ecs/20140526/RevokeSecurityGroupEgressResponse"

	"aliyun-go-sdk-ecs/20140526/RevokeSecurityGroupResponse"

	"aliyun-go-sdk-ecs/20140526/ResizeDiskResponse"

	"aliyun-go-sdk-ecs/20140526/ResetDiskResponse"

	"aliyun-go-sdk-ecs/20140526/ReplaceSystemDiskResponse"

	"aliyun-go-sdk-ecs/20140526/ReleaseEipAddressResponse"

	"aliyun-go-sdk-ecs/20140526/ReInitDiskResponse"

	"aliyun-go-sdk-ecs/20140526/RebootInstanceResponse"

	"aliyun-go-sdk-ecs/20140526/ModifyVSwitchAttributeResponse"

	"aliyun-go-sdk-ecs/20140526/ModifyVRouterAttributeResponse"

	"aliyun-go-sdk-ecs/20140526/ModifyVpcAttributeResponse"

	"aliyun-go-sdk-ecs/20140526/ModifySnapshotAttributeResponse"

	"aliyun-go-sdk-ecs/20140526/ModifySecurityGroupAttributeResponse"

	"aliyun-go-sdk-ecs/20140526/ModifyInstanceVpcAttributeResponse"

	"aliyun-go-sdk-ecs/20140526/ModifyInstanceVncPasswdResponse"

	"aliyun-go-sdk-ecs/20140526/ModifyInstanceNetworkSpecResponse"

	"aliyun-go-sdk-ecs/20140526/ModifyInstanceAttributeResponse"

	"aliyun-go-sdk-ecs/20140526/ModifyImageSharePermissionResponse"

	"aliyun-go-sdk-ecs/20140526/ModifyImageShareGroupPermissionResponse"

	"aliyun-go-sdk-ecs/20140526/ModifyImageAttributeResponse"

	"aliyun-go-sdk-ecs/20140526/ModifyEipAddressAttributeResponse"

	"aliyun-go-sdk-ecs/20140526/ModifyDiskAttributeResponse"

	"aliyun-go-sdk-ecs/20140526/ModifyAutoSnapshotPolicyResponse"

	"aliyun-go-sdk-ecs/20140526/LeaveSecurityGroupResponse"

	"aliyun-go-sdk-ecs/20140526/JoinSecurityGroupResponse"

	"aliyun-go-sdk-ecs/20140526/DetachDiskResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeZonesResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeVSwitchesResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeVRoutersResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeVpcsResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeSnapshotsResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeSecurityGroupsResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeSecurityGroupAttributeResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeRouteTablesResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeRegionsResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeLimitationResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeInstanceVncUrlResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeInstanceVncPasswdResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeInstanceTypesResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeInstanceStatusResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeInstancesResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeInstanceMonitorDataResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeInstanceAttributeResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeImageSharePermissionResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeImagesResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeEipMonitorDataResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeEipAddressesResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeDisksResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeDiskMonitorDataResponse"

	"aliyun-go-sdk-ecs/20140526/DescribeAutoSnapshotPolicyResponse"

	"aliyun-go-sdk-ecs/20140526/DeleteVSwitchResponse"

	"aliyun-go-sdk-ecs/20140526/DeleteVpcResponse"

	"aliyun-go-sdk-ecs/20140526/DeleteSnapshotResponse"

	"aliyun-go-sdk-ecs/20140526/DeleteSecurityGroupResponse"

	"aliyun-go-sdk-ecs/20140526/DeleteRouteEntryResponse"

	"aliyun-go-sdk-ecs/20140526/DeleteInstanceResponse"

	"aliyun-go-sdk-ecs/20140526/DeleteImageResponse"

	"aliyun-go-sdk-ecs/20140526/DeleteDiskResponse"

	"aliyun-go-sdk-ecs/20140526/CreateVSwitchResponse"

	"aliyun-go-sdk-ecs/20140526/CreateVpcResponse"

	"aliyun-go-sdk-ecs/20140526/CreateSnapshotResponse"

	"aliyun-go-sdk-ecs/20140526/CreateSecurityGroupResponse"

	"aliyun-go-sdk-ecs/20140526/CreateRouteEntryResponse"

	"aliyun-go-sdk-ecs/20140526/CreateInstanceResponse"

	"aliyun-go-sdk-ecs/20140526/CreateImageResponse"

	"aliyun-go-sdk-ecs/20140526/CreateDiskResponse"

	"aliyun-go-sdk-ecs/20140526/CopyImageResponse"

	"aliyun-go-sdk-ecs/20140526/CheckDiskEnableAutoSnapshotValidationResponse"

	"aliyun-go-sdk-ecs/20140526/CheckAutoSnapshotPolicyResponse"

	"aliyun-go-sdk-ecs/20140526/CancelCopyImageResponse"

	"aliyun-go-sdk-ecs/20140526/AuthorizeSecurityGroupEgressResponse"

	"aliyun-go-sdk-ecs/20140526/AuthorizeSecurityGroupResponse"

	"aliyun-go-sdk-ecs/20140526/AttachDiskResponse"

	"aliyun-go-sdk-ecs/20140526/AssociateEipAddressResponse"

	"aliyun-go-sdk-ecs/20140526/AllocatePublicIpAddressResponse"

	"aliyun-go-sdk-ecs/20140526/AllocateEipAddressResponse"
)

type EcsClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *EcsClient {
	p := &EcsClient{}
	p.Version = "2014-05-26"
	p.ProductName = "Ecs"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *EcsClient) UnassociateHaVip(unassociateHaVipRequest *UnassociateHaVipRequest) (ecs_UnassociateHaVipResponse.UnassociateHaVipResponse, error) {
	var resObj ecs_UnassociateHaVipResponse.UnassociateHaVipResponse

	if unassociateHaVipRequest == nil {
		unassociateHaVipRequest = new(UnassociateHaVipRequest)
	}
	err := c.DoAction(unassociateHaVipRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) ModifyHaVipAttribute(modifyHaVipAttributeRequest *ModifyHaVipAttributeRequest) (ecs_ModifyHaVipAttributeResponse.ModifyHaVipAttributeResponse, error) {
	var resObj ecs_ModifyHaVipAttributeResponse.ModifyHaVipAttributeResponse

	if modifyHaVipAttributeRequest == nil {
		modifyHaVipAttributeRequest = new(ModifyHaVipAttributeRequest)
	}
	err := c.DoAction(modifyHaVipAttributeRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeHaVips(describeHaVipsRequest *DescribeHaVipsRequest) (ecs_DescribeHaVipsResponse.DescribeHaVipsResponse, error) {
	var resObj ecs_DescribeHaVipsResponse.DescribeHaVipsResponse

	if describeHaVipsRequest == nil {
		describeHaVipsRequest = new(DescribeHaVipsRequest)
	}
	err := c.DoAction(describeHaVipsRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DeleteHaVip(deleteHaVipRequest *DeleteHaVipRequest) (ecs_DeleteHaVipResponse.DeleteHaVipResponse, error) {
	var resObj ecs_DeleteHaVipResponse.DeleteHaVipResponse

	if deleteHaVipRequest == nil {
		deleteHaVipRequest = new(DeleteHaVipRequest)
	}
	err := c.DoAction(deleteHaVipRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) CreateHaVip(createHaVipRequest *CreateHaVipRequest) (ecs_CreateHaVipResponse.CreateHaVipResponse, error) {
	var resObj ecs_CreateHaVipResponse.CreateHaVipResponse

	if createHaVipRequest == nil {
		createHaVipRequest = new(CreateHaVipRequest)
	}
	err := c.DoAction(createHaVipRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) AssociateHaVip(associateHaVipRequest *AssociateHaVipRequest) (ecs_AssociateHaVipResponse.AssociateHaVipResponse, error) {
	var resObj ecs_AssociateHaVipResponse.AssociateHaVipResponse

	if associateHaVipRequest == nil {
		associateHaVipRequest = new(AssociateHaVipRequest)
	}
	err := c.DoAction(associateHaVipRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) RenewInstance(renewInstanceRequest *RenewInstanceRequest) (ecs_RenewInstanceResponse.RenewInstanceResponse, error) {
	var resObj ecs_RenewInstanceResponse.RenewInstanceResponse

	if renewInstanceRequest == nil {
		renewInstanceRequest = new(RenewInstanceRequest)
	}
	err := c.DoAction(renewInstanceRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeTagKeys(describeTagKeysRequest *DescribeTagKeysRequest) (ecs_DescribeTagKeysResponse.DescribeTagKeysResponse, error) {
	var resObj ecs_DescribeTagKeysResponse.DescribeTagKeysResponse

	if describeTagKeysRequest == nil {
		describeTagKeysRequest = new(DescribeTagKeysRequest)
	}
	err := c.DoAction(describeTagKeysRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) RemoveTags(removeTagsRequest *RemoveTagsRequest) (ecs_RemoveTagsResponse.RemoveTagsResponse, error) {
	var resObj ecs_RemoveTagsResponse.RemoveTagsResponse

	if removeTagsRequest == nil {
		removeTagsRequest = new(RemoveTagsRequest)
	}
	err := c.DoAction(removeTagsRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeTags(describeTagsRequest *DescribeTagsRequest) (ecs_DescribeTagsResponse.DescribeTagsResponse, error) {
	var resObj ecs_DescribeTagsResponse.DescribeTagsResponse

	if describeTagsRequest == nil {
		describeTagsRequest = new(DescribeTagsRequest)
	}
	err := c.DoAction(describeTagsRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeResourceByTags(describeResourceByTagsRequest *DescribeResourceByTagsRequest) (ecs_DescribeResourceByTagsResponse.DescribeResourceByTagsResponse, error) {
	var resObj ecs_DescribeResourceByTagsResponse.DescribeResourceByTagsResponse

	if describeResourceByTagsRequest == nil {
		describeResourceByTagsRequest = new(DescribeResourceByTagsRequest)
	}
	err := c.DoAction(describeResourceByTagsRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) AddTags(addTagsRequest *AddTagsRequest) (ecs_AddTagsResponse.AddTagsResponse, error) {
	var resObj ecs_AddTagsResponse.AddTagsResponse

	if addTagsRequest == nil {
		addTagsRequest = new(AddTagsRequest)
	}
	err := c.DoAction(addTagsRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) UnassociateEipAddress(unassociateEipAddressRequest *UnassociateEipAddressRequest) (ecs_UnassociateEipAddressResponse.UnassociateEipAddressResponse, error) {
	var resObj ecs_UnassociateEipAddressResponse.UnassociateEipAddressResponse

	if unassociateEipAddressRequest == nil {
		unassociateEipAddressRequest = new(UnassociateEipAddressRequest)
	}
	err := c.DoAction(unassociateEipAddressRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) StopInstance(stopInstanceRequest *StopInstanceRequest) (ecs_StopInstanceResponse.StopInstanceResponse, error) {
	var resObj ecs_StopInstanceResponse.StopInstanceResponse

	if stopInstanceRequest == nil {
		stopInstanceRequest = new(StopInstanceRequest)
	}
	err := c.DoAction(stopInstanceRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) StartInstance(startInstanceRequest *StartInstanceRequest) (ecs_StartInstanceResponse.StartInstanceResponse, error) {
	var resObj ecs_StartInstanceResponse.StartInstanceResponse

	if startInstanceRequest == nil {
		startInstanceRequest = new(StartInstanceRequest)
	}
	err := c.DoAction(startInstanceRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) RevokeSecurityGroupEgress(revokeSecurityGroupEgressRequest *RevokeSecurityGroupEgressRequest) (ecs_RevokeSecurityGroupEgressResponse.RevokeSecurityGroupEgressResponse, error) {
	var resObj ecs_RevokeSecurityGroupEgressResponse.RevokeSecurityGroupEgressResponse

	if revokeSecurityGroupEgressRequest == nil {
		revokeSecurityGroupEgressRequest = new(RevokeSecurityGroupEgressRequest)
	}
	err := c.DoAction(revokeSecurityGroupEgressRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) RevokeSecurityGroup(revokeSecurityGroupRequest *RevokeSecurityGroupRequest) (ecs_RevokeSecurityGroupResponse.RevokeSecurityGroupResponse, error) {
	var resObj ecs_RevokeSecurityGroupResponse.RevokeSecurityGroupResponse

	if revokeSecurityGroupRequest == nil {
		revokeSecurityGroupRequest = new(RevokeSecurityGroupRequest)
	}
	err := c.DoAction(revokeSecurityGroupRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) ResizeDisk(resizeDiskRequest *ResizeDiskRequest) (ecs_ResizeDiskResponse.ResizeDiskResponse, error) {
	var resObj ecs_ResizeDiskResponse.ResizeDiskResponse

	if resizeDiskRequest == nil {
		resizeDiskRequest = new(ResizeDiskRequest)
	}
	err := c.DoAction(resizeDiskRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) ResetDisk(resetDiskRequest *ResetDiskRequest) (ecs_ResetDiskResponse.ResetDiskResponse, error) {
	var resObj ecs_ResetDiskResponse.ResetDiskResponse

	if resetDiskRequest == nil {
		resetDiskRequest = new(ResetDiskRequest)
	}
	err := c.DoAction(resetDiskRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) ReplaceSystemDisk(replaceSystemDiskRequest *ReplaceSystemDiskRequest) (ecs_ReplaceSystemDiskResponse.ReplaceSystemDiskResponse, error) {
	var resObj ecs_ReplaceSystemDiskResponse.ReplaceSystemDiskResponse

	if replaceSystemDiskRequest == nil {
		replaceSystemDiskRequest = new(ReplaceSystemDiskRequest)
	}
	err := c.DoAction(replaceSystemDiskRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) ReleaseEipAddress(releaseEipAddressRequest *ReleaseEipAddressRequest) (ecs_ReleaseEipAddressResponse.ReleaseEipAddressResponse, error) {
	var resObj ecs_ReleaseEipAddressResponse.ReleaseEipAddressResponse

	if releaseEipAddressRequest == nil {
		releaseEipAddressRequest = new(ReleaseEipAddressRequest)
	}
	err := c.DoAction(releaseEipAddressRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) ReInitDisk(reInitDiskRequest *ReInitDiskRequest) (ecs_ReInitDiskResponse.ReInitDiskResponse, error) {
	var resObj ecs_ReInitDiskResponse.ReInitDiskResponse

	if reInitDiskRequest == nil {
		reInitDiskRequest = new(ReInitDiskRequest)
	}
	err := c.DoAction(reInitDiskRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) RebootInstance(rebootInstanceRequest *RebootInstanceRequest) (ecs_RebootInstanceResponse.RebootInstanceResponse, error) {
	var resObj ecs_RebootInstanceResponse.RebootInstanceResponse

	if rebootInstanceRequest == nil {
		rebootInstanceRequest = new(RebootInstanceRequest)
	}
	err := c.DoAction(rebootInstanceRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) ModifyVSwitchAttribute(modifyVSwitchAttributeRequest *ModifyVSwitchAttributeRequest) (ecs_ModifyVSwitchAttributeResponse.ModifyVSwitchAttributeResponse, error) {
	var resObj ecs_ModifyVSwitchAttributeResponse.ModifyVSwitchAttributeResponse

	if modifyVSwitchAttributeRequest == nil {
		modifyVSwitchAttributeRequest = new(ModifyVSwitchAttributeRequest)
	}
	err := c.DoAction(modifyVSwitchAttributeRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) ModifyVRouterAttribute(modifyVRouterAttributeRequest *ModifyVRouterAttributeRequest) (ecs_ModifyVRouterAttributeResponse.ModifyVRouterAttributeResponse, error) {
	var resObj ecs_ModifyVRouterAttributeResponse.ModifyVRouterAttributeResponse

	if modifyVRouterAttributeRequest == nil {
		modifyVRouterAttributeRequest = new(ModifyVRouterAttributeRequest)
	}
	err := c.DoAction(modifyVRouterAttributeRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) ModifyVpcAttribute(modifyVpcAttributeRequest *ModifyVpcAttributeRequest) (ecs_ModifyVpcAttributeResponse.ModifyVpcAttributeResponse, error) {
	var resObj ecs_ModifyVpcAttributeResponse.ModifyVpcAttributeResponse

	if modifyVpcAttributeRequest == nil {
		modifyVpcAttributeRequest = new(ModifyVpcAttributeRequest)
	}
	err := c.DoAction(modifyVpcAttributeRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) ModifySnapshotAttribute(modifySnapshotAttributeRequest *ModifySnapshotAttributeRequest) (ecs_ModifySnapshotAttributeResponse.ModifySnapshotAttributeResponse, error) {
	var resObj ecs_ModifySnapshotAttributeResponse.ModifySnapshotAttributeResponse

	if modifySnapshotAttributeRequest == nil {
		modifySnapshotAttributeRequest = new(ModifySnapshotAttributeRequest)
	}
	err := c.DoAction(modifySnapshotAttributeRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) ModifySecurityGroupAttribute(modifySecurityGroupAttributeRequest *ModifySecurityGroupAttributeRequest) (ecs_ModifySecurityGroupAttributeResponse.ModifySecurityGroupAttributeResponse, error) {
	var resObj ecs_ModifySecurityGroupAttributeResponse.ModifySecurityGroupAttributeResponse

	if modifySecurityGroupAttributeRequest == nil {
		modifySecurityGroupAttributeRequest = new(ModifySecurityGroupAttributeRequest)
	}
	err := c.DoAction(modifySecurityGroupAttributeRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) ModifyInstanceVpcAttribute(modifyInstanceVpcAttributeRequest *ModifyInstanceVpcAttributeRequest) (ecs_ModifyInstanceVpcAttributeResponse.ModifyInstanceVpcAttributeResponse, error) {
	var resObj ecs_ModifyInstanceVpcAttributeResponse.ModifyInstanceVpcAttributeResponse

	if modifyInstanceVpcAttributeRequest == nil {
		modifyInstanceVpcAttributeRequest = new(ModifyInstanceVpcAttributeRequest)
	}
	err := c.DoAction(modifyInstanceVpcAttributeRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) ModifyInstanceVncPasswd(modifyInstanceVncPasswdRequest *ModifyInstanceVncPasswdRequest) (ecs_ModifyInstanceVncPasswdResponse.ModifyInstanceVncPasswdResponse, error) {
	var resObj ecs_ModifyInstanceVncPasswdResponse.ModifyInstanceVncPasswdResponse

	if modifyInstanceVncPasswdRequest == nil {
		modifyInstanceVncPasswdRequest = new(ModifyInstanceVncPasswdRequest)
	}
	err := c.DoAction(modifyInstanceVncPasswdRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) ModifyInstanceNetworkSpec(modifyInstanceNetworkSpecRequest *ModifyInstanceNetworkSpecRequest) (ecs_ModifyInstanceNetworkSpecResponse.ModifyInstanceNetworkSpecResponse, error) {
	var resObj ecs_ModifyInstanceNetworkSpecResponse.ModifyInstanceNetworkSpecResponse

	if modifyInstanceNetworkSpecRequest == nil {
		modifyInstanceNetworkSpecRequest = new(ModifyInstanceNetworkSpecRequest)
	}
	err := c.DoAction(modifyInstanceNetworkSpecRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) ModifyInstanceAttribute(modifyInstanceAttributeRequest *ModifyInstanceAttributeRequest) (ecs_ModifyInstanceAttributeResponse.ModifyInstanceAttributeResponse, error) {
	var resObj ecs_ModifyInstanceAttributeResponse.ModifyInstanceAttributeResponse

	if modifyInstanceAttributeRequest == nil {
		modifyInstanceAttributeRequest = new(ModifyInstanceAttributeRequest)
	}
	err := c.DoAction(modifyInstanceAttributeRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) ModifyImageSharePermission(modifyImageSharePermissionRequest *ModifyImageSharePermissionRequest) (ecs_ModifyImageSharePermissionResponse.ModifyImageSharePermissionResponse, error) {
	var resObj ecs_ModifyImageSharePermissionResponse.ModifyImageSharePermissionResponse

	if modifyImageSharePermissionRequest == nil {
		modifyImageSharePermissionRequest = new(ModifyImageSharePermissionRequest)
	}
	err := c.DoAction(modifyImageSharePermissionRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) ModifyImageShareGroupPermission(modifyImageShareGroupPermissionRequest *ModifyImageShareGroupPermissionRequest) (ecs_ModifyImageShareGroupPermissionResponse.ModifyImageShareGroupPermissionResponse, error) {
	var resObj ecs_ModifyImageShareGroupPermissionResponse.ModifyImageShareGroupPermissionResponse

	if modifyImageShareGroupPermissionRequest == nil {
		modifyImageShareGroupPermissionRequest = new(ModifyImageShareGroupPermissionRequest)
	}
	err := c.DoAction(modifyImageShareGroupPermissionRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) ModifyImageAttribute(modifyImageAttributeRequest *ModifyImageAttributeRequest) (ecs_ModifyImageAttributeResponse.ModifyImageAttributeResponse, error) {
	var resObj ecs_ModifyImageAttributeResponse.ModifyImageAttributeResponse

	if modifyImageAttributeRequest == nil {
		modifyImageAttributeRequest = new(ModifyImageAttributeRequest)
	}
	err := c.DoAction(modifyImageAttributeRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) ModifyEipAddressAttribute(modifyEipAddressAttributeRequest *ModifyEipAddressAttributeRequest) (ecs_ModifyEipAddressAttributeResponse.ModifyEipAddressAttributeResponse, error) {
	var resObj ecs_ModifyEipAddressAttributeResponse.ModifyEipAddressAttributeResponse

	if modifyEipAddressAttributeRequest == nil {
		modifyEipAddressAttributeRequest = new(ModifyEipAddressAttributeRequest)
	}
	err := c.DoAction(modifyEipAddressAttributeRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) ModifyDiskAttribute(modifyDiskAttributeRequest *ModifyDiskAttributeRequest) (ecs_ModifyDiskAttributeResponse.ModifyDiskAttributeResponse, error) {
	var resObj ecs_ModifyDiskAttributeResponse.ModifyDiskAttributeResponse

	if modifyDiskAttributeRequest == nil {
		modifyDiskAttributeRequest = new(ModifyDiskAttributeRequest)
	}
	err := c.DoAction(modifyDiskAttributeRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) ModifyAutoSnapshotPolicy(modifyAutoSnapshotPolicyRequest *ModifyAutoSnapshotPolicyRequest) (ecs_ModifyAutoSnapshotPolicyResponse.ModifyAutoSnapshotPolicyResponse, error) {
	var resObj ecs_ModifyAutoSnapshotPolicyResponse.ModifyAutoSnapshotPolicyResponse

	if modifyAutoSnapshotPolicyRequest == nil {
		modifyAutoSnapshotPolicyRequest = new(ModifyAutoSnapshotPolicyRequest)
	}
	err := c.DoAction(modifyAutoSnapshotPolicyRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) LeaveSecurityGroup(leaveSecurityGroupRequest *LeaveSecurityGroupRequest) (ecs_LeaveSecurityGroupResponse.LeaveSecurityGroupResponse, error) {
	var resObj ecs_LeaveSecurityGroupResponse.LeaveSecurityGroupResponse

	if leaveSecurityGroupRequest == nil {
		leaveSecurityGroupRequest = new(LeaveSecurityGroupRequest)
	}
	err := c.DoAction(leaveSecurityGroupRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) JoinSecurityGroup(joinSecurityGroupRequest *JoinSecurityGroupRequest) (ecs_JoinSecurityGroupResponse.JoinSecurityGroupResponse, error) {
	var resObj ecs_JoinSecurityGroupResponse.JoinSecurityGroupResponse

	if joinSecurityGroupRequest == nil {
		joinSecurityGroupRequest = new(JoinSecurityGroupRequest)
	}
	err := c.DoAction(joinSecurityGroupRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DetachDisk(detachDiskRequest *DetachDiskRequest) (ecs_DetachDiskResponse.DetachDiskResponse, error) {
	var resObj ecs_DetachDiskResponse.DetachDiskResponse

	if detachDiskRequest == nil {
		detachDiskRequest = new(DetachDiskRequest)
	}
	err := c.DoAction(detachDiskRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeZones(describeZonesRequest *DescribeZonesRequest) (ecs_DescribeZonesResponse.DescribeZonesResponse, error) {
	var resObj ecs_DescribeZonesResponse.DescribeZonesResponse

	if describeZonesRequest == nil {
		describeZonesRequest = new(DescribeZonesRequest)
	}
	err := c.DoAction(describeZonesRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeVSwitches(describeVSwitchesRequest *DescribeVSwitchesRequest) (ecs_DescribeVSwitchesResponse.DescribeVSwitchesResponse, error) {
	var resObj ecs_DescribeVSwitchesResponse.DescribeVSwitchesResponse

	if describeVSwitchesRequest == nil {
		describeVSwitchesRequest = new(DescribeVSwitchesRequest)
	}
	err := c.DoAction(describeVSwitchesRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeVRouters(describeVRoutersRequest *DescribeVRoutersRequest) (ecs_DescribeVRoutersResponse.DescribeVRoutersResponse, error) {
	var resObj ecs_DescribeVRoutersResponse.DescribeVRoutersResponse

	if describeVRoutersRequest == nil {
		describeVRoutersRequest = new(DescribeVRoutersRequest)
	}
	err := c.DoAction(describeVRoutersRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeVpcs(describeVpcsRequest *DescribeVpcsRequest) (ecs_DescribeVpcsResponse.DescribeVpcsResponse, error) {
	var resObj ecs_DescribeVpcsResponse.DescribeVpcsResponse

	if describeVpcsRequest == nil {
		describeVpcsRequest = new(DescribeVpcsRequest)
	}
	err := c.DoAction(describeVpcsRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeSnapshots(describeSnapshotsRequest *DescribeSnapshotsRequest) (ecs_DescribeSnapshotsResponse.DescribeSnapshotsResponse, error) {
	var resObj ecs_DescribeSnapshotsResponse.DescribeSnapshotsResponse

	if describeSnapshotsRequest == nil {
		describeSnapshotsRequest = new(DescribeSnapshotsRequest)
	}
	err := c.DoAction(describeSnapshotsRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeSecurityGroups(describeSecurityGroupsRequest *DescribeSecurityGroupsRequest) (ecs_DescribeSecurityGroupsResponse.DescribeSecurityGroupsResponse, error) {
	var resObj ecs_DescribeSecurityGroupsResponse.DescribeSecurityGroupsResponse

	if describeSecurityGroupsRequest == nil {
		describeSecurityGroupsRequest = new(DescribeSecurityGroupsRequest)
	}
	err := c.DoAction(describeSecurityGroupsRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeSecurityGroupAttribute(describeSecurityGroupAttributeRequest *DescribeSecurityGroupAttributeRequest) (ecs_DescribeSecurityGroupAttributeResponse.DescribeSecurityGroupAttributeResponse, error) {
	var resObj ecs_DescribeSecurityGroupAttributeResponse.DescribeSecurityGroupAttributeResponse

	if describeSecurityGroupAttributeRequest == nil {
		describeSecurityGroupAttributeRequest = new(DescribeSecurityGroupAttributeRequest)
	}
	err := c.DoAction(describeSecurityGroupAttributeRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeRouteTables(describeRouteTablesRequest *DescribeRouteTablesRequest) (ecs_DescribeRouteTablesResponse.DescribeRouteTablesResponse, error) {
	var resObj ecs_DescribeRouteTablesResponse.DescribeRouteTablesResponse

	if describeRouteTablesRequest == nil {
		describeRouteTablesRequest = new(DescribeRouteTablesRequest)
	}
	err := c.DoAction(describeRouteTablesRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeRegions(describeRegionsRequest *DescribeRegionsRequest) (ecs_DescribeRegionsResponse.DescribeRegionsResponse, error) {
	var resObj ecs_DescribeRegionsResponse.DescribeRegionsResponse

	if describeRegionsRequest == nil {
		describeRegionsRequest = new(DescribeRegionsRequest)
	}
	err := c.DoAction(describeRegionsRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeLimitation(describeLimitationRequest *DescribeLimitationRequest) (ecs_DescribeLimitationResponse.DescribeLimitationResponse, error) {
	var resObj ecs_DescribeLimitationResponse.DescribeLimitationResponse

	if describeLimitationRequest == nil {
		describeLimitationRequest = new(DescribeLimitationRequest)
	}
	err := c.DoAction(describeLimitationRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeInstanceVncUrl(describeInstanceVncUrlRequest *DescribeInstanceVncUrlRequest) (ecs_DescribeInstanceVncUrlResponse.DescribeInstanceVncUrlResponse, error) {
	var resObj ecs_DescribeInstanceVncUrlResponse.DescribeInstanceVncUrlResponse

	if describeInstanceVncUrlRequest == nil {
		describeInstanceVncUrlRequest = new(DescribeInstanceVncUrlRequest)
	}
	err := c.DoAction(describeInstanceVncUrlRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeInstanceVncPasswd(describeInstanceVncPasswdRequest *DescribeInstanceVncPasswdRequest) (ecs_DescribeInstanceVncPasswdResponse.DescribeInstanceVncPasswdResponse, error) {
	var resObj ecs_DescribeInstanceVncPasswdResponse.DescribeInstanceVncPasswdResponse

	if describeInstanceVncPasswdRequest == nil {
		describeInstanceVncPasswdRequest = new(DescribeInstanceVncPasswdRequest)
	}
	err := c.DoAction(describeInstanceVncPasswdRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeInstanceTypes(describeInstanceTypesRequest *DescribeInstanceTypesRequest) (ecs_DescribeInstanceTypesResponse.DescribeInstanceTypesResponse, error) {
	var resObj ecs_DescribeInstanceTypesResponse.DescribeInstanceTypesResponse

	if describeInstanceTypesRequest == nil {
		describeInstanceTypesRequest = new(DescribeInstanceTypesRequest)
	}
	err := c.DoAction(describeInstanceTypesRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeInstanceStatus(describeInstanceStatusRequest *DescribeInstanceStatusRequest) (ecs_DescribeInstanceStatusResponse.DescribeInstanceStatusResponse, error) {
	var resObj ecs_DescribeInstanceStatusResponse.DescribeInstanceStatusResponse

	if describeInstanceStatusRequest == nil {
		describeInstanceStatusRequest = new(DescribeInstanceStatusRequest)
	}
	err := c.DoAction(describeInstanceStatusRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeInstances(describeInstancesRequest *DescribeInstancesRequest) (ecs_DescribeInstancesResponse.DescribeInstancesResponse, error) {
	var resObj ecs_DescribeInstancesResponse.DescribeInstancesResponse

	if describeInstancesRequest == nil {
		describeInstancesRequest = new(DescribeInstancesRequest)
	}
	err := c.DoAction(describeInstancesRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeInstanceMonitorData(describeInstanceMonitorDataRequest *DescribeInstanceMonitorDataRequest) (ecs_DescribeInstanceMonitorDataResponse.DescribeInstanceMonitorDataResponse, error) {
	var resObj ecs_DescribeInstanceMonitorDataResponse.DescribeInstanceMonitorDataResponse

	if describeInstanceMonitorDataRequest == nil {
		describeInstanceMonitorDataRequest = new(DescribeInstanceMonitorDataRequest)
	}
	err := c.DoAction(describeInstanceMonitorDataRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeInstanceAttribute(describeInstanceAttributeRequest *DescribeInstanceAttributeRequest) (ecs_DescribeInstanceAttributeResponse.DescribeInstanceAttributeResponse, error) {
	var resObj ecs_DescribeInstanceAttributeResponse.DescribeInstanceAttributeResponse

	if describeInstanceAttributeRequest == nil {
		describeInstanceAttributeRequest = new(DescribeInstanceAttributeRequest)
	}
	err := c.DoAction(describeInstanceAttributeRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeImageSharePermission(describeImageSharePermissionRequest *DescribeImageSharePermissionRequest) (ecs_DescribeImageSharePermissionResponse.DescribeImageSharePermissionResponse, error) {
	var resObj ecs_DescribeImageSharePermissionResponse.DescribeImageSharePermissionResponse

	if describeImageSharePermissionRequest == nil {
		describeImageSharePermissionRequest = new(DescribeImageSharePermissionRequest)
	}
	err := c.DoAction(describeImageSharePermissionRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeImages(describeImagesRequest *DescribeImagesRequest) (ecs_DescribeImagesResponse.DescribeImagesResponse, error) {
	var resObj ecs_DescribeImagesResponse.DescribeImagesResponse

	if describeImagesRequest == nil {
		describeImagesRequest = new(DescribeImagesRequest)
	}
	err := c.DoAction(describeImagesRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeEipMonitorData(describeEipMonitorDataRequest *DescribeEipMonitorDataRequest) (ecs_DescribeEipMonitorDataResponse.DescribeEipMonitorDataResponse, error) {
	var resObj ecs_DescribeEipMonitorDataResponse.DescribeEipMonitorDataResponse

	if describeEipMonitorDataRequest == nil {
		describeEipMonitorDataRequest = new(DescribeEipMonitorDataRequest)
	}
	err := c.DoAction(describeEipMonitorDataRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeEipAddresses(describeEipAddressesRequest *DescribeEipAddressesRequest) (ecs_DescribeEipAddressesResponse.DescribeEipAddressesResponse, error) {
	var resObj ecs_DescribeEipAddressesResponse.DescribeEipAddressesResponse

	if describeEipAddressesRequest == nil {
		describeEipAddressesRequest = new(DescribeEipAddressesRequest)
	}
	err := c.DoAction(describeEipAddressesRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeDisks(describeDisksRequest *DescribeDisksRequest) (ecs_DescribeDisksResponse.DescribeDisksResponse, error) {
	var resObj ecs_DescribeDisksResponse.DescribeDisksResponse

	if describeDisksRequest == nil {
		describeDisksRequest = new(DescribeDisksRequest)
	}
	err := c.DoAction(describeDisksRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeDiskMonitorData(describeDiskMonitorDataRequest *DescribeDiskMonitorDataRequest) (ecs_DescribeDiskMonitorDataResponse.DescribeDiskMonitorDataResponse, error) {
	var resObj ecs_DescribeDiskMonitorDataResponse.DescribeDiskMonitorDataResponse

	if describeDiskMonitorDataRequest == nil {
		describeDiskMonitorDataRequest = new(DescribeDiskMonitorDataRequest)
	}
	err := c.DoAction(describeDiskMonitorDataRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DescribeAutoSnapshotPolicy(describeAutoSnapshotPolicyRequest *DescribeAutoSnapshotPolicyRequest) (ecs_DescribeAutoSnapshotPolicyResponse.DescribeAutoSnapshotPolicyResponse, error) {
	var resObj ecs_DescribeAutoSnapshotPolicyResponse.DescribeAutoSnapshotPolicyResponse

	if describeAutoSnapshotPolicyRequest == nil {
		describeAutoSnapshotPolicyRequest = new(DescribeAutoSnapshotPolicyRequest)
	}
	err := c.DoAction(describeAutoSnapshotPolicyRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DeleteVSwitch(deleteVSwitchRequest *DeleteVSwitchRequest) (ecs_DeleteVSwitchResponse.DeleteVSwitchResponse, error) {
	var resObj ecs_DeleteVSwitchResponse.DeleteVSwitchResponse

	if deleteVSwitchRequest == nil {
		deleteVSwitchRequest = new(DeleteVSwitchRequest)
	}
	err := c.DoAction(deleteVSwitchRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DeleteVpc(deleteVpcRequest *DeleteVpcRequest) (ecs_DeleteVpcResponse.DeleteVpcResponse, error) {
	var resObj ecs_DeleteVpcResponse.DeleteVpcResponse

	if deleteVpcRequest == nil {
		deleteVpcRequest = new(DeleteVpcRequest)
	}
	err := c.DoAction(deleteVpcRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DeleteSnapshot(deleteSnapshotRequest *DeleteSnapshotRequest) (ecs_DeleteSnapshotResponse.DeleteSnapshotResponse, error) {
	var resObj ecs_DeleteSnapshotResponse.DeleteSnapshotResponse

	if deleteSnapshotRequest == nil {
		deleteSnapshotRequest = new(DeleteSnapshotRequest)
	}
	err := c.DoAction(deleteSnapshotRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DeleteSecurityGroup(deleteSecurityGroupRequest *DeleteSecurityGroupRequest) (ecs_DeleteSecurityGroupResponse.DeleteSecurityGroupResponse, error) {
	var resObj ecs_DeleteSecurityGroupResponse.DeleteSecurityGroupResponse

	if deleteSecurityGroupRequest == nil {
		deleteSecurityGroupRequest = new(DeleteSecurityGroupRequest)
	}
	err := c.DoAction(deleteSecurityGroupRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DeleteRouteEntry(deleteRouteEntryRequest *DeleteRouteEntryRequest) (ecs_DeleteRouteEntryResponse.DeleteRouteEntryResponse, error) {
	var resObj ecs_DeleteRouteEntryResponse.DeleteRouteEntryResponse

	if deleteRouteEntryRequest == nil {
		deleteRouteEntryRequest = new(DeleteRouteEntryRequest)
	}
	err := c.DoAction(deleteRouteEntryRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DeleteInstance(deleteInstanceRequest *DeleteInstanceRequest) (ecs_DeleteInstanceResponse.DeleteInstanceResponse, error) {
	var resObj ecs_DeleteInstanceResponse.DeleteInstanceResponse

	if deleteInstanceRequest == nil {
		deleteInstanceRequest = new(DeleteInstanceRequest)
	}
	err := c.DoAction(deleteInstanceRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DeleteImage(deleteImageRequest *DeleteImageRequest) (ecs_DeleteImageResponse.DeleteImageResponse, error) {
	var resObj ecs_DeleteImageResponse.DeleteImageResponse

	if deleteImageRequest == nil {
		deleteImageRequest = new(DeleteImageRequest)
	}
	err := c.DoAction(deleteImageRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) DeleteDisk(deleteDiskRequest *DeleteDiskRequest) (ecs_DeleteDiskResponse.DeleteDiskResponse, error) {
	var resObj ecs_DeleteDiskResponse.DeleteDiskResponse

	if deleteDiskRequest == nil {
		deleteDiskRequest = new(DeleteDiskRequest)
	}
	err := c.DoAction(deleteDiskRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) CreateVSwitch(createVSwitchRequest *CreateVSwitchRequest) (ecs_CreateVSwitchResponse.CreateVSwitchResponse, error) {
	var resObj ecs_CreateVSwitchResponse.CreateVSwitchResponse

	if createVSwitchRequest == nil {
		createVSwitchRequest = new(CreateVSwitchRequest)
	}
	err := c.DoAction(createVSwitchRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) CreateVpc(createVpcRequest *CreateVpcRequest) (ecs_CreateVpcResponse.CreateVpcResponse, error) {
	var resObj ecs_CreateVpcResponse.CreateVpcResponse

	if createVpcRequest == nil {
		createVpcRequest = new(CreateVpcRequest)
	}
	err := c.DoAction(createVpcRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) CreateSnapshot(createSnapshotRequest *CreateSnapshotRequest) (ecs_CreateSnapshotResponse.CreateSnapshotResponse, error) {
	var resObj ecs_CreateSnapshotResponse.CreateSnapshotResponse

	if createSnapshotRequest == nil {
		createSnapshotRequest = new(CreateSnapshotRequest)
	}
	err := c.DoAction(createSnapshotRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) CreateSecurityGroup(createSecurityGroupRequest *CreateSecurityGroupRequest) (ecs_CreateSecurityGroupResponse.CreateSecurityGroupResponse, error) {
	var resObj ecs_CreateSecurityGroupResponse.CreateSecurityGroupResponse

	if createSecurityGroupRequest == nil {
		createSecurityGroupRequest = new(CreateSecurityGroupRequest)
	}
	err := c.DoAction(createSecurityGroupRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) CreateRouteEntry(createRouteEntryRequest *CreateRouteEntryRequest) (ecs_CreateRouteEntryResponse.CreateRouteEntryResponse, error) {
	var resObj ecs_CreateRouteEntryResponse.CreateRouteEntryResponse

	if createRouteEntryRequest == nil {
		createRouteEntryRequest = new(CreateRouteEntryRequest)
	}
	err := c.DoAction(createRouteEntryRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) CreateInstance(createInstanceRequest *CreateInstanceRequest) (ecs_CreateInstanceResponse.CreateInstanceResponse, error) {
	var resObj ecs_CreateInstanceResponse.CreateInstanceResponse

	if createInstanceRequest == nil {
		createInstanceRequest = new(CreateInstanceRequest)
	}
	err := c.DoAction(createInstanceRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) CreateImage(createImageRequest *CreateImageRequest) (ecs_CreateImageResponse.CreateImageResponse, error) {
	var resObj ecs_CreateImageResponse.CreateImageResponse

	if createImageRequest == nil {
		createImageRequest = new(CreateImageRequest)
	}
	err := c.DoAction(createImageRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) CreateDisk(createDiskRequest *CreateDiskRequest) (ecs_CreateDiskResponse.CreateDiskResponse, error) {
	var resObj ecs_CreateDiskResponse.CreateDiskResponse

	if createDiskRequest == nil {
		createDiskRequest = new(CreateDiskRequest)
	}
	err := c.DoAction(createDiskRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) CopyImage(copyImageRequest *CopyImageRequest) (ecs_CopyImageResponse.CopyImageResponse, error) {
	var resObj ecs_CopyImageResponse.CopyImageResponse

	if copyImageRequest == nil {
		copyImageRequest = new(CopyImageRequest)
	}
	err := c.DoAction(copyImageRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) CheckDiskEnableAutoSnapshotValidation(checkDiskEnableAutoSnapshotValidationRequest *CheckDiskEnableAutoSnapshotValidationRequest) (ecs_CheckDiskEnableAutoSnapshotValidationResponse.CheckDiskEnableAutoSnapshotValidationResponse, error) {
	var resObj ecs_CheckDiskEnableAutoSnapshotValidationResponse.CheckDiskEnableAutoSnapshotValidationResponse

	if checkDiskEnableAutoSnapshotValidationRequest == nil {
		checkDiskEnableAutoSnapshotValidationRequest = new(CheckDiskEnableAutoSnapshotValidationRequest)
	}
	err := c.DoAction(checkDiskEnableAutoSnapshotValidationRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) CheckAutoSnapshotPolicy(checkAutoSnapshotPolicyRequest *CheckAutoSnapshotPolicyRequest) (ecs_CheckAutoSnapshotPolicyResponse.CheckAutoSnapshotPolicyResponse, error) {
	var resObj ecs_CheckAutoSnapshotPolicyResponse.CheckAutoSnapshotPolicyResponse

	if checkAutoSnapshotPolicyRequest == nil {
		checkAutoSnapshotPolicyRequest = new(CheckAutoSnapshotPolicyRequest)
	}
	err := c.DoAction(checkAutoSnapshotPolicyRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) CancelCopyImage(cancelCopyImageRequest *CancelCopyImageRequest) (ecs_CancelCopyImageResponse.CancelCopyImageResponse, error) {
	var resObj ecs_CancelCopyImageResponse.CancelCopyImageResponse

	if cancelCopyImageRequest == nil {
		cancelCopyImageRequest = new(CancelCopyImageRequest)
	}
	err := c.DoAction(cancelCopyImageRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) AuthorizeSecurityGroupEgress(authorizeSecurityGroupEgressRequest *AuthorizeSecurityGroupEgressRequest) (ecs_AuthorizeSecurityGroupEgressResponse.AuthorizeSecurityGroupEgressResponse, error) {
	var resObj ecs_AuthorizeSecurityGroupEgressResponse.AuthorizeSecurityGroupEgressResponse

	if authorizeSecurityGroupEgressRequest == nil {
		authorizeSecurityGroupEgressRequest = new(AuthorizeSecurityGroupEgressRequest)
	}
	err := c.DoAction(authorizeSecurityGroupEgressRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) AuthorizeSecurityGroup(authorizeSecurityGroupRequest *AuthorizeSecurityGroupRequest) (ecs_AuthorizeSecurityGroupResponse.AuthorizeSecurityGroupResponse, error) {
	var resObj ecs_AuthorizeSecurityGroupResponse.AuthorizeSecurityGroupResponse

	if authorizeSecurityGroupRequest == nil {
		authorizeSecurityGroupRequest = new(AuthorizeSecurityGroupRequest)
	}
	err := c.DoAction(authorizeSecurityGroupRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) AttachDisk(attachDiskRequest *AttachDiskRequest) (ecs_AttachDiskResponse.AttachDiskResponse, error) {
	var resObj ecs_AttachDiskResponse.AttachDiskResponse

	if attachDiskRequest == nil {
		attachDiskRequest = new(AttachDiskRequest)
	}
	err := c.DoAction(attachDiskRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) AssociateEipAddress(associateEipAddressRequest *AssociateEipAddressRequest) (ecs_AssociateEipAddressResponse.AssociateEipAddressResponse, error) {
	var resObj ecs_AssociateEipAddressResponse.AssociateEipAddressResponse

	if associateEipAddressRequest == nil {
		associateEipAddressRequest = new(AssociateEipAddressRequest)
	}
	err := c.DoAction(associateEipAddressRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) AllocatePublicIpAddress(allocatePublicIpAddressRequest *AllocatePublicIpAddressRequest) (ecs_AllocatePublicIpAddressResponse.AllocatePublicIpAddressResponse, error) {
	var resObj ecs_AllocatePublicIpAddressResponse.AllocatePublicIpAddressResponse

	if allocatePublicIpAddressRequest == nil {
		allocatePublicIpAddressRequest = new(AllocatePublicIpAddressRequest)
	}
	err := c.DoAction(allocatePublicIpAddressRequest, &resObj)
	return resObj, err
}

func (c *EcsClient) AllocateEipAddress(allocateEipAddressRequest *AllocateEipAddressRequest) (ecs_AllocateEipAddressResponse.AllocateEipAddressResponse, error) {
	var resObj ecs_AllocateEipAddressResponse.AllocateEipAddressResponse

	if allocateEipAddressRequest == nil {
		allocateEipAddressRequest = new(AllocateEipAddressRequest)
	}
	err := c.DoAction(allocateEipAddressRequest, &resObj)
	return resObj, err
}
