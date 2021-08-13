# SecurityConfigPasswordReuse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumDisallowedOldPasswords** | Pointer to **NullableInt32** | Specifies the minimum number of old passwords that are not allowed for changing the password. | [optional] 
**NumDifferentChars** | Pointer to **NullableInt32** | Specifies the number of characters in the new password that needs to be different from the old password (only applicable when changing the user&#39;s own password). | [optional] 

## Methods

### NewSecurityConfigPasswordReuse

`func NewSecurityConfigPasswordReuse() *SecurityConfigPasswordReuse`

NewSecurityConfigPasswordReuse instantiates a new SecurityConfigPasswordReuse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityConfigPasswordReuseWithDefaults

`func NewSecurityConfigPasswordReuseWithDefaults() *SecurityConfigPasswordReuse`

NewSecurityConfigPasswordReuseWithDefaults instantiates a new SecurityConfigPasswordReuse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumDisallowedOldPasswords

`func (o *SecurityConfigPasswordReuse) GetNumDisallowedOldPasswords() int32`

GetNumDisallowedOldPasswords returns the NumDisallowedOldPasswords field if non-nil, zero value otherwise.

### GetNumDisallowedOldPasswordsOk

`func (o *SecurityConfigPasswordReuse) GetNumDisallowedOldPasswordsOk() (*int32, bool)`

GetNumDisallowedOldPasswordsOk returns a tuple with the NumDisallowedOldPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDisallowedOldPasswords

`func (o *SecurityConfigPasswordReuse) SetNumDisallowedOldPasswords(v int32)`

SetNumDisallowedOldPasswords sets NumDisallowedOldPasswords field to given value.

### HasNumDisallowedOldPasswords

`func (o *SecurityConfigPasswordReuse) HasNumDisallowedOldPasswords() bool`

HasNumDisallowedOldPasswords returns a boolean if a field has been set.

### SetNumDisallowedOldPasswordsNil

`func (o *SecurityConfigPasswordReuse) SetNumDisallowedOldPasswordsNil(b bool)`

 SetNumDisallowedOldPasswordsNil sets the value for NumDisallowedOldPasswords to be an explicit nil

### UnsetNumDisallowedOldPasswords
`func (o *SecurityConfigPasswordReuse) UnsetNumDisallowedOldPasswords()`

UnsetNumDisallowedOldPasswords ensures that no value is present for NumDisallowedOldPasswords, not even an explicit nil
### GetNumDifferentChars

`func (o *SecurityConfigPasswordReuse) GetNumDifferentChars() int32`

GetNumDifferentChars returns the NumDifferentChars field if non-nil, zero value otherwise.

### GetNumDifferentCharsOk

`func (o *SecurityConfigPasswordReuse) GetNumDifferentCharsOk() (*int32, bool)`

GetNumDifferentCharsOk returns a tuple with the NumDifferentChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDifferentChars

`func (o *SecurityConfigPasswordReuse) SetNumDifferentChars(v int32)`

SetNumDifferentChars sets NumDifferentChars field to given value.

### HasNumDifferentChars

`func (o *SecurityConfigPasswordReuse) HasNumDifferentChars() bool`

HasNumDifferentChars returns a boolean if a field has been set.

### SetNumDifferentCharsNil

`func (o *SecurityConfigPasswordReuse) SetNumDifferentCharsNil(b bool)`

 SetNumDifferentCharsNil sets the value for NumDifferentChars to be an explicit nil

### UnsetNumDifferentChars
`func (o *SecurityConfigPasswordReuse) UnsetNumDifferentChars()`

UnsetNumDifferentChars ensures that no value is present for NumDifferentChars, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


