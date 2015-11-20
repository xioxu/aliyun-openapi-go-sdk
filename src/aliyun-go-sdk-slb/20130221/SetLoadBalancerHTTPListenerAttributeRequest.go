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
	LoadBalancerId     string
	ListenerPort       int32
	XForwardedFor      string
	Scheduler          string
	StickySession      string
	StickySessionType  string
	CookieTimeout      int32
	Cookie             string
	HealthCheck        string
	Domain             string
	URI                string
	HealthyThreshold   int32
	UnhealthyThreshold int32
	HealthCheckTimeout int32
	Interval           int32
	HostId             string
	OwnerAccount       string
}

func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 17)

	core.AddToMapIfNotDefaultValueStr(queryVals, "LoadBalancerId", r.LoadBalancerId)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "ListenerPort", r.ListenerPort)
	core.AddToMapIfNotDefaultValueStr(queryVals, "XForwardedFor", r.XForwardedFor)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Scheduler", r.Scheduler)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StickySession", r.StickySession)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StickySessionType", r.StickySessionType)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "CookieTimeout", r.CookieTimeout)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Cookie", r.Cookie)
	core.AddToMapIfNotDefaultValueStr(queryVals, "HealthCheck", r.HealthCheck)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Domain", r.Domain)
	core.AddToMapIfNotDefaultValueStr(queryVals, "URI", r.URI)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "HealthyThreshold", r.HealthyThreshold)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "UnhealthyThreshold", r.UnhealthyThreshold)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "HealthCheckTimeout", r.HealthCheckTimeout)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "Interval", r.Interval)
	core.AddToMapIfNotDefaultValueStr(queryVals, "HostId", r.HostId)
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
