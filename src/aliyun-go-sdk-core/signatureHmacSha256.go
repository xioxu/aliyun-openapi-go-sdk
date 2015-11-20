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
	"aliyun-go-sdk-core/parameterHelper"
	"crypto/hmac"
	"crypto/sha256"
)

type SignatureHmacSha256 struct {
}

func (*SignatureHmacSha256) MakeSignature(secret string, strToSignature string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(strToSignature))
	decodeBytes := parameterHelper.Base64Encode((h.Sum(nil)))

	return decodeBytes
}

func (*SignatureHmacSha256) GetVersion() string {
	return "1.0"
}

func (*SignatureHmacSha256) GetSignatureMethod() string {
	return "HMAC-SHA1"
}
