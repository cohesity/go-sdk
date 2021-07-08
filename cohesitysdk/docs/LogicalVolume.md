# LogicalVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceRootNode** | Pointer to [**DeviceTreeDetails**](DeviceTreeDetails.md) |  | [optional] 
**GroupName** | Pointer to **NullableString** | Specifies the group name of the logical volume. | [optional] 
**GroupUuid** | Pointer to **NullableString** | Specifies the group uuid of the logical volume. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the logical volume. | [optional] 
**Uuid** | Pointer to **NullableString** | Specifies the uuid of the logical volume. | [optional] 

## Methods

### NewLogicalVolume

`func NewLogicalVolume() *LogicalVolume`

NewLogicalVolume instantiates a new LogicalVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogicalVolumeWithDefaults

`func NewLogicalVolumeWithDefaults() *LogicalVolume`

NewLogicalVolumeWithDefaults instantiates a new LogicalVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceRootNode

`func (o *LogicalVolume) GetDeviceRootNode() DeviceTreeDetails`

GetDeviceRootNode returns the DeviceRootNode field if non-nil, zero value otherwise.

### GetDeviceRootNodeOk

`func (o *LogicalVolume) GetDeviceRootNodeOk() (*DeviceTreeDetails, bool)`

GetDeviceRootNodeOk returns a tuple with the DeviceRootNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRootNode

`func (o *LogicalVolume) SetDeviceRootNode(v DeviceTreeDetails)`

SetDeviceRootNode sets DeviceRootNode field to given value.

### HasDeviceRootNode

`func (o *LogicalVolume) HasDeviceRootNode() bool`

HasDeviceRootNode returns a boolean if a field has been set.

### GetGroupName

`func (o *LogicalVolume) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *LogicalVolume) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *LogicalVolume) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *LogicalVolume) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *LogicalVolume) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *LogicalVolume) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetGroupUuid

`func (o *LogicalVolume) GetGroupUuid() string`

GetGroupUuid returns the GroupUuid field if non-nil, zero value otherwise.

### GetGroupUuidOk

`func (o *LogicalVolume) GetGroupUuidOk() (*string, bool)`

GetGroupUuidOk returns a tuple with the GroupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupUuid

`func (o *LogicalVolume) SetGroupUuid(v string)`

SetGroupUuid sets GroupUuid field to given value.

### HasGroupUuid

`func (o *LogicalVolume) HasGroupUuid() bool`

HasGroupUuid returns a boolean if a field has been set.

### SetGroupUuidNil

`func (o *LogicalVolume) SetGroupUuidNil(b bool)`

 SetGroupUuidNil sets the value for GroupUuid to be an explicit nil

### UnsetGroupUuid
`func (o *LogicalVolume) UnsetGroupUuid()`

UnsetGroupUuid ensures that no value is present for GroupUuid, not even an explicit nil
### GetName

`func (o *LogicalVolume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogicalVolume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogicalVolume) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LogicalVolume) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *LogicalVolume) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *LogicalVolume) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUuid

`func (o *LogicalVolume) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *LogicalVolume) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *LogicalVolume) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *LogicalVolume) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *LogicalVolume) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *LogicalVolume) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


