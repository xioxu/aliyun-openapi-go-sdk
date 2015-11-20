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

package ossadmin

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-ossadmin/20150520/GetImgVpcInfoResponse"

	"aliyun-go-sdk-ossadmin/20150520/DeleteImgVpcResponse"

	"aliyun-go-sdk-ossadmin/20150520/CreateImgVpcResponse"

	"aliyun-go-sdk-ossadmin/20150520/UnBindBucketVipResponse"

	"aliyun-go-sdk-ossadmin/20150520/GetOssVpcInfoResponse"

	"aliyun-go-sdk-ossadmin/20150520/GetBucketVipsResponse"

	"aliyun-go-sdk-ossadmin/20150520/DeleteOssVpcResponse"

	"aliyun-go-sdk-ossadmin/20150520/CreateOssVpcResponse"

	"aliyun-go-sdk-ossadmin/20150520/BindBucketVipResponse"
)

type OssadminClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *OssadminClient {
	p := &OssadminClient{}
	p.Version = "2015-05-20"
	p.ProductName = "OssAdmin"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *OssadminClient) GetImgVpcInfo(getImgVpcInfoRequest *GetImgVpcInfoRequest) (ossadmin_GetImgVpcInfoResponse.GetImgVpcInfoResponse, error) {
	var resObj ossadmin_GetImgVpcInfoResponse.GetImgVpcInfoResponse

	if getImgVpcInfoRequest == nil {
		getImgVpcInfoRequest = new(GetImgVpcInfoRequest)
	}
	err := c.DoAction(getImgVpcInfoRequest, &resObj)
	return resObj, err
}

func (c *OssadminClient) DeleteImgVpc(deleteImgVpcRequest *DeleteImgVpcRequest) (ossadmin_DeleteImgVpcResponse.DeleteImgVpcResponse, error) {
	var resObj ossadmin_DeleteImgVpcResponse.DeleteImgVpcResponse

	if deleteImgVpcRequest == nil {
		deleteImgVpcRequest = new(DeleteImgVpcRequest)
	}
	err := c.DoAction(deleteImgVpcRequest, &resObj)
	return resObj, err
}

func (c *OssadminClient) CreateImgVpc(createImgVpcRequest *CreateImgVpcRequest) (ossadmin_CreateImgVpcResponse.CreateImgVpcResponse, error) {
	var resObj ossadmin_CreateImgVpcResponse.CreateImgVpcResponse

	if createImgVpcRequest == nil {
		createImgVpcRequest = new(CreateImgVpcRequest)
	}
	err := c.DoAction(createImgVpcRequest, &resObj)
	return resObj, err
}

func (c *OssadminClient) UnBindBucketVip(unBindBucketVipRequest *UnBindBucketVipRequest) (ossadmin_UnBindBucketVipResponse.UnBindBucketVipResponse, error) {
	var resObj ossadmin_UnBindBucketVipResponse.UnBindBucketVipResponse

	if unBindBucketVipRequest == nil {
		unBindBucketVipRequest = new(UnBindBucketVipRequest)
	}
	err := c.DoAction(unBindBucketVipRequest, &resObj)
	return resObj, err
}

func (c *OssadminClient) GetOssVpcInfo(getOssVpcInfoRequest *GetOssVpcInfoRequest) (ossadmin_GetOssVpcInfoResponse.GetOssVpcInfoResponse, error) {
	var resObj ossadmin_GetOssVpcInfoResponse.GetOssVpcInfoResponse

	if getOssVpcInfoRequest == nil {
		getOssVpcInfoRequest = new(GetOssVpcInfoRequest)
	}
	err := c.DoAction(getOssVpcInfoRequest, &resObj)
	return resObj, err
}

func (c *OssadminClient) GetBucketVips(getBucketVipsRequest *GetBucketVipsRequest) (ossadmin_GetBucketVipsResponse.GetBucketVipsResponse, error) {
	var resObj ossadmin_GetBucketVipsResponse.GetBucketVipsResponse

	if getBucketVipsRequest == nil {
		getBucketVipsRequest = new(GetBucketVipsRequest)
	}
	err := c.DoAction(getBucketVipsRequest, &resObj)
	return resObj, err
}

func (c *OssadminClient) DeleteOssVpc(deleteOssVpcRequest *DeleteOssVpcRequest) (ossadmin_DeleteOssVpcResponse.DeleteOssVpcResponse, error) {
	var resObj ossadmin_DeleteOssVpcResponse.DeleteOssVpcResponse

	if deleteOssVpcRequest == nil {
		deleteOssVpcRequest = new(DeleteOssVpcRequest)
	}
	err := c.DoAction(deleteOssVpcRequest, &resObj)
	return resObj, err
}

func (c *OssadminClient) CreateOssVpc(createOssVpcRequest *CreateOssVpcRequest) (ossadmin_CreateOssVpcResponse.CreateOssVpcResponse, error) {
	var resObj ossadmin_CreateOssVpcResponse.CreateOssVpcResponse

	if createOssVpcRequest == nil {
		createOssVpcRequest = new(CreateOssVpcRequest)
	}
	err := c.DoAction(createOssVpcRequest, &resObj)
	return resObj, err
}

func (c *OssadminClient) BindBucketVip(bindBucketVipRequest *BindBucketVipRequest) (ossadmin_BindBucketVipResponse.BindBucketVipResponse, error) {
	var resObj ossadmin_BindBucketVipResponse.BindBucketVipResponse

	if bindBucketVipRequest == nil {
		bindBucketVipRequest = new(BindBucketVipRequest)
	}
	err := c.DoAction(bindBucketVipRequest, &resObj)
	return resObj, err
}
