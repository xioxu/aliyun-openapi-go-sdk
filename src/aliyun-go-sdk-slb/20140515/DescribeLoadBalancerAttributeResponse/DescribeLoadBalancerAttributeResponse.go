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
	ListenerPortsAndProtocal []ListenerPortAndProtocal
	ListenerPortsAndProtocol []ListenerPortAndProtocol
	BackendServers           []BackendServer
	LoadBalancerId           string
	LoadBalancerName         string
	LoadBalancerStatus       string
	RegionId                 string
	RegionIdAlias            string
	Address                  string
	AddressType              string
	VpcId                    string
	VSwitchId                string
	NetworkType              string
	InternetChargeType       string
	Bandwidth                int32
	SysBandwidth             int32
	CreateTime               string
	CreateTimeStamp          int64
	MasterZoneId             string
	SlaveZoneId              string
	MaxConnLimit             int32
	SecurityStatus           string
	RequestId                string
	ListenerPorts            []string
}

func (m *DescribeLoadBalancerAttributeResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.ListenerPortsAndProtocal = make([]ListenerPortAndProtocal, 0)
	for i := 0; i < (flatObj[parentKey+".ListenerPortsAndProtocal.length"]).(int); i++ {
		tmp := ListenerPortAndProtocal{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".ListenerPortsAndProtocal[%d]", i), flatObj)
		m.ListenerPortsAndProtocal = append(m.ListenerPortsAndProtocal, tmp)
	}

	m.ListenerPortsAndProtocol = make([]ListenerPortAndProtocol, 0)
	for i := 0; i < (flatObj[parentKey+".ListenerPortsAndProtocol.length"]).(int); i++ {
		tmp := ListenerPortAndProtocol{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".ListenerPortsAndProtocol[%d]", i), flatObj)
		m.ListenerPortsAndProtocol = append(m.ListenerPortsAndProtocol, tmp)
	}

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

	m.RegionIdAlias = builder.GetString(parentKey+".RegionIdAlias", flatObj)

	m.Address = builder.GetString(parentKey+".Address", flatObj)

	m.AddressType = builder.GetString(parentKey+".AddressType", flatObj)

	m.VpcId = builder.GetString(parentKey+".VpcId", flatObj)

	m.VSwitchId = builder.GetString(parentKey+".VSwitchId", flatObj)

	m.NetworkType = builder.GetString(parentKey+".NetworkType", flatObj)

	m.InternetChargeType = builder.GetString(parentKey+".InternetChargeType", flatObj)

	m.Bandwidth = builder.GetInt32(parentKey+".Bandwidth", flatObj)

	m.SysBandwidth = builder.GetInt32(parentKey+".SysBandwidth", flatObj)

	m.CreateTime = builder.GetString(parentKey+".CreateTime", flatObj)

	m.CreateTimeStamp = builder.GetInt64(parentKey+".CreateTimeStamp", flatObj)

	m.MasterZoneId = builder.GetString(parentKey+".MasterZoneId", flatObj)

	m.SlaveZoneId = builder.GetString(parentKey+".SlaveZoneId", flatObj)

	m.MaxConnLimit = builder.GetInt32(parentKey+".MaxConnLimit", flatObj)

	m.SecurityStatus = builder.GetString(parentKey+".SecurityStatus", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

	m.ListenerPorts = make([]string, 0)
	for i := 0; i < (flatObj[parentKey+".ListenerPorts.length"]).(int); i++ {
		m.ListenerPorts = append(m.ListenerPorts, builder.GetString(fmt.Sprintf("%s[%d]", parentKey+".ListenerPorts", i), flatObj))
	}

}
