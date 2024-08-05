# AwsKmsConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | Pointer to **NullableString** | AWS account access key id. Required when &#39;iamRoleArn&#39; is not given. | [optional] 
**CmkAlias** | Pointer to **NullableString** | AWS CMK alias. Only need one of cmkAlias, cmkArn, cmkKeyId to connect to AWS KMS. | [optional] 
**CmkArn** | Pointer to **NullableString** | AWS CMK Amazon resource number. Only need one of cmkAlias, cmkArn, cmkKeyId to connect to AWS KMS. | [optional] 
**CmkKeyId** | Pointer to **NullableString** | AWS CMK key id. Only need one of cmkAlias, cmkArn, cmkKeyId to connect to AWS KMS. | [optional] 
**IamRoleArn** | Pointer to **NullableString** | The IAM role which will be used to authenticate with AWS KMS. Required when &#39;accessKeyId&#39; and &#39;secretAccessKey&#39; fields are not provided. | [optional] 
**Region** | Pointer to **NullableString** | AWS region, e.g. us-east-1, us-west-2, for the AWS Glacier service to be used to authenticate resources within this region by the configured AWS account. | [optional] 
**VerifySSL** | Pointer to **NullableBool** | Enable SSL verification or not. | [optional] [default to true]

## Methods

### NewAwsKmsConfigurationResponse

`func NewAwsKmsConfigurationResponse() *AwsKmsConfigurationResponse`

NewAwsKmsConfigurationResponse instantiates a new AwsKmsConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsKmsConfigurationResponseWithDefaults

`func NewAwsKmsConfigurationResponseWithDefaults() *AwsKmsConfigurationResponse`

NewAwsKmsConfigurationResponseWithDefaults instantiates a new AwsKmsConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyId

`func (o *AwsKmsConfigurationResponse) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *AwsKmsConfigurationResponse) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *AwsKmsConfigurationResponse) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *AwsKmsConfigurationResponse) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### SetAccessKeyIdNil

`func (o *AwsKmsConfigurationResponse) SetAccessKeyIdNil(b bool)`

 SetAccessKeyIdNil sets the value for AccessKeyId to be an explicit nil

### UnsetAccessKeyId
`func (o *AwsKmsConfigurationResponse) UnsetAccessKeyId()`

UnsetAccessKeyId ensures that no value is present for AccessKeyId, not even an explicit nil
### GetCmkAlias

`func (o *AwsKmsConfigurationResponse) GetCmkAlias() string`

GetCmkAlias returns the CmkAlias field if non-nil, zero value otherwise.

### GetCmkAliasOk

`func (o *AwsKmsConfigurationResponse) GetCmkAliasOk() (*string, bool)`

GetCmkAliasOk returns a tuple with the CmkAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmkAlias

`func (o *AwsKmsConfigurationResponse) SetCmkAlias(v string)`

SetCmkAlias sets CmkAlias field to given value.

### HasCmkAlias

`func (o *AwsKmsConfigurationResponse) HasCmkAlias() bool`

HasCmkAlias returns a boolean if a field has been set.

### SetCmkAliasNil

`func (o *AwsKmsConfigurationResponse) SetCmkAliasNil(b bool)`

 SetCmkAliasNil sets the value for CmkAlias to be an explicit nil

### UnsetCmkAlias
`func (o *AwsKmsConfigurationResponse) UnsetCmkAlias()`

UnsetCmkAlias ensures that no value is present for CmkAlias, not even an explicit nil
### GetCmkArn

`func (o *AwsKmsConfigurationResponse) GetCmkArn() string`

GetCmkArn returns the CmkArn field if non-nil, zero value otherwise.

### GetCmkArnOk

`func (o *AwsKmsConfigurationResponse) GetCmkArnOk() (*string, bool)`

GetCmkArnOk returns a tuple with the CmkArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmkArn

`func (o *AwsKmsConfigurationResponse) SetCmkArn(v string)`

SetCmkArn sets CmkArn field to given value.

### HasCmkArn

`func (o *AwsKmsConfigurationResponse) HasCmkArn() bool`

HasCmkArn returns a boolean if a field has been set.

### SetCmkArnNil

`func (o *AwsKmsConfigurationResponse) SetCmkArnNil(b bool)`

 SetCmkArnNil sets the value for CmkArn to be an explicit nil

### UnsetCmkArn
`func (o *AwsKmsConfigurationResponse) UnsetCmkArn()`

UnsetCmkArn ensures that no value is present for CmkArn, not even an explicit nil
### GetCmkKeyId

`func (o *AwsKmsConfigurationResponse) GetCmkKeyId() string`

GetCmkKeyId returns the CmkKeyId field if non-nil, zero value otherwise.

### GetCmkKeyIdOk

`func (o *AwsKmsConfigurationResponse) GetCmkKeyIdOk() (*string, bool)`

GetCmkKeyIdOk returns a tuple with the CmkKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmkKeyId

`func (o *AwsKmsConfigurationResponse) SetCmkKeyId(v string)`

SetCmkKeyId sets CmkKeyId field to given value.

### HasCmkKeyId

`func (o *AwsKmsConfigurationResponse) HasCmkKeyId() bool`

HasCmkKeyId returns a boolean if a field has been set.

### SetCmkKeyIdNil

`func (o *AwsKmsConfigurationResponse) SetCmkKeyIdNil(b bool)`

 SetCmkKeyIdNil sets the value for CmkKeyId to be an explicit nil

### UnsetCmkKeyId
`func (o *AwsKmsConfigurationResponse) UnsetCmkKeyId()`

UnsetCmkKeyId ensures that no value is present for CmkKeyId, not even an explicit nil
### GetIamRoleArn

`func (o *AwsKmsConfigurationResponse) GetIamRoleArn() string`

GetIamRoleArn returns the IamRoleArn field if non-nil, zero value otherwise.

### GetIamRoleArnOk

`func (o *AwsKmsConfigurationResponse) GetIamRoleArnOk() (*string, bool)`

GetIamRoleArnOk returns a tuple with the IamRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRoleArn

`func (o *AwsKmsConfigurationResponse) SetIamRoleArn(v string)`

SetIamRoleArn sets IamRoleArn field to given value.

### HasIamRoleArn

`func (o *AwsKmsConfigurationResponse) HasIamRoleArn() bool`

HasIamRoleArn returns a boolean if a field has been set.

### SetIamRoleArnNil

`func (o *AwsKmsConfigurationResponse) SetIamRoleArnNil(b bool)`

 SetIamRoleArnNil sets the value for IamRoleArn to be an explicit nil

### UnsetIamRoleArn
`func (o *AwsKmsConfigurationResponse) UnsetIamRoleArn()`

UnsetIamRoleArn ensures that no value is present for IamRoleArn, not even an explicit nil
### GetRegion

`func (o *AwsKmsConfigurationResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AwsKmsConfigurationResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AwsKmsConfigurationResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AwsKmsConfigurationResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *AwsKmsConfigurationResponse) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *AwsKmsConfigurationResponse) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetVerifySSL

`func (o *AwsKmsConfigurationResponse) GetVerifySSL() bool`

GetVerifySSL returns the VerifySSL field if non-nil, zero value otherwise.

### GetVerifySSLOk

`func (o *AwsKmsConfigurationResponse) GetVerifySSLOk() (*bool, bool)`

GetVerifySSLOk returns a tuple with the VerifySSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySSL

`func (o *AwsKmsConfigurationResponse) SetVerifySSL(v bool)`

SetVerifySSL sets VerifySSL field to given value.

### HasVerifySSL

`func (o *AwsKmsConfigurationResponse) HasVerifySSL() bool`

HasVerifySSL returns a boolean if a field has been set.

### SetVerifySSLNil

`func (o *AwsKmsConfigurationResponse) SetVerifySSLNil(b bool)`

 SetVerifySSLNil sets the value for VerifySSL to be an explicit nil

### UnsetVerifySSL
`func (o *AwsKmsConfigurationResponse) UnsetVerifySSL()`

UnsetVerifySSL ensures that no value is present for VerifySSL, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


