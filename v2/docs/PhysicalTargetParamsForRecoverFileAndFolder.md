# PhysicalTargetParamsForRecoverFileAndFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternateRestoreDirectory** | Pointer to **NullableString** | Specifies the directory path where restore should happen if restore_to_original_paths is set to false. | [optional] 
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other volumes if one of the volumes fails to recover. Default value is false. | [optional] 
**OverwriteExisting** | Pointer to **NullableBool** | Specifies whether to overwrite existing file/folder during recovery. | [optional] 
**PreserveAcls** | Pointer to **NullableBool** | Whether to preserve the ACLs of the original file. | [optional] 
**PreserveAttributes** | Pointer to **NullableBool** | Specifies whether to preserve file/folder attributes during recovery. | [optional] 
**PreserveTimestamps** | Pointer to **NullableBool** | Whether to preserve the original time stamps. | [optional] 
**RecoverTarget** | [**PhysicalTargetParamsForRecoverFileAndFolderRecoverTarget**](PhysicalTargetParamsForRecoverFileAndFolderRecoverTarget.md) |  | 
**RestoreEntityType** | Pointer to **NullableString** | Specifies the restore type (restore everything or ACLs only) when restoring or downloading files or folders from a Physical file based or block based backup snapshot. | [optional] 
**RestoreExclusionParam** | Pointer to **[]string** | Specifies list of exclusions that needs to be applicable for the current file level restores. Applicable only for file level restores. | [optional] 
**RestoreToOriginalPaths** | Pointer to **NullableBool** | If this is true, then files will be restored to original paths. | [optional] 
**SaveSuccessFiles** | Pointer to **NullableBool** | Specifies whether to save success files or not. Default value is false | [optional] 
**VlanConfig** | Pointer to [**NullableHyperVTargetParamsForRecoverVmVlanConfig**](HyperVTargetParamsForRecoverVmVlanConfig.md) |  | [optional] 

## Methods

### NewPhysicalTargetParamsForRecoverFileAndFolder

`func NewPhysicalTargetParamsForRecoverFileAndFolder(recoverTarget PhysicalTargetParamsForRecoverFileAndFolderRecoverTarget, ) *PhysicalTargetParamsForRecoverFileAndFolder`

NewPhysicalTargetParamsForRecoverFileAndFolder instantiates a new PhysicalTargetParamsForRecoverFileAndFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalTargetParamsForRecoverFileAndFolderWithDefaults

`func NewPhysicalTargetParamsForRecoverFileAndFolderWithDefaults() *PhysicalTargetParamsForRecoverFileAndFolder`

NewPhysicalTargetParamsForRecoverFileAndFolderWithDefaults instantiates a new PhysicalTargetParamsForRecoverFileAndFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternateRestoreDirectory

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetAlternateRestoreDirectory() string`

GetAlternateRestoreDirectory returns the AlternateRestoreDirectory field if non-nil, zero value otherwise.

### GetAlternateRestoreDirectoryOk

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetAlternateRestoreDirectoryOk() (*string, bool)`

GetAlternateRestoreDirectoryOk returns a tuple with the AlternateRestoreDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateRestoreDirectory

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetAlternateRestoreDirectory(v string)`

SetAlternateRestoreDirectory sets AlternateRestoreDirectory field to given value.

### HasAlternateRestoreDirectory

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) HasAlternateRestoreDirectory() bool`

HasAlternateRestoreDirectory returns a boolean if a field has been set.

### SetAlternateRestoreDirectoryNil

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetAlternateRestoreDirectoryNil(b bool)`

 SetAlternateRestoreDirectoryNil sets the value for AlternateRestoreDirectory to be an explicit nil

### UnsetAlternateRestoreDirectory
`func (o *PhysicalTargetParamsForRecoverFileAndFolder) UnsetAlternateRestoreDirectory()`

UnsetAlternateRestoreDirectory ensures that no value is present for AlternateRestoreDirectory, not even an explicit nil
### GetContinueOnError

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *PhysicalTargetParamsForRecoverFileAndFolder) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetOverwriteExisting

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetOverwriteExisting() bool`

GetOverwriteExisting returns the OverwriteExisting field if non-nil, zero value otherwise.

### GetOverwriteExistingOk

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetOverwriteExistingOk() (*bool, bool)`

GetOverwriteExistingOk returns a tuple with the OverwriteExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteExisting

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetOverwriteExisting(v bool)`

SetOverwriteExisting sets OverwriteExisting field to given value.

### HasOverwriteExisting

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) HasOverwriteExisting() bool`

HasOverwriteExisting returns a boolean if a field has been set.

### SetOverwriteExistingNil

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetOverwriteExistingNil(b bool)`

 SetOverwriteExistingNil sets the value for OverwriteExisting to be an explicit nil

### UnsetOverwriteExisting
`func (o *PhysicalTargetParamsForRecoverFileAndFolder) UnsetOverwriteExisting()`

UnsetOverwriteExisting ensures that no value is present for OverwriteExisting, not even an explicit nil
### GetPreserveAcls

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetPreserveAcls() bool`

GetPreserveAcls returns the PreserveAcls field if non-nil, zero value otherwise.

### GetPreserveAclsOk

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetPreserveAclsOk() (*bool, bool)`

GetPreserveAclsOk returns a tuple with the PreserveAcls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveAcls

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetPreserveAcls(v bool)`

SetPreserveAcls sets PreserveAcls field to given value.

### HasPreserveAcls

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) HasPreserveAcls() bool`

HasPreserveAcls returns a boolean if a field has been set.

### SetPreserveAclsNil

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetPreserveAclsNil(b bool)`

 SetPreserveAclsNil sets the value for PreserveAcls to be an explicit nil

### UnsetPreserveAcls
`func (o *PhysicalTargetParamsForRecoverFileAndFolder) UnsetPreserveAcls()`

UnsetPreserveAcls ensures that no value is present for PreserveAcls, not even an explicit nil
### GetPreserveAttributes

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetPreserveAttributes() bool`

GetPreserveAttributes returns the PreserveAttributes field if non-nil, zero value otherwise.

### GetPreserveAttributesOk

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetPreserveAttributesOk() (*bool, bool)`

GetPreserveAttributesOk returns a tuple with the PreserveAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveAttributes

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetPreserveAttributes(v bool)`

SetPreserveAttributes sets PreserveAttributes field to given value.

### HasPreserveAttributes

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) HasPreserveAttributes() bool`

HasPreserveAttributes returns a boolean if a field has been set.

### SetPreserveAttributesNil

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetPreserveAttributesNil(b bool)`

 SetPreserveAttributesNil sets the value for PreserveAttributes to be an explicit nil

### UnsetPreserveAttributes
`func (o *PhysicalTargetParamsForRecoverFileAndFolder) UnsetPreserveAttributes()`

UnsetPreserveAttributes ensures that no value is present for PreserveAttributes, not even an explicit nil
### GetPreserveTimestamps

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetPreserveTimestamps() bool`

GetPreserveTimestamps returns the PreserveTimestamps field if non-nil, zero value otherwise.

### GetPreserveTimestampsOk

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetPreserveTimestampsOk() (*bool, bool)`

GetPreserveTimestampsOk returns a tuple with the PreserveTimestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveTimestamps

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetPreserveTimestamps(v bool)`

SetPreserveTimestamps sets PreserveTimestamps field to given value.

### HasPreserveTimestamps

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) HasPreserveTimestamps() bool`

HasPreserveTimestamps returns a boolean if a field has been set.

### SetPreserveTimestampsNil

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetPreserveTimestampsNil(b bool)`

 SetPreserveTimestampsNil sets the value for PreserveTimestamps to be an explicit nil

### UnsetPreserveTimestamps
`func (o *PhysicalTargetParamsForRecoverFileAndFolder) UnsetPreserveTimestamps()`

UnsetPreserveTimestamps ensures that no value is present for PreserveTimestamps, not even an explicit nil
### GetRecoverTarget

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetRecoverTarget() PhysicalTargetParamsForRecoverFileAndFolderRecoverTarget`

GetRecoverTarget returns the RecoverTarget field if non-nil, zero value otherwise.

### GetRecoverTargetOk

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetRecoverTargetOk() (*PhysicalTargetParamsForRecoverFileAndFolderRecoverTarget, bool)`

GetRecoverTargetOk returns a tuple with the RecoverTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverTarget

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetRecoverTarget(v PhysicalTargetParamsForRecoverFileAndFolderRecoverTarget)`

SetRecoverTarget sets RecoverTarget field to given value.


### GetRestoreEntityType

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetRestoreEntityType() string`

GetRestoreEntityType returns the RestoreEntityType field if non-nil, zero value otherwise.

### GetRestoreEntityTypeOk

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetRestoreEntityTypeOk() (*string, bool)`

GetRestoreEntityTypeOk returns a tuple with the RestoreEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreEntityType

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetRestoreEntityType(v string)`

SetRestoreEntityType sets RestoreEntityType field to given value.

### HasRestoreEntityType

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) HasRestoreEntityType() bool`

HasRestoreEntityType returns a boolean if a field has been set.

### SetRestoreEntityTypeNil

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetRestoreEntityTypeNil(b bool)`

 SetRestoreEntityTypeNil sets the value for RestoreEntityType to be an explicit nil

### UnsetRestoreEntityType
`func (o *PhysicalTargetParamsForRecoverFileAndFolder) UnsetRestoreEntityType()`

UnsetRestoreEntityType ensures that no value is present for RestoreEntityType, not even an explicit nil
### GetRestoreExclusionParam

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetRestoreExclusionParam() []string`

GetRestoreExclusionParam returns the RestoreExclusionParam field if non-nil, zero value otherwise.

### GetRestoreExclusionParamOk

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetRestoreExclusionParamOk() (*[]string, bool)`

GetRestoreExclusionParamOk returns a tuple with the RestoreExclusionParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreExclusionParam

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetRestoreExclusionParam(v []string)`

SetRestoreExclusionParam sets RestoreExclusionParam field to given value.

### HasRestoreExclusionParam

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) HasRestoreExclusionParam() bool`

HasRestoreExclusionParam returns a boolean if a field has been set.

### SetRestoreExclusionParamNil

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetRestoreExclusionParamNil(b bool)`

 SetRestoreExclusionParamNil sets the value for RestoreExclusionParam to be an explicit nil

### UnsetRestoreExclusionParam
`func (o *PhysicalTargetParamsForRecoverFileAndFolder) UnsetRestoreExclusionParam()`

UnsetRestoreExclusionParam ensures that no value is present for RestoreExclusionParam, not even an explicit nil
### GetRestoreToOriginalPaths

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetRestoreToOriginalPaths() bool`

GetRestoreToOriginalPaths returns the RestoreToOriginalPaths field if non-nil, zero value otherwise.

### GetRestoreToOriginalPathsOk

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetRestoreToOriginalPathsOk() (*bool, bool)`

GetRestoreToOriginalPathsOk returns a tuple with the RestoreToOriginalPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreToOriginalPaths

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetRestoreToOriginalPaths(v bool)`

SetRestoreToOriginalPaths sets RestoreToOriginalPaths field to given value.

### HasRestoreToOriginalPaths

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) HasRestoreToOriginalPaths() bool`

HasRestoreToOriginalPaths returns a boolean if a field has been set.

### SetRestoreToOriginalPathsNil

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetRestoreToOriginalPathsNil(b bool)`

 SetRestoreToOriginalPathsNil sets the value for RestoreToOriginalPaths to be an explicit nil

### UnsetRestoreToOriginalPaths
`func (o *PhysicalTargetParamsForRecoverFileAndFolder) UnsetRestoreToOriginalPaths()`

UnsetRestoreToOriginalPaths ensures that no value is present for RestoreToOriginalPaths, not even an explicit nil
### GetSaveSuccessFiles

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetSaveSuccessFiles() bool`

GetSaveSuccessFiles returns the SaveSuccessFiles field if non-nil, zero value otherwise.

### GetSaveSuccessFilesOk

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetSaveSuccessFilesOk() (*bool, bool)`

GetSaveSuccessFilesOk returns a tuple with the SaveSuccessFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveSuccessFiles

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetSaveSuccessFiles(v bool)`

SetSaveSuccessFiles sets SaveSuccessFiles field to given value.

### HasSaveSuccessFiles

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) HasSaveSuccessFiles() bool`

HasSaveSuccessFiles returns a boolean if a field has been set.

### SetSaveSuccessFilesNil

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetSaveSuccessFilesNil(b bool)`

 SetSaveSuccessFilesNil sets the value for SaveSuccessFiles to be an explicit nil

### UnsetSaveSuccessFiles
`func (o *PhysicalTargetParamsForRecoverFileAndFolder) UnsetSaveSuccessFiles()`

UnsetSaveSuccessFiles ensures that no value is present for SaveSuccessFiles, not even an explicit nil
### GetVlanConfig

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetVlanConfig() HyperVTargetParamsForRecoverVmVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) GetVlanConfigOk() (*HyperVTargetParamsForRecoverVmVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetVlanConfig(v HyperVTargetParamsForRecoverVmVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### SetVlanConfigNil

`func (o *PhysicalTargetParamsForRecoverFileAndFolder) SetVlanConfigNil(b bool)`

 SetVlanConfigNil sets the value for VlanConfig to be an explicit nil

### UnsetVlanConfig
`func (o *PhysicalTargetParamsForRecoverFileAndFolder) UnsetVlanConfig()`

UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


