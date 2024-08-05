# IPConfigParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interface** | Pointer to **NullableString** | Specifies the network interface name. IPs would be assigned to the specified interface. | [optional] 
**IpFamily** | Pointer to **NullableInt64** | Specifies the IP family of the config. | [optional] 
**Ips** | Pointer to **[]string** | Specifies a list of IP addresses to be assigned. | [optional] 
**NodeIds** | Pointer to **[]string** | Specifies the cluster node ids. | [optional] 
**Role** | Pointer to **NullableString** | Specifies the interface role. | [optional] 
**SubnetGateway** | Pointer to **NullableString** | Specifies the interface gateway. | [optional] 
**SubnetMaskBits** | Pointer to **NullableInt64** | Specifies the interface subnet mask bits. | [optional] 

## Methods

### NewIPConfigParams

`func NewIPConfigParams() *IPConfigParams`

NewIPConfigParams instantiates a new IPConfigParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPConfigParamsWithDefaults

`func NewIPConfigParamsWithDefaults() *IPConfigParams`

NewIPConfigParamsWithDefaults instantiates a new IPConfigParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterface

`func (o *IPConfigParams) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *IPConfigParams) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *IPConfigParams) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *IPConfigParams) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### SetInterfaceNil

`func (o *IPConfigParams) SetInterfaceNil(b bool)`

 SetInterfaceNil sets the value for Interface to be an explicit nil

### UnsetInterface
`func (o *IPConfigParams) UnsetInterface()`

UnsetInterface ensures that no value is present for Interface, not even an explicit nil
### GetIpFamily

`func (o *IPConfigParams) GetIpFamily() int64`

GetIpFamily returns the IpFamily field if non-nil, zero value otherwise.

### GetIpFamilyOk

`func (o *IPConfigParams) GetIpFamilyOk() (*int64, bool)`

GetIpFamilyOk returns a tuple with the IpFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFamily

`func (o *IPConfigParams) SetIpFamily(v int64)`

SetIpFamily sets IpFamily field to given value.

### HasIpFamily

`func (o *IPConfigParams) HasIpFamily() bool`

HasIpFamily returns a boolean if a field has been set.

### SetIpFamilyNil

`func (o *IPConfigParams) SetIpFamilyNil(b bool)`

 SetIpFamilyNil sets the value for IpFamily to be an explicit nil

### UnsetIpFamily
`func (o *IPConfigParams) UnsetIpFamily()`

UnsetIpFamily ensures that no value is present for IpFamily, not even an explicit nil
### GetIps

`func (o *IPConfigParams) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *IPConfigParams) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *IPConfigParams) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *IPConfigParams) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetNodeIds

`func (o *IPConfigParams) GetNodeIds() []string`

GetNodeIds returns the NodeIds field if non-nil, zero value otherwise.

### GetNodeIdsOk

`func (o *IPConfigParams) GetNodeIdsOk() (*[]string, bool)`

GetNodeIdsOk returns a tuple with the NodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIds

`func (o *IPConfigParams) SetNodeIds(v []string)`

SetNodeIds sets NodeIds field to given value.

### HasNodeIds

`func (o *IPConfigParams) HasNodeIds() bool`

HasNodeIds returns a boolean if a field has been set.

### GetRole

`func (o *IPConfigParams) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *IPConfigParams) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *IPConfigParams) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *IPConfigParams) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *IPConfigParams) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *IPConfigParams) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetSubnetGateway

`func (o *IPConfigParams) GetSubnetGateway() string`

GetSubnetGateway returns the SubnetGateway field if non-nil, zero value otherwise.

### GetSubnetGatewayOk

`func (o *IPConfigParams) GetSubnetGatewayOk() (*string, bool)`

GetSubnetGatewayOk returns a tuple with the SubnetGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetGateway

`func (o *IPConfigParams) SetSubnetGateway(v string)`

SetSubnetGateway sets SubnetGateway field to given value.

### HasSubnetGateway

`func (o *IPConfigParams) HasSubnetGateway() bool`

HasSubnetGateway returns a boolean if a field has been set.

### SetSubnetGatewayNil

`func (o *IPConfigParams) SetSubnetGatewayNil(b bool)`

 SetSubnetGatewayNil sets the value for SubnetGateway to be an explicit nil

### UnsetSubnetGateway
`func (o *IPConfigParams) UnsetSubnetGateway()`

UnsetSubnetGateway ensures that no value is present for SubnetGateway, not even an explicit nil
### GetSubnetMaskBits

`func (o *IPConfigParams) GetSubnetMaskBits() int64`

GetSubnetMaskBits returns the SubnetMaskBits field if non-nil, zero value otherwise.

### GetSubnetMaskBitsOk

`func (o *IPConfigParams) GetSubnetMaskBitsOk() (*int64, bool)`

GetSubnetMaskBitsOk returns a tuple with the SubnetMaskBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMaskBits

`func (o *IPConfigParams) SetSubnetMaskBits(v int64)`

SetSubnetMaskBits sets SubnetMaskBits field to given value.

### HasSubnetMaskBits

`func (o *IPConfigParams) HasSubnetMaskBits() bool`

HasSubnetMaskBits returns a boolean if a field has been set.

### SetSubnetMaskBitsNil

`func (o *IPConfigParams) SetSubnetMaskBitsNil(b bool)`

 SetSubnetMaskBitsNil sets the value for SubnetMaskBits to be an explicit nil

### UnsetSubnetMaskBits
`func (o *IPConfigParams) UnsetSubnetMaskBits()`

UnsetSubnetMaskBits ensures that no value is present for SubnetMaskBits, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


