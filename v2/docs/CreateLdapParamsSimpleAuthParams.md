# CreateLdapParamsSimpleAuthParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseSsl** | Pointer to **NullableBool** | Specifies whether to use SSL for LDAP connections. | [optional] 
**UserDistinguishedName** | **NullableString** | Specifies the user distinguished name that is used for LDAP authentication. | 
**UserPassword** | Pointer to **NullableString** | Specifies the user password that is used for LDAP authentication. | [optional] 

## Methods

### NewCreateLdapParamsSimpleAuthParams

`func NewCreateLdapParamsSimpleAuthParams(userDistinguishedName NullableString, ) *CreateLdapParamsSimpleAuthParams`

NewCreateLdapParamsSimpleAuthParams instantiates a new CreateLdapParamsSimpleAuthParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLdapParamsSimpleAuthParamsWithDefaults

`func NewCreateLdapParamsSimpleAuthParamsWithDefaults() *CreateLdapParamsSimpleAuthParams`

NewCreateLdapParamsSimpleAuthParamsWithDefaults instantiates a new CreateLdapParamsSimpleAuthParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseSsl

`func (o *CreateLdapParamsSimpleAuthParams) GetUseSsl() bool`

GetUseSsl returns the UseSsl field if non-nil, zero value otherwise.

### GetUseSslOk

`func (o *CreateLdapParamsSimpleAuthParams) GetUseSslOk() (*bool, bool)`

GetUseSslOk returns a tuple with the UseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsl

`func (o *CreateLdapParamsSimpleAuthParams) SetUseSsl(v bool)`

SetUseSsl sets UseSsl field to given value.

### HasUseSsl

`func (o *CreateLdapParamsSimpleAuthParams) HasUseSsl() bool`

HasUseSsl returns a boolean if a field has been set.

### SetUseSslNil

`func (o *CreateLdapParamsSimpleAuthParams) SetUseSslNil(b bool)`

 SetUseSslNil sets the value for UseSsl to be an explicit nil

### UnsetUseSsl
`func (o *CreateLdapParamsSimpleAuthParams) UnsetUseSsl()`

UnsetUseSsl ensures that no value is present for UseSsl, not even an explicit nil
### GetUserDistinguishedName

`func (o *CreateLdapParamsSimpleAuthParams) GetUserDistinguishedName() string`

GetUserDistinguishedName returns the UserDistinguishedName field if non-nil, zero value otherwise.

### GetUserDistinguishedNameOk

`func (o *CreateLdapParamsSimpleAuthParams) GetUserDistinguishedNameOk() (*string, bool)`

GetUserDistinguishedNameOk returns a tuple with the UserDistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDistinguishedName

`func (o *CreateLdapParamsSimpleAuthParams) SetUserDistinguishedName(v string)`

SetUserDistinguishedName sets UserDistinguishedName field to given value.


### SetUserDistinguishedNameNil

`func (o *CreateLdapParamsSimpleAuthParams) SetUserDistinguishedNameNil(b bool)`

 SetUserDistinguishedNameNil sets the value for UserDistinguishedName to be an explicit nil

### UnsetUserDistinguishedName
`func (o *CreateLdapParamsSimpleAuthParams) UnsetUserDistinguishedName()`

UnsetUserDistinguishedName ensures that no value is present for UserDistinguishedName, not even an explicit nil
### GetUserPassword

`func (o *CreateLdapParamsSimpleAuthParams) GetUserPassword() string`

GetUserPassword returns the UserPassword field if non-nil, zero value otherwise.

### GetUserPasswordOk

`func (o *CreateLdapParamsSimpleAuthParams) GetUserPasswordOk() (*string, bool)`

GetUserPasswordOk returns a tuple with the UserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPassword

`func (o *CreateLdapParamsSimpleAuthParams) SetUserPassword(v string)`

SetUserPassword sets UserPassword field to given value.

### HasUserPassword

`func (o *CreateLdapParamsSimpleAuthParams) HasUserPassword() bool`

HasUserPassword returns a boolean if a field has been set.

### SetUserPasswordNil

`func (o *CreateLdapParamsSimpleAuthParams) SetUserPasswordNil(b bool)`

 SetUserPasswordNil sets the value for UserPassword to be an explicit nil

### UnsetUserPassword
`func (o *CreateLdapParamsSimpleAuthParams) UnsetUserPassword()`

UnsetUserPassword ensures that no value is present for UserPassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


