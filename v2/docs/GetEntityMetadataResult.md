# GetEntityMetadataResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityList** | Pointer to [**[]EntityMetadataParams**](EntityMetadataParams.md) | Specifies the list of entities with thier metadata. | [optional] 
**SourceId** | Pointer to **int64** | Specifies the source id of the entities whose metadata is being requested. | [optional] 

## Methods

### NewGetEntityMetadataResult

`func NewGetEntityMetadataResult() *GetEntityMetadataResult`

NewGetEntityMetadataResult instantiates a new GetEntityMetadataResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEntityMetadataResultWithDefaults

`func NewGetEntityMetadataResultWithDefaults() *GetEntityMetadataResult`

NewGetEntityMetadataResultWithDefaults instantiates a new GetEntityMetadataResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityList

`func (o *GetEntityMetadataResult) GetEntityList() []EntityMetadataParams`

GetEntityList returns the EntityList field if non-nil, zero value otherwise.

### GetEntityListOk

`func (o *GetEntityMetadataResult) GetEntityListOk() (*[]EntityMetadataParams, bool)`

GetEntityListOk returns a tuple with the EntityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityList

`func (o *GetEntityMetadataResult) SetEntityList(v []EntityMetadataParams)`

SetEntityList sets EntityList field to given value.

### HasEntityList

`func (o *GetEntityMetadataResult) HasEntityList() bool`

HasEntityList returns a boolean if a field has been set.

### SetEntityListNil

`func (o *GetEntityMetadataResult) SetEntityListNil(b bool)`

 SetEntityListNil sets the value for EntityList to be an explicit nil

### UnsetEntityList
`func (o *GetEntityMetadataResult) UnsetEntityList()`

UnsetEntityList ensures that no value is present for EntityList, not even an explicit nil
### GetSourceId

`func (o *GetEntityMetadataResult) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *GetEntityMetadataResult) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *GetEntityMetadataResult) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *GetEntityMetadataResult) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


