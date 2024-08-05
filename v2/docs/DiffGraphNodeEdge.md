# DiffGraphNodeEdge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseRelation** | Pointer to [**GraphEdge**](GraphEdge.md) | This will be present if either the pair of edge relation and destination node is modified or deleted. If diff type is unmodified this field will not be present. | [optional] 
**CurrentRelation** | Pointer to [**GraphEdge**](GraphEdge.md) | This will be present if either the pair of edge relation and destination node is added, modified or unmodified. | [optional] 
**DiffType** | Pointer to **NullableString** | Specifies the diff type for the base node. | [optional] 

## Methods

### NewDiffGraphNodeEdge

`func NewDiffGraphNodeEdge() *DiffGraphNodeEdge`

NewDiffGraphNodeEdge instantiates a new DiffGraphNodeEdge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiffGraphNodeEdgeWithDefaults

`func NewDiffGraphNodeEdgeWithDefaults() *DiffGraphNodeEdge`

NewDiffGraphNodeEdgeWithDefaults instantiates a new DiffGraphNodeEdge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseRelation

`func (o *DiffGraphNodeEdge) GetBaseRelation() GraphEdge`

GetBaseRelation returns the BaseRelation field if non-nil, zero value otherwise.

### GetBaseRelationOk

`func (o *DiffGraphNodeEdge) GetBaseRelationOk() (*GraphEdge, bool)`

GetBaseRelationOk returns a tuple with the BaseRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseRelation

`func (o *DiffGraphNodeEdge) SetBaseRelation(v GraphEdge)`

SetBaseRelation sets BaseRelation field to given value.

### HasBaseRelation

`func (o *DiffGraphNodeEdge) HasBaseRelation() bool`

HasBaseRelation returns a boolean if a field has been set.

### GetCurrentRelation

`func (o *DiffGraphNodeEdge) GetCurrentRelation() GraphEdge`

GetCurrentRelation returns the CurrentRelation field if non-nil, zero value otherwise.

### GetCurrentRelationOk

`func (o *DiffGraphNodeEdge) GetCurrentRelationOk() (*GraphEdge, bool)`

GetCurrentRelationOk returns a tuple with the CurrentRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRelation

`func (o *DiffGraphNodeEdge) SetCurrentRelation(v GraphEdge)`

SetCurrentRelation sets CurrentRelation field to given value.

### HasCurrentRelation

`func (o *DiffGraphNodeEdge) HasCurrentRelation() bool`

HasCurrentRelation returns a boolean if a field has been set.

### GetDiffType

`func (o *DiffGraphNodeEdge) GetDiffType() string`

GetDiffType returns the DiffType field if non-nil, zero value otherwise.

### GetDiffTypeOk

`func (o *DiffGraphNodeEdge) GetDiffTypeOk() (*string, bool)`

GetDiffTypeOk returns a tuple with the DiffType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffType

`func (o *DiffGraphNodeEdge) SetDiffType(v string)`

SetDiffType sets DiffType field to given value.

### HasDiffType

`func (o *DiffGraphNodeEdge) HasDiffType() bool`

HasDiffType returns a boolean if a field has been set.

### SetDiffTypeNil

`func (o *DiffGraphNodeEdge) SetDiffTypeNil(b bool)`

 SetDiffTypeNil sets the value for DiffType to be an explicit nil

### UnsetDiffType
`func (o *DiffGraphNodeEdge) UnsetDiffType()`

UnsetDiffType ensures that no value is present for DiffType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


