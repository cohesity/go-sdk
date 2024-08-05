# AssociateEntityMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityList** | [**[]EntityMetadataParams**](EntityMetadataParams.md) | Specifies a list of entity and associated metadata mappings. | 
**SourceId** | **int64** | Specifies the source id of the entities vector whose metadata is being updated. | 

## Methods

### NewAssociateEntityMetadataRequest

`func NewAssociateEntityMetadataRequest(entityList []EntityMetadataParams, sourceId int64, ) *AssociateEntityMetadataRequest`

NewAssociateEntityMetadataRequest instantiates a new AssociateEntityMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociateEntityMetadataRequestWithDefaults

`func NewAssociateEntityMetadataRequestWithDefaults() *AssociateEntityMetadataRequest`

NewAssociateEntityMetadataRequestWithDefaults instantiates a new AssociateEntityMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityList

`func (o *AssociateEntityMetadataRequest) GetEntityList() []EntityMetadataParams`

GetEntityList returns the EntityList field if non-nil, zero value otherwise.

### GetEntityListOk

`func (o *AssociateEntityMetadataRequest) GetEntityListOk() (*[]EntityMetadataParams, bool)`

GetEntityListOk returns a tuple with the EntityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityList

`func (o *AssociateEntityMetadataRequest) SetEntityList(v []EntityMetadataParams)`

SetEntityList sets EntityList field to given value.


### GetSourceId

`func (o *AssociateEntityMetadataRequest) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *AssociateEntityMetadataRequest) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *AssociateEntityMetadataRequest) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


