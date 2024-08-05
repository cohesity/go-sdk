# FleetConfigFleetNetworkParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnet** | **NullableString** | Specifies subnet for the fleet. | 
**Vpc** | **NullableString** | Specifies vpc for the fleet. | 

## Methods

### NewFleetConfigFleetNetworkParams

`func NewFleetConfigFleetNetworkParams(subnet NullableString, vpc NullableString, ) *FleetConfigFleetNetworkParams`

NewFleetConfigFleetNetworkParams instantiates a new FleetConfigFleetNetworkParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetConfigFleetNetworkParamsWithDefaults

`func NewFleetConfigFleetNetworkParamsWithDefaults() *FleetConfigFleetNetworkParams`

NewFleetConfigFleetNetworkParamsWithDefaults instantiates a new FleetConfigFleetNetworkParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnet

`func (o *FleetConfigFleetNetworkParams) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *FleetConfigFleetNetworkParams) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *FleetConfigFleetNetworkParams) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.


### SetSubnetNil

`func (o *FleetConfigFleetNetworkParams) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *FleetConfigFleetNetworkParams) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetVpc

`func (o *FleetConfigFleetNetworkParams) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *FleetConfigFleetNetworkParams) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *FleetConfigFleetNetworkParams) SetVpc(v string)`

SetVpc sets Vpc field to given value.


### SetVpcNil

`func (o *FleetConfigFleetNetworkParams) SetVpcNil(b bool)`

 SetVpcNil sets the value for Vpc to be an explicit nil

### UnsetVpc
`func (o *FleetConfigFleetNetworkParams) UnsetVpc()`

UnsetVpc ensures that no value is present for Vpc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


