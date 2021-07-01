# AwsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | Pointer to **NullableInt64** | Specifies id of the AWS instance type in which to deploy the VM. | [optional] 
**NetworkSecurityGroupIds** | Pointer to **[]int64** | Specifies ids of the network security groups within above VPC. | [optional] 
**RdsParams** | Pointer to [**RdsParams**](RdsParams.md) |  | [optional] 
**Region** | Pointer to **NullableInt64** | Specifies id of the AWS region in which to deploy the VM. | [optional] 
**SubnetId** | Pointer to **NullableInt64** | Specifies id of the subnet within above VPC. | [optional] 
**VirtualPrivateCloudId** | Pointer to **NullableInt64** | Specifies id of the Virtual Private Cloud to chose for the instance type. | [optional] 

## Methods

### NewAwsParams

`func NewAwsParams() *AwsParams`

NewAwsParams instantiates a new AwsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsParamsWithDefaults

`func NewAwsParamsWithDefaults() *AwsParams`

NewAwsParamsWithDefaults instantiates a new AwsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *AwsParams) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *AwsParams) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *AwsParams) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *AwsParams) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### SetInstanceIdNil

`func (o *AwsParams) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *AwsParams) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetNetworkSecurityGroupIds

`func (o *AwsParams) GetNetworkSecurityGroupIds() []int64`

GetNetworkSecurityGroupIds returns the NetworkSecurityGroupIds field if non-nil, zero value otherwise.

### GetNetworkSecurityGroupIdsOk

`func (o *AwsParams) GetNetworkSecurityGroupIdsOk() (*[]int64, bool)`

GetNetworkSecurityGroupIdsOk returns a tuple with the NetworkSecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSecurityGroupIds

`func (o *AwsParams) SetNetworkSecurityGroupIds(v []int64)`

SetNetworkSecurityGroupIds sets NetworkSecurityGroupIds field to given value.

### HasNetworkSecurityGroupIds

`func (o *AwsParams) HasNetworkSecurityGroupIds() bool`

HasNetworkSecurityGroupIds returns a boolean if a field has been set.

### SetNetworkSecurityGroupIdsNil

`func (o *AwsParams) SetNetworkSecurityGroupIdsNil(b bool)`

 SetNetworkSecurityGroupIdsNil sets the value for NetworkSecurityGroupIds to be an explicit nil

### UnsetNetworkSecurityGroupIds
`func (o *AwsParams) UnsetNetworkSecurityGroupIds()`

UnsetNetworkSecurityGroupIds ensures that no value is present for NetworkSecurityGroupIds, not even an explicit nil
### GetRdsParams

`func (o *AwsParams) GetRdsParams() RdsParams`

GetRdsParams returns the RdsParams field if non-nil, zero value otherwise.

### GetRdsParamsOk

`func (o *AwsParams) GetRdsParamsOk() (*RdsParams, bool)`

GetRdsParamsOk returns a tuple with the RdsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsParams

`func (o *AwsParams) SetRdsParams(v RdsParams)`

SetRdsParams sets RdsParams field to given value.

### HasRdsParams

`func (o *AwsParams) HasRdsParams() bool`

HasRdsParams returns a boolean if a field has been set.

### GetRegion

`func (o *AwsParams) GetRegion() int64`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AwsParams) GetRegionOk() (*int64, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AwsParams) SetRegion(v int64)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AwsParams) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *AwsParams) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *AwsParams) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetSubnetId

`func (o *AwsParams) GetSubnetId() int64`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *AwsParams) GetSubnetIdOk() (*int64, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *AwsParams) SetSubnetId(v int64)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *AwsParams) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *AwsParams) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *AwsParams) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetVirtualPrivateCloudId

`func (o *AwsParams) GetVirtualPrivateCloudId() int64`

GetVirtualPrivateCloudId returns the VirtualPrivateCloudId field if non-nil, zero value otherwise.

### GetVirtualPrivateCloudIdOk

`func (o *AwsParams) GetVirtualPrivateCloudIdOk() (*int64, bool)`

GetVirtualPrivateCloudIdOk returns a tuple with the VirtualPrivateCloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPrivateCloudId

`func (o *AwsParams) SetVirtualPrivateCloudId(v int64)`

SetVirtualPrivateCloudId sets VirtualPrivateCloudId field to given value.

### HasVirtualPrivateCloudId

`func (o *AwsParams) HasVirtualPrivateCloudId() bool`

HasVirtualPrivateCloudId returns a boolean if a field has been set.

### SetVirtualPrivateCloudIdNil

`func (o *AwsParams) SetVirtualPrivateCloudIdNil(b bool)`

 SetVirtualPrivateCloudIdNil sets the value for VirtualPrivateCloudId to be an explicit nil

### UnsetVirtualPrivateCloudId
`func (o *AwsParams) UnsetVirtualPrivateCloudId()`

UnsetVirtualPrivateCloudId ensures that no value is present for VirtualPrivateCloudId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


