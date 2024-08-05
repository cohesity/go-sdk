# NodeBondInterfaceParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | **[]string** | Names of the secondary interfaces. | 
**NodeId** | Pointer to **NullableInt64** | Id of the node, this is required when node is part of a cluster i.e when nodeType is &#39;ClusterNode&#39;. | [optional] 
**NodeType** | **string** | Type of the node where the interface has to be created. &#39;ClusterNode&#39; indicates a node that is part of a cluster. &#39;FreeNode&#39; indicates a node that is not part of cluster. | 
**BondingMode** | **string** | Bonding mode of the interface. | 
**Name** | **string** | Name of the bond interface. | 

## Methods

### NewNodeBondInterfaceParams

`func NewNodeBondInterfaceParams(members []string, nodeType string, bondingMode string, name string, ) *NodeBondInterfaceParams`

NewNodeBondInterfaceParams instantiates a new NodeBondInterfaceParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeBondInterfaceParamsWithDefaults

`func NewNodeBondInterfaceParamsWithDefaults() *NodeBondInterfaceParams`

NewNodeBondInterfaceParamsWithDefaults instantiates a new NodeBondInterfaceParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *NodeBondInterfaceParams) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *NodeBondInterfaceParams) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *NodeBondInterfaceParams) SetMembers(v []string)`

SetMembers sets Members field to given value.


### GetNodeId

`func (o *NodeBondInterfaceParams) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *NodeBondInterfaceParams) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *NodeBondInterfaceParams) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *NodeBondInterfaceParams) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### SetNodeIdNil

`func (o *NodeBondInterfaceParams) SetNodeIdNil(b bool)`

 SetNodeIdNil sets the value for NodeId to be an explicit nil

### UnsetNodeId
`func (o *NodeBondInterfaceParams) UnsetNodeId()`

UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
### GetNodeType

`func (o *NodeBondInterfaceParams) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *NodeBondInterfaceParams) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *NodeBondInterfaceParams) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.


### GetBondingMode

`func (o *NodeBondInterfaceParams) GetBondingMode() string`

GetBondingMode returns the BondingMode field if non-nil, zero value otherwise.

### GetBondingModeOk

`func (o *NodeBondInterfaceParams) GetBondingModeOk() (*string, bool)`

GetBondingModeOk returns a tuple with the BondingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondingMode

`func (o *NodeBondInterfaceParams) SetBondingMode(v string)`

SetBondingMode sets BondingMode field to given value.


### GetName

`func (o *NodeBondInterfaceParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeBondInterfaceParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeBondInterfaceParams) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


