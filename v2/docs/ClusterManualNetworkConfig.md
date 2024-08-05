# ClusterManualNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsServers** | **[]string** | Specifies the list of Dns Servers new cluster should be configured with. | 
**Gateway** | **NullableString** | Specifies the gateway of the new cluster network. | 
**SubnetIp** | **NullableString** | Specifies the ip subnet ip of the cluster network. | 
**SubnetMask** | **NullableString** | Specifies the ip subnet mask of the cluster network. | 

## Methods

### NewClusterManualNetworkConfig

`func NewClusterManualNetworkConfig(dnsServers []string, gateway NullableString, subnetIp NullableString, subnetMask NullableString, ) *ClusterManualNetworkConfig`

NewClusterManualNetworkConfig instantiates a new ClusterManualNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterManualNetworkConfigWithDefaults

`func NewClusterManualNetworkConfigWithDefaults() *ClusterManualNetworkConfig`

NewClusterManualNetworkConfigWithDefaults instantiates a new ClusterManualNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsServers

`func (o *ClusterManualNetworkConfig) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *ClusterManualNetworkConfig) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *ClusterManualNetworkConfig) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.


### SetDnsServersNil

`func (o *ClusterManualNetworkConfig) SetDnsServersNil(b bool)`

 SetDnsServersNil sets the value for DnsServers to be an explicit nil

### UnsetDnsServers
`func (o *ClusterManualNetworkConfig) UnsetDnsServers()`

UnsetDnsServers ensures that no value is present for DnsServers, not even an explicit nil
### GetGateway

`func (o *ClusterManualNetworkConfig) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *ClusterManualNetworkConfig) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *ClusterManualNetworkConfig) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### SetGatewayNil

`func (o *ClusterManualNetworkConfig) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *ClusterManualNetworkConfig) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetSubnetIp

`func (o *ClusterManualNetworkConfig) GetSubnetIp() string`

GetSubnetIp returns the SubnetIp field if non-nil, zero value otherwise.

### GetSubnetIpOk

`func (o *ClusterManualNetworkConfig) GetSubnetIpOk() (*string, bool)`

GetSubnetIpOk returns a tuple with the SubnetIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIp

`func (o *ClusterManualNetworkConfig) SetSubnetIp(v string)`

SetSubnetIp sets SubnetIp field to given value.


### SetSubnetIpNil

`func (o *ClusterManualNetworkConfig) SetSubnetIpNil(b bool)`

 SetSubnetIpNil sets the value for SubnetIp to be an explicit nil

### UnsetSubnetIp
`func (o *ClusterManualNetworkConfig) UnsetSubnetIp()`

UnsetSubnetIp ensures that no value is present for SubnetIp, not even an explicit nil
### GetSubnetMask

`func (o *ClusterManualNetworkConfig) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *ClusterManualNetworkConfig) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *ClusterManualNetworkConfig) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.


### SetSubnetMaskNil

`func (o *ClusterManualNetworkConfig) SetSubnetMaskNil(b bool)`

 SetSubnetMaskNil sets the value for SubnetMask to be an explicit nil

### UnsetSubnetMask
`func (o *ClusterManualNetworkConfig) UnsetSubnetMask()`

UnsetSubnetMask ensures that no value is present for SubnetMask, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


