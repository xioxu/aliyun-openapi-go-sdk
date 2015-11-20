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

package rds_DescribeAbnormalDBInstancesResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type InstanceResult struct {
	AbnormalItems         []AbnormalItem
	DBInstanceDescription string
	DBInstanceId          string
}

func (m *InstanceResult) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.AbnormalItems = make([]AbnormalItem, 0)
	for i := 0; i < (flatObj[parentKey+".AbnormalItems.length"]).(int); i++ {
		tmp := AbnormalItem{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".AbnormalItems[%d]", i), flatObj)
		m.AbnormalItems = append(m.AbnormalItems, tmp)
	}

	m.DBInstanceDescription = builder.GetString(parentKey+".DBInstanceDescription", flatObj)

	m.DBInstanceId = builder.GetString(parentKey+".DBInstanceId", flatObj)

}
