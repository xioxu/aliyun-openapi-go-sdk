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

package ecs_DescribeZonesResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type Zone struct {
	ZoneId                    string
	LocalName                 string
	AvailableResourceCreation []string
	AvailableDiskCategories   []string
}

func (m *Zone) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.ZoneId = builder.GetString(parentKey+".ZoneId", flatObj)

	m.LocalName = builder.GetString(parentKey+".LocalName", flatObj)

	m.AvailableResourceCreation = make([]string, 0)
	for i := 0; i < (flatObj[parentKey+".AvailableResourceCreation.length"]).(int); i++ {
		m.AvailableResourceCreation = append(m.AvailableResourceCreation, builder.GetString(fmt.Sprintf("%s[%d]", parentKey+".AvailableResourceCreation", i), flatObj))
	}

	m.AvailableDiskCategories = make([]string, 0)
	for i := 0; i < (flatObj[parentKey+".AvailableDiskCategories.length"]).(int); i++ {
		m.AvailableDiskCategories = append(m.AvailableDiskCategories, builder.GetString(fmt.Sprintf("%s[%d]", parentKey+".AvailableDiskCategories", i), flatObj))
	}

}
