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

package rds_DescribeOptimizeAdviceOnBigTableResponse

import (
	"aliyun-go-sdk-core/objectBuilder"
)

type AdviceOnBigTable struct {
	DBName    string
	TableName string
	TableSize int64
	DataSize  int64
	IndexSize int64
}

func (m *AdviceOnBigTable) BuildProperties(parentKey string, flatObj map[string]interface{}) {

	m.DBName = builder.GetString(parentKey+".DBName", flatObj)

	m.TableName = builder.GetString(parentKey+".TableName", flatObj)

	m.TableSize = builder.GetInt64(parentKey+".TableSize", flatObj)

	m.DataSize = builder.GetInt64(parentKey+".DataSize", flatObj)

	m.IndexSize = builder.GetInt64(parentKey+".IndexSize", flatObj)

}
