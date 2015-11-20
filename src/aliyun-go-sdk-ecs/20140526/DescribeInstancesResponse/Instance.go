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

package ecs_DescribeInstancesResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type Instance struct {
	OperationLocks          []LockReason
	Tags                    []Tag
	InstanceId              string
	InstanceName            string
	Description             string
	ImageId                 string
	RegionId                string
	ZoneId                  string
	ClusterId               string
	InstanceType            string
	HostName                string
	Status                  string
	SerialNumber            string
	InternetChargeType      string
	InternetMaxBandwidthIn  int32
	InternetMaxBandwidthOut int32
	VlanId                  string
	CreationTime            string
	InstanceNetworkType     string
	InstanceChargeType      string
	ExpiredTime             string
	IoOptimized             bool
	DeviceAvailable         bool
	VpcAttributes           VpcAttributes
	EipAddress              EipAddress
	SecurityGroupIds        []string
	PublicIpAddress         []string
	InnerIpAddress          []string
}

func (m *Instance) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.OperationLocks = make([]LockReason, 0)
	for i := 0; i < (flatObj[parentKey+".OperationLocks.length"]).(int); i++ {
		tmp := LockReason{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".OperationLocks[%d]", i), flatObj)
		m.OperationLocks = append(m.OperationLocks, tmp)
	}

	m.Tags = make([]Tag, 0)
	for i := 0; i < (flatObj[parentKey+".Tags.length"]).(int); i++ {
		tmp := Tag{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".Tags[%d]", i), flatObj)
		m.Tags = append(m.Tags, tmp)
	}

	m.InstanceId = builder.GetString(parentKey+".InstanceId", flatObj)

	m.InstanceName = builder.GetString(parentKey+".InstanceName", flatObj)

	m.Description = builder.GetString(parentKey+".Description", flatObj)

	m.ImageId = builder.GetString(parentKey+".ImageId", flatObj)

	m.RegionId = builder.GetString(parentKey+".RegionId", flatObj)

	m.ZoneId = builder.GetString(parentKey+".ZoneId", flatObj)

	m.ClusterId = builder.GetString(parentKey+".ClusterId", flatObj)

	m.InstanceType = builder.GetString(parentKey+".InstanceType", flatObj)

	m.HostName = builder.GetString(parentKey+".HostName", flatObj)

	m.Status = builder.GetString(parentKey+".Status", flatObj)

	m.SerialNumber = builder.GetString(parentKey+".SerialNumber", flatObj)

	m.InternetChargeType = builder.GetString(parentKey+".InternetChargeType", flatObj)

	m.InternetMaxBandwidthIn = builder.GetInt32(parentKey+".InternetMaxBandwidthIn", flatObj)

	m.InternetMaxBandwidthOut = builder.GetInt32(parentKey+".InternetMaxBandwidthOut", flatObj)

	m.VlanId = builder.GetString(parentKey+".VlanId", flatObj)

	m.CreationTime = builder.GetString(parentKey+".CreationTime", flatObj)

	m.InstanceNetworkType = builder.GetString(parentKey+".InstanceNetworkType", flatObj)

	m.InstanceChargeType = builder.GetString(parentKey+".InstanceChargeType", flatObj)

	m.ExpiredTime = builder.GetString(parentKey+".ExpiredTime", flatObj)

	m.IoOptimized = builder.GetBool(parentKey+".IoOptimized", flatObj)

	m.DeviceAvailable = builder.GetBool(parentKey+".DeviceAvailable", flatObj)

	m.VpcAttributes = VpcAttributes{}
	m.VpcAttributes.BuildProperties(parentKey+".VpcAttributes", flatObj)

	m.EipAddress = EipAddress{}
	m.EipAddress.BuildProperties(parentKey+".EipAddress", flatObj)

	m.SecurityGroupIds = make([]string, 0)
	for i := 0; i < (flatObj[parentKey+".SecurityGroupIds.length"]).(int); i++ {
		m.SecurityGroupIds = append(m.SecurityGroupIds, builder.GetString(fmt.Sprintf("%s[%d]", parentKey+".SecurityGroupIds", i), flatObj))
	}

	m.PublicIpAddress = make([]string, 0)
	for i := 0; i < (flatObj[parentKey+".PublicIpAddress.length"]).(int); i++ {
		m.PublicIpAddress = append(m.PublicIpAddress, builder.GetString(fmt.Sprintf("%s[%d]", parentKey+".PublicIpAddress", i), flatObj))
	}

	m.InnerIpAddress = make([]string, 0)
	for i := 0; i < (flatObj[parentKey+".InnerIpAddress.length"]).(int); i++ {
		m.InnerIpAddress = append(m.InnerIpAddress, builder.GetString(fmt.Sprintf("%s[%d]", parentKey+".InnerIpAddress", i), flatObj))
	}

}
