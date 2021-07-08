# VolumeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskVec** | Pointer to [**[]VolumeInfoDiskInfo**](VolumeInfoDiskInfo.md) | Information about all the disks and partitions needed to mount this logical volume. | [optional] 
**DisplayName** | Pointer to **NullableString** | Display name. | [optional] 
**FilesystemType** | Pointer to **NullableString** | Filesystem on this volume. | [optional] 
**FsLabel** | Pointer to **NullableString** | Filesystem label. | [optional] 
**FsUuid** | Pointer to **NullableString** | Filesystem uuid. | [optional] 
**IsBootable** | Pointer to **NullableBool** | Is this volume bootable? | [optional] 
**IsDedup** | Pointer to **NullableBool** | Is this a dedup volume? Currently, set to true only for ntfs dedup volume. | [optional] 
**IsSupported** | Pointer to **NullableBool** | Is this a supported Volume (filesystem)? | [optional] 
**LvInfo** | Pointer to [**VolumeInfoLogicalVolumeInfo**](VolumeInfoLogicalVolumeInfo.md) |  | [optional] 
**VolumeGuid** | Pointer to **NullableString** | The guid of the volume represented by this virtual disk. This information will be originally populated by magneto for physical environments. | [optional] 
**VolumeType** | Pointer to **NullableInt32** | Whether this volume is simple, lvm or ldm. | [optional] 

## Methods

### NewVolumeInfo

`func NewVolumeInfo() *VolumeInfo`

NewVolumeInfo instantiates a new VolumeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeInfoWithDefaults

`func NewVolumeInfoWithDefaults() *VolumeInfo`

NewVolumeInfoWithDefaults instantiates a new VolumeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskVec

`func (o *VolumeInfo) GetDiskVec() []VolumeInfoDiskInfo`

GetDiskVec returns the DiskVec field if non-nil, zero value otherwise.

### GetDiskVecOk

`func (o *VolumeInfo) GetDiskVecOk() (*[]VolumeInfoDiskInfo, bool)`

GetDiskVecOk returns a tuple with the DiskVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskVec

`func (o *VolumeInfo) SetDiskVec(v []VolumeInfoDiskInfo)`

SetDiskVec sets DiskVec field to given value.

### HasDiskVec

`func (o *VolumeInfo) HasDiskVec() bool`

HasDiskVec returns a boolean if a field has been set.

### SetDiskVecNil

`func (o *VolumeInfo) SetDiskVecNil(b bool)`

 SetDiskVecNil sets the value for DiskVec to be an explicit nil

### UnsetDiskVec
`func (o *VolumeInfo) UnsetDiskVec()`

UnsetDiskVec ensures that no value is present for DiskVec, not even an explicit nil
### GetDisplayName

`func (o *VolumeInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VolumeInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VolumeInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *VolumeInfo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *VolumeInfo) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *VolumeInfo) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetFilesystemType

`func (o *VolumeInfo) GetFilesystemType() string`

GetFilesystemType returns the FilesystemType field if non-nil, zero value otherwise.

### GetFilesystemTypeOk

`func (o *VolumeInfo) GetFilesystemTypeOk() (*string, bool)`

GetFilesystemTypeOk returns a tuple with the FilesystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemType

`func (o *VolumeInfo) SetFilesystemType(v string)`

SetFilesystemType sets FilesystemType field to given value.

### HasFilesystemType

`func (o *VolumeInfo) HasFilesystemType() bool`

HasFilesystemType returns a boolean if a field has been set.

### SetFilesystemTypeNil

`func (o *VolumeInfo) SetFilesystemTypeNil(b bool)`

 SetFilesystemTypeNil sets the value for FilesystemType to be an explicit nil

### UnsetFilesystemType
`func (o *VolumeInfo) UnsetFilesystemType()`

UnsetFilesystemType ensures that no value is present for FilesystemType, not even an explicit nil
### GetFsLabel

`func (o *VolumeInfo) GetFsLabel() string`

GetFsLabel returns the FsLabel field if non-nil, zero value otherwise.

### GetFsLabelOk

`func (o *VolumeInfo) GetFsLabelOk() (*string, bool)`

GetFsLabelOk returns a tuple with the FsLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsLabel

`func (o *VolumeInfo) SetFsLabel(v string)`

SetFsLabel sets FsLabel field to given value.

### HasFsLabel

`func (o *VolumeInfo) HasFsLabel() bool`

HasFsLabel returns a boolean if a field has been set.

### SetFsLabelNil

`func (o *VolumeInfo) SetFsLabelNil(b bool)`

 SetFsLabelNil sets the value for FsLabel to be an explicit nil

### UnsetFsLabel
`func (o *VolumeInfo) UnsetFsLabel()`

UnsetFsLabel ensures that no value is present for FsLabel, not even an explicit nil
### GetFsUuid

`func (o *VolumeInfo) GetFsUuid() string`

GetFsUuid returns the FsUuid field if non-nil, zero value otherwise.

### GetFsUuidOk

`func (o *VolumeInfo) GetFsUuidOk() (*string, bool)`

GetFsUuidOk returns a tuple with the FsUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsUuid

`func (o *VolumeInfo) SetFsUuid(v string)`

SetFsUuid sets FsUuid field to given value.

### HasFsUuid

`func (o *VolumeInfo) HasFsUuid() bool`

HasFsUuid returns a boolean if a field has been set.

### SetFsUuidNil

`func (o *VolumeInfo) SetFsUuidNil(b bool)`

 SetFsUuidNil sets the value for FsUuid to be an explicit nil

### UnsetFsUuid
`func (o *VolumeInfo) UnsetFsUuid()`

UnsetFsUuid ensures that no value is present for FsUuid, not even an explicit nil
### GetIsBootable

`func (o *VolumeInfo) GetIsBootable() bool`

GetIsBootable returns the IsBootable field if non-nil, zero value otherwise.

### GetIsBootableOk

`func (o *VolumeInfo) GetIsBootableOk() (*bool, bool)`

GetIsBootableOk returns a tuple with the IsBootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBootable

`func (o *VolumeInfo) SetIsBootable(v bool)`

SetIsBootable sets IsBootable field to given value.

### HasIsBootable

`func (o *VolumeInfo) HasIsBootable() bool`

HasIsBootable returns a boolean if a field has been set.

### SetIsBootableNil

`func (o *VolumeInfo) SetIsBootableNil(b bool)`

 SetIsBootableNil sets the value for IsBootable to be an explicit nil

### UnsetIsBootable
`func (o *VolumeInfo) UnsetIsBootable()`

UnsetIsBootable ensures that no value is present for IsBootable, not even an explicit nil
### GetIsDedup

`func (o *VolumeInfo) GetIsDedup() bool`

GetIsDedup returns the IsDedup field if non-nil, zero value otherwise.

### GetIsDedupOk

`func (o *VolumeInfo) GetIsDedupOk() (*bool, bool)`

GetIsDedupOk returns a tuple with the IsDedup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDedup

`func (o *VolumeInfo) SetIsDedup(v bool)`

SetIsDedup sets IsDedup field to given value.

### HasIsDedup

`func (o *VolumeInfo) HasIsDedup() bool`

HasIsDedup returns a boolean if a field has been set.

### SetIsDedupNil

`func (o *VolumeInfo) SetIsDedupNil(b bool)`

 SetIsDedupNil sets the value for IsDedup to be an explicit nil

### UnsetIsDedup
`func (o *VolumeInfo) UnsetIsDedup()`

UnsetIsDedup ensures that no value is present for IsDedup, not even an explicit nil
### GetIsSupported

`func (o *VolumeInfo) GetIsSupported() bool`

GetIsSupported returns the IsSupported field if non-nil, zero value otherwise.

### GetIsSupportedOk

`func (o *VolumeInfo) GetIsSupportedOk() (*bool, bool)`

GetIsSupportedOk returns a tuple with the IsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupported

`func (o *VolumeInfo) SetIsSupported(v bool)`

SetIsSupported sets IsSupported field to given value.

### HasIsSupported

`func (o *VolumeInfo) HasIsSupported() bool`

HasIsSupported returns a boolean if a field has been set.

### SetIsSupportedNil

`func (o *VolumeInfo) SetIsSupportedNil(b bool)`

 SetIsSupportedNil sets the value for IsSupported to be an explicit nil

### UnsetIsSupported
`func (o *VolumeInfo) UnsetIsSupported()`

UnsetIsSupported ensures that no value is present for IsSupported, not even an explicit nil
### GetLvInfo

`func (o *VolumeInfo) GetLvInfo() VolumeInfoLogicalVolumeInfo`

GetLvInfo returns the LvInfo field if non-nil, zero value otherwise.

### GetLvInfoOk

`func (o *VolumeInfo) GetLvInfoOk() (*VolumeInfoLogicalVolumeInfo, bool)`

GetLvInfoOk returns a tuple with the LvInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvInfo

`func (o *VolumeInfo) SetLvInfo(v VolumeInfoLogicalVolumeInfo)`

SetLvInfo sets LvInfo field to given value.

### HasLvInfo

`func (o *VolumeInfo) HasLvInfo() bool`

HasLvInfo returns a boolean if a field has been set.

### GetVolumeGuid

`func (o *VolumeInfo) GetVolumeGuid() string`

GetVolumeGuid returns the VolumeGuid field if non-nil, zero value otherwise.

### GetVolumeGuidOk

`func (o *VolumeInfo) GetVolumeGuidOk() (*string, bool)`

GetVolumeGuidOk returns a tuple with the VolumeGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGuid

`func (o *VolumeInfo) SetVolumeGuid(v string)`

SetVolumeGuid sets VolumeGuid field to given value.

### HasVolumeGuid

`func (o *VolumeInfo) HasVolumeGuid() bool`

HasVolumeGuid returns a boolean if a field has been set.

### SetVolumeGuidNil

`func (o *VolumeInfo) SetVolumeGuidNil(b bool)`

 SetVolumeGuidNil sets the value for VolumeGuid to be an explicit nil

### UnsetVolumeGuid
`func (o *VolumeInfo) UnsetVolumeGuid()`

UnsetVolumeGuid ensures that no value is present for VolumeGuid, not even an explicit nil
### GetVolumeType

`func (o *VolumeInfo) GetVolumeType() int32`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *VolumeInfo) GetVolumeTypeOk() (*int32, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *VolumeInfo) SetVolumeType(v int32)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *VolumeInfo) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### SetVolumeTypeNil

`func (o *VolumeInfo) SetVolumeTypeNil(b bool)`

 SetVolumeTypeNil sets the value for VolumeType to be an explicit nil

### UnsetVolumeType
`func (o *VolumeInfo) UnsetVolumeType()`

UnsetVolumeType ensures that no value is present for VolumeType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


