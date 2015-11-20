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

package ecs_DescribeAutoSnapshotPolicyResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type AutoSnapshotPolicy struct {
	SystemDiskPolicyEnabled           string
	SystemDiskPolicyTimePeriod        string
	SystemDiskPolicyRetentionDays     string
	SystemDiskPolicyRetentionLastWeek string
	DataDiskPolicyEnabled             string
	DataDiskPolicyTimePeriod          string
	DataDiskPolicyRetentionDays       string
	DataDiskPolicyRetentionLastWeek   string
}

func (m *AutoSnapshotPolicy) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.SystemDiskPolicyEnabled = builder.GetString(parentKey+".SystemDiskPolicyEnabled", flatObj)

	m.SystemDiskPolicyTimePeriod = builder.GetString(parentKey+".SystemDiskPolicyTimePeriod", flatObj)

	m.SystemDiskPolicyRetentionDays = builder.GetString(parentKey+".SystemDiskPolicyRetentionDays", flatObj)

	m.SystemDiskPolicyRetentionLastWeek = builder.GetString(parentKey+".SystemDiskPolicyRetentionLastWeek", flatObj)

	m.DataDiskPolicyEnabled = builder.GetString(parentKey+".DataDiskPolicyEnabled", flatObj)

	m.DataDiskPolicyTimePeriod = builder.GetString(parentKey+".DataDiskPolicyTimePeriod", flatObj)

	m.DataDiskPolicyRetentionDays = builder.GetString(parentKey+".DataDiskPolicyRetentionDays", flatObj)

	m.DataDiskPolicyRetentionLastWeek = builder.GetString(parentKey+".DataDiskPolicyRetentionLastWeek", flatObj)

}
