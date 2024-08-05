# RecoverGcpVmNewSourceNetworkConfigSubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubnetId** | **NullableInt64** | Specifies the id of the subnet. | 
**SubnetName** | Pointer to **NullableString** | Specifies the name of the subnet. | [optional] [readonly] 
**VpcName** | Pointer to **NullableString** | Specifies the name of the vpc network. | [optional] [readonly] 

## Methods

### NewRecoverGcpVmNewSourceNetworkConfigSubnet

`func NewRecoverGcpVmNewSourceNetworkConfigSubnet(subnetId NullableInt64, ) *RecoverGcpVmNewSourceNetworkConfigSubnet`

NewRecoverGcpVmNewSourceNetworkConfigSubnet instantiates a new RecoverGcpVmNewSourceNetworkConfigSubnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverGcpVmNewSourceNetworkConfigSubnetWithDefaults

`func NewRecoverGcpVmNewSourceNetworkConfigSubnetWithDefaults() *RecoverGcpVmNewSourceNetworkConfigSubnet`

NewRecoverGcpVmNewSourceNetworkConfigSubnetWithDefaults instantiates a new RecoverGcpVmNewSourceNetworkConfigSubnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnetId

`func (o *RecoverGcpVmNewSourceNetworkConfigSubnet) GetSubnetId() int64`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *RecoverGcpVmNewSourceNetworkConfigSubnet) GetSubnetIdOk() (*int64, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *RecoverGcpVmNewSourceNetworkConfigSubnet) SetSubnetId(v int64)`

SetSubnetId sets SubnetId field to given value.


### SetSubnetIdNil

`func (o *RecoverGcpVmNewSourceNetworkConfigSubnet) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *RecoverGcpVmNewSourceNetworkConfigSubnet) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetSubnetName

`func (o *RecoverGcpVmNewSourceNetworkConfigSubnet) GetSubnetName() string`

GetSubnetName returns the SubnetName field if non-nil, zero value otherwise.

### GetSubnetNameOk

`func (o *RecoverGcpVmNewSourceNetworkConfigSubnet) GetSubnetNameOk() (*string, bool)`

GetSubnetNameOk returns a tuple with the SubnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetName

`func (o *RecoverGcpVmNewSourceNetworkConfigSubnet) SetSubnetName(v string)`

SetSubnetName sets SubnetName field to given value.

### HasSubnetName

`func (o *RecoverGcpVmNewSourceNetworkConfigSubnet) HasSubnetName() bool`

HasSubnetName returns a boolean if a field has been set.

### SetSubnetNameNil

`func (o *RecoverGcpVmNewSourceNetworkConfigSubnet) SetSubnetNameNil(b bool)`

 SetSubnetNameNil sets the value for SubnetName to be an explicit nil

### UnsetSubnetName
`func (o *RecoverGcpVmNewSourceNetworkConfigSubnet) UnsetSubnetName()`

UnsetSubnetName ensures that no value is present for SubnetName, not even an explicit nil
### GetVpcName

`func (o *RecoverGcpVmNewSourceNetworkConfigSubnet) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *RecoverGcpVmNewSourceNetworkConfigSubnet) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *RecoverGcpVmNewSourceNetworkConfigSubnet) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.

### HasVpcName

`func (o *RecoverGcpVmNewSourceNetworkConfigSubnet) HasVpcName() bool`

HasVpcName returns a boolean if a field has been set.

### SetVpcNameNil

`func (o *RecoverGcpVmNewSourceNetworkConfigSubnet) SetVpcNameNil(b bool)`

 SetVpcNameNil sets the value for VpcName to be an explicit nil

### UnsetVpcName
`func (o *RecoverGcpVmNewSourceNetworkConfigSubnet) UnsetVpcName()`

UnsetVpcName ensures that no value is present for VpcName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


