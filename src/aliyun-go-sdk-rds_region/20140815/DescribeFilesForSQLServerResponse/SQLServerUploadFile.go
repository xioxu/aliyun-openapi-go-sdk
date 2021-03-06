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

package rds_region_DescribeFilesForSQLServerResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type SQLServerUploadFile struct {
	DBName            string
	FileName          string
	FileSize          int64
	InternetFtpServer string
	InternetPort      int32
	IntranetFtpserver string
	Intranetport      int32
	UserName          string
	Password          string
	FileStatus        string
	Description       string
	CreationTime      string
}

func (m *SQLServerUploadFile) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.DBName = builder.GetString(parentKey+".DBName", flatObj)

	m.FileName = builder.GetString(parentKey+".FileName", flatObj)

	m.FileSize = builder.GetInt64(parentKey+".FileSize", flatObj)

	m.InternetFtpServer = builder.GetString(parentKey+".InternetFtpServer", flatObj)

	m.InternetPort = builder.GetInt32(parentKey+".InternetPort", flatObj)

	m.IntranetFtpserver = builder.GetString(parentKey+".IntranetFtpserver", flatObj)

	m.Intranetport = builder.GetInt32(parentKey+".Intranetport", flatObj)

	m.UserName = builder.GetString(parentKey+".UserName", flatObj)

	m.Password = builder.GetString(parentKey+".Password", flatObj)

	m.FileStatus = builder.GetString(parentKey+".FileStatus", flatObj)

	m.Description = builder.GetString(parentKey+".Description", flatObj)

	m.CreationTime = builder.GetString(parentKey+".CreationTime", flatObj)

}
