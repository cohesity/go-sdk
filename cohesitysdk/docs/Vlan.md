# Vlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddToClusterPartition** | Pointer to **NullableBool** | Specifies whether to add the VLAN IPs to the cluster partition that already has one or more IPs from this VLAN. | [optional] 
**AllTenantAccess** | Pointer to **NullableBool** | Specifies if this VLAN can be used by all tenants without explicit assignment to them. This option can only be set true for VLANs that are not assigned to any tenant. | [optional] 
**AppIpVecInUse** | Pointer to **NullableBool** | Set to true when ips are in use by Athena Apps. Note: If it is true then vlan interface can&#39;t be deleted. | [optional] 
**Appsips** | Pointer to **[]string** | Array of Athena Apps IPs.  Specifies a list of Athena IPs in the VLAN. | [optional] 
**Description** | Pointer to **NullableString** | Specifies a description of the VLAN. | [optional] 
**Gateway** | Pointer to **NullableString** | Specifies the Gateway of the VLAN. It can carry V4 or V6 in case of requests, and carrises V4 in case of response. | [optional] 
**GatewayV6** | Pointer to **NullableString** | Specifies the Gateway of the VLAN. | [optional] 
**Hostname** | Pointer to **NullableString** | Specifies the hostname of the VLAN. | [optional] 
**Id** | Pointer to **NullableInt32** | Specifies the id of the VLAN. | [optional] 
**IfaceGroupName** | Pointer to **NullableString** | Specifies the interface group name of the VLAN. It is in the format of &lt;base_interface_group_name&gt;.&lt;vlan_id&gt;. | [optional] 
**InterfaceName** | Pointer to **NullableString** | Specifies the interface name of the VLAN. | [optional] 
**IpFamily** | Pointer to **NullableInt32** | Specifies IP family. Based on this, subnet/gateway field contains V4 or V6 values. Used in Request. | [optional] 
**IpPoolMap** | Pointer to **map[string][]string** | IpPoolMap.  Pool IPs to program VIP followers. | [optional] 
**IpRange** | Pointer to [**IpRange**](IpRange.md) |  | [optional] 
**Ips** | Pointer to **[]string** | Array of IPs.  Specifies a list of IPs in the VLAN. | [optional] 
**Mtu** | Pointer to **NullableInt32** |  | [optional] 
**Subnet** | Pointer to [**NullableSubnet**](Subnet.md) | Specifies the subnet of the VLAN. The netmask can be specified by setting netmaskBits or netmaskIp4. The netmask can only be set using netmaskIp4 if the IP address is an IPv4 address. It can carry V4 or V6 in case of requests, and carries V4 in case of response. | [optional] 
**SubnetV6** | Pointer to [**NullableSubnet**](Subnet.md) | Specifies the subnet of the VLAN. The netmask can be specified by setting netmaskBits or netmaskIp4. The netmask can only be set using netmaskIp4 if the IP address is an IPv4 address. | [optional] 
**TenantId** | Pointer to **NullableString** | Optional tenant id that this vlan belongs to. | [optional] 
**VlanName** | Pointer to **NullableString** | Specifies the VLAN name of the vlanId. | [optional] 

## Methods

### NewVlan

`func NewVlan() *Vlan`

NewVlan instantiates a new Vlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanWithDefaults

`func NewVlanWithDefaults() *Vlan`

NewVlanWithDefaults instantiates a new Vlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddToClusterPartition

`func (o *Vlan) GetAddToClusterPartition() bool`

GetAddToClusterPartition returns the AddToClusterPartition field if non-nil, zero value otherwise.

### GetAddToClusterPartitionOk

`func (o *Vlan) GetAddToClusterPartitionOk() (*bool, bool)`

GetAddToClusterPartitionOk returns a tuple with the AddToClusterPartition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddToClusterPartition

`func (o *Vlan) SetAddToClusterPartition(v bool)`

SetAddToClusterPartition sets AddToClusterPartition field to given value.

### HasAddToClusterPartition

`func (o *Vlan) HasAddToClusterPartition() bool`

HasAddToClusterPartition returns a boolean if a field has been set.

### SetAddToClusterPartitionNil

`func (o *Vlan) SetAddToClusterPartitionNil(b bool)`

 SetAddToClusterPartitionNil sets the value for AddToClusterPartition to be an explicit nil

### UnsetAddToClusterPartition
`func (o *Vlan) UnsetAddToClusterPartition()`

UnsetAddToClusterPartition ensures that no value is present for AddToClusterPartition, not even an explicit nil
### GetAllTenantAccess

`func (o *Vlan) GetAllTenantAccess() bool`

GetAllTenantAccess returns the AllTenantAccess field if non-nil, zero value otherwise.

### GetAllTenantAccessOk

`func (o *Vlan) GetAllTenantAccessOk() (*bool, bool)`

GetAllTenantAccessOk returns a tuple with the AllTenantAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllTenantAccess

`func (o *Vlan) SetAllTenantAccess(v bool)`

SetAllTenantAccess sets AllTenantAccess field to given value.

### HasAllTenantAccess

`func (o *Vlan) HasAllTenantAccess() bool`

HasAllTenantAccess returns a boolean if a field has been set.

### SetAllTenantAccessNil

`func (o *Vlan) SetAllTenantAccessNil(b bool)`

 SetAllTenantAccessNil sets the value for AllTenantAccess to be an explicit nil

### UnsetAllTenantAccess
`func (o *Vlan) UnsetAllTenantAccess()`

UnsetAllTenantAccess ensures that no value is present for AllTenantAccess, not even an explicit nil
### GetAppIpVecInUse

`func (o *Vlan) GetAppIpVecInUse() bool`

GetAppIpVecInUse returns the AppIpVecInUse field if non-nil, zero value otherwise.

### GetAppIpVecInUseOk

`func (o *Vlan) GetAppIpVecInUseOk() (*bool, bool)`

GetAppIpVecInUseOk returns a tuple with the AppIpVecInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIpVecInUse

`func (o *Vlan) SetAppIpVecInUse(v bool)`

SetAppIpVecInUse sets AppIpVecInUse field to given value.

### HasAppIpVecInUse

`func (o *Vlan) HasAppIpVecInUse() bool`

HasAppIpVecInUse returns a boolean if a field has been set.

### SetAppIpVecInUseNil

`func (o *Vlan) SetAppIpVecInUseNil(b bool)`

 SetAppIpVecInUseNil sets the value for AppIpVecInUse to be an explicit nil

### UnsetAppIpVecInUse
`func (o *Vlan) UnsetAppIpVecInUse()`

UnsetAppIpVecInUse ensures that no value is present for AppIpVecInUse, not even an explicit nil
### GetAppsips

`func (o *Vlan) GetAppsips() []string`

GetAppsips returns the Appsips field if non-nil, zero value otherwise.

### GetAppsipsOk

`func (o *Vlan) GetAppsipsOk() (*[]string, bool)`

GetAppsipsOk returns a tuple with the Appsips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppsips

`func (o *Vlan) SetAppsips(v []string)`

SetAppsips sets Appsips field to given value.

### HasAppsips

`func (o *Vlan) HasAppsips() bool`

HasAppsips returns a boolean if a field has been set.

### SetAppsipsNil

`func (o *Vlan) SetAppsipsNil(b bool)`

 SetAppsipsNil sets the value for Appsips to be an explicit nil

### UnsetAppsips
`func (o *Vlan) UnsetAppsips()`

UnsetAppsips ensures that no value is present for Appsips, not even an explicit nil
### GetDescription

`func (o *Vlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Vlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Vlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Vlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Vlan) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Vlan) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetGateway

`func (o *Vlan) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *Vlan) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *Vlan) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *Vlan) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *Vlan) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *Vlan) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetGatewayV6

`func (o *Vlan) GetGatewayV6() string`

GetGatewayV6 returns the GatewayV6 field if non-nil, zero value otherwise.

### GetGatewayV6Ok

`func (o *Vlan) GetGatewayV6Ok() (*string, bool)`

GetGatewayV6Ok returns a tuple with the GatewayV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayV6

`func (o *Vlan) SetGatewayV6(v string)`

SetGatewayV6 sets GatewayV6 field to given value.

### HasGatewayV6

`func (o *Vlan) HasGatewayV6() bool`

HasGatewayV6 returns a boolean if a field has been set.

### SetGatewayV6Nil

`func (o *Vlan) SetGatewayV6Nil(b bool)`

 SetGatewayV6Nil sets the value for GatewayV6 to be an explicit nil

### UnsetGatewayV6
`func (o *Vlan) UnsetGatewayV6()`

UnsetGatewayV6 ensures that no value is present for GatewayV6, not even an explicit nil
### GetHostname

`func (o *Vlan) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Vlan) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Vlan) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *Vlan) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostnameNil

`func (o *Vlan) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *Vlan) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
### GetId

`func (o *Vlan) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Vlan) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Vlan) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Vlan) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Vlan) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Vlan) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIfaceGroupName

`func (o *Vlan) GetIfaceGroupName() string`

GetIfaceGroupName returns the IfaceGroupName field if non-nil, zero value otherwise.

### GetIfaceGroupNameOk

`func (o *Vlan) GetIfaceGroupNameOk() (*string, bool)`

GetIfaceGroupNameOk returns a tuple with the IfaceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfaceGroupName

`func (o *Vlan) SetIfaceGroupName(v string)`

SetIfaceGroupName sets IfaceGroupName field to given value.

### HasIfaceGroupName

`func (o *Vlan) HasIfaceGroupName() bool`

HasIfaceGroupName returns a boolean if a field has been set.

### SetIfaceGroupNameNil

`func (o *Vlan) SetIfaceGroupNameNil(b bool)`

 SetIfaceGroupNameNil sets the value for IfaceGroupName to be an explicit nil

### UnsetIfaceGroupName
`func (o *Vlan) UnsetIfaceGroupName()`

UnsetIfaceGroupName ensures that no value is present for IfaceGroupName, not even an explicit nil
### GetInterfaceName

`func (o *Vlan) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *Vlan) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *Vlan) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *Vlan) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### SetInterfaceNameNil

`func (o *Vlan) SetInterfaceNameNil(b bool)`

 SetInterfaceNameNil sets the value for InterfaceName to be an explicit nil

### UnsetInterfaceName
`func (o *Vlan) UnsetInterfaceName()`

UnsetInterfaceName ensures that no value is present for InterfaceName, not even an explicit nil
### GetIpFamily

`func (o *Vlan) GetIpFamily() int32`

GetIpFamily returns the IpFamily field if non-nil, zero value otherwise.

### GetIpFamilyOk

`func (o *Vlan) GetIpFamilyOk() (*int32, bool)`

GetIpFamilyOk returns a tuple with the IpFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFamily

`func (o *Vlan) SetIpFamily(v int32)`

SetIpFamily sets IpFamily field to given value.

### HasIpFamily

`func (o *Vlan) HasIpFamily() bool`

HasIpFamily returns a boolean if a field has been set.

### SetIpFamilyNil

`func (o *Vlan) SetIpFamilyNil(b bool)`

 SetIpFamilyNil sets the value for IpFamily to be an explicit nil

### UnsetIpFamily
`func (o *Vlan) UnsetIpFamily()`

UnsetIpFamily ensures that no value is present for IpFamily, not even an explicit nil
### GetIpPoolMap

`func (o *Vlan) GetIpPoolMap() map[string][]string`

GetIpPoolMap returns the IpPoolMap field if non-nil, zero value otherwise.

### GetIpPoolMapOk

`func (o *Vlan) GetIpPoolMapOk() (*map[string][]string, bool)`

GetIpPoolMapOk returns a tuple with the IpPoolMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolMap

`func (o *Vlan) SetIpPoolMap(v map[string][]string)`

SetIpPoolMap sets IpPoolMap field to given value.

### HasIpPoolMap

`func (o *Vlan) HasIpPoolMap() bool`

HasIpPoolMap returns a boolean if a field has been set.

### SetIpPoolMapNil

`func (o *Vlan) SetIpPoolMapNil(b bool)`

 SetIpPoolMapNil sets the value for IpPoolMap to be an explicit nil

### UnsetIpPoolMap
`func (o *Vlan) UnsetIpPoolMap()`

UnsetIpPoolMap ensures that no value is present for IpPoolMap, not even an explicit nil
### GetIpRange

`func (o *Vlan) GetIpRange() IpRange`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *Vlan) GetIpRangeOk() (*IpRange, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *Vlan) SetIpRange(v IpRange)`

SetIpRange sets IpRange field to given value.

### HasIpRange

`func (o *Vlan) HasIpRange() bool`

HasIpRange returns a boolean if a field has been set.

### GetIps

`func (o *Vlan) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *Vlan) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *Vlan) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *Vlan) HasIps() bool`

HasIps returns a boolean if a field has been set.

### SetIpsNil

`func (o *Vlan) SetIpsNil(b bool)`

 SetIpsNil sets the value for Ips to be an explicit nil

### UnsetIps
`func (o *Vlan) UnsetIps()`

UnsetIps ensures that no value is present for Ips, not even an explicit nil
### GetMtu

`func (o *Vlan) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *Vlan) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *Vlan) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *Vlan) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *Vlan) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *Vlan) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetSubnet

`func (o *Vlan) GetSubnet() Subnet`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *Vlan) GetSubnetOk() (*Subnet, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *Vlan) SetSubnet(v Subnet)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *Vlan) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### SetSubnetNil

`func (o *Vlan) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *Vlan) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetSubnetV6

`func (o *Vlan) GetSubnetV6() Subnet`

GetSubnetV6 returns the SubnetV6 field if non-nil, zero value otherwise.

### GetSubnetV6Ok

`func (o *Vlan) GetSubnetV6Ok() (*Subnet, bool)`

GetSubnetV6Ok returns a tuple with the SubnetV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetV6

`func (o *Vlan) SetSubnetV6(v Subnet)`

SetSubnetV6 sets SubnetV6 field to given value.

### HasSubnetV6

`func (o *Vlan) HasSubnetV6() bool`

HasSubnetV6 returns a boolean if a field has been set.

### SetSubnetV6Nil

`func (o *Vlan) SetSubnetV6Nil(b bool)`

 SetSubnetV6Nil sets the value for SubnetV6 to be an explicit nil

### UnsetSubnetV6
`func (o *Vlan) UnsetSubnetV6()`

UnsetSubnetV6 ensures that no value is present for SubnetV6, not even an explicit nil
### GetTenantId

`func (o *Vlan) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Vlan) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Vlan) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Vlan) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *Vlan) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *Vlan) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetVlanName

`func (o *Vlan) GetVlanName() string`

GetVlanName returns the VlanName field if non-nil, zero value otherwise.

### GetVlanNameOk

`func (o *Vlan) GetVlanNameOk() (*string, bool)`

GetVlanNameOk returns a tuple with the VlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanName

`func (o *Vlan) SetVlanName(v string)`

SetVlanName sets VlanName field to given value.

### HasVlanName

`func (o *Vlan) HasVlanName() bool`

HasVlanName returns a boolean if a field has been set.

### SetVlanNameNil

`func (o *Vlan) SetVlanNameNil(b bool)`

 SetVlanNameNil sets the value for VlanName to be an explicit nil

### UnsetVlanName
`func (o *Vlan) UnsetVlanName()`

UnsetVlanName ensures that no value is present for VlanName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


