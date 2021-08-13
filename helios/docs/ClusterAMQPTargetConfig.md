# ClusterAMQPTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerIp** | Pointer to **NullableString** | Specifies the server ip. | [optional] 
**Username** | Pointer to **NullableString** | Specifies the username. | [optional] 
**Password** | Pointer to **NullableString** | Specifies the password. | [optional] 
**VirtualHost** | Pointer to **NullableString** | Specifies the virtual host. | [optional] 
**Exchange** | Pointer to **NullableString** | Specifies the exchange. | [optional] 
**FilerId** | Pointer to **NullableInt64** | Specifies the filer id. | [optional] 
**Certificate** | Pointer to **NullableString** | Specifies the certificate. | [optional] 

## Methods

### NewClusterAMQPTargetConfig

`func NewClusterAMQPTargetConfig() *ClusterAMQPTargetConfig`

NewClusterAMQPTargetConfig instantiates a new ClusterAMQPTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAMQPTargetConfigWithDefaults

`func NewClusterAMQPTargetConfigWithDefaults() *ClusterAMQPTargetConfig`

NewClusterAMQPTargetConfigWithDefaults instantiates a new ClusterAMQPTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerIp

`func (o *ClusterAMQPTargetConfig) GetServerIp() string`

GetServerIp returns the ServerIp field if non-nil, zero value otherwise.

### GetServerIpOk

`func (o *ClusterAMQPTargetConfig) GetServerIpOk() (*string, bool)`

GetServerIpOk returns a tuple with the ServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIp

`func (o *ClusterAMQPTargetConfig) SetServerIp(v string)`

SetServerIp sets ServerIp field to given value.

### HasServerIp

`func (o *ClusterAMQPTargetConfig) HasServerIp() bool`

HasServerIp returns a boolean if a field has been set.

### SetServerIpNil

`func (o *ClusterAMQPTargetConfig) SetServerIpNil(b bool)`

 SetServerIpNil sets the value for ServerIp to be an explicit nil

### UnsetServerIp
`func (o *ClusterAMQPTargetConfig) UnsetServerIp()`

UnsetServerIp ensures that no value is present for ServerIp, not even an explicit nil
### GetUsername

`func (o *ClusterAMQPTargetConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ClusterAMQPTargetConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ClusterAMQPTargetConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ClusterAMQPTargetConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ClusterAMQPTargetConfig) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ClusterAMQPTargetConfig) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *ClusterAMQPTargetConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ClusterAMQPTargetConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ClusterAMQPTargetConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ClusterAMQPTargetConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ClusterAMQPTargetConfig) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ClusterAMQPTargetConfig) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetVirtualHost

`func (o *ClusterAMQPTargetConfig) GetVirtualHost() string`

GetVirtualHost returns the VirtualHost field if non-nil, zero value otherwise.

### GetVirtualHostOk

`func (o *ClusterAMQPTargetConfig) GetVirtualHostOk() (*string, bool)`

GetVirtualHostOk returns a tuple with the VirtualHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualHost

`func (o *ClusterAMQPTargetConfig) SetVirtualHost(v string)`

SetVirtualHost sets VirtualHost field to given value.

### HasVirtualHost

`func (o *ClusterAMQPTargetConfig) HasVirtualHost() bool`

HasVirtualHost returns a boolean if a field has been set.

### SetVirtualHostNil

`func (o *ClusterAMQPTargetConfig) SetVirtualHostNil(b bool)`

 SetVirtualHostNil sets the value for VirtualHost to be an explicit nil

### UnsetVirtualHost
`func (o *ClusterAMQPTargetConfig) UnsetVirtualHost()`

UnsetVirtualHost ensures that no value is present for VirtualHost, not even an explicit nil
### GetExchange

`func (o *ClusterAMQPTargetConfig) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *ClusterAMQPTargetConfig) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *ClusterAMQPTargetConfig) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *ClusterAMQPTargetConfig) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### SetExchangeNil

`func (o *ClusterAMQPTargetConfig) SetExchangeNil(b bool)`

 SetExchangeNil sets the value for Exchange to be an explicit nil

### UnsetExchange
`func (o *ClusterAMQPTargetConfig) UnsetExchange()`

UnsetExchange ensures that no value is present for Exchange, not even an explicit nil
### GetFilerId

`func (o *ClusterAMQPTargetConfig) GetFilerId() int64`

GetFilerId returns the FilerId field if non-nil, zero value otherwise.

### GetFilerIdOk

`func (o *ClusterAMQPTargetConfig) GetFilerIdOk() (*int64, bool)`

GetFilerIdOk returns a tuple with the FilerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilerId

`func (o *ClusterAMQPTargetConfig) SetFilerId(v int64)`

SetFilerId sets FilerId field to given value.

### HasFilerId

`func (o *ClusterAMQPTargetConfig) HasFilerId() bool`

HasFilerId returns a boolean if a field has been set.

### SetFilerIdNil

`func (o *ClusterAMQPTargetConfig) SetFilerIdNil(b bool)`

 SetFilerIdNil sets the value for FilerId to be an explicit nil

### UnsetFilerId
`func (o *ClusterAMQPTargetConfig) UnsetFilerId()`

UnsetFilerId ensures that no value is present for FilerId, not even an explicit nil
### GetCertificate

`func (o *ClusterAMQPTargetConfig) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ClusterAMQPTargetConfig) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ClusterAMQPTargetConfig) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *ClusterAMQPTargetConfig) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *ClusterAMQPTargetConfig) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *ClusterAMQPTargetConfig) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


