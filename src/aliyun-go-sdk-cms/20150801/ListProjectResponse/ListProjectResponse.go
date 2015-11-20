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

package cms_ListProjectResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type ListProjectResponse struct {
	Code       string
	Message    string
	Success    string
	TraceId    string
	RequestId  string
	Datapoints []string
}

func (m *ListProjectResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.Code = builder.GetString(parentKey+".Code", flatObj)

	m.Message = builder.GetString(parentKey+".Message", flatObj)

	m.Success = builder.GetString(parentKey+".Success", flatObj)

	m.TraceId = builder.GetString(parentKey+".TraceId", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

	m.Datapoints = make([]string, 0)
	for i := 0; i < (flatObj[parentKey+".Datapoints.length"]).(int); i++ {
		m.Datapoints = append(m.Datapoints, builder.GetString(fmt.Sprintf("%s[%d]", parentKey+".Datapoints", i), flatObj))
	}

}
