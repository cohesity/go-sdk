# NetworkInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies the name of the network interface. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the network interface. | [optional] 
**StaticIP** | Pointer to **NullableString** | Specifies the static IP of the network interface. | [optional] 
**VirtualIP** | Pointer to **NullableString** | Specifies the virtual IP of the network interface. | [optional] 
**Gateway** | Pointer to **NullableString** | Specifies the gateway of the network interface. | [optional] 
**Mtu** | Pointer to **NullableInt32** | Specifies the MTU of the network interface. | [optional] 
**Subnet** | Pointer to **NullableString** | Specifies the subnet of the network interface. | [optional] 
**IsUp** | Pointer to **NullableBool** | Specifies whether or not the interface is up. | [optional] 
**Group** | Pointer to **NullableString** | Specifies the group to which this interface belongs. | [optional] 
**Role** | Pointer to **NullableString** | Specifies the interface role. | [optional] 
**DefaultRoute** | Pointer to **NullableBool** | Specifies whether or not this interface is the default route. | [optional] 
**BondSlaveNames** | Pointer to **[]string** | Specifies the names of the bond slaves for this interface. | [optional] 
**BondSlaveSlots** | Pointer to **[]string** | Specifies the slots of the bond slaves for this interface. | [optional] 
**BondingMode** | Pointer to **NullableString** | Specifies the bonding mode of this interface. | [optional] 
**MacAddress** | Pointer to **NullableString** | Specifies the MAC address of this interface. | [optional] 
**IsConnected** | Pointer to **NullableBool** | Specifies whether or not this interface is connected. | [optional] 
**Speed** | Pointer to **NullableString** | Specifies the speed of this interface. | [optional] 

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


