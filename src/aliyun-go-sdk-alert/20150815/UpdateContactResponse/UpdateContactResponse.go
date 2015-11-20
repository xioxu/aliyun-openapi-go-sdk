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

package alert_UpdateContactResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type UpdateContactResponse struct {
	code      string
	message   string
	success   string
	traceId   string
	RequestId string
}

func (m *UpdateContactResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.code = builder.GetString(parentKey+".code", flatObj)

	m.message = builder.GetString(parentKey+".message", flatObj)

	m.success = builder.GetString(parentKey+".success", flatObj)

	m.traceId = builder.GetString(parentKey+".traceId", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

}
