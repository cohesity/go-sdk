# IpmiLanInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **NullableInt32** | Specifies the channel through which the IPMI interface communicates on the network. | [optional] 
**DefaultGatewayIp** | Pointer to **NullableString** | Specifies the default gateway ip for the given ipmi lan. | [optional] 
**DefaultGatewayMac** | Pointer to **NullableString** | Specifies the default gateway mac address for the given ipmi lan. | [optional] 
**IpAddrSource** | Pointer to **NullableString** | Specifies the ipmi lan address source. | [optional] 
**LanIp** | Pointer to **NullableString** | Specifies the ip address for the given ipmi lan. | [optional] 
**SubnetMask** | Pointer to **NullableString** | Specifies the subnet mask for the given ipmi lan. | [optional] 

## Methods

### NewIpmiLanInfo

`func NewIpmiLanInfo() *IpmiLanInfo`

NewIpmiLanInfo instantiates a new IpmiLanInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpmiLanInfoWithDefaults

`func NewIpmiLanInfoWithDefaults() *IpmiLanInfo`

NewIpmiLanInfoWithDefaults instantiates a new IpmiLanInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *IpmiLanInfo) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *IpmiLanInfo) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *IpmiLanInfo) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *IpmiLanInfo) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### SetChannelNil

`func (o *IpmiLanInfo) SetChannelNil(b bool)`

 SetChannelNil sets the value for Channel to be an explicit nil

### UnsetChannel
`func (o *IpmiLanInfo) UnsetChannel()`

UnsetChannel ensures that no value is present for Channel, not even an explicit nil
### GetDefaultGatewayIp

`func (o *IpmiLanInfo) GetDefaultGatewayIp() string`

GetDefaultGatewayIp returns the DefaultGatewayIp field if non-nil, zero value otherwise.

### GetDefaultGatewayIpOk

`func (o *IpmiLanInfo) GetDefaultGatewayIpOk() (*string, bool)`

GetDefaultGatewayIpOk returns a tuple with the DefaultGatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGatewayIp

`func (o *IpmiLanInfo) SetDefaultGatewayIp(v string)`

SetDefaultGatewayIp sets DefaultGatewayIp field to given value.

### HasDefaultGatewayIp

`func (o *IpmiLanInfo) HasDefaultGatewayIp() bool`

HasDefaultGatewayIp returns a boolean if a field has been set.

### SetDefaultGatewayIpNil

`func (o *IpmiLanInfo) SetDefaultGatewayIpNil(b bool)`

 SetDefaultGatewayIpNil sets the value for DefaultGatewayIp to be an explicit nil

### UnsetDefaultGatewayIp
`func (o *IpmiLanInfo) UnsetDefaultGatewayIp()`

UnsetDefaultGatewayIp ensures that no value is present for DefaultGatewayIp, not even an explicit nil
### GetDefaultGatewayMac

`func (o *IpmiLanInfo) GetDefaultGatewayMac() string`

GetDefaultGatewayMac returns the DefaultGatewayMac field if non-nil, zero value otherwise.

### GetDefaultGatewayMacOk

`func (o *IpmiLanInfo) GetDefaultGatewayMacOk() (*string, bool)`

GetDefaultGatewayMacOk returns a tuple with the DefaultGatewayMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGatewayMac

`func (o *IpmiLanInfo) SetDefaultGatewayMac(v string)`

SetDefaultGatewayMac sets DefaultGatewayMac field to given value.

### HasDefaultGatewayMac

`func (o *IpmiLanInfo) HasDefaultGatewayMac() bool`

HasDefaultGatewayMac returns a boolean if a field has been set.

### SetDefaultGatewayMacNil

`func (o *IpmiLanInfo) SetDefaultGatewayMacNil(b bool)`

 SetDefaultGatewayMacNil sets the value for DefaultGatewayMac to be an explicit nil

### UnsetDefaultGatewayMac
`func (o *IpmiLanInfo) UnsetDefaultGatewayMac()`

UnsetDefaultGatewayMac ensures that no value is present for DefaultGatewayMac, not even an explicit nil
### GetIpAddrSource

`func (o *IpmiLanInfo) GetIpAddrSource() string`

GetIpAddrSource returns the IpAddrSource field if non-nil, zero value otherwise.

### GetIpAddrSourceOk

`func (o *IpmiLanInfo) GetIpAddrSourceOk() (*string, bool)`

GetIpAddrSourceOk returns a tuple with the IpAddrSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddrSource

`func (o *IpmiLanInfo) SetIpAddrSource(v string)`

SetIpAddrSource sets IpAddrSource field to given value.

### HasIpAddrSource

`func (o *IpmiLanInfo) HasIpAddrSource() bool`

HasIpAddrSource returns a boolean if a field has been set.

### SetIpAddrSourceNil

`func (o *IpmiLanInfo) SetIpAddrSourceNil(b bool)`

 SetIpAddrSourceNil sets the value for IpAddrSource to be an explicit nil

### UnsetIpAddrSource
`func (o *IpmiLanInfo) UnsetIpAddrSource()`

UnsetIpAddrSource ensures that no value is present for IpAddrSource, not even an explicit nil
### GetLanIp

`func (o *IpmiLanInfo) GetLanIp() string`

GetLanIp returns the LanIp field if non-nil, zero value otherwise.

### GetLanIpOk

`func (o *IpmiLanInfo) GetLanIpOk() (*string, bool)`

GetLanIpOk returns a tuple with the LanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanIp

`func (o *IpmiLanInfo) SetLanIp(v string)`

SetLanIp sets LanIp field to given value.

### HasLanIp

`func (o *IpmiLanInfo) HasLanIp() bool`

HasLanIp returns a boolean if a field has been set.

### SetLanIpNil

`func (o *IpmiLanInfo) SetLanIpNil(b bool)`

 SetLanIpNil sets the value for LanIp to be an explicit nil

### UnsetLanIp
`func (o *IpmiLanInfo) UnsetLanIp()`

UnsetLanIp ensures that no value is present for LanIp, not even an explicit nil
### GetSubnetMask

`func (o *IpmiLanInfo) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *IpmiLanInfo) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *IpmiLanInfo) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.

### HasSubnetMask

`func (o *IpmiLanInfo) HasSubnetMask() bool`

HasSubnetMask returns a boolean if a field has been set.

### SetSubnetMaskNil

`func (o *IpmiLanInfo) SetSubnetMaskNil(b bool)`

 SetSubnetMaskNil sets the value for SubnetMask to be an explicit nil

### UnsetSubnetMask
`func (o *IpmiLanInfo) UnsetSubnetMask()`

UnsetSubnetMask ensures that no value is present for SubnetMask, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


