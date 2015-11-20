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

package ram_UpdateGroupResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type Group struct {
	GroupName  string
	Comments   string
	CreateDate string
	UpdateDate string
}

func (m *Group) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.GroupName = builder.GetString(parentKey+".GroupName", flatObj)

	m.Comments = builder.GetString(parentKey+".Comments", flatObj)

	m.CreateDate = builder.GetString(parentKey+".CreateDate", flatObj)

	m.UpdateDate = builder.GetString(parentKey+".UpdateDate", flatObj)

}
