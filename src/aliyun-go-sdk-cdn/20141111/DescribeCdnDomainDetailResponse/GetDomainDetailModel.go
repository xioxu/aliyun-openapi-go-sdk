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

package cdn_DescribeCdnDomainDetailResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type GetDomainDetailModel struct {
	GmtCreated   string
	GmtModified  string
	SourceType   string
	DomainStatus string
	CdnType      string
	Cname        string
	HttpsCname   string
	DomainName   string
	Remark       string
	Sources      []string
}

func (m *GetDomainDetailModel) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.GmtCreated = builder.GetString(parentKey+".GmtCreated", flatObj)

	m.GmtModified = builder.GetString(parentKey+".GmtModified", flatObj)

	m.SourceType = builder.GetString(parentKey+".SourceType", flatObj)

	m.DomainStatus = builder.GetString(parentKey+".DomainStatus", flatObj)

	m.CdnType = builder.GetString(parentKey+".CdnType", flatObj)

	m.Cname = builder.GetString(parentKey+".Cname", flatObj)

	m.HttpsCname = builder.GetString(parentKey+".HttpsCname", flatObj)

	m.DomainName = builder.GetString(parentKey+".DomainName", flatObj)

	m.Remark = builder.GetString(parentKey+".Remark", flatObj)

	m.Sources = make([]string, 0)
	for i := 0; i < (flatObj[parentKey+".Sources.length"]).(int); i++ {
		m.Sources = append(m.Sources, builder.GetString(fmt.Sprintf("%s[%d]", parentKey+".Sources", i), flatObj))
	}

}
