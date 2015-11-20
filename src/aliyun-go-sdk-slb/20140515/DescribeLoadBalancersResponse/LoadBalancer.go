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

package slb_DescribeLoadBalancersResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type LoadBalancer struct {
	LoadBalancerId     string
	LoadBalancerName   string
	LoadBalancerStatus string
	Address            string
	AddressType        string
	RegionId           string
	RegionIdAlias      string
	VSwitchId          string
	VpcId              string
	NetworkType        string
	MasterZoneId       string
	SlaveZoneId        string
	MaxConnLimit       int32
	SecurityStatus     string
	InternetChargeType string
	SysBandwidth       int32
}

func (m *LoadBalancer) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.LoadBalancerId = builder.GetString(parentKey+".LoadBalancerId", flatObj)

	m.LoadBalancerName = builder.GetString(parentKey+".LoadBalancerName", flatObj)

	m.LoadBalancerStatus = builder.GetString(parentKey+".LoadBalancerStatus", flatObj)

	m.Address = builder.GetString(parentKey+".Address", flatObj)

	m.AddressType = builder.GetString(parentKey+".AddressType", flatObj)

	m.RegionId = builder.GetString(parentKey+".RegionId", flatObj)

	m.RegionIdAlias = builder.GetString(parentKey+".RegionIdAlias", flatObj)

	m.VSwitchId = builder.GetString(parentKey+".VSwitchId", flatObj)

	m.VpcId = builder.GetString(parentKey+".VpcId", flatObj)

	m.NetworkType = builder.GetString(parentKey+".NetworkType", flatObj)

	m.MasterZoneId = builder.GetString(parentKey+".MasterZoneId", flatObj)

	m.SlaveZoneId = builder.GetString(parentKey+".SlaveZoneId", flatObj)

	m.MaxConnLimit = builder.GetInt32(parentKey+".MaxConnLimit", flatObj)

	m.SecurityStatus = builder.GetString(parentKey+".SecurityStatus", flatObj)

	m.InternetChargeType = builder.GetString(parentKey+".InternetChargeType", flatObj)

	m.SysBandwidth = builder.GetInt32(parentKey+".SysBandwidth", flatObj)

}
