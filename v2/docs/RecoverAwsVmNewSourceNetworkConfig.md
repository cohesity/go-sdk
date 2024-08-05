# RecoverAwsVmNewSourceNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroups** | [**[]RecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the network security groups within above VPC. | 
**Subnet** | [**NullableRecoverAwsAuroraNewSourceNetworkConfigSubnet**](RecoverAwsAuroraNewSourceNetworkConfigSubnet.md) |  | 
**Vpc** | [**NullableRecoverAwsAuroraNewSourceNetworkConfigVpc**](RecoverAwsAuroraNewSourceNetworkConfigVpc.md) |  | 

## Methods

### NewRecoverAwsVmNewSourceNetworkConfig

`func NewRecoverAwsVmNewSourceNetworkConfig(securityGroups []RecoveryObjectIdentifier, subnet NullableRecoverAwsAuroraNewSourceNetworkConfigSubnet, vpc NullableRecoverAwsAuroraNewSourceNetworkConfigVpc, ) *RecoverAwsVmNewSourceNetworkConfig`

NewRecoverAwsVmNewSourceNetworkConfig instantiates a new RecoverAwsVmNewSourceNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAwsVmNewSourceNetworkConfigWithDefaults

`func NewRecoverAwsVmNewSourceNetworkConfigWithDefaults() *RecoverAwsVmNewSourceNetworkConfig`

NewRecoverAwsVmNewSourceNetworkConfigWithDefaults instantiates a new RecoverAwsVmNewSourceNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroups

`func (o *RecoverAwsVmNewSourceNetworkConfig) GetSecurityGroups() []RecoveryObjectIdentifier`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *RecoverAwsVmNewSourceNetworkConfig) GetSecurityGroupsOk() (*[]RecoveryObjectIdentifier, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *RecoverAwsVmNewSourceNetworkConfig) SetSecurityGroups(v []RecoveryObjectIdentifier)`

SetSecurityGroups sets SecurityGroups field to given value.


### SetSecurityGroupsNil

`func (o *RecoverAwsVmNewSourceNetworkConfig) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *RecoverAwsVmNewSourceNetworkConfig) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetSubnet

`func (o *RecoverAwsVmNewSourceNetworkConfig) GetSubnet() RecoverAwsAuroraNewSourceNetworkConfigSubnet`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *RecoverAwsVmNewSourceNetworkConfig) GetSubnetOk() (*RecoverAwsAuroraNewSourceNetworkConfigSubnet, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *RecoverAwsVmNewSourceNetworkConfig) SetSubnet(v RecoverAwsAuroraNewSourceNetworkConfigSubnet)`

SetSubnet sets Subnet field to given value.


### SetSubnetNil

`func (o *RecoverAwsVmNewSourceNetworkConfig) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *RecoverAwsVmNewSourceNetworkConfig) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetVpc

`func (o *RecoverAwsVmNewSourceNetworkConfig) GetVpc() RecoverAwsAuroraNewSourceNetworkConfigVpc`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *RecoverAwsVmNewSourceNetworkConfig) GetVpcOk() (*RecoverAwsAuroraNewSourceNetworkConfigVpc, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *RecoverAwsVmNewSourceNetworkConfig) SetVpc(v RecoverAwsAuroraNewSourceNetworkConfigVpc)`

SetVpc sets Vpc field to given value.


### SetVpcNil

`func (o *RecoverAwsVmNewSourceNetworkConfig) SetVpcNil(b bool)`

 SetVpcNil sets the value for Vpc to be an explicit nil

### UnsetVpc
`func (o *RecoverAwsVmNewSourceNetworkConfig) UnsetVpc()`

UnsetVpc ensures that no value is present for Vpc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


