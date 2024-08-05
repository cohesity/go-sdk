# ObjectUniqueIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to **NullableString** | Environment of the object | [optional] 
**ObjectIdentifier** | Pointer to **NullableString** | Unique identifier of the object. This is assigned by the third party adapter and saved with the object data. | [optional] 
**ParentIdentifier** | Pointer to **NullableString** | Unique identifier of the parent for the object. This is assigned by the third party adapter and saved with the object data. | [optional] 

## Methods

### NewObjectUniqueIdentifier

`func NewObjectUniqueIdentifier() *ObjectUniqueIdentifier`

NewObjectUniqueIdentifier instantiates a new ObjectUniqueIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectUniqueIdentifierWithDefaults

`func NewObjectUniqueIdentifierWithDefaults() *ObjectUniqueIdentifier`

NewObjectUniqueIdentifierWithDefaults instantiates a new ObjectUniqueIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *ObjectUniqueIdentifier) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ObjectUniqueIdentifier) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ObjectUniqueIdentifier) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ObjectUniqueIdentifier) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *ObjectUniqueIdentifier) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ObjectUniqueIdentifier) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetObjectIdentifier

`func (o *ObjectUniqueIdentifier) GetObjectIdentifier() string`

GetObjectIdentifier returns the ObjectIdentifier field if non-nil, zero value otherwise.

### GetObjectIdentifierOk

`func (o *ObjectUniqueIdentifier) GetObjectIdentifierOk() (*string, bool)`

GetObjectIdentifierOk returns a tuple with the ObjectIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIdentifier

`func (o *ObjectUniqueIdentifier) SetObjectIdentifier(v string)`

SetObjectIdentifier sets ObjectIdentifier field to given value.

### HasObjectIdentifier

`func (o *ObjectUniqueIdentifier) HasObjectIdentifier() bool`

HasObjectIdentifier returns a boolean if a field has been set.

### SetObjectIdentifierNil

`func (o *ObjectUniqueIdentifier) SetObjectIdentifierNil(b bool)`

 SetObjectIdentifierNil sets the value for ObjectIdentifier to be an explicit nil

### UnsetObjectIdentifier
`func (o *ObjectUniqueIdentifier) UnsetObjectIdentifier()`

UnsetObjectIdentifier ensures that no value is present for ObjectIdentifier, not even an explicit nil
### GetParentIdentifier

`func (o *ObjectUniqueIdentifier) GetParentIdentifier() string`

GetParentIdentifier returns the ParentIdentifier field if non-nil, zero value otherwise.

### GetParentIdentifierOk

`func (o *ObjectUniqueIdentifier) GetParentIdentifierOk() (*string, bool)`

GetParentIdentifierOk returns a tuple with the ParentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIdentifier

`func (o *ObjectUniqueIdentifier) SetParentIdentifier(v string)`

SetParentIdentifier sets ParentIdentifier field to given value.

### HasParentIdentifier

`func (o *ObjectUniqueIdentifier) HasParentIdentifier() bool`

HasParentIdentifier returns a boolean if a field has been set.

### SetParentIdentifierNil

`func (o *ObjectUniqueIdentifier) SetParentIdentifierNil(b bool)`

 SetParentIdentifierNil sets the value for ParentIdentifier to be an explicit nil

### UnsetParentIdentifier
`func (o *ObjectUniqueIdentifier) UnsetParentIdentifier()`

UnsetParentIdentifier ensures that no value is present for ParentIdentifier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


