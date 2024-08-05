# AwsCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | Pointer to **NullableString** | Specifies the type of authentication being used in the request. | [optional] 
**DirectoryDNSAddress** | Pointer to **NullableString** | Specifies the DNS address of the AWS managed active directory in which. Currently is set only for kerberos authentication. | [optional] 
**Password** | Pointer to **NullableString** | Specifies the password to access target entity. | [optional] 
**RealmName** | Pointer to **NullableString** | Specifies the Kerberos realm name for a Kerberos-secured target. | [optional] 
**Username** | Pointer to **NullableString** | Specifies the username to access target entity. | [optional] 

## Methods

### NewAwsCredentials

`func NewAwsCredentials() *AwsCredentials`

NewAwsCredentials instantiates a new AwsCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsCredentialsWithDefaults

`func NewAwsCredentialsWithDefaults() *AwsCredentials`

NewAwsCredentialsWithDefaults instantiates a new AwsCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *AwsCredentials) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *AwsCredentials) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *AwsCredentials) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *AwsCredentials) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### SetAuthTypeNil

`func (o *AwsCredentials) SetAuthTypeNil(b bool)`

 SetAuthTypeNil sets the value for AuthType to be an explicit nil

### UnsetAuthType
`func (o *AwsCredentials) UnsetAuthType()`

UnsetAuthType ensures that no value is present for AuthType, not even an explicit nil
### GetDirectoryDNSAddress

`func (o *AwsCredentials) GetDirectoryDNSAddress() string`

GetDirectoryDNSAddress returns the DirectoryDNSAddress field if non-nil, zero value otherwise.

### GetDirectoryDNSAddressOk

`func (o *AwsCredentials) GetDirectoryDNSAddressOk() (*string, bool)`

GetDirectoryDNSAddressOk returns a tuple with the DirectoryDNSAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryDNSAddress

`func (o *AwsCredentials) SetDirectoryDNSAddress(v string)`

SetDirectoryDNSAddress sets DirectoryDNSAddress field to given value.

### HasDirectoryDNSAddress

`func (o *AwsCredentials) HasDirectoryDNSAddress() bool`

HasDirectoryDNSAddress returns a boolean if a field has been set.

### SetDirectoryDNSAddressNil

`func (o *AwsCredentials) SetDirectoryDNSAddressNil(b bool)`

 SetDirectoryDNSAddressNil sets the value for DirectoryDNSAddress to be an explicit nil

### UnsetDirectoryDNSAddress
`func (o *AwsCredentials) UnsetDirectoryDNSAddress()`

UnsetDirectoryDNSAddress ensures that no value is present for DirectoryDNSAddress, not even an explicit nil
### GetPassword

`func (o *AwsCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AwsCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AwsCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AwsCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *AwsCredentials) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *AwsCredentials) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetRealmName

`func (o *AwsCredentials) GetRealmName() string`

GetRealmName returns the RealmName field if non-nil, zero value otherwise.

### GetRealmNameOk

`func (o *AwsCredentials) GetRealmNameOk() (*string, bool)`

GetRealmNameOk returns a tuple with the RealmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmName

`func (o *AwsCredentials) SetRealmName(v string)`

SetRealmName sets RealmName field to given value.

### HasRealmName

`func (o *AwsCredentials) HasRealmName() bool`

HasRealmName returns a boolean if a field has been set.

### SetRealmNameNil

`func (o *AwsCredentials) SetRealmNameNil(b bool)`

 SetRealmNameNil sets the value for RealmName to be an explicit nil

### UnsetRealmName
`func (o *AwsCredentials) UnsetRealmName()`

UnsetRealmName ensures that no value is present for RealmName, not even an explicit nil
### GetUsername

`func (o *AwsCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AwsCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AwsCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AwsCredentials) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *AwsCredentials) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *AwsCredentials) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


