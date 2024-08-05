# GraphNodeRelation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GraphNodeInfo** | Pointer to [**GraphNode**](GraphNode.md) | Specifies the graph node info for which edges are provided. | [optional] 
**Relations** | Pointer to [**[]GraphEdge**](GraphEdge.md) | Specifies the list of related edges/neighbours/relations of the source graph node. | [optional] 

## Methods

### NewGraphNodeRelation

`func NewGraphNodeRelation() *GraphNodeRelation`

NewGraphNodeRelation instantiates a new GraphNodeRelation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphNodeRelationWithDefaults

`func NewGraphNodeRelationWithDefaults() *GraphNodeRelation`

NewGraphNodeRelationWithDefaults instantiates a new GraphNodeRelation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGraphNodeInfo

`func (o *GraphNodeRelation) GetGraphNodeInfo() GraphNode`

GetGraphNodeInfo returns the GraphNodeInfo field if non-nil, zero value otherwise.

### GetGraphNodeInfoOk

`func (o *GraphNodeRelation) GetGraphNodeInfoOk() (*GraphNode, bool)`

GetGraphNodeInfoOk returns a tuple with the GraphNodeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphNodeInfo

`func (o *GraphNodeRelation) SetGraphNodeInfo(v GraphNode)`

SetGraphNodeInfo sets GraphNodeInfo field to given value.

### HasGraphNodeInfo

`func (o *GraphNodeRelation) HasGraphNodeInfo() bool`

HasGraphNodeInfo returns a boolean if a field has been set.

### GetRelations

`func (o *GraphNodeRelation) GetRelations() []GraphEdge`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *GraphNodeRelation) GetRelationsOk() (*[]GraphEdge, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *GraphNodeRelation) SetRelations(v []GraphEdge)`

SetRelations sets Relations field to given value.

### HasRelations

`func (o *GraphNodeRelation) HasRelations() bool`

HasRelations returns a boolean if a field has been set.

### SetRelationsNil

`func (o *GraphNodeRelation) SetRelationsNil(b bool)`

 SetRelationsNil sets the value for Relations to be an explicit nil

### UnsetRelations
`func (o *GraphNodeRelation) UnsetRelations()`

UnsetRelations ensures that no value is present for Relations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


