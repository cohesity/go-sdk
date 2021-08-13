# RecoverGcpVmNewSourceNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnet** | [**NullableGcpVpcSubnetConfig**](GcpVpcSubnetConfig.md) | Specifies the subnet. | 

## Methods

### NewRecoverGcpVmNewSourceNetworkConfig

`func NewRecoverGcpVmNewSourceNetworkConfig(subnet NullableGcpVpcSubnetConfig, ) *RecoverGcpVmNewSourceNetworkConfig`

NewRecoverGcpVmNewSourceNetworkConfig instantiates a new RecoverGcpVmNewSourceNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverGcpVmNewSourceNetworkConfigWithDefaults

`func NewRecoverGcpVmNewSourceNetworkConfigWithDefaults() *RecoverGcpVmNewSourceNetworkConfig`

NewRecoverGcpVmNewSourceNetworkConfigWithDefaults instantiates a new RecoverGcpVmNewSourceNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnet

`func (o *RecoverGcpVmNewSourceNetworkConfig) GetSubnet() GcpVpcSubnetConfig`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *RecoverGcpVmNewSourceNetworkConfig) GetSubnetOk() (*GcpVpcSubnetConfig, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *RecoverGcpVmNewSourceNetworkConfig) SetSubnet(v GcpVpcSubnetConfig)`

SetSubnet sets Subnet field to given value.


### SetSubnetNil

`func (o *RecoverGcpVmNewSourceNetworkConfig) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *RecoverGcpVmNewSourceNetworkConfig) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


