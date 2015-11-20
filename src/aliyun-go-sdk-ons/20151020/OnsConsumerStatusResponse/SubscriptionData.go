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

package ons_OnsConsumerStatusResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type SubscriptionData struct {
	ClassFilterMode bool
	Topic           string
	SubString       string
	SubVersion      int64
	TagsSet         []string
	CodeSet         []string
}

func (m *SubscriptionData) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.ClassFilterMode = builder.GetBool(parentKey+".ClassFilterMode", flatObj)

	m.Topic = builder.GetString(parentKey+".Topic", flatObj)

	m.SubString = builder.GetString(parentKey+".SubString", flatObj)

	m.SubVersion = builder.GetInt64(parentKey+".SubVersion", flatObj)

	m.TagsSet = make([]string, 0)
	for i := 0; i < (flatObj[parentKey+".TagsSet.length"]).(int); i++ {
		m.TagsSet = append(m.TagsSet, builder.GetString(fmt.Sprintf("%s[%d]", parentKey+".TagsSet", i), flatObj))
	}

	m.CodeSet = make([]string, 0)
	for i := 0; i < (flatObj[parentKey+".CodeSet.length"]).(int); i++ {
		m.CodeSet = append(m.CodeSet, builder.GetString(fmt.Sprintf("%s[%d]", parentKey+".CodeSet", i), flatObj))
	}

}
