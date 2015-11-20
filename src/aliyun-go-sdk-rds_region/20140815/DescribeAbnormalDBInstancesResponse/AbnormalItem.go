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

package rds_region_DescribeAbnormalDBInstancesResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type AbnormalItem struct {
	CheckTime      string
	CheckItem      string
	AbnormalReason string
	AbnormalValue  string
	AbnormalDetail string
	AdviceKey      string
	AdviseValue    []string
}

func (m *AbnormalItem) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.CheckTime = builder.GetString(parentKey+".CheckTime", flatObj)

	m.CheckItem = builder.GetString(parentKey+".CheckItem", flatObj)

	m.AbnormalReason = builder.GetString(parentKey+".AbnormalReason", flatObj)

	m.AbnormalValue = builder.GetString(parentKey+".AbnormalValue", flatObj)

	m.AbnormalDetail = builder.GetString(parentKey+".AbnormalDetail", flatObj)

	m.AdviceKey = builder.GetString(parentKey+".AdviceKey", flatObj)

	m.AdviseValue = make([]string, 0)
	for i := 0; i < (flatObj[parentKey+".AdviseValue.length"]).(int); i++ {
		m.AdviseValue = append(m.AdviseValue, builder.GetString(fmt.Sprintf("%s[%d]", parentKey+".AdviseValue", i), flatObj))
	}

}
