# DeviceTreeChildDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to [**DeviceTree**](DeviceTree.md) |  | [optional] 
**DeviceType** | Pointer to **NullableInt32** | This specifies how the parent device is using this child device. | [optional] 
**PartitionSlice** | Pointer to [**DeviceTreePartitionSlice**](DeviceTreePartitionSlice.md) |  | [optional] 

## Methods

### NewDeviceTreeChildDevice

`func NewDeviceTreeChildDevice() *DeviceTreeChildDevice`

NewDeviceTreeChildDevice instantiates a new DeviceTreeChildDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceTreeChildDeviceWithDefaults

`func NewDeviceTreeChildDeviceWithDefaults() *DeviceTreeChildDevice`

NewDeviceTreeChildDeviceWithDefaults instantiates a new DeviceTreeChildDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *DeviceTreeChildDevice) GetDevice() DeviceTree`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *DeviceTreeChildDevice) GetDeviceOk() (*DeviceTree, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *DeviceTreeChildDevice) SetDevice(v DeviceTree)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *DeviceTreeChildDevice) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDeviceType

`func (o *DeviceTreeChildDevice) GetDeviceType() int32`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DeviceTreeChildDevice) GetDeviceTypeOk() (*int32, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DeviceTreeChildDevice) SetDeviceType(v int32)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *DeviceTreeChildDevice) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *DeviceTreeChildDevice) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *DeviceTreeChildDevice) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetPartitionSlice

`func (o *DeviceTreeChildDevice) GetPartitionSlice() DeviceTreePartitionSlice`

GetPartitionSlice returns the PartitionSlice field if non-nil, zero value otherwise.

### GetPartitionSliceOk

`func (o *DeviceTreeChildDevice) GetPartitionSliceOk() (*DeviceTreePartitionSlice, bool)`

GetPartitionSliceOk returns a tuple with the PartitionSlice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionSlice

`func (o *DeviceTreeChildDevice) SetPartitionSlice(v DeviceTreePartitionSlice)`

SetPartitionSlice sets PartitionSlice field to given value.

### HasPartitionSlice

`func (o *DeviceTreeChildDevice) HasPartitionSlice() bool`

HasPartitionSlice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


