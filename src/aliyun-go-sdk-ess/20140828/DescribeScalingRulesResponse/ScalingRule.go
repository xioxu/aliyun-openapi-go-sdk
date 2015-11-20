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

package ess_DescribeScalingRulesResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type ScalingRule struct {
	ScalingRuleId   string
	ScalingGroupId  string
	ScalingRuleName string
	Cooldown        int32
	AdjustmentType  string
	AdjustmentValue int32
	MinSize         int32
	MaxSize         int32
	ScalingRuleAri  string
}

func (m *ScalingRule) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.ScalingRuleId = builder.GetString(parentKey+".ScalingRuleId", flatObj)

	m.ScalingGroupId = builder.GetString(parentKey+".ScalingGroupId", flatObj)

	m.ScalingRuleName = builder.GetString(parentKey+".ScalingRuleName", flatObj)

	m.Cooldown = builder.GetInt32(parentKey+".Cooldown", flatObj)

	m.AdjustmentType = builder.GetString(parentKey+".AdjustmentType", flatObj)

	m.AdjustmentValue = builder.GetInt32(parentKey+".AdjustmentValue", flatObj)

	m.MinSize = builder.GetInt32(parentKey+".MinSize", flatObj)

	m.MaxSize = builder.GetInt32(parentKey+".MaxSize", flatObj)

	m.ScalingRuleAri = builder.GetString(parentKey+".ScalingRuleAri", flatObj)

}
