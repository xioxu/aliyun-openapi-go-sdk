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

package kvstore_DescribePriceResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type Order struct {
	OriginalAmount float32
	TradeAmount    float32
	DiscountAmount float32
	RuleIds        []string
}

func (m *Order) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.OriginalAmount = builder.GetFloat32(parentKey+".OriginalAmount", flatObj)

	m.TradeAmount = builder.GetFloat32(parentKey+".TradeAmount", flatObj)

	m.DiscountAmount = builder.GetFloat32(parentKey+".DiscountAmount", flatObj)

	m.RuleIds = make([]string, 0)
	for i := 0; i < (flatObj[parentKey+".RuleIds.length"]).(int); i++ {
		m.RuleIds = append(m.RuleIds, builder.GetString(fmt.Sprintf("%s[%d]", parentKey+".RuleIds", i), flatObj))
	}

}
