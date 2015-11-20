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
	"fmt"
)

type EcsRole struct {
	Nodes          []Node
	IsMaster       bool
	InstanceType   string
	Payment        string
	CpuCore        string
	MemoryCapacity string
	DiskType       int32
	DiskCapacity   int32
	BandWidth      string
	NetPayType     string
	EcsPayType     string
}

func (m *EcsRole) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.Nodes = make([]Node, 0)
	for i := 0; i < (flatObj[parentKey+".Nodes.length"]).(int); i++ {
		tmp := Node{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".Nodes[%d]", i), flatObj)
		m.Nodes = append(m.Nodes, tmp)
	}

	m.IsMaster = builder.GetBool(parentKey+".IsMaster", flatObj)

	m.InstanceType = builder.GetString(parentKey+".InstanceType", flatObj)

	m.Payment = builder.GetString(parentKey+".Payment", flatObj)

	m.CpuCore = builder.GetString(parentKey+".CpuCore", flatObj)

	m.MemoryCapacity = builder.GetString(parentKey+".MemoryCapacity", flatObj)

	m.DiskType = builder.GetInt32(parentKey+".DiskType", flatObj)

	m.DiskCapacity = builder.GetInt32(parentKey+".DiskCapacity", flatObj)

	m.BandWidth = builder.GetString(parentKey+".BandWidth", flatObj)

	m.NetPayType = builder.GetString(parentKey+".NetPayType", flatObj)

	m.EcsPayType = builder.GetString(parentKey+".EcsPayType", flatObj)

}
