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

import "aliyun-go-sdk-core"

type CreateClusterRequest struct {
	Name              string
	ZoneId            string
	LogEnable         bool
	LogPath           string
	SecurityGroup     string
	IsOpenPublicIp    bool
	SecurityGroupName string
	EmrVer            string
	ClusterType       string
	Install           string
	MasterIndex       int32
	EcsOrder          string
	PayType           int32
	Period            int32
	EmrRole4ECS       string
	EmrRole4Oss       string
}

func (r *CreateClusterRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 16)

	core.AddToMapIfNotDefaultValueStr(queryVals, "Name", r.Name)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ZoneId", r.ZoneId)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "LogEnable", r.LogEnable)
	core.AddToMapIfNotDefaultValueStr(queryVals, "LogPath", r.LogPath)
	core.AddToMapIfNotDefaultValueStr(queryVals, "SecurityGroup", r.SecurityGroup)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "IsOpenPublicIp", r.IsOpenPublicIp)
	core.AddToMapIfNotDefaultValueStr(queryVals, "SecurityGroupName", r.SecurityGroupName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "EmrVer", r.EmrVer)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ClusterType", r.ClusterType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Install", r.Install)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "MasterIndex", r.MasterIndex)
	core.AddToMapIfNotDefaultValueStr(queryVals, "EcsOrder", r.EcsOrder)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PayType", r.PayType)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "Period", r.Period)
	core.AddToMapIfNotDefaultValueStr(queryVals, "EmrRole4ECS", r.EmrRole4ECS)
	core.AddToMapIfNotDefaultValueStr(queryVals, "EmrRole4Oss", r.EmrRole4Oss)

	return queryVals
}

func (r *CreateClusterRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *CreateClusterRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *CreateClusterRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *CreateClusterRequest) GetPath() string {
	return ""
}

func (r *CreateClusterRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *CreateClusterRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *CreateClusterRequest) GetActionName() string {
	return "CreateCluster"
}
