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

package ram_UpdateRoleResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type Role struct {
	RoleId                   string
	RoleName                 string
	Arn                      string
	Description              string
	AssumeRolePolicyDocument string
	CreateDate               string
	UpdateDate               string
}

func (m *Role) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.RoleId = builder.GetString(parentKey+".RoleId", flatObj)

	m.RoleName = builder.GetString(parentKey+".RoleName", flatObj)

	m.Arn = builder.GetString(parentKey+".Arn", flatObj)

	m.Description = builder.GetString(parentKey+".Description", flatObj)

	m.AssumeRolePolicyDocument = builder.GetString(parentKey+".AssumeRolePolicyDocument", flatObj)

	m.CreateDate = builder.GetString(parentKey+".CreateDate", flatObj)

	m.UpdateDate = builder.GetString(parentKey+".UpdateDate", flatObj)

}
