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

package ess

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-ess/20140828/RemoveInstancesResponse"

	"aliyun-go-sdk-ess/20140828/ModifyScheduledTaskResponse"

	"aliyun-go-sdk-ess/20140828/ModifyScalingRuleResponse"

	"aliyun-go-sdk-ess/20140828/ModifyScalingGroupResponse"

	"aliyun-go-sdk-ess/20140828/ExecuteScalingRuleResponse"

	"aliyun-go-sdk-ess/20140828/EnableScalingGroupResponse"

	"aliyun-go-sdk-ess/20140828/DisableScalingGroupResponse"

	"aliyun-go-sdk-ess/20140828/DetachInstancesResponse"

	"aliyun-go-sdk-ess/20140828/DescribeScheduledTasksResponse"

	"aliyun-go-sdk-ess/20140828/DescribeScalingRulesResponse"

	"aliyun-go-sdk-ess/20140828/DescribeScalingInstancesResponse"

	"aliyun-go-sdk-ess/20140828/DescribeScalingGroupsResponse"

	"aliyun-go-sdk-ess/20140828/DescribeScalingConfigurationsResponse"

	"aliyun-go-sdk-ess/20140828/DescribeScalingActivitiesResponse"

	"aliyun-go-sdk-ess/20140828/DeleteScheduledTaskResponse"

	"aliyun-go-sdk-ess/20140828/DeleteScalingRuleResponse"

	"aliyun-go-sdk-ess/20140828/DeleteScalingGroupResponse"

	"aliyun-go-sdk-ess/20140828/DeleteScalingConfigurationResponse"

	"aliyun-go-sdk-ess/20140828/CreateScheduledTaskResponse"

	"aliyun-go-sdk-ess/20140828/CreateScalingRuleResponse"

	"aliyun-go-sdk-ess/20140828/CreateScalingGroupResponse"

	"aliyun-go-sdk-ess/20140828/CreateScalingConfigurationResponse"

	"aliyun-go-sdk-ess/20140828/AttachInstancesResponse"
)

type EssClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *EssClient {
	p := &EssClient{}
	p.Version = "2014-08-28"
	p.ProductName = "Ess"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *EssClient) RemoveInstances(removeInstancesRequest *RemoveInstancesRequest) (ess_RemoveInstancesResponse.RemoveInstancesResponse, error) {
	var resObj ess_RemoveInstancesResponse.RemoveInstancesResponse

	if removeInstancesRequest == nil {
		removeInstancesRequest = new(RemoveInstancesRequest)
	}
	err := c.DoAction(removeInstancesRequest, &resObj)
	return resObj, err
}

func (c *EssClient) ModifyScheduledTask(modifyScheduledTaskRequest *ModifyScheduledTaskRequest) (ess_ModifyScheduledTaskResponse.ModifyScheduledTaskResponse, error) {
	var resObj ess_ModifyScheduledTaskResponse.ModifyScheduledTaskResponse

	if modifyScheduledTaskRequest == nil {
		modifyScheduledTaskRequest = new(ModifyScheduledTaskRequest)
	}
	err := c.DoAction(modifyScheduledTaskRequest, &resObj)
	return resObj, err
}

func (c *EssClient) ModifyScalingRule(modifyScalingRuleRequest *ModifyScalingRuleRequest) (ess_ModifyScalingRuleResponse.ModifyScalingRuleResponse, error) {
	var resObj ess_ModifyScalingRuleResponse.ModifyScalingRuleResponse

	if modifyScalingRuleRequest == nil {
		modifyScalingRuleRequest = new(ModifyScalingRuleRequest)
	}
	err := c.DoAction(modifyScalingRuleRequest, &resObj)
	return resObj, err
}

func (c *EssClient) ModifyScalingGroup(modifyScalingGroupRequest *ModifyScalingGroupRequest) (ess_ModifyScalingGroupResponse.ModifyScalingGroupResponse, error) {
	var resObj ess_ModifyScalingGroupResponse.ModifyScalingGroupResponse

	if modifyScalingGroupRequest == nil {
		modifyScalingGroupRequest = new(ModifyScalingGroupRequest)
	}
	err := c.DoAction(modifyScalingGroupRequest, &resObj)
	return resObj, err
}

func (c *EssClient) ExecuteScalingRule(executeScalingRuleRequest *ExecuteScalingRuleRequest) (ess_ExecuteScalingRuleResponse.ExecuteScalingRuleResponse, error) {
	var resObj ess_ExecuteScalingRuleResponse.ExecuteScalingRuleResponse

	if executeScalingRuleRequest == nil {
		executeScalingRuleRequest = new(ExecuteScalingRuleRequest)
	}
	err := c.DoAction(executeScalingRuleRequest, &resObj)
	return resObj, err
}

func (c *EssClient) EnableScalingGroup(enableScalingGroupRequest *EnableScalingGroupRequest) (ess_EnableScalingGroupResponse.EnableScalingGroupResponse, error) {
	var resObj ess_EnableScalingGroupResponse.EnableScalingGroupResponse

	if enableScalingGroupRequest == nil {
		enableScalingGroupRequest = new(EnableScalingGroupRequest)
	}
	err := c.DoAction(enableScalingGroupRequest, &resObj)
	return resObj, err
}

func (c *EssClient) DisableScalingGroup(disableScalingGroupRequest *DisableScalingGroupRequest) (ess_DisableScalingGroupResponse.DisableScalingGroupResponse, error) {
	var resObj ess_DisableScalingGroupResponse.DisableScalingGroupResponse

	if disableScalingGroupRequest == nil {
		disableScalingGroupRequest = new(DisableScalingGroupRequest)
	}
	err := c.DoAction(disableScalingGroupRequest, &resObj)
	return resObj, err
}

func (c *EssClient) DetachInstances(detachInstancesRequest *DetachInstancesRequest) (ess_DetachInstancesResponse.DetachInstancesResponse, error) {
	var resObj ess_DetachInstancesResponse.DetachInstancesResponse

	if detachInstancesRequest == nil {
		detachInstancesRequest = new(DetachInstancesRequest)
	}
	err := c.DoAction(detachInstancesRequest, &resObj)
	return resObj, err
}

func (c *EssClient) DescribeScheduledTasks(describeScheduledTasksRequest *DescribeScheduledTasksRequest) (ess_DescribeScheduledTasksResponse.DescribeScheduledTasksResponse, error) {
	var resObj ess_DescribeScheduledTasksResponse.DescribeScheduledTasksResponse

	if describeScheduledTasksRequest == nil {
		describeScheduledTasksRequest = new(DescribeScheduledTasksRequest)
	}
	err := c.DoAction(describeScheduledTasksRequest, &resObj)
	return resObj, err
}

func (c *EssClient) DescribeScalingRules(describeScalingRulesRequest *DescribeScalingRulesRequest) (ess_DescribeScalingRulesResponse.DescribeScalingRulesResponse, error) {
	var resObj ess_DescribeScalingRulesResponse.DescribeScalingRulesResponse

	if describeScalingRulesRequest == nil {
		describeScalingRulesRequest = new(DescribeScalingRulesRequest)
	}
	err := c.DoAction(describeScalingRulesRequest, &resObj)
	return resObj, err
}

func (c *EssClient) DescribeScalingInstances(describeScalingInstancesRequest *DescribeScalingInstancesRequest) (ess_DescribeScalingInstancesResponse.DescribeScalingInstancesResponse, error) {
	var resObj ess_DescribeScalingInstancesResponse.DescribeScalingInstancesResponse

	if describeScalingInstancesRequest == nil {
		describeScalingInstancesRequest = new(DescribeScalingInstancesRequest)
	}
	err := c.DoAction(describeScalingInstancesRequest, &resObj)
	return resObj, err
}

func (c *EssClient) DescribeScalingGroups(describeScalingGroupsRequest *DescribeScalingGroupsRequest) (ess_DescribeScalingGroupsResponse.DescribeScalingGroupsResponse, error) {
	var resObj ess_DescribeScalingGroupsResponse.DescribeScalingGroupsResponse

	if describeScalingGroupsRequest == nil {
		describeScalingGroupsRequest = new(DescribeScalingGroupsRequest)
	}
	err := c.DoAction(describeScalingGroupsRequest, &resObj)
	return resObj, err
}

func (c *EssClient) DescribeScalingConfigurations(describeScalingConfigurationsRequest *DescribeScalingConfigurationsRequest) (ess_DescribeScalingConfigurationsResponse.DescribeScalingConfigurationsResponse, error) {
	var resObj ess_DescribeScalingConfigurationsResponse.DescribeScalingConfigurationsResponse

	if describeScalingConfigurationsRequest == nil {
		describeScalingConfigurationsRequest = new(DescribeScalingConfigurationsRequest)
	}
	err := c.DoAction(describeScalingConfigurationsRequest, &resObj)
	return resObj, err
}

func (c *EssClient) DescribeScalingActivities(describeScalingActivitiesRequest *DescribeScalingActivitiesRequest) (ess_DescribeScalingActivitiesResponse.DescribeScalingActivitiesResponse, error) {
	var resObj ess_DescribeScalingActivitiesResponse.DescribeScalingActivitiesResponse

	if describeScalingActivitiesRequest == nil {
		describeScalingActivitiesRequest = new(DescribeScalingActivitiesRequest)
	}
	err := c.DoAction(describeScalingActivitiesRequest, &resObj)
	return resObj, err
}

func (c *EssClient) DeleteScheduledTask(deleteScheduledTaskRequest *DeleteScheduledTaskRequest) (ess_DeleteScheduledTaskResponse.DeleteScheduledTaskResponse, error) {
	var resObj ess_DeleteScheduledTaskResponse.DeleteScheduledTaskResponse

	if deleteScheduledTaskRequest == nil {
		deleteScheduledTaskRequest = new(DeleteScheduledTaskRequest)
	}
	err := c.DoAction(deleteScheduledTaskRequest, &resObj)
	return resObj, err
}

func (c *EssClient) DeleteScalingRule(deleteScalingRuleRequest *DeleteScalingRuleRequest) (ess_DeleteScalingRuleResponse.DeleteScalingRuleResponse, error) {
	var resObj ess_DeleteScalingRuleResponse.DeleteScalingRuleResponse

	if deleteScalingRuleRequest == nil {
		deleteScalingRuleRequest = new(DeleteScalingRuleRequest)
	}
	err := c.DoAction(deleteScalingRuleRequest, &resObj)
	return resObj, err
}

func (c *EssClient) DeleteScalingGroup(deleteScalingGroupRequest *DeleteScalingGroupRequest) (ess_DeleteScalingGroupResponse.DeleteScalingGroupResponse, error) {
	var resObj ess_DeleteScalingGroupResponse.DeleteScalingGroupResponse

	if deleteScalingGroupRequest == nil {
		deleteScalingGroupRequest = new(DeleteScalingGroupRequest)
	}
	err := c.DoAction(deleteScalingGroupRequest, &resObj)
	return resObj, err
}

func (c *EssClient) DeleteScalingConfiguration(deleteScalingConfigurationRequest *DeleteScalingConfigurationRequest) (ess_DeleteScalingConfigurationResponse.DeleteScalingConfigurationResponse, error) {
	var resObj ess_DeleteScalingConfigurationResponse.DeleteScalingConfigurationResponse

	if deleteScalingConfigurationRequest == nil {
		deleteScalingConfigurationRequest = new(DeleteScalingConfigurationRequest)
	}
	err := c.DoAction(deleteScalingConfigurationRequest, &resObj)
	return resObj, err
}

func (c *EssClient) CreateScheduledTask(createScheduledTaskRequest *CreateScheduledTaskRequest) (ess_CreateScheduledTaskResponse.CreateScheduledTaskResponse, error) {
	var resObj ess_CreateScheduledTaskResponse.CreateScheduledTaskResponse

	if createScheduledTaskRequest == nil {
		createScheduledTaskRequest = new(CreateScheduledTaskRequest)
	}
	err := c.DoAction(createScheduledTaskRequest, &resObj)
	return resObj, err
}

func (c *EssClient) CreateScalingRule(createScalingRuleRequest *CreateScalingRuleRequest) (ess_CreateScalingRuleResponse.CreateScalingRuleResponse, error) {
	var resObj ess_CreateScalingRuleResponse.CreateScalingRuleResponse

	if createScalingRuleRequest == nil {
		createScalingRuleRequest = new(CreateScalingRuleRequest)
	}
	err := c.DoAction(createScalingRuleRequest, &resObj)
	return resObj, err
}

func (c *EssClient) CreateScalingGroup(createScalingGroupRequest *CreateScalingGroupRequest) (ess_CreateScalingGroupResponse.CreateScalingGroupResponse, error) {
	var resObj ess_CreateScalingGroupResponse.CreateScalingGroupResponse

	if createScalingGroupRequest == nil {
		createScalingGroupRequest = new(CreateScalingGroupRequest)
	}
	err := c.DoAction(createScalingGroupRequest, &resObj)
	return resObj, err
}

func (c *EssClient) CreateScalingConfiguration(createScalingConfigurationRequest *CreateScalingConfigurationRequest) (ess_CreateScalingConfigurationResponse.CreateScalingConfigurationResponse, error) {
	var resObj ess_CreateScalingConfigurationResponse.CreateScalingConfigurationResponse

	if createScalingConfigurationRequest == nil {
		createScalingConfigurationRequest = new(CreateScalingConfigurationRequest)
	}
	err := c.DoAction(createScalingConfigurationRequest, &resObj)
	return resObj, err
}

func (c *EssClient) AttachInstances(attachInstancesRequest *AttachInstancesRequest) (ess_AttachInstancesResponse.AttachInstancesResponse, error) {
	var resObj ess_AttachInstancesResponse.AttachInstancesResponse

	if attachInstancesRequest == nil {
		attachInstancesRequest = new(AttachInstancesRequest)
	}
	err := c.DoAction(attachInstancesRequest, &resObj)
	return resObj, err
}
