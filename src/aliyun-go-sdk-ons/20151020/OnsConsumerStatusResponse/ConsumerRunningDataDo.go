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

type ConsumerRunningDataDo struct {
	ConsumerId         string
	Topic              string
	Rt                 float32
	OkTps              float32
	FailedTps          int64
	FailedCountPerHour int64
}

func (m *ConsumerRunningDataDo) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.ConsumerId = builder.GetString(parentKey+".ConsumerId", flatObj)

	m.Topic = builder.GetString(parentKey+".Topic", flatObj)

	m.Rt = builder.GetFloat32(parentKey+".Rt", flatObj)

	m.OkTps = builder.GetFloat32(parentKey+".OkTps", flatObj)

	m.FailedTps = builder.GetInt64(parentKey+".FailedTps", flatObj)

	m.FailedCountPerHour = builder.GetInt64(parentKey+".FailedCountPerHour", flatObj)

}
