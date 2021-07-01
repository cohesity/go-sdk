# FleetNetworkParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **NullableString** | Specifies the region for the fleet. | [optional] 
**Subnet** | Pointer to **NullableString** | Specifies the subnet for the fleet. | [optional] 
**Vpc** | Pointer to **NullableString** | Specifies the vpc for the fleet. | [optional] 

## Methods

### NewFleetNetworkParams

`func NewFleetNetworkParams() *FleetNetworkParams`

NewFleetNetworkParams instantiates a new FleetNetworkParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetNetworkParamsWithDefaults

`func NewFleetNetworkParamsWithDefaults() *FleetNetworkParams`

NewFleetNetworkParamsWithDefaults instantiates a new FleetNetworkParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *FleetNetworkParams) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FleetNetworkParams) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FleetNetworkParams) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *FleetNetworkParams) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *FleetNetworkParams) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *FleetNetworkParams) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetSubnet

`func (o *FleetNetworkParams) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *FleetNetworkParams) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *FleetNetworkParams) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *FleetNetworkParams) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### SetSubnetNil

`func (o *FleetNetworkParams) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *FleetNetworkParams) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetVpc

`func (o *FleetNetworkParams) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *FleetNetworkParams) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *FleetNetworkParams) SetVpc(v string)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *FleetNetworkParams) HasVpc() bool`

HasVpc returns a boolean if a field has been set.

### SetVpcNil

`func (o *FleetNetworkParams) SetVpcNil(b bool)`

 SetVpcNil sets the value for Vpc to be an explicit nil

### UnsetVpc
`func (o *FleetNetworkParams) UnsetVpc()`

UnsetVpc ensures that no value is present for Vpc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


