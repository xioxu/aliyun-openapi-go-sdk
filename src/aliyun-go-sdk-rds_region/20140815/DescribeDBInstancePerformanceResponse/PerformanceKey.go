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

package rds_region_DescribeDBInstancePerformanceResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type PerformanceKey struct {
	Values      []PerformanceValue
	Key         string
	Unit        string
	ValueFormat string
}

func (m *PerformanceKey) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.Values = make([]PerformanceValue, 0)
	for i := 0; i < (flatObj[parentKey+".Values.length"]).(int); i++ {
		tmp := PerformanceValue{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".Values[%d]", i), flatObj)
		m.Values = append(m.Values, tmp)
	}

	m.Key = builder.GetString(parentKey+".Key", flatObj)

	m.Unit = builder.GetString(parentKey+".Unit", flatObj)

	m.ValueFormat = builder.GetString(parentKey+".ValueFormat", flatObj)

}
