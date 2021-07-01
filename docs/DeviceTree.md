# DeviceTree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChildVec** | Pointer to [**[]DeviceTreeChildDevice**](DeviceTreeChildDevice.md) |  | [optional] 
**DeviceId** | Pointer to **NullableInt64** | Internal device identifier of the device to be activated as a thin volume. | [optional] 
**DeviceLength** | Pointer to **NullableInt64** | The length of this device. This should match the length which is computable based on children and combining strategy.  e.g. if there is only one partition slice in an LVM volume, &#39;length&#39; in the partition slice is equal to &#39;device_length&#39;. | [optional] 
**StripeSize** | Pointer to **NullableInt32** | In case data is striped, this represents the length of the stripe. The number of stripes is defined by the size of child_vec above. | [optional] 
**ThinPoolChunkSize** | Pointer to **NullableInt64** | Chunk size. Only populated if device type is thin pool. | [optional] 
**Type** | Pointer to **NullableInt32** | How to combine the children. | [optional] 

## Methods

### NewDeviceTree

`func NewDeviceTree() *DeviceTree`

NewDeviceTree instantiates a new DeviceTree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceTreeWithDefaults

`func NewDeviceTreeWithDefaults() *DeviceTree`

NewDeviceTreeWithDefaults instantiates a new DeviceTree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildVec

`func (o *DeviceTree) GetChildVec() []DeviceTreeChildDevice`

GetChildVec returns the ChildVec field if non-nil, zero value otherwise.

### GetChildVecOk

`func (o *DeviceTree) GetChildVecOk() (*[]DeviceTreeChildDevice, bool)`

GetChildVecOk returns a tuple with the ChildVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildVec

`func (o *DeviceTree) SetChildVec(v []DeviceTreeChildDevice)`

SetChildVec sets ChildVec field to given value.

### HasChildVec

`func (o *DeviceTree) HasChildVec() bool`

HasChildVec returns a boolean if a field has been set.

### SetChildVecNil

`func (o *DeviceTree) SetChildVecNil(b bool)`

 SetChildVecNil sets the value for ChildVec to be an explicit nil

### UnsetChildVec
`func (o *DeviceTree) UnsetChildVec()`

UnsetChildVec ensures that no value is present for ChildVec, not even an explicit nil
### GetDeviceId

`func (o *DeviceTree) GetDeviceId() int64`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DeviceTree) GetDeviceIdOk() (*int64, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DeviceTree) SetDeviceId(v int64)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *DeviceTree) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *DeviceTree) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *DeviceTree) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetDeviceLength

`func (o *DeviceTree) GetDeviceLength() int64`

GetDeviceLength returns the DeviceLength field if non-nil, zero value otherwise.

### GetDeviceLengthOk

`func (o *DeviceTree) GetDeviceLengthOk() (*int64, bool)`

GetDeviceLengthOk returns a tuple with the DeviceLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceLength

`func (o *DeviceTree) SetDeviceLength(v int64)`

SetDeviceLength sets DeviceLength field to given value.

### HasDeviceLength

`func (o *DeviceTree) HasDeviceLength() bool`

HasDeviceLength returns a boolean if a field has been set.

### SetDeviceLengthNil

`func (o *DeviceTree) SetDeviceLengthNil(b bool)`

 SetDeviceLengthNil sets the value for DeviceLength to be an explicit nil

### UnsetDeviceLength
`func (o *DeviceTree) UnsetDeviceLength()`

UnsetDeviceLength ensures that no value is present for DeviceLength, not even an explicit nil
### GetStripeSize

`func (o *DeviceTree) GetStripeSize() int32`

GetStripeSize returns the StripeSize field if non-nil, zero value otherwise.

### GetStripeSizeOk

`func (o *DeviceTree) GetStripeSizeOk() (*int32, bool)`

GetStripeSizeOk returns a tuple with the StripeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeSize

`func (o *DeviceTree) SetStripeSize(v int32)`

SetStripeSize sets StripeSize field to given value.

### HasStripeSize

`func (o *DeviceTree) HasStripeSize() bool`

HasStripeSize returns a boolean if a field has been set.

### SetStripeSizeNil

`func (o *DeviceTree) SetStripeSizeNil(b bool)`

 SetStripeSizeNil sets the value for StripeSize to be an explicit nil

### UnsetStripeSize
`func (o *DeviceTree) UnsetStripeSize()`

UnsetStripeSize ensures that no value is present for StripeSize, not even an explicit nil
### GetThinPoolChunkSize

`func (o *DeviceTree) GetThinPoolChunkSize() int64`

GetThinPoolChunkSize returns the ThinPoolChunkSize field if non-nil, zero value otherwise.

### GetThinPoolChunkSizeOk

`func (o *DeviceTree) GetThinPoolChunkSizeOk() (*int64, bool)`

GetThinPoolChunkSizeOk returns a tuple with the ThinPoolChunkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinPoolChunkSize

`func (o *DeviceTree) SetThinPoolChunkSize(v int64)`

SetThinPoolChunkSize sets ThinPoolChunkSize field to given value.

### HasThinPoolChunkSize

`func (o *DeviceTree) HasThinPoolChunkSize() bool`

HasThinPoolChunkSize returns a boolean if a field has been set.

### SetThinPoolChunkSizeNil

`func (o *DeviceTree) SetThinPoolChunkSizeNil(b bool)`

 SetThinPoolChunkSizeNil sets the value for ThinPoolChunkSize to be an explicit nil

### UnsetThinPoolChunkSize
`func (o *DeviceTree) UnsetThinPoolChunkSize()`

UnsetThinPoolChunkSize ensures that no value is present for ThinPoolChunkSize, not even an explicit nil
### GetType

`func (o *DeviceTree) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceTree) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceTree) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *DeviceTree) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *DeviceTree) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *DeviceTree) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


