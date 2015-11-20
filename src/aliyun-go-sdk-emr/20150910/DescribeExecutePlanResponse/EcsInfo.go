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

package emr_DescribeExecutePlanResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type EcsInfo struct {
	EcsRoles        []EcsRole
	RegionId        string
	ZoneId          string
	ImageId         string
	ImageName       string
	SparkVersion    string
	SecurityGroupId string
	TotalCount      int32
	MasterCount     int32
	SlaveCount      int32
}

func (m *EcsInfo) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.EcsRoles = make([]EcsRole, 0)
	for i := 0; i < (flatObj[parentKey+".EcsRoles.length"]).(int); i++ {
		tmp := EcsRole{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".EcsRoles[%d]", i), flatObj)
		m.EcsRoles = append(m.EcsRoles, tmp)
	}

	m.RegionId = builder.GetString(parentKey+".RegionId", flatObj)

	m.ZoneId = builder.GetString(parentKey+".ZoneId", flatObj)

	m.ImageId = builder.GetString(parentKey+".ImageId", flatObj)

	m.ImageName = builder.GetString(parentKey+".ImageName", flatObj)

	m.SparkVersion = builder.GetString(parentKey+".SparkVersion", flatObj)

	m.SecurityGroupId = builder.GetString(parentKey+".SecurityGroupId", flatObj)

	m.TotalCount = builder.GetInt32(parentKey+".TotalCount", flatObj)

	m.MasterCount = builder.GetInt32(parentKey+".MasterCount", flatObj)

	m.SlaveCount = builder.GetInt32(parentKey+".SlaveCount", flatObj)

}
