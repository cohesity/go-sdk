# NodeNetworkInterfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChassisSerialNumber** | Pointer to **NullableString** | Chassis serial number. | [optional] 
**Interfaces** | Pointer to [**[]Interface**](Interface.md) | List of interfaces on the node. | [optional] 
**NodeId** | Pointer to **NullableInt64** | Id of the node. | [optional] 
**Slot** | Pointer to **NullableInt64** | Slot number of the node. | [optional] 

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

### GetChassisSerialNumber

`func (o *NodeNetworkInterfaces) GetChassisSerialNumber() string`

GetChassisSerialNumber returns the ChassisSerialNumber field if non-nil, zero value otherwise.

### GetChassisSerialNumberOk

`func (o *NodeNetworkInterfaces) GetChassisSerialNumberOk() (*string, bool)`

GetChassisSerialNumberOk returns a tuple with the ChassisSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisSerialNumber

`func (o *NodeNetworkInterfaces) SetChassisSerialNumber(v string)`

SetChassisSerialNumber sets ChassisSerialNumber field to given value.

### HasChassisSerialNumber

`func (o *NodeNetworkInterfaces) HasChassisSerialNumber() bool`

HasChassisSerialNumber returns a boolean if a field has been set.

### SetChassisSerialNumberNil

`func (o *NodeNetworkInterfaces) SetChassisSerialNumberNil(b bool)`

 SetChassisSerialNumberNil sets the value for ChassisSerialNumber to be an explicit nil

### UnsetChassisSerialNumber
`func (o *NodeNetworkInterfaces) UnsetChassisSerialNumber()`

UnsetChassisSerialNumber ensures that no value is present for ChassisSerialNumber, not even an explicit nil
### GetInterfaces

`func (o *NodeNetworkInterfaces) GetInterfaces() []Interface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *NodeNetworkInterfaces) GetInterfacesOk() (*[]Interface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *NodeNetworkInterfaces) SetInterfaces(v []Interface)`

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


