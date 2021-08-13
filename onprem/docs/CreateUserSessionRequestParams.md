# CreateUserSessionRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **NullableString** | Specifies the login name of the Cohesity user | [optional] 
**Password** | Pointer to **NullableString** | Specifies the password of the Cohesity user | [optional] 
**Domain** | Pointer to **NullableString** | Specifies the domain the user is logging in to. For a local user the domain is LOCAL. For LDAP/AD user, the domain will map to a LDAP connection string. A user is uniquely identified by a combination of username and domain. LOCAL is the default domain. | [optional] 
**Certificate** | Pointer to **NullableString** | Specifies the certificate for cert based authentication. | [optional] 
**PrivateKey** | Pointer to **NullableString** | Specifies the private key for cert based authentication. | [optional] 
**OtpCode** | Pointer to **NullableString** | Specifies OTP code for MFA verification. | [optional] 
**OtpType** | Pointer to **NullableString** | Specifies OTP Type for MFA verification. | [optional] 

## Methods

### NewCreateUserSessionRequestParams

`func NewCreateUserSessionRequestParams() *CreateUserSessionRequestParams`

NewCreateUserSessionRequestParams instantiates a new CreateUserSessionRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserSessionRequestParamsWithDefaults

`func NewCreateUserSessionRequestParamsWithDefaults() *CreateUserSessionRequestParams`

NewCreateUserSessionRequestParamsWithDefaults instantiates a new CreateUserSessionRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *CreateUserSessionRequestParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateUserSessionRequestParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateUserSessionRequestParams) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateUserSessionRequestParams) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *CreateUserSessionRequestParams) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *CreateUserSessionRequestParams) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *CreateUserSessionRequestParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateUserSessionRequestParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateUserSessionRequestParams) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateUserSessionRequestParams) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *CreateUserSessionRequestParams) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *CreateUserSessionRequestParams) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetDomain

`func (o *CreateUserSessionRequestParams) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreateUserSessionRequestParams) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreateUserSessionRequestParams) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *CreateUserSessionRequestParams) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *CreateUserSessionRequestParams) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *CreateUserSessionRequestParams) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetCertificate

`func (o *CreateUserSessionRequestParams) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CreateUserSessionRequestParams) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CreateUserSessionRequestParams) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CreateUserSessionRequestParams) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *CreateUserSessionRequestParams) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *CreateUserSessionRequestParams) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetPrivateKey

`func (o *CreateUserSessionRequestParams) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CreateUserSessionRequestParams) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CreateUserSessionRequestParams) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *CreateUserSessionRequestParams) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *CreateUserSessionRequestParams) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *CreateUserSessionRequestParams) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
### GetOtpCode

`func (o *CreateUserSessionRequestParams) GetOtpCode() string`

GetOtpCode returns the OtpCode field if non-nil, zero value otherwise.

### GetOtpCodeOk

`func (o *CreateUserSessionRequestParams) GetOtpCodeOk() (*string, bool)`

GetOtpCodeOk returns a tuple with the OtpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpCode

`func (o *CreateUserSessionRequestParams) SetOtpCode(v string)`

SetOtpCode sets OtpCode field to given value.

### HasOtpCode

`func (o *CreateUserSessionRequestParams) HasOtpCode() bool`

HasOtpCode returns a boolean if a field has been set.

### SetOtpCodeNil

`func (o *CreateUserSessionRequestParams) SetOtpCodeNil(b bool)`

 SetOtpCodeNil sets the value for OtpCode to be an explicit nil

### UnsetOtpCode
`func (o *CreateUserSessionRequestParams) UnsetOtpCode()`

UnsetOtpCode ensures that no value is present for OtpCode, not even an explicit nil
### GetOtpType

`func (o *CreateUserSessionRequestParams) GetOtpType() string`

GetOtpType returns the OtpType field if non-nil, zero value otherwise.

### GetOtpTypeOk

`func (o *CreateUserSessionRequestParams) GetOtpTypeOk() (*string, bool)`

GetOtpTypeOk returns a tuple with the OtpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpType

`func (o *CreateUserSessionRequestParams) SetOtpType(v string)`

SetOtpType sets OtpType field to given value.

### HasOtpType

`func (o *CreateUserSessionRequestParams) HasOtpType() bool`

HasOtpType returns a boolean if a field has been set.

### SetOtpTypeNil

`func (o *CreateUserSessionRequestParams) SetOtpTypeNil(b bool)`

 SetOtpTypeNil sets the value for OtpType to be an explicit nil

### UnsetOtpType
`func (o *CreateUserSessionRequestParams) UnsetOtpType()`

UnsetOtpType ensures that no value is present for OtpType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


