# NetworkInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BondSlaveSlotTypes** | Pointer to **[]string** | Specifies the types of the slots of any slaves if this interface is a bond. | [optional] 
**BondSlaves** | Pointer to **[]string** | Specifies the names of any slaves if this interface is a bond. | [optional] 
**BondingMode** | Pointer to **NullableInt32** | Specifies the bonding mode if this interface is a bond. | [optional] 
**Gateway** | Pointer to **NullableString** | Specifies the gateway of the interface. | [optional] 
**Gateway6** | Pointer to **NullableString** | Specifies the gateway6 of the interface. | [optional] 
**Group** | Pointer to **NullableString** | Specifies the group that this interface belongs to. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the ID of this network interface. | [optional] 
**IsConnected** | Pointer to **NullableBool** | Specifies whether or not the Interface is connected. | [optional] 
**IsDefaultRoute** | Pointer to **NullableBool** | Specifies whether or not to use this interface as the default route. | [optional] 
**IsUp** | Pointer to **NullableBool** | Specifies whether or not the interface is currently  up. | [optional] 
**MacAddress** | Pointer to **NullableString** | Specifies the Mac address of the Interface. | [optional] 
**Mtu** | Pointer to **NullableInt32** | Specifies the MTU of the interface. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the interface port. | [optional] 
**Role** | Pointer to **NullableString** | Specifies the role of this interface. | [optional] 
**Services** | Pointer to **[]string** | Specifies the types of services this interface is used for. | [optional] 
**Speed** | Pointer to **NullableString** | Specifies the speed of the Interface. | [optional] 
**StaticIp** | Pointer to **NullableString** | Specifies the static IP of the interface. | [optional] 
**StaticIp6** | Pointer to **NullableString** | Specifies the static IPv6 of the interface. | [optional] 
**Subnet** | Pointer to **NullableString** | Specifies the subnet mask of the interface. | [optional] 
**Subnet6** | Pointer to **NullableString** | Specifies the subnet6 mask of the interface. | [optional] 
**Type** | Pointer to **NullableInt32** | Specifies the type of interface. | [optional] 
**VirtualIp** | Pointer to **NullableString** | Specifies the virtual IP of the interface. | [optional] 

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

### GetBondSlaveSlotTypes

`func (o *NetworkInterface) GetBondSlaveSlotTypes() []string`

GetBondSlaveSlotTypes returns the BondSlaveSlotTypes field if non-nil, zero value otherwise.

### GetBondSlaveSlotTypesOk

`func (o *NetworkInterface) GetBondSlaveSlotTypesOk() (*[]string, bool)`

GetBondSlaveSlotTypesOk returns a tuple with the BondSlaveSlotTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondSlaveSlotTypes

`func (o *NetworkInterface) SetBondSlaveSlotTypes(v []string)`

SetBondSlaveSlotTypes sets BondSlaveSlotTypes field to given value.

### HasBondSlaveSlotTypes

`func (o *NetworkInterface) HasBondSlaveSlotTypes() bool`

HasBondSlaveSlotTypes returns a boolean if a field has been set.

### SetBondSlaveSlotTypesNil

`func (o *NetworkInterface) SetBondSlaveSlotTypesNil(b bool)`

 SetBondSlaveSlotTypesNil sets the value for BondSlaveSlotTypes to be an explicit nil

### UnsetBondSlaveSlotTypes
`func (o *NetworkInterface) UnsetBondSlaveSlotTypes()`

UnsetBondSlaveSlotTypes ensures that no value is present for BondSlaveSlotTypes, not even an explicit nil
### GetBondSlaves

`func (o *NetworkInterface) GetBondSlaves() []string`

GetBondSlaves returns the BondSlaves field if non-nil, zero value otherwise.

### GetBondSlavesOk

`func (o *NetworkInterface) GetBondSlavesOk() (*[]string, bool)`

GetBondSlavesOk returns a tuple with the BondSlaves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondSlaves

`func (o *NetworkInterface) SetBondSlaves(v []string)`

SetBondSlaves sets BondSlaves field to given value.

### HasBondSlaves

`func (o *NetworkInterface) HasBondSlaves() bool`

HasBondSlaves returns a boolean if a field has been set.

### SetBondSlavesNil

`func (o *NetworkInterface) SetBondSlavesNil(b bool)`

 SetBondSlavesNil sets the value for BondSlaves to be an explicit nil

### UnsetBondSlaves
`func (o *NetworkInterface) UnsetBondSlaves()`

UnsetBondSlaves ensures that no value is present for BondSlaves, not even an explicit nil
### GetBondingMode

`func (o *NetworkInterface) GetBondingMode() int32`

GetBondingMode returns the BondingMode field if non-nil, zero value otherwise.

### GetBondingModeOk

`func (o *NetworkInterface) GetBondingModeOk() (*int32, bool)`

GetBondingModeOk returns a tuple with the BondingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondingMode

`func (o *NetworkInterface) SetBondingMode(v int32)`

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
### GetGateway6

`func (o *NetworkInterface) GetGateway6() string`

GetGateway6 returns the Gateway6 field if non-nil, zero value otherwise.

### GetGateway6Ok

`func (o *NetworkInterface) GetGateway6Ok() (*string, bool)`

GetGateway6Ok returns a tuple with the Gateway6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway6

`func (o *NetworkInterface) SetGateway6(v string)`

SetGateway6 sets Gateway6 field to given value.

### HasGateway6

`func (o *NetworkInterface) HasGateway6() bool`

HasGateway6 returns a boolean if a field has been set.

### SetGateway6Nil

`func (o *NetworkInterface) SetGateway6Nil(b bool)`

 SetGateway6Nil sets the value for Gateway6 to be an explicit nil

### UnsetGateway6
`func (o *NetworkInterface) UnsetGateway6()`

UnsetGateway6 ensures that no value is present for Gateway6, not even an explicit nil
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
### GetId

`func (o *NetworkInterface) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkInterface) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkInterface) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkInterface) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *NetworkInterface) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *NetworkInterface) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
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
### GetIsDefaultRoute

`func (o *NetworkInterface) GetIsDefaultRoute() bool`

GetIsDefaultRoute returns the IsDefaultRoute field if non-nil, zero value otherwise.

### GetIsDefaultRouteOk

`func (o *NetworkInterface) GetIsDefaultRouteOk() (*bool, bool)`

GetIsDefaultRouteOk returns a tuple with the IsDefaultRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultRoute

`func (o *NetworkInterface) SetIsDefaultRoute(v bool)`

SetIsDefaultRoute sets IsDefaultRoute field to given value.

### HasIsDefaultRoute

`func (o *NetworkInterface) HasIsDefaultRoute() bool`

HasIsDefaultRoute returns a boolean if a field has been set.

### SetIsDefaultRouteNil

`func (o *NetworkInterface) SetIsDefaultRouteNil(b bool)`

 SetIsDefaultRouteNil sets the value for IsDefaultRoute to be an explicit nil

### UnsetIsDefaultRoute
`func (o *NetworkInterface) UnsetIsDefaultRoute()`

UnsetIsDefaultRoute ensures that no value is present for IsDefaultRoute, not even an explicit nil
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
### GetStaticIp

`func (o *NetworkInterface) GetStaticIp() string`

GetStaticIp returns the StaticIp field if non-nil, zero value otherwise.

### GetStaticIpOk

`func (o *NetworkInterface) GetStaticIpOk() (*string, bool)`

GetStaticIpOk returns a tuple with the StaticIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIp

`func (o *NetworkInterface) SetStaticIp(v string)`

SetStaticIp sets StaticIp field to given value.

### HasStaticIp

`func (o *NetworkInterface) HasStaticIp() bool`

HasStaticIp returns a boolean if a field has been set.

### SetStaticIpNil

`func (o *NetworkInterface) SetStaticIpNil(b bool)`

 SetStaticIpNil sets the value for StaticIp to be an explicit nil

### UnsetStaticIp
`func (o *NetworkInterface) UnsetStaticIp()`

UnsetStaticIp ensures that no value is present for StaticIp, not even an explicit nil
### GetStaticIp6

`func (o *NetworkInterface) GetStaticIp6() string`

GetStaticIp6 returns the StaticIp6 field if non-nil, zero value otherwise.

### GetStaticIp6Ok

`func (o *NetworkInterface) GetStaticIp6Ok() (*string, bool)`

GetStaticIp6Ok returns a tuple with the StaticIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIp6

`func (o *NetworkInterface) SetStaticIp6(v string)`

SetStaticIp6 sets StaticIp6 field to given value.

### HasStaticIp6

`func (o *NetworkInterface) HasStaticIp6() bool`

HasStaticIp6 returns a boolean if a field has been set.

### SetStaticIp6Nil

`func (o *NetworkInterface) SetStaticIp6Nil(b bool)`

 SetStaticIp6Nil sets the value for StaticIp6 to be an explicit nil

### UnsetStaticIp6
`func (o *NetworkInterface) UnsetStaticIp6()`

UnsetStaticIp6 ensures that no value is present for StaticIp6, not even an explicit nil
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
### GetSubnet6

`func (o *NetworkInterface) GetSubnet6() string`

GetSubnet6 returns the Subnet6 field if non-nil, zero value otherwise.

### GetSubnet6Ok

`func (o *NetworkInterface) GetSubnet6Ok() (*string, bool)`

GetSubnet6Ok returns a tuple with the Subnet6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet6

`func (o *NetworkInterface) SetSubnet6(v string)`

SetSubnet6 sets Subnet6 field to given value.

### HasSubnet6

`func (o *NetworkInterface) HasSubnet6() bool`

HasSubnet6 returns a boolean if a field has been set.

### SetSubnet6Nil

`func (o *NetworkInterface) SetSubnet6Nil(b bool)`

 SetSubnet6Nil sets the value for Subnet6 to be an explicit nil

### UnsetSubnet6
`func (o *NetworkInterface) UnsetSubnet6()`

UnsetSubnet6 ensures that no value is present for Subnet6, not even an explicit nil
### GetType

`func (o *NetworkInterface) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkInterface) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkInterface) SetType(v int32)`

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
### GetVirtualIp

`func (o *NetworkInterface) GetVirtualIp() string`

GetVirtualIp returns the VirtualIp field if non-nil, zero value otherwise.

### GetVirtualIpOk

`func (o *NetworkInterface) GetVirtualIpOk() (*string, bool)`

GetVirtualIpOk returns a tuple with the VirtualIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualIp

`func (o *NetworkInterface) SetVirtualIp(v string)`

SetVirtualIp sets VirtualIp field to given value.

### HasVirtualIp

`func (o *NetworkInterface) HasVirtualIp() bool`

HasVirtualIp returns a boolean if a field has been set.

### SetVirtualIpNil

`func (o *NetworkInterface) SetVirtualIpNil(b bool)`

 SetVirtualIpNil sets the value for VirtualIp to be an explicit nil

### UnsetVirtualIp
`func (o *NetworkInterface) UnsetVirtualIp()`

UnsetVirtualIp ensures that no value is present for VirtualIp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


