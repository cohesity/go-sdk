# IpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceName** | Pointer to **NullableString** | The interface name.  Specifies which interface to assign IP to. | [optional] 
**IpFamily** | Pointer to **NullableInt32** | IpFamily of this config. | [optional] 
**Ips** | Pointer to **[]string** | The interface ips. | [optional] 
**NodeIds** | Pointer to **[]int64** | Node ids. | [optional] 
**Role** | Pointer to **NullableString** | The interface role. | [optional] 
**SubnetGateway** | Pointer to **NullableString** | The interface gateway. | [optional] 
**SubnetMaskBits** | Pointer to **NullableInt32** | The interface subnet mask bits. | [optional] 

## Methods

### NewIpConfig

`func NewIpConfig() *IpConfig`

NewIpConfig instantiates a new IpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpConfigWithDefaults

`func NewIpConfigWithDefaults() *IpConfig`

NewIpConfigWithDefaults instantiates a new IpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceName

`func (o *IpConfig) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *IpConfig) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *IpConfig) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *IpConfig) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### SetInterfaceNameNil

`func (o *IpConfig) SetInterfaceNameNil(b bool)`

 SetInterfaceNameNil sets the value for InterfaceName to be an explicit nil

### UnsetInterfaceName
`func (o *IpConfig) UnsetInterfaceName()`

UnsetInterfaceName ensures that no value is present for InterfaceName, not even an explicit nil
### GetIpFamily

`func (o *IpConfig) GetIpFamily() int32`

GetIpFamily returns the IpFamily field if non-nil, zero value otherwise.

### GetIpFamilyOk

`func (o *IpConfig) GetIpFamilyOk() (*int32, bool)`

GetIpFamilyOk returns a tuple with the IpFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFamily

`func (o *IpConfig) SetIpFamily(v int32)`

SetIpFamily sets IpFamily field to given value.

### HasIpFamily

`func (o *IpConfig) HasIpFamily() bool`

HasIpFamily returns a boolean if a field has been set.

### SetIpFamilyNil

`func (o *IpConfig) SetIpFamilyNil(b bool)`

 SetIpFamilyNil sets the value for IpFamily to be an explicit nil

### UnsetIpFamily
`func (o *IpConfig) UnsetIpFamily()`

UnsetIpFamily ensures that no value is present for IpFamily, not even an explicit nil
### GetIps

`func (o *IpConfig) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *IpConfig) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *IpConfig) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *IpConfig) HasIps() bool`

HasIps returns a boolean if a field has been set.

### SetIpsNil

`func (o *IpConfig) SetIpsNil(b bool)`

 SetIpsNil sets the value for Ips to be an explicit nil

### UnsetIps
`func (o *IpConfig) UnsetIps()`

UnsetIps ensures that no value is present for Ips, not even an explicit nil
### GetNodeIds

`func (o *IpConfig) GetNodeIds() []int64`

GetNodeIds returns the NodeIds field if non-nil, zero value otherwise.

### GetNodeIdsOk

`func (o *IpConfig) GetNodeIdsOk() (*[]int64, bool)`

GetNodeIdsOk returns a tuple with the NodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIds

`func (o *IpConfig) SetNodeIds(v []int64)`

SetNodeIds sets NodeIds field to given value.

### HasNodeIds

`func (o *IpConfig) HasNodeIds() bool`

HasNodeIds returns a boolean if a field has been set.

### SetNodeIdsNil

`func (o *IpConfig) SetNodeIdsNil(b bool)`

 SetNodeIdsNil sets the value for NodeIds to be an explicit nil

### UnsetNodeIds
`func (o *IpConfig) UnsetNodeIds()`

UnsetNodeIds ensures that no value is present for NodeIds, not even an explicit nil
### GetRole

`func (o *IpConfig) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *IpConfig) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *IpConfig) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *IpConfig) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *IpConfig) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *IpConfig) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetSubnetGateway

`func (o *IpConfig) GetSubnetGateway() string`

GetSubnetGateway returns the SubnetGateway field if non-nil, zero value otherwise.

### GetSubnetGatewayOk

`func (o *IpConfig) GetSubnetGatewayOk() (*string, bool)`

GetSubnetGatewayOk returns a tuple with the SubnetGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetGateway

`func (o *IpConfig) SetSubnetGateway(v string)`

SetSubnetGateway sets SubnetGateway field to given value.

### HasSubnetGateway

`func (o *IpConfig) HasSubnetGateway() bool`

HasSubnetGateway returns a boolean if a field has been set.

### SetSubnetGatewayNil

`func (o *IpConfig) SetSubnetGatewayNil(b bool)`

 SetSubnetGatewayNil sets the value for SubnetGateway to be an explicit nil

### UnsetSubnetGateway
`func (o *IpConfig) UnsetSubnetGateway()`

UnsetSubnetGateway ensures that no value is present for SubnetGateway, not even an explicit nil
### GetSubnetMaskBits

`func (o *IpConfig) GetSubnetMaskBits() int32`

GetSubnetMaskBits returns the SubnetMaskBits field if non-nil, zero value otherwise.

### GetSubnetMaskBitsOk

`func (o *IpConfig) GetSubnetMaskBitsOk() (*int32, bool)`

GetSubnetMaskBitsOk returns a tuple with the SubnetMaskBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMaskBits

`func (o *IpConfig) SetSubnetMaskBits(v int32)`

SetSubnetMaskBits sets SubnetMaskBits field to given value.

### HasSubnetMaskBits

`func (o *IpConfig) HasSubnetMaskBits() bool`

HasSubnetMaskBits returns a boolean if a field has been set.

### SetSubnetMaskBitsNil

`func (o *IpConfig) SetSubnetMaskBitsNil(b bool)`

 SetSubnetMaskBitsNil sets the value for SubnetMaskBits to be an explicit nil

### UnsetSubnetMaskBits
`func (o *IpConfig) UnsetSubnetMaskBits()`

UnsetSubnetMaskBits ensures that no value is present for SubnetMaskBits, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


