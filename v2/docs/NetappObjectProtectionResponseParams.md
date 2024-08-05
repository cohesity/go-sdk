# NetappObjectProtectionResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether or not the backup should continue regardless of whether or not an error was encountered. | [optional] 
**EncryptionEnabled** | Pointer to **NullableBool** | Specifies whether the encryption should be used while backup or not. | [optional] 
**FileFilters** | Pointer to [**FileFilteringPolicy**](FileFilteringPolicy.md) |  | [optional] 
**FileLockConfig** | Pointer to [**FileLevelDataLockConfig**](FileLevelDataLockConfig.md) |  | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**PrePostScript** | Pointer to [**HostBasedBackupScriptParams**](HostBasedBackupScriptParams.md) |  | [optional] 
**ThrottlingConfig** | Pointer to [**NasThrottlingConfig**](NasThrottlingConfig.md) |  | [optional] 
**BackupExistingSnapshot** | Pointer to **NullableBool** | Specifies that snapshot label is not set for Data-Protect Netapp Volumes backup. If field is set to true, existing oldest snapshot is used for backup and subsequent incremental will be selected in ascending order of snapshot create time on the source. If snapshot label is set, this field is set to false. | [optional] 
**ContinuousSnapshots** | Pointer to [**ContinuousSnapshotParams**](ContinuousSnapshotParams.md) |  | [optional] 
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the objects to be excluded in the Protection. | [optional] 
**Protocol** | Pointer to **NullableString** | Specifies the protocol of the NAS device being backed up. | [optional] 
**SnapshotLabel** | Pointer to [**SnapshotLabel**](SnapshotLabel.md) |  | [optional] 

## Methods

### NewNetappObjectProtectionResponseParams

`func NewNetappObjectProtectionResponseParams() *NetappObjectProtectionResponseParams`

NewNetappObjectProtectionResponseParams instantiates a new NetappObjectProtectionResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetappObjectProtectionResponseParamsWithDefaults

`func NewNetappObjectProtectionResponseParamsWithDefaults() *NetappObjectProtectionResponseParams`

NewNetappObjectProtectionResponseParamsWithDefaults instantiates a new NetappObjectProtectionResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *NetappObjectProtectionResponseParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *NetappObjectProtectionResponseParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *NetappObjectProtectionResponseParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *NetappObjectProtectionResponseParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *NetappObjectProtectionResponseParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *NetappObjectProtectionResponseParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetEncryptionEnabled

`func (o *NetappObjectProtectionResponseParams) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *NetappObjectProtectionResponseParams) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *NetappObjectProtectionResponseParams) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *NetappObjectProtectionResponseParams) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### SetEncryptionEnabledNil

`func (o *NetappObjectProtectionResponseParams) SetEncryptionEnabledNil(b bool)`

 SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil

### UnsetEncryptionEnabled
`func (o *NetappObjectProtectionResponseParams) UnsetEncryptionEnabled()`

UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
### GetFileFilters

`func (o *NetappObjectProtectionResponseParams) GetFileFilters() FileFilteringPolicy`

GetFileFilters returns the FileFilters field if non-nil, zero value otherwise.

### GetFileFiltersOk

`func (o *NetappObjectProtectionResponseParams) GetFileFiltersOk() (*FileFilteringPolicy, bool)`

GetFileFiltersOk returns a tuple with the FileFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFilters

`func (o *NetappObjectProtectionResponseParams) SetFileFilters(v FileFilteringPolicy)`

SetFileFilters sets FileFilters field to given value.

### HasFileFilters

`func (o *NetappObjectProtectionResponseParams) HasFileFilters() bool`

HasFileFilters returns a boolean if a field has been set.

### GetFileLockConfig

`func (o *NetappObjectProtectionResponseParams) GetFileLockConfig() FileLevelDataLockConfig`

GetFileLockConfig returns the FileLockConfig field if non-nil, zero value otherwise.

### GetFileLockConfigOk

`func (o *NetappObjectProtectionResponseParams) GetFileLockConfigOk() (*FileLevelDataLockConfig, bool)`

GetFileLockConfigOk returns a tuple with the FileLockConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLockConfig

`func (o *NetappObjectProtectionResponseParams) SetFileLockConfig(v FileLevelDataLockConfig)`

SetFileLockConfig sets FileLockConfig field to given value.

### HasFileLockConfig

`func (o *NetappObjectProtectionResponseParams) HasFileLockConfig() bool`

HasFileLockConfig returns a boolean if a field has been set.

### GetIndexingPolicy

`func (o *NetappObjectProtectionResponseParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *NetappObjectProtectionResponseParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *NetappObjectProtectionResponseParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *NetappObjectProtectionResponseParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetPrePostScript

`func (o *NetappObjectProtectionResponseParams) GetPrePostScript() HostBasedBackupScriptParams`

GetPrePostScript returns the PrePostScript field if non-nil, zero value otherwise.

### GetPrePostScriptOk

`func (o *NetappObjectProtectionResponseParams) GetPrePostScriptOk() (*HostBasedBackupScriptParams, bool)`

GetPrePostScriptOk returns a tuple with the PrePostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePostScript

`func (o *NetappObjectProtectionResponseParams) SetPrePostScript(v HostBasedBackupScriptParams)`

SetPrePostScript sets PrePostScript field to given value.

### HasPrePostScript

`func (o *NetappObjectProtectionResponseParams) HasPrePostScript() bool`

HasPrePostScript returns a boolean if a field has been set.

### GetThrottlingConfig

`func (o *NetappObjectProtectionResponseParams) GetThrottlingConfig() NasThrottlingConfig`

GetThrottlingConfig returns the ThrottlingConfig field if non-nil, zero value otherwise.

### GetThrottlingConfigOk

`func (o *NetappObjectProtectionResponseParams) GetThrottlingConfigOk() (*NasThrottlingConfig, bool)`

GetThrottlingConfigOk returns a tuple with the ThrottlingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingConfig

`func (o *NetappObjectProtectionResponseParams) SetThrottlingConfig(v NasThrottlingConfig)`

SetThrottlingConfig sets ThrottlingConfig field to given value.

### HasThrottlingConfig

`func (o *NetappObjectProtectionResponseParams) HasThrottlingConfig() bool`

HasThrottlingConfig returns a boolean if a field has been set.

### GetBackupExistingSnapshot

`func (o *NetappObjectProtectionResponseParams) GetBackupExistingSnapshot() bool`

GetBackupExistingSnapshot returns the BackupExistingSnapshot field if non-nil, zero value otherwise.

### GetBackupExistingSnapshotOk

`func (o *NetappObjectProtectionResponseParams) GetBackupExistingSnapshotOk() (*bool, bool)`

GetBackupExistingSnapshotOk returns a tuple with the BackupExistingSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupExistingSnapshot

`func (o *NetappObjectProtectionResponseParams) SetBackupExistingSnapshot(v bool)`

SetBackupExistingSnapshot sets BackupExistingSnapshot field to given value.

### HasBackupExistingSnapshot

`func (o *NetappObjectProtectionResponseParams) HasBackupExistingSnapshot() bool`

HasBackupExistingSnapshot returns a boolean if a field has been set.

### SetBackupExistingSnapshotNil

`func (o *NetappObjectProtectionResponseParams) SetBackupExistingSnapshotNil(b bool)`

 SetBackupExistingSnapshotNil sets the value for BackupExistingSnapshot to be an explicit nil

### UnsetBackupExistingSnapshot
`func (o *NetappObjectProtectionResponseParams) UnsetBackupExistingSnapshot()`

UnsetBackupExistingSnapshot ensures that no value is present for BackupExistingSnapshot, not even an explicit nil
### GetContinuousSnapshots

`func (o *NetappObjectProtectionResponseParams) GetContinuousSnapshots() ContinuousSnapshotParams`

GetContinuousSnapshots returns the ContinuousSnapshots field if non-nil, zero value otherwise.

### GetContinuousSnapshotsOk

`func (o *NetappObjectProtectionResponseParams) GetContinuousSnapshotsOk() (*ContinuousSnapshotParams, bool)`

GetContinuousSnapshotsOk returns a tuple with the ContinuousSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuousSnapshots

`func (o *NetappObjectProtectionResponseParams) SetContinuousSnapshots(v ContinuousSnapshotParams)`

SetContinuousSnapshots sets ContinuousSnapshots field to given value.

### HasContinuousSnapshots

`func (o *NetappObjectProtectionResponseParams) HasContinuousSnapshots() bool`

HasContinuousSnapshots returns a boolean if a field has been set.

### GetExcludeObjectIds

`func (o *NetappObjectProtectionResponseParams) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *NetappObjectProtectionResponseParams) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *NetappObjectProtectionResponseParams) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *NetappObjectProtectionResponseParams) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### SetExcludeObjectIdsNil

`func (o *NetappObjectProtectionResponseParams) SetExcludeObjectIdsNil(b bool)`

 SetExcludeObjectIdsNil sets the value for ExcludeObjectIds to be an explicit nil

### UnsetExcludeObjectIds
`func (o *NetappObjectProtectionResponseParams) UnsetExcludeObjectIds()`

UnsetExcludeObjectIds ensures that no value is present for ExcludeObjectIds, not even an explicit nil
### GetProtocol

`func (o *NetappObjectProtectionResponseParams) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NetappObjectProtectionResponseParams) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NetappObjectProtectionResponseParams) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *NetappObjectProtectionResponseParams) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *NetappObjectProtectionResponseParams) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *NetappObjectProtectionResponseParams) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetSnapshotLabel

`func (o *NetappObjectProtectionResponseParams) GetSnapshotLabel() SnapshotLabel`

GetSnapshotLabel returns the SnapshotLabel field if non-nil, zero value otherwise.

### GetSnapshotLabelOk

`func (o *NetappObjectProtectionResponseParams) GetSnapshotLabelOk() (*SnapshotLabel, bool)`

GetSnapshotLabelOk returns a tuple with the SnapshotLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotLabel

`func (o *NetappObjectProtectionResponseParams) SetSnapshotLabel(v SnapshotLabel)`

SetSnapshotLabel sets SnapshotLabel field to given value.

### HasSnapshotLabel

`func (o *NetappObjectProtectionResponseParams) HasSnapshotLabel() bool`

HasSnapshotLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


