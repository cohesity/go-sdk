# S3ConfigForIndexing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**S3BucketName** | **NullableString** | Name of S3 bucket to store indexed data. | 
**S3IamRoleArn** | **NullableString** | IAM role ARN which has access to S3 instance. | 
**S3Prefix** | **NullableString** | Prefix under S3 bucket where data will be stored. | 

## Methods

### NewS3ConfigForIndexing

`func NewS3ConfigForIndexing(s3BucketName NullableString, s3IamRoleArn NullableString, s3Prefix NullableString, ) *S3ConfigForIndexing`

NewS3ConfigForIndexing instantiates a new S3ConfigForIndexing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3ConfigForIndexingWithDefaults

`func NewS3ConfigForIndexingWithDefaults() *S3ConfigForIndexing`

NewS3ConfigForIndexingWithDefaults instantiates a new S3ConfigForIndexing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetS3BucketName

`func (o *S3ConfigForIndexing) GetS3BucketName() string`

GetS3BucketName returns the S3BucketName field if non-nil, zero value otherwise.

### GetS3BucketNameOk

`func (o *S3ConfigForIndexing) GetS3BucketNameOk() (*string, bool)`

GetS3BucketNameOk returns a tuple with the S3BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3BucketName

`func (o *S3ConfigForIndexing) SetS3BucketName(v string)`

SetS3BucketName sets S3BucketName field to given value.


### SetS3BucketNameNil

`func (o *S3ConfigForIndexing) SetS3BucketNameNil(b bool)`

 SetS3BucketNameNil sets the value for S3BucketName to be an explicit nil

### UnsetS3BucketName
`func (o *S3ConfigForIndexing) UnsetS3BucketName()`

UnsetS3BucketName ensures that no value is present for S3BucketName, not even an explicit nil
### GetS3IamRoleArn

`func (o *S3ConfigForIndexing) GetS3IamRoleArn() string`

GetS3IamRoleArn returns the S3IamRoleArn field if non-nil, zero value otherwise.

### GetS3IamRoleArnOk

`func (o *S3ConfigForIndexing) GetS3IamRoleArnOk() (*string, bool)`

GetS3IamRoleArnOk returns a tuple with the S3IamRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3IamRoleArn

`func (o *S3ConfigForIndexing) SetS3IamRoleArn(v string)`

SetS3IamRoleArn sets S3IamRoleArn field to given value.


### SetS3IamRoleArnNil

`func (o *S3ConfigForIndexing) SetS3IamRoleArnNil(b bool)`

 SetS3IamRoleArnNil sets the value for S3IamRoleArn to be an explicit nil

### UnsetS3IamRoleArn
`func (o *S3ConfigForIndexing) UnsetS3IamRoleArn()`

UnsetS3IamRoleArn ensures that no value is present for S3IamRoleArn, not even an explicit nil
### GetS3Prefix

`func (o *S3ConfigForIndexing) GetS3Prefix() string`

GetS3Prefix returns the S3Prefix field if non-nil, zero value otherwise.

### GetS3PrefixOk

`func (o *S3ConfigForIndexing) GetS3PrefixOk() (*string, bool)`

GetS3PrefixOk returns a tuple with the S3Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Prefix

`func (o *S3ConfigForIndexing) SetS3Prefix(v string)`

SetS3Prefix sets S3Prefix field to given value.


### SetS3PrefixNil

`func (o *S3ConfigForIndexing) SetS3PrefixNil(b bool)`

 SetS3PrefixNil sets the value for S3Prefix to be an explicit nil

### UnsetS3Prefix
`func (o *S3ConfigForIndexing) UnsetS3Prefix()`

UnsetS3Prefix ensures that no value is present for S3Prefix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


