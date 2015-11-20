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

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-crm/20150408/DeleteLabelResponse"

	"aliyun-go-sdk-crm/20150408/CheckLabelResponse"

	"aliyun-go-sdk-crm/20150408/AddLabelResponse"

	"aliyun-go-sdk-crm/20150408/QueryCustomerLabelResponse"
)

type CrmClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *CrmClient {
	p := &CrmClient{}
	p.Version = "2015-04-08"
	p.ProductName = "Crm"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *CrmClient) DeleteLabel(deleteLabelRequest *DeleteLabelRequest) (crm_DeleteLabelResponse.DeleteLabelResponse, error) {
	var resObj crm_DeleteLabelResponse.DeleteLabelResponse

	if deleteLabelRequest == nil {
		deleteLabelRequest = new(DeleteLabelRequest)
	}
	err := c.DoAction(deleteLabelRequest, &resObj)
	return resObj, err
}

func (c *CrmClient) CheckLabel(checkLabelRequest *CheckLabelRequest) (crm_CheckLabelResponse.CheckLabelResponse, error) {
	var resObj crm_CheckLabelResponse.CheckLabelResponse

	if checkLabelRequest == nil {
		checkLabelRequest = new(CheckLabelRequest)
	}
	err := c.DoAction(checkLabelRequest, &resObj)
	return resObj, err
}

func (c *CrmClient) AddLabel(addLabelRequest *AddLabelRequest) (crm_AddLabelResponse.AddLabelResponse, error) {
	var resObj crm_AddLabelResponse.AddLabelResponse

	if addLabelRequest == nil {
		addLabelRequest = new(AddLabelRequest)
	}
	err := c.DoAction(addLabelRequest, &resObj)
	return resObj, err
}

func (c *CrmClient) QueryCustomerLabel(queryCustomerLabelRequest *QueryCustomerLabelRequest) (crm_QueryCustomerLabelResponse.QueryCustomerLabelResponse, error) {
	var resObj crm_QueryCustomerLabelResponse.QueryCustomerLabelResponse

	if queryCustomerLabelRequest == nil {
		queryCustomerLabelRequest = new(QueryCustomerLabelRequest)
	}
	err := c.DoAction(queryCustomerLabelRequest, &resObj)
	return resObj, err
}
