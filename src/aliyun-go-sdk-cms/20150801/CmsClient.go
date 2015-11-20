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

package cms

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-cms/20150801/GetMetricStatisticsResponse"

	"aliyun-go-sdk-cms/20150801/BatchQueryMetricResponse"

	"aliyun-go-sdk-cms/20150801/QueryStatisticsResponse"

	"aliyun-go-sdk-cms/20150801/QueryMetricTopNResponse"

	"aliyun-go-sdk-cms/20150801/QueryMetricResponse"

	"aliyun-go-sdk-cms/20150801/QueryListMetricResponse"

	"aliyun-go-sdk-cms/20150801/QueryIncrementalResponse"

	"aliyun-go-sdk-cms/20150801/DescribeMetricResponse"

	"aliyun-go-sdk-cms/20150801/StopProjectResponse"

	"aliyun-go-sdk-cms/20150801/StatusProjectResponse"

	"aliyun-go-sdk-cms/20150801/StartProjectResponse"

	"aliyun-go-sdk-cms/20150801/BatchCreateSqlMetricsResponse"

	"aliyun-go-sdk-cms/20150801/UpdateMetricStreamResponse"

	"aliyun-go-sdk-cms/20150801/StopMetricStreamResponse"

	"aliyun-go-sdk-cms/20150801/GetMetricStreamStatusResponse"

	"aliyun-go-sdk-cms/20150801/StartMetricStreamResponse"

	"aliyun-go-sdk-cms/20150801/GetMetricStreamMetaResponse"

	"aliyun-go-sdk-cms/20150801/ListMetricStreamResponse"

	"aliyun-go-sdk-cms/20150801/GetMetricStreamResponse"

	"aliyun-go-sdk-cms/20150801/DeleteMetricStreamResponse"

	"aliyun-go-sdk-cms/20150801/CreateMetricStreamResponse"

	"aliyun-go-sdk-cms/20150801/DescribeMetricDatumResponse"

	"aliyun-go-sdk-cms/20150801/ListDimTableDataResponse"

	"aliyun-go-sdk-cms/20150801/PutDimTableDataResponse"

	"aliyun-go-sdk-cms/20150801/UpdateProjectResponse"

	"aliyun-go-sdk-cms/20150801/ListProjectResponse"

	"aliyun-go-sdk-cms/20150801/GetProjectResponse"

	"aliyun-go-sdk-cms/20150801/DeleteProjectResponse"

	"aliyun-go-sdk-cms/20150801/CreateProjectResponse"

	"aliyun-go-sdk-cms/20150801/UpdateMetricsResponse"

	"aliyun-go-sdk-cms/20150801/UpdateSQLMetricsResponse"

	"aliyun-go-sdk-cms/20150801/ListSQLMetricsResponse"

	"aliyun-go-sdk-cms/20150801/GetSQLMetricsResponse"

	"aliyun-go-sdk-cms/20150801/CreateSQLMetricsResponse"

	"aliyun-go-sdk-cms/20150801/ListMetricsPlanResponse"

	"aliyun-go-sdk-cms/20150801/GetMetricsMetaResponse"

	"aliyun-go-sdk-cms/20150801/ListMetricsResponse"

	"aliyun-go-sdk-cms/20150801/GetMetricsResponse"

	"aliyun-go-sdk-cms/20150801/DeleteMetricsResponse"

	"aliyun-go-sdk-cms/20150801/CreateMetricsResponse"

	"aliyun-go-sdk-cms/20150801/BatchCreateMetricsResponse"

	"aliyun-go-sdk-cms/20150801/DescribeMetricListResponse"

	"aliyun-go-sdk-cms/20150801/UpdateDimTableResponse"

	"aliyun-go-sdk-cms/20150801/ListDimTableResponse"

	"aliyun-go-sdk-cms/20150801/GetDimTableResponse"

	"aliyun-go-sdk-cms/20150801/DeleteDimTableResponse"

	"aliyun-go-sdk-cms/20150801/CreateDimTableResponse"

	"aliyun-go-sdk-cms/20150801/DeleteDimTableDataResponse"

	"aliyun-go-sdk-cms/20150801/BatchPutDimTableDataResponse"
)

type CmsClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *CmsClient {
	p := &CmsClient{}
	p.Version = "2015-08-01"
	p.ProductName = "Cms"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *CmsClient) GetMetricStatistics(getMetricStatisticsRequest *GetMetricStatisticsRequest) (cms_GetMetricStatisticsResponse.GetMetricStatisticsResponse, error) {
	var resObj cms_GetMetricStatisticsResponse.GetMetricStatisticsResponse

	if getMetricStatisticsRequest == nil {
		getMetricStatisticsRequest = new(GetMetricStatisticsRequest)
	}
	err := c.DoAction(getMetricStatisticsRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) BatchQueryMetric(batchQueryMetricRequest *BatchQueryMetricRequest) (cms_BatchQueryMetricResponse.BatchQueryMetricResponse, error) {
	var resObj cms_BatchQueryMetricResponse.BatchQueryMetricResponse

	if batchQueryMetricRequest == nil {
		batchQueryMetricRequest = new(BatchQueryMetricRequest)
	}
	err := c.DoAction(batchQueryMetricRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) QueryStatistics(queryStatisticsRequest *QueryStatisticsRequest) (cms_QueryStatisticsResponse.QueryStatisticsResponse, error) {
	var resObj cms_QueryStatisticsResponse.QueryStatisticsResponse

	if queryStatisticsRequest == nil {
		queryStatisticsRequest = new(QueryStatisticsRequest)
	}
	err := c.DoAction(queryStatisticsRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) QueryMetricTopN(queryMetricTopNRequest *QueryMetricTopNRequest) (cms_QueryMetricTopNResponse.QueryMetricTopNResponse, error) {
	var resObj cms_QueryMetricTopNResponse.QueryMetricTopNResponse

	if queryMetricTopNRequest == nil {
		queryMetricTopNRequest = new(QueryMetricTopNRequest)
	}
	err := c.DoAction(queryMetricTopNRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) QueryMetric(queryMetricRequest *QueryMetricRequest) (cms_QueryMetricResponse.QueryMetricResponse, error) {
	var resObj cms_QueryMetricResponse.QueryMetricResponse

	if queryMetricRequest == nil {
		queryMetricRequest = new(QueryMetricRequest)
	}
	err := c.DoAction(queryMetricRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) QueryListMetric(queryListMetricRequest *QueryListMetricRequest) (cms_QueryListMetricResponse.QueryListMetricResponse, error) {
	var resObj cms_QueryListMetricResponse.QueryListMetricResponse

	if queryListMetricRequest == nil {
		queryListMetricRequest = new(QueryListMetricRequest)
	}
	err := c.DoAction(queryListMetricRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) QueryIncremental(queryIncrementalRequest *QueryIncrementalRequest) (cms_QueryIncrementalResponse.QueryIncrementalResponse, error) {
	var resObj cms_QueryIncrementalResponse.QueryIncrementalResponse

	if queryIncrementalRequest == nil {
		queryIncrementalRequest = new(QueryIncrementalRequest)
	}
	err := c.DoAction(queryIncrementalRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) DescribeMetric(describeMetricRequest *DescribeMetricRequest) (cms_DescribeMetricResponse.DescribeMetricResponse, error) {
	var resObj cms_DescribeMetricResponse.DescribeMetricResponse

	if describeMetricRequest == nil {
		describeMetricRequest = new(DescribeMetricRequest)
	}
	err := c.DoAction(describeMetricRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) StopProject(stopProjectRequest *StopProjectRequest) (cms_StopProjectResponse.StopProjectResponse, error) {
	var resObj cms_StopProjectResponse.StopProjectResponse

	if stopProjectRequest == nil {
		stopProjectRequest = new(StopProjectRequest)
	}
	err := c.DoAction(stopProjectRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) StatusProject(statusProjectRequest *StatusProjectRequest) (cms_StatusProjectResponse.StatusProjectResponse, error) {
	var resObj cms_StatusProjectResponse.StatusProjectResponse

	if statusProjectRequest == nil {
		statusProjectRequest = new(StatusProjectRequest)
	}
	err := c.DoAction(statusProjectRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) StartProject(startProjectRequest *StartProjectRequest) (cms_StartProjectResponse.StartProjectResponse, error) {
	var resObj cms_StartProjectResponse.StartProjectResponse

	if startProjectRequest == nil {
		startProjectRequest = new(StartProjectRequest)
	}
	err := c.DoAction(startProjectRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) BatchCreateSqlMetrics(batchCreateSqlMetricsRequest *BatchCreateSqlMetricsRequest) (cms_BatchCreateSqlMetricsResponse.BatchCreateSqlMetricsResponse, error) {
	var resObj cms_BatchCreateSqlMetricsResponse.BatchCreateSqlMetricsResponse

	if batchCreateSqlMetricsRequest == nil {
		batchCreateSqlMetricsRequest = new(BatchCreateSqlMetricsRequest)
	}
	err := c.DoAction(batchCreateSqlMetricsRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) UpdateMetricStream(updateMetricStreamRequest *UpdateMetricStreamRequest) (cms_UpdateMetricStreamResponse.UpdateMetricStreamResponse, error) {
	var resObj cms_UpdateMetricStreamResponse.UpdateMetricStreamResponse

	if updateMetricStreamRequest == nil {
		updateMetricStreamRequest = new(UpdateMetricStreamRequest)
	}
	err := c.DoAction(updateMetricStreamRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) StopMetricStream(stopMetricStreamRequest *StopMetricStreamRequest) (cms_StopMetricStreamResponse.StopMetricStreamResponse, error) {
	var resObj cms_StopMetricStreamResponse.StopMetricStreamResponse

	if stopMetricStreamRequest == nil {
		stopMetricStreamRequest = new(StopMetricStreamRequest)
	}
	err := c.DoAction(stopMetricStreamRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) GetMetricStreamStatus(getMetricStreamStatusRequest *GetMetricStreamStatusRequest) (cms_GetMetricStreamStatusResponse.GetMetricStreamStatusResponse, error) {
	var resObj cms_GetMetricStreamStatusResponse.GetMetricStreamStatusResponse

	if getMetricStreamStatusRequest == nil {
		getMetricStreamStatusRequest = new(GetMetricStreamStatusRequest)
	}
	err := c.DoAction(getMetricStreamStatusRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) StartMetricStream(startMetricStreamRequest *StartMetricStreamRequest) (cms_StartMetricStreamResponse.StartMetricStreamResponse, error) {
	var resObj cms_StartMetricStreamResponse.StartMetricStreamResponse

	if startMetricStreamRequest == nil {
		startMetricStreamRequest = new(StartMetricStreamRequest)
	}
	err := c.DoAction(startMetricStreamRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) GetMetricStreamMeta(getMetricStreamMetaRequest *GetMetricStreamMetaRequest) (cms_GetMetricStreamMetaResponse.GetMetricStreamMetaResponse, error) {
	var resObj cms_GetMetricStreamMetaResponse.GetMetricStreamMetaResponse

	if getMetricStreamMetaRequest == nil {
		getMetricStreamMetaRequest = new(GetMetricStreamMetaRequest)
	}
	err := c.DoAction(getMetricStreamMetaRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) ListMetricStream(listMetricStreamRequest *ListMetricStreamRequest) (cms_ListMetricStreamResponse.ListMetricStreamResponse, error) {
	var resObj cms_ListMetricStreamResponse.ListMetricStreamResponse

	if listMetricStreamRequest == nil {
		listMetricStreamRequest = new(ListMetricStreamRequest)
	}
	err := c.DoAction(listMetricStreamRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) GetMetricStream(getMetricStreamRequest *GetMetricStreamRequest) (cms_GetMetricStreamResponse.GetMetricStreamResponse, error) {
	var resObj cms_GetMetricStreamResponse.GetMetricStreamResponse

	if getMetricStreamRequest == nil {
		getMetricStreamRequest = new(GetMetricStreamRequest)
	}
	err := c.DoAction(getMetricStreamRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) DeleteMetricStream(deleteMetricStreamRequest *DeleteMetricStreamRequest) (cms_DeleteMetricStreamResponse.DeleteMetricStreamResponse, error) {
	var resObj cms_DeleteMetricStreamResponse.DeleteMetricStreamResponse

	if deleteMetricStreamRequest == nil {
		deleteMetricStreamRequest = new(DeleteMetricStreamRequest)
	}
	err := c.DoAction(deleteMetricStreamRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) CreateMetricStream(createMetricStreamRequest *CreateMetricStreamRequest) (cms_CreateMetricStreamResponse.CreateMetricStreamResponse, error) {
	var resObj cms_CreateMetricStreamResponse.CreateMetricStreamResponse

	if createMetricStreamRequest == nil {
		createMetricStreamRequest = new(CreateMetricStreamRequest)
	}
	err := c.DoAction(createMetricStreamRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) DescribeMetricDatum(describeMetricDatumRequest *DescribeMetricDatumRequest) (cms_DescribeMetricDatumResponse.DescribeMetricDatumResponse, error) {
	var resObj cms_DescribeMetricDatumResponse.DescribeMetricDatumResponse

	if describeMetricDatumRequest == nil {
		describeMetricDatumRequest = new(DescribeMetricDatumRequest)
	}
	err := c.DoAction(describeMetricDatumRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) ListDimTableData(listDimTableDataRequest *ListDimTableDataRequest) (cms_ListDimTableDataResponse.ListDimTableDataResponse, error) {
	var resObj cms_ListDimTableDataResponse.ListDimTableDataResponse

	if listDimTableDataRequest == nil {
		listDimTableDataRequest = new(ListDimTableDataRequest)
	}
	err := c.DoAction(listDimTableDataRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) PutDimTableData(putDimTableDataRequest *PutDimTableDataRequest) (cms_PutDimTableDataResponse.PutDimTableDataResponse, error) {
	var resObj cms_PutDimTableDataResponse.PutDimTableDataResponse

	if putDimTableDataRequest == nil {
		putDimTableDataRequest = new(PutDimTableDataRequest)
	}
	err := c.DoAction(putDimTableDataRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) UpdateProject(updateProjectRequest *UpdateProjectRequest) (cms_UpdateProjectResponse.UpdateProjectResponse, error) {
	var resObj cms_UpdateProjectResponse.UpdateProjectResponse

	if updateProjectRequest == nil {
		updateProjectRequest = new(UpdateProjectRequest)
	}
	err := c.DoAction(updateProjectRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) ListProject(listProjectRequest *ListProjectRequest) (cms_ListProjectResponse.ListProjectResponse, error) {
	var resObj cms_ListProjectResponse.ListProjectResponse

	if listProjectRequest == nil {
		listProjectRequest = new(ListProjectRequest)
	}
	err := c.DoAction(listProjectRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) GetProject(getProjectRequest *GetProjectRequest) (cms_GetProjectResponse.GetProjectResponse, error) {
	var resObj cms_GetProjectResponse.GetProjectResponse

	if getProjectRequest == nil {
		getProjectRequest = new(GetProjectRequest)
	}
	err := c.DoAction(getProjectRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) DeleteProject(deleteProjectRequest *DeleteProjectRequest) (cms_DeleteProjectResponse.DeleteProjectResponse, error) {
	var resObj cms_DeleteProjectResponse.DeleteProjectResponse

	if deleteProjectRequest == nil {
		deleteProjectRequest = new(DeleteProjectRequest)
	}
	err := c.DoAction(deleteProjectRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) CreateProject(createProjectRequest *CreateProjectRequest) (cms_CreateProjectResponse.CreateProjectResponse, error) {
	var resObj cms_CreateProjectResponse.CreateProjectResponse

	if createProjectRequest == nil {
		createProjectRequest = new(CreateProjectRequest)
	}
	err := c.DoAction(createProjectRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) UpdateMetrics(updateMetricsRequest *UpdateMetricsRequest) (cms_UpdateMetricsResponse.UpdateMetricsResponse, error) {
	var resObj cms_UpdateMetricsResponse.UpdateMetricsResponse

	if updateMetricsRequest == nil {
		updateMetricsRequest = new(UpdateMetricsRequest)
	}
	err := c.DoAction(updateMetricsRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) UpdateSQLMetrics(updateSQLMetricsRequest *UpdateSQLMetricsRequest) (cms_UpdateSQLMetricsResponse.UpdateSQLMetricsResponse, error) {
	var resObj cms_UpdateSQLMetricsResponse.UpdateSQLMetricsResponse

	if updateSQLMetricsRequest == nil {
		updateSQLMetricsRequest = new(UpdateSQLMetricsRequest)
	}
	err := c.DoAction(updateSQLMetricsRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) ListSQLMetrics(listSQLMetricsRequest *ListSQLMetricsRequest) (cms_ListSQLMetricsResponse.ListSQLMetricsResponse, error) {
	var resObj cms_ListSQLMetricsResponse.ListSQLMetricsResponse

	if listSQLMetricsRequest == nil {
		listSQLMetricsRequest = new(ListSQLMetricsRequest)
	}
	err := c.DoAction(listSQLMetricsRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) GetSQLMetrics(getSQLMetricsRequest *GetSQLMetricsRequest) (cms_GetSQLMetricsResponse.GetSQLMetricsResponse, error) {
	var resObj cms_GetSQLMetricsResponse.GetSQLMetricsResponse

	if getSQLMetricsRequest == nil {
		getSQLMetricsRequest = new(GetSQLMetricsRequest)
	}
	err := c.DoAction(getSQLMetricsRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) CreateSQLMetrics(createSQLMetricsRequest *CreateSQLMetricsRequest) (cms_CreateSQLMetricsResponse.CreateSQLMetricsResponse, error) {
	var resObj cms_CreateSQLMetricsResponse.CreateSQLMetricsResponse

	if createSQLMetricsRequest == nil {
		createSQLMetricsRequest = new(CreateSQLMetricsRequest)
	}
	err := c.DoAction(createSQLMetricsRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) ListMetricsPlan(listMetricsPlanRequest *ListMetricsPlanRequest) (cms_ListMetricsPlanResponse.ListMetricsPlanResponse, error) {
	var resObj cms_ListMetricsPlanResponse.ListMetricsPlanResponse

	if listMetricsPlanRequest == nil {
		listMetricsPlanRequest = new(ListMetricsPlanRequest)
	}
	err := c.DoAction(listMetricsPlanRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) GetMetricsMeta(getMetricsMetaRequest *GetMetricsMetaRequest) (cms_GetMetricsMetaResponse.GetMetricsMetaResponse, error) {
	var resObj cms_GetMetricsMetaResponse.GetMetricsMetaResponse

	if getMetricsMetaRequest == nil {
		getMetricsMetaRequest = new(GetMetricsMetaRequest)
	}
	err := c.DoAction(getMetricsMetaRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) ListMetrics(listMetricsRequest *ListMetricsRequest) (cms_ListMetricsResponse.ListMetricsResponse, error) {
	var resObj cms_ListMetricsResponse.ListMetricsResponse

	if listMetricsRequest == nil {
		listMetricsRequest = new(ListMetricsRequest)
	}
	err := c.DoAction(listMetricsRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) GetMetrics(getMetricsRequest *GetMetricsRequest) (cms_GetMetricsResponse.GetMetricsResponse, error) {
	var resObj cms_GetMetricsResponse.GetMetricsResponse

	if getMetricsRequest == nil {
		getMetricsRequest = new(GetMetricsRequest)
	}
	err := c.DoAction(getMetricsRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) DeleteMetrics(deleteMetricsRequest *DeleteMetricsRequest) (cms_DeleteMetricsResponse.DeleteMetricsResponse, error) {
	var resObj cms_DeleteMetricsResponse.DeleteMetricsResponse

	if deleteMetricsRequest == nil {
		deleteMetricsRequest = new(DeleteMetricsRequest)
	}
	err := c.DoAction(deleteMetricsRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) CreateMetrics(createMetricsRequest *CreateMetricsRequest) (cms_CreateMetricsResponse.CreateMetricsResponse, error) {
	var resObj cms_CreateMetricsResponse.CreateMetricsResponse

	if createMetricsRequest == nil {
		createMetricsRequest = new(CreateMetricsRequest)
	}
	err := c.DoAction(createMetricsRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) BatchCreateMetrics(batchCreateMetricsRequest *BatchCreateMetricsRequest) (cms_BatchCreateMetricsResponse.BatchCreateMetricsResponse, error) {
	var resObj cms_BatchCreateMetricsResponse.BatchCreateMetricsResponse

	if batchCreateMetricsRequest == nil {
		batchCreateMetricsRequest = new(BatchCreateMetricsRequest)
	}
	err := c.DoAction(batchCreateMetricsRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) DescribeMetricList(describeMetricListRequest *DescribeMetricListRequest) (cms_DescribeMetricListResponse.DescribeMetricListResponse, error) {
	var resObj cms_DescribeMetricListResponse.DescribeMetricListResponse

	if describeMetricListRequest == nil {
		describeMetricListRequest = new(DescribeMetricListRequest)
	}
	err := c.DoAction(describeMetricListRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) UpdateDimTable(updateDimTableRequest *UpdateDimTableRequest) (cms_UpdateDimTableResponse.UpdateDimTableResponse, error) {
	var resObj cms_UpdateDimTableResponse.UpdateDimTableResponse

	if updateDimTableRequest == nil {
		updateDimTableRequest = new(UpdateDimTableRequest)
	}
	err := c.DoAction(updateDimTableRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) ListDimTable(listDimTableRequest *ListDimTableRequest) (cms_ListDimTableResponse.ListDimTableResponse, error) {
	var resObj cms_ListDimTableResponse.ListDimTableResponse

	if listDimTableRequest == nil {
		listDimTableRequest = new(ListDimTableRequest)
	}
	err := c.DoAction(listDimTableRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) GetDimTable(getDimTableRequest *GetDimTableRequest) (cms_GetDimTableResponse.GetDimTableResponse, error) {
	var resObj cms_GetDimTableResponse.GetDimTableResponse

	if getDimTableRequest == nil {
		getDimTableRequest = new(GetDimTableRequest)
	}
	err := c.DoAction(getDimTableRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) DeleteDimTable(deleteDimTableRequest *DeleteDimTableRequest) (cms_DeleteDimTableResponse.DeleteDimTableResponse, error) {
	var resObj cms_DeleteDimTableResponse.DeleteDimTableResponse

	if deleteDimTableRequest == nil {
		deleteDimTableRequest = new(DeleteDimTableRequest)
	}
	err := c.DoAction(deleteDimTableRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) CreateDimTable(createDimTableRequest *CreateDimTableRequest) (cms_CreateDimTableResponse.CreateDimTableResponse, error) {
	var resObj cms_CreateDimTableResponse.CreateDimTableResponse

	if createDimTableRequest == nil {
		createDimTableRequest = new(CreateDimTableRequest)
	}
	err := c.DoAction(createDimTableRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) DeleteDimTableData(deleteDimTableDataRequest *DeleteDimTableDataRequest) (cms_DeleteDimTableDataResponse.DeleteDimTableDataResponse, error) {
	var resObj cms_DeleteDimTableDataResponse.DeleteDimTableDataResponse

	if deleteDimTableDataRequest == nil {
		deleteDimTableDataRequest = new(DeleteDimTableDataRequest)
	}
	err := c.DoAction(deleteDimTableDataRequest, &resObj)
	return resObj, err
}

func (c *CmsClient) BatchPutDimTableData(batchPutDimTableDataRequest *BatchPutDimTableDataRequest) (cms_BatchPutDimTableDataResponse.BatchPutDimTableDataResponse, error) {
	var resObj cms_BatchPutDimTableDataResponse.BatchPutDimTableDataResponse

	if batchPutDimTableDataRequest == nil {
		batchPutDimTableDataRequest = new(BatchPutDimTableDataRequest)
	}
	err := c.DoAction(batchPutDimTableDataRequest, &resObj)
	return resObj, err
}
