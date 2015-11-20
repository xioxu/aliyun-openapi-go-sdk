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

package sts_AssumeRoleResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type Credentials struct {
	SecurityToken   string
	AccessKeySecret string
	AccessKeyId     string
	Expiration      string
}

func (m *Credentials) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.SecurityToken = builder.GetString(parentKey+".SecurityToken", flatObj)

	m.AccessKeySecret = builder.GetString(parentKey+".AccessKeySecret", flatObj)

	m.AccessKeyId = builder.GetString(parentKey+".AccessKeyId", flatObj)

	m.Expiration = builder.GetString(parentKey+".Expiration", flatObj)

}
