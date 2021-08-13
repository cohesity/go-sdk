# ArchivalGcpExternalTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **NullableString** | Specifies the bucket name of the external target. | 
**ProjectId** | **NullableString** | Specifies the project Id of the external target. | 
**ClientEmailAddress** | **NullableString** | Specifies the client email address of the external target. | 
**ClientPrivateKey** | **NullableString** | Specifies the client private key of the external target. | 
**StorageClass** | **NullableString** | Specifies the GCP External Target storage class. | 
**SourceSideDeduplication** | Pointer to **NullableBool** | Specifies the Source Side Deduplication setting for the GCP external target | [optional] 
**IsIncrementalArchivalEnabled** | Pointer to **NullableBool** | Specifies if Incremental Archival setting is enabled or not. | [optional] 
**IsForeverIncrementalArchivalEnabled** | Pointer to **NullableBool** | Specifies if Forever Incremental Archival setting is enabled or not. | [optional] 

## Methods

### NewArchivalGcpExternalTargetParams

`func NewArchivalGcpExternalTargetParams(bucketName NullableString, projectId NullableString, clientEmailAddress NullableString, clientPrivateKey NullableString, storageClass NullableString, ) *ArchivalGcpExternalTargetParams`

NewArchivalGcpExternalTargetParams instantiates a new ArchivalGcpExternalTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchivalGcpExternalTargetParamsWithDefaults

`func NewArchivalGcpExternalTargetParamsWithDefaults() *ArchivalGcpExternalTargetParams`

NewArchivalGcpExternalTargetParamsWithDefaults instantiates a new ArchivalGcpExternalTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *ArchivalGcpExternalTargetParams) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *ArchivalGcpExternalTargetParams) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *ArchivalGcpExternalTargetParams) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### SetBucketNameNil

`func (o *ArchivalGcpExternalTargetParams) SetBucketNameNil(b bool)`

 SetBucketNameNil sets the value for BucketName to be an explicit nil

### UnsetBucketName
`func (o *ArchivalGcpExternalTargetParams) UnsetBucketName()`

UnsetBucketName ensures that no value is present for BucketName, not even an explicit nil
### GetProjectId

`func (o *ArchivalGcpExternalTargetParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ArchivalGcpExternalTargetParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ArchivalGcpExternalTargetParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### SetProjectIdNil

`func (o *ArchivalGcpExternalTargetParams) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *ArchivalGcpExternalTargetParams) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetClientEmailAddress

`func (o *ArchivalGcpExternalTargetParams) GetClientEmailAddress() string`

GetClientEmailAddress returns the ClientEmailAddress field if non-nil, zero value otherwise.

### GetClientEmailAddressOk

`func (o *ArchivalGcpExternalTargetParams) GetClientEmailAddressOk() (*string, bool)`

GetClientEmailAddressOk returns a tuple with the ClientEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEmailAddress

`func (o *ArchivalGcpExternalTargetParams) SetClientEmailAddress(v string)`

SetClientEmailAddress sets ClientEmailAddress field to given value.


### SetClientEmailAddressNil

`func (o *ArchivalGcpExternalTargetParams) SetClientEmailAddressNil(b bool)`

 SetClientEmailAddressNil sets the value for ClientEmailAddress to be an explicit nil

### UnsetClientEmailAddress
`func (o *ArchivalGcpExternalTargetParams) UnsetClientEmailAddress()`

UnsetClientEmailAddress ensures that no value is present for ClientEmailAddress, not even an explicit nil
### GetClientPrivateKey

`func (o *ArchivalGcpExternalTargetParams) GetClientPrivateKey() string`

GetClientPrivateKey returns the ClientPrivateKey field if non-nil, zero value otherwise.

### GetClientPrivateKeyOk

`func (o *ArchivalGcpExternalTargetParams) GetClientPrivateKeyOk() (*string, bool)`

GetClientPrivateKeyOk returns a tuple with the ClientPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPrivateKey

`func (o *ArchivalGcpExternalTargetParams) SetClientPrivateKey(v string)`

SetClientPrivateKey sets ClientPrivateKey field to given value.


### SetClientPrivateKeyNil

`func (o *ArchivalGcpExternalTargetParams) SetClientPrivateKeyNil(b bool)`

 SetClientPrivateKeyNil sets the value for ClientPrivateKey to be an explicit nil

### UnsetClientPrivateKey
`func (o *ArchivalGcpExternalTargetParams) UnsetClientPrivateKey()`

UnsetClientPrivateKey ensures that no value is present for ClientPrivateKey, not even an explicit nil
### GetStorageClass

`func (o *ArchivalGcpExternalTargetParams) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *ArchivalGcpExternalTargetParams) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *ArchivalGcpExternalTargetParams) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.


### SetStorageClassNil

`func (o *ArchivalGcpExternalTargetParams) SetStorageClassNil(b bool)`

 SetStorageClassNil sets the value for StorageClass to be an explicit nil

### UnsetStorageClass
`func (o *ArchivalGcpExternalTargetParams) UnsetStorageClass()`

UnsetStorageClass ensures that no value is present for StorageClass, not even an explicit nil
### GetSourceSideDeduplication

`func (o *ArchivalGcpExternalTargetParams) GetSourceSideDeduplication() bool`

GetSourceSideDeduplication returns the SourceSideDeduplication field if non-nil, zero value otherwise.

### GetSourceSideDeduplicationOk

`func (o *ArchivalGcpExternalTargetParams) GetSourceSideDeduplicationOk() (*bool, bool)`

GetSourceSideDeduplicationOk returns a tuple with the SourceSideDeduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSideDeduplication

`func (o *ArchivalGcpExternalTargetParams) SetSourceSideDeduplication(v bool)`

SetSourceSideDeduplication sets SourceSideDeduplication field to given value.

### HasSourceSideDeduplication

`func (o *ArchivalGcpExternalTargetParams) HasSourceSideDeduplication() bool`

HasSourceSideDeduplication returns a boolean if a field has been set.

### SetSourceSideDeduplicationNil

`func (o *ArchivalGcpExternalTargetParams) SetSourceSideDeduplicationNil(b bool)`

 SetSourceSideDeduplicationNil sets the value for SourceSideDeduplication to be an explicit nil

### UnsetSourceSideDeduplication
`func (o *ArchivalGcpExternalTargetParams) UnsetSourceSideDeduplication()`

UnsetSourceSideDeduplication ensures that no value is present for SourceSideDeduplication, not even an explicit nil
### GetIsIncrementalArchivalEnabled

`func (o *ArchivalGcpExternalTargetParams) GetIsIncrementalArchivalEnabled() bool`

GetIsIncrementalArchivalEnabled returns the IsIncrementalArchivalEnabled field if non-nil, zero value otherwise.

### GetIsIncrementalArchivalEnabledOk

`func (o *ArchivalGcpExternalTargetParams) GetIsIncrementalArchivalEnabledOk() (*bool, bool)`

GetIsIncrementalArchivalEnabledOk returns a tuple with the IsIncrementalArchivalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncrementalArchivalEnabled

`func (o *ArchivalGcpExternalTargetParams) SetIsIncrementalArchivalEnabled(v bool)`

SetIsIncrementalArchivalEnabled sets IsIncrementalArchivalEnabled field to given value.

### HasIsIncrementalArchivalEnabled

`func (o *ArchivalGcpExternalTargetParams) HasIsIncrementalArchivalEnabled() bool`

HasIsIncrementalArchivalEnabled returns a boolean if a field has been set.

### SetIsIncrementalArchivalEnabledNil

`func (o *ArchivalGcpExternalTargetParams) SetIsIncrementalArchivalEnabledNil(b bool)`

 SetIsIncrementalArchivalEnabledNil sets the value for IsIncrementalArchivalEnabled to be an explicit nil

### UnsetIsIncrementalArchivalEnabled
`func (o *ArchivalGcpExternalTargetParams) UnsetIsIncrementalArchivalEnabled()`

UnsetIsIncrementalArchivalEnabled ensures that no value is present for IsIncrementalArchivalEnabled, not even an explicit nil
### GetIsForeverIncrementalArchivalEnabled

`func (o *ArchivalGcpExternalTargetParams) GetIsForeverIncrementalArchivalEnabled() bool`

GetIsForeverIncrementalArchivalEnabled returns the IsForeverIncrementalArchivalEnabled field if non-nil, zero value otherwise.

### GetIsForeverIncrementalArchivalEnabledOk

`func (o *ArchivalGcpExternalTargetParams) GetIsForeverIncrementalArchivalEnabledOk() (*bool, bool)`

GetIsForeverIncrementalArchivalEnabledOk returns a tuple with the IsForeverIncrementalArchivalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForeverIncrementalArchivalEnabled

`func (o *ArchivalGcpExternalTargetParams) SetIsForeverIncrementalArchivalEnabled(v bool)`

SetIsForeverIncrementalArchivalEnabled sets IsForeverIncrementalArchivalEnabled field to given value.

### HasIsForeverIncrementalArchivalEnabled

`func (o *ArchivalGcpExternalTargetParams) HasIsForeverIncrementalArchivalEnabled() bool`

HasIsForeverIncrementalArchivalEnabled returns a boolean if a field has been set.

### SetIsForeverIncrementalArchivalEnabledNil

`func (o *ArchivalGcpExternalTargetParams) SetIsForeverIncrementalArchivalEnabledNil(b bool)`

 SetIsForeverIncrementalArchivalEnabledNil sets the value for IsForeverIncrementalArchivalEnabled to be an explicit nil

### UnsetIsForeverIncrementalArchivalEnabled
`func (o *ArchivalGcpExternalTargetParams) UnsetIsForeverIncrementalArchivalEnabled()`

UnsetIsForeverIncrementalArchivalEnabled ensures that no value is present for IsForeverIncrementalArchivalEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


