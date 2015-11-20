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

package pts

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-pts/20150801/SetScenarioStatusResponse"

	"aliyun-go-sdk-pts/20150801/SendWangWangResponse"

	"aliyun-go-sdk-pts/20150801/ReportVuserResponse"

	"aliyun-go-sdk-pts/20150801/ReportTestSampleResponse"

	"aliyun-go-sdk-pts/20150801/ReportLogSampleResponse"

	"aliyun-go-sdk-pts/20150801/GetTasksResponse"

	"aliyun-go-sdk-pts/20150801/GetScriptResponse"

	"aliyun-go-sdk-pts/20150801/GetKeySecretResponse"

	"aliyun-go-sdk-pts/20150801/CreateTransactionResponse"

	"aliyun-go-sdk-pts/20150801/StopTaskResponse"

	"aliyun-go-sdk-pts/20150801/SetTaskStatusResponse"
)

type PtsClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *PtsClient {
	p := &PtsClient{}
	p.Version = "2015-08-01"
	p.ProductName = "PTS"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *PtsClient) SetScenarioStatus(setScenarioStatusRequest *SetScenarioStatusRequest) (pts_SetScenarioStatusResponse.SetScenarioStatusResponse, error) {
	var resObj pts_SetScenarioStatusResponse.SetScenarioStatusResponse

	if setScenarioStatusRequest == nil {
		setScenarioStatusRequest = new(SetScenarioStatusRequest)
	}
	err := c.DoAction(setScenarioStatusRequest, &resObj)
	return resObj, err
}

func (c *PtsClient) SendWangWang(sendWangWangRequest *SendWangWangRequest) (pts_SendWangWangResponse.SendWangWangResponse, error) {
	var resObj pts_SendWangWangResponse.SendWangWangResponse

	if sendWangWangRequest == nil {
		sendWangWangRequest = new(SendWangWangRequest)
	}
	err := c.DoAction(sendWangWangRequest, &resObj)
	return resObj, err
}

func (c *PtsClient) ReportVuser(reportVuserRequest *ReportVuserRequest) (pts_ReportVuserResponse.ReportVuserResponse, error) {
	var resObj pts_ReportVuserResponse.ReportVuserResponse

	if reportVuserRequest == nil {
		reportVuserRequest = new(ReportVuserRequest)
	}
	err := c.DoAction(reportVuserRequest, &resObj)
	return resObj, err
}

func (c *PtsClient) ReportTestSample(reportTestSampleRequest *ReportTestSampleRequest) (pts_ReportTestSampleResponse.ReportTestSampleResponse, error) {
	var resObj pts_ReportTestSampleResponse.ReportTestSampleResponse

	if reportTestSampleRequest == nil {
		reportTestSampleRequest = new(ReportTestSampleRequest)
	}
	err := c.DoAction(reportTestSampleRequest, &resObj)
	return resObj, err
}

func (c *PtsClient) ReportLogSample(reportLogSampleRequest *ReportLogSampleRequest) (pts_ReportLogSampleResponse.ReportLogSampleResponse, error) {
	var resObj pts_ReportLogSampleResponse.ReportLogSampleResponse

	if reportLogSampleRequest == nil {
		reportLogSampleRequest = new(ReportLogSampleRequest)
	}
	err := c.DoAction(reportLogSampleRequest, &resObj)
	return resObj, err
}

func (c *PtsClient) GetTasks(getTasksRequest *GetTasksRequest) (pts_GetTasksResponse.GetTasksResponse, error) {
	var resObj pts_GetTasksResponse.GetTasksResponse

	if getTasksRequest == nil {
		getTasksRequest = new(GetTasksRequest)
	}
	err := c.DoAction(getTasksRequest, &resObj)
	return resObj, err
}

func (c *PtsClient) GetScript(getScriptRequest *GetScriptRequest) (pts_GetScriptResponse.GetScriptResponse, error) {
	var resObj pts_GetScriptResponse.GetScriptResponse

	if getScriptRequest == nil {
		getScriptRequest = new(GetScriptRequest)
	}
	err := c.DoAction(getScriptRequest, &resObj)
	return resObj, err
}

func (c *PtsClient) GetKeySecret(getKeySecretRequest *GetKeySecretRequest) (pts_GetKeySecretResponse.GetKeySecretResponse, error) {
	var resObj pts_GetKeySecretResponse.GetKeySecretResponse

	if getKeySecretRequest == nil {
		getKeySecretRequest = new(GetKeySecretRequest)
	}
	err := c.DoAction(getKeySecretRequest, &resObj)
	return resObj, err
}

func (c *PtsClient) CreateTransaction(createTransactionRequest *CreateTransactionRequest) (pts_CreateTransactionResponse.CreateTransactionResponse, error) {
	var resObj pts_CreateTransactionResponse.CreateTransactionResponse

	if createTransactionRequest == nil {
		createTransactionRequest = new(CreateTransactionRequest)
	}
	err := c.DoAction(createTransactionRequest, &resObj)
	return resObj, err
}

func (c *PtsClient) StopTask(stopTaskRequest *StopTaskRequest) (pts_StopTaskResponse.StopTaskResponse, error) {
	var resObj pts_StopTaskResponse.StopTaskResponse

	if stopTaskRequest == nil {
		stopTaskRequest = new(StopTaskRequest)
	}
	err := c.DoAction(stopTaskRequest, &resObj)
	return resObj, err
}

func (c *PtsClient) SetTaskStatus(setTaskStatusRequest *SetTaskStatusRequest) (pts_SetTaskStatusResponse.SetTaskStatusResponse, error) {
	var resObj pts_SetTaskStatusResponse.SetTaskStatusResponse

	if setTaskStatusRequest == nil {
		setTaskStatusRequest = new(SetTaskStatusRequest)
	}
	err := c.DoAction(setTaskStatusRequest, &resObj)
	return resObj, err
}
