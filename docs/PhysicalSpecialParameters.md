# PhysicalSpecialParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationParameters** | Pointer to [**ApplicationParameters**](ApplicationParameters.md) |  | [optional] 
**EnableSystemBackup** | Pointer to **NullableBool** | Specifies whether to allow system backup using 3rd party tools installed on the Protection Host. System backups are used for doing bare metal recovery later. This field is applicable only for System backups. | [optional] 
**FilePaths** | Pointer to [**[]FilePathParameters**](FilePathParameters.md) | Array of File Paths to Back Up.  Specifies a list of directories or files to protect in a Physical Server. | [optional] 
**MetadataFilePath** | Pointer to **NullableString** | Specifies metadata path on source. This file contains absolute paths of files that needs to be backed up on the same source. | [optional] 
**SkipNestedVolumesVec** | Pointer to **[]string** | Specifies mounttypes of nested volumes to be skipped. | [optional] 
**UsesSkipNestedVolumesVec** | Pointer to **NullableBool** | Specifies whether to use SkipNestedVolumes vec to skip nested mounts. | [optional] 
**VolumeGuid** | Pointer to **[]string** | Array of Mounted Volumes to Back Up.  Specifies the subset of mounted volumes to protect in a Physical Server. If not specified, all mounted volumes on a Physical Server are protected. | [optional] 
**WindowsParameters** | Pointer to [**WindowsHostSnapshotParameters**](WindowsHostSnapshotParameters.md) |  | [optional] 

## Methods

### NewPhysicalSpecialParameters

`func NewPhysicalSpecialParameters() *PhysicalSpecialParameters`

NewPhysicalSpecialParameters instantiates a new PhysicalSpecialParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalSpecialParametersWithDefaults

`func NewPhysicalSpecialParametersWithDefaults() *PhysicalSpecialParameters`

NewPhysicalSpecialParametersWithDefaults instantiates a new PhysicalSpecialParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationParameters

`func (o *PhysicalSpecialParameters) GetApplicationParameters() ApplicationParameters`

GetApplicationParameters returns the ApplicationParameters field if non-nil, zero value otherwise.

### GetApplicationParametersOk

`func (o *PhysicalSpecialParameters) GetApplicationParametersOk() (*ApplicationParameters, bool)`

GetApplicationParametersOk returns a tuple with the ApplicationParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationParameters

`func (o *PhysicalSpecialParameters) SetApplicationParameters(v ApplicationParameters)`

SetApplicationParameters sets ApplicationParameters field to given value.

### HasApplicationParameters

`func (o *PhysicalSpecialParameters) HasApplicationParameters() bool`

HasApplicationParameters returns a boolean if a field has been set.

### GetEnableSystemBackup

`func (o *PhysicalSpecialParameters) GetEnableSystemBackup() bool`

GetEnableSystemBackup returns the EnableSystemBackup field if non-nil, zero value otherwise.

### GetEnableSystemBackupOk

`func (o *PhysicalSpecialParameters) GetEnableSystemBackupOk() (*bool, bool)`

GetEnableSystemBackupOk returns a tuple with the EnableSystemBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSystemBackup

`func (o *PhysicalSpecialParameters) SetEnableSystemBackup(v bool)`

SetEnableSystemBackup sets EnableSystemBackup field to given value.

### HasEnableSystemBackup

`func (o *PhysicalSpecialParameters) HasEnableSystemBackup() bool`

HasEnableSystemBackup returns a boolean if a field has been set.

### SetEnableSystemBackupNil

`func (o *PhysicalSpecialParameters) SetEnableSystemBackupNil(b bool)`

 SetEnableSystemBackupNil sets the value for EnableSystemBackup to be an explicit nil

### UnsetEnableSystemBackup
`func (o *PhysicalSpecialParameters) UnsetEnableSystemBackup()`

UnsetEnableSystemBackup ensures that no value is present for EnableSystemBackup, not even an explicit nil
### GetFilePaths

`func (o *PhysicalSpecialParameters) GetFilePaths() []FilePathParameters`

GetFilePaths returns the FilePaths field if non-nil, zero value otherwise.

### GetFilePathsOk

`func (o *PhysicalSpecialParameters) GetFilePathsOk() (*[]FilePathParameters, bool)`

GetFilePathsOk returns a tuple with the FilePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePaths

`func (o *PhysicalSpecialParameters) SetFilePaths(v []FilePathParameters)`

SetFilePaths sets FilePaths field to given value.

### HasFilePaths

`func (o *PhysicalSpecialParameters) HasFilePaths() bool`

HasFilePaths returns a boolean if a field has been set.

### SetFilePathsNil

`func (o *PhysicalSpecialParameters) SetFilePathsNil(b bool)`

 SetFilePathsNil sets the value for FilePaths to be an explicit nil

### UnsetFilePaths
`func (o *PhysicalSpecialParameters) UnsetFilePaths()`

UnsetFilePaths ensures that no value is present for FilePaths, not even an explicit nil
### GetMetadataFilePath

`func (o *PhysicalSpecialParameters) GetMetadataFilePath() string`

GetMetadataFilePath returns the MetadataFilePath field if non-nil, zero value otherwise.

### GetMetadataFilePathOk

`func (o *PhysicalSpecialParameters) GetMetadataFilePathOk() (*string, bool)`

GetMetadataFilePathOk returns a tuple with the MetadataFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFilePath

`func (o *PhysicalSpecialParameters) SetMetadataFilePath(v string)`

SetMetadataFilePath sets MetadataFilePath field to given value.

### HasMetadataFilePath

`func (o *PhysicalSpecialParameters) HasMetadataFilePath() bool`

HasMetadataFilePath returns a boolean if a field has been set.

### SetMetadataFilePathNil

`func (o *PhysicalSpecialParameters) SetMetadataFilePathNil(b bool)`

 SetMetadataFilePathNil sets the value for MetadataFilePath to be an explicit nil

### UnsetMetadataFilePath
`func (o *PhysicalSpecialParameters) UnsetMetadataFilePath()`

UnsetMetadataFilePath ensures that no value is present for MetadataFilePath, not even an explicit nil
### GetSkipNestedVolumesVec

`func (o *PhysicalSpecialParameters) GetSkipNestedVolumesVec() []string`

GetSkipNestedVolumesVec returns the SkipNestedVolumesVec field if non-nil, zero value otherwise.

### GetSkipNestedVolumesVecOk

`func (o *PhysicalSpecialParameters) GetSkipNestedVolumesVecOk() (*[]string, bool)`

GetSkipNestedVolumesVecOk returns a tuple with the SkipNestedVolumesVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipNestedVolumesVec

`func (o *PhysicalSpecialParameters) SetSkipNestedVolumesVec(v []string)`

SetSkipNestedVolumesVec sets SkipNestedVolumesVec field to given value.

### HasSkipNestedVolumesVec

`func (o *PhysicalSpecialParameters) HasSkipNestedVolumesVec() bool`

HasSkipNestedVolumesVec returns a boolean if a field has been set.

### SetSkipNestedVolumesVecNil

`func (o *PhysicalSpecialParameters) SetSkipNestedVolumesVecNil(b bool)`

 SetSkipNestedVolumesVecNil sets the value for SkipNestedVolumesVec to be an explicit nil

### UnsetSkipNestedVolumesVec
`func (o *PhysicalSpecialParameters) UnsetSkipNestedVolumesVec()`

UnsetSkipNestedVolumesVec ensures that no value is present for SkipNestedVolumesVec, not even an explicit nil
### GetUsesSkipNestedVolumesVec

`func (o *PhysicalSpecialParameters) GetUsesSkipNestedVolumesVec() bool`

GetUsesSkipNestedVolumesVec returns the UsesSkipNestedVolumesVec field if non-nil, zero value otherwise.

### GetUsesSkipNestedVolumesVecOk

`func (o *PhysicalSpecialParameters) GetUsesSkipNestedVolumesVecOk() (*bool, bool)`

GetUsesSkipNestedVolumesVecOk returns a tuple with the UsesSkipNestedVolumesVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesSkipNestedVolumesVec

`func (o *PhysicalSpecialParameters) SetUsesSkipNestedVolumesVec(v bool)`

SetUsesSkipNestedVolumesVec sets UsesSkipNestedVolumesVec field to given value.

### HasUsesSkipNestedVolumesVec

`func (o *PhysicalSpecialParameters) HasUsesSkipNestedVolumesVec() bool`

HasUsesSkipNestedVolumesVec returns a boolean if a field has been set.

### SetUsesSkipNestedVolumesVecNil

`func (o *PhysicalSpecialParameters) SetUsesSkipNestedVolumesVecNil(b bool)`

 SetUsesSkipNestedVolumesVecNil sets the value for UsesSkipNestedVolumesVec to be an explicit nil

### UnsetUsesSkipNestedVolumesVec
`func (o *PhysicalSpecialParameters) UnsetUsesSkipNestedVolumesVec()`

UnsetUsesSkipNestedVolumesVec ensures that no value is present for UsesSkipNestedVolumesVec, not even an explicit nil
### GetVolumeGuid

`func (o *PhysicalSpecialParameters) GetVolumeGuid() []string`

GetVolumeGuid returns the VolumeGuid field if non-nil, zero value otherwise.

### GetVolumeGuidOk

`func (o *PhysicalSpecialParameters) GetVolumeGuidOk() (*[]string, bool)`

GetVolumeGuidOk returns a tuple with the VolumeGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGuid

`func (o *PhysicalSpecialParameters) SetVolumeGuid(v []string)`

SetVolumeGuid sets VolumeGuid field to given value.

### HasVolumeGuid

`func (o *PhysicalSpecialParameters) HasVolumeGuid() bool`

HasVolumeGuid returns a boolean if a field has been set.

### SetVolumeGuidNil

`func (o *PhysicalSpecialParameters) SetVolumeGuidNil(b bool)`

 SetVolumeGuidNil sets the value for VolumeGuid to be an explicit nil

### UnsetVolumeGuid
`func (o *PhysicalSpecialParameters) UnsetVolumeGuid()`

UnsetVolumeGuid ensures that no value is present for VolumeGuid, not even an explicit nil
### GetWindowsParameters

`func (o *PhysicalSpecialParameters) GetWindowsParameters() WindowsHostSnapshotParameters`

GetWindowsParameters returns the WindowsParameters field if non-nil, zero value otherwise.

### GetWindowsParametersOk

`func (o *PhysicalSpecialParameters) GetWindowsParametersOk() (*WindowsHostSnapshotParameters, bool)`

GetWindowsParametersOk returns a tuple with the WindowsParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsParameters

`func (o *PhysicalSpecialParameters) SetWindowsParameters(v WindowsHostSnapshotParameters)`

SetWindowsParameters sets WindowsParameters field to given value.

### HasWindowsParameters

`func (o *PhysicalSpecialParameters) HasWindowsParameters() bool`

HasWindowsParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


