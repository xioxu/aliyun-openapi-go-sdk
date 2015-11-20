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

package ons_OnsConsumerStatusResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type ConsumeQueueOffset struct {
	Topic          string
	BrokerName     string
	QueueId        int32
	BrokerOffset   int64
	ConsumerOffset int64
	LastTimestamp  int64
}

func (m *ConsumeQueueOffset) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.Topic = builder.GetString(parentKey+".Topic", flatObj)

	m.BrokerName = builder.GetString(parentKey+".BrokerName", flatObj)

	m.QueueId = builder.GetInt32(parentKey+".QueueId", flatObj)

	m.BrokerOffset = builder.GetInt64(parentKey+".BrokerOffset", flatObj)

	m.ConsumerOffset = builder.GetInt64(parentKey+".ConsumerOffset", flatObj)

	m.LastTimestamp = builder.GetInt64(parentKey+".LastTimestamp", flatObj)

}
