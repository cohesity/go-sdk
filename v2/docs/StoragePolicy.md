# StoragePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AesEncryptionMode** | Pointer to **NullableString** | Specifies the encryption mode for a Storage Domain. | [optional] 
**AppMarkerDetectionEnabled** | Pointer to **NullableBool** | Specifies whether app marker detection is enabled. When enabled, app markers will be removed from data and put in separate chunks. | [optional] 
**CloudSpillVaultId** | Pointer to **NullableInt64** | Specifies the vault id assigned for cloud spill for a Storage Domain. | [optional] 
**CompressionParams** | Pointer to [**StoragePolicyCompressionParams**](StoragePolicyCompressionParams.md) |  | [optional] 
**DeduplicationCompressionDelaySecs** | Pointer to **NullableInt32** | Specifies the time in seconds when deduplication and compression of the Storage Domain starts. | [optional] 
**DeduplicationParams** | Pointer to [**StoragePolicyDeduplicationParams**](StoragePolicyDeduplicationParams.md) |  | [optional] 
**EncryptionType** | Pointer to **NullableString** | Specifies the encryption type for a Storage Domain. | [optional] 
**ErasureCodingParams** | Pointer to [**StoragePolicyErasureCodingParams**](StoragePolicyErasureCodingParams.md) |  | [optional] 
**NumDiskFailuresTolerated** | Pointer to **NullableInt32** | Specifies the number of disk failures to tolerate for a Storage Domain. By default, this field is 1 for cluster with three or more nodes. If erasure coding is enabled, this field will be the same as numCodedStripes. | [optional] 
**NumNodeFailuresTolerated** | Pointer to **NullableInt32** | Specifies the number of node failures to tolerate for a Storage Domain. By default this field is replication factor minus 1 for replication chunk files and is the same as numCodedStripes for erasure coding chunk files. | [optional] 

## Methods

### NewStoragePolicy

`func NewStoragePolicy() *StoragePolicy`

NewStoragePolicy instantiates a new StoragePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePolicyWithDefaults

`func NewStoragePolicyWithDefaults() *StoragePolicy`

NewStoragePolicyWithDefaults instantiates a new StoragePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAesEncryptionMode

`func (o *StoragePolicy) GetAesEncryptionMode() string`

GetAesEncryptionMode returns the AesEncryptionMode field if non-nil, zero value otherwise.

### GetAesEncryptionModeOk

`func (o *StoragePolicy) GetAesEncryptionModeOk() (*string, bool)`

GetAesEncryptionModeOk returns a tuple with the AesEncryptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAesEncryptionMode

`func (o *StoragePolicy) SetAesEncryptionMode(v string)`

SetAesEncryptionMode sets AesEncryptionMode field to given value.

### HasAesEncryptionMode

`func (o *StoragePolicy) HasAesEncryptionMode() bool`

HasAesEncryptionMode returns a boolean if a field has been set.

### SetAesEncryptionModeNil

`func (o *StoragePolicy) SetAesEncryptionModeNil(b bool)`

 SetAesEncryptionModeNil sets the value for AesEncryptionMode to be an explicit nil

### UnsetAesEncryptionMode
`func (o *StoragePolicy) UnsetAesEncryptionMode()`

UnsetAesEncryptionMode ensures that no value is present for AesEncryptionMode, not even an explicit nil
### GetAppMarkerDetectionEnabled

`func (o *StoragePolicy) GetAppMarkerDetectionEnabled() bool`

GetAppMarkerDetectionEnabled returns the AppMarkerDetectionEnabled field if non-nil, zero value otherwise.

### GetAppMarkerDetectionEnabledOk

`func (o *StoragePolicy) GetAppMarkerDetectionEnabledOk() (*bool, bool)`

GetAppMarkerDetectionEnabledOk returns a tuple with the AppMarkerDetectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppMarkerDetectionEnabled

`func (o *StoragePolicy) SetAppMarkerDetectionEnabled(v bool)`

SetAppMarkerDetectionEnabled sets AppMarkerDetectionEnabled field to given value.

### HasAppMarkerDetectionEnabled

`func (o *StoragePolicy) HasAppMarkerDetectionEnabled() bool`

HasAppMarkerDetectionEnabled returns a boolean if a field has been set.

### SetAppMarkerDetectionEnabledNil

`func (o *StoragePolicy) SetAppMarkerDetectionEnabledNil(b bool)`

 SetAppMarkerDetectionEnabledNil sets the value for AppMarkerDetectionEnabled to be an explicit nil

### UnsetAppMarkerDetectionEnabled
`func (o *StoragePolicy) UnsetAppMarkerDetectionEnabled()`

UnsetAppMarkerDetectionEnabled ensures that no value is present for AppMarkerDetectionEnabled, not even an explicit nil
### GetCloudSpillVaultId

`func (o *StoragePolicy) GetCloudSpillVaultId() int64`

GetCloudSpillVaultId returns the CloudSpillVaultId field if non-nil, zero value otherwise.

### GetCloudSpillVaultIdOk

`func (o *StoragePolicy) GetCloudSpillVaultIdOk() (*int64, bool)`

GetCloudSpillVaultIdOk returns a tuple with the CloudSpillVaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSpillVaultId

`func (o *StoragePolicy) SetCloudSpillVaultId(v int64)`

SetCloudSpillVaultId sets CloudSpillVaultId field to given value.

### HasCloudSpillVaultId

`func (o *StoragePolicy) HasCloudSpillVaultId() bool`

HasCloudSpillVaultId returns a boolean if a field has been set.

### SetCloudSpillVaultIdNil

`func (o *StoragePolicy) SetCloudSpillVaultIdNil(b bool)`

 SetCloudSpillVaultIdNil sets the value for CloudSpillVaultId to be an explicit nil

### UnsetCloudSpillVaultId
`func (o *StoragePolicy) UnsetCloudSpillVaultId()`

UnsetCloudSpillVaultId ensures that no value is present for CloudSpillVaultId, not even an explicit nil
### GetCompressionParams

`func (o *StoragePolicy) GetCompressionParams() StoragePolicyCompressionParams`

GetCompressionParams returns the CompressionParams field if non-nil, zero value otherwise.

### GetCompressionParamsOk

`func (o *StoragePolicy) GetCompressionParamsOk() (*StoragePolicyCompressionParams, bool)`

GetCompressionParamsOk returns a tuple with the CompressionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionParams

`func (o *StoragePolicy) SetCompressionParams(v StoragePolicyCompressionParams)`

SetCompressionParams sets CompressionParams field to given value.

### HasCompressionParams

`func (o *StoragePolicy) HasCompressionParams() bool`

HasCompressionParams returns a boolean if a field has been set.

### GetDeduplicationCompressionDelaySecs

`func (o *StoragePolicy) GetDeduplicationCompressionDelaySecs() int32`

GetDeduplicationCompressionDelaySecs returns the DeduplicationCompressionDelaySecs field if non-nil, zero value otherwise.

### GetDeduplicationCompressionDelaySecsOk

`func (o *StoragePolicy) GetDeduplicationCompressionDelaySecsOk() (*int32, bool)`

GetDeduplicationCompressionDelaySecsOk returns a tuple with the DeduplicationCompressionDelaySecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplicationCompressionDelaySecs

`func (o *StoragePolicy) SetDeduplicationCompressionDelaySecs(v int32)`

SetDeduplicationCompressionDelaySecs sets DeduplicationCompressionDelaySecs field to given value.

### HasDeduplicationCompressionDelaySecs

`func (o *StoragePolicy) HasDeduplicationCompressionDelaySecs() bool`

HasDeduplicationCompressionDelaySecs returns a boolean if a field has been set.

### SetDeduplicationCompressionDelaySecsNil

`func (o *StoragePolicy) SetDeduplicationCompressionDelaySecsNil(b bool)`

 SetDeduplicationCompressionDelaySecsNil sets the value for DeduplicationCompressionDelaySecs to be an explicit nil

### UnsetDeduplicationCompressionDelaySecs
`func (o *StoragePolicy) UnsetDeduplicationCompressionDelaySecs()`

UnsetDeduplicationCompressionDelaySecs ensures that no value is present for DeduplicationCompressionDelaySecs, not even an explicit nil
### GetDeduplicationParams

`func (o *StoragePolicy) GetDeduplicationParams() StoragePolicyDeduplicationParams`

GetDeduplicationParams returns the DeduplicationParams field if non-nil, zero value otherwise.

### GetDeduplicationParamsOk

`func (o *StoragePolicy) GetDeduplicationParamsOk() (*StoragePolicyDeduplicationParams, bool)`

GetDeduplicationParamsOk returns a tuple with the DeduplicationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplicationParams

`func (o *StoragePolicy) SetDeduplicationParams(v StoragePolicyDeduplicationParams)`

SetDeduplicationParams sets DeduplicationParams field to given value.

### HasDeduplicationParams

`func (o *StoragePolicy) HasDeduplicationParams() bool`

HasDeduplicationParams returns a boolean if a field has been set.

### GetEncryptionType

`func (o *StoragePolicy) GetEncryptionType() string`

GetEncryptionType returns the EncryptionType field if non-nil, zero value otherwise.

### GetEncryptionTypeOk

`func (o *StoragePolicy) GetEncryptionTypeOk() (*string, bool)`

GetEncryptionTypeOk returns a tuple with the EncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionType

`func (o *StoragePolicy) SetEncryptionType(v string)`

SetEncryptionType sets EncryptionType field to given value.

### HasEncryptionType

`func (o *StoragePolicy) HasEncryptionType() bool`

HasEncryptionType returns a boolean if a field has been set.

### SetEncryptionTypeNil

`func (o *StoragePolicy) SetEncryptionTypeNil(b bool)`

 SetEncryptionTypeNil sets the value for EncryptionType to be an explicit nil

### UnsetEncryptionType
`func (o *StoragePolicy) UnsetEncryptionType()`

UnsetEncryptionType ensures that no value is present for EncryptionType, not even an explicit nil
### GetErasureCodingParams

`func (o *StoragePolicy) GetErasureCodingParams() StoragePolicyErasureCodingParams`

GetErasureCodingParams returns the ErasureCodingParams field if non-nil, zero value otherwise.

### GetErasureCodingParamsOk

`func (o *StoragePolicy) GetErasureCodingParamsOk() (*StoragePolicyErasureCodingParams, bool)`

GetErasureCodingParamsOk returns a tuple with the ErasureCodingParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErasureCodingParams

`func (o *StoragePolicy) SetErasureCodingParams(v StoragePolicyErasureCodingParams)`

SetErasureCodingParams sets ErasureCodingParams field to given value.

### HasErasureCodingParams

`func (o *StoragePolicy) HasErasureCodingParams() bool`

HasErasureCodingParams returns a boolean if a field has been set.

### GetNumDiskFailuresTolerated

`func (o *StoragePolicy) GetNumDiskFailuresTolerated() int32`

GetNumDiskFailuresTolerated returns the NumDiskFailuresTolerated field if non-nil, zero value otherwise.

### GetNumDiskFailuresToleratedOk

`func (o *StoragePolicy) GetNumDiskFailuresToleratedOk() (*int32, bool)`

GetNumDiskFailuresToleratedOk returns a tuple with the NumDiskFailuresTolerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDiskFailuresTolerated

`func (o *StoragePolicy) SetNumDiskFailuresTolerated(v int32)`

SetNumDiskFailuresTolerated sets NumDiskFailuresTolerated field to given value.

### HasNumDiskFailuresTolerated

`func (o *StoragePolicy) HasNumDiskFailuresTolerated() bool`

HasNumDiskFailuresTolerated returns a boolean if a field has been set.

### SetNumDiskFailuresToleratedNil

`func (o *StoragePolicy) SetNumDiskFailuresToleratedNil(b bool)`

 SetNumDiskFailuresToleratedNil sets the value for NumDiskFailuresTolerated to be an explicit nil

### UnsetNumDiskFailuresTolerated
`func (o *StoragePolicy) UnsetNumDiskFailuresTolerated()`

UnsetNumDiskFailuresTolerated ensures that no value is present for NumDiskFailuresTolerated, not even an explicit nil
### GetNumNodeFailuresTolerated

`func (o *StoragePolicy) GetNumNodeFailuresTolerated() int32`

GetNumNodeFailuresTolerated returns the NumNodeFailuresTolerated field if non-nil, zero value otherwise.

### GetNumNodeFailuresToleratedOk

`func (o *StoragePolicy) GetNumNodeFailuresToleratedOk() (*int32, bool)`

GetNumNodeFailuresToleratedOk returns a tuple with the NumNodeFailuresTolerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumNodeFailuresTolerated

`func (o *StoragePolicy) SetNumNodeFailuresTolerated(v int32)`

SetNumNodeFailuresTolerated sets NumNodeFailuresTolerated field to given value.

### HasNumNodeFailuresTolerated

`func (o *StoragePolicy) HasNumNodeFailuresTolerated() bool`

HasNumNodeFailuresTolerated returns a boolean if a field has been set.

### SetNumNodeFailuresToleratedNil

`func (o *StoragePolicy) SetNumNodeFailuresToleratedNil(b bool)`

 SetNumNodeFailuresToleratedNil sets the value for NumNodeFailuresTolerated to be an explicit nil

### UnsetNumNodeFailuresTolerated
`func (o *StoragePolicy) UnsetNumNodeFailuresTolerated()`

UnsetNumNodeFailuresTolerated ensures that no value is present for NumNodeFailuresTolerated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


