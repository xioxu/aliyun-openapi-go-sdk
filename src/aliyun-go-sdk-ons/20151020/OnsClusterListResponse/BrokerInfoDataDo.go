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

package ons_OnsClusterListResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type BrokerInfoDataDo struct {
	ClusterName   string
	BrokerName    string
	BrokerId      int32
	BrokerAddr    string
	BrokerIp      string
	Version       string
	InTPS         float32
	OutTPS        float32
	InTotalYest   float32
	OutTotalYest  float32
	InTotalToday  float32
	OutTotalToday float32
}

func (m *BrokerInfoDataDo) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.ClusterName = builder.GetString(parentKey+".ClusterName", flatObj)

	m.BrokerName = builder.GetString(parentKey+".BrokerName", flatObj)

	m.BrokerId = builder.GetInt32(parentKey+".BrokerId", flatObj)

	m.BrokerAddr = builder.GetString(parentKey+".BrokerAddr", flatObj)

	m.BrokerIp = builder.GetString(parentKey+".BrokerIp", flatObj)

	m.Version = builder.GetString(parentKey+".Version", flatObj)

	m.InTPS = builder.GetFloat32(parentKey+".InTPS", flatObj)

	m.OutTPS = builder.GetFloat32(parentKey+".OutTPS", flatObj)

	m.InTotalYest = builder.GetFloat32(parentKey+".InTotalYest", flatObj)

	m.OutTotalYest = builder.GetFloat32(parentKey+".OutTotalYest", flatObj)

	m.InTotalToday = builder.GetFloat32(parentKey+".InTotalToday", flatObj)

	m.OutTotalToday = builder.GetFloat32(parentKey+".OutTotalToday", flatObj)

}
