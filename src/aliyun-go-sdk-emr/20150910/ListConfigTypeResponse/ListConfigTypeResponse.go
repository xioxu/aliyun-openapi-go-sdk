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
	"fmt"
)

type ListConfigTypeResponse struct {
	securityGroupTypes []SecurityGroupType
	instanceTypes      []InstanceType
	EmrVerTypes        []EmrVerType
	ZoneTypes          []ZoneType
	RequestId          string
}

func (m *ListConfigTypeResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.securityGroupTypes = make([]SecurityGroupType, 0)
	for i := 0; i < (flatObj[parentKey+".securityGroupTypes.length"]).(int); i++ {
		tmp := SecurityGroupType{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".securityGroupTypes[%d]", i), flatObj)
		m.securityGroupTypes = append(m.securityGroupTypes, tmp)
	}

	m.instanceTypes = make([]InstanceType, 0)
	for i := 0; i < (flatObj[parentKey+".instanceTypes.length"]).(int); i++ {
		tmp := InstanceType{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".instanceTypes[%d]", i), flatObj)
		m.instanceTypes = append(m.instanceTypes, tmp)
	}

	m.EmrVerTypes = make([]EmrVerType, 0)
	for i := 0; i < (flatObj[parentKey+".EmrVerTypes.length"]).(int); i++ {
		tmp := EmrVerType{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".EmrVerTypes[%d]", i), flatObj)
		m.EmrVerTypes = append(m.EmrVerTypes, tmp)
	}

	m.ZoneTypes = make([]ZoneType, 0)
	for i := 0; i < (flatObj[parentKey+".ZoneTypes.length"]).(int); i++ {
		tmp := ZoneType{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".ZoneTypes[%d]", i), flatObj)
		m.ZoneTypes = append(m.ZoneTypes, tmp)
	}

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

}
