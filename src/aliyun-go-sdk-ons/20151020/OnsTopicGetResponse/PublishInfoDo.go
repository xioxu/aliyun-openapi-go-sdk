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

package ons_OnsTopicGetResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type PublishInfoDo struct {
	Id           int64
	ChannelId    int32
	ChannelName  string
	RegionId     string
	RegionName   string
	Topic        string
	Owner        string
	Relation     int32
	RelationName string
	Status       int32
	StatusName   string
	Appkey       int32
	CreateTime   int64
	UpdateTime   int64
	Remark       string
}

func (m *PublishInfoDo) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.Id = builder.GetInt64(parentKey+".Id", flatObj)

	m.ChannelId = builder.GetInt32(parentKey+".ChannelId", flatObj)

	m.ChannelName = builder.GetString(parentKey+".ChannelName", flatObj)

	m.RegionId = builder.GetString(parentKey+".RegionId", flatObj)

	m.RegionName = builder.GetString(parentKey+".RegionName", flatObj)

	m.Topic = builder.GetString(parentKey+".Topic", flatObj)

	m.Owner = builder.GetString(parentKey+".Owner", flatObj)

	m.Relation = builder.GetInt32(parentKey+".Relation", flatObj)

	m.RelationName = builder.GetString(parentKey+".RelationName", flatObj)

	m.Status = builder.GetInt32(parentKey+".Status", flatObj)

	m.StatusName = builder.GetString(parentKey+".StatusName", flatObj)

	m.Appkey = builder.GetInt32(parentKey+".Appkey", flatObj)

	m.CreateTime = builder.GetInt64(parentKey+".CreateTime", flatObj)

	m.UpdateTime = builder.GetInt64(parentKey+".UpdateTime", flatObj)

	m.Remark = builder.GetString(parentKey+".Remark", flatObj)

}
