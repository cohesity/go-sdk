# ArchivalAwsExternalTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **NullableString** | Specifies bucket name of the External Target. | 
**Region** | **NullableString** | Specifies region of the External Target. | 
**StorageClass** | **NullableString** | Specifies the AWS External Target storage class. | 
**SourceSideDeduplication** | Pointer to **NullableBool** | Specifies the Source Side Deduplication setting for the AWS external target | [optional] 
**IsIncrementalArchivalEnabled** | Pointer to **NullableBool** | Specifies if Incremental Archival setting is enabled or not. | [optional] 
**IsForeverIncrementalArchivalEnabled** | Pointer to **NullableBool** | Specifies if Forever Incremental Archival setting is enabled or not. | [optional] 
**AwsS3StandardParams** | Pointer to [**AwsS3StandardParams**](AwsS3StandardParams.md) |  | [optional] 
**AwsS3StandardIAParams** | Pointer to [**AwsS3StandardIAParams**](AwsS3StandardIAParams.md) |  | [optional] 
**AwsS3OneZoneIAParams** | Pointer to [**AwsS3OneZoneIAParams**](AwsS3OneZoneIAParams.md) |  | [optional] 
**AwsS3IntelligentParams** | Pointer to [**AwsS3IntelligentParams**](AwsS3IntelligentParams.md) |  | [optional] 
**AwsS3GlacierParams** | Pointer to [**AwsS3GlacierParams**](AwsS3GlacierParams.md) |  | [optional] 
**AwsS3GlacierDeepArchiveParams** | Pointer to [**AwsS3GlacierDeepArchiveParams**](AwsS3GlacierDeepArchiveParams.md) |  | [optional] 
**AwsGlacierParams** | Pointer to [**AwsGlacierParams**](AwsGlacierParams.md) |  | [optional] 

## Methods

### NewArchivalAwsExternalTargetParams

`func NewArchivalAwsExternalTargetParams(bucketName NullableString, region NullableString, storageClass NullableString, ) *ArchivalAwsExternalTargetParams`

NewArchivalAwsExternalTargetParams instantiates a new ArchivalAwsExternalTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchivalAwsExternalTargetParamsWithDefaults

`func NewArchivalAwsExternalTargetParamsWithDefaults() *ArchivalAwsExternalTargetParams`

NewArchivalAwsExternalTargetParamsWithDefaults instantiates a new ArchivalAwsExternalTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *ArchivalAwsExternalTargetParams) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *ArchivalAwsExternalTargetParams) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *ArchivalAwsExternalTargetParams) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### SetBucketNameNil

`func (o *ArchivalAwsExternalTargetParams) SetBucketNameNil(b bool)`

 SetBucketNameNil sets the value for BucketName to be an explicit nil

### UnsetBucketName
`func (o *ArchivalAwsExternalTargetParams) UnsetBucketName()`

UnsetBucketName ensures that no value is present for BucketName, not even an explicit nil
### GetRegion

`func (o *ArchivalAwsExternalTargetParams) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ArchivalAwsExternalTargetParams) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ArchivalAwsExternalTargetParams) SetRegion(v string)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *ArchivalAwsExternalTargetParams) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *ArchivalAwsExternalTargetParams) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetStorageClass

`func (o *ArchivalAwsExternalTargetParams) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *ArchivalAwsExternalTargetParams) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *ArchivalAwsExternalTargetParams) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.


### SetStorageClassNil

`func (o *ArchivalAwsExternalTargetParams) SetStorageClassNil(b bool)`

 SetStorageClassNil sets the value for StorageClass to be an explicit nil

### UnsetStorageClass
`func (o *ArchivalAwsExternalTargetParams) UnsetStorageClass()`

UnsetStorageClass ensures that no value is present for StorageClass, not even an explicit nil
### GetSourceSideDeduplication

`func (o *ArchivalAwsExternalTargetParams) GetSourceSideDeduplication() bool`

GetSourceSideDeduplication returns the SourceSideDeduplication field if non-nil, zero value otherwise.

### GetSourceSideDeduplicationOk

`func (o *ArchivalAwsExternalTargetParams) GetSourceSideDeduplicationOk() (*bool, bool)`

GetSourceSideDeduplicationOk returns a tuple with the SourceSideDeduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSideDeduplication

`func (o *ArchivalAwsExternalTargetParams) SetSourceSideDeduplication(v bool)`

SetSourceSideDeduplication sets SourceSideDeduplication field to given value.

### HasSourceSideDeduplication

`func (o *ArchivalAwsExternalTargetParams) HasSourceSideDeduplication() bool`

HasSourceSideDeduplication returns a boolean if a field has been set.

### SetSourceSideDeduplicationNil

`func (o *ArchivalAwsExternalTargetParams) SetSourceSideDeduplicationNil(b bool)`

 SetSourceSideDeduplicationNil sets the value for SourceSideDeduplication to be an explicit nil

### UnsetSourceSideDeduplication
`func (o *ArchivalAwsExternalTargetParams) UnsetSourceSideDeduplication()`

UnsetSourceSideDeduplication ensures that no value is present for SourceSideDeduplication, not even an explicit nil
### GetIsIncrementalArchivalEnabled

`func (o *ArchivalAwsExternalTargetParams) GetIsIncrementalArchivalEnabled() bool`

GetIsIncrementalArchivalEnabled returns the IsIncrementalArchivalEnabled field if non-nil, zero value otherwise.

### GetIsIncrementalArchivalEnabledOk

`func (o *ArchivalAwsExternalTargetParams) GetIsIncrementalArchivalEnabledOk() (*bool, bool)`

GetIsIncrementalArchivalEnabledOk returns a tuple with the IsIncrementalArchivalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncrementalArchivalEnabled

`func (o *ArchivalAwsExternalTargetParams) SetIsIncrementalArchivalEnabled(v bool)`

SetIsIncrementalArchivalEnabled sets IsIncrementalArchivalEnabled field to given value.

### HasIsIncrementalArchivalEnabled

`func (o *ArchivalAwsExternalTargetParams) HasIsIncrementalArchivalEnabled() bool`

HasIsIncrementalArchivalEnabled returns a boolean if a field has been set.

### SetIsIncrementalArchivalEnabledNil

`func (o *ArchivalAwsExternalTargetParams) SetIsIncrementalArchivalEnabledNil(b bool)`

 SetIsIncrementalArchivalEnabledNil sets the value for IsIncrementalArchivalEnabled to be an explicit nil

### UnsetIsIncrementalArchivalEnabled
`func (o *ArchivalAwsExternalTargetParams) UnsetIsIncrementalArchivalEnabled()`

UnsetIsIncrementalArchivalEnabled ensures that no value is present for IsIncrementalArchivalEnabled, not even an explicit nil
### GetIsForeverIncrementalArchivalEnabled

`func (o *ArchivalAwsExternalTargetParams) GetIsForeverIncrementalArchivalEnabled() bool`

GetIsForeverIncrementalArchivalEnabled returns the IsForeverIncrementalArchivalEnabled field if non-nil, zero value otherwise.

### GetIsForeverIncrementalArchivalEnabledOk

`func (o *ArchivalAwsExternalTargetParams) GetIsForeverIncrementalArchivalEnabledOk() (*bool, bool)`

GetIsForeverIncrementalArchivalEnabledOk returns a tuple with the IsForeverIncrementalArchivalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForeverIncrementalArchivalEnabled

`func (o *ArchivalAwsExternalTargetParams) SetIsForeverIncrementalArchivalEnabled(v bool)`

SetIsForeverIncrementalArchivalEnabled sets IsForeverIncrementalArchivalEnabled field to given value.

### HasIsForeverIncrementalArchivalEnabled

`func (o *ArchivalAwsExternalTargetParams) HasIsForeverIncrementalArchivalEnabled() bool`

HasIsForeverIncrementalArchivalEnabled returns a boolean if a field has been set.

### SetIsForeverIncrementalArchivalEnabledNil

`func (o *ArchivalAwsExternalTargetParams) SetIsForeverIncrementalArchivalEnabledNil(b bool)`

 SetIsForeverIncrementalArchivalEnabledNil sets the value for IsForeverIncrementalArchivalEnabled to be an explicit nil

### UnsetIsForeverIncrementalArchivalEnabled
`func (o *ArchivalAwsExternalTargetParams) UnsetIsForeverIncrementalArchivalEnabled()`

UnsetIsForeverIncrementalArchivalEnabled ensures that no value is present for IsForeverIncrementalArchivalEnabled, not even an explicit nil
### GetAwsS3StandardParams

`func (o *ArchivalAwsExternalTargetParams) GetAwsS3StandardParams() AwsS3StandardParams`

GetAwsS3StandardParams returns the AwsS3StandardParams field if non-nil, zero value otherwise.

### GetAwsS3StandardParamsOk

`func (o *ArchivalAwsExternalTargetParams) GetAwsS3StandardParamsOk() (*AwsS3StandardParams, bool)`

GetAwsS3StandardParamsOk returns a tuple with the AwsS3StandardParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsS3StandardParams

`func (o *ArchivalAwsExternalTargetParams) SetAwsS3StandardParams(v AwsS3StandardParams)`

SetAwsS3StandardParams sets AwsS3StandardParams field to given value.

### HasAwsS3StandardParams

`func (o *ArchivalAwsExternalTargetParams) HasAwsS3StandardParams() bool`

HasAwsS3StandardParams returns a boolean if a field has been set.

### GetAwsS3StandardIAParams

`func (o *ArchivalAwsExternalTargetParams) GetAwsS3StandardIAParams() AwsS3StandardIAParams`

GetAwsS3StandardIAParams returns the AwsS3StandardIAParams field if non-nil, zero value otherwise.

### GetAwsS3StandardIAParamsOk

`func (o *ArchivalAwsExternalTargetParams) GetAwsS3StandardIAParamsOk() (*AwsS3StandardIAParams, bool)`

GetAwsS3StandardIAParamsOk returns a tuple with the AwsS3StandardIAParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsS3StandardIAParams

`func (o *ArchivalAwsExternalTargetParams) SetAwsS3StandardIAParams(v AwsS3StandardIAParams)`

SetAwsS3StandardIAParams sets AwsS3StandardIAParams field to given value.

### HasAwsS3StandardIAParams

`func (o *ArchivalAwsExternalTargetParams) HasAwsS3StandardIAParams() bool`

HasAwsS3StandardIAParams returns a boolean if a field has been set.

### GetAwsS3OneZoneIAParams

`func (o *ArchivalAwsExternalTargetParams) GetAwsS3OneZoneIAParams() AwsS3OneZoneIAParams`

GetAwsS3OneZoneIAParams returns the AwsS3OneZoneIAParams field if non-nil, zero value otherwise.

### GetAwsS3OneZoneIAParamsOk

`func (o *ArchivalAwsExternalTargetParams) GetAwsS3OneZoneIAParamsOk() (*AwsS3OneZoneIAParams, bool)`

GetAwsS3OneZoneIAParamsOk returns a tuple with the AwsS3OneZoneIAParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsS3OneZoneIAParams

`func (o *ArchivalAwsExternalTargetParams) SetAwsS3OneZoneIAParams(v AwsS3OneZoneIAParams)`

SetAwsS3OneZoneIAParams sets AwsS3OneZoneIAParams field to given value.

### HasAwsS3OneZoneIAParams

`func (o *ArchivalAwsExternalTargetParams) HasAwsS3OneZoneIAParams() bool`

HasAwsS3OneZoneIAParams returns a boolean if a field has been set.

### GetAwsS3IntelligentParams

`func (o *ArchivalAwsExternalTargetParams) GetAwsS3IntelligentParams() AwsS3IntelligentParams`

GetAwsS3IntelligentParams returns the AwsS3IntelligentParams field if non-nil, zero value otherwise.

### GetAwsS3IntelligentParamsOk

`func (o *ArchivalAwsExternalTargetParams) GetAwsS3IntelligentParamsOk() (*AwsS3IntelligentParams, bool)`

GetAwsS3IntelligentParamsOk returns a tuple with the AwsS3IntelligentParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsS3IntelligentParams

`func (o *ArchivalAwsExternalTargetParams) SetAwsS3IntelligentParams(v AwsS3IntelligentParams)`

SetAwsS3IntelligentParams sets AwsS3IntelligentParams field to given value.

### HasAwsS3IntelligentParams

`func (o *ArchivalAwsExternalTargetParams) HasAwsS3IntelligentParams() bool`

HasAwsS3IntelligentParams returns a boolean if a field has been set.

### GetAwsS3GlacierParams

`func (o *ArchivalAwsExternalTargetParams) GetAwsS3GlacierParams() AwsS3GlacierParams`

GetAwsS3GlacierParams returns the AwsS3GlacierParams field if non-nil, zero value otherwise.

### GetAwsS3GlacierParamsOk

`func (o *ArchivalAwsExternalTargetParams) GetAwsS3GlacierParamsOk() (*AwsS3GlacierParams, bool)`

GetAwsS3GlacierParamsOk returns a tuple with the AwsS3GlacierParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsS3GlacierParams

`func (o *ArchivalAwsExternalTargetParams) SetAwsS3GlacierParams(v AwsS3GlacierParams)`

SetAwsS3GlacierParams sets AwsS3GlacierParams field to given value.

### HasAwsS3GlacierParams

`func (o *ArchivalAwsExternalTargetParams) HasAwsS3GlacierParams() bool`

HasAwsS3GlacierParams returns a boolean if a field has been set.

### GetAwsS3GlacierDeepArchiveParams

`func (o *ArchivalAwsExternalTargetParams) GetAwsS3GlacierDeepArchiveParams() AwsS3GlacierDeepArchiveParams`

GetAwsS3GlacierDeepArchiveParams returns the AwsS3GlacierDeepArchiveParams field if non-nil, zero value otherwise.

### GetAwsS3GlacierDeepArchiveParamsOk

`func (o *ArchivalAwsExternalTargetParams) GetAwsS3GlacierDeepArchiveParamsOk() (*AwsS3GlacierDeepArchiveParams, bool)`

GetAwsS3GlacierDeepArchiveParamsOk returns a tuple with the AwsS3GlacierDeepArchiveParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsS3GlacierDeepArchiveParams

`func (o *ArchivalAwsExternalTargetParams) SetAwsS3GlacierDeepArchiveParams(v AwsS3GlacierDeepArchiveParams)`

SetAwsS3GlacierDeepArchiveParams sets AwsS3GlacierDeepArchiveParams field to given value.

### HasAwsS3GlacierDeepArchiveParams

`func (o *ArchivalAwsExternalTargetParams) HasAwsS3GlacierDeepArchiveParams() bool`

HasAwsS3GlacierDeepArchiveParams returns a boolean if a field has been set.

### GetAwsGlacierParams

`func (o *ArchivalAwsExternalTargetParams) GetAwsGlacierParams() AwsGlacierParams`

GetAwsGlacierParams returns the AwsGlacierParams field if non-nil, zero value otherwise.

### GetAwsGlacierParamsOk

`func (o *ArchivalAwsExternalTargetParams) GetAwsGlacierParamsOk() (*AwsGlacierParams, bool)`

GetAwsGlacierParamsOk returns a tuple with the AwsGlacierParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsGlacierParams

`func (o *ArchivalAwsExternalTargetParams) SetAwsGlacierParams(v AwsGlacierParams)`

SetAwsGlacierParams sets AwsGlacierParams field to given value.

### HasAwsGlacierParams

`func (o *ArchivalAwsExternalTargetParams) HasAwsGlacierParams() bool`

HasAwsGlacierParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


