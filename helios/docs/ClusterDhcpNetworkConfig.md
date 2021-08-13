# ClusterDhcpNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateway** | Pointer to **NullableString** | Specifies the gateway of the new cluster network. | [optional] [readonly] 
**SubnetIp** | Pointer to **NullableString** | Specifies the ip subnet ip of the cluster network. | [optional] [readonly] 
**SubnetMask** | Pointer to **NullableString** | Specifies the ip subnet mask of the cluster network. | [optional] [readonly] 
**DnsServers** | **[]string** | Specifies the list of Dns Servers new cluster should be configured with. | 

## Methods

### NewClusterDhcpNetworkConfig

`func NewClusterDhcpNetworkConfig(dnsServers []string, ) *ClusterDhcpNetworkConfig`

NewClusterDhcpNetworkConfig instantiates a new ClusterDhcpNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDhcpNetworkConfigWithDefaults

`func NewClusterDhcpNetworkConfigWithDefaults() *ClusterDhcpNetworkConfig`

NewClusterDhcpNetworkConfigWithDefaults instantiates a new ClusterDhcpNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateway

`func (o *ClusterDhcpNetworkConfig) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *ClusterDhcpNetworkConfig) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *ClusterDhcpNetworkConfig) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *ClusterDhcpNetworkConfig) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *ClusterDhcpNetworkConfig) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *ClusterDhcpNetworkConfig) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetSubnetIp

`func (o *ClusterDhcpNetworkConfig) GetSubnetIp() string`

GetSubnetIp returns the SubnetIp field if non-nil, zero value otherwise.

### GetSubnetIpOk

`func (o *ClusterDhcpNetworkConfig) GetSubnetIpOk() (*string, bool)`

GetSubnetIpOk returns a tuple with the SubnetIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIp

`func (o *ClusterDhcpNetworkConfig) SetSubnetIp(v string)`

SetSubnetIp sets SubnetIp field to given value.

### HasSubnetIp

`func (o *ClusterDhcpNetworkConfig) HasSubnetIp() bool`

HasSubnetIp returns a boolean if a field has been set.

### SetSubnetIpNil

`func (o *ClusterDhcpNetworkConfig) SetSubnetIpNil(b bool)`

 SetSubnetIpNil sets the value for SubnetIp to be an explicit nil

### UnsetSubnetIp
`func (o *ClusterDhcpNetworkConfig) UnsetSubnetIp()`

UnsetSubnetIp ensures that no value is present for SubnetIp, not even an explicit nil
### GetSubnetMask

`func (o *ClusterDhcpNetworkConfig) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *ClusterDhcpNetworkConfig) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *ClusterDhcpNetworkConfig) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.

### HasSubnetMask

`func (o *ClusterDhcpNetworkConfig) HasSubnetMask() bool`

HasSubnetMask returns a boolean if a field has been set.

### SetSubnetMaskNil

`func (o *ClusterDhcpNetworkConfig) SetSubnetMaskNil(b bool)`

 SetSubnetMaskNil sets the value for SubnetMask to be an explicit nil

### UnsetSubnetMask
`func (o *ClusterDhcpNetworkConfig) UnsetSubnetMask()`

UnsetSubnetMask ensures that no value is present for SubnetMask, not even an explicit nil
### GetDnsServers

`func (o *ClusterDhcpNetworkConfig) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *ClusterDhcpNetworkConfig) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *ClusterDhcpNetworkConfig) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.


### SetDnsServersNil

`func (o *ClusterDhcpNetworkConfig) SetDnsServersNil(b bool)`

 SetDnsServersNil sets the value for DnsServers to be an explicit nil

### UnsetDnsServers
`func (o *ClusterDhcpNetworkConfig) UnsetDnsServers()`

UnsetDnsServers ensures that no value is present for DnsServers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


