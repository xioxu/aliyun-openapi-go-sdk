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

package yundun

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-yundun/20150416/VulScanLogResponse"

	"aliyun-go-sdk-yundun/20150416/WafInfoResponse"

	"aliyun-go-sdk-yundun/20150416/SummaryResponse"

	"aliyun-go-sdk-yundun/20150416/SetDdosQpsResponse"

	"aliyun-go-sdk-yundun/20150416/SetDdosAutoResponse"

	"aliyun-go-sdk-yundun/20150416/ServiceStatusResponse"

	"aliyun-go-sdk-yundun/20150416/QueryDdosConfigResponse"

	"aliyun-go-sdk-yundun/20150416/OpenVulScanResponse"

	"aliyun-go-sdk-yundun/20150416/OpenPortScanResponse"

	"aliyun-go-sdk-yundun/20150416/OpenCCProtectResponse"

	"aliyun-go-sdk-yundun/20150416/GetDdosConfigOptionsResponse"

	"aliyun-go-sdk-yundun/20150416/DetectVulByIpResponse"

	"aliyun-go-sdk-yundun/20150416/DetectVulByIdResponse"

	"aliyun-go-sdk-yundun/20150416/DeleteCNameWafResponse"

	"aliyun-go-sdk-yundun/20150416/DdosFlowGraphResponse"

	"aliyun-go-sdk-yundun/20150416/ConfirmLoginResponse"

	"aliyun-go-sdk-yundun/20150416/ConfigDdosResponse"

	"aliyun-go-sdk-yundun/20150416/CloseVulScanResponse"

	"aliyun-go-sdk-yundun/20150416/ClosePortScanResponse"

	"aliyun-go-sdk-yundun/20150416/CloseCCProtectResponse"

	"aliyun-go-sdk-yundun/20150416/AddCNameWafResponse"

	"aliyun-go-sdk-yundun/20150416/WebshellLogResponse"

	"aliyun-go-sdk-yundun/20150416/WafLogResponse"

	"aliyun-go-sdk-yundun/20150416/SecureCheckResponse"

	"aliyun-go-sdk-yundun/20150416/LogineventLogResponse"

	"aliyun-go-sdk-yundun/20150416/ListInstanceInfosResponse"

	"aliyun-go-sdk-yundun/20150416/DeleteBackDoorFileResponse"

	"aliyun-go-sdk-yundun/20150416/DdosLogResponse"

	"aliyun-go-sdk-yundun/20150416/BruteforceLogResponse"
)

type YundunClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *YundunClient {
	p := &YundunClient{}
	p.Version = "2015-04-16"
	p.ProductName = "Yundun"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *YundunClient) VulScanLog(vulScanLogRequest *VulScanLogRequest) (yundun_VulScanLogResponse.VulScanLogResponse, error) {
	var resObj yundun_VulScanLogResponse.VulScanLogResponse

	if vulScanLogRequest == nil {
		vulScanLogRequest = new(VulScanLogRequest)
	}
	err := c.DoAction(vulScanLogRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) WafInfo(wafInfoRequest *WafInfoRequest) (yundun_WafInfoResponse.WafInfoResponse, error) {
	var resObj yundun_WafInfoResponse.WafInfoResponse

	if wafInfoRequest == nil {
		wafInfoRequest = new(WafInfoRequest)
	}
	err := c.DoAction(wafInfoRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) Summary(summaryRequest *SummaryRequest) (yundun_SummaryResponse.SummaryResponse, error) {
	var resObj yundun_SummaryResponse.SummaryResponse

	if summaryRequest == nil {
		summaryRequest = new(SummaryRequest)
	}
	err := c.DoAction(summaryRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) SetDdosQps(setDdosQpsRequest *SetDdosQpsRequest) (yundun_SetDdosQpsResponse.SetDdosQpsResponse, error) {
	var resObj yundun_SetDdosQpsResponse.SetDdosQpsResponse

	if setDdosQpsRequest == nil {
		setDdosQpsRequest = new(SetDdosQpsRequest)
	}
	err := c.DoAction(setDdosQpsRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) SetDdosAuto(setDdosAutoRequest *SetDdosAutoRequest) (yundun_SetDdosAutoResponse.SetDdosAutoResponse, error) {
	var resObj yundun_SetDdosAutoResponse.SetDdosAutoResponse

	if setDdosAutoRequest == nil {
		setDdosAutoRequest = new(SetDdosAutoRequest)
	}
	err := c.DoAction(setDdosAutoRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) ServiceStatus(serviceStatusRequest *ServiceStatusRequest) (yundun_ServiceStatusResponse.ServiceStatusResponse, error) {
	var resObj yundun_ServiceStatusResponse.ServiceStatusResponse

	if serviceStatusRequest == nil {
		serviceStatusRequest = new(ServiceStatusRequest)
	}
	err := c.DoAction(serviceStatusRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) QueryDdosConfig(queryDdosConfigRequest *QueryDdosConfigRequest) (yundun_QueryDdosConfigResponse.QueryDdosConfigResponse, error) {
	var resObj yundun_QueryDdosConfigResponse.QueryDdosConfigResponse

	if queryDdosConfigRequest == nil {
		queryDdosConfigRequest = new(QueryDdosConfigRequest)
	}
	err := c.DoAction(queryDdosConfigRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) OpenVulScan(openVulScanRequest *OpenVulScanRequest) (yundun_OpenVulScanResponse.OpenVulScanResponse, error) {
	var resObj yundun_OpenVulScanResponse.OpenVulScanResponse

	if openVulScanRequest == nil {
		openVulScanRequest = new(OpenVulScanRequest)
	}
	err := c.DoAction(openVulScanRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) OpenPortScan(openPortScanRequest *OpenPortScanRequest) (yundun_OpenPortScanResponse.OpenPortScanResponse, error) {
	var resObj yundun_OpenPortScanResponse.OpenPortScanResponse

	if openPortScanRequest == nil {
		openPortScanRequest = new(OpenPortScanRequest)
	}
	err := c.DoAction(openPortScanRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) OpenCCProtect(openCCProtectRequest *OpenCCProtectRequest) (yundun_OpenCCProtectResponse.OpenCCProtectResponse, error) {
	var resObj yundun_OpenCCProtectResponse.OpenCCProtectResponse

	if openCCProtectRequest == nil {
		openCCProtectRequest = new(OpenCCProtectRequest)
	}
	err := c.DoAction(openCCProtectRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) GetDdosConfigOptions(getDdosConfigOptionsRequest *GetDdosConfigOptionsRequest) (yundun_GetDdosConfigOptionsResponse.GetDdosConfigOptionsResponse, error) {
	var resObj yundun_GetDdosConfigOptionsResponse.GetDdosConfigOptionsResponse

	if getDdosConfigOptionsRequest == nil {
		getDdosConfigOptionsRequest = new(GetDdosConfigOptionsRequest)
	}
	err := c.DoAction(getDdosConfigOptionsRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) DetectVulByIp(detectVulByIpRequest *DetectVulByIpRequest) (yundun_DetectVulByIpResponse.DetectVulByIpResponse, error) {
	var resObj yundun_DetectVulByIpResponse.DetectVulByIpResponse

	if detectVulByIpRequest == nil {
		detectVulByIpRequest = new(DetectVulByIpRequest)
	}
	err := c.DoAction(detectVulByIpRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) DetectVulById(detectVulByIdRequest *DetectVulByIdRequest) (yundun_DetectVulByIdResponse.DetectVulByIdResponse, error) {
	var resObj yundun_DetectVulByIdResponse.DetectVulByIdResponse

	if detectVulByIdRequest == nil {
		detectVulByIdRequest = new(DetectVulByIdRequest)
	}
	err := c.DoAction(detectVulByIdRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) DeleteCNameWaf(deleteCNameWafRequest *DeleteCNameWafRequest) (yundun_DeleteCNameWafResponse.DeleteCNameWafResponse, error) {
	var resObj yundun_DeleteCNameWafResponse.DeleteCNameWafResponse

	if deleteCNameWafRequest == nil {
		deleteCNameWafRequest = new(DeleteCNameWafRequest)
	}
	err := c.DoAction(deleteCNameWafRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) DdosFlowGraph(ddosFlowGraphRequest *DdosFlowGraphRequest) (yundun_DdosFlowGraphResponse.DdosFlowGraphResponse, error) {
	var resObj yundun_DdosFlowGraphResponse.DdosFlowGraphResponse

	if ddosFlowGraphRequest == nil {
		ddosFlowGraphRequest = new(DdosFlowGraphRequest)
	}
	err := c.DoAction(ddosFlowGraphRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) ConfirmLogin(confirmLoginRequest *ConfirmLoginRequest) (yundun_ConfirmLoginResponse.ConfirmLoginResponse, error) {
	var resObj yundun_ConfirmLoginResponse.ConfirmLoginResponse

	if confirmLoginRequest == nil {
		confirmLoginRequest = new(ConfirmLoginRequest)
	}
	err := c.DoAction(confirmLoginRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) ConfigDdos(configDdosRequest *ConfigDdosRequest) (yundun_ConfigDdosResponse.ConfigDdosResponse, error) {
	var resObj yundun_ConfigDdosResponse.ConfigDdosResponse

	if configDdosRequest == nil {
		configDdosRequest = new(ConfigDdosRequest)
	}
	err := c.DoAction(configDdosRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) CloseVulScan(closeVulScanRequest *CloseVulScanRequest) (yundun_CloseVulScanResponse.CloseVulScanResponse, error) {
	var resObj yundun_CloseVulScanResponse.CloseVulScanResponse

	if closeVulScanRequest == nil {
		closeVulScanRequest = new(CloseVulScanRequest)
	}
	err := c.DoAction(closeVulScanRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) ClosePortScan(closePortScanRequest *ClosePortScanRequest) (yundun_ClosePortScanResponse.ClosePortScanResponse, error) {
	var resObj yundun_ClosePortScanResponse.ClosePortScanResponse

	if closePortScanRequest == nil {
		closePortScanRequest = new(ClosePortScanRequest)
	}
	err := c.DoAction(closePortScanRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) CloseCCProtect(closeCCProtectRequest *CloseCCProtectRequest) (yundun_CloseCCProtectResponse.CloseCCProtectResponse, error) {
	var resObj yundun_CloseCCProtectResponse.CloseCCProtectResponse

	if closeCCProtectRequest == nil {
		closeCCProtectRequest = new(CloseCCProtectRequest)
	}
	err := c.DoAction(closeCCProtectRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) AddCNameWaf(addCNameWafRequest *AddCNameWafRequest) (yundun_AddCNameWafResponse.AddCNameWafResponse, error) {
	var resObj yundun_AddCNameWafResponse.AddCNameWafResponse

	if addCNameWafRequest == nil {
		addCNameWafRequest = new(AddCNameWafRequest)
	}
	err := c.DoAction(addCNameWafRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) WebshellLog(webshellLogRequest *WebshellLogRequest) (yundun_WebshellLogResponse.WebshellLogResponse, error) {
	var resObj yundun_WebshellLogResponse.WebshellLogResponse

	if webshellLogRequest == nil {
		webshellLogRequest = new(WebshellLogRequest)
	}
	err := c.DoAction(webshellLogRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) WafLog(wafLogRequest *WafLogRequest) (yundun_WafLogResponse.WafLogResponse, error) {
	var resObj yundun_WafLogResponse.WafLogResponse

	if wafLogRequest == nil {
		wafLogRequest = new(WafLogRequest)
	}
	err := c.DoAction(wafLogRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) SecureCheck(secureCheckRequest *SecureCheckRequest) (yundun_SecureCheckResponse.SecureCheckResponse, error) {
	var resObj yundun_SecureCheckResponse.SecureCheckResponse

	if secureCheckRequest == nil {
		secureCheckRequest = new(SecureCheckRequest)
	}
	err := c.DoAction(secureCheckRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) LogineventLog(logineventLogRequest *LogineventLogRequest) (yundun_LogineventLogResponse.LogineventLogResponse, error) {
	var resObj yundun_LogineventLogResponse.LogineventLogResponse

	if logineventLogRequest == nil {
		logineventLogRequest = new(LogineventLogRequest)
	}
	err := c.DoAction(logineventLogRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) ListInstanceInfos(listInstanceInfosRequest *ListInstanceInfosRequest) (yundun_ListInstanceInfosResponse.ListInstanceInfosResponse, error) {
	var resObj yundun_ListInstanceInfosResponse.ListInstanceInfosResponse

	if listInstanceInfosRequest == nil {
		listInstanceInfosRequest = new(ListInstanceInfosRequest)
	}
	err := c.DoAction(listInstanceInfosRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) DeleteBackDoorFile(deleteBackDoorFileRequest *DeleteBackDoorFileRequest) (yundun_DeleteBackDoorFileResponse.DeleteBackDoorFileResponse, error) {
	var resObj yundun_DeleteBackDoorFileResponse.DeleteBackDoorFileResponse

	if deleteBackDoorFileRequest == nil {
		deleteBackDoorFileRequest = new(DeleteBackDoorFileRequest)
	}
	err := c.DoAction(deleteBackDoorFileRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) DdosLog(ddosLogRequest *DdosLogRequest) (yundun_DdosLogResponse.DdosLogResponse, error) {
	var resObj yundun_DdosLogResponse.DdosLogResponse

	if ddosLogRequest == nil {
		ddosLogRequest = new(DdosLogRequest)
	}
	err := c.DoAction(ddosLogRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) BruteforceLog(bruteforceLogRequest *BruteforceLogRequest) (yundun_BruteforceLogResponse.BruteforceLogResponse, error) {
	var resObj yundun_BruteforceLogResponse.BruteforceLogResponse

	if bruteforceLogRequest == nil {
		bruteforceLogRequest = new(BruteforceLogRequest)
	}
	err := c.DoAction(bruteforceLogRequest, &resObj)
	return resObj, err
}
