# Interface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BondMembers** | Pointer to [**[]BondMember**](BondMember.md) | Bond member details for bond interface. | [optional] 
**BondingMode** | Pointer to **NullableString** | Bonding mode if this interface is a bond. | [optional] 
**DefaultRoute** | Pointer to **NullableBool** | Specifies whether or not this interface is the default route. | [optional] 
**Gateway** | Pointer to **NullableString** | Gateway of the interface. | [optional] 
**Group** | Pointer to **NullableString** | Group to which this interface belongs. | [optional] 
**Id** | Pointer to **NullableInt64** | Id of the interface. | [optional] 
**Ipv6Gateway** | Pointer to **NullableString** | The IPv6 gateway of the interface. | [optional] 
**Ipv6Static** | Pointer to **NullableString** | Static IPv6 of the interface. | [optional] 
**Ipv6Subnet** | Pointer to **NullableString** | The IPv6 subnet of the interface. | [optional] 
**IsConnected** | Pointer to **NullableBool** | Specifies whether or not this interface is connected. | [optional] 
**IsUp** | Pointer to **NullableBool** | Specifies whether or not the interface is up. | [optional] 
**MacAddress** | Pointer to **NullableString** | MAC address of the interface. | [optional] 
**Mtu** | Pointer to **NullableInt32** | MTU of the interface. | [optional] 
**Name** | Pointer to **NullableString** | The name of the interface. | [optional] 
**Role** | Pointer to **NullableString** | Role of the interface. | [optional] 
**Services** | Pointer to **[]string** | Types of services this interface is used for. | [optional] 
**Speed** | Pointer to **NullableString** | Speed of the interface. | [optional] 
**StaticIp** | Pointer to **NullableString** | Static IP of the interface. | [optional] 
**Stats** | Pointer to [**InterfaceStats**](InterfaceStats.md) |  | [optional] 
**Subnet** | Pointer to **NullableString** | Subnet of the interface. | [optional] 
**Type** | Pointer to **NullableString** | The type of the interface. | [optional] 
**VirtualIp** | Pointer to **NullableString** | Virtual IP of the interface. | [optional] 

## Methods

### NewInterface

`func NewInterface() *Interface`

NewInterface instantiates a new Interface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceWithDefaults

`func NewInterfaceWithDefaults() *Interface`

NewInterfaceWithDefaults instantiates a new Interface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBondMembers

`func (o *Interface) GetBondMembers() []BondMember`

GetBondMembers returns the BondMembers field if non-nil, zero value otherwise.

### GetBondMembersOk

`func (o *Interface) GetBondMembersOk() (*[]BondMember, bool)`

GetBondMembersOk returns a tuple with the BondMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondMembers

`func (o *Interface) SetBondMembers(v []BondMember)`

SetBondMembers sets BondMembers field to given value.

### HasBondMembers

`func (o *Interface) HasBondMembers() bool`

HasBondMembers returns a boolean if a field has been set.

### SetBondMembersNil

`func (o *Interface) SetBondMembersNil(b bool)`

 SetBondMembersNil sets the value for BondMembers to be an explicit nil

### UnsetBondMembers
`func (o *Interface) UnsetBondMembers()`

UnsetBondMembers ensures that no value is present for BondMembers, not even an explicit nil
### GetBondingMode

`func (o *Interface) GetBondingMode() string`

GetBondingMode returns the BondingMode field if non-nil, zero value otherwise.

### GetBondingModeOk

`func (o *Interface) GetBondingModeOk() (*string, bool)`

GetBondingModeOk returns a tuple with the BondingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondingMode

`func (o *Interface) SetBondingMode(v string)`

SetBondingMode sets BondingMode field to given value.

### HasBondingMode

`func (o *Interface) HasBondingMode() bool`

HasBondingMode returns a boolean if a field has been set.

### SetBondingModeNil

`func (o *Interface) SetBondingModeNil(b bool)`

 SetBondingModeNil sets the value for BondingMode to be an explicit nil

### UnsetBondingMode
`func (o *Interface) UnsetBondingMode()`

UnsetBondingMode ensures that no value is present for BondingMode, not even an explicit nil
### GetDefaultRoute

`func (o *Interface) GetDefaultRoute() bool`

GetDefaultRoute returns the DefaultRoute field if non-nil, zero value otherwise.

### GetDefaultRouteOk

`func (o *Interface) GetDefaultRouteOk() (*bool, bool)`

GetDefaultRouteOk returns a tuple with the DefaultRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoute

`func (o *Interface) SetDefaultRoute(v bool)`

SetDefaultRoute sets DefaultRoute field to given value.

### HasDefaultRoute

`func (o *Interface) HasDefaultRoute() bool`

HasDefaultRoute returns a boolean if a field has been set.

### SetDefaultRouteNil

`func (o *Interface) SetDefaultRouteNil(b bool)`

 SetDefaultRouteNil sets the value for DefaultRoute to be an explicit nil

### UnsetDefaultRoute
`func (o *Interface) UnsetDefaultRoute()`

UnsetDefaultRoute ensures that no value is present for DefaultRoute, not even an explicit nil
### GetGateway

`func (o *Interface) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *Interface) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *Interface) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *Interface) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *Interface) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *Interface) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetGroup

`func (o *Interface) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Interface) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Interface) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Interface) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *Interface) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *Interface) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetId

`func (o *Interface) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Interface) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Interface) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Interface) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Interface) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Interface) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIpv6Gateway

`func (o *Interface) GetIpv6Gateway() string`

GetIpv6Gateway returns the Ipv6Gateway field if non-nil, zero value otherwise.

### GetIpv6GatewayOk

`func (o *Interface) GetIpv6GatewayOk() (*string, bool)`

GetIpv6GatewayOk returns a tuple with the Ipv6Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Gateway

`func (o *Interface) SetIpv6Gateway(v string)`

SetIpv6Gateway sets Ipv6Gateway field to given value.

### HasIpv6Gateway

`func (o *Interface) HasIpv6Gateway() bool`

HasIpv6Gateway returns a boolean if a field has been set.

### SetIpv6GatewayNil

`func (o *Interface) SetIpv6GatewayNil(b bool)`

 SetIpv6GatewayNil sets the value for Ipv6Gateway to be an explicit nil

### UnsetIpv6Gateway
`func (o *Interface) UnsetIpv6Gateway()`

UnsetIpv6Gateway ensures that no value is present for Ipv6Gateway, not even an explicit nil
### GetIpv6Static

`func (o *Interface) GetIpv6Static() string`

GetIpv6Static returns the Ipv6Static field if non-nil, zero value otherwise.

### GetIpv6StaticOk

`func (o *Interface) GetIpv6StaticOk() (*string, bool)`

GetIpv6StaticOk returns a tuple with the Ipv6Static field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Static

`func (o *Interface) SetIpv6Static(v string)`

SetIpv6Static sets Ipv6Static field to given value.

### HasIpv6Static

`func (o *Interface) HasIpv6Static() bool`

HasIpv6Static returns a boolean if a field has been set.

### SetIpv6StaticNil

`func (o *Interface) SetIpv6StaticNil(b bool)`

 SetIpv6StaticNil sets the value for Ipv6Static to be an explicit nil

### UnsetIpv6Static
`func (o *Interface) UnsetIpv6Static()`

UnsetIpv6Static ensures that no value is present for Ipv6Static, not even an explicit nil
### GetIpv6Subnet

`func (o *Interface) GetIpv6Subnet() string`

GetIpv6Subnet returns the Ipv6Subnet field if non-nil, zero value otherwise.

### GetIpv6SubnetOk

`func (o *Interface) GetIpv6SubnetOk() (*string, bool)`

GetIpv6SubnetOk returns a tuple with the Ipv6Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Subnet

`func (o *Interface) SetIpv6Subnet(v string)`

SetIpv6Subnet sets Ipv6Subnet field to given value.

### HasIpv6Subnet

`func (o *Interface) HasIpv6Subnet() bool`

HasIpv6Subnet returns a boolean if a field has been set.

### SetIpv6SubnetNil

`func (o *Interface) SetIpv6SubnetNil(b bool)`

 SetIpv6SubnetNil sets the value for Ipv6Subnet to be an explicit nil

### UnsetIpv6Subnet
`func (o *Interface) UnsetIpv6Subnet()`

UnsetIpv6Subnet ensures that no value is present for Ipv6Subnet, not even an explicit nil
### GetIsConnected

`func (o *Interface) GetIsConnected() bool`

GetIsConnected returns the IsConnected field if non-nil, zero value otherwise.

### GetIsConnectedOk

`func (o *Interface) GetIsConnectedOk() (*bool, bool)`

GetIsConnectedOk returns a tuple with the IsConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConnected

`func (o *Interface) SetIsConnected(v bool)`

SetIsConnected sets IsConnected field to given value.

### HasIsConnected

`func (o *Interface) HasIsConnected() bool`

HasIsConnected returns a boolean if a field has been set.

### SetIsConnectedNil

`func (o *Interface) SetIsConnectedNil(b bool)`

 SetIsConnectedNil sets the value for IsConnected to be an explicit nil

### UnsetIsConnected
`func (o *Interface) UnsetIsConnected()`

UnsetIsConnected ensures that no value is present for IsConnected, not even an explicit nil
### GetIsUp

`func (o *Interface) GetIsUp() bool`

GetIsUp returns the IsUp field if non-nil, zero value otherwise.

### GetIsUpOk

`func (o *Interface) GetIsUpOk() (*bool, bool)`

GetIsUpOk returns a tuple with the IsUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUp

`func (o *Interface) SetIsUp(v bool)`

SetIsUp sets IsUp field to given value.

### HasIsUp

`func (o *Interface) HasIsUp() bool`

HasIsUp returns a boolean if a field has been set.

### SetIsUpNil

`func (o *Interface) SetIsUpNil(b bool)`

 SetIsUpNil sets the value for IsUp to be an explicit nil

### UnsetIsUp
`func (o *Interface) UnsetIsUp()`

UnsetIsUp ensures that no value is present for IsUp, not even an explicit nil
### GetMacAddress

`func (o *Interface) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *Interface) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *Interface) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *Interface) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *Interface) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *Interface) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
### GetMtu

`func (o *Interface) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *Interface) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *Interface) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *Interface) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *Interface) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *Interface) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetName

`func (o *Interface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Interface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Interface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Interface) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Interface) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Interface) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRole

`func (o *Interface) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Interface) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Interface) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *Interface) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *Interface) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *Interface) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetServices

`func (o *Interface) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *Interface) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *Interface) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *Interface) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *Interface) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *Interface) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil
### GetSpeed

`func (o *Interface) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *Interface) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *Interface) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *Interface) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### SetSpeedNil

`func (o *Interface) SetSpeedNil(b bool)`

 SetSpeedNil sets the value for Speed to be an explicit nil

### UnsetSpeed
`func (o *Interface) UnsetSpeed()`

UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
### GetStaticIp

`func (o *Interface) GetStaticIp() string`

GetStaticIp returns the StaticIp field if non-nil, zero value otherwise.

### GetStaticIpOk

`func (o *Interface) GetStaticIpOk() (*string, bool)`

GetStaticIpOk returns a tuple with the StaticIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIp

`func (o *Interface) SetStaticIp(v string)`

SetStaticIp sets StaticIp field to given value.

### HasStaticIp

`func (o *Interface) HasStaticIp() bool`

HasStaticIp returns a boolean if a field has been set.

### SetStaticIpNil

`func (o *Interface) SetStaticIpNil(b bool)`

 SetStaticIpNil sets the value for StaticIp to be an explicit nil

### UnsetStaticIp
`func (o *Interface) UnsetStaticIp()`

UnsetStaticIp ensures that no value is present for StaticIp, not even an explicit nil
### GetStats

`func (o *Interface) GetStats() InterfaceStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Interface) GetStatsOk() (*InterfaceStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Interface) SetStats(v InterfaceStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *Interface) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetSubnet

`func (o *Interface) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *Interface) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *Interface) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *Interface) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### SetSubnetNil

`func (o *Interface) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *Interface) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetType

`func (o *Interface) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Interface) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Interface) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Interface) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Interface) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Interface) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetVirtualIp

`func (o *Interface) GetVirtualIp() string`

GetVirtualIp returns the VirtualIp field if non-nil, zero value otherwise.

### GetVirtualIpOk

`func (o *Interface) GetVirtualIpOk() (*string, bool)`

GetVirtualIpOk returns a tuple with the VirtualIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualIp

`func (o *Interface) SetVirtualIp(v string)`

SetVirtualIp sets VirtualIp field to given value.

### HasVirtualIp

`func (o *Interface) HasVirtualIp() bool`

HasVirtualIp returns a boolean if a field has been set.

### SetVirtualIpNil

`func (o *Interface) SetVirtualIpNil(b bool)`

 SetVirtualIpNil sets the value for VirtualIp to be an explicit nil

### UnsetVirtualIp
`func (o *Interface) UnsetVirtualIp()`

UnsetVirtualIp ensures that no value is present for VirtualIp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


