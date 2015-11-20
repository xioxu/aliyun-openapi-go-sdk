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

package ecs_DescribeDisksResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type Disk struct {
	OperationLocks     []OperationLock
	Tags               []Tag
	DiskId             string
	RegionId           string
	ZoneId             string
	DiskName           string
	Description        string
	Type               string
	Category           string
	Size               int32
	ImageId            string
	SourceSnapshotId   string
	ProductCode        string
	Portable           bool
	Status             string
	InstanceId         string
	Device             string
	DeleteWithInstance bool
	DeleteAutoSnapshot bool
	EnableAutoSnapshot bool
	CreationTime       string
	AttachedTime       string
	DetachedTime       string
	DiskChargeType     string
	ExpiredTime        string
}

func (m *Disk) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.OperationLocks = make([]OperationLock, 0)
	for i := 0; i < (flatObj[parentKey+".OperationLocks.length"]).(int); i++ {
		tmp := OperationLock{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".OperationLocks[%d]", i), flatObj)
		m.OperationLocks = append(m.OperationLocks, tmp)
	}

	m.Tags = make([]Tag, 0)
	for i := 0; i < (flatObj[parentKey+".Tags.length"]).(int); i++ {
		tmp := Tag{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".Tags[%d]", i), flatObj)
		m.Tags = append(m.Tags, tmp)
	}

	m.DiskId = builder.GetString(parentKey+".DiskId", flatObj)

	m.RegionId = builder.GetString(parentKey+".RegionId", flatObj)

	m.ZoneId = builder.GetString(parentKey+".ZoneId", flatObj)

	m.DiskName = builder.GetString(parentKey+".DiskName", flatObj)

	m.Description = builder.GetString(parentKey+".Description", flatObj)

	m.Type = builder.GetString(parentKey+".Type", flatObj)

	m.Category = builder.GetString(parentKey+".Category", flatObj)

	m.Size = builder.GetInt32(parentKey+".Size", flatObj)

	m.ImageId = builder.GetString(parentKey+".ImageId", flatObj)

	m.SourceSnapshotId = builder.GetString(parentKey+".SourceSnapshotId", flatObj)

	m.ProductCode = builder.GetString(parentKey+".ProductCode", flatObj)

	m.Portable = builder.GetBool(parentKey+".Portable", flatObj)

	m.Status = builder.GetString(parentKey+".Status", flatObj)

	m.InstanceId = builder.GetString(parentKey+".InstanceId", flatObj)

	m.Device = builder.GetString(parentKey+".Device", flatObj)

	m.DeleteWithInstance = builder.GetBool(parentKey+".DeleteWithInstance", flatObj)

	m.DeleteAutoSnapshot = builder.GetBool(parentKey+".DeleteAutoSnapshot", flatObj)

	m.EnableAutoSnapshot = builder.GetBool(parentKey+".EnableAutoSnapshot", flatObj)

	m.CreationTime = builder.GetString(parentKey+".CreationTime", flatObj)

	m.AttachedTime = builder.GetString(parentKey+".AttachedTime", flatObj)

	m.DetachedTime = builder.GetString(parentKey+".DetachedTime", flatObj)

	m.DiskChargeType = builder.GetString(parentKey+".DiskChargeType", flatObj)

	m.ExpiredTime = builder.GetString(parentKey+".ExpiredTime", flatObj)

}
