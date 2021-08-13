# IsilonObjectProtectionUpdateRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether or not the backup should continue regardless of whether or not an error was encountered. | [optional] 
**EncryptionEnabled** | Pointer to **NullableBool** | Specifies whether the encryption should be used while backup or not. | [optional] 
**FileLockConfig** | Pointer to [**FileLevelDataLockConfig**](FileLevelDataLockConfig.md) |  | [optional] 
**FileFilters** | Pointer to [**FileFilteringPolicy**](FileFilteringPolicy.md) |  | [optional] 
**PrePostScript** | Pointer to [**HostBasedBackupScriptParams**](HostBasedBackupScriptParams.md) |  | [optional] 
**ThrottlingConfig** | Pointer to [**NasThrottlingConfig**](NasThrottlingConfig.md) |  | [optional] 
**Protocol** | Pointer to **NullableString** | Specifies the protocol of the NAS device being backed up. | [optional] 
**ContinuousSnapshots** | Pointer to [**ContinuousSnapshotParams**](ContinuousSnapshotParams.md) |  | [optional] 
**UseChangelist** | Pointer to **NullableBool** | Specify whether to use the Isilon Changelist API to directly discover changed files/directories for faster incremental backup. Cohesity will keep an extra snapshot which will be deleted by the next successful backup. | [optional] 

## Methods

### NewIsilonObjectProtectionUpdateRequestParams

`func NewIsilonObjectProtectionUpdateRequestParams() *IsilonObjectProtectionUpdateRequestParams`

NewIsilonObjectProtectionUpdateRequestParams instantiates a new IsilonObjectProtectionUpdateRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIsilonObjectProtectionUpdateRequestParamsWithDefaults

`func NewIsilonObjectProtectionUpdateRequestParamsWithDefaults() *IsilonObjectProtectionUpdateRequestParams`

NewIsilonObjectProtectionUpdateRequestParamsWithDefaults instantiates a new IsilonObjectProtectionUpdateRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexingPolicy

`func (o *IsilonObjectProtectionUpdateRequestParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *IsilonObjectProtectionUpdateRequestParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *IsilonObjectProtectionUpdateRequestParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *IsilonObjectProtectionUpdateRequestParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetContinueOnError

`func (o *IsilonObjectProtectionUpdateRequestParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *IsilonObjectProtectionUpdateRequestParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *IsilonObjectProtectionUpdateRequestParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *IsilonObjectProtectionUpdateRequestParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *IsilonObjectProtectionUpdateRequestParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *IsilonObjectProtectionUpdateRequestParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetEncryptionEnabled

`func (o *IsilonObjectProtectionUpdateRequestParams) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *IsilonObjectProtectionUpdateRequestParams) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *IsilonObjectProtectionUpdateRequestParams) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *IsilonObjectProtectionUpdateRequestParams) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### SetEncryptionEnabledNil

`func (o *IsilonObjectProtectionUpdateRequestParams) SetEncryptionEnabledNil(b bool)`

 SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil

### UnsetEncryptionEnabled
`func (o *IsilonObjectProtectionUpdateRequestParams) UnsetEncryptionEnabled()`

UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
### GetFileLockConfig

`func (o *IsilonObjectProtectionUpdateRequestParams) GetFileLockConfig() FileLevelDataLockConfig`

GetFileLockConfig returns the FileLockConfig field if non-nil, zero value otherwise.

### GetFileLockConfigOk

`func (o *IsilonObjectProtectionUpdateRequestParams) GetFileLockConfigOk() (*FileLevelDataLockConfig, bool)`

GetFileLockConfigOk returns a tuple with the FileLockConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLockConfig

`func (o *IsilonObjectProtectionUpdateRequestParams) SetFileLockConfig(v FileLevelDataLockConfig)`

SetFileLockConfig sets FileLockConfig field to given value.

### HasFileLockConfig

`func (o *IsilonObjectProtectionUpdateRequestParams) HasFileLockConfig() bool`

HasFileLockConfig returns a boolean if a field has been set.

### GetFileFilters

`func (o *IsilonObjectProtectionUpdateRequestParams) GetFileFilters() FileFilteringPolicy`

GetFileFilters returns the FileFilters field if non-nil, zero value otherwise.

### GetFileFiltersOk

`func (o *IsilonObjectProtectionUpdateRequestParams) GetFileFiltersOk() (*FileFilteringPolicy, bool)`

GetFileFiltersOk returns a tuple with the FileFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFilters

`func (o *IsilonObjectProtectionUpdateRequestParams) SetFileFilters(v FileFilteringPolicy)`

SetFileFilters sets FileFilters field to given value.

### HasFileFilters

`func (o *IsilonObjectProtectionUpdateRequestParams) HasFileFilters() bool`

HasFileFilters returns a boolean if a field has been set.

### GetPrePostScript

`func (o *IsilonObjectProtectionUpdateRequestParams) GetPrePostScript() HostBasedBackupScriptParams`

GetPrePostScript returns the PrePostScript field if non-nil, zero value otherwise.

### GetPrePostScriptOk

`func (o *IsilonObjectProtectionUpdateRequestParams) GetPrePostScriptOk() (*HostBasedBackupScriptParams, bool)`

GetPrePostScriptOk returns a tuple with the PrePostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePostScript

`func (o *IsilonObjectProtectionUpdateRequestParams) SetPrePostScript(v HostBasedBackupScriptParams)`

SetPrePostScript sets PrePostScript field to given value.

### HasPrePostScript

`func (o *IsilonObjectProtectionUpdateRequestParams) HasPrePostScript() bool`

HasPrePostScript returns a boolean if a field has been set.

### GetThrottlingConfig

`func (o *IsilonObjectProtectionUpdateRequestParams) GetThrottlingConfig() NasThrottlingConfig`

GetThrottlingConfig returns the ThrottlingConfig field if non-nil, zero value otherwise.

### GetThrottlingConfigOk

`func (o *IsilonObjectProtectionUpdateRequestParams) GetThrottlingConfigOk() (*NasThrottlingConfig, bool)`

GetThrottlingConfigOk returns a tuple with the ThrottlingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingConfig

`func (o *IsilonObjectProtectionUpdateRequestParams) SetThrottlingConfig(v NasThrottlingConfig)`

SetThrottlingConfig sets ThrottlingConfig field to given value.

### HasThrottlingConfig

`func (o *IsilonObjectProtectionUpdateRequestParams) HasThrottlingConfig() bool`

HasThrottlingConfig returns a boolean if a field has been set.

### GetProtocol

`func (o *IsilonObjectProtectionUpdateRequestParams) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *IsilonObjectProtectionUpdateRequestParams) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *IsilonObjectProtectionUpdateRequestParams) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *IsilonObjectProtectionUpdateRequestParams) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *IsilonObjectProtectionUpdateRequestParams) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *IsilonObjectProtectionUpdateRequestParams) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetContinuousSnapshots

`func (o *IsilonObjectProtectionUpdateRequestParams) GetContinuousSnapshots() ContinuousSnapshotParams`

GetContinuousSnapshots returns the ContinuousSnapshots field if non-nil, zero value otherwise.

### GetContinuousSnapshotsOk

`func (o *IsilonObjectProtectionUpdateRequestParams) GetContinuousSnapshotsOk() (*ContinuousSnapshotParams, bool)`

GetContinuousSnapshotsOk returns a tuple with the ContinuousSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuousSnapshots

`func (o *IsilonObjectProtectionUpdateRequestParams) SetContinuousSnapshots(v ContinuousSnapshotParams)`

SetContinuousSnapshots sets ContinuousSnapshots field to given value.

### HasContinuousSnapshots

`func (o *IsilonObjectProtectionUpdateRequestParams) HasContinuousSnapshots() bool`

HasContinuousSnapshots returns a boolean if a field has been set.

### GetUseChangelist

`func (o *IsilonObjectProtectionUpdateRequestParams) GetUseChangelist() bool`

GetUseChangelist returns the UseChangelist field if non-nil, zero value otherwise.

### GetUseChangelistOk

`func (o *IsilonObjectProtectionUpdateRequestParams) GetUseChangelistOk() (*bool, bool)`

GetUseChangelistOk returns a tuple with the UseChangelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseChangelist

`func (o *IsilonObjectProtectionUpdateRequestParams) SetUseChangelist(v bool)`

SetUseChangelist sets UseChangelist field to given value.

### HasUseChangelist

`func (o *IsilonObjectProtectionUpdateRequestParams) HasUseChangelist() bool`

HasUseChangelist returns a boolean if a field has been set.

### SetUseChangelistNil

`func (o *IsilonObjectProtectionUpdateRequestParams) SetUseChangelistNil(b bool)`

 SetUseChangelistNil sets the value for UseChangelist to be an explicit nil

### UnsetUseChangelist
`func (o *IsilonObjectProtectionUpdateRequestParams) UnsetUseChangelist()`

UnsetUseChangelist ensures that no value is present for UseChangelist, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


