# S3ConfigBucketPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specifies the identifier of the bucket policy. This is a read-only property. | [optional] 
**RawPolicy** | **NullableString** | Specifies the raw JSON string of the store policy. | 
**Version** | Pointer to **NullableString** | Specifies the language syntax rules that are to be used to process the policy. This is a read-only property. | [optional] 

## Methods

### NewS3ConfigBucketPolicy

`func NewS3ConfigBucketPolicy(rawPolicy NullableString, ) *S3ConfigBucketPolicy`

NewS3ConfigBucketPolicy instantiates a new S3ConfigBucketPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3ConfigBucketPolicyWithDefaults

`func NewS3ConfigBucketPolicyWithDefaults() *S3ConfigBucketPolicy`

NewS3ConfigBucketPolicyWithDefaults instantiates a new S3ConfigBucketPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *S3ConfigBucketPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *S3ConfigBucketPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *S3ConfigBucketPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *S3ConfigBucketPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *S3ConfigBucketPolicy) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *S3ConfigBucketPolicy) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetRawPolicy

`func (o *S3ConfigBucketPolicy) GetRawPolicy() string`

GetRawPolicy returns the RawPolicy field if non-nil, zero value otherwise.

### GetRawPolicyOk

`func (o *S3ConfigBucketPolicy) GetRawPolicyOk() (*string, bool)`

GetRawPolicyOk returns a tuple with the RawPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawPolicy

`func (o *S3ConfigBucketPolicy) SetRawPolicy(v string)`

SetRawPolicy sets RawPolicy field to given value.


### SetRawPolicyNil

`func (o *S3ConfigBucketPolicy) SetRawPolicyNil(b bool)`

 SetRawPolicyNil sets the value for RawPolicy to be an explicit nil

### UnsetRawPolicy
`func (o *S3ConfigBucketPolicy) UnsetRawPolicy()`

UnsetRawPolicy ensures that no value is present for RawPolicy, not even an explicit nil
### GetVersion

`func (o *S3ConfigBucketPolicy) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *S3ConfigBucketPolicy) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *S3ConfigBucketPolicy) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *S3ConfigBucketPolicy) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *S3ConfigBucketPolicy) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *S3ConfigBucketPolicy) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


