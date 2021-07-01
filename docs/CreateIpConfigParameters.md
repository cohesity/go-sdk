# CreateIpConfigParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ips** | Pointer to **[]string** | Specifies the interface ips. | [optional] 
**Mtu** | Pointer to **NullableInt32** | Specifies the interface mtu. | [optional] 
**Name** | **NullableString** | Specifies the interface name. | 
**Role** | Pointer to **NullableString** | Specifies the interface role. &#39;kPrimary&#39; indicates a primary role. &#39;kSecondary&#39; indicates a secondary role. | [optional] 
**SubnetGateway** | Pointer to **NullableString** | Specifies the interface gateway. | [optional] 
**SubnetMask** | Pointer to **NullableString** | Specifies the interface subnet mask. | [optional] 

## Methods

### NewCreateIpConfigParameters

`func NewCreateIpConfigParameters(name NullableString, ) *CreateIpConfigParameters`

NewCreateIpConfigParameters instantiates a new CreateIpConfigParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIpConfigParametersWithDefaults

`func NewCreateIpConfigParametersWithDefaults() *CreateIpConfigParameters`

NewCreateIpConfigParametersWithDefaults instantiates a new CreateIpConfigParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIps

`func (o *CreateIpConfigParameters) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *CreateIpConfigParameters) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *CreateIpConfigParameters) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *CreateIpConfigParameters) HasIps() bool`

HasIps returns a boolean if a field has been set.

### SetIpsNil

`func (o *CreateIpConfigParameters) SetIpsNil(b bool)`

 SetIpsNil sets the value for Ips to be an explicit nil

### UnsetIps
`func (o *CreateIpConfigParameters) UnsetIps()`

UnsetIps ensures that no value is present for Ips, not even an explicit nil
### GetMtu

`func (o *CreateIpConfigParameters) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *CreateIpConfigParameters) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *CreateIpConfigParameters) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *CreateIpConfigParameters) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *CreateIpConfigParameters) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *CreateIpConfigParameters) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetName

`func (o *CreateIpConfigParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateIpConfigParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateIpConfigParameters) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateIpConfigParameters) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateIpConfigParameters) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRole

`func (o *CreateIpConfigParameters) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CreateIpConfigParameters) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CreateIpConfigParameters) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *CreateIpConfigParameters) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *CreateIpConfigParameters) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *CreateIpConfigParameters) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetSubnetGateway

`func (o *CreateIpConfigParameters) GetSubnetGateway() string`

GetSubnetGateway returns the SubnetGateway field if non-nil, zero value otherwise.

### GetSubnetGatewayOk

`func (o *CreateIpConfigParameters) GetSubnetGatewayOk() (*string, bool)`

GetSubnetGatewayOk returns a tuple with the SubnetGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetGateway

`func (o *CreateIpConfigParameters) SetSubnetGateway(v string)`

SetSubnetGateway sets SubnetGateway field to given value.

### HasSubnetGateway

`func (o *CreateIpConfigParameters) HasSubnetGateway() bool`

HasSubnetGateway returns a boolean if a field has been set.

### SetSubnetGatewayNil

`func (o *CreateIpConfigParameters) SetSubnetGatewayNil(b bool)`

 SetSubnetGatewayNil sets the value for SubnetGateway to be an explicit nil

### UnsetSubnetGateway
`func (o *CreateIpConfigParameters) UnsetSubnetGateway()`

UnsetSubnetGateway ensures that no value is present for SubnetGateway, not even an explicit nil
### GetSubnetMask

`func (o *CreateIpConfigParameters) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *CreateIpConfigParameters) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *CreateIpConfigParameters) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.

### HasSubnetMask

`func (o *CreateIpConfigParameters) HasSubnetMask() bool`

HasSubnetMask returns a boolean if a field has been set.

### SetSubnetMaskNil

`func (o *CreateIpConfigParameters) SetSubnetMaskNil(b bool)`

 SetSubnetMaskNil sets the value for SubnetMask to be an explicit nil

### UnsetSubnetMask
`func (o *CreateIpConfigParameters) UnsetSubnetMask()`

UnsetSubnetMask ensures that no value is present for SubnetMask, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


