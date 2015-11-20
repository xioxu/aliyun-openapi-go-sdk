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

//Package builder prviders some help methods for construct a response modal
package builder

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
)

func UnmarshalJson(content []byte, obj interface{}) error {
	return json.Unmarshal(content, obj)
}

func UnmarshalXml(content []byte, obj interface{}) error {
	return xml.Unmarshal(content, obj)
}

type node struct {
	XMLName xml.Name
	Content []byte `xml:",innerxml"`
	Nodes   []node `xml:",any"`
}

//Convert xml content to map[string]string
func FlatXmlContentToMap(k string, xmlContent []byte, reconstructedMap map[string]interface{}) {

	var xmlObj node
	err := xml.Unmarshal(xmlContent, &xmlObj)
	if err != nil {
		panic(err)
	}

	walk([]node{xmlObj}, func(n node) {
		fmt.Println("---")
		fmt.Println(n.XMLName.Local)
		fmt.Println(string(n.Content))
	})
}

func walk(nodes []node, f func(node)) {

	for _, n := range nodes {
		f(n)
		walk(n.Nodes, f)
	}
}

//Convert json content to map[string]interface{}
func FlatJsonObjToMap(k string, v interface{}, reconstructedMap map[string]interface{}) {

	switch val := v.(type) {
	case int:
		reconstructedMap[k] = val
		break
	case int32:
		reconstructedMap[k] = val
		break
	case int64:
		reconstructedMap[k] = val
		break
	case float32:
		reconstructedMap[k] = val
		break
	case float64:
		reconstructedMap[k] = val
		break
	case string:
		reconstructedMap[k] = val
		break
	case []interface{}:
		reconstructedMap[getParentKey(k)+".length"] = len(val)
		for idx, ele := range val {
			FlatJsonObjToMap(fmt.Sprintf("%s[%d]", getParentKey(k), idx), ele, reconstructedMap)
		}
		break
	case map[string]interface{}:
		for kMap, vMap := range val {
			FlatJsonObjToMap(k+"."+kMap, vMap, reconstructedMap)
		}
		break
	default:
		reconstructedMap[k] = val
		break
	}
}

//Get the index key, e.g: if the input is 'a.b.c', the output should be 'a.b'
func getParentKey(k string) string {
	if k == "" {
		return ""
	}

	keys := strings.Split(k, ".")

	return strings.Join(keys[:len(keys)-1], ".")
}
