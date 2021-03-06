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

package ons_OnsConsumerAccumulateResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type DetailInTopicDo struct {
	OffsetList    []ConsumeQueueOffset
	Topic         string
	ConsumeTps    float32
	LastTimestamp int64
	DelayTime     int64
}

func (m *DetailInTopicDo) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.OffsetList = make([]ConsumeQueueOffset, 0)
	for i := 0; i < (flatObj[parentKey+".OffsetList.length"]).(int); i++ {
		tmp := ConsumeQueueOffset{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".OffsetList[%d]", i), flatObj)
		m.OffsetList = append(m.OffsetList, tmp)
	}

	m.Topic = builder.GetString(parentKey+".Topic", flatObj)

	m.ConsumeTps = builder.GetFloat32(parentKey+".ConsumeTps", flatObj)

	m.LastTimestamp = builder.GetInt64(parentKey+".LastTimestamp", flatObj)

	m.DelayTime = builder.GetInt64(parentKey+".DelayTime", flatObj)

}
