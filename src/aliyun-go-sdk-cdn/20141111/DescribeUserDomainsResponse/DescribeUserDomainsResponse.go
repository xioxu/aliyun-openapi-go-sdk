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
	"fmt"
)

type DescribeUserDomainsResponse struct {
	Domains    []PageData
	PageNumber int64
	PageSize   int64
	TotalCount int64
	RequestId  string
}

func (m *DescribeUserDomainsResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.Domains = make([]PageData, 0)
	for i := 0; i < (flatObj[parentKey+".Domains.length"]).(int); i++ {
		tmp := PageData{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".Domains[%d]", i), flatObj)
		m.Domains = append(m.Domains, tmp)
	}

	m.PageNumber = builder.GetInt64(parentKey+".PageNumber", flatObj)

	m.PageSize = builder.GetInt64(parentKey+".PageSize", flatObj)

	m.TotalCount = builder.GetInt64(parentKey+".TotalCount", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

}
