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

package alert

import "aliyun-go-sdk-core"
import "fmt"

type ListDimensionsRequest struct {
	ProjectName string
	AlertName   string
	Page        int32
	PageSize    int32
}

func (r *ListDimensionsRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 2)

	core.AddToMapIfNotDefaultValueInt32(queryVals, "Page", r.Page)
	core.AddToMapIfNotDefaultValueInt32(queryVals, "PageSize", r.PageSize)

	return queryVals
}

func (r *ListDimensionsRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *ListDimensionsRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *ListDimensionsRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *ListDimensionsRequest) GetPath() string {
	return fmt.Sprintf("/projects/%s/alerts/%s/dimensions", r.ProjectName, r.AlertName)
}

func (r *ListDimensionsRequest) GetRequestMethod() string {
	return "GET"
}

func (r *ListDimensionsRequest) GetProtocol() string {
	return "HTTP"
}

func (r *ListDimensionsRequest) GetActionName() string {
	return "ListDimensions"
}
