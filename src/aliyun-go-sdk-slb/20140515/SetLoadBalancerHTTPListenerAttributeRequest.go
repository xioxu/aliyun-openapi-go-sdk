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

type SetLoadBalancerHTTPListenerAttributeRequest struct {
	OwnerId                int64
	ResourceOwnerAccount   string
	ResourceOwnerId        int64
	LoadBalancerId         string
	ListenerPort           int32
	Bandwidth              int32
	XForwardedFor          string
	Scheduler              string
	StickySession          string
	StickySessionType      string
	CookieTimeout          int32
	Cookie                 string
	HealthCheck            string
	HealthCheckDomain      string
	HealthCheckURI         string
	HealthyThreshold       int32
	UnhealthyThreshold     int32
	HealthCheckTimeout     int32
	HealthCheckInterval    int32
	HealthCheckConnectPort int32
	HealthCheckHttpCode    string
	MaxConnLimit           int32
	OwnerAccount           string
}

func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 23)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "LoadBalancerId", r.LoadBalancerId)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "ListenerPort", r.ListenerPort)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "Bandwidth", r.Bandwidth)
	core.AddToMapIfNotDefaultValueStr(queryVals, "XForwardedFor", r.XForwardedFor)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Scheduler", r.Scheduler)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StickySession", r.StickySession)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StickySessionType", r.StickySessionType)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "CookieTimeout", r.CookieTimeout)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Cookie", r.Cookie)
	core.AddToMapIfNotDefaultValueStr(queryVals, "HealthCheck", r.HealthCheck)
	core.AddToMapIfNotDefaultValueStr(queryVals, "HealthCheckDomain", r.HealthCheckDomain)
	core.AddToMapIfNotDefaultValueStr(queryVals, "HealthCheckURI", r.HealthCheckURI)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "HealthyThreshold", r.HealthyThreshold)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "UnhealthyThreshold", r.UnhealthyThreshold)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "HealthCheckTimeout", r.HealthCheckTimeout)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "HealthCheckInterval", r.HealthCheckInterval)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "HealthCheckConnectPort", r.HealthCheckConnectPort)
	core.AddToMapIfNotDefaultValueStr(queryVals, "HealthCheckHttpCode", r.HealthCheckHttpCode)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "MaxConnLimit", r.MaxConnLimit)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetPath() string {
	return ""
}

func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetActionName() string {
	return "SetLoadBalancerHTTPListenerAttribute"
}
