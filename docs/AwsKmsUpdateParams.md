# AwsKmsUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | Pointer to **NullableString** | Access key id needed to access the cloud account. When update cluster config, should encrypte accessKeyId with cluster ID. | [optional] 
**CaCertificatePath** | Pointer to **NullableString** | Specify the ca certificate path. | [optional] 
**SecretAccessKey** | Pointer to **NullableString** | Secret access key needed to access the cloud account. This is encrypted with the cluster id. | [optional] 
**VerifySSL** | Pointer to **NullableBool** | Specify whether to verify SSL when connect with AWS KMS. Default is true. | [optional] 

## Methods

### NewAwsKmsUpdateParams

`func NewAwsKmsUpdateParams() *AwsKmsUpdateParams`

NewAwsKmsUpdateParams instantiates a new AwsKmsUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsKmsUpdateParamsWithDefaults

`func NewAwsKmsUpdateParamsWithDefaults() *AwsKmsUpdateParams`

NewAwsKmsUpdateParamsWithDefaults instantiates a new AwsKmsUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyId

`func (o *AwsKmsUpdateParams) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *AwsKmsUpdateParams) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *AwsKmsUpdateParams) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *AwsKmsUpdateParams) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### SetAccessKeyIdNil

`func (o *AwsKmsUpdateParams) SetAccessKeyIdNil(b bool)`

 SetAccessKeyIdNil sets the value for AccessKeyId to be an explicit nil

### UnsetAccessKeyId
`func (o *AwsKmsUpdateParams) UnsetAccessKeyId()`

UnsetAccessKeyId ensures that no value is present for AccessKeyId, not even an explicit nil
### GetCaCertificatePath

`func (o *AwsKmsUpdateParams) GetCaCertificatePath() string`

GetCaCertificatePath returns the CaCertificatePath field if non-nil, zero value otherwise.

### GetCaCertificatePathOk

`func (o *AwsKmsUpdateParams) GetCaCertificatePathOk() (*string, bool)`

GetCaCertificatePathOk returns a tuple with the CaCertificatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificatePath

`func (o *AwsKmsUpdateParams) SetCaCertificatePath(v string)`

SetCaCertificatePath sets CaCertificatePath field to given value.

### HasCaCertificatePath

`func (o *AwsKmsUpdateParams) HasCaCertificatePath() bool`

HasCaCertificatePath returns a boolean if a field has been set.

### SetCaCertificatePathNil

`func (o *AwsKmsUpdateParams) SetCaCertificatePathNil(b bool)`

 SetCaCertificatePathNil sets the value for CaCertificatePath to be an explicit nil

### UnsetCaCertificatePath
`func (o *AwsKmsUpdateParams) UnsetCaCertificatePath()`

UnsetCaCertificatePath ensures that no value is present for CaCertificatePath, not even an explicit nil
### GetSecretAccessKey

`func (o *AwsKmsUpdateParams) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *AwsKmsUpdateParams) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *AwsKmsUpdateParams) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *AwsKmsUpdateParams) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### SetSecretAccessKeyNil

`func (o *AwsKmsUpdateParams) SetSecretAccessKeyNil(b bool)`

 SetSecretAccessKeyNil sets the value for SecretAccessKey to be an explicit nil

### UnsetSecretAccessKey
`func (o *AwsKmsUpdateParams) UnsetSecretAccessKey()`

UnsetSecretAccessKey ensures that no value is present for SecretAccessKey, not even an explicit nil
### GetVerifySSL

`func (o *AwsKmsUpdateParams) GetVerifySSL() bool`

GetVerifySSL returns the VerifySSL field if non-nil, zero value otherwise.

### GetVerifySSLOk

`func (o *AwsKmsUpdateParams) GetVerifySSLOk() (*bool, bool)`

GetVerifySSLOk returns a tuple with the VerifySSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySSL

`func (o *AwsKmsUpdateParams) SetVerifySSL(v bool)`

SetVerifySSL sets VerifySSL field to given value.

### HasVerifySSL

`func (o *AwsKmsUpdateParams) HasVerifySSL() bool`

HasVerifySSL returns a boolean if a field has been set.

### SetVerifySSLNil

`func (o *AwsKmsUpdateParams) SetVerifySSLNil(b bool)`

 SetVerifySSLNil sets the value for VerifySSL to be an explicit nil

### UnsetVerifySSL
`func (o *AwsKmsUpdateParams) UnsetVerifySSL()`

UnsetVerifySSL ensures that no value is present for VerifySSL, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


