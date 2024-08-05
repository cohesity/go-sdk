# GetGraphNodeDetailsRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeAttributes** | Pointer to **NullableBool** | If set to false the response will only return name, type and is_root fields filled in each node/relation. If set to true all the attributes for the nodes and its relations are also returned. Defaults to true. | [optional] 
**QueryRelations** | Pointer to **NullableBool** | If set to false only the node info will be returned, else the relations matching the below relation filters will be returned. Defaults to false. | [optional] 
**RelationFilter** | Pointer to [**NullableGraphNodeRelationFilterParams**](GraphNodeRelationFilterParams.md) | Specifies the filter params for graph node relations. | [optional] 

## Methods

### NewGetGraphNodeDetailsRequestParams

`func NewGetGraphNodeDetailsRequestParams() *GetGraphNodeDetailsRequestParams`

NewGetGraphNodeDetailsRequestParams instantiates a new GetGraphNodeDetailsRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGraphNodeDetailsRequestParamsWithDefaults

`func NewGetGraphNodeDetailsRequestParamsWithDefaults() *GetGraphNodeDetailsRequestParams`

NewGetGraphNodeDetailsRequestParamsWithDefaults instantiates a new GetGraphNodeDetailsRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeAttributes

`func (o *GetGraphNodeDetailsRequestParams) GetIncludeAttributes() bool`

GetIncludeAttributes returns the IncludeAttributes field if non-nil, zero value otherwise.

### GetIncludeAttributesOk

`func (o *GetGraphNodeDetailsRequestParams) GetIncludeAttributesOk() (*bool, bool)`

GetIncludeAttributesOk returns a tuple with the IncludeAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttributes

`func (o *GetGraphNodeDetailsRequestParams) SetIncludeAttributes(v bool)`

SetIncludeAttributes sets IncludeAttributes field to given value.

### HasIncludeAttributes

`func (o *GetGraphNodeDetailsRequestParams) HasIncludeAttributes() bool`

HasIncludeAttributes returns a boolean if a field has been set.

### SetIncludeAttributesNil

`func (o *GetGraphNodeDetailsRequestParams) SetIncludeAttributesNil(b bool)`

 SetIncludeAttributesNil sets the value for IncludeAttributes to be an explicit nil

### UnsetIncludeAttributes
`func (o *GetGraphNodeDetailsRequestParams) UnsetIncludeAttributes()`

UnsetIncludeAttributes ensures that no value is present for IncludeAttributes, not even an explicit nil
### GetQueryRelations

`func (o *GetGraphNodeDetailsRequestParams) GetQueryRelations() bool`

GetQueryRelations returns the QueryRelations field if non-nil, zero value otherwise.

### GetQueryRelationsOk

`func (o *GetGraphNodeDetailsRequestParams) GetQueryRelationsOk() (*bool, bool)`

GetQueryRelationsOk returns a tuple with the QueryRelations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryRelations

`func (o *GetGraphNodeDetailsRequestParams) SetQueryRelations(v bool)`

SetQueryRelations sets QueryRelations field to given value.

### HasQueryRelations

`func (o *GetGraphNodeDetailsRequestParams) HasQueryRelations() bool`

HasQueryRelations returns a boolean if a field has been set.

### SetQueryRelationsNil

`func (o *GetGraphNodeDetailsRequestParams) SetQueryRelationsNil(b bool)`

 SetQueryRelationsNil sets the value for QueryRelations to be an explicit nil

### UnsetQueryRelations
`func (o *GetGraphNodeDetailsRequestParams) UnsetQueryRelations()`

UnsetQueryRelations ensures that no value is present for QueryRelations, not even an explicit nil
### GetRelationFilter

`func (o *GetGraphNodeDetailsRequestParams) GetRelationFilter() GraphNodeRelationFilterParams`

GetRelationFilter returns the RelationFilter field if non-nil, zero value otherwise.

### GetRelationFilterOk

`func (o *GetGraphNodeDetailsRequestParams) GetRelationFilterOk() (*GraphNodeRelationFilterParams, bool)`

GetRelationFilterOk returns a tuple with the RelationFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationFilter

`func (o *GetGraphNodeDetailsRequestParams) SetRelationFilter(v GraphNodeRelationFilterParams)`

SetRelationFilter sets RelationFilter field to given value.

### HasRelationFilter

`func (o *GetGraphNodeDetailsRequestParams) HasRelationFilter() bool`

HasRelationFilter returns a boolean if a field has been set.

### SetRelationFilterNil

`func (o *GetGraphNodeDetailsRequestParams) SetRelationFilterNil(b bool)`

 SetRelationFilterNil sets the value for RelationFilter to be an explicit nil

### UnsetRelationFilter
`func (o *GetGraphNodeDetailsRequestParams) UnsetRelationFilter()`

UnsetRelationFilter ensures that no value is present for RelationFilter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


