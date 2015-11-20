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

package ecs_DescribeEipAddressesResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type EipAddress struct {
	OperationLocks     []LockReason
	RegionId           string
	IpAddress          string
	AllocationId       string
	Status             string
	InstanceId         string
	Bandwidth          string
	InternetChargeType string
	AllocationTime     string
	InstanceType       string
}

func (m *EipAddress) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.OperationLocks = make([]LockReason, 0)
	for i := 0; i < (flatObj[parentKey+".OperationLocks.length"]).(int); i++ {
		tmp := LockReason{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".OperationLocks[%d]", i), flatObj)
		m.OperationLocks = append(m.OperationLocks, tmp)
	}

	m.RegionId = builder.GetString(parentKey+".RegionId", flatObj)

	m.IpAddress = builder.GetString(parentKey+".IpAddress", flatObj)

	m.AllocationId = builder.GetString(parentKey+".AllocationId", flatObj)

	m.Status = builder.GetString(parentKey+".Status", flatObj)

	m.InstanceId = builder.GetString(parentKey+".InstanceId", flatObj)

	m.Bandwidth = builder.GetString(parentKey+".Bandwidth", flatObj)

	m.InternetChargeType = builder.GetString(parentKey+".InternetChargeType", flatObj)

	m.AllocationTime = builder.GetString(parentKey+".AllocationTime", flatObj)

	m.InstanceType = builder.GetString(parentKey+".InstanceType", flatObj)

}
