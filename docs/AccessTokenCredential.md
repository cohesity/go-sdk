# AccessTokenCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **NullableString** | Specifies the domain the user is logging in to. For a Local user model, the domain is always LOCAL. For LDAP/AD user models, the domain will map to an LDAP connection string. A user is uniquely identified by a combination of username and domain. If this is not set, LOCAL is assumed. | [optional] 
**Password** | **NullableString** | Specifies the password of the Cohesity user account. | 
**Username** | **NullableString** | Specifies the login name of the Cohesity user. | 

## Methods

### NewAccessTokenCredential

`func NewAccessTokenCredential(password NullableString, username NullableString, ) *AccessTokenCredential`

NewAccessTokenCredential instantiates a new AccessTokenCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenCredentialWithDefaults

`func NewAccessTokenCredentialWithDefaults() *AccessTokenCredential`

NewAccessTokenCredentialWithDefaults instantiates a new AccessTokenCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *AccessTokenCredential) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *AccessTokenCredential) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *AccessTokenCredential) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *AccessTokenCredential) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *AccessTokenCredential) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *AccessTokenCredential) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetPassword

`func (o *AccessTokenCredential) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AccessTokenCredential) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AccessTokenCredential) SetPassword(v string)`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *AccessTokenCredential) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *AccessTokenCredential) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetUsername

`func (o *AccessTokenCredential) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AccessTokenCredential) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AccessTokenCredential) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *AccessTokenCredential) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *AccessTokenCredential) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


