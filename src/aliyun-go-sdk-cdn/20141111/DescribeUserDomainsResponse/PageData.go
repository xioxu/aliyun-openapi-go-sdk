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

package cdn_DescribeUserDomainsResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type PageData struct {
	DomainName   string
	Cname        string
	CdnType      string
	DomainStatus string
	GmtCreated   string
	GmtModified  string
}

func (m *PageData) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.DomainName = builder.GetString(parentKey+".DomainName", flatObj)

	m.Cname = builder.GetString(parentKey+".Cname", flatObj)

	m.CdnType = builder.GetString(parentKey+".CdnType", flatObj)

	m.DomainStatus = builder.GetString(parentKey+".DomainStatus", flatObj)

	m.GmtCreated = builder.GetString(parentKey+".GmtCreated", flatObj)

	m.GmtModified = builder.GetString(parentKey+".GmtModified", flatObj)

}
