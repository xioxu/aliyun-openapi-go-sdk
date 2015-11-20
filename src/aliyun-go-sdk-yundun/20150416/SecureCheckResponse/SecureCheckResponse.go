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

package yundun_SecureCheckResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type SecureCheckResponse struct {
	ProblemList      []Info
	NoProblemList    []Info
	NoScanList       []Info
	ScanningList     []Info
	InnerIpList      []Info
	RecentInstanceId string
	RequestId        string
}

func (m *SecureCheckResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.ProblemList = make([]Info, 0)
	for i := 0; i < (flatObj[parentKey+".ProblemList.length"]).(int); i++ {
		tmp := Info{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".ProblemList[%d]", i), flatObj)
		m.ProblemList = append(m.ProblemList, tmp)
	}

	m.NoProblemList = make([]Info, 0)
	for i := 0; i < (flatObj[parentKey+".NoProblemList.length"]).(int); i++ {
		tmp := Info{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".NoProblemList[%d]", i), flatObj)
		m.NoProblemList = append(m.NoProblemList, tmp)
	}

	m.NoScanList = make([]Info, 0)
	for i := 0; i < (flatObj[parentKey+".NoScanList.length"]).(int); i++ {
		tmp := Info{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".NoScanList[%d]", i), flatObj)
		m.NoScanList = append(m.NoScanList, tmp)
	}

	m.ScanningList = make([]Info, 0)
	for i := 0; i < (flatObj[parentKey+".ScanningList.length"]).(int); i++ {
		tmp := Info{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".ScanningList[%d]", i), flatObj)
		m.ScanningList = append(m.ScanningList, tmp)
	}

	m.InnerIpList = make([]Info, 0)
	for i := 0; i < (flatObj[parentKey+".InnerIpList.length"]).(int); i++ {
		tmp := Info{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".InnerIpList[%d]", i), flatObj)
		m.InnerIpList = append(m.InnerIpList, tmp)
	}

	m.RecentInstanceId = builder.GetString(parentKey+".RecentInstanceId", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

}
