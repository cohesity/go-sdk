# StorageDomainStoragePolicy

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

### NewStorageDomainStoragePolicy

`func NewStorageDomainStoragePolicy() *StorageDomainStoragePolicy`

NewStorageDomainStoragePolicy instantiates a new StorageDomainStoragePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageDomainStoragePolicyWithDefaults

`func NewStorageDomainStoragePolicyWithDefaults() *StorageDomainStoragePolicy`

NewStorageDomainStoragePolicyWithDefaults instantiates a new StorageDomainStoragePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAesEncryptionMode

`func (o *StorageDomainStoragePolicy) GetAesEncryptionMode() string`

GetAesEncryptionMode returns the AesEncryptionMode field if non-nil, zero value otherwise.

### GetAesEncryptionModeOk

`func (o *StorageDomainStoragePolicy) GetAesEncryptionModeOk() (*string, bool)`

GetAesEncryptionModeOk returns a tuple with the AesEncryptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAesEncryptionMode

`func (o *StorageDomainStoragePolicy) SetAesEncryptionMode(v string)`

SetAesEncryptionMode sets AesEncryptionMode field to given value.

### HasAesEncryptionMode

`func (o *StorageDomainStoragePolicy) HasAesEncryptionMode() bool`

HasAesEncryptionMode returns a boolean if a field has been set.

### SetAesEncryptionModeNil

`func (o *StorageDomainStoragePolicy) SetAesEncryptionModeNil(b bool)`

 SetAesEncryptionModeNil sets the value for AesEncryptionMode to be an explicit nil

### UnsetAesEncryptionMode
`func (o *StorageDomainStoragePolicy) UnsetAesEncryptionMode()`

UnsetAesEncryptionMode ensures that no value is present for AesEncryptionMode, not even an explicit nil
### GetAppMarkerDetectionEnabled

`func (o *StorageDomainStoragePolicy) GetAppMarkerDetectionEnabled() bool`

GetAppMarkerDetectionEnabled returns the AppMarkerDetectionEnabled field if non-nil, zero value otherwise.

### GetAppMarkerDetectionEnabledOk

`func (o *StorageDomainStoragePolicy) GetAppMarkerDetectionEnabledOk() (*bool, bool)`

GetAppMarkerDetectionEnabledOk returns a tuple with the AppMarkerDetectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppMarkerDetectionEnabled

`func (o *StorageDomainStoragePolicy) SetAppMarkerDetectionEnabled(v bool)`

SetAppMarkerDetectionEnabled sets AppMarkerDetectionEnabled field to given value.

### HasAppMarkerDetectionEnabled

`func (o *StorageDomainStoragePolicy) HasAppMarkerDetectionEnabled() bool`

HasAppMarkerDetectionEnabled returns a boolean if a field has been set.

### SetAppMarkerDetectionEnabledNil

`func (o *StorageDomainStoragePolicy) SetAppMarkerDetectionEnabledNil(b bool)`

 SetAppMarkerDetectionEnabledNil sets the value for AppMarkerDetectionEnabled to be an explicit nil

### UnsetAppMarkerDetectionEnabled
`func (o *StorageDomainStoragePolicy) UnsetAppMarkerDetectionEnabled()`

UnsetAppMarkerDetectionEnabled ensures that no value is present for AppMarkerDetectionEnabled, not even an explicit nil
### GetCloudSpillVaultId

`func (o *StorageDomainStoragePolicy) GetCloudSpillVaultId() int64`

GetCloudSpillVaultId returns the CloudSpillVaultId field if non-nil, zero value otherwise.

### GetCloudSpillVaultIdOk

`func (o *StorageDomainStoragePolicy) GetCloudSpillVaultIdOk() (*int64, bool)`

GetCloudSpillVaultIdOk returns a tuple with the CloudSpillVaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSpillVaultId

`func (o *StorageDomainStoragePolicy) SetCloudSpillVaultId(v int64)`

SetCloudSpillVaultId sets CloudSpillVaultId field to given value.

### HasCloudSpillVaultId

`func (o *StorageDomainStoragePolicy) HasCloudSpillVaultId() bool`

HasCloudSpillVaultId returns a boolean if a field has been set.

### SetCloudSpillVaultIdNil

`func (o *StorageDomainStoragePolicy) SetCloudSpillVaultIdNil(b bool)`

 SetCloudSpillVaultIdNil sets the value for CloudSpillVaultId to be an explicit nil

### UnsetCloudSpillVaultId
`func (o *StorageDomainStoragePolicy) UnsetCloudSpillVaultId()`

UnsetCloudSpillVaultId ensures that no value is present for CloudSpillVaultId, not even an explicit nil
### GetCompressionParams

`func (o *StorageDomainStoragePolicy) GetCompressionParams() StoragePolicyCompressionParams`

GetCompressionParams returns the CompressionParams field if non-nil, zero value otherwise.

### GetCompressionParamsOk

`func (o *StorageDomainStoragePolicy) GetCompressionParamsOk() (*StoragePolicyCompressionParams, bool)`

GetCompressionParamsOk returns a tuple with the CompressionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionParams

`func (o *StorageDomainStoragePolicy) SetCompressionParams(v StoragePolicyCompressionParams)`

SetCompressionParams sets CompressionParams field to given value.

### HasCompressionParams

`func (o *StorageDomainStoragePolicy) HasCompressionParams() bool`

HasCompressionParams returns a boolean if a field has been set.

### GetDeduplicationCompressionDelaySecs

`func (o *StorageDomainStoragePolicy) GetDeduplicationCompressionDelaySecs() int32`

GetDeduplicationCompressionDelaySecs returns the DeduplicationCompressionDelaySecs field if non-nil, zero value otherwise.

### GetDeduplicationCompressionDelaySecsOk

`func (o *StorageDomainStoragePolicy) GetDeduplicationCompressionDelaySecsOk() (*int32, bool)`

GetDeduplicationCompressionDelaySecsOk returns a tuple with the DeduplicationCompressionDelaySecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplicationCompressionDelaySecs

`func (o *StorageDomainStoragePolicy) SetDeduplicationCompressionDelaySecs(v int32)`

SetDeduplicationCompressionDelaySecs sets DeduplicationCompressionDelaySecs field to given value.

### HasDeduplicationCompressionDelaySecs

`func (o *StorageDomainStoragePolicy) HasDeduplicationCompressionDelaySecs() bool`

HasDeduplicationCompressionDelaySecs returns a boolean if a field has been set.

### SetDeduplicationCompressionDelaySecsNil

`func (o *StorageDomainStoragePolicy) SetDeduplicationCompressionDelaySecsNil(b bool)`

 SetDeduplicationCompressionDelaySecsNil sets the value for DeduplicationCompressionDelaySecs to be an explicit nil

### UnsetDeduplicationCompressionDelaySecs
`func (o *StorageDomainStoragePolicy) UnsetDeduplicationCompressionDelaySecs()`

UnsetDeduplicationCompressionDelaySecs ensures that no value is present for DeduplicationCompressionDelaySecs, not even an explicit nil
### GetDeduplicationParams

`func (o *StorageDomainStoragePolicy) GetDeduplicationParams() StoragePolicyDeduplicationParams`

GetDeduplicationParams returns the DeduplicationParams field if non-nil, zero value otherwise.

### GetDeduplicationParamsOk

`func (o *StorageDomainStoragePolicy) GetDeduplicationParamsOk() (*StoragePolicyDeduplicationParams, bool)`

GetDeduplicationParamsOk returns a tuple with the DeduplicationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplicationParams

`func (o *StorageDomainStoragePolicy) SetDeduplicationParams(v StoragePolicyDeduplicationParams)`

SetDeduplicationParams sets DeduplicationParams field to given value.

### HasDeduplicationParams

`func (o *StorageDomainStoragePolicy) HasDeduplicationParams() bool`

HasDeduplicationParams returns a boolean if a field has been set.

### GetEncryptionType

`func (o *StorageDomainStoragePolicy) GetEncryptionType() string`

GetEncryptionType returns the EncryptionType field if non-nil, zero value otherwise.

### GetEncryptionTypeOk

`func (o *StorageDomainStoragePolicy) GetEncryptionTypeOk() (*string, bool)`

GetEncryptionTypeOk returns a tuple with the EncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionType

`func (o *StorageDomainStoragePolicy) SetEncryptionType(v string)`

SetEncryptionType sets EncryptionType field to given value.

### HasEncryptionType

`func (o *StorageDomainStoragePolicy) HasEncryptionType() bool`

HasEncryptionType returns a boolean if a field has been set.

### SetEncryptionTypeNil

`func (o *StorageDomainStoragePolicy) SetEncryptionTypeNil(b bool)`

 SetEncryptionTypeNil sets the value for EncryptionType to be an explicit nil

### UnsetEncryptionType
`func (o *StorageDomainStoragePolicy) UnsetEncryptionType()`

UnsetEncryptionType ensures that no value is present for EncryptionType, not even an explicit nil
### GetErasureCodingParams

`func (o *StorageDomainStoragePolicy) GetErasureCodingParams() StoragePolicyErasureCodingParams`

GetErasureCodingParams returns the ErasureCodingParams field if non-nil, zero value otherwise.

### GetErasureCodingParamsOk

`func (o *StorageDomainStoragePolicy) GetErasureCodingParamsOk() (*StoragePolicyErasureCodingParams, bool)`

GetErasureCodingParamsOk returns a tuple with the ErasureCodingParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErasureCodingParams

`func (o *StorageDomainStoragePolicy) SetErasureCodingParams(v StoragePolicyErasureCodingParams)`

SetErasureCodingParams sets ErasureCodingParams field to given value.

### HasErasureCodingParams

`func (o *StorageDomainStoragePolicy) HasErasureCodingParams() bool`

HasErasureCodingParams returns a boolean if a field has been set.

### GetNumDiskFailuresTolerated

`func (o *StorageDomainStoragePolicy) GetNumDiskFailuresTolerated() int32`

GetNumDiskFailuresTolerated returns the NumDiskFailuresTolerated field if non-nil, zero value otherwise.

### GetNumDiskFailuresToleratedOk

`func (o *StorageDomainStoragePolicy) GetNumDiskFailuresToleratedOk() (*int32, bool)`

GetNumDiskFailuresToleratedOk returns a tuple with the NumDiskFailuresTolerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDiskFailuresTolerated

`func (o *StorageDomainStoragePolicy) SetNumDiskFailuresTolerated(v int32)`

SetNumDiskFailuresTolerated sets NumDiskFailuresTolerated field to given value.

### HasNumDiskFailuresTolerated

`func (o *StorageDomainStoragePolicy) HasNumDiskFailuresTolerated() bool`

HasNumDiskFailuresTolerated returns a boolean if a field has been set.

### SetNumDiskFailuresToleratedNil

`func (o *StorageDomainStoragePolicy) SetNumDiskFailuresToleratedNil(b bool)`

 SetNumDiskFailuresToleratedNil sets the value for NumDiskFailuresTolerated to be an explicit nil

### UnsetNumDiskFailuresTolerated
`func (o *StorageDomainStoragePolicy) UnsetNumDiskFailuresTolerated()`

UnsetNumDiskFailuresTolerated ensures that no value is present for NumDiskFailuresTolerated, not even an explicit nil
### GetNumNodeFailuresTolerated

`func (o *StorageDomainStoragePolicy) GetNumNodeFailuresTolerated() int32`

GetNumNodeFailuresTolerated returns the NumNodeFailuresTolerated field if non-nil, zero value otherwise.

### GetNumNodeFailuresToleratedOk

`func (o *StorageDomainStoragePolicy) GetNumNodeFailuresToleratedOk() (*int32, bool)`

GetNumNodeFailuresToleratedOk returns a tuple with the NumNodeFailuresTolerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumNodeFailuresTolerated

`func (o *StorageDomainStoragePolicy) SetNumNodeFailuresTolerated(v int32)`

SetNumNodeFailuresTolerated sets NumNodeFailuresTolerated field to given value.

### HasNumNodeFailuresTolerated

`func (o *StorageDomainStoragePolicy) HasNumNodeFailuresTolerated() bool`

HasNumNodeFailuresTolerated returns a boolean if a field has been set.

### SetNumNodeFailuresToleratedNil

`func (o *StorageDomainStoragePolicy) SetNumNodeFailuresToleratedNil(b bool)`

 SetNumNodeFailuresToleratedNil sets the value for NumNodeFailuresTolerated to be an explicit nil

### UnsetNumNodeFailuresTolerated
`func (o *StorageDomainStoragePolicy) UnsetNumNodeFailuresTolerated()`

UnsetNumNodeFailuresTolerated ensures that no value is present for NumNodeFailuresTolerated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


