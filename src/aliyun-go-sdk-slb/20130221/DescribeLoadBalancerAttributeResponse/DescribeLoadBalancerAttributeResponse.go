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

package slb_DescribeLoadBalancerAttributeResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type DescribeLoadBalancerAttributeResponse struct {
	BackendServers     []BackendServer
	LoadBalancerId     string
	LoadBalancerName   string
	LoadBalancerStatus string
	RegionId           string
	Address            string
	IsPublicAddress    string
	RequestId          string
	ListenerPorts      []string
}

func (m *DescribeLoadBalancerAttributeResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.BackendServers = make([]BackendServer, 0)
	for i := 0; i < (flatObj[parentKey+".BackendServers.length"]).(int); i++ {
		tmp := BackendServer{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".BackendServers[%d]", i), flatObj)
		m.BackendServers = append(m.BackendServers, tmp)
	}

	m.LoadBalancerId = builder.GetString(parentKey+".LoadBalancerId", flatObj)

	m.LoadBalancerName = builder.GetString(parentKey+".LoadBalancerName", flatObj)

	m.LoadBalancerStatus = builder.GetString(parentKey+".LoadBalancerStatus", flatObj)

	m.RegionId = builder.GetString(parentKey+".RegionId", flatObj)

	m.Address = builder.GetString(parentKey+".Address", flatObj)

	m.IsPublicAddress = builder.GetString(parentKey+".IsPublicAddress", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

	m.ListenerPorts = make([]string, 0)
	for i := 0; i < (flatObj[parentKey+".ListenerPorts.length"]).(int); i++ {
		m.ListenerPorts = append(m.ListenerPorts, builder.GetString(fmt.Sprintf("%s[%d]", parentKey+".ListenerPorts", i), flatObj))
	}

}
