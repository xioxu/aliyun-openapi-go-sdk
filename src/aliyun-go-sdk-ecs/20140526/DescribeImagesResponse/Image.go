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

package ecs_DescribeImagesResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type Image struct {
	DiskDeviceMappings []DiskDeviceMapping
	Tags               []Tag
	Progress           string
	ImageId            string
	ImageName          string
	ImageVersion       string
	Description        string
	Size               int32
	ImageOwnerAlias    string
	OSName             string
	Architecture       string
	Status             string
	ProductCode        string
	IsSubscribed       bool
	CreationTime       string
	IsSelfShared       string
	OSType             string
	Platform           string
	Usage              string
	IsCopied           bool
}

func (m *Image) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.DiskDeviceMappings = make([]DiskDeviceMapping, 0)
	for i := 0; i < (flatObj[parentKey+".DiskDeviceMappings.length"]).(int); i++ {
		tmp := DiskDeviceMapping{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".DiskDeviceMappings[%d]", i), flatObj)
		m.DiskDeviceMappings = append(m.DiskDeviceMappings, tmp)
	}

	m.Tags = make([]Tag, 0)
	for i := 0; i < (flatObj[parentKey+".Tags.length"]).(int); i++ {
		tmp := Tag{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".Tags[%d]", i), flatObj)
		m.Tags = append(m.Tags, tmp)
	}

	m.Progress = builder.GetString(parentKey+".Progress", flatObj)

	m.ImageId = builder.GetString(parentKey+".ImageId", flatObj)

	m.ImageName = builder.GetString(parentKey+".ImageName", flatObj)

	m.ImageVersion = builder.GetString(parentKey+".ImageVersion", flatObj)

	m.Description = builder.GetString(parentKey+".Description", flatObj)

	m.Size = builder.GetInt32(parentKey+".Size", flatObj)

	m.ImageOwnerAlias = builder.GetString(parentKey+".ImageOwnerAlias", flatObj)

	m.OSName = builder.GetString(parentKey+".OSName", flatObj)

	m.Architecture = builder.GetString(parentKey+".Architecture", flatObj)

	m.Status = builder.GetString(parentKey+".Status", flatObj)

	m.ProductCode = builder.GetString(parentKey+".ProductCode", flatObj)

	m.IsSubscribed = builder.GetBool(parentKey+".IsSubscribed", flatObj)

	m.CreationTime = builder.GetString(parentKey+".CreationTime", flatObj)

	m.IsSelfShared = builder.GetString(parentKey+".IsSelfShared", flatObj)

	m.OSType = builder.GetString(parentKey+".OSType", flatObj)

	m.Platform = builder.GetString(parentKey+".Platform", flatObj)

	m.Usage = builder.GetString(parentKey+".Usage", flatObj)

	m.IsCopied = builder.GetBool(parentKey+".IsCopied", flatObj)

}
