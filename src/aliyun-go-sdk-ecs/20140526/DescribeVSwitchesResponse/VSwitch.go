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

package ecs_DescribeVSwitchesResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type VSwitch struct {
	VSwitchId               string
	VpcId                   string
	Status                  string
	CidrBlock               string
	ZoneId                  string
	AvailableIpAddressCount int64
	Description             string
	VSwitchName             string
	CreationTime            string
}

func (m *VSwitch) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.VSwitchId = builder.GetString(parentKey+".VSwitchId", flatObj)

	m.VpcId = builder.GetString(parentKey+".VpcId", flatObj)

	m.Status = builder.GetString(parentKey+".Status", flatObj)

	m.CidrBlock = builder.GetString(parentKey+".CidrBlock", flatObj)

	m.ZoneId = builder.GetString(parentKey+".ZoneId", flatObj)

	m.AvailableIpAddressCount = builder.GetInt64(parentKey+".AvailableIpAddressCount", flatObj)

	m.Description = builder.GetString(parentKey+".Description", flatObj)

	m.VSwitchName = builder.GetString(parentKey+".VSwitchName", flatObj)

	m.CreationTime = builder.GetString(parentKey+".CreationTime", flatObj)

}
