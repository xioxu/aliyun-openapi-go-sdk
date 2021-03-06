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

package rds_region_DescribeRealtimeDiagnosesResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type DescribeRealtimeDiagnosesResponse struct {
	Tasks            []RealtimeDiagnoseTasks
	Engine           string
	TotalRecordCount int32
	PageNumber       int32
	PageRecordCount  int32
	RequestId        string
}

func (m *DescribeRealtimeDiagnosesResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.Tasks = make([]RealtimeDiagnoseTasks, 0)
	for i := 0; i < (flatObj[parentKey+".Tasks.length"]).(int); i++ {
		tmp := RealtimeDiagnoseTasks{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".Tasks[%d]", i), flatObj)
		m.Tasks = append(m.Tasks, tmp)
	}

	m.Engine = builder.GetString(parentKey+".Engine", flatObj)

	m.TotalRecordCount = builder.GetInt32(parentKey+".TotalRecordCount", flatObj)

	m.PageNumber = builder.GetInt32(parentKey+".PageNumber", flatObj)

	m.PageRecordCount = builder.GetInt32(parentKey+".PageRecordCount", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

}
