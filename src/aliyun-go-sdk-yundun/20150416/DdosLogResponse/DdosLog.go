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

package yundun_DdosLogResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type DdosLog struct {
	StartTime    string
	EndTime      string
	Reason       string
	Status       int32
	Bps          int64
	Pps          int64
	Qps          int64
	AttackType   string
	AttackIpList string
	Type         int32
}

func (m *DdosLog) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.StartTime = builder.GetString(parentKey+".StartTime", flatObj)

	m.EndTime = builder.GetString(parentKey+".EndTime", flatObj)

	m.Reason = builder.GetString(parentKey+".Reason", flatObj)

	m.Status = builder.GetInt32(parentKey+".Status", flatObj)

	m.Bps = builder.GetInt64(parentKey+".Bps", flatObj)

	m.Pps = builder.GetInt64(parentKey+".Pps", flatObj)

	m.Qps = builder.GetInt64(parentKey+".Qps", flatObj)

	m.AttackType = builder.GetString(parentKey+".AttackType", flatObj)

	m.AttackIpList = builder.GetString(parentKey+".AttackIpList", flatObj)

	m.Type = builder.GetInt32(parentKey+".Type", flatObj)

}
