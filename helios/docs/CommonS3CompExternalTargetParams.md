# CommonS3CompExternalTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **NullableString** | Specifies the bucket name of the external target. | 
**AccessKeyId** | **NullableString** | Specifies the access key id of the external target. | 
**SecretAccessKey** | **NullableString** | Specifies the secret access key of the external target. | 
**EndPoint** | **NullableString** | Specifies the endpoint of the external target. | 
**SecureConnection** | Pointer to **NullableBool** | Specifies the secure connection(https) is enabled or not. | [optional] 
**SignatureVersion** | Pointer to **NullableInt32** | Specifies the aws signature version of the external target. | [optional] 

## Methods

### NewCommonS3CompExternalTargetParams

`func NewCommonS3CompExternalTargetParams(bucketName NullableString, accessKeyId NullableString, secretAccessKey NullableString, endPoint NullableString, ) *CommonS3CompExternalTargetParams`

NewCommonS3CompExternalTargetParams instantiates a new CommonS3CompExternalTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonS3CompExternalTargetParamsWithDefaults

`func NewCommonS3CompExternalTargetParamsWithDefaults() *CommonS3CompExternalTargetParams`

NewCommonS3CompExternalTargetParamsWithDefaults instantiates a new CommonS3CompExternalTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *CommonS3CompExternalTargetParams) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *CommonS3CompExternalTargetParams) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *CommonS3CompExternalTargetParams) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### SetBucketNameNil

`func (o *CommonS3CompExternalTargetParams) SetBucketNameNil(b bool)`

 SetBucketNameNil sets the value for BucketName to be an explicit nil

### UnsetBucketName
`func (o *CommonS3CompExternalTargetParams) UnsetBucketName()`

UnsetBucketName ensures that no value is present for BucketName, not even an explicit nil
### GetAccessKeyId

`func (o *CommonS3CompExternalTargetParams) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *CommonS3CompExternalTargetParams) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *CommonS3CompExternalTargetParams) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.


### SetAccessKeyIdNil

`func (o *CommonS3CompExternalTargetParams) SetAccessKeyIdNil(b bool)`

 SetAccessKeyIdNil sets the value for AccessKeyId to be an explicit nil

### UnsetAccessKeyId
`func (o *CommonS3CompExternalTargetParams) UnsetAccessKeyId()`

UnsetAccessKeyId ensures that no value is present for AccessKeyId, not even an explicit nil
### GetSecretAccessKey

`func (o *CommonS3CompExternalTargetParams) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *CommonS3CompExternalTargetParams) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *CommonS3CompExternalTargetParams) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.


### SetSecretAccessKeyNil

`func (o *CommonS3CompExternalTargetParams) SetSecretAccessKeyNil(b bool)`

 SetSecretAccessKeyNil sets the value for SecretAccessKey to be an explicit nil

### UnsetSecretAccessKey
`func (o *CommonS3CompExternalTargetParams) UnsetSecretAccessKey()`

UnsetSecretAccessKey ensures that no value is present for SecretAccessKey, not even an explicit nil
### GetEndPoint

`func (o *CommonS3CompExternalTargetParams) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *CommonS3CompExternalTargetParams) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *CommonS3CompExternalTargetParams) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.


### SetEndPointNil

`func (o *CommonS3CompExternalTargetParams) SetEndPointNil(b bool)`

 SetEndPointNil sets the value for EndPoint to be an explicit nil

### UnsetEndPoint
`func (o *CommonS3CompExternalTargetParams) UnsetEndPoint()`

UnsetEndPoint ensures that no value is present for EndPoint, not even an explicit nil
### GetSecureConnection

`func (o *CommonS3CompExternalTargetParams) GetSecureConnection() bool`

GetSecureConnection returns the SecureConnection field if non-nil, zero value otherwise.

### GetSecureConnectionOk

`func (o *CommonS3CompExternalTargetParams) GetSecureConnectionOk() (*bool, bool)`

GetSecureConnectionOk returns a tuple with the SecureConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureConnection

`func (o *CommonS3CompExternalTargetParams) SetSecureConnection(v bool)`

SetSecureConnection sets SecureConnection field to given value.

### HasSecureConnection

`func (o *CommonS3CompExternalTargetParams) HasSecureConnection() bool`

HasSecureConnection returns a boolean if a field has been set.

### SetSecureConnectionNil

`func (o *CommonS3CompExternalTargetParams) SetSecureConnectionNil(b bool)`

 SetSecureConnectionNil sets the value for SecureConnection to be an explicit nil

### UnsetSecureConnection
`func (o *CommonS3CompExternalTargetParams) UnsetSecureConnection()`

UnsetSecureConnection ensures that no value is present for SecureConnection, not even an explicit nil
### GetSignatureVersion

`func (o *CommonS3CompExternalTargetParams) GetSignatureVersion() int32`

GetSignatureVersion returns the SignatureVersion field if non-nil, zero value otherwise.

### GetSignatureVersionOk

`func (o *CommonS3CompExternalTargetParams) GetSignatureVersionOk() (*int32, bool)`

GetSignatureVersionOk returns a tuple with the SignatureVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureVersion

`func (o *CommonS3CompExternalTargetParams) SetSignatureVersion(v int32)`

SetSignatureVersion sets SignatureVersion field to given value.

### HasSignatureVersion

`func (o *CommonS3CompExternalTargetParams) HasSignatureVersion() bool`

HasSignatureVersion returns a boolean if a field has been set.

### SetSignatureVersionNil

`func (o *CommonS3CompExternalTargetParams) SetSignatureVersionNil(b bool)`

 SetSignatureVersionNil sets the value for SignatureVersion to be an explicit nil

### UnsetSignatureVersion
`func (o *CommonS3CompExternalTargetParams) UnsetSignatureVersion()`

UnsetSignatureVersion ensures that no value is present for SignatureVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


