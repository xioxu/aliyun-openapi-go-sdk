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

package drds_DescribeDrdsInstanceResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type Data struct {
	Vips           []Vip
	DrdsInstanceId string
	Type           string
	RegionId       string
	ZoneId         string
	Description    string
	NetworkType    string
	Status         string
	CreateTime     int64
	Version        int64
}

func (m *Data) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.Vips = make([]Vip, 0)
	for i := 0; i < (flatObj[parentKey+".Vips.length"]).(int); i++ {
		tmp := Vip{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".Vips[%d]", i), flatObj)
		m.Vips = append(m.Vips, tmp)
	}

	m.DrdsInstanceId = builder.GetString(parentKey+".DrdsInstanceId", flatObj)

	m.Type = builder.GetString(parentKey+".Type", flatObj)

	m.RegionId = builder.GetString(parentKey+".RegionId", flatObj)

	m.ZoneId = builder.GetString(parentKey+".ZoneId", flatObj)

	m.Description = builder.GetString(parentKey+".Description", flatObj)

	m.NetworkType = builder.GetString(parentKey+".NetworkType", flatObj)

	m.Status = builder.GetString(parentKey+".Status", flatObj)

	m.CreateTime = builder.GetInt64(parentKey+".CreateTime", flatObj)

	m.Version = builder.GetInt64(parentKey+".Version", flatObj)

}
