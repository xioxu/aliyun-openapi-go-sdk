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

//This package contains the helper methods which is used during http request buiding.
package parameterHelper

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/url"
	"time"
)

const (
	base64CodeTbl = `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/`
)

var base64Encoding *base64.Encoding

func ParseToISO8601Time(t time.Time) string {
	return t.UTC().Format("2006-01-02T15:04:05Z")
}

func ParseToRFC2616Date(t time.Time) string {
	gmt, _ := time.LoadLocation("GMT")
	return t.In(gmt).Format(time.RFC1123)
}

func ParseFromISO8601Time(t string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05Z", t)
}

func ParseFromRFC2616Date(t string) (time.Time, error) {
	return time.Parse(time.RFC1123, t)
}

func GetISO8601Time() string {
	return ParseToISO8601Time(time.Now())
}

func GetRFC2616Date() string {
	return ParseToRFC2616Date(time.Now())
}

func Md5(values []byte) []byte {
	p := md5.New()
	p.Write(values)
	return p.Sum(nil)
}

func Md5Str(values []byte) string {
	return Base64Encode(Md5(values))
}

func Base64Encode(values []byte) string {
	if base64Encoding == nil {
		base64Encoding = base64.NewEncoding(base64CodeTbl)
	}

	return base64Encoding.EncodeToString(values)
}

func ConvertMapToUrlValues(mp map[string]string) url.Values {
	urlVals := make(url.Values)
	for k, v := range mp {
		urlVals.Add(k, v)
	}

	return urlVals
}

func GetSignatureNonce() string {
	r := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, r); err != nil {
		return ""
	}

	md5Bytes := Md5(r)
	timeBytes, _ := time.Now().GobEncode()
	onceBytes := append(md5Bytes, timeBytes...)

	return base64.URLEncoding.EncodeToString(onceBytes)

}
