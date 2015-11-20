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

package ossadmin

import "aliyun-go-sdk-core"

type PutBucketPolicyRequest struct {
	Uid                string
	Bid                string
	BucketName         string
	IamPolicy          string
	DisallowEmptyRefer bool
	EnableDualCluster  bool
	ErrorFile          string
	IndexFile          string
	Location           string
	LogBucket          string
	LogPrefix          string
	WhiteReferList     string
	OwnerAccount       string
}

func (r *PutBucketPolicyRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 13)

	core.AddToMapIfNotDefaultValueStr(queryVals, "uid", r.Uid)
	core.AddToMapIfNotDefaultValueStr(queryVals, "bid", r.Bid)
	core.AddToMapIfNotDefaultValueStr(queryVals, "BucketName", r.BucketName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "IamPolicy", r.IamPolicy)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "DisallowEmptyRefer", r.DisallowEmptyRefer)
	core.AddToMapIfNotDefaultValueBoolean(queryVals, "EnableDualCluster", r.EnableDualCluster)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ErrorFile", r.ErrorFile)
	core.AddToMapIfNotDefaultValueStr(queryVals, "IndexFile", r.IndexFile)
	core.AddToMapIfNotDefaultValueStr(queryVals, "Location", r.Location)
	core.AddToMapIfNotDefaultValueStr(queryVals, "LogBucket", r.LogBucket)
	core.AddToMapIfNotDefaultValueStr(queryVals, "LogPrefix", r.LogPrefix)
	core.AddToMapIfNotDefaultValueStr(queryVals, "WhiteReferList", r.WhiteReferList)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *PutBucketPolicyRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *PutBucketPolicyRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *PutBucketPolicyRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *PutBucketPolicyRequest) GetPath() string {
	return ""
}

func (r *PutBucketPolicyRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *PutBucketPolicyRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *PutBucketPolicyRequest) GetActionName() string {
	return "PutBucketPolicy"
}
