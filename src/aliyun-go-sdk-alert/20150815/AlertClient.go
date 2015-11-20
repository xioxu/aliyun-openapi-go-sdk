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

package alert

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-alert/20150815/UpdateLogHubMetricResponse"

	"aliyun-go-sdk-alert/20150815/ListLogHubMetricResponse"

	"aliyun-go-sdk-alert/20150815/GetLogHubMetricResponse"

	"aliyun-go-sdk-alert/20150815/DeleteLogHubMetricResponse"

	"aliyun-go-sdk-alert/20150815/CreateLogHubMetricResponse"

	"aliyun-go-sdk-alert/20150815/UpdateDimensionsResponse"

	"aliyun-go-sdk-alert/20150815/ListDimensionsResponse"

	"aliyun-go-sdk-alert/20150815/GetDimensionsResponse"

	"aliyun-go-sdk-alert/20150815/DeleteDimensionsResponse"

	"aliyun-go-sdk-alert/20150815/CreateDimensionsResponse"

	"aliyun-go-sdk-alert/20150815/EnableAlertResponse"

	"aliyun-go-sdk-alert/20150815/DisableAlertResponse"

	"aliyun-go-sdk-alert/20150815/ListAlertStateResponse"

	"aliyun-go-sdk-alert/20150815/ListAlertResponse"

	"aliyun-go-sdk-alert/20150815/GetAlertResponse"

	"aliyun-go-sdk-alert/20150815/DeleteAlertResponse"

	"aliyun-go-sdk-alert/20150815/CreateAlertResponse"

	"aliyun-go-sdk-alert/20150815/UpdateContactGroupResponse"

	"aliyun-go-sdk-alert/20150815/ListContactGroupResponse"

	"aliyun-go-sdk-alert/20150815/GetContactGroupResponse"

	"aliyun-go-sdk-alert/20150815/DeleteContactGroupResponse"

	"aliyun-go-sdk-alert/20150815/CreateContactGroupResponse"

	"aliyun-go-sdk-alert/20150815/GetContactResponse"

	"aliyun-go-sdk-alert/20150815/DeleteContactResponse"

	"aliyun-go-sdk-alert/20150815/CreateContactResponse"

	"aliyun-go-sdk-alert/20150815/UpdateAlertResponse"

	"aliyun-go-sdk-alert/20150815/DeleteDBSourceResponse"

	"aliyun-go-sdk-alert/20150815/CreateDBSourceResponse"

	"aliyun-go-sdk-alert/20150815/UpdateDBMetricResponse"

	"aliyun-go-sdk-alert/20150815/ListDBMetricResponse"

	"aliyun-go-sdk-alert/20150815/GetDBMetricResponse"

	"aliyun-go-sdk-alert/20150815/DeleteDBMetricResponse"

	"aliyun-go-sdk-alert/20150815/CreateDBMetricResponse"

	"aliyun-go-sdk-alert/20150815/UpdateContactResponse"

	"aliyun-go-sdk-alert/20150815/ListContactResponse"

	"aliyun-go-sdk-alert/20150815/BatchQueryProjectResponse"

	"aliyun-go-sdk-alert/20150815/ListNotifyHistoryResponse"

	"aliyun-go-sdk-alert/20150815/UpdateLevelChannelResponse"

	"aliyun-go-sdk-alert/20150815/ListLevelChannelResponse"

	"aliyun-go-sdk-alert/20150815/GetLevelChannelResponse"

	"aliyun-go-sdk-alert/20150815/DeleteLevelChannelResponse"

	"aliyun-go-sdk-alert/20150815/CreateLevelChannelResponse"

	"aliyun-go-sdk-alert/20150815/UpdateDBSourceResponse"

	"aliyun-go-sdk-alert/20150815/ListDBSourceResponse"

	"aliyun-go-sdk-alert/20150815/GetDBSourceResponse"

	"aliyun-go-sdk-alert/20150815/UpdateProjectResponse"

	"aliyun-go-sdk-alert/20150815/ListProjectResponse"

	"aliyun-go-sdk-alert/20150815/GetProjectResponse"

	"aliyun-go-sdk-alert/20150815/DeleteProjectResponse"

	"aliyun-go-sdk-alert/20150815/CreateProjectResponse"

	"aliyun-go-sdk-alert/20150815/RemoveProjectOwnerResponse"

	"aliyun-go-sdk-alert/20150815/GrantProjectOwnerResponse"
)

type AlertClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *AlertClient {
	p := &AlertClient{}
	p.Version = "2015-08-15"
	p.ProductName = "Alert"
	p.Profile = profile
	p.ApiStyle = "ROA"
	p.HttpRequestBuilder = &core.RoaHttpRequestBuilder{}
	return p
}

func (c *AlertClient) UpdateLogHubMetric(updateLogHubMetricRequest *UpdateLogHubMetricRequest) (alert_UpdateLogHubMetricResponse.UpdateLogHubMetricResponse, error) {
	var resObj alert_UpdateLogHubMetricResponse.UpdateLogHubMetricResponse

	if updateLogHubMetricRequest == nil {
		updateLogHubMetricRequest = new(UpdateLogHubMetricRequest)
	}
	err := c.DoAction(updateLogHubMetricRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) ListLogHubMetric(listLogHubMetricRequest *ListLogHubMetricRequest) (alert_ListLogHubMetricResponse.ListLogHubMetricResponse, error) {
	var resObj alert_ListLogHubMetricResponse.ListLogHubMetricResponse

	if listLogHubMetricRequest == nil {
		listLogHubMetricRequest = new(ListLogHubMetricRequest)
	}
	err := c.DoAction(listLogHubMetricRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) GetLogHubMetric(getLogHubMetricRequest *GetLogHubMetricRequest) (alert_GetLogHubMetricResponse.GetLogHubMetricResponse, error) {
	var resObj alert_GetLogHubMetricResponse.GetLogHubMetricResponse

	if getLogHubMetricRequest == nil {
		getLogHubMetricRequest = new(GetLogHubMetricRequest)
	}
	err := c.DoAction(getLogHubMetricRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) DeleteLogHubMetric(deleteLogHubMetricRequest *DeleteLogHubMetricRequest) (alert_DeleteLogHubMetricResponse.DeleteLogHubMetricResponse, error) {
	var resObj alert_DeleteLogHubMetricResponse.DeleteLogHubMetricResponse

	if deleteLogHubMetricRequest == nil {
		deleteLogHubMetricRequest = new(DeleteLogHubMetricRequest)
	}
	err := c.DoAction(deleteLogHubMetricRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) CreateLogHubMetric(createLogHubMetricRequest *CreateLogHubMetricRequest) (alert_CreateLogHubMetricResponse.CreateLogHubMetricResponse, error) {
	var resObj alert_CreateLogHubMetricResponse.CreateLogHubMetricResponse

	if createLogHubMetricRequest == nil {
		createLogHubMetricRequest = new(CreateLogHubMetricRequest)
	}
	err := c.DoAction(createLogHubMetricRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) UpdateDimensions(updateDimensionsRequest *UpdateDimensionsRequest) (alert_UpdateDimensionsResponse.UpdateDimensionsResponse, error) {
	var resObj alert_UpdateDimensionsResponse.UpdateDimensionsResponse

	if updateDimensionsRequest == nil {
		updateDimensionsRequest = new(UpdateDimensionsRequest)
	}
	err := c.DoAction(updateDimensionsRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) ListDimensions(listDimensionsRequest *ListDimensionsRequest) (alert_ListDimensionsResponse.ListDimensionsResponse, error) {
	var resObj alert_ListDimensionsResponse.ListDimensionsResponse

	if listDimensionsRequest == nil {
		listDimensionsRequest = new(ListDimensionsRequest)
	}
	err := c.DoAction(listDimensionsRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) GetDimensions(getDimensionsRequest *GetDimensionsRequest) (alert_GetDimensionsResponse.GetDimensionsResponse, error) {
	var resObj alert_GetDimensionsResponse.GetDimensionsResponse

	if getDimensionsRequest == nil {
		getDimensionsRequest = new(GetDimensionsRequest)
	}
	err := c.DoAction(getDimensionsRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) DeleteDimensions(deleteDimensionsRequest *DeleteDimensionsRequest) (alert_DeleteDimensionsResponse.DeleteDimensionsResponse, error) {
	var resObj alert_DeleteDimensionsResponse.DeleteDimensionsResponse

	if deleteDimensionsRequest == nil {
		deleteDimensionsRequest = new(DeleteDimensionsRequest)
	}
	err := c.DoAction(deleteDimensionsRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) CreateDimensions(createDimensionsRequest *CreateDimensionsRequest) (alert_CreateDimensionsResponse.CreateDimensionsResponse, error) {
	var resObj alert_CreateDimensionsResponse.CreateDimensionsResponse

	if createDimensionsRequest == nil {
		createDimensionsRequest = new(CreateDimensionsRequest)
	}
	err := c.DoAction(createDimensionsRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) EnableAlert(enableAlertRequest *EnableAlertRequest) (alert_EnableAlertResponse.EnableAlertResponse, error) {
	var resObj alert_EnableAlertResponse.EnableAlertResponse

	if enableAlertRequest == nil {
		enableAlertRequest = new(EnableAlertRequest)
	}
	err := c.DoAction(enableAlertRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) DisableAlert(disableAlertRequest *DisableAlertRequest) (alert_DisableAlertResponse.DisableAlertResponse, error) {
	var resObj alert_DisableAlertResponse.DisableAlertResponse

	if disableAlertRequest == nil {
		disableAlertRequest = new(DisableAlertRequest)
	}
	err := c.DoAction(disableAlertRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) ListAlertState(listAlertStateRequest *ListAlertStateRequest) (alert_ListAlertStateResponse.ListAlertStateResponse, error) {
	var resObj alert_ListAlertStateResponse.ListAlertStateResponse

	if listAlertStateRequest == nil {
		listAlertStateRequest = new(ListAlertStateRequest)
	}
	err := c.DoAction(listAlertStateRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) ListAlert(listAlertRequest *ListAlertRequest) (alert_ListAlertResponse.ListAlertResponse, error) {
	var resObj alert_ListAlertResponse.ListAlertResponse

	if listAlertRequest == nil {
		listAlertRequest = new(ListAlertRequest)
	}
	err := c.DoAction(listAlertRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) GetAlert(getAlertRequest *GetAlertRequest) (alert_GetAlertResponse.GetAlertResponse, error) {
	var resObj alert_GetAlertResponse.GetAlertResponse

	if getAlertRequest == nil {
		getAlertRequest = new(GetAlertRequest)
	}
	err := c.DoAction(getAlertRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) DeleteAlert(deleteAlertRequest *DeleteAlertRequest) (alert_DeleteAlertResponse.DeleteAlertResponse, error) {
	var resObj alert_DeleteAlertResponse.DeleteAlertResponse

	if deleteAlertRequest == nil {
		deleteAlertRequest = new(DeleteAlertRequest)
	}
	err := c.DoAction(deleteAlertRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) CreateAlert(createAlertRequest *CreateAlertRequest) (alert_CreateAlertResponse.CreateAlertResponse, error) {
	var resObj alert_CreateAlertResponse.CreateAlertResponse

	if createAlertRequest == nil {
		createAlertRequest = new(CreateAlertRequest)
	}
	err := c.DoAction(createAlertRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) UpdateContactGroup(updateContactGroupRequest *UpdateContactGroupRequest) (alert_UpdateContactGroupResponse.UpdateContactGroupResponse, error) {
	var resObj alert_UpdateContactGroupResponse.UpdateContactGroupResponse

	if updateContactGroupRequest == nil {
		updateContactGroupRequest = new(UpdateContactGroupRequest)
	}
	err := c.DoAction(updateContactGroupRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) ListContactGroup(listContactGroupRequest *ListContactGroupRequest) (alert_ListContactGroupResponse.ListContactGroupResponse, error) {
	var resObj alert_ListContactGroupResponse.ListContactGroupResponse

	if listContactGroupRequest == nil {
		listContactGroupRequest = new(ListContactGroupRequest)
	}
	err := c.DoAction(listContactGroupRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) GetContactGroup(getContactGroupRequest *GetContactGroupRequest) (alert_GetContactGroupResponse.GetContactGroupResponse, error) {
	var resObj alert_GetContactGroupResponse.GetContactGroupResponse

	if getContactGroupRequest == nil {
		getContactGroupRequest = new(GetContactGroupRequest)
	}
	err := c.DoAction(getContactGroupRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) DeleteContactGroup(deleteContactGroupRequest *DeleteContactGroupRequest) (alert_DeleteContactGroupResponse.DeleteContactGroupResponse, error) {
	var resObj alert_DeleteContactGroupResponse.DeleteContactGroupResponse

	if deleteContactGroupRequest == nil {
		deleteContactGroupRequest = new(DeleteContactGroupRequest)
	}
	err := c.DoAction(deleteContactGroupRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) CreateContactGroup(createContactGroupRequest *CreateContactGroupRequest) (alert_CreateContactGroupResponse.CreateContactGroupResponse, error) {
	var resObj alert_CreateContactGroupResponse.CreateContactGroupResponse

	if createContactGroupRequest == nil {
		createContactGroupRequest = new(CreateContactGroupRequest)
	}
	err := c.DoAction(createContactGroupRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) GetContact(getContactRequest *GetContactRequest) (alert_GetContactResponse.GetContactResponse, error) {
	var resObj alert_GetContactResponse.GetContactResponse

	if getContactRequest == nil {
		getContactRequest = new(GetContactRequest)
	}
	err := c.DoAction(getContactRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) DeleteContact(deleteContactRequest *DeleteContactRequest) (alert_DeleteContactResponse.DeleteContactResponse, error) {
	var resObj alert_DeleteContactResponse.DeleteContactResponse

	if deleteContactRequest == nil {
		deleteContactRequest = new(DeleteContactRequest)
	}
	err := c.DoAction(deleteContactRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) CreateContact(createContactRequest *CreateContactRequest) (alert_CreateContactResponse.CreateContactResponse, error) {
	var resObj alert_CreateContactResponse.CreateContactResponse

	if createContactRequest == nil {
		createContactRequest = new(CreateContactRequest)
	}
	err := c.DoAction(createContactRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) UpdateAlert(updateAlertRequest *UpdateAlertRequest) (alert_UpdateAlertResponse.UpdateAlertResponse, error) {
	var resObj alert_UpdateAlertResponse.UpdateAlertResponse

	if updateAlertRequest == nil {
		updateAlertRequest = new(UpdateAlertRequest)
	}
	err := c.DoAction(updateAlertRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) DeleteDBSource(deleteDBSourceRequest *DeleteDBSourceRequest) (alert_DeleteDBSourceResponse.DeleteDBSourceResponse, error) {
	var resObj alert_DeleteDBSourceResponse.DeleteDBSourceResponse

	if deleteDBSourceRequest == nil {
		deleteDBSourceRequest = new(DeleteDBSourceRequest)
	}
	err := c.DoAction(deleteDBSourceRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) CreateDBSource(createDBSourceRequest *CreateDBSourceRequest) (alert_CreateDBSourceResponse.CreateDBSourceResponse, error) {
	var resObj alert_CreateDBSourceResponse.CreateDBSourceResponse

	if createDBSourceRequest == nil {
		createDBSourceRequest = new(CreateDBSourceRequest)
	}
	err := c.DoAction(createDBSourceRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) UpdateDBMetric(updateDBMetricRequest *UpdateDBMetricRequest) (alert_UpdateDBMetricResponse.UpdateDBMetricResponse, error) {
	var resObj alert_UpdateDBMetricResponse.UpdateDBMetricResponse

	if updateDBMetricRequest == nil {
		updateDBMetricRequest = new(UpdateDBMetricRequest)
	}
	err := c.DoAction(updateDBMetricRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) ListDBMetric(listDBMetricRequest *ListDBMetricRequest) (alert_ListDBMetricResponse.ListDBMetricResponse, error) {
	var resObj alert_ListDBMetricResponse.ListDBMetricResponse

	if listDBMetricRequest == nil {
		listDBMetricRequest = new(ListDBMetricRequest)
	}
	err := c.DoAction(listDBMetricRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) GetDBMetric(getDBMetricRequest *GetDBMetricRequest) (alert_GetDBMetricResponse.GetDBMetricResponse, error) {
	var resObj alert_GetDBMetricResponse.GetDBMetricResponse

	if getDBMetricRequest == nil {
		getDBMetricRequest = new(GetDBMetricRequest)
	}
	err := c.DoAction(getDBMetricRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) DeleteDBMetric(deleteDBMetricRequest *DeleteDBMetricRequest) (alert_DeleteDBMetricResponse.DeleteDBMetricResponse, error) {
	var resObj alert_DeleteDBMetricResponse.DeleteDBMetricResponse

	if deleteDBMetricRequest == nil {
		deleteDBMetricRequest = new(DeleteDBMetricRequest)
	}
	err := c.DoAction(deleteDBMetricRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) CreateDBMetric(createDBMetricRequest *CreateDBMetricRequest) (alert_CreateDBMetricResponse.CreateDBMetricResponse, error) {
	var resObj alert_CreateDBMetricResponse.CreateDBMetricResponse

	if createDBMetricRequest == nil {
		createDBMetricRequest = new(CreateDBMetricRequest)
	}
	err := c.DoAction(createDBMetricRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) UpdateContact(updateContactRequest *UpdateContactRequest) (alert_UpdateContactResponse.UpdateContactResponse, error) {
	var resObj alert_UpdateContactResponse.UpdateContactResponse

	if updateContactRequest == nil {
		updateContactRequest = new(UpdateContactRequest)
	}
	err := c.DoAction(updateContactRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) ListContact(listContactRequest *ListContactRequest) (alert_ListContactResponse.ListContactResponse, error) {
	var resObj alert_ListContactResponse.ListContactResponse

	if listContactRequest == nil {
		listContactRequest = new(ListContactRequest)
	}
	err := c.DoAction(listContactRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) BatchQueryProject(batchQueryProjectRequest *BatchQueryProjectRequest) (alert_BatchQueryProjectResponse.BatchQueryProjectResponse, error) {
	var resObj alert_BatchQueryProjectResponse.BatchQueryProjectResponse

	if batchQueryProjectRequest == nil {
		batchQueryProjectRequest = new(BatchQueryProjectRequest)
	}
	err := c.DoAction(batchQueryProjectRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) ListNotifyHistory(listNotifyHistoryRequest *ListNotifyHistoryRequest) (alert_ListNotifyHistoryResponse.ListNotifyHistoryResponse, error) {
	var resObj alert_ListNotifyHistoryResponse.ListNotifyHistoryResponse

	if listNotifyHistoryRequest == nil {
		listNotifyHistoryRequest = new(ListNotifyHistoryRequest)
	}
	err := c.DoAction(listNotifyHistoryRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) UpdateLevelChannel(updateLevelChannelRequest *UpdateLevelChannelRequest) (alert_UpdateLevelChannelResponse.UpdateLevelChannelResponse, error) {
	var resObj alert_UpdateLevelChannelResponse.UpdateLevelChannelResponse

	if updateLevelChannelRequest == nil {
		updateLevelChannelRequest = new(UpdateLevelChannelRequest)
	}
	err := c.DoAction(updateLevelChannelRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) ListLevelChannel(listLevelChannelRequest *ListLevelChannelRequest) (alert_ListLevelChannelResponse.ListLevelChannelResponse, error) {
	var resObj alert_ListLevelChannelResponse.ListLevelChannelResponse

	if listLevelChannelRequest == nil {
		listLevelChannelRequest = new(ListLevelChannelRequest)
	}
	err := c.DoAction(listLevelChannelRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) GetLevelChannel(getLevelChannelRequest *GetLevelChannelRequest) (alert_GetLevelChannelResponse.GetLevelChannelResponse, error) {
	var resObj alert_GetLevelChannelResponse.GetLevelChannelResponse

	if getLevelChannelRequest == nil {
		getLevelChannelRequest = new(GetLevelChannelRequest)
	}
	err := c.DoAction(getLevelChannelRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) DeleteLevelChannel(deleteLevelChannelRequest *DeleteLevelChannelRequest) (alert_DeleteLevelChannelResponse.DeleteLevelChannelResponse, error) {
	var resObj alert_DeleteLevelChannelResponse.DeleteLevelChannelResponse

	if deleteLevelChannelRequest == nil {
		deleteLevelChannelRequest = new(DeleteLevelChannelRequest)
	}
	err := c.DoAction(deleteLevelChannelRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) CreateLevelChannel(createLevelChannelRequest *CreateLevelChannelRequest) (alert_CreateLevelChannelResponse.CreateLevelChannelResponse, error) {
	var resObj alert_CreateLevelChannelResponse.CreateLevelChannelResponse

	if createLevelChannelRequest == nil {
		createLevelChannelRequest = new(CreateLevelChannelRequest)
	}
	err := c.DoAction(createLevelChannelRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) UpdateDBSource(updateDBSourceRequest *UpdateDBSourceRequest) (alert_UpdateDBSourceResponse.UpdateDBSourceResponse, error) {
	var resObj alert_UpdateDBSourceResponse.UpdateDBSourceResponse

	if updateDBSourceRequest == nil {
		updateDBSourceRequest = new(UpdateDBSourceRequest)
	}
	err := c.DoAction(updateDBSourceRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) ListDBSource(listDBSourceRequest *ListDBSourceRequest) (alert_ListDBSourceResponse.ListDBSourceResponse, error) {
	var resObj alert_ListDBSourceResponse.ListDBSourceResponse

	if listDBSourceRequest == nil {
		listDBSourceRequest = new(ListDBSourceRequest)
	}
	err := c.DoAction(listDBSourceRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) GetDBSource(getDBSourceRequest *GetDBSourceRequest) (alert_GetDBSourceResponse.GetDBSourceResponse, error) {
	var resObj alert_GetDBSourceResponse.GetDBSourceResponse

	if getDBSourceRequest == nil {
		getDBSourceRequest = new(GetDBSourceRequest)
	}
	err := c.DoAction(getDBSourceRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) UpdateProject(updateProjectRequest *UpdateProjectRequest) (alert_UpdateProjectResponse.UpdateProjectResponse, error) {
	var resObj alert_UpdateProjectResponse.UpdateProjectResponse

	if updateProjectRequest == nil {
		updateProjectRequest = new(UpdateProjectRequest)
	}
	err := c.DoAction(updateProjectRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) ListProject(listProjectRequest *ListProjectRequest) (alert_ListProjectResponse.ListProjectResponse, error) {
	var resObj alert_ListProjectResponse.ListProjectResponse

	if listProjectRequest == nil {
		listProjectRequest = new(ListProjectRequest)
	}
	err := c.DoAction(listProjectRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) GetProject(getProjectRequest *GetProjectRequest) (alert_GetProjectResponse.GetProjectResponse, error) {
	var resObj alert_GetProjectResponse.GetProjectResponse

	if getProjectRequest == nil {
		getProjectRequest = new(GetProjectRequest)
	}
	err := c.DoAction(getProjectRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) DeleteProject(deleteProjectRequest *DeleteProjectRequest) (alert_DeleteProjectResponse.DeleteProjectResponse, error) {
	var resObj alert_DeleteProjectResponse.DeleteProjectResponse

	if deleteProjectRequest == nil {
		deleteProjectRequest = new(DeleteProjectRequest)
	}
	err := c.DoAction(deleteProjectRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) CreateProject(createProjectRequest *CreateProjectRequest) (alert_CreateProjectResponse.CreateProjectResponse, error) {
	var resObj alert_CreateProjectResponse.CreateProjectResponse

	if createProjectRequest == nil {
		createProjectRequest = new(CreateProjectRequest)
	}
	err := c.DoAction(createProjectRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) RemoveProjectOwner(removeProjectOwnerRequest *RemoveProjectOwnerRequest) (alert_RemoveProjectOwnerResponse.RemoveProjectOwnerResponse, error) {
	var resObj alert_RemoveProjectOwnerResponse.RemoveProjectOwnerResponse

	if removeProjectOwnerRequest == nil {
		removeProjectOwnerRequest = new(RemoveProjectOwnerRequest)
	}
	err := c.DoAction(removeProjectOwnerRequest, &resObj)
	return resObj, err
}

func (c *AlertClient) GrantProjectOwner(grantProjectOwnerRequest *GrantProjectOwnerRequest) (alert_GrantProjectOwnerResponse.GrantProjectOwnerResponse, error) {
	var resObj alert_GrantProjectOwnerResponse.GrantProjectOwnerResponse

	if grantProjectOwnerRequest == nil {
		grantProjectOwnerRequest = new(GrantProjectOwnerRequest)
	}
	err := c.DoAction(grantProjectOwnerRequest, &resObj)
	return resObj, err
}
