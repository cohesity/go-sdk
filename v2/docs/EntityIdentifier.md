# EntityIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **NullableString** | Specifies the Id of an Entity. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of an Entity. | [optional] 

## Methods

### NewEntityIdentifier

`func NewEntityIdentifier() *EntityIdentifier`

NewEntityIdentifier instantiates a new EntityIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityIdentifierWithDefaults

`func NewEntityIdentifierWithDefaults() *EntityIdentifier`

NewEntityIdentifierWithDefaults instantiates a new EntityIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *EntityIdentifier) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *EntityIdentifier) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *EntityIdentifier) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *EntityIdentifier) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *EntityIdentifier) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *EntityIdentifier) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetName

`func (o *EntityIdentifier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityIdentifier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityIdentifier) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntityIdentifier) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EntityIdentifier) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EntityIdentifier) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


