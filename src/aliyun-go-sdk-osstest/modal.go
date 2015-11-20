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

package oss

import (
	"aliyun-go-sdk-core/objectBuilder"
	"fmt"
)

type ListBucketRequest struct {
}

func (r *ListBucketRequest) GetQuery() map[string]string {
	var queryVals = make(map[string]string, 0)
	return queryVals
}

func (r *ListBucketRequest) GetBody() map[string]string {
	var bodyVals = make(map[string]string, 0)

	return bodyVals
}

func (r *ListBucketRequest) GetHeaders() map[string]string {
	var headerVals = make(map[string]string, 0)

	return headerVals
}

func (r *ListBucketRequest) GetDomainValues() map[string]string {
	var domainVals = make(map[string]string, 0)

	return domainVals
}

func (r *ListBucketRequest) GetPath() string {
	return ""
}

func (r *ListBucketRequest) GetRequestMethod() string {
	return "GET"
}

func (r *ListBucketRequest) GetProtocol() string {
	return "HTTP"
}

func (r *ListBucketRequest) GetActionName() string {
	return "ListBucket"
}

type Owner struct {
	Id          string
	DisplayName string
}

func (m *Owner) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.Id = builder.GetString(parentKey+".Id", flatObj)
	m.DisplayName = builder.GetString(parentKey+".DisplayName", flatObj)

}

type Bucket struct {
	Location     string
	Name         string
	CreationDate string
}

func (m *Bucket) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.Location = builder.GetString(parentKey+".Location", flatObj)
	m.Name = builder.GetString(parentKey+".Name", flatObj)
	m.CreationDate = builder.GetString(parentKey+".CreationDate", flatObj)

}

type ListBucketResponse struct {
	Owner   Owner
	Buckets []Bucket
}

func (m *ListBucketResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.Owner = Owner{}
	m.Owner.BuildProperties(parentKey, flatObj)

	m.Buckets = make([]Bucket, 0)

	for i := 0; i < (flatObj[".Buckets.length"]).(int); i++ {
		tmp := Bucket{}
		tmp.BuildProperties(fmt.Sprintf(".Buckets[%d]", i), flatObj)
		m.Buckets = append(m.Buckets, tmp)
	}

}
