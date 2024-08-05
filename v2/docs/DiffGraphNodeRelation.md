# DiffGraphNodeRelation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiffGraphNode** | Pointer to [**DiffGraphNode**](DiffGraphNode.md) | Specifies the difference in the graph node info. | [optional] 
**DiffRelations** | Pointer to [**[]DiffGraphNodeEdge**](DiffGraphNodeEdge.md) | Specifies the all pair of edges/node relations which are added, deleted or modified | [optional] 
**SrcNodeId** | Pointer to **string** | Specifies Unique ID of the source node. | [optional] [readonly] 
**UnmodifiedRelations** | Pointer to [**[]GraphEdge**](GraphEdge.md) | Specifies the list of all the edges which are unmodified. | [optional] 

## Methods

### NewDiffGraphNodeRelation

`func NewDiffGraphNodeRelation() *DiffGraphNodeRelation`

NewDiffGraphNodeRelation instantiates a new DiffGraphNodeRelation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiffGraphNodeRelationWithDefaults

`func NewDiffGraphNodeRelationWithDefaults() *DiffGraphNodeRelation`

NewDiffGraphNodeRelationWithDefaults instantiates a new DiffGraphNodeRelation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiffGraphNode

`func (o *DiffGraphNodeRelation) GetDiffGraphNode() DiffGraphNode`

GetDiffGraphNode returns the DiffGraphNode field if non-nil, zero value otherwise.

### GetDiffGraphNodeOk

`func (o *DiffGraphNodeRelation) GetDiffGraphNodeOk() (*DiffGraphNode, bool)`

GetDiffGraphNodeOk returns a tuple with the DiffGraphNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffGraphNode

`func (o *DiffGraphNodeRelation) SetDiffGraphNode(v DiffGraphNode)`

SetDiffGraphNode sets DiffGraphNode field to given value.

### HasDiffGraphNode

`func (o *DiffGraphNodeRelation) HasDiffGraphNode() bool`

HasDiffGraphNode returns a boolean if a field has been set.

### GetDiffRelations

`func (o *DiffGraphNodeRelation) GetDiffRelations() []DiffGraphNodeEdge`

GetDiffRelations returns the DiffRelations field if non-nil, zero value otherwise.

### GetDiffRelationsOk

`func (o *DiffGraphNodeRelation) GetDiffRelationsOk() (*[]DiffGraphNodeEdge, bool)`

GetDiffRelationsOk returns a tuple with the DiffRelations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffRelations

`func (o *DiffGraphNodeRelation) SetDiffRelations(v []DiffGraphNodeEdge)`

SetDiffRelations sets DiffRelations field to given value.

### HasDiffRelations

`func (o *DiffGraphNodeRelation) HasDiffRelations() bool`

HasDiffRelations returns a boolean if a field has been set.

### SetDiffRelationsNil

`func (o *DiffGraphNodeRelation) SetDiffRelationsNil(b bool)`

 SetDiffRelationsNil sets the value for DiffRelations to be an explicit nil

### UnsetDiffRelations
`func (o *DiffGraphNodeRelation) UnsetDiffRelations()`

UnsetDiffRelations ensures that no value is present for DiffRelations, not even an explicit nil
### GetSrcNodeId

`func (o *DiffGraphNodeRelation) GetSrcNodeId() string`

GetSrcNodeId returns the SrcNodeId field if non-nil, zero value otherwise.

### GetSrcNodeIdOk

`func (o *DiffGraphNodeRelation) GetSrcNodeIdOk() (*string, bool)`

GetSrcNodeIdOk returns a tuple with the SrcNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcNodeId

`func (o *DiffGraphNodeRelation) SetSrcNodeId(v string)`

SetSrcNodeId sets SrcNodeId field to given value.

### HasSrcNodeId

`func (o *DiffGraphNodeRelation) HasSrcNodeId() bool`

HasSrcNodeId returns a boolean if a field has been set.

### GetUnmodifiedRelations

`func (o *DiffGraphNodeRelation) GetUnmodifiedRelations() []GraphEdge`

GetUnmodifiedRelations returns the UnmodifiedRelations field if non-nil, zero value otherwise.

### GetUnmodifiedRelationsOk

`func (o *DiffGraphNodeRelation) GetUnmodifiedRelationsOk() (*[]GraphEdge, bool)`

GetUnmodifiedRelationsOk returns a tuple with the UnmodifiedRelations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmodifiedRelations

`func (o *DiffGraphNodeRelation) SetUnmodifiedRelations(v []GraphEdge)`

SetUnmodifiedRelations sets UnmodifiedRelations field to given value.

### HasUnmodifiedRelations

`func (o *DiffGraphNodeRelation) HasUnmodifiedRelations() bool`

HasUnmodifiedRelations returns a boolean if a field has been set.

### SetUnmodifiedRelationsNil

`func (o *DiffGraphNodeRelation) SetUnmodifiedRelationsNil(b bool)`

 SetUnmodifiedRelationsNil sets the value for UnmodifiedRelations to be an explicit nil

### UnsetUnmodifiedRelations
`func (o *DiffGraphNodeRelation) UnsetUnmodifiedRelations()`

UnsetUnmodifiedRelations ensures that no value is present for UnmodifiedRelations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


