# HeliosOnPremVMNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | Pointer to **NullableInt64** | Specifies the ID of the Node. | [optional] 
**NodeIp** | **NullableString** | Specifies the IP address of the Node. | 

## Methods

### NewHeliosOnPremVMNode

`func NewHeliosOnPremVMNode(nodeIp NullableString, ) *HeliosOnPremVMNode`

NewHeliosOnPremVMNode instantiates a new HeliosOnPremVMNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosOnPremVMNodeWithDefaults

`func NewHeliosOnPremVMNodeWithDefaults() *HeliosOnPremVMNode`

NewHeliosOnPremVMNodeWithDefaults instantiates a new HeliosOnPremVMNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *HeliosOnPremVMNode) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *HeliosOnPremVMNode) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *HeliosOnPremVMNode) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *HeliosOnPremVMNode) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### SetNodeIdNil

`func (o *HeliosOnPremVMNode) SetNodeIdNil(b bool)`

 SetNodeIdNil sets the value for NodeId to be an explicit nil

### UnsetNodeId
`func (o *HeliosOnPremVMNode) UnsetNodeId()`

UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
### GetNodeIp

`func (o *HeliosOnPremVMNode) GetNodeIp() string`

GetNodeIp returns the NodeIp field if non-nil, zero value otherwise.

### GetNodeIpOk

`func (o *HeliosOnPremVMNode) GetNodeIpOk() (*string, bool)`

GetNodeIpOk returns a tuple with the NodeIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIp

`func (o *HeliosOnPremVMNode) SetNodeIp(v string)`

SetNodeIp sets NodeIp field to given value.


### SetNodeIpNil

`func (o *HeliosOnPremVMNode) SetNodeIpNil(b bool)`

 SetNodeIpNil sets the value for NodeIp to be an explicit nil

### UnsetNodeIp
`func (o *HeliosOnPremVMNode) UnsetNodeIp()`

UnsetNodeIp ensures that no value is present for NodeIp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


