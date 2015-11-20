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

package inner_DescribeBusinessStatusResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type DescribeBusinessStatusResponse struct {
	UserBusinessStatusList []UserBusinessStatus
	Success                bool
	RequestId              string
}

func (m *DescribeBusinessStatusResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.UserBusinessStatusList = make([]UserBusinessStatus, 0)
	for i := 0; i < (flatObj[parentKey+".UserBusinessStatusList.length"]).(int); i++ {
		tmp := UserBusinessStatus{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".UserBusinessStatusList[%d]", i), flatObj)
		m.UserBusinessStatusList = append(m.UserBusinessStatusList, tmp)
	}

	m.Success = builder.GetBool(parentKey+".Success", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

}
