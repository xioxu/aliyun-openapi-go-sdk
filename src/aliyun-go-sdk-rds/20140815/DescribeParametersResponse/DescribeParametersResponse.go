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

package rds_DescribeParametersResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type DescribeParametersResponse struct {
	ConfigParameters  []DBInstanceParameter
	RunningParameters []DBInstanceParameter
	Engine            string
	EngineVersion     string
	RequestId         string
}

func (m *DescribeParametersResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.ConfigParameters = make([]DBInstanceParameter, 0)
	for i := 0; i < (flatObj[parentKey+".ConfigParameters.length"]).(int); i++ {
		tmp := DBInstanceParameter{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".ConfigParameters[%d]", i), flatObj)
		m.ConfigParameters = append(m.ConfigParameters, tmp)
	}

	m.RunningParameters = make([]DBInstanceParameter, 0)
	for i := 0; i < (flatObj[parentKey+".RunningParameters.length"]).(int); i++ {
		tmp := DBInstanceParameter{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".RunningParameters[%d]", i), flatObj)
		m.RunningParameters = append(m.RunningParameters, tmp)
	}

	m.Engine = builder.GetString(parentKey+".Engine", flatObj)

	m.EngineVersion = builder.GetString(parentKey+".EngineVersion", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

}
