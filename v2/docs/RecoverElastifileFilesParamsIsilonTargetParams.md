# RecoverElastifileFilesParamsIsilonTargetParams

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

### NewRecoverElastifileFilesParamsIsilonTargetParams

`func NewRecoverElastifileFilesParamsIsilonTargetParams(alternatePath NullableString, volume RecoverOtherNasToElastifileFilesTargetParamsVolume, ) *RecoverElastifileFilesParamsIsilonTargetParams`

NewRecoverElastifileFilesParamsIsilonTargetParams instantiates a new RecoverElastifileFilesParamsIsilonTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverElastifileFilesParamsIsilonTargetParamsWithDefaults

`func NewRecoverElastifileFilesParamsIsilonTargetParamsWithDefaults() *RecoverElastifileFilesParamsIsilonTargetParams`

NewRecoverElastifileFilesParamsIsilonTargetParamsWithDefaults instantiates a new RecoverElastifileFilesParamsIsilonTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternatePath

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) GetAlternatePath() string`

GetAlternatePath returns the AlternatePath field if non-nil, zero value otherwise.

### GetAlternatePathOk

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) GetAlternatePathOk() (*string, bool)`

GetAlternatePathOk returns a tuple with the AlternatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternatePath

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) SetAlternatePath(v string)`

SetAlternatePath sets AlternatePath field to given value.


### SetAlternatePathNil

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) SetAlternatePathNil(b bool)`

 SetAlternatePathNil sets the value for AlternatePath to be an explicit nil

### UnsetAlternatePath
`func (o *RecoverElastifileFilesParamsIsilonTargetParams) UnsetAlternatePath()`

UnsetAlternatePath ensures that no value is present for AlternatePath, not even an explicit nil
### GetContinueOnError

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverElastifileFilesParamsIsilonTargetParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetEncryptionEnabled

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### SetEncryptionEnabledNil

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) SetEncryptionEnabledNil(b bool)`

 SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil

### UnsetEncryptionEnabled
`func (o *RecoverElastifileFilesParamsIsilonTargetParams) UnsetEncryptionEnabled()`

UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
### GetFilterIpConfig

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) GetFilterIpConfig() FilterIpConfig`

GetFilterIpConfig returns the FilterIpConfig field if non-nil, zero value otherwise.

### GetFilterIpConfigOk

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) GetFilterIpConfigOk() (*FilterIpConfig, bool)`

GetFilterIpConfigOk returns a tuple with the FilterIpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterIpConfig

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) SetFilterIpConfig(v FilterIpConfig)`

SetFilterIpConfig sets FilterIpConfig field to given value.

### HasFilterIpConfig

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) HasFilterIpConfig() bool`

HasFilterIpConfig returns a boolean if a field has been set.

### GetOverwriteExistingFile

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) GetOverwriteExistingFile() bool`

GetOverwriteExistingFile returns the OverwriteExistingFile field if non-nil, zero value otherwise.

### GetOverwriteExistingFileOk

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) GetOverwriteExistingFileOk() (*bool, bool)`

GetOverwriteExistingFileOk returns a tuple with the OverwriteExistingFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteExistingFile

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) SetOverwriteExistingFile(v bool)`

SetOverwriteExistingFile sets OverwriteExistingFile field to given value.

### HasOverwriteExistingFile

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) HasOverwriteExistingFile() bool`

HasOverwriteExistingFile returns a boolean if a field has been set.

### SetOverwriteExistingFileNil

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) SetOverwriteExistingFileNil(b bool)`

 SetOverwriteExistingFileNil sets the value for OverwriteExistingFile to be an explicit nil

### UnsetOverwriteExistingFile
`func (o *RecoverElastifileFilesParamsIsilonTargetParams) UnsetOverwriteExistingFile()`

UnsetOverwriteExistingFile ensures that no value is present for OverwriteExistingFile, not even an explicit nil
### GetParentSource

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) GetParentSource() RecoverOtherNasToElastifileFilesTargetParamsParentSource`

GetParentSource returns the ParentSource field if non-nil, zero value otherwise.

### GetParentSourceOk

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) GetParentSourceOk() (*RecoverOtherNasToElastifileFilesTargetParamsParentSource, bool)`

GetParentSourceOk returns a tuple with the ParentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSource

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) SetParentSource(v RecoverOtherNasToElastifileFilesTargetParamsParentSource)`

SetParentSource sets ParentSource field to given value.

### HasParentSource

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) HasParentSource() bool`

HasParentSource returns a boolean if a field has been set.

### GetPreserveFileAttributes

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) GetPreserveFileAttributes() bool`

GetPreserveFileAttributes returns the PreserveFileAttributes field if non-nil, zero value otherwise.

### GetPreserveFileAttributesOk

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) GetPreserveFileAttributesOk() (*bool, bool)`

GetPreserveFileAttributesOk returns a tuple with the PreserveFileAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveFileAttributes

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) SetPreserveFileAttributes(v bool)`

SetPreserveFileAttributes sets PreserveFileAttributes field to given value.

### HasPreserveFileAttributes

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) HasPreserveFileAttributes() bool`

HasPreserveFileAttributes returns a boolean if a field has been set.

### SetPreserveFileAttributesNil

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) SetPreserveFileAttributesNil(b bool)`

 SetPreserveFileAttributesNil sets the value for PreserveFileAttributes to be an explicit nil

### UnsetPreserveFileAttributes
`func (o *RecoverElastifileFilesParamsIsilonTargetParams) UnsetPreserveFileAttributes()`

UnsetPreserveFileAttributes ensures that no value is present for PreserveFileAttributes, not even an explicit nil
### GetVlanConfig

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) GetVlanConfig() RecoveryVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) GetVlanConfigOk() (*RecoveryVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) SetVlanConfig(v RecoveryVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### GetVolume

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) GetVolume() RecoverOtherNasToElastifileFilesTargetParamsVolume`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) GetVolumeOk() (*RecoverOtherNasToElastifileFilesTargetParamsVolume, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *RecoverElastifileFilesParamsIsilonTargetParams) SetVolume(v RecoverOtherNasToElastifileFilesTargetParamsVolume)`

SetVolume sets Volume field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


