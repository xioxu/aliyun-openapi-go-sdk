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

package rds_region_DescribeResourceUsageResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type DescribeResourceUsageResponse struct {
	DBInstanceId   string
	Engine         string
	DiskUsed       int64
	DataSize       int64
	LogSize        int64
	BackupSize     int64
	SQLSize        int64
	ColdBackupSize int64
	RequestId      string
}

func (m *DescribeResourceUsageResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.DBInstanceId = builder.GetString(parentKey+".DBInstanceId", flatObj)

	m.Engine = builder.GetString(parentKey+".Engine", flatObj)

	m.DiskUsed = builder.GetInt64(parentKey+".DiskUsed", flatObj)

	m.DataSize = builder.GetInt64(parentKey+".DataSize", flatObj)

	m.LogSize = builder.GetInt64(parentKey+".LogSize", flatObj)

	m.BackupSize = builder.GetInt64(parentKey+".BackupSize", flatObj)

	m.SQLSize = builder.GetInt64(parentKey+".SQLSize", flatObj)

	m.ColdBackupSize = builder.GetInt64(parentKey+".ColdBackupSize", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

}
