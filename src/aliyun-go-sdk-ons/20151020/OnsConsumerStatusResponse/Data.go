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

type Data struct {
	ConnectionSet              []ConnectionDo
	DetailInTopicList          []DetailInTopicDo
	ConsumerConnectionInfoList []ConsumerConnectionInfoDo
	Online                     bool
	TotalDiff                  int64
	ConsumeTps                 float32
	LastTimestamp              int64
	DelayTime                  int64
	ConsumeModel               string
	SubscriptionSame           bool
	RebalanceOK                bool
}

func (m *Data) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.ConnectionSet = make([]ConnectionDo, 0)
	for i := 0; i < (flatObj[parentKey+".ConnectionSet.length"]).(int); i++ {
		tmp := ConnectionDo{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".ConnectionSet[%d]", i), flatObj)
		m.ConnectionSet = append(m.ConnectionSet, tmp)
	}

	m.DetailInTopicList = make([]DetailInTopicDo, 0)
	for i := 0; i < (flatObj[parentKey+".DetailInTopicList.length"]).(int); i++ {
		tmp := DetailInTopicDo{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".DetailInTopicList[%d]", i), flatObj)
		m.DetailInTopicList = append(m.DetailInTopicList, tmp)
	}

	m.ConsumerConnectionInfoList = make([]ConsumerConnectionInfoDo, 0)
	for i := 0; i < (flatObj[parentKey+".ConsumerConnectionInfoList.length"]).(int); i++ {
		tmp := ConsumerConnectionInfoDo{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".ConsumerConnectionInfoList[%d]", i), flatObj)
		m.ConsumerConnectionInfoList = append(m.ConsumerConnectionInfoList, tmp)
	}

	m.Online = builder.GetBool(parentKey+".Online", flatObj)

	m.TotalDiff = builder.GetInt64(parentKey+".TotalDiff", flatObj)

	m.ConsumeTps = builder.GetFloat32(parentKey+".ConsumeTps", flatObj)

	m.LastTimestamp = builder.GetInt64(parentKey+".LastTimestamp", flatObj)

	m.DelayTime = builder.GetInt64(parentKey+".DelayTime", flatObj)

	m.ConsumeModel = builder.GetString(parentKey+".ConsumeModel", flatObj)

	m.SubscriptionSame = builder.GetBool(parentKey+".SubscriptionSame", flatObj)

	m.RebalanceOK = builder.GetBool(parentKey+".RebalanceOK", flatObj)

}
