# IsilonObjectProtectionResponseParams

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
**ContinuousSnapshots** | Pointer to [**ContinuousSnapshotParams**](ContinuousSnapshotParams.md) |  | [optional] 
**Protocol** | Pointer to **NullableString** | Specifies the protocol of the NAS device being backed up. | [optional] 
**UseChangelist** | Pointer to **NullableBool** | Specify whether to use the Isilon Changelist API to directly discover changed files/directories for faster incremental backup. Cohesity will keep an extra snapshot which will be deleted by the next successful backup. | [optional] 

## Methods

### NewIsilonObjectProtectionResponseParams

`func NewIsilonObjectProtectionResponseParams() *IsilonObjectProtectionResponseParams`

NewIsilonObjectProtectionResponseParams instantiates a new IsilonObjectProtectionResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIsilonObjectProtectionResponseParamsWithDefaults

`func NewIsilonObjectProtectionResponseParamsWithDefaults() *IsilonObjectProtectionResponseParams`

NewIsilonObjectProtectionResponseParamsWithDefaults instantiates a new IsilonObjectProtectionResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *IsilonObjectProtectionResponseParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *IsilonObjectProtectionResponseParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *IsilonObjectProtectionResponseParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *IsilonObjectProtectionResponseParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *IsilonObjectProtectionResponseParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *IsilonObjectProtectionResponseParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetEncryptionEnabled

`func (o *IsilonObjectProtectionResponseParams) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *IsilonObjectProtectionResponseParams) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *IsilonObjectProtectionResponseParams) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *IsilonObjectProtectionResponseParams) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### SetEncryptionEnabledNil

`func (o *IsilonObjectProtectionResponseParams) SetEncryptionEnabledNil(b bool)`

 SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil

### UnsetEncryptionEnabled
`func (o *IsilonObjectProtectionResponseParams) UnsetEncryptionEnabled()`

UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
### GetFileFilters

`func (o *IsilonObjectProtectionResponseParams) GetFileFilters() FileFilteringPolicy`

GetFileFilters returns the FileFilters field if non-nil, zero value otherwise.

### GetFileFiltersOk

`func (o *IsilonObjectProtectionResponseParams) GetFileFiltersOk() (*FileFilteringPolicy, bool)`

GetFileFiltersOk returns a tuple with the FileFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFilters

`func (o *IsilonObjectProtectionResponseParams) SetFileFilters(v FileFilteringPolicy)`

SetFileFilters sets FileFilters field to given value.

### HasFileFilters

`func (o *IsilonObjectProtectionResponseParams) HasFileFilters() bool`

HasFileFilters returns a boolean if a field has been set.

### GetFileLockConfig

`func (o *IsilonObjectProtectionResponseParams) GetFileLockConfig() FileLevelDataLockConfig`

GetFileLockConfig returns the FileLockConfig field if non-nil, zero value otherwise.

### GetFileLockConfigOk

`func (o *IsilonObjectProtectionResponseParams) GetFileLockConfigOk() (*FileLevelDataLockConfig, bool)`

GetFileLockConfigOk returns a tuple with the FileLockConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLockConfig

`func (o *IsilonObjectProtectionResponseParams) SetFileLockConfig(v FileLevelDataLockConfig)`

SetFileLockConfig sets FileLockConfig field to given value.

### HasFileLockConfig

`func (o *IsilonObjectProtectionResponseParams) HasFileLockConfig() bool`

HasFileLockConfig returns a boolean if a field has been set.

### GetIndexingPolicy

`func (o *IsilonObjectProtectionResponseParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *IsilonObjectProtectionResponseParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *IsilonObjectProtectionResponseParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *IsilonObjectProtectionResponseParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetPrePostScript

`func (o *IsilonObjectProtectionResponseParams) GetPrePostScript() HostBasedBackupScriptParams`

GetPrePostScript returns the PrePostScript field if non-nil, zero value otherwise.

### GetPrePostScriptOk

`func (o *IsilonObjectProtectionResponseParams) GetPrePostScriptOk() (*HostBasedBackupScriptParams, bool)`

GetPrePostScriptOk returns a tuple with the PrePostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePostScript

`func (o *IsilonObjectProtectionResponseParams) SetPrePostScript(v HostBasedBackupScriptParams)`

SetPrePostScript sets PrePostScript field to given value.

### HasPrePostScript

`func (o *IsilonObjectProtectionResponseParams) HasPrePostScript() bool`

HasPrePostScript returns a boolean if a field has been set.

### GetThrottlingConfig

`func (o *IsilonObjectProtectionResponseParams) GetThrottlingConfig() NasThrottlingConfig`

GetThrottlingConfig returns the ThrottlingConfig field if non-nil, zero value otherwise.

### GetThrottlingConfigOk

`func (o *IsilonObjectProtectionResponseParams) GetThrottlingConfigOk() (*NasThrottlingConfig, bool)`

GetThrottlingConfigOk returns a tuple with the ThrottlingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingConfig

`func (o *IsilonObjectProtectionResponseParams) SetThrottlingConfig(v NasThrottlingConfig)`

SetThrottlingConfig sets ThrottlingConfig field to given value.

### HasThrottlingConfig

`func (o *IsilonObjectProtectionResponseParams) HasThrottlingConfig() bool`

HasThrottlingConfig returns a boolean if a field has been set.

### GetContinuousSnapshots

`func (o *IsilonObjectProtectionResponseParams) GetContinuousSnapshots() ContinuousSnapshotParams`

GetContinuousSnapshots returns the ContinuousSnapshots field if non-nil, zero value otherwise.

### GetContinuousSnapshotsOk

`func (o *IsilonObjectProtectionResponseParams) GetContinuousSnapshotsOk() (*ContinuousSnapshotParams, bool)`

GetContinuousSnapshotsOk returns a tuple with the ContinuousSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuousSnapshots

`func (o *IsilonObjectProtectionResponseParams) SetContinuousSnapshots(v ContinuousSnapshotParams)`

SetContinuousSnapshots sets ContinuousSnapshots field to given value.

### HasContinuousSnapshots

`func (o *IsilonObjectProtectionResponseParams) HasContinuousSnapshots() bool`

HasContinuousSnapshots returns a boolean if a field has been set.

### GetProtocol

`func (o *IsilonObjectProtectionResponseParams) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *IsilonObjectProtectionResponseParams) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *IsilonObjectProtectionResponseParams) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *IsilonObjectProtectionResponseParams) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *IsilonObjectProtectionResponseParams) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *IsilonObjectProtectionResponseParams) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetUseChangelist

`func (o *IsilonObjectProtectionResponseParams) GetUseChangelist() bool`

GetUseChangelist returns the UseChangelist field if non-nil, zero value otherwise.

### GetUseChangelistOk

`func (o *IsilonObjectProtectionResponseParams) GetUseChangelistOk() (*bool, bool)`

GetUseChangelistOk returns a tuple with the UseChangelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseChangelist

`func (o *IsilonObjectProtectionResponseParams) SetUseChangelist(v bool)`

SetUseChangelist sets UseChangelist field to given value.

### HasUseChangelist

`func (o *IsilonObjectProtectionResponseParams) HasUseChangelist() bool`

HasUseChangelist returns a boolean if a field has been set.

### SetUseChangelistNil

`func (o *IsilonObjectProtectionResponseParams) SetUseChangelistNil(b bool)`

 SetUseChangelistNil sets the value for UseChangelist to be an explicit nil

### UnsetUseChangelist
`func (o *IsilonObjectProtectionResponseParams) UnsetUseChangelist()`

UnsetUseChangelist ensures that no value is present for UseChangelist, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


