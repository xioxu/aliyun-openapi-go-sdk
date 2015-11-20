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

package rds_DescribeSlowLogRecordsResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type SQLSlowRecord struct {
	HostAddress        string
	DBName             string
	SQLText            string
	QueryTimes         int64
	LockTimes          int64
	ParseRowCounts     int64
	ReturnRowCounts    int64
	ExecutionStartTime string
}

func (m *SQLSlowRecord) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.HostAddress = builder.GetString(parentKey+".HostAddress", flatObj)

	m.DBName = builder.GetString(parentKey+".DBName", flatObj)

	m.SQLText = builder.GetString(parentKey+".SQLText", flatObj)

	m.QueryTimes = builder.GetInt64(parentKey+".QueryTimes", flatObj)

	m.LockTimes = builder.GetInt64(parentKey+".LockTimes", flatObj)

	m.ParseRowCounts = builder.GetInt64(parentKey+".ParseRowCounts", flatObj)

	m.ReturnRowCounts = builder.GetInt64(parentKey+".ReturnRowCounts", flatObj)

	m.ExecutionStartTime = builder.GetString(parentKey+".ExecutionStartTime", flatObj)

}
