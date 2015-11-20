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

package ons_OnsMessageGetByTopicResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type OnsRestMessageDo struct {
	PropertyList              []MessageProperty
	Topic                     string
	Flag                      int32
	Body                      string
	QueueId                   int32
	StoreSize                 int32
	QueueOffset               int64
	SysFlag                   int32
	BornTimestamp             int64
	BornHost                  string
	StoreTimestamp            int64
	StoreHost                 string
	MsgId                     string
	CommitLogOffset           int64
	BodyCRC                   int32
	ReconsumeTimes            int32
	PreparedTransactionOffset int64
}

func (m *OnsRestMessageDo) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.PropertyList = make([]MessageProperty, 0)
	for i := 0; i < (flatObj[parentKey+".PropertyList.length"]).(int); i++ {
		tmp := MessageProperty{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".PropertyList[%d]", i), flatObj)
		m.PropertyList = append(m.PropertyList, tmp)
	}

	m.Topic = builder.GetString(parentKey+".Topic", flatObj)

	m.Flag = builder.GetInt32(parentKey+".Flag", flatObj)

	m.Body = builder.GetString(parentKey+".Body", flatObj)

	m.QueueId = builder.GetInt32(parentKey+".QueueId", flatObj)

	m.StoreSize = builder.GetInt32(parentKey+".StoreSize", flatObj)

	m.QueueOffset = builder.GetInt64(parentKey+".QueueOffset", flatObj)

	m.SysFlag = builder.GetInt32(parentKey+".SysFlag", flatObj)

	m.BornTimestamp = builder.GetInt64(parentKey+".BornTimestamp", flatObj)

	m.BornHost = builder.GetString(parentKey+".BornHost", flatObj)

	m.StoreTimestamp = builder.GetInt64(parentKey+".StoreTimestamp", flatObj)

	m.StoreHost = builder.GetString(parentKey+".StoreHost", flatObj)

	m.MsgId = builder.GetString(parentKey+".MsgId", flatObj)

	m.CommitLogOffset = builder.GetInt64(parentKey+".CommitLogOffset", flatObj)

	m.BodyCRC = builder.GetInt32(parentKey+".BodyCRC", flatObj)

	m.ReconsumeTimes = builder.GetInt32(parentKey+".ReconsumeTimes", flatObj)

	m.PreparedTransactionOffset = builder.GetInt64(parentKey+".PreparedTransactionOffset", flatObj)

}
