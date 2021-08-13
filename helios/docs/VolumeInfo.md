# VolumeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies the volume name. | [optional] 
**IsSupported** | Pointer to **NullableBool** | Specifies if this volume is supported. | [optional] 
**VolumeType** | Pointer to **NullableString** | Specifies the volume type. | [optional] 
**FilesystemType** | Pointer to **NullableString** | Specifies the filesystem type. | [optional] 
**FilesystemUuid** | Pointer to **NullableString** | Specifies the filesystem uuid. | [optional] 
**VolumeGuid** | Pointer to **NullableString** | Specifies the volume guid. | [optional] 
**VolumeSizeInBytes** | Pointer to **NullableInt64** | Specifies volume size in bytes. | [optional] 
**LogicalVolumeInfo** | Pointer to [**LogicalVolumeInfo**](LogicalVolumeInfo.md) | Specifies the logical volume info. This fields is for &#39;LVM&#39; and &#39;LDM&#39; volume type only. | [optional] 

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

### GetName

`func (o *VolumeInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VolumeInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VolumeInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VolumeInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
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
### GetVolumeType

`func (o *VolumeInfo) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *VolumeInfo) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *VolumeInfo) SetVolumeType(v string)`

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
### GetFilesystemUuid

`func (o *VolumeInfo) GetFilesystemUuid() string`

GetFilesystemUuid returns the FilesystemUuid field if non-nil, zero value otherwise.

### GetFilesystemUuidOk

`func (o *VolumeInfo) GetFilesystemUuidOk() (*string, bool)`

GetFilesystemUuidOk returns a tuple with the FilesystemUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemUuid

`func (o *VolumeInfo) SetFilesystemUuid(v string)`

SetFilesystemUuid sets FilesystemUuid field to given value.

### HasFilesystemUuid

`func (o *VolumeInfo) HasFilesystemUuid() bool`

HasFilesystemUuid returns a boolean if a field has been set.

### SetFilesystemUuidNil

`func (o *VolumeInfo) SetFilesystemUuidNil(b bool)`

 SetFilesystemUuidNil sets the value for FilesystemUuid to be an explicit nil

### UnsetFilesystemUuid
`func (o *VolumeInfo) UnsetFilesystemUuid()`

UnsetFilesystemUuid ensures that no value is present for FilesystemUuid, not even an explicit nil
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
### GetVolumeSizeInBytes

`func (o *VolumeInfo) GetVolumeSizeInBytes() int64`

GetVolumeSizeInBytes returns the VolumeSizeInBytes field if non-nil, zero value otherwise.

### GetVolumeSizeInBytesOk

`func (o *VolumeInfo) GetVolumeSizeInBytesOk() (*int64, bool)`

GetVolumeSizeInBytesOk returns a tuple with the VolumeSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSizeInBytes

`func (o *VolumeInfo) SetVolumeSizeInBytes(v int64)`

SetVolumeSizeInBytes sets VolumeSizeInBytes field to given value.

### HasVolumeSizeInBytes

`func (o *VolumeInfo) HasVolumeSizeInBytes() bool`

HasVolumeSizeInBytes returns a boolean if a field has been set.

### SetVolumeSizeInBytesNil

`func (o *VolumeInfo) SetVolumeSizeInBytesNil(b bool)`

 SetVolumeSizeInBytesNil sets the value for VolumeSizeInBytes to be an explicit nil

### UnsetVolumeSizeInBytes
`func (o *VolumeInfo) UnsetVolumeSizeInBytes()`

UnsetVolumeSizeInBytes ensures that no value is present for VolumeSizeInBytes, not even an explicit nil
### GetLogicalVolumeInfo

`func (o *VolumeInfo) GetLogicalVolumeInfo() LogicalVolumeInfo`

GetLogicalVolumeInfo returns the LogicalVolumeInfo field if non-nil, zero value otherwise.

### GetLogicalVolumeInfoOk

`func (o *VolumeInfo) GetLogicalVolumeInfoOk() (*LogicalVolumeInfo, bool)`

GetLogicalVolumeInfoOk returns a tuple with the LogicalVolumeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalVolumeInfo

`func (o *VolumeInfo) SetLogicalVolumeInfo(v LogicalVolumeInfo)`

SetLogicalVolumeInfo sets LogicalVolumeInfo field to given value.

### HasLogicalVolumeInfo

`func (o *VolumeInfo) HasLogicalVolumeInfo() bool`

HasLogicalVolumeInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


