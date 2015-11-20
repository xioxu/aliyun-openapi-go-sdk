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

package ess_DescribeScalingGroupsResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type ScalingGroup struct {
	DefaultCooldown              int32
	MaxSize                      int32
	PendingCapacity              int32
	LoadBalancerId               string
	RemovingCapacity             int32
	ScalingGroupName             string
	ActiveCapacity               int32
	ActiveScalingConfigurationId string
	ScalingGroupId               string
	RegionId                     string
	TotalCapacity                int32
	MinSize                      int32
	LifecycleState               string
	CreationTime                 string
	RemovalPolicies              []string
	DBInstanceIds                []string
}

func (m *ScalingGroup) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.DefaultCooldown = builder.GetInt32(parentKey+".DefaultCooldown", flatObj)

	m.MaxSize = builder.GetInt32(parentKey+".MaxSize", flatObj)

	m.PendingCapacity = builder.GetInt32(parentKey+".PendingCapacity", flatObj)

	m.LoadBalancerId = builder.GetString(parentKey+".LoadBalancerId", flatObj)

	m.RemovingCapacity = builder.GetInt32(parentKey+".RemovingCapacity", flatObj)

	m.ScalingGroupName = builder.GetString(parentKey+".ScalingGroupName", flatObj)

	m.ActiveCapacity = builder.GetInt32(parentKey+".ActiveCapacity", flatObj)

	m.ActiveScalingConfigurationId = builder.GetString(parentKey+".ActiveScalingConfigurationId", flatObj)

	m.ScalingGroupId = builder.GetString(parentKey+".ScalingGroupId", flatObj)

	m.RegionId = builder.GetString(parentKey+".RegionId", flatObj)

	m.TotalCapacity = builder.GetInt32(parentKey+".TotalCapacity", flatObj)

	m.MinSize = builder.GetInt32(parentKey+".MinSize", flatObj)

	m.LifecycleState = builder.GetString(parentKey+".LifecycleState", flatObj)

	m.CreationTime = builder.GetString(parentKey+".CreationTime", flatObj)

	m.RemovalPolicies = make([]string, 0)
	for i := 0; i < (flatObj[parentKey+".RemovalPolicies.length"]).(int); i++ {
		m.RemovalPolicies = append(m.RemovalPolicies, builder.GetString(fmt.Sprintf("%s[%d]", parentKey+".RemovalPolicies", i), flatObj))
	}

	m.DBInstanceIds = make([]string, 0)
	for i := 0; i < (flatObj[parentKey+".DBInstanceIds.length"]).(int); i++ {
		m.DBInstanceIds = append(m.DBInstanceIds, builder.GetString(fmt.Sprintf("%s[%d]", parentKey+".DBInstanceIds", i), flatObj))
	}

}
