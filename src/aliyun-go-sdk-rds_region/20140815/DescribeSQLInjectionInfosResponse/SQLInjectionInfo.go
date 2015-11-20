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

package rds_region_DescribeSQLInjectionInfosResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type SQLInjectionInfo struct {
	DBName         string
	SQLText        string
	LatencyTime    string
	HostAddress    string
	ExecuteTime    string
	AccountName    string
	EffectRowCount string
}

func (m *SQLInjectionInfo) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.DBName = builder.GetString(parentKey+".DBName", flatObj)

	m.SQLText = builder.GetString(parentKey+".SQLText", flatObj)

	m.LatencyTime = builder.GetString(parentKey+".LatencyTime", flatObj)

	m.HostAddress = builder.GetString(parentKey+".HostAddress", flatObj)

	m.ExecuteTime = builder.GetString(parentKey+".ExecuteTime", flatObj)

	m.AccountName = builder.GetString(parentKey+".AccountName", flatObj)

	m.EffectRowCount = builder.GetString(parentKey+".EffectRowCount", flatObj)

}
