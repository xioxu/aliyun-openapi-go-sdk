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

type VpcAttributes struct {
	VpcId            string
	VSwitchId        string
	NatIpAddress     string
	PrivateIpAddress []string
}

func (m *VpcAttributes) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.VpcId = builder.GetString(parentKey+".VpcId", flatObj)

	m.VSwitchId = builder.GetString(parentKey+".VSwitchId", flatObj)

	m.NatIpAddress = builder.GetString(parentKey+".NatIpAddress", flatObj)

	m.PrivateIpAddress = make([]string, 0)
	for i := 0; i < (flatObj[parentKey+".PrivateIpAddress.length"]).(int); i++ {
		m.PrivateIpAddress = append(m.PrivateIpAddress, builder.GetString(fmt.Sprintf("%s[%d]", parentKey+".PrivateIpAddress", i), flatObj))
	}

}
