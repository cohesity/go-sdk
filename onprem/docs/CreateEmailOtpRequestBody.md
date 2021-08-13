# CreateEmailOtpRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **NullableString** | Specifies the login name of the Cohesity user | [optional] 
**Password** | Pointer to **NullableString** | Specifies the password of the Cohesity user | [optional] 
**Domain** | Pointer to **NullableString** | Specifies the domain the user is logging in to. For a local user the domain is LOCAL. For LDAP/AD user, the domain will map to a LDAP connection string. A user is uniquely identified by a combination of username and domain. LOCAL is the default domain. | [optional] 

## Methods

### NewCreateEmailOtpRequestBody

`func NewCreateEmailOtpRequestBody() *CreateEmailOtpRequestBody`

NewCreateEmailOtpRequestBody instantiates a new CreateEmailOtpRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEmailOtpRequestBodyWithDefaults

`func NewCreateEmailOtpRequestBodyWithDefaults() *CreateEmailOtpRequestBody`

NewCreateEmailOtpRequestBodyWithDefaults instantiates a new CreateEmailOtpRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *CreateEmailOtpRequestBody) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateEmailOtpRequestBody) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateEmailOtpRequestBody) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateEmailOtpRequestBody) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *CreateEmailOtpRequestBody) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *CreateEmailOtpRequestBody) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *CreateEmailOtpRequestBody) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateEmailOtpRequestBody) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateEmailOtpRequestBody) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateEmailOtpRequestBody) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *CreateEmailOtpRequestBody) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *CreateEmailOtpRequestBody) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetDomain

`func (o *CreateEmailOtpRequestBody) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreateEmailOtpRequestBody) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreateEmailOtpRequestBody) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *CreateEmailOtpRequestBody) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *CreateEmailOtpRequestBody) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *CreateEmailOtpRequestBody) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


