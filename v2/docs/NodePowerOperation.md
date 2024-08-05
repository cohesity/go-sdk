# NodePowerOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | Pointer to **NullableInt64** | Id of the node to do the specified operation. | [optional] 
**Operation** | **NullableString** | The operation clould be poweroff, reboot. | 

## Methods

### NewNodePowerOperation

`func NewNodePowerOperation(operation NullableString, ) *NodePowerOperation`

NewNodePowerOperation instantiates a new NodePowerOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodePowerOperationWithDefaults

`func NewNodePowerOperationWithDefaults() *NodePowerOperation`

NewNodePowerOperationWithDefaults instantiates a new NodePowerOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *NodePowerOperation) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *NodePowerOperation) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *NodePowerOperation) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *NodePowerOperation) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### SetNodeIdNil

`func (o *NodePowerOperation) SetNodeIdNil(b bool)`

 SetNodeIdNil sets the value for NodeId to be an explicit nil

### UnsetNodeId
`func (o *NodePowerOperation) UnsetNodeId()`

UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
### GetOperation

`func (o *NodePowerOperation) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *NodePowerOperation) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *NodePowerOperation) SetOperation(v string)`

SetOperation sets Operation field to given value.


### SetOperationNil

`func (o *NodePowerOperation) SetOperationNil(b bool)`

 SetOperationNil sets the value for Operation to be an explicit nil

### UnsetOperation
`func (o *NodePowerOperation) UnsetOperation()`

UnsetOperation ensures that no value is present for Operation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


