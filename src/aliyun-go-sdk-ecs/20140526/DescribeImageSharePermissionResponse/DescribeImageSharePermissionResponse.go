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

package ecs_DescribeImageSharePermissionResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type DescribeImageSharePermissionResponse struct {
	ShareGroups []ShareGroup
	Accounts    []Account
	RegionId    string
	TotalCount  int32
	PageNumber  int32
	PageSize    int32
	ImageId     string
	RequestId   string
}

func (m *DescribeImageSharePermissionResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.ShareGroups = make([]ShareGroup, 0)
	for i := 0; i < (flatObj[parentKey+".ShareGroups.length"]).(int); i++ {
		tmp := ShareGroup{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".ShareGroups[%d]", i), flatObj)
		m.ShareGroups = append(m.ShareGroups, tmp)
	}

	m.Accounts = make([]Account, 0)
	for i := 0; i < (flatObj[parentKey+".Accounts.length"]).(int); i++ {
		tmp := Account{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".Accounts[%d]", i), flatObj)
		m.Accounts = append(m.Accounts, tmp)
	}

	m.RegionId = builder.GetString(parentKey+".RegionId", flatObj)

	m.TotalCount = builder.GetInt32(parentKey+".TotalCount", flatObj)

	m.PageNumber = builder.GetInt32(parentKey+".PageNumber", flatObj)

	m.PageSize = builder.GetInt32(parentKey+".PageSize", flatObj)

	m.ImageId = builder.GetString(parentKey+".ImageId", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

}
