# RecoverElastifileNasVolumeParamsGpfsTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other volumes if one of the volumes fails to recover. Default value is false. | [optional] 
**EncryptionEnabled** | Pointer to **NullableBool** | Specifies whether encryption should be enabled during recovery. | [optional] 
**FilterIpConfig** | Pointer to [**FilterIpConfig**](FilterIpConfig.md) |  | [optional] 
**OverwriteExistingFile** | Pointer to **NullableBool** | Specifies whether to overwrite existing file/folder during recovery. | [optional] 
**ParentSource** | Pointer to [**RecoverOtherNasToElastifileFilesTargetParamsParentSource**](RecoverOtherNasToElastifileFilesTargetParamsParentSource.md) |  | [optional] 
**PreserveFileAttributes** | Pointer to **NullableBool** | Specifies whether to preserve file/folder attributes during recovery. | [optional] 
**VlanConfig** | Pointer to [**RecoveryVlanConfig**](RecoveryVlanConfig.md) |  | [optional] 
**Volume** | [**RecoverOtherNasToElastifileVolumeTargetParamsVolume**](RecoverOtherNasToElastifileVolumeTargetParamsVolume.md) |  | 

## Methods

### NewRecoverElastifileNasVolumeParamsGpfsTargetParams

`func NewRecoverElastifileNasVolumeParamsGpfsTargetParams(volume RecoverOtherNasToElastifileVolumeTargetParamsVolume, ) *RecoverElastifileNasVolumeParamsGpfsTargetParams`

NewRecoverElastifileNasVolumeParamsGpfsTargetParams instantiates a new RecoverElastifileNasVolumeParamsGpfsTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverElastifileNasVolumeParamsGpfsTargetParamsWithDefaults

`func NewRecoverElastifileNasVolumeParamsGpfsTargetParamsWithDefaults() *RecoverElastifileNasVolumeParamsGpfsTargetParams`

NewRecoverElastifileNasVolumeParamsGpfsTargetParamsWithDefaults instantiates a new RecoverElastifileNasVolumeParamsGpfsTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetEncryptionEnabled

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### SetEncryptionEnabledNil

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) SetEncryptionEnabledNil(b bool)`

 SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil

### UnsetEncryptionEnabled
`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) UnsetEncryptionEnabled()`

UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
### GetFilterIpConfig

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) GetFilterIpConfig() FilterIpConfig`

GetFilterIpConfig returns the FilterIpConfig field if non-nil, zero value otherwise.

### GetFilterIpConfigOk

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) GetFilterIpConfigOk() (*FilterIpConfig, bool)`

GetFilterIpConfigOk returns a tuple with the FilterIpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterIpConfig

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) SetFilterIpConfig(v FilterIpConfig)`

SetFilterIpConfig sets FilterIpConfig field to given value.

### HasFilterIpConfig

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) HasFilterIpConfig() bool`

HasFilterIpConfig returns a boolean if a field has been set.

### GetOverwriteExistingFile

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) GetOverwriteExistingFile() bool`

GetOverwriteExistingFile returns the OverwriteExistingFile field if non-nil, zero value otherwise.

### GetOverwriteExistingFileOk

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) GetOverwriteExistingFileOk() (*bool, bool)`

GetOverwriteExistingFileOk returns a tuple with the OverwriteExistingFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteExistingFile

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) SetOverwriteExistingFile(v bool)`

SetOverwriteExistingFile sets OverwriteExistingFile field to given value.

### HasOverwriteExistingFile

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) HasOverwriteExistingFile() bool`

HasOverwriteExistingFile returns a boolean if a field has been set.

### SetOverwriteExistingFileNil

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) SetOverwriteExistingFileNil(b bool)`

 SetOverwriteExistingFileNil sets the value for OverwriteExistingFile to be an explicit nil

### UnsetOverwriteExistingFile
`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) UnsetOverwriteExistingFile()`

UnsetOverwriteExistingFile ensures that no value is present for OverwriteExistingFile, not even an explicit nil
### GetParentSource

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) GetParentSource() RecoverOtherNasToElastifileFilesTargetParamsParentSource`

GetParentSource returns the ParentSource field if non-nil, zero value otherwise.

### GetParentSourceOk

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) GetParentSourceOk() (*RecoverOtherNasToElastifileFilesTargetParamsParentSource, bool)`

GetParentSourceOk returns a tuple with the ParentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSource

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) SetParentSource(v RecoverOtherNasToElastifileFilesTargetParamsParentSource)`

SetParentSource sets ParentSource field to given value.

### HasParentSource

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) HasParentSource() bool`

HasParentSource returns a boolean if a field has been set.

### GetPreserveFileAttributes

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) GetPreserveFileAttributes() bool`

GetPreserveFileAttributes returns the PreserveFileAttributes field if non-nil, zero value otherwise.

### GetPreserveFileAttributesOk

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) GetPreserveFileAttributesOk() (*bool, bool)`

GetPreserveFileAttributesOk returns a tuple with the PreserveFileAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveFileAttributes

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) SetPreserveFileAttributes(v bool)`

SetPreserveFileAttributes sets PreserveFileAttributes field to given value.

### HasPreserveFileAttributes

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) HasPreserveFileAttributes() bool`

HasPreserveFileAttributes returns a boolean if a field has been set.

### SetPreserveFileAttributesNil

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) SetPreserveFileAttributesNil(b bool)`

 SetPreserveFileAttributesNil sets the value for PreserveFileAttributes to be an explicit nil

### UnsetPreserveFileAttributes
`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) UnsetPreserveFileAttributes()`

UnsetPreserveFileAttributes ensures that no value is present for PreserveFileAttributes, not even an explicit nil
### GetVlanConfig

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) GetVlanConfig() RecoveryVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) GetVlanConfigOk() (*RecoveryVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) SetVlanConfig(v RecoveryVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### GetVolume

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) GetVolume() RecoverOtherNasToElastifileVolumeTargetParamsVolume`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) GetVolumeOk() (*RecoverOtherNasToElastifileVolumeTargetParamsVolume, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *RecoverElastifileNasVolumeParamsGpfsTargetParams) SetVolume(v RecoverOtherNasToElastifileVolumeTargetParamsVolume)`

SetVolume sets Volume field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


