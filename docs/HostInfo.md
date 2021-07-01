# HostInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **NullableString** | Specifies the password of the host to establish SSH connection. The certificate is copied to the host after generating the certificate on the cluster. | [optional] 
**ServerName** | Pointer to **NullableString** | Specifies the servername of the host where certificate is to be deployed. | [optional] 
**Target** | Pointer to **NullableString** | Specifies the target location on the host where the certificate is deployed. | [optional] 
**UserName** | Pointer to **NullableString** | Specifies the username of the host. | [optional] 

## Methods

### NewHostInfo

`func NewHostInfo() *HostInfo`

NewHostInfo instantiates a new HostInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostInfoWithDefaults

`func NewHostInfoWithDefaults() *HostInfo`

NewHostInfoWithDefaults instantiates a new HostInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *HostInfo) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HostInfo) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HostInfo) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *HostInfo) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *HostInfo) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *HostInfo) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetServerName

`func (o *HostInfo) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *HostInfo) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *HostInfo) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *HostInfo) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### SetServerNameNil

`func (o *HostInfo) SetServerNameNil(b bool)`

 SetServerNameNil sets the value for ServerName to be an explicit nil

### UnsetServerName
`func (o *HostInfo) UnsetServerName()`

UnsetServerName ensures that no value is present for ServerName, not even an explicit nil
### GetTarget

`func (o *HostInfo) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *HostInfo) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *HostInfo) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *HostInfo) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *HostInfo) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *HostInfo) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetUserName

`func (o *HostInfo) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *HostInfo) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *HostInfo) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *HostInfo) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *HostInfo) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *HostInfo) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


