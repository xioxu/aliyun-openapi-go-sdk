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

package ecs

import "aliyun-go-sdk-core"

type RevokeSecurityGroupEgressRequest struct {
	OwnerId               int64
	ResourceOwnerAccount  string
	ResourceOwnerId       int64
	SecurityGroupId       string
	IpProtocol            string
	PortRange             string
	DestGroupId           string
	DestGroupOwnerAccount string
	DestCidrIp            string
	Policy                string
	NicType               string
	OwnerAccount          string
}

func (r *RevokeSecurityGroupEgressRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 12)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "SecurityGroupId", r.SecurityGroupId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "IpProtocol", r.IpProtocol)
	core.AddToMapIfNotDefaultValueStr(queryVals, "PortRange", r.PortRange)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DestGroupId", r.DestGroupId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DestGroupOwnerAccount", r.DestGroupOwnerAccount)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DestCidrIp", r.DestCidrIp)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Policy", r.Policy)
	core.AddToMapIfNotDefaultValueStr(queryVals, "NicType", r.NicType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *RevokeSecurityGroupEgressRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *RevokeSecurityGroupEgressRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *RevokeSecurityGroupEgressRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *RevokeSecurityGroupEgressRequest) GetPath() string {
	return ""
}

func (r *RevokeSecurityGroupEgressRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *RevokeSecurityGroupEgressRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *RevokeSecurityGroupEgressRequest) GetActionName() string {
	return "RevokeSecurityGroupEgress"
}
