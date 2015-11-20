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

package ecs_DescribeDiskMonitorDataResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type DiskMonitorData struct {
	DiskId    string
	IOPSRead  int32
	IOPSWrite int32
	IOPSTotal int32
	BPSRead   int32
	BPSWrite  int32
	BPSTotal  int32
	TimeStamp string
}

func (m *DiskMonitorData) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.DiskId = builder.GetString(parentKey+".DiskId", flatObj)

	m.IOPSRead = builder.GetInt32(parentKey+".IOPSRead", flatObj)

	m.IOPSWrite = builder.GetInt32(parentKey+".IOPSWrite", flatObj)

	m.IOPSTotal = builder.GetInt32(parentKey+".IOPSTotal", flatObj)

	m.BPSRead = builder.GetInt32(parentKey+".BPSRead", flatObj)

	m.BPSWrite = builder.GetInt32(parentKey+".BPSWrite", flatObj)

	m.BPSTotal = builder.GetInt32(parentKey+".BPSTotal", flatObj)

	m.TimeStamp = builder.GetString(parentKey+".TimeStamp", flatObj)

}
