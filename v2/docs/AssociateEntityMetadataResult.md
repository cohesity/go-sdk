# AssociateEntityMetadataResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityList** | Pointer to [**[]AssociateEntityMetadataResultParams**](AssociateEntityMetadataResultParams.md) | Specifies the list of entity ids mapped to errors (if any). | [optional] 
**SourceId** | Pointer to **int64** | Specifies the source id of the entities vector whose metadata is being updated. | [optional] 

## Methods

### NewAssociateEntityMetadataResult

`func NewAssociateEntityMetadataResult() *AssociateEntityMetadataResult`

NewAssociateEntityMetadataResult instantiates a new AssociateEntityMetadataResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociateEntityMetadataResultWithDefaults

`func NewAssociateEntityMetadataResultWithDefaults() *AssociateEntityMetadataResult`

NewAssociateEntityMetadataResultWithDefaults instantiates a new AssociateEntityMetadataResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityList

`func (o *AssociateEntityMetadataResult) GetEntityList() []AssociateEntityMetadataResultParams`

GetEntityList returns the EntityList field if non-nil, zero value otherwise.

### GetEntityListOk

`func (o *AssociateEntityMetadataResult) GetEntityListOk() (*[]AssociateEntityMetadataResultParams, bool)`

GetEntityListOk returns a tuple with the EntityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityList

`func (o *AssociateEntityMetadataResult) SetEntityList(v []AssociateEntityMetadataResultParams)`

SetEntityList sets EntityList field to given value.

### HasEntityList

`func (o *AssociateEntityMetadataResult) HasEntityList() bool`

HasEntityList returns a boolean if a field has been set.

### SetEntityListNil

`func (o *AssociateEntityMetadataResult) SetEntityListNil(b bool)`

 SetEntityListNil sets the value for EntityList to be an explicit nil

### UnsetEntityList
`func (o *AssociateEntityMetadataResult) UnsetEntityList()`

UnsetEntityList ensures that no value is present for EntityList, not even an explicit nil
### GetSourceId

`func (o *AssociateEntityMetadataResult) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *AssociateEntityMetadataResult) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *AssociateEntityMetadataResult) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *AssociateEntityMetadataResult) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


