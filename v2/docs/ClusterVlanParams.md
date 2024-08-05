# ClusterVlanParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllTenantAccess** | Pointer to **NullableBool** | Allow vlan to be used by all tenants without explicit assignment. Set to true only when the vlan is not assigned to any tenant. | [optional] [default to false]
**AppIps** | Pointer to **[]string** | Vlan IP addresses for apps. | [optional] 
**Description** | Pointer to **NullableString** | Description of the vlan. | [optional] 
**DnsDelegationZones** | Pointer to [**[]DnsDelegationZone**](DnsDelegationZone.md) | DNS delegation zones of the vlan. | [optional] 
**EcmpEnabled** | Pointer to **NullableBool** | Set to true to enable ECMP in the vlan. | [optional] [default to false]
**Fqdn** | Pointer to **NullableString** | FQDN of the vlan. | [optional] 
**Gateway** | Pointer to **NullableString** | Subnet gateway of the vlan. This can be Ipv4 or Ipv6 gateway based on the IP addresses type. | [optional] 
**IpAddressesType** | Pointer to **NullableString** | Type of IP addresses. The default value is Ipv4. | [optional] 
**IpPools** | Pointer to [**[]IpPool**](IpPool.md) | IP pools from the vlan ip addresses, the IPs in a pool goes together. One IP from each pool forms a VIP group. | [optional] 
**IpRanges** | Pointer to [**[]IpRange**](IpRange.md) | Vlan IP address ranges, only one of ips or ipRanges parameters should be given. | [optional] 
**Ips** | Pointer to **[]string** | Vlan IP addresses, only one of ips or ipRanges parameters should be given. | [optional] 
**Mtu** | Pointer to **NullableInt32** | MTU of the vlan. | [optional] 
**Subnet** | Pointer to **NullableString** | IPv6 or IPv6 subnet in CIDR format i.e ip-address/prefix. Examples: IPv4 subnet&#39;192.168.0.101/24&#39;, &#39;10.10.1.32/27&#39;. IPv6 subnet &#39;3005:1231:2006:0025::0/96&#39;, 3005:1231:2006:0025::0/128 | [optional] 
**TenantId** | Pointer to **NullableString** | Tenant id to assign vlan to a tenant. | [optional] 
**VlanName** | Pointer to **NullableString** | Name of the Vlan. | [optional] 
**InterfaceName** | **string** | Vlan interface name, it should be in interface_group_name.vlan_id format. | 
**AppIpsInUse** | Pointer to **NullableBool** | Set to true when vlan app IP addresses are being used by apps. When this is set to true, the vlan interface can&#39;t be deleted. | [optional] 

## Methods

### NewClusterVlanParams

`func NewClusterVlanParams(interfaceName string, ) *ClusterVlanParams`

NewClusterVlanParams instantiates a new ClusterVlanParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterVlanParamsWithDefaults

`func NewClusterVlanParamsWithDefaults() *ClusterVlanParams`

NewClusterVlanParamsWithDefaults instantiates a new ClusterVlanParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllTenantAccess

`func (o *ClusterVlanParams) GetAllTenantAccess() bool`

GetAllTenantAccess returns the AllTenantAccess field if non-nil, zero value otherwise.

### GetAllTenantAccessOk

`func (o *ClusterVlanParams) GetAllTenantAccessOk() (*bool, bool)`

GetAllTenantAccessOk returns a tuple with the AllTenantAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllTenantAccess

`func (o *ClusterVlanParams) SetAllTenantAccess(v bool)`

SetAllTenantAccess sets AllTenantAccess field to given value.

### HasAllTenantAccess

`func (o *ClusterVlanParams) HasAllTenantAccess() bool`

HasAllTenantAccess returns a boolean if a field has been set.

### SetAllTenantAccessNil

`func (o *ClusterVlanParams) SetAllTenantAccessNil(b bool)`

 SetAllTenantAccessNil sets the value for AllTenantAccess to be an explicit nil

### UnsetAllTenantAccess
`func (o *ClusterVlanParams) UnsetAllTenantAccess()`

UnsetAllTenantAccess ensures that no value is present for AllTenantAccess, not even an explicit nil
### GetAppIps

`func (o *ClusterVlanParams) GetAppIps() []string`

GetAppIps returns the AppIps field if non-nil, zero value otherwise.

### GetAppIpsOk

`func (o *ClusterVlanParams) GetAppIpsOk() (*[]string, bool)`

GetAppIpsOk returns a tuple with the AppIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIps

`func (o *ClusterVlanParams) SetAppIps(v []string)`

SetAppIps sets AppIps field to given value.

### HasAppIps

`func (o *ClusterVlanParams) HasAppIps() bool`

HasAppIps returns a boolean if a field has been set.

### SetAppIpsNil

`func (o *ClusterVlanParams) SetAppIpsNil(b bool)`

 SetAppIpsNil sets the value for AppIps to be an explicit nil

### UnsetAppIps
`func (o *ClusterVlanParams) UnsetAppIps()`

UnsetAppIps ensures that no value is present for AppIps, not even an explicit nil
### GetDescription

`func (o *ClusterVlanParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterVlanParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterVlanParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterVlanParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ClusterVlanParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ClusterVlanParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDnsDelegationZones

`func (o *ClusterVlanParams) GetDnsDelegationZones() []DnsDelegationZone`

GetDnsDelegationZones returns the DnsDelegationZones field if non-nil, zero value otherwise.

### GetDnsDelegationZonesOk

`func (o *ClusterVlanParams) GetDnsDelegationZonesOk() (*[]DnsDelegationZone, bool)`

GetDnsDelegationZonesOk returns a tuple with the DnsDelegationZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDelegationZones

`func (o *ClusterVlanParams) SetDnsDelegationZones(v []DnsDelegationZone)`

SetDnsDelegationZones sets DnsDelegationZones field to given value.

### HasDnsDelegationZones

`func (o *ClusterVlanParams) HasDnsDelegationZones() bool`

HasDnsDelegationZones returns a boolean if a field has been set.

### SetDnsDelegationZonesNil

`func (o *ClusterVlanParams) SetDnsDelegationZonesNil(b bool)`

 SetDnsDelegationZonesNil sets the value for DnsDelegationZones to be an explicit nil

### UnsetDnsDelegationZones
`func (o *ClusterVlanParams) UnsetDnsDelegationZones()`

UnsetDnsDelegationZones ensures that no value is present for DnsDelegationZones, not even an explicit nil
### GetEcmpEnabled

`func (o *ClusterVlanParams) GetEcmpEnabled() bool`

GetEcmpEnabled returns the EcmpEnabled field if non-nil, zero value otherwise.

### GetEcmpEnabledOk

`func (o *ClusterVlanParams) GetEcmpEnabledOk() (*bool, bool)`

GetEcmpEnabledOk returns a tuple with the EcmpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcmpEnabled

`func (o *ClusterVlanParams) SetEcmpEnabled(v bool)`

SetEcmpEnabled sets EcmpEnabled field to given value.

### HasEcmpEnabled

`func (o *ClusterVlanParams) HasEcmpEnabled() bool`

HasEcmpEnabled returns a boolean if a field has been set.

### SetEcmpEnabledNil

`func (o *ClusterVlanParams) SetEcmpEnabledNil(b bool)`

 SetEcmpEnabledNil sets the value for EcmpEnabled to be an explicit nil

### UnsetEcmpEnabled
`func (o *ClusterVlanParams) UnsetEcmpEnabled()`

UnsetEcmpEnabled ensures that no value is present for EcmpEnabled, not even an explicit nil
### GetFqdn

`func (o *ClusterVlanParams) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *ClusterVlanParams) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *ClusterVlanParams) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *ClusterVlanParams) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### SetFqdnNil

`func (o *ClusterVlanParams) SetFqdnNil(b bool)`

 SetFqdnNil sets the value for Fqdn to be an explicit nil

### UnsetFqdn
`func (o *ClusterVlanParams) UnsetFqdn()`

UnsetFqdn ensures that no value is present for Fqdn, not even an explicit nil
### GetGateway

`func (o *ClusterVlanParams) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *ClusterVlanParams) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *ClusterVlanParams) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *ClusterVlanParams) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *ClusterVlanParams) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *ClusterVlanParams) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetIpAddressesType

`func (o *ClusterVlanParams) GetIpAddressesType() string`

GetIpAddressesType returns the IpAddressesType field if non-nil, zero value otherwise.

### GetIpAddressesTypeOk

`func (o *ClusterVlanParams) GetIpAddressesTypeOk() (*string, bool)`

GetIpAddressesTypeOk returns a tuple with the IpAddressesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressesType

`func (o *ClusterVlanParams) SetIpAddressesType(v string)`

SetIpAddressesType sets IpAddressesType field to given value.

### HasIpAddressesType

`func (o *ClusterVlanParams) HasIpAddressesType() bool`

HasIpAddressesType returns a boolean if a field has been set.

### SetIpAddressesTypeNil

`func (o *ClusterVlanParams) SetIpAddressesTypeNil(b bool)`

 SetIpAddressesTypeNil sets the value for IpAddressesType to be an explicit nil

### UnsetIpAddressesType
`func (o *ClusterVlanParams) UnsetIpAddressesType()`

UnsetIpAddressesType ensures that no value is present for IpAddressesType, not even an explicit nil
### GetIpPools

`func (o *ClusterVlanParams) GetIpPools() []IpPool`

GetIpPools returns the IpPools field if non-nil, zero value otherwise.

### GetIpPoolsOk

`func (o *ClusterVlanParams) GetIpPoolsOk() (*[]IpPool, bool)`

GetIpPoolsOk returns a tuple with the IpPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPools

`func (o *ClusterVlanParams) SetIpPools(v []IpPool)`

SetIpPools sets IpPools field to given value.

### HasIpPools

`func (o *ClusterVlanParams) HasIpPools() bool`

HasIpPools returns a boolean if a field has been set.

### SetIpPoolsNil

`func (o *ClusterVlanParams) SetIpPoolsNil(b bool)`

 SetIpPoolsNil sets the value for IpPools to be an explicit nil

### UnsetIpPools
`func (o *ClusterVlanParams) UnsetIpPools()`

UnsetIpPools ensures that no value is present for IpPools, not even an explicit nil
### GetIpRanges

`func (o *ClusterVlanParams) GetIpRanges() []IpRange`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *ClusterVlanParams) GetIpRangesOk() (*[]IpRange, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *ClusterVlanParams) SetIpRanges(v []IpRange)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *ClusterVlanParams) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.

### SetIpRangesNil

`func (o *ClusterVlanParams) SetIpRangesNil(b bool)`

 SetIpRangesNil sets the value for IpRanges to be an explicit nil

### UnsetIpRanges
`func (o *ClusterVlanParams) UnsetIpRanges()`

UnsetIpRanges ensures that no value is present for IpRanges, not even an explicit nil
### GetIps

`func (o *ClusterVlanParams) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *ClusterVlanParams) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *ClusterVlanParams) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *ClusterVlanParams) HasIps() bool`

HasIps returns a boolean if a field has been set.

### SetIpsNil

`func (o *ClusterVlanParams) SetIpsNil(b bool)`

 SetIpsNil sets the value for Ips to be an explicit nil

### UnsetIps
`func (o *ClusterVlanParams) UnsetIps()`

UnsetIps ensures that no value is present for Ips, not even an explicit nil
### GetMtu

`func (o *ClusterVlanParams) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *ClusterVlanParams) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *ClusterVlanParams) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *ClusterVlanParams) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *ClusterVlanParams) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *ClusterVlanParams) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetSubnet

`func (o *ClusterVlanParams) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *ClusterVlanParams) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *ClusterVlanParams) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *ClusterVlanParams) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### SetSubnetNil

`func (o *ClusterVlanParams) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *ClusterVlanParams) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetTenantId

`func (o *ClusterVlanParams) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ClusterVlanParams) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ClusterVlanParams) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ClusterVlanParams) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ClusterVlanParams) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ClusterVlanParams) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetVlanName

`func (o *ClusterVlanParams) GetVlanName() string`

GetVlanName returns the VlanName field if non-nil, zero value otherwise.

### GetVlanNameOk

`func (o *ClusterVlanParams) GetVlanNameOk() (*string, bool)`

GetVlanNameOk returns a tuple with the VlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanName

`func (o *ClusterVlanParams) SetVlanName(v string)`

SetVlanName sets VlanName field to given value.

### HasVlanName

`func (o *ClusterVlanParams) HasVlanName() bool`

HasVlanName returns a boolean if a field has been set.

### SetVlanNameNil

`func (o *ClusterVlanParams) SetVlanNameNil(b bool)`

 SetVlanNameNil sets the value for VlanName to be an explicit nil

### UnsetVlanName
`func (o *ClusterVlanParams) UnsetVlanName()`

UnsetVlanName ensures that no value is present for VlanName, not even an explicit nil
### GetInterfaceName

`func (o *ClusterVlanParams) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *ClusterVlanParams) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *ClusterVlanParams) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.


### GetAppIpsInUse

`func (o *ClusterVlanParams) GetAppIpsInUse() bool`

GetAppIpsInUse returns the AppIpsInUse field if non-nil, zero value otherwise.

### GetAppIpsInUseOk

`func (o *ClusterVlanParams) GetAppIpsInUseOk() (*bool, bool)`

GetAppIpsInUseOk returns a tuple with the AppIpsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIpsInUse

`func (o *ClusterVlanParams) SetAppIpsInUse(v bool)`

SetAppIpsInUse sets AppIpsInUse field to given value.

### HasAppIpsInUse

`func (o *ClusterVlanParams) HasAppIpsInUse() bool`

HasAppIpsInUse returns a boolean if a field has been set.

### SetAppIpsInUseNil

`func (o *ClusterVlanParams) SetAppIpsInUseNil(b bool)`

 SetAppIpsInUseNil sets the value for AppIpsInUse to be an explicit nil

### UnsetAppIpsInUse
`func (o *ClusterVlanParams) UnsetAppIpsInUse()`

UnsetAppIpsInUse ensures that no value is present for AppIpsInUse, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


