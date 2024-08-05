# OriginalViewFilesTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternatePath** | Pointer to **NullableString** | Specifies the alternate path location to recover files to. | [optional] 
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other files if one of the files fails to recover. Default value is false. | [optional] 
**OverwriteExistingFile** | Pointer to **NullableBool** | Specifies whether to overwrite existing file/folder during recovery. | [optional] 
**PreserveFileAttributes** | Pointer to **NullableBool** | Specifies whether to preserve file/folder attributes during recovery. | [optional] 
**RecoverToOriginalPath** | **NullableBool** | Specifies whether to recover files and folders to the original path location. If false, alternatePath must be specified. | 
**VlanConfig** | Pointer to [**RecoveryVlanConfig**](RecoveryVlanConfig.md) |  | [optional] 

## Methods

### NewOriginalViewFilesTargetParams

`func NewOriginalViewFilesTargetParams(recoverToOriginalPath NullableBool, ) *OriginalViewFilesTargetParams`

NewOriginalViewFilesTargetParams instantiates a new OriginalViewFilesTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginalViewFilesTargetParamsWithDefaults

`func NewOriginalViewFilesTargetParamsWithDefaults() *OriginalViewFilesTargetParams`

NewOriginalViewFilesTargetParamsWithDefaults instantiates a new OriginalViewFilesTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternatePath

`func (o *OriginalViewFilesTargetParams) GetAlternatePath() string`

GetAlternatePath returns the AlternatePath field if non-nil, zero value otherwise.

### GetAlternatePathOk

`func (o *OriginalViewFilesTargetParams) GetAlternatePathOk() (*string, bool)`

GetAlternatePathOk returns a tuple with the AlternatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternatePath

`func (o *OriginalViewFilesTargetParams) SetAlternatePath(v string)`

SetAlternatePath sets AlternatePath field to given value.

### HasAlternatePath

`func (o *OriginalViewFilesTargetParams) HasAlternatePath() bool`

HasAlternatePath returns a boolean if a field has been set.

### SetAlternatePathNil

`func (o *OriginalViewFilesTargetParams) SetAlternatePathNil(b bool)`

 SetAlternatePathNil sets the value for AlternatePath to be an explicit nil

### UnsetAlternatePath
`func (o *OriginalViewFilesTargetParams) UnsetAlternatePath()`

UnsetAlternatePath ensures that no value is present for AlternatePath, not even an explicit nil
### GetContinueOnError

`func (o *OriginalViewFilesTargetParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *OriginalViewFilesTargetParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *OriginalViewFilesTargetParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *OriginalViewFilesTargetParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *OriginalViewFilesTargetParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *OriginalViewFilesTargetParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetOverwriteExistingFile

`func (o *OriginalViewFilesTargetParams) GetOverwriteExistingFile() bool`

GetOverwriteExistingFile returns the OverwriteExistingFile field if non-nil, zero value otherwise.

### GetOverwriteExistingFileOk

`func (o *OriginalViewFilesTargetParams) GetOverwriteExistingFileOk() (*bool, bool)`

GetOverwriteExistingFileOk returns a tuple with the OverwriteExistingFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteExistingFile

`func (o *OriginalViewFilesTargetParams) SetOverwriteExistingFile(v bool)`

SetOverwriteExistingFile sets OverwriteExistingFile field to given value.

### HasOverwriteExistingFile

`func (o *OriginalViewFilesTargetParams) HasOverwriteExistingFile() bool`

HasOverwriteExistingFile returns a boolean if a field has been set.

### SetOverwriteExistingFileNil

`func (o *OriginalViewFilesTargetParams) SetOverwriteExistingFileNil(b bool)`

 SetOverwriteExistingFileNil sets the value for OverwriteExistingFile to be an explicit nil

### UnsetOverwriteExistingFile
`func (o *OriginalViewFilesTargetParams) UnsetOverwriteExistingFile()`

UnsetOverwriteExistingFile ensures that no value is present for OverwriteExistingFile, not even an explicit nil
### GetPreserveFileAttributes

`func (o *OriginalViewFilesTargetParams) GetPreserveFileAttributes() bool`

GetPreserveFileAttributes returns the PreserveFileAttributes field if non-nil, zero value otherwise.

### GetPreserveFileAttributesOk

`func (o *OriginalViewFilesTargetParams) GetPreserveFileAttributesOk() (*bool, bool)`

GetPreserveFileAttributesOk returns a tuple with the PreserveFileAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveFileAttributes

`func (o *OriginalViewFilesTargetParams) SetPreserveFileAttributes(v bool)`

SetPreserveFileAttributes sets PreserveFileAttributes field to given value.

### HasPreserveFileAttributes

`func (o *OriginalViewFilesTargetParams) HasPreserveFileAttributes() bool`

HasPreserveFileAttributes returns a boolean if a field has been set.

### SetPreserveFileAttributesNil

`func (o *OriginalViewFilesTargetParams) SetPreserveFileAttributesNil(b bool)`

 SetPreserveFileAttributesNil sets the value for PreserveFileAttributes to be an explicit nil

### UnsetPreserveFileAttributes
`func (o *OriginalViewFilesTargetParams) UnsetPreserveFileAttributes()`

UnsetPreserveFileAttributes ensures that no value is present for PreserveFileAttributes, not even an explicit nil
### GetRecoverToOriginalPath

`func (o *OriginalViewFilesTargetParams) GetRecoverToOriginalPath() bool`

GetRecoverToOriginalPath returns the RecoverToOriginalPath field if non-nil, zero value otherwise.

### GetRecoverToOriginalPathOk

`func (o *OriginalViewFilesTargetParams) GetRecoverToOriginalPathOk() (*bool, bool)`

GetRecoverToOriginalPathOk returns a tuple with the RecoverToOriginalPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToOriginalPath

`func (o *OriginalViewFilesTargetParams) SetRecoverToOriginalPath(v bool)`

SetRecoverToOriginalPath sets RecoverToOriginalPath field to given value.


### SetRecoverToOriginalPathNil

`func (o *OriginalViewFilesTargetParams) SetRecoverToOriginalPathNil(b bool)`

 SetRecoverToOriginalPathNil sets the value for RecoverToOriginalPath to be an explicit nil

### UnsetRecoverToOriginalPath
`func (o *OriginalViewFilesTargetParams) UnsetRecoverToOriginalPath()`

UnsetRecoverToOriginalPath ensures that no value is present for RecoverToOriginalPath, not even an explicit nil
### GetVlanConfig

`func (o *OriginalViewFilesTargetParams) GetVlanConfig() RecoveryVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *OriginalViewFilesTargetParams) GetVlanConfigOk() (*RecoveryVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *OriginalViewFilesTargetParams) SetVlanConfig(v RecoveryVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *OriginalViewFilesTargetParams) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


