# GetGraphNodeRelationsDiffParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiffRelation** | Pointer to **NullableBool** | If set to false only the diff of node info will be returned else the diff of relations matching the below edge filters will also be returned. Defaults to false. | [optional] 
**DiffTypes** | Pointer to **[]string** | Specifies an optional mask to filter only certain kinds of diffs. Supported diff types - Added/Modified/Deleted/Unmodified | [optional] 
**RelationFilter** | Pointer to [**NullableGraphNodeRelationFilterParams**](GraphNodeRelationFilterParams.md) | Specifies the edges filter params for a graph node. | [optional] 

## Methods

### NewGetGraphNodeRelationsDiffParams

`func NewGetGraphNodeRelationsDiffParams() *GetGraphNodeRelationsDiffParams`

NewGetGraphNodeRelationsDiffParams instantiates a new GetGraphNodeRelationsDiffParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGraphNodeRelationsDiffParamsWithDefaults

`func NewGetGraphNodeRelationsDiffParamsWithDefaults() *GetGraphNodeRelationsDiffParams`

NewGetGraphNodeRelationsDiffParamsWithDefaults instantiates a new GetGraphNodeRelationsDiffParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiffRelation

`func (o *GetGraphNodeRelationsDiffParams) GetDiffRelation() bool`

GetDiffRelation returns the DiffRelation field if non-nil, zero value otherwise.

### GetDiffRelationOk

`func (o *GetGraphNodeRelationsDiffParams) GetDiffRelationOk() (*bool, bool)`

GetDiffRelationOk returns a tuple with the DiffRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffRelation

`func (o *GetGraphNodeRelationsDiffParams) SetDiffRelation(v bool)`

SetDiffRelation sets DiffRelation field to given value.

### HasDiffRelation

`func (o *GetGraphNodeRelationsDiffParams) HasDiffRelation() bool`

HasDiffRelation returns a boolean if a field has been set.

### SetDiffRelationNil

`func (o *GetGraphNodeRelationsDiffParams) SetDiffRelationNil(b bool)`

 SetDiffRelationNil sets the value for DiffRelation to be an explicit nil

### UnsetDiffRelation
`func (o *GetGraphNodeRelationsDiffParams) UnsetDiffRelation()`

UnsetDiffRelation ensures that no value is present for DiffRelation, not even an explicit nil
### GetDiffTypes

`func (o *GetGraphNodeRelationsDiffParams) GetDiffTypes() []string`

GetDiffTypes returns the DiffTypes field if non-nil, zero value otherwise.

### GetDiffTypesOk

`func (o *GetGraphNodeRelationsDiffParams) GetDiffTypesOk() (*[]string, bool)`

GetDiffTypesOk returns a tuple with the DiffTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffTypes

`func (o *GetGraphNodeRelationsDiffParams) SetDiffTypes(v []string)`

SetDiffTypes sets DiffTypes field to given value.

### HasDiffTypes

`func (o *GetGraphNodeRelationsDiffParams) HasDiffTypes() bool`

HasDiffTypes returns a boolean if a field has been set.

### SetDiffTypesNil

`func (o *GetGraphNodeRelationsDiffParams) SetDiffTypesNil(b bool)`

 SetDiffTypesNil sets the value for DiffTypes to be an explicit nil

### UnsetDiffTypes
`func (o *GetGraphNodeRelationsDiffParams) UnsetDiffTypes()`

UnsetDiffTypes ensures that no value is present for DiffTypes, not even an explicit nil
### GetRelationFilter

`func (o *GetGraphNodeRelationsDiffParams) GetRelationFilter() GraphNodeRelationFilterParams`

GetRelationFilter returns the RelationFilter field if non-nil, zero value otherwise.

### GetRelationFilterOk

`func (o *GetGraphNodeRelationsDiffParams) GetRelationFilterOk() (*GraphNodeRelationFilterParams, bool)`

GetRelationFilterOk returns a tuple with the RelationFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationFilter

`func (o *GetGraphNodeRelationsDiffParams) SetRelationFilter(v GraphNodeRelationFilterParams)`

SetRelationFilter sets RelationFilter field to given value.

### HasRelationFilter

`func (o *GetGraphNodeRelationsDiffParams) HasRelationFilter() bool`

HasRelationFilter returns a boolean if a field has been set.

### SetRelationFilterNil

`func (o *GetGraphNodeRelationsDiffParams) SetRelationFilterNil(b bool)`

 SetRelationFilterNil sets the value for RelationFilter to be an explicit nil

### UnsetRelationFilter
`func (o *GetGraphNodeRelationsDiffParams) UnsetRelationFilter()`

UnsetRelationFilter ensures that no value is present for RelationFilter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


