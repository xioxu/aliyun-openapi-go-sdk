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

package emr_ListExecutePlansResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type ExecutePlanInfo struct {
	Id              int64
	Name            string
	LastRunStatus   string
	RunTime         int64
	ClusterName     string
	IsCreateCluster bool
	Stragety        int32
	Status          string
	TimeInterval    int32
	StartTime       string
	TimeUnit        string
}

func (m *ExecutePlanInfo) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.Id = builder.GetInt64(parentKey+".Id", flatObj)

	m.Name = builder.GetString(parentKey+".Name", flatObj)

	m.LastRunStatus = builder.GetString(parentKey+".LastRunStatus", flatObj)

	m.RunTime = builder.GetInt64(parentKey+".RunTime", flatObj)

	m.ClusterName = builder.GetString(parentKey+".ClusterName", flatObj)

	m.IsCreateCluster = builder.GetBool(parentKey+".IsCreateCluster", flatObj)

	m.Stragety = builder.GetInt32(parentKey+".Stragety", flatObj)

	m.Status = builder.GetString(parentKey+".Status", flatObj)

	m.TimeInterval = builder.GetInt32(parentKey+".TimeInterval", flatObj)

	m.StartTime = builder.GetString(parentKey+".StartTime", flatObj)

	m.TimeUnit = builder.GetString(parentKey+".TimeUnit", flatObj)

}
