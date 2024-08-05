# PreCheckValidation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPassed** | Pointer to **NullableBool** | Specifies validation passed or failed | [optional] 
**Message** | Pointer to **NullableString** | Specifies the validation failure message | [optional] 
**Validation** | Pointer to **NullableString** | Specifies validation type | [optional] 

## Methods

### NewPreCheckValidation

`func NewPreCheckValidation() *PreCheckValidation`

NewPreCheckValidation instantiates a new PreCheckValidation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreCheckValidationWithDefaults

`func NewPreCheckValidationWithDefaults() *PreCheckValidation`

NewPreCheckValidationWithDefaults instantiates a new PreCheckValidation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPassed

`func (o *PreCheckValidation) GetIsPassed() bool`

GetIsPassed returns the IsPassed field if non-nil, zero value otherwise.

### GetIsPassedOk

`func (o *PreCheckValidation) GetIsPassedOk() (*bool, bool)`

GetIsPassedOk returns a tuple with the IsPassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPassed

`func (o *PreCheckValidation) SetIsPassed(v bool)`

SetIsPassed sets IsPassed field to given value.

### HasIsPassed

`func (o *PreCheckValidation) HasIsPassed() bool`

HasIsPassed returns a boolean if a field has been set.

### SetIsPassedNil

`func (o *PreCheckValidation) SetIsPassedNil(b bool)`

 SetIsPassedNil sets the value for IsPassed to be an explicit nil

### UnsetIsPassed
`func (o *PreCheckValidation) UnsetIsPassed()`

UnsetIsPassed ensures that no value is present for IsPassed, not even an explicit nil
### GetMessage

`func (o *PreCheckValidation) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PreCheckValidation) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PreCheckValidation) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PreCheckValidation) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *PreCheckValidation) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *PreCheckValidation) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetValidation

`func (o *PreCheckValidation) GetValidation() string`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *PreCheckValidation) GetValidationOk() (*string, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *PreCheckValidation) SetValidation(v string)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *PreCheckValidation) HasValidation() bool`

HasValidation returns a boolean if a field has been set.

### SetValidationNil

`func (o *PreCheckValidation) SetValidationNil(b bool)`

 SetValidationNil sets the value for Validation to be an explicit nil

### UnsetValidation
`func (o *PreCheckValidation) UnsetValidation()`

UnsetValidation ensures that no value is present for Validation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


