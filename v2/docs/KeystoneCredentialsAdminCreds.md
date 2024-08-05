# KeystoneCredentialsAdminCreds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **NullableString** | Specifies the administrator domain name. | 
**Password** | Pointer to **NullableString** | Specifies the password of Keystone administrator. | [optional] 
**Username** | **NullableString** | Specifies the username of Keystone administrator. | 

## Methods

### NewKeystoneCredentialsAdminCreds

`func NewKeystoneCredentialsAdminCreds(domain NullableString, username NullableString, ) *KeystoneCredentialsAdminCreds`

NewKeystoneCredentialsAdminCreds instantiates a new KeystoneCredentialsAdminCreds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeystoneCredentialsAdminCredsWithDefaults

`func NewKeystoneCredentialsAdminCredsWithDefaults() *KeystoneCredentialsAdminCreds`

NewKeystoneCredentialsAdminCredsWithDefaults instantiates a new KeystoneCredentialsAdminCreds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *KeystoneCredentialsAdminCreds) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *KeystoneCredentialsAdminCreds) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *KeystoneCredentialsAdminCreds) SetDomain(v string)`

SetDomain sets Domain field to given value.


### SetDomainNil

`func (o *KeystoneCredentialsAdminCreds) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *KeystoneCredentialsAdminCreds) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetPassword

`func (o *KeystoneCredentialsAdminCreds) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *KeystoneCredentialsAdminCreds) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *KeystoneCredentialsAdminCreds) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *KeystoneCredentialsAdminCreds) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *KeystoneCredentialsAdminCreds) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *KeystoneCredentialsAdminCreds) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetUsername

`func (o *KeystoneCredentialsAdminCreds) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *KeystoneCredentialsAdminCreds) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *KeystoneCredentialsAdminCreds) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *KeystoneCredentialsAdminCreds) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *KeystoneCredentialsAdminCreds) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


