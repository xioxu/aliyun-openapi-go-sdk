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
	"strings"
	"time"
)

var int32Default int32
var int64Default int64
var timeDefault time.Time
var float32Default float32
var boolDefault bool
var float64Default float64

func AddToMapIfNotDefaultValueStr(mapVals map[string]string, k string, v string) {
	if v != "" {
		mapVals[k] = v
	}
}

func AddToMapIfNotDefaultValueStrArr(mapVals map[string]string, k string, v []string) {
	if len(v) > 0 {
		mapVals[k] = strings.Join(v, ",")
	}
}

func AddToMapIfNotDefaultValueInt32(mapVals map[string]string, k string, v int32) {

	if v != int32Default {
		mapVals[k] = builder.ToString(v)
	}
}

func AddToMapIfNotDefaultValueInt64(mapVals map[string]string, k string, v int64) {
	if v != int64Default {
		mapVals[k] = builder.ToString(v)
	}
}

func AddToMapIfNotDefaultValueBoolean(mapVals map[string]string, k string, v bool) {
	if v != boolDefault {
		mapVals[k] = builder.ToString(v)
	}
}

func AddToMapIfNotDefaultValueTime(mapVals map[string]string, k string, v time.Time) {
	if v != timeDefault {
		mapVals[k] = builder.ToString(v)
	}
}

func AddToMapIfNotDefaultValueFloat32(mapVals map[string]string, k string, v float32) {
	if v != float32Default {
		mapVals[k] = builder.ToString(v)
	}
}

func AddToMapIfNotDefaultValueFloat64(mapVals map[string]string, k string, v float64) {
	if v != float64Default {
		mapVals[k] = builder.ToString(v)
	}
}

func CheckError(err error, msg string) {
	if err != nil {
		panic(msg + " " + err.Error())
	}
}
