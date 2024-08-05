# CommonArchivalAwsExternalTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **NullableString** | Specifies bucket name of the External Target. | 
**Region** | **NullableString** | Specifies region of the External Target. | 
**BucketOwnerAccountId** | Pointer to **NullableString** | Specifies the account Id of the S3 bucket owner. | [optional] 
**IsForeverIncrementalArchivalEnabled** | Pointer to **NullableBool** | Specifies if Forever Incremental Archival setting is enabled or not. | [optional] 
**IsIncrementalArchivalEnabled** | Pointer to **NullableBool** | Specifies if Incremental Archival setting is enabled or not. | [optional] 
**SourceSideDeduplication** | Pointer to **NullableBool** | Specifies the Source Side Deduplication setting for the AWS external target | [optional] 
**StorageClass** | **NullableString** | Specifies the AWS External Target storage class. | 

## Methods

### NewCommonArchivalAwsExternalTargetParams

`func NewCommonArchivalAwsExternalTargetParams(bucketName NullableString, region NullableString, storageClass NullableString, ) *CommonArchivalAwsExternalTargetParams`

NewCommonArchivalAwsExternalTargetParams instantiates a new CommonArchivalAwsExternalTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonArchivalAwsExternalTargetParamsWithDefaults

`func NewCommonArchivalAwsExternalTargetParamsWithDefaults() *CommonArchivalAwsExternalTargetParams`

NewCommonArchivalAwsExternalTargetParamsWithDefaults instantiates a new CommonArchivalAwsExternalTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *CommonArchivalAwsExternalTargetParams) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *CommonArchivalAwsExternalTargetParams) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *CommonArchivalAwsExternalTargetParams) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### SetBucketNameNil

`func (o *CommonArchivalAwsExternalTargetParams) SetBucketNameNil(b bool)`

 SetBucketNameNil sets the value for BucketName to be an explicit nil

### UnsetBucketName
`func (o *CommonArchivalAwsExternalTargetParams) UnsetBucketName()`

UnsetBucketName ensures that no value is present for BucketName, not even an explicit nil
### GetRegion

`func (o *CommonArchivalAwsExternalTargetParams) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CommonArchivalAwsExternalTargetParams) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CommonArchivalAwsExternalTargetParams) SetRegion(v string)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *CommonArchivalAwsExternalTargetParams) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *CommonArchivalAwsExternalTargetParams) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetBucketOwnerAccountId

`func (o *CommonArchivalAwsExternalTargetParams) GetBucketOwnerAccountId() string`

GetBucketOwnerAccountId returns the BucketOwnerAccountId field if non-nil, zero value otherwise.

### GetBucketOwnerAccountIdOk

`func (o *CommonArchivalAwsExternalTargetParams) GetBucketOwnerAccountIdOk() (*string, bool)`

GetBucketOwnerAccountIdOk returns a tuple with the BucketOwnerAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketOwnerAccountId

`func (o *CommonArchivalAwsExternalTargetParams) SetBucketOwnerAccountId(v string)`

SetBucketOwnerAccountId sets BucketOwnerAccountId field to given value.

### HasBucketOwnerAccountId

`func (o *CommonArchivalAwsExternalTargetParams) HasBucketOwnerAccountId() bool`

HasBucketOwnerAccountId returns a boolean if a field has been set.

### SetBucketOwnerAccountIdNil

`func (o *CommonArchivalAwsExternalTargetParams) SetBucketOwnerAccountIdNil(b bool)`

 SetBucketOwnerAccountIdNil sets the value for BucketOwnerAccountId to be an explicit nil

### UnsetBucketOwnerAccountId
`func (o *CommonArchivalAwsExternalTargetParams) UnsetBucketOwnerAccountId()`

UnsetBucketOwnerAccountId ensures that no value is present for BucketOwnerAccountId, not even an explicit nil
### GetIsForeverIncrementalArchivalEnabled

`func (o *CommonArchivalAwsExternalTargetParams) GetIsForeverIncrementalArchivalEnabled() bool`

GetIsForeverIncrementalArchivalEnabled returns the IsForeverIncrementalArchivalEnabled field if non-nil, zero value otherwise.

### GetIsForeverIncrementalArchivalEnabledOk

`func (o *CommonArchivalAwsExternalTargetParams) GetIsForeverIncrementalArchivalEnabledOk() (*bool, bool)`

GetIsForeverIncrementalArchivalEnabledOk returns a tuple with the IsForeverIncrementalArchivalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForeverIncrementalArchivalEnabled

`func (o *CommonArchivalAwsExternalTargetParams) SetIsForeverIncrementalArchivalEnabled(v bool)`

SetIsForeverIncrementalArchivalEnabled sets IsForeverIncrementalArchivalEnabled field to given value.

### HasIsForeverIncrementalArchivalEnabled

`func (o *CommonArchivalAwsExternalTargetParams) HasIsForeverIncrementalArchivalEnabled() bool`

HasIsForeverIncrementalArchivalEnabled returns a boolean if a field has been set.

### SetIsForeverIncrementalArchivalEnabledNil

`func (o *CommonArchivalAwsExternalTargetParams) SetIsForeverIncrementalArchivalEnabledNil(b bool)`

 SetIsForeverIncrementalArchivalEnabledNil sets the value for IsForeverIncrementalArchivalEnabled to be an explicit nil

### UnsetIsForeverIncrementalArchivalEnabled
`func (o *CommonArchivalAwsExternalTargetParams) UnsetIsForeverIncrementalArchivalEnabled()`

UnsetIsForeverIncrementalArchivalEnabled ensures that no value is present for IsForeverIncrementalArchivalEnabled, not even an explicit nil
### GetIsIncrementalArchivalEnabled

`func (o *CommonArchivalAwsExternalTargetParams) GetIsIncrementalArchivalEnabled() bool`

GetIsIncrementalArchivalEnabled returns the IsIncrementalArchivalEnabled field if non-nil, zero value otherwise.

### GetIsIncrementalArchivalEnabledOk

`func (o *CommonArchivalAwsExternalTargetParams) GetIsIncrementalArchivalEnabledOk() (*bool, bool)`

GetIsIncrementalArchivalEnabledOk returns a tuple with the IsIncrementalArchivalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncrementalArchivalEnabled

`func (o *CommonArchivalAwsExternalTargetParams) SetIsIncrementalArchivalEnabled(v bool)`

SetIsIncrementalArchivalEnabled sets IsIncrementalArchivalEnabled field to given value.

### HasIsIncrementalArchivalEnabled

`func (o *CommonArchivalAwsExternalTargetParams) HasIsIncrementalArchivalEnabled() bool`

HasIsIncrementalArchivalEnabled returns a boolean if a field has been set.

### SetIsIncrementalArchivalEnabledNil

`func (o *CommonArchivalAwsExternalTargetParams) SetIsIncrementalArchivalEnabledNil(b bool)`

 SetIsIncrementalArchivalEnabledNil sets the value for IsIncrementalArchivalEnabled to be an explicit nil

### UnsetIsIncrementalArchivalEnabled
`func (o *CommonArchivalAwsExternalTargetParams) UnsetIsIncrementalArchivalEnabled()`

UnsetIsIncrementalArchivalEnabled ensures that no value is present for IsIncrementalArchivalEnabled, not even an explicit nil
### GetSourceSideDeduplication

`func (o *CommonArchivalAwsExternalTargetParams) GetSourceSideDeduplication() bool`

GetSourceSideDeduplication returns the SourceSideDeduplication field if non-nil, zero value otherwise.

### GetSourceSideDeduplicationOk

`func (o *CommonArchivalAwsExternalTargetParams) GetSourceSideDeduplicationOk() (*bool, bool)`

GetSourceSideDeduplicationOk returns a tuple with the SourceSideDeduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSideDeduplication

`func (o *CommonArchivalAwsExternalTargetParams) SetSourceSideDeduplication(v bool)`

SetSourceSideDeduplication sets SourceSideDeduplication field to given value.

### HasSourceSideDeduplication

`func (o *CommonArchivalAwsExternalTargetParams) HasSourceSideDeduplication() bool`

HasSourceSideDeduplication returns a boolean if a field has been set.

### SetSourceSideDeduplicationNil

`func (o *CommonArchivalAwsExternalTargetParams) SetSourceSideDeduplicationNil(b bool)`

 SetSourceSideDeduplicationNil sets the value for SourceSideDeduplication to be an explicit nil

### UnsetSourceSideDeduplication
`func (o *CommonArchivalAwsExternalTargetParams) UnsetSourceSideDeduplication()`

UnsetSourceSideDeduplication ensures that no value is present for SourceSideDeduplication, not even an explicit nil
### GetStorageClass

`func (o *CommonArchivalAwsExternalTargetParams) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *CommonArchivalAwsExternalTargetParams) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *CommonArchivalAwsExternalTargetParams) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.


### SetStorageClassNil

`func (o *CommonArchivalAwsExternalTargetParams) SetStorageClassNil(b bool)`

 SetStorageClassNil sets the value for StorageClass to be an explicit nil

### UnsetStorageClass
`func (o *CommonArchivalAwsExternalTargetParams) UnsetStorageClass()`

UnsetStorageClass ensures that no value is present for StorageClass, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


