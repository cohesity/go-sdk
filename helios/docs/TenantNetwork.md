# TenantNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorEnabled** | **NullableBool** | Whether connector (hybrid extender) is enabled. | 
**ClusterHostname** | Pointer to **NullableString** | The hostname for Cohesity cluster as seen by tenants and as is routable from the tenant&#39;s network. Tenant&#39;s VLAN&#39;s hostname, if available can be used instead but it is mandatory to provide this value if there&#39;s no VLAN hostname to use. Also, when set, this field would take precedence over VLAN hostname. | [optional] 
**ClusterIps** | Pointer to **[]string** | Set of IPs as seen from the tenant&#39;s network for the Cohesity cluster. Only one from &#39;clusterHostname&#39; and &#39;clusterIps&#39; is needed. | [optional] 

## Methods

### NewTenantNetwork

`func NewTenantNetwork(connectorEnabled NullableBool, ) *TenantNetwork`

NewTenantNetwork instantiates a new TenantNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantNetworkWithDefaults

`func NewTenantNetworkWithDefaults() *TenantNetwork`

NewTenantNetworkWithDefaults instantiates a new TenantNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorEnabled

`func (o *TenantNetwork) GetConnectorEnabled() bool`

GetConnectorEnabled returns the ConnectorEnabled field if non-nil, zero value otherwise.

### GetConnectorEnabledOk

`func (o *TenantNetwork) GetConnectorEnabledOk() (*bool, bool)`

GetConnectorEnabledOk returns a tuple with the ConnectorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorEnabled

`func (o *TenantNetwork) SetConnectorEnabled(v bool)`

SetConnectorEnabled sets ConnectorEnabled field to given value.


### SetConnectorEnabledNil

`func (o *TenantNetwork) SetConnectorEnabledNil(b bool)`

 SetConnectorEnabledNil sets the value for ConnectorEnabled to be an explicit nil

### UnsetConnectorEnabled
`func (o *TenantNetwork) UnsetConnectorEnabled()`

UnsetConnectorEnabled ensures that no value is present for ConnectorEnabled, not even an explicit nil
### GetClusterHostname

`func (o *TenantNetwork) GetClusterHostname() string`

GetClusterHostname returns the ClusterHostname field if non-nil, zero value otherwise.

### GetClusterHostnameOk

`func (o *TenantNetwork) GetClusterHostnameOk() (*string, bool)`

GetClusterHostnameOk returns a tuple with the ClusterHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterHostname

`func (o *TenantNetwork) SetClusterHostname(v string)`

SetClusterHostname sets ClusterHostname field to given value.

### HasClusterHostname

`func (o *TenantNetwork) HasClusterHostname() bool`

HasClusterHostname returns a boolean if a field has been set.

### SetClusterHostnameNil

`func (o *TenantNetwork) SetClusterHostnameNil(b bool)`

 SetClusterHostnameNil sets the value for ClusterHostname to be an explicit nil

### UnsetClusterHostname
`func (o *TenantNetwork) UnsetClusterHostname()`

UnsetClusterHostname ensures that no value is present for ClusterHostname, not even an explicit nil
### GetClusterIps

`func (o *TenantNetwork) GetClusterIps() []string`

GetClusterIps returns the ClusterIps field if non-nil, zero value otherwise.

### GetClusterIpsOk

`func (o *TenantNetwork) GetClusterIpsOk() (*[]string, bool)`

GetClusterIpsOk returns a tuple with the ClusterIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIps

`func (o *TenantNetwork) SetClusterIps(v []string)`

SetClusterIps sets ClusterIps field to given value.

### HasClusterIps

`func (o *TenantNetwork) HasClusterIps() bool`

HasClusterIps returns a boolean if a field has been set.

### SetClusterIpsNil

`func (o *TenantNetwork) SetClusterIpsNil(b bool)`

 SetClusterIpsNil sets the value for ClusterIps to be an explicit nil

### UnsetClusterIps
`func (o *TenantNetwork) UnsetClusterIps()`

UnsetClusterIps ensures that no value is present for ClusterIps, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


