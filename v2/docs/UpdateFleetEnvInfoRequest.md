# UpdateFleetEnvInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsFleetInfo** | Pointer to [**AwsFleetInfo**](AwsFleetInfo.md) |  | [optional] 
**AzureFleetInfo** | Pointer to [**AzureFleetInfo**](AzureFleetInfo.md) |  | [optional] 
**IamRole** | Pointer to **NullableString** | Specifies the IAM role used to create instances.This field is now deprecated, use awsFleetInfo instead. | [optional] 
**Region** | Pointer to **NullableString** | Specifies the Region of the CE cluster. This field is now deprecated, use awsFleetInfo instead. | [optional] 
**SecurityGroupId** | Pointer to **NullableString** | Specifies the security group Id.This field is now deprecated, use awsFleetInfo instead. | [optional] 
**SubnetId** | Pointer to **NullableString** | Specifies the Subnet of the CE cluster.This field is now deprecated, use awsFleetInfo instead. | [optional] 
**VpcId** | Pointer to **NullableString** | Specifies the VPC of the CE cluster.This field is now deprecated, use awsFleetInfo instead. | [optional] 

## Methods

### NewUpdateFleetEnvInfoRequest

`func NewUpdateFleetEnvInfoRequest() *UpdateFleetEnvInfoRequest`

NewUpdateFleetEnvInfoRequest instantiates a new UpdateFleetEnvInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFleetEnvInfoRequestWithDefaults

`func NewUpdateFleetEnvInfoRequestWithDefaults() *UpdateFleetEnvInfoRequest`

NewUpdateFleetEnvInfoRequestWithDefaults instantiates a new UpdateFleetEnvInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsFleetInfo

`func (o *UpdateFleetEnvInfoRequest) GetAwsFleetInfo() AwsFleetInfo`

GetAwsFleetInfo returns the AwsFleetInfo field if non-nil, zero value otherwise.

### GetAwsFleetInfoOk

`func (o *UpdateFleetEnvInfoRequest) GetAwsFleetInfoOk() (*AwsFleetInfo, bool)`

GetAwsFleetInfoOk returns a tuple with the AwsFleetInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsFleetInfo

`func (o *UpdateFleetEnvInfoRequest) SetAwsFleetInfo(v AwsFleetInfo)`

SetAwsFleetInfo sets AwsFleetInfo field to given value.

### HasAwsFleetInfo

`func (o *UpdateFleetEnvInfoRequest) HasAwsFleetInfo() bool`

HasAwsFleetInfo returns a boolean if a field has been set.

### GetAzureFleetInfo

`func (o *UpdateFleetEnvInfoRequest) GetAzureFleetInfo() AzureFleetInfo`

GetAzureFleetInfo returns the AzureFleetInfo field if non-nil, zero value otherwise.

### GetAzureFleetInfoOk

`func (o *UpdateFleetEnvInfoRequest) GetAzureFleetInfoOk() (*AzureFleetInfo, bool)`

GetAzureFleetInfoOk returns a tuple with the AzureFleetInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureFleetInfo

`func (o *UpdateFleetEnvInfoRequest) SetAzureFleetInfo(v AzureFleetInfo)`

SetAzureFleetInfo sets AzureFleetInfo field to given value.

### HasAzureFleetInfo

`func (o *UpdateFleetEnvInfoRequest) HasAzureFleetInfo() bool`

HasAzureFleetInfo returns a boolean if a field has been set.

### GetIamRole

`func (o *UpdateFleetEnvInfoRequest) GetIamRole() string`

GetIamRole returns the IamRole field if non-nil, zero value otherwise.

### GetIamRoleOk

`func (o *UpdateFleetEnvInfoRequest) GetIamRoleOk() (*string, bool)`

GetIamRoleOk returns a tuple with the IamRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRole

`func (o *UpdateFleetEnvInfoRequest) SetIamRole(v string)`

SetIamRole sets IamRole field to given value.

### HasIamRole

`func (o *UpdateFleetEnvInfoRequest) HasIamRole() bool`

HasIamRole returns a boolean if a field has been set.

### SetIamRoleNil

`func (o *UpdateFleetEnvInfoRequest) SetIamRoleNil(b bool)`

 SetIamRoleNil sets the value for IamRole to be an explicit nil

### UnsetIamRole
`func (o *UpdateFleetEnvInfoRequest) UnsetIamRole()`

UnsetIamRole ensures that no value is present for IamRole, not even an explicit nil
### GetRegion

`func (o *UpdateFleetEnvInfoRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *UpdateFleetEnvInfoRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *UpdateFleetEnvInfoRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *UpdateFleetEnvInfoRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *UpdateFleetEnvInfoRequest) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *UpdateFleetEnvInfoRequest) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetSecurityGroupId

`func (o *UpdateFleetEnvInfoRequest) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *UpdateFleetEnvInfoRequest) GetSecurityGroupIdOk() (*string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *UpdateFleetEnvInfoRequest) SetSecurityGroupId(v string)`

SetSecurityGroupId sets SecurityGroupId field to given value.

### HasSecurityGroupId

`func (o *UpdateFleetEnvInfoRequest) HasSecurityGroupId() bool`

HasSecurityGroupId returns a boolean if a field has been set.

### SetSecurityGroupIdNil

`func (o *UpdateFleetEnvInfoRequest) SetSecurityGroupIdNil(b bool)`

 SetSecurityGroupIdNil sets the value for SecurityGroupId to be an explicit nil

### UnsetSecurityGroupId
`func (o *UpdateFleetEnvInfoRequest) UnsetSecurityGroupId()`

UnsetSecurityGroupId ensures that no value is present for SecurityGroupId, not even an explicit nil
### GetSubnetId

`func (o *UpdateFleetEnvInfoRequest) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *UpdateFleetEnvInfoRequest) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *UpdateFleetEnvInfoRequest) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *UpdateFleetEnvInfoRequest) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *UpdateFleetEnvInfoRequest) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *UpdateFleetEnvInfoRequest) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetVpcId

`func (o *UpdateFleetEnvInfoRequest) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *UpdateFleetEnvInfoRequest) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *UpdateFleetEnvInfoRequest) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *UpdateFleetEnvInfoRequest) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *UpdateFleetEnvInfoRequest) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *UpdateFleetEnvInfoRequest) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


