# HeliosAwsCloudSpinParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **NullableInt64** | Specifies id of the AWS region in which to deploy the VM. | 
**SubnetId** | Pointer to **NullableInt64** | Specifies id of the subnet within above VPC. | [optional] 
**VpcId** | Pointer to **NullableInt64** | Specifies id of the Virtual Private Cloud to chose for the instance type. | [optional] 

## Methods

### NewHeliosAwsCloudSpinParams

`func NewHeliosAwsCloudSpinParams(region NullableInt64, ) *HeliosAwsCloudSpinParams`

NewHeliosAwsCloudSpinParams instantiates a new HeliosAwsCloudSpinParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosAwsCloudSpinParamsWithDefaults

`func NewHeliosAwsCloudSpinParamsWithDefaults() *HeliosAwsCloudSpinParams`

NewHeliosAwsCloudSpinParamsWithDefaults instantiates a new HeliosAwsCloudSpinParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *HeliosAwsCloudSpinParams) GetRegion() int64`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HeliosAwsCloudSpinParams) GetRegionOk() (*int64, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HeliosAwsCloudSpinParams) SetRegion(v int64)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *HeliosAwsCloudSpinParams) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *HeliosAwsCloudSpinParams) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetSubnetId

`func (o *HeliosAwsCloudSpinParams) GetSubnetId() int64`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *HeliosAwsCloudSpinParams) GetSubnetIdOk() (*int64, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *HeliosAwsCloudSpinParams) SetSubnetId(v int64)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *HeliosAwsCloudSpinParams) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *HeliosAwsCloudSpinParams) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *HeliosAwsCloudSpinParams) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetVpcId

`func (o *HeliosAwsCloudSpinParams) GetVpcId() int64`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *HeliosAwsCloudSpinParams) GetVpcIdOk() (*int64, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *HeliosAwsCloudSpinParams) SetVpcId(v int64)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *HeliosAwsCloudSpinParams) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *HeliosAwsCloudSpinParams) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *HeliosAwsCloudSpinParams) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


