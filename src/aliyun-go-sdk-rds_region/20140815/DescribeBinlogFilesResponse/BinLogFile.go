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

package rds_region_DescribeBinlogFilesResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type BinLogFile struct {
	FileSize        int64
	LogBeginTime    string
	LogEndTime      string
	DownloadLink    string
	LinkExpiredTime string
}

func (m *BinLogFile) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.FileSize = builder.GetInt64(parentKey+".FileSize", flatObj)

	m.LogBeginTime = builder.GetString(parentKey+".LogBeginTime", flatObj)

	m.LogEndTime = builder.GetString(parentKey+".LogEndTime", flatObj)

	m.DownloadLink = builder.GetString(parentKey+".DownloadLink", flatObj)

	m.LinkExpiredTime = builder.GetString(parentKey+".LinkExpiredTime", flatObj)

}
