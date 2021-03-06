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

type DisableAlertRequest struct {
	ProjectName string
	AlertName   string
}

func (r *DisableAlertRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 0)

	return queryVals
}

func (r *DisableAlertRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *DisableAlertRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *DisableAlertRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *DisableAlertRequest) GetPath() string {
	return fmt.Sprintf("/projects/%s/alerts/%s/disable", r.ProjectName, r.AlertName)
}

func (r *DisableAlertRequest) GetRequestMethod() string {
	return "PUT"
}

func (r *DisableAlertRequest) GetProtocol() string {
	return "HTTP"
}

func (r *DisableAlertRequest) GetActionName() string {
	return "DisableAlert"
}
