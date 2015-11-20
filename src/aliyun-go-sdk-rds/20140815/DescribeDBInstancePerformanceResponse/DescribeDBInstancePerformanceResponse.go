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

package rds_DescribeDBInstancePerformanceResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type DescribeDBInstancePerformanceResponse struct {
	PerformanceKeys []PerformanceKey
	DBInstanceId    string
	Engine          string
	StartTime       string
	EndTime         string
	RequestId       string
}

func (m *DescribeDBInstancePerformanceResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.PerformanceKeys = make([]PerformanceKey, 0)
	for i := 0; i < (flatObj[parentKey+".PerformanceKeys.length"]).(int); i++ {
		tmp := PerformanceKey{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".PerformanceKeys[%d]", i), flatObj)
		m.PerformanceKeys = append(m.PerformanceKeys, tmp)
	}

	m.DBInstanceId = builder.GetString(parentKey+".DBInstanceId", flatObj)

	m.Engine = builder.GetString(parentKey+".Engine", flatObj)

	m.StartTime = builder.GetString(parentKey+".StartTime", flatObj)

	m.EndTime = builder.GetString(parentKey+".EndTime", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

}