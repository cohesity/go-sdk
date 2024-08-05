# RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternatePath** | **NullableString** | Specifies the path location to recover files to. | 
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other files if one of the files fails to recover. Default value is false. | [optional] 
**EncryptionEnabled** | Pointer to **NullableBool** | Specifies whether encryption should be enabled during recovery. | [optional] 
**FilterIpConfig** | Pointer to [**FilterIpConfig**](FilterIpConfig.md) |  | [optional] 
**OverwriteExistingFile** | Pointer to **NullableBool** | Specifies whether to overwrite existing file/folder during recovery. | [optional] 
**ParentSource** | Pointer to [**RecoverOtherNasToElastifileFilesTargetParamsParentSource**](RecoverOtherNasToElastifileFilesTargetParamsParentSource.md) |  | [optional] 
**PreserveFileAttributes** | Pointer to **NullableBool** | Specifies whether to preserve file/folder attributes during recovery. | [optional] 
**VlanConfig** | Pointer to [**RecoveryVlanConfig**](RecoveryVlanConfig.md) |  | [optional] 
**Volume** | [**RecoverOtherNasToElastifileFilesTargetParamsVolume**](RecoverOtherNasToElastifileFilesTargetParamsVolume.md) |  | 

## Methods

### NewRecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig

`func NewRecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig(alternatePath NullableString, volume RecoverOtherNasToElastifileFilesTargetParamsVolume, ) *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig`

NewRecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig instantiates a new RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfigWithDefaults

`func NewRecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfigWithDefaults() *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig`

NewRecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfigWithDefaults instantiates a new RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternatePath

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) GetAlternatePath() string`

GetAlternatePath returns the AlternatePath field if non-nil, zero value otherwise.

### GetAlternatePathOk

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) GetAlternatePathOk() (*string, bool)`

GetAlternatePathOk returns a tuple with the AlternatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternatePath

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) SetAlternatePath(v string)`

SetAlternatePath sets AlternatePath field to given value.


### SetAlternatePathNil

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) SetAlternatePathNil(b bool)`

 SetAlternatePathNil sets the value for AlternatePath to be an explicit nil

### UnsetAlternatePath
`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) UnsetAlternatePath()`

UnsetAlternatePath ensures that no value is present for AlternatePath, not even an explicit nil
### GetContinueOnError

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetEncryptionEnabled

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### SetEncryptionEnabledNil

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) SetEncryptionEnabledNil(b bool)`

 SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil

### UnsetEncryptionEnabled
`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) UnsetEncryptionEnabled()`

UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
### GetFilterIpConfig

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) GetFilterIpConfig() FilterIpConfig`

GetFilterIpConfig returns the FilterIpConfig field if non-nil, zero value otherwise.

### GetFilterIpConfigOk

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) GetFilterIpConfigOk() (*FilterIpConfig, bool)`

GetFilterIpConfigOk returns a tuple with the FilterIpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterIpConfig

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) SetFilterIpConfig(v FilterIpConfig)`

SetFilterIpConfig sets FilterIpConfig field to given value.

### HasFilterIpConfig

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) HasFilterIpConfig() bool`

HasFilterIpConfig returns a boolean if a field has been set.

### GetOverwriteExistingFile

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) GetOverwriteExistingFile() bool`

GetOverwriteExistingFile returns the OverwriteExistingFile field if non-nil, zero value otherwise.

### GetOverwriteExistingFileOk

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) GetOverwriteExistingFileOk() (*bool, bool)`

GetOverwriteExistingFileOk returns a tuple with the OverwriteExistingFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteExistingFile

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) SetOverwriteExistingFile(v bool)`

SetOverwriteExistingFile sets OverwriteExistingFile field to given value.

### HasOverwriteExistingFile

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) HasOverwriteExistingFile() bool`

HasOverwriteExistingFile returns a boolean if a field has been set.

### SetOverwriteExistingFileNil

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) SetOverwriteExistingFileNil(b bool)`

 SetOverwriteExistingFileNil sets the value for OverwriteExistingFile to be an explicit nil

### UnsetOverwriteExistingFile
`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) UnsetOverwriteExistingFile()`

UnsetOverwriteExistingFile ensures that no value is present for OverwriteExistingFile, not even an explicit nil
### GetParentSource

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) GetParentSource() RecoverOtherNasToElastifileFilesTargetParamsParentSource`

GetParentSource returns the ParentSource field if non-nil, zero value otherwise.

### GetParentSourceOk

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) GetParentSourceOk() (*RecoverOtherNasToElastifileFilesTargetParamsParentSource, bool)`

GetParentSourceOk returns a tuple with the ParentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSource

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) SetParentSource(v RecoverOtherNasToElastifileFilesTargetParamsParentSource)`

SetParentSource sets ParentSource field to given value.

### HasParentSource

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) HasParentSource() bool`

HasParentSource returns a boolean if a field has been set.

### GetPreserveFileAttributes

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) GetPreserveFileAttributes() bool`

GetPreserveFileAttributes returns the PreserveFileAttributes field if non-nil, zero value otherwise.

### GetPreserveFileAttributesOk

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) GetPreserveFileAttributesOk() (*bool, bool)`

GetPreserveFileAttributesOk returns a tuple with the PreserveFileAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveFileAttributes

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) SetPreserveFileAttributes(v bool)`

SetPreserveFileAttributes sets PreserveFileAttributes field to given value.

### HasPreserveFileAttributes

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) HasPreserveFileAttributes() bool`

HasPreserveFileAttributes returns a boolean if a field has been set.

### SetPreserveFileAttributesNil

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) SetPreserveFileAttributesNil(b bool)`

 SetPreserveFileAttributesNil sets the value for PreserveFileAttributes to be an explicit nil

### UnsetPreserveFileAttributes
`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) UnsetPreserveFileAttributes()`

UnsetPreserveFileAttributes ensures that no value is present for PreserveFileAttributes, not even an explicit nil
### GetVlanConfig

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) GetVlanConfig() RecoveryVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) GetVlanConfigOk() (*RecoveryVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) SetVlanConfig(v RecoveryVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### GetVolume

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) GetVolume() RecoverOtherNasToElastifileFilesTargetParamsVolume`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) GetVolumeOk() (*RecoverOtherNasToElastifileFilesTargetParamsVolume, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *RecoverFlashbladeToFlashbladeFilesTargetParamsNewSourceConfig) SetVolume(v RecoverOtherNasToElastifileFilesTargetParamsVolume)`

SetVolume sets Volume field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


