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

package bss_DescribeCouponDetailResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type DescribeCouponDetailResponse struct {
	CouponTemplateId int64
	TotalAmount      string
	BalanceAmount    string
	FrozenAmount     string
	ExpiredAmount    string
	DeliveryTime     string
	ExpiredTime      string
	CouponNumber     string
	Status           string
	Description      string
	CreationTime     string
	ModificationTime string
	PriceLimit       string
	Application      string
	RequestId        string
	ProductCodes     []string
	TradeTypes       []string
}

func (m *DescribeCouponDetailResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.CouponTemplateId = builder.GetInt64(parentKey+".CouponTemplateId", flatObj)

	m.TotalAmount = builder.GetString(parentKey+".TotalAmount", flatObj)

	m.BalanceAmount = builder.GetString(parentKey+".BalanceAmount", flatObj)

	m.FrozenAmount = builder.GetString(parentKey+".FrozenAmount", flatObj)

	m.ExpiredAmount = builder.GetString(parentKey+".ExpiredAmount", flatObj)

	m.DeliveryTime = builder.GetString(parentKey+".DeliveryTime", flatObj)

	m.ExpiredTime = builder.GetString(parentKey+".ExpiredTime", flatObj)

	m.CouponNumber = builder.GetString(parentKey+".CouponNumber", flatObj)

	m.Status = builder.GetString(parentKey+".Status", flatObj)

	m.Description = builder.GetString(parentKey+".Description", flatObj)

	m.CreationTime = builder.GetString(parentKey+".CreationTime", flatObj)

	m.ModificationTime = builder.GetString(parentKey+".ModificationTime", flatObj)

	m.PriceLimit = builder.GetString(parentKey+".PriceLimit", flatObj)

	m.Application = builder.GetString(parentKey+".Application", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

	m.ProductCodes = make([]string, 0)
	for i := 0; i < (flatObj[parentKey+".ProductCodes.length"]).(int); i++ {
		m.ProductCodes = append(m.ProductCodes, builder.GetString(fmt.Sprintf("%s[%d]", parentKey+".ProductCodes", i), flatObj))
	}

	m.TradeTypes = make([]string, 0)
	for i := 0; i < (flatObj[parentKey+".TradeTypes.length"]).(int); i++ {
		m.TradeTypes = append(m.TradeTypes, builder.GetString(fmt.Sprintf("%s[%d]", parentKey+".TradeTypes", i), flatObj))
	}

}
