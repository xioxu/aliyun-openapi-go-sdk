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

package ess_DescribeScalingConfigurationsResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type ScalingConfiguration struct {
	DataDisks                []DataDisk
	ScalingConfigurationId   string
	ScalingConfigurationName string
	ScalingGroupId           string
	ImageId                  string
	InstanceType             string
	SecurityGroupId          string
	InternetChargeType       string
	InternetMaxBandwidthIn   int32
	InternetMaxBandwidthOut  int32
	SystemDiskCategory       string
	LifecycleState           string
	CreationTime             string
}

func (m *ScalingConfiguration) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.DataDisks = make([]DataDisk, 0)
	for i := 0; i < (flatObj[parentKey+".DataDisks.length"]).(int); i++ {
		tmp := DataDisk{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".DataDisks[%d]", i), flatObj)
		m.DataDisks = append(m.DataDisks, tmp)
	}

	m.ScalingConfigurationId = builder.GetString(parentKey+".ScalingConfigurationId", flatObj)

	m.ScalingConfigurationName = builder.GetString(parentKey+".ScalingConfigurationName", flatObj)

	m.ScalingGroupId = builder.GetString(parentKey+".ScalingGroupId", flatObj)

	m.ImageId = builder.GetString(parentKey+".ImageId", flatObj)

	m.InstanceType = builder.GetString(parentKey+".InstanceType", flatObj)

	m.SecurityGroupId = builder.GetString(parentKey+".SecurityGroupId", flatObj)

	m.InternetChargeType = builder.GetString(parentKey+".InternetChargeType", flatObj)

	m.InternetMaxBandwidthIn = builder.GetInt32(parentKey+".InternetMaxBandwidthIn", flatObj)

	m.InternetMaxBandwidthOut = builder.GetInt32(parentKey+".InternetMaxBandwidthOut", flatObj)

	m.SystemDiskCategory = builder.GetString(parentKey+".SystemDiskCategory", flatObj)

	m.LifecycleState = builder.GetString(parentKey+".LifecycleState", flatObj)

	m.CreationTime = builder.GetString(parentKey+".CreationTime", flatObj)

}
