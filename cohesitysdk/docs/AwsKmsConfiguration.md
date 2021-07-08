# AwsKmsConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | Pointer to **NullableString** | Access key id needed to access the cloud account. When update cluster config, should encrypte accessKeyId with cluster ID. | [optional] 
**CaCertificate** | Pointer to **NullableString** | Specify the ca certificate path. | [optional] 
**CmkAlias** | Pointer to **NullableString** | The string alias of the CMK. | [optional] 
**CmkArn** | Pointer to **NullableString** | The Amazon Resource Number of AWS Customer Managed Key. | [optional] 
**CmkKeyId** | Pointer to **NullableString** | AWS keyId, and alias. Only need one of them to connect AWS. Alias is better, because keyId maybe rotated by AWS. The unique key id of the CMK. | [optional] 
**Region** | Pointer to **NullableString** | AWS region, e.g. us-east-1, us-west-2, for the AWS Glacier service to be used to authenticate resources within this region by the configured AWS account. | [optional] 
**SecretAccessKey** | Pointer to **NullableString** | Secret access key needed to access the cloud account. This is encrypted with the cluster id. | [optional] 
**VerifySSL** | Pointer to **NullableBool** | Specify whether to verify SSL when connect with AWS KMS. Default is true. | [optional] 

## Methods

### NewAwsKmsConfiguration

`func NewAwsKmsConfiguration() *AwsKmsConfiguration`

NewAwsKmsConfiguration instantiates a new AwsKmsConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsKmsConfigurationWithDefaults

`func NewAwsKmsConfigurationWithDefaults() *AwsKmsConfiguration`

NewAwsKmsConfigurationWithDefaults instantiates a new AwsKmsConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyId

`func (o *AwsKmsConfiguration) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *AwsKmsConfiguration) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *AwsKmsConfiguration) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *AwsKmsConfiguration) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### SetAccessKeyIdNil

`func (o *AwsKmsConfiguration) SetAccessKeyIdNil(b bool)`

 SetAccessKeyIdNil sets the value for AccessKeyId to be an explicit nil

### UnsetAccessKeyId
`func (o *AwsKmsConfiguration) UnsetAccessKeyId()`

UnsetAccessKeyId ensures that no value is present for AccessKeyId, not even an explicit nil
### GetCaCertificate

`func (o *AwsKmsConfiguration) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *AwsKmsConfiguration) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *AwsKmsConfiguration) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *AwsKmsConfiguration) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.

### SetCaCertificateNil

`func (o *AwsKmsConfiguration) SetCaCertificateNil(b bool)`

 SetCaCertificateNil sets the value for CaCertificate to be an explicit nil

### UnsetCaCertificate
`func (o *AwsKmsConfiguration) UnsetCaCertificate()`

UnsetCaCertificate ensures that no value is present for CaCertificate, not even an explicit nil
### GetCmkAlias

`func (o *AwsKmsConfiguration) GetCmkAlias() string`

GetCmkAlias returns the CmkAlias field if non-nil, zero value otherwise.

### GetCmkAliasOk

`func (o *AwsKmsConfiguration) GetCmkAliasOk() (*string, bool)`

GetCmkAliasOk returns a tuple with the CmkAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmkAlias

`func (o *AwsKmsConfiguration) SetCmkAlias(v string)`

SetCmkAlias sets CmkAlias field to given value.

### HasCmkAlias

`func (o *AwsKmsConfiguration) HasCmkAlias() bool`

HasCmkAlias returns a boolean if a field has been set.

### SetCmkAliasNil

`func (o *AwsKmsConfiguration) SetCmkAliasNil(b bool)`

 SetCmkAliasNil sets the value for CmkAlias to be an explicit nil

### UnsetCmkAlias
`func (o *AwsKmsConfiguration) UnsetCmkAlias()`

UnsetCmkAlias ensures that no value is present for CmkAlias, not even an explicit nil
### GetCmkArn

`func (o *AwsKmsConfiguration) GetCmkArn() string`

GetCmkArn returns the CmkArn field if non-nil, zero value otherwise.

### GetCmkArnOk

`func (o *AwsKmsConfiguration) GetCmkArnOk() (*string, bool)`

GetCmkArnOk returns a tuple with the CmkArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmkArn

`func (o *AwsKmsConfiguration) SetCmkArn(v string)`

SetCmkArn sets CmkArn field to given value.

### HasCmkArn

`func (o *AwsKmsConfiguration) HasCmkArn() bool`

HasCmkArn returns a boolean if a field has been set.

### SetCmkArnNil

`func (o *AwsKmsConfiguration) SetCmkArnNil(b bool)`

 SetCmkArnNil sets the value for CmkArn to be an explicit nil

### UnsetCmkArn
`func (o *AwsKmsConfiguration) UnsetCmkArn()`

UnsetCmkArn ensures that no value is present for CmkArn, not even an explicit nil
### GetCmkKeyId

`func (o *AwsKmsConfiguration) GetCmkKeyId() string`

GetCmkKeyId returns the CmkKeyId field if non-nil, zero value otherwise.

### GetCmkKeyIdOk

`func (o *AwsKmsConfiguration) GetCmkKeyIdOk() (*string, bool)`

GetCmkKeyIdOk returns a tuple with the CmkKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmkKeyId

`func (o *AwsKmsConfiguration) SetCmkKeyId(v string)`

SetCmkKeyId sets CmkKeyId field to given value.

### HasCmkKeyId

`func (o *AwsKmsConfiguration) HasCmkKeyId() bool`

HasCmkKeyId returns a boolean if a field has been set.

### SetCmkKeyIdNil

`func (o *AwsKmsConfiguration) SetCmkKeyIdNil(b bool)`

 SetCmkKeyIdNil sets the value for CmkKeyId to be an explicit nil

### UnsetCmkKeyId
`func (o *AwsKmsConfiguration) UnsetCmkKeyId()`

UnsetCmkKeyId ensures that no value is present for CmkKeyId, not even an explicit nil
### GetRegion

`func (o *AwsKmsConfiguration) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AwsKmsConfiguration) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AwsKmsConfiguration) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AwsKmsConfiguration) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *AwsKmsConfiguration) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *AwsKmsConfiguration) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetSecretAccessKey

`func (o *AwsKmsConfiguration) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *AwsKmsConfiguration) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *AwsKmsConfiguration) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *AwsKmsConfiguration) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### SetSecretAccessKeyNil

`func (o *AwsKmsConfiguration) SetSecretAccessKeyNil(b bool)`

 SetSecretAccessKeyNil sets the value for SecretAccessKey to be an explicit nil

### UnsetSecretAccessKey
`func (o *AwsKmsConfiguration) UnsetSecretAccessKey()`

UnsetSecretAccessKey ensures that no value is present for SecretAccessKey, not even an explicit nil
### GetVerifySSL

`func (o *AwsKmsConfiguration) GetVerifySSL() bool`

GetVerifySSL returns the VerifySSL field if non-nil, zero value otherwise.

### GetVerifySSLOk

`func (o *AwsKmsConfiguration) GetVerifySSLOk() (*bool, bool)`

GetVerifySSLOk returns a tuple with the VerifySSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySSL

`func (o *AwsKmsConfiguration) SetVerifySSL(v bool)`

SetVerifySSL sets VerifySSL field to given value.

### HasVerifySSL

`func (o *AwsKmsConfiguration) HasVerifySSL() bool`

HasVerifySSL returns a boolean if a field has been set.

### SetVerifySSLNil

`func (o *AwsKmsConfiguration) SetVerifySSLNil(b bool)`

 SetVerifySSLNil sets the value for VerifySSL to be an explicit nil

### UnsetVerifySSL
`func (o *AwsKmsConfiguration) UnsetVerifySSL()`

UnsetVerifySSL ensures that no value is present for VerifySSL, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


