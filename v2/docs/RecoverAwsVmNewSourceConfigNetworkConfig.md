# RecoverAwsVmNewSourceConfigNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroups** | [**[]RecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the network security groups within above VPC. | 
**Subnet** | [**NullableRecoverAwsAuroraNewSourceNetworkConfigSubnet**](RecoverAwsAuroraNewSourceNetworkConfigSubnet.md) |  | 
**Vpc** | [**NullableRecoverAwsAuroraNewSourceNetworkConfigVpc**](RecoverAwsAuroraNewSourceNetworkConfigVpc.md) |  | 

## Methods

### NewRecoverAwsVmNewSourceConfigNetworkConfig

`func NewRecoverAwsVmNewSourceConfigNetworkConfig(securityGroups []RecoveryObjectIdentifier, subnet NullableRecoverAwsAuroraNewSourceNetworkConfigSubnet, vpc NullableRecoverAwsAuroraNewSourceNetworkConfigVpc, ) *RecoverAwsVmNewSourceConfigNetworkConfig`

NewRecoverAwsVmNewSourceConfigNetworkConfig instantiates a new RecoverAwsVmNewSourceConfigNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAwsVmNewSourceConfigNetworkConfigWithDefaults

`func NewRecoverAwsVmNewSourceConfigNetworkConfigWithDefaults() *RecoverAwsVmNewSourceConfigNetworkConfig`

NewRecoverAwsVmNewSourceConfigNetworkConfigWithDefaults instantiates a new RecoverAwsVmNewSourceConfigNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroups

`func (o *RecoverAwsVmNewSourceConfigNetworkConfig) GetSecurityGroups() []RecoveryObjectIdentifier`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *RecoverAwsVmNewSourceConfigNetworkConfig) GetSecurityGroupsOk() (*[]RecoveryObjectIdentifier, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *RecoverAwsVmNewSourceConfigNetworkConfig) SetSecurityGroups(v []RecoveryObjectIdentifier)`

SetSecurityGroups sets SecurityGroups field to given value.


### SetSecurityGroupsNil

`func (o *RecoverAwsVmNewSourceConfigNetworkConfig) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *RecoverAwsVmNewSourceConfigNetworkConfig) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetSubnet

`func (o *RecoverAwsVmNewSourceConfigNetworkConfig) GetSubnet() RecoverAwsAuroraNewSourceNetworkConfigSubnet`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *RecoverAwsVmNewSourceConfigNetworkConfig) GetSubnetOk() (*RecoverAwsAuroraNewSourceNetworkConfigSubnet, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *RecoverAwsVmNewSourceConfigNetworkConfig) SetSubnet(v RecoverAwsAuroraNewSourceNetworkConfigSubnet)`

SetSubnet sets Subnet field to given value.


### SetSubnetNil

`func (o *RecoverAwsVmNewSourceConfigNetworkConfig) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *RecoverAwsVmNewSourceConfigNetworkConfig) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetVpc

`func (o *RecoverAwsVmNewSourceConfigNetworkConfig) GetVpc() RecoverAwsAuroraNewSourceNetworkConfigVpc`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *RecoverAwsVmNewSourceConfigNetworkConfig) GetVpcOk() (*RecoverAwsAuroraNewSourceNetworkConfigVpc, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *RecoverAwsVmNewSourceConfigNetworkConfig) SetVpc(v RecoverAwsAuroraNewSourceNetworkConfigVpc)`

SetVpc sets Vpc field to given value.


### SetVpcNil

`func (o *RecoverAwsVmNewSourceConfigNetworkConfig) SetVpcNil(b bool)`

 SetVpcNil sets the value for Vpc to be an explicit nil

### UnsetVpc
`func (o *RecoverAwsVmNewSourceConfigNetworkConfig) UnsetVpc()`

UnsetVpc ensures that no value is present for Vpc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


