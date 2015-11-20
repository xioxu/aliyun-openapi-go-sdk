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

package ocs_DescribeInstancesResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type OcsInstance struct {
	InstanceId       string
	InstanceName     string
	ConnectionDomain string
	Port             int32
	UserName         string
	InstanceStatus   string
	Capacity         int64
	QPS              int64
	Bandwidth        int64
	Connections      int64
	RegionId         string
	ZoneId           string
	NetworkType      string
	VpcId            string
	VSwitchId        string
	PrivateIpAddress string
	CreationTime     string
}

func (m *OcsInstance) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.InstanceId = builder.GetString(parentKey+".InstanceId", flatObj)

	m.InstanceName = builder.GetString(parentKey+".InstanceName", flatObj)

	m.ConnectionDomain = builder.GetString(parentKey+".ConnectionDomain", flatObj)

	m.Port = builder.GetInt32(parentKey+".Port", flatObj)

	m.UserName = builder.GetString(parentKey+".UserName", flatObj)

	m.InstanceStatus = builder.GetString(parentKey+".InstanceStatus", flatObj)

	m.Capacity = builder.GetInt64(parentKey+".Capacity", flatObj)

	m.QPS = builder.GetInt64(parentKey+".QPS", flatObj)

	m.Bandwidth = builder.GetInt64(parentKey+".Bandwidth", flatObj)

	m.Connections = builder.GetInt64(parentKey+".Connections", flatObj)

	m.RegionId = builder.GetString(parentKey+".RegionId", flatObj)

	m.ZoneId = builder.GetString(parentKey+".ZoneId", flatObj)

	m.NetworkType = builder.GetString(parentKey+".NetworkType", flatObj)

	m.VpcId = builder.GetString(parentKey+".VpcId", flatObj)

	m.VSwitchId = builder.GetString(parentKey+".VSwitchId", flatObj)

	m.PrivateIpAddress = builder.GetString(parentKey+".PrivateIpAddress", flatObj)

	m.CreationTime = builder.GetString(parentKey+".CreationTime", flatObj)

}
