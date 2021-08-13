# LogicalVolumeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumeGroupUuid** | Pointer to **NullableString** | Specifies the volume group uuid. | [optional] 
**VolumeGroupName** | Pointer to **NullableString** | Specifies the volume group name. | [optional] 
**LogicalVolumeUuid** | Pointer to **NullableString** | Specifies the logical volume uuid. | [optional] 
**LogicalVolumeName** | Pointer to **NullableString** | Specifies the logical volume name. | [optional] 
**DeviceTree** | Pointer to [**DeviceTreeNode**](DeviceTreeNode.md) | Specifies the tree structure of the logical volume. | [optional] 

## Methods

### NewLogicalVolumeInfo

`func NewLogicalVolumeInfo() *LogicalVolumeInfo`

NewLogicalVolumeInfo instantiates a new LogicalVolumeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogicalVolumeInfoWithDefaults

`func NewLogicalVolumeInfoWithDefaults() *LogicalVolumeInfo`

NewLogicalVolumeInfoWithDefaults instantiates a new LogicalVolumeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumeGroupUuid

`func (o *LogicalVolumeInfo) GetVolumeGroupUuid() string`

GetVolumeGroupUuid returns the VolumeGroupUuid field if non-nil, zero value otherwise.

### GetVolumeGroupUuidOk

`func (o *LogicalVolumeInfo) GetVolumeGroupUuidOk() (*string, bool)`

GetVolumeGroupUuidOk returns a tuple with the VolumeGroupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGroupUuid

`func (o *LogicalVolumeInfo) SetVolumeGroupUuid(v string)`

SetVolumeGroupUuid sets VolumeGroupUuid field to given value.

### HasVolumeGroupUuid

`func (o *LogicalVolumeInfo) HasVolumeGroupUuid() bool`

HasVolumeGroupUuid returns a boolean if a field has been set.

### SetVolumeGroupUuidNil

`func (o *LogicalVolumeInfo) SetVolumeGroupUuidNil(b bool)`

 SetVolumeGroupUuidNil sets the value for VolumeGroupUuid to be an explicit nil

### UnsetVolumeGroupUuid
`func (o *LogicalVolumeInfo) UnsetVolumeGroupUuid()`

UnsetVolumeGroupUuid ensures that no value is present for VolumeGroupUuid, not even an explicit nil
### GetVolumeGroupName

`func (o *LogicalVolumeInfo) GetVolumeGroupName() string`

GetVolumeGroupName returns the VolumeGroupName field if non-nil, zero value otherwise.

### GetVolumeGroupNameOk

`func (o *LogicalVolumeInfo) GetVolumeGroupNameOk() (*string, bool)`

GetVolumeGroupNameOk returns a tuple with the VolumeGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGroupName

`func (o *LogicalVolumeInfo) SetVolumeGroupName(v string)`

SetVolumeGroupName sets VolumeGroupName field to given value.

### HasVolumeGroupName

`func (o *LogicalVolumeInfo) HasVolumeGroupName() bool`

HasVolumeGroupName returns a boolean if a field has been set.

### SetVolumeGroupNameNil

`func (o *LogicalVolumeInfo) SetVolumeGroupNameNil(b bool)`

 SetVolumeGroupNameNil sets the value for VolumeGroupName to be an explicit nil

### UnsetVolumeGroupName
`func (o *LogicalVolumeInfo) UnsetVolumeGroupName()`

UnsetVolumeGroupName ensures that no value is present for VolumeGroupName, not even an explicit nil
### GetLogicalVolumeUuid

`func (o *LogicalVolumeInfo) GetLogicalVolumeUuid() string`

GetLogicalVolumeUuid returns the LogicalVolumeUuid field if non-nil, zero value otherwise.

### GetLogicalVolumeUuidOk

`func (o *LogicalVolumeInfo) GetLogicalVolumeUuidOk() (*string, bool)`

GetLogicalVolumeUuidOk returns a tuple with the LogicalVolumeUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalVolumeUuid

`func (o *LogicalVolumeInfo) SetLogicalVolumeUuid(v string)`

SetLogicalVolumeUuid sets LogicalVolumeUuid field to given value.

### HasLogicalVolumeUuid

`func (o *LogicalVolumeInfo) HasLogicalVolumeUuid() bool`

HasLogicalVolumeUuid returns a boolean if a field has been set.

### SetLogicalVolumeUuidNil

`func (o *LogicalVolumeInfo) SetLogicalVolumeUuidNil(b bool)`

 SetLogicalVolumeUuidNil sets the value for LogicalVolumeUuid to be an explicit nil

### UnsetLogicalVolumeUuid
`func (o *LogicalVolumeInfo) UnsetLogicalVolumeUuid()`

UnsetLogicalVolumeUuid ensures that no value is present for LogicalVolumeUuid, not even an explicit nil
### GetLogicalVolumeName

`func (o *LogicalVolumeInfo) GetLogicalVolumeName() string`

GetLogicalVolumeName returns the LogicalVolumeName field if non-nil, zero value otherwise.

### GetLogicalVolumeNameOk

`func (o *LogicalVolumeInfo) GetLogicalVolumeNameOk() (*string, bool)`

GetLogicalVolumeNameOk returns a tuple with the LogicalVolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalVolumeName

`func (o *LogicalVolumeInfo) SetLogicalVolumeName(v string)`

SetLogicalVolumeName sets LogicalVolumeName field to given value.

### HasLogicalVolumeName

`func (o *LogicalVolumeInfo) HasLogicalVolumeName() bool`

HasLogicalVolumeName returns a boolean if a field has been set.

### SetLogicalVolumeNameNil

`func (o *LogicalVolumeInfo) SetLogicalVolumeNameNil(b bool)`

 SetLogicalVolumeNameNil sets the value for LogicalVolumeName to be an explicit nil

### UnsetLogicalVolumeName
`func (o *LogicalVolumeInfo) UnsetLogicalVolumeName()`

UnsetLogicalVolumeName ensures that no value is present for LogicalVolumeName, not even an explicit nil
### GetDeviceTree

`func (o *LogicalVolumeInfo) GetDeviceTree() DeviceTreeNode`

GetDeviceTree returns the DeviceTree field if non-nil, zero value otherwise.

### GetDeviceTreeOk

`func (o *LogicalVolumeInfo) GetDeviceTreeOk() (*DeviceTreeNode, bool)`

GetDeviceTreeOk returns a tuple with the DeviceTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTree

`func (o *LogicalVolumeInfo) SetDeviceTree(v DeviceTreeNode)`

SetDeviceTree sets DeviceTree field to given value.

### HasDeviceTree

`func (o *LogicalVolumeInfo) HasDeviceTree() bool`

HasDeviceTree returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


