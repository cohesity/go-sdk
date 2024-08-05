# EntitySelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllEntities** | Pointer to **NullableBool** | If set to true all the entities ever known to this cluster will be exported. These include but are not limited to registered, unregistered, and replicated entities | [optional] 
**RegisteredEntities** | Pointer to **NullableBool** | If this is true, the entities that are regisered will be qualified. Else, all the entities will be matched. | [optional] 

## Methods

### NewEntitySelector

`func NewEntitySelector() *EntitySelector`

NewEntitySelector instantiates a new EntitySelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitySelectorWithDefaults

`func NewEntitySelectorWithDefaults() *EntitySelector`

NewEntitySelectorWithDefaults instantiates a new EntitySelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllEntities

`func (o *EntitySelector) GetAllEntities() bool`

GetAllEntities returns the AllEntities field if non-nil, zero value otherwise.

### GetAllEntitiesOk

`func (o *EntitySelector) GetAllEntitiesOk() (*bool, bool)`

GetAllEntitiesOk returns a tuple with the AllEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllEntities

`func (o *EntitySelector) SetAllEntities(v bool)`

SetAllEntities sets AllEntities field to given value.

### HasAllEntities

`func (o *EntitySelector) HasAllEntities() bool`

HasAllEntities returns a boolean if a field has been set.

### SetAllEntitiesNil

`func (o *EntitySelector) SetAllEntitiesNil(b bool)`

 SetAllEntitiesNil sets the value for AllEntities to be an explicit nil

### UnsetAllEntities
`func (o *EntitySelector) UnsetAllEntities()`

UnsetAllEntities ensures that no value is present for AllEntities, not even an explicit nil
### GetRegisteredEntities

`func (o *EntitySelector) GetRegisteredEntities() bool`

GetRegisteredEntities returns the RegisteredEntities field if non-nil, zero value otherwise.

### GetRegisteredEntitiesOk

`func (o *EntitySelector) GetRegisteredEntitiesOk() (*bool, bool)`

GetRegisteredEntitiesOk returns a tuple with the RegisteredEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredEntities

`func (o *EntitySelector) SetRegisteredEntities(v bool)`

SetRegisteredEntities sets RegisteredEntities field to given value.

### HasRegisteredEntities

`func (o *EntitySelector) HasRegisteredEntities() bool`

HasRegisteredEntities returns a boolean if a field has been set.

### SetRegisteredEntitiesNil

`func (o *EntitySelector) SetRegisteredEntitiesNil(b bool)`

 SetRegisteredEntitiesNil sets the value for RegisteredEntities to be an explicit nil

### UnsetRegisteredEntities
`func (o *EntitySelector) UnsetRegisteredEntities()`

UnsetRegisteredEntities ensures that no value is present for RegisteredEntities, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


