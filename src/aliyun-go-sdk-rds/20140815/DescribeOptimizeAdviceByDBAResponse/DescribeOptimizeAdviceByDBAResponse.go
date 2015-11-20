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

package rds_DescribeOptimizeAdviceByDBAResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type DescribeOptimizeAdviceByDBAResponse struct {
	Items             []AdviceByDBA
	TotalRecordsCount int32
	PageNumber        int32
	PageRecordCount   int32
	RequestId         string
}

func (m *DescribeOptimizeAdviceByDBAResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.Items = make([]AdviceByDBA, 0)
	for i := 0; i < (flatObj[parentKey+".Items.length"]).(int); i++ {
		tmp := AdviceByDBA{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".Items[%d]", i), flatObj)
		m.Items = append(m.Items, tmp)
	}

	m.TotalRecordsCount = builder.GetInt32(parentKey+".TotalRecordsCount", flatObj)

	m.PageNumber = builder.GetInt32(parentKey+".PageNumber", flatObj)

	m.PageRecordCount = builder.GetInt32(parentKey+".PageRecordCount", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

}
