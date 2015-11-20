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

package ram_ListUserPoliciesResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type ListUserPoliciesResponse struct {
	RequestId string
	Policies  []string
}

func (m *ListUserPoliciesResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

	m.Policies = make([]string, 0)
	for i := 0; i < (flatObj[parentKey+".Policies.length"]).(int); i++ {
		m.Policies = append(m.Policies, builder.GetString(fmt.Sprintf("%s[%d]", parentKey+".Policies", i), flatObj))
	}

}
