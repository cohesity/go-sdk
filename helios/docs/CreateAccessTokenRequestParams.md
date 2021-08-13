# CreateAccessTokenRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **NullableString** | Specifies the login name of the Cohesity user. | 
**Password** | **NullableString** | Specifies the password of the Cohesity user account. | 
**Domain** | Pointer to **NullableString** | Specifies the domain the user is logging in to. For a local user the domain is LOCAL. For LDAP/AD user, the domain will map to a LDAP connection string. A user is uniquely identified by a combination of username and domain. LOCAL is the default domain. | [optional] 

## Methods

### NewCreateAccessTokenRequestParams

`func NewCreateAccessTokenRequestParams(username NullableString, password NullableString, ) *CreateAccessTokenRequestParams`

NewCreateAccessTokenRequestParams instantiates a new CreateAccessTokenRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccessTokenRequestParamsWithDefaults

`func NewCreateAccessTokenRequestParamsWithDefaults() *CreateAccessTokenRequestParams`

NewCreateAccessTokenRequestParamsWithDefaults instantiates a new CreateAccessTokenRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *CreateAccessTokenRequestParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateAccessTokenRequestParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateAccessTokenRequestParams) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *CreateAccessTokenRequestParams) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *CreateAccessTokenRequestParams) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *CreateAccessTokenRequestParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateAccessTokenRequestParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateAccessTokenRequestParams) SetPassword(v string)`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *CreateAccessTokenRequestParams) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *CreateAccessTokenRequestParams) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetDomain

`func (o *CreateAccessTokenRequestParams) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreateAccessTokenRequestParams) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreateAccessTokenRequestParams) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *CreateAccessTokenRequestParams) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *CreateAccessTokenRequestParams) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *CreateAccessTokenRequestParams) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


