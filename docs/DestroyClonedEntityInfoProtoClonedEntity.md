# DestroyClonedEntityInfoProtoClonedEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**RelativeRestorePathVec** | Pointer to **[]string** | Path of all files created by the clone operation. Each path is relative to the clone view. | [optional] 

## Methods

### NewDestroyClonedEntityInfoProtoClonedEntity

`func NewDestroyClonedEntityInfoProtoClonedEntity() *DestroyClonedEntityInfoProtoClonedEntity`

NewDestroyClonedEntityInfoProtoClonedEntity instantiates a new DestroyClonedEntityInfoProtoClonedEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestroyClonedEntityInfoProtoClonedEntityWithDefaults

`func NewDestroyClonedEntityInfoProtoClonedEntityWithDefaults() *DestroyClonedEntityInfoProtoClonedEntity`

NewDestroyClonedEntityInfoProtoClonedEntityWithDefaults instantiates a new DestroyClonedEntityInfoProtoClonedEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *DestroyClonedEntityInfoProtoClonedEntity) GetEntity() EntityProto`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *DestroyClonedEntityInfoProtoClonedEntity) GetEntityOk() (*EntityProto, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *DestroyClonedEntityInfoProtoClonedEntity) SetEntity(v EntityProto)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *DestroyClonedEntityInfoProtoClonedEntity) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetRelativeRestorePathVec

`func (o *DestroyClonedEntityInfoProtoClonedEntity) GetRelativeRestorePathVec() []string`

GetRelativeRestorePathVec returns the RelativeRestorePathVec field if non-nil, zero value otherwise.

### GetRelativeRestorePathVecOk

`func (o *DestroyClonedEntityInfoProtoClonedEntity) GetRelativeRestorePathVecOk() (*[]string, bool)`

GetRelativeRestorePathVecOk returns a tuple with the RelativeRestorePathVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeRestorePathVec

`func (o *DestroyClonedEntityInfoProtoClonedEntity) SetRelativeRestorePathVec(v []string)`

SetRelativeRestorePathVec sets RelativeRestorePathVec field to given value.

### HasRelativeRestorePathVec

`func (o *DestroyClonedEntityInfoProtoClonedEntity) HasRelativeRestorePathVec() bool`

HasRelativeRestorePathVec returns a boolean if a field has been set.

### SetRelativeRestorePathVecNil

`func (o *DestroyClonedEntityInfoProtoClonedEntity) SetRelativeRestorePathVecNil(b bool)`

 SetRelativeRestorePathVecNil sets the value for RelativeRestorePathVec to be an explicit nil

### UnsetRelativeRestorePathVec
`func (o *DestroyClonedEntityInfoProtoClonedEntity) UnsetRelativeRestorePathVec()`

UnsetRelativeRestorePathVec ensures that no value is present for RelativeRestorePathVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


