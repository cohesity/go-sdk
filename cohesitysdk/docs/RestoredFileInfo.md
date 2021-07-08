# RestoredFileInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsolutePath** | Pointer to **NullableString** | Full path of the file being restored: the actual file path without the disk. E.g.: \\Program Files\\App\\file.txt | [optional] 
**AttachedDiskId** | Pointer to **NullableInt32** | Disk information of where the source file is currently located. | [optional] 
**DiskPartitionId** | Pointer to **NullableInt32** | Disk partition to which the file belongs to. | [optional] 
**FsUuid** | Pointer to **NullableString** | File system UUID on which file resides. | [optional] 
**InodeNumber** | Pointer to **NullableInt64** | Inode number of the file. This is needed for snapmirror restore workflow. | [optional] 
**IsDirectory** | Pointer to **NullableBool** | Whether the path points to a directory. | [optional] 
**IsNonSimpleLdmVol** | Pointer to **NullableBool** | This will be set to true for recovery workflows for non-simple volumes on Windows Dynamic Disks. In that case, we will use VolumeInfo instead of some of the details captured here (e.g. virtual_disk_file) for determining disk and volume related details. | [optional] 
**RestoreBaseDirectory** | Pointer to **NullableString** | This must be set to a directory path if restore_to_original_paths is false and restore task has multiple files which are not desired to be restore to one common location. If this filed is populated, &#39;absolute_path&#39; will be restored under this location. If this field is not populated all files in restore task will be restored to location specified in RestoreFilesPreferences. | [optional] 
**RestoreMountPoint** | Pointer to **NullableString** | Mount point of the volume on which the file to be restored is located. E.g.: c:\\temp\\vhd_mount_1234 | [optional] 
**SizeBytes** | Pointer to **NullableInt64** | Size of the file in bytes. Required in FLR in GCP using Cloud Functions. | [optional] 
**VirtualDiskFile** | Pointer to **NullableString** | Virtual disk file to which this file belongs to. | [optional] 
**VolumeId** | Pointer to **NullableString** | Id of the volume. | [optional] 
**VolumePath** | Pointer to **NullableString** | Original volume name (or drive letter). This is used while performing the copy to the original paths. E.g.: c: | [optional] 

## Methods

### NewRestoredFileInfo

`func NewRestoredFileInfo() *RestoredFileInfo`

NewRestoredFileInfo instantiates a new RestoredFileInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoredFileInfoWithDefaults

`func NewRestoredFileInfoWithDefaults() *RestoredFileInfo`

NewRestoredFileInfoWithDefaults instantiates a new RestoredFileInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsolutePath

`func (o *RestoredFileInfo) GetAbsolutePath() string`

GetAbsolutePath returns the AbsolutePath field if non-nil, zero value otherwise.

### GetAbsolutePathOk

`func (o *RestoredFileInfo) GetAbsolutePathOk() (*string, bool)`

GetAbsolutePathOk returns a tuple with the AbsolutePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsolutePath

`func (o *RestoredFileInfo) SetAbsolutePath(v string)`

SetAbsolutePath sets AbsolutePath field to given value.

### HasAbsolutePath

`func (o *RestoredFileInfo) HasAbsolutePath() bool`

HasAbsolutePath returns a boolean if a field has been set.

### SetAbsolutePathNil

`func (o *RestoredFileInfo) SetAbsolutePathNil(b bool)`

 SetAbsolutePathNil sets the value for AbsolutePath to be an explicit nil

### UnsetAbsolutePath
`func (o *RestoredFileInfo) UnsetAbsolutePath()`

UnsetAbsolutePath ensures that no value is present for AbsolutePath, not even an explicit nil
### GetAttachedDiskId

`func (o *RestoredFileInfo) GetAttachedDiskId() int32`

GetAttachedDiskId returns the AttachedDiskId field if non-nil, zero value otherwise.

### GetAttachedDiskIdOk

`func (o *RestoredFileInfo) GetAttachedDiskIdOk() (*int32, bool)`

GetAttachedDiskIdOk returns a tuple with the AttachedDiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedDiskId

`func (o *RestoredFileInfo) SetAttachedDiskId(v int32)`

SetAttachedDiskId sets AttachedDiskId field to given value.

### HasAttachedDiskId

`func (o *RestoredFileInfo) HasAttachedDiskId() bool`

HasAttachedDiskId returns a boolean if a field has been set.

### SetAttachedDiskIdNil

`func (o *RestoredFileInfo) SetAttachedDiskIdNil(b bool)`

 SetAttachedDiskIdNil sets the value for AttachedDiskId to be an explicit nil

### UnsetAttachedDiskId
`func (o *RestoredFileInfo) UnsetAttachedDiskId()`

UnsetAttachedDiskId ensures that no value is present for AttachedDiskId, not even an explicit nil
### GetDiskPartitionId

`func (o *RestoredFileInfo) GetDiskPartitionId() int32`

GetDiskPartitionId returns the DiskPartitionId field if non-nil, zero value otherwise.

### GetDiskPartitionIdOk

`func (o *RestoredFileInfo) GetDiskPartitionIdOk() (*int32, bool)`

GetDiskPartitionIdOk returns a tuple with the DiskPartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskPartitionId

`func (o *RestoredFileInfo) SetDiskPartitionId(v int32)`

SetDiskPartitionId sets DiskPartitionId field to given value.

### HasDiskPartitionId

`func (o *RestoredFileInfo) HasDiskPartitionId() bool`

HasDiskPartitionId returns a boolean if a field has been set.

### SetDiskPartitionIdNil

`func (o *RestoredFileInfo) SetDiskPartitionIdNil(b bool)`

 SetDiskPartitionIdNil sets the value for DiskPartitionId to be an explicit nil

### UnsetDiskPartitionId
`func (o *RestoredFileInfo) UnsetDiskPartitionId()`

UnsetDiskPartitionId ensures that no value is present for DiskPartitionId, not even an explicit nil
### GetFsUuid

`func (o *RestoredFileInfo) GetFsUuid() string`

GetFsUuid returns the FsUuid field if non-nil, zero value otherwise.

### GetFsUuidOk

`func (o *RestoredFileInfo) GetFsUuidOk() (*string, bool)`

GetFsUuidOk returns a tuple with the FsUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUuid

`func (o *RestoredFileInfo) SetFsUuid(v string)`

SetFsUuid sets FsUuid field to given value.

### HasFsUuid

`func (o *RestoredFileInfo) HasFsUuid() bool`

HasFsUuid returns a boolean if a field has been set.

### SetFsUuidNil

`func (o *RestoredFileInfo) SetFsUuidNil(b bool)`

 SetFsUuidNil sets the value for FsUuid to be an explicit nil

### UnsetFsUuid
`func (o *RestoredFileInfo) UnsetFsUuid()`

UnsetFsUuid ensures that no value is present for FsUuid, not even an explicit nil
### GetInodeNumber

`func (o *RestoredFileInfo) GetInodeNumber() int64`

GetInodeNumber returns the InodeNumber field if non-nil, zero value otherwise.

### GetInodeNumberOk

`func (o *RestoredFileInfo) GetInodeNumberOk() (*int64, bool)`

GetInodeNumberOk returns a tuple with the InodeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInodeNumber

`func (o *RestoredFileInfo) SetInodeNumber(v int64)`

SetInodeNumber sets InodeNumber field to given value.

### HasInodeNumber

`func (o *RestoredFileInfo) HasInodeNumber() bool`

HasInodeNumber returns a boolean if a field has been set.

### SetInodeNumberNil

`func (o *RestoredFileInfo) SetInodeNumberNil(b bool)`

 SetInodeNumberNil sets the value for InodeNumber to be an explicit nil

### UnsetInodeNumber
`func (o *RestoredFileInfo) UnsetInodeNumber()`

UnsetInodeNumber ensures that no value is present for InodeNumber, not even an explicit nil
### GetIsDirectory

`func (o *RestoredFileInfo) GetIsDirectory() bool`

GetIsDirectory returns the IsDirectory field if non-nil, zero value otherwise.

### GetIsDirectoryOk

`func (o *RestoredFileInfo) GetIsDirectoryOk() (*bool, bool)`

GetIsDirectoryOk returns a tuple with the IsDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectory

`func (o *RestoredFileInfo) SetIsDirectory(v bool)`

SetIsDirectory sets IsDirectory field to given value.

### HasIsDirectory

`func (o *RestoredFileInfo) HasIsDirectory() bool`

HasIsDirectory returns a boolean if a field has been set.

### SetIsDirectoryNil

`func (o *RestoredFileInfo) SetIsDirectoryNil(b bool)`

 SetIsDirectoryNil sets the value for IsDirectory to be an explicit nil

### UnsetIsDirectory
`func (o *RestoredFileInfo) UnsetIsDirectory()`

UnsetIsDirectory ensures that no value is present for IsDirectory, not even an explicit nil
### GetIsNonSimpleLdmVol

`func (o *RestoredFileInfo) GetIsNonSimpleLdmVol() bool`

GetIsNonSimpleLdmVol returns the IsNonSimpleLdmVol field if non-nil, zero value otherwise.

### GetIsNonSimpleLdmVolOk

`func (o *RestoredFileInfo) GetIsNonSimpleLdmVolOk() (*bool, bool)`

GetIsNonSimpleLdmVolOk returns a tuple with the IsNonSimpleLdmVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNonSimpleLdmVol

`func (o *RestoredFileInfo) SetIsNonSimpleLdmVol(v bool)`

SetIsNonSimpleLdmVol sets IsNonSimpleLdmVol field to given value.

### HasIsNonSimpleLdmVol

`func (o *RestoredFileInfo) HasIsNonSimpleLdmVol() bool`

HasIsNonSimpleLdmVol returns a boolean if a field has been set.

### SetIsNonSimpleLdmVolNil

`func (o *RestoredFileInfo) SetIsNonSimpleLdmVolNil(b bool)`

 SetIsNonSimpleLdmVolNil sets the value for IsNonSimpleLdmVol to be an explicit nil

### UnsetIsNonSimpleLdmVol
`func (o *RestoredFileInfo) UnsetIsNonSimpleLdmVol()`

UnsetIsNonSimpleLdmVol ensures that no value is present for IsNonSimpleLdmVol, not even an explicit nil
### GetRestoreBaseDirectory

`func (o *RestoredFileInfo) GetRestoreBaseDirectory() string`

GetRestoreBaseDirectory returns the RestoreBaseDirectory field if non-nil, zero value otherwise.

### GetRestoreBaseDirectoryOk

`func (o *RestoredFileInfo) GetRestoreBaseDirectoryOk() (*string, bool)`

GetRestoreBaseDirectoryOk returns a tuple with the RestoreBaseDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreBaseDirectory

`func (o *RestoredFileInfo) SetRestoreBaseDirectory(v string)`

SetRestoreBaseDirectory sets RestoreBaseDirectory field to given value.

### HasRestoreBaseDirectory

`func (o *RestoredFileInfo) HasRestoreBaseDirectory() bool`

HasRestoreBaseDirectory returns a boolean if a field has been set.

### SetRestoreBaseDirectoryNil

`func (o *RestoredFileInfo) SetRestoreBaseDirectoryNil(b bool)`

 SetRestoreBaseDirectoryNil sets the value for RestoreBaseDirectory to be an explicit nil

### UnsetRestoreBaseDirectory
`func (o *RestoredFileInfo) UnsetRestoreBaseDirectory()`

UnsetRestoreBaseDirectory ensures that no value is present for RestoreBaseDirectory, not even an explicit nil
### GetRestoreMountPoint

`func (o *RestoredFileInfo) GetRestoreMountPoint() string`

GetRestoreMountPoint returns the RestoreMountPoint field if non-nil, zero value otherwise.

### GetRestoreMountPointOk

`func (o *RestoredFileInfo) GetRestoreMountPointOk() (*string, bool)`

GetRestoreMountPointOk returns a tuple with the RestoreMountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreMountPoint

`func (o *RestoredFileInfo) SetRestoreMountPoint(v string)`

SetRestoreMountPoint sets RestoreMountPoint field to given value.

### HasRestoreMountPoint

`func (o *RestoredFileInfo) HasRestoreMountPoint() bool`

HasRestoreMountPoint returns a boolean if a field has been set.

### SetRestoreMountPointNil

`func (o *RestoredFileInfo) SetRestoreMountPointNil(b bool)`

 SetRestoreMountPointNil sets the value for RestoreMountPoint to be an explicit nil

### UnsetRestoreMountPoint
`func (o *RestoredFileInfo) UnsetRestoreMountPoint()`

UnsetRestoreMountPoint ensures that no value is present for RestoreMountPoint, not even an explicit nil
### GetSizeBytes

`func (o *RestoredFileInfo) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *RestoredFileInfo) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *RestoredFileInfo) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *RestoredFileInfo) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### SetSizeBytesNil

`func (o *RestoredFileInfo) SetSizeBytesNil(b bool)`

 SetSizeBytesNil sets the value for SizeBytes to be an explicit nil

### UnsetSizeBytes
`func (o *RestoredFileInfo) UnsetSizeBytes()`

UnsetSizeBytes ensures that no value is present for SizeBytes, not even an explicit nil
### GetVirtualDiskFile

`func (o *RestoredFileInfo) GetVirtualDiskFile() string`

GetVirtualDiskFile returns the VirtualDiskFile field if non-nil, zero value otherwise.

### GetVirtualDiskFileOk

`func (o *RestoredFileInfo) GetVirtualDiskFileOk() (*string, bool)`

GetVirtualDiskFileOk returns a tuple with the VirtualDiskFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDiskFile

`func (o *RestoredFileInfo) SetVirtualDiskFile(v string)`

SetVirtualDiskFile sets VirtualDiskFile field to given value.

### HasVirtualDiskFile

`func (o *RestoredFileInfo) HasVirtualDiskFile() bool`

HasVirtualDiskFile returns a boolean if a field has been set.

### SetVirtualDiskFileNil

`func (o *RestoredFileInfo) SetVirtualDiskFileNil(b bool)`

 SetVirtualDiskFileNil sets the value for VirtualDiskFile to be an explicit nil

### UnsetVirtualDiskFile
`func (o *RestoredFileInfo) UnsetVirtualDiskFile()`

UnsetVirtualDiskFile ensures that no value is present for VirtualDiskFile, not even an explicit nil
### GetVolumeId

`func (o *RestoredFileInfo) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *RestoredFileInfo) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *RestoredFileInfo) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *RestoredFileInfo) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### SetVolumeIdNil

`func (o *RestoredFileInfo) SetVolumeIdNil(b bool)`

 SetVolumeIdNil sets the value for VolumeId to be an explicit nil

### UnsetVolumeId
`func (o *RestoredFileInfo) UnsetVolumeId()`

UnsetVolumeId ensures that no value is present for VolumeId, not even an explicit nil
### GetVolumePath

`func (o *RestoredFileInfo) GetVolumePath() string`

GetVolumePath returns the VolumePath field if non-nil, zero value otherwise.

### GetVolumePathOk

`func (o *RestoredFileInfo) GetVolumePathOk() (*string, bool)`

GetVolumePathOk returns a tuple with the VolumePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumePath

`func (o *RestoredFileInfo) SetVolumePath(v string)`

SetVolumePath sets VolumePath field to given value.

### HasVolumePath

`func (o *RestoredFileInfo) HasVolumePath() bool`

HasVolumePath returns a boolean if a field has been set.

### SetVolumePathNil

`func (o *RestoredFileInfo) SetVolumePathNil(b bool)`

 SetVolumePathNil sets the value for VolumePath to be an explicit nil

### UnsetVolumePath
`func (o *RestoredFileInfo) UnsetVolumePath()`

UnsetVolumePath ensures that no value is present for VolumePath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


