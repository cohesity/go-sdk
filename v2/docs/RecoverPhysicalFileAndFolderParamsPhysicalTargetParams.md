# RecoverPhysicalFileAndFolderParamsPhysicalTargetParams

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

### NewRecoverPhysicalFileAndFolderParamsPhysicalTargetParams

`func NewRecoverPhysicalFileAndFolderParamsPhysicalTargetParams(recoverTarget PhysicalTargetParamsForRecoverFileAndFolderRecoverTarget, ) *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams`

NewRecoverPhysicalFileAndFolderParamsPhysicalTargetParams instantiates a new RecoverPhysicalFileAndFolderParamsPhysicalTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverPhysicalFileAndFolderParamsPhysicalTargetParamsWithDefaults

`func NewRecoverPhysicalFileAndFolderParamsPhysicalTargetParamsWithDefaults() *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams`

NewRecoverPhysicalFileAndFolderParamsPhysicalTargetParamsWithDefaults instantiates a new RecoverPhysicalFileAndFolderParamsPhysicalTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternateRestoreDirectory

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) GetAlternateRestoreDirectory() string`

GetAlternateRestoreDirectory returns the AlternateRestoreDirectory field if non-nil, zero value otherwise.

### GetAlternateRestoreDirectoryOk

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) GetAlternateRestoreDirectoryOk() (*string, bool)`

GetAlternateRestoreDirectoryOk returns a tuple with the AlternateRestoreDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateRestoreDirectory

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) SetAlternateRestoreDirectory(v string)`

SetAlternateRestoreDirectory sets AlternateRestoreDirectory field to given value.

### HasAlternateRestoreDirectory

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) HasAlternateRestoreDirectory() bool`

HasAlternateRestoreDirectory returns a boolean if a field has been set.

### SetAlternateRestoreDirectoryNil

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) SetAlternateRestoreDirectoryNil(b bool)`

 SetAlternateRestoreDirectoryNil sets the value for AlternateRestoreDirectory to be an explicit nil

### UnsetAlternateRestoreDirectory
`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) UnsetAlternateRestoreDirectory()`

UnsetAlternateRestoreDirectory ensures that no value is present for AlternateRestoreDirectory, not even an explicit nil
### GetContinueOnError

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetOverwriteExisting

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) GetOverwriteExisting() bool`

GetOverwriteExisting returns the OverwriteExisting field if non-nil, zero value otherwise.

### GetOverwriteExistingOk

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) GetOverwriteExistingOk() (*bool, bool)`

GetOverwriteExistingOk returns a tuple with the OverwriteExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteExisting

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) SetOverwriteExisting(v bool)`

SetOverwriteExisting sets OverwriteExisting field to given value.

### HasOverwriteExisting

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) HasOverwriteExisting() bool`

HasOverwriteExisting returns a boolean if a field has been set.

### SetOverwriteExistingNil

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) SetOverwriteExistingNil(b bool)`

 SetOverwriteExistingNil sets the value for OverwriteExisting to be an explicit nil

### UnsetOverwriteExisting
`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) UnsetOverwriteExisting()`

UnsetOverwriteExisting ensures that no value is present for OverwriteExisting, not even an explicit nil
### GetPreserveAcls

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) GetPreserveAcls() bool`

GetPreserveAcls returns the PreserveAcls field if non-nil, zero value otherwise.

### GetPreserveAclsOk

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) GetPreserveAclsOk() (*bool, bool)`

GetPreserveAclsOk returns a tuple with the PreserveAcls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveAcls

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) SetPreserveAcls(v bool)`

SetPreserveAcls sets PreserveAcls field to given value.

### HasPreserveAcls

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) HasPreserveAcls() bool`

HasPreserveAcls returns a boolean if a field has been set.

### SetPreserveAclsNil

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) SetPreserveAclsNil(b bool)`

 SetPreserveAclsNil sets the value for PreserveAcls to be an explicit nil

### UnsetPreserveAcls
`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) UnsetPreserveAcls()`

UnsetPreserveAcls ensures that no value is present for PreserveAcls, not even an explicit nil
### GetPreserveAttributes

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) GetPreserveAttributes() bool`

GetPreserveAttributes returns the PreserveAttributes field if non-nil, zero value otherwise.

### GetPreserveAttributesOk

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) GetPreserveAttributesOk() (*bool, bool)`

GetPreserveAttributesOk returns a tuple with the PreserveAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveAttributes

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) SetPreserveAttributes(v bool)`

SetPreserveAttributes sets PreserveAttributes field to given value.

### HasPreserveAttributes

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) HasPreserveAttributes() bool`

HasPreserveAttributes returns a boolean if a field has been set.

### SetPreserveAttributesNil

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) SetPreserveAttributesNil(b bool)`

 SetPreserveAttributesNil sets the value for PreserveAttributes to be an explicit nil

### UnsetPreserveAttributes
`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) UnsetPreserveAttributes()`

UnsetPreserveAttributes ensures that no value is present for PreserveAttributes, not even an explicit nil
### GetPreserveTimestamps

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) GetPreserveTimestamps() bool`

GetPreserveTimestamps returns the PreserveTimestamps field if non-nil, zero value otherwise.

### GetPreserveTimestampsOk

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) GetPreserveTimestampsOk() (*bool, bool)`

GetPreserveTimestampsOk returns a tuple with the PreserveTimestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveTimestamps

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) SetPreserveTimestamps(v bool)`

SetPreserveTimestamps sets PreserveTimestamps field to given value.

### HasPreserveTimestamps

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) HasPreserveTimestamps() bool`

HasPreserveTimestamps returns a boolean if a field has been set.

### SetPreserveTimestampsNil

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) SetPreserveTimestampsNil(b bool)`

 SetPreserveTimestampsNil sets the value for PreserveTimestamps to be an explicit nil

### UnsetPreserveTimestamps
`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) UnsetPreserveTimestamps()`

UnsetPreserveTimestamps ensures that no value is present for PreserveTimestamps, not even an explicit nil
### GetRecoverTarget

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) GetRecoverTarget() PhysicalTargetParamsForRecoverFileAndFolderRecoverTarget`

GetRecoverTarget returns the RecoverTarget field if non-nil, zero value otherwise.

### GetRecoverTargetOk

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) GetRecoverTargetOk() (*PhysicalTargetParamsForRecoverFileAndFolderRecoverTarget, bool)`

GetRecoverTargetOk returns a tuple with the RecoverTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverTarget

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) SetRecoverTarget(v PhysicalTargetParamsForRecoverFileAndFolderRecoverTarget)`

SetRecoverTarget sets RecoverTarget field to given value.


### GetRestoreEntityType

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) GetRestoreEntityType() string`

GetRestoreEntityType returns the RestoreEntityType field if non-nil, zero value otherwise.

### GetRestoreEntityTypeOk

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) GetRestoreEntityTypeOk() (*string, bool)`

GetRestoreEntityTypeOk returns a tuple with the RestoreEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreEntityType

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) SetRestoreEntityType(v string)`

SetRestoreEntityType sets RestoreEntityType field to given value.

### HasRestoreEntityType

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) HasRestoreEntityType() bool`

HasRestoreEntityType returns a boolean if a field has been set.

### SetRestoreEntityTypeNil

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) SetRestoreEntityTypeNil(b bool)`

 SetRestoreEntityTypeNil sets the value for RestoreEntityType to be an explicit nil

### UnsetRestoreEntityType
`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) UnsetRestoreEntityType()`

UnsetRestoreEntityType ensures that no value is present for RestoreEntityType, not even an explicit nil
### GetRestoreExclusionParam

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) GetRestoreExclusionParam() []string`

GetRestoreExclusionParam returns the RestoreExclusionParam field if non-nil, zero value otherwise.

### GetRestoreExclusionParamOk

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) GetRestoreExclusionParamOk() (*[]string, bool)`

GetRestoreExclusionParamOk returns a tuple with the RestoreExclusionParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreExclusionParam

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) SetRestoreExclusionParam(v []string)`

SetRestoreExclusionParam sets RestoreExclusionParam field to given value.

### HasRestoreExclusionParam

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) HasRestoreExclusionParam() bool`

HasRestoreExclusionParam returns a boolean if a field has been set.

### SetRestoreExclusionParamNil

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) SetRestoreExclusionParamNil(b bool)`

 SetRestoreExclusionParamNil sets the value for RestoreExclusionParam to be an explicit nil

### UnsetRestoreExclusionParam
`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) UnsetRestoreExclusionParam()`

UnsetRestoreExclusionParam ensures that no value is present for RestoreExclusionParam, not even an explicit nil
### GetRestoreToOriginalPaths

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) GetRestoreToOriginalPaths() bool`

GetRestoreToOriginalPaths returns the RestoreToOriginalPaths field if non-nil, zero value otherwise.

### GetRestoreToOriginalPathsOk

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) GetRestoreToOriginalPathsOk() (*bool, bool)`

GetRestoreToOriginalPathsOk returns a tuple with the RestoreToOriginalPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreToOriginalPaths

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) SetRestoreToOriginalPaths(v bool)`

SetRestoreToOriginalPaths sets RestoreToOriginalPaths field to given value.

### HasRestoreToOriginalPaths

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) HasRestoreToOriginalPaths() bool`

HasRestoreToOriginalPaths returns a boolean if a field has been set.

### SetRestoreToOriginalPathsNil

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) SetRestoreToOriginalPathsNil(b bool)`

 SetRestoreToOriginalPathsNil sets the value for RestoreToOriginalPaths to be an explicit nil

### UnsetRestoreToOriginalPaths
`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) UnsetRestoreToOriginalPaths()`

UnsetRestoreToOriginalPaths ensures that no value is present for RestoreToOriginalPaths, not even an explicit nil
### GetSaveSuccessFiles

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) GetSaveSuccessFiles() bool`

GetSaveSuccessFiles returns the SaveSuccessFiles field if non-nil, zero value otherwise.

### GetSaveSuccessFilesOk

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) GetSaveSuccessFilesOk() (*bool, bool)`

GetSaveSuccessFilesOk returns a tuple with the SaveSuccessFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveSuccessFiles

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) SetSaveSuccessFiles(v bool)`

SetSaveSuccessFiles sets SaveSuccessFiles field to given value.

### HasSaveSuccessFiles

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) HasSaveSuccessFiles() bool`

HasSaveSuccessFiles returns a boolean if a field has been set.

### SetSaveSuccessFilesNil

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) SetSaveSuccessFilesNil(b bool)`

 SetSaveSuccessFilesNil sets the value for SaveSuccessFiles to be an explicit nil

### UnsetSaveSuccessFiles
`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) UnsetSaveSuccessFiles()`

UnsetSaveSuccessFiles ensures that no value is present for SaveSuccessFiles, not even an explicit nil
### GetVlanConfig

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) GetVlanConfig() HyperVTargetParamsForRecoverVmVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) GetVlanConfigOk() (*HyperVTargetParamsForRecoverVmVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) SetVlanConfig(v HyperVTargetParamsForRecoverVmVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### SetVlanConfigNil

`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) SetVlanConfigNil(b bool)`

 SetVlanConfigNil sets the value for VlanConfig to be an explicit nil

### UnsetVlanConfig
`func (o *RecoverPhysicalFileAndFolderParamsPhysicalTargetParams) UnsetVlanConfig()`

UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


