# SimpleAuthParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserDistinguishedName** | **NullableString** | Specifies the user distinguished name that is used for LDAP authentication. | 
**UserPassword** | Pointer to **NullableString** | Specifies the user password that is used for LDAP authentication. | [optional] 
**UseSsl** | Pointer to **NullableBool** | Specifies whether to use SSL for LDAP connections. | [optional] 

## Methods

### NewSimpleAuthParams

`func NewSimpleAuthParams(userDistinguishedName NullableString, ) *SimpleAuthParams`

NewSimpleAuthParams instantiates a new SimpleAuthParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleAuthParamsWithDefaults

`func NewSimpleAuthParamsWithDefaults() *SimpleAuthParams`

NewSimpleAuthParamsWithDefaults instantiates a new SimpleAuthParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserDistinguishedName

`func (o *SimpleAuthParams) GetUserDistinguishedName() string`

GetUserDistinguishedName returns the UserDistinguishedName field if non-nil, zero value otherwise.

### GetUserDistinguishedNameOk

`func (o *SimpleAuthParams) GetUserDistinguishedNameOk() (*string, bool)`

GetUserDistinguishedNameOk returns a tuple with the UserDistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDistinguishedName

`func (o *SimpleAuthParams) SetUserDistinguishedName(v string)`

SetUserDistinguishedName sets UserDistinguishedName field to given value.


### SetUserDistinguishedNameNil

`func (o *SimpleAuthParams) SetUserDistinguishedNameNil(b bool)`

 SetUserDistinguishedNameNil sets the value for UserDistinguishedName to be an explicit nil

### UnsetUserDistinguishedName
`func (o *SimpleAuthParams) UnsetUserDistinguishedName()`

UnsetUserDistinguishedName ensures that no value is present for UserDistinguishedName, not even an explicit nil
### GetUserPassword

`func (o *SimpleAuthParams) GetUserPassword() string`

GetUserPassword returns the UserPassword field if non-nil, zero value otherwise.

### GetUserPasswordOk

`func (o *SimpleAuthParams) GetUserPasswordOk() (*string, bool)`

GetUserPasswordOk returns a tuple with the UserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPassword

`func (o *SimpleAuthParams) SetUserPassword(v string)`

SetUserPassword sets UserPassword field to given value.

### HasUserPassword

`func (o *SimpleAuthParams) HasUserPassword() bool`

HasUserPassword returns a boolean if a field has been set.

### SetUserPasswordNil

`func (o *SimpleAuthParams) SetUserPasswordNil(b bool)`

 SetUserPasswordNil sets the value for UserPassword to be an explicit nil

### UnsetUserPassword
`func (o *SimpleAuthParams) UnsetUserPassword()`

UnsetUserPassword ensures that no value is present for UserPassword, not even an explicit nil
### GetUseSsl

`func (o *SimpleAuthParams) GetUseSsl() bool`

GetUseSsl returns the UseSsl field if non-nil, zero value otherwise.

### GetUseSslOk

`func (o *SimpleAuthParams) GetUseSslOk() (*bool, bool)`

GetUseSslOk returns a tuple with the UseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsl

`func (o *SimpleAuthParams) SetUseSsl(v bool)`

SetUseSsl sets UseSsl field to given value.

### HasUseSsl

`func (o *SimpleAuthParams) HasUseSsl() bool`

HasUseSsl returns a boolean if a field has been set.

### SetUseSslNil

`func (o *SimpleAuthParams) SetUseSslNil(b bool)`

 SetUseSslNil sets the value for UseSsl to be an explicit nil

### UnsetUseSsl
`func (o *SimpleAuthParams) UnsetUseSsl()`

UnsetUseSsl ensures that no value is present for UseSsl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


