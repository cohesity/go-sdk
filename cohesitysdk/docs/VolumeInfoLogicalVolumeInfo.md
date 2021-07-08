# VolumeInfoLogicalVolumeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceTree** | Pointer to [**DeviceTree**](DeviceTree.md) |  | [optional] 
**LogicalVolumeName** | Pointer to **NullableString** | Logical volume name. | [optional] 
**LogicalVolumeUuid** | Pointer to **NullableString** | Logical volume uuid. | [optional] 
**VolumeGroupName** | Pointer to **NullableString** | Volume group name. | [optional] 
**VolumeGroupUuid** | Pointer to **NullableString** | Volume group uuid. | [optional] 

## Methods

### NewVolumeInfoLogicalVolumeInfo

`func NewVolumeInfoLogicalVolumeInfo() *VolumeInfoLogicalVolumeInfo`

NewVolumeInfoLogicalVolumeInfo instantiates a new VolumeInfoLogicalVolumeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeInfoLogicalVolumeInfoWithDefaults

`func NewVolumeInfoLogicalVolumeInfoWithDefaults() *VolumeInfoLogicalVolumeInfo`

NewVolumeInfoLogicalVolumeInfoWithDefaults instantiates a new VolumeInfoLogicalVolumeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceTree

`func (o *VolumeInfoLogicalVolumeInfo) GetDeviceTree() DeviceTree`

GetDeviceTree returns the DeviceTree field if non-nil, zero value otherwise.

### GetDeviceTreeOk

`func (o *VolumeInfoLogicalVolumeInfo) GetDeviceTreeOk() (*DeviceTree, bool)`

GetDeviceTreeOk returns a tuple with the DeviceTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTree

`func (o *VolumeInfoLogicalVolumeInfo) SetDeviceTree(v DeviceTree)`

SetDeviceTree sets DeviceTree field to given value.

### HasDeviceTree

`func (o *VolumeInfoLogicalVolumeInfo) HasDeviceTree() bool`

HasDeviceTree returns a boolean if a field has been set.

### GetLogicalVolumeName

`func (o *VolumeInfoLogicalVolumeInfo) GetLogicalVolumeName() string`

GetLogicalVolumeName returns the LogicalVolumeName field if non-nil, zero value otherwise.

### GetLogicalVolumeNameOk

`func (o *VolumeInfoLogicalVolumeInfo) GetLogicalVolumeNameOk() (*string, bool)`

GetLogicalVolumeNameOk returns a tuple with the LogicalVolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalVolumeName

`func (o *VolumeInfoLogicalVolumeInfo) SetLogicalVolumeName(v string)`

SetLogicalVolumeName sets LogicalVolumeName field to given value.

### HasLogicalVolumeName

`func (o *VolumeInfoLogicalVolumeInfo) HasLogicalVolumeName() bool`

HasLogicalVolumeName returns a boolean if a field has been set.

### SetLogicalVolumeNameNil

`func (o *VolumeInfoLogicalVolumeInfo) SetLogicalVolumeNameNil(b bool)`

 SetLogicalVolumeNameNil sets the value for LogicalVolumeName to be an explicit nil

### UnsetLogicalVolumeName
`func (o *VolumeInfoLogicalVolumeInfo) UnsetLogicalVolumeName()`

UnsetLogicalVolumeName ensures that no value is present for LogicalVolumeName, not even an explicit nil
### GetLogicalVolumeUuid

`func (o *VolumeInfoLogicalVolumeInfo) GetLogicalVolumeUuid() string`

GetLogicalVolumeUuid returns the LogicalVolumeUuid field if non-nil, zero value otherwise.

### GetLogicalVolumeUuidOk

`func (o *VolumeInfoLogicalVolumeInfo) GetLogicalVolumeUuidOk() (*string, bool)`

GetLogicalVolumeUuidOk returns a tuple with the LogicalVolumeUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalVolumeUuid

`func (o *VolumeInfoLogicalVolumeInfo) SetLogicalVolumeUuid(v string)`

SetLogicalVolumeUuid sets LogicalVolumeUuid field to given value.

### HasLogicalVolumeUuid

`func (o *VolumeInfoLogicalVolumeInfo) HasLogicalVolumeUuid() bool`

HasLogicalVolumeUuid returns a boolean if a field has been set.

### SetLogicalVolumeUuidNil

`func (o *VolumeInfoLogicalVolumeInfo) SetLogicalVolumeUuidNil(b bool)`

 SetLogicalVolumeUuidNil sets the value for LogicalVolumeUuid to be an explicit nil

### UnsetLogicalVolumeUuid
`func (o *VolumeInfoLogicalVolumeInfo) UnsetLogicalVolumeUuid()`

UnsetLogicalVolumeUuid ensures that no value is present for LogicalVolumeUuid, not even an explicit nil
### GetVolumeGroupName

`func (o *VolumeInfoLogicalVolumeInfo) GetVolumeGroupName() string`

GetVolumeGroupName returns the VolumeGroupName field if non-nil, zero value otherwise.

### GetVolumeGroupNameOk

`func (o *VolumeInfoLogicalVolumeInfo) GetVolumeGroupNameOk() (*string, bool)`

GetVolumeGroupNameOk returns a tuple with the VolumeGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGroupName

`func (o *VolumeInfoLogicalVolumeInfo) SetVolumeGroupName(v string)`

SetVolumeGroupName sets VolumeGroupName field to given value.

### HasVolumeGroupName

`func (o *VolumeInfoLogicalVolumeInfo) HasVolumeGroupName() bool`

HasVolumeGroupName returns a boolean if a field has been set.

### SetVolumeGroupNameNil

`func (o *VolumeInfoLogicalVolumeInfo) SetVolumeGroupNameNil(b bool)`

 SetVolumeGroupNameNil sets the value for VolumeGroupName to be an explicit nil

### UnsetVolumeGroupName
`func (o *VolumeInfoLogicalVolumeInfo) UnsetVolumeGroupName()`

UnsetVolumeGroupName ensures that no value is present for VolumeGroupName, not even an explicit nil
### GetVolumeGroupUuid

`func (o *VolumeInfoLogicalVolumeInfo) GetVolumeGroupUuid() string`

GetVolumeGroupUuid returns the VolumeGroupUuid field if non-nil, zero value otherwise.

### GetVolumeGroupUuidOk

`func (o *VolumeInfoLogicalVolumeInfo) GetVolumeGroupUuidOk() (*string, bool)`

GetVolumeGroupUuidOk returns a tuple with the VolumeGroupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGroupUuid

`func (o *VolumeInfoLogicalVolumeInfo) SetVolumeGroupUuid(v string)`

SetVolumeGroupUuid sets VolumeGroupUuid field to given value.

### HasVolumeGroupUuid

`func (o *VolumeInfoLogicalVolumeInfo) HasVolumeGroupUuid() bool`

HasVolumeGroupUuid returns a boolean if a field has been set.

### SetVolumeGroupUuidNil

`func (o *VolumeInfoLogicalVolumeInfo) SetVolumeGroupUuidNil(b bool)`

 SetVolumeGroupUuidNil sets the value for VolumeGroupUuid to be an explicit nil

### UnsetVolumeGroupUuid
`func (o *VolumeInfoLogicalVolumeInfo) UnsetVolumeGroupUuid()`

UnsetVolumeGroupUuid ensures that no value is present for VolumeGroupUuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


