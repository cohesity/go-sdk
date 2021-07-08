# VirtualNodeConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | Pointer to **NullableInt64** | Specifies the Node ID for this node. | [optional] 
**NodeIp** | Pointer to **NullableString** | Specifies the Node IP address for this node. | [optional] 
**UseAsComputeNode** | Pointer to **NullableBool** | Specifies whether to use the Node for compute only. | [optional] 

## Methods

### NewVirtualNodeConfiguration

`func NewVirtualNodeConfiguration() *VirtualNodeConfiguration`

NewVirtualNodeConfiguration instantiates a new VirtualNodeConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualNodeConfigurationWithDefaults

`func NewVirtualNodeConfigurationWithDefaults() *VirtualNodeConfiguration`

NewVirtualNodeConfigurationWithDefaults instantiates a new VirtualNodeConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *VirtualNodeConfiguration) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *VirtualNodeConfiguration) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *VirtualNodeConfiguration) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *VirtualNodeConfiguration) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### SetNodeIdNil

`func (o *VirtualNodeConfiguration) SetNodeIdNil(b bool)`

 SetNodeIdNil sets the value for NodeId to be an explicit nil

### UnsetNodeId
`func (o *VirtualNodeConfiguration) UnsetNodeId()`

UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
### GetNodeIp

`func (o *VirtualNodeConfiguration) GetNodeIp() string`

GetNodeIp returns the NodeIp field if non-nil, zero value otherwise.

### GetNodeIpOk

`func (o *VirtualNodeConfiguration) GetNodeIpOk() (*string, bool)`

GetNodeIpOk returns a tuple with the NodeIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIp

`func (o *VirtualNodeConfiguration) SetNodeIp(v string)`

SetNodeIp sets NodeIp field to given value.

### HasNodeIp

`func (o *VirtualNodeConfiguration) HasNodeIp() bool`

HasNodeIp returns a boolean if a field has been set.

### SetNodeIpNil

`func (o *VirtualNodeConfiguration) SetNodeIpNil(b bool)`

 SetNodeIpNil sets the value for NodeIp to be an explicit nil

### UnsetNodeIp
`func (o *VirtualNodeConfiguration) UnsetNodeIp()`

UnsetNodeIp ensures that no value is present for NodeIp, not even an explicit nil
### GetUseAsComputeNode

`func (o *VirtualNodeConfiguration) GetUseAsComputeNode() bool`

GetUseAsComputeNode returns the UseAsComputeNode field if non-nil, zero value otherwise.

### GetUseAsComputeNodeOk

`func (o *VirtualNodeConfiguration) GetUseAsComputeNodeOk() (*bool, bool)`

GetUseAsComputeNodeOk returns a tuple with the UseAsComputeNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAsComputeNode

`func (o *VirtualNodeConfiguration) SetUseAsComputeNode(v bool)`

SetUseAsComputeNode sets UseAsComputeNode field to given value.

### HasUseAsComputeNode

`func (o *VirtualNodeConfiguration) HasUseAsComputeNode() bool`

HasUseAsComputeNode returns a boolean if a field has been set.

### SetUseAsComputeNodeNil

`func (o *VirtualNodeConfiguration) SetUseAsComputeNodeNil(b bool)`

 SetUseAsComputeNodeNil sets the value for UseAsComputeNode to be an explicit nil

### UnsetUseAsComputeNode
`func (o *VirtualNodeConfiguration) UnsetUseAsComputeNode()`

UnsetUseAsComputeNode ensures that no value is present for UseAsComputeNode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


