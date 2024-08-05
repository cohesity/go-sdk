# UpdateClusterVlanParams

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

## Methods

### NewUpdateClusterVlanParams

`func NewUpdateClusterVlanParams() *UpdateClusterVlanParams`

NewUpdateClusterVlanParams instantiates a new UpdateClusterVlanParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterVlanParamsWithDefaults

`func NewUpdateClusterVlanParamsWithDefaults() *UpdateClusterVlanParams`

NewUpdateClusterVlanParamsWithDefaults instantiates a new UpdateClusterVlanParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllTenantAccess

`func (o *UpdateClusterVlanParams) GetAllTenantAccess() bool`

GetAllTenantAccess returns the AllTenantAccess field if non-nil, zero value otherwise.

### GetAllTenantAccessOk

`func (o *UpdateClusterVlanParams) GetAllTenantAccessOk() (*bool, bool)`

GetAllTenantAccessOk returns a tuple with the AllTenantAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllTenantAccess

`func (o *UpdateClusterVlanParams) SetAllTenantAccess(v bool)`

SetAllTenantAccess sets AllTenantAccess field to given value.

### HasAllTenantAccess

`func (o *UpdateClusterVlanParams) HasAllTenantAccess() bool`

HasAllTenantAccess returns a boolean if a field has been set.

### SetAllTenantAccessNil

`func (o *UpdateClusterVlanParams) SetAllTenantAccessNil(b bool)`

 SetAllTenantAccessNil sets the value for AllTenantAccess to be an explicit nil

### UnsetAllTenantAccess
`func (o *UpdateClusterVlanParams) UnsetAllTenantAccess()`

UnsetAllTenantAccess ensures that no value is present for AllTenantAccess, not even an explicit nil
### GetAppIps

`func (o *UpdateClusterVlanParams) GetAppIps() []string`

GetAppIps returns the AppIps field if non-nil, zero value otherwise.

### GetAppIpsOk

`func (o *UpdateClusterVlanParams) GetAppIpsOk() (*[]string, bool)`

GetAppIpsOk returns a tuple with the AppIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIps

`func (o *UpdateClusterVlanParams) SetAppIps(v []string)`

SetAppIps sets AppIps field to given value.

### HasAppIps

`func (o *UpdateClusterVlanParams) HasAppIps() bool`

HasAppIps returns a boolean if a field has been set.

### SetAppIpsNil

`func (o *UpdateClusterVlanParams) SetAppIpsNil(b bool)`

 SetAppIpsNil sets the value for AppIps to be an explicit nil

### UnsetAppIps
`func (o *UpdateClusterVlanParams) UnsetAppIps()`

UnsetAppIps ensures that no value is present for AppIps, not even an explicit nil
### GetDescription

`func (o *UpdateClusterVlanParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateClusterVlanParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateClusterVlanParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateClusterVlanParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateClusterVlanParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateClusterVlanParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDnsDelegationZones

`func (o *UpdateClusterVlanParams) GetDnsDelegationZones() []DnsDelegationZone`

GetDnsDelegationZones returns the DnsDelegationZones field if non-nil, zero value otherwise.

### GetDnsDelegationZonesOk

`func (o *UpdateClusterVlanParams) GetDnsDelegationZonesOk() (*[]DnsDelegationZone, bool)`

GetDnsDelegationZonesOk returns a tuple with the DnsDelegationZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDelegationZones

`func (o *UpdateClusterVlanParams) SetDnsDelegationZones(v []DnsDelegationZone)`

SetDnsDelegationZones sets DnsDelegationZones field to given value.

### HasDnsDelegationZones

`func (o *UpdateClusterVlanParams) HasDnsDelegationZones() bool`

HasDnsDelegationZones returns a boolean if a field has been set.

### SetDnsDelegationZonesNil

`func (o *UpdateClusterVlanParams) SetDnsDelegationZonesNil(b bool)`

 SetDnsDelegationZonesNil sets the value for DnsDelegationZones to be an explicit nil

### UnsetDnsDelegationZones
`func (o *UpdateClusterVlanParams) UnsetDnsDelegationZones()`

UnsetDnsDelegationZones ensures that no value is present for DnsDelegationZones, not even an explicit nil
### GetEcmpEnabled

`func (o *UpdateClusterVlanParams) GetEcmpEnabled() bool`

GetEcmpEnabled returns the EcmpEnabled field if non-nil, zero value otherwise.

### GetEcmpEnabledOk

`func (o *UpdateClusterVlanParams) GetEcmpEnabledOk() (*bool, bool)`

GetEcmpEnabledOk returns a tuple with the EcmpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcmpEnabled

`func (o *UpdateClusterVlanParams) SetEcmpEnabled(v bool)`

SetEcmpEnabled sets EcmpEnabled field to given value.

### HasEcmpEnabled

`func (o *UpdateClusterVlanParams) HasEcmpEnabled() bool`

HasEcmpEnabled returns a boolean if a field has been set.

### SetEcmpEnabledNil

`func (o *UpdateClusterVlanParams) SetEcmpEnabledNil(b bool)`

 SetEcmpEnabledNil sets the value for EcmpEnabled to be an explicit nil

### UnsetEcmpEnabled
`func (o *UpdateClusterVlanParams) UnsetEcmpEnabled()`

UnsetEcmpEnabled ensures that no value is present for EcmpEnabled, not even an explicit nil
### GetFqdn

`func (o *UpdateClusterVlanParams) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *UpdateClusterVlanParams) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *UpdateClusterVlanParams) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *UpdateClusterVlanParams) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### SetFqdnNil

`func (o *UpdateClusterVlanParams) SetFqdnNil(b bool)`

 SetFqdnNil sets the value for Fqdn to be an explicit nil

### UnsetFqdn
`func (o *UpdateClusterVlanParams) UnsetFqdn()`

UnsetFqdn ensures that no value is present for Fqdn, not even an explicit nil
### GetGateway

`func (o *UpdateClusterVlanParams) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UpdateClusterVlanParams) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UpdateClusterVlanParams) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UpdateClusterVlanParams) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *UpdateClusterVlanParams) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *UpdateClusterVlanParams) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetIpAddressesType

`func (o *UpdateClusterVlanParams) GetIpAddressesType() string`

GetIpAddressesType returns the IpAddressesType field if non-nil, zero value otherwise.

### GetIpAddressesTypeOk

`func (o *UpdateClusterVlanParams) GetIpAddressesTypeOk() (*string, bool)`

GetIpAddressesTypeOk returns a tuple with the IpAddressesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressesType

`func (o *UpdateClusterVlanParams) SetIpAddressesType(v string)`

SetIpAddressesType sets IpAddressesType field to given value.

### HasIpAddressesType

`func (o *UpdateClusterVlanParams) HasIpAddressesType() bool`

HasIpAddressesType returns a boolean if a field has been set.

### SetIpAddressesTypeNil

`func (o *UpdateClusterVlanParams) SetIpAddressesTypeNil(b bool)`

 SetIpAddressesTypeNil sets the value for IpAddressesType to be an explicit nil

### UnsetIpAddressesType
`func (o *UpdateClusterVlanParams) UnsetIpAddressesType()`

UnsetIpAddressesType ensures that no value is present for IpAddressesType, not even an explicit nil
### GetIpPools

`func (o *UpdateClusterVlanParams) GetIpPools() []IpPool`

GetIpPools returns the IpPools field if non-nil, zero value otherwise.

### GetIpPoolsOk

`func (o *UpdateClusterVlanParams) GetIpPoolsOk() (*[]IpPool, bool)`

GetIpPoolsOk returns a tuple with the IpPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPools

`func (o *UpdateClusterVlanParams) SetIpPools(v []IpPool)`

SetIpPools sets IpPools field to given value.

### HasIpPools

`func (o *UpdateClusterVlanParams) HasIpPools() bool`

HasIpPools returns a boolean if a field has been set.

### SetIpPoolsNil

`func (o *UpdateClusterVlanParams) SetIpPoolsNil(b bool)`

 SetIpPoolsNil sets the value for IpPools to be an explicit nil

### UnsetIpPools
`func (o *UpdateClusterVlanParams) UnsetIpPools()`

UnsetIpPools ensures that no value is present for IpPools, not even an explicit nil
### GetIpRanges

`func (o *UpdateClusterVlanParams) GetIpRanges() []IpRange`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *UpdateClusterVlanParams) GetIpRangesOk() (*[]IpRange, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *UpdateClusterVlanParams) SetIpRanges(v []IpRange)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *UpdateClusterVlanParams) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.

### SetIpRangesNil

`func (o *UpdateClusterVlanParams) SetIpRangesNil(b bool)`

 SetIpRangesNil sets the value for IpRanges to be an explicit nil

### UnsetIpRanges
`func (o *UpdateClusterVlanParams) UnsetIpRanges()`

UnsetIpRanges ensures that no value is present for IpRanges, not even an explicit nil
### GetIps

`func (o *UpdateClusterVlanParams) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *UpdateClusterVlanParams) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *UpdateClusterVlanParams) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *UpdateClusterVlanParams) HasIps() bool`

HasIps returns a boolean if a field has been set.

### SetIpsNil

`func (o *UpdateClusterVlanParams) SetIpsNil(b bool)`

 SetIpsNil sets the value for Ips to be an explicit nil

### UnsetIps
`func (o *UpdateClusterVlanParams) UnsetIps()`

UnsetIps ensures that no value is present for Ips, not even an explicit nil
### GetMtu

`func (o *UpdateClusterVlanParams) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *UpdateClusterVlanParams) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *UpdateClusterVlanParams) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *UpdateClusterVlanParams) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *UpdateClusterVlanParams) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *UpdateClusterVlanParams) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetSubnet

`func (o *UpdateClusterVlanParams) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *UpdateClusterVlanParams) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *UpdateClusterVlanParams) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *UpdateClusterVlanParams) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### SetSubnetNil

`func (o *UpdateClusterVlanParams) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *UpdateClusterVlanParams) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetTenantId

`func (o *UpdateClusterVlanParams) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UpdateClusterVlanParams) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UpdateClusterVlanParams) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *UpdateClusterVlanParams) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *UpdateClusterVlanParams) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *UpdateClusterVlanParams) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetVlanName

`func (o *UpdateClusterVlanParams) GetVlanName() string`

GetVlanName returns the VlanName field if non-nil, zero value otherwise.

### GetVlanNameOk

`func (o *UpdateClusterVlanParams) GetVlanNameOk() (*string, bool)`

GetVlanNameOk returns a tuple with the VlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanName

`func (o *UpdateClusterVlanParams) SetVlanName(v string)`

SetVlanName sets VlanName field to given value.

### HasVlanName

`func (o *UpdateClusterVlanParams) HasVlanName() bool`

HasVlanName returns a boolean if a field has been set.

### SetVlanNameNil

`func (o *UpdateClusterVlanParams) SetVlanNameNil(b bool)`

 SetVlanNameNil sets the value for VlanName to be an explicit nil

### UnsetVlanName
`func (o *UpdateClusterVlanParams) UnsetVlanName()`

UnsetVlanName ensures that no value is present for VlanName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


