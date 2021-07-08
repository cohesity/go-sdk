# StoragePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppMarkerDetection** | Pointer to **NullableBool** | Specifies Whether to support app marker detection. When this is set to true, app markers (like commvault markers) will be removed from data and put in separate chunks. This way deduplication is improved as it is done on data that has no app markers. | [optional] 
**CloudSpillVaultId** | Pointer to **NullableInt64** | Specifies the vault id assigned for an external Storage Target to facilitate cloud spill. | [optional] 
**CompressionPolicy** | Pointer to **NullableString** | Specifies the compression setting to be applied to a Storage Domain (View Box). &#39;kCompressionNone&#39; indicates that data is not compressed. &#39;kCompressionLow&#39; indicates that data is compressed using LZ4 or Snappy. &#39;kCompressionHigh&#39; indicates that data is compressed in Gzip. | [optional] 
**DeduplicateCompressDelaySecs** | Pointer to **NullableInt32** | Specifies the time in seconds when deduplication and compression of data on the Storage Domain (View Box) starts. If set to 0, deduplication and compression is done inline (as the data is being written). Otherwise, post-process deduplication and compression is done after the specified delay. | [optional] 
**DeduplicationEnabled** | Pointer to **NullableBool** | Specifies if deduplication is enabled for the Storage Domain (View Box). If deduplication is enabled, the Cohesity Cluster eliminates duplicate blocks of repeating data stored on the Cluster thus reducing the amount of storage space needed to store data. | [optional] 
**EncryptionPolicy** | Pointer to **NullableString** | Specifies the encryption setting for the Storage Domain (View Box). &#39;kEncryptionNone&#39; indicates the data is not encrypted. &#39;kEncryptionStrong&#39; indicates the data is encrypted. | [optional] 
**ErasureCodingInfo** | Pointer to [**ErasureCodingInfo**](ErasureCodingInfo.md) |  | [optional] 
**InlineCompress** | Pointer to **NullableBool** | Specifies if compression should occur inline (as the data is being written). This field is only relevant if compression is enabled. If deduplication is set to inline, Cohesity recommends setting compression to inline. | [optional] 
**InlineDeduplicate** | Pointer to **NullableBool** | Specifies if deduplication should occur inline (as the data is being written). This field is only relevant if deduplication is enabled. | [optional] 
**NumFailuresTolerated** | Pointer to **NullableInt32** | Number of disk failures to tolerate. This is an optional field. Default value is 1 for cluster having 3 or more nodes. If erasure coding is not enabled, then this specifies the replication factor for the Storage Domain (View Box). For RF&#x3D;2, number of failures to tolerate should be specified as 1. If erasure coding is enabled, then this value will be same as number of coded stripes. | [optional] 
**NumNodeFailuresTolerated** | Pointer to **NullableInt32** | Number of node failures to tolerate. If NumNodeFailuresTolerated is set to 2, then we would tolerate up to two node failures. If the following is not set, then the number of node failures tolerated would be same as replication factor - 1 for replicated chunk files or number of coded stripes for erasure coding chunk files. | [optional] 

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

### GetAppMarkerDetection

`func (o *StoragePolicy) GetAppMarkerDetection() bool`

GetAppMarkerDetection returns the AppMarkerDetection field if non-nil, zero value otherwise.

### GetAppMarkerDetectionOk

`func (o *StoragePolicy) GetAppMarkerDetectionOk() (*bool, bool)`

GetAppMarkerDetectionOk returns a tuple with the AppMarkerDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppMarkerDetection

`func (o *StoragePolicy) SetAppMarkerDetection(v bool)`

SetAppMarkerDetection sets AppMarkerDetection field to given value.

### HasAppMarkerDetection

`func (o *StoragePolicy) HasAppMarkerDetection() bool`

HasAppMarkerDetection returns a boolean if a field has been set.

### SetAppMarkerDetectionNil

`func (o *StoragePolicy) SetAppMarkerDetectionNil(b bool)`

 SetAppMarkerDetectionNil sets the value for AppMarkerDetection to be an explicit nil

### UnsetAppMarkerDetection
`func (o *StoragePolicy) UnsetAppMarkerDetection()`

UnsetAppMarkerDetection ensures that no value is present for AppMarkerDetection, not even an explicit nil
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
### GetCompressionPolicy

`func (o *StoragePolicy) GetCompressionPolicy() string`

GetCompressionPolicy returns the CompressionPolicy field if non-nil, zero value otherwise.

### GetCompressionPolicyOk

`func (o *StoragePolicy) GetCompressionPolicyOk() (*string, bool)`

GetCompressionPolicyOk returns a tuple with the CompressionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionPolicy

`func (o *StoragePolicy) SetCompressionPolicy(v string)`

SetCompressionPolicy sets CompressionPolicy field to given value.

### HasCompressionPolicy

`func (o *StoragePolicy) HasCompressionPolicy() bool`

HasCompressionPolicy returns a boolean if a field has been set.

### SetCompressionPolicyNil

`func (o *StoragePolicy) SetCompressionPolicyNil(b bool)`

 SetCompressionPolicyNil sets the value for CompressionPolicy to be an explicit nil

### UnsetCompressionPolicy
`func (o *StoragePolicy) UnsetCompressionPolicy()`

UnsetCompressionPolicy ensures that no value is present for CompressionPolicy, not even an explicit nil
### GetDeduplicateCompressDelaySecs

`func (o *StoragePolicy) GetDeduplicateCompressDelaySecs() int32`

GetDeduplicateCompressDelaySecs returns the DeduplicateCompressDelaySecs field if non-nil, zero value otherwise.

### GetDeduplicateCompressDelaySecsOk

`func (o *StoragePolicy) GetDeduplicateCompressDelaySecsOk() (*int32, bool)`

GetDeduplicateCompressDelaySecsOk returns a tuple with the DeduplicateCompressDelaySecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplicateCompressDelaySecs

`func (o *StoragePolicy) SetDeduplicateCompressDelaySecs(v int32)`

SetDeduplicateCompressDelaySecs sets DeduplicateCompressDelaySecs field to given value.

### HasDeduplicateCompressDelaySecs

`func (o *StoragePolicy) HasDeduplicateCompressDelaySecs() bool`

HasDeduplicateCompressDelaySecs returns a boolean if a field has been set.

### SetDeduplicateCompressDelaySecsNil

`func (o *StoragePolicy) SetDeduplicateCompressDelaySecsNil(b bool)`

 SetDeduplicateCompressDelaySecsNil sets the value for DeduplicateCompressDelaySecs to be an explicit nil

### UnsetDeduplicateCompressDelaySecs
`func (o *StoragePolicy) UnsetDeduplicateCompressDelaySecs()`

UnsetDeduplicateCompressDelaySecs ensures that no value is present for DeduplicateCompressDelaySecs, not even an explicit nil
### GetDeduplicationEnabled

`func (o *StoragePolicy) GetDeduplicationEnabled() bool`

GetDeduplicationEnabled returns the DeduplicationEnabled field if non-nil, zero value otherwise.

### GetDeduplicationEnabledOk

`func (o *StoragePolicy) GetDeduplicationEnabledOk() (*bool, bool)`

GetDeduplicationEnabledOk returns a tuple with the DeduplicationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplicationEnabled

`func (o *StoragePolicy) SetDeduplicationEnabled(v bool)`

SetDeduplicationEnabled sets DeduplicationEnabled field to given value.

### HasDeduplicationEnabled

`func (o *StoragePolicy) HasDeduplicationEnabled() bool`

HasDeduplicationEnabled returns a boolean if a field has been set.

### SetDeduplicationEnabledNil

`func (o *StoragePolicy) SetDeduplicationEnabledNil(b bool)`

 SetDeduplicationEnabledNil sets the value for DeduplicationEnabled to be an explicit nil

### UnsetDeduplicationEnabled
`func (o *StoragePolicy) UnsetDeduplicationEnabled()`

UnsetDeduplicationEnabled ensures that no value is present for DeduplicationEnabled, not even an explicit nil
### GetEncryptionPolicy

`func (o *StoragePolicy) GetEncryptionPolicy() string`

GetEncryptionPolicy returns the EncryptionPolicy field if non-nil, zero value otherwise.

### GetEncryptionPolicyOk

`func (o *StoragePolicy) GetEncryptionPolicyOk() (*string, bool)`

GetEncryptionPolicyOk returns a tuple with the EncryptionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionPolicy

`func (o *StoragePolicy) SetEncryptionPolicy(v string)`

SetEncryptionPolicy sets EncryptionPolicy field to given value.

### HasEncryptionPolicy

`func (o *StoragePolicy) HasEncryptionPolicy() bool`

HasEncryptionPolicy returns a boolean if a field has been set.

### SetEncryptionPolicyNil

`func (o *StoragePolicy) SetEncryptionPolicyNil(b bool)`

 SetEncryptionPolicyNil sets the value for EncryptionPolicy to be an explicit nil

### UnsetEncryptionPolicy
`func (o *StoragePolicy) UnsetEncryptionPolicy()`

UnsetEncryptionPolicy ensures that no value is present for EncryptionPolicy, not even an explicit nil
### GetErasureCodingInfo

`func (o *StoragePolicy) GetErasureCodingInfo() ErasureCodingInfo`

GetErasureCodingInfo returns the ErasureCodingInfo field if non-nil, zero value otherwise.

### GetErasureCodingInfoOk

`func (o *StoragePolicy) GetErasureCodingInfoOk() (*ErasureCodingInfo, bool)`

GetErasureCodingInfoOk returns a tuple with the ErasureCodingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErasureCodingInfo

`func (o *StoragePolicy) SetErasureCodingInfo(v ErasureCodingInfo)`

SetErasureCodingInfo sets ErasureCodingInfo field to given value.

### HasErasureCodingInfo

`func (o *StoragePolicy) HasErasureCodingInfo() bool`

HasErasureCodingInfo returns a boolean if a field has been set.

### GetInlineCompress

`func (o *StoragePolicy) GetInlineCompress() bool`

GetInlineCompress returns the InlineCompress field if non-nil, zero value otherwise.

### GetInlineCompressOk

`func (o *StoragePolicy) GetInlineCompressOk() (*bool, bool)`

GetInlineCompressOk returns a tuple with the InlineCompress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineCompress

`func (o *StoragePolicy) SetInlineCompress(v bool)`

SetInlineCompress sets InlineCompress field to given value.

### HasInlineCompress

`func (o *StoragePolicy) HasInlineCompress() bool`

HasInlineCompress returns a boolean if a field has been set.

### SetInlineCompressNil

`func (o *StoragePolicy) SetInlineCompressNil(b bool)`

 SetInlineCompressNil sets the value for InlineCompress to be an explicit nil

### UnsetInlineCompress
`func (o *StoragePolicy) UnsetInlineCompress()`

UnsetInlineCompress ensures that no value is present for InlineCompress, not even an explicit nil
### GetInlineDeduplicate

`func (o *StoragePolicy) GetInlineDeduplicate() bool`

GetInlineDeduplicate returns the InlineDeduplicate field if non-nil, zero value otherwise.

### GetInlineDeduplicateOk

`func (o *StoragePolicy) GetInlineDeduplicateOk() (*bool, bool)`

GetInlineDeduplicateOk returns a tuple with the InlineDeduplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineDeduplicate

`func (o *StoragePolicy) SetInlineDeduplicate(v bool)`

SetInlineDeduplicate sets InlineDeduplicate field to given value.

### HasInlineDeduplicate

`func (o *StoragePolicy) HasInlineDeduplicate() bool`

HasInlineDeduplicate returns a boolean if a field has been set.

### SetInlineDeduplicateNil

`func (o *StoragePolicy) SetInlineDeduplicateNil(b bool)`

 SetInlineDeduplicateNil sets the value for InlineDeduplicate to be an explicit nil

### UnsetInlineDeduplicate
`func (o *StoragePolicy) UnsetInlineDeduplicate()`

UnsetInlineDeduplicate ensures that no value is present for InlineDeduplicate, not even an explicit nil
### GetNumFailuresTolerated

`func (o *StoragePolicy) GetNumFailuresTolerated() int32`

GetNumFailuresTolerated returns the NumFailuresTolerated field if non-nil, zero value otherwise.

### GetNumFailuresToleratedOk

`func (o *StoragePolicy) GetNumFailuresToleratedOk() (*int32, bool)`

GetNumFailuresToleratedOk returns a tuple with the NumFailuresTolerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFailuresTolerated

`func (o *StoragePolicy) SetNumFailuresTolerated(v int32)`

SetNumFailuresTolerated sets NumFailuresTolerated field to given value.

### HasNumFailuresTolerated

`func (o *StoragePolicy) HasNumFailuresTolerated() bool`

HasNumFailuresTolerated returns a boolean if a field has been set.

### SetNumFailuresToleratedNil

`func (o *StoragePolicy) SetNumFailuresToleratedNil(b bool)`

 SetNumFailuresToleratedNil sets the value for NumFailuresTolerated to be an explicit nil

### UnsetNumFailuresTolerated
`func (o *StoragePolicy) UnsetNumFailuresTolerated()`

UnsetNumFailuresTolerated ensures that no value is present for NumFailuresTolerated, not even an explicit nil
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


