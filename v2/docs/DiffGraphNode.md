# DiffGraphNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseGraphNode** | Pointer to [**GraphNode**](GraphNode.md) | Specifies the information about the node from base snapshot. | [optional] 
**CurrentGraphNode** | Pointer to [**GraphNode**](GraphNode.md) | Specifies the information about the graph node from current snapshot. If the same node is deleted in live view, this could be empty. | [optional] 
**DiffType** | Pointer to **NullableString** | Specifies the diff type for the base node. | [optional] 

## Methods

### NewDiffGraphNode

`func NewDiffGraphNode() *DiffGraphNode`

NewDiffGraphNode instantiates a new DiffGraphNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiffGraphNodeWithDefaults

`func NewDiffGraphNodeWithDefaults() *DiffGraphNode`

NewDiffGraphNodeWithDefaults instantiates a new DiffGraphNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseGraphNode

`func (o *DiffGraphNode) GetBaseGraphNode() GraphNode`

GetBaseGraphNode returns the BaseGraphNode field if non-nil, zero value otherwise.

### GetBaseGraphNodeOk

`func (o *DiffGraphNode) GetBaseGraphNodeOk() (*GraphNode, bool)`

GetBaseGraphNodeOk returns a tuple with the BaseGraphNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseGraphNode

`func (o *DiffGraphNode) SetBaseGraphNode(v GraphNode)`

SetBaseGraphNode sets BaseGraphNode field to given value.

### HasBaseGraphNode

`func (o *DiffGraphNode) HasBaseGraphNode() bool`

HasBaseGraphNode returns a boolean if a field has been set.

### GetCurrentGraphNode

`func (o *DiffGraphNode) GetCurrentGraphNode() GraphNode`

GetCurrentGraphNode returns the CurrentGraphNode field if non-nil, zero value otherwise.

### GetCurrentGraphNodeOk

`func (o *DiffGraphNode) GetCurrentGraphNodeOk() (*GraphNode, bool)`

GetCurrentGraphNodeOk returns a tuple with the CurrentGraphNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentGraphNode

`func (o *DiffGraphNode) SetCurrentGraphNode(v GraphNode)`

SetCurrentGraphNode sets CurrentGraphNode field to given value.

### HasCurrentGraphNode

`func (o *DiffGraphNode) HasCurrentGraphNode() bool`

HasCurrentGraphNode returns a boolean if a field has been set.

### GetDiffType

`func (o *DiffGraphNode) GetDiffType() string`

GetDiffType returns the DiffType field if non-nil, zero value otherwise.

### GetDiffTypeOk

`func (o *DiffGraphNode) GetDiffTypeOk() (*string, bool)`

GetDiffTypeOk returns a tuple with the DiffType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffType

`func (o *DiffGraphNode) SetDiffType(v string)`

SetDiffType sets DiffType field to given value.

### HasDiffType

`func (o *DiffGraphNode) HasDiffType() bool`

HasDiffType returns a boolean if a field has been set.

### SetDiffTypeNil

`func (o *DiffGraphNode) SetDiffTypeNil(b bool)`

 SetDiffTypeNil sets the value for DiffType to be an explicit nil

### UnsetDiffType
`func (o *DiffGraphNode) UnsetDiffType()`

UnsetDiffType ensures that no value is present for DiffType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


