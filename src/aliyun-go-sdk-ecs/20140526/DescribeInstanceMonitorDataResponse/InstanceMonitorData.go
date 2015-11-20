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

package ecs_DescribeInstanceMonitorDataResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type InstanceMonitorData struct {
	InstanceId        string
	CPU               int32
	IntranetRX        int32
	IntranetTX        int32
	IntranetBandwidth int32
	InternetRX        int32
	InternetTX        int32
	InternetBandwidth int32
	IOPSRead          int32
	IOPSWrite         int32
	BPSRead           int32
	BPSWrite          int32
	TimeStamp         string
}

func (m *InstanceMonitorData) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.InstanceId = builder.GetString(parentKey+".InstanceId", flatObj)

	m.CPU = builder.GetInt32(parentKey+".CPU", flatObj)

	m.IntranetRX = builder.GetInt32(parentKey+".IntranetRX", flatObj)

	m.IntranetTX = builder.GetInt32(parentKey+".IntranetTX", flatObj)

	m.IntranetBandwidth = builder.GetInt32(parentKey+".IntranetBandwidth", flatObj)

	m.InternetRX = builder.GetInt32(parentKey+".InternetRX", flatObj)

	m.InternetTX = builder.GetInt32(parentKey+".InternetTX", flatObj)

	m.InternetBandwidth = builder.GetInt32(parentKey+".InternetBandwidth", flatObj)

	m.IOPSRead = builder.GetInt32(parentKey+".IOPSRead", flatObj)

	m.IOPSWrite = builder.GetInt32(parentKey+".IOPSWrite", flatObj)

	m.BPSRead = builder.GetInt32(parentKey+".BPSRead", flatObj)

	m.BPSWrite = builder.GetInt32(parentKey+".BPSWrite", flatObj)

	m.TimeStamp = builder.GetString(parentKey+".TimeStamp", flatObj)

}
