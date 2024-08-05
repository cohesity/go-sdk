# ClusterCreateNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpNetworkConfig** | Pointer to [**ClusterDhcpNetworkConfig**](ClusterDhcpNetworkConfig.md) |  | [optional] 
**DomainNames** | **[]string** | Specifies the list of Domain Names new cluster should be configured with. | 
**IpPreference** | Pointer to **NullableString** | Specifies IP preference of the cluster to be Ipv4/Ipv6. It is Ipv4 by default. | [optional] 
**ManualNetworkConfig** | Pointer to [**ClusterManualNetworkConfig**](ClusterManualNetworkConfig.md) |  | [optional] 
**NtpServers** | **[]string** | Specifies the list of NTP Servers new cluster should be configured with. | 
**SecondaryDhcpNetworkConfig** | Pointer to [**ClusterDhcpNetworkConfig**](ClusterDhcpNetworkConfig.md) |  | [optional] 
**SecondaryManualNetworkConfig** | Pointer to [**ClusterManualNetworkConfig**](ClusterManualNetworkConfig.md) |  | [optional] 
**UseDhcp** | **NullableBool** | Specifies whether or not to use DHCP to configure the network of the Cluster. | 
**VipHostName** | Pointer to **NullableString** | Specifies the FQDN hostname of the cluster. | [optional] 

## Methods

### NewClusterCreateNetworkConfig

`func NewClusterCreateNetworkConfig(domainNames []string, ntpServers []string, useDhcp NullableBool, ) *ClusterCreateNetworkConfig`

NewClusterCreateNetworkConfig instantiates a new ClusterCreateNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCreateNetworkConfigWithDefaults

`func NewClusterCreateNetworkConfigWithDefaults() *ClusterCreateNetworkConfig`

NewClusterCreateNetworkConfigWithDefaults instantiates a new ClusterCreateNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpNetworkConfig

`func (o *ClusterCreateNetworkConfig) GetDhcpNetworkConfig() ClusterDhcpNetworkConfig`

GetDhcpNetworkConfig returns the DhcpNetworkConfig field if non-nil, zero value otherwise.

### GetDhcpNetworkConfigOk

`func (o *ClusterCreateNetworkConfig) GetDhcpNetworkConfigOk() (*ClusterDhcpNetworkConfig, bool)`

GetDhcpNetworkConfigOk returns a tuple with the DhcpNetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpNetworkConfig

`func (o *ClusterCreateNetworkConfig) SetDhcpNetworkConfig(v ClusterDhcpNetworkConfig)`

SetDhcpNetworkConfig sets DhcpNetworkConfig field to given value.

### HasDhcpNetworkConfig

`func (o *ClusterCreateNetworkConfig) HasDhcpNetworkConfig() bool`

HasDhcpNetworkConfig returns a boolean if a field has been set.

### GetDomainNames

`func (o *ClusterCreateNetworkConfig) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *ClusterCreateNetworkConfig) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *ClusterCreateNetworkConfig) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.


### SetDomainNamesNil

`func (o *ClusterCreateNetworkConfig) SetDomainNamesNil(b bool)`

 SetDomainNamesNil sets the value for DomainNames to be an explicit nil

### UnsetDomainNames
`func (o *ClusterCreateNetworkConfig) UnsetDomainNames()`

UnsetDomainNames ensures that no value is present for DomainNames, not even an explicit nil
### GetIpPreference

`func (o *ClusterCreateNetworkConfig) GetIpPreference() string`

GetIpPreference returns the IpPreference field if non-nil, zero value otherwise.

### GetIpPreferenceOk

`func (o *ClusterCreateNetworkConfig) GetIpPreferenceOk() (*string, bool)`

GetIpPreferenceOk returns a tuple with the IpPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPreference

`func (o *ClusterCreateNetworkConfig) SetIpPreference(v string)`

SetIpPreference sets IpPreference field to given value.

### HasIpPreference

`func (o *ClusterCreateNetworkConfig) HasIpPreference() bool`

HasIpPreference returns a boolean if a field has been set.

### SetIpPreferenceNil

`func (o *ClusterCreateNetworkConfig) SetIpPreferenceNil(b bool)`

 SetIpPreferenceNil sets the value for IpPreference to be an explicit nil

### UnsetIpPreference
`func (o *ClusterCreateNetworkConfig) UnsetIpPreference()`

UnsetIpPreference ensures that no value is present for IpPreference, not even an explicit nil
### GetManualNetworkConfig

`func (o *ClusterCreateNetworkConfig) GetManualNetworkConfig() ClusterManualNetworkConfig`

GetManualNetworkConfig returns the ManualNetworkConfig field if non-nil, zero value otherwise.

### GetManualNetworkConfigOk

`func (o *ClusterCreateNetworkConfig) GetManualNetworkConfigOk() (*ClusterManualNetworkConfig, bool)`

GetManualNetworkConfigOk returns a tuple with the ManualNetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualNetworkConfig

`func (o *ClusterCreateNetworkConfig) SetManualNetworkConfig(v ClusterManualNetworkConfig)`

SetManualNetworkConfig sets ManualNetworkConfig field to given value.

### HasManualNetworkConfig

`func (o *ClusterCreateNetworkConfig) HasManualNetworkConfig() bool`

HasManualNetworkConfig returns a boolean if a field has been set.

### GetNtpServers

`func (o *ClusterCreateNetworkConfig) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *ClusterCreateNetworkConfig) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *ClusterCreateNetworkConfig) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.


### SetNtpServersNil

`func (o *ClusterCreateNetworkConfig) SetNtpServersNil(b bool)`

 SetNtpServersNil sets the value for NtpServers to be an explicit nil

### UnsetNtpServers
`func (o *ClusterCreateNetworkConfig) UnsetNtpServers()`

UnsetNtpServers ensures that no value is present for NtpServers, not even an explicit nil
### GetSecondaryDhcpNetworkConfig

`func (o *ClusterCreateNetworkConfig) GetSecondaryDhcpNetworkConfig() ClusterDhcpNetworkConfig`

GetSecondaryDhcpNetworkConfig returns the SecondaryDhcpNetworkConfig field if non-nil, zero value otherwise.

### GetSecondaryDhcpNetworkConfigOk

`func (o *ClusterCreateNetworkConfig) GetSecondaryDhcpNetworkConfigOk() (*ClusterDhcpNetworkConfig, bool)`

GetSecondaryDhcpNetworkConfigOk returns a tuple with the SecondaryDhcpNetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDhcpNetworkConfig

`func (o *ClusterCreateNetworkConfig) SetSecondaryDhcpNetworkConfig(v ClusterDhcpNetworkConfig)`

SetSecondaryDhcpNetworkConfig sets SecondaryDhcpNetworkConfig field to given value.

### HasSecondaryDhcpNetworkConfig

`func (o *ClusterCreateNetworkConfig) HasSecondaryDhcpNetworkConfig() bool`

HasSecondaryDhcpNetworkConfig returns a boolean if a field has been set.

### GetSecondaryManualNetworkConfig

`func (o *ClusterCreateNetworkConfig) GetSecondaryManualNetworkConfig() ClusterManualNetworkConfig`

GetSecondaryManualNetworkConfig returns the SecondaryManualNetworkConfig field if non-nil, zero value otherwise.

### GetSecondaryManualNetworkConfigOk

`func (o *ClusterCreateNetworkConfig) GetSecondaryManualNetworkConfigOk() (*ClusterManualNetworkConfig, bool)`

GetSecondaryManualNetworkConfigOk returns a tuple with the SecondaryManualNetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryManualNetworkConfig

`func (o *ClusterCreateNetworkConfig) SetSecondaryManualNetworkConfig(v ClusterManualNetworkConfig)`

SetSecondaryManualNetworkConfig sets SecondaryManualNetworkConfig field to given value.

### HasSecondaryManualNetworkConfig

`func (o *ClusterCreateNetworkConfig) HasSecondaryManualNetworkConfig() bool`

HasSecondaryManualNetworkConfig returns a boolean if a field has been set.

### GetUseDhcp

`func (o *ClusterCreateNetworkConfig) GetUseDhcp() bool`

GetUseDhcp returns the UseDhcp field if non-nil, zero value otherwise.

### GetUseDhcpOk

`func (o *ClusterCreateNetworkConfig) GetUseDhcpOk() (*bool, bool)`

GetUseDhcpOk returns a tuple with the UseDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDhcp

`func (o *ClusterCreateNetworkConfig) SetUseDhcp(v bool)`

SetUseDhcp sets UseDhcp field to given value.


### SetUseDhcpNil

`func (o *ClusterCreateNetworkConfig) SetUseDhcpNil(b bool)`

 SetUseDhcpNil sets the value for UseDhcp to be an explicit nil

### UnsetUseDhcp
`func (o *ClusterCreateNetworkConfig) UnsetUseDhcp()`

UnsetUseDhcp ensures that no value is present for UseDhcp, not even an explicit nil
### GetVipHostName

`func (o *ClusterCreateNetworkConfig) GetVipHostName() string`

GetVipHostName returns the VipHostName field if non-nil, zero value otherwise.

### GetVipHostNameOk

`func (o *ClusterCreateNetworkConfig) GetVipHostNameOk() (*string, bool)`

GetVipHostNameOk returns a tuple with the VipHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipHostName

`func (o *ClusterCreateNetworkConfig) SetVipHostName(v string)`

SetVipHostName sets VipHostName field to given value.

### HasVipHostName

`func (o *ClusterCreateNetworkConfig) HasVipHostName() bool`

HasVipHostName returns a boolean if a field has been set.

### SetVipHostNameNil

`func (o *ClusterCreateNetworkConfig) SetVipHostNameNil(b bool)`

 SetVipHostNameNil sets the value for VipHostName to be an explicit nil

### UnsetVipHostName
`func (o *ClusterCreateNetworkConfig) UnsetVipHostName()`

UnsetVipHostName ensures that no value is present for VipHostName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


