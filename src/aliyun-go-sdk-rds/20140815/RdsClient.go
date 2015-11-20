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

package rds

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-rds/20140815/DescribeDBInstanceIPArrayListResponse"

	"aliyun-go-sdk-rds/20140815/DescribeSQLLogReportListResponse"

	"aliyun-go-sdk-rds/20140815/ResetAccountForPGResponse"

	"aliyun-go-sdk-rds/20140815/AllocateInstancePrivateConnectionResponse"

	"aliyun-go-sdk-rds/20140815/UnlockDBInstanceResponse"

	"aliyun-go-sdk-rds/20140815/LockDBInstanceResponse"

	"aliyun-go-sdk-rds/20140815/CreateDatabaseForInnerResponse"

	"aliyun-go-sdk-rds/20140815/CreateAccountForInnerResponse"

	"aliyun-go-sdk-rds/20140815/UpgradeDBInstanceEngineVersionResponse"

	"aliyun-go-sdk-rds/20140815/StopSyncingResponse"

	"aliyun-go-sdk-rds/20140815/RevokeAccountPrivilegeResponse"

	"aliyun-go-sdk-rds/20140815/RestoreDBInstanceResponse"

	"aliyun-go-sdk-rds/20140815/RestartDBInstanceResponse"

	"aliyun-go-sdk-rds/20140815/ResetAccountPasswordResponse"

	"aliyun-go-sdk-rds/20140815/RemoveTagsFromResourceResponse"

	"aliyun-go-sdk-rds/20140815/PurgeDBInstanceLogResponse"

	"aliyun-go-sdk-rds/20140815/PreCheckBeforeImportDataResponse"

	"aliyun-go-sdk-rds/20140815/ModifySecurityIpsResponse"

	"aliyun-go-sdk-rds/20140815/ModifyPostpaidDBInstanceSpecResponse"

	"aliyun-go-sdk-rds/20140815/ModifyParameterResponse"

	"aliyun-go-sdk-rds/20140815/ModifyDBInstanceSpecResponse"

	"aliyun-go-sdk-rds/20140815/ModifyDBInstanceMaintainTimeResponse"

	"aliyun-go-sdk-rds/20140815/ModifyDBInstanceDescriptionResponse"

	"aliyun-go-sdk-rds/20140815/ModifyDBDescriptionResponse"

	"aliyun-go-sdk-rds/20140815/ModifyBackupPolicyResponse"

	"aliyun-go-sdk-rds/20140815/ModifyAccountDescriptionResponse"

	"aliyun-go-sdk-rds/20140815/MigrateToOtherZoneResponse"

	"aliyun-go-sdk-rds/20140815/ImportDataFromDatabaseResponse"

	"aliyun-go-sdk-rds/20140815/ImportDataForSQLServerResponse"

	"aliyun-go-sdk-rds/20140815/ImportDatabaseBetweenInstancesResponse"

	"aliyun-go-sdk-rds/20140815/GrantAccountPrivilegeResponse"

	"aliyun-go-sdk-rds/20140815/ExtractBackupFromOASResponse"

	"aliyun-go-sdk-rds/20140815/DescribeVpcZoneNosResponse"

	"aliyun-go-sdk-rds/20140815/DescribeTasksResponse"

	"aliyun-go-sdk-rds/20140815/DescribeSQLLogReportsResponse"

	"aliyun-go-sdk-rds/20140815/DescribeSQLLogRecordsResponse"

	"aliyun-go-sdk-rds/20140815/DescribeSQLInjectionInfosResponse"

	"aliyun-go-sdk-rds/20140815/DescribeSlowLogsResponse"

	"aliyun-go-sdk-rds/20140815/DescribeSlowLogRecordsResponse"

	"aliyun-go-sdk-rds/20140815/DescribeResourceUsageResponse"

	"aliyun-go-sdk-rds/20140815/DescribeRegionsResponse"

	"aliyun-go-sdk-rds/20140815/DescribeRealtimeDiagnosesResponse"

	"aliyun-go-sdk-rds/20140815/DescribePreCheckResultsResponse"

	"aliyun-go-sdk-rds/20140815/DescribeParameterTemplatesResponse"

	"aliyun-go-sdk-rds/20140815/DescribeParametersResponse"

	"aliyun-go-sdk-rds/20140815/DescribeOptimizeAdviceOnStorageResponse"

	"aliyun-go-sdk-rds/20140815/DescribeOptimizeAdviceOnMissPKResponse"

	"aliyun-go-sdk-rds/20140815/DescribeOptimizeAdviceOnMissIndexResponse"

	"aliyun-go-sdk-rds/20140815/DescribeOptimizeAdviceOnExcessIndexResponse"

	"aliyun-go-sdk-rds/20140815/DescribeOptimizeAdviceOnBigTableResponse"

	"aliyun-go-sdk-rds/20140815/DescribeOptimizeAdviceByDBAResponse"

	"aliyun-go-sdk-rds/20140815/DescribeOperationLogsResponse"

	"aliyun-go-sdk-rds/20140815/DescribeModifyParameterLogResponse"

	"aliyun-go-sdk-rds/20140815/DescribeImportsForSQLServerResponse"

	"aliyun-go-sdk-rds/20140815/DescribeFilesForSQLServerResponse"

	"aliyun-go-sdk-rds/20140815/DescribeErrorLogsResponse"

	"aliyun-go-sdk-rds/20140815/DescribeDBInstancePerformanceResponse"

	"aliyun-go-sdk-rds/20140815/DescribeDatabasesResponse"

	"aliyun-go-sdk-rds/20140815/DescribeBinlogFilesResponse"

	"aliyun-go-sdk-rds/20140815/DescribeBackupTasksResponse"

	"aliyun-go-sdk-rds/20140815/DescribeBackupsResponse"

	"aliyun-go-sdk-rds/20140815/DescribeBackupPolicyResponse"

	"aliyun-go-sdk-rds/20140815/DescribeAccountsResponse"

	"aliyun-go-sdk-rds/20140815/DescibeImportsFromDatabaseResponse"

	"aliyun-go-sdk-rds/20140815/DeleteDBInstanceResponse"

	"aliyun-go-sdk-rds/20140815/DeleteDatabaseResponse"

	"aliyun-go-sdk-rds/20140815/DeleteAccountResponse"

	"aliyun-go-sdk-rds/20140815/CreateUploadPathForSQLServerResponse"

	"aliyun-go-sdk-rds/20140815/CreateTempDBInstanceResponse"

	"aliyun-go-sdk-rds/20140815/CreatePostpaidDBInstanceForChannelResponse"

	"aliyun-go-sdk-rds/20140815/CreatePostpaidDBInstanceResponse"

	"aliyun-go-sdk-rds/20140815/CreateDBInstanceforFirstPayResponse"

	"aliyun-go-sdk-rds/20140815/CreateDBInstanceForChannelResponse"

	"aliyun-go-sdk-rds/20140815/CreateDatabaseResponse"

	"aliyun-go-sdk-rds/20140815/CreateBackupResponse"

	"aliyun-go-sdk-rds/20140815/CreateAccountResponse"

	"aliyun-go-sdk-rds/20140815/CheckDBNameAvailableResponse"

	"aliyun-go-sdk-rds/20140815/CheckAccountNameAvailableResponse"

	"aliyun-go-sdk-rds/20140815/CancelImportResponse"

	"aliyun-go-sdk-rds/20140815/BatchRevokeAccountPrivilegeResponse"

	"aliyun-go-sdk-rds/20140815/BatchGrantAccountPrivilegeResponse"

	"aliyun-go-sdk-rds/20140815/AddTagsToResourceResponse"

	"aliyun-go-sdk-rds/20140815/SwitchDBInstanceNetTypeResponse"

	"aliyun-go-sdk-rds/20140815/ReleaseInstancePublicConnectionResponse"

	"aliyun-go-sdk-rds/20140815/ModifyDBInstanceNetworkTypeResponse"

	"aliyun-go-sdk-rds/20140815/ModifyDBInstanceConnectionStringResponse"

	"aliyun-go-sdk-rds/20140815/ModifyDBInstanceConnectionModeResponse"

	"aliyun-go-sdk-rds/20140815/DescribeDBInstanceNetInfoForChannelResponse"

	"aliyun-go-sdk-rds/20140815/DescribeDBInstanceNetInfoResponse"

	"aliyun-go-sdk-rds/20140815/CreateReadOnlyDBInstanceResponse"

	"aliyun-go-sdk-rds/20140815/CreateDBInstanceResponse"

	"aliyun-go-sdk-rds/20140815/AllocateInstancePublicConnectionResponse"

	"aliyun-go-sdk-rds/20140815/DescribeDBInstancesByPerformanceResponse"

	"aliyun-go-sdk-rds/20140815/DescribeDBInstancesByExpireTimeResponse"

	"aliyun-go-sdk-rds/20140815/DescribeDBInstancesResponse"

	"aliyun-go-sdk-rds/20140815/DescribeDBInstanceAttributeResponse"

	"aliyun-go-sdk-rds/20140815/DescribeAbnormalDBInstancesResponse"

	"aliyun-go-sdk-rds/20140815/StartDBInstanceDiagnoseResponse"
)

type RdsClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *RdsClient {
	p := &RdsClient{}
	p.Version = "2014-08-15"
	p.ProductName = "Rds"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *RdsClient) DescribeDBInstanceIPArrayList(describeDBInstanceIPArrayListRequest *DescribeDBInstanceIPArrayListRequest) (rds_DescribeDBInstanceIPArrayListResponse.DescribeDBInstanceIPArrayListResponse, error) {
	var resObj rds_DescribeDBInstanceIPArrayListResponse.DescribeDBInstanceIPArrayListResponse

	if describeDBInstanceIPArrayListRequest == nil {
		describeDBInstanceIPArrayListRequest = new(DescribeDBInstanceIPArrayListRequest)
	}
	err := c.DoAction(describeDBInstanceIPArrayListRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeSQLLogReportList(describeSQLLogReportListRequest *DescribeSQLLogReportListRequest) (rds_DescribeSQLLogReportListResponse.DescribeSQLLogReportListResponse, error) {
	var resObj rds_DescribeSQLLogReportListResponse.DescribeSQLLogReportListResponse

	if describeSQLLogReportListRequest == nil {
		describeSQLLogReportListRequest = new(DescribeSQLLogReportListRequest)
	}
	err := c.DoAction(describeSQLLogReportListRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) ResetAccountForPG(resetAccountForPGRequest *ResetAccountForPGRequest) (rds_ResetAccountForPGResponse.ResetAccountForPGResponse, error) {
	var resObj rds_ResetAccountForPGResponse.ResetAccountForPGResponse

	if resetAccountForPGRequest == nil {
		resetAccountForPGRequest = new(ResetAccountForPGRequest)
	}
	err := c.DoAction(resetAccountForPGRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) AllocateInstancePrivateConnection(allocateInstancePrivateConnectionRequest *AllocateInstancePrivateConnectionRequest) (rds_AllocateInstancePrivateConnectionResponse.AllocateInstancePrivateConnectionResponse, error) {
	var resObj rds_AllocateInstancePrivateConnectionResponse.AllocateInstancePrivateConnectionResponse

	if allocateInstancePrivateConnectionRequest == nil {
		allocateInstancePrivateConnectionRequest = new(AllocateInstancePrivateConnectionRequest)
	}
	err := c.DoAction(allocateInstancePrivateConnectionRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) UnlockDBInstance(unlockDBInstanceRequest *UnlockDBInstanceRequest) (rds_UnlockDBInstanceResponse.UnlockDBInstanceResponse, error) {
	var resObj rds_UnlockDBInstanceResponse.UnlockDBInstanceResponse

	if unlockDBInstanceRequest == nil {
		unlockDBInstanceRequest = new(UnlockDBInstanceRequest)
	}
	err := c.DoAction(unlockDBInstanceRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) LockDBInstance(lockDBInstanceRequest *LockDBInstanceRequest) (rds_LockDBInstanceResponse.LockDBInstanceResponse, error) {
	var resObj rds_LockDBInstanceResponse.LockDBInstanceResponse

	if lockDBInstanceRequest == nil {
		lockDBInstanceRequest = new(LockDBInstanceRequest)
	}
	err := c.DoAction(lockDBInstanceRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) CreateDatabaseForInner(createDatabaseForInnerRequest *CreateDatabaseForInnerRequest) (rds_CreateDatabaseForInnerResponse.CreateDatabaseForInnerResponse, error) {
	var resObj rds_CreateDatabaseForInnerResponse.CreateDatabaseForInnerResponse

	if createDatabaseForInnerRequest == nil {
		createDatabaseForInnerRequest = new(CreateDatabaseForInnerRequest)
	}
	err := c.DoAction(createDatabaseForInnerRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) CreateAccountForInner(createAccountForInnerRequest *CreateAccountForInnerRequest) (rds_CreateAccountForInnerResponse.CreateAccountForInnerResponse, error) {
	var resObj rds_CreateAccountForInnerResponse.CreateAccountForInnerResponse

	if createAccountForInnerRequest == nil {
		createAccountForInnerRequest = new(CreateAccountForInnerRequest)
	}
	err := c.DoAction(createAccountForInnerRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) UpgradeDBInstanceEngineVersion(upgradeDBInstanceEngineVersionRequest *UpgradeDBInstanceEngineVersionRequest) (rds_UpgradeDBInstanceEngineVersionResponse.UpgradeDBInstanceEngineVersionResponse, error) {
	var resObj rds_UpgradeDBInstanceEngineVersionResponse.UpgradeDBInstanceEngineVersionResponse

	if upgradeDBInstanceEngineVersionRequest == nil {
		upgradeDBInstanceEngineVersionRequest = new(UpgradeDBInstanceEngineVersionRequest)
	}
	err := c.DoAction(upgradeDBInstanceEngineVersionRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) StopSyncing(stopSyncingRequest *StopSyncingRequest) (rds_StopSyncingResponse.StopSyncingResponse, error) {
	var resObj rds_StopSyncingResponse.StopSyncingResponse

	if stopSyncingRequest == nil {
		stopSyncingRequest = new(StopSyncingRequest)
	}
	err := c.DoAction(stopSyncingRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) RevokeAccountPrivilege(revokeAccountPrivilegeRequest *RevokeAccountPrivilegeRequest) (rds_RevokeAccountPrivilegeResponse.RevokeAccountPrivilegeResponse, error) {
	var resObj rds_RevokeAccountPrivilegeResponse.RevokeAccountPrivilegeResponse

	if revokeAccountPrivilegeRequest == nil {
		revokeAccountPrivilegeRequest = new(RevokeAccountPrivilegeRequest)
	}
	err := c.DoAction(revokeAccountPrivilegeRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) RestoreDBInstance(restoreDBInstanceRequest *RestoreDBInstanceRequest) (rds_RestoreDBInstanceResponse.RestoreDBInstanceResponse, error) {
	var resObj rds_RestoreDBInstanceResponse.RestoreDBInstanceResponse

	if restoreDBInstanceRequest == nil {
		restoreDBInstanceRequest = new(RestoreDBInstanceRequest)
	}
	err := c.DoAction(restoreDBInstanceRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) RestartDBInstance(restartDBInstanceRequest *RestartDBInstanceRequest) (rds_RestartDBInstanceResponse.RestartDBInstanceResponse, error) {
	var resObj rds_RestartDBInstanceResponse.RestartDBInstanceResponse

	if restartDBInstanceRequest == nil {
		restartDBInstanceRequest = new(RestartDBInstanceRequest)
	}
	err := c.DoAction(restartDBInstanceRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) ResetAccountPassword(resetAccountPasswordRequest *ResetAccountPasswordRequest) (rds_ResetAccountPasswordResponse.ResetAccountPasswordResponse, error) {
	var resObj rds_ResetAccountPasswordResponse.ResetAccountPasswordResponse

	if resetAccountPasswordRequest == nil {
		resetAccountPasswordRequest = new(ResetAccountPasswordRequest)
	}
	err := c.DoAction(resetAccountPasswordRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) RemoveTagsFromResource(removeTagsFromResourceRequest *RemoveTagsFromResourceRequest) (rds_RemoveTagsFromResourceResponse.RemoveTagsFromResourceResponse, error) {
	var resObj rds_RemoveTagsFromResourceResponse.RemoveTagsFromResourceResponse

	if removeTagsFromResourceRequest == nil {
		removeTagsFromResourceRequest = new(RemoveTagsFromResourceRequest)
	}
	err := c.DoAction(removeTagsFromResourceRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) PurgeDBInstanceLog(purgeDBInstanceLogRequest *PurgeDBInstanceLogRequest) (rds_PurgeDBInstanceLogResponse.PurgeDBInstanceLogResponse, error) {
	var resObj rds_PurgeDBInstanceLogResponse.PurgeDBInstanceLogResponse

	if purgeDBInstanceLogRequest == nil {
		purgeDBInstanceLogRequest = new(PurgeDBInstanceLogRequest)
	}
	err := c.DoAction(purgeDBInstanceLogRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) PreCheckBeforeImportData(preCheckBeforeImportDataRequest *PreCheckBeforeImportDataRequest) (rds_PreCheckBeforeImportDataResponse.PreCheckBeforeImportDataResponse, error) {
	var resObj rds_PreCheckBeforeImportDataResponse.PreCheckBeforeImportDataResponse

	if preCheckBeforeImportDataRequest == nil {
		preCheckBeforeImportDataRequest = new(PreCheckBeforeImportDataRequest)
	}
	err := c.DoAction(preCheckBeforeImportDataRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) ModifySecurityIps(modifySecurityIpsRequest *ModifySecurityIpsRequest) (rds_ModifySecurityIpsResponse.ModifySecurityIpsResponse, error) {
	var resObj rds_ModifySecurityIpsResponse.ModifySecurityIpsResponse

	if modifySecurityIpsRequest == nil {
		modifySecurityIpsRequest = new(ModifySecurityIpsRequest)
	}
	err := c.DoAction(modifySecurityIpsRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) ModifyPostpaidDBInstanceSpec(modifyPostpaidDBInstanceSpecRequest *ModifyPostpaidDBInstanceSpecRequest) (rds_ModifyPostpaidDBInstanceSpecResponse.ModifyPostpaidDBInstanceSpecResponse, error) {
	var resObj rds_ModifyPostpaidDBInstanceSpecResponse.ModifyPostpaidDBInstanceSpecResponse

	if modifyPostpaidDBInstanceSpecRequest == nil {
		modifyPostpaidDBInstanceSpecRequest = new(ModifyPostpaidDBInstanceSpecRequest)
	}
	err := c.DoAction(modifyPostpaidDBInstanceSpecRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) ModifyParameter(modifyParameterRequest *ModifyParameterRequest) (rds_ModifyParameterResponse.ModifyParameterResponse, error) {
	var resObj rds_ModifyParameterResponse.ModifyParameterResponse

	if modifyParameterRequest == nil {
		modifyParameterRequest = new(ModifyParameterRequest)
	}
	err := c.DoAction(modifyParameterRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) ModifyDBInstanceSpec(modifyDBInstanceSpecRequest *ModifyDBInstanceSpecRequest) (rds_ModifyDBInstanceSpecResponse.ModifyDBInstanceSpecResponse, error) {
	var resObj rds_ModifyDBInstanceSpecResponse.ModifyDBInstanceSpecResponse

	if modifyDBInstanceSpecRequest == nil {
		modifyDBInstanceSpecRequest = new(ModifyDBInstanceSpecRequest)
	}
	err := c.DoAction(modifyDBInstanceSpecRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) ModifyDBInstanceMaintainTime(modifyDBInstanceMaintainTimeRequest *ModifyDBInstanceMaintainTimeRequest) (rds_ModifyDBInstanceMaintainTimeResponse.ModifyDBInstanceMaintainTimeResponse, error) {
	var resObj rds_ModifyDBInstanceMaintainTimeResponse.ModifyDBInstanceMaintainTimeResponse

	if modifyDBInstanceMaintainTimeRequest == nil {
		modifyDBInstanceMaintainTimeRequest = new(ModifyDBInstanceMaintainTimeRequest)
	}
	err := c.DoAction(modifyDBInstanceMaintainTimeRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) ModifyDBInstanceDescription(modifyDBInstanceDescriptionRequest *ModifyDBInstanceDescriptionRequest) (rds_ModifyDBInstanceDescriptionResponse.ModifyDBInstanceDescriptionResponse, error) {
	var resObj rds_ModifyDBInstanceDescriptionResponse.ModifyDBInstanceDescriptionResponse

	if modifyDBInstanceDescriptionRequest == nil {
		modifyDBInstanceDescriptionRequest = new(ModifyDBInstanceDescriptionRequest)
	}
	err := c.DoAction(modifyDBInstanceDescriptionRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) ModifyDBDescription(modifyDBDescriptionRequest *ModifyDBDescriptionRequest) (rds_ModifyDBDescriptionResponse.ModifyDBDescriptionResponse, error) {
	var resObj rds_ModifyDBDescriptionResponse.ModifyDBDescriptionResponse

	if modifyDBDescriptionRequest == nil {
		modifyDBDescriptionRequest = new(ModifyDBDescriptionRequest)
	}
	err := c.DoAction(modifyDBDescriptionRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) ModifyBackupPolicy(modifyBackupPolicyRequest *ModifyBackupPolicyRequest) (rds_ModifyBackupPolicyResponse.ModifyBackupPolicyResponse, error) {
	var resObj rds_ModifyBackupPolicyResponse.ModifyBackupPolicyResponse

	if modifyBackupPolicyRequest == nil {
		modifyBackupPolicyRequest = new(ModifyBackupPolicyRequest)
	}
	err := c.DoAction(modifyBackupPolicyRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) ModifyAccountDescription(modifyAccountDescriptionRequest *ModifyAccountDescriptionRequest) (rds_ModifyAccountDescriptionResponse.ModifyAccountDescriptionResponse, error) {
	var resObj rds_ModifyAccountDescriptionResponse.ModifyAccountDescriptionResponse

	if modifyAccountDescriptionRequest == nil {
		modifyAccountDescriptionRequest = new(ModifyAccountDescriptionRequest)
	}
	err := c.DoAction(modifyAccountDescriptionRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) MigrateToOtherZone(migrateToOtherZoneRequest *MigrateToOtherZoneRequest) (rds_MigrateToOtherZoneResponse.MigrateToOtherZoneResponse, error) {
	var resObj rds_MigrateToOtherZoneResponse.MigrateToOtherZoneResponse

	if migrateToOtherZoneRequest == nil {
		migrateToOtherZoneRequest = new(MigrateToOtherZoneRequest)
	}
	err := c.DoAction(migrateToOtherZoneRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) ImportDataFromDatabase(importDataFromDatabaseRequest *ImportDataFromDatabaseRequest) (rds_ImportDataFromDatabaseResponse.ImportDataFromDatabaseResponse, error) {
	var resObj rds_ImportDataFromDatabaseResponse.ImportDataFromDatabaseResponse

	if importDataFromDatabaseRequest == nil {
		importDataFromDatabaseRequest = new(ImportDataFromDatabaseRequest)
	}
	err := c.DoAction(importDataFromDatabaseRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) ImportDataForSQLServer(importDataForSQLServerRequest *ImportDataForSQLServerRequest) (rds_ImportDataForSQLServerResponse.ImportDataForSQLServerResponse, error) {
	var resObj rds_ImportDataForSQLServerResponse.ImportDataForSQLServerResponse

	if importDataForSQLServerRequest == nil {
		importDataForSQLServerRequest = new(ImportDataForSQLServerRequest)
	}
	err := c.DoAction(importDataForSQLServerRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) ImportDatabaseBetweenInstances(importDatabaseBetweenInstancesRequest *ImportDatabaseBetweenInstancesRequest) (rds_ImportDatabaseBetweenInstancesResponse.ImportDatabaseBetweenInstancesResponse, error) {
	var resObj rds_ImportDatabaseBetweenInstancesResponse.ImportDatabaseBetweenInstancesResponse

	if importDatabaseBetweenInstancesRequest == nil {
		importDatabaseBetweenInstancesRequest = new(ImportDatabaseBetweenInstancesRequest)
	}
	err := c.DoAction(importDatabaseBetweenInstancesRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) GrantAccountPrivilege(grantAccountPrivilegeRequest *GrantAccountPrivilegeRequest) (rds_GrantAccountPrivilegeResponse.GrantAccountPrivilegeResponse, error) {
	var resObj rds_GrantAccountPrivilegeResponse.GrantAccountPrivilegeResponse

	if grantAccountPrivilegeRequest == nil {
		grantAccountPrivilegeRequest = new(GrantAccountPrivilegeRequest)
	}
	err := c.DoAction(grantAccountPrivilegeRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) ExtractBackupFromOAS(extractBackupFromOASRequest *ExtractBackupFromOASRequest) (rds_ExtractBackupFromOASResponse.ExtractBackupFromOASResponse, error) {
	var resObj rds_ExtractBackupFromOASResponse.ExtractBackupFromOASResponse

	if extractBackupFromOASRequest == nil {
		extractBackupFromOASRequest = new(ExtractBackupFromOASRequest)
	}
	err := c.DoAction(extractBackupFromOASRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeVpcZoneNos(describeVpcZoneNosRequest *DescribeVpcZoneNosRequest) (rds_DescribeVpcZoneNosResponse.DescribeVpcZoneNosResponse, error) {
	var resObj rds_DescribeVpcZoneNosResponse.DescribeVpcZoneNosResponse

	if describeVpcZoneNosRequest == nil {
		describeVpcZoneNosRequest = new(DescribeVpcZoneNosRequest)
	}
	err := c.DoAction(describeVpcZoneNosRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeTasks(describeTasksRequest *DescribeTasksRequest) (rds_DescribeTasksResponse.DescribeTasksResponse, error) {
	var resObj rds_DescribeTasksResponse.DescribeTasksResponse

	if describeTasksRequest == nil {
		describeTasksRequest = new(DescribeTasksRequest)
	}
	err := c.DoAction(describeTasksRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeSQLLogReports(describeSQLLogReportsRequest *DescribeSQLLogReportsRequest) (rds_DescribeSQLLogReportsResponse.DescribeSQLLogReportsResponse, error) {
	var resObj rds_DescribeSQLLogReportsResponse.DescribeSQLLogReportsResponse

	if describeSQLLogReportsRequest == nil {
		describeSQLLogReportsRequest = new(DescribeSQLLogReportsRequest)
	}
	err := c.DoAction(describeSQLLogReportsRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeSQLLogRecords(describeSQLLogRecordsRequest *DescribeSQLLogRecordsRequest) (rds_DescribeSQLLogRecordsResponse.DescribeSQLLogRecordsResponse, error) {
	var resObj rds_DescribeSQLLogRecordsResponse.DescribeSQLLogRecordsResponse

	if describeSQLLogRecordsRequest == nil {
		describeSQLLogRecordsRequest = new(DescribeSQLLogRecordsRequest)
	}
	err := c.DoAction(describeSQLLogRecordsRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeSQLInjectionInfos(describeSQLInjectionInfosRequest *DescribeSQLInjectionInfosRequest) (rds_DescribeSQLInjectionInfosResponse.DescribeSQLInjectionInfosResponse, error) {
	var resObj rds_DescribeSQLInjectionInfosResponse.DescribeSQLInjectionInfosResponse

	if describeSQLInjectionInfosRequest == nil {
		describeSQLInjectionInfosRequest = new(DescribeSQLInjectionInfosRequest)
	}
	err := c.DoAction(describeSQLInjectionInfosRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeSlowLogs(describeSlowLogsRequest *DescribeSlowLogsRequest) (rds_DescribeSlowLogsResponse.DescribeSlowLogsResponse, error) {
	var resObj rds_DescribeSlowLogsResponse.DescribeSlowLogsResponse

	if describeSlowLogsRequest == nil {
		describeSlowLogsRequest = new(DescribeSlowLogsRequest)
	}
	err := c.DoAction(describeSlowLogsRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeSlowLogRecords(describeSlowLogRecordsRequest *DescribeSlowLogRecordsRequest) (rds_DescribeSlowLogRecordsResponse.DescribeSlowLogRecordsResponse, error) {
	var resObj rds_DescribeSlowLogRecordsResponse.DescribeSlowLogRecordsResponse

	if describeSlowLogRecordsRequest == nil {
		describeSlowLogRecordsRequest = new(DescribeSlowLogRecordsRequest)
	}
	err := c.DoAction(describeSlowLogRecordsRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeResourceUsage(describeResourceUsageRequest *DescribeResourceUsageRequest) (rds_DescribeResourceUsageResponse.DescribeResourceUsageResponse, error) {
	var resObj rds_DescribeResourceUsageResponse.DescribeResourceUsageResponse

	if describeResourceUsageRequest == nil {
		describeResourceUsageRequest = new(DescribeResourceUsageRequest)
	}
	err := c.DoAction(describeResourceUsageRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeRegions(describeRegionsRequest *DescribeRegionsRequest) (rds_DescribeRegionsResponse.DescribeRegionsResponse, error) {
	var resObj rds_DescribeRegionsResponse.DescribeRegionsResponse

	if describeRegionsRequest == nil {
		describeRegionsRequest = new(DescribeRegionsRequest)
	}
	err := c.DoAction(describeRegionsRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeRealtimeDiagnoses(describeRealtimeDiagnosesRequest *DescribeRealtimeDiagnosesRequest) (rds_DescribeRealtimeDiagnosesResponse.DescribeRealtimeDiagnosesResponse, error) {
	var resObj rds_DescribeRealtimeDiagnosesResponse.DescribeRealtimeDiagnosesResponse

	if describeRealtimeDiagnosesRequest == nil {
		describeRealtimeDiagnosesRequest = new(DescribeRealtimeDiagnosesRequest)
	}
	err := c.DoAction(describeRealtimeDiagnosesRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribePreCheckResults(describePreCheckResultsRequest *DescribePreCheckResultsRequest) (rds_DescribePreCheckResultsResponse.DescribePreCheckResultsResponse, error) {
	var resObj rds_DescribePreCheckResultsResponse.DescribePreCheckResultsResponse

	if describePreCheckResultsRequest == nil {
		describePreCheckResultsRequest = new(DescribePreCheckResultsRequest)
	}
	err := c.DoAction(describePreCheckResultsRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeParameterTemplates(describeParameterTemplatesRequest *DescribeParameterTemplatesRequest) (rds_DescribeParameterTemplatesResponse.DescribeParameterTemplatesResponse, error) {
	var resObj rds_DescribeParameterTemplatesResponse.DescribeParameterTemplatesResponse

	if describeParameterTemplatesRequest == nil {
		describeParameterTemplatesRequest = new(DescribeParameterTemplatesRequest)
	}
	err := c.DoAction(describeParameterTemplatesRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeParameters(describeParametersRequest *DescribeParametersRequest) (rds_DescribeParametersResponse.DescribeParametersResponse, error) {
	var resObj rds_DescribeParametersResponse.DescribeParametersResponse

	if describeParametersRequest == nil {
		describeParametersRequest = new(DescribeParametersRequest)
	}
	err := c.DoAction(describeParametersRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeOptimizeAdviceOnStorage(describeOptimizeAdviceOnStorageRequest *DescribeOptimizeAdviceOnStorageRequest) (rds_DescribeOptimizeAdviceOnStorageResponse.DescribeOptimizeAdviceOnStorageResponse, error) {
	var resObj rds_DescribeOptimizeAdviceOnStorageResponse.DescribeOptimizeAdviceOnStorageResponse

	if describeOptimizeAdviceOnStorageRequest == nil {
		describeOptimizeAdviceOnStorageRequest = new(DescribeOptimizeAdviceOnStorageRequest)
	}
	err := c.DoAction(describeOptimizeAdviceOnStorageRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeOptimizeAdviceOnMissPK(describeOptimizeAdviceOnMissPKRequest *DescribeOptimizeAdviceOnMissPKRequest) (rds_DescribeOptimizeAdviceOnMissPKResponse.DescribeOptimizeAdviceOnMissPKResponse, error) {
	var resObj rds_DescribeOptimizeAdviceOnMissPKResponse.DescribeOptimizeAdviceOnMissPKResponse

	if describeOptimizeAdviceOnMissPKRequest == nil {
		describeOptimizeAdviceOnMissPKRequest = new(DescribeOptimizeAdviceOnMissPKRequest)
	}
	err := c.DoAction(describeOptimizeAdviceOnMissPKRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeOptimizeAdviceOnMissIndex(describeOptimizeAdviceOnMissIndexRequest *DescribeOptimizeAdviceOnMissIndexRequest) (rds_DescribeOptimizeAdviceOnMissIndexResponse.DescribeOptimizeAdviceOnMissIndexResponse, error) {
	var resObj rds_DescribeOptimizeAdviceOnMissIndexResponse.DescribeOptimizeAdviceOnMissIndexResponse

	if describeOptimizeAdviceOnMissIndexRequest == nil {
		describeOptimizeAdviceOnMissIndexRequest = new(DescribeOptimizeAdviceOnMissIndexRequest)
	}
	err := c.DoAction(describeOptimizeAdviceOnMissIndexRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeOptimizeAdviceOnExcessIndex(describeOptimizeAdviceOnExcessIndexRequest *DescribeOptimizeAdviceOnExcessIndexRequest) (rds_DescribeOptimizeAdviceOnExcessIndexResponse.DescribeOptimizeAdviceOnExcessIndexResponse, error) {
	var resObj rds_DescribeOptimizeAdviceOnExcessIndexResponse.DescribeOptimizeAdviceOnExcessIndexResponse

	if describeOptimizeAdviceOnExcessIndexRequest == nil {
		describeOptimizeAdviceOnExcessIndexRequest = new(DescribeOptimizeAdviceOnExcessIndexRequest)
	}
	err := c.DoAction(describeOptimizeAdviceOnExcessIndexRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeOptimizeAdviceOnBigTable(describeOptimizeAdviceOnBigTableRequest *DescribeOptimizeAdviceOnBigTableRequest) (rds_DescribeOptimizeAdviceOnBigTableResponse.DescribeOptimizeAdviceOnBigTableResponse, error) {
	var resObj rds_DescribeOptimizeAdviceOnBigTableResponse.DescribeOptimizeAdviceOnBigTableResponse

	if describeOptimizeAdviceOnBigTableRequest == nil {
		describeOptimizeAdviceOnBigTableRequest = new(DescribeOptimizeAdviceOnBigTableRequest)
	}
	err := c.DoAction(describeOptimizeAdviceOnBigTableRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeOptimizeAdviceByDBA(describeOptimizeAdviceByDBARequest *DescribeOptimizeAdviceByDBARequest) (rds_DescribeOptimizeAdviceByDBAResponse.DescribeOptimizeAdviceByDBAResponse, error) {
	var resObj rds_DescribeOptimizeAdviceByDBAResponse.DescribeOptimizeAdviceByDBAResponse

	if describeOptimizeAdviceByDBARequest == nil {
		describeOptimizeAdviceByDBARequest = new(DescribeOptimizeAdviceByDBARequest)
	}
	err := c.DoAction(describeOptimizeAdviceByDBARequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeOperationLogs(describeOperationLogsRequest *DescribeOperationLogsRequest) (rds_DescribeOperationLogsResponse.DescribeOperationLogsResponse, error) {
	var resObj rds_DescribeOperationLogsResponse.DescribeOperationLogsResponse

	if describeOperationLogsRequest == nil {
		describeOperationLogsRequest = new(DescribeOperationLogsRequest)
	}
	err := c.DoAction(describeOperationLogsRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeModifyParameterLog(describeModifyParameterLogRequest *DescribeModifyParameterLogRequest) (rds_DescribeModifyParameterLogResponse.DescribeModifyParameterLogResponse, error) {
	var resObj rds_DescribeModifyParameterLogResponse.DescribeModifyParameterLogResponse

	if describeModifyParameterLogRequest == nil {
		describeModifyParameterLogRequest = new(DescribeModifyParameterLogRequest)
	}
	err := c.DoAction(describeModifyParameterLogRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeImportsForSQLServer(describeImportsForSQLServerRequest *DescribeImportsForSQLServerRequest) (rds_DescribeImportsForSQLServerResponse.DescribeImportsForSQLServerResponse, error) {
	var resObj rds_DescribeImportsForSQLServerResponse.DescribeImportsForSQLServerResponse

	if describeImportsForSQLServerRequest == nil {
		describeImportsForSQLServerRequest = new(DescribeImportsForSQLServerRequest)
	}
	err := c.DoAction(describeImportsForSQLServerRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeFilesForSQLServer(describeFilesForSQLServerRequest *DescribeFilesForSQLServerRequest) (rds_DescribeFilesForSQLServerResponse.DescribeFilesForSQLServerResponse, error) {
	var resObj rds_DescribeFilesForSQLServerResponse.DescribeFilesForSQLServerResponse

	if describeFilesForSQLServerRequest == nil {
		describeFilesForSQLServerRequest = new(DescribeFilesForSQLServerRequest)
	}
	err := c.DoAction(describeFilesForSQLServerRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeErrorLogs(describeErrorLogsRequest *DescribeErrorLogsRequest) (rds_DescribeErrorLogsResponse.DescribeErrorLogsResponse, error) {
	var resObj rds_DescribeErrorLogsResponse.DescribeErrorLogsResponse

	if describeErrorLogsRequest == nil {
		describeErrorLogsRequest = new(DescribeErrorLogsRequest)
	}
	err := c.DoAction(describeErrorLogsRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeDBInstancePerformance(describeDBInstancePerformanceRequest *DescribeDBInstancePerformanceRequest) (rds_DescribeDBInstancePerformanceResponse.DescribeDBInstancePerformanceResponse, error) {
	var resObj rds_DescribeDBInstancePerformanceResponse.DescribeDBInstancePerformanceResponse

	if describeDBInstancePerformanceRequest == nil {
		describeDBInstancePerformanceRequest = new(DescribeDBInstancePerformanceRequest)
	}
	err := c.DoAction(describeDBInstancePerformanceRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeDatabases(describeDatabasesRequest *DescribeDatabasesRequest) (rds_DescribeDatabasesResponse.DescribeDatabasesResponse, error) {
	var resObj rds_DescribeDatabasesResponse.DescribeDatabasesResponse

	if describeDatabasesRequest == nil {
		describeDatabasesRequest = new(DescribeDatabasesRequest)
	}
	err := c.DoAction(describeDatabasesRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeBinlogFiles(describeBinlogFilesRequest *DescribeBinlogFilesRequest) (rds_DescribeBinlogFilesResponse.DescribeBinlogFilesResponse, error) {
	var resObj rds_DescribeBinlogFilesResponse.DescribeBinlogFilesResponse

	if describeBinlogFilesRequest == nil {
		describeBinlogFilesRequest = new(DescribeBinlogFilesRequest)
	}
	err := c.DoAction(describeBinlogFilesRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeBackupTasks(describeBackupTasksRequest *DescribeBackupTasksRequest) (rds_DescribeBackupTasksResponse.DescribeBackupTasksResponse, error) {
	var resObj rds_DescribeBackupTasksResponse.DescribeBackupTasksResponse

	if describeBackupTasksRequest == nil {
		describeBackupTasksRequest = new(DescribeBackupTasksRequest)
	}
	err := c.DoAction(describeBackupTasksRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeBackups(describeBackupsRequest *DescribeBackupsRequest) (rds_DescribeBackupsResponse.DescribeBackupsResponse, error) {
	var resObj rds_DescribeBackupsResponse.DescribeBackupsResponse

	if describeBackupsRequest == nil {
		describeBackupsRequest = new(DescribeBackupsRequest)
	}
	err := c.DoAction(describeBackupsRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeBackupPolicy(describeBackupPolicyRequest *DescribeBackupPolicyRequest) (rds_DescribeBackupPolicyResponse.DescribeBackupPolicyResponse, error) {
	var resObj rds_DescribeBackupPolicyResponse.DescribeBackupPolicyResponse

	if describeBackupPolicyRequest == nil {
		describeBackupPolicyRequest = new(DescribeBackupPolicyRequest)
	}
	err := c.DoAction(describeBackupPolicyRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeAccounts(describeAccountsRequest *DescribeAccountsRequest) (rds_DescribeAccountsResponse.DescribeAccountsResponse, error) {
	var resObj rds_DescribeAccountsResponse.DescribeAccountsResponse

	if describeAccountsRequest == nil {
		describeAccountsRequest = new(DescribeAccountsRequest)
	}
	err := c.DoAction(describeAccountsRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescibeImportsFromDatabase(descibeImportsFromDatabaseRequest *DescibeImportsFromDatabaseRequest) (rds_DescibeImportsFromDatabaseResponse.DescibeImportsFromDatabaseResponse, error) {
	var resObj rds_DescibeImportsFromDatabaseResponse.DescibeImportsFromDatabaseResponse

	if descibeImportsFromDatabaseRequest == nil {
		descibeImportsFromDatabaseRequest = new(DescibeImportsFromDatabaseRequest)
	}
	err := c.DoAction(descibeImportsFromDatabaseRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DeleteDBInstance(deleteDBInstanceRequest *DeleteDBInstanceRequest) (rds_DeleteDBInstanceResponse.DeleteDBInstanceResponse, error) {
	var resObj rds_DeleteDBInstanceResponse.DeleteDBInstanceResponse

	if deleteDBInstanceRequest == nil {
		deleteDBInstanceRequest = new(DeleteDBInstanceRequest)
	}
	err := c.DoAction(deleteDBInstanceRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DeleteDatabase(deleteDatabaseRequest *DeleteDatabaseRequest) (rds_DeleteDatabaseResponse.DeleteDatabaseResponse, error) {
	var resObj rds_DeleteDatabaseResponse.DeleteDatabaseResponse

	if deleteDatabaseRequest == nil {
		deleteDatabaseRequest = new(DeleteDatabaseRequest)
	}
	err := c.DoAction(deleteDatabaseRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DeleteAccount(deleteAccountRequest *DeleteAccountRequest) (rds_DeleteAccountResponse.DeleteAccountResponse, error) {
	var resObj rds_DeleteAccountResponse.DeleteAccountResponse

	if deleteAccountRequest == nil {
		deleteAccountRequest = new(DeleteAccountRequest)
	}
	err := c.DoAction(deleteAccountRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) CreateUploadPathForSQLServer(createUploadPathForSQLServerRequest *CreateUploadPathForSQLServerRequest) (rds_CreateUploadPathForSQLServerResponse.CreateUploadPathForSQLServerResponse, error) {
	var resObj rds_CreateUploadPathForSQLServerResponse.CreateUploadPathForSQLServerResponse

	if createUploadPathForSQLServerRequest == nil {
		createUploadPathForSQLServerRequest = new(CreateUploadPathForSQLServerRequest)
	}
	err := c.DoAction(createUploadPathForSQLServerRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) CreateTempDBInstance(createTempDBInstanceRequest *CreateTempDBInstanceRequest) (rds_CreateTempDBInstanceResponse.CreateTempDBInstanceResponse, error) {
	var resObj rds_CreateTempDBInstanceResponse.CreateTempDBInstanceResponse

	if createTempDBInstanceRequest == nil {
		createTempDBInstanceRequest = new(CreateTempDBInstanceRequest)
	}
	err := c.DoAction(createTempDBInstanceRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) CreatePostpaidDBInstanceForChannel(createPostpaidDBInstanceForChannelRequest *CreatePostpaidDBInstanceForChannelRequest) (rds_CreatePostpaidDBInstanceForChannelResponse.CreatePostpaidDBInstanceForChannelResponse, error) {
	var resObj rds_CreatePostpaidDBInstanceForChannelResponse.CreatePostpaidDBInstanceForChannelResponse

	if createPostpaidDBInstanceForChannelRequest == nil {
		createPostpaidDBInstanceForChannelRequest = new(CreatePostpaidDBInstanceForChannelRequest)
	}
	err := c.DoAction(createPostpaidDBInstanceForChannelRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) CreatePostpaidDBInstance(createPostpaidDBInstanceRequest *CreatePostpaidDBInstanceRequest) (rds_CreatePostpaidDBInstanceResponse.CreatePostpaidDBInstanceResponse, error) {
	var resObj rds_CreatePostpaidDBInstanceResponse.CreatePostpaidDBInstanceResponse

	if createPostpaidDBInstanceRequest == nil {
		createPostpaidDBInstanceRequest = new(CreatePostpaidDBInstanceRequest)
	}
	err := c.DoAction(createPostpaidDBInstanceRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) CreateDBInstanceforFirstPay(createDBInstanceforFirstPayRequest *CreateDBInstanceforFirstPayRequest) (rds_CreateDBInstanceforFirstPayResponse.CreateDBInstanceforFirstPayResponse, error) {
	var resObj rds_CreateDBInstanceforFirstPayResponse.CreateDBInstanceforFirstPayResponse

	if createDBInstanceforFirstPayRequest == nil {
		createDBInstanceforFirstPayRequest = new(CreateDBInstanceforFirstPayRequest)
	}
	err := c.DoAction(createDBInstanceforFirstPayRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) CreateDBInstanceForChannel(createDBInstanceForChannelRequest *CreateDBInstanceForChannelRequest) (rds_CreateDBInstanceForChannelResponse.CreateDBInstanceForChannelResponse, error) {
	var resObj rds_CreateDBInstanceForChannelResponse.CreateDBInstanceForChannelResponse

	if createDBInstanceForChannelRequest == nil {
		createDBInstanceForChannelRequest = new(CreateDBInstanceForChannelRequest)
	}
	err := c.DoAction(createDBInstanceForChannelRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) CreateDatabase(createDatabaseRequest *CreateDatabaseRequest) (rds_CreateDatabaseResponse.CreateDatabaseResponse, error) {
	var resObj rds_CreateDatabaseResponse.CreateDatabaseResponse

	if createDatabaseRequest == nil {
		createDatabaseRequest = new(CreateDatabaseRequest)
	}
	err := c.DoAction(createDatabaseRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) CreateBackup(createBackupRequest *CreateBackupRequest) (rds_CreateBackupResponse.CreateBackupResponse, error) {
	var resObj rds_CreateBackupResponse.CreateBackupResponse

	if createBackupRequest == nil {
		createBackupRequest = new(CreateBackupRequest)
	}
	err := c.DoAction(createBackupRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) CreateAccount(createAccountRequest *CreateAccountRequest) (rds_CreateAccountResponse.CreateAccountResponse, error) {
	var resObj rds_CreateAccountResponse.CreateAccountResponse

	if createAccountRequest == nil {
		createAccountRequest = new(CreateAccountRequest)
	}
	err := c.DoAction(createAccountRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) CheckDBNameAvailable(checkDBNameAvailableRequest *CheckDBNameAvailableRequest) (rds_CheckDBNameAvailableResponse.CheckDBNameAvailableResponse, error) {
	var resObj rds_CheckDBNameAvailableResponse.CheckDBNameAvailableResponse

	if checkDBNameAvailableRequest == nil {
		checkDBNameAvailableRequest = new(CheckDBNameAvailableRequest)
	}
	err := c.DoAction(checkDBNameAvailableRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) CheckAccountNameAvailable(checkAccountNameAvailableRequest *CheckAccountNameAvailableRequest) (rds_CheckAccountNameAvailableResponse.CheckAccountNameAvailableResponse, error) {
	var resObj rds_CheckAccountNameAvailableResponse.CheckAccountNameAvailableResponse

	if checkAccountNameAvailableRequest == nil {
		checkAccountNameAvailableRequest = new(CheckAccountNameAvailableRequest)
	}
	err := c.DoAction(checkAccountNameAvailableRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) CancelImport(cancelImportRequest *CancelImportRequest) (rds_CancelImportResponse.CancelImportResponse, error) {
	var resObj rds_CancelImportResponse.CancelImportResponse

	if cancelImportRequest == nil {
		cancelImportRequest = new(CancelImportRequest)
	}
	err := c.DoAction(cancelImportRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) BatchRevokeAccountPrivilege(batchRevokeAccountPrivilegeRequest *BatchRevokeAccountPrivilegeRequest) (rds_BatchRevokeAccountPrivilegeResponse.BatchRevokeAccountPrivilegeResponse, error) {
	var resObj rds_BatchRevokeAccountPrivilegeResponse.BatchRevokeAccountPrivilegeResponse

	if batchRevokeAccountPrivilegeRequest == nil {
		batchRevokeAccountPrivilegeRequest = new(BatchRevokeAccountPrivilegeRequest)
	}
	err := c.DoAction(batchRevokeAccountPrivilegeRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) BatchGrantAccountPrivilege(batchGrantAccountPrivilegeRequest *BatchGrantAccountPrivilegeRequest) (rds_BatchGrantAccountPrivilegeResponse.BatchGrantAccountPrivilegeResponse, error) {
	var resObj rds_BatchGrantAccountPrivilegeResponse.BatchGrantAccountPrivilegeResponse

	if batchGrantAccountPrivilegeRequest == nil {
		batchGrantAccountPrivilegeRequest = new(BatchGrantAccountPrivilegeRequest)
	}
	err := c.DoAction(batchGrantAccountPrivilegeRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) AddTagsToResource(addTagsToResourceRequest *AddTagsToResourceRequest) (rds_AddTagsToResourceResponse.AddTagsToResourceResponse, error) {
	var resObj rds_AddTagsToResourceResponse.AddTagsToResourceResponse

	if addTagsToResourceRequest == nil {
		addTagsToResourceRequest = new(AddTagsToResourceRequest)
	}
	err := c.DoAction(addTagsToResourceRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) SwitchDBInstanceNetType(switchDBInstanceNetTypeRequest *SwitchDBInstanceNetTypeRequest) (rds_SwitchDBInstanceNetTypeResponse.SwitchDBInstanceNetTypeResponse, error) {
	var resObj rds_SwitchDBInstanceNetTypeResponse.SwitchDBInstanceNetTypeResponse

	if switchDBInstanceNetTypeRequest == nil {
		switchDBInstanceNetTypeRequest = new(SwitchDBInstanceNetTypeRequest)
	}
	err := c.DoAction(switchDBInstanceNetTypeRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) ReleaseInstancePublicConnection(releaseInstancePublicConnectionRequest *ReleaseInstancePublicConnectionRequest) (rds_ReleaseInstancePublicConnectionResponse.ReleaseInstancePublicConnectionResponse, error) {
	var resObj rds_ReleaseInstancePublicConnectionResponse.ReleaseInstancePublicConnectionResponse

	if releaseInstancePublicConnectionRequest == nil {
		releaseInstancePublicConnectionRequest = new(ReleaseInstancePublicConnectionRequest)
	}
	err := c.DoAction(releaseInstancePublicConnectionRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) ModifyDBInstanceNetworkType(modifyDBInstanceNetworkTypeRequest *ModifyDBInstanceNetworkTypeRequest) (rds_ModifyDBInstanceNetworkTypeResponse.ModifyDBInstanceNetworkTypeResponse, error) {
	var resObj rds_ModifyDBInstanceNetworkTypeResponse.ModifyDBInstanceNetworkTypeResponse

	if modifyDBInstanceNetworkTypeRequest == nil {
		modifyDBInstanceNetworkTypeRequest = new(ModifyDBInstanceNetworkTypeRequest)
	}
	err := c.DoAction(modifyDBInstanceNetworkTypeRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) ModifyDBInstanceConnectionString(modifyDBInstanceConnectionStringRequest *ModifyDBInstanceConnectionStringRequest) (rds_ModifyDBInstanceConnectionStringResponse.ModifyDBInstanceConnectionStringResponse, error) {
	var resObj rds_ModifyDBInstanceConnectionStringResponse.ModifyDBInstanceConnectionStringResponse

	if modifyDBInstanceConnectionStringRequest == nil {
		modifyDBInstanceConnectionStringRequest = new(ModifyDBInstanceConnectionStringRequest)
	}
	err := c.DoAction(modifyDBInstanceConnectionStringRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) ModifyDBInstanceConnectionMode(modifyDBInstanceConnectionModeRequest *ModifyDBInstanceConnectionModeRequest) (rds_ModifyDBInstanceConnectionModeResponse.ModifyDBInstanceConnectionModeResponse, error) {
	var resObj rds_ModifyDBInstanceConnectionModeResponse.ModifyDBInstanceConnectionModeResponse

	if modifyDBInstanceConnectionModeRequest == nil {
		modifyDBInstanceConnectionModeRequest = new(ModifyDBInstanceConnectionModeRequest)
	}
	err := c.DoAction(modifyDBInstanceConnectionModeRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeDBInstanceNetInfoForChannel(describeDBInstanceNetInfoForChannelRequest *DescribeDBInstanceNetInfoForChannelRequest) (rds_DescribeDBInstanceNetInfoForChannelResponse.DescribeDBInstanceNetInfoForChannelResponse, error) {
	var resObj rds_DescribeDBInstanceNetInfoForChannelResponse.DescribeDBInstanceNetInfoForChannelResponse

	if describeDBInstanceNetInfoForChannelRequest == nil {
		describeDBInstanceNetInfoForChannelRequest = new(DescribeDBInstanceNetInfoForChannelRequest)
	}
	err := c.DoAction(describeDBInstanceNetInfoForChannelRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeDBInstanceNetInfo(describeDBInstanceNetInfoRequest *DescribeDBInstanceNetInfoRequest) (rds_DescribeDBInstanceNetInfoResponse.DescribeDBInstanceNetInfoResponse, error) {
	var resObj rds_DescribeDBInstanceNetInfoResponse.DescribeDBInstanceNetInfoResponse

	if describeDBInstanceNetInfoRequest == nil {
		describeDBInstanceNetInfoRequest = new(DescribeDBInstanceNetInfoRequest)
	}
	err := c.DoAction(describeDBInstanceNetInfoRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) CreateReadOnlyDBInstance(createReadOnlyDBInstanceRequest *CreateReadOnlyDBInstanceRequest) (rds_CreateReadOnlyDBInstanceResponse.CreateReadOnlyDBInstanceResponse, error) {
	var resObj rds_CreateReadOnlyDBInstanceResponse.CreateReadOnlyDBInstanceResponse

	if createReadOnlyDBInstanceRequest == nil {
		createReadOnlyDBInstanceRequest = new(CreateReadOnlyDBInstanceRequest)
	}
	err := c.DoAction(createReadOnlyDBInstanceRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) CreateDBInstance(createDBInstanceRequest *CreateDBInstanceRequest) (rds_CreateDBInstanceResponse.CreateDBInstanceResponse, error) {
	var resObj rds_CreateDBInstanceResponse.CreateDBInstanceResponse

	if createDBInstanceRequest == nil {
		createDBInstanceRequest = new(CreateDBInstanceRequest)
	}
	err := c.DoAction(createDBInstanceRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) AllocateInstancePublicConnection(allocateInstancePublicConnectionRequest *AllocateInstancePublicConnectionRequest) (rds_AllocateInstancePublicConnectionResponse.AllocateInstancePublicConnectionResponse, error) {
	var resObj rds_AllocateInstancePublicConnectionResponse.AllocateInstancePublicConnectionResponse

	if allocateInstancePublicConnectionRequest == nil {
		allocateInstancePublicConnectionRequest = new(AllocateInstancePublicConnectionRequest)
	}
	err := c.DoAction(allocateInstancePublicConnectionRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeDBInstancesByPerformance(describeDBInstancesByPerformanceRequest *DescribeDBInstancesByPerformanceRequest) (rds_DescribeDBInstancesByPerformanceResponse.DescribeDBInstancesByPerformanceResponse, error) {
	var resObj rds_DescribeDBInstancesByPerformanceResponse.DescribeDBInstancesByPerformanceResponse

	if describeDBInstancesByPerformanceRequest == nil {
		describeDBInstancesByPerformanceRequest = new(DescribeDBInstancesByPerformanceRequest)
	}
	err := c.DoAction(describeDBInstancesByPerformanceRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeDBInstancesByExpireTime(describeDBInstancesByExpireTimeRequest *DescribeDBInstancesByExpireTimeRequest) (rds_DescribeDBInstancesByExpireTimeResponse.DescribeDBInstancesByExpireTimeResponse, error) {
	var resObj rds_DescribeDBInstancesByExpireTimeResponse.DescribeDBInstancesByExpireTimeResponse

	if describeDBInstancesByExpireTimeRequest == nil {
		describeDBInstancesByExpireTimeRequest = new(DescribeDBInstancesByExpireTimeRequest)
	}
	err := c.DoAction(describeDBInstancesByExpireTimeRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeDBInstances(describeDBInstancesRequest *DescribeDBInstancesRequest) (rds_DescribeDBInstancesResponse.DescribeDBInstancesResponse, error) {
	var resObj rds_DescribeDBInstancesResponse.DescribeDBInstancesResponse

	if describeDBInstancesRequest == nil {
		describeDBInstancesRequest = new(DescribeDBInstancesRequest)
	}
	err := c.DoAction(describeDBInstancesRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeDBInstanceAttribute(describeDBInstanceAttributeRequest *DescribeDBInstanceAttributeRequest) (rds_DescribeDBInstanceAttributeResponse.DescribeDBInstanceAttributeResponse, error) {
	var resObj rds_DescribeDBInstanceAttributeResponse.DescribeDBInstanceAttributeResponse

	if describeDBInstanceAttributeRequest == nil {
		describeDBInstanceAttributeRequest = new(DescribeDBInstanceAttributeRequest)
	}
	err := c.DoAction(describeDBInstanceAttributeRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) DescribeAbnormalDBInstances(describeAbnormalDBInstancesRequest *DescribeAbnormalDBInstancesRequest) (rds_DescribeAbnormalDBInstancesResponse.DescribeAbnormalDBInstancesResponse, error) {
	var resObj rds_DescribeAbnormalDBInstancesResponse.DescribeAbnormalDBInstancesResponse

	if describeAbnormalDBInstancesRequest == nil {
		describeAbnormalDBInstancesRequest = new(DescribeAbnormalDBInstancesRequest)
	}
	err := c.DoAction(describeAbnormalDBInstancesRequest, &resObj)
	return resObj, err
}

func (c *RdsClient) StartDBInstanceDiagnose(startDBInstanceDiagnoseRequest *StartDBInstanceDiagnoseRequest) (rds_StartDBInstanceDiagnoseResponse.StartDBInstanceDiagnoseResponse, error) {
	var resObj rds_StartDBInstanceDiagnoseResponse.StartDBInstanceDiagnoseResponse

	if startDBInstanceDiagnoseRequest == nil {
		startDBInstanceDiagnoseRequest = new(StartDBInstanceDiagnoseRequest)
	}
	err := c.DoAction(startDBInstanceDiagnoseRequest, &resObj)
	return resObj, err
}
