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

package rds_DescribeDBInstanceAttributeResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type DBInstanceAttribute struct {
	ReadOnlyDBInstanceIds       []ReadOnlyDBInstanceId
	InsId                       int32
	DBInstanceId                string
	PayType                     string
	DBInstanceClassType         string
	DBInstanceType              string
	RegionId                    string
	ConnectionString            string
	Port                        string
	Engine                      string
	EngineVersion               string
	DBInstanceClass             string
	DBInstanceMemory            int64
	DBInstanceStorage           int32
	DBInstanceNetType           string
	DBInstanceStatus            string
	DBInstanceDescription       string
	LockMode                    string
	LockReason                  string
	ReadDelayTime               string
	DBMaxQuantity               int32
	AccountMaxQuantity          int32
	CreationTime                string
	ExpireTime                  string
	MaintainTime                string
	AvailabilityValue           string
	MaxIOPS                     int32
	MaxConnections              int32
	MasterInstanceId            string
	IncrementSourceDBInstanceId string
	GuardDBInstanceId           string
	TempDBInstanceId            string
	SecurityIPList              string
	ZoneId                      string
	InstanceNetworkType         string
	VpcId                       string
	ConnectionMode              string
}

func (m *DBInstanceAttribute) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.ReadOnlyDBInstanceIds = make([]ReadOnlyDBInstanceId, 0)
	for i := 0; i < (flatObj[parentKey+".ReadOnlyDBInstanceIds.length"]).(int); i++ {
		tmp := ReadOnlyDBInstanceId{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".ReadOnlyDBInstanceIds[%d]", i), flatObj)
		m.ReadOnlyDBInstanceIds = append(m.ReadOnlyDBInstanceIds, tmp)
	}

	m.InsId = builder.GetInt32(parentKey+".InsId", flatObj)

	m.DBInstanceId = builder.GetString(parentKey+".DBInstanceId", flatObj)

	m.PayType = builder.GetString(parentKey+".PayType", flatObj)

	m.DBInstanceClassType = builder.GetString(parentKey+".DBInstanceClassType", flatObj)

	m.DBInstanceType = builder.GetString(parentKey+".DBInstanceType", flatObj)

	m.RegionId = builder.GetString(parentKey+".RegionId", flatObj)

	m.ConnectionString = builder.GetString(parentKey+".ConnectionString", flatObj)

	m.Port = builder.GetString(parentKey+".Port", flatObj)

	m.Engine = builder.GetString(parentKey+".Engine", flatObj)

	m.EngineVersion = builder.GetString(parentKey+".EngineVersion", flatObj)

	m.DBInstanceClass = builder.GetString(parentKey+".DBInstanceClass", flatObj)

	m.DBInstanceMemory = builder.GetInt64(parentKey+".DBInstanceMemory", flatObj)

	m.DBInstanceStorage = builder.GetInt32(parentKey+".DBInstanceStorage", flatObj)

	m.DBInstanceNetType = builder.GetString(parentKey+".DBInstanceNetType", flatObj)

	m.DBInstanceStatus = builder.GetString(parentKey+".DBInstanceStatus", flatObj)

	m.DBInstanceDescription = builder.GetString(parentKey+".DBInstanceDescription", flatObj)

	m.LockMode = builder.GetString(parentKey+".LockMode", flatObj)

	m.LockReason = builder.GetString(parentKey+".LockReason", flatObj)

	m.ReadDelayTime = builder.GetString(parentKey+".ReadDelayTime", flatObj)

	m.DBMaxQuantity = builder.GetInt32(parentKey+".DBMaxQuantity", flatObj)

	m.AccountMaxQuantity = builder.GetInt32(parentKey+".AccountMaxQuantity", flatObj)

	m.CreationTime = builder.GetString(parentKey+".CreationTime", flatObj)

	m.ExpireTime = builder.GetString(parentKey+".ExpireTime", flatObj)

	m.MaintainTime = builder.GetString(parentKey+".MaintainTime", flatObj)

	m.AvailabilityValue = builder.GetString(parentKey+".AvailabilityValue", flatObj)

	m.MaxIOPS = builder.GetInt32(parentKey+".MaxIOPS", flatObj)

	m.MaxConnections = builder.GetInt32(parentKey+".MaxConnections", flatObj)

	m.MasterInstanceId = builder.GetString(parentKey+".MasterInstanceId", flatObj)

	m.IncrementSourceDBInstanceId = builder.GetString(parentKey+".IncrementSourceDBInstanceId", flatObj)

	m.GuardDBInstanceId = builder.GetString(parentKey+".GuardDBInstanceId", flatObj)

	m.TempDBInstanceId = builder.GetString(parentKey+".TempDBInstanceId", flatObj)

	m.SecurityIPList = builder.GetString(parentKey+".SecurityIPList", flatObj)

	m.ZoneId = builder.GetString(parentKey+".ZoneId", flatObj)

	m.InstanceNetworkType = builder.GetString(parentKey+".InstanceNetworkType", flatObj)

	m.VpcId = builder.GetString(parentKey+".VpcId", flatObj)

	m.ConnectionMode = builder.GetString(parentKey+".ConnectionMode", flatObj)

}
