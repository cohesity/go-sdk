# ArchivalS3CompExternalTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **NullableString** | Specifies the bucket name of the external target. | 
**AccessKeyId** | **NullableString** | Specifies the access key id of the external target. | 
**SecretAccessKey** | **NullableString** | Specifies the secret access key of the external target. | 
**EndPoint** | **NullableString** | Specifies the endpoint of the external target. | 
**SecureConnection** | Pointer to **NullableBool** | Specifies the secure connection(https) is enabled or not. | [optional] 
**SignatureVersion** | Pointer to **NullableInt32** | Specifies the aws signature version of the external target. | [optional] 
**SourceSideDeduplication** | Pointer to **NullableBool** | Specifies the Source Side Deduplication setting for the S3 Compatible external target | [optional] 
**IsIncrementalArchivalEnabled** | Pointer to **NullableBool** | Specifies if Incremental Archival setting is enabled or not. | [optional] 
**IsForeverIncrementalArchivalEnabled** | Pointer to **NullableBool** | Specifies if Forever Incremental Archival setting is enabled or not. | [optional] 

## Methods

### NewArchivalS3CompExternalTargetParams

`func NewArchivalS3CompExternalTargetParams(bucketName NullableString, accessKeyId NullableString, secretAccessKey NullableString, endPoint NullableString, ) *ArchivalS3CompExternalTargetParams`

NewArchivalS3CompExternalTargetParams instantiates a new ArchivalS3CompExternalTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchivalS3CompExternalTargetParamsWithDefaults

`func NewArchivalS3CompExternalTargetParamsWithDefaults() *ArchivalS3CompExternalTargetParams`

NewArchivalS3CompExternalTargetParamsWithDefaults instantiates a new ArchivalS3CompExternalTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *ArchivalS3CompExternalTargetParams) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *ArchivalS3CompExternalTargetParams) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *ArchivalS3CompExternalTargetParams) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### SetBucketNameNil

`func (o *ArchivalS3CompExternalTargetParams) SetBucketNameNil(b bool)`

 SetBucketNameNil sets the value for BucketName to be an explicit nil

### UnsetBucketName
`func (o *ArchivalS3CompExternalTargetParams) UnsetBucketName()`

UnsetBucketName ensures that no value is present for BucketName, not even an explicit nil
### GetAccessKeyId

`func (o *ArchivalS3CompExternalTargetParams) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *ArchivalS3CompExternalTargetParams) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *ArchivalS3CompExternalTargetParams) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.


### SetAccessKeyIdNil

`func (o *ArchivalS3CompExternalTargetParams) SetAccessKeyIdNil(b bool)`

 SetAccessKeyIdNil sets the value for AccessKeyId to be an explicit nil

### UnsetAccessKeyId
`func (o *ArchivalS3CompExternalTargetParams) UnsetAccessKeyId()`

UnsetAccessKeyId ensures that no value is present for AccessKeyId, not even an explicit nil
### GetSecretAccessKey

`func (o *ArchivalS3CompExternalTargetParams) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *ArchivalS3CompExternalTargetParams) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *ArchivalS3CompExternalTargetParams) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.


### SetSecretAccessKeyNil

`func (o *ArchivalS3CompExternalTargetParams) SetSecretAccessKeyNil(b bool)`

 SetSecretAccessKeyNil sets the value for SecretAccessKey to be an explicit nil

### UnsetSecretAccessKey
`func (o *ArchivalS3CompExternalTargetParams) UnsetSecretAccessKey()`

UnsetSecretAccessKey ensures that no value is present for SecretAccessKey, not even an explicit nil
### GetEndPoint

`func (o *ArchivalS3CompExternalTargetParams) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *ArchivalS3CompExternalTargetParams) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *ArchivalS3CompExternalTargetParams) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.


### SetEndPointNil

`func (o *ArchivalS3CompExternalTargetParams) SetEndPointNil(b bool)`

 SetEndPointNil sets the value for EndPoint to be an explicit nil

### UnsetEndPoint
`func (o *ArchivalS3CompExternalTargetParams) UnsetEndPoint()`

UnsetEndPoint ensures that no value is present for EndPoint, not even an explicit nil
### GetSecureConnection

`func (o *ArchivalS3CompExternalTargetParams) GetSecureConnection() bool`

GetSecureConnection returns the SecureConnection field if non-nil, zero value otherwise.

### GetSecureConnectionOk

`func (o *ArchivalS3CompExternalTargetParams) GetSecureConnectionOk() (*bool, bool)`

GetSecureConnectionOk returns a tuple with the SecureConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureConnection

`func (o *ArchivalS3CompExternalTargetParams) SetSecureConnection(v bool)`

SetSecureConnection sets SecureConnection field to given value.

### HasSecureConnection

`func (o *ArchivalS3CompExternalTargetParams) HasSecureConnection() bool`

HasSecureConnection returns a boolean if a field has been set.

### SetSecureConnectionNil

`func (o *ArchivalS3CompExternalTargetParams) SetSecureConnectionNil(b bool)`

 SetSecureConnectionNil sets the value for SecureConnection to be an explicit nil

### UnsetSecureConnection
`func (o *ArchivalS3CompExternalTargetParams) UnsetSecureConnection()`

UnsetSecureConnection ensures that no value is present for SecureConnection, not even an explicit nil
### GetSignatureVersion

`func (o *ArchivalS3CompExternalTargetParams) GetSignatureVersion() int32`

GetSignatureVersion returns the SignatureVersion field if non-nil, zero value otherwise.

### GetSignatureVersionOk

`func (o *ArchivalS3CompExternalTargetParams) GetSignatureVersionOk() (*int32, bool)`

GetSignatureVersionOk returns a tuple with the SignatureVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureVersion

`func (o *ArchivalS3CompExternalTargetParams) SetSignatureVersion(v int32)`

SetSignatureVersion sets SignatureVersion field to given value.

### HasSignatureVersion

`func (o *ArchivalS3CompExternalTargetParams) HasSignatureVersion() bool`

HasSignatureVersion returns a boolean if a field has been set.

### SetSignatureVersionNil

`func (o *ArchivalS3CompExternalTargetParams) SetSignatureVersionNil(b bool)`

 SetSignatureVersionNil sets the value for SignatureVersion to be an explicit nil

### UnsetSignatureVersion
`func (o *ArchivalS3CompExternalTargetParams) UnsetSignatureVersion()`

UnsetSignatureVersion ensures that no value is present for SignatureVersion, not even an explicit nil
### GetSourceSideDeduplication

`func (o *ArchivalS3CompExternalTargetParams) GetSourceSideDeduplication() bool`

GetSourceSideDeduplication returns the SourceSideDeduplication field if non-nil, zero value otherwise.

### GetSourceSideDeduplicationOk

`func (o *ArchivalS3CompExternalTargetParams) GetSourceSideDeduplicationOk() (*bool, bool)`

GetSourceSideDeduplicationOk returns a tuple with the SourceSideDeduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSideDeduplication

`func (o *ArchivalS3CompExternalTargetParams) SetSourceSideDeduplication(v bool)`

SetSourceSideDeduplication sets SourceSideDeduplication field to given value.

### HasSourceSideDeduplication

`func (o *ArchivalS3CompExternalTargetParams) HasSourceSideDeduplication() bool`

HasSourceSideDeduplication returns a boolean if a field has been set.

### SetSourceSideDeduplicationNil

`func (o *ArchivalS3CompExternalTargetParams) SetSourceSideDeduplicationNil(b bool)`

 SetSourceSideDeduplicationNil sets the value for SourceSideDeduplication to be an explicit nil

### UnsetSourceSideDeduplication
`func (o *ArchivalS3CompExternalTargetParams) UnsetSourceSideDeduplication()`

UnsetSourceSideDeduplication ensures that no value is present for SourceSideDeduplication, not even an explicit nil
### GetIsIncrementalArchivalEnabled

`func (o *ArchivalS3CompExternalTargetParams) GetIsIncrementalArchivalEnabled() bool`

GetIsIncrementalArchivalEnabled returns the IsIncrementalArchivalEnabled field if non-nil, zero value otherwise.

### GetIsIncrementalArchivalEnabledOk

`func (o *ArchivalS3CompExternalTargetParams) GetIsIncrementalArchivalEnabledOk() (*bool, bool)`

GetIsIncrementalArchivalEnabledOk returns a tuple with the IsIncrementalArchivalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncrementalArchivalEnabled

`func (o *ArchivalS3CompExternalTargetParams) SetIsIncrementalArchivalEnabled(v bool)`

SetIsIncrementalArchivalEnabled sets IsIncrementalArchivalEnabled field to given value.

### HasIsIncrementalArchivalEnabled

`func (o *ArchivalS3CompExternalTargetParams) HasIsIncrementalArchivalEnabled() bool`

HasIsIncrementalArchivalEnabled returns a boolean if a field has been set.

### SetIsIncrementalArchivalEnabledNil

`func (o *ArchivalS3CompExternalTargetParams) SetIsIncrementalArchivalEnabledNil(b bool)`

 SetIsIncrementalArchivalEnabledNil sets the value for IsIncrementalArchivalEnabled to be an explicit nil

### UnsetIsIncrementalArchivalEnabled
`func (o *ArchivalS3CompExternalTargetParams) UnsetIsIncrementalArchivalEnabled()`

UnsetIsIncrementalArchivalEnabled ensures that no value is present for IsIncrementalArchivalEnabled, not even an explicit nil
### GetIsForeverIncrementalArchivalEnabled

`func (o *ArchivalS3CompExternalTargetParams) GetIsForeverIncrementalArchivalEnabled() bool`

GetIsForeverIncrementalArchivalEnabled returns the IsForeverIncrementalArchivalEnabled field if non-nil, zero value otherwise.

### GetIsForeverIncrementalArchivalEnabledOk

`func (o *ArchivalS3CompExternalTargetParams) GetIsForeverIncrementalArchivalEnabledOk() (*bool, bool)`

GetIsForeverIncrementalArchivalEnabledOk returns a tuple with the IsForeverIncrementalArchivalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForeverIncrementalArchivalEnabled

`func (o *ArchivalS3CompExternalTargetParams) SetIsForeverIncrementalArchivalEnabled(v bool)`

SetIsForeverIncrementalArchivalEnabled sets IsForeverIncrementalArchivalEnabled field to given value.

### HasIsForeverIncrementalArchivalEnabled

`func (o *ArchivalS3CompExternalTargetParams) HasIsForeverIncrementalArchivalEnabled() bool`

HasIsForeverIncrementalArchivalEnabled returns a boolean if a field has been set.

### SetIsForeverIncrementalArchivalEnabledNil

`func (o *ArchivalS3CompExternalTargetParams) SetIsForeverIncrementalArchivalEnabledNil(b bool)`

 SetIsForeverIncrementalArchivalEnabledNil sets the value for IsForeverIncrementalArchivalEnabled to be an explicit nil

### UnsetIsForeverIncrementalArchivalEnabled
`func (o *ArchivalS3CompExternalTargetParams) UnsetIsForeverIncrementalArchivalEnabled()`

UnsetIsForeverIncrementalArchivalEnabled ensures that no value is present for IsForeverIncrementalArchivalEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


