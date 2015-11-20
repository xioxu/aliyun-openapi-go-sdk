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

package rds_region_DescribeTasksResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type TaskProgressInfo struct {
	DBName             string
	BeginTime          string
	ProgressInfo       string
	FinishTime         string
	TaskAction         string
	TaskId             string
	Progress           string
	ExpectedFinishTime string
	Status             string
}

func (m *TaskProgressInfo) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.DBName = builder.GetString(parentKey+".DBName", flatObj)

	m.BeginTime = builder.GetString(parentKey+".BeginTime", flatObj)

	m.ProgressInfo = builder.GetString(parentKey+".ProgressInfo", flatObj)

	m.FinishTime = builder.GetString(parentKey+".FinishTime", flatObj)

	m.TaskAction = builder.GetString(parentKey+".TaskAction", flatObj)

	m.TaskId = builder.GetString(parentKey+".TaskId", flatObj)

	m.Progress = builder.GetString(parentKey+".Progress", flatObj)

	m.ExpectedFinishTime = builder.GetString(parentKey+".ExpectedFinishTime", flatObj)

	m.Status = builder.GetString(parentKey+".Status", flatObj)

}
