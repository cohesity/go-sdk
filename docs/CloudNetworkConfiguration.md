# CloudNetworkConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterGateway** | Pointer to **NullableString** | Specifies the default gateway IP address (or addresses) for the Cluster network. | [optional] 
**ClusterSubnetMask** | Pointer to **NullableString** | Specifies the subnet mask (or masks) of the Cluster network. | [optional] 
**DnsServers** | Pointer to **[]string** | Specifies the list of DNS Servers this cluster should be configured with. | [optional] 
**DomainNames** | Pointer to **[]string** | Specifies the list of domain names this cluster should be configured with. | [optional] 
**NtpServers** | Pointer to **[]string** | Specifies the list of NTP Servers this cluster should be configured with. | [optional] 

## Methods

### NewCloudNetworkConfiguration

`func NewCloudNetworkConfiguration() *CloudNetworkConfiguration`

NewCloudNetworkConfiguration instantiates a new CloudNetworkConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudNetworkConfigurationWithDefaults

`func NewCloudNetworkConfigurationWithDefaults() *CloudNetworkConfiguration`

NewCloudNetworkConfigurationWithDefaults instantiates a new CloudNetworkConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterGateway

`func (o *CloudNetworkConfiguration) GetClusterGateway() string`

GetClusterGateway returns the ClusterGateway field if non-nil, zero value otherwise.

### GetClusterGatewayOk

`func (o *CloudNetworkConfiguration) GetClusterGatewayOk() (*string, bool)`

GetClusterGatewayOk returns a tuple with the ClusterGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterGateway

`func (o *CloudNetworkConfiguration) SetClusterGateway(v string)`

SetClusterGateway sets ClusterGateway field to given value.

### HasClusterGateway

`func (o *CloudNetworkConfiguration) HasClusterGateway() bool`

HasClusterGateway returns a boolean if a field has been set.

### SetClusterGatewayNil

`func (o *CloudNetworkConfiguration) SetClusterGatewayNil(b bool)`

 SetClusterGatewayNil sets the value for ClusterGateway to be an explicit nil

### UnsetClusterGateway
`func (o *CloudNetworkConfiguration) UnsetClusterGateway()`

UnsetClusterGateway ensures that no value is present for ClusterGateway, not even an explicit nil
### GetClusterSubnetMask

`func (o *CloudNetworkConfiguration) GetClusterSubnetMask() string`

GetClusterSubnetMask returns the ClusterSubnetMask field if non-nil, zero value otherwise.

### GetClusterSubnetMaskOk

`func (o *CloudNetworkConfiguration) GetClusterSubnetMaskOk() (*string, bool)`

GetClusterSubnetMaskOk returns a tuple with the ClusterSubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterSubnetMask

`func (o *CloudNetworkConfiguration) SetClusterSubnetMask(v string)`

SetClusterSubnetMask sets ClusterSubnetMask field to given value.

### HasClusterSubnetMask

`func (o *CloudNetworkConfiguration) HasClusterSubnetMask() bool`

HasClusterSubnetMask returns a boolean if a field has been set.

### SetClusterSubnetMaskNil

`func (o *CloudNetworkConfiguration) SetClusterSubnetMaskNil(b bool)`

 SetClusterSubnetMaskNil sets the value for ClusterSubnetMask to be an explicit nil

### UnsetClusterSubnetMask
`func (o *CloudNetworkConfiguration) UnsetClusterSubnetMask()`

UnsetClusterSubnetMask ensures that no value is present for ClusterSubnetMask, not even an explicit nil
### GetDnsServers

`func (o *CloudNetworkConfiguration) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *CloudNetworkConfiguration) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *CloudNetworkConfiguration) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *CloudNetworkConfiguration) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### SetDnsServersNil

`func (o *CloudNetworkConfiguration) SetDnsServersNil(b bool)`

 SetDnsServersNil sets the value for DnsServers to be an explicit nil

### UnsetDnsServers
`func (o *CloudNetworkConfiguration) UnsetDnsServers()`

UnsetDnsServers ensures that no value is present for DnsServers, not even an explicit nil
### GetDomainNames

`func (o *CloudNetworkConfiguration) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *CloudNetworkConfiguration) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *CloudNetworkConfiguration) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *CloudNetworkConfiguration) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### SetDomainNamesNil

`func (o *CloudNetworkConfiguration) SetDomainNamesNil(b bool)`

 SetDomainNamesNil sets the value for DomainNames to be an explicit nil

### UnsetDomainNames
`func (o *CloudNetworkConfiguration) UnsetDomainNames()`

UnsetDomainNames ensures that no value is present for DomainNames, not even an explicit nil
### GetNtpServers

`func (o *CloudNetworkConfiguration) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *CloudNetworkConfiguration) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *CloudNetworkConfiguration) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *CloudNetworkConfiguration) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### SetNtpServersNil

`func (o *CloudNetworkConfiguration) SetNtpServersNil(b bool)`

 SetNtpServersNil sets the value for NtpServers to be an explicit nil

### UnsetNtpServers
`func (o *CloudNetworkConfiguration) UnsetNtpServers()`

UnsetNtpServers ensures that no value is present for NtpServers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


