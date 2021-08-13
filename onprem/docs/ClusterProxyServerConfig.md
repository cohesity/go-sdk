# ClusterProxyServerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies the unique name of the proxy server. | [optional] [readonly] 
**Ip** | **NullableString** | Specifies the IP address of the proxy server. | 
**Port** | **NullableInt32** | Specifies the port on which the server is listening. | 
**Username** | Pointer to **NullableString** | Specifies the username for the proxy. | [optional] 
**Password** | Pointer to **NullableString** | Specifies the password for the proxy. | [optional] 

## Methods

### NewClusterProxyServerConfig

`func NewClusterProxyServerConfig(ip NullableString, port NullableInt32, ) *ClusterProxyServerConfig`

NewClusterProxyServerConfig instantiates a new ClusterProxyServerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterProxyServerConfigWithDefaults

`func NewClusterProxyServerConfigWithDefaults() *ClusterProxyServerConfig`

NewClusterProxyServerConfigWithDefaults instantiates a new ClusterProxyServerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterProxyServerConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterProxyServerConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterProxyServerConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterProxyServerConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ClusterProxyServerConfig) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ClusterProxyServerConfig) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIp

`func (o *ClusterProxyServerConfig) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ClusterProxyServerConfig) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ClusterProxyServerConfig) SetIp(v string)`

SetIp sets Ip field to given value.


### SetIpNil

`func (o *ClusterProxyServerConfig) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *ClusterProxyServerConfig) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetPort

`func (o *ClusterProxyServerConfig) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ClusterProxyServerConfig) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ClusterProxyServerConfig) SetPort(v int32)`

SetPort sets Port field to given value.


### SetPortNil

`func (o *ClusterProxyServerConfig) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *ClusterProxyServerConfig) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetUsername

`func (o *ClusterProxyServerConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ClusterProxyServerConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ClusterProxyServerConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ClusterProxyServerConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ClusterProxyServerConfig) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ClusterProxyServerConfig) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *ClusterProxyServerConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ClusterProxyServerConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ClusterProxyServerConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ClusterProxyServerConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ClusterProxyServerConfig) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ClusterProxyServerConfig) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


