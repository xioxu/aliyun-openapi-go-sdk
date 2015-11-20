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

package emr_ListExecutePlanExecuteRecordsResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type ExecutePlanRecordInfo struct {
	Id              int64
	ExecutePlanId   int64
	ExecutePlanName string
	StartTime       string
	RunTime         int64
	ClusterId       int64
	ClusterName     string
	ClusterType     int32
	Status          int32
	LogEnable       bool
	LogPath         string
}

func (m *ExecutePlanRecordInfo) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.Id = builder.GetInt64(parentKey+".Id", flatObj)

	m.ExecutePlanId = builder.GetInt64(parentKey+".ExecutePlanId", flatObj)

	m.ExecutePlanName = builder.GetString(parentKey+".ExecutePlanName", flatObj)

	m.StartTime = builder.GetString(parentKey+".StartTime", flatObj)

	m.RunTime = builder.GetInt64(parentKey+".RunTime", flatObj)

	m.ClusterId = builder.GetInt64(parentKey+".ClusterId", flatObj)

	m.ClusterName = builder.GetString(parentKey+".ClusterName", flatObj)

	m.ClusterType = builder.GetInt32(parentKey+".ClusterType", flatObj)

	m.Status = builder.GetInt32(parentKey+".Status", flatObj)

	m.LogEnable = builder.GetBool(parentKey+".LogEnable", flatObj)

	m.LogPath = builder.GetString(parentKey+".LogPath", flatObj)

}
