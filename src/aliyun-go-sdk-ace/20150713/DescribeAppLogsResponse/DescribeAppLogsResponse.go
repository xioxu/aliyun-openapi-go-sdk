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

package ace_DescribeAppLogsResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type DescribeAppLogsResponse struct {
	Items           []AppLog
	PageRecordCount int32
	RequestId       string
}

func (m *DescribeAppLogsResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.Items = make([]AppLog, 0)
	for i := 0; i < (flatObj[parentKey+".Items.length"]).(int); i++ {
		tmp := AppLog{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".Items[%d]", i), flatObj)
		m.Items = append(m.Items, tmp)
	}

	m.PageRecordCount = builder.GetInt32(parentKey+".PageRecordCount", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

}
