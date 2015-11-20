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

package drds_DescribeDDLTaskResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type Data struct {
	RequestId   string
	TargetId    string
	TaskDetail  string
	TaskStatus  int32
	TaskPhase   string
	TaskType    int32
	TaskName    string
	GmtCreate   int64
	AllowCancel string
	ErrMsg      string
}

func (m *Data) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

	m.TargetId = builder.GetString(parentKey+".TargetId", flatObj)

	m.TaskDetail = builder.GetString(parentKey+".TaskDetail", flatObj)

	m.TaskStatus = builder.GetInt32(parentKey+".TaskStatus", flatObj)

	m.TaskPhase = builder.GetString(parentKey+".TaskPhase", flatObj)

	m.TaskType = builder.GetInt32(parentKey+".TaskType", flatObj)

	m.TaskName = builder.GetString(parentKey+".TaskName", flatObj)

	m.GmtCreate = builder.GetInt64(parentKey+".GmtCreate", flatObj)

	m.AllowCancel = builder.GetString(parentKey+".AllowCancel", flatObj)

	m.ErrMsg = builder.GetString(parentKey+".ErrMsg", flatObj)

}
