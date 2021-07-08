# Pattern

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSystemDefined** | Pointer to **NullableBool** | Whether this pattern is system defined. | [optional] 
**Name** | Pointer to **NullableString** | Name of the pattern. This is marked optional but is required. | [optional] 
**Type** | Pointer to **NullableInt32** | Pattern type. | [optional] 
**Value** | Pointer to **NullableString** | Value of the pattern. This is marked optional but is required. | [optional] 

## Methods

### NewPattern

`func NewPattern() *Pattern`

NewPattern instantiates a new Pattern object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatternWithDefaults

`func NewPatternWithDefaults() *Pattern`

NewPatternWithDefaults instantiates a new Pattern object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSystemDefined

`func (o *Pattern) GetIsSystemDefined() bool`

GetIsSystemDefined returns the IsSystemDefined field if non-nil, zero value otherwise.

### GetIsSystemDefinedOk

`func (o *Pattern) GetIsSystemDefinedOk() (*bool, bool)`

GetIsSystemDefinedOk returns a tuple with the IsSystemDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemDefined

`func (o *Pattern) SetIsSystemDefined(v bool)`

SetIsSystemDefined sets IsSystemDefined field to given value.

### HasIsSystemDefined

`func (o *Pattern) HasIsSystemDefined() bool`

HasIsSystemDefined returns a boolean if a field has been set.

### SetIsSystemDefinedNil

`func (o *Pattern) SetIsSystemDefinedNil(b bool)`

 SetIsSystemDefinedNil sets the value for IsSystemDefined to be an explicit nil

### UnsetIsSystemDefined
`func (o *Pattern) UnsetIsSystemDefined()`

UnsetIsSystemDefined ensures that no value is present for IsSystemDefined, not even an explicit nil
### GetName

`func (o *Pattern) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Pattern) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Pattern) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Pattern) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Pattern) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Pattern) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *Pattern) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Pattern) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Pattern) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *Pattern) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Pattern) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Pattern) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetValue

`func (o *Pattern) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Pattern) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Pattern) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Pattern) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *Pattern) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Pattern) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


