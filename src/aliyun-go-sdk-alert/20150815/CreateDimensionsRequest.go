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

type CreateDimensionsRequest struct {
	ProjectName string
	AlertName   string
	Dimensions  string
}

func (r *CreateDimensionsRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 0)

	return queryVals
}

func (r *CreateDimensionsRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 1)

	core.AddToMapIfNotDefaultValueStr(bodyVals, "Dimensions", r.Dimensions)

	return bodyVals
}

func (r *CreateDimensionsRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *CreateDimensionsRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *CreateDimensionsRequest) GetPath() string {
	return fmt.Sprintf("/projects/%s/alerts/%s/dimensions", r.ProjectName, r.AlertName)
}

func (r *CreateDimensionsRequest) GetRequestMethod() string {
	return "POST"
}

func (r *CreateDimensionsRequest) GetProtocol() string {
	return "HTTP"
}

func (r *CreateDimensionsRequest) GetActionName() string {
	return "CreateDimensions"
}
