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
	"errors"
	"testing"
)

func TestAddToMapIfNotDefaultValueStr(t *testing.T) {
	mapVals := make(map[string]string)

	AddToMapIfNotDefaultValueStr(mapVals, "k1", "")
	AddToMapIfNotDefaultValueStr(mapVals, "k2", "not empty")

	if _, ok := mapVals["k1"]; ok {
		t.Error("")
	}

	if _, ok := mapVals["k2"]; !ok {
		t.Error("")
	}
}

func TestAddToMapIfNotDefaultValueInt32(t *testing.T) {
	mapVals := make(map[string]string)

	AddToMapIfNotDefaultValueInt32(mapVals, "k1", 0)
	AddToMapIfNotDefaultValueInt32(mapVals, "k2", 1)

	if _, ok := mapVals["k1"]; ok {
		t.Error("")
	}

	if _, ok := mapVals["k2"]; !ok {
		t.Error("")
	}
}

func TestAddToMapIfNotDefaultValueInt64(t *testing.T) {
	mapVals := make(map[string]string)

	AddToMapIfNotDefaultValueInt64(mapVals, "k1", int64(0))
	AddToMapIfNotDefaultValueInt64(mapVals, "k2", int64(1))

	if _, ok := mapVals["k1"]; ok {
		t.Error("")
	}

	if _, ok := mapVals["k2"]; !ok {
		t.Error("")
	}
}

func TestAddToMapIfNotDefaultValueBool(t *testing.T) {
	mapVals := make(map[string]string)

	AddToMapIfNotDefaultValueBoolean(mapVals, "k1", false)
	AddToMapIfNotDefaultValueBoolean(mapVals, "k2", true)

	if _, ok := mapVals["k1"]; ok {
		t.Error("")
	}

	if _, ok := mapVals["k2"]; !ok {
		t.Error("")
	}
}

func TestAddToMapIfNotDefaultValueFloat32(t *testing.T) {
	mapVals := make(map[string]string)

	AddToMapIfNotDefaultValueFloat32(mapVals, "k1", 0)
	AddToMapIfNotDefaultValueFloat32(mapVals, "k2", 0.1)

	if _, ok := mapVals["k1"]; ok {
		t.Error("")
	}

	if _, ok := mapVals["k2"]; !ok {
		t.Error("")
	}
}

func TestAddToMapIfNotDefaultValueFloat64(t *testing.T) {
	mapVals := make(map[string]string)

	AddToMapIfNotDefaultValueFloat64(mapVals, "k1", 0)
	AddToMapIfNotDefaultValueFloat64(mapVals, "k2", 0.1)

	if _, ok := mapVals["k1"]; ok {
		t.Error("")
	}

	if _, ok := mapVals["k2"]; !ok {
		t.Error("")
	}
}

func TestCheckError(t *testing.T) {
	CheckError(nil, "")

	expectPanic(t, func() {
		CheckError(errors.New("test"), "")
	})
}
