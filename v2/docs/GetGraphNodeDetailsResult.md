# GetGraphNodeDetailsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GraphNodeRelation** | Pointer to [**NullableGraphNodeRelation**](GraphNodeRelation.md) | Specifies fetched graph node and their edges information. | [optional] 

## Methods

### NewGetGraphNodeDetailsResult

`func NewGetGraphNodeDetailsResult() *GetGraphNodeDetailsResult`

NewGetGraphNodeDetailsResult instantiates a new GetGraphNodeDetailsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGraphNodeDetailsResultWithDefaults

`func NewGetGraphNodeDetailsResultWithDefaults() *GetGraphNodeDetailsResult`

NewGetGraphNodeDetailsResultWithDefaults instantiates a new GetGraphNodeDetailsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGraphNodeRelation

`func (o *GetGraphNodeDetailsResult) GetGraphNodeRelation() GraphNodeRelation`

GetGraphNodeRelation returns the GraphNodeRelation field if non-nil, zero value otherwise.

### GetGraphNodeRelationOk

`func (o *GetGraphNodeDetailsResult) GetGraphNodeRelationOk() (*GraphNodeRelation, bool)`

GetGraphNodeRelationOk returns a tuple with the GraphNodeRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphNodeRelation

`func (o *GetGraphNodeDetailsResult) SetGraphNodeRelation(v GraphNodeRelation)`

SetGraphNodeRelation sets GraphNodeRelation field to given value.

### HasGraphNodeRelation

`func (o *GetGraphNodeDetailsResult) HasGraphNodeRelation() bool`

HasGraphNodeRelation returns a boolean if a field has been set.

### SetGraphNodeRelationNil

`func (o *GetGraphNodeDetailsResult) SetGraphNodeRelationNil(b bool)`

 SetGraphNodeRelationNil sets the value for GraphNodeRelation to be an explicit nil

### UnsetGraphNodeRelation
`func (o *GetGraphNodeDetailsResult) UnsetGraphNodeRelation()`

UnsetGraphNodeRelation ensures that no value is present for GraphNodeRelation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


