# IamRoleAwsCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpIamRoleArn** | Pointer to **NullableString** | This is only applicable in case of DMaaS. Control plane IAM role ARN, this is first assumed by the dataplane(cluster). Then we assume the iam_role_arn which is tenant&#39;s IAM role with all required permissions. | [optional] 
**IamRoleArn** | **NullableString** | Specifies the IAM role which will be used to access the security credentials required for API calls. This should have all the permissions required for the tenant&#39;s use case. In case of DMaaS this will be the Tenant&#39;s IAM role ARN. This is assumed only after the cp_iam_role_arn(control plane role) is assumed | 

## Methods

### NewIamRoleAwsCredentials

`func NewIamRoleAwsCredentials(iamRoleArn NullableString, ) *IamRoleAwsCredentials`

NewIamRoleAwsCredentials instantiates a new IamRoleAwsCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamRoleAwsCredentialsWithDefaults

`func NewIamRoleAwsCredentialsWithDefaults() *IamRoleAwsCredentials`

NewIamRoleAwsCredentialsWithDefaults instantiates a new IamRoleAwsCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpIamRoleArn

`func (o *IamRoleAwsCredentials) GetCpIamRoleArn() string`

GetCpIamRoleArn returns the CpIamRoleArn field if non-nil, zero value otherwise.

### GetCpIamRoleArnOk

`func (o *IamRoleAwsCredentials) GetCpIamRoleArnOk() (*string, bool)`

GetCpIamRoleArnOk returns a tuple with the CpIamRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpIamRoleArn

`func (o *IamRoleAwsCredentials) SetCpIamRoleArn(v string)`

SetCpIamRoleArn sets CpIamRoleArn field to given value.

### HasCpIamRoleArn

`func (o *IamRoleAwsCredentials) HasCpIamRoleArn() bool`

HasCpIamRoleArn returns a boolean if a field has been set.

### SetCpIamRoleArnNil

`func (o *IamRoleAwsCredentials) SetCpIamRoleArnNil(b bool)`

 SetCpIamRoleArnNil sets the value for CpIamRoleArn to be an explicit nil

### UnsetCpIamRoleArn
`func (o *IamRoleAwsCredentials) UnsetCpIamRoleArn()`

UnsetCpIamRoleArn ensures that no value is present for CpIamRoleArn, not even an explicit nil
### GetIamRoleArn

`func (o *IamRoleAwsCredentials) GetIamRoleArn() string`

GetIamRoleArn returns the IamRoleArn field if non-nil, zero value otherwise.

### GetIamRoleArnOk

`func (o *IamRoleAwsCredentials) GetIamRoleArnOk() (*string, bool)`

GetIamRoleArnOk returns a tuple with the IamRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRoleArn

`func (o *IamRoleAwsCredentials) SetIamRoleArn(v string)`

SetIamRoleArn sets IamRoleArn field to given value.


### SetIamRoleArnNil

`func (o *IamRoleAwsCredentials) SetIamRoleArnNil(b bool)`

 SetIamRoleArnNil sets the value for IamRoleArn to be an explicit nil

### UnsetIamRoleArn
`func (o *IamRoleAwsCredentials) UnsetIamRoleArn()`

UnsetIamRoleArn ensures that no value is present for IamRoleArn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


