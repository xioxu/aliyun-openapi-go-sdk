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

type SetLoadBalancerUDPListenerAttributeRequest struct {
	OwnerId                   int64
	ResourceOwnerAccount      string
	ResourceOwnerId           int64
	LoadBalancerId            string
	ListenerPort              int32
	Bandwidth                 int32
	Scheduler                 string
	PersistenceTimeout        int32
	HealthyThreshold          int32
	UnhealthyThreshold        int32
	HealthCheckConnectTimeout int32
	HealthCheckConnectPort    int32
	HealthCheckInterval       int32
	MaxConnLimit              int32
	OwnerAccount              string
}

func (r *SetLoadBalancerUDPListenerAttributeRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 15)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "LoadBalancerId", r.LoadBalancerId)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "ListenerPort", r.ListenerPort)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "Bandwidth", r.Bandwidth)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Scheduler", r.Scheduler)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PersistenceTimeout", r.PersistenceTimeout)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "HealthyThreshold", r.HealthyThreshold)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "UnhealthyThreshold", r.UnhealthyThreshold)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "HealthCheckConnectTimeout", r.HealthCheckConnectTimeout)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "HealthCheckConnectPort", r.HealthCheckConnectPort)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "HealthCheckInterval", r.HealthCheckInterval)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "MaxConnLimit", r.MaxConnLimit)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *SetLoadBalancerUDPListenerAttributeRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *SetLoadBalancerUDPListenerAttributeRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *SetLoadBalancerUDPListenerAttributeRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *SetLoadBalancerUDPListenerAttributeRequest) GetPath() string {
	return ""
}

func (r *SetLoadBalancerUDPListenerAttributeRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *SetLoadBalancerUDPListenerAttributeRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *SetLoadBalancerUDPListenerAttributeRequest) GetActionName() string {
	return "SetLoadBalancerUDPListenerAttribute"
}
