# NodeFreeDisks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChassisSerial** | Pointer to **NullableString** | Chassis serial number. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | Error message of disks assimilation request. | [optional] [readonly] 
**FreeDisks** | [**[]FreeDisk**](FreeDisk.md) | Specifies list of free disks of node. | 
**NodeId** | **NullableInt64** | Specifies the id of a node. | 
**Slot** | Pointer to **NullableInt64** | Slot number of node | [optional] 

## Methods

### NewNodeFreeDisks

`func NewNodeFreeDisks(freeDisks []FreeDisk, nodeId NullableInt64, ) *NodeFreeDisks`

NewNodeFreeDisks instantiates a new NodeFreeDisks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeFreeDisksWithDefaults

`func NewNodeFreeDisksWithDefaults() *NodeFreeDisks`

NewNodeFreeDisksWithDefaults instantiates a new NodeFreeDisks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChassisSerial

`func (o *NodeFreeDisks) GetChassisSerial() string`

GetChassisSerial returns the ChassisSerial field if non-nil, zero value otherwise.

### GetChassisSerialOk

`func (o *NodeFreeDisks) GetChassisSerialOk() (*string, bool)`

GetChassisSerialOk returns a tuple with the ChassisSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisSerial

`func (o *NodeFreeDisks) SetChassisSerial(v string)`

SetChassisSerial sets ChassisSerial field to given value.

### HasChassisSerial

`func (o *NodeFreeDisks) HasChassisSerial() bool`

HasChassisSerial returns a boolean if a field has been set.

### SetChassisSerialNil

`func (o *NodeFreeDisks) SetChassisSerialNil(b bool)`

 SetChassisSerialNil sets the value for ChassisSerial to be an explicit nil

### UnsetChassisSerial
`func (o *NodeFreeDisks) UnsetChassisSerial()`

UnsetChassisSerial ensures that no value is present for ChassisSerial, not even an explicit nil
### GetErrorMessage

`func (o *NodeFreeDisks) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *NodeFreeDisks) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *NodeFreeDisks) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *NodeFreeDisks) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *NodeFreeDisks) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *NodeFreeDisks) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetFreeDisks

`func (o *NodeFreeDisks) GetFreeDisks() []FreeDisk`

GetFreeDisks returns the FreeDisks field if non-nil, zero value otherwise.

### GetFreeDisksOk

`func (o *NodeFreeDisks) GetFreeDisksOk() (*[]FreeDisk, bool)`

GetFreeDisksOk returns a tuple with the FreeDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeDisks

`func (o *NodeFreeDisks) SetFreeDisks(v []FreeDisk)`

SetFreeDisks sets FreeDisks field to given value.


### SetFreeDisksNil

`func (o *NodeFreeDisks) SetFreeDisksNil(b bool)`

 SetFreeDisksNil sets the value for FreeDisks to be an explicit nil

### UnsetFreeDisks
`func (o *NodeFreeDisks) UnsetFreeDisks()`

UnsetFreeDisks ensures that no value is present for FreeDisks, not even an explicit nil
### GetNodeId

`func (o *NodeFreeDisks) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *NodeFreeDisks) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *NodeFreeDisks) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.


### SetNodeIdNil

`func (o *NodeFreeDisks) SetNodeIdNil(b bool)`

 SetNodeIdNil sets the value for NodeId to be an explicit nil

### UnsetNodeId
`func (o *NodeFreeDisks) UnsetNodeId()`

UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
### GetSlot

`func (o *NodeFreeDisks) GetSlot() int64`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *NodeFreeDisks) GetSlotOk() (*int64, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *NodeFreeDisks) SetSlot(v int64)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *NodeFreeDisks) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### SetSlotNil

`func (o *NodeFreeDisks) SetSlotNil(b bool)`

 SetSlotNil sets the value for Slot to be an explicit nil

### UnsetSlot
`func (o *NodeFreeDisks) UnsetSlot()`

UnsetSlot ensures that no value is present for Slot, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


