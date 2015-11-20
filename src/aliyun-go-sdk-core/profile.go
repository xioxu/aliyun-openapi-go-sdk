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
	"strings"
)

type AccessProfile struct {
	Signer       SignatureAlgorithm
	FormatType   string
	RegionId     string
	AccessKeyId  string
	AccessSecret string
	Protocol     string
	endpoints    map[string]map[string]string
}

var defaultProfile *AccessProfile

func GetDefaultProfile(accessKey string, accessSecret string) *AccessProfile {
	if defaultProfile == nil {
		p := &AccessProfile{FormatType: "JSON"}
		p.Signer = &SignatureHmacSha1{}
		p.Protocol = "http"
		endpointFileItems := ReadEncpointConfigFile()

		p.endpoints = make(map[string]map[string]string, 0)

		for _, endpoint := range endpointFileItems {
			for _, regionId := range endpoint.RegionIds {

				if _, exist := p.endpoints[regionId]; !exist {
					//Since there are about 35 products in aliyun openapi currently
					p.endpoints[regionId] = make(map[string]string, 35)
				}

				for _, product := range endpoint.Products {
					p.endpoints[regionId][strings.ToLower(product.ProductName)] = product.DomainName
				}
			}

		}
		defaultProfile = p
	}

	defaultProfile.RegionId = "cn-hangzhou"
	defaultProfile.AccessKeyId = accessKey
	defaultProfile.AccessSecret = accessSecret
	return defaultProfile
}

func (p *AccessProfile) AddEndpoint(regionId string, productName string, domain string) {
	if _, exist := p.endpoints[regionId]; !exist {
		//Since there are about 35 products in aliyun openapi currently
		p.endpoints[regionId] = make(map[string]string, 35)
	}

	p.endpoints[regionId][strings.ToLower(productName)] = domain
}

func (p *AccessProfile) GetServiceUrl(productName string) string {
	if p.endpoints == nil || productName == "" {
		return ""
	}

	regionProducts, exists := p.endpoints[p.RegionId]

	if exists {
		domain, productExist := regionProducts[strings.ToLower(productName)]

		if !productExist {
			panic("Cannot find the specified product name in encpoint:" + productName)
		}

		return domain
	} else {
		panic("There is no valid endpoint in retion:" + p.RegionId)
	}
}
