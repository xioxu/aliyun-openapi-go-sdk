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

package ecs_DescribeEipMonitorDataResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type EipMonitorData struct {
	EipRX        int32
	EipTX        int32
	EipFlow      int32
	EipBandwidth int32
	EipPackets   int32
	TimeStamp    string
}

func (m *EipMonitorData) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.EipRX = builder.GetInt32(parentKey+".EipRX", flatObj)

	m.EipTX = builder.GetInt32(parentKey+".EipTX", flatObj)

	m.EipFlow = builder.GetInt32(parentKey+".EipFlow", flatObj)

	m.EipBandwidth = builder.GetInt32(parentKey+".EipBandwidth", flatObj)

	m.EipPackets = builder.GetInt32(parentKey+".EipPackets", flatObj)

	m.TimeStamp = builder.GetString(parentKey+".TimeStamp", flatObj)

}
