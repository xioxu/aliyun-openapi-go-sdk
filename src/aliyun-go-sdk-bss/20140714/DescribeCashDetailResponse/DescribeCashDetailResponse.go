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

package bss_DescribeCashDetailResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type DescribeCashDetailResponse struct {
	BalanceAmount        string
	AmountOwed           string
	EnableThresholdAlert string
	MiniAlertThreshold   int64
	FrozenAmount         string
	CreditCardAmount     string
	RemmitanceAmount     string
	CreditLimit          string
	AvailableCredit      string
	RequestId            string
}

func (m *DescribeCashDetailResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.BalanceAmount = builder.GetString(parentKey+".BalanceAmount", flatObj)

	m.AmountOwed = builder.GetString(parentKey+".AmountOwed", flatObj)

	m.EnableThresholdAlert = builder.GetString(parentKey+".EnableThresholdAlert", flatObj)

	m.MiniAlertThreshold = builder.GetInt64(parentKey+".MiniAlertThreshold", flatObj)

	m.FrozenAmount = builder.GetString(parentKey+".FrozenAmount", flatObj)

	m.CreditCardAmount = builder.GetString(parentKey+".CreditCardAmount", flatObj)

	m.RemmitanceAmount = builder.GetString(parentKey+".RemmitanceAmount", flatObj)

	m.CreditLimit = builder.GetString(parentKey+".CreditLimit", flatObj)

	m.AvailableCredit = builder.GetString(parentKey+".AvailableCredit", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

}
