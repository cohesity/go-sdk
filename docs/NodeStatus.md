# NodeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessage** | Pointer to **NullableString** | Specifies an optional message relating to the node status. | [optional] 
**IpmiIp** | Pointer to **NullableString** | Specifies the IPMI IP of the node (if physical cluster). | [optional] 
**NodeId** | Pointer to **NullableInt64** | Specifies the ID of the node. | [optional] 
**NodeIp** | Pointer to **NullableString** | For physical nodes this will specify the IP address of the node. | [optional] 

## Methods

### NewNodeStatus

`func NewNodeStatus() *NodeStatus`

NewNodeStatus instantiates a new NodeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeStatusWithDefaults

`func NewNodeStatusWithDefaults() *NodeStatus`

NewNodeStatusWithDefaults instantiates a new NodeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessage

`func (o *NodeStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *NodeStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *NodeStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *NodeStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *NodeStatus) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *NodeStatus) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetIpmiIp

`func (o *NodeStatus) GetIpmiIp() string`

GetIpmiIp returns the IpmiIp field if non-nil, zero value otherwise.

### GetIpmiIpOk

`func (o *NodeStatus) GetIpmiIpOk() (*string, bool)`

GetIpmiIpOk returns a tuple with the IpmiIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpmiIp

`func (o *NodeStatus) SetIpmiIp(v string)`

SetIpmiIp sets IpmiIp field to given value.

### HasIpmiIp

`func (o *NodeStatus) HasIpmiIp() bool`

HasIpmiIp returns a boolean if a field has been set.

### SetIpmiIpNil

`func (o *NodeStatus) SetIpmiIpNil(b bool)`

 SetIpmiIpNil sets the value for IpmiIp to be an explicit nil

### UnsetIpmiIp
`func (o *NodeStatus) UnsetIpmiIp()`

UnsetIpmiIp ensures that no value is present for IpmiIp, not even an explicit nil
### GetNodeId

`func (o *NodeStatus) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *NodeStatus) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *NodeStatus) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *NodeStatus) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### SetNodeIdNil

`func (o *NodeStatus) SetNodeIdNil(b bool)`

 SetNodeIdNil sets the value for NodeId to be an explicit nil

### UnsetNodeId
`func (o *NodeStatus) UnsetNodeId()`

UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
### GetNodeIp

`func (o *NodeStatus) GetNodeIp() string`

GetNodeIp returns the NodeIp field if non-nil, zero value otherwise.

### GetNodeIpOk

`func (o *NodeStatus) GetNodeIpOk() (*string, bool)`

GetNodeIpOk returns a tuple with the NodeIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIp

`func (o *NodeStatus) SetNodeIp(v string)`

SetNodeIp sets NodeIp field to given value.

### HasNodeIp

`func (o *NodeStatus) HasNodeIp() bool`

HasNodeIp returns a boolean if a field has been set.

### SetNodeIpNil

`func (o *NodeStatus) SetNodeIpNil(b bool)`

 SetNodeIpNil sets the value for NodeIp to be an explicit nil

### UnsetNodeIp
`func (o *NodeStatus) UnsetNodeIp()`

UnsetNodeIp ensures that no value is present for NodeIp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


