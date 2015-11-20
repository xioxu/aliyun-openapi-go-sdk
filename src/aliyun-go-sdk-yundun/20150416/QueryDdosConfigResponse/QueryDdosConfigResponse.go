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

package yundun_QueryDdosConfigResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type QueryDdosConfigResponse struct {
	Bps              int64
	Pps              int64
	Qps              int64
	Sipconn          int64
	Sipnew           int64
	Layer7Config     bool
	FlowPosition     int32
	QpsPosition      int32
	StrategyPosition int32
	Level            int32
	HoleBps          string
	ConfigType       string
	RequestId        string
}

func (m *QueryDdosConfigResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.Bps = builder.GetInt64(parentKey+".Bps", flatObj)

	m.Pps = builder.GetInt64(parentKey+".Pps", flatObj)

	m.Qps = builder.GetInt64(parentKey+".Qps", flatObj)

	m.Sipconn = builder.GetInt64(parentKey+".Sipconn", flatObj)

	m.Sipnew = builder.GetInt64(parentKey+".Sipnew", flatObj)

	m.Layer7Config = builder.GetBool(parentKey+".Layer7Config", flatObj)

	m.FlowPosition = builder.GetInt32(parentKey+".FlowPosition", flatObj)

	m.QpsPosition = builder.GetInt32(parentKey+".QpsPosition", flatObj)

	m.StrategyPosition = builder.GetInt32(parentKey+".StrategyPosition", flatObj)

	m.Level = builder.GetInt32(parentKey+".Level", flatObj)

	m.HoleBps = builder.GetString(parentKey+".HoleBps", flatObj)

	m.ConfigType = builder.GetString(parentKey+".ConfigType", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

}
