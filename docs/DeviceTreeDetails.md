# DeviceTreeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CombineMethod** | Pointer to **NullableString** | Specifies how to combine the children of this node. The combining strategy for child devices. Some of these strategies imply constraint on the number of child devices. e.g. RAID5 will have 5 children. &#39;LINEAR&#39; indicates children are juxtaposed to form this device. &#39;STRIPE&#39; indicates children are striped. &#39;MIRROR&#39; indicates children are mirrored. &#39;RAID5&#39; &#39;RAID6&#39; &#39;ZERO&#39; &#39;THIN&#39; &#39;THINPOOL&#39; &#39;SNAPSHOT&#39; &#39;CACHE&#39; &#39;CACHEPOOL&#39; | [optional] 
**DeviceLength** | Pointer to **NullableInt64** | Specifies the length of this device. This number should match the length that is calculated from the children and combining method. | [optional] 
**DeviceNodes** | Pointer to [**[]DeviceNode**](DeviceNode.md) | Specifies the children of this node in the device tree. | [optional] 
**StripeSize** | Pointer to **NullableInt32** | Specifies the size of the striped data if the data is striped. | [optional] 

## Methods

### NewDeviceTreeDetails

`func NewDeviceTreeDetails() *DeviceTreeDetails`

NewDeviceTreeDetails instantiates a new DeviceTreeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceTreeDetailsWithDefaults

`func NewDeviceTreeDetailsWithDefaults() *DeviceTreeDetails`

NewDeviceTreeDetailsWithDefaults instantiates a new DeviceTreeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCombineMethod

`func (o *DeviceTreeDetails) GetCombineMethod() string`

GetCombineMethod returns the CombineMethod field if non-nil, zero value otherwise.

### GetCombineMethodOk

`func (o *DeviceTreeDetails) GetCombineMethodOk() (*string, bool)`

GetCombineMethodOk returns a tuple with the CombineMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombineMethod

`func (o *DeviceTreeDetails) SetCombineMethod(v string)`

SetCombineMethod sets CombineMethod field to given value.

### HasCombineMethod

`func (o *DeviceTreeDetails) HasCombineMethod() bool`

HasCombineMethod returns a boolean if a field has been set.

### SetCombineMethodNil

`func (o *DeviceTreeDetails) SetCombineMethodNil(b bool)`

 SetCombineMethodNil sets the value for CombineMethod to be an explicit nil

### UnsetCombineMethod
`func (o *DeviceTreeDetails) UnsetCombineMethod()`

UnsetCombineMethod ensures that no value is present for CombineMethod, not even an explicit nil
### GetDeviceLength

`func (o *DeviceTreeDetails) GetDeviceLength() int64`

GetDeviceLength returns the DeviceLength field if non-nil, zero value otherwise.

### GetDeviceLengthOk

`func (o *DeviceTreeDetails) GetDeviceLengthOk() (*int64, bool)`

GetDeviceLengthOk returns a tuple with the DeviceLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceLength

`func (o *DeviceTreeDetails) SetDeviceLength(v int64)`

SetDeviceLength sets DeviceLength field to given value.

### HasDeviceLength

`func (o *DeviceTreeDetails) HasDeviceLength() bool`

HasDeviceLength returns a boolean if a field has been set.

### SetDeviceLengthNil

`func (o *DeviceTreeDetails) SetDeviceLengthNil(b bool)`

 SetDeviceLengthNil sets the value for DeviceLength to be an explicit nil

### UnsetDeviceLength
`func (o *DeviceTreeDetails) UnsetDeviceLength()`

UnsetDeviceLength ensures that no value is present for DeviceLength, not even an explicit nil
### GetDeviceNodes

`func (o *DeviceTreeDetails) GetDeviceNodes() []DeviceNode`

GetDeviceNodes returns the DeviceNodes field if non-nil, zero value otherwise.

### GetDeviceNodesOk

`func (o *DeviceTreeDetails) GetDeviceNodesOk() (*[]DeviceNode, bool)`

GetDeviceNodesOk returns a tuple with the DeviceNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceNodes

`func (o *DeviceTreeDetails) SetDeviceNodes(v []DeviceNode)`

SetDeviceNodes sets DeviceNodes field to given value.

### HasDeviceNodes

`func (o *DeviceTreeDetails) HasDeviceNodes() bool`

HasDeviceNodes returns a boolean if a field has been set.

### SetDeviceNodesNil

`func (o *DeviceTreeDetails) SetDeviceNodesNil(b bool)`

 SetDeviceNodesNil sets the value for DeviceNodes to be an explicit nil

### UnsetDeviceNodes
`func (o *DeviceTreeDetails) UnsetDeviceNodes()`

UnsetDeviceNodes ensures that no value is present for DeviceNodes, not even an explicit nil
### GetStripeSize

`func (o *DeviceTreeDetails) GetStripeSize() int32`

GetStripeSize returns the StripeSize field if non-nil, zero value otherwise.

### GetStripeSizeOk

`func (o *DeviceTreeDetails) GetStripeSizeOk() (*int32, bool)`

GetStripeSizeOk returns a tuple with the StripeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeSize

`func (o *DeviceTreeDetails) SetStripeSize(v int32)`

SetStripeSize sets StripeSize field to given value.

### HasStripeSize

`func (o *DeviceTreeDetails) HasStripeSize() bool`

HasStripeSize returns a boolean if a field has been set.

### SetStripeSizeNil

`func (o *DeviceTreeDetails) SetStripeSizeNil(b bool)`

 SetStripeSizeNil sets the value for StripeSize to be an explicit nil

### UnsetStripeSize
`func (o *DeviceTreeDetails) UnsetStripeSize()`

UnsetStripeSize ensures that no value is present for StripeSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


