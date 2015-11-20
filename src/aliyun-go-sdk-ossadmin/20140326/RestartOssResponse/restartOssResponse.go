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

package ossadmin_RestartOssResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type RestartOssResponse struct {
	Code           string
	Message        string
	Success        bool
	aliUid         int64
	instanceId     string
	instacneStatus string
	instanceName   string
	startTime      string
	endTime        string
	RequestId      string
}

func (m *RestartOssResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.Code = builder.GetString(parentKey+".Code", flatObj)

	m.Message = builder.GetString(parentKey+".Message", flatObj)

	m.Success = builder.GetBool(parentKey+".Success", flatObj)

	m.aliUid = builder.GetInt64(parentKey+".aliUid", flatObj)

	m.instanceId = builder.GetString(parentKey+".instanceId", flatObj)

	m.instacneStatus = builder.GetString(parentKey+".instacneStatus", flatObj)

	m.instanceName = builder.GetString(parentKey+".instanceName", flatObj)

	m.startTime = builder.GetString(parentKey+".startTime", flatObj)

	m.endTime = builder.GetString(parentKey+".endTime", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

}
