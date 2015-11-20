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

type DeleteDBSourceRequest struct {
	ProjectName string
	SourceName  string
}

func (r *DeleteDBSourceRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 0)

	return queryVals
}

func (r *DeleteDBSourceRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *DeleteDBSourceRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *DeleteDBSourceRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *DeleteDBSourceRequest) GetPath() string {
	return fmt.Sprintf("/projects/%s/sources/%s", r.ProjectName, r.SourceName)
}

func (r *DeleteDBSourceRequest) GetRequestMethod() string {
	return "DELETE"
}

func (r *DeleteDBSourceRequest) GetProtocol() string {
	return "HTTP"
}

func (r *DeleteDBSourceRequest) GetActionName() string {
	return "DeleteDBSource"
}
