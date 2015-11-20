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

package rds_DescribeDatabasesResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type Database struct {
	Accounts         []AccountPrivilegeInfo
	DBName           string
	DBInstanceId     string
	Engine           string
	DBStatus         string
	CharacterSetName string
	DBDescription    string
}

func (m *Database) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.Accounts = make([]AccountPrivilegeInfo, 0)
	for i := 0; i < (flatObj[parentKey+".Accounts.length"]).(int); i++ {
		tmp := AccountPrivilegeInfo{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".Accounts[%d]", i), flatObj)
		m.Accounts = append(m.Accounts, tmp)
	}

	m.DBName = builder.GetString(parentKey+".DBName", flatObj)

	m.DBInstanceId = builder.GetString(parentKey+".DBInstanceId", flatObj)

	m.Engine = builder.GetString(parentKey+".Engine", flatObj)

	m.DBStatus = builder.GetString(parentKey+".DBStatus", flatObj)

	m.CharacterSetName = builder.GetString(parentKey+".CharacterSetName", flatObj)

	m.DBDescription = builder.GetString(parentKey+".DBDescription", flatObj)

}
