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

package ons

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-ons/20151020/OnsConsumerAccumulateResponse"

	"aliyun-go-sdk-ons/20151020/OnsClusterNamesResponse"

	"aliyun-go-sdk-ons/20151020/OnsClusterListResponse"

	"aliyun-go-sdk-ons/20151020/OnsCloudGetAppkeyListResponse"

	"aliyun-go-sdk-ons/20151020/OnsEmpowerListResponse"

	"aliyun-go-sdk-ons/20151020/OnsEmpowerDeleteResponse"

	"aliyun-go-sdk-ons/20151020/OnsEmpowerCreateResponse"

	"aliyun-go-sdk-ons/20151020/OnsConsumerTimeSpanResponse"

	"aliyun-go-sdk-ons/20151020/OnsConsumerStatusResponse"

	"aliyun-go-sdk-ons/20151020/OnsConsumerResetOffsetResponse"

	"aliyun-go-sdk-ons/20151020/OnsConsumerGetConnectionResponse"

	"aliyun-go-sdk-ons/20151020/OnsPublishDeleteResponse"

	"aliyun-go-sdk-ons/20151020/OnsPublishCreateResponse"

	"aliyun-go-sdk-ons/20151020/OnsMessageTraceResponse"

	"aliyun-go-sdk-ons/20151020/OnsMessageSendResponse"

	"aliyun-go-sdk-ons/20151020/OnsMessagePushResponse"

	"aliyun-go-sdk-ons/20151020/OnsMessageGetByTopicResponse"

	"aliyun-go-sdk-ons/20151020/OnsMessageGetByMsgIdResponse"

	"aliyun-go-sdk-ons/20151020/OnsMessageGetByKeyResponse"

	"aliyun-go-sdk-ons/20151020/OnsSubscriptionListResponse"

	"aliyun-go-sdk-ons/20151020/OnsSubscriptionGetResponse"

	"aliyun-go-sdk-ons/20151020/OnsSubscriptionDeleteResponse"

	"aliyun-go-sdk-ons/20151020/OnsSubscriptionCreateResponse"

	"aliyun-go-sdk-ons/20151020/OnsRegionListResponse"

	"aliyun-go-sdk-ons/20151020/OnsPublishSearchResponse"

	"aliyun-go-sdk-ons/20151020/OnsPublishListResponse"

	"aliyun-go-sdk-ons/20151020/OnsPublishGetResponse"

	"aliyun-go-sdk-ons/20151020/OnsTopicStatusResponse"

	"aliyun-go-sdk-ons/20151020/OnsTopicSearchResponse"

	"aliyun-go-sdk-ons/20151020/OnsTopicListResponse"

	"aliyun-go-sdk-ons/20151020/OnsTopicGetResponse"

	"aliyun-go-sdk-ons/20151020/OnsTopicDeleteResponse"

	"aliyun-go-sdk-ons/20151020/OnsTopicCreateResponse"

	"aliyun-go-sdk-ons/20151020/OnsSubscriptionSearchResponse"

	"aliyun-go-sdk-ons/20151020/OnsWarnCreateResponse"

	"aliyun-go-sdk-ons/20151020/OnsWarnAdminResponse"

	"aliyun-go-sdk-ons/20151020/OnsTrendTopicInputTpsResponse"

	"aliyun-go-sdk-ons/20151020/OnsTrendGroupOutputTpsResponse"

	"aliyun-go-sdk-ons/20151020/OnsTrendGetMachineSarResponse"

	"aliyun-go-sdk-ons/20151020/OnsTrendClusterOutputTpsResponse"

	"aliyun-go-sdk-ons/20151020/OnsTrendClusterInputTpsResponse"

	"aliyun-go-sdk-ons/20151020/OnsWarnListResponse"

	"aliyun-go-sdk-ons/20151020/OnsWarnEnableResponse"

	"aliyun-go-sdk-ons/20151020/OnsWarnEditorResponse"

	"aliyun-go-sdk-ons/20151020/OnsWarnDisableResponse"

	"aliyun-go-sdk-ons/20151020/OnsWarnDeleteResponse"
)

type OnsClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *OnsClient {
	p := &OnsClient{}
	p.Version = "2015-10-20"
	p.ProductName = "Ons"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *OnsClient) OnsConsumerAccumulate(onsConsumerAccumulateRequest *OnsConsumerAccumulateRequest) (ons_OnsConsumerAccumulateResponse.OnsConsumerAccumulateResponse, error) {
	var resObj ons_OnsConsumerAccumulateResponse.OnsConsumerAccumulateResponse

	if onsConsumerAccumulateRequest == nil {
		onsConsumerAccumulateRequest = new(OnsConsumerAccumulateRequest)
	}
	err := c.DoAction(onsConsumerAccumulateRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsClusterNames(onsClusterNamesRequest *OnsClusterNamesRequest) (ons_OnsClusterNamesResponse.OnsClusterNamesResponse, error) {
	var resObj ons_OnsClusterNamesResponse.OnsClusterNamesResponse

	if onsClusterNamesRequest == nil {
		onsClusterNamesRequest = new(OnsClusterNamesRequest)
	}
	err := c.DoAction(onsClusterNamesRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsClusterList(onsClusterListRequest *OnsClusterListRequest) (ons_OnsClusterListResponse.OnsClusterListResponse, error) {
	var resObj ons_OnsClusterListResponse.OnsClusterListResponse

	if onsClusterListRequest == nil {
		onsClusterListRequest = new(OnsClusterListRequest)
	}
	err := c.DoAction(onsClusterListRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsCloudGetAppkeyList(onsCloudGetAppkeyListRequest *OnsCloudGetAppkeyListRequest) (ons_OnsCloudGetAppkeyListResponse.OnsCloudGetAppkeyListResponse, error) {
	var resObj ons_OnsCloudGetAppkeyListResponse.OnsCloudGetAppkeyListResponse

	if onsCloudGetAppkeyListRequest == nil {
		onsCloudGetAppkeyListRequest = new(OnsCloudGetAppkeyListRequest)
	}
	err := c.DoAction(onsCloudGetAppkeyListRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsEmpowerList(onsEmpowerListRequest *OnsEmpowerListRequest) (ons_OnsEmpowerListResponse.OnsEmpowerListResponse, error) {
	var resObj ons_OnsEmpowerListResponse.OnsEmpowerListResponse

	if onsEmpowerListRequest == nil {
		onsEmpowerListRequest = new(OnsEmpowerListRequest)
	}
	err := c.DoAction(onsEmpowerListRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsEmpowerDelete(onsEmpowerDeleteRequest *OnsEmpowerDeleteRequest) (ons_OnsEmpowerDeleteResponse.OnsEmpowerDeleteResponse, error) {
	var resObj ons_OnsEmpowerDeleteResponse.OnsEmpowerDeleteResponse

	if onsEmpowerDeleteRequest == nil {
		onsEmpowerDeleteRequest = new(OnsEmpowerDeleteRequest)
	}
	err := c.DoAction(onsEmpowerDeleteRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsEmpowerCreate(onsEmpowerCreateRequest *OnsEmpowerCreateRequest) (ons_OnsEmpowerCreateResponse.OnsEmpowerCreateResponse, error) {
	var resObj ons_OnsEmpowerCreateResponse.OnsEmpowerCreateResponse

	if onsEmpowerCreateRequest == nil {
		onsEmpowerCreateRequest = new(OnsEmpowerCreateRequest)
	}
	err := c.DoAction(onsEmpowerCreateRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsConsumerTimeSpan(onsConsumerTimeSpanRequest *OnsConsumerTimeSpanRequest) (ons_OnsConsumerTimeSpanResponse.OnsConsumerTimeSpanResponse, error) {
	var resObj ons_OnsConsumerTimeSpanResponse.OnsConsumerTimeSpanResponse

	if onsConsumerTimeSpanRequest == nil {
		onsConsumerTimeSpanRequest = new(OnsConsumerTimeSpanRequest)
	}
	err := c.DoAction(onsConsumerTimeSpanRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsConsumerStatus(onsConsumerStatusRequest *OnsConsumerStatusRequest) (ons_OnsConsumerStatusResponse.OnsConsumerStatusResponse, error) {
	var resObj ons_OnsConsumerStatusResponse.OnsConsumerStatusResponse

	if onsConsumerStatusRequest == nil {
		onsConsumerStatusRequest = new(OnsConsumerStatusRequest)
	}
	err := c.DoAction(onsConsumerStatusRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsConsumerResetOffset(onsConsumerResetOffsetRequest *OnsConsumerResetOffsetRequest) (ons_OnsConsumerResetOffsetResponse.OnsConsumerResetOffsetResponse, error) {
	var resObj ons_OnsConsumerResetOffsetResponse.OnsConsumerResetOffsetResponse

	if onsConsumerResetOffsetRequest == nil {
		onsConsumerResetOffsetRequest = new(OnsConsumerResetOffsetRequest)
	}
	err := c.DoAction(onsConsumerResetOffsetRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsConsumerGetConnection(onsConsumerGetConnectionRequest *OnsConsumerGetConnectionRequest) (ons_OnsConsumerGetConnectionResponse.OnsConsumerGetConnectionResponse, error) {
	var resObj ons_OnsConsumerGetConnectionResponse.OnsConsumerGetConnectionResponse

	if onsConsumerGetConnectionRequest == nil {
		onsConsumerGetConnectionRequest = new(OnsConsumerGetConnectionRequest)
	}
	err := c.DoAction(onsConsumerGetConnectionRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsPublishDelete(onsPublishDeleteRequest *OnsPublishDeleteRequest) (ons_OnsPublishDeleteResponse.OnsPublishDeleteResponse, error) {
	var resObj ons_OnsPublishDeleteResponse.OnsPublishDeleteResponse

	if onsPublishDeleteRequest == nil {
		onsPublishDeleteRequest = new(OnsPublishDeleteRequest)
	}
	err := c.DoAction(onsPublishDeleteRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsPublishCreate(onsPublishCreateRequest *OnsPublishCreateRequest) (ons_OnsPublishCreateResponse.OnsPublishCreateResponse, error) {
	var resObj ons_OnsPublishCreateResponse.OnsPublishCreateResponse

	if onsPublishCreateRequest == nil {
		onsPublishCreateRequest = new(OnsPublishCreateRequest)
	}
	err := c.DoAction(onsPublishCreateRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsMessageTrace(onsMessageTraceRequest *OnsMessageTraceRequest) (ons_OnsMessageTraceResponse.OnsMessageTraceResponse, error) {
	var resObj ons_OnsMessageTraceResponse.OnsMessageTraceResponse

	if onsMessageTraceRequest == nil {
		onsMessageTraceRequest = new(OnsMessageTraceRequest)
	}
	err := c.DoAction(onsMessageTraceRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsMessageSend(onsMessageSendRequest *OnsMessageSendRequest) (ons_OnsMessageSendResponse.OnsMessageSendResponse, error) {
	var resObj ons_OnsMessageSendResponse.OnsMessageSendResponse

	if onsMessageSendRequest == nil {
		onsMessageSendRequest = new(OnsMessageSendRequest)
	}
	err := c.DoAction(onsMessageSendRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsMessagePush(onsMessagePushRequest *OnsMessagePushRequest) (ons_OnsMessagePushResponse.OnsMessagePushResponse, error) {
	var resObj ons_OnsMessagePushResponse.OnsMessagePushResponse

	if onsMessagePushRequest == nil {
		onsMessagePushRequest = new(OnsMessagePushRequest)
	}
	err := c.DoAction(onsMessagePushRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsMessageGetByTopic(onsMessageGetByTopicRequest *OnsMessageGetByTopicRequest) (ons_OnsMessageGetByTopicResponse.OnsMessageGetByTopicResponse, error) {
	var resObj ons_OnsMessageGetByTopicResponse.OnsMessageGetByTopicResponse

	if onsMessageGetByTopicRequest == nil {
		onsMessageGetByTopicRequest = new(OnsMessageGetByTopicRequest)
	}
	err := c.DoAction(onsMessageGetByTopicRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsMessageGetByMsgId(onsMessageGetByMsgIdRequest *OnsMessageGetByMsgIdRequest) (ons_OnsMessageGetByMsgIdResponse.OnsMessageGetByMsgIdResponse, error) {
	var resObj ons_OnsMessageGetByMsgIdResponse.OnsMessageGetByMsgIdResponse

	if onsMessageGetByMsgIdRequest == nil {
		onsMessageGetByMsgIdRequest = new(OnsMessageGetByMsgIdRequest)
	}
	err := c.DoAction(onsMessageGetByMsgIdRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsMessageGetByKey(onsMessageGetByKeyRequest *OnsMessageGetByKeyRequest) (ons_OnsMessageGetByKeyResponse.OnsMessageGetByKeyResponse, error) {
	var resObj ons_OnsMessageGetByKeyResponse.OnsMessageGetByKeyResponse

	if onsMessageGetByKeyRequest == nil {
		onsMessageGetByKeyRequest = new(OnsMessageGetByKeyRequest)
	}
	err := c.DoAction(onsMessageGetByKeyRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsSubscriptionList(onsSubscriptionListRequest *OnsSubscriptionListRequest) (ons_OnsSubscriptionListResponse.OnsSubscriptionListResponse, error) {
	var resObj ons_OnsSubscriptionListResponse.OnsSubscriptionListResponse

	if onsSubscriptionListRequest == nil {
		onsSubscriptionListRequest = new(OnsSubscriptionListRequest)
	}
	err := c.DoAction(onsSubscriptionListRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsSubscriptionGet(onsSubscriptionGetRequest *OnsSubscriptionGetRequest) (ons_OnsSubscriptionGetResponse.OnsSubscriptionGetResponse, error) {
	var resObj ons_OnsSubscriptionGetResponse.OnsSubscriptionGetResponse

	if onsSubscriptionGetRequest == nil {
		onsSubscriptionGetRequest = new(OnsSubscriptionGetRequest)
	}
	err := c.DoAction(onsSubscriptionGetRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsSubscriptionDelete(onsSubscriptionDeleteRequest *OnsSubscriptionDeleteRequest) (ons_OnsSubscriptionDeleteResponse.OnsSubscriptionDeleteResponse, error) {
	var resObj ons_OnsSubscriptionDeleteResponse.OnsSubscriptionDeleteResponse

	if onsSubscriptionDeleteRequest == nil {
		onsSubscriptionDeleteRequest = new(OnsSubscriptionDeleteRequest)
	}
	err := c.DoAction(onsSubscriptionDeleteRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsSubscriptionCreate(onsSubscriptionCreateRequest *OnsSubscriptionCreateRequest) (ons_OnsSubscriptionCreateResponse.OnsSubscriptionCreateResponse, error) {
	var resObj ons_OnsSubscriptionCreateResponse.OnsSubscriptionCreateResponse

	if onsSubscriptionCreateRequest == nil {
		onsSubscriptionCreateRequest = new(OnsSubscriptionCreateRequest)
	}
	err := c.DoAction(onsSubscriptionCreateRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsRegionList(onsRegionListRequest *OnsRegionListRequest) (ons_OnsRegionListResponse.OnsRegionListResponse, error) {
	var resObj ons_OnsRegionListResponse.OnsRegionListResponse

	if onsRegionListRequest == nil {
		onsRegionListRequest = new(OnsRegionListRequest)
	}
	err := c.DoAction(onsRegionListRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsPublishSearch(onsPublishSearchRequest *OnsPublishSearchRequest) (ons_OnsPublishSearchResponse.OnsPublishSearchResponse, error) {
	var resObj ons_OnsPublishSearchResponse.OnsPublishSearchResponse

	if onsPublishSearchRequest == nil {
		onsPublishSearchRequest = new(OnsPublishSearchRequest)
	}
	err := c.DoAction(onsPublishSearchRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsPublishList(onsPublishListRequest *OnsPublishListRequest) (ons_OnsPublishListResponse.OnsPublishListResponse, error) {
	var resObj ons_OnsPublishListResponse.OnsPublishListResponse

	if onsPublishListRequest == nil {
		onsPublishListRequest = new(OnsPublishListRequest)
	}
	err := c.DoAction(onsPublishListRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsPublishGet(onsPublishGetRequest *OnsPublishGetRequest) (ons_OnsPublishGetResponse.OnsPublishGetResponse, error) {
	var resObj ons_OnsPublishGetResponse.OnsPublishGetResponse

	if onsPublishGetRequest == nil {
		onsPublishGetRequest = new(OnsPublishGetRequest)
	}
	err := c.DoAction(onsPublishGetRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsTopicStatus(onsTopicStatusRequest *OnsTopicStatusRequest) (ons_OnsTopicStatusResponse.OnsTopicStatusResponse, error) {
	var resObj ons_OnsTopicStatusResponse.OnsTopicStatusResponse

	if onsTopicStatusRequest == nil {
		onsTopicStatusRequest = new(OnsTopicStatusRequest)
	}
	err := c.DoAction(onsTopicStatusRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsTopicSearch(onsTopicSearchRequest *OnsTopicSearchRequest) (ons_OnsTopicSearchResponse.OnsTopicSearchResponse, error) {
	var resObj ons_OnsTopicSearchResponse.OnsTopicSearchResponse

	if onsTopicSearchRequest == nil {
		onsTopicSearchRequest = new(OnsTopicSearchRequest)
	}
	err := c.DoAction(onsTopicSearchRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsTopicList(onsTopicListRequest *OnsTopicListRequest) (ons_OnsTopicListResponse.OnsTopicListResponse, error) {
	var resObj ons_OnsTopicListResponse.OnsTopicListResponse

	if onsTopicListRequest == nil {
		onsTopicListRequest = new(OnsTopicListRequest)
	}
	err := c.DoAction(onsTopicListRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsTopicGet(onsTopicGetRequest *OnsTopicGetRequest) (ons_OnsTopicGetResponse.OnsTopicGetResponse, error) {
	var resObj ons_OnsTopicGetResponse.OnsTopicGetResponse

	if onsTopicGetRequest == nil {
		onsTopicGetRequest = new(OnsTopicGetRequest)
	}
	err := c.DoAction(onsTopicGetRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsTopicDelete(onsTopicDeleteRequest *OnsTopicDeleteRequest) (ons_OnsTopicDeleteResponse.OnsTopicDeleteResponse, error) {
	var resObj ons_OnsTopicDeleteResponse.OnsTopicDeleteResponse

	if onsTopicDeleteRequest == nil {
		onsTopicDeleteRequest = new(OnsTopicDeleteRequest)
	}
	err := c.DoAction(onsTopicDeleteRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsTopicCreate(onsTopicCreateRequest *OnsTopicCreateRequest) (ons_OnsTopicCreateResponse.OnsTopicCreateResponse, error) {
	var resObj ons_OnsTopicCreateResponse.OnsTopicCreateResponse

	if onsTopicCreateRequest == nil {
		onsTopicCreateRequest = new(OnsTopicCreateRequest)
	}
	err := c.DoAction(onsTopicCreateRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsSubscriptionSearch(onsSubscriptionSearchRequest *OnsSubscriptionSearchRequest) (ons_OnsSubscriptionSearchResponse.OnsSubscriptionSearchResponse, error) {
	var resObj ons_OnsSubscriptionSearchResponse.OnsSubscriptionSearchResponse

	if onsSubscriptionSearchRequest == nil {
		onsSubscriptionSearchRequest = new(OnsSubscriptionSearchRequest)
	}
	err := c.DoAction(onsSubscriptionSearchRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsWarnCreate(onsWarnCreateRequest *OnsWarnCreateRequest) (ons_OnsWarnCreateResponse.OnsWarnCreateResponse, error) {
	var resObj ons_OnsWarnCreateResponse.OnsWarnCreateResponse

	if onsWarnCreateRequest == nil {
		onsWarnCreateRequest = new(OnsWarnCreateRequest)
	}
	err := c.DoAction(onsWarnCreateRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsWarnAdmin(onsWarnAdminRequest *OnsWarnAdminRequest) (ons_OnsWarnAdminResponse.OnsWarnAdminResponse, error) {
	var resObj ons_OnsWarnAdminResponse.OnsWarnAdminResponse

	if onsWarnAdminRequest == nil {
		onsWarnAdminRequest = new(OnsWarnAdminRequest)
	}
	err := c.DoAction(onsWarnAdminRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsTrendTopicInputTps(onsTrendTopicInputTpsRequest *OnsTrendTopicInputTpsRequest) (ons_OnsTrendTopicInputTpsResponse.OnsTrendTopicInputTpsResponse, error) {
	var resObj ons_OnsTrendTopicInputTpsResponse.OnsTrendTopicInputTpsResponse

	if onsTrendTopicInputTpsRequest == nil {
		onsTrendTopicInputTpsRequest = new(OnsTrendTopicInputTpsRequest)
	}
	err := c.DoAction(onsTrendTopicInputTpsRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsTrendGroupOutputTps(onsTrendGroupOutputTpsRequest *OnsTrendGroupOutputTpsRequest) (ons_OnsTrendGroupOutputTpsResponse.OnsTrendGroupOutputTpsResponse, error) {
	var resObj ons_OnsTrendGroupOutputTpsResponse.OnsTrendGroupOutputTpsResponse

	if onsTrendGroupOutputTpsRequest == nil {
		onsTrendGroupOutputTpsRequest = new(OnsTrendGroupOutputTpsRequest)
	}
	err := c.DoAction(onsTrendGroupOutputTpsRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsTrendGetMachineSar(onsTrendGetMachineSarRequest *OnsTrendGetMachineSarRequest) (ons_OnsTrendGetMachineSarResponse.OnsTrendGetMachineSarResponse, error) {
	var resObj ons_OnsTrendGetMachineSarResponse.OnsTrendGetMachineSarResponse

	if onsTrendGetMachineSarRequest == nil {
		onsTrendGetMachineSarRequest = new(OnsTrendGetMachineSarRequest)
	}
	err := c.DoAction(onsTrendGetMachineSarRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsTrendClusterOutputTps(onsTrendClusterOutputTpsRequest *OnsTrendClusterOutputTpsRequest) (ons_OnsTrendClusterOutputTpsResponse.OnsTrendClusterOutputTpsResponse, error) {
	var resObj ons_OnsTrendClusterOutputTpsResponse.OnsTrendClusterOutputTpsResponse

	if onsTrendClusterOutputTpsRequest == nil {
		onsTrendClusterOutputTpsRequest = new(OnsTrendClusterOutputTpsRequest)
	}
	err := c.DoAction(onsTrendClusterOutputTpsRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsTrendClusterInputTps(onsTrendClusterInputTpsRequest *OnsTrendClusterInputTpsRequest) (ons_OnsTrendClusterInputTpsResponse.OnsTrendClusterInputTpsResponse, error) {
	var resObj ons_OnsTrendClusterInputTpsResponse.OnsTrendClusterInputTpsResponse

	if onsTrendClusterInputTpsRequest == nil {
		onsTrendClusterInputTpsRequest = new(OnsTrendClusterInputTpsRequest)
	}
	err := c.DoAction(onsTrendClusterInputTpsRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsWarnList(onsWarnListRequest *OnsWarnListRequest) (ons_OnsWarnListResponse.OnsWarnListResponse, error) {
	var resObj ons_OnsWarnListResponse.OnsWarnListResponse

	if onsWarnListRequest == nil {
		onsWarnListRequest = new(OnsWarnListRequest)
	}
	err := c.DoAction(onsWarnListRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsWarnEnable(onsWarnEnableRequest *OnsWarnEnableRequest) (ons_OnsWarnEnableResponse.OnsWarnEnableResponse, error) {
	var resObj ons_OnsWarnEnableResponse.OnsWarnEnableResponse

	if onsWarnEnableRequest == nil {
		onsWarnEnableRequest = new(OnsWarnEnableRequest)
	}
	err := c.DoAction(onsWarnEnableRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsWarnEditor(onsWarnEditorRequest *OnsWarnEditorRequest) (ons_OnsWarnEditorResponse.OnsWarnEditorResponse, error) {
	var resObj ons_OnsWarnEditorResponse.OnsWarnEditorResponse

	if onsWarnEditorRequest == nil {
		onsWarnEditorRequest = new(OnsWarnEditorRequest)
	}
	err := c.DoAction(onsWarnEditorRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsWarnDisable(onsWarnDisableRequest *OnsWarnDisableRequest) (ons_OnsWarnDisableResponse.OnsWarnDisableResponse, error) {
	var resObj ons_OnsWarnDisableResponse.OnsWarnDisableResponse

	if onsWarnDisableRequest == nil {
		onsWarnDisableRequest = new(OnsWarnDisableRequest)
	}
	err := c.DoAction(onsWarnDisableRequest, &resObj)
	return resObj, err
}

func (c *OnsClient) OnsWarnDelete(onsWarnDeleteRequest *OnsWarnDeleteRequest) (ons_OnsWarnDeleteResponse.OnsWarnDeleteResponse, error) {
	var resObj ons_OnsWarnDeleteResponse.OnsWarnDeleteResponse

	if onsWarnDeleteRequest == nil {
		onsWarnDeleteRequest = new(OnsWarnDeleteRequest)
	}
	err := c.DoAction(onsWarnDeleteRequest, &resObj)
	return resObj, err
}
