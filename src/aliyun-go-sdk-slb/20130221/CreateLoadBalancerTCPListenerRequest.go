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

type CreateLoadBalancerTCPListenerRequest struct {
	LoadBalancerId     string
	ListenerPort       int32
	BackendServerPort  int32
	ListenerStatus     string
	Scheduler          string
	PersistenceTimeout int32
	HealthCheck        string
	ConnectTimeout     int32
	ConnectPort        int32
	Interval           int32
	HostId             string
	OwnerAccount       string
}

func (r *CreateLoadBalancerTCPListenerRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 12)

	core.AddToMapIfNotDefaultValueStr(queryVals, "LoadBalancerId", r.LoadBalancerId)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "ListenerPort", r.ListenerPort)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "BackendServerPort", r.BackendServerPort)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ListenerStatus", r.ListenerStatus)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Scheduler", r.Scheduler)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PersistenceTimeout", r.PersistenceTimeout)
	core.AddToMapIfNotDefaultValueStr(queryVals, "HealthCheck", r.HealthCheck)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "ConnectTimeout", r.ConnectTimeout)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "ConnectPort", r.ConnectPort)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "Interval", r.Interval)
	core.AddToMapIfNotDefaultValueStr(queryVals, "HostId", r.HostId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *CreateLoadBalancerTCPListenerRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *CreateLoadBalancerTCPListenerRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *CreateLoadBalancerTCPListenerRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *CreateLoadBalancerTCPListenerRequest) GetPath() string {
	return ""
}

func (r *CreateLoadBalancerTCPListenerRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *CreateLoadBalancerTCPListenerRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *CreateLoadBalancerTCPListenerRequest) GetActionName() string {
	return "CreateLoadBalancerTCPListener"
}
