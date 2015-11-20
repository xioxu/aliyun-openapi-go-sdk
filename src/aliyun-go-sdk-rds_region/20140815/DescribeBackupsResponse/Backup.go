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

package rds_region_DescribeBackupsResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type Backup struct {
	BackupId               string
	DBInstanceId           string
	BackupStatus           string
	BackupStartTime        string
	BackupEndTime          string
	BackupType             string
	BackupMode             string
	BackupMethod           string
	BackupDownloadURL      string
	BackupLocation         string
	BackupExtractionStatus string
	BackupScale            string
	BackupDBNames          string
	BackupSize             int64
}

func (m *Backup) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.BackupId = builder.GetString(parentKey+".BackupId", flatObj)

	m.DBInstanceId = builder.GetString(parentKey+".DBInstanceId", flatObj)

	m.BackupStatus = builder.GetString(parentKey+".BackupStatus", flatObj)

	m.BackupStartTime = builder.GetString(parentKey+".BackupStartTime", flatObj)

	m.BackupEndTime = builder.GetString(parentKey+".BackupEndTime", flatObj)

	m.BackupType = builder.GetString(parentKey+".BackupType", flatObj)

	m.BackupMode = builder.GetString(parentKey+".BackupMode", flatObj)

	m.BackupMethod = builder.GetString(parentKey+".BackupMethod", flatObj)

	m.BackupDownloadURL = builder.GetString(parentKey+".BackupDownloadURL", flatObj)

	m.BackupLocation = builder.GetString(parentKey+".BackupLocation", flatObj)

	m.BackupExtractionStatus = builder.GetString(parentKey+".BackupExtractionStatus", flatObj)

	m.BackupScale = builder.GetString(parentKey+".BackupScale", flatObj)

	m.BackupDBNames = builder.GetString(parentKey+".BackupDBNames", flatObj)

	m.BackupSize = builder.GetInt64(parentKey+".BackupSize", flatObj)

}
