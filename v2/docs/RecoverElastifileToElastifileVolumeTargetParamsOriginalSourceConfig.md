# RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other volumes if one of the volumes fails to recover. Default value is false. | [optional] 
**EncryptionEnabled** | Pointer to **NullableBool** | Specifies whether encryption should be enabled during recovery. | [optional] 
**FilterIpConfig** | Pointer to [**FilterIpConfig**](FilterIpConfig.md) |  | [optional] 
**OverwriteExistingFile** | Pointer to **NullableBool** | Specifies whether to overwrite existing file/folder during recovery. | [optional] 
**PreserveFileAttributes** | Pointer to **NullableBool** | Specifies whether to preserve file/folder attributes during recovery. | [optional] 
**VlanConfig** | Pointer to [**RecoveryVlanConfig**](RecoveryVlanConfig.md) |  | [optional] 

## Methods

### NewRecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig

`func NewRecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig() *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig`

NewRecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig instantiates a new RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfigWithDefaults

`func NewRecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfigWithDefaults() *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig`

NewRecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfigWithDefaults instantiates a new RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetEncryptionEnabled

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### SetEncryptionEnabledNil

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) SetEncryptionEnabledNil(b bool)`

 SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil

### UnsetEncryptionEnabled
`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) UnsetEncryptionEnabled()`

UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
### GetFilterIpConfig

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) GetFilterIpConfig() FilterIpConfig`

GetFilterIpConfig returns the FilterIpConfig field if non-nil, zero value otherwise.

### GetFilterIpConfigOk

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) GetFilterIpConfigOk() (*FilterIpConfig, bool)`

GetFilterIpConfigOk returns a tuple with the FilterIpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterIpConfig

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) SetFilterIpConfig(v FilterIpConfig)`

SetFilterIpConfig sets FilterIpConfig field to given value.

### HasFilterIpConfig

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) HasFilterIpConfig() bool`

HasFilterIpConfig returns a boolean if a field has been set.

### GetOverwriteExistingFile

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) GetOverwriteExistingFile() bool`

GetOverwriteExistingFile returns the OverwriteExistingFile field if non-nil, zero value otherwise.

### GetOverwriteExistingFileOk

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) GetOverwriteExistingFileOk() (*bool, bool)`

GetOverwriteExistingFileOk returns a tuple with the OverwriteExistingFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteExistingFile

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) SetOverwriteExistingFile(v bool)`

SetOverwriteExistingFile sets OverwriteExistingFile field to given value.

### HasOverwriteExistingFile

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) HasOverwriteExistingFile() bool`

HasOverwriteExistingFile returns a boolean if a field has been set.

### SetOverwriteExistingFileNil

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) SetOverwriteExistingFileNil(b bool)`

 SetOverwriteExistingFileNil sets the value for OverwriteExistingFile to be an explicit nil

### UnsetOverwriteExistingFile
`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) UnsetOverwriteExistingFile()`

UnsetOverwriteExistingFile ensures that no value is present for OverwriteExistingFile, not even an explicit nil
### GetPreserveFileAttributes

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) GetPreserveFileAttributes() bool`

GetPreserveFileAttributes returns the PreserveFileAttributes field if non-nil, zero value otherwise.

### GetPreserveFileAttributesOk

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) GetPreserveFileAttributesOk() (*bool, bool)`

GetPreserveFileAttributesOk returns a tuple with the PreserveFileAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveFileAttributes

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) SetPreserveFileAttributes(v bool)`

SetPreserveFileAttributes sets PreserveFileAttributes field to given value.

### HasPreserveFileAttributes

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) HasPreserveFileAttributes() bool`

HasPreserveFileAttributes returns a boolean if a field has been set.

### SetPreserveFileAttributesNil

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) SetPreserveFileAttributesNil(b bool)`

 SetPreserveFileAttributesNil sets the value for PreserveFileAttributes to be an explicit nil

### UnsetPreserveFileAttributes
`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) UnsetPreserveFileAttributes()`

UnsetPreserveFileAttributes ensures that no value is present for PreserveFileAttributes, not even an explicit nil
### GetVlanConfig

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) GetVlanConfig() RecoveryVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) GetVlanConfigOk() (*RecoveryVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) SetVlanConfig(v RecoveryVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *RecoverElastifileToElastifileVolumeTargetParamsOriginalSourceConfig) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


