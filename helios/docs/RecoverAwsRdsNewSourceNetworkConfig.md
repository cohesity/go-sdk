# RecoverAwsRdsNewSourceNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vpc** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the Virtual Private Cloud to choose for the instance type. | 
**Subnet** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the subnet within above VPC. | 
**AvailabilityZone** | Pointer to [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the entity representing the availability zone to use while restoring the DB. | [optional] 
**SecurityGroups** | Pointer to [**[]RecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the network security groups within above VPC. | [optional] 

## Methods

### NewRecoverAwsRdsNewSourceNetworkConfig

`func NewRecoverAwsRdsNewSourceNetworkConfig(vpc NullableRecoveryObjectIdentifier, subnet NullableRecoveryObjectIdentifier, ) *RecoverAwsRdsNewSourceNetworkConfig`

NewRecoverAwsRdsNewSourceNetworkConfig instantiates a new RecoverAwsRdsNewSourceNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAwsRdsNewSourceNetworkConfigWithDefaults

`func NewRecoverAwsRdsNewSourceNetworkConfigWithDefaults() *RecoverAwsRdsNewSourceNetworkConfig`

NewRecoverAwsRdsNewSourceNetworkConfigWithDefaults instantiates a new RecoverAwsRdsNewSourceNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpc

`func (o *RecoverAwsRdsNewSourceNetworkConfig) GetVpc() RecoveryObjectIdentifier`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *RecoverAwsRdsNewSourceNetworkConfig) GetVpcOk() (*RecoveryObjectIdentifier, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *RecoverAwsRdsNewSourceNetworkConfig) SetVpc(v RecoveryObjectIdentifier)`

SetVpc sets Vpc field to given value.


### SetVpcNil

`func (o *RecoverAwsRdsNewSourceNetworkConfig) SetVpcNil(b bool)`

 SetVpcNil sets the value for Vpc to be an explicit nil

### UnsetVpc
`func (o *RecoverAwsRdsNewSourceNetworkConfig) UnsetVpc()`

UnsetVpc ensures that no value is present for Vpc, not even an explicit nil
### GetSubnet

`func (o *RecoverAwsRdsNewSourceNetworkConfig) GetSubnet() RecoveryObjectIdentifier`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *RecoverAwsRdsNewSourceNetworkConfig) GetSubnetOk() (*RecoveryObjectIdentifier, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *RecoverAwsRdsNewSourceNetworkConfig) SetSubnet(v RecoveryObjectIdentifier)`

SetSubnet sets Subnet field to given value.


### SetSubnetNil

`func (o *RecoverAwsRdsNewSourceNetworkConfig) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *RecoverAwsRdsNewSourceNetworkConfig) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetAvailabilityZone

`func (o *RecoverAwsRdsNewSourceNetworkConfig) GetAvailabilityZone() RecoveryObjectIdentifier`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *RecoverAwsRdsNewSourceNetworkConfig) GetAvailabilityZoneOk() (*RecoveryObjectIdentifier, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *RecoverAwsRdsNewSourceNetworkConfig) SetAvailabilityZone(v RecoveryObjectIdentifier)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *RecoverAwsRdsNewSourceNetworkConfig) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *RecoverAwsRdsNewSourceNetworkConfig) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *RecoverAwsRdsNewSourceNetworkConfig) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetSecurityGroups

`func (o *RecoverAwsRdsNewSourceNetworkConfig) GetSecurityGroups() []RecoveryObjectIdentifier`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *RecoverAwsRdsNewSourceNetworkConfig) GetSecurityGroupsOk() (*[]RecoveryObjectIdentifier, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *RecoverAwsRdsNewSourceNetworkConfig) SetSecurityGroups(v []RecoveryObjectIdentifier)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *RecoverAwsRdsNewSourceNetworkConfig) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *RecoverAwsRdsNewSourceNetworkConfig) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *RecoverAwsRdsNewSourceNetworkConfig) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


