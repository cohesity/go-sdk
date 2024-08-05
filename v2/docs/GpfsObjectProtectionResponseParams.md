# GpfsObjectProtectionResponseParams

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
**Protocol** | Pointer to **NullableString** | Specifies the protocol of the NAS device being backed up. | [optional] 

## Methods

### NewGpfsObjectProtectionResponseParams

`func NewGpfsObjectProtectionResponseParams() *GpfsObjectProtectionResponseParams`

NewGpfsObjectProtectionResponseParams instantiates a new GpfsObjectProtectionResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpfsObjectProtectionResponseParamsWithDefaults

`func NewGpfsObjectProtectionResponseParamsWithDefaults() *GpfsObjectProtectionResponseParams`

NewGpfsObjectProtectionResponseParamsWithDefaults instantiates a new GpfsObjectProtectionResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *GpfsObjectProtectionResponseParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *GpfsObjectProtectionResponseParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *GpfsObjectProtectionResponseParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *GpfsObjectProtectionResponseParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *GpfsObjectProtectionResponseParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *GpfsObjectProtectionResponseParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetEncryptionEnabled

`func (o *GpfsObjectProtectionResponseParams) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *GpfsObjectProtectionResponseParams) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *GpfsObjectProtectionResponseParams) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *GpfsObjectProtectionResponseParams) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### SetEncryptionEnabledNil

`func (o *GpfsObjectProtectionResponseParams) SetEncryptionEnabledNil(b bool)`

 SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil

### UnsetEncryptionEnabled
`func (o *GpfsObjectProtectionResponseParams) UnsetEncryptionEnabled()`

UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
### GetFileFilters

`func (o *GpfsObjectProtectionResponseParams) GetFileFilters() FileFilteringPolicy`

GetFileFilters returns the FileFilters field if non-nil, zero value otherwise.

### GetFileFiltersOk

`func (o *GpfsObjectProtectionResponseParams) GetFileFiltersOk() (*FileFilteringPolicy, bool)`

GetFileFiltersOk returns a tuple with the FileFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFilters

`func (o *GpfsObjectProtectionResponseParams) SetFileFilters(v FileFilteringPolicy)`

SetFileFilters sets FileFilters field to given value.

### HasFileFilters

`func (o *GpfsObjectProtectionResponseParams) HasFileFilters() bool`

HasFileFilters returns a boolean if a field has been set.

### GetFileLockConfig

`func (o *GpfsObjectProtectionResponseParams) GetFileLockConfig() FileLevelDataLockConfig`

GetFileLockConfig returns the FileLockConfig field if non-nil, zero value otherwise.

### GetFileLockConfigOk

`func (o *GpfsObjectProtectionResponseParams) GetFileLockConfigOk() (*FileLevelDataLockConfig, bool)`

GetFileLockConfigOk returns a tuple with the FileLockConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLockConfig

`func (o *GpfsObjectProtectionResponseParams) SetFileLockConfig(v FileLevelDataLockConfig)`

SetFileLockConfig sets FileLockConfig field to given value.

### HasFileLockConfig

`func (o *GpfsObjectProtectionResponseParams) HasFileLockConfig() bool`

HasFileLockConfig returns a boolean if a field has been set.

### GetIndexingPolicy

`func (o *GpfsObjectProtectionResponseParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *GpfsObjectProtectionResponseParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *GpfsObjectProtectionResponseParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *GpfsObjectProtectionResponseParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetPrePostScript

`func (o *GpfsObjectProtectionResponseParams) GetPrePostScript() HostBasedBackupScriptParams`

GetPrePostScript returns the PrePostScript field if non-nil, zero value otherwise.

### GetPrePostScriptOk

`func (o *GpfsObjectProtectionResponseParams) GetPrePostScriptOk() (*HostBasedBackupScriptParams, bool)`

GetPrePostScriptOk returns a tuple with the PrePostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePostScript

`func (o *GpfsObjectProtectionResponseParams) SetPrePostScript(v HostBasedBackupScriptParams)`

SetPrePostScript sets PrePostScript field to given value.

### HasPrePostScript

`func (o *GpfsObjectProtectionResponseParams) HasPrePostScript() bool`

HasPrePostScript returns a boolean if a field has been set.

### GetThrottlingConfig

`func (o *GpfsObjectProtectionResponseParams) GetThrottlingConfig() NasThrottlingConfig`

GetThrottlingConfig returns the ThrottlingConfig field if non-nil, zero value otherwise.

### GetThrottlingConfigOk

`func (o *GpfsObjectProtectionResponseParams) GetThrottlingConfigOk() (*NasThrottlingConfig, bool)`

GetThrottlingConfigOk returns a tuple with the ThrottlingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingConfig

`func (o *GpfsObjectProtectionResponseParams) SetThrottlingConfig(v NasThrottlingConfig)`

SetThrottlingConfig sets ThrottlingConfig field to given value.

### HasThrottlingConfig

`func (o *GpfsObjectProtectionResponseParams) HasThrottlingConfig() bool`

HasThrottlingConfig returns a boolean if a field has been set.

### GetProtocol

`func (o *GpfsObjectProtectionResponseParams) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GpfsObjectProtectionResponseParams) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GpfsObjectProtectionResponseParams) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GpfsObjectProtectionResponseParams) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *GpfsObjectProtectionResponseParams) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *GpfsObjectProtectionResponseParams) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


