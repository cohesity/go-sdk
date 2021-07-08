# NetworkConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterGateway** | Pointer to **NullableString** | Specifies the default gateway IP address (or addresses) for the Cluster network. | [optional] 
**ClusterSubnetMask** | Pointer to **NullableString** | Specifies the subnet mask (or masks) of the Cluster network. | [optional] 
**DnsServers** | Pointer to **[]string** | Specifies the list of DNS Servers this cluster should be configured with. | [optional] 
**DomainNames** | Pointer to **[]string** | Specifies the list of domain names this cluster should be configured with. | [optional] 
**NtpServers** | Pointer to **[]string** | Specifies the list of NTP Servers this cluster should be configured with. | [optional] 
**VipHostname** | Pointer to **NullableString** | Specifies the virtual IP hostname. | [optional] 
**Vips** | Pointer to **[]string** | Specifies the list of virtual IPs for the new cluster. | [optional] 

## Methods

### NewNetworkConfiguration

`func NewNetworkConfiguration() *NetworkConfiguration`

NewNetworkConfiguration instantiates a new NetworkConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkConfigurationWithDefaults

`func NewNetworkConfigurationWithDefaults() *NetworkConfiguration`

NewNetworkConfigurationWithDefaults instantiates a new NetworkConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterGateway

`func (o *NetworkConfiguration) GetClusterGateway() string`

GetClusterGateway returns the ClusterGateway field if non-nil, zero value otherwise.

### GetClusterGatewayOk

`func (o *NetworkConfiguration) GetClusterGatewayOk() (*string, bool)`

GetClusterGatewayOk returns a tuple with the ClusterGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterGateway

`func (o *NetworkConfiguration) SetClusterGateway(v string)`

SetClusterGateway sets ClusterGateway field to given value.

### HasClusterGateway

`func (o *NetworkConfiguration) HasClusterGateway() bool`

HasClusterGateway returns a boolean if a field has been set.

### SetClusterGatewayNil

`func (o *NetworkConfiguration) SetClusterGatewayNil(b bool)`

 SetClusterGatewayNil sets the value for ClusterGateway to be an explicit nil

### UnsetClusterGateway
`func (o *NetworkConfiguration) UnsetClusterGateway()`

UnsetClusterGateway ensures that no value is present for ClusterGateway, not even an explicit nil
### GetClusterSubnetMask

`func (o *NetworkConfiguration) GetClusterSubnetMask() string`

GetClusterSubnetMask returns the ClusterSubnetMask field if non-nil, zero value otherwise.

### GetClusterSubnetMaskOk

`func (o *NetworkConfiguration) GetClusterSubnetMaskOk() (*string, bool)`

GetClusterSubnetMaskOk returns a tuple with the ClusterSubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterSubnetMask

`func (o *NetworkConfiguration) SetClusterSubnetMask(v string)`

SetClusterSubnetMask sets ClusterSubnetMask field to given value.

### HasClusterSubnetMask

`func (o *NetworkConfiguration) HasClusterSubnetMask() bool`

HasClusterSubnetMask returns a boolean if a field has been set.

### SetClusterSubnetMaskNil

`func (o *NetworkConfiguration) SetClusterSubnetMaskNil(b bool)`

 SetClusterSubnetMaskNil sets the value for ClusterSubnetMask to be an explicit nil

### UnsetClusterSubnetMask
`func (o *NetworkConfiguration) UnsetClusterSubnetMask()`

UnsetClusterSubnetMask ensures that no value is present for ClusterSubnetMask, not even an explicit nil
### GetDnsServers

`func (o *NetworkConfiguration) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *NetworkConfiguration) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *NetworkConfiguration) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *NetworkConfiguration) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### SetDnsServersNil

`func (o *NetworkConfiguration) SetDnsServersNil(b bool)`

 SetDnsServersNil sets the value for DnsServers to be an explicit nil

### UnsetDnsServers
`func (o *NetworkConfiguration) UnsetDnsServers()`

UnsetDnsServers ensures that no value is present for DnsServers, not even an explicit nil
### GetDomainNames

`func (o *NetworkConfiguration) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *NetworkConfiguration) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *NetworkConfiguration) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *NetworkConfiguration) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### SetDomainNamesNil

`func (o *NetworkConfiguration) SetDomainNamesNil(b bool)`

 SetDomainNamesNil sets the value for DomainNames to be an explicit nil

### UnsetDomainNames
`func (o *NetworkConfiguration) UnsetDomainNames()`

UnsetDomainNames ensures that no value is present for DomainNames, not even an explicit nil
### GetNtpServers

`func (o *NetworkConfiguration) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *NetworkConfiguration) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *NetworkConfiguration) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *NetworkConfiguration) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### SetNtpServersNil

`func (o *NetworkConfiguration) SetNtpServersNil(b bool)`

 SetNtpServersNil sets the value for NtpServers to be an explicit nil

### UnsetNtpServers
`func (o *NetworkConfiguration) UnsetNtpServers()`

UnsetNtpServers ensures that no value is present for NtpServers, not even an explicit nil
### GetVipHostname

`func (o *NetworkConfiguration) GetVipHostname() string`

GetVipHostname returns the VipHostname field if non-nil, zero value otherwise.

### GetVipHostnameOk

`func (o *NetworkConfiguration) GetVipHostnameOk() (*string, bool)`

GetVipHostnameOk returns a tuple with the VipHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipHostname

`func (o *NetworkConfiguration) SetVipHostname(v string)`

SetVipHostname sets VipHostname field to given value.

### HasVipHostname

`func (o *NetworkConfiguration) HasVipHostname() bool`

HasVipHostname returns a boolean if a field has been set.

### SetVipHostnameNil

`func (o *NetworkConfiguration) SetVipHostnameNil(b bool)`

 SetVipHostnameNil sets the value for VipHostname to be an explicit nil

### UnsetVipHostname
`func (o *NetworkConfiguration) UnsetVipHostname()`

UnsetVipHostname ensures that no value is present for VipHostname, not even an explicit nil
### GetVips

`func (o *NetworkConfiguration) GetVips() []string`

GetVips returns the Vips field if non-nil, zero value otherwise.

### GetVipsOk

`func (o *NetworkConfiguration) GetVipsOk() (*[]string, bool)`

GetVipsOk returns a tuple with the Vips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVips

`func (o *NetworkConfiguration) SetVips(v []string)`

SetVips sets Vips field to given value.

### HasVips

`func (o *NetworkConfiguration) HasVips() bool`

HasVips returns a boolean if a field has been set.

### SetVipsNil

`func (o *NetworkConfiguration) SetVipsNil(b bool)`

 SetVipsNil sets the value for Vips to be an explicit nil

### UnsetVips
`func (o *NetworkConfiguration) UnsetVips()`

UnsetVips ensures that no value is present for Vips, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


