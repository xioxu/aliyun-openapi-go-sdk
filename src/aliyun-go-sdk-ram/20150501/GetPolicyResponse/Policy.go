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

package ram_GetPolicyResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type Policy struct {
	PolicyName      string
	PolicyType      string
	Description     string
	DefaultVersion  string
	PolicyDocument  string
	CreateDate      string
	UpdateDate      string
	AttachmentCount int32
}

func (m *Policy) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.PolicyName = builder.GetString(parentKey+".PolicyName", flatObj)

	m.PolicyType = builder.GetString(parentKey+".PolicyType", flatObj)

	m.Description = builder.GetString(parentKey+".Description", flatObj)

	m.DefaultVersion = builder.GetString(parentKey+".DefaultVersion", flatObj)

	m.PolicyDocument = builder.GetString(parentKey+".PolicyDocument", flatObj)

	m.CreateDate = builder.GetString(parentKey+".CreateDate", flatObj)

	m.UpdateDate = builder.GetString(parentKey+".UpdateDate", flatObj)

	m.AttachmentCount = builder.GetInt32(parentKey+".AttachmentCount", flatObj)

}
