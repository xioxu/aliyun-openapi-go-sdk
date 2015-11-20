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

package ecs_DescribeVpcsResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type Vpc struct {
	VpcId        string
	RegionId     string
	Status       string
	VpcName      string
	CreationTime string
	CidrBlock    string
	VRouterId    string
	Description  string
	VSwitchIds   []string
	UserCidrs    []string
}

func (m *Vpc) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.VpcId = builder.GetString(parentKey+".VpcId", flatObj)

	m.RegionId = builder.GetString(parentKey+".RegionId", flatObj)

	m.Status = builder.GetString(parentKey+".Status", flatObj)

	m.VpcName = builder.GetString(parentKey+".VpcName", flatObj)

	m.CreationTime = builder.GetString(parentKey+".CreationTime", flatObj)

	m.CidrBlock = builder.GetString(parentKey+".CidrBlock", flatObj)

	m.VRouterId = builder.GetString(parentKey+".VRouterId", flatObj)

	m.Description = builder.GetString(parentKey+".Description", flatObj)

	m.VSwitchIds = make([]string, 0)
	for i := 0; i < (flatObj[parentKey+".VSwitchIds.length"]).(int); i++ {
		m.VSwitchIds = append(m.VSwitchIds, builder.GetString(fmt.Sprintf("%s[%d]", parentKey+".VSwitchIds", i), flatObj))
	}

	m.UserCidrs = make([]string, 0)
	for i := 0; i < (flatObj[parentKey+".UserCidrs.length"]).(int); i++ {
		m.UserCidrs = append(m.UserCidrs, builder.GetString(fmt.Sprintf("%s[%d]", parentKey+".UserCidrs", i), flatObj))
	}

}
