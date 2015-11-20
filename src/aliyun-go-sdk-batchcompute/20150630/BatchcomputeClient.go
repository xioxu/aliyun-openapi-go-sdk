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

package batchcompute

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-batchcompute/20150630/StopJobResponse"

	"aliyun-go-sdk-batchcompute/20150630/StartJobResponse"

	"aliyun-go-sdk-batchcompute/20150630/ReleaseJobResponse"

	"aliyun-go-sdk-batchcompute/20150630/PutJobResponse"

	"aliyun-go-sdk-batchcompute/20150630/PostJobResponse"

	"aliyun-go-sdk-batchcompute/20150630/ListSnapshotsResponse"

	"aliyun-go-sdk-batchcompute/20150630/ListJobsResponse"

	"aliyun-go-sdk-batchcompute/20150630/ListImagesResponse"

	"aliyun-go-sdk-batchcompute/20150630/GetTasksResponse"

	"aliyun-go-sdk-batchcompute/20150630/GetJobDescriptionResponse"

	"aliyun-go-sdk-batchcompute/20150630/GetJobResponse"

	"aliyun-go-sdk-batchcompute/20150630/DeleteJobResponse"
)

type BatchcomputeClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *BatchcomputeClient {
	p := &BatchcomputeClient{}
	p.Version = "2015-06-30"
	p.ProductName = "BatchCompute"
	p.Profile = profile
	p.ApiStyle = "ROA"
	p.HttpRequestBuilder = &core.RoaHttpRequestBuilder{}
	return p
}

func (c *BatchcomputeClient) StopJob(stopJobRequest *StopJobRequest) (batchcompute_StopJobResponse.StopJobResponse, error) {
	var resObj batchcompute_StopJobResponse.StopJobResponse

	if stopJobRequest == nil {
		stopJobRequest = new(StopJobRequest)
	}
	err := c.DoAction(stopJobRequest, &resObj)
	return resObj, err
}

func (c *BatchcomputeClient) StartJob(startJobRequest *StartJobRequest) (batchcompute_StartJobResponse.StartJobResponse, error) {
	var resObj batchcompute_StartJobResponse.StartJobResponse

	if startJobRequest == nil {
		startJobRequest = new(StartJobRequest)
	}
	err := c.DoAction(startJobRequest, &resObj)
	return resObj, err
}

func (c *BatchcomputeClient) ReleaseJob(releaseJobRequest *ReleaseJobRequest) (batchcompute_ReleaseJobResponse.ReleaseJobResponse, error) {
	var resObj batchcompute_ReleaseJobResponse.ReleaseJobResponse

	if releaseJobRequest == nil {
		releaseJobRequest = new(ReleaseJobRequest)
	}
	err := c.DoAction(releaseJobRequest, &resObj)
	return resObj, err
}

func (c *BatchcomputeClient) PutJob(putJobRequest *PutJobRequest) (batchcompute_PutJobResponse.PutJobResponse, error) {
	var resObj batchcompute_PutJobResponse.PutJobResponse

	if putJobRequest == nil {
		putJobRequest = new(PutJobRequest)
	}
	err := c.DoAction(putJobRequest, &resObj)
	return resObj, err
}

func (c *BatchcomputeClient) PostJob(postJobRequest *PostJobRequest) (batchcompute_PostJobResponse.PostJobResponse, error) {
	var resObj batchcompute_PostJobResponse.PostJobResponse

	if postJobRequest == nil {
		postJobRequest = new(PostJobRequest)
	}
	err := c.DoAction(postJobRequest, &resObj)
	return resObj, err
}

func (c *BatchcomputeClient) ListSnapshots(listSnapshotsRequest *ListSnapshotsRequest) (batchcompute_ListSnapshotsResponse.ListSnapshotsResponse, error) {
	var resObj batchcompute_ListSnapshotsResponse.ListSnapshotsResponse

	if listSnapshotsRequest == nil {
		listSnapshotsRequest = new(ListSnapshotsRequest)
	}
	err := c.DoAction(listSnapshotsRequest, &resObj)
	return resObj, err
}

func (c *BatchcomputeClient) ListJobs(listJobsRequest *ListJobsRequest) (batchcompute_ListJobsResponse.ListJobsResponse, error) {
	var resObj batchcompute_ListJobsResponse.ListJobsResponse

	if listJobsRequest == nil {
		listJobsRequest = new(ListJobsRequest)
	}
	err := c.DoAction(listJobsRequest, &resObj)
	return resObj, err
}

func (c *BatchcomputeClient) ListImages(listImagesRequest *ListImagesRequest) (batchcompute_ListImagesResponse.ListImagesResponse, error) {
	var resObj batchcompute_ListImagesResponse.ListImagesResponse

	if listImagesRequest == nil {
		listImagesRequest = new(ListImagesRequest)
	}
	err := c.DoAction(listImagesRequest, &resObj)
	return resObj, err
}

func (c *BatchcomputeClient) GetTasks(getTasksRequest *GetTasksRequest) (batchcompute_GetTasksResponse.GetTasksResponse, error) {
	var resObj batchcompute_GetTasksResponse.GetTasksResponse

	if getTasksRequest == nil {
		getTasksRequest = new(GetTasksRequest)
	}
	err := c.DoAction(getTasksRequest, &resObj)
	return resObj, err
}

func (c *BatchcomputeClient) GetJobDescription(getJobDescriptionRequest *GetJobDescriptionRequest) (batchcompute_GetJobDescriptionResponse.GetJobDescriptionResponse, error) {
	var resObj batchcompute_GetJobDescriptionResponse.GetJobDescriptionResponse

	if getJobDescriptionRequest == nil {
		getJobDescriptionRequest = new(GetJobDescriptionRequest)
	}
	err := c.DoAction(getJobDescriptionRequest, &resObj)
	return resObj, err
}

func (c *BatchcomputeClient) GetJob(getJobRequest *GetJobRequest) (batchcompute_GetJobResponse.GetJobResponse, error) {
	var resObj batchcompute_GetJobResponse.GetJobResponse

	if getJobRequest == nil {
		getJobRequest = new(GetJobRequest)
	}
	err := c.DoAction(getJobRequest, &resObj)
	return resObj, err
}

func (c *BatchcomputeClient) DeleteJob(deleteJobRequest *DeleteJobRequest) (batchcompute_DeleteJobResponse.DeleteJobResponse, error) {
	var resObj batchcompute_DeleteJobResponse.DeleteJobResponse

	if deleteJobRequest == nil {
		deleteJobRequest = new(DeleteJobRequest)
	}
	err := c.DoAction(deleteJobRequest, &resObj)
	return resObj, err
}
