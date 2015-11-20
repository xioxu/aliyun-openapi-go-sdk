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

package rds_region_DescribeDBInstancesResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type DBInstance struct {
	ReadOnlyDBInstanceIds []ReadOnlyDBInstanceId
	InsId                 int32
	DBInstanceId          string
	DBInstanceDescription string
	PayType               string
	DBInstanceType        string
	RegionId              string
	ExpireTime            string
	DBInstanceStatus      string
	Engine                string
	DBInstanceNetType     string
	ConnectionMode        string
	LockMode              string
	InstanceNetworkType   string
	LockReason            string
	ZoneId                string
	MutriORsignle         bool
	CreateTime            string
	EngineVersion         string
	GuardDBInstanceId     string
	TempDBInstanceId      string
	MasterInstanceId      string
	VpcId                 string
}

func (m *DBInstance) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.ReadOnlyDBInstanceIds = make([]ReadOnlyDBInstanceId, 0)
	for i := 0; i < (flatObj[parentKey+".ReadOnlyDBInstanceIds.length"]).(int); i++ {
		tmp := ReadOnlyDBInstanceId{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".ReadOnlyDBInstanceIds[%d]", i), flatObj)
		m.ReadOnlyDBInstanceIds = append(m.ReadOnlyDBInstanceIds, tmp)
	}

	m.InsId = builder.GetInt32(parentKey+".InsId", flatObj)

	m.DBInstanceId = builder.GetString(parentKey+".DBInstanceId", flatObj)

	m.DBInstanceDescription = builder.GetString(parentKey+".DBInstanceDescription", flatObj)

	m.PayType = builder.GetString(parentKey+".PayType", flatObj)

	m.DBInstanceType = builder.GetString(parentKey+".DBInstanceType", flatObj)

	m.RegionId = builder.GetString(parentKey+".RegionId", flatObj)

	m.ExpireTime = builder.GetString(parentKey+".ExpireTime", flatObj)

	m.DBInstanceStatus = builder.GetString(parentKey+".DBInstanceStatus", flatObj)

	m.Engine = builder.GetString(parentKey+".Engine", flatObj)

	m.DBInstanceNetType = builder.GetString(parentKey+".DBInstanceNetType", flatObj)

	m.ConnectionMode = builder.GetString(parentKey+".ConnectionMode", flatObj)

	m.LockMode = builder.GetString(parentKey+".LockMode", flatObj)

	m.InstanceNetworkType = builder.GetString(parentKey+".InstanceNetworkType", flatObj)

	m.LockReason = builder.GetString(parentKey+".LockReason", flatObj)

	m.ZoneId = builder.GetString(parentKey+".ZoneId", flatObj)

	m.MutriORsignle = builder.GetBool(parentKey+".MutriORsignle", flatObj)

	m.CreateTime = builder.GetString(parentKey+".CreateTime", flatObj)

	m.EngineVersion = builder.GetString(parentKey+".EngineVersion", flatObj)

	m.GuardDBInstanceId = builder.GetString(parentKey+".GuardDBInstanceId", flatObj)

	m.TempDBInstanceId = builder.GetString(parentKey+".TempDBInstanceId", flatObj)

	m.MasterInstanceId = builder.GetString(parentKey+".MasterInstanceId", flatObj)

	m.VpcId = builder.GetString(parentKey+".VpcId", flatObj)

}
