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

package ecs_DescribeTagKeysResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type DescribeTagKeysResponse struct {
	PageSize   int32
	PageNumber int32
	TotalCount int32
	RequestId  string
	TagKeys    []string
}

func (m *DescribeTagKeysResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.PageSize = builder.GetInt32(parentKey+".PageSize", flatObj)

	m.PageNumber = builder.GetInt32(parentKey+".PageNumber", flatObj)

	m.TotalCount = builder.GetInt32(parentKey+".TotalCount", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

	m.TagKeys = make([]string, 0)
	for i := 0; i < (flatObj[parentKey+".TagKeys.length"]).(int); i++ {
		m.TagKeys = append(m.TagKeys, builder.GetString(fmt.Sprintf("%s[%d]", parentKey+".TagKeys", i), flatObj))
	}

}
