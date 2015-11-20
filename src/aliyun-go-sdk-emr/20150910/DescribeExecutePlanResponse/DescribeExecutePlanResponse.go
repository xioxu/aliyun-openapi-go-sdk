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

package emr_DescribeExecutePlanResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type DescribeExecutePlanResponse struct {
	JobInfos        []JobInfo
	Id              int64
	Name            string
	Strategy        int32
	TimeInterval    int32
	StartTime       string
	TimeUnit        string
	IsCreateCluster bool
	RequestId       string
	ClusterInfo     ClusterInfo
	SoftwareInfo    SoftwareInfo
	EcsInfo         EcsInfo
}

func (m *DescribeExecutePlanResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.JobInfos = make([]JobInfo, 0)
	for i := 0; i < (flatObj[parentKey+".JobInfos.length"]).(int); i++ {
		tmp := JobInfo{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".JobInfos[%d]", i), flatObj)
		m.JobInfos = append(m.JobInfos, tmp)
	}

	m.Id = builder.GetInt64(parentKey+".Id", flatObj)

	m.Name = builder.GetString(parentKey+".Name", flatObj)

	m.Strategy = builder.GetInt32(parentKey+".Strategy", flatObj)

	m.TimeInterval = builder.GetInt32(parentKey+".TimeInterval", flatObj)

	m.StartTime = builder.GetString(parentKey+".StartTime", flatObj)

	m.TimeUnit = builder.GetString(parentKey+".TimeUnit", flatObj)

	m.IsCreateCluster = builder.GetBool(parentKey+".IsCreateCluster", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

	m.ClusterInfo = ClusterInfo{}
	m.ClusterInfo.BuildProperties(parentKey+".ClusterInfo", flatObj)

	m.SoftwareInfo = SoftwareInfo{}
	m.SoftwareInfo.BuildProperties(parentKey+".SoftwareInfo", flatObj)

	m.EcsInfo = EcsInfo{}
	m.EcsInfo.BuildProperties(parentKey+".EcsInfo", flatObj)

}
