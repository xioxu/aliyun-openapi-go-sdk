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

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
	"testing"
)

type DescribeAppLogsRequest struct {
	AppId      string
	StartTime  string
	EndTime    string
	PageSize   int32
	PageNumber int32
}

func (r *DescribeAppLogsRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 5)

	AddToMapIfNotDefaultValueStr(queryVals, "AppId", r.AppId)
	AddToMapIfNotDefaultValueStr(queryVals, "StartTime", r.StartTime)
	AddToMapIfNotDefaultValueStr(queryVals, "EndTime", r.EndTime)
	AddToMapIfNotDefaultValueInt32(queryVals, "PageSize", r.PageSize)
	AddToMapIfNotDefaultValueInt32(queryVals, "PageNumber", r.PageNumber)

	return queryVals
}

func (r *DescribeAppLogsRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *DescribeAppLogsRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *DescribeAppLogsRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *DescribeAppLogsRequest) GetPath() string {
	return ""
}

func (r *DescribeAppLogsRequest) GetRequestMethod() string {
	return "POST|GET"
}

func (r *DescribeAppLogsRequest) GetProtocol() string {
	return "HTTP|HTTPS"
}

func (r *DescribeAppLogsRequest) GetActionName() string {
	return "DescribeAppLogs"
}

type DescribeAppLogsResponse struct {
	Items           []AppLog
	PageRecordCount int32
	RequestId       string
}

func (m *DescribeAppLogsResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.Items = make([]AppLog, 0)
	for i := 0; i < (flatObj[parentKey+".Items.length"]).(int); i++ {
		tmp := AppLog{}
		tmp.BuildProperties(fmt.Sprintf(parentKey+".Items[%d]", i), flatObj)
		m.Items = append(m.Items, tmp)
	}

	m.PageRecordCount = builder.GetInt32(parentKey+".PageRecordCount", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

}

type AppLog struct {
	LogTime    string
	LogContent string
}

func (m *AppLog) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.LogTime = builder.GetString(parentKey+".LogTime", flatObj)

	m.LogContent = builder.GetString(parentKey+".LogContent", flatObj)

}

type AceClient struct {
	BaseClient
}

func New(profile *AccessProfile) *AceClient {
	p := &AceClient{}
	p.Version = "2015-07-13"
	p.ProductName = "Ace"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &RpcHttpRequestBuilder{}
	return p
}

func (c *AceClient) DescribeAppLogs(describeAppLogsRequest *DescribeAppLogsRequest) (DescribeAppLogsResponse, error) {
	var resObj DescribeAppLogsResponse

	if describeAppLogsRequest == nil {
		describeAppLogsRequest = new(DescribeAppLogsRequest)
	}
	err := c.DoAction(describeAppLogsRequest, &resObj)
	return resObj, err
}

func TestDoAction(t *testing.T) {
	accessProfile := GetDefaultProfile("accessKey", "secret")
	client := New(accessProfile)

	req := &DescribeAppLogsRequest{}
	_, err := client.DescribeAppLogs(req)

	if err == nil {
		t.Error("TestDoAction failed")
	}
}
