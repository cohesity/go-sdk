# NetworkConnectionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dns** | Pointer to **NullableString** | Specifies the DNS Server of the network connection. | [optional] 
**DomainName** | Pointer to **NullableString** | Specifies the domain name of the network connection. | [optional] 
**NetworkGateway** | Pointer to **NullableString** | Specifies the network Gateway of the network connection. | [optional] 
**Ntp** | Pointer to **NullableString** | Specifies the NTP Server of the network connection. | [optional] 

## Methods

### NewNetworkConnectionInfo

`func NewNetworkConnectionInfo() *NetworkConnectionInfo`

NewNetworkConnectionInfo instantiates a new NetworkConnectionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkConnectionInfoWithDefaults

`func NewNetworkConnectionInfoWithDefaults() *NetworkConnectionInfo`

NewNetworkConnectionInfoWithDefaults instantiates a new NetworkConnectionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDns

`func (o *NetworkConnectionInfo) GetDns() string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *NetworkConnectionInfo) GetDnsOk() (*string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *NetworkConnectionInfo) SetDns(v string)`

SetDns sets Dns field to given value.

### HasDns

`func (o *NetworkConnectionInfo) HasDns() bool`

HasDns returns a boolean if a field has been set.

### SetDnsNil

`func (o *NetworkConnectionInfo) SetDnsNil(b bool)`

 SetDnsNil sets the value for Dns to be an explicit nil

### UnsetDns
`func (o *NetworkConnectionInfo) UnsetDns()`

UnsetDns ensures that no value is present for Dns, not even an explicit nil
### GetDomainName

`func (o *NetworkConnectionInfo) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *NetworkConnectionInfo) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *NetworkConnectionInfo) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *NetworkConnectionInfo) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *NetworkConnectionInfo) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *NetworkConnectionInfo) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetNetworkGateway

`func (o *NetworkConnectionInfo) GetNetworkGateway() string`

GetNetworkGateway returns the NetworkGateway field if non-nil, zero value otherwise.

### GetNetworkGatewayOk

`func (o *NetworkConnectionInfo) GetNetworkGatewayOk() (*string, bool)`

GetNetworkGatewayOk returns a tuple with the NetworkGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkGateway

`func (o *NetworkConnectionInfo) SetNetworkGateway(v string)`

SetNetworkGateway sets NetworkGateway field to given value.

### HasNetworkGateway

`func (o *NetworkConnectionInfo) HasNetworkGateway() bool`

HasNetworkGateway returns a boolean if a field has been set.

### SetNetworkGatewayNil

`func (o *NetworkConnectionInfo) SetNetworkGatewayNil(b bool)`

 SetNetworkGatewayNil sets the value for NetworkGateway to be an explicit nil

### UnsetNetworkGateway
`func (o *NetworkConnectionInfo) UnsetNetworkGateway()`

UnsetNetworkGateway ensures that no value is present for NetworkGateway, not even an explicit nil
### GetNtp

`func (o *NetworkConnectionInfo) GetNtp() string`

GetNtp returns the Ntp field if non-nil, zero value otherwise.

### GetNtpOk

`func (o *NetworkConnectionInfo) GetNtpOk() (*string, bool)`

GetNtpOk returns a tuple with the Ntp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtp

`func (o *NetworkConnectionInfo) SetNtp(v string)`

SetNtp sets Ntp field to given value.

### HasNtp

`func (o *NetworkConnectionInfo) HasNtp() bool`

HasNtp returns a boolean if a field has been set.

### SetNtpNil

`func (o *NetworkConnectionInfo) SetNtpNil(b bool)`

 SetNtpNil sets the value for Ntp to be an explicit nil

### UnsetNtp
`func (o *NetworkConnectionInfo) UnsetNtp()`

UnsetNtp ensures that no value is present for Ntp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


