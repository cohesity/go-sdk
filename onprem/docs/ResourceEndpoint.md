# ResourceEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdn** | Pointer to **NullableString** | Specifies the fqdn of this endpoint. | [optional] 
**Ipv4addr** | Pointer to **NullableString** | Specifies the ipv4 address of this endpoint. | [optional] 
**Ipv6addr** | Pointer to **NullableString** | Specifies the ipv6 address of this endpoint. | [optional] 
**SubnetIp4addr** | Pointer to **NullableString** | Specifies the subnet Ip4 address of this endpoint. | [optional] 
**PreferredAddress** | Pointer to **NullableString** | Specifies the preferred address to use for connecting. | [optional] 
**IsPreferredEndpoint** | Pointer to **NullableBool** | Whether to use this endpoint to connect. | [optional] 

## Methods

### NewResourceEndpoint

`func NewResourceEndpoint() *ResourceEndpoint`

NewResourceEndpoint instantiates a new ResourceEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceEndpointWithDefaults

`func NewResourceEndpointWithDefaults() *ResourceEndpoint`

NewResourceEndpointWithDefaults instantiates a new ResourceEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqdn

`func (o *ResourceEndpoint) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *ResourceEndpoint) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *ResourceEndpoint) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *ResourceEndpoint) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### SetFqdnNil

`func (o *ResourceEndpoint) SetFqdnNil(b bool)`

 SetFqdnNil sets the value for Fqdn to be an explicit nil

### UnsetFqdn
`func (o *ResourceEndpoint) UnsetFqdn()`

UnsetFqdn ensures that no value is present for Fqdn, not even an explicit nil
### GetIpv4addr

`func (o *ResourceEndpoint) GetIpv4addr() string`

GetIpv4addr returns the Ipv4addr field if non-nil, zero value otherwise.

### GetIpv4addrOk

`func (o *ResourceEndpoint) GetIpv4addrOk() (*string, bool)`

GetIpv4addrOk returns a tuple with the Ipv4addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4addr

`func (o *ResourceEndpoint) SetIpv4addr(v string)`

SetIpv4addr sets Ipv4addr field to given value.

### HasIpv4addr

`func (o *ResourceEndpoint) HasIpv4addr() bool`

HasIpv4addr returns a boolean if a field has been set.

### SetIpv4addrNil

`func (o *ResourceEndpoint) SetIpv4addrNil(b bool)`

 SetIpv4addrNil sets the value for Ipv4addr to be an explicit nil

### UnsetIpv4addr
`func (o *ResourceEndpoint) UnsetIpv4addr()`

UnsetIpv4addr ensures that no value is present for Ipv4addr, not even an explicit nil
### GetIpv6addr

`func (o *ResourceEndpoint) GetIpv6addr() string`

GetIpv6addr returns the Ipv6addr field if non-nil, zero value otherwise.

### GetIpv6addrOk

`func (o *ResourceEndpoint) GetIpv6addrOk() (*string, bool)`

GetIpv6addrOk returns a tuple with the Ipv6addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6addr

`func (o *ResourceEndpoint) SetIpv6addr(v string)`

SetIpv6addr sets Ipv6addr field to given value.

### HasIpv6addr

`func (o *ResourceEndpoint) HasIpv6addr() bool`

HasIpv6addr returns a boolean if a field has been set.

### SetIpv6addrNil

`func (o *ResourceEndpoint) SetIpv6addrNil(b bool)`

 SetIpv6addrNil sets the value for Ipv6addr to be an explicit nil

### UnsetIpv6addr
`func (o *ResourceEndpoint) UnsetIpv6addr()`

UnsetIpv6addr ensures that no value is present for Ipv6addr, not even an explicit nil
### GetSubnetIp4addr

`func (o *ResourceEndpoint) GetSubnetIp4addr() string`

GetSubnetIp4addr returns the SubnetIp4addr field if non-nil, zero value otherwise.

### GetSubnetIp4addrOk

`func (o *ResourceEndpoint) GetSubnetIp4addrOk() (*string, bool)`

GetSubnetIp4addrOk returns a tuple with the SubnetIp4addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIp4addr

`func (o *ResourceEndpoint) SetSubnetIp4addr(v string)`

SetSubnetIp4addr sets SubnetIp4addr field to given value.

### HasSubnetIp4addr

`func (o *ResourceEndpoint) HasSubnetIp4addr() bool`

HasSubnetIp4addr returns a boolean if a field has been set.

### SetSubnetIp4addrNil

`func (o *ResourceEndpoint) SetSubnetIp4addrNil(b bool)`

 SetSubnetIp4addrNil sets the value for SubnetIp4addr to be an explicit nil

### UnsetSubnetIp4addr
`func (o *ResourceEndpoint) UnsetSubnetIp4addr()`

UnsetSubnetIp4addr ensures that no value is present for SubnetIp4addr, not even an explicit nil
### GetPreferredAddress

`func (o *ResourceEndpoint) GetPreferredAddress() string`

GetPreferredAddress returns the PreferredAddress field if non-nil, zero value otherwise.

### GetPreferredAddressOk

`func (o *ResourceEndpoint) GetPreferredAddressOk() (*string, bool)`

GetPreferredAddressOk returns a tuple with the PreferredAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredAddress

`func (o *ResourceEndpoint) SetPreferredAddress(v string)`

SetPreferredAddress sets PreferredAddress field to given value.

### HasPreferredAddress

`func (o *ResourceEndpoint) HasPreferredAddress() bool`

HasPreferredAddress returns a boolean if a field has been set.

### SetPreferredAddressNil

`func (o *ResourceEndpoint) SetPreferredAddressNil(b bool)`

 SetPreferredAddressNil sets the value for PreferredAddress to be an explicit nil

### UnsetPreferredAddress
`func (o *ResourceEndpoint) UnsetPreferredAddress()`

UnsetPreferredAddress ensures that no value is present for PreferredAddress, not even an explicit nil
### GetIsPreferredEndpoint

`func (o *ResourceEndpoint) GetIsPreferredEndpoint() bool`

GetIsPreferredEndpoint returns the IsPreferredEndpoint field if non-nil, zero value otherwise.

### GetIsPreferredEndpointOk

`func (o *ResourceEndpoint) GetIsPreferredEndpointOk() (*bool, bool)`

GetIsPreferredEndpointOk returns a tuple with the IsPreferredEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreferredEndpoint

`func (o *ResourceEndpoint) SetIsPreferredEndpoint(v bool)`

SetIsPreferredEndpoint sets IsPreferredEndpoint field to given value.

### HasIsPreferredEndpoint

`func (o *ResourceEndpoint) HasIsPreferredEndpoint() bool`

HasIsPreferredEndpoint returns a boolean if a field has been set.

### SetIsPreferredEndpointNil

`func (o *ResourceEndpoint) SetIsPreferredEndpointNil(b bool)`

 SetIsPreferredEndpointNil sets the value for IsPreferredEndpoint to be an explicit nil

### UnsetIsPreferredEndpoint
`func (o *ResourceEndpoint) UnsetIsPreferredEndpoint()`

UnsetIsPreferredEndpoint ensures that no value is present for IsPreferredEndpoint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


