# AwsKmsConfigurationUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | Pointer to **NullableString** | AWS account access key id. Required when &#39;iamRoleArn&#39; is not given. | [optional] 
**CaCertificate** | Pointer to **NullableString** | Specify the ca certificate. | [optional] 
**IamRoleArn** | Pointer to **NullableString** | The IAM role which will be used to authenticate with AWS KMS. Required when &#39;accessKeyId&#39; and &#39;secretAccessKey&#39; fields are not provided. | [optional] 
**SecretAccessKey** | Pointer to **NullableString** | AWS account secret access key. Required when &#39;iamRoleArn&#39; is not given. | [optional] 
**VerifySSL** | Pointer to **NullableBool** | Enable SSL verification or not. | [optional] [default to true]

## Methods

### NewAwsKmsConfigurationUpdateParams

`func NewAwsKmsConfigurationUpdateParams() *AwsKmsConfigurationUpdateParams`

NewAwsKmsConfigurationUpdateParams instantiates a new AwsKmsConfigurationUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsKmsConfigurationUpdateParamsWithDefaults

`func NewAwsKmsConfigurationUpdateParamsWithDefaults() *AwsKmsConfigurationUpdateParams`

NewAwsKmsConfigurationUpdateParamsWithDefaults instantiates a new AwsKmsConfigurationUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyId

`func (o *AwsKmsConfigurationUpdateParams) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *AwsKmsConfigurationUpdateParams) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *AwsKmsConfigurationUpdateParams) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *AwsKmsConfigurationUpdateParams) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### SetAccessKeyIdNil

`func (o *AwsKmsConfigurationUpdateParams) SetAccessKeyIdNil(b bool)`

 SetAccessKeyIdNil sets the value for AccessKeyId to be an explicit nil

### UnsetAccessKeyId
`func (o *AwsKmsConfigurationUpdateParams) UnsetAccessKeyId()`

UnsetAccessKeyId ensures that no value is present for AccessKeyId, not even an explicit nil
### GetCaCertificate

`func (o *AwsKmsConfigurationUpdateParams) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *AwsKmsConfigurationUpdateParams) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *AwsKmsConfigurationUpdateParams) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *AwsKmsConfigurationUpdateParams) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.

### SetCaCertificateNil

`func (o *AwsKmsConfigurationUpdateParams) SetCaCertificateNil(b bool)`

 SetCaCertificateNil sets the value for CaCertificate to be an explicit nil

### UnsetCaCertificate
`func (o *AwsKmsConfigurationUpdateParams) UnsetCaCertificate()`

UnsetCaCertificate ensures that no value is present for CaCertificate, not even an explicit nil
### GetIamRoleArn

`func (o *AwsKmsConfigurationUpdateParams) GetIamRoleArn() string`

GetIamRoleArn returns the IamRoleArn field if non-nil, zero value otherwise.

### GetIamRoleArnOk

`func (o *AwsKmsConfigurationUpdateParams) GetIamRoleArnOk() (*string, bool)`

GetIamRoleArnOk returns a tuple with the IamRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRoleArn

`func (o *AwsKmsConfigurationUpdateParams) SetIamRoleArn(v string)`

SetIamRoleArn sets IamRoleArn field to given value.

### HasIamRoleArn

`func (o *AwsKmsConfigurationUpdateParams) HasIamRoleArn() bool`

HasIamRoleArn returns a boolean if a field has been set.

### SetIamRoleArnNil

`func (o *AwsKmsConfigurationUpdateParams) SetIamRoleArnNil(b bool)`

 SetIamRoleArnNil sets the value for IamRoleArn to be an explicit nil

### UnsetIamRoleArn
`func (o *AwsKmsConfigurationUpdateParams) UnsetIamRoleArn()`

UnsetIamRoleArn ensures that no value is present for IamRoleArn, not even an explicit nil
### GetSecretAccessKey

`func (o *AwsKmsConfigurationUpdateParams) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *AwsKmsConfigurationUpdateParams) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *AwsKmsConfigurationUpdateParams) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *AwsKmsConfigurationUpdateParams) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### SetSecretAccessKeyNil

`func (o *AwsKmsConfigurationUpdateParams) SetSecretAccessKeyNil(b bool)`

 SetSecretAccessKeyNil sets the value for SecretAccessKey to be an explicit nil

### UnsetSecretAccessKey
`func (o *AwsKmsConfigurationUpdateParams) UnsetSecretAccessKey()`

UnsetSecretAccessKey ensures that no value is present for SecretAccessKey, not even an explicit nil
### GetVerifySSL

`func (o *AwsKmsConfigurationUpdateParams) GetVerifySSL() bool`

GetVerifySSL returns the VerifySSL field if non-nil, zero value otherwise.

### GetVerifySSLOk

`func (o *AwsKmsConfigurationUpdateParams) GetVerifySSLOk() (*bool, bool)`

GetVerifySSLOk returns a tuple with the VerifySSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySSL

`func (o *AwsKmsConfigurationUpdateParams) SetVerifySSL(v bool)`

SetVerifySSL sets VerifySSL field to given value.

### HasVerifySSL

`func (o *AwsKmsConfigurationUpdateParams) HasVerifySSL() bool`

HasVerifySSL returns a boolean if a field has been set.

### SetVerifySSLNil

`func (o *AwsKmsConfigurationUpdateParams) SetVerifySSLNil(b bool)`

 SetVerifySSLNil sets the value for VerifySSL to be an explicit nil

### UnsetVerifySSL
`func (o *AwsKmsConfigurationUpdateParams) UnsetVerifySSL()`

UnsetVerifySSL ensures that no value is present for VerifySSL, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


