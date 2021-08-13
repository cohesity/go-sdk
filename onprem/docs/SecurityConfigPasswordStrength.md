# SecurityConfigPasswordStrength

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinLength** | Pointer to **NullableInt32** | Specifies the password minimum length. | [optional] 
**IncludeUpperLetter** | Pointer to **NullableBool** | Specifies if the password needs to have at least one uppercase letter. | [optional] 
**IncludeLowerLetter** | Pointer to **NullableBool** | Specifies if the password needs to have at least one lowercase letter. | [optional] 
**IncludeNumber** | Pointer to **NullableBool** | Specifies if the password needs to have at least one number. | [optional] 
**IncludeSpecialChar** | Pointer to **NullableBool** | Specifies if the password needs to have at least one special character. | [optional] 

## Methods

### NewSecurityConfigPasswordStrength

`func NewSecurityConfigPasswordStrength() *SecurityConfigPasswordStrength`

NewSecurityConfigPasswordStrength instantiates a new SecurityConfigPasswordStrength object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityConfigPasswordStrengthWithDefaults

`func NewSecurityConfigPasswordStrengthWithDefaults() *SecurityConfigPasswordStrength`

NewSecurityConfigPasswordStrengthWithDefaults instantiates a new SecurityConfigPasswordStrength object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinLength

`func (o *SecurityConfigPasswordStrength) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *SecurityConfigPasswordStrength) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *SecurityConfigPasswordStrength) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *SecurityConfigPasswordStrength) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### SetMinLengthNil

`func (o *SecurityConfigPasswordStrength) SetMinLengthNil(b bool)`

 SetMinLengthNil sets the value for MinLength to be an explicit nil

### UnsetMinLength
`func (o *SecurityConfigPasswordStrength) UnsetMinLength()`

UnsetMinLength ensures that no value is present for MinLength, not even an explicit nil
### GetIncludeUpperLetter

`func (o *SecurityConfigPasswordStrength) GetIncludeUpperLetter() bool`

GetIncludeUpperLetter returns the IncludeUpperLetter field if non-nil, zero value otherwise.

### GetIncludeUpperLetterOk

`func (o *SecurityConfigPasswordStrength) GetIncludeUpperLetterOk() (*bool, bool)`

GetIncludeUpperLetterOk returns a tuple with the IncludeUpperLetter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeUpperLetter

`func (o *SecurityConfigPasswordStrength) SetIncludeUpperLetter(v bool)`

SetIncludeUpperLetter sets IncludeUpperLetter field to given value.

### HasIncludeUpperLetter

`func (o *SecurityConfigPasswordStrength) HasIncludeUpperLetter() bool`

HasIncludeUpperLetter returns a boolean if a field has been set.

### SetIncludeUpperLetterNil

`func (o *SecurityConfigPasswordStrength) SetIncludeUpperLetterNil(b bool)`

 SetIncludeUpperLetterNil sets the value for IncludeUpperLetter to be an explicit nil

### UnsetIncludeUpperLetter
`func (o *SecurityConfigPasswordStrength) UnsetIncludeUpperLetter()`

UnsetIncludeUpperLetter ensures that no value is present for IncludeUpperLetter, not even an explicit nil
### GetIncludeLowerLetter

`func (o *SecurityConfigPasswordStrength) GetIncludeLowerLetter() bool`

GetIncludeLowerLetter returns the IncludeLowerLetter field if non-nil, zero value otherwise.

### GetIncludeLowerLetterOk

`func (o *SecurityConfigPasswordStrength) GetIncludeLowerLetterOk() (*bool, bool)`

GetIncludeLowerLetterOk returns a tuple with the IncludeLowerLetter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLowerLetter

`func (o *SecurityConfigPasswordStrength) SetIncludeLowerLetter(v bool)`

SetIncludeLowerLetter sets IncludeLowerLetter field to given value.

### HasIncludeLowerLetter

`func (o *SecurityConfigPasswordStrength) HasIncludeLowerLetter() bool`

HasIncludeLowerLetter returns a boolean if a field has been set.

### SetIncludeLowerLetterNil

`func (o *SecurityConfigPasswordStrength) SetIncludeLowerLetterNil(b bool)`

 SetIncludeLowerLetterNil sets the value for IncludeLowerLetter to be an explicit nil

### UnsetIncludeLowerLetter
`func (o *SecurityConfigPasswordStrength) UnsetIncludeLowerLetter()`

UnsetIncludeLowerLetter ensures that no value is present for IncludeLowerLetter, not even an explicit nil
### GetIncludeNumber

`func (o *SecurityConfigPasswordStrength) GetIncludeNumber() bool`

GetIncludeNumber returns the IncludeNumber field if non-nil, zero value otherwise.

### GetIncludeNumberOk

`func (o *SecurityConfigPasswordStrength) GetIncludeNumberOk() (*bool, bool)`

GetIncludeNumberOk returns a tuple with the IncludeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNumber

`func (o *SecurityConfigPasswordStrength) SetIncludeNumber(v bool)`

SetIncludeNumber sets IncludeNumber field to given value.

### HasIncludeNumber

`func (o *SecurityConfigPasswordStrength) HasIncludeNumber() bool`

HasIncludeNumber returns a boolean if a field has been set.

### SetIncludeNumberNil

`func (o *SecurityConfigPasswordStrength) SetIncludeNumberNil(b bool)`

 SetIncludeNumberNil sets the value for IncludeNumber to be an explicit nil

### UnsetIncludeNumber
`func (o *SecurityConfigPasswordStrength) UnsetIncludeNumber()`

UnsetIncludeNumber ensures that no value is present for IncludeNumber, not even an explicit nil
### GetIncludeSpecialChar

`func (o *SecurityConfigPasswordStrength) GetIncludeSpecialChar() bool`

GetIncludeSpecialChar returns the IncludeSpecialChar field if non-nil, zero value otherwise.

### GetIncludeSpecialCharOk

`func (o *SecurityConfigPasswordStrength) GetIncludeSpecialCharOk() (*bool, bool)`

GetIncludeSpecialCharOk returns a tuple with the IncludeSpecialChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSpecialChar

`func (o *SecurityConfigPasswordStrength) SetIncludeSpecialChar(v bool)`

SetIncludeSpecialChar sets IncludeSpecialChar field to given value.

### HasIncludeSpecialChar

`func (o *SecurityConfigPasswordStrength) HasIncludeSpecialChar() bool`

HasIncludeSpecialChar returns a boolean if a field has been set.

### SetIncludeSpecialCharNil

`func (o *SecurityConfigPasswordStrength) SetIncludeSpecialCharNil(b bool)`

 SetIncludeSpecialCharNil sets the value for IncludeSpecialChar to be an explicit nil

### UnsetIncludeSpecialChar
`func (o *SecurityConfigPasswordStrength) UnsetIncludeSpecialChar()`

UnsetIncludeSpecialChar ensures that no value is present for IncludeSpecialChar, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


