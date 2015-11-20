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

package ecs_DescribeRouteTablesResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type RouteTable struct {
	RouteEntrys    []RouteEntry
	VRouterId      string
	RouteTableId   string
	RouteTableType string
	CreationTime   string
}

func (m *RouteTable) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.RouteEntrys = make([]RouteEntry, 0)
	for i := 0; i < (flatObj[parentKey+".RouteEntrys.length"]).(int); i++ {
		tmp := RouteEntry{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".RouteEntrys[%d]", i), flatObj)
		m.RouteEntrys = append(m.RouteEntrys, tmp)
	}

	m.VRouterId = builder.GetString(parentKey+".VRouterId", flatObj)

	m.RouteTableId = builder.GetString(parentKey+".RouteTableId", flatObj)

	m.RouteTableType = builder.GetString(parentKey+".RouteTableType", flatObj)

	m.CreationTime = builder.GetString(parentKey+".CreationTime", flatObj)

}
