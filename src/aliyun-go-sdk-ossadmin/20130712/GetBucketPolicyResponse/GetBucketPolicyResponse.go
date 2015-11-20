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

package ossadmin_GetBucketPolicyResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type GetBucketPolicyResponse struct {
	DisallowEmptyRefer bool
	EnableDualCluster  bool
	ErrorFile          string
	IndexFile          string
	Location           string
	LogBucket          string
	LogPrefix          string
	IamPolicy          string
	RequestId          string
	WhiteReferList     WhiteReferList
}

func (m *GetBucketPolicyResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.DisallowEmptyRefer = builder.GetBool(parentKey+".DisallowEmptyRefer", flatObj)

	m.EnableDualCluster = builder.GetBool(parentKey+".EnableDualCluster", flatObj)

	m.ErrorFile = builder.GetString(parentKey+".ErrorFile", flatObj)

	m.IndexFile = builder.GetString(parentKey+".IndexFile", flatObj)

	m.Location = builder.GetString(parentKey+".Location", flatObj)

	m.LogBucket = builder.GetString(parentKey+".LogBucket", flatObj)

	m.LogPrefix = builder.GetString(parentKey+".LogPrefix", flatObj)

	m.IamPolicy = builder.GetString(parentKey+".IamPolicy", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

	m.WhiteReferList = WhiteReferList{}
	m.WhiteReferList.BuildProperties(parentKey+".WhiteReferList", flatObj)

}
