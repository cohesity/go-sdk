# PhysicalFileBackupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupPathInfoVec** | Pointer to [**[]PhysicalFileBackupParamsBackupPathInfo**](PhysicalFileBackupParamsBackupPathInfo.md) | Specifies the paths to backup on the Physical source. | [optional] 
**MetadataFilePath** | Pointer to **NullableString** | Specifies metadata path on source. This file contains absolute paths of files that needs to be backed up on the same source. If this field is set, backup_path_info_vec will be ignored. | [optional] 
**SkipNestedVolumesVec** | Pointer to **[]string** | Mount types of nested volumes to be skipped. | [optional] 
**SymlinkFollowNasTarget** | Pointer to **NullableBool** | Specifies whether to follow nas target pointed by symlink. Set to true only for windows physical file based job. | [optional] 
**UsesSkipNestedVolumesVec** | Pointer to **NullableBool** | Specifies whether to use skip_nested_volumes_vec to skip nested mounts. Before 6.4, BackupPathInfo.skip_nested_volumes boolean was used to skip nested volumes. So we use this boolean to support older jobs. | [optional] 

## Methods

### NewPhysicalFileBackupParams

`func NewPhysicalFileBackupParams() *PhysicalFileBackupParams`

NewPhysicalFileBackupParams instantiates a new PhysicalFileBackupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalFileBackupParamsWithDefaults

`func NewPhysicalFileBackupParamsWithDefaults() *PhysicalFileBackupParams`

NewPhysicalFileBackupParamsWithDefaults instantiates a new PhysicalFileBackupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupPathInfoVec

`func (o *PhysicalFileBackupParams) GetBackupPathInfoVec() []PhysicalFileBackupParamsBackupPathInfo`

GetBackupPathInfoVec returns the BackupPathInfoVec field if non-nil, zero value otherwise.

### GetBackupPathInfoVecOk

`func (o *PhysicalFileBackupParams) GetBackupPathInfoVecOk() (*[]PhysicalFileBackupParamsBackupPathInfo, bool)`

GetBackupPathInfoVecOk returns a tuple with the BackupPathInfoVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupPathInfoVec

`func (o *PhysicalFileBackupParams) SetBackupPathInfoVec(v []PhysicalFileBackupParamsBackupPathInfo)`

SetBackupPathInfoVec sets BackupPathInfoVec field to given value.

### HasBackupPathInfoVec

`func (o *PhysicalFileBackupParams) HasBackupPathInfoVec() bool`

HasBackupPathInfoVec returns a boolean if a field has been set.

### SetBackupPathInfoVecNil

`func (o *PhysicalFileBackupParams) SetBackupPathInfoVecNil(b bool)`

 SetBackupPathInfoVecNil sets the value for BackupPathInfoVec to be an explicit nil

### UnsetBackupPathInfoVec
`func (o *PhysicalFileBackupParams) UnsetBackupPathInfoVec()`

UnsetBackupPathInfoVec ensures that no value is present for BackupPathInfoVec, not even an explicit nil
### GetMetadataFilePath

`func (o *PhysicalFileBackupParams) GetMetadataFilePath() string`

GetMetadataFilePath returns the MetadataFilePath field if non-nil, zero value otherwise.

### GetMetadataFilePathOk

`func (o *PhysicalFileBackupParams) GetMetadataFilePathOk() (*string, bool)`

GetMetadataFilePathOk returns a tuple with the MetadataFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFilePath

`func (o *PhysicalFileBackupParams) SetMetadataFilePath(v string)`

SetMetadataFilePath sets MetadataFilePath field to given value.

### HasMetadataFilePath

`func (o *PhysicalFileBackupParams) HasMetadataFilePath() bool`

HasMetadataFilePath returns a boolean if a field has been set.

### SetMetadataFilePathNil

`func (o *PhysicalFileBackupParams) SetMetadataFilePathNil(b bool)`

 SetMetadataFilePathNil sets the value for MetadataFilePath to be an explicit nil

### UnsetMetadataFilePath
`func (o *PhysicalFileBackupParams) UnsetMetadataFilePath()`

UnsetMetadataFilePath ensures that no value is present for MetadataFilePath, not even an explicit nil
### GetSkipNestedVolumesVec

`func (o *PhysicalFileBackupParams) GetSkipNestedVolumesVec() []string`

GetSkipNestedVolumesVec returns the SkipNestedVolumesVec field if non-nil, zero value otherwise.

### GetSkipNestedVolumesVecOk

`func (o *PhysicalFileBackupParams) GetSkipNestedVolumesVecOk() (*[]string, bool)`

GetSkipNestedVolumesVecOk returns a tuple with the SkipNestedVolumesVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipNestedVolumesVec

`func (o *PhysicalFileBackupParams) SetSkipNestedVolumesVec(v []string)`

SetSkipNestedVolumesVec sets SkipNestedVolumesVec field to given value.

### HasSkipNestedVolumesVec

`func (o *PhysicalFileBackupParams) HasSkipNestedVolumesVec() bool`

HasSkipNestedVolumesVec returns a boolean if a field has been set.

### SetSkipNestedVolumesVecNil

`func (o *PhysicalFileBackupParams) SetSkipNestedVolumesVecNil(b bool)`

 SetSkipNestedVolumesVecNil sets the value for SkipNestedVolumesVec to be an explicit nil

### UnsetSkipNestedVolumesVec
`func (o *PhysicalFileBackupParams) UnsetSkipNestedVolumesVec()`

UnsetSkipNestedVolumesVec ensures that no value is present for SkipNestedVolumesVec, not even an explicit nil
### GetSymlinkFollowNasTarget

`func (o *PhysicalFileBackupParams) GetSymlinkFollowNasTarget() bool`

GetSymlinkFollowNasTarget returns the SymlinkFollowNasTarget field if non-nil, zero value otherwise.

### GetSymlinkFollowNasTargetOk

`func (o *PhysicalFileBackupParams) GetSymlinkFollowNasTargetOk() (*bool, bool)`

GetSymlinkFollowNasTargetOk returns a tuple with the SymlinkFollowNasTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymlinkFollowNasTarget

`func (o *PhysicalFileBackupParams) SetSymlinkFollowNasTarget(v bool)`

SetSymlinkFollowNasTarget sets SymlinkFollowNasTarget field to given value.

### HasSymlinkFollowNasTarget

`func (o *PhysicalFileBackupParams) HasSymlinkFollowNasTarget() bool`

HasSymlinkFollowNasTarget returns a boolean if a field has been set.

### SetSymlinkFollowNasTargetNil

`func (o *PhysicalFileBackupParams) SetSymlinkFollowNasTargetNil(b bool)`

 SetSymlinkFollowNasTargetNil sets the value for SymlinkFollowNasTarget to be an explicit nil

### UnsetSymlinkFollowNasTarget
`func (o *PhysicalFileBackupParams) UnsetSymlinkFollowNasTarget()`

UnsetSymlinkFollowNasTarget ensures that no value is present for SymlinkFollowNasTarget, not even an explicit nil
### GetUsesSkipNestedVolumesVec

`func (o *PhysicalFileBackupParams) GetUsesSkipNestedVolumesVec() bool`

GetUsesSkipNestedVolumesVec returns the UsesSkipNestedVolumesVec field if non-nil, zero value otherwise.

### GetUsesSkipNestedVolumesVecOk

`func (o *PhysicalFileBackupParams) GetUsesSkipNestedVolumesVecOk() (*bool, bool)`

GetUsesSkipNestedVolumesVecOk returns a tuple with the UsesSkipNestedVolumesVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesSkipNestedVolumesVec

`func (o *PhysicalFileBackupParams) SetUsesSkipNestedVolumesVec(v bool)`

SetUsesSkipNestedVolumesVec sets UsesSkipNestedVolumesVec field to given value.

### HasUsesSkipNestedVolumesVec

`func (o *PhysicalFileBackupParams) HasUsesSkipNestedVolumesVec() bool`

HasUsesSkipNestedVolumesVec returns a boolean if a field has been set.

### SetUsesSkipNestedVolumesVecNil

`func (o *PhysicalFileBackupParams) SetUsesSkipNestedVolumesVecNil(b bool)`

 SetUsesSkipNestedVolumesVecNil sets the value for UsesSkipNestedVolumesVec to be an explicit nil

### UnsetUsesSkipNestedVolumesVec
`func (o *PhysicalFileBackupParams) UnsetUsesSkipNestedVolumesVec()`

UnsetUsesSkipNestedVolumesVec ensures that no value is present for UsesSkipNestedVolumesVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


