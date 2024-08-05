# AwsRecoverS3NewTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | [**NullableAwsRecoverS3NewTargetConfigBucket**](AwsRecoverS3NewTargetConfigBucket.md) |  | 
**Region** | [**NullableAwsRecoverS3NewTargetConfigRegion**](AwsRecoverS3NewTargetConfigRegion.md) |  | 
**Source** | [**NullableAwsRecoverS3NewTargetConfigSource**](AwsRecoverS3NewTargetConfigSource.md) |  | 

## Methods

### NewAwsRecoverS3NewTargetConfig

`func NewAwsRecoverS3NewTargetConfig(bucket NullableAwsRecoverS3NewTargetConfigBucket, region NullableAwsRecoverS3NewTargetConfigRegion, source NullableAwsRecoverS3NewTargetConfigSource, ) *AwsRecoverS3NewTargetConfig`

NewAwsRecoverS3NewTargetConfig instantiates a new AwsRecoverS3NewTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsRecoverS3NewTargetConfigWithDefaults

`func NewAwsRecoverS3NewTargetConfigWithDefaults() *AwsRecoverS3NewTargetConfig`

NewAwsRecoverS3NewTargetConfigWithDefaults instantiates a new AwsRecoverS3NewTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *AwsRecoverS3NewTargetConfig) GetBucket() AwsRecoverS3NewTargetConfigBucket`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *AwsRecoverS3NewTargetConfig) GetBucketOk() (*AwsRecoverS3NewTargetConfigBucket, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *AwsRecoverS3NewTargetConfig) SetBucket(v AwsRecoverS3NewTargetConfigBucket)`

SetBucket sets Bucket field to given value.


### SetBucketNil

`func (o *AwsRecoverS3NewTargetConfig) SetBucketNil(b bool)`

 SetBucketNil sets the value for Bucket to be an explicit nil

### UnsetBucket
`func (o *AwsRecoverS3NewTargetConfig) UnsetBucket()`

UnsetBucket ensures that no value is present for Bucket, not even an explicit nil
### GetRegion

`func (o *AwsRecoverS3NewTargetConfig) GetRegion() AwsRecoverS3NewTargetConfigRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AwsRecoverS3NewTargetConfig) GetRegionOk() (*AwsRecoverS3NewTargetConfigRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AwsRecoverS3NewTargetConfig) SetRegion(v AwsRecoverS3NewTargetConfigRegion)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *AwsRecoverS3NewTargetConfig) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *AwsRecoverS3NewTargetConfig) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetSource

`func (o *AwsRecoverS3NewTargetConfig) GetSource() AwsRecoverS3NewTargetConfigSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AwsRecoverS3NewTargetConfig) GetSourceOk() (*AwsRecoverS3NewTargetConfigSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AwsRecoverS3NewTargetConfig) SetSource(v AwsRecoverS3NewTargetConfigSource)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *AwsRecoverS3NewTargetConfig) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *AwsRecoverS3NewTargetConfig) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


