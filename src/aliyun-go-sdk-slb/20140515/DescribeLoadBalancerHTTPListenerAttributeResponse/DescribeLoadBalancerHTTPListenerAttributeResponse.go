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

package slb_DescribeLoadBalancerHTTPListenerAttributeResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type DescribeLoadBalancerHTTPListenerAttributeResponse struct {
	ListenerPort           int32
	BackendServerPort      int32
	Bandwidth              int32
	Status                 string
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
	RequestId              string
}

func (m *DescribeLoadBalancerHTTPListenerAttributeResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.ListenerPort = builder.GetInt32(parentKey+".ListenerPort", flatObj)

	m.BackendServerPort = builder.GetInt32(parentKey+".BackendServerPort", flatObj)

	m.Bandwidth = builder.GetInt32(parentKey+".Bandwidth", flatObj)

	m.Status = builder.GetString(parentKey+".Status", flatObj)

	m.XForwardedFor = builder.GetString(parentKey+".XForwardedFor", flatObj)

	m.Scheduler = builder.GetString(parentKey+".Scheduler", flatObj)

	m.StickySession = builder.GetString(parentKey+".StickySession", flatObj)

	m.StickySessionType = builder.GetString(parentKey+".StickySessionType", flatObj)

	m.CookieTimeout = builder.GetInt32(parentKey+".CookieTimeout", flatObj)

	m.Cookie = builder.GetString(parentKey+".Cookie", flatObj)

	m.HealthCheck = builder.GetString(parentKey+".HealthCheck", flatObj)

	m.HealthCheckDomain = builder.GetString(parentKey+".HealthCheckDomain", flatObj)

	m.HealthCheckURI = builder.GetString(parentKey+".HealthCheckURI", flatObj)

	m.HealthyThreshold = builder.GetInt32(parentKey+".HealthyThreshold", flatObj)

	m.UnhealthyThreshold = builder.GetInt32(parentKey+".UnhealthyThreshold", flatObj)

	m.HealthCheckTimeout = builder.GetInt32(parentKey+".HealthCheckTimeout", flatObj)

	m.HealthCheckInterval = builder.GetInt32(parentKey+".HealthCheckInterval", flatObj)

	m.HealthCheckConnectPort = builder.GetInt32(parentKey+".HealthCheckConnectPort", flatObj)

	m.HealthCheckHttpCode = builder.GetString(parentKey+".HealthCheckHttpCode", flatObj)

	m.MaxConnLimit = builder.GetInt32(parentKey+".MaxConnLimit", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

}
