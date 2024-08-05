# CommonAwsExternalTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **NullableString** | Specifies bucket name of the External Target. | 
**Region** | **NullableString** | Specifies region of the External Target. | 

## Methods

### NewCommonAwsExternalTargetParams

`func NewCommonAwsExternalTargetParams(bucketName NullableString, region NullableString, ) *CommonAwsExternalTargetParams`

NewCommonAwsExternalTargetParams instantiates a new CommonAwsExternalTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonAwsExternalTargetParamsWithDefaults

`func NewCommonAwsExternalTargetParamsWithDefaults() *CommonAwsExternalTargetParams`

NewCommonAwsExternalTargetParamsWithDefaults instantiates a new CommonAwsExternalTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *CommonAwsExternalTargetParams) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *CommonAwsExternalTargetParams) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *CommonAwsExternalTargetParams) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### SetBucketNameNil

`func (o *CommonAwsExternalTargetParams) SetBucketNameNil(b bool)`

 SetBucketNameNil sets the value for BucketName to be an explicit nil

### UnsetBucketName
`func (o *CommonAwsExternalTargetParams) UnsetBucketName()`

UnsetBucketName ensures that no value is present for BucketName, not even an explicit nil
### GetRegion

`func (o *CommonAwsExternalTargetParams) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CommonAwsExternalTargetParams) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CommonAwsExternalTargetParams) SetRegion(v string)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *CommonAwsExternalTargetParams) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *CommonAwsExternalTargetParams) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


