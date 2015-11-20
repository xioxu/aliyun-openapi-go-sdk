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

package emr_DescribeClusterResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type ClusterInfo struct {
	ClusterId         int64
	BizId             string
	ClusterName       string
	UtcStartTime      string
	UtcStopTime       string
	TimeOutEnable     int32
	TimeOutTime       string
	TimeOutWarningWay int32
	TimeOutOperate    int32
	ReleaseSetting    int32
	IsOpenOssLog      bool
	Status            int32
	LastJobStatus     int32
	RunningTime       int32
	FailReason        string
	OssPath           string
	EmrRole4ECS       string
	EmrRole4Oss       string
	IsOpenPublicIp    bool
}

func (m *ClusterInfo) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.ClusterId = builder.GetInt64(parentKey+".ClusterId", flatObj)

	m.BizId = builder.GetString(parentKey+".BizId", flatObj)

	m.ClusterName = builder.GetString(parentKey+".ClusterName", flatObj)

	m.UtcStartTime = builder.GetString(parentKey+".UtcStartTime", flatObj)

	m.UtcStopTime = builder.GetString(parentKey+".UtcStopTime", flatObj)

	m.TimeOutEnable = builder.GetInt32(parentKey+".TimeOutEnable", flatObj)

	m.TimeOutTime = builder.GetString(parentKey+".TimeOutTime", flatObj)

	m.TimeOutWarningWay = builder.GetInt32(parentKey+".TimeOutWarningWay", flatObj)

	m.TimeOutOperate = builder.GetInt32(parentKey+".TimeOutOperate", flatObj)

	m.ReleaseSetting = builder.GetInt32(parentKey+".ReleaseSetting", flatObj)

	m.IsOpenOssLog = builder.GetBool(parentKey+".IsOpenOssLog", flatObj)

	m.Status = builder.GetInt32(parentKey+".Status", flatObj)

	m.LastJobStatus = builder.GetInt32(parentKey+".LastJobStatus", flatObj)

	m.RunningTime = builder.GetInt32(parentKey+".RunningTime", flatObj)

	m.FailReason = builder.GetString(parentKey+".FailReason", flatObj)

	m.OssPath = builder.GetString(parentKey+".OssPath", flatObj)

	m.EmrRole4ECS = builder.GetString(parentKey+".EmrRole4ECS", flatObj)

	m.EmrRole4Oss = builder.GetString(parentKey+".EmrRole4Oss", flatObj)

	m.IsOpenPublicIp = builder.GetBool(parentKey+".IsOpenPublicIp", flatObj)

}
