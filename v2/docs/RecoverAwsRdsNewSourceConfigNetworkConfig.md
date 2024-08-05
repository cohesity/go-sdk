# RecoverAwsRdsNewSourceConfigNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZone** | Pointer to [**NullableRecoverAwsAuroraNewSourceNetworkConfigAvailabilityZone**](RecoverAwsAuroraNewSourceNetworkConfigAvailabilityZone.md) |  | [optional] 
**SecurityGroups** | Pointer to [**[]RecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the network security groups within above VPC. | [optional] 
**Subnet** | [**NullableRecoverAwsAuroraNewSourceNetworkConfigSubnet**](RecoverAwsAuroraNewSourceNetworkConfigSubnet.md) |  | 
**Vpc** | [**NullableRecoverAwsAuroraNewSourceNetworkConfigVpc**](RecoverAwsAuroraNewSourceNetworkConfigVpc.md) |  | 

## Methods

### NewRecoverAwsRdsNewSourceConfigNetworkConfig

`func NewRecoverAwsRdsNewSourceConfigNetworkConfig(subnet NullableRecoverAwsAuroraNewSourceNetworkConfigSubnet, vpc NullableRecoverAwsAuroraNewSourceNetworkConfigVpc, ) *RecoverAwsRdsNewSourceConfigNetworkConfig`

NewRecoverAwsRdsNewSourceConfigNetworkConfig instantiates a new RecoverAwsRdsNewSourceConfigNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAwsRdsNewSourceConfigNetworkConfigWithDefaults

`func NewRecoverAwsRdsNewSourceConfigNetworkConfigWithDefaults() *RecoverAwsRdsNewSourceConfigNetworkConfig`

NewRecoverAwsRdsNewSourceConfigNetworkConfigWithDefaults instantiates a new RecoverAwsRdsNewSourceConfigNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZone

`func (o *RecoverAwsRdsNewSourceConfigNetworkConfig) GetAvailabilityZone() RecoverAwsAuroraNewSourceNetworkConfigAvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *RecoverAwsRdsNewSourceConfigNetworkConfig) GetAvailabilityZoneOk() (*RecoverAwsAuroraNewSourceNetworkConfigAvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *RecoverAwsRdsNewSourceConfigNetworkConfig) SetAvailabilityZone(v RecoverAwsAuroraNewSourceNetworkConfigAvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *RecoverAwsRdsNewSourceConfigNetworkConfig) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *RecoverAwsRdsNewSourceConfigNetworkConfig) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *RecoverAwsRdsNewSourceConfigNetworkConfig) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetSecurityGroups

`func (o *RecoverAwsRdsNewSourceConfigNetworkConfig) GetSecurityGroups() []RecoveryObjectIdentifier`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *RecoverAwsRdsNewSourceConfigNetworkConfig) GetSecurityGroupsOk() (*[]RecoveryObjectIdentifier, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *RecoverAwsRdsNewSourceConfigNetworkConfig) SetSecurityGroups(v []RecoveryObjectIdentifier)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *RecoverAwsRdsNewSourceConfigNetworkConfig) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *RecoverAwsRdsNewSourceConfigNetworkConfig) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *RecoverAwsRdsNewSourceConfigNetworkConfig) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetSubnet

`func (o *RecoverAwsRdsNewSourceConfigNetworkConfig) GetSubnet() RecoverAwsAuroraNewSourceNetworkConfigSubnet`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *RecoverAwsRdsNewSourceConfigNetworkConfig) GetSubnetOk() (*RecoverAwsAuroraNewSourceNetworkConfigSubnet, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *RecoverAwsRdsNewSourceConfigNetworkConfig) SetSubnet(v RecoverAwsAuroraNewSourceNetworkConfigSubnet)`

SetSubnet sets Subnet field to given value.


### SetSubnetNil

`func (o *RecoverAwsRdsNewSourceConfigNetworkConfig) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *RecoverAwsRdsNewSourceConfigNetworkConfig) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetVpc

`func (o *RecoverAwsRdsNewSourceConfigNetworkConfig) GetVpc() RecoverAwsAuroraNewSourceNetworkConfigVpc`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *RecoverAwsRdsNewSourceConfigNetworkConfig) GetVpcOk() (*RecoverAwsAuroraNewSourceNetworkConfigVpc, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *RecoverAwsRdsNewSourceConfigNetworkConfig) SetVpc(v RecoverAwsAuroraNewSourceNetworkConfigVpc)`

SetVpc sets Vpc field to given value.


### SetVpcNil

`func (o *RecoverAwsRdsNewSourceConfigNetworkConfig) SetVpcNil(b bool)`

 SetVpcNil sets the value for Vpc to be an explicit nil

### UnsetVpc
`func (o *RecoverAwsRdsNewSourceConfigNetworkConfig) UnsetVpc()`

UnsetVpc ensures that no value is present for Vpc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


