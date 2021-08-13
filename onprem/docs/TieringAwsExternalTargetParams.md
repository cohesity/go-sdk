# TieringAwsExternalTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **NullableString** | Specifies bucket name of the External Target. | 
**Region** | **NullableString** | Specifies region of the External Target. | 
**StorageClass** | **NullableString** | Specifies the AWS External Target storage class. | 
**AwsS3StandardParams** | Pointer to [**AwsS3StandardParams**](AwsS3StandardParams.md) |  | [optional] 
**AwsS3IntelligentParams** | Pointer to [**AwsS3IntelligentParams**](AwsS3IntelligentParams.md) |  | [optional] 

## Methods

### NewTieringAwsExternalTargetParams

`func NewTieringAwsExternalTargetParams(bucketName NullableString, region NullableString, storageClass NullableString, ) *TieringAwsExternalTargetParams`

NewTieringAwsExternalTargetParams instantiates a new TieringAwsExternalTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTieringAwsExternalTargetParamsWithDefaults

`func NewTieringAwsExternalTargetParamsWithDefaults() *TieringAwsExternalTargetParams`

NewTieringAwsExternalTargetParamsWithDefaults instantiates a new TieringAwsExternalTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *TieringAwsExternalTargetParams) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *TieringAwsExternalTargetParams) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *TieringAwsExternalTargetParams) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### SetBucketNameNil

`func (o *TieringAwsExternalTargetParams) SetBucketNameNil(b bool)`

 SetBucketNameNil sets the value for BucketName to be an explicit nil

### UnsetBucketName
`func (o *TieringAwsExternalTargetParams) UnsetBucketName()`

UnsetBucketName ensures that no value is present for BucketName, not even an explicit nil
### GetRegion

`func (o *TieringAwsExternalTargetParams) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *TieringAwsExternalTargetParams) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *TieringAwsExternalTargetParams) SetRegion(v string)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *TieringAwsExternalTargetParams) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *TieringAwsExternalTargetParams) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetStorageClass

`func (o *TieringAwsExternalTargetParams) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *TieringAwsExternalTargetParams) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *TieringAwsExternalTargetParams) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.


### SetStorageClassNil

`func (o *TieringAwsExternalTargetParams) SetStorageClassNil(b bool)`

 SetStorageClassNil sets the value for StorageClass to be an explicit nil

### UnsetStorageClass
`func (o *TieringAwsExternalTargetParams) UnsetStorageClass()`

UnsetStorageClass ensures that no value is present for StorageClass, not even an explicit nil
### GetAwsS3StandardParams

`func (o *TieringAwsExternalTargetParams) GetAwsS3StandardParams() AwsS3StandardParams`

GetAwsS3StandardParams returns the AwsS3StandardParams field if non-nil, zero value otherwise.

### GetAwsS3StandardParamsOk

`func (o *TieringAwsExternalTargetParams) GetAwsS3StandardParamsOk() (*AwsS3StandardParams, bool)`

GetAwsS3StandardParamsOk returns a tuple with the AwsS3StandardParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsS3StandardParams

`func (o *TieringAwsExternalTargetParams) SetAwsS3StandardParams(v AwsS3StandardParams)`

SetAwsS3StandardParams sets AwsS3StandardParams field to given value.

### HasAwsS3StandardParams

`func (o *TieringAwsExternalTargetParams) HasAwsS3StandardParams() bool`

HasAwsS3StandardParams returns a boolean if a field has been set.

### GetAwsS3IntelligentParams

`func (o *TieringAwsExternalTargetParams) GetAwsS3IntelligentParams() AwsS3IntelligentParams`

GetAwsS3IntelligentParams returns the AwsS3IntelligentParams field if non-nil, zero value otherwise.

### GetAwsS3IntelligentParamsOk

`func (o *TieringAwsExternalTargetParams) GetAwsS3IntelligentParamsOk() (*AwsS3IntelligentParams, bool)`

GetAwsS3IntelligentParamsOk returns a tuple with the AwsS3IntelligentParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsS3IntelligentParams

`func (o *TieringAwsExternalTargetParams) SetAwsS3IntelligentParams(v AwsS3IntelligentParams)`

SetAwsS3IntelligentParams sets AwsS3IntelligentParams field to given value.

### HasAwsS3IntelligentParams

`func (o *TieringAwsExternalTargetParams) HasAwsS3IntelligentParams() bool`

HasAwsS3IntelligentParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


