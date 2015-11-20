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

	"aliyun-go-sdk-ossadmin/20130712/PutBucketStatusResponse"

	"aliyun-go-sdk-ossadmin/20130712/PutBucketPolicyResponse"

	"aliyun-go-sdk-ossadmin/20130712/PutBucketLimitResponse"

	"aliyun-go-sdk-ossadmin/20130712/GetBucketPolicyResponse"

	"aliyun-go-sdk-ossadmin/20130712/CreateOssInstanceResponse"
)

type OssadminClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *OssadminClient {
	p := &OssadminClient{}
	p.Version = "2013-07-12"
	p.ProductName = "OssAdmin"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *OssadminClient) PutBucketStatus(putBucketStatusRequest *PutBucketStatusRequest) (ossadmin_PutBucketStatusResponse.PutBucketStatusResponse, error) {
	var resObj ossadmin_PutBucketStatusResponse.PutBucketStatusResponse

	if putBucketStatusRequest == nil {
		putBucketStatusRequest = new(PutBucketStatusRequest)
	}
	err := c.DoAction(putBucketStatusRequest, &resObj)
	return resObj, err
}

func (c *OssadminClient) PutBucketPolicy(putBucketPolicyRequest *PutBucketPolicyRequest) (ossadmin_PutBucketPolicyResponse.PutBucketPolicyResponse, error) {
	var resObj ossadmin_PutBucketPolicyResponse.PutBucketPolicyResponse

	if putBucketPolicyRequest == nil {
		putBucketPolicyRequest = new(PutBucketPolicyRequest)
	}
	err := c.DoAction(putBucketPolicyRequest, &resObj)
	return resObj, err
}

func (c *OssadminClient) PutBucketLimit(putBucketLimitRequest *PutBucketLimitRequest) (ossadmin_PutBucketLimitResponse.PutBucketLimitResponse, error) {
	var resObj ossadmin_PutBucketLimitResponse.PutBucketLimitResponse

	if putBucketLimitRequest == nil {
		putBucketLimitRequest = new(PutBucketLimitRequest)
	}
	err := c.DoAction(putBucketLimitRequest, &resObj)
	return resObj, err
}

func (c *OssadminClient) GetBucketPolicy(getBucketPolicyRequest *GetBucketPolicyRequest) (ossadmin_GetBucketPolicyResponse.GetBucketPolicyResponse, error) {
	var resObj ossadmin_GetBucketPolicyResponse.GetBucketPolicyResponse

	if getBucketPolicyRequest == nil {
		getBucketPolicyRequest = new(GetBucketPolicyRequest)
	}
	err := c.DoAction(getBucketPolicyRequest, &resObj)
	return resObj, err
}

func (c *OssadminClient) CreateOssInstance(createOssInstanceRequest *CreateOssInstanceRequest) (ossadmin_CreateOssInstanceResponse.CreateOssInstanceResponse, error) {
	var resObj ossadmin_CreateOssInstanceResponse.CreateOssInstanceResponse

	if createOssInstanceRequest == nil {
		createOssInstanceRequest = new(CreateOssInstanceRequest)
	}
	err := c.DoAction(createOssInstanceRequest, &resObj)
	return resObj, err
}
