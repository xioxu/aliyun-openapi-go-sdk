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

package yundun_LogineventLogResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type LoginEventLog struct {
	BlockTimes int32
	SourceIp   string
	Status     int32
	Time       string
	Location   string
}

func (m *LoginEventLog) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.BlockTimes = builder.GetInt32(parentKey+".BlockTimes", flatObj)

	m.SourceIp = builder.GetString(parentKey+".SourceIp", flatObj)

	m.Status = builder.GetInt32(parentKey+".Status", flatObj)

	m.Time = builder.GetString(parentKey+".Time", flatObj)

	m.Location = builder.GetString(parentKey+".Location", flatObj)

}
