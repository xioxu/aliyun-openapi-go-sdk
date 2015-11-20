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

package rds

import "aliyun-go-sdk-core"

type DescribeBackupsRequest struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	DBInstanceId         string
	BackupId             string
	BackupLocation       string
	BackupStatus         string
	BackupMode           string
	StartTime            string
	EndTime              string
	PageSize             int32
	PageNumber           int32
	OwnerAccount         string
}

func (r *DescribeBackupsRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 13)

	core.AddToMapIfNotDefaultValueInt64(queryVals, "OwnerId", r.OwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "ResourceOwnerAccount", r.ResourceOwnerAccount)
	core.AddToMapIfNotDefaultValueInt64(queryVals, "ResourceOwnerId", r.ResourceOwnerId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "DBInstanceId", r.DBInstanceId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "BackupId", r.BackupId)
	core.AddToMapIfNotDefaultValueStr(queryVals, "BackupLocation", r.BackupLocation)
	core.AddToMapIfNotDefaultValueStr(queryVals, "BackupStatus", r.BackupStatus)
	core.AddToMapIfNotDefaultValueStr(queryVals, "BackupMode", r.BackupMode)
	core.AddToMapIfNotDefaultValueStr(queryVals, "StartTime", r.StartTime)
	core.AddToMapIfNotDefaultValueStr(queryVals, "EndTime", r.EndTime)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageSize", r.PageSize)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageNumber", r.PageNumber)
	core.AddToMapIfNotDefaultValueStr(queryVals, "OwnerAccount", r.OwnerAccount)

	return queryVals
}

func (r *DescribeBackupsRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *DescribeBackupsRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *DescribeBackupsRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *DescribeBackupsRequest) GetPath() string {
	return ""
}

func (r *DescribeBackupsRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *DescribeBackupsRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *DescribeBackupsRequest) GetActionName() string {
	return "DescribeBackups"
}
