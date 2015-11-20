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

package rds_DescribeSQLLogReportListResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type Item struct {
	LatencyTopNItems []LatencyTopNItem
	QPSTopNItems     []QPSTopNItem
	ReportTime       string
}

func (m *Item) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.LatencyTopNItems = make([]LatencyTopNItem, 0)
	for i := 0; i < (flatObj[parentKey+".LatencyTopNItems.length"]).(int); i++ {
		tmp := LatencyTopNItem{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".LatencyTopNItems[%d]", i), flatObj)
		m.LatencyTopNItems = append(m.LatencyTopNItems, tmp)
	}

	m.QPSTopNItems = make([]QPSTopNItem, 0)
	for i := 0; i < (flatObj[parentKey+".QPSTopNItems.length"]).(int); i++ {
		tmp := QPSTopNItem{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".QPSTopNItems[%d]", i), flatObj)
		m.QPSTopNItems = append(m.QPSTopNItems, tmp)
	}

	m.ReportTime = builder.GetString(parentKey+".ReportTime", flatObj)

}
