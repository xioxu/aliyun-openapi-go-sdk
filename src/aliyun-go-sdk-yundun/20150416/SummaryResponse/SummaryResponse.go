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

package yundun_SummaryResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type SummaryResponse struct {
	Status            int64
	AbnormalHostCount int64
	RequestId         string
	Ddos              Ddos
	BruteForce        BruteForce
	Webshell          Webshell
	RemoteLogin       RemoteLogin
	WebAttack         WebAttack
	WebLeak           WebLeak
}

func (m *SummaryResponse) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.Status = builder.GetInt64(parentKey+".Status", flatObj)

	m.AbnormalHostCount = builder.GetInt64(parentKey+".AbnormalHostCount", flatObj)

	m.RequestId = builder.GetString(parentKey+".RequestId", flatObj)

	m.Ddos = Ddos{}
	m.Ddos.BuildProperties(parentKey+".Ddos", flatObj)

	m.BruteForce = BruteForce{}
	m.BruteForce.BuildProperties(parentKey+".BruteForce", flatObj)

	m.Webshell = Webshell{}
	m.Webshell.BuildProperties(parentKey+".Webshell", flatObj)

	m.RemoteLogin = RemoteLogin{}
	m.RemoteLogin.BuildProperties(parentKey+".RemoteLogin", flatObj)

	m.WebAttack = WebAttack{}
	m.WebAttack.BuildProperties(parentKey+".WebAttack", flatObj)

	m.WebLeak = WebLeak{}
	m.WebLeak.BuildProperties(parentKey+".WebLeak", flatObj)

}
