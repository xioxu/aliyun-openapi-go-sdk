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
)

type QueueTimeSpan struct {
	MinTimeStamp     int64
	MaxTimeStamp     int64
	ConsumeTimeStamp int64
	MessageQueue     MessageQueue
}

func (m *QueueTimeSpan) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.MinTimeStamp = builder.GetInt64(parentKey+".MinTimeStamp", flatObj)

	m.MaxTimeStamp = builder.GetInt64(parentKey+".MaxTimeStamp", flatObj)

	m.ConsumeTimeStamp = builder.GetInt64(parentKey+".ConsumeTimeStamp", flatObj)

	m.MessageQueue = MessageQueue{}
	m.MessageQueue.BuildProperties(parentKey+".MessageQueue", flatObj)

}
