# KeystoneAdminParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **NullableString** | Specifies the administrator domain name. | 
**Username** | **NullableString** | Specifies the username of Keystone administrator. | 
**Password** | Pointer to **NullableString** | Specifies the password of Keystone administrator. | [optional] 

## Methods

### NewKeystoneAdminParams

`func NewKeystoneAdminParams(domain NullableString, username NullableString, ) *KeystoneAdminParams`

NewKeystoneAdminParams instantiates a new KeystoneAdminParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeystoneAdminParamsWithDefaults

`func NewKeystoneAdminParamsWithDefaults() *KeystoneAdminParams`

NewKeystoneAdminParamsWithDefaults instantiates a new KeystoneAdminParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *KeystoneAdminParams) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *KeystoneAdminParams) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *KeystoneAdminParams) SetDomain(v string)`

SetDomain sets Domain field to given value.


### SetDomainNil

`func (o *KeystoneAdminParams) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *KeystoneAdminParams) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetUsername

`func (o *KeystoneAdminParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *KeystoneAdminParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *KeystoneAdminParams) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *KeystoneAdminParams) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *KeystoneAdminParams) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *KeystoneAdminParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *KeystoneAdminParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *KeystoneAdminParams) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *KeystoneAdminParams) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *KeystoneAdminParams) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *KeystoneAdminParams) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


