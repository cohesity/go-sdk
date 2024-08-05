# RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other volumes if one of the volumes fails to recover. Default value is false. | [optional] 
**EncryptionEnabled** | Pointer to **NullableBool** | Specifies whether encryption should be enabled during recovery. | [optional] 
**OverwriteExistingFile** | Pointer to **NullableBool** | Specifies whether to overwrite existing file/folder during recovery. | [optional] 
**PreserveFileAttributes** | Pointer to **NullableBool** | Specifies whether to preserve file/folder attributes during recovery. | [optional] 
**VlanConfig** | Pointer to [**RecoveryVlanConfig**](RecoveryVlanConfig.md) |  | [optional] 

## Methods

### NewRecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig

`func NewRecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig() *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig`

NewRecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig instantiates a new RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfigWithDefaults

`func NewRecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfigWithDefaults() *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig`

NewRecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfigWithDefaults instantiates a new RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetEncryptionEnabled

`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### SetEncryptionEnabledNil

`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) SetEncryptionEnabledNil(b bool)`

 SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil

### UnsetEncryptionEnabled
`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) UnsetEncryptionEnabled()`

UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
### GetOverwriteExistingFile

`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) GetOverwriteExistingFile() bool`

GetOverwriteExistingFile returns the OverwriteExistingFile field if non-nil, zero value otherwise.

### GetOverwriteExistingFileOk

`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) GetOverwriteExistingFileOk() (*bool, bool)`

GetOverwriteExistingFileOk returns a tuple with the OverwriteExistingFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteExistingFile

`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) SetOverwriteExistingFile(v bool)`

SetOverwriteExistingFile sets OverwriteExistingFile field to given value.

### HasOverwriteExistingFile

`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) HasOverwriteExistingFile() bool`

HasOverwriteExistingFile returns a boolean if a field has been set.

### SetOverwriteExistingFileNil

`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) SetOverwriteExistingFileNil(b bool)`

 SetOverwriteExistingFileNil sets the value for OverwriteExistingFile to be an explicit nil

### UnsetOverwriteExistingFile
`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) UnsetOverwriteExistingFile()`

UnsetOverwriteExistingFile ensures that no value is present for OverwriteExistingFile, not even an explicit nil
### GetPreserveFileAttributes

`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) GetPreserveFileAttributes() bool`

GetPreserveFileAttributes returns the PreserveFileAttributes field if non-nil, zero value otherwise.

### GetPreserveFileAttributesOk

`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) GetPreserveFileAttributesOk() (*bool, bool)`

GetPreserveFileAttributesOk returns a tuple with the PreserveFileAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveFileAttributes

`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) SetPreserveFileAttributes(v bool)`

SetPreserveFileAttributes sets PreserveFileAttributes field to given value.

### HasPreserveFileAttributes

`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) HasPreserveFileAttributes() bool`

HasPreserveFileAttributes returns a boolean if a field has been set.

### SetPreserveFileAttributesNil

`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) SetPreserveFileAttributesNil(b bool)`

 SetPreserveFileAttributesNil sets the value for PreserveFileAttributes to be an explicit nil

### UnsetPreserveFileAttributes
`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) UnsetPreserveFileAttributes()`

UnsetPreserveFileAttributes ensures that no value is present for PreserveFileAttributes, not even an explicit nil
### GetVlanConfig

`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) GetVlanConfig() RecoveryVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) GetVlanConfigOk() (*RecoveryVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) SetVlanConfig(v RecoveryVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *RecoverGenericNasToGenericNasVolumeTargetParamsOriginalSourceConfig) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


