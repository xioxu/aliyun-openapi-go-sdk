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

package ons_OnsRegionListResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type RegionDo struct {
	Id          int64
	RegionId    string
	RegionName  string
	ChannelId   int32
	ChannelName string
	Owner       string
	Cluster     string
	Status      int32
	IsShare     int32
	CreateTime  int64
	UpdateTime  int64
}

func (m *RegionDo) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.Id = builder.GetInt64(parentKey+".Id", flatObj)

	m.RegionId = builder.GetString(parentKey+".RegionId", flatObj)

	m.RegionName = builder.GetString(parentKey+".RegionName", flatObj)

	m.ChannelId = builder.GetInt32(parentKey+".ChannelId", flatObj)

	m.ChannelName = builder.GetString(parentKey+".ChannelName", flatObj)

	m.Owner = builder.GetString(parentKey+".Owner", flatObj)

	m.Cluster = builder.GetString(parentKey+".Cluster", flatObj)

	m.Status = builder.GetInt32(parentKey+".Status", flatObj)

	m.IsShare = builder.GetInt32(parentKey+".IsShare", flatObj)

	m.CreateTime = builder.GetInt64(parentKey+".CreateTime", flatObj)

	m.UpdateTime = builder.GetInt64(parentKey+".UpdateTime", flatObj)

}
