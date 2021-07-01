# IscsiSanPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | Pointer to **NullableString** | Specifies the IP address of the SAN port. | [optional] 
**Iqn** | Pointer to **NullableString** | Specifies the qualified name of the iSCSI port (IQN). | [optional] 
**TcpPort** | Pointer to **NullableInt32** | Specifies the listening port(tcp) of the SAN port. | [optional] 

## Methods

### NewIscsiSanPort

`func NewIscsiSanPort() *IscsiSanPort`

NewIscsiSanPort instantiates a new IscsiSanPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIscsiSanPortWithDefaults

`func NewIscsiSanPortWithDefaults() *IscsiSanPort`

NewIscsiSanPortWithDefaults instantiates a new IscsiSanPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *IscsiSanPort) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *IscsiSanPort) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *IscsiSanPort) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *IscsiSanPort) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *IscsiSanPort) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *IscsiSanPort) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetIqn

`func (o *IscsiSanPort) GetIqn() string`

GetIqn returns the Iqn field if non-nil, zero value otherwise.

### GetIqnOk

`func (o *IscsiSanPort) GetIqnOk() (*string, bool)`

GetIqnOk returns a tuple with the Iqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqn

`func (o *IscsiSanPort) SetIqn(v string)`

SetIqn sets Iqn field to given value.

### HasIqn

`func (o *IscsiSanPort) HasIqn() bool`

HasIqn returns a boolean if a field has been set.

### SetIqnNil

`func (o *IscsiSanPort) SetIqnNil(b bool)`

 SetIqnNil sets the value for Iqn to be an explicit nil

### UnsetIqn
`func (o *IscsiSanPort) UnsetIqn()`

UnsetIqn ensures that no value is present for Iqn, not even an explicit nil
### GetTcpPort

`func (o *IscsiSanPort) GetTcpPort() int32`

GetTcpPort returns the TcpPort field if non-nil, zero value otherwise.

### GetTcpPortOk

`func (o *IscsiSanPort) GetTcpPortOk() (*int32, bool)`

GetTcpPortOk returns a tuple with the TcpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpPort

`func (o *IscsiSanPort) SetTcpPort(v int32)`

SetTcpPort sets TcpPort field to given value.

### HasTcpPort

`func (o *IscsiSanPort) HasTcpPort() bool`

HasTcpPort returns a boolean if a field has been set.

### SetTcpPortNil

`func (o *IscsiSanPort) SetTcpPortNil(b bool)`

 SetTcpPortNil sets the value for TcpPort to be an explicit nil

### UnsetTcpPort
`func (o *IscsiSanPort) UnsetTcpPort()`

UnsetTcpPort ensures that no value is present for TcpPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


