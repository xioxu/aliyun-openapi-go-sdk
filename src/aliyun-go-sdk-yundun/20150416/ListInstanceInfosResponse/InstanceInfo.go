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

package yundun_ListInstanceInfosResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type InstanceInfo struct {
	Region       string
	RegionName   string
	RegionEnName string
	InstanceName string
	InstanceId   string
	Ip           string
	InternetIp   string
	IntranetIp   string
	Ddos         int32
	HostEvent    int32
	SecureCheck  int32
	AegisStatus  int32
	Waf          int32
	IsLock       bool
	LockType     string
	UnLockTimes  int32
	TriggerTime  string
}

func (m *InstanceInfo) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.Region = builder.GetString(parentKey+".Region", flatObj)

	m.RegionName = builder.GetString(parentKey+".RegionName", flatObj)

	m.RegionEnName = builder.GetString(parentKey+".RegionEnName", flatObj)

	m.InstanceName = builder.GetString(parentKey+".InstanceName", flatObj)

	m.InstanceId = builder.GetString(parentKey+".InstanceId", flatObj)

	m.Ip = builder.GetString(parentKey+".Ip", flatObj)

	m.InternetIp = builder.GetString(parentKey+".InternetIp", flatObj)

	m.IntranetIp = builder.GetString(parentKey+".IntranetIp", flatObj)

	m.Ddos = builder.GetInt32(parentKey+".Ddos", flatObj)

	m.HostEvent = builder.GetInt32(parentKey+".HostEvent", flatObj)

	m.SecureCheck = builder.GetInt32(parentKey+".SecureCheck", flatObj)

	m.AegisStatus = builder.GetInt32(parentKey+".AegisStatus", flatObj)

	m.Waf = builder.GetInt32(parentKey+".Waf", flatObj)

	m.IsLock = builder.GetBool(parentKey+".IsLock", flatObj)

	m.LockType = builder.GetString(parentKey+".LockType", flatObj)

	m.UnLockTimes = builder.GetInt32(parentKey+".UnLockTimes", flatObj)

	m.TriggerTime = builder.GetString(parentKey+".TriggerTime", flatObj)

}
