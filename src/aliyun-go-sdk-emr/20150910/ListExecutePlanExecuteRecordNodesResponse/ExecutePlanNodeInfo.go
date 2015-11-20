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

package emr_ListExecutePlanExecuteRecordNodesResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type ExecutePlanNodeInfo struct {
	WorkNodeId   string
	WorkNodeName string
	StartTime    string
	RunTime      int64
	JobType      int32
	JobId        int64
	ClusterId    int64
	Status       int32
}

func (m *ExecutePlanNodeInfo) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.WorkNodeId = builder.GetString(parentKey+".WorkNodeId", flatObj)

	m.WorkNodeName = builder.GetString(parentKey+".WorkNodeName", flatObj)

	m.StartTime = builder.GetString(parentKey+".StartTime", flatObj)

	m.RunTime = builder.GetInt64(parentKey+".RunTime", flatObj)

	m.JobType = builder.GetInt32(parentKey+".JobType", flatObj)

	m.JobId = builder.GetInt64(parentKey+".JobId", flatObj)

	m.ClusterId = builder.GetInt64(parentKey+".ClusterId", flatObj)

	m.Status = builder.GetInt32(parentKey+".Status", flatObj)

}
