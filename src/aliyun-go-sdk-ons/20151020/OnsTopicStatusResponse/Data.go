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

package ons_OnsTopicStatusResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type Data struct {
	OffsetTable   []TopicQueueOffset
	TotalCount    int64
	LastTimeStamp int64
}

func (m *Data) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.OffsetTable = make([]TopicQueueOffset, 0)
	for i := 0; i < (flatObj[parentKey+".OffsetTable.length"]).(int); i++ {
		tmp := TopicQueueOffset{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".OffsetTable[%d]", i), flatObj)
		m.OffsetTable = append(m.OffsetTable, tmp)
	}

	m.TotalCount = builder.GetInt64(parentKey+".TotalCount", flatObj)

	m.LastTimeStamp = builder.GetInt64(parentKey+".LastTimeStamp", flatObj)

}
