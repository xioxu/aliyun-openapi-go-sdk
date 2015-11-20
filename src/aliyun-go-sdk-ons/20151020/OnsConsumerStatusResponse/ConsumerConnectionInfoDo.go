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
	"fmt"
)

type ConsumerConnectionInfoDo struct {
	SubscriptionSet        []SubscriptionData
	RunningDataList        []ConsumerRunningDataDo
	Jstack                 []ThreadTrackDo
	ProcessQueueInfoDoList []ProcessQueueInfoDo
	InstanceId             string
	Connection             string
	Language               string
	Version                string
	ConsumeModel           string
	ConsumeType            string
	ThreadCount            int32
	StartTimeStamp         int64
	LastTimeStamp          int64
}

func (m *ConsumerConnectionInfoDo) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.SubscriptionSet = make([]SubscriptionData, 0)
	for i := 0; i < (flatObj[parentKey+".SubscriptionSet.length"]).(int); i++ {
		tmp := SubscriptionData{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".SubscriptionSet[%d]", i), flatObj)
		m.SubscriptionSet = append(m.SubscriptionSet, tmp)
	}

	m.RunningDataList = make([]ConsumerRunningDataDo, 0)
	for i := 0; i < (flatObj[parentKey+".RunningDataList.length"]).(int); i++ {
		tmp := ConsumerRunningDataDo{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".RunningDataList[%d]", i), flatObj)
		m.RunningDataList = append(m.RunningDataList, tmp)
	}

	m.Jstack = make([]ThreadTrackDo, 0)
	for i := 0; i < (flatObj[parentKey+".Jstack.length"]).(int); i++ {
		tmp := ThreadTrackDo{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".Jstack[%d]", i), flatObj)
		m.Jstack = append(m.Jstack, tmp)
	}

	m.ProcessQueueInfoDoList = make([]ProcessQueueInfoDo, 0)
	for i := 0; i < (flatObj[parentKey+".ProcessQueueInfoDoList.length"]).(int); i++ {
		tmp := ProcessQueueInfoDo{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".ProcessQueueInfoDoList[%d]", i), flatObj)
		m.ProcessQueueInfoDoList = append(m.ProcessQueueInfoDoList, tmp)
	}

	m.InstanceId = builder.GetString(parentKey+".InstanceId", flatObj)

	m.Connection = builder.GetString(parentKey+".Connection", flatObj)

	m.Language = builder.GetString(parentKey+".Language", flatObj)

	m.Version = builder.GetString(parentKey+".Version", flatObj)

	m.ConsumeModel = builder.GetString(parentKey+".ConsumeModel", flatObj)

	m.ConsumeType = builder.GetString(parentKey+".ConsumeType", flatObj)

	m.ThreadCount = builder.GetInt32(parentKey+".ThreadCount", flatObj)

	m.StartTimeStamp = builder.GetInt64(parentKey+".StartTimeStamp", flatObj)

	m.LastTimeStamp = builder.GetInt64(parentKey+".LastTimeStamp", flatObj)

}
