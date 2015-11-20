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

package crm

import "aliyun-go-sdk-core"

type AddLabelRequest struct {
	PK          string
	LabelName   string
	LabelSeries string
	EndTime     string
}

func (r *AddLabelRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 4)

	core.AddToMapIfNotDefaultValueStr(queryVals, "PK", r.PK)
	core.AddToMapIfNotDefaultValueStr(queryVals, "LabelName", r.LabelName)
	core.AddToMapIfNotDefaultValueStr(queryVals, "LabelSeries", r.LabelSeries)
	core.AddToMapIfNotDefaultValueStr(queryVals, "EndTime", r.EndTime)

	return queryVals
}

func (r *AddLabelRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *AddLabelRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *AddLabelRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *AddLabelRequest) GetPath() string {
	return ""
}

func (r *AddLabelRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *AddLabelRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *AddLabelRequest) GetActionName() string {
	return "AddLabel"
}
