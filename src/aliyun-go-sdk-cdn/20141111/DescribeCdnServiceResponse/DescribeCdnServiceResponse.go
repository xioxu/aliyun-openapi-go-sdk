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

package cdn_DescribeCdnServiceResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type DescribeCdnServiceResponse struct {
	OperationLocks     []LockReason
	InternetChargeType string
	OpeningTime        string
	ChangingChargeType string
	ChangingAffectTime string
	RequestId          string
}

func (m *DescribeCdnServiceResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.OperationLocks = make([]LockReason, 0)
	for i := 0; i < (flatObj[parentKey+".OperationLocks.length"]).(int); i++ {
		tmp := LockReason{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".OperationLocks[%d]", i), flatObj)
		m.OperationLocks = append(m.OperationLocks, tmp)
	}

	m.InternetChargeType = builder.GetString(parentKey+".InternetChargeType", flatObj)

	m.OpeningTime = builder.GetString(parentKey+".OpeningTime", flatObj)

	m.ChangingChargeType = builder.GetString(parentKey+".ChangingChargeType", flatObj)

	m.ChangingAffectTime = builder.GetString(parentKey+".ChangingAffectTime", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

}
