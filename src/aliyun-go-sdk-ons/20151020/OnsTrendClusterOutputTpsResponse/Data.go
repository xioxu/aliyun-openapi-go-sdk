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

package ons_OnsTrendClusterOutputTpsResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type Data struct {
	Records []StatsDataDo
	Title   string
	XUnit   string
	YUnit   string
}

func (m *Data) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.Records = make([]StatsDataDo, 0)
	for i := 0; i < (flatObj[parentKey+".Records.length"]).(int); i++ {
		tmp := StatsDataDo{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".Records[%d]", i), flatObj)
		m.Records = append(m.Records, tmp)
	}

	m.Title = builder.GetString(parentKey+".Title", flatObj)

	m.XUnit = builder.GetString(parentKey+".XUnit", flatObj)

	m.YUnit = builder.GetString(parentKey+".YUnit", flatObj)

}
