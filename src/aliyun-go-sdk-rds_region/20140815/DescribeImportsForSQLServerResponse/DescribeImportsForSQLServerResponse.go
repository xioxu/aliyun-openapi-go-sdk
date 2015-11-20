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

package rds_region_DescribeImportsForSQLServerResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type DescribeImportsForSQLServerResponse struct {
	Items             []SQLServerImport
	TotalRecordCounts int32
	PageNumber        int32
	SQLItemsCounts    int32
	RequestId         string
}

func (m *DescribeImportsForSQLServerResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.Items = make([]SQLServerImport, 0)
	for i := 0; i < (flatObj[parentKey+".Items.length"]).(int); i++ {
		tmp := SQLServerImport{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".Items[%d]", i), flatObj)
		m.Items = append(m.Items, tmp)
	}

	m.TotalRecordCounts = builder.GetInt32(parentKey+".TotalRecordCounts", flatObj)

	m.PageNumber = builder.GetInt32(parentKey+".PageNumber", flatObj)

	m.SQLItemsCounts = builder.GetInt32(parentKey+".SQLItemsCounts", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

}
