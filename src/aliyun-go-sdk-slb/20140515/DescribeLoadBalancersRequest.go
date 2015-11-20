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

package slb

import "aliyun-go-sdk-core"

type DescribeLoadBalancersRequest struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	ServerId             string
	LoadBalancerId       string
	AddressType          string
	InternetChargeType   string
	VpcId                string
	VSwitchId            string
	NetworkType          string
	Address              string
	SecurityStatus       string
	MasterZoneId         string
	SlaveZoneId          string
	OwnerAccount         string
}

func (r *DescribeLoadBalancersRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 15)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ServerId", r.ServerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "LoadBalancerId", r.LoadBalancerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "AddressType", r.AddressType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "InternetChargeType", r.InternetChargeType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "VpcId", r.VpcId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "VSwitchId", r.VSwitchId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "NetworkType", r.NetworkType)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Address", r.Address)
	core.AddToMapIfNotDefaultValueStr(queryVals, "SecurityStatus", r.SecurityStatus)
	core.AddToMapIfNotDefaultValueStr(queryVals, "MasterZoneId", r.MasterZoneId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "SlaveZoneId", r.SlaveZoneId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *DescribeLoadBalancersRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *DescribeLoadBalancersRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *DescribeLoadBalancersRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *DescribeLoadBalancersRequest) GetPath() string {
	return ""
}

func (r *DescribeLoadBalancersRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *DescribeLoadBalancersRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *DescribeLoadBalancersRequest) GetActionName() string {
	return "DescribeLoadBalancers"
}
