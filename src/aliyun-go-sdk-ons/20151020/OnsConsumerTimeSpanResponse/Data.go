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

package ons_OnsConsumerTimeSpanResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type Data struct {
	QueueTimeSpans   []QueueTimeSpan
	Topic            string
	MinTimeStamp     int64
	MaxTimeStamp     int64
	ConsumeTimeStamp int64
}

func (m *Data) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.QueueTimeSpans = make([]QueueTimeSpan, 0)
	for i := 0; i < (flatObj[parentKey+".QueueTimeSpans.length"]).(int); i++ {
		tmp := QueueTimeSpan{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".QueueTimeSpans[%d]", i), flatObj)
		m.QueueTimeSpans = append(m.QueueTimeSpans, tmp)
	}

	m.Topic = builder.GetString(parentKey+".Topic", flatObj)

	m.MinTimeStamp = builder.GetInt64(parentKey+".MinTimeStamp", flatObj)

	m.MaxTimeStamp = builder.GetInt64(parentKey+".MaxTimeStamp", flatObj)

	m.ConsumeTimeStamp = builder.GetInt64(parentKey+".ConsumeTimeStamp", flatObj)

}
