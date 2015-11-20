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

package core

type BaseRequest interface {
	//Get the request values which should be submitted in query
	GetQuery() map[string]string
	//Get the request sub url
	GetPath() string
	//Get the supported request protocol
	GetProtocol() string
	//Get the request values which should be submitted in body
	GetBody() map[string]string
	//Get the request values which should be submitted in header
	GetHeaders() map[string]string
	//Get the request values which should be submitted in domain
	GetDomainValues() map[string]string
	//Get the supported request method for this api
	GetRequestMethod() string
	//Get the request action name
	GetActionName() string
}
