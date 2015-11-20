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

package slb

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-slb/20140515/DescribeLocationsResponse"

	"aliyun-go-sdk-slb/20140515/DescribeZonesResponse"

	"aliyun-go-sdk-slb/20140515/DescribeRegions4LocationResponse"

	"aliyun-go-sdk-slb/20140515/SetLoadBalancerUDPListenerAttributeResponse"

	"aliyun-go-sdk-slb/20140515/DescribeLoadBalancerUDPListenerAttributeResponse"

	"aliyun-go-sdk-slb/20140515/CreateLoadBalancerUDPListenerResponse"

	"aliyun-go-sdk-slb/20140515/UploadServerCertificateResponse"

	"aliyun-go-sdk-slb/20140515/StopLoadBalancerListenerResponse"

	"aliyun-go-sdk-slb/20140515/StartLoadBalancerListenerResponse"

	"aliyun-go-sdk-slb/20140515/SetServerCertificateNameResponse"

	"aliyun-go-sdk-slb/20140515/SetLoadBalancerTCPListenerAttributeResponse"

	"aliyun-go-sdk-slb/20140515/SetLoadBalancerStatusResponse"

	"aliyun-go-sdk-slb/20140515/SetLoadBalancerNameResponse"

	"aliyun-go-sdk-slb/20140515/SetLoadBalancerHTTPSListenerAttributeResponse"

	"aliyun-go-sdk-slb/20140515/SetLoadBalancerHTTPListenerAttributeResponse"

	"aliyun-go-sdk-slb/20140515/SetListenerAccessControlStatusResponse"

	"aliyun-go-sdk-slb/20140515/SetBackendServersResponse"

	"aliyun-go-sdk-slb/20140515/RemoveListenerWhiteListItemResponse"

	"aliyun-go-sdk-slb/20140515/RemoveBackendServersResponse"

	"aliyun-go-sdk-slb/20140515/ModifyLoadBalancerInternetSpecResponse"

	"aliyun-go-sdk-slb/20140515/DescribeServerCertificatesResponse"

	"aliyun-go-sdk-slb/20140515/DescribeServerCertificateResponse"

	"aliyun-go-sdk-slb/20140515/DescribeRegionsResponse"

	"aliyun-go-sdk-slb/20140515/DescribeLoadBalancerTCPListenerAttributeResponse"

	"aliyun-go-sdk-slb/20140515/DescribeLoadBalancersResponse"

	"aliyun-go-sdk-slb/20140515/DescribeLoadBalancerHTTPSListenerAttributeResponse"

	"aliyun-go-sdk-slb/20140515/DescribeLoadBalancerHTTPListenerAttributeResponse"

	"aliyun-go-sdk-slb/20140515/DescribeLoadBalancerAttributeResponse"

	"aliyun-go-sdk-slb/20140515/DescribeListenerAccessControlAttributeResponse"

	"aliyun-go-sdk-slb/20140515/DescribeHealthStatusResponse"

	"aliyun-go-sdk-slb/20140515/DeleteServerCertificateResponse"

	"aliyun-go-sdk-slb/20140515/DeleteLoadBalancerListenerResponse"

	"aliyun-go-sdk-slb/20140515/DeleteLoadBalancerResponse"

	"aliyun-go-sdk-slb/20140515/CreateLoadBalancerTCPListenerResponse"

	"aliyun-go-sdk-slb/20140515/CreateLoadBalancerHTTPSListenerResponse"

	"aliyun-go-sdk-slb/20140515/CreateLoadBalancerHTTPListenerResponse"

	"aliyun-go-sdk-slb/20140515/CreateLoadBalancerResponse"

	"aliyun-go-sdk-slb/20140515/AddListenerWhiteListItemResponse"

	"aliyun-go-sdk-slb/20140515/AddBackendServersResponse"
)

type SlbClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *SlbClient {
	p := &SlbClient{}
	p.Version = "2014-05-15"
	p.ProductName = "Slb"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *SlbClient) DescribeLocations(describeLocationsRequest *DescribeLocationsRequest) (slb_DescribeLocationsResponse.DescribeLocationsResponse, error) {
	var resObj slb_DescribeLocationsResponse.DescribeLocationsResponse

	if describeLocationsRequest == nil {
		describeLocationsRequest = new(DescribeLocationsRequest)
	}
	err := c.DoAction(describeLocationsRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) DescribeZones(describeZonesRequest *DescribeZonesRequest) (slb_DescribeZonesResponse.DescribeZonesResponse, error) {
	var resObj slb_DescribeZonesResponse.DescribeZonesResponse

	if describeZonesRequest == nil {
		describeZonesRequest = new(DescribeZonesRequest)
	}
	err := c.DoAction(describeZonesRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) DescribeRegions4Location(describeRegions4LocationRequest *DescribeRegions4LocationRequest) (slb_DescribeRegions4LocationResponse.DescribeRegions4LocationResponse, error) {
	var resObj slb_DescribeRegions4LocationResponse.DescribeRegions4LocationResponse

	if describeRegions4LocationRequest == nil {
		describeRegions4LocationRequest = new(DescribeRegions4LocationRequest)
	}
	err := c.DoAction(describeRegions4LocationRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) SetLoadBalancerUDPListenerAttribute(setLoadBalancerUDPListenerAttributeRequest *SetLoadBalancerUDPListenerAttributeRequest) (slb_SetLoadBalancerUDPListenerAttributeResponse.SetLoadBalancerUDPListenerAttributeResponse, error) {
	var resObj slb_SetLoadBalancerUDPListenerAttributeResponse.SetLoadBalancerUDPListenerAttributeResponse

	if setLoadBalancerUDPListenerAttributeRequest == nil {
		setLoadBalancerUDPListenerAttributeRequest = new(SetLoadBalancerUDPListenerAttributeRequest)
	}
	err := c.DoAction(setLoadBalancerUDPListenerAttributeRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) DescribeLoadBalancerUDPListenerAttribute(describeLoadBalancerUDPListenerAttributeRequest *DescribeLoadBalancerUDPListenerAttributeRequest) (slb_DescribeLoadBalancerUDPListenerAttributeResponse.DescribeLoadBalancerUDPListenerAttributeResponse, error) {
	var resObj slb_DescribeLoadBalancerUDPListenerAttributeResponse.DescribeLoadBalancerUDPListenerAttributeResponse

	if describeLoadBalancerUDPListenerAttributeRequest == nil {
		describeLoadBalancerUDPListenerAttributeRequest = new(DescribeLoadBalancerUDPListenerAttributeRequest)
	}
	err := c.DoAction(describeLoadBalancerUDPListenerAttributeRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) CreateLoadBalancerUDPListener(createLoadBalancerUDPListenerRequest *CreateLoadBalancerUDPListenerRequest) (slb_CreateLoadBalancerUDPListenerResponse.CreateLoadBalancerUDPListenerResponse, error) {
	var resObj slb_CreateLoadBalancerUDPListenerResponse.CreateLoadBalancerUDPListenerResponse

	if createLoadBalancerUDPListenerRequest == nil {
		createLoadBalancerUDPListenerRequest = new(CreateLoadBalancerUDPListenerRequest)
	}
	err := c.DoAction(createLoadBalancerUDPListenerRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) UploadServerCertificate(uploadServerCertificateRequest *UploadServerCertificateRequest) (slb_UploadServerCertificateResponse.UploadServerCertificateResponse, error) {
	var resObj slb_UploadServerCertificateResponse.UploadServerCertificateResponse

	if uploadServerCertificateRequest == nil {
		uploadServerCertificateRequest = new(UploadServerCertificateRequest)
	}
	err := c.DoAction(uploadServerCertificateRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) StopLoadBalancerListener(stopLoadBalancerListenerRequest *StopLoadBalancerListenerRequest) (slb_StopLoadBalancerListenerResponse.StopLoadBalancerListenerResponse, error) {
	var resObj slb_StopLoadBalancerListenerResponse.StopLoadBalancerListenerResponse

	if stopLoadBalancerListenerRequest == nil {
		stopLoadBalancerListenerRequest = new(StopLoadBalancerListenerRequest)
	}
	err := c.DoAction(stopLoadBalancerListenerRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) StartLoadBalancerListener(startLoadBalancerListenerRequest *StartLoadBalancerListenerRequest) (slb_StartLoadBalancerListenerResponse.StartLoadBalancerListenerResponse, error) {
	var resObj slb_StartLoadBalancerListenerResponse.StartLoadBalancerListenerResponse

	if startLoadBalancerListenerRequest == nil {
		startLoadBalancerListenerRequest = new(StartLoadBalancerListenerRequest)
	}
	err := c.DoAction(startLoadBalancerListenerRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) SetServerCertificateName(setServerCertificateNameRequest *SetServerCertificateNameRequest) (slb_SetServerCertificateNameResponse.SetServerCertificateNameResponse, error) {
	var resObj slb_SetServerCertificateNameResponse.SetServerCertificateNameResponse

	if setServerCertificateNameRequest == nil {
		setServerCertificateNameRequest = new(SetServerCertificateNameRequest)
	}
	err := c.DoAction(setServerCertificateNameRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) SetLoadBalancerTCPListenerAttribute(setLoadBalancerTCPListenerAttributeRequest *SetLoadBalancerTCPListenerAttributeRequest) (slb_SetLoadBalancerTCPListenerAttributeResponse.SetLoadBalancerTCPListenerAttributeResponse, error) {
	var resObj slb_SetLoadBalancerTCPListenerAttributeResponse.SetLoadBalancerTCPListenerAttributeResponse

	if setLoadBalancerTCPListenerAttributeRequest == nil {
		setLoadBalancerTCPListenerAttributeRequest = new(SetLoadBalancerTCPListenerAttributeRequest)
	}
	err := c.DoAction(setLoadBalancerTCPListenerAttributeRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) SetLoadBalancerStatus(setLoadBalancerStatusRequest *SetLoadBalancerStatusRequest) (slb_SetLoadBalancerStatusResponse.SetLoadBalancerStatusResponse, error) {
	var resObj slb_SetLoadBalancerStatusResponse.SetLoadBalancerStatusResponse

	if setLoadBalancerStatusRequest == nil {
		setLoadBalancerStatusRequest = new(SetLoadBalancerStatusRequest)
	}
	err := c.DoAction(setLoadBalancerStatusRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) SetLoadBalancerName(setLoadBalancerNameRequest *SetLoadBalancerNameRequest) (slb_SetLoadBalancerNameResponse.SetLoadBalancerNameResponse, error) {
	var resObj slb_SetLoadBalancerNameResponse.SetLoadBalancerNameResponse

	if setLoadBalancerNameRequest == nil {
		setLoadBalancerNameRequest = new(SetLoadBalancerNameRequest)
	}
	err := c.DoAction(setLoadBalancerNameRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) SetLoadBalancerHTTPSListenerAttribute(setLoadBalancerHTTPSListenerAttributeRequest *SetLoadBalancerHTTPSListenerAttributeRequest) (slb_SetLoadBalancerHTTPSListenerAttributeResponse.SetLoadBalancerHTTPSListenerAttributeResponse, error) {
	var resObj slb_SetLoadBalancerHTTPSListenerAttributeResponse.SetLoadBalancerHTTPSListenerAttributeResponse

	if setLoadBalancerHTTPSListenerAttributeRequest == nil {
		setLoadBalancerHTTPSListenerAttributeRequest = new(SetLoadBalancerHTTPSListenerAttributeRequest)
	}
	err := c.DoAction(setLoadBalancerHTTPSListenerAttributeRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) SetLoadBalancerHTTPListenerAttribute(setLoadBalancerHTTPListenerAttributeRequest *SetLoadBalancerHTTPListenerAttributeRequest) (slb_SetLoadBalancerHTTPListenerAttributeResponse.SetLoadBalancerHTTPListenerAttributeResponse, error) {
	var resObj slb_SetLoadBalancerHTTPListenerAttributeResponse.SetLoadBalancerHTTPListenerAttributeResponse

	if setLoadBalancerHTTPListenerAttributeRequest == nil {
		setLoadBalancerHTTPListenerAttributeRequest = new(SetLoadBalancerHTTPListenerAttributeRequest)
	}
	err := c.DoAction(setLoadBalancerHTTPListenerAttributeRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) SetListenerAccessControlStatus(setListenerAccessControlStatusRequest *SetListenerAccessControlStatusRequest) (slb_SetListenerAccessControlStatusResponse.SetListenerAccessControlStatusResponse, error) {
	var resObj slb_SetListenerAccessControlStatusResponse.SetListenerAccessControlStatusResponse

	if setListenerAccessControlStatusRequest == nil {
		setListenerAccessControlStatusRequest = new(SetListenerAccessControlStatusRequest)
	}
	err := c.DoAction(setListenerAccessControlStatusRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) SetBackendServers(setBackendServersRequest *SetBackendServersRequest) (slb_SetBackendServersResponse.SetBackendServersResponse, error) {
	var resObj slb_SetBackendServersResponse.SetBackendServersResponse

	if setBackendServersRequest == nil {
		setBackendServersRequest = new(SetBackendServersRequest)
	}
	err := c.DoAction(setBackendServersRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) RemoveListenerWhiteListItem(removeListenerWhiteListItemRequest *RemoveListenerWhiteListItemRequest) (slb_RemoveListenerWhiteListItemResponse.RemoveListenerWhiteListItemResponse, error) {
	var resObj slb_RemoveListenerWhiteListItemResponse.RemoveListenerWhiteListItemResponse

	if removeListenerWhiteListItemRequest == nil {
		removeListenerWhiteListItemRequest = new(RemoveListenerWhiteListItemRequest)
	}
	err := c.DoAction(removeListenerWhiteListItemRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) RemoveBackendServers(removeBackendServersRequest *RemoveBackendServersRequest) (slb_RemoveBackendServersResponse.RemoveBackendServersResponse, error) {
	var resObj slb_RemoveBackendServersResponse.RemoveBackendServersResponse

	if removeBackendServersRequest == nil {
		removeBackendServersRequest = new(RemoveBackendServersRequest)
	}
	err := c.DoAction(removeBackendServersRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) ModifyLoadBalancerInternetSpec(modifyLoadBalancerInternetSpecRequest *ModifyLoadBalancerInternetSpecRequest) (slb_ModifyLoadBalancerInternetSpecResponse.ModifyLoadBalancerInternetSpecResponse, error) {
	var resObj slb_ModifyLoadBalancerInternetSpecResponse.ModifyLoadBalancerInternetSpecResponse

	if modifyLoadBalancerInternetSpecRequest == nil {
		modifyLoadBalancerInternetSpecRequest = new(ModifyLoadBalancerInternetSpecRequest)
	}
	err := c.DoAction(modifyLoadBalancerInternetSpecRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) DescribeServerCertificates(describeServerCertificatesRequest *DescribeServerCertificatesRequest) (slb_DescribeServerCertificatesResponse.DescribeServerCertificatesResponse, error) {
	var resObj slb_DescribeServerCertificatesResponse.DescribeServerCertificatesResponse

	if describeServerCertificatesRequest == nil {
		describeServerCertificatesRequest = new(DescribeServerCertificatesRequest)
	}
	err := c.DoAction(describeServerCertificatesRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) DescribeServerCertificate(describeServerCertificateRequest *DescribeServerCertificateRequest) (slb_DescribeServerCertificateResponse.DescribeServerCertificateResponse, error) {
	var resObj slb_DescribeServerCertificateResponse.DescribeServerCertificateResponse

	if describeServerCertificateRequest == nil {
		describeServerCertificateRequest = new(DescribeServerCertificateRequest)
	}
	err := c.DoAction(describeServerCertificateRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) DescribeRegions(describeRegionsRequest *DescribeRegionsRequest) (slb_DescribeRegionsResponse.DescribeRegionsResponse, error) {
	var resObj slb_DescribeRegionsResponse.DescribeRegionsResponse

	if describeRegionsRequest == nil {
		describeRegionsRequest = new(DescribeRegionsRequest)
	}
	err := c.DoAction(describeRegionsRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) DescribeLoadBalancerTCPListenerAttribute(describeLoadBalancerTCPListenerAttributeRequest *DescribeLoadBalancerTCPListenerAttributeRequest) (slb_DescribeLoadBalancerTCPListenerAttributeResponse.DescribeLoadBalancerTCPListenerAttributeResponse, error) {
	var resObj slb_DescribeLoadBalancerTCPListenerAttributeResponse.DescribeLoadBalancerTCPListenerAttributeResponse

	if describeLoadBalancerTCPListenerAttributeRequest == nil {
		describeLoadBalancerTCPListenerAttributeRequest = new(DescribeLoadBalancerTCPListenerAttributeRequest)
	}
	err := c.DoAction(describeLoadBalancerTCPListenerAttributeRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) DescribeLoadBalancers(describeLoadBalancersRequest *DescribeLoadBalancersRequest) (slb_DescribeLoadBalancersResponse.DescribeLoadBalancersResponse, error) {
	var resObj slb_DescribeLoadBalancersResponse.DescribeLoadBalancersResponse

	if describeLoadBalancersRequest == nil {
		describeLoadBalancersRequest = new(DescribeLoadBalancersRequest)
	}
	err := c.DoAction(describeLoadBalancersRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) DescribeLoadBalancerHTTPSListenerAttribute(describeLoadBalancerHTTPSListenerAttributeRequest *DescribeLoadBalancerHTTPSListenerAttributeRequest) (slb_DescribeLoadBalancerHTTPSListenerAttributeResponse.DescribeLoadBalancerHTTPSListenerAttributeResponse, error) {
	var resObj slb_DescribeLoadBalancerHTTPSListenerAttributeResponse.DescribeLoadBalancerHTTPSListenerAttributeResponse

	if describeLoadBalancerHTTPSListenerAttributeRequest == nil {
		describeLoadBalancerHTTPSListenerAttributeRequest = new(DescribeLoadBalancerHTTPSListenerAttributeRequest)
	}
	err := c.DoAction(describeLoadBalancerHTTPSListenerAttributeRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) DescribeLoadBalancerHTTPListenerAttribute(describeLoadBalancerHTTPListenerAttributeRequest *DescribeLoadBalancerHTTPListenerAttributeRequest) (slb_DescribeLoadBalancerHTTPListenerAttributeResponse.DescribeLoadBalancerHTTPListenerAttributeResponse, error) {
	var resObj slb_DescribeLoadBalancerHTTPListenerAttributeResponse.DescribeLoadBalancerHTTPListenerAttributeResponse

	if describeLoadBalancerHTTPListenerAttributeRequest == nil {
		describeLoadBalancerHTTPListenerAttributeRequest = new(DescribeLoadBalancerHTTPListenerAttributeRequest)
	}
	err := c.DoAction(describeLoadBalancerHTTPListenerAttributeRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) DescribeLoadBalancerAttribute(describeLoadBalancerAttributeRequest *DescribeLoadBalancerAttributeRequest) (slb_DescribeLoadBalancerAttributeResponse.DescribeLoadBalancerAttributeResponse, error) {
	var resObj slb_DescribeLoadBalancerAttributeResponse.DescribeLoadBalancerAttributeResponse

	if describeLoadBalancerAttributeRequest == nil {
		describeLoadBalancerAttributeRequest = new(DescribeLoadBalancerAttributeRequest)
	}
	err := c.DoAction(describeLoadBalancerAttributeRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) DescribeListenerAccessControlAttribute(describeListenerAccessControlAttributeRequest *DescribeListenerAccessControlAttributeRequest) (slb_DescribeListenerAccessControlAttributeResponse.DescribeListenerAccessControlAttributeResponse, error) {
	var resObj slb_DescribeListenerAccessControlAttributeResponse.DescribeListenerAccessControlAttributeResponse

	if describeListenerAccessControlAttributeRequest == nil {
		describeListenerAccessControlAttributeRequest = new(DescribeListenerAccessControlAttributeRequest)
	}
	err := c.DoAction(describeListenerAccessControlAttributeRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) DescribeHealthStatus(describeHealthStatusRequest *DescribeHealthStatusRequest) (slb_DescribeHealthStatusResponse.DescribeHealthStatusResponse, error) {
	var resObj slb_DescribeHealthStatusResponse.DescribeHealthStatusResponse

	if describeHealthStatusRequest == nil {
		describeHealthStatusRequest = new(DescribeHealthStatusRequest)
	}
	err := c.DoAction(describeHealthStatusRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) DeleteServerCertificate(deleteServerCertificateRequest *DeleteServerCertificateRequest) (slb_DeleteServerCertificateResponse.DeleteServerCertificateResponse, error) {
	var resObj slb_DeleteServerCertificateResponse.DeleteServerCertificateResponse

	if deleteServerCertificateRequest == nil {
		deleteServerCertificateRequest = new(DeleteServerCertificateRequest)
	}
	err := c.DoAction(deleteServerCertificateRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) DeleteLoadBalancerListener(deleteLoadBalancerListenerRequest *DeleteLoadBalancerListenerRequest) (slb_DeleteLoadBalancerListenerResponse.DeleteLoadBalancerListenerResponse, error) {
	var resObj slb_DeleteLoadBalancerListenerResponse.DeleteLoadBalancerListenerResponse

	if deleteLoadBalancerListenerRequest == nil {
		deleteLoadBalancerListenerRequest = new(DeleteLoadBalancerListenerRequest)
	}
	err := c.DoAction(deleteLoadBalancerListenerRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) DeleteLoadBalancer(deleteLoadBalancerRequest *DeleteLoadBalancerRequest) (slb_DeleteLoadBalancerResponse.DeleteLoadBalancerResponse, error) {
	var resObj slb_DeleteLoadBalancerResponse.DeleteLoadBalancerResponse

	if deleteLoadBalancerRequest == nil {
		deleteLoadBalancerRequest = new(DeleteLoadBalancerRequest)
	}
	err := c.DoAction(deleteLoadBalancerRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) CreateLoadBalancerTCPListener(createLoadBalancerTCPListenerRequest *CreateLoadBalancerTCPListenerRequest) (slb_CreateLoadBalancerTCPListenerResponse.CreateLoadBalancerTCPListenerResponse, error) {
	var resObj slb_CreateLoadBalancerTCPListenerResponse.CreateLoadBalancerTCPListenerResponse

	if createLoadBalancerTCPListenerRequest == nil {
		createLoadBalancerTCPListenerRequest = new(CreateLoadBalancerTCPListenerRequest)
	}
	err := c.DoAction(createLoadBalancerTCPListenerRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) CreateLoadBalancerHTTPSListener(createLoadBalancerHTTPSListenerRequest *CreateLoadBalancerHTTPSListenerRequest) (slb_CreateLoadBalancerHTTPSListenerResponse.CreateLoadBalancerHTTPSListenerResponse, error) {
	var resObj slb_CreateLoadBalancerHTTPSListenerResponse.CreateLoadBalancerHTTPSListenerResponse

	if createLoadBalancerHTTPSListenerRequest == nil {
		createLoadBalancerHTTPSListenerRequest = new(CreateLoadBalancerHTTPSListenerRequest)
	}
	err := c.DoAction(createLoadBalancerHTTPSListenerRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) CreateLoadBalancerHTTPListener(createLoadBalancerHTTPListenerRequest *CreateLoadBalancerHTTPListenerRequest) (slb_CreateLoadBalancerHTTPListenerResponse.CreateLoadBalancerHTTPListenerResponse, error) {
	var resObj slb_CreateLoadBalancerHTTPListenerResponse.CreateLoadBalancerHTTPListenerResponse

	if createLoadBalancerHTTPListenerRequest == nil {
		createLoadBalancerHTTPListenerRequest = new(CreateLoadBalancerHTTPListenerRequest)
	}
	err := c.DoAction(createLoadBalancerHTTPListenerRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) CreateLoadBalancer(createLoadBalancerRequest *CreateLoadBalancerRequest) (slb_CreateLoadBalancerResponse.CreateLoadBalancerResponse, error) {
	var resObj slb_CreateLoadBalancerResponse.CreateLoadBalancerResponse

	if createLoadBalancerRequest == nil {
		createLoadBalancerRequest = new(CreateLoadBalancerRequest)
	}
	err := c.DoAction(createLoadBalancerRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) AddListenerWhiteListItem(addListenerWhiteListItemRequest *AddListenerWhiteListItemRequest) (slb_AddListenerWhiteListItemResponse.AddListenerWhiteListItemResponse, error) {
	var resObj slb_AddListenerWhiteListItemResponse.AddListenerWhiteListItemResponse

	if addListenerWhiteListItemRequest == nil {
		addListenerWhiteListItemRequest = new(AddListenerWhiteListItemRequest)
	}
	err := c.DoAction(addListenerWhiteListItemRequest, &resObj)
	return resObj, err
}

func (c *SlbClient) AddBackendServers(addBackendServersRequest *AddBackendServersRequest) (slb_AddBackendServersResponse.AddBackendServersResponse, error) {
	var resObj slb_AddBackendServersResponse.AddBackendServersResponse

	if addBackendServersRequest == nil {
		addBackendServersRequest = new(AddBackendServersRequest)
	}
	err := c.DoAction(addBackendServersRequest, &resObj)
	return resObj, err
}
