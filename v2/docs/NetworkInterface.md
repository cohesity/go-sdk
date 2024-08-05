# NetworkInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveBondSlave** | Pointer to **NullableString** | Current active slave. This is only valid in active-backup mode. | [optional] 
**BondSlaveNames** | Pointer to **[]string** | Specifies the names of the bond slaves for this interface. | [optional] 
**BondSlaveSlots** | Pointer to **[]string** | Specifies the slots of the bond slaves for this interface. | [optional] 
**BondSlavesDetails** | Pointer to [**[]BondMember**](BondMember.md) | Bond member details for bond interface. | [optional] 
**BondingMode** | Pointer to **NullableString** | Specifies the bonding mode of this interface. | [optional] 
**DefaultRoute** | Pointer to **NullableBool** | Specifies whether or not this interface is the default route. | [optional] 
**Gateway** | Pointer to **NullableString** | Specifies the gateway of the network interface. | [optional] 
**GatewayV6** | Pointer to **NullableString** | Specifies the gatewayV6 of the network interface. | [optional] 
**Group** | Pointer to **NullableString** | Specifies the group to which this interface belongs. | [optional] 
**IsConnected** | Pointer to **NullableBool** | Specifies whether or not this interface is connected. | [optional] 
**IsUp** | Pointer to **NullableBool** | Specifies whether or not the interface is up. | [optional] 
**MacAddress** | Pointer to **NullableString** | Specifies the MAC address of this interface. | [optional] 
**Mtu** | Pointer to **NullableInt32** | Specifies the MTU of the network interface. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the network interface. | [optional] 
**Role** | Pointer to **NullableString** | Specifies the interface role. | [optional] 
**Services** | Pointer to **[]string** | Services which use this interface. | [optional] 
**Speed** | Pointer to **NullableString** | Specifies the speed of this interface. | [optional] 
**StaticIP** | Pointer to **NullableString** | Specifies the static IP of the network interface. | [optional] 
**StaticIpV6** | Pointer to **NullableString** | Specifies the static IPV6 of the network interface. | [optional] 
**Stats** | Pointer to [**InterfaceStats**](InterfaceStats.md) |  | [optional] 
**Subnet** | Pointer to **NullableString** | Specifies the subnet of the network interface. | [optional] 
**SubnetV6** | Pointer to **NullableString** | Specifies the subnetV6 of the network interface. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the network interface. | [optional] 
**VirtualIP** | Pointer to **NullableString** | Specifies the virtual IP of the network interface. | [optional] 

## Methods

### NewNetworkInterface

`func NewNetworkInterface() *NetworkInterface`

NewNetworkInterface instantiates a new NetworkInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInterfaceWithDefaults

`func NewNetworkInterfaceWithDefaults() *NetworkInterface`

NewNetworkInterfaceWithDefaults instantiates a new NetworkInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveBondSlave

`func (o *NetworkInterface) GetActiveBondSlave() string`

GetActiveBondSlave returns the ActiveBondSlave field if non-nil, zero value otherwise.

### GetActiveBondSlaveOk

`func (o *NetworkInterface) GetActiveBondSlaveOk() (*string, bool)`

GetActiveBondSlaveOk returns a tuple with the ActiveBondSlave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveBondSlave

`func (o *NetworkInterface) SetActiveBondSlave(v string)`

SetActiveBondSlave sets ActiveBondSlave field to given value.

### HasActiveBondSlave

`func (o *NetworkInterface) HasActiveBondSlave() bool`

HasActiveBondSlave returns a boolean if a field has been set.

### SetActiveBondSlaveNil

`func (o *NetworkInterface) SetActiveBondSlaveNil(b bool)`

 SetActiveBondSlaveNil sets the value for ActiveBondSlave to be an explicit nil

### UnsetActiveBondSlave
`func (o *NetworkInterface) UnsetActiveBondSlave()`

UnsetActiveBondSlave ensures that no value is present for ActiveBondSlave, not even an explicit nil
### GetBondSlaveNames

`func (o *NetworkInterface) GetBondSlaveNames() []string`

GetBondSlaveNames returns the BondSlaveNames field if non-nil, zero value otherwise.

### GetBondSlaveNamesOk

`func (o *NetworkInterface) GetBondSlaveNamesOk() (*[]string, bool)`

GetBondSlaveNamesOk returns a tuple with the BondSlaveNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondSlaveNames

`func (o *NetworkInterface) SetBondSlaveNames(v []string)`

SetBondSlaveNames sets BondSlaveNames field to given value.

### HasBondSlaveNames

`func (o *NetworkInterface) HasBondSlaveNames() bool`

HasBondSlaveNames returns a boolean if a field has been set.

### SetBondSlaveNamesNil

`func (o *NetworkInterface) SetBondSlaveNamesNil(b bool)`

 SetBondSlaveNamesNil sets the value for BondSlaveNames to be an explicit nil

### UnsetBondSlaveNames
`func (o *NetworkInterface) UnsetBondSlaveNames()`

UnsetBondSlaveNames ensures that no value is present for BondSlaveNames, not even an explicit nil
### GetBondSlaveSlots

`func (o *NetworkInterface) GetBondSlaveSlots() []string`

GetBondSlaveSlots returns the BondSlaveSlots field if non-nil, zero value otherwise.

### GetBondSlaveSlotsOk

`func (o *NetworkInterface) GetBondSlaveSlotsOk() (*[]string, bool)`

GetBondSlaveSlotsOk returns a tuple with the BondSlaveSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondSlaveSlots

`func (o *NetworkInterface) SetBondSlaveSlots(v []string)`

SetBondSlaveSlots sets BondSlaveSlots field to given value.

### HasBondSlaveSlots

`func (o *NetworkInterface) HasBondSlaveSlots() bool`

HasBondSlaveSlots returns a boolean if a field has been set.

### SetBondSlaveSlotsNil

`func (o *NetworkInterface) SetBondSlaveSlotsNil(b bool)`

 SetBondSlaveSlotsNil sets the value for BondSlaveSlots to be an explicit nil

### UnsetBondSlaveSlots
`func (o *NetworkInterface) UnsetBondSlaveSlots()`

UnsetBondSlaveSlots ensures that no value is present for BondSlaveSlots, not even an explicit nil
### GetBondSlavesDetails

`func (o *NetworkInterface) GetBondSlavesDetails() []BondMember`

GetBondSlavesDetails returns the BondSlavesDetails field if non-nil, zero value otherwise.

### GetBondSlavesDetailsOk

`func (o *NetworkInterface) GetBondSlavesDetailsOk() (*[]BondMember, bool)`

GetBondSlavesDetailsOk returns a tuple with the BondSlavesDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondSlavesDetails

`func (o *NetworkInterface) SetBondSlavesDetails(v []BondMember)`

SetBondSlavesDetails sets BondSlavesDetails field to given value.

### HasBondSlavesDetails

`func (o *NetworkInterface) HasBondSlavesDetails() bool`

HasBondSlavesDetails returns a boolean if a field has been set.

### SetBondSlavesDetailsNil

`func (o *NetworkInterface) SetBondSlavesDetailsNil(b bool)`

 SetBondSlavesDetailsNil sets the value for BondSlavesDetails to be an explicit nil

### UnsetBondSlavesDetails
`func (o *NetworkInterface) UnsetBondSlavesDetails()`

UnsetBondSlavesDetails ensures that no value is present for BondSlavesDetails, not even an explicit nil
### GetBondingMode

`func (o *NetworkInterface) GetBondingMode() string`

GetBondingMode returns the BondingMode field if non-nil, zero value otherwise.

### GetBondingModeOk

`func (o *NetworkInterface) GetBondingModeOk() (*string, bool)`

GetBondingModeOk returns a tuple with the BondingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondingMode

`func (o *NetworkInterface) SetBondingMode(v string)`

SetBondingMode sets BondingMode field to given value.

### HasBondingMode

`func (o *NetworkInterface) HasBondingMode() bool`

HasBondingMode returns a boolean if a field has been set.

### SetBondingModeNil

`func (o *NetworkInterface) SetBondingModeNil(b bool)`

 SetBondingModeNil sets the value for BondingMode to be an explicit nil

### UnsetBondingMode
`func (o *NetworkInterface) UnsetBondingMode()`

UnsetBondingMode ensures that no value is present for BondingMode, not even an explicit nil
### GetDefaultRoute

`func (o *NetworkInterface) GetDefaultRoute() bool`

GetDefaultRoute returns the DefaultRoute field if non-nil, zero value otherwise.

### GetDefaultRouteOk

`func (o *NetworkInterface) GetDefaultRouteOk() (*bool, bool)`

GetDefaultRouteOk returns a tuple with the DefaultRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoute

`func (o *NetworkInterface) SetDefaultRoute(v bool)`

SetDefaultRoute sets DefaultRoute field to given value.

### HasDefaultRoute

`func (o *NetworkInterface) HasDefaultRoute() bool`

HasDefaultRoute returns a boolean if a field has been set.

### SetDefaultRouteNil

`func (o *NetworkInterface) SetDefaultRouteNil(b bool)`

 SetDefaultRouteNil sets the value for DefaultRoute to be an explicit nil

### UnsetDefaultRoute
`func (o *NetworkInterface) UnsetDefaultRoute()`

UnsetDefaultRoute ensures that no value is present for DefaultRoute, not even an explicit nil
### GetGateway

`func (o *NetworkInterface) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *NetworkInterface) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *NetworkInterface) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *NetworkInterface) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *NetworkInterface) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *NetworkInterface) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetGatewayV6

`func (o *NetworkInterface) GetGatewayV6() string`

GetGatewayV6 returns the GatewayV6 field if non-nil, zero value otherwise.

### GetGatewayV6Ok

`func (o *NetworkInterface) GetGatewayV6Ok() (*string, bool)`

GetGatewayV6Ok returns a tuple with the GatewayV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayV6

`func (o *NetworkInterface) SetGatewayV6(v string)`

SetGatewayV6 sets GatewayV6 field to given value.

### HasGatewayV6

`func (o *NetworkInterface) HasGatewayV6() bool`

HasGatewayV6 returns a boolean if a field has been set.

### SetGatewayV6Nil

`func (o *NetworkInterface) SetGatewayV6Nil(b bool)`

 SetGatewayV6Nil sets the value for GatewayV6 to be an explicit nil

### UnsetGatewayV6
`func (o *NetworkInterface) UnsetGatewayV6()`

UnsetGatewayV6 ensures that no value is present for GatewayV6, not even an explicit nil
### GetGroup

`func (o *NetworkInterface) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *NetworkInterface) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *NetworkInterface) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *NetworkInterface) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *NetworkInterface) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *NetworkInterface) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetIsConnected

`func (o *NetworkInterface) GetIsConnected() bool`

GetIsConnected returns the IsConnected field if non-nil, zero value otherwise.

### GetIsConnectedOk

`func (o *NetworkInterface) GetIsConnectedOk() (*bool, bool)`

GetIsConnectedOk returns a tuple with the IsConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConnected

`func (o *NetworkInterface) SetIsConnected(v bool)`

SetIsConnected sets IsConnected field to given value.

### HasIsConnected

`func (o *NetworkInterface) HasIsConnected() bool`

HasIsConnected returns a boolean if a field has been set.

### SetIsConnectedNil

`func (o *NetworkInterface) SetIsConnectedNil(b bool)`

 SetIsConnectedNil sets the value for IsConnected to be an explicit nil

### UnsetIsConnected
`func (o *NetworkInterface) UnsetIsConnected()`

UnsetIsConnected ensures that no value is present for IsConnected, not even an explicit nil
### GetIsUp

`func (o *NetworkInterface) GetIsUp() bool`

GetIsUp returns the IsUp field if non-nil, zero value otherwise.

### GetIsUpOk

`func (o *NetworkInterface) GetIsUpOk() (*bool, bool)`

GetIsUpOk returns a tuple with the IsUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUp

`func (o *NetworkInterface) SetIsUp(v bool)`

SetIsUp sets IsUp field to given value.

### HasIsUp

`func (o *NetworkInterface) HasIsUp() bool`

HasIsUp returns a boolean if a field has been set.

### SetIsUpNil

`func (o *NetworkInterface) SetIsUpNil(b bool)`

 SetIsUpNil sets the value for IsUp to be an explicit nil

### UnsetIsUp
`func (o *NetworkInterface) UnsetIsUp()`

UnsetIsUp ensures that no value is present for IsUp, not even an explicit nil
### GetMacAddress

`func (o *NetworkInterface) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *NetworkInterface) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *NetworkInterface) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *NetworkInterface) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *NetworkInterface) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *NetworkInterface) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
### GetMtu

`func (o *NetworkInterface) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *NetworkInterface) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *NetworkInterface) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *NetworkInterface) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *NetworkInterface) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *NetworkInterface) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetName

`func (o *NetworkInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkInterface) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *NetworkInterface) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *NetworkInterface) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRole

`func (o *NetworkInterface) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *NetworkInterface) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *NetworkInterface) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *NetworkInterface) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *NetworkInterface) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *NetworkInterface) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetServices

`func (o *NetworkInterface) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *NetworkInterface) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *NetworkInterface) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *NetworkInterface) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *NetworkInterface) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *NetworkInterface) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil
### GetSpeed

`func (o *NetworkInterface) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *NetworkInterface) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *NetworkInterface) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *NetworkInterface) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### SetSpeedNil

`func (o *NetworkInterface) SetSpeedNil(b bool)`

 SetSpeedNil sets the value for Speed to be an explicit nil

### UnsetSpeed
`func (o *NetworkInterface) UnsetSpeed()`

UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
### GetStaticIP

`func (o *NetworkInterface) GetStaticIP() string`

GetStaticIP returns the StaticIP field if non-nil, zero value otherwise.

### GetStaticIPOk

`func (o *NetworkInterface) GetStaticIPOk() (*string, bool)`

GetStaticIPOk returns a tuple with the StaticIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIP

`func (o *NetworkInterface) SetStaticIP(v string)`

SetStaticIP sets StaticIP field to given value.

### HasStaticIP

`func (o *NetworkInterface) HasStaticIP() bool`

HasStaticIP returns a boolean if a field has been set.

### SetStaticIPNil

`func (o *NetworkInterface) SetStaticIPNil(b bool)`

 SetStaticIPNil sets the value for StaticIP to be an explicit nil

### UnsetStaticIP
`func (o *NetworkInterface) UnsetStaticIP()`

UnsetStaticIP ensures that no value is present for StaticIP, not even an explicit nil
### GetStaticIpV6

`func (o *NetworkInterface) GetStaticIpV6() string`

GetStaticIpV6 returns the StaticIpV6 field if non-nil, zero value otherwise.

### GetStaticIpV6Ok

`func (o *NetworkInterface) GetStaticIpV6Ok() (*string, bool)`

GetStaticIpV6Ok returns a tuple with the StaticIpV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIpV6

`func (o *NetworkInterface) SetStaticIpV6(v string)`

SetStaticIpV6 sets StaticIpV6 field to given value.

### HasStaticIpV6

`func (o *NetworkInterface) HasStaticIpV6() bool`

HasStaticIpV6 returns a boolean if a field has been set.

### SetStaticIpV6Nil

`func (o *NetworkInterface) SetStaticIpV6Nil(b bool)`

 SetStaticIpV6Nil sets the value for StaticIpV6 to be an explicit nil

### UnsetStaticIpV6
`func (o *NetworkInterface) UnsetStaticIpV6()`

UnsetStaticIpV6 ensures that no value is present for StaticIpV6, not even an explicit nil
### GetStats

`func (o *NetworkInterface) GetStats() InterfaceStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *NetworkInterface) GetStatsOk() (*InterfaceStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *NetworkInterface) SetStats(v InterfaceStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *NetworkInterface) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetSubnet

`func (o *NetworkInterface) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *NetworkInterface) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *NetworkInterface) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *NetworkInterface) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### SetSubnetNil

`func (o *NetworkInterface) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *NetworkInterface) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetSubnetV6

`func (o *NetworkInterface) GetSubnetV6() string`

GetSubnetV6 returns the SubnetV6 field if non-nil, zero value otherwise.

### GetSubnetV6Ok

`func (o *NetworkInterface) GetSubnetV6Ok() (*string, bool)`

GetSubnetV6Ok returns a tuple with the SubnetV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetV6

`func (o *NetworkInterface) SetSubnetV6(v string)`

SetSubnetV6 sets SubnetV6 field to given value.

### HasSubnetV6

`func (o *NetworkInterface) HasSubnetV6() bool`

HasSubnetV6 returns a boolean if a field has been set.

### SetSubnetV6Nil

`func (o *NetworkInterface) SetSubnetV6Nil(b bool)`

 SetSubnetV6Nil sets the value for SubnetV6 to be an explicit nil

### UnsetSubnetV6
`func (o *NetworkInterface) UnsetSubnetV6()`

UnsetSubnetV6 ensures that no value is present for SubnetV6, not even an explicit nil
### GetType

`func (o *NetworkInterface) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkInterface) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkInterface) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkInterface) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *NetworkInterface) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *NetworkInterface) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetVirtualIP

`func (o *NetworkInterface) GetVirtualIP() string`

GetVirtualIP returns the VirtualIP field if non-nil, zero value otherwise.

### GetVirtualIPOk

`func (o *NetworkInterface) GetVirtualIPOk() (*string, bool)`

GetVirtualIPOk returns a tuple with the VirtualIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualIP

`func (o *NetworkInterface) SetVirtualIP(v string)`

SetVirtualIP sets VirtualIP field to given value.

### HasVirtualIP

`func (o *NetworkInterface) HasVirtualIP() bool`

HasVirtualIP returns a boolean if a field has been set.

### SetVirtualIPNil

`func (o *NetworkInterface) SetVirtualIPNil(b bool)`

 SetVirtualIPNil sets the value for VirtualIP to be an explicit nil

### UnsetVirtualIP
`func (o *NetworkInterface) UnsetVirtualIP()`

UnsetVirtualIP ensures that no value is present for VirtualIP, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


