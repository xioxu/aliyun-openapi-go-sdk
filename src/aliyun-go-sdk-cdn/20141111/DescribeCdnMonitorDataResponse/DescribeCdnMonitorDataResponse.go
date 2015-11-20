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

package cdn_DescribeCdnMonitorDataResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type DescribeCdnMonitorDataResponse struct {
	MonitorDatas    []CDNMonitorData
	DomainName      string
	MonitorInterval int64
	StartTime       string
	EndTime         string
	RequestId       string
}

func (m *DescribeCdnMonitorDataResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.MonitorDatas = make([]CDNMonitorData, 0)
	for i := 0; i < (flatObj[parentKey+".MonitorDatas.length"]).(int); i++ {
		tmp := CDNMonitorData{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".MonitorDatas[%d]", i), flatObj)
		m.MonitorDatas = append(m.MonitorDatas, tmp)
	}

	m.DomainName = builder.GetString(parentKey+".DomainName", flatObj)

	m.MonitorInterval = builder.GetInt64(parentKey+".MonitorInterval", flatObj)

	m.StartTime = builder.GetString(parentKey+".StartTime", flatObj)

	m.EndTime = builder.GetString(parentKey+".EndTime", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

}
