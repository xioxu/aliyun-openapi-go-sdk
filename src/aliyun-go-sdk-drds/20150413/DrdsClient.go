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

package drds

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-drds/20150413/DescribeDrdsDBIpWhiteListResponse"

	"aliyun-go-sdk-drds/20150413/ModifyDrdsIpWhiteListResponse"

	"aliyun-go-sdk-drds/20150413/RemoveDrdsInstanceResponse"

	"aliyun-go-sdk-drds/20150413/DescribeCreateDrdsInstanceStatusResponse"

	"aliyun-go-sdk-drds/20150413/ModifyDrdsInstanceDescriptionResponse"

	"aliyun-go-sdk-drds/20150413/ModifyDrdsDBPasswdResponse"

	"aliyun-go-sdk-drds/20150413/ListUnCompleteTasksResponse"

	"aliyun-go-sdk-drds/20150413/DropTablesResponse"

	"aliyun-go-sdk-drds/20150413/DropIndexesResponse"

	"aliyun-go-sdk-drds/20150413/DescribeDrdsInstancesResponse"

	"aliyun-go-sdk-drds/20150413/DescribeDrdsInstanceResponse"

	"aliyun-go-sdk-drds/20150413/DescribeDrdsDBsResponse"

	"aliyun-go-sdk-drds/20150413/DescribeDrdsDBResponse"

	"aliyun-go-sdk-drds/20150413/DescribeDDLTaskResponse"

	"aliyun-go-sdk-drds/20150413/DeleteDrdsDBResponse"

	"aliyun-go-sdk-drds/20150413/CreateTableResponse"

	"aliyun-go-sdk-drds/20150413/CreateIndexResponse"

	"aliyun-go-sdk-drds/20150413/CreateDrdsInstanceResponse"

	"aliyun-go-sdk-drds/20150413/CreateDrdsDBResponse"

	"aliyun-go-sdk-drds/20150413/CancelDDLTaskResponse"

	"aliyun-go-sdk-drds/20150413/AlterTableResponse"
)

type DrdsClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *DrdsClient {
	p := &DrdsClient{}
	p.Version = "2015-04-13"
	p.ProductName = "Drds"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *DrdsClient) DescribeDrdsDBIpWhiteList(describeDrdsDBIpWhiteListRequest *DescribeDrdsDBIpWhiteListRequest) (drds_DescribeDrdsDBIpWhiteListResponse.DescribeDrdsDBIpWhiteListResponse, error) {
	var resObj drds_DescribeDrdsDBIpWhiteListResponse.DescribeDrdsDBIpWhiteListResponse

	if describeDrdsDBIpWhiteListRequest == nil {
		describeDrdsDBIpWhiteListRequest = new(DescribeDrdsDBIpWhiteListRequest)
	}
	err := c.DoAction(describeDrdsDBIpWhiteListRequest, &resObj)
	return resObj, err
}

func (c *DrdsClient) ModifyDrdsIpWhiteList(modifyDrdsIpWhiteListRequest *ModifyDrdsIpWhiteListRequest) (drds_ModifyDrdsIpWhiteListResponse.ModifyDrdsIpWhiteListResponse, error) {
	var resObj drds_ModifyDrdsIpWhiteListResponse.ModifyDrdsIpWhiteListResponse

	if modifyDrdsIpWhiteListRequest == nil {
		modifyDrdsIpWhiteListRequest = new(ModifyDrdsIpWhiteListRequest)
	}
	err := c.DoAction(modifyDrdsIpWhiteListRequest, &resObj)
	return resObj, err
}

func (c *DrdsClient) RemoveDrdsInstance(removeDrdsInstanceRequest *RemoveDrdsInstanceRequest) (drds_RemoveDrdsInstanceResponse.RemoveDrdsInstanceResponse, error) {
	var resObj drds_RemoveDrdsInstanceResponse.RemoveDrdsInstanceResponse

	if removeDrdsInstanceRequest == nil {
		removeDrdsInstanceRequest = new(RemoveDrdsInstanceRequest)
	}
	err := c.DoAction(removeDrdsInstanceRequest, &resObj)
	return resObj, err
}

func (c *DrdsClient) DescribeCreateDrdsInstanceStatus(describeCreateDrdsInstanceStatusRequest *DescribeCreateDrdsInstanceStatusRequest) (drds_DescribeCreateDrdsInstanceStatusResponse.DescribeCreateDrdsInstanceStatusResponse, error) {
	var resObj drds_DescribeCreateDrdsInstanceStatusResponse.DescribeCreateDrdsInstanceStatusResponse

	if describeCreateDrdsInstanceStatusRequest == nil {
		describeCreateDrdsInstanceStatusRequest = new(DescribeCreateDrdsInstanceStatusRequest)
	}
	err := c.DoAction(describeCreateDrdsInstanceStatusRequest, &resObj)
	return resObj, err
}

func (c *DrdsClient) ModifyDrdsInstanceDescription(modifyDrdsInstanceDescriptionRequest *ModifyDrdsInstanceDescriptionRequest) (drds_ModifyDrdsInstanceDescriptionResponse.ModifyDrdsInstanceDescriptionResponse, error) {
	var resObj drds_ModifyDrdsInstanceDescriptionResponse.ModifyDrdsInstanceDescriptionResponse

	if modifyDrdsInstanceDescriptionRequest == nil {
		modifyDrdsInstanceDescriptionRequest = new(ModifyDrdsInstanceDescriptionRequest)
	}
	err := c.DoAction(modifyDrdsInstanceDescriptionRequest, &resObj)
	return resObj, err
}

func (c *DrdsClient) ModifyDrdsDBPasswd(modifyDrdsDBPasswdRequest *ModifyDrdsDBPasswdRequest) (drds_ModifyDrdsDBPasswdResponse.ModifyDrdsDBPasswdResponse, error) {
	var resObj drds_ModifyDrdsDBPasswdResponse.ModifyDrdsDBPasswdResponse

	if modifyDrdsDBPasswdRequest == nil {
		modifyDrdsDBPasswdRequest = new(ModifyDrdsDBPasswdRequest)
	}
	err := c.DoAction(modifyDrdsDBPasswdRequest, &resObj)
	return resObj, err
}

func (c *DrdsClient) ListUnCompleteTasks(listUnCompleteTasksRequest *ListUnCompleteTasksRequest) (drds_ListUnCompleteTasksResponse.ListUnCompleteTasksResponse, error) {
	var resObj drds_ListUnCompleteTasksResponse.ListUnCompleteTasksResponse

	if listUnCompleteTasksRequest == nil {
		listUnCompleteTasksRequest = new(ListUnCompleteTasksRequest)
	}
	err := c.DoAction(listUnCompleteTasksRequest, &resObj)
	return resObj, err
}

func (c *DrdsClient) DropTables(dropTablesRequest *DropTablesRequest) (drds_DropTablesResponse.DropTablesResponse, error) {
	var resObj drds_DropTablesResponse.DropTablesResponse

	if dropTablesRequest == nil {
		dropTablesRequest = new(DropTablesRequest)
	}
	err := c.DoAction(dropTablesRequest, &resObj)
	return resObj, err
}

func (c *DrdsClient) DropIndexes(dropIndexesRequest *DropIndexesRequest) (drds_DropIndexesResponse.DropIndexesResponse, error) {
	var resObj drds_DropIndexesResponse.DropIndexesResponse

	if dropIndexesRequest == nil {
		dropIndexesRequest = new(DropIndexesRequest)
	}
	err := c.DoAction(dropIndexesRequest, &resObj)
	return resObj, err
}

func (c *DrdsClient) DescribeDrdsInstances(describeDrdsInstancesRequest *DescribeDrdsInstancesRequest) (drds_DescribeDrdsInstancesResponse.DescribeDrdsInstancesResponse, error) {
	var resObj drds_DescribeDrdsInstancesResponse.DescribeDrdsInstancesResponse

	if describeDrdsInstancesRequest == nil {
		describeDrdsInstancesRequest = new(DescribeDrdsInstancesRequest)
	}
	err := c.DoAction(describeDrdsInstancesRequest, &resObj)
	return resObj, err
}

func (c *DrdsClient) DescribeDrdsInstance(describeDrdsInstanceRequest *DescribeDrdsInstanceRequest) (drds_DescribeDrdsInstanceResponse.DescribeDrdsInstanceResponse, error) {
	var resObj drds_DescribeDrdsInstanceResponse.DescribeDrdsInstanceResponse

	if describeDrdsInstanceRequest == nil {
		describeDrdsInstanceRequest = new(DescribeDrdsInstanceRequest)
	}
	err := c.DoAction(describeDrdsInstanceRequest, &resObj)
	return resObj, err
}

func (c *DrdsClient) DescribeDrdsDBs(describeDrdsDBsRequest *DescribeDrdsDBsRequest) (drds_DescribeDrdsDBsResponse.DescribeDrdsDBsResponse, error) {
	var resObj drds_DescribeDrdsDBsResponse.DescribeDrdsDBsResponse

	if describeDrdsDBsRequest == nil {
		describeDrdsDBsRequest = new(DescribeDrdsDBsRequest)
	}
	err := c.DoAction(describeDrdsDBsRequest, &resObj)
	return resObj, err
}

func (c *DrdsClient) DescribeDrdsDB(describeDrdsDBRequest *DescribeDrdsDBRequest) (drds_DescribeDrdsDBResponse.DescribeDrdsDBResponse, error) {
	var resObj drds_DescribeDrdsDBResponse.DescribeDrdsDBResponse

	if describeDrdsDBRequest == nil {
		describeDrdsDBRequest = new(DescribeDrdsDBRequest)
	}
	err := c.DoAction(describeDrdsDBRequest, &resObj)
	return resObj, err
}

func (c *DrdsClient) DescribeDDLTask(describeDDLTaskRequest *DescribeDDLTaskRequest) (drds_DescribeDDLTaskResponse.DescribeDDLTaskResponse, error) {
	var resObj drds_DescribeDDLTaskResponse.DescribeDDLTaskResponse

	if describeDDLTaskRequest == nil {
		describeDDLTaskRequest = new(DescribeDDLTaskRequest)
	}
	err := c.DoAction(describeDDLTaskRequest, &resObj)
	return resObj, err
}

func (c *DrdsClient) DeleteDrdsDB(deleteDrdsDBRequest *DeleteDrdsDBRequest) (drds_DeleteDrdsDBResponse.DeleteDrdsDBResponse, error) {
	var resObj drds_DeleteDrdsDBResponse.DeleteDrdsDBResponse

	if deleteDrdsDBRequest == nil {
		deleteDrdsDBRequest = new(DeleteDrdsDBRequest)
	}
	err := c.DoAction(deleteDrdsDBRequest, &resObj)
	return resObj, err
}

func (c *DrdsClient) CreateTable(createTableRequest *CreateTableRequest) (drds_CreateTableResponse.CreateTableResponse, error) {
	var resObj drds_CreateTableResponse.CreateTableResponse

	if createTableRequest == nil {
		createTableRequest = new(CreateTableRequest)
	}
	err := c.DoAction(createTableRequest, &resObj)
	return resObj, err
}

func (c *DrdsClient) CreateIndex(createIndexRequest *CreateIndexRequest) (drds_CreateIndexResponse.CreateIndexResponse, error) {
	var resObj drds_CreateIndexResponse.CreateIndexResponse

	if createIndexRequest == nil {
		createIndexRequest = new(CreateIndexRequest)
	}
	err := c.DoAction(createIndexRequest, &resObj)
	return resObj, err
}

func (c *DrdsClient) CreateDrdsInstance(createDrdsInstanceRequest *CreateDrdsInstanceRequest) (drds_CreateDrdsInstanceResponse.CreateDrdsInstanceResponse, error) {
	var resObj drds_CreateDrdsInstanceResponse.CreateDrdsInstanceResponse

	if createDrdsInstanceRequest == nil {
		createDrdsInstanceRequest = new(CreateDrdsInstanceRequest)
	}
	err := c.DoAction(createDrdsInstanceRequest, &resObj)
	return resObj, err
}

func (c *DrdsClient) CreateDrdsDB(createDrdsDBRequest *CreateDrdsDBRequest) (drds_CreateDrdsDBResponse.CreateDrdsDBResponse, error) {
	var resObj drds_CreateDrdsDBResponse.CreateDrdsDBResponse

	if createDrdsDBRequest == nil {
		createDrdsDBRequest = new(CreateDrdsDBRequest)
	}
	err := c.DoAction(createDrdsDBRequest, &resObj)
	return resObj, err
}

func (c *DrdsClient) CancelDDLTask(cancelDDLTaskRequest *CancelDDLTaskRequest) (drds_CancelDDLTaskResponse.CancelDDLTaskResponse, error) {
	var resObj drds_CancelDDLTaskResponse.CancelDDLTaskResponse

	if cancelDDLTaskRequest == nil {
		cancelDDLTaskRequest = new(CancelDDLTaskRequest)
	}
	err := c.DoAction(cancelDDLTaskRequest, &resObj)
	return resObj, err
}

func (c *DrdsClient) AlterTable(alterTableRequest *AlterTableRequest) (drds_AlterTableResponse.AlterTableResponse, error) {
	var resObj drds_AlterTableResponse.AlterTableResponse

	if alterTableRequest == nil {
		alterTableRequest = new(AlterTableRequest)
	}
	err := c.DoAction(alterTableRequest, &resObj)
	return resObj, err
}
