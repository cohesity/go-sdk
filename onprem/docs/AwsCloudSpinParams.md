# AwsCloudSpinParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **NullableInt64** | Specifies id of the AWS region in which to deploy the VM. | 
**VpcId** | Pointer to **NullableInt64** | Specifies id of the Virtual Private Cloud to chose for the instance type. | [optional] 
**SubnetId** | Pointer to **NullableInt64** | Specifies id of the subnet within above VPC. | [optional] 

## Methods

### NewAwsCloudSpinParams

`func NewAwsCloudSpinParams(region NullableInt64, ) *AwsCloudSpinParams`

NewAwsCloudSpinParams instantiates a new AwsCloudSpinParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsCloudSpinParamsWithDefaults

`func NewAwsCloudSpinParamsWithDefaults() *AwsCloudSpinParams`

NewAwsCloudSpinParamsWithDefaults instantiates a new AwsCloudSpinParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *AwsCloudSpinParams) GetRegion() int64`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AwsCloudSpinParams) GetRegionOk() (*int64, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AwsCloudSpinParams) SetRegion(v int64)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *AwsCloudSpinParams) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *AwsCloudSpinParams) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetVpcId

`func (o *AwsCloudSpinParams) GetVpcId() int64`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AwsCloudSpinParams) GetVpcIdOk() (*int64, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AwsCloudSpinParams) SetVpcId(v int64)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *AwsCloudSpinParams) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *AwsCloudSpinParams) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *AwsCloudSpinParams) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetSubnetId

`func (o *AwsCloudSpinParams) GetSubnetId() int64`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *AwsCloudSpinParams) GetSubnetIdOk() (*int64, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *AwsCloudSpinParams) SetSubnetId(v int64)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *AwsCloudSpinParams) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *AwsCloudSpinParams) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *AwsCloudSpinParams) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


