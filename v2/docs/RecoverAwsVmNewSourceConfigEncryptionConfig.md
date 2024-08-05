# RecoverAwsVmNewSourceConfigEncryptionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomKmsKeyArn** | Pointer to **NullableString** | Specifies custom KMS key arn. It will be of form arn:aws:kms:&lt;region&gt;:&lt;account_id&gt;:key/&lt;key_id&gt; | [optional] 
**KmsKey** | Pointer to [**NullableEncryptionConfigKmsKey**](EncryptionConfigKmsKey.md) |  | [optional] 
**ShouldEncrypt** | **NullableBool** | Specifies whether to encrypt recovered volumes or not. Default value is true. | 

## Methods

### NewRecoverAwsVmNewSourceConfigEncryptionConfig

`func NewRecoverAwsVmNewSourceConfigEncryptionConfig(shouldEncrypt NullableBool, ) *RecoverAwsVmNewSourceConfigEncryptionConfig`

NewRecoverAwsVmNewSourceConfigEncryptionConfig instantiates a new RecoverAwsVmNewSourceConfigEncryptionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAwsVmNewSourceConfigEncryptionConfigWithDefaults

`func NewRecoverAwsVmNewSourceConfigEncryptionConfigWithDefaults() *RecoverAwsVmNewSourceConfigEncryptionConfig`

NewRecoverAwsVmNewSourceConfigEncryptionConfigWithDefaults instantiates a new RecoverAwsVmNewSourceConfigEncryptionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomKmsKeyArn

`func (o *RecoverAwsVmNewSourceConfigEncryptionConfig) GetCustomKmsKeyArn() string`

GetCustomKmsKeyArn returns the CustomKmsKeyArn field if non-nil, zero value otherwise.

### GetCustomKmsKeyArnOk

`func (o *RecoverAwsVmNewSourceConfigEncryptionConfig) GetCustomKmsKeyArnOk() (*string, bool)`

GetCustomKmsKeyArnOk returns a tuple with the CustomKmsKeyArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomKmsKeyArn

`func (o *RecoverAwsVmNewSourceConfigEncryptionConfig) SetCustomKmsKeyArn(v string)`

SetCustomKmsKeyArn sets CustomKmsKeyArn field to given value.

### HasCustomKmsKeyArn

`func (o *RecoverAwsVmNewSourceConfigEncryptionConfig) HasCustomKmsKeyArn() bool`

HasCustomKmsKeyArn returns a boolean if a field has been set.

### SetCustomKmsKeyArnNil

`func (o *RecoverAwsVmNewSourceConfigEncryptionConfig) SetCustomKmsKeyArnNil(b bool)`

 SetCustomKmsKeyArnNil sets the value for CustomKmsKeyArn to be an explicit nil

### UnsetCustomKmsKeyArn
`func (o *RecoverAwsVmNewSourceConfigEncryptionConfig) UnsetCustomKmsKeyArn()`

UnsetCustomKmsKeyArn ensures that no value is present for CustomKmsKeyArn, not even an explicit nil
### GetKmsKey

`func (o *RecoverAwsVmNewSourceConfigEncryptionConfig) GetKmsKey() EncryptionConfigKmsKey`

GetKmsKey returns the KmsKey field if non-nil, zero value otherwise.

### GetKmsKeyOk

`func (o *RecoverAwsVmNewSourceConfigEncryptionConfig) GetKmsKeyOk() (*EncryptionConfigKmsKey, bool)`

GetKmsKeyOk returns a tuple with the KmsKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsKey

`func (o *RecoverAwsVmNewSourceConfigEncryptionConfig) SetKmsKey(v EncryptionConfigKmsKey)`

SetKmsKey sets KmsKey field to given value.

### HasKmsKey

`func (o *RecoverAwsVmNewSourceConfigEncryptionConfig) HasKmsKey() bool`

HasKmsKey returns a boolean if a field has been set.

### SetKmsKeyNil

`func (o *RecoverAwsVmNewSourceConfigEncryptionConfig) SetKmsKeyNil(b bool)`

 SetKmsKeyNil sets the value for KmsKey to be an explicit nil

### UnsetKmsKey
`func (o *RecoverAwsVmNewSourceConfigEncryptionConfig) UnsetKmsKey()`

UnsetKmsKey ensures that no value is present for KmsKey, not even an explicit nil
### GetShouldEncrypt

`func (o *RecoverAwsVmNewSourceConfigEncryptionConfig) GetShouldEncrypt() bool`

GetShouldEncrypt returns the ShouldEncrypt field if non-nil, zero value otherwise.

### GetShouldEncryptOk

`func (o *RecoverAwsVmNewSourceConfigEncryptionConfig) GetShouldEncryptOk() (*bool, bool)`

GetShouldEncryptOk returns a tuple with the ShouldEncrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldEncrypt

`func (o *RecoverAwsVmNewSourceConfigEncryptionConfig) SetShouldEncrypt(v bool)`

SetShouldEncrypt sets ShouldEncrypt field to given value.


### SetShouldEncryptNil

`func (o *RecoverAwsVmNewSourceConfigEncryptionConfig) SetShouldEncryptNil(b bool)`

 SetShouldEncryptNil sets the value for ShouldEncrypt to be an explicit nil

### UnsetShouldEncrypt
`func (o *RecoverAwsVmNewSourceConfigEncryptionConfig) UnsetShouldEncrypt()`

UnsetShouldEncrypt ensures that no value is present for ShouldEncrypt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


