# RecoverNetappNasVolumeParamsGenericNasTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other volumes if one of the volumes fails to recover. Default value is false. | [optional] 
**EncryptionEnabled** | Pointer to **NullableBool** | Specifies whether encryption should be enabled during recovery. | [optional] 
**OverwriteExistingFile** | Pointer to **NullableBool** | Specifies whether to overwrite existing file/folder during recovery. | [optional] 
**ParentSource** | Pointer to [**RecoverOtherNasToElastifileFilesTargetParamsParentSource**](RecoverOtherNasToElastifileFilesTargetParamsParentSource.md) |  | [optional] 
**PreserveFileAttributes** | Pointer to **NullableBool** | Specifies whether to preserve file/folder attributes during recovery. | [optional] 
**VlanConfig** | Pointer to [**RecoveryVlanConfig**](RecoveryVlanConfig.md) |  | [optional] 
**Volume** | [**RecoverOtherNasToElastifileVolumeTargetParamsVolume**](RecoverOtherNasToElastifileVolumeTargetParamsVolume.md) |  | 

## Methods

### NewRecoverNetappNasVolumeParamsGenericNasTargetParams

`func NewRecoverNetappNasVolumeParamsGenericNasTargetParams(volume RecoverOtherNasToElastifileVolumeTargetParamsVolume, ) *RecoverNetappNasVolumeParamsGenericNasTargetParams`

NewRecoverNetappNasVolumeParamsGenericNasTargetParams instantiates a new RecoverNetappNasVolumeParamsGenericNasTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverNetappNasVolumeParamsGenericNasTargetParamsWithDefaults

`func NewRecoverNetappNasVolumeParamsGenericNasTargetParamsWithDefaults() *RecoverNetappNasVolumeParamsGenericNasTargetParams`

NewRecoverNetappNasVolumeParamsGenericNasTargetParamsWithDefaults instantiates a new RecoverNetappNasVolumeParamsGenericNasTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetEncryptionEnabled

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### SetEncryptionEnabledNil

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) SetEncryptionEnabledNil(b bool)`

 SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil

### UnsetEncryptionEnabled
`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) UnsetEncryptionEnabled()`

UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
### GetOverwriteExistingFile

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) GetOverwriteExistingFile() bool`

GetOverwriteExistingFile returns the OverwriteExistingFile field if non-nil, zero value otherwise.

### GetOverwriteExistingFileOk

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) GetOverwriteExistingFileOk() (*bool, bool)`

GetOverwriteExistingFileOk returns a tuple with the OverwriteExistingFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteExistingFile

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) SetOverwriteExistingFile(v bool)`

SetOverwriteExistingFile sets OverwriteExistingFile field to given value.

### HasOverwriteExistingFile

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) HasOverwriteExistingFile() bool`

HasOverwriteExistingFile returns a boolean if a field has been set.

### SetOverwriteExistingFileNil

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) SetOverwriteExistingFileNil(b bool)`

 SetOverwriteExistingFileNil sets the value for OverwriteExistingFile to be an explicit nil

### UnsetOverwriteExistingFile
`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) UnsetOverwriteExistingFile()`

UnsetOverwriteExistingFile ensures that no value is present for OverwriteExistingFile, not even an explicit nil
### GetParentSource

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) GetParentSource() RecoverOtherNasToElastifileFilesTargetParamsParentSource`

GetParentSource returns the ParentSource field if non-nil, zero value otherwise.

### GetParentSourceOk

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) GetParentSourceOk() (*RecoverOtherNasToElastifileFilesTargetParamsParentSource, bool)`

GetParentSourceOk returns a tuple with the ParentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSource

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) SetParentSource(v RecoverOtherNasToElastifileFilesTargetParamsParentSource)`

SetParentSource sets ParentSource field to given value.

### HasParentSource

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) HasParentSource() bool`

HasParentSource returns a boolean if a field has been set.

### GetPreserveFileAttributes

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) GetPreserveFileAttributes() bool`

GetPreserveFileAttributes returns the PreserveFileAttributes field if non-nil, zero value otherwise.

### GetPreserveFileAttributesOk

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) GetPreserveFileAttributesOk() (*bool, bool)`

GetPreserveFileAttributesOk returns a tuple with the PreserveFileAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveFileAttributes

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) SetPreserveFileAttributes(v bool)`

SetPreserveFileAttributes sets PreserveFileAttributes field to given value.

### HasPreserveFileAttributes

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) HasPreserveFileAttributes() bool`

HasPreserveFileAttributes returns a boolean if a field has been set.

### SetPreserveFileAttributesNil

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) SetPreserveFileAttributesNil(b bool)`

 SetPreserveFileAttributesNil sets the value for PreserveFileAttributes to be an explicit nil

### UnsetPreserveFileAttributes
`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) UnsetPreserveFileAttributes()`

UnsetPreserveFileAttributes ensures that no value is present for PreserveFileAttributes, not even an explicit nil
### GetVlanConfig

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) GetVlanConfig() RecoveryVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) GetVlanConfigOk() (*RecoveryVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) SetVlanConfig(v RecoveryVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### GetVolume

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) GetVolume() RecoverOtherNasToElastifileVolumeTargetParamsVolume`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) GetVolumeOk() (*RecoverOtherNasToElastifileVolumeTargetParamsVolume, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *RecoverNetappNasVolumeParamsGenericNasTargetParams) SetVolume(v RecoverOtherNasToElastifileVolumeTargetParamsVolume)`

SetVolume sets Volume field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


