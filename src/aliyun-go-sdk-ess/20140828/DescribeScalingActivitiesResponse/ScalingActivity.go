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

package ess_DescribeScalingActivitiesResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type ScalingActivity struct {
	ScalingActivityId string
	ScalingGroupId    string
	Description       string
	Cause             string
	StartTime         string
	EndTime           string
	Progress          int32
	StatusCode        string
	StatusMessage     string
}

func (m *ScalingActivity) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.ScalingActivityId = builder.GetString(parentKey+".ScalingActivityId", flatObj)

	m.ScalingGroupId = builder.GetString(parentKey+".ScalingGroupId", flatObj)

	m.Description = builder.GetString(parentKey+".Description", flatObj)

	m.Cause = builder.GetString(parentKey+".Cause", flatObj)

	m.StartTime = builder.GetString(parentKey+".StartTime", flatObj)

	m.EndTime = builder.GetString(parentKey+".EndTime", flatObj)

	m.Progress = builder.GetInt32(parentKey+".Progress", flatObj)

	m.StatusCode = builder.GetString(parentKey+".StatusCode", flatObj)

	m.StatusMessage = builder.GetString(parentKey+".StatusMessage", flatObj)

}
