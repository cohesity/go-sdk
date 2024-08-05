# CommonTieringAwsExternalTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **NullableString** | Specifies bucket name of the External Target. | 
**Region** | **NullableString** | Specifies region of the External Target. | 
**StorageClass** | **NullableString** | Specifies the AWS External Target storage class. | 

## Methods

### NewCommonTieringAwsExternalTargetParams

`func NewCommonTieringAwsExternalTargetParams(bucketName NullableString, region NullableString, storageClass NullableString, ) *CommonTieringAwsExternalTargetParams`

NewCommonTieringAwsExternalTargetParams instantiates a new CommonTieringAwsExternalTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonTieringAwsExternalTargetParamsWithDefaults

`func NewCommonTieringAwsExternalTargetParamsWithDefaults() *CommonTieringAwsExternalTargetParams`

NewCommonTieringAwsExternalTargetParamsWithDefaults instantiates a new CommonTieringAwsExternalTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *CommonTieringAwsExternalTargetParams) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *CommonTieringAwsExternalTargetParams) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *CommonTieringAwsExternalTargetParams) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### SetBucketNameNil

`func (o *CommonTieringAwsExternalTargetParams) SetBucketNameNil(b bool)`

 SetBucketNameNil sets the value for BucketName to be an explicit nil

### UnsetBucketName
`func (o *CommonTieringAwsExternalTargetParams) UnsetBucketName()`

UnsetBucketName ensures that no value is present for BucketName, not even an explicit nil
### GetRegion

`func (o *CommonTieringAwsExternalTargetParams) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CommonTieringAwsExternalTargetParams) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CommonTieringAwsExternalTargetParams) SetRegion(v string)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *CommonTieringAwsExternalTargetParams) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *CommonTieringAwsExternalTargetParams) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetStorageClass

`func (o *CommonTieringAwsExternalTargetParams) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *CommonTieringAwsExternalTargetParams) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *CommonTieringAwsExternalTargetParams) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.


### SetStorageClassNil

`func (o *CommonTieringAwsExternalTargetParams) SetStorageClassNil(b bool)`

 SetStorageClassNil sets the value for StorageClass to be an explicit nil

### UnsetStorageClass
`func (o *CommonTieringAwsExternalTargetParams) UnsetStorageClass()`

UnsetStorageClass ensures that no value is present for StorageClass, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


