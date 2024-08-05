# SecondaryId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Specifies name of the secondary ID for an object. | 
**Value** | Pointer to **NullableString** | Specifies value of the secondary ID for an object. | [optional] 

## Methods

### NewSecondaryId

`func NewSecondaryId(name NullableString, ) *SecondaryId`

NewSecondaryId instantiates a new SecondaryId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecondaryIdWithDefaults

`func NewSecondaryIdWithDefaults() *SecondaryId`

NewSecondaryIdWithDefaults instantiates a new SecondaryId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SecondaryId) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecondaryId) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecondaryId) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *SecondaryId) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SecondaryId) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetValue

`func (o *SecondaryId) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SecondaryId) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SecondaryId) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SecondaryId) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *SecondaryId) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *SecondaryId) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


