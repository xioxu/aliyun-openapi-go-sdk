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

package kvstore

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-r-kvstore/20150101/TransformToPrePaidResponse"

	"aliyun-go-sdk-r-kvstore/20150101/RenewInstanceResponse"

	"aliyun-go-sdk-r-kvstore/20150101/DescribeUserInfoResponse"

	"aliyun-go-sdk-r-kvstore/20150101/DescribePriceResponse"

	"aliyun-go-sdk-r-kvstore/20150101/DescribeCommodityResponse"

	"aliyun-go-sdk-r-kvstore/20150101/CreateInstancesResponse"

	"aliyun-go-sdk-r-kvstore/20150101/DescribeInstanceCountResponse"

	"aliyun-go-sdk-r-kvstore/20150101/DescribeConnectionDomainResponse"

	"aliyun-go-sdk-r-kvstore/20150101/VerifyPasswordResponse"

	"aliyun-go-sdk-r-kvstore/20150101/ModifyInstanceConfigResponse"

	"aliyun-go-sdk-r-kvstore/20150101/ModifyInstanceCapacityResponse"

	"aliyun-go-sdk-r-kvstore/20150101/ModifyInstanceAttributeResponse"

	"aliyun-go-sdk-r-kvstore/20150101/FlushInstanceResponse"

	"aliyun-go-sdk-r-kvstore/20150101/DescribeRegionsResponse"

	"aliyun-go-sdk-r-kvstore/20150101/DescribeMonitorValuesResponse"

	"aliyun-go-sdk-r-kvstore/20150101/DescribeMonitorItemsResponse"

	"aliyun-go-sdk-r-kvstore/20150101/DescribeInstancesResponse"

	"aliyun-go-sdk-r-kvstore/20150101/DescribeInstanceConfigResponse"

	"aliyun-go-sdk-r-kvstore/20150101/DescribeHistoryMonitorValuesResponse"

	"aliyun-go-sdk-r-kvstore/20150101/DeleteInstanceResponse"

	"aliyun-go-sdk-r-kvstore/20150101/DeactivateInstanceResponse"

	"aliyun-go-sdk-r-kvstore/20150101/DataOperateResponse"

	"aliyun-go-sdk-r-kvstore/20150101/CreateInstanceResponse"

	"aliyun-go-sdk-r-kvstore/20150101/ActivateInstanceResponse"
)

type KvstoreClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *KvstoreClient {
	p := &KvstoreClient{}
	p.Version = "2015-01-01"
	p.ProductName = "R-kvstore"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *KvstoreClient) TransformToPrePaid(transformToPrePaidRequest *TransformToPrePaidRequest) (kvstore_TransformToPrePaidResponse.TransformToPrePaidResponse, error) {
	var resObj kvstore_TransformToPrePaidResponse.TransformToPrePaidResponse

	if transformToPrePaidRequest == nil {
		transformToPrePaidRequest = new(TransformToPrePaidRequest)
	}
	err := c.DoAction(transformToPrePaidRequest, &resObj)
	return resObj, err
}

func (c *KvstoreClient) RenewInstance(renewInstanceRequest *RenewInstanceRequest) (kvstore_RenewInstanceResponse.RenewInstanceResponse, error) {
	var resObj kvstore_RenewInstanceResponse.RenewInstanceResponse

	if renewInstanceRequest == nil {
		renewInstanceRequest = new(RenewInstanceRequest)
	}
	err := c.DoAction(renewInstanceRequest, &resObj)
	return resObj, err
}

func (c *KvstoreClient) DescribeUserInfo(describeUserInfoRequest *DescribeUserInfoRequest) (kvstore_DescribeUserInfoResponse.DescribeUserInfoResponse, error) {
	var resObj kvstore_DescribeUserInfoResponse.DescribeUserInfoResponse

	if describeUserInfoRequest == nil {
		describeUserInfoRequest = new(DescribeUserInfoRequest)
	}
	err := c.DoAction(describeUserInfoRequest, &resObj)
	return resObj, err
}

func (c *KvstoreClient) DescribePrice(describePriceRequest *DescribePriceRequest) (kvstore_DescribePriceResponse.DescribePriceResponse, error) {
	var resObj kvstore_DescribePriceResponse.DescribePriceResponse

	if describePriceRequest == nil {
		describePriceRequest = new(DescribePriceRequest)
	}
	err := c.DoAction(describePriceRequest, &resObj)
	return resObj, err
}

func (c *KvstoreClient) DescribeCommodity(describeCommodityRequest *DescribeCommodityRequest) (kvstore_DescribeCommodityResponse.DescribeCommodityResponse, error) {
	var resObj kvstore_DescribeCommodityResponse.DescribeCommodityResponse

	if describeCommodityRequest == nil {
		describeCommodityRequest = new(DescribeCommodityRequest)
	}
	err := c.DoAction(describeCommodityRequest, &resObj)
	return resObj, err
}

func (c *KvstoreClient) CreateInstances(createInstancesRequest *CreateInstancesRequest) (kvstore_CreateInstancesResponse.CreateInstancesResponse, error) {
	var resObj kvstore_CreateInstancesResponse.CreateInstancesResponse

	if createInstancesRequest == nil {
		createInstancesRequest = new(CreateInstancesRequest)
	}
	err := c.DoAction(createInstancesRequest, &resObj)
	return resObj, err
}

func (c *KvstoreClient) DescribeInstanceCount(describeInstanceCountRequest *DescribeInstanceCountRequest) (kvstore_DescribeInstanceCountResponse.DescribeInstanceCountResponse, error) {
	var resObj kvstore_DescribeInstanceCountResponse.DescribeInstanceCountResponse

	if describeInstanceCountRequest == nil {
		describeInstanceCountRequest = new(DescribeInstanceCountRequest)
	}
	err := c.DoAction(describeInstanceCountRequest, &resObj)
	return resObj, err
}

func (c *KvstoreClient) DescribeConnectionDomain(describeConnectionDomainRequest *DescribeConnectionDomainRequest) (kvstore_DescribeConnectionDomainResponse.DescribeConnectionDomainResponse, error) {
	var resObj kvstore_DescribeConnectionDomainResponse.DescribeConnectionDomainResponse

	if describeConnectionDomainRequest == nil {
		describeConnectionDomainRequest = new(DescribeConnectionDomainRequest)
	}
	err := c.DoAction(describeConnectionDomainRequest, &resObj)
	return resObj, err
}

func (c *KvstoreClient) VerifyPassword(verifyPasswordRequest *VerifyPasswordRequest) (kvstore_VerifyPasswordResponse.VerifyPasswordResponse, error) {
	var resObj kvstore_VerifyPasswordResponse.VerifyPasswordResponse

	if verifyPasswordRequest == nil {
		verifyPasswordRequest = new(VerifyPasswordRequest)
	}
	err := c.DoAction(verifyPasswordRequest, &resObj)
	return resObj, err
}

func (c *KvstoreClient) ModifyInstanceConfig(modifyInstanceConfigRequest *ModifyInstanceConfigRequest) (kvstore_ModifyInstanceConfigResponse.ModifyInstanceConfigResponse, error) {
	var resObj kvstore_ModifyInstanceConfigResponse.ModifyInstanceConfigResponse

	if modifyInstanceConfigRequest == nil {
		modifyInstanceConfigRequest = new(ModifyInstanceConfigRequest)
	}
	err := c.DoAction(modifyInstanceConfigRequest, &resObj)
	return resObj, err
}

func (c *KvstoreClient) ModifyInstanceCapacity(modifyInstanceCapacityRequest *ModifyInstanceCapacityRequest) (kvstore_ModifyInstanceCapacityResponse.ModifyInstanceCapacityResponse, error) {
	var resObj kvstore_ModifyInstanceCapacityResponse.ModifyInstanceCapacityResponse

	if modifyInstanceCapacityRequest == nil {
		modifyInstanceCapacityRequest = new(ModifyInstanceCapacityRequest)
	}
	err := c.DoAction(modifyInstanceCapacityRequest, &resObj)
	return resObj, err
}

func (c *KvstoreClient) ModifyInstanceAttribute(modifyInstanceAttributeRequest *ModifyInstanceAttributeRequest) (kvstore_ModifyInstanceAttributeResponse.ModifyInstanceAttributeResponse, error) {
	var resObj kvstore_ModifyInstanceAttributeResponse.ModifyInstanceAttributeResponse

	if modifyInstanceAttributeRequest == nil {
		modifyInstanceAttributeRequest = new(ModifyInstanceAttributeRequest)
	}
	err := c.DoAction(modifyInstanceAttributeRequest, &resObj)
	return resObj, err
}

func (c *KvstoreClient) FlushInstance(flushInstanceRequest *FlushInstanceRequest) (kvstore_FlushInstanceResponse.FlushInstanceResponse, error) {
	var resObj kvstore_FlushInstanceResponse.FlushInstanceResponse

	if flushInstanceRequest == nil {
		flushInstanceRequest = new(FlushInstanceRequest)
	}
	err := c.DoAction(flushInstanceRequest, &resObj)
	return resObj, err
}

func (c *KvstoreClient) DescribeRegions(describeRegionsRequest *DescribeRegionsRequest) (kvstore_DescribeRegionsResponse.DescribeRegionsResponse, error) {
	var resObj kvstore_DescribeRegionsResponse.DescribeRegionsResponse

	if describeRegionsRequest == nil {
		describeRegionsRequest = new(DescribeRegionsRequest)
	}
	err := c.DoAction(describeRegionsRequest, &resObj)
	return resObj, err
}

func (c *KvstoreClient) DescribeMonitorValues(describeMonitorValuesRequest *DescribeMonitorValuesRequest) (kvstore_DescribeMonitorValuesResponse.DescribeMonitorValuesResponse, error) {
	var resObj kvstore_DescribeMonitorValuesResponse.DescribeMonitorValuesResponse

	if describeMonitorValuesRequest == nil {
		describeMonitorValuesRequest = new(DescribeMonitorValuesRequest)
	}
	err := c.DoAction(describeMonitorValuesRequest, &resObj)
	return resObj, err
}

func (c *KvstoreClient) DescribeMonitorItems(describeMonitorItemsRequest *DescribeMonitorItemsRequest) (kvstore_DescribeMonitorItemsResponse.DescribeMonitorItemsResponse, error) {
	var resObj kvstore_DescribeMonitorItemsResponse.DescribeMonitorItemsResponse

	if describeMonitorItemsRequest == nil {
		describeMonitorItemsRequest = new(DescribeMonitorItemsRequest)
	}
	err := c.DoAction(describeMonitorItemsRequest, &resObj)
	return resObj, err
}

func (c *KvstoreClient) DescribeInstances(describeInstancesRequest *DescribeInstancesRequest) (kvstore_DescribeInstancesResponse.DescribeInstancesResponse, error) {
	var resObj kvstore_DescribeInstancesResponse.DescribeInstancesResponse

	if describeInstancesRequest == nil {
		describeInstancesRequest = new(DescribeInstancesRequest)
	}
	err := c.DoAction(describeInstancesRequest, &resObj)
	return resObj, err
}

func (c *KvstoreClient) DescribeInstanceConfig(describeInstanceConfigRequest *DescribeInstanceConfigRequest) (kvstore_DescribeInstanceConfigResponse.DescribeInstanceConfigResponse, error) {
	var resObj kvstore_DescribeInstanceConfigResponse.DescribeInstanceConfigResponse

	if describeInstanceConfigRequest == nil {
		describeInstanceConfigRequest = new(DescribeInstanceConfigRequest)
	}
	err := c.DoAction(describeInstanceConfigRequest, &resObj)
	return resObj, err
}

func (c *KvstoreClient) DescribeHistoryMonitorValues(describeHistoryMonitorValuesRequest *DescribeHistoryMonitorValuesRequest) (kvstore_DescribeHistoryMonitorValuesResponse.DescribeHistoryMonitorValuesResponse, error) {
	var resObj kvstore_DescribeHistoryMonitorValuesResponse.DescribeHistoryMonitorValuesResponse

	if describeHistoryMonitorValuesRequest == nil {
		describeHistoryMonitorValuesRequest = new(DescribeHistoryMonitorValuesRequest)
	}
	err := c.DoAction(describeHistoryMonitorValuesRequest, &resObj)
	return resObj, err
}

func (c *KvstoreClient) DeleteInstance(deleteInstanceRequest *DeleteInstanceRequest) (kvstore_DeleteInstanceResponse.DeleteInstanceResponse, error) {
	var resObj kvstore_DeleteInstanceResponse.DeleteInstanceResponse

	if deleteInstanceRequest == nil {
		deleteInstanceRequest = new(DeleteInstanceRequest)
	}
	err := c.DoAction(deleteInstanceRequest, &resObj)
	return resObj, err
}

func (c *KvstoreClient) DeactivateInstance(deactivateInstanceRequest *DeactivateInstanceRequest) (kvstore_DeactivateInstanceResponse.DeactivateInstanceResponse, error) {
	var resObj kvstore_DeactivateInstanceResponse.DeactivateInstanceResponse

	if deactivateInstanceRequest == nil {
		deactivateInstanceRequest = new(DeactivateInstanceRequest)
	}
	err := c.DoAction(deactivateInstanceRequest, &resObj)
	return resObj, err
}

func (c *KvstoreClient) DataOperate(dataOperateRequest *DataOperateRequest) (kvstore_DataOperateResponse.DataOperateResponse, error) {
	var resObj kvstore_DataOperateResponse.DataOperateResponse

	if dataOperateRequest == nil {
		dataOperateRequest = new(DataOperateRequest)
	}
	err := c.DoAction(dataOperateRequest, &resObj)
	return resObj, err
}

func (c *KvstoreClient) CreateInstance(createInstanceRequest *CreateInstanceRequest) (kvstore_CreateInstanceResponse.CreateInstanceResponse, error) {
	var resObj kvstore_CreateInstanceResponse.CreateInstanceResponse

	if createInstanceRequest == nil {
		createInstanceRequest = new(CreateInstanceRequest)
	}
	err := c.DoAction(createInstanceRequest, &resObj)
	return resObj, err
}

func (c *KvstoreClient) ActivateInstance(activateInstanceRequest *ActivateInstanceRequest) (kvstore_ActivateInstanceResponse.ActivateInstanceResponse, error) {
	var resObj kvstore_ActivateInstanceResponse.ActivateInstanceResponse

	if activateInstanceRequest == nil {
		activateInstanceRequest = new(ActivateInstanceRequest)
	}
	err := c.DoAction(activateInstanceRequest, &resObj)
	return resObj, err
}
