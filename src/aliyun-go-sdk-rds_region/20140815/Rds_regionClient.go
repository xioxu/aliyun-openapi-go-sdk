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

package rds_region

import (
	"aliyun-go-sdk-core"

	"aliyun-go-sdk-rds_region/20140815/DescribeDBInstanceIPArrayListResponse"

	"aliyun-go-sdk-rds_region/20140815/UpgradeDBInstanceEngineVersionResponse"

	"aliyun-go-sdk-rds_region/20140815/UnlockDBInstanceResponse"

	"aliyun-go-sdk-rds_region/20140815/SwitchDBInstanceNetTypeResponse"

	"aliyun-go-sdk-rds_region/20140815/StopSyncingResponse"

	"aliyun-go-sdk-rds_region/20140815/StartDBInstanceDiagnoseResponse"

	"aliyun-go-sdk-rds_region/20140815/RevokeAccountPrivilegeResponse"

	"aliyun-go-sdk-rds_region/20140815/RestoreDBInstanceResponse"

	"aliyun-go-sdk-rds_region/20140815/RestartDBInstanceResponse"

	"aliyun-go-sdk-rds_region/20140815/ResetAccountPasswordResponse"

	"aliyun-go-sdk-rds_region/20140815/ResetAccountForPGResponse"

	"aliyun-go-sdk-rds_region/20140815/RemoveTagsFromResourceResponse"

	"aliyun-go-sdk-rds_region/20140815/ReleaseInstancePublicConnectionResponse"

	"aliyun-go-sdk-rds_region/20140815/PurgeDBInstanceLogResponse"

	"aliyun-go-sdk-rds_region/20140815/PreCheckBeforeImportDataResponse"

	"aliyun-go-sdk-rds_region/20140815/ModifySecurityIpsResponse"

	"aliyun-go-sdk-rds_region/20140815/ModifyPostpaidDBInstanceSpecResponse"

	"aliyun-go-sdk-rds_region/20140815/ModifyParameterResponse"

	"aliyun-go-sdk-rds_region/20140815/ModifyDBInstanceSpecResponse"

	"aliyun-go-sdk-rds_region/20140815/ModifyDBInstanceNetworkTypeResponse"

	"aliyun-go-sdk-rds_region/20140815/ModifyDBInstanceMaintainTimeResponse"

	"aliyun-go-sdk-rds_region/20140815/ModifyDBInstanceDescriptionResponse"

	"aliyun-go-sdk-rds_region/20140815/ModifyDBInstanceConnectionStringResponse"

	"aliyun-go-sdk-rds_region/20140815/ModifyDBInstanceConnectionModeResponse"

	"aliyun-go-sdk-rds_region/20140815/ModifyDBDescriptionResponse"

	"aliyun-go-sdk-rds_region/20140815/ModifyBackupPolicyResponse"

	"aliyun-go-sdk-rds_region/20140815/ModifyAccountDescriptionResponse"

	"aliyun-go-sdk-rds_region/20140815/MigrateToOtherZoneResponse"

	"aliyun-go-sdk-rds_region/20140815/LockDBInstanceResponse"

	"aliyun-go-sdk-rds_region/20140815/ImportDataFromDatabaseResponse"

	"aliyun-go-sdk-rds_region/20140815/ImportDataForSQLServerResponse"

	"aliyun-go-sdk-rds_region/20140815/ImportDatabaseBetweenInstancesResponse"

	"aliyun-go-sdk-rds_region/20140815/GrantAccountPrivilegeResponse"

	"aliyun-go-sdk-rds_region/20140815/ExtractBackupFromOASResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeVpcZoneNosResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeTasksResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeSQLLogReportsResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeSQLLogRecordsResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeSQLInjectionInfosResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeSlowLogsResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeSlowLogRecordsResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeResourceUsageResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeRegionsResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeRealtimeDiagnosesResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribePreCheckResultsResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeParameterTemplatesResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeParametersResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeOptimizeAdviceOnStorageResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeOptimizeAdviceOnMissPKResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeOptimizeAdviceOnMissIndexResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeOptimizeAdviceOnExcessIndexResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeOptimizeAdviceOnBigTableResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeOptimizeAdviceByDBAResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeOperationLogsResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeModifyParameterLogResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeImportsForSQLServerResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeFilesForSQLServerResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeErrorLogsResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeDBInstancesByPerformanceResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeDBInstancesByExpireTimeResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeDBInstancesResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeDBInstancePerformanceResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeDBInstanceNetInfoForChannelResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeDBInstanceNetInfoResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeDBInstanceAttributeResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeDatabasesResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeBinlogFilesResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeBackupTasksResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeBackupsResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeBackupPolicyResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeAccountsResponse"

	"aliyun-go-sdk-rds_region/20140815/DescribeAbnormalDBInstancesResponse"

	"aliyun-go-sdk-rds_region/20140815/DescibeImportsFromDatabaseResponse"

	"aliyun-go-sdk-rds_region/20140815/DeleteDBInstanceResponse"

	"aliyun-go-sdk-rds_region/20140815/DeleteDatabaseResponse"

	"aliyun-go-sdk-rds_region/20140815/DeleteAccountResponse"

	"aliyun-go-sdk-rds_region/20140815/CreateUploadPathForSQLServerResponse"

	"aliyun-go-sdk-rds_region/20140815/CreateTempDBInstanceResponse"

	"aliyun-go-sdk-rds_region/20140815/CreateReadOnlyDBInstanceResponse"

	"aliyun-go-sdk-rds_region/20140815/CreatePostpaidDBInstanceForChannelResponse"

	"aliyun-go-sdk-rds_region/20140815/CreatePostpaidDBInstanceResponse"

	"aliyun-go-sdk-rds_region/20140815/CreateDBInstanceforFirstPayResponse"

	"aliyun-go-sdk-rds_region/20140815/CreateDBInstanceForChannelResponse"

	"aliyun-go-sdk-rds_region/20140815/CreateDBInstanceResponse"

	"aliyun-go-sdk-rds_region/20140815/CreateDatabaseForInnerResponse"

	"aliyun-go-sdk-rds_region/20140815/CreateDatabaseResponse"

	"aliyun-go-sdk-rds_region/20140815/CreateBackupResponse"

	"aliyun-go-sdk-rds_region/20140815/CreateAccountForInnerResponse"

	"aliyun-go-sdk-rds_region/20140815/CreateAccountResponse"

	"aliyun-go-sdk-rds_region/20140815/CheckDBNameAvailableResponse"

	"aliyun-go-sdk-rds_region/20140815/CheckAccountNameAvailableResponse"

	"aliyun-go-sdk-rds_region/20140815/CancelImportResponse"

	"aliyun-go-sdk-rds_region/20140815/BatchRevokeAccountPrivilegeResponse"

	"aliyun-go-sdk-rds_region/20140815/BatchGrantAccountPrivilegeResponse"

	"aliyun-go-sdk-rds_region/20140815/AllocateInstancePublicConnectionResponse"

	"aliyun-go-sdk-rds_region/20140815/AllocateInstancePrivateConnectionResponse"

	"aliyun-go-sdk-rds_region/20140815/AddTagsToResourceResponse"
)

type Rds_regionClient struct {
	core.BaseClient
}

func New(profile *core.AccessProfile) *Rds_regionClient {
	p := &Rds_regionClient{}
	p.Version = "2014-08-15"
	p.ProductName = "Rds_Region"
	p.Profile = profile
	p.ApiStyle = "RPC"
	p.HttpRequestBuilder = &core.RpcHttpRequestBuilder{}
	return p
}

func (c *Rds_regionClient) DescribeDBInstanceIPArrayList(describeDBInstanceIPArrayListRequest *DescribeDBInstanceIPArrayListRequest) (rds_region_DescribeDBInstanceIPArrayListResponse.DescribeDBInstanceIPArrayListResponse, error) {
	var resObj rds_region_DescribeDBInstanceIPArrayListResponse.DescribeDBInstanceIPArrayListResponse

	if describeDBInstanceIPArrayListRequest == nil {
		describeDBInstanceIPArrayListRequest = new(DescribeDBInstanceIPArrayListRequest)
	}
	err := c.DoAction(describeDBInstanceIPArrayListRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) UpgradeDBInstanceEngineVersion(upgradeDBInstanceEngineVersionRequest *UpgradeDBInstanceEngineVersionRequest) (rds_region_UpgradeDBInstanceEngineVersionResponse.UpgradeDBInstanceEngineVersionResponse, error) {
	var resObj rds_region_UpgradeDBInstanceEngineVersionResponse.UpgradeDBInstanceEngineVersionResponse

	if upgradeDBInstanceEngineVersionRequest == nil {
		upgradeDBInstanceEngineVersionRequest = new(UpgradeDBInstanceEngineVersionRequest)
	}
	err := c.DoAction(upgradeDBInstanceEngineVersionRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) UnlockDBInstance(unlockDBInstanceRequest *UnlockDBInstanceRequest) (rds_region_UnlockDBInstanceResponse.UnlockDBInstanceResponse, error) {
	var resObj rds_region_UnlockDBInstanceResponse.UnlockDBInstanceResponse

	if unlockDBInstanceRequest == nil {
		unlockDBInstanceRequest = new(UnlockDBInstanceRequest)
	}
	err := c.DoAction(unlockDBInstanceRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) SwitchDBInstanceNetType(switchDBInstanceNetTypeRequest *SwitchDBInstanceNetTypeRequest) (rds_region_SwitchDBInstanceNetTypeResponse.SwitchDBInstanceNetTypeResponse, error) {
	var resObj rds_region_SwitchDBInstanceNetTypeResponse.SwitchDBInstanceNetTypeResponse

	if switchDBInstanceNetTypeRequest == nil {
		switchDBInstanceNetTypeRequest = new(SwitchDBInstanceNetTypeRequest)
	}
	err := c.DoAction(switchDBInstanceNetTypeRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) StopSyncing(stopSyncingRequest *StopSyncingRequest) (rds_region_StopSyncingResponse.StopSyncingResponse, error) {
	var resObj rds_region_StopSyncingResponse.StopSyncingResponse

	if stopSyncingRequest == nil {
		stopSyncingRequest = new(StopSyncingRequest)
	}
	err := c.DoAction(stopSyncingRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) StartDBInstanceDiagnose(startDBInstanceDiagnoseRequest *StartDBInstanceDiagnoseRequest) (rds_region_StartDBInstanceDiagnoseResponse.StartDBInstanceDiagnoseResponse, error) {
	var resObj rds_region_StartDBInstanceDiagnoseResponse.StartDBInstanceDiagnoseResponse

	if startDBInstanceDiagnoseRequest == nil {
		startDBInstanceDiagnoseRequest = new(StartDBInstanceDiagnoseRequest)
	}
	err := c.DoAction(startDBInstanceDiagnoseRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) RevokeAccountPrivilege(revokeAccountPrivilegeRequest *RevokeAccountPrivilegeRequest) (rds_region_RevokeAccountPrivilegeResponse.RevokeAccountPrivilegeResponse, error) {
	var resObj rds_region_RevokeAccountPrivilegeResponse.RevokeAccountPrivilegeResponse

	if revokeAccountPrivilegeRequest == nil {
		revokeAccountPrivilegeRequest = new(RevokeAccountPrivilegeRequest)
	}
	err := c.DoAction(revokeAccountPrivilegeRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) RestoreDBInstance(restoreDBInstanceRequest *RestoreDBInstanceRequest) (rds_region_RestoreDBInstanceResponse.RestoreDBInstanceResponse, error) {
	var resObj rds_region_RestoreDBInstanceResponse.RestoreDBInstanceResponse

	if restoreDBInstanceRequest == nil {
		restoreDBInstanceRequest = new(RestoreDBInstanceRequest)
	}
	err := c.DoAction(restoreDBInstanceRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) RestartDBInstance(restartDBInstanceRequest *RestartDBInstanceRequest) (rds_region_RestartDBInstanceResponse.RestartDBInstanceResponse, error) {
	var resObj rds_region_RestartDBInstanceResponse.RestartDBInstanceResponse

	if restartDBInstanceRequest == nil {
		restartDBInstanceRequest = new(RestartDBInstanceRequest)
	}
	err := c.DoAction(restartDBInstanceRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) ResetAccountPassword(resetAccountPasswordRequest *ResetAccountPasswordRequest) (rds_region_ResetAccountPasswordResponse.ResetAccountPasswordResponse, error) {
	var resObj rds_region_ResetAccountPasswordResponse.ResetAccountPasswordResponse

	if resetAccountPasswordRequest == nil {
		resetAccountPasswordRequest = new(ResetAccountPasswordRequest)
	}
	err := c.DoAction(resetAccountPasswordRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) ResetAccountForPG(resetAccountForPGRequest *ResetAccountForPGRequest) (rds_region_ResetAccountForPGResponse.ResetAccountForPGResponse, error) {
	var resObj rds_region_ResetAccountForPGResponse.ResetAccountForPGResponse

	if resetAccountForPGRequest == nil {
		resetAccountForPGRequest = new(ResetAccountForPGRequest)
	}
	err := c.DoAction(resetAccountForPGRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) RemoveTagsFromResource(removeTagsFromResourceRequest *RemoveTagsFromResourceRequest) (rds_region_RemoveTagsFromResourceResponse.RemoveTagsFromResourceResponse, error) {
	var resObj rds_region_RemoveTagsFromResourceResponse.RemoveTagsFromResourceResponse

	if removeTagsFromResourceRequest == nil {
		removeTagsFromResourceRequest = new(RemoveTagsFromResourceRequest)
	}
	err := c.DoAction(removeTagsFromResourceRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) ReleaseInstancePublicConnection(releaseInstancePublicConnectionRequest *ReleaseInstancePublicConnectionRequest) (rds_region_ReleaseInstancePublicConnectionResponse.ReleaseInstancePublicConnectionResponse, error) {
	var resObj rds_region_ReleaseInstancePublicConnectionResponse.ReleaseInstancePublicConnectionResponse

	if releaseInstancePublicConnectionRequest == nil {
		releaseInstancePublicConnectionRequest = new(ReleaseInstancePublicConnectionRequest)
	}
	err := c.DoAction(releaseInstancePublicConnectionRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) PurgeDBInstanceLog(purgeDBInstanceLogRequest *PurgeDBInstanceLogRequest) (rds_region_PurgeDBInstanceLogResponse.PurgeDBInstanceLogResponse, error) {
	var resObj rds_region_PurgeDBInstanceLogResponse.PurgeDBInstanceLogResponse

	if purgeDBInstanceLogRequest == nil {
		purgeDBInstanceLogRequest = new(PurgeDBInstanceLogRequest)
	}
	err := c.DoAction(purgeDBInstanceLogRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) PreCheckBeforeImportData(preCheckBeforeImportDataRequest *PreCheckBeforeImportDataRequest) (rds_region_PreCheckBeforeImportDataResponse.PreCheckBeforeImportDataResponse, error) {
	var resObj rds_region_PreCheckBeforeImportDataResponse.PreCheckBeforeImportDataResponse

	if preCheckBeforeImportDataRequest == nil {
		preCheckBeforeImportDataRequest = new(PreCheckBeforeImportDataRequest)
	}
	err := c.DoAction(preCheckBeforeImportDataRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) ModifySecurityIps(modifySecurityIpsRequest *ModifySecurityIpsRequest) (rds_region_ModifySecurityIpsResponse.ModifySecurityIpsResponse, error) {
	var resObj rds_region_ModifySecurityIpsResponse.ModifySecurityIpsResponse

	if modifySecurityIpsRequest == nil {
		modifySecurityIpsRequest = new(ModifySecurityIpsRequest)
	}
	err := c.DoAction(modifySecurityIpsRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) ModifyPostpaidDBInstanceSpec(modifyPostpaidDBInstanceSpecRequest *ModifyPostpaidDBInstanceSpecRequest) (rds_region_ModifyPostpaidDBInstanceSpecResponse.ModifyPostpaidDBInstanceSpecResponse, error) {
	var resObj rds_region_ModifyPostpaidDBInstanceSpecResponse.ModifyPostpaidDBInstanceSpecResponse

	if modifyPostpaidDBInstanceSpecRequest == nil {
		modifyPostpaidDBInstanceSpecRequest = new(ModifyPostpaidDBInstanceSpecRequest)
	}
	err := c.DoAction(modifyPostpaidDBInstanceSpecRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) ModifyParameter(modifyParameterRequest *ModifyParameterRequest) (rds_region_ModifyParameterResponse.ModifyParameterResponse, error) {
	var resObj rds_region_ModifyParameterResponse.ModifyParameterResponse

	if modifyParameterRequest == nil {
		modifyParameterRequest = new(ModifyParameterRequest)
	}
	err := c.DoAction(modifyParameterRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) ModifyDBInstanceSpec(modifyDBInstanceSpecRequest *ModifyDBInstanceSpecRequest) (rds_region_ModifyDBInstanceSpecResponse.ModifyDBInstanceSpecResponse, error) {
	var resObj rds_region_ModifyDBInstanceSpecResponse.ModifyDBInstanceSpecResponse

	if modifyDBInstanceSpecRequest == nil {
		modifyDBInstanceSpecRequest = new(ModifyDBInstanceSpecRequest)
	}
	err := c.DoAction(modifyDBInstanceSpecRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) ModifyDBInstanceNetworkType(modifyDBInstanceNetworkTypeRequest *ModifyDBInstanceNetworkTypeRequest) (rds_region_ModifyDBInstanceNetworkTypeResponse.ModifyDBInstanceNetworkTypeResponse, error) {
	var resObj rds_region_ModifyDBInstanceNetworkTypeResponse.ModifyDBInstanceNetworkTypeResponse

	if modifyDBInstanceNetworkTypeRequest == nil {
		modifyDBInstanceNetworkTypeRequest = new(ModifyDBInstanceNetworkTypeRequest)
	}
	err := c.DoAction(modifyDBInstanceNetworkTypeRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) ModifyDBInstanceMaintainTime(modifyDBInstanceMaintainTimeRequest *ModifyDBInstanceMaintainTimeRequest) (rds_region_ModifyDBInstanceMaintainTimeResponse.ModifyDBInstanceMaintainTimeResponse, error) {
	var resObj rds_region_ModifyDBInstanceMaintainTimeResponse.ModifyDBInstanceMaintainTimeResponse

	if modifyDBInstanceMaintainTimeRequest == nil {
		modifyDBInstanceMaintainTimeRequest = new(ModifyDBInstanceMaintainTimeRequest)
	}
	err := c.DoAction(modifyDBInstanceMaintainTimeRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) ModifyDBInstanceDescription(modifyDBInstanceDescriptionRequest *ModifyDBInstanceDescriptionRequest) (rds_region_ModifyDBInstanceDescriptionResponse.ModifyDBInstanceDescriptionResponse, error) {
	var resObj rds_region_ModifyDBInstanceDescriptionResponse.ModifyDBInstanceDescriptionResponse

	if modifyDBInstanceDescriptionRequest == nil {
		modifyDBInstanceDescriptionRequest = new(ModifyDBInstanceDescriptionRequest)
	}
	err := c.DoAction(modifyDBInstanceDescriptionRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) ModifyDBInstanceConnectionString(modifyDBInstanceConnectionStringRequest *ModifyDBInstanceConnectionStringRequest) (rds_region_ModifyDBInstanceConnectionStringResponse.ModifyDBInstanceConnectionStringResponse, error) {
	var resObj rds_region_ModifyDBInstanceConnectionStringResponse.ModifyDBInstanceConnectionStringResponse

	if modifyDBInstanceConnectionStringRequest == nil {
		modifyDBInstanceConnectionStringRequest = new(ModifyDBInstanceConnectionStringRequest)
	}
	err := c.DoAction(modifyDBInstanceConnectionStringRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) ModifyDBInstanceConnectionMode(modifyDBInstanceConnectionModeRequest *ModifyDBInstanceConnectionModeRequest) (rds_region_ModifyDBInstanceConnectionModeResponse.ModifyDBInstanceConnectionModeResponse, error) {
	var resObj rds_region_ModifyDBInstanceConnectionModeResponse.ModifyDBInstanceConnectionModeResponse

	if modifyDBInstanceConnectionModeRequest == nil {
		modifyDBInstanceConnectionModeRequest = new(ModifyDBInstanceConnectionModeRequest)
	}
	err := c.DoAction(modifyDBInstanceConnectionModeRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) ModifyDBDescription(modifyDBDescriptionRequest *ModifyDBDescriptionRequest) (rds_region_ModifyDBDescriptionResponse.ModifyDBDescriptionResponse, error) {
	var resObj rds_region_ModifyDBDescriptionResponse.ModifyDBDescriptionResponse

	if modifyDBDescriptionRequest == nil {
		modifyDBDescriptionRequest = new(ModifyDBDescriptionRequest)
	}
	err := c.DoAction(modifyDBDescriptionRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) ModifyBackupPolicy(modifyBackupPolicyRequest *ModifyBackupPolicyRequest) (rds_region_ModifyBackupPolicyResponse.ModifyBackupPolicyResponse, error) {
	var resObj rds_region_ModifyBackupPolicyResponse.ModifyBackupPolicyResponse

	if modifyBackupPolicyRequest == nil {
		modifyBackupPolicyRequest = new(ModifyBackupPolicyRequest)
	}
	err := c.DoAction(modifyBackupPolicyRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) ModifyAccountDescription(modifyAccountDescriptionRequest *ModifyAccountDescriptionRequest) (rds_region_ModifyAccountDescriptionResponse.ModifyAccountDescriptionResponse, error) {
	var resObj rds_region_ModifyAccountDescriptionResponse.ModifyAccountDescriptionResponse

	if modifyAccountDescriptionRequest == nil {
		modifyAccountDescriptionRequest = new(ModifyAccountDescriptionRequest)
	}
	err := c.DoAction(modifyAccountDescriptionRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) MigrateToOtherZone(migrateToOtherZoneRequest *MigrateToOtherZoneRequest) (rds_region_MigrateToOtherZoneResponse.MigrateToOtherZoneResponse, error) {
	var resObj rds_region_MigrateToOtherZoneResponse.MigrateToOtherZoneResponse

	if migrateToOtherZoneRequest == nil {
		migrateToOtherZoneRequest = new(MigrateToOtherZoneRequest)
	}
	err := c.DoAction(migrateToOtherZoneRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) LockDBInstance(lockDBInstanceRequest *LockDBInstanceRequest) (rds_region_LockDBInstanceResponse.LockDBInstanceResponse, error) {
	var resObj rds_region_LockDBInstanceResponse.LockDBInstanceResponse

	if lockDBInstanceRequest == nil {
		lockDBInstanceRequest = new(LockDBInstanceRequest)
	}
	err := c.DoAction(lockDBInstanceRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) ImportDataFromDatabase(importDataFromDatabaseRequest *ImportDataFromDatabaseRequest) (rds_region_ImportDataFromDatabaseResponse.ImportDataFromDatabaseResponse, error) {
	var resObj rds_region_ImportDataFromDatabaseResponse.ImportDataFromDatabaseResponse

	if importDataFromDatabaseRequest == nil {
		importDataFromDatabaseRequest = new(ImportDataFromDatabaseRequest)
	}
	err := c.DoAction(importDataFromDatabaseRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) ImportDataForSQLServer(importDataForSQLServerRequest *ImportDataForSQLServerRequest) (rds_region_ImportDataForSQLServerResponse.ImportDataForSQLServerResponse, error) {
	var resObj rds_region_ImportDataForSQLServerResponse.ImportDataForSQLServerResponse

	if importDataForSQLServerRequest == nil {
		importDataForSQLServerRequest = new(ImportDataForSQLServerRequest)
	}
	err := c.DoAction(importDataForSQLServerRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) ImportDatabaseBetweenInstances(importDatabaseBetweenInstancesRequest *ImportDatabaseBetweenInstancesRequest) (rds_region_ImportDatabaseBetweenInstancesResponse.ImportDatabaseBetweenInstancesResponse, error) {
	var resObj rds_region_ImportDatabaseBetweenInstancesResponse.ImportDatabaseBetweenInstancesResponse

	if importDatabaseBetweenInstancesRequest == nil {
		importDatabaseBetweenInstancesRequest = new(ImportDatabaseBetweenInstancesRequest)
	}
	err := c.DoAction(importDatabaseBetweenInstancesRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) GrantAccountPrivilege(grantAccountPrivilegeRequest *GrantAccountPrivilegeRequest) (rds_region_GrantAccountPrivilegeResponse.GrantAccountPrivilegeResponse, error) {
	var resObj rds_region_GrantAccountPrivilegeResponse.GrantAccountPrivilegeResponse

	if grantAccountPrivilegeRequest == nil {
		grantAccountPrivilegeRequest = new(GrantAccountPrivilegeRequest)
	}
	err := c.DoAction(grantAccountPrivilegeRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) ExtractBackupFromOAS(extractBackupFromOASRequest *ExtractBackupFromOASRequest) (rds_region_ExtractBackupFromOASResponse.ExtractBackupFromOASResponse, error) {
	var resObj rds_region_ExtractBackupFromOASResponse.ExtractBackupFromOASResponse

	if extractBackupFromOASRequest == nil {
		extractBackupFromOASRequest = new(ExtractBackupFromOASRequest)
	}
	err := c.DoAction(extractBackupFromOASRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeVpcZoneNos(describeVpcZoneNosRequest *DescribeVpcZoneNosRequest) (rds_region_DescribeVpcZoneNosResponse.DescribeVpcZoneNosResponse, error) {
	var resObj rds_region_DescribeVpcZoneNosResponse.DescribeVpcZoneNosResponse

	if describeVpcZoneNosRequest == nil {
		describeVpcZoneNosRequest = new(DescribeVpcZoneNosRequest)
	}
	err := c.DoAction(describeVpcZoneNosRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeTasks(describeTasksRequest *DescribeTasksRequest) (rds_region_DescribeTasksResponse.DescribeTasksResponse, error) {
	var resObj rds_region_DescribeTasksResponse.DescribeTasksResponse

	if describeTasksRequest == nil {
		describeTasksRequest = new(DescribeTasksRequest)
	}
	err := c.DoAction(describeTasksRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeSQLLogReports(describeSQLLogReportsRequest *DescribeSQLLogReportsRequest) (rds_region_DescribeSQLLogReportsResponse.DescribeSQLLogReportsResponse, error) {
	var resObj rds_region_DescribeSQLLogReportsResponse.DescribeSQLLogReportsResponse

	if describeSQLLogReportsRequest == nil {
		describeSQLLogReportsRequest = new(DescribeSQLLogReportsRequest)
	}
	err := c.DoAction(describeSQLLogReportsRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeSQLLogRecords(describeSQLLogRecordsRequest *DescribeSQLLogRecordsRequest) (rds_region_DescribeSQLLogRecordsResponse.DescribeSQLLogRecordsResponse, error) {
	var resObj rds_region_DescribeSQLLogRecordsResponse.DescribeSQLLogRecordsResponse

	if describeSQLLogRecordsRequest == nil {
		describeSQLLogRecordsRequest = new(DescribeSQLLogRecordsRequest)
	}
	err := c.DoAction(describeSQLLogRecordsRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeSQLInjectionInfos(describeSQLInjectionInfosRequest *DescribeSQLInjectionInfosRequest) (rds_region_DescribeSQLInjectionInfosResponse.DescribeSQLInjectionInfosResponse, error) {
	var resObj rds_region_DescribeSQLInjectionInfosResponse.DescribeSQLInjectionInfosResponse

	if describeSQLInjectionInfosRequest == nil {
		describeSQLInjectionInfosRequest = new(DescribeSQLInjectionInfosRequest)
	}
	err := c.DoAction(describeSQLInjectionInfosRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeSlowLogs(describeSlowLogsRequest *DescribeSlowLogsRequest) (rds_region_DescribeSlowLogsResponse.DescribeSlowLogsResponse, error) {
	var resObj rds_region_DescribeSlowLogsResponse.DescribeSlowLogsResponse

	if describeSlowLogsRequest == nil {
		describeSlowLogsRequest = new(DescribeSlowLogsRequest)
	}
	err := c.DoAction(describeSlowLogsRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeSlowLogRecords(describeSlowLogRecordsRequest *DescribeSlowLogRecordsRequest) (rds_region_DescribeSlowLogRecordsResponse.DescribeSlowLogRecordsResponse, error) {
	var resObj rds_region_DescribeSlowLogRecordsResponse.DescribeSlowLogRecordsResponse

	if describeSlowLogRecordsRequest == nil {
		describeSlowLogRecordsRequest = new(DescribeSlowLogRecordsRequest)
	}
	err := c.DoAction(describeSlowLogRecordsRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeResourceUsage(describeResourceUsageRequest *DescribeResourceUsageRequest) (rds_region_DescribeResourceUsageResponse.DescribeResourceUsageResponse, error) {
	var resObj rds_region_DescribeResourceUsageResponse.DescribeResourceUsageResponse

	if describeResourceUsageRequest == nil {
		describeResourceUsageRequest = new(DescribeResourceUsageRequest)
	}
	err := c.DoAction(describeResourceUsageRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeRegions(describeRegionsRequest *DescribeRegionsRequest) (rds_region_DescribeRegionsResponse.DescribeRegionsResponse, error) {
	var resObj rds_region_DescribeRegionsResponse.DescribeRegionsResponse

	if describeRegionsRequest == nil {
		describeRegionsRequest = new(DescribeRegionsRequest)
	}
	err := c.DoAction(describeRegionsRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeRealtimeDiagnoses(describeRealtimeDiagnosesRequest *DescribeRealtimeDiagnosesRequest) (rds_region_DescribeRealtimeDiagnosesResponse.DescribeRealtimeDiagnosesResponse, error) {
	var resObj rds_region_DescribeRealtimeDiagnosesResponse.DescribeRealtimeDiagnosesResponse

	if describeRealtimeDiagnosesRequest == nil {
		describeRealtimeDiagnosesRequest = new(DescribeRealtimeDiagnosesRequest)
	}
	err := c.DoAction(describeRealtimeDiagnosesRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribePreCheckResults(describePreCheckResultsRequest *DescribePreCheckResultsRequest) (rds_region_DescribePreCheckResultsResponse.DescribePreCheckResultsResponse, error) {
	var resObj rds_region_DescribePreCheckResultsResponse.DescribePreCheckResultsResponse

	if describePreCheckResultsRequest == nil {
		describePreCheckResultsRequest = new(DescribePreCheckResultsRequest)
	}
	err := c.DoAction(describePreCheckResultsRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeParameterTemplates(describeParameterTemplatesRequest *DescribeParameterTemplatesRequest) (rds_region_DescribeParameterTemplatesResponse.DescribeParameterTemplatesResponse, error) {
	var resObj rds_region_DescribeParameterTemplatesResponse.DescribeParameterTemplatesResponse

	if describeParameterTemplatesRequest == nil {
		describeParameterTemplatesRequest = new(DescribeParameterTemplatesRequest)
	}
	err := c.DoAction(describeParameterTemplatesRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeParameters(describeParametersRequest *DescribeParametersRequest) (rds_region_DescribeParametersResponse.DescribeParametersResponse, error) {
	var resObj rds_region_DescribeParametersResponse.DescribeParametersResponse

	if describeParametersRequest == nil {
		describeParametersRequest = new(DescribeParametersRequest)
	}
	err := c.DoAction(describeParametersRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeOptimizeAdviceOnStorage(describeOptimizeAdviceOnStorageRequest *DescribeOptimizeAdviceOnStorageRequest) (rds_region_DescribeOptimizeAdviceOnStorageResponse.DescribeOptimizeAdviceOnStorageResponse, error) {
	var resObj rds_region_DescribeOptimizeAdviceOnStorageResponse.DescribeOptimizeAdviceOnStorageResponse

	if describeOptimizeAdviceOnStorageRequest == nil {
		describeOptimizeAdviceOnStorageRequest = new(DescribeOptimizeAdviceOnStorageRequest)
	}
	err := c.DoAction(describeOptimizeAdviceOnStorageRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeOptimizeAdviceOnMissPK(describeOptimizeAdviceOnMissPKRequest *DescribeOptimizeAdviceOnMissPKRequest) (rds_region_DescribeOptimizeAdviceOnMissPKResponse.DescribeOptimizeAdviceOnMissPKResponse, error) {
	var resObj rds_region_DescribeOptimizeAdviceOnMissPKResponse.DescribeOptimizeAdviceOnMissPKResponse

	if describeOptimizeAdviceOnMissPKRequest == nil {
		describeOptimizeAdviceOnMissPKRequest = new(DescribeOptimizeAdviceOnMissPKRequest)
	}
	err := c.DoAction(describeOptimizeAdviceOnMissPKRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeOptimizeAdviceOnMissIndex(describeOptimizeAdviceOnMissIndexRequest *DescribeOptimizeAdviceOnMissIndexRequest) (rds_region_DescribeOptimizeAdviceOnMissIndexResponse.DescribeOptimizeAdviceOnMissIndexResponse, error) {
	var resObj rds_region_DescribeOptimizeAdviceOnMissIndexResponse.DescribeOptimizeAdviceOnMissIndexResponse

	if describeOptimizeAdviceOnMissIndexRequest == nil {
		describeOptimizeAdviceOnMissIndexRequest = new(DescribeOptimizeAdviceOnMissIndexRequest)
	}
	err := c.DoAction(describeOptimizeAdviceOnMissIndexRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeOptimizeAdviceOnExcessIndex(describeOptimizeAdviceOnExcessIndexRequest *DescribeOptimizeAdviceOnExcessIndexRequest) (rds_region_DescribeOptimizeAdviceOnExcessIndexResponse.DescribeOptimizeAdviceOnExcessIndexResponse, error) {
	var resObj rds_region_DescribeOptimizeAdviceOnExcessIndexResponse.DescribeOptimizeAdviceOnExcessIndexResponse

	if describeOptimizeAdviceOnExcessIndexRequest == nil {
		describeOptimizeAdviceOnExcessIndexRequest = new(DescribeOptimizeAdviceOnExcessIndexRequest)
	}
	err := c.DoAction(describeOptimizeAdviceOnExcessIndexRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeOptimizeAdviceOnBigTable(describeOptimizeAdviceOnBigTableRequest *DescribeOptimizeAdviceOnBigTableRequest) (rds_region_DescribeOptimizeAdviceOnBigTableResponse.DescribeOptimizeAdviceOnBigTableResponse, error) {
	var resObj rds_region_DescribeOptimizeAdviceOnBigTableResponse.DescribeOptimizeAdviceOnBigTableResponse

	if describeOptimizeAdviceOnBigTableRequest == nil {
		describeOptimizeAdviceOnBigTableRequest = new(DescribeOptimizeAdviceOnBigTableRequest)
	}
	err := c.DoAction(describeOptimizeAdviceOnBigTableRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeOptimizeAdviceByDBA(describeOptimizeAdviceByDBARequest *DescribeOptimizeAdviceByDBARequest) (rds_region_DescribeOptimizeAdviceByDBAResponse.DescribeOptimizeAdviceByDBAResponse, error) {
	var resObj rds_region_DescribeOptimizeAdviceByDBAResponse.DescribeOptimizeAdviceByDBAResponse

	if describeOptimizeAdviceByDBARequest == nil {
		describeOptimizeAdviceByDBARequest = new(DescribeOptimizeAdviceByDBARequest)
	}
	err := c.DoAction(describeOptimizeAdviceByDBARequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeOperationLogs(describeOperationLogsRequest *DescribeOperationLogsRequest) (rds_region_DescribeOperationLogsResponse.DescribeOperationLogsResponse, error) {
	var resObj rds_region_DescribeOperationLogsResponse.DescribeOperationLogsResponse

	if describeOperationLogsRequest == nil {
		describeOperationLogsRequest = new(DescribeOperationLogsRequest)
	}
	err := c.DoAction(describeOperationLogsRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeModifyParameterLog(describeModifyParameterLogRequest *DescribeModifyParameterLogRequest) (rds_region_DescribeModifyParameterLogResponse.DescribeModifyParameterLogResponse, error) {
	var resObj rds_region_DescribeModifyParameterLogResponse.DescribeModifyParameterLogResponse

	if describeModifyParameterLogRequest == nil {
		describeModifyParameterLogRequest = new(DescribeModifyParameterLogRequest)
	}
	err := c.DoAction(describeModifyParameterLogRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeImportsForSQLServer(describeImportsForSQLServerRequest *DescribeImportsForSQLServerRequest) (rds_region_DescribeImportsForSQLServerResponse.DescribeImportsForSQLServerResponse, error) {
	var resObj rds_region_DescribeImportsForSQLServerResponse.DescribeImportsForSQLServerResponse

	if describeImportsForSQLServerRequest == nil {
		describeImportsForSQLServerRequest = new(DescribeImportsForSQLServerRequest)
	}
	err := c.DoAction(describeImportsForSQLServerRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeFilesForSQLServer(describeFilesForSQLServerRequest *DescribeFilesForSQLServerRequest) (rds_region_DescribeFilesForSQLServerResponse.DescribeFilesForSQLServerResponse, error) {
	var resObj rds_region_DescribeFilesForSQLServerResponse.DescribeFilesForSQLServerResponse

	if describeFilesForSQLServerRequest == nil {
		describeFilesForSQLServerRequest = new(DescribeFilesForSQLServerRequest)
	}
	err := c.DoAction(describeFilesForSQLServerRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeErrorLogs(describeErrorLogsRequest *DescribeErrorLogsRequest) (rds_region_DescribeErrorLogsResponse.DescribeErrorLogsResponse, error) {
	var resObj rds_region_DescribeErrorLogsResponse.DescribeErrorLogsResponse

	if describeErrorLogsRequest == nil {
		describeErrorLogsRequest = new(DescribeErrorLogsRequest)
	}
	err := c.DoAction(describeErrorLogsRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeDBInstancesByPerformance(describeDBInstancesByPerformanceRequest *DescribeDBInstancesByPerformanceRequest) (rds_region_DescribeDBInstancesByPerformanceResponse.DescribeDBInstancesByPerformanceResponse, error) {
	var resObj rds_region_DescribeDBInstancesByPerformanceResponse.DescribeDBInstancesByPerformanceResponse

	if describeDBInstancesByPerformanceRequest == nil {
		describeDBInstancesByPerformanceRequest = new(DescribeDBInstancesByPerformanceRequest)
	}
	err := c.DoAction(describeDBInstancesByPerformanceRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeDBInstancesByExpireTime(describeDBInstancesByExpireTimeRequest *DescribeDBInstancesByExpireTimeRequest) (rds_region_DescribeDBInstancesByExpireTimeResponse.DescribeDBInstancesByExpireTimeResponse, error) {
	var resObj rds_region_DescribeDBInstancesByExpireTimeResponse.DescribeDBInstancesByExpireTimeResponse

	if describeDBInstancesByExpireTimeRequest == nil {
		describeDBInstancesByExpireTimeRequest = new(DescribeDBInstancesByExpireTimeRequest)
	}
	err := c.DoAction(describeDBInstancesByExpireTimeRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeDBInstances(describeDBInstancesRequest *DescribeDBInstancesRequest) (rds_region_DescribeDBInstancesResponse.DescribeDBInstancesResponse, error) {
	var resObj rds_region_DescribeDBInstancesResponse.DescribeDBInstancesResponse

	if describeDBInstancesRequest == nil {
		describeDBInstancesRequest = new(DescribeDBInstancesRequest)
	}
	err := c.DoAction(describeDBInstancesRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeDBInstancePerformance(describeDBInstancePerformanceRequest *DescribeDBInstancePerformanceRequest) (rds_region_DescribeDBInstancePerformanceResponse.DescribeDBInstancePerformanceResponse, error) {
	var resObj rds_region_DescribeDBInstancePerformanceResponse.DescribeDBInstancePerformanceResponse

	if describeDBInstancePerformanceRequest == nil {
		describeDBInstancePerformanceRequest = new(DescribeDBInstancePerformanceRequest)
	}
	err := c.DoAction(describeDBInstancePerformanceRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeDBInstanceNetInfoForChannel(describeDBInstanceNetInfoForChannelRequest *DescribeDBInstanceNetInfoForChannelRequest) (rds_region_DescribeDBInstanceNetInfoForChannelResponse.DescribeDBInstanceNetInfoForChannelResponse, error) {
	var resObj rds_region_DescribeDBInstanceNetInfoForChannelResponse.DescribeDBInstanceNetInfoForChannelResponse

	if describeDBInstanceNetInfoForChannelRequest == nil {
		describeDBInstanceNetInfoForChannelRequest = new(DescribeDBInstanceNetInfoForChannelRequest)
	}
	err := c.DoAction(describeDBInstanceNetInfoForChannelRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeDBInstanceNetInfo(describeDBInstanceNetInfoRequest *DescribeDBInstanceNetInfoRequest) (rds_region_DescribeDBInstanceNetInfoResponse.DescribeDBInstanceNetInfoResponse, error) {
	var resObj rds_region_DescribeDBInstanceNetInfoResponse.DescribeDBInstanceNetInfoResponse

	if describeDBInstanceNetInfoRequest == nil {
		describeDBInstanceNetInfoRequest = new(DescribeDBInstanceNetInfoRequest)
	}
	err := c.DoAction(describeDBInstanceNetInfoRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeDBInstanceAttribute(describeDBInstanceAttributeRequest *DescribeDBInstanceAttributeRequest) (rds_region_DescribeDBInstanceAttributeResponse.DescribeDBInstanceAttributeResponse, error) {
	var resObj rds_region_DescribeDBInstanceAttributeResponse.DescribeDBInstanceAttributeResponse

	if describeDBInstanceAttributeRequest == nil {
		describeDBInstanceAttributeRequest = new(DescribeDBInstanceAttributeRequest)
	}
	err := c.DoAction(describeDBInstanceAttributeRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeDatabases(describeDatabasesRequest *DescribeDatabasesRequest) (rds_region_DescribeDatabasesResponse.DescribeDatabasesResponse, error) {
	var resObj rds_region_DescribeDatabasesResponse.DescribeDatabasesResponse

	if describeDatabasesRequest == nil {
		describeDatabasesRequest = new(DescribeDatabasesRequest)
	}
	err := c.DoAction(describeDatabasesRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeBinlogFiles(describeBinlogFilesRequest *DescribeBinlogFilesRequest) (rds_region_DescribeBinlogFilesResponse.DescribeBinlogFilesResponse, error) {
	var resObj rds_region_DescribeBinlogFilesResponse.DescribeBinlogFilesResponse

	if describeBinlogFilesRequest == nil {
		describeBinlogFilesRequest = new(DescribeBinlogFilesRequest)
	}
	err := c.DoAction(describeBinlogFilesRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeBackupTasks(describeBackupTasksRequest *DescribeBackupTasksRequest) (rds_region_DescribeBackupTasksResponse.DescribeBackupTasksResponse, error) {
	var resObj rds_region_DescribeBackupTasksResponse.DescribeBackupTasksResponse

	if describeBackupTasksRequest == nil {
		describeBackupTasksRequest = new(DescribeBackupTasksRequest)
	}
	err := c.DoAction(describeBackupTasksRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeBackups(describeBackupsRequest *DescribeBackupsRequest) (rds_region_DescribeBackupsResponse.DescribeBackupsResponse, error) {
	var resObj rds_region_DescribeBackupsResponse.DescribeBackupsResponse

	if describeBackupsRequest == nil {
		describeBackupsRequest = new(DescribeBackupsRequest)
	}
	err := c.DoAction(describeBackupsRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeBackupPolicy(describeBackupPolicyRequest *DescribeBackupPolicyRequest) (rds_region_DescribeBackupPolicyResponse.DescribeBackupPolicyResponse, error) {
	var resObj rds_region_DescribeBackupPolicyResponse.DescribeBackupPolicyResponse

	if describeBackupPolicyRequest == nil {
		describeBackupPolicyRequest = new(DescribeBackupPolicyRequest)
	}
	err := c.DoAction(describeBackupPolicyRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeAccounts(describeAccountsRequest *DescribeAccountsRequest) (rds_region_DescribeAccountsResponse.DescribeAccountsResponse, error) {
	var resObj rds_region_DescribeAccountsResponse.DescribeAccountsResponse

	if describeAccountsRequest == nil {
		describeAccountsRequest = new(DescribeAccountsRequest)
	}
	err := c.DoAction(describeAccountsRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescribeAbnormalDBInstances(describeAbnormalDBInstancesRequest *DescribeAbnormalDBInstancesRequest) (rds_region_DescribeAbnormalDBInstancesResponse.DescribeAbnormalDBInstancesResponse, error) {
	var resObj rds_region_DescribeAbnormalDBInstancesResponse.DescribeAbnormalDBInstancesResponse

	if describeAbnormalDBInstancesRequest == nil {
		describeAbnormalDBInstancesRequest = new(DescribeAbnormalDBInstancesRequest)
	}
	err := c.DoAction(describeAbnormalDBInstancesRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DescibeImportsFromDatabase(descibeImportsFromDatabaseRequest *DescibeImportsFromDatabaseRequest) (rds_region_DescibeImportsFromDatabaseResponse.DescibeImportsFromDatabaseResponse, error) {
	var resObj rds_region_DescibeImportsFromDatabaseResponse.DescibeImportsFromDatabaseResponse

	if descibeImportsFromDatabaseRequest == nil {
		descibeImportsFromDatabaseRequest = new(DescibeImportsFromDatabaseRequest)
	}
	err := c.DoAction(descibeImportsFromDatabaseRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DeleteDBInstance(deleteDBInstanceRequest *DeleteDBInstanceRequest) (rds_region_DeleteDBInstanceResponse.DeleteDBInstanceResponse, error) {
	var resObj rds_region_DeleteDBInstanceResponse.DeleteDBInstanceResponse

	if deleteDBInstanceRequest == nil {
		deleteDBInstanceRequest = new(DeleteDBInstanceRequest)
	}
	err := c.DoAction(deleteDBInstanceRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DeleteDatabase(deleteDatabaseRequest *DeleteDatabaseRequest) (rds_region_DeleteDatabaseResponse.DeleteDatabaseResponse, error) {
	var resObj rds_region_DeleteDatabaseResponse.DeleteDatabaseResponse

	if deleteDatabaseRequest == nil {
		deleteDatabaseRequest = new(DeleteDatabaseRequest)
	}
	err := c.DoAction(deleteDatabaseRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) DeleteAccount(deleteAccountRequest *DeleteAccountRequest) (rds_region_DeleteAccountResponse.DeleteAccountResponse, error) {
	var resObj rds_region_DeleteAccountResponse.DeleteAccountResponse

	if deleteAccountRequest == nil {
		deleteAccountRequest = new(DeleteAccountRequest)
	}
	err := c.DoAction(deleteAccountRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) CreateUploadPathForSQLServer(createUploadPathForSQLServerRequest *CreateUploadPathForSQLServerRequest) (rds_region_CreateUploadPathForSQLServerResponse.CreateUploadPathForSQLServerResponse, error) {
	var resObj rds_region_CreateUploadPathForSQLServerResponse.CreateUploadPathForSQLServerResponse

	if createUploadPathForSQLServerRequest == nil {
		createUploadPathForSQLServerRequest = new(CreateUploadPathForSQLServerRequest)
	}
	err := c.DoAction(createUploadPathForSQLServerRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) CreateTempDBInstance(createTempDBInstanceRequest *CreateTempDBInstanceRequest) (rds_region_CreateTempDBInstanceResponse.CreateTempDBInstanceResponse, error) {
	var resObj rds_region_CreateTempDBInstanceResponse.CreateTempDBInstanceResponse

	if createTempDBInstanceRequest == nil {
		createTempDBInstanceRequest = new(CreateTempDBInstanceRequest)
	}
	err := c.DoAction(createTempDBInstanceRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) CreateReadOnlyDBInstance(createReadOnlyDBInstanceRequest *CreateReadOnlyDBInstanceRequest) (rds_region_CreateReadOnlyDBInstanceResponse.CreateReadOnlyDBInstanceResponse, error) {
	var resObj rds_region_CreateReadOnlyDBInstanceResponse.CreateReadOnlyDBInstanceResponse

	if createReadOnlyDBInstanceRequest == nil {
		createReadOnlyDBInstanceRequest = new(CreateReadOnlyDBInstanceRequest)
	}
	err := c.DoAction(createReadOnlyDBInstanceRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) CreatePostpaidDBInstanceForChannel(createPostpaidDBInstanceForChannelRequest *CreatePostpaidDBInstanceForChannelRequest) (rds_region_CreatePostpaidDBInstanceForChannelResponse.CreatePostpaidDBInstanceForChannelResponse, error) {
	var resObj rds_region_CreatePostpaidDBInstanceForChannelResponse.CreatePostpaidDBInstanceForChannelResponse

	if createPostpaidDBInstanceForChannelRequest == nil {
		createPostpaidDBInstanceForChannelRequest = new(CreatePostpaidDBInstanceForChannelRequest)
	}
	err := c.DoAction(createPostpaidDBInstanceForChannelRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) CreatePostpaidDBInstance(createPostpaidDBInstanceRequest *CreatePostpaidDBInstanceRequest) (rds_region_CreatePostpaidDBInstanceResponse.CreatePostpaidDBInstanceResponse, error) {
	var resObj rds_region_CreatePostpaidDBInstanceResponse.CreatePostpaidDBInstanceResponse

	if createPostpaidDBInstanceRequest == nil {
		createPostpaidDBInstanceRequest = new(CreatePostpaidDBInstanceRequest)
	}
	err := c.DoAction(createPostpaidDBInstanceRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) CreateDBInstanceforFirstPay(createDBInstanceforFirstPayRequest *CreateDBInstanceforFirstPayRequest) (rds_region_CreateDBInstanceforFirstPayResponse.CreateDBInstanceforFirstPayResponse, error) {
	var resObj rds_region_CreateDBInstanceforFirstPayResponse.CreateDBInstanceforFirstPayResponse

	if createDBInstanceforFirstPayRequest == nil {
		createDBInstanceforFirstPayRequest = new(CreateDBInstanceforFirstPayRequest)
	}
	err := c.DoAction(createDBInstanceforFirstPayRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) CreateDBInstanceForChannel(createDBInstanceForChannelRequest *CreateDBInstanceForChannelRequest) (rds_region_CreateDBInstanceForChannelResponse.CreateDBInstanceForChannelResponse, error) {
	var resObj rds_region_CreateDBInstanceForChannelResponse.CreateDBInstanceForChannelResponse

	if createDBInstanceForChannelRequest == nil {
		createDBInstanceForChannelRequest = new(CreateDBInstanceForChannelRequest)
	}
	err := c.DoAction(createDBInstanceForChannelRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) CreateDBInstance(createDBInstanceRequest *CreateDBInstanceRequest) (rds_region_CreateDBInstanceResponse.CreateDBInstanceResponse, error) {
	var resObj rds_region_CreateDBInstanceResponse.CreateDBInstanceResponse

	if createDBInstanceRequest == nil {
		createDBInstanceRequest = new(CreateDBInstanceRequest)
	}
	err := c.DoAction(createDBInstanceRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) CreateDatabaseForInner(createDatabaseForInnerRequest *CreateDatabaseForInnerRequest) (rds_region_CreateDatabaseForInnerResponse.CreateDatabaseForInnerResponse, error) {
	var resObj rds_region_CreateDatabaseForInnerResponse.CreateDatabaseForInnerResponse

	if createDatabaseForInnerRequest == nil {
		createDatabaseForInnerRequest = new(CreateDatabaseForInnerRequest)
	}
	err := c.DoAction(createDatabaseForInnerRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) CreateDatabase(createDatabaseRequest *CreateDatabaseRequest) (rds_region_CreateDatabaseResponse.CreateDatabaseResponse, error) {
	var resObj rds_region_CreateDatabaseResponse.CreateDatabaseResponse

	if createDatabaseRequest == nil {
		createDatabaseRequest = new(CreateDatabaseRequest)
	}
	err := c.DoAction(createDatabaseRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) CreateBackup(createBackupRequest *CreateBackupRequest) (rds_region_CreateBackupResponse.CreateBackupResponse, error) {
	var resObj rds_region_CreateBackupResponse.CreateBackupResponse

	if createBackupRequest == nil {
		createBackupRequest = new(CreateBackupRequest)
	}
	err := c.DoAction(createBackupRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) CreateAccountForInner(createAccountForInnerRequest *CreateAccountForInnerRequest) (rds_region_CreateAccountForInnerResponse.CreateAccountForInnerResponse, error) {
	var resObj rds_region_CreateAccountForInnerResponse.CreateAccountForInnerResponse

	if createAccountForInnerRequest == nil {
		createAccountForInnerRequest = new(CreateAccountForInnerRequest)
	}
	err := c.DoAction(createAccountForInnerRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) CreateAccount(createAccountRequest *CreateAccountRequest) (rds_region_CreateAccountResponse.CreateAccountResponse, error) {
	var resObj rds_region_CreateAccountResponse.CreateAccountResponse

	if createAccountRequest == nil {
		createAccountRequest = new(CreateAccountRequest)
	}
	err := c.DoAction(createAccountRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) CheckDBNameAvailable(checkDBNameAvailableRequest *CheckDBNameAvailableRequest) (rds_region_CheckDBNameAvailableResponse.CheckDBNameAvailableResponse, error) {
	var resObj rds_region_CheckDBNameAvailableResponse.CheckDBNameAvailableResponse

	if checkDBNameAvailableRequest == nil {
		checkDBNameAvailableRequest = new(CheckDBNameAvailableRequest)
	}
	err := c.DoAction(checkDBNameAvailableRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) CheckAccountNameAvailable(checkAccountNameAvailableRequest *CheckAccountNameAvailableRequest) (rds_region_CheckAccountNameAvailableResponse.CheckAccountNameAvailableResponse, error) {
	var resObj rds_region_CheckAccountNameAvailableResponse.CheckAccountNameAvailableResponse

	if checkAccountNameAvailableRequest == nil {
		checkAccountNameAvailableRequest = new(CheckAccountNameAvailableRequest)
	}
	err := c.DoAction(checkAccountNameAvailableRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) CancelImport(cancelImportRequest *CancelImportRequest) (rds_region_CancelImportResponse.CancelImportResponse, error) {
	var resObj rds_region_CancelImportResponse.CancelImportResponse

	if cancelImportRequest == nil {
		cancelImportRequest = new(CancelImportRequest)
	}
	err := c.DoAction(cancelImportRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) BatchRevokeAccountPrivilege(batchRevokeAccountPrivilegeRequest *BatchRevokeAccountPrivilegeRequest) (rds_region_BatchRevokeAccountPrivilegeResponse.BatchRevokeAccountPrivilegeResponse, error) {
	var resObj rds_region_BatchRevokeAccountPrivilegeResponse.BatchRevokeAccountPrivilegeResponse

	if batchRevokeAccountPrivilegeRequest == nil {
		batchRevokeAccountPrivilegeRequest = new(BatchRevokeAccountPrivilegeRequest)
	}
	err := c.DoAction(batchRevokeAccountPrivilegeRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) BatchGrantAccountPrivilege(batchGrantAccountPrivilegeRequest *BatchGrantAccountPrivilegeRequest) (rds_region_BatchGrantAccountPrivilegeResponse.BatchGrantAccountPrivilegeResponse, error) {
	var resObj rds_region_BatchGrantAccountPrivilegeResponse.BatchGrantAccountPrivilegeResponse

	if batchGrantAccountPrivilegeRequest == nil {
		batchGrantAccountPrivilegeRequest = new(BatchGrantAccountPrivilegeRequest)
	}
	err := c.DoAction(batchGrantAccountPrivilegeRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) AllocateInstancePublicConnection(allocateInstancePublicConnectionRequest *AllocateInstancePublicConnectionRequest) (rds_region_AllocateInstancePublicConnectionResponse.AllocateInstancePublicConnectionResponse, error) {
	var resObj rds_region_AllocateInstancePublicConnectionResponse.AllocateInstancePublicConnectionResponse

	if allocateInstancePublicConnectionRequest == nil {
		allocateInstancePublicConnectionRequest = new(AllocateInstancePublicConnectionRequest)
	}
	err := c.DoAction(allocateInstancePublicConnectionRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) AllocateInstancePrivateConnection(allocateInstancePrivateConnectionRequest *AllocateInstancePrivateConnectionRequest) (rds_region_AllocateInstancePrivateConnectionResponse.AllocateInstancePrivateConnectionResponse, error) {
	var resObj rds_region_AllocateInstancePrivateConnectionResponse.AllocateInstancePrivateConnectionResponse

	if allocateInstancePrivateConnectionRequest == nil {
		allocateInstancePrivateConnectionRequest = new(AllocateInstancePrivateConnectionRequest)
	}
	err := c.DoAction(allocateInstancePrivateConnectionRequest, &resObj)
	return resObj, err
}

func (c *Rds_regionClient) AddTagsToResource(addTagsToResourceRequest *AddTagsToResourceRequest) (rds_region_AddTagsToResourceResponse.AddTagsToResourceResponse, error) {
	var resObj rds_region_AddTagsToResourceResponse.AddTagsToResourceResponse

	if addTagsToResourceRequest == nil {
		addTagsToResourceRequest = new(AddTagsToResourceRequest)
	}
	err := c.DoAction(addTagsToResourceRequest, &resObj)
	return resObj, err
}
