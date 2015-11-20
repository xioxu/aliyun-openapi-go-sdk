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

package emr

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-emr/20150910/ModifyExecutePlanWithClusterResponse"

	"aliyun-go-sdk-emr/20150910/ModifyExecutePlanResponse"

	"aliyun-go-sdk-emr/20150910/ListExecutePlanJobsResponse"

	"aliyun-go-sdk-emr/20150910/CreateExecutePlanResponse"

	"aliyun-go-sdk-emr/20150910/CreateClusterResponse"

	"aliyun-go-sdk-emr/20150910/CreateJobResponse"

	"aliyun-go-sdk-emr/20150910/CreateExecutePlanWithClusterResponse"

	"aliyun-go-sdk-emr/20150910/DeleteExecutePlanResponse"

	"aliyun-go-sdk-emr/20150910/DeleteJobResponse"

	"aliyun-go-sdk-emr/20150910/DescribeJobResponse"

	"aliyun-go-sdk-emr/20150910/DescribeClusterResponse"

	"aliyun-go-sdk-emr/20150910/ListConfigTypeResponse"

	"aliyun-go-sdk-emr/20150910/ListClustersResponse"

	"aliyun-go-sdk-emr/20150910/ListExecutePlanExecuteRecordsResponse"

	"aliyun-go-sdk-emr/20150910/ListExecutePlanExecuteRecordNodesResponse"

	"aliyun-go-sdk-emr/20150910/ListExecutePlansResponse"

	"aliyun-go-sdk-emr/20150910/ListRegionsResponse"

	"aliyun-go-sdk-emr/20150910/ListJobsResponse"

	"aliyun-go-sdk-emr/20150910/ModifyJobResponse"

	"aliyun-go-sdk-emr/20150910/ModifyClusterNameResponse"

	"aliyun-go-sdk-emr/20150910/StopExecutePlanResponse"

	"aliyun-go-sdk-emr/20150910/RunExecutePlanResponse"

	"aliyun-go-sdk-emr/20150910/ReleaseClusterResponse"

	"aliyun-go-sdk-emr/20150910/ListExecutePlanNodeInstancesResponse"

	"aliyun-go-sdk-emr/20150910/ResumeExecutePlanResponse"

	"aliyun-go-sdk-emr/20150910/JobResourceResponse"

	"aliyun-go-sdk-emr/20150910/DescribeExecutePlanResponse"

	"aliyun-go-sdk-emr/20150910/KillExecutePlanRecordNodeResponse"
)

type EmrClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *EmrClient {
	p := &EmrClient{}
	p.Version = "2015-09-10"
	p.ProductName = "Emr"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *EmrClient) ModifyExecutePlanWithCluster(modifyExecutePlanWithClusterRequest *ModifyExecutePlanWithClusterRequest) (emr_ModifyExecutePlanWithClusterResponse.ModifyExecutePlanWithClusterResponse, error) {
	var resObj emr_ModifyExecutePlanWithClusterResponse.ModifyExecutePlanWithClusterResponse

	if modifyExecutePlanWithClusterRequest == nil {
		modifyExecutePlanWithClusterRequest = new(ModifyExecutePlanWithClusterRequest)
	}
	err := c.DoAction(modifyExecutePlanWithClusterRequest, &resObj)
	return resObj, err
}

func (c *EmrClient) ModifyExecutePlan(modifyExecutePlanRequest *ModifyExecutePlanRequest) (emr_ModifyExecutePlanResponse.ModifyExecutePlanResponse, error) {
	var resObj emr_ModifyExecutePlanResponse.ModifyExecutePlanResponse

	if modifyExecutePlanRequest == nil {
		modifyExecutePlanRequest = new(ModifyExecutePlanRequest)
	}
	err := c.DoAction(modifyExecutePlanRequest, &resObj)
	return resObj, err
}

func (c *EmrClient) ListExecutePlanJobs(listExecutePlanJobsRequest *ListExecutePlanJobsRequest) (emr_ListExecutePlanJobsResponse.ListExecutePlanJobsResponse, error) {
	var resObj emr_ListExecutePlanJobsResponse.ListExecutePlanJobsResponse

	if listExecutePlanJobsRequest == nil {
		listExecutePlanJobsRequest = new(ListExecutePlanJobsRequest)
	}
	err := c.DoAction(listExecutePlanJobsRequest, &resObj)
	return resObj, err
}

func (c *EmrClient) CreateExecutePlan(createExecutePlanRequest *CreateExecutePlanRequest) (emr_CreateExecutePlanResponse.CreateExecutePlanResponse, error) {
	var resObj emr_CreateExecutePlanResponse.CreateExecutePlanResponse

	if createExecutePlanRequest == nil {
		createExecutePlanRequest = new(CreateExecutePlanRequest)
	}
	err := c.DoAction(createExecutePlanRequest, &resObj)
	return resObj, err
}

func (c *EmrClient) CreateCluster(createClusterRequest *CreateClusterRequest) (emr_CreateClusterResponse.CreateClusterResponse, error) {
	var resObj emr_CreateClusterResponse.CreateClusterResponse

	if createClusterRequest == nil {
		createClusterRequest = new(CreateClusterRequest)
	}
	err := c.DoAction(createClusterRequest, &resObj)
	return resObj, err
}

func (c *EmrClient) CreateJob(createJobRequest *CreateJobRequest) (emr_CreateJobResponse.CreateJobResponse, error) {
	var resObj emr_CreateJobResponse.CreateJobResponse

	if createJobRequest == nil {
		createJobRequest = new(CreateJobRequest)
	}
	err := c.DoAction(createJobRequest, &resObj)
	return resObj, err
}

func (c *EmrClient) CreateExecutePlanWithCluster(createExecutePlanWithClusterRequest *CreateExecutePlanWithClusterRequest) (emr_CreateExecutePlanWithClusterResponse.CreateExecutePlanWithClusterResponse, error) {
	var resObj emr_CreateExecutePlanWithClusterResponse.CreateExecutePlanWithClusterResponse

	if createExecutePlanWithClusterRequest == nil {
		createExecutePlanWithClusterRequest = new(CreateExecutePlanWithClusterRequest)
	}
	err := c.DoAction(createExecutePlanWithClusterRequest, &resObj)
	return resObj, err
}

func (c *EmrClient) DeleteExecutePlan(deleteExecutePlanRequest *DeleteExecutePlanRequest) (emr_DeleteExecutePlanResponse.DeleteExecutePlanResponse, error) {
	var resObj emr_DeleteExecutePlanResponse.DeleteExecutePlanResponse

	if deleteExecutePlanRequest == nil {
		deleteExecutePlanRequest = new(DeleteExecutePlanRequest)
	}
	err := c.DoAction(deleteExecutePlanRequest, &resObj)
	return resObj, err
}

func (c *EmrClient) DeleteJob(deleteJobRequest *DeleteJobRequest) (emr_DeleteJobResponse.DeleteJobResponse, error) {
	var resObj emr_DeleteJobResponse.DeleteJobResponse

	if deleteJobRequest == nil {
		deleteJobRequest = new(DeleteJobRequest)
	}
	err := c.DoAction(deleteJobRequest, &resObj)
	return resObj, err
}

func (c *EmrClient) DescribeJob(describeJobRequest *DescribeJobRequest) (emr_DescribeJobResponse.DescribeJobResponse, error) {
	var resObj emr_DescribeJobResponse.DescribeJobResponse

	if describeJobRequest == nil {
		describeJobRequest = new(DescribeJobRequest)
	}
	err := c.DoAction(describeJobRequest, &resObj)
	return resObj, err
}

func (c *EmrClient) DescribeCluster(describeClusterRequest *DescribeClusterRequest) (emr_DescribeClusterResponse.DescribeClusterResponse, error) {
	var resObj emr_DescribeClusterResponse.DescribeClusterResponse

	if describeClusterRequest == nil {
		describeClusterRequest = new(DescribeClusterRequest)
	}
	err := c.DoAction(describeClusterRequest, &resObj)
	return resObj, err
}

func (c *EmrClient) ListConfigType(listConfigTypeRequest *ListConfigTypeRequest) (emr_ListConfigTypeResponse.ListConfigTypeResponse, error) {
	var resObj emr_ListConfigTypeResponse.ListConfigTypeResponse

	if listConfigTypeRequest == nil {
		listConfigTypeRequest = new(ListConfigTypeRequest)
	}
	err := c.DoAction(listConfigTypeRequest, &resObj)
	return resObj, err
}

func (c *EmrClient) ListClusters(listClustersRequest *ListClustersRequest) (emr_ListClustersResponse.ListClustersResponse, error) {
	var resObj emr_ListClustersResponse.ListClustersResponse

	if listClustersRequest == nil {
		listClustersRequest = new(ListClustersRequest)
	}
	err := c.DoAction(listClustersRequest, &resObj)
	return resObj, err
}

func (c *EmrClient) ListExecutePlanExecuteRecords(listExecutePlanExecuteRecordsRequest *ListExecutePlanExecuteRecordsRequest) (emr_ListExecutePlanExecuteRecordsResponse.ListExecutePlanExecuteRecordsResponse, error) {
	var resObj emr_ListExecutePlanExecuteRecordsResponse.ListExecutePlanExecuteRecordsResponse

	if listExecutePlanExecuteRecordsRequest == nil {
		listExecutePlanExecuteRecordsRequest = new(ListExecutePlanExecuteRecordsRequest)
	}
	err := c.DoAction(listExecutePlanExecuteRecordsRequest, &resObj)
	return resObj, err
}

func (c *EmrClient) ListExecutePlanExecuteRecordNodes(listExecutePlanExecuteRecordNodesRequest *ListExecutePlanExecuteRecordNodesRequest) (emr_ListExecutePlanExecuteRecordNodesResponse.ListExecutePlanExecuteRecordNodesResponse, error) {
	var resObj emr_ListExecutePlanExecuteRecordNodesResponse.ListExecutePlanExecuteRecordNodesResponse

	if listExecutePlanExecuteRecordNodesRequest == nil {
		listExecutePlanExecuteRecordNodesRequest = new(ListExecutePlanExecuteRecordNodesRequest)
	}
	err := c.DoAction(listExecutePlanExecuteRecordNodesRequest, &resObj)
	return resObj, err
}

func (c *EmrClient) ListExecutePlans(listExecutePlansRequest *ListExecutePlansRequest) (emr_ListExecutePlansResponse.ListExecutePlansResponse, error) {
	var resObj emr_ListExecutePlansResponse.ListExecutePlansResponse

	if listExecutePlansRequest == nil {
		listExecutePlansRequest = new(ListExecutePlansRequest)
	}
	err := c.DoAction(listExecutePlansRequest, &resObj)
	return resObj, err
}

func (c *EmrClient) ListRegions(listRegionsRequest *ListRegionsRequest) (emr_ListRegionsResponse.ListRegionsResponse, error) {
	var resObj emr_ListRegionsResponse.ListRegionsResponse

	if listRegionsRequest == nil {
		listRegionsRequest = new(ListRegionsRequest)
	}
	err := c.DoAction(listRegionsRequest, &resObj)
	return resObj, err
}

func (c *EmrClient) ListJobs(listJobsRequest *ListJobsRequest) (emr_ListJobsResponse.ListJobsResponse, error) {
	var resObj emr_ListJobsResponse.ListJobsResponse

	if listJobsRequest == nil {
		listJobsRequest = new(ListJobsRequest)
	}
	err := c.DoAction(listJobsRequest, &resObj)
	return resObj, err
}

func (c *EmrClient) ModifyJob(modifyJobRequest *ModifyJobRequest) (emr_ModifyJobResponse.ModifyJobResponse, error) {
	var resObj emr_ModifyJobResponse.ModifyJobResponse

	if modifyJobRequest == nil {
		modifyJobRequest = new(ModifyJobRequest)
	}
	err := c.DoAction(modifyJobRequest, &resObj)
	return resObj, err
}

func (c *EmrClient) ModifyClusterName(modifyClusterNameRequest *ModifyClusterNameRequest) (emr_ModifyClusterNameResponse.ModifyClusterNameResponse, error) {
	var resObj emr_ModifyClusterNameResponse.ModifyClusterNameResponse

	if modifyClusterNameRequest == nil {
		modifyClusterNameRequest = new(ModifyClusterNameRequest)
	}
	err := c.DoAction(modifyClusterNameRequest, &resObj)
	return resObj, err
}

func (c *EmrClient) StopExecutePlan(stopExecutePlanRequest *StopExecutePlanRequest) (emr_StopExecutePlanResponse.StopExecutePlanResponse, error) {
	var resObj emr_StopExecutePlanResponse.StopExecutePlanResponse

	if stopExecutePlanRequest == nil {
		stopExecutePlanRequest = new(StopExecutePlanRequest)
	}
	err := c.DoAction(stopExecutePlanRequest, &resObj)
	return resObj, err
}

func (c *EmrClient) RunExecutePlan(runExecutePlanRequest *RunExecutePlanRequest) (emr_RunExecutePlanResponse.RunExecutePlanResponse, error) {
	var resObj emr_RunExecutePlanResponse.RunExecutePlanResponse

	if runExecutePlanRequest == nil {
		runExecutePlanRequest = new(RunExecutePlanRequest)
	}
	err := c.DoAction(runExecutePlanRequest, &resObj)
	return resObj, err
}

func (c *EmrClient) ReleaseCluster(releaseClusterRequest *ReleaseClusterRequest) (emr_ReleaseClusterResponse.ReleaseClusterResponse, error) {
	var resObj emr_ReleaseClusterResponse.ReleaseClusterResponse

	if releaseClusterRequest == nil {
		releaseClusterRequest = new(ReleaseClusterRequest)
	}
	err := c.DoAction(releaseClusterRequest, &resObj)
	return resObj, err
}

func (c *EmrClient) ListExecutePlanNodeInstances(listExecutePlanNodeInstancesRequest *ListExecutePlanNodeInstancesRequest) (emr_ListExecutePlanNodeInstancesResponse.ListExecutePlanNodeInstancesResponse, error) {
	var resObj emr_ListExecutePlanNodeInstancesResponse.ListExecutePlanNodeInstancesResponse

	if listExecutePlanNodeInstancesRequest == nil {
		listExecutePlanNodeInstancesRequest = new(ListExecutePlanNodeInstancesRequest)
	}
	err := c.DoAction(listExecutePlanNodeInstancesRequest, &resObj)
	return resObj, err
}

func (c *EmrClient) ResumeExecutePlan(resumeExecutePlanRequest *ResumeExecutePlanRequest) (emr_ResumeExecutePlanResponse.ResumeExecutePlanResponse, error) {
	var resObj emr_ResumeExecutePlanResponse.ResumeExecutePlanResponse

	if resumeExecutePlanRequest == nil {
		resumeExecutePlanRequest = new(ResumeExecutePlanRequest)
	}
	err := c.DoAction(resumeExecutePlanRequest, &resObj)
	return resObj, err
}

func (c *EmrClient) JobResource(jobResourceRequest *JobResourceRequest) (emr_JobResourceResponse.JobResourceResponse, error) {
	var resObj emr_JobResourceResponse.JobResourceResponse

	if jobResourceRequest == nil {
		jobResourceRequest = new(JobResourceRequest)
	}
	err := c.DoAction(jobResourceRequest, &resObj)
	return resObj, err
}

func (c *EmrClient) DescribeExecutePlan(describeExecutePlanRequest *DescribeExecutePlanRequest) (emr_DescribeExecutePlanResponse.DescribeExecutePlanResponse, error) {
	var resObj emr_DescribeExecutePlanResponse.DescribeExecutePlanResponse

	if describeExecutePlanRequest == nil {
		describeExecutePlanRequest = new(DescribeExecutePlanRequest)
	}
	err := c.DoAction(describeExecutePlanRequest, &resObj)
	return resObj, err
}

func (c *EmrClient) KillExecutePlanRecordNode(killExecutePlanRecordNodeRequest *KillExecutePlanRecordNodeRequest) (emr_KillExecutePlanRecordNodeResponse.KillExecutePlanRecordNodeResponse, error) {
	var resObj emr_KillExecutePlanRecordNodeResponse.KillExecutePlanRecordNodeResponse

	if killExecutePlanRecordNodeRequest == nil {
		killExecutePlanRecordNodeRequest = new(KillExecutePlanRecordNodeRequest)
	}
	err := c.DoAction(killExecutePlanRecordNodeRequest, &resObj)
	return resObj, err
}
