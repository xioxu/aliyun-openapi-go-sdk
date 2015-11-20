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

package ecs_DescribeVRoutersResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type VRouter struct {
	RegionId      string
	VpcId         string
	VRouterName   string
	Description   string
	VRouterId     string
	CreationTime  string
	RouteTableIds []string
}

func (m *VRouter) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.RegionId = builder.GetString(parentKey+".RegionId", flatObj)

	m.VpcId = builder.GetString(parentKey+".VpcId", flatObj)

	m.VRouterName = builder.GetString(parentKey+".VRouterName", flatObj)

	m.Description = builder.GetString(parentKey+".Description", flatObj)

	m.VRouterId = builder.GetString(parentKey+".VRouterId", flatObj)

	m.CreationTime = builder.GetString(parentKey+".CreationTime", flatObj)

	m.RouteTableIds = make([]string, 0)
	for i := 0; i < (flatObj[parentKey+".RouteTableIds.length"]).(int); i++ {
		m.RouteTableIds = append(m.RouteTableIds, builder.GetString(fmt.Sprintf("%s[%d]", parentKey+".RouteTableIds", i), flatObj))
	}

}
