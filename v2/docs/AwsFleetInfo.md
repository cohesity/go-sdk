# AwsFleetInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IamRole** | **NullableString** | Specifies the IAM role used to create instances. | 
**Region** | **NullableString** | Specifies the Region of the CE cluster. | 
**SecurityGroupId** | **NullableString** | Specifies the security group Id. | 
**SubnetId** | **NullableString** | Specifies the Subnet of the CE cluster. | 
**VpcId** | **NullableString** | Specifies the VPC of the CE cluster. | 

## Methods

### NewAwsFleetInfo

`func NewAwsFleetInfo(iamRole NullableString, region NullableString, securityGroupId NullableString, subnetId NullableString, vpcId NullableString, ) *AwsFleetInfo`

NewAwsFleetInfo instantiates a new AwsFleetInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsFleetInfoWithDefaults

`func NewAwsFleetInfoWithDefaults() *AwsFleetInfo`

NewAwsFleetInfoWithDefaults instantiates a new AwsFleetInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIamRole

`func (o *AwsFleetInfo) GetIamRole() string`

GetIamRole returns the IamRole field if non-nil, zero value otherwise.

### GetIamRoleOk

`func (o *AwsFleetInfo) GetIamRoleOk() (*string, bool)`

GetIamRoleOk returns a tuple with the IamRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRole

`func (o *AwsFleetInfo) SetIamRole(v string)`

SetIamRole sets IamRole field to given value.


### SetIamRoleNil

`func (o *AwsFleetInfo) SetIamRoleNil(b bool)`

 SetIamRoleNil sets the value for IamRole to be an explicit nil

### UnsetIamRole
`func (o *AwsFleetInfo) UnsetIamRole()`

UnsetIamRole ensures that no value is present for IamRole, not even an explicit nil
### GetRegion

`func (o *AwsFleetInfo) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AwsFleetInfo) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AwsFleetInfo) SetRegion(v string)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *AwsFleetInfo) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *AwsFleetInfo) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetSecurityGroupId

`func (o *AwsFleetInfo) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *AwsFleetInfo) GetSecurityGroupIdOk() (*string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *AwsFleetInfo) SetSecurityGroupId(v string)`

SetSecurityGroupId sets SecurityGroupId field to given value.


### SetSecurityGroupIdNil

`func (o *AwsFleetInfo) SetSecurityGroupIdNil(b bool)`

 SetSecurityGroupIdNil sets the value for SecurityGroupId to be an explicit nil

### UnsetSecurityGroupId
`func (o *AwsFleetInfo) UnsetSecurityGroupId()`

UnsetSecurityGroupId ensures that no value is present for SecurityGroupId, not even an explicit nil
### GetSubnetId

`func (o *AwsFleetInfo) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *AwsFleetInfo) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *AwsFleetInfo) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### SetSubnetIdNil

`func (o *AwsFleetInfo) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *AwsFleetInfo) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetVpcId

`func (o *AwsFleetInfo) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AwsFleetInfo) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AwsFleetInfo) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### SetVpcIdNil

`func (o *AwsFleetInfo) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *AwsFleetInfo) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


