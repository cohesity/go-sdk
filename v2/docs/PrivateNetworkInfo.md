# PrivateNetworkInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to **NullableString** | Specifies the subnet for creating a private endpoint. | [optional] 
**Region** | Pointer to [**NullablePrivateNetworkInfoRegion**](PrivateNetworkInfoRegion.md) |  | [optional] 
**Subnet** | Pointer to [**NullablePrivateNetworkInfoSubnet**](PrivateNetworkInfoSubnet.md) |  | [optional] 
**Vpn** | Pointer to [**NullablePrivateNetworkInfoVpn**](PrivateNetworkInfoVpn.md) |  | [optional] 

## Methods

### NewPrivateNetworkInfo

`func NewPrivateNetworkInfo() *PrivateNetworkInfo`

NewPrivateNetworkInfo instantiates a new PrivateNetworkInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateNetworkInfoWithDefaults

`func NewPrivateNetworkInfoWithDefaults() *PrivateNetworkInfo`

NewPrivateNetworkInfoWithDefaults instantiates a new PrivateNetworkInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *PrivateNetworkInfo) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PrivateNetworkInfo) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PrivateNetworkInfo) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PrivateNetworkInfo) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *PrivateNetworkInfo) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *PrivateNetworkInfo) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetRegion

`func (o *PrivateNetworkInfo) GetRegion() PrivateNetworkInfoRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PrivateNetworkInfo) GetRegionOk() (*PrivateNetworkInfoRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PrivateNetworkInfo) SetRegion(v PrivateNetworkInfoRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PrivateNetworkInfo) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *PrivateNetworkInfo) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *PrivateNetworkInfo) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetSubnet

`func (o *PrivateNetworkInfo) GetSubnet() PrivateNetworkInfoSubnet`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *PrivateNetworkInfo) GetSubnetOk() (*PrivateNetworkInfoSubnet, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *PrivateNetworkInfo) SetSubnet(v PrivateNetworkInfoSubnet)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *PrivateNetworkInfo) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### SetSubnetNil

`func (o *PrivateNetworkInfo) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *PrivateNetworkInfo) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetVpn

`func (o *PrivateNetworkInfo) GetVpn() PrivateNetworkInfoVpn`

GetVpn returns the Vpn field if non-nil, zero value otherwise.

### GetVpnOk

`func (o *PrivateNetworkInfo) GetVpnOk() (*PrivateNetworkInfoVpn, bool)`

GetVpnOk returns a tuple with the Vpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpn

`func (o *PrivateNetworkInfo) SetVpn(v PrivateNetworkInfoVpn)`

SetVpn sets Vpn field to given value.

### HasVpn

`func (o *PrivateNetworkInfo) HasVpn() bool`

HasVpn returns a boolean if a field has been set.

### SetVpnNil

`func (o *PrivateNetworkInfo) SetVpnNil(b bool)`

 SetVpnNil sets the value for Vpn to be an explicit nil

### UnsetVpn
`func (o *PrivateNetworkInfo) UnsetVpn()`

UnsetVpn ensures that no value is present for Vpn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


