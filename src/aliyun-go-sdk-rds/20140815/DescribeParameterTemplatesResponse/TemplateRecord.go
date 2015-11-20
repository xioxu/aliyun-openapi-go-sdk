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

package rds_DescribeParameterTemplatesResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type TemplateRecord struct {
	ParameterName        string
	ParameterValue       string
	ForceModify          string
	ForceRestart         string
	CheckingCode         string
	ParameterDescription string
}

func (m *TemplateRecord) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.ParameterName = builder.GetString(parentKey+".ParameterName", flatObj)

	m.ParameterValue = builder.GetString(parentKey+".ParameterValue", flatObj)

	m.ForceModify = builder.GetString(parentKey+".ForceModify", flatObj)

	m.ForceRestart = builder.GetString(parentKey+".ForceRestart", flatObj)

	m.CheckingCode = builder.GetString(parentKey+".CheckingCode", flatObj)

	m.ParameterDescription = builder.GetString(parentKey+".ParameterDescription", flatObj)

}
