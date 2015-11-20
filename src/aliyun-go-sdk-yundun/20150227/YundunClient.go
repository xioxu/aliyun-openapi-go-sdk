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

package yundun

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-yundun/20150227/WebAttackNumResponse"

	"aliyun-go-sdk-yundun/20150227/TodayqpsByRegionResponse"

	"aliyun-go-sdk-yundun/20150227/TodayMalwareNumResponse"

	"aliyun-go-sdk-yundun/20150227/TodayCrackInterceptResponse"

	"aliyun-go-sdk-yundun/20150227/TodayBackdoorResponse"

	"aliyun-go-sdk-yundun/20150227/TodayAllppsResponse"

	"aliyun-go-sdk-yundun/20150227/TodayAllkbpsResponse"

	"aliyun-go-sdk-yundun/20150227/TodayAegisOnlineRateResponse"

	"aliyun-go-sdk-yundun/20150227/CurrentDdosAttackNumResponse"

	"aliyun-go-sdk-yundun/20150227/AllMalwareNumResponse"
)

type YundunClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *YundunClient {
	p := &YundunClient{}
	p.Version = "2015-02-27"
	p.ProductName = "Yundun"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *YundunClient) WebAttackNum(webAttackNumRequest *WebAttackNumRequest) (yundun_WebAttackNumResponse.WebAttackNumResponse, error) {
	var resObj yundun_WebAttackNumResponse.WebAttackNumResponse

	if webAttackNumRequest == nil {
		webAttackNumRequest = new(WebAttackNumRequest)
	}
	err := c.DoAction(webAttackNumRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) TodayqpsByRegion(todayqpsByRegionRequest *TodayqpsByRegionRequest) (yundun_TodayqpsByRegionResponse.TodayqpsByRegionResponse, error) {
	var resObj yundun_TodayqpsByRegionResponse.TodayqpsByRegionResponse

	if todayqpsByRegionRequest == nil {
		todayqpsByRegionRequest = new(TodayqpsByRegionRequest)
	}
	err := c.DoAction(todayqpsByRegionRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) TodayMalwareNum(todayMalwareNumRequest *TodayMalwareNumRequest) (yundun_TodayMalwareNumResponse.TodayMalwareNumResponse, error) {
	var resObj yundun_TodayMalwareNumResponse.TodayMalwareNumResponse

	if todayMalwareNumRequest == nil {
		todayMalwareNumRequest = new(TodayMalwareNumRequest)
	}
	err := c.DoAction(todayMalwareNumRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) TodayCrackIntercept(todayCrackInterceptRequest *TodayCrackInterceptRequest) (yundun_TodayCrackInterceptResponse.TodayCrackInterceptResponse, error) {
	var resObj yundun_TodayCrackInterceptResponse.TodayCrackInterceptResponse

	if todayCrackInterceptRequest == nil {
		todayCrackInterceptRequest = new(TodayCrackInterceptRequest)
	}
	err := c.DoAction(todayCrackInterceptRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) TodayBackdoor(todayBackdoorRequest *TodayBackdoorRequest) (yundun_TodayBackdoorResponse.TodayBackdoorResponse, error) {
	var resObj yundun_TodayBackdoorResponse.TodayBackdoorResponse

	if todayBackdoorRequest == nil {
		todayBackdoorRequest = new(TodayBackdoorRequest)
	}
	err := c.DoAction(todayBackdoorRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) TodayAllpps(todayAllppsRequest *TodayAllppsRequest) (yundun_TodayAllppsResponse.TodayAllppsResponse, error) {
	var resObj yundun_TodayAllppsResponse.TodayAllppsResponse

	if todayAllppsRequest == nil {
		todayAllppsRequest = new(TodayAllppsRequest)
	}
	err := c.DoAction(todayAllppsRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) TodayAllkbps(todayAllkbpsRequest *TodayAllkbpsRequest) (yundun_TodayAllkbpsResponse.TodayAllkbpsResponse, error) {
	var resObj yundun_TodayAllkbpsResponse.TodayAllkbpsResponse

	if todayAllkbpsRequest == nil {
		todayAllkbpsRequest = new(TodayAllkbpsRequest)
	}
	err := c.DoAction(todayAllkbpsRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) TodayAegisOnlineRate(todayAegisOnlineRateRequest *TodayAegisOnlineRateRequest) (yundun_TodayAegisOnlineRateResponse.TodayAegisOnlineRateResponse, error) {
	var resObj yundun_TodayAegisOnlineRateResponse.TodayAegisOnlineRateResponse

	if todayAegisOnlineRateRequest == nil {
		todayAegisOnlineRateRequest = new(TodayAegisOnlineRateRequest)
	}
	err := c.DoAction(todayAegisOnlineRateRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) CurrentDdosAttackNum(currentDdosAttackNumRequest *CurrentDdosAttackNumRequest) (yundun_CurrentDdosAttackNumResponse.CurrentDdosAttackNumResponse, error) {
	var resObj yundun_CurrentDdosAttackNumResponse.CurrentDdosAttackNumResponse

	if currentDdosAttackNumRequest == nil {
		currentDdosAttackNumRequest = new(CurrentDdosAttackNumRequest)
	}
	err := c.DoAction(currentDdosAttackNumRequest, &resObj)
	return resObj, err
}

func (c *YundunClient) AllMalwareNum(allMalwareNumRequest *AllMalwareNumRequest) (yundun_AllMalwareNumResponse.AllMalwareNumResponse, error) {
	var resObj yundun_AllMalwareNumResponse.AllMalwareNumResponse

	if allMalwareNumRequest == nil {
		allMalwareNumRequest = new(AllMalwareNumRequest)
	}
	err := c.DoAction(allMalwareNumRequest, &resObj)
	return resObj, err
}
