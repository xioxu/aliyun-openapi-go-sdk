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

package yundun_DdosFlowGraphResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type DdosFlowGraphResponse struct {
	NormalFlows []NormalFlow
	TotalFlows  []TotalFlow
	RequestId   string
}

func (m *DdosFlowGraphResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.NormalFlows = make([]NormalFlow, 0)
	for i := 0; i < (flatObj[parentKey+".NormalFlows.length"]).(int); i++ {
		tmp := NormalFlow{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".NormalFlows[%d]", i), flatObj)
		m.NormalFlows = append(m.NormalFlows, tmp)
	}

	m.TotalFlows = make([]TotalFlow, 0)
	for i := 0; i < (flatObj[parentKey+".TotalFlows.length"]).(int); i++ {
		tmp := TotalFlow{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".TotalFlows[%d]", i), flatObj)
		m.TotalFlows = append(m.TotalFlows, tmp)
	}

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

}
