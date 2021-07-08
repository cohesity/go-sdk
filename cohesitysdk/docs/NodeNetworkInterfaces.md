# NodeNetworkInterfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChassisSerial** | Pointer to **NullableString** | Specifies the serial number of Chassis. | [optional] 
**Interfaces** | Pointer to [**[]NetworkInterface**](NetworkInterface.md) | Specifies the list of network interfaces present on this Node. | [optional] 
**Message** | Pointer to **NullableString** | Specifies an optional message describing the result of the request pertaining to this Node. | [optional] 
**NodeId** | Pointer to **NullableInt64** | Specifies the ID of the Node. | [optional] 
**Slot** | Pointer to **NullableInt64** | Specifies the slot number the Node is located in. | [optional] 

## Methods

### NewNodeNetworkInterfaces

`func NewNodeNetworkInterfaces() *NodeNetworkInterfaces`

NewNodeNetworkInterfaces instantiates a new NodeNetworkInterfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeNetworkInterfacesWithDefaults

`func NewNodeNetworkInterfacesWithDefaults() *NodeNetworkInterfaces`

NewNodeNetworkInterfacesWithDefaults instantiates a new NodeNetworkInterfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChassisSerial

`func (o *NodeNetworkInterfaces) GetChassisSerial() string`

GetChassisSerial returns the ChassisSerial field if non-nil, zero value otherwise.

### GetChassisSerialOk

`func (o *NodeNetworkInterfaces) GetChassisSerialOk() (*string, bool)`

GetChassisSerialOk returns a tuple with the ChassisSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisSerial

`func (o *NodeNetworkInterfaces) SetChassisSerial(v string)`

SetChassisSerial sets ChassisSerial field to given value.

### HasChassisSerial

`func (o *NodeNetworkInterfaces) HasChassisSerial() bool`

HasChassisSerial returns a boolean if a field has been set.

### SetChassisSerialNil

`func (o *NodeNetworkInterfaces) SetChassisSerialNil(b bool)`

 SetChassisSerialNil sets the value for ChassisSerial to be an explicit nil

### UnsetChassisSerial
`func (o *NodeNetworkInterfaces) UnsetChassisSerial()`

UnsetChassisSerial ensures that no value is present for ChassisSerial, not even an explicit nil
### GetInterfaces

`func (o *NodeNetworkInterfaces) GetInterfaces() []NetworkInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *NodeNetworkInterfaces) GetInterfacesOk() (*[]NetworkInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *NodeNetworkInterfaces) SetInterfaces(v []NetworkInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *NodeNetworkInterfaces) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### SetInterfacesNil

`func (o *NodeNetworkInterfaces) SetInterfacesNil(b bool)`

 SetInterfacesNil sets the value for Interfaces to be an explicit nil

### UnsetInterfaces
`func (o *NodeNetworkInterfaces) UnsetInterfaces()`

UnsetInterfaces ensures that no value is present for Interfaces, not even an explicit nil
### GetMessage

`func (o *NodeNetworkInterfaces) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NodeNetworkInterfaces) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NodeNetworkInterfaces) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *NodeNetworkInterfaces) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *NodeNetworkInterfaces) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *NodeNetworkInterfaces) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetNodeId

`func (o *NodeNetworkInterfaces) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *NodeNetworkInterfaces) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *NodeNetworkInterfaces) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *NodeNetworkInterfaces) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### SetNodeIdNil

`func (o *NodeNetworkInterfaces) SetNodeIdNil(b bool)`

 SetNodeIdNil sets the value for NodeId to be an explicit nil

### UnsetNodeId
`func (o *NodeNetworkInterfaces) UnsetNodeId()`

UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
### GetSlot

`func (o *NodeNetworkInterfaces) GetSlot() int64`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *NodeNetworkInterfaces) GetSlotOk() (*int64, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *NodeNetworkInterfaces) SetSlot(v int64)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *NodeNetworkInterfaces) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### SetSlotNil

`func (o *NodeNetworkInterfaces) SetSlotNil(b bool)`

 SetSlotNil sets the value for Slot to be an explicit nil

### UnsetSlot
`func (o *NodeNetworkInterfaces) UnsetSlot()`

UnsetSlot ensures that no value is present for Slot, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


