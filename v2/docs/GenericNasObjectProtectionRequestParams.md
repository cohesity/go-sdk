# GenericNasObjectProtectionRequestParams

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
**Objects** | [**[]ProtectionObjectInput**](ProtectionObjectInput.md) | Specifies the objects to be protected. | 

## Methods

### NewGenericNasObjectProtectionRequestParams

`func NewGenericNasObjectProtectionRequestParams(objects []ProtectionObjectInput, ) *GenericNasObjectProtectionRequestParams`

NewGenericNasObjectProtectionRequestParams instantiates a new GenericNasObjectProtectionRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericNasObjectProtectionRequestParamsWithDefaults

`func NewGenericNasObjectProtectionRequestParamsWithDefaults() *GenericNasObjectProtectionRequestParams`

NewGenericNasObjectProtectionRequestParamsWithDefaults instantiates a new GenericNasObjectProtectionRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *GenericNasObjectProtectionRequestParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *GenericNasObjectProtectionRequestParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *GenericNasObjectProtectionRequestParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *GenericNasObjectProtectionRequestParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *GenericNasObjectProtectionRequestParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *GenericNasObjectProtectionRequestParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetEncryptionEnabled

`func (o *GenericNasObjectProtectionRequestParams) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *GenericNasObjectProtectionRequestParams) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *GenericNasObjectProtectionRequestParams) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *GenericNasObjectProtectionRequestParams) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### SetEncryptionEnabledNil

`func (o *GenericNasObjectProtectionRequestParams) SetEncryptionEnabledNil(b bool)`

 SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil

### UnsetEncryptionEnabled
`func (o *GenericNasObjectProtectionRequestParams) UnsetEncryptionEnabled()`

UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
### GetFileFilters

`func (o *GenericNasObjectProtectionRequestParams) GetFileFilters() FileFilteringPolicy`

GetFileFilters returns the FileFilters field if non-nil, zero value otherwise.

### GetFileFiltersOk

`func (o *GenericNasObjectProtectionRequestParams) GetFileFiltersOk() (*FileFilteringPolicy, bool)`

GetFileFiltersOk returns a tuple with the FileFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFilters

`func (o *GenericNasObjectProtectionRequestParams) SetFileFilters(v FileFilteringPolicy)`

SetFileFilters sets FileFilters field to given value.

### HasFileFilters

`func (o *GenericNasObjectProtectionRequestParams) HasFileFilters() bool`

HasFileFilters returns a boolean if a field has been set.

### GetFileLockConfig

`func (o *GenericNasObjectProtectionRequestParams) GetFileLockConfig() FileLevelDataLockConfig`

GetFileLockConfig returns the FileLockConfig field if non-nil, zero value otherwise.

### GetFileLockConfigOk

`func (o *GenericNasObjectProtectionRequestParams) GetFileLockConfigOk() (*FileLevelDataLockConfig, bool)`

GetFileLockConfigOk returns a tuple with the FileLockConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLockConfig

`func (o *GenericNasObjectProtectionRequestParams) SetFileLockConfig(v FileLevelDataLockConfig)`

SetFileLockConfig sets FileLockConfig field to given value.

### HasFileLockConfig

`func (o *GenericNasObjectProtectionRequestParams) HasFileLockConfig() bool`

HasFileLockConfig returns a boolean if a field has been set.

### GetIndexingPolicy

`func (o *GenericNasObjectProtectionRequestParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *GenericNasObjectProtectionRequestParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *GenericNasObjectProtectionRequestParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *GenericNasObjectProtectionRequestParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetPrePostScript

`func (o *GenericNasObjectProtectionRequestParams) GetPrePostScript() HostBasedBackupScriptParams`

GetPrePostScript returns the PrePostScript field if non-nil, zero value otherwise.

### GetPrePostScriptOk

`func (o *GenericNasObjectProtectionRequestParams) GetPrePostScriptOk() (*HostBasedBackupScriptParams, bool)`

GetPrePostScriptOk returns a tuple with the PrePostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePostScript

`func (o *GenericNasObjectProtectionRequestParams) SetPrePostScript(v HostBasedBackupScriptParams)`

SetPrePostScript sets PrePostScript field to given value.

### HasPrePostScript

`func (o *GenericNasObjectProtectionRequestParams) HasPrePostScript() bool`

HasPrePostScript returns a boolean if a field has been set.

### GetThrottlingConfig

`func (o *GenericNasObjectProtectionRequestParams) GetThrottlingConfig() NasThrottlingConfig`

GetThrottlingConfig returns the ThrottlingConfig field if non-nil, zero value otherwise.

### GetThrottlingConfigOk

`func (o *GenericNasObjectProtectionRequestParams) GetThrottlingConfigOk() (*NasThrottlingConfig, bool)`

GetThrottlingConfigOk returns a tuple with the ThrottlingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingConfig

`func (o *GenericNasObjectProtectionRequestParams) SetThrottlingConfig(v NasThrottlingConfig)`

SetThrottlingConfig sets ThrottlingConfig field to given value.

### HasThrottlingConfig

`func (o *GenericNasObjectProtectionRequestParams) HasThrottlingConfig() bool`

HasThrottlingConfig returns a boolean if a field has been set.

### GetProtocol

`func (o *GenericNasObjectProtectionRequestParams) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GenericNasObjectProtectionRequestParams) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GenericNasObjectProtectionRequestParams) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GenericNasObjectProtectionRequestParams) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *GenericNasObjectProtectionRequestParams) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *GenericNasObjectProtectionRequestParams) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetObjects

`func (o *GenericNasObjectProtectionRequestParams) GetObjects() []ProtectionObjectInput`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *GenericNasObjectProtectionRequestParams) GetObjectsOk() (*[]ProtectionObjectInput, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *GenericNasObjectProtectionRequestParams) SetObjects(v []ProtectionObjectInput)`

SetObjects sets Objects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


