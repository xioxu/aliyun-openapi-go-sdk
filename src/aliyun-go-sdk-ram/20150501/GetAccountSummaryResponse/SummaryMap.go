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

package ram_GetAccountSummaryResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type SummaryMap struct {
	UsersQuota                    int32
	Users                         int32
	AccessKeysPerUserQuota        int32
	VirtualMFADevicesQuota        int32
	MFADevices                    int32
	MFADevicesInUse               int32
	GroupsQuota                   int32
	Groups                        int32
	GroupsPerUserQuota            int32
	RolesQuota                    int32
	Roles                         int32
	PoliciesQuota                 int32
	Policies                      int32
	PolicySizeQuota               int32
	VersionsPerPolicyQuota        int32
	AttachedPoliciesPerUserQuota  int32
	AttachedPoliciesPerGroupQuota int32
	AttachedPoliciesPerRoleQuota  int32
}

func (m *SummaryMap) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.UsersQuota = builder.GetInt32(parentKey+".UsersQuota", flatObj)

	m.Users = builder.GetInt32(parentKey+".Users", flatObj)

	m.AccessKeysPerUserQuota = builder.GetInt32(parentKey+".AccessKeysPerUserQuota", flatObj)

	m.VirtualMFADevicesQuota = builder.GetInt32(parentKey+".VirtualMFADevicesQuota", flatObj)

	m.MFADevices = builder.GetInt32(parentKey+".MFADevices", flatObj)

	m.MFADevicesInUse = builder.GetInt32(parentKey+".MFADevicesInUse", flatObj)

	m.GroupsQuota = builder.GetInt32(parentKey+".GroupsQuota", flatObj)

	m.Groups = builder.GetInt32(parentKey+".Groups", flatObj)

	m.GroupsPerUserQuota = builder.GetInt32(parentKey+".GroupsPerUserQuota", flatObj)

	m.RolesQuota = builder.GetInt32(parentKey+".RolesQuota", flatObj)

	m.Roles = builder.GetInt32(parentKey+".Roles", flatObj)

	m.PoliciesQuota = builder.GetInt32(parentKey+".PoliciesQuota", flatObj)

	m.Policies = builder.GetInt32(parentKey+".Policies", flatObj)

	m.PolicySizeQuota = builder.GetInt32(parentKey+".PolicySizeQuota", flatObj)

	m.VersionsPerPolicyQuota = builder.GetInt32(parentKey+".VersionsPerPolicyQuota", flatObj)

	m.AttachedPoliciesPerUserQuota = builder.GetInt32(parentKey+".AttachedPoliciesPerUserQuota", flatObj)

	m.AttachedPoliciesPerGroupQuota = builder.GetInt32(parentKey+".AttachedPoliciesPerGroupQuota", flatObj)

	m.AttachedPoliciesPerRoleQuota = builder.GetInt32(parentKey+".AttachedPoliciesPerRoleQuota", flatObj)

}
