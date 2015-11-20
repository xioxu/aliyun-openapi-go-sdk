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

package rds_DescribeDBInstanceNetInfoForChannelResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type DBInstanceNetInfo struct {
	SecurityIPGroups []SecurityIPGroups
	ConnectionString string
	IPAddress        string
	IPType           string
	Port             string
	VPCId            string
	VSwitchId        string
}

func (m *DBInstanceNetInfo) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.SecurityIPGroups = make([]SecurityIPGroups, 0)
	for i := 0; i < (flatObj[parentKey+".SecurityIPGroups.length"]).(int); i++ {
		tmp := SecurityIPGroups{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".SecurityIPGroups[%d]", i), flatObj)
		m.SecurityIPGroups = append(m.SecurityIPGroups, tmp)
	}

	m.ConnectionString = builder.GetString(parentKey+".ConnectionString", flatObj)

	m.IPAddress = builder.GetString(parentKey+".IPAddress", flatObj)

	m.IPType = builder.GetString(parentKey+".IPType", flatObj)

	m.Port = builder.GetString(parentKey+".Port", flatObj)

	m.VPCId = builder.GetString(parentKey+".VPCId", flatObj)

	m.VSwitchId = builder.GetString(parentKey+".VSwitchId", flatObj)

}
