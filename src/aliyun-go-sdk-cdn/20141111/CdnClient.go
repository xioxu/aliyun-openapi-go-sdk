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

package cdn

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-cdn/20141111/DescribeCdnDomainLogsResponse"

	"aliyun-go-sdk-cdn/20141111/DescribeCdnDomainDetailResponse"

	"aliyun-go-sdk-cdn/20141111/DescribeCdnDomainBaseDetailResponse"

	"aliyun-go-sdk-cdn/20141111/DeleteCdnDomainResponse"

	"aliyun-go-sdk-cdn/20141111/AddCdnDomainResponse"

	"aliyun-go-sdk-cdn/20141111/PushObjectCacheResponse"

	"aliyun-go-sdk-cdn/20141111/OpenCdnServiceResponse"

	"aliyun-go-sdk-cdn/20141111/ModifyCdnServiceResponse"

	"aliyun-go-sdk-cdn/20141111/DescribeUserDomainsResponse"

	"aliyun-go-sdk-cdn/20141111/DescribeRefreshTasksResponse"

	"aliyun-go-sdk-cdn/20141111/DescribeOneMinuteDataResponse"

	"aliyun-go-sdk-cdn/20141111/DescribeCdnServiceResponse"

	"aliyun-go-sdk-cdn/20141111/DescribeCdnMonitorDataResponse"

	"aliyun-go-sdk-cdn/20141111/StopCdnDomainResponse"

	"aliyun-go-sdk-cdn/20141111/StartCdnDomainResponse"

	"aliyun-go-sdk-cdn/20141111/RefreshObjectCachesResponse"
)

type CdnClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *CdnClient {
	p := &CdnClient{}
	p.Version = "2014-11-11"
	p.ProductName = "Cdn"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *CdnClient) DescribeCdnDomainLogs(describeCdnDomainLogsRequest *DescribeCdnDomainLogsRequest) (cdn_DescribeCdnDomainLogsResponse.DescribeCdnDomainLogsResponse, error) {
	var resObj cdn_DescribeCdnDomainLogsResponse.DescribeCdnDomainLogsResponse

	if describeCdnDomainLogsRequest == nil {
		describeCdnDomainLogsRequest = new(DescribeCdnDomainLogsRequest)
	}
	err := c.DoAction(describeCdnDomainLogsRequest, &resObj)
	return resObj, err
}

func (c *CdnClient) DescribeCdnDomainDetail(describeCdnDomainDetailRequest *DescribeCdnDomainDetailRequest) (cdn_DescribeCdnDomainDetailResponse.DescribeCdnDomainDetailResponse, error) {
	var resObj cdn_DescribeCdnDomainDetailResponse.DescribeCdnDomainDetailResponse

	if describeCdnDomainDetailRequest == nil {
		describeCdnDomainDetailRequest = new(DescribeCdnDomainDetailRequest)
	}
	err := c.DoAction(describeCdnDomainDetailRequest, &resObj)
	return resObj, err
}

func (c *CdnClient) DescribeCdnDomainBaseDetail(describeCdnDomainBaseDetailRequest *DescribeCdnDomainBaseDetailRequest) (cdn_DescribeCdnDomainBaseDetailResponse.DescribeCdnDomainBaseDetailResponse, error) {
	var resObj cdn_DescribeCdnDomainBaseDetailResponse.DescribeCdnDomainBaseDetailResponse

	if describeCdnDomainBaseDetailRequest == nil {
		describeCdnDomainBaseDetailRequest = new(DescribeCdnDomainBaseDetailRequest)
	}
	err := c.DoAction(describeCdnDomainBaseDetailRequest, &resObj)
	return resObj, err
}

func (c *CdnClient) DeleteCdnDomain(deleteCdnDomainRequest *DeleteCdnDomainRequest) (cdn_DeleteCdnDomainResponse.DeleteCdnDomainResponse, error) {
	var resObj cdn_DeleteCdnDomainResponse.DeleteCdnDomainResponse

	if deleteCdnDomainRequest == nil {
		deleteCdnDomainRequest = new(DeleteCdnDomainRequest)
	}
	err := c.DoAction(deleteCdnDomainRequest, &resObj)
	return resObj, err
}

func (c *CdnClient) AddCdnDomain(addCdnDomainRequest *AddCdnDomainRequest) (cdn_AddCdnDomainResponse.AddCdnDomainResponse, error) {
	var resObj cdn_AddCdnDomainResponse.AddCdnDomainResponse

	if addCdnDomainRequest == nil {
		addCdnDomainRequest = new(AddCdnDomainRequest)
	}
	err := c.DoAction(addCdnDomainRequest, &resObj)
	return resObj, err
}

func (c *CdnClient) PushObjectCache(pushObjectCacheRequest *PushObjectCacheRequest) (cdn_PushObjectCacheResponse.PushObjectCacheResponse, error) {
	var resObj cdn_PushObjectCacheResponse.PushObjectCacheResponse

	if pushObjectCacheRequest == nil {
		pushObjectCacheRequest = new(PushObjectCacheRequest)
	}
	err := c.DoAction(pushObjectCacheRequest, &resObj)
	return resObj, err
}

func (c *CdnClient) OpenCdnService(openCdnServiceRequest *OpenCdnServiceRequest) (cdn_OpenCdnServiceResponse.OpenCdnServiceResponse, error) {
	var resObj cdn_OpenCdnServiceResponse.OpenCdnServiceResponse

	if openCdnServiceRequest == nil {
		openCdnServiceRequest = new(OpenCdnServiceRequest)
	}
	err := c.DoAction(openCdnServiceRequest, &resObj)
	return resObj, err
}

func (c *CdnClient) ModifyCdnService(modifyCdnServiceRequest *ModifyCdnServiceRequest) (cdn_ModifyCdnServiceResponse.ModifyCdnServiceResponse, error) {
	var resObj cdn_ModifyCdnServiceResponse.ModifyCdnServiceResponse

	if modifyCdnServiceRequest == nil {
		modifyCdnServiceRequest = new(ModifyCdnServiceRequest)
	}
	err := c.DoAction(modifyCdnServiceRequest, &resObj)
	return resObj, err
}

func (c *CdnClient) DescribeUserDomains(describeUserDomainsRequest *DescribeUserDomainsRequest) (cdn_DescribeUserDomainsResponse.DescribeUserDomainsResponse, error) {
	var resObj cdn_DescribeUserDomainsResponse.DescribeUserDomainsResponse

	if describeUserDomainsRequest == nil {
		describeUserDomainsRequest = new(DescribeUserDomainsRequest)
	}
	err := c.DoAction(describeUserDomainsRequest, &resObj)
	return resObj, err
}

func (c *CdnClient) DescribeRefreshTasks(describeRefreshTasksRequest *DescribeRefreshTasksRequest) (cdn_DescribeRefreshTasksResponse.DescribeRefreshTasksResponse, error) {
	var resObj cdn_DescribeRefreshTasksResponse.DescribeRefreshTasksResponse

	if describeRefreshTasksRequest == nil {
		describeRefreshTasksRequest = new(DescribeRefreshTasksRequest)
	}
	err := c.DoAction(describeRefreshTasksRequest, &resObj)
	return resObj, err
}

func (c *CdnClient) DescribeOneMinuteData(describeOneMinuteDataRequest *DescribeOneMinuteDataRequest) (cdn_DescribeOneMinuteDataResponse.DescribeOneMinuteDataResponse, error) {
	var resObj cdn_DescribeOneMinuteDataResponse.DescribeOneMinuteDataResponse

	if describeOneMinuteDataRequest == nil {
		describeOneMinuteDataRequest = new(DescribeOneMinuteDataRequest)
	}
	err := c.DoAction(describeOneMinuteDataRequest, &resObj)
	return resObj, err
}

func (c *CdnClient) DescribeCdnService(describeCdnServiceRequest *DescribeCdnServiceRequest) (cdn_DescribeCdnServiceResponse.DescribeCdnServiceResponse, error) {
	var resObj cdn_DescribeCdnServiceResponse.DescribeCdnServiceResponse

	if describeCdnServiceRequest == nil {
		describeCdnServiceRequest = new(DescribeCdnServiceRequest)
	}
	err := c.DoAction(describeCdnServiceRequest, &resObj)
	return resObj, err
}

func (c *CdnClient) DescribeCdnMonitorData(describeCdnMonitorDataRequest *DescribeCdnMonitorDataRequest) (cdn_DescribeCdnMonitorDataResponse.DescribeCdnMonitorDataResponse, error) {
	var resObj cdn_DescribeCdnMonitorDataResponse.DescribeCdnMonitorDataResponse

	if describeCdnMonitorDataRequest == nil {
		describeCdnMonitorDataRequest = new(DescribeCdnMonitorDataRequest)
	}
	err := c.DoAction(describeCdnMonitorDataRequest, &resObj)
	return resObj, err
}

func (c *CdnClient) StopCdnDomain(stopCdnDomainRequest *StopCdnDomainRequest) (cdn_StopCdnDomainResponse.StopCdnDomainResponse, error) {
	var resObj cdn_StopCdnDomainResponse.StopCdnDomainResponse

	if stopCdnDomainRequest == nil {
		stopCdnDomainRequest = new(StopCdnDomainRequest)
	}
	err := c.DoAction(stopCdnDomainRequest, &resObj)
	return resObj, err
}

func (c *CdnClient) StartCdnDomain(startCdnDomainRequest *StartCdnDomainRequest) (cdn_StartCdnDomainResponse.StartCdnDomainResponse, error) {
	var resObj cdn_StartCdnDomainResponse.StartCdnDomainResponse

	if startCdnDomainRequest == nil {
		startCdnDomainRequest = new(StartCdnDomainRequest)
	}
	err := c.DoAction(startCdnDomainRequest, &resObj)
	return resObj, err
}

func (c *CdnClient) RefreshObjectCaches(refreshObjectCachesRequest *RefreshObjectCachesRequest) (cdn_RefreshObjectCachesResponse.RefreshObjectCachesResponse, error) {
	var resObj cdn_RefreshObjectCachesResponse.RefreshObjectCachesResponse

	if refreshObjectCachesRequest == nil {
		refreshObjectCachesRequest = new(RefreshObjectCachesRequest)
	}
	err := c.DoAction(refreshObjectCachesRequest, &resObj)
	return resObj, err
}
