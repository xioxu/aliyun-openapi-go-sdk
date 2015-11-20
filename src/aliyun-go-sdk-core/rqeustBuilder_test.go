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
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestGetRequestMethod(t *testing.T) {
	reqMtd1 := "get"
	reqMtd2 := "post"
	reqMtd3 := "PUT"

	if "GET" != GetRequestMethod(reqMtd1) {
		t.Error("")
	}

	if "POST" != GetRequestMethod(reqMtd2) {
		t.Error("")
	}

	if "PUT" != GetRequestMethod(reqMtd3) {
		t.Error("")
	}
}

func TestGetRequestMethodForPanic(t *testing.T) {
	expectPanic(t, func() {
		reqMtd1 := "put,get"

		GetRequestMethod(reqMtd1)
	})
}

func TestConcatUrl(t *testing.T) {

	if concatUrl("http://a.com/b", "subpath/aaa") != "http://a.com/b/subpath/aaa" {
		t.Error("")
	}

	if concatUrl("http://a.com/b/", "/subpath/aaa") != "http://a.com/b/subpath/aaa" {
		t.Error("")
	}

	if concatUrl("http://a.com/b/", "subpath/aaa") != "http://a.com/b/subpath/aaa" {
		t.Error("")
	}

	if concatUrl("http://a.com/b", "/subpath/aaa") != "http://a.com/b/subpath/aaa" {
		t.Error("")
	}
}

func TestmergeMap(t *testing.T) {

	m1 := make(map[string]string)
	m2 := make(map[string]string)

	m1["k1"] = "v1"
	m2["k2"] = "v2"
	mergeMap(m1, m2)

	if _, ok := m1["k2"]; !ok {
		t.Error("")
	}
}

func TestgetAcceptContentType(t *testing.T) {
	if getAcceptContentType("", "put") != "application/octet-stream" {
		t.Error("")
	}

	if getAcceptContentType("xml", "get") != "application/xml" {
		t.Error("")
	}

	if getAcceptContentType("json", "post") != "application/json" {
		t.Error("")
	}
}

func TestGetProductUrl(t *testing.T) {
	client := new(BaseClient)
	client.Profile = GetDefaultProfile("", "")

	e := client.Profile.endpoints

	for k, _ := range e {
		fmt.Println(k)
	}

	client.ProductName = "Ecs"

	if GetProductUrl(client) != "ecs-cn-hangzhou.aliyuncs.com" {
		t.Error("TestGetProductUrl failed ")
	}
}

func TestGetProductUrlExpectException(t *testing.T) {
	client := new(BaseClient)
	client.Profile = GetDefaultProfile("", "")
	client.ProductName = "abc"

	expectPanic(t, func() {
		GetProductUrl(client)
	})
}

func TestMergeMapToUrlValues(t *testing.T) {

	u := make(url.Values)
	m := make(map[string]string)
	m["k1"] = "v1"
	m["k2"] = "v2"

	mergeMapToUrlValues(u, m)
	_, ok := u["k1"]

	if len(u) != 2 || !ok {
		t.Error("TestMergeMapToUrlValues failed")
	}
}

func TestMergeUrlValues(t *testing.T) {
	u1 := make(url.Values)
	u2 := make(url.Values)

	u2.Add("k1", "v1")
	u2.Add("k2", "v2")

	mergeUrlValues(u1, u2)

	_, ok := u1["k2"]

	if !ok || len(u1) != 2 {
		t.Error("TestMergeUrlValues failed")
	}
}

func TestAddSdkVersion(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://aa", nil)

	addSdkVersion(req)

	v := req.Header.Get("x-sdk-client")

	if v == "" {
		t.Error("TestAddSdkVersion failed")
	}
}
