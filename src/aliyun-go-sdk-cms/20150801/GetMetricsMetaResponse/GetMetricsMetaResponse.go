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

package cms_GetMetricsMetaResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type GetMetricsMetaResponse struct {
	Code      string
	Message   string
	Success   string
	TraceId   string
	Result    string
	RequestId string
}

func (m *GetMetricsMetaResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.Code = builder.GetString(parentKey+".Code", flatObj)

	m.Message = builder.GetString(parentKey+".Message", flatObj)

	m.Success = builder.GetString(parentKey+".Success", flatObj)

	m.TraceId = builder.GetString(parentKey+".TraceId", flatObj)

	m.Result = builder.GetString(parentKey+".Result", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

}
