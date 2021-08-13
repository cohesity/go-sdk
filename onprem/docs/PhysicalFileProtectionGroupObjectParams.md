# PhysicalFileProtectionGroupObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Specifies the ID of the object protected. | 
**Name** | Pointer to **NullableString** | Specifies the name of the object protected. | [optional] [readonly] 
**FilePaths** | Pointer to [**[]PhysicalFileBackupPathParams**](PhysicalFileBackupPathParams.md) | Specifies a list of file paths to be protected by this Protection Group. | [optional] 
**UsesPathLevelSkipNestedVolumeSetting** | Pointer to **NullableBool** | Specifies whether path level or object level skip nested volume setting will be used. | [optional] 
**NestedVolumeTypesToSkip** | Pointer to **[]string** | Specifies mount types of nested volumes to be skipped. | [optional] 
**FollowNasSymlinkTarget** | Pointer to **NullableBool** | Specifies whether to follow NAS target pointed by symlink for windows sources. | [optional] 
**MetadataFilePath** | Pointer to **NullableString** | Specifies the path of metadatafile on source. This file contains absolute paths of files that needs to be backed up on the same source. | [optional] 

## Methods

### NewPhysicalFileProtectionGroupObjectParams

`func NewPhysicalFileProtectionGroupObjectParams(id int64, ) *PhysicalFileProtectionGroupObjectParams`

NewPhysicalFileProtectionGroupObjectParams instantiates a new PhysicalFileProtectionGroupObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalFileProtectionGroupObjectParamsWithDefaults

`func NewPhysicalFileProtectionGroupObjectParamsWithDefaults() *PhysicalFileProtectionGroupObjectParams`

NewPhysicalFileProtectionGroupObjectParamsWithDefaults instantiates a new PhysicalFileProtectionGroupObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PhysicalFileProtectionGroupObjectParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PhysicalFileProtectionGroupObjectParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PhysicalFileProtectionGroupObjectParams) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *PhysicalFileProtectionGroupObjectParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PhysicalFileProtectionGroupObjectParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PhysicalFileProtectionGroupObjectParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PhysicalFileProtectionGroupObjectParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PhysicalFileProtectionGroupObjectParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PhysicalFileProtectionGroupObjectParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFilePaths

`func (o *PhysicalFileProtectionGroupObjectParams) GetFilePaths() []PhysicalFileBackupPathParams`

GetFilePaths returns the FilePaths field if non-nil, zero value otherwise.

### GetFilePathsOk

`func (o *PhysicalFileProtectionGroupObjectParams) GetFilePathsOk() (*[]PhysicalFileBackupPathParams, bool)`

GetFilePathsOk returns a tuple with the FilePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePaths

`func (o *PhysicalFileProtectionGroupObjectParams) SetFilePaths(v []PhysicalFileBackupPathParams)`

SetFilePaths sets FilePaths field to given value.

### HasFilePaths

`func (o *PhysicalFileProtectionGroupObjectParams) HasFilePaths() bool`

HasFilePaths returns a boolean if a field has been set.

### GetUsesPathLevelSkipNestedVolumeSetting

`func (o *PhysicalFileProtectionGroupObjectParams) GetUsesPathLevelSkipNestedVolumeSetting() bool`

GetUsesPathLevelSkipNestedVolumeSetting returns the UsesPathLevelSkipNestedVolumeSetting field if non-nil, zero value otherwise.

### GetUsesPathLevelSkipNestedVolumeSettingOk

`func (o *PhysicalFileProtectionGroupObjectParams) GetUsesPathLevelSkipNestedVolumeSettingOk() (*bool, bool)`

GetUsesPathLevelSkipNestedVolumeSettingOk returns a tuple with the UsesPathLevelSkipNestedVolumeSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesPathLevelSkipNestedVolumeSetting

`func (o *PhysicalFileProtectionGroupObjectParams) SetUsesPathLevelSkipNestedVolumeSetting(v bool)`

SetUsesPathLevelSkipNestedVolumeSetting sets UsesPathLevelSkipNestedVolumeSetting field to given value.

### HasUsesPathLevelSkipNestedVolumeSetting

`func (o *PhysicalFileProtectionGroupObjectParams) HasUsesPathLevelSkipNestedVolumeSetting() bool`

HasUsesPathLevelSkipNestedVolumeSetting returns a boolean if a field has been set.

### SetUsesPathLevelSkipNestedVolumeSettingNil

`func (o *PhysicalFileProtectionGroupObjectParams) SetUsesPathLevelSkipNestedVolumeSettingNil(b bool)`

 SetUsesPathLevelSkipNestedVolumeSettingNil sets the value for UsesPathLevelSkipNestedVolumeSetting to be an explicit nil

### UnsetUsesPathLevelSkipNestedVolumeSetting
`func (o *PhysicalFileProtectionGroupObjectParams) UnsetUsesPathLevelSkipNestedVolumeSetting()`

UnsetUsesPathLevelSkipNestedVolumeSetting ensures that no value is present for UsesPathLevelSkipNestedVolumeSetting, not even an explicit nil
### GetNestedVolumeTypesToSkip

`func (o *PhysicalFileProtectionGroupObjectParams) GetNestedVolumeTypesToSkip() []string`

GetNestedVolumeTypesToSkip returns the NestedVolumeTypesToSkip field if non-nil, zero value otherwise.

### GetNestedVolumeTypesToSkipOk

`func (o *PhysicalFileProtectionGroupObjectParams) GetNestedVolumeTypesToSkipOk() (*[]string, bool)`

GetNestedVolumeTypesToSkipOk returns a tuple with the NestedVolumeTypesToSkip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedVolumeTypesToSkip

`func (o *PhysicalFileProtectionGroupObjectParams) SetNestedVolumeTypesToSkip(v []string)`

SetNestedVolumeTypesToSkip sets NestedVolumeTypesToSkip field to given value.

### HasNestedVolumeTypesToSkip

`func (o *PhysicalFileProtectionGroupObjectParams) HasNestedVolumeTypesToSkip() bool`

HasNestedVolumeTypesToSkip returns a boolean if a field has been set.

### GetFollowNasSymlinkTarget

`func (o *PhysicalFileProtectionGroupObjectParams) GetFollowNasSymlinkTarget() bool`

GetFollowNasSymlinkTarget returns the FollowNasSymlinkTarget field if non-nil, zero value otherwise.

### GetFollowNasSymlinkTargetOk

`func (o *PhysicalFileProtectionGroupObjectParams) GetFollowNasSymlinkTargetOk() (*bool, bool)`

GetFollowNasSymlinkTargetOk returns a tuple with the FollowNasSymlinkTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowNasSymlinkTarget

`func (o *PhysicalFileProtectionGroupObjectParams) SetFollowNasSymlinkTarget(v bool)`

SetFollowNasSymlinkTarget sets FollowNasSymlinkTarget field to given value.

### HasFollowNasSymlinkTarget

`func (o *PhysicalFileProtectionGroupObjectParams) HasFollowNasSymlinkTarget() bool`

HasFollowNasSymlinkTarget returns a boolean if a field has been set.

### SetFollowNasSymlinkTargetNil

`func (o *PhysicalFileProtectionGroupObjectParams) SetFollowNasSymlinkTargetNil(b bool)`

 SetFollowNasSymlinkTargetNil sets the value for FollowNasSymlinkTarget to be an explicit nil

### UnsetFollowNasSymlinkTarget
`func (o *PhysicalFileProtectionGroupObjectParams) UnsetFollowNasSymlinkTarget()`

UnsetFollowNasSymlinkTarget ensures that no value is present for FollowNasSymlinkTarget, not even an explicit nil
### GetMetadataFilePath

`func (o *PhysicalFileProtectionGroupObjectParams) GetMetadataFilePath() string`

GetMetadataFilePath returns the MetadataFilePath field if non-nil, zero value otherwise.

### GetMetadataFilePathOk

`func (o *PhysicalFileProtectionGroupObjectParams) GetMetadataFilePathOk() (*string, bool)`

GetMetadataFilePathOk returns a tuple with the MetadataFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFilePath

`func (o *PhysicalFileProtectionGroupObjectParams) SetMetadataFilePath(v string)`

SetMetadataFilePath sets MetadataFilePath field to given value.

### HasMetadataFilePath

`func (o *PhysicalFileProtectionGroupObjectParams) HasMetadataFilePath() bool`

HasMetadataFilePath returns a boolean if a field has been set.

### SetMetadataFilePathNil

`func (o *PhysicalFileProtectionGroupObjectParams) SetMetadataFilePathNil(b bool)`

 SetMetadataFilePathNil sets the value for MetadataFilePath to be an explicit nil

### UnsetMetadataFilePath
`func (o *PhysicalFileProtectionGroupObjectParams) UnsetMetadataFilePath()`

UnsetMetadataFilePath ensures that no value is present for MetadataFilePath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


