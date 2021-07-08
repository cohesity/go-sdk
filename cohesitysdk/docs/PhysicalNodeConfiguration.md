# PhysicalNodeConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | Pointer to **NullableInt64** | Specifies the Node ID for this node. | [optional] 
**NodeIp** | Pointer to **NullableString** | Specifies the Node IP address for this node. | [optional] 
**NodeIpmiIp** | Pointer to **NullableString** | Specifies IPMI IP for this node. | [optional] 
**UseAsComputeNode** | Pointer to **NullableBool** | Specifies whether to use the Node for compute only. | [optional] 

## Methods

### NewPhysicalNodeConfiguration

`func NewPhysicalNodeConfiguration() *PhysicalNodeConfiguration`

NewPhysicalNodeConfiguration instantiates a new PhysicalNodeConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalNodeConfigurationWithDefaults

`func NewPhysicalNodeConfigurationWithDefaults() *PhysicalNodeConfiguration`

NewPhysicalNodeConfigurationWithDefaults instantiates a new PhysicalNodeConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *PhysicalNodeConfiguration) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *PhysicalNodeConfiguration) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *PhysicalNodeConfiguration) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *PhysicalNodeConfiguration) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### SetNodeIdNil

`func (o *PhysicalNodeConfiguration) SetNodeIdNil(b bool)`

 SetNodeIdNil sets the value for NodeId to be an explicit nil

### UnsetNodeId
`func (o *PhysicalNodeConfiguration) UnsetNodeId()`

UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
### GetNodeIp

`func (o *PhysicalNodeConfiguration) GetNodeIp() string`

GetNodeIp returns the NodeIp field if non-nil, zero value otherwise.

### GetNodeIpOk

`func (o *PhysicalNodeConfiguration) GetNodeIpOk() (*string, bool)`

GetNodeIpOk returns a tuple with the NodeIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIp

`func (o *PhysicalNodeConfiguration) SetNodeIp(v string)`

SetNodeIp sets NodeIp field to given value.

### HasNodeIp

`func (o *PhysicalNodeConfiguration) HasNodeIp() bool`

HasNodeIp returns a boolean if a field has been set.

### SetNodeIpNil

`func (o *PhysicalNodeConfiguration) SetNodeIpNil(b bool)`

 SetNodeIpNil sets the value for NodeIp to be an explicit nil

### UnsetNodeIp
`func (o *PhysicalNodeConfiguration) UnsetNodeIp()`

UnsetNodeIp ensures that no value is present for NodeIp, not even an explicit nil
### GetNodeIpmiIp

`func (o *PhysicalNodeConfiguration) GetNodeIpmiIp() string`

GetNodeIpmiIp returns the NodeIpmiIp field if non-nil, zero value otherwise.

### GetNodeIpmiIpOk

`func (o *PhysicalNodeConfiguration) GetNodeIpmiIpOk() (*string, bool)`

GetNodeIpmiIpOk returns a tuple with the NodeIpmiIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIpmiIp

`func (o *PhysicalNodeConfiguration) SetNodeIpmiIp(v string)`

SetNodeIpmiIp sets NodeIpmiIp field to given value.

### HasNodeIpmiIp

`func (o *PhysicalNodeConfiguration) HasNodeIpmiIp() bool`

HasNodeIpmiIp returns a boolean if a field has been set.

### SetNodeIpmiIpNil

`func (o *PhysicalNodeConfiguration) SetNodeIpmiIpNil(b bool)`

 SetNodeIpmiIpNil sets the value for NodeIpmiIp to be an explicit nil

### UnsetNodeIpmiIp
`func (o *PhysicalNodeConfiguration) UnsetNodeIpmiIp()`

UnsetNodeIpmiIp ensures that no value is present for NodeIpmiIp, not even an explicit nil
### GetUseAsComputeNode

`func (o *PhysicalNodeConfiguration) GetUseAsComputeNode() bool`

GetUseAsComputeNode returns the UseAsComputeNode field if non-nil, zero value otherwise.

### GetUseAsComputeNodeOk

`func (o *PhysicalNodeConfiguration) GetUseAsComputeNodeOk() (*bool, bool)`

GetUseAsComputeNodeOk returns a tuple with the UseAsComputeNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAsComputeNode

`func (o *PhysicalNodeConfiguration) SetUseAsComputeNode(v bool)`

SetUseAsComputeNode sets UseAsComputeNode field to given value.

### HasUseAsComputeNode

`func (o *PhysicalNodeConfiguration) HasUseAsComputeNode() bool`

HasUseAsComputeNode returns a boolean if a field has been set.

### SetUseAsComputeNodeNil

`func (o *PhysicalNodeConfiguration) SetUseAsComputeNodeNil(b bool)`

 SetUseAsComputeNodeNil sets the value for UseAsComputeNode to be an explicit nil

### UnsetUseAsComputeNode
`func (o *PhysicalNodeConfiguration) UnsetUseAsComputeNode()`

UnsetUseAsComputeNode ensures that no value is present for UseAsComputeNode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


