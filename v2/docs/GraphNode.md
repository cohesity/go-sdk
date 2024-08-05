# GraphNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AadParams** | Pointer to [**NullableAadNodeInfo**](AadNodeInfo.md) | Specifies the aad specific node information. | [optional] 
**IsRootNode** | Pointer to **NullableBool** | Boolean to indicate if this is a root node or not. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the display name of the node. | [optional] 
**NodeId** | **string** | Specifies the unique id of the node. | 

## Methods

### NewGraphNode

`func NewGraphNode(nodeId string, ) *GraphNode`

NewGraphNode instantiates a new GraphNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphNodeWithDefaults

`func NewGraphNodeWithDefaults() *GraphNode`

NewGraphNodeWithDefaults instantiates a new GraphNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAadParams

`func (o *GraphNode) GetAadParams() AadNodeInfo`

GetAadParams returns the AadParams field if non-nil, zero value otherwise.

### GetAadParamsOk

`func (o *GraphNode) GetAadParamsOk() (*AadNodeInfo, bool)`

GetAadParamsOk returns a tuple with the AadParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAadParams

`func (o *GraphNode) SetAadParams(v AadNodeInfo)`

SetAadParams sets AadParams field to given value.

### HasAadParams

`func (o *GraphNode) HasAadParams() bool`

HasAadParams returns a boolean if a field has been set.

### SetAadParamsNil

`func (o *GraphNode) SetAadParamsNil(b bool)`

 SetAadParamsNil sets the value for AadParams to be an explicit nil

### UnsetAadParams
`func (o *GraphNode) UnsetAadParams()`

UnsetAadParams ensures that no value is present for AadParams, not even an explicit nil
### GetIsRootNode

`func (o *GraphNode) GetIsRootNode() bool`

GetIsRootNode returns the IsRootNode field if non-nil, zero value otherwise.

### GetIsRootNodeOk

`func (o *GraphNode) GetIsRootNodeOk() (*bool, bool)`

GetIsRootNodeOk returns a tuple with the IsRootNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRootNode

`func (o *GraphNode) SetIsRootNode(v bool)`

SetIsRootNode sets IsRootNode field to given value.

### HasIsRootNode

`func (o *GraphNode) HasIsRootNode() bool`

HasIsRootNode returns a boolean if a field has been set.

### SetIsRootNodeNil

`func (o *GraphNode) SetIsRootNodeNil(b bool)`

 SetIsRootNodeNil sets the value for IsRootNode to be an explicit nil

### UnsetIsRootNode
`func (o *GraphNode) UnsetIsRootNode()`

UnsetIsRootNode ensures that no value is present for IsRootNode, not even an explicit nil
### GetName

`func (o *GraphNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GraphNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GraphNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GraphNode) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GraphNode) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GraphNode) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNodeId

`func (o *GraphNode) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *GraphNode) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *GraphNode) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


