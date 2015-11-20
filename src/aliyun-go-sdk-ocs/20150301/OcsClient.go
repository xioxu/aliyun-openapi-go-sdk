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

package ocs

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-ocs/20150301/DescribeZonesResponse"

	"aliyun-go-sdk-ocs/20150301/ActivateInstanceResponse"

	"aliyun-go-sdk-ocs/20150301/AddAuthenticIPResponse"

	"aliyun-go-sdk-ocs/20150301/CreateInstanceResponse"

	"aliyun-go-sdk-ocs/20150301/DeactivateInstanceResponse"

	"aliyun-go-sdk-ocs/20150301/DeleteInstanceResponse"

	"aliyun-go-sdk-ocs/20150301/DescribeInstancesResponse"

	"aliyun-go-sdk-ocs/20150301/DescribeHistoryMonitorValuesResponse"

	"aliyun-go-sdk-ocs/20150301/DescribeAuthenticIPResponse"

	"aliyun-go-sdk-ocs/20150301/DescribeSecurityIpsResponse"

	"aliyun-go-sdk-ocs/20150301/DescribeRegionsResponse"

	"aliyun-go-sdk-ocs/20150301/DescribeMonitorValuesResponse"

	"aliyun-go-sdk-ocs/20150301/DescribeMonitorItemsResponse"

	"aliyun-go-sdk-ocs/20150301/ModifySecurityIpsResponse"

	"aliyun-go-sdk-ocs/20150301/ModifyInstanceCapacityResponse"

	"aliyun-go-sdk-ocs/20150301/ModifyInstanceAttributeResponse"

	"aliyun-go-sdk-ocs/20150301/FlushInstanceResponse"

	"aliyun-go-sdk-ocs/20150301/VerifyPasswordResponse"

	"aliyun-go-sdk-ocs/20150301/ReplaceAuthenticIPResponse"

	"aliyun-go-sdk-ocs/20150301/RemoveAuthenticIPResponse"

	"aliyun-go-sdk-ocs/20150301/DataOperateResponse"
)

type OcsClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *OcsClient {
	p := &OcsClient{}
	p.Version = "2015-03-01"
	p.ProductName = "Ocs"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *OcsClient) DescribeZones(describeZonesRequest *DescribeZonesRequest) (ocs_DescribeZonesResponse.DescribeZonesResponse, error) {
	var resObj ocs_DescribeZonesResponse.DescribeZonesResponse

	if describeZonesRequest == nil {
		describeZonesRequest = new(DescribeZonesRequest)
	}
	err := c.DoAction(describeZonesRequest, &resObj)
	return resObj, err
}

func (c *OcsClient) ActivateInstance(activateInstanceRequest *ActivateInstanceRequest) (ocs_ActivateInstanceResponse.ActivateInstanceResponse, error) {
	var resObj ocs_ActivateInstanceResponse.ActivateInstanceResponse

	if activateInstanceRequest == nil {
		activateInstanceRequest = new(ActivateInstanceRequest)
	}
	err := c.DoAction(activateInstanceRequest, &resObj)
	return resObj, err
}

func (c *OcsClient) AddAuthenticIP(addAuthenticIPRequest *AddAuthenticIPRequest) (ocs_AddAuthenticIPResponse.AddAuthenticIPResponse, error) {
	var resObj ocs_AddAuthenticIPResponse.AddAuthenticIPResponse

	if addAuthenticIPRequest == nil {
		addAuthenticIPRequest = new(AddAuthenticIPRequest)
	}
	err := c.DoAction(addAuthenticIPRequest, &resObj)
	return resObj, err
}

func (c *OcsClient) CreateInstance(createInstanceRequest *CreateInstanceRequest) (ocs_CreateInstanceResponse.CreateInstanceResponse, error) {
	var resObj ocs_CreateInstanceResponse.CreateInstanceResponse

	if createInstanceRequest == nil {
		createInstanceRequest = new(CreateInstanceRequest)
	}
	err := c.DoAction(createInstanceRequest, &resObj)
	return resObj, err
}

func (c *OcsClient) DeactivateInstance(deactivateInstanceRequest *DeactivateInstanceRequest) (ocs_DeactivateInstanceResponse.DeactivateInstanceResponse, error) {
	var resObj ocs_DeactivateInstanceResponse.DeactivateInstanceResponse

	if deactivateInstanceRequest == nil {
		deactivateInstanceRequest = new(DeactivateInstanceRequest)
	}
	err := c.DoAction(deactivateInstanceRequest, &resObj)
	return resObj, err
}

func (c *OcsClient) DeleteInstance(deleteInstanceRequest *DeleteInstanceRequest) (ocs_DeleteInstanceResponse.DeleteInstanceResponse, error) {
	var resObj ocs_DeleteInstanceResponse.DeleteInstanceResponse

	if deleteInstanceRequest == nil {
		deleteInstanceRequest = new(DeleteInstanceRequest)
	}
	err := c.DoAction(deleteInstanceRequest, &resObj)
	return resObj, err
}

func (c *OcsClient) DescribeInstances(describeInstancesRequest *DescribeInstancesRequest) (ocs_DescribeInstancesResponse.DescribeInstancesResponse, error) {
	var resObj ocs_DescribeInstancesResponse.DescribeInstancesResponse

	if describeInstancesRequest == nil {
		describeInstancesRequest = new(DescribeInstancesRequest)
	}
	err := c.DoAction(describeInstancesRequest, &resObj)
	return resObj, err
}

func (c *OcsClient) DescribeHistoryMonitorValues(describeHistoryMonitorValuesRequest *DescribeHistoryMonitorValuesRequest) (ocs_DescribeHistoryMonitorValuesResponse.DescribeHistoryMonitorValuesResponse, error) {
	var resObj ocs_DescribeHistoryMonitorValuesResponse.DescribeHistoryMonitorValuesResponse

	if describeHistoryMonitorValuesRequest == nil {
		describeHistoryMonitorValuesRequest = new(DescribeHistoryMonitorValuesRequest)
	}
	err := c.DoAction(describeHistoryMonitorValuesRequest, &resObj)
	return resObj, err
}

func (c *OcsClient) DescribeAuthenticIP(describeAuthenticIPRequest *DescribeAuthenticIPRequest) (ocs_DescribeAuthenticIPResponse.DescribeAuthenticIPResponse, error) {
	var resObj ocs_DescribeAuthenticIPResponse.DescribeAuthenticIPResponse

	if describeAuthenticIPRequest == nil {
		describeAuthenticIPRequest = new(DescribeAuthenticIPRequest)
	}
	err := c.DoAction(describeAuthenticIPRequest, &resObj)
	return resObj, err
}

func (c *OcsClient) DescribeSecurityIps(describeSecurityIpsRequest *DescribeSecurityIpsRequest) (ocs_DescribeSecurityIpsResponse.DescribeSecurityIpsResponse, error) {
	var resObj ocs_DescribeSecurityIpsResponse.DescribeSecurityIpsResponse

	if describeSecurityIpsRequest == nil {
		describeSecurityIpsRequest = new(DescribeSecurityIpsRequest)
	}
	err := c.DoAction(describeSecurityIpsRequest, &resObj)
	return resObj, err
}

func (c *OcsClient) DescribeRegions(describeRegionsRequest *DescribeRegionsRequest) (ocs_DescribeRegionsResponse.DescribeRegionsResponse, error) {
	var resObj ocs_DescribeRegionsResponse.DescribeRegionsResponse

	if describeRegionsRequest == nil {
		describeRegionsRequest = new(DescribeRegionsRequest)
	}
	err := c.DoAction(describeRegionsRequest, &resObj)
	return resObj, err
}

func (c *OcsClient) DescribeMonitorValues(describeMonitorValuesRequest *DescribeMonitorValuesRequest) (ocs_DescribeMonitorValuesResponse.DescribeMonitorValuesResponse, error) {
	var resObj ocs_DescribeMonitorValuesResponse.DescribeMonitorValuesResponse

	if describeMonitorValuesRequest == nil {
		describeMonitorValuesRequest = new(DescribeMonitorValuesRequest)
	}
	err := c.DoAction(describeMonitorValuesRequest, &resObj)
	return resObj, err
}

func (c *OcsClient) DescribeMonitorItems(describeMonitorItemsRequest *DescribeMonitorItemsRequest) (ocs_DescribeMonitorItemsResponse.DescribeMonitorItemsResponse, error) {
	var resObj ocs_DescribeMonitorItemsResponse.DescribeMonitorItemsResponse

	if describeMonitorItemsRequest == nil {
		describeMonitorItemsRequest = new(DescribeMonitorItemsRequest)
	}
	err := c.DoAction(describeMonitorItemsRequest, &resObj)
	return resObj, err
}

func (c *OcsClient) ModifySecurityIps(modifySecurityIpsRequest *ModifySecurityIpsRequest) (ocs_ModifySecurityIpsResponse.ModifySecurityIpsResponse, error) {
	var resObj ocs_ModifySecurityIpsResponse.ModifySecurityIpsResponse

	if modifySecurityIpsRequest == nil {
		modifySecurityIpsRequest = new(ModifySecurityIpsRequest)
	}
	err := c.DoAction(modifySecurityIpsRequest, &resObj)
	return resObj, err
}

func (c *OcsClient) ModifyInstanceCapacity(modifyInstanceCapacityRequest *ModifyInstanceCapacityRequest) (ocs_ModifyInstanceCapacityResponse.ModifyInstanceCapacityResponse, error) {
	var resObj ocs_ModifyInstanceCapacityResponse.ModifyInstanceCapacityResponse

	if modifyInstanceCapacityRequest == nil {
		modifyInstanceCapacityRequest = new(ModifyInstanceCapacityRequest)
	}
	err := c.DoAction(modifyInstanceCapacityRequest, &resObj)
	return resObj, err
}

func (c *OcsClient) ModifyInstanceAttribute(modifyInstanceAttributeRequest *ModifyInstanceAttributeRequest) (ocs_ModifyInstanceAttributeResponse.ModifyInstanceAttributeResponse, error) {
	var resObj ocs_ModifyInstanceAttributeResponse.ModifyInstanceAttributeResponse

	if modifyInstanceAttributeRequest == nil {
		modifyInstanceAttributeRequest = new(ModifyInstanceAttributeRequest)
	}
	err := c.DoAction(modifyInstanceAttributeRequest, &resObj)
	return resObj, err
}

func (c *OcsClient) FlushInstance(flushInstanceRequest *FlushInstanceRequest) (ocs_FlushInstanceResponse.FlushInstanceResponse, error) {
	var resObj ocs_FlushInstanceResponse.FlushInstanceResponse

	if flushInstanceRequest == nil {
		flushInstanceRequest = new(FlushInstanceRequest)
	}
	err := c.DoAction(flushInstanceRequest, &resObj)
	return resObj, err
}

func (c *OcsClient) VerifyPassword(verifyPasswordRequest *VerifyPasswordRequest) (ocs_VerifyPasswordResponse.VerifyPasswordResponse, error) {
	var resObj ocs_VerifyPasswordResponse.VerifyPasswordResponse

	if verifyPasswordRequest == nil {
		verifyPasswordRequest = new(VerifyPasswordRequest)
	}
	err := c.DoAction(verifyPasswordRequest, &resObj)
	return resObj, err
}

func (c *OcsClient) ReplaceAuthenticIP(replaceAuthenticIPRequest *ReplaceAuthenticIPRequest) (ocs_ReplaceAuthenticIPResponse.ReplaceAuthenticIPResponse, error) {
	var resObj ocs_ReplaceAuthenticIPResponse.ReplaceAuthenticIPResponse

	if replaceAuthenticIPRequest == nil {
		replaceAuthenticIPRequest = new(ReplaceAuthenticIPRequest)
	}
	err := c.DoAction(replaceAuthenticIPRequest, &resObj)
	return resObj, err
}

func (c *OcsClient) RemoveAuthenticIP(removeAuthenticIPRequest *RemoveAuthenticIPRequest) (ocs_RemoveAuthenticIPResponse.RemoveAuthenticIPResponse, error) {
	var resObj ocs_RemoveAuthenticIPResponse.RemoveAuthenticIPResponse

	if removeAuthenticIPRequest == nil {
		removeAuthenticIPRequest = new(RemoveAuthenticIPRequest)
	}
	err := c.DoAction(removeAuthenticIPRequest, &resObj)
	return resObj, err
}

func (c *OcsClient) DataOperate(dataOperateRequest *DataOperateRequest) (ocs_DataOperateResponse.DataOperateResponse, error) {
	var resObj ocs_DataOperateResponse.DataOperateResponse

	if dataOperateRequest == nil {
		dataOperateRequest = new(DataOperateRequest)
	}
	err := c.DoAction(dataOperateRequest, &resObj)
	return resObj, err
}
