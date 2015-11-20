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

type DescribeEipAddressesResponse struct {
	EipAddresses []EipAddress
	TotalCount   int32
	PageNumber   int32
	PageSize     int32
	RequestId    string
}

func (m *DescribeEipAddressesResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.EipAddresses = make([]EipAddress, 0)
	for i := 0; i < (flatObj[parentKey+".EipAddresses.length"]).(int); i++ {
		tmp := EipAddress{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".EipAddresses[%d]", i), flatObj)
		m.EipAddresses = append(m.EipAddresses, tmp)
	}

	m.TotalCount = builder.GetInt32(parentKey+".TotalCount", flatObj)

	m.PageNumber = builder.GetInt32(parentKey+".PageNumber", flatObj)

	m.PageSize = builder.GetInt32(parentKey+".PageSize", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

}
