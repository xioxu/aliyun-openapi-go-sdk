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

package emr_ListConfigTypeResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type InstanceType struct {
	Classify     string
	State        string
	Type         string
	CpuCore      string
	MemSize      string
	HasCloudDisk bool
	HasLocalDisk bool
	HasLocalSSD  bool
	HasCloudSSD  bool
}

func (m *InstanceType) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.Classify = builder.GetString(parentKey+".Classify", flatObj)

	m.State = builder.GetString(parentKey+".State", flatObj)

	m.Type = builder.GetString(parentKey+".Type", flatObj)

	m.CpuCore = builder.GetString(parentKey+".CpuCore", flatObj)

	m.MemSize = builder.GetString(parentKey+".MemSize", flatObj)

	m.HasCloudDisk = builder.GetBool(parentKey+".HasCloudDisk", flatObj)

	m.HasLocalDisk = builder.GetBool(parentKey+".HasLocalDisk", flatObj)

	m.HasLocalSSD = builder.GetBool(parentKey+".HasLocalSSD", flatObj)

	m.HasCloudSSD = builder.GetBool(parentKey+".HasCloudSSD", flatObj)

}
