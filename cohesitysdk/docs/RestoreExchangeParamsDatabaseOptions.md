# RestoreExchangeParamsDatabaseOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **NullableInt64** | The windows machine to which the database will be restored. | [optional] 

## Methods

### NewRestoreExchangeParamsDatabaseOptions

`func NewRestoreExchangeParamsDatabaseOptions() *RestoreExchangeParamsDatabaseOptions`

NewRestoreExchangeParamsDatabaseOptions instantiates a new RestoreExchangeParamsDatabaseOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreExchangeParamsDatabaseOptionsWithDefaults

`func NewRestoreExchangeParamsDatabaseOptionsWithDefaults() *RestoreExchangeParamsDatabaseOptions`

NewRestoreExchangeParamsDatabaseOptionsWithDefaults instantiates a new RestoreExchangeParamsDatabaseOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *RestoreExchangeParamsDatabaseOptions) GetEntityId() int64`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *RestoreExchangeParamsDatabaseOptions) GetEntityIdOk() (*int64, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *RestoreExchangeParamsDatabaseOptions) SetEntityId(v int64)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *RestoreExchangeParamsDatabaseOptions) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *RestoreExchangeParamsDatabaseOptions) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *RestoreExchangeParamsDatabaseOptions) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


