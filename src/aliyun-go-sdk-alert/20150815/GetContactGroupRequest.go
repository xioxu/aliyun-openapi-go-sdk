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

import "fmt"

type GetContactGroupRequest struct {
	ProjectName string
	GroupName   string
}

func (r *GetContactGroupRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 0)

	return queryVals
}

func (r *GetContactGroupRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *GetContactGroupRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *GetContactGroupRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *GetContactGroupRequest) GetPath() string {
	return fmt.Sprintf("/projects/%s/groups/%s", r.ProjectName, r.GroupName)
}

func (r *GetContactGroupRequest) GetRequestMethod() string {
	return "GET"
}

func (r *GetContactGroupRequest) GetProtocol() string {
	return "HTTP"
}

func (r *GetContactGroupRequest) GetActionName() string {
	return "GetContactGroup"
}
