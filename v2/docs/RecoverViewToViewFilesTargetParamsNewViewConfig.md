# RecoverViewToViewFilesTargetParamsNewViewConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternatePath** | **NullableString** | Specifies the path location to recover files to. | 
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other files if one of the files fails to recover. Default value is false. | [optional] 
**OverwriteExistingFile** | Pointer to **NullableBool** | Specifies whether to overwrite existing file/folder during recovery. | [optional] 
**PreserveFileAttributes** | Pointer to **NullableBool** | Specifies whether to preserve file/folder attributes during recovery. | [optional] 
**VlanConfig** | Pointer to [**RecoveryVlanConfig**](RecoveryVlanConfig.md) |  | [optional] 

## Methods

### NewRecoverViewToViewFilesTargetParamsNewViewConfig

`func NewRecoverViewToViewFilesTargetParamsNewViewConfig(alternatePath NullableString, ) *RecoverViewToViewFilesTargetParamsNewViewConfig`

NewRecoverViewToViewFilesTargetParamsNewViewConfig instantiates a new RecoverViewToViewFilesTargetParamsNewViewConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverViewToViewFilesTargetParamsNewViewConfigWithDefaults

`func NewRecoverViewToViewFilesTargetParamsNewViewConfigWithDefaults() *RecoverViewToViewFilesTargetParamsNewViewConfig`

NewRecoverViewToViewFilesTargetParamsNewViewConfigWithDefaults instantiates a new RecoverViewToViewFilesTargetParamsNewViewConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternatePath

`func (o *RecoverViewToViewFilesTargetParamsNewViewConfig) GetAlternatePath() string`

GetAlternatePath returns the AlternatePath field if non-nil, zero value otherwise.

### GetAlternatePathOk

`func (o *RecoverViewToViewFilesTargetParamsNewViewConfig) GetAlternatePathOk() (*string, bool)`

GetAlternatePathOk returns a tuple with the AlternatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternatePath

`func (o *RecoverViewToViewFilesTargetParamsNewViewConfig) SetAlternatePath(v string)`

SetAlternatePath sets AlternatePath field to given value.


### SetAlternatePathNil

`func (o *RecoverViewToViewFilesTargetParamsNewViewConfig) SetAlternatePathNil(b bool)`

 SetAlternatePathNil sets the value for AlternatePath to be an explicit nil

### UnsetAlternatePath
`func (o *RecoverViewToViewFilesTargetParamsNewViewConfig) UnsetAlternatePath()`

UnsetAlternatePath ensures that no value is present for AlternatePath, not even an explicit nil
### GetContinueOnError

`func (o *RecoverViewToViewFilesTargetParamsNewViewConfig) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverViewToViewFilesTargetParamsNewViewConfig) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverViewToViewFilesTargetParamsNewViewConfig) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverViewToViewFilesTargetParamsNewViewConfig) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverViewToViewFilesTargetParamsNewViewConfig) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverViewToViewFilesTargetParamsNewViewConfig) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetOverwriteExistingFile

`func (o *RecoverViewToViewFilesTargetParamsNewViewConfig) GetOverwriteExistingFile() bool`

GetOverwriteExistingFile returns the OverwriteExistingFile field if non-nil, zero value otherwise.

### GetOverwriteExistingFileOk

`func (o *RecoverViewToViewFilesTargetParamsNewViewConfig) GetOverwriteExistingFileOk() (*bool, bool)`

GetOverwriteExistingFileOk returns a tuple with the OverwriteExistingFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteExistingFile

`func (o *RecoverViewToViewFilesTargetParamsNewViewConfig) SetOverwriteExistingFile(v bool)`

SetOverwriteExistingFile sets OverwriteExistingFile field to given value.

### HasOverwriteExistingFile

`func (o *RecoverViewToViewFilesTargetParamsNewViewConfig) HasOverwriteExistingFile() bool`

HasOverwriteExistingFile returns a boolean if a field has been set.

### SetOverwriteExistingFileNil

`func (o *RecoverViewToViewFilesTargetParamsNewViewConfig) SetOverwriteExistingFileNil(b bool)`

 SetOverwriteExistingFileNil sets the value for OverwriteExistingFile to be an explicit nil

### UnsetOverwriteExistingFile
`func (o *RecoverViewToViewFilesTargetParamsNewViewConfig) UnsetOverwriteExistingFile()`

UnsetOverwriteExistingFile ensures that no value is present for OverwriteExistingFile, not even an explicit nil
### GetPreserveFileAttributes

`func (o *RecoverViewToViewFilesTargetParamsNewViewConfig) GetPreserveFileAttributes() bool`

GetPreserveFileAttributes returns the PreserveFileAttributes field if non-nil, zero value otherwise.

### GetPreserveFileAttributesOk

`func (o *RecoverViewToViewFilesTargetParamsNewViewConfig) GetPreserveFileAttributesOk() (*bool, bool)`

GetPreserveFileAttributesOk returns a tuple with the PreserveFileAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveFileAttributes

`func (o *RecoverViewToViewFilesTargetParamsNewViewConfig) SetPreserveFileAttributes(v bool)`

SetPreserveFileAttributes sets PreserveFileAttributes field to given value.

### HasPreserveFileAttributes

`func (o *RecoverViewToViewFilesTargetParamsNewViewConfig) HasPreserveFileAttributes() bool`

HasPreserveFileAttributes returns a boolean if a field has been set.

### SetPreserveFileAttributesNil

`func (o *RecoverViewToViewFilesTargetParamsNewViewConfig) SetPreserveFileAttributesNil(b bool)`

 SetPreserveFileAttributesNil sets the value for PreserveFileAttributes to be an explicit nil

### UnsetPreserveFileAttributes
`func (o *RecoverViewToViewFilesTargetParamsNewViewConfig) UnsetPreserveFileAttributes()`

UnsetPreserveFileAttributes ensures that no value is present for PreserveFileAttributes, not even an explicit nil
### GetVlanConfig

`func (o *RecoverViewToViewFilesTargetParamsNewViewConfig) GetVlanConfig() RecoveryVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *RecoverViewToViewFilesTargetParamsNewViewConfig) GetVlanConfigOk() (*RecoveryVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *RecoverViewToViewFilesTargetParamsNewViewConfig) SetVlanConfig(v RecoveryVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *RecoverViewToViewFilesTargetParamsNewViewConfig) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


