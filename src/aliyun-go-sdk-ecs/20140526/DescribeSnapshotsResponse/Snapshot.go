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

package ecs_DescribeSnapshotsResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type Snapshot struct {
	Tags           []Tag
	SnapshotId     string
	SnapshotName   string
	Progress       string
	ProductCode    string
	SourceDiskId   string
	SourceDiskType string
	SourceDiskSize string
	Description    string
	CreationTime   string
	Status         string
	Usage          string
}

func (m *Snapshot) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.Tags = make([]Tag, 0)
	for i := 0; i < (flatObj[parentKey+".Tags.length"]).(int); i++ {
		tmp := Tag{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".Tags[%d]", i), flatObj)
		m.Tags = append(m.Tags, tmp)
	}

	m.SnapshotId = builder.GetString(parentKey+".SnapshotId", flatObj)

	m.SnapshotName = builder.GetString(parentKey+".SnapshotName", flatObj)

	m.Progress = builder.GetString(parentKey+".Progress", flatObj)

	m.ProductCode = builder.GetString(parentKey+".ProductCode", flatObj)

	m.SourceDiskId = builder.GetString(parentKey+".SourceDiskId", flatObj)

	m.SourceDiskType = builder.GetString(parentKey+".SourceDiskType", flatObj)

	m.SourceDiskSize = builder.GetString(parentKey+".SourceDiskSize", flatObj)

	m.Description = builder.GetString(parentKey+".Description", flatObj)

	m.CreationTime = builder.GetString(parentKey+".CreationTime", flatObj)

	m.Status = builder.GetString(parentKey+".Status", flatObj)

	m.Usage = builder.GetString(parentKey+".Usage", flatObj)

}
