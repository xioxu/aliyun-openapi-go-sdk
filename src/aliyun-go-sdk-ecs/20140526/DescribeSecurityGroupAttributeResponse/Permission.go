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

package ecs_DescribeSecurityGroupAttributeResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type Permission struct {
	IpProtocol              string
	PortRange               string
	SourceGroupId           string
	SourceCidrIp            string
	Policy                  string
	NicType                 string
	SourceGroupOwnerAccount string
	DestGroupId             string
	DestCidrIp              string
	DestGroupOwnerAccount   string
	Priority                string
	Direction               string
}

func (m *Permission) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.IpProtocol = builder.GetString(parentKey+".IpProtocol", flatObj)

	m.PortRange = builder.GetString(parentKey+".PortRange", flatObj)

	m.SourceGroupId = builder.GetString(parentKey+".SourceGroupId", flatObj)

	m.SourceCidrIp = builder.GetString(parentKey+".SourceCidrIp", flatObj)

	m.Policy = builder.GetString(parentKey+".Policy", flatObj)

	m.NicType = builder.GetString(parentKey+".NicType", flatObj)

	m.SourceGroupOwnerAccount = builder.GetString(parentKey+".SourceGroupOwnerAccount", flatObj)

	m.DestGroupId = builder.GetString(parentKey+".DestGroupId", flatObj)

	m.DestCidrIp = builder.GetString(parentKey+".DestCidrIp", flatObj)

	m.DestGroupOwnerAccount = builder.GetString(parentKey+".DestGroupOwnerAccount", flatObj)

	m.Priority = builder.GetString(parentKey+".Priority", flatObj)

	m.Direction = builder.GetString(parentKey+".Direction", flatObj)

}
