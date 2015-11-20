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

package yundun_GetDdosConfigOptionsResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type GetDdosConfigOptionsResponse struct {
	RequestThresholdOptions1   []RequestThresholdOption
	RequestThresholdOptions2   []RequestThresholdOption
	ConnectionThresholdOptions []ConnectionThresholdOption
	RequestId                  string
	QpsOptions1                []string
	QpsOptions2                []string
}

func (m *GetDdosConfigOptionsResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.RequestThresholdOptions1 = make([]RequestThresholdOption, 0)
	for i := 0; i < (flatObj[parentKey+".RequestThresholdOptions1.length"]).(int); i++ {
		tmp := RequestThresholdOption{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".RequestThresholdOptions1[%d]", i), flatObj)
		m.RequestThresholdOptions1 = append(m.RequestThresholdOptions1, tmp)
	}

	m.RequestThresholdOptions2 = make([]RequestThresholdOption, 0)
	for i := 0; i < (flatObj[parentKey+".RequestThresholdOptions2.length"]).(int); i++ {
		tmp := RequestThresholdOption{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".RequestThresholdOptions2[%d]", i), flatObj)
		m.RequestThresholdOptions2 = append(m.RequestThresholdOptions2, tmp)
	}

	m.ConnectionThresholdOptions = make([]ConnectionThresholdOption, 0)
	for i := 0; i < (flatObj[parentKey+".ConnectionThresholdOptions.length"]).(int); i++ {
		tmp := ConnectionThresholdOption{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".ConnectionThresholdOptions[%d]", i), flatObj)
		m.ConnectionThresholdOptions = append(m.ConnectionThresholdOptions, tmp)
	}

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

	m.QpsOptions1 = make([]string, 0)
	for i := 0; i < (flatObj[parentKey+".QpsOptions1.length"]).(int); i++ {
		m.QpsOptions1 = append(m.QpsOptions1, builder.GetString(fmt.Sprintf("%s[%d]", parentKey+".QpsOptions1", i), flatObj))
	}

	m.QpsOptions2 = make([]string, 0)
	for i := 0; i < (flatObj[parentKey+".QpsOptions2.length"]).(int); i++ {
		m.QpsOptions2 = append(m.QpsOptions2, builder.GetString(fmt.Sprintf("%s[%d]", parentKey+".QpsOptions2", i), flatObj))
	}

}
