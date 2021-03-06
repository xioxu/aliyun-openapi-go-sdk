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

package yundun_DdosLogResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type DdosLogResponse struct {
	LogList      []DdosLog
	AttackStatus int32
	StartTime    string
	EndTime      string
	PageNumber   int32
	PageSize     int32
	TotalCount   int32
	RequestId    string
}

func (m *DdosLogResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.LogList = make([]DdosLog, 0)
	for i := 0; i < (flatObj[parentKey+".LogList.length"]).(int); i++ {
		tmp := DdosLog{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".LogList[%d]", i), flatObj)
		m.LogList = append(m.LogList, tmp)
	}

	m.AttackStatus = builder.GetInt32(parentKey+".AttackStatus", flatObj)

	m.StartTime = builder.GetString(parentKey+".StartTime", flatObj)

	m.EndTime = builder.GetString(parentKey+".EndTime", flatObj)

	m.PageNumber = builder.GetInt32(parentKey+".PageNumber", flatObj)

	m.PageSize = builder.GetInt32(parentKey+".PageSize", flatObj)

	m.TotalCount = builder.GetInt32(parentKey+".TotalCount", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

}
