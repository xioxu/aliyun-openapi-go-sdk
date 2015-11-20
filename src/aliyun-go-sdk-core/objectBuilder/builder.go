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

//Package 'builder' contails the helper method for converting response content to response modal.
package builder

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//Convert int/bool/time to string
func ToString(v interface{}) string {
	switch val := v.(type) {
	case int:
		return fmt.Sprintf("%d", val)
	case int32:
		return fmt.Sprintf("%d", val)
	case int64:
		return strconv.FormatInt(val, 10)
	case bool:
		return strconv.FormatBool(val)
	case string:
		return val
	case float32:
		return fmt.Sprintf("%0.2f", val)
	case float64:
		return fmt.Sprintf("%0.2f", val)
	case time.Time:
		return val.Format("2006-01-02 15:04:05")
	default:
		return fmt.Sprintf("%s", val)
	}
}

//Get the string value according the key in a map
func GetString(key string, resVals map[string]interface{}) string {
	v, existed := resVals[key]

	if existed {
		return ToString(v)

	} else {
		var fromSubObj []string = make([]string, 0)

		for k, v := range resVals {
			if strings.HasPrefix(k, key+".") {
				fromSubObj = append(fromSubObj, fmt.Sprintf("%s:%s", strings.Replace(k, key+".", "", 1), ToString(v)))
			}
		}

		return strings.Join(fromSubObj, ",")
	}
}

//Get the int value according the key in a map
func GetInt(key string, resVals map[string]interface{}) int {
	switch val := resVals[key].(type) {
	case int:
		return val
	case int64:
		return int(val)
	case string:
		r, _ := strconv.Atoi(val)
		return r
	case float64:
		return int(val)
	default:
		return -1
	}
}

//Get the int32 value according the key in a map
func GetInt32(key string, resVals map[string]interface{}) int32 {
	switch val := resVals[key].(type) {
	case int:
		return int32(val)
	case int64:
		return int32(val)
	case string:
		r, _ := strconv.Atoi(val)
		return int32(r)
	case float64:
		return int32(val)
	default:
		return -1
	}
}

//Get the int64 value according the key in a map
func GetInt64(key string, resVals map[string]interface{}) int64 {
	switch val := resVals[key].(type) {
	case int:
		return int64(val)
	case int64:
		return int64(val)
	case string:
		r, _ := strconv.ParseInt(val, 10, 64)
		return r
	case float64:
		return int64(val)
	default:
		return -1
	}
}

//Get the float32 value according the key in a map
func GetFloat32(key string, resVals map[string]interface{}) float32 {
	switch val := resVals[key].(type) {
	case int:
		return float32(val)
	case int64:
		return float32(val)
	case string:
		r, _ := strconv.ParseFloat(val, 32)
		return float32(r)
	case float64:
		return float32(val)
	default:
		return -1
	}
}

//Get the float64 value according the key in a map
func GetFloat64(key string, resVals map[string]interface{}) float64 {
	switch val := resVals[key].(type) {
	case int:
		return float64(val)
	case int64:
		return float64(val)
	case string:
		r, _ := strconv.ParseFloat(val, 64)
		return float64(r)
	case float64:
		return val
	default:
		return -1
	}
}

//Get the bool value according the key in a map
func GetBool(key string, resVals map[string]interface{}) bool {
	return resVals[key].(bool)
}

//Get the time value according the key in a map
func GetTime(key string, resVals map[string]interface{}) time.Time {
	return resVals[key].(time.Time)
}
