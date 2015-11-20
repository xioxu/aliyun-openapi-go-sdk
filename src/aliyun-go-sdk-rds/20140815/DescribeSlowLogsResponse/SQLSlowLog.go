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

package rds_DescribeSlowLogsResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type SQLSlowLog struct {
	SlowLogId                     int64
	SQLId                         int64
	DBName                        string
	SQLText                       string
	MySQLTotalExecutionCounts     int64
	MySQLTotalExecutionTimes      int64
	TotalLockTimes                int64
	MaxLockTime                   int64
	ParseTotalRowCounts           int64
	ParseMaxRowCount              int64
	ReturnTotalRowCounts          int64
	ReturnMaxRowCount             int64
	CreateTime                    string
	SQLServerTotalExecutionCounts int64
	SQLServerTotalExecutionTimes  int64
	TotalLogicalReadCounts        int64
	TotalPhysicalReadCounts       int64
	ReportTime                    string
	MaxExecutionTime              int64
	AvgExecutionTime              int64
}

func (m *SQLSlowLog) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.SlowLogId = builder.GetInt64(parentKey+".SlowLogId", flatObj)

	m.SQLId = builder.GetInt64(parentKey+".SQLId", flatObj)

	m.DBName = builder.GetString(parentKey+".DBName", flatObj)

	m.SQLText = builder.GetString(parentKey+".SQLText", flatObj)

	m.MySQLTotalExecutionCounts = builder.GetInt64(parentKey+".MySQLTotalExecutionCounts", flatObj)

	m.MySQLTotalExecutionTimes = builder.GetInt64(parentKey+".MySQLTotalExecutionTimes", flatObj)

	m.TotalLockTimes = builder.GetInt64(parentKey+".TotalLockTimes", flatObj)

	m.MaxLockTime = builder.GetInt64(parentKey+".MaxLockTime", flatObj)

	m.ParseTotalRowCounts = builder.GetInt64(parentKey+".ParseTotalRowCounts", flatObj)

	m.ParseMaxRowCount = builder.GetInt64(parentKey+".ParseMaxRowCount", flatObj)

	m.ReturnTotalRowCounts = builder.GetInt64(parentKey+".ReturnTotalRowCounts", flatObj)

	m.ReturnMaxRowCount = builder.GetInt64(parentKey+".ReturnMaxRowCount", flatObj)

	m.CreateTime = builder.GetString(parentKey+".CreateTime", flatObj)

	m.SQLServerTotalExecutionCounts = builder.GetInt64(parentKey+".SQLServerTotalExecutionCounts", flatObj)

	m.SQLServerTotalExecutionTimes = builder.GetInt64(parentKey+".SQLServerTotalExecutionTimes", flatObj)

	m.TotalLogicalReadCounts = builder.GetInt64(parentKey+".TotalLogicalReadCounts", flatObj)

	m.TotalPhysicalReadCounts = builder.GetInt64(parentKey+".TotalPhysicalReadCounts", flatObj)

	m.ReportTime = builder.GetString(parentKey+".ReportTime", flatObj)

	m.MaxExecutionTime = builder.GetInt64(parentKey+".MaxExecutionTime", flatObj)

	m.AvgExecutionTime = builder.GetInt64(parentKey+".AvgExecutionTime", flatObj)

}
