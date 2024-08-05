# Recovery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanTearDown** | Pointer to **NullableBool** | Specifies whether it&#39;s possible to tear down the objects created by the recovery. | [optional] 
**CreationInfo** | Pointer to [**CreationInfo**](CreationInfo.md) |  | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of the Recovery in Unix timestamp epoch in microseconds. This field will be populated only after Recovery is finished. | [optional] 
**Id** | Pointer to **NullableString** | Specifies the id of the Recovery. | [optional] 
**IsMultiStageRestore** | Pointer to **NullableBool** | Specifies whether the current recovery operation is a multi-stage restore operation. This is currently used by VMware recoveres for the migration/hot-standby use case. | [optional] 
**IsParentRecovery** | Pointer to **NullableBool** | Specifies whether the current recovery operation has created child recoveries. This is currently used in SQL recovery where multiple child recoveries can be tracked under a common/parent recovery. | [optional] 
**Messages** | Pointer to **[]string** | Specifies messages about the recovery. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the Recovery. | [optional] 
**ParentRecoveryId** | Pointer to **NullableString** | If current recovery is child recovery triggered by another parent recovery operation, then this field willt specify the id of the parent recovery. | [optional] 
**Permissions** | Pointer to [**[]Tenant**](Tenant.md) | Specifies the list of tenants that have permissions for this recovery. | [optional] 
**ProgressTaskId** | Pointer to **NullableString** | Progress monitor task id for Recovery. | [optional] 
**RecoveryAction** | Pointer to **string** | Specifies the type of recover action. | [optional] 
**RetrieveArchiveTasks** | Pointer to [**[]RetrieveArchiveTask**](RetrieveArchiveTask.md) | Specifies the list of persistent state of a retrieve of an archive task. | [optional] 
**SnapshotEnvironment** | Pointer to **string** | Specifies the type of snapshot environment for which the Recovery was performed. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the Recovery in Unix timestamp epoch in microseconds. | [optional] 
**Status** | Pointer to **NullableString** | Status of the Recovery. &#39;Running&#39; indicates that the Recovery is still running. &#39;Canceled&#39; indicates that the Recovery has been cancelled. &#39;Canceling&#39; indicates that the Recovery is in the process of being cancelled. &#39;Failed&#39; indicates that the Recovery has failed. &#39;Succeeded&#39; indicates that the Recovery has finished successfully. &#39;SucceededWithWarning&#39; indicates that the Recovery finished successfully, but there were some warning messages. &#39;Skipped&#39; indicates that the Recovery task was skipped. | [optional] 
**TearDownMessage** | Pointer to **NullableString** | Specifies the error message about the tear down operation if it fails. | [optional] 
**TearDownStatus** | Pointer to **NullableString** | Specifies the status of the tear down operation. This is only set when the canTearDown is set to true. &#39;DestroyScheduled&#39; indicates that the tear down is ready to schedule. &#39;Destroying&#39; indicates that the tear down is still running. &#39;Destroyed&#39; indicates that the tear down succeeded. &#39;DestroyError&#39; indicates that the tear down failed. | [optional] 
**AcropolisParams** | Pointer to [**RecoverAcropolisParams**](RecoverAcropolisParams.md) |  | [optional] 
**AwsParams** | Pointer to [**RecoverAwsParams**](RecoverAwsParams.md) |  | [optional] 
**AzureParams** | Pointer to [**RecoverAzureParams**](RecoverAzureParams.md) |  | [optional] 
**CassandraParams** | Pointer to [**CassandraParams**](CassandraParams.md) |  | [optional] 
**CouchbaseParams** | Pointer to [**CouchbaseParams**](CouchbaseParams.md) |  | [optional] 
**ElastifileParams** | Pointer to [**RecoverElastifileParams**](RecoverElastifileParams.md) |  | [optional] 
**ExchangeParams** | Pointer to [**RecoverExchangeParams**](RecoverExchangeParams.md) |  | [optional] 
**ExperimentalAdapterParams** | Pointer to [**ExperimentalAdapterParams**](ExperimentalAdapterParams.md) |  | [optional] 
**FlashbladeParams** | Pointer to [**RecoverFlashbladeParams**](RecoverFlashbladeParams.md) |  | [optional] 
**GcpParams** | Pointer to [**RecoverGcpParams**](RecoverGcpParams.md) |  | [optional] 
**GenericNasParams** | Pointer to [**RecoverGenericNasParams**](RecoverGenericNasParams.md) |  | [optional] 
**GpfsParams** | Pointer to [**RecoverGpfsParams**](RecoverGpfsParams.md) |  | [optional] 
**HbaseParams** | Pointer to [**HbaseParams**](HbaseParams.md) |  | [optional] 
**HdfsParams** | Pointer to [**HdfsParams**](HdfsParams.md) |  | [optional] 
**HiveParams** | Pointer to [**HiveParams**](HiveParams.md) |  | [optional] 
**HypervParams** | Pointer to [**RecoverHyperVParams**](RecoverHyperVParams.md) |  | [optional] 
**IbmFlashSystemParams** | Pointer to [**RecoverPureParams**](RecoverPureParams.md) |  | [optional] 
**IsilonParams** | Pointer to [**RecoverIsilonParams**](RecoverIsilonParams.md) |  | [optional] 
**KubernetesParams** | Pointer to [**RecoverKubernetesParams**](RecoverKubernetesParams.md) |  | [optional] 
**KvmParams** | Pointer to [**RecoverKvmParams**](RecoverKvmParams.md) |  | [optional] 
**MongodbParams** | Pointer to [**MongodbParams**](MongodbParams.md) |  | [optional] 
**MssqlParams** | Pointer to [**RecoverSqlParams**](RecoverSqlParams.md) |  | [optional] 
**NetappParams** | Pointer to [**RecoverNetappParams**](RecoverNetappParams.md) |  | [optional] 
**Office365Params** | Pointer to [**RecoverO365Params**](RecoverO365Params.md) |  | [optional] 
**OracleParams** | Pointer to [**RecoverOracleParams**](RecoverOracleParams.md) |  | [optional] 
**PhysicalParams** | Pointer to [**RecoverPhysicalParams**](RecoverPhysicalParams.md) |  | [optional] 
**PureParams** | Pointer to [**RecoverPureParams**](RecoverPureParams.md) |  | [optional] 
**SapHanaParams** | Pointer to [**SapHanaParams**](SapHanaParams.md) |  | [optional] 
**SfdcParams** | Pointer to [**RecoverSalesforceParams**](RecoverSalesforceParams.md) |  | [optional] 
**UdaParams** | Pointer to [**UdaParams**](UdaParams.md) |  | [optional] 
**ViewParams** | Pointer to [**RecoverViewParams**](RecoverViewParams.md) |  | [optional] 
**VmwareParams** | Pointer to [**RecoverVmwareParams**](RecoverVmwareParams.md) |  | [optional] 

## Methods

### NewRecovery

`func NewRecovery() *Recovery`

NewRecovery instantiates a new Recovery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryWithDefaults

`func NewRecoveryWithDefaults() *Recovery`

NewRecoveryWithDefaults instantiates a new Recovery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanTearDown

`func (o *Recovery) GetCanTearDown() bool`

GetCanTearDown returns the CanTearDown field if non-nil, zero value otherwise.

### GetCanTearDownOk

`func (o *Recovery) GetCanTearDownOk() (*bool, bool)`

GetCanTearDownOk returns a tuple with the CanTearDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTearDown

`func (o *Recovery) SetCanTearDown(v bool)`

SetCanTearDown sets CanTearDown field to given value.

### HasCanTearDown

`func (o *Recovery) HasCanTearDown() bool`

HasCanTearDown returns a boolean if a field has been set.

### SetCanTearDownNil

`func (o *Recovery) SetCanTearDownNil(b bool)`

 SetCanTearDownNil sets the value for CanTearDown to be an explicit nil

### UnsetCanTearDown
`func (o *Recovery) UnsetCanTearDown()`

UnsetCanTearDown ensures that no value is present for CanTearDown, not even an explicit nil
### GetCreationInfo

`func (o *Recovery) GetCreationInfo() CreationInfo`

GetCreationInfo returns the CreationInfo field if non-nil, zero value otherwise.

### GetCreationInfoOk

`func (o *Recovery) GetCreationInfoOk() (*CreationInfo, bool)`

GetCreationInfoOk returns a tuple with the CreationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationInfo

`func (o *Recovery) SetCreationInfo(v CreationInfo)`

SetCreationInfo sets CreationInfo field to given value.

### HasCreationInfo

`func (o *Recovery) HasCreationInfo() bool`

HasCreationInfo returns a boolean if a field has been set.

### GetEndTimeUsecs

`func (o *Recovery) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *Recovery) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *Recovery) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *Recovery) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *Recovery) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *Recovery) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetId

`func (o *Recovery) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Recovery) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Recovery) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Recovery) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Recovery) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Recovery) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsMultiStageRestore

`func (o *Recovery) GetIsMultiStageRestore() bool`

GetIsMultiStageRestore returns the IsMultiStageRestore field if non-nil, zero value otherwise.

### GetIsMultiStageRestoreOk

`func (o *Recovery) GetIsMultiStageRestoreOk() (*bool, bool)`

GetIsMultiStageRestoreOk returns a tuple with the IsMultiStageRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiStageRestore

`func (o *Recovery) SetIsMultiStageRestore(v bool)`

SetIsMultiStageRestore sets IsMultiStageRestore field to given value.

### HasIsMultiStageRestore

`func (o *Recovery) HasIsMultiStageRestore() bool`

HasIsMultiStageRestore returns a boolean if a field has been set.

### SetIsMultiStageRestoreNil

`func (o *Recovery) SetIsMultiStageRestoreNil(b bool)`

 SetIsMultiStageRestoreNil sets the value for IsMultiStageRestore to be an explicit nil

### UnsetIsMultiStageRestore
`func (o *Recovery) UnsetIsMultiStageRestore()`

UnsetIsMultiStageRestore ensures that no value is present for IsMultiStageRestore, not even an explicit nil
### GetIsParentRecovery

`func (o *Recovery) GetIsParentRecovery() bool`

GetIsParentRecovery returns the IsParentRecovery field if non-nil, zero value otherwise.

### GetIsParentRecoveryOk

`func (o *Recovery) GetIsParentRecoveryOk() (*bool, bool)`

GetIsParentRecoveryOk returns a tuple with the IsParentRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsParentRecovery

`func (o *Recovery) SetIsParentRecovery(v bool)`

SetIsParentRecovery sets IsParentRecovery field to given value.

### HasIsParentRecovery

`func (o *Recovery) HasIsParentRecovery() bool`

HasIsParentRecovery returns a boolean if a field has been set.

### SetIsParentRecoveryNil

`func (o *Recovery) SetIsParentRecoveryNil(b bool)`

 SetIsParentRecoveryNil sets the value for IsParentRecovery to be an explicit nil

### UnsetIsParentRecovery
`func (o *Recovery) UnsetIsParentRecovery()`

UnsetIsParentRecovery ensures that no value is present for IsParentRecovery, not even an explicit nil
### GetMessages

`func (o *Recovery) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *Recovery) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *Recovery) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *Recovery) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessagesNil

`func (o *Recovery) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *Recovery) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
### GetName

`func (o *Recovery) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Recovery) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Recovery) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Recovery) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Recovery) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Recovery) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParentRecoveryId

`func (o *Recovery) GetParentRecoveryId() string`

GetParentRecoveryId returns the ParentRecoveryId field if non-nil, zero value otherwise.

### GetParentRecoveryIdOk

`func (o *Recovery) GetParentRecoveryIdOk() (*string, bool)`

GetParentRecoveryIdOk returns a tuple with the ParentRecoveryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRecoveryId

`func (o *Recovery) SetParentRecoveryId(v string)`

SetParentRecoveryId sets ParentRecoveryId field to given value.

### HasParentRecoveryId

`func (o *Recovery) HasParentRecoveryId() bool`

HasParentRecoveryId returns a boolean if a field has been set.

### SetParentRecoveryIdNil

`func (o *Recovery) SetParentRecoveryIdNil(b bool)`

 SetParentRecoveryIdNil sets the value for ParentRecoveryId to be an explicit nil

### UnsetParentRecoveryId
`func (o *Recovery) UnsetParentRecoveryId()`

UnsetParentRecoveryId ensures that no value is present for ParentRecoveryId, not even an explicit nil
### GetPermissions

`func (o *Recovery) GetPermissions() []Tenant`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Recovery) GetPermissionsOk() (*[]Tenant, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Recovery) SetPermissions(v []Tenant)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *Recovery) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *Recovery) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *Recovery) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetProgressTaskId

`func (o *Recovery) GetProgressTaskId() string`

GetProgressTaskId returns the ProgressTaskId field if non-nil, zero value otherwise.

### GetProgressTaskIdOk

`func (o *Recovery) GetProgressTaskIdOk() (*string, bool)`

GetProgressTaskIdOk returns a tuple with the ProgressTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressTaskId

`func (o *Recovery) SetProgressTaskId(v string)`

SetProgressTaskId sets ProgressTaskId field to given value.

### HasProgressTaskId

`func (o *Recovery) HasProgressTaskId() bool`

HasProgressTaskId returns a boolean if a field has been set.

### SetProgressTaskIdNil

`func (o *Recovery) SetProgressTaskIdNil(b bool)`

 SetProgressTaskIdNil sets the value for ProgressTaskId to be an explicit nil

### UnsetProgressTaskId
`func (o *Recovery) UnsetProgressTaskId()`

UnsetProgressTaskId ensures that no value is present for ProgressTaskId, not even an explicit nil
### GetRecoveryAction

`func (o *Recovery) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *Recovery) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *Recovery) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.

### HasRecoveryAction

`func (o *Recovery) HasRecoveryAction() bool`

HasRecoveryAction returns a boolean if a field has been set.

### GetRetrieveArchiveTasks

`func (o *Recovery) GetRetrieveArchiveTasks() []RetrieveArchiveTask`

GetRetrieveArchiveTasks returns the RetrieveArchiveTasks field if non-nil, zero value otherwise.

### GetRetrieveArchiveTasksOk

`func (o *Recovery) GetRetrieveArchiveTasksOk() (*[]RetrieveArchiveTask, bool)`

GetRetrieveArchiveTasksOk returns a tuple with the RetrieveArchiveTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrieveArchiveTasks

`func (o *Recovery) SetRetrieveArchiveTasks(v []RetrieveArchiveTask)`

SetRetrieveArchiveTasks sets RetrieveArchiveTasks field to given value.

### HasRetrieveArchiveTasks

`func (o *Recovery) HasRetrieveArchiveTasks() bool`

HasRetrieveArchiveTasks returns a boolean if a field has been set.

### SetRetrieveArchiveTasksNil

`func (o *Recovery) SetRetrieveArchiveTasksNil(b bool)`

 SetRetrieveArchiveTasksNil sets the value for RetrieveArchiveTasks to be an explicit nil

### UnsetRetrieveArchiveTasks
`func (o *Recovery) UnsetRetrieveArchiveTasks()`

UnsetRetrieveArchiveTasks ensures that no value is present for RetrieveArchiveTasks, not even an explicit nil
### GetSnapshotEnvironment

`func (o *Recovery) GetSnapshotEnvironment() string`

GetSnapshotEnvironment returns the SnapshotEnvironment field if non-nil, zero value otherwise.

### GetSnapshotEnvironmentOk

`func (o *Recovery) GetSnapshotEnvironmentOk() (*string, bool)`

GetSnapshotEnvironmentOk returns a tuple with the SnapshotEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotEnvironment

`func (o *Recovery) SetSnapshotEnvironment(v string)`

SetSnapshotEnvironment sets SnapshotEnvironment field to given value.

### HasSnapshotEnvironment

`func (o *Recovery) HasSnapshotEnvironment() bool`

HasSnapshotEnvironment returns a boolean if a field has been set.

### GetStartTimeUsecs

`func (o *Recovery) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *Recovery) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *Recovery) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *Recovery) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *Recovery) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *Recovery) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetStatus

`func (o *Recovery) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Recovery) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Recovery) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Recovery) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Recovery) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Recovery) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTearDownMessage

`func (o *Recovery) GetTearDownMessage() string`

GetTearDownMessage returns the TearDownMessage field if non-nil, zero value otherwise.

### GetTearDownMessageOk

`func (o *Recovery) GetTearDownMessageOk() (*string, bool)`

GetTearDownMessageOk returns a tuple with the TearDownMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTearDownMessage

`func (o *Recovery) SetTearDownMessage(v string)`

SetTearDownMessage sets TearDownMessage field to given value.

### HasTearDownMessage

`func (o *Recovery) HasTearDownMessage() bool`

HasTearDownMessage returns a boolean if a field has been set.

### SetTearDownMessageNil

`func (o *Recovery) SetTearDownMessageNil(b bool)`

 SetTearDownMessageNil sets the value for TearDownMessage to be an explicit nil

### UnsetTearDownMessage
`func (o *Recovery) UnsetTearDownMessage()`

UnsetTearDownMessage ensures that no value is present for TearDownMessage, not even an explicit nil
### GetTearDownStatus

`func (o *Recovery) GetTearDownStatus() string`

GetTearDownStatus returns the TearDownStatus field if non-nil, zero value otherwise.

### GetTearDownStatusOk

`func (o *Recovery) GetTearDownStatusOk() (*string, bool)`

GetTearDownStatusOk returns a tuple with the TearDownStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTearDownStatus

`func (o *Recovery) SetTearDownStatus(v string)`

SetTearDownStatus sets TearDownStatus field to given value.

### HasTearDownStatus

`func (o *Recovery) HasTearDownStatus() bool`

HasTearDownStatus returns a boolean if a field has been set.

### SetTearDownStatusNil

`func (o *Recovery) SetTearDownStatusNil(b bool)`

 SetTearDownStatusNil sets the value for TearDownStatus to be an explicit nil

### UnsetTearDownStatus
`func (o *Recovery) UnsetTearDownStatus()`

UnsetTearDownStatus ensures that no value is present for TearDownStatus, not even an explicit nil
### GetAcropolisParams

`func (o *Recovery) GetAcropolisParams() RecoverAcropolisParams`

GetAcropolisParams returns the AcropolisParams field if non-nil, zero value otherwise.

### GetAcropolisParamsOk

`func (o *Recovery) GetAcropolisParamsOk() (*RecoverAcropolisParams, bool)`

GetAcropolisParamsOk returns a tuple with the AcropolisParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcropolisParams

`func (o *Recovery) SetAcropolisParams(v RecoverAcropolisParams)`

SetAcropolisParams sets AcropolisParams field to given value.

### HasAcropolisParams

`func (o *Recovery) HasAcropolisParams() bool`

HasAcropolisParams returns a boolean if a field has been set.

### GetAwsParams

`func (o *Recovery) GetAwsParams() RecoverAwsParams`

GetAwsParams returns the AwsParams field if non-nil, zero value otherwise.

### GetAwsParamsOk

`func (o *Recovery) GetAwsParamsOk() (*RecoverAwsParams, bool)`

GetAwsParamsOk returns a tuple with the AwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsParams

`func (o *Recovery) SetAwsParams(v RecoverAwsParams)`

SetAwsParams sets AwsParams field to given value.

### HasAwsParams

`func (o *Recovery) HasAwsParams() bool`

HasAwsParams returns a boolean if a field has been set.

### GetAzureParams

`func (o *Recovery) GetAzureParams() RecoverAzureParams`

GetAzureParams returns the AzureParams field if non-nil, zero value otherwise.

### GetAzureParamsOk

`func (o *Recovery) GetAzureParamsOk() (*RecoverAzureParams, bool)`

GetAzureParamsOk returns a tuple with the AzureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureParams

`func (o *Recovery) SetAzureParams(v RecoverAzureParams)`

SetAzureParams sets AzureParams field to given value.

### HasAzureParams

`func (o *Recovery) HasAzureParams() bool`

HasAzureParams returns a boolean if a field has been set.

### GetCassandraParams

`func (o *Recovery) GetCassandraParams() CassandraParams`

GetCassandraParams returns the CassandraParams field if non-nil, zero value otherwise.

### GetCassandraParamsOk

`func (o *Recovery) GetCassandraParamsOk() (*CassandraParams, bool)`

GetCassandraParamsOk returns a tuple with the CassandraParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraParams

`func (o *Recovery) SetCassandraParams(v CassandraParams)`

SetCassandraParams sets CassandraParams field to given value.

### HasCassandraParams

`func (o *Recovery) HasCassandraParams() bool`

HasCassandraParams returns a boolean if a field has been set.

### GetCouchbaseParams

`func (o *Recovery) GetCouchbaseParams() CouchbaseParams`

GetCouchbaseParams returns the CouchbaseParams field if non-nil, zero value otherwise.

### GetCouchbaseParamsOk

`func (o *Recovery) GetCouchbaseParamsOk() (*CouchbaseParams, bool)`

GetCouchbaseParamsOk returns a tuple with the CouchbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseParams

`func (o *Recovery) SetCouchbaseParams(v CouchbaseParams)`

SetCouchbaseParams sets CouchbaseParams field to given value.

### HasCouchbaseParams

`func (o *Recovery) HasCouchbaseParams() bool`

HasCouchbaseParams returns a boolean if a field has been set.

### GetElastifileParams

`func (o *Recovery) GetElastifileParams() RecoverElastifileParams`

GetElastifileParams returns the ElastifileParams field if non-nil, zero value otherwise.

### GetElastifileParamsOk

`func (o *Recovery) GetElastifileParamsOk() (*RecoverElastifileParams, bool)`

GetElastifileParamsOk returns a tuple with the ElastifileParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileParams

`func (o *Recovery) SetElastifileParams(v RecoverElastifileParams)`

SetElastifileParams sets ElastifileParams field to given value.

### HasElastifileParams

`func (o *Recovery) HasElastifileParams() bool`

HasElastifileParams returns a boolean if a field has been set.

### GetExchangeParams

`func (o *Recovery) GetExchangeParams() RecoverExchangeParams`

GetExchangeParams returns the ExchangeParams field if non-nil, zero value otherwise.

### GetExchangeParamsOk

`func (o *Recovery) GetExchangeParamsOk() (*RecoverExchangeParams, bool)`

GetExchangeParamsOk returns a tuple with the ExchangeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeParams

`func (o *Recovery) SetExchangeParams(v RecoverExchangeParams)`

SetExchangeParams sets ExchangeParams field to given value.

### HasExchangeParams

`func (o *Recovery) HasExchangeParams() bool`

HasExchangeParams returns a boolean if a field has been set.

### GetExperimentalAdapterParams

`func (o *Recovery) GetExperimentalAdapterParams() ExperimentalAdapterParams`

GetExperimentalAdapterParams returns the ExperimentalAdapterParams field if non-nil, zero value otherwise.

### GetExperimentalAdapterParamsOk

`func (o *Recovery) GetExperimentalAdapterParamsOk() (*ExperimentalAdapterParams, bool)`

GetExperimentalAdapterParamsOk returns a tuple with the ExperimentalAdapterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentalAdapterParams

`func (o *Recovery) SetExperimentalAdapterParams(v ExperimentalAdapterParams)`

SetExperimentalAdapterParams sets ExperimentalAdapterParams field to given value.

### HasExperimentalAdapterParams

`func (o *Recovery) HasExperimentalAdapterParams() bool`

HasExperimentalAdapterParams returns a boolean if a field has been set.

### GetFlashbladeParams

`func (o *Recovery) GetFlashbladeParams() RecoverFlashbladeParams`

GetFlashbladeParams returns the FlashbladeParams field if non-nil, zero value otherwise.

### GetFlashbladeParamsOk

`func (o *Recovery) GetFlashbladeParamsOk() (*RecoverFlashbladeParams, bool)`

GetFlashbladeParamsOk returns a tuple with the FlashbladeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashbladeParams

`func (o *Recovery) SetFlashbladeParams(v RecoverFlashbladeParams)`

SetFlashbladeParams sets FlashbladeParams field to given value.

### HasFlashbladeParams

`func (o *Recovery) HasFlashbladeParams() bool`

HasFlashbladeParams returns a boolean if a field has been set.

### GetGcpParams

`func (o *Recovery) GetGcpParams() RecoverGcpParams`

GetGcpParams returns the GcpParams field if non-nil, zero value otherwise.

### GetGcpParamsOk

`func (o *Recovery) GetGcpParamsOk() (*RecoverGcpParams, bool)`

GetGcpParamsOk returns a tuple with the GcpParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpParams

`func (o *Recovery) SetGcpParams(v RecoverGcpParams)`

SetGcpParams sets GcpParams field to given value.

### HasGcpParams

`func (o *Recovery) HasGcpParams() bool`

HasGcpParams returns a boolean if a field has been set.

### GetGenericNasParams

`func (o *Recovery) GetGenericNasParams() RecoverGenericNasParams`

GetGenericNasParams returns the GenericNasParams field if non-nil, zero value otherwise.

### GetGenericNasParamsOk

`func (o *Recovery) GetGenericNasParamsOk() (*RecoverGenericNasParams, bool)`

GetGenericNasParamsOk returns a tuple with the GenericNasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericNasParams

`func (o *Recovery) SetGenericNasParams(v RecoverGenericNasParams)`

SetGenericNasParams sets GenericNasParams field to given value.

### HasGenericNasParams

`func (o *Recovery) HasGenericNasParams() bool`

HasGenericNasParams returns a boolean if a field has been set.

### GetGpfsParams

`func (o *Recovery) GetGpfsParams() RecoverGpfsParams`

GetGpfsParams returns the GpfsParams field if non-nil, zero value otherwise.

### GetGpfsParamsOk

`func (o *Recovery) GetGpfsParamsOk() (*RecoverGpfsParams, bool)`

GetGpfsParamsOk returns a tuple with the GpfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsParams

`func (o *Recovery) SetGpfsParams(v RecoverGpfsParams)`

SetGpfsParams sets GpfsParams field to given value.

### HasGpfsParams

`func (o *Recovery) HasGpfsParams() bool`

HasGpfsParams returns a boolean if a field has been set.

### GetHbaseParams

`func (o *Recovery) GetHbaseParams() HbaseParams`

GetHbaseParams returns the HbaseParams field if non-nil, zero value otherwise.

### GetHbaseParamsOk

`func (o *Recovery) GetHbaseParamsOk() (*HbaseParams, bool)`

GetHbaseParamsOk returns a tuple with the HbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseParams

`func (o *Recovery) SetHbaseParams(v HbaseParams)`

SetHbaseParams sets HbaseParams field to given value.

### HasHbaseParams

`func (o *Recovery) HasHbaseParams() bool`

HasHbaseParams returns a boolean if a field has been set.

### GetHdfsParams

`func (o *Recovery) GetHdfsParams() HdfsParams`

GetHdfsParams returns the HdfsParams field if non-nil, zero value otherwise.

### GetHdfsParamsOk

`func (o *Recovery) GetHdfsParamsOk() (*HdfsParams, bool)`

GetHdfsParamsOk returns a tuple with the HdfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsParams

`func (o *Recovery) SetHdfsParams(v HdfsParams)`

SetHdfsParams sets HdfsParams field to given value.

### HasHdfsParams

`func (o *Recovery) HasHdfsParams() bool`

HasHdfsParams returns a boolean if a field has been set.

### GetHiveParams

`func (o *Recovery) GetHiveParams() HiveParams`

GetHiveParams returns the HiveParams field if non-nil, zero value otherwise.

### GetHiveParamsOk

`func (o *Recovery) GetHiveParamsOk() (*HiveParams, bool)`

GetHiveParamsOk returns a tuple with the HiveParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveParams

`func (o *Recovery) SetHiveParams(v HiveParams)`

SetHiveParams sets HiveParams field to given value.

### HasHiveParams

`func (o *Recovery) HasHiveParams() bool`

HasHiveParams returns a boolean if a field has been set.

### GetHypervParams

`func (o *Recovery) GetHypervParams() RecoverHyperVParams`

GetHypervParams returns the HypervParams field if non-nil, zero value otherwise.

### GetHypervParamsOk

`func (o *Recovery) GetHypervParamsOk() (*RecoverHyperVParams, bool)`

GetHypervParamsOk returns a tuple with the HypervParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervParams

`func (o *Recovery) SetHypervParams(v RecoverHyperVParams)`

SetHypervParams sets HypervParams field to given value.

### HasHypervParams

`func (o *Recovery) HasHypervParams() bool`

HasHypervParams returns a boolean if a field has been set.

### GetIbmFlashSystemParams

`func (o *Recovery) GetIbmFlashSystemParams() RecoverPureParams`

GetIbmFlashSystemParams returns the IbmFlashSystemParams field if non-nil, zero value otherwise.

### GetIbmFlashSystemParamsOk

`func (o *Recovery) GetIbmFlashSystemParamsOk() (*RecoverPureParams, bool)`

GetIbmFlashSystemParamsOk returns a tuple with the IbmFlashSystemParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbmFlashSystemParams

`func (o *Recovery) SetIbmFlashSystemParams(v RecoverPureParams)`

SetIbmFlashSystemParams sets IbmFlashSystemParams field to given value.

### HasIbmFlashSystemParams

`func (o *Recovery) HasIbmFlashSystemParams() bool`

HasIbmFlashSystemParams returns a boolean if a field has been set.

### GetIsilonParams

`func (o *Recovery) GetIsilonParams() RecoverIsilonParams`

GetIsilonParams returns the IsilonParams field if non-nil, zero value otherwise.

### GetIsilonParamsOk

`func (o *Recovery) GetIsilonParamsOk() (*RecoverIsilonParams, bool)`

GetIsilonParamsOk returns a tuple with the IsilonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonParams

`func (o *Recovery) SetIsilonParams(v RecoverIsilonParams)`

SetIsilonParams sets IsilonParams field to given value.

### HasIsilonParams

`func (o *Recovery) HasIsilonParams() bool`

HasIsilonParams returns a boolean if a field has been set.

### GetKubernetesParams

`func (o *Recovery) GetKubernetesParams() RecoverKubernetesParams`

GetKubernetesParams returns the KubernetesParams field if non-nil, zero value otherwise.

### GetKubernetesParamsOk

`func (o *Recovery) GetKubernetesParamsOk() (*RecoverKubernetesParams, bool)`

GetKubernetesParamsOk returns a tuple with the KubernetesParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesParams

`func (o *Recovery) SetKubernetesParams(v RecoverKubernetesParams)`

SetKubernetesParams sets KubernetesParams field to given value.

### HasKubernetesParams

`func (o *Recovery) HasKubernetesParams() bool`

HasKubernetesParams returns a boolean if a field has been set.

### GetKvmParams

`func (o *Recovery) GetKvmParams() RecoverKvmParams`

GetKvmParams returns the KvmParams field if non-nil, zero value otherwise.

### GetKvmParamsOk

`func (o *Recovery) GetKvmParamsOk() (*RecoverKvmParams, bool)`

GetKvmParamsOk returns a tuple with the KvmParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmParams

`func (o *Recovery) SetKvmParams(v RecoverKvmParams)`

SetKvmParams sets KvmParams field to given value.

### HasKvmParams

`func (o *Recovery) HasKvmParams() bool`

HasKvmParams returns a boolean if a field has been set.

### GetMongodbParams

`func (o *Recovery) GetMongodbParams() MongodbParams`

GetMongodbParams returns the MongodbParams field if non-nil, zero value otherwise.

### GetMongodbParamsOk

`func (o *Recovery) GetMongodbParamsOk() (*MongodbParams, bool)`

GetMongodbParamsOk returns a tuple with the MongodbParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbParams

`func (o *Recovery) SetMongodbParams(v MongodbParams)`

SetMongodbParams sets MongodbParams field to given value.

### HasMongodbParams

`func (o *Recovery) HasMongodbParams() bool`

HasMongodbParams returns a boolean if a field has been set.

### GetMssqlParams

`func (o *Recovery) GetMssqlParams() RecoverSqlParams`

GetMssqlParams returns the MssqlParams field if non-nil, zero value otherwise.

### GetMssqlParamsOk

`func (o *Recovery) GetMssqlParamsOk() (*RecoverSqlParams, bool)`

GetMssqlParamsOk returns a tuple with the MssqlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlParams

`func (o *Recovery) SetMssqlParams(v RecoverSqlParams)`

SetMssqlParams sets MssqlParams field to given value.

### HasMssqlParams

`func (o *Recovery) HasMssqlParams() bool`

HasMssqlParams returns a boolean if a field has been set.

### GetNetappParams

`func (o *Recovery) GetNetappParams() RecoverNetappParams`

GetNetappParams returns the NetappParams field if non-nil, zero value otherwise.

### GetNetappParamsOk

`func (o *Recovery) GetNetappParamsOk() (*RecoverNetappParams, bool)`

GetNetappParamsOk returns a tuple with the NetappParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappParams

`func (o *Recovery) SetNetappParams(v RecoverNetappParams)`

SetNetappParams sets NetappParams field to given value.

### HasNetappParams

`func (o *Recovery) HasNetappParams() bool`

HasNetappParams returns a boolean if a field has been set.

### GetOffice365Params

`func (o *Recovery) GetOffice365Params() RecoverO365Params`

GetOffice365Params returns the Office365Params field if non-nil, zero value otherwise.

### GetOffice365ParamsOk

`func (o *Recovery) GetOffice365ParamsOk() (*RecoverO365Params, bool)`

GetOffice365ParamsOk returns a tuple with the Office365Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365Params

`func (o *Recovery) SetOffice365Params(v RecoverO365Params)`

SetOffice365Params sets Office365Params field to given value.

### HasOffice365Params

`func (o *Recovery) HasOffice365Params() bool`

HasOffice365Params returns a boolean if a field has been set.

### GetOracleParams

`func (o *Recovery) GetOracleParams() RecoverOracleParams`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *Recovery) GetOracleParamsOk() (*RecoverOracleParams, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *Recovery) SetOracleParams(v RecoverOracleParams)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *Recovery) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.

### GetPhysicalParams

`func (o *Recovery) GetPhysicalParams() RecoverPhysicalParams`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *Recovery) GetPhysicalParamsOk() (*RecoverPhysicalParams, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *Recovery) SetPhysicalParams(v RecoverPhysicalParams)`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *Recovery) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.

### GetPureParams

`func (o *Recovery) GetPureParams() RecoverPureParams`

GetPureParams returns the PureParams field if non-nil, zero value otherwise.

### GetPureParamsOk

`func (o *Recovery) GetPureParamsOk() (*RecoverPureParams, bool)`

GetPureParamsOk returns a tuple with the PureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPureParams

`func (o *Recovery) SetPureParams(v RecoverPureParams)`

SetPureParams sets PureParams field to given value.

### HasPureParams

`func (o *Recovery) HasPureParams() bool`

HasPureParams returns a boolean if a field has been set.

### GetSapHanaParams

`func (o *Recovery) GetSapHanaParams() SapHanaParams`

GetSapHanaParams returns the SapHanaParams field if non-nil, zero value otherwise.

### GetSapHanaParamsOk

`func (o *Recovery) GetSapHanaParamsOk() (*SapHanaParams, bool)`

GetSapHanaParamsOk returns a tuple with the SapHanaParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSapHanaParams

`func (o *Recovery) SetSapHanaParams(v SapHanaParams)`

SetSapHanaParams sets SapHanaParams field to given value.

### HasSapHanaParams

`func (o *Recovery) HasSapHanaParams() bool`

HasSapHanaParams returns a boolean if a field has been set.

### GetSfdcParams

`func (o *Recovery) GetSfdcParams() RecoverSalesforceParams`

GetSfdcParams returns the SfdcParams field if non-nil, zero value otherwise.

### GetSfdcParamsOk

`func (o *Recovery) GetSfdcParamsOk() (*RecoverSalesforceParams, bool)`

GetSfdcParamsOk returns a tuple with the SfdcParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfdcParams

`func (o *Recovery) SetSfdcParams(v RecoverSalesforceParams)`

SetSfdcParams sets SfdcParams field to given value.

### HasSfdcParams

`func (o *Recovery) HasSfdcParams() bool`

HasSfdcParams returns a boolean if a field has been set.

### GetUdaParams

`func (o *Recovery) GetUdaParams() UdaParams`

GetUdaParams returns the UdaParams field if non-nil, zero value otherwise.

### GetUdaParamsOk

`func (o *Recovery) GetUdaParamsOk() (*UdaParams, bool)`

GetUdaParamsOk returns a tuple with the UdaParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdaParams

`func (o *Recovery) SetUdaParams(v UdaParams)`

SetUdaParams sets UdaParams field to given value.

### HasUdaParams

`func (o *Recovery) HasUdaParams() bool`

HasUdaParams returns a boolean if a field has been set.

### GetViewParams

`func (o *Recovery) GetViewParams() RecoverViewParams`

GetViewParams returns the ViewParams field if non-nil, zero value otherwise.

### GetViewParamsOk

`func (o *Recovery) GetViewParamsOk() (*RecoverViewParams, bool)`

GetViewParamsOk returns a tuple with the ViewParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewParams

`func (o *Recovery) SetViewParams(v RecoverViewParams)`

SetViewParams sets ViewParams field to given value.

### HasViewParams

`func (o *Recovery) HasViewParams() bool`

HasViewParams returns a boolean if a field has been set.

### GetVmwareParams

`func (o *Recovery) GetVmwareParams() RecoverVmwareParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *Recovery) GetVmwareParamsOk() (*RecoverVmwareParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *Recovery) SetVmwareParams(v RecoverVmwareParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *Recovery) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


