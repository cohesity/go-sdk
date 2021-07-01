# UsageAndPerformanceStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataInBytes** | Pointer to **NullableInt64** | Specifies the data read from the protected objects by the Cohesity Cluster before any data reduction using deduplication and compression. | [optional] 
**DataInBytesAfterReduction** | Pointer to **NullableInt64** | Morphed Usage before data is replicated to other nodes as per RF or Erasure Coding policy. | [optional] 
**MinUsablePhysicalCapacityBytes** | Pointer to **NullableInt64** | Specifies the minimum usable capacity available after erasure coding or RF. This will only be populated for cluster. If a cluster has multiple Domains (View Boxes) with different RF or erasure coding, this metric will be computed using the scheme that will provide least saving. | [optional] 
**NumBytesRead** | Pointer to **NullableInt64** | Provides the total number of bytes read in the last 30 seconds. | [optional] 
**NumBytesWritten** | Pointer to **NullableInt64** | Provides the total number of bytes written in the last 30 second. | [optional] 
**PhysicalCapacityBytes** | Pointer to **NullableInt64** | Provides the total physical capacity in bytes of all the storage devices, after subtracting space reserved for cluster services | [optional] 
**ReadIos** | Pointer to **NullableInt64** | Provides the number of Read IOs that occurred in the last 30 seconds. | [optional] 
**ReadLatencyMsecs** | Pointer to **NullableFloat64** | Provides the Read latency in milliseconds for the Read IOs that occurred during the last 30 seconds. | [optional] 
**SystemCapacityBytes** | Pointer to **NullableInt64** | Provides the total available capacity as computed by the Linux &#39;statfs&#39; command. | [optional] 
**SystemUsageBytes** | Pointer to **NullableInt64** | Provides the usage of bytes, as computed by the Linux &#39;statfs&#39; command, after the size of the data is reduced by change-block tracking, compression and deduplication. | [optional] 
**TotalPhysicalRawUsageBytes** | Pointer to **NullableInt64** | Provides the usage of bytes, as computed by the Cohesity Cluster, before the size of the data is reduced by change-block tracking, compression and deduplication. | [optional] 
**TotalPhysicalUsageBytes** | Pointer to **NullableInt64** | Provides the data stored locally, after the data has been reduced by deduplication and compression, including the space required for honoring the resiliency settings (EC/RF). | [optional] 
**WriteIos** | Pointer to **NullableInt64** | Provides the number of Write IOs that occurred in the last 30 seconds. | [optional] 
**WriteLatencyMsecs** | Pointer to **NullableFloat64** | Provides the Write latency in milliseconds for the Write IOs that occurred during the last 30 seconds. | [optional] 

## Methods

### NewUsageAndPerformanceStats

`func NewUsageAndPerformanceStats() *UsageAndPerformanceStats`

NewUsageAndPerformanceStats instantiates a new UsageAndPerformanceStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageAndPerformanceStatsWithDefaults

`func NewUsageAndPerformanceStatsWithDefaults() *UsageAndPerformanceStats`

NewUsageAndPerformanceStatsWithDefaults instantiates a new UsageAndPerformanceStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataInBytes

`func (o *UsageAndPerformanceStats) GetDataInBytes() int64`

GetDataInBytes returns the DataInBytes field if non-nil, zero value otherwise.

### GetDataInBytesOk

`func (o *UsageAndPerformanceStats) GetDataInBytesOk() (*int64, bool)`

GetDataInBytesOk returns a tuple with the DataInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataInBytes

`func (o *UsageAndPerformanceStats) SetDataInBytes(v int64)`

SetDataInBytes sets DataInBytes field to given value.

### HasDataInBytes

`func (o *UsageAndPerformanceStats) HasDataInBytes() bool`

HasDataInBytes returns a boolean if a field has been set.

### SetDataInBytesNil

`func (o *UsageAndPerformanceStats) SetDataInBytesNil(b bool)`

 SetDataInBytesNil sets the value for DataInBytes to be an explicit nil

### UnsetDataInBytes
`func (o *UsageAndPerformanceStats) UnsetDataInBytes()`

UnsetDataInBytes ensures that no value is present for DataInBytes, not even an explicit nil
### GetDataInBytesAfterReduction

`func (o *UsageAndPerformanceStats) GetDataInBytesAfterReduction() int64`

GetDataInBytesAfterReduction returns the DataInBytesAfterReduction field if non-nil, zero value otherwise.

### GetDataInBytesAfterReductionOk

`func (o *UsageAndPerformanceStats) GetDataInBytesAfterReductionOk() (*int64, bool)`

GetDataInBytesAfterReductionOk returns a tuple with the DataInBytesAfterReduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataInBytesAfterReduction

`func (o *UsageAndPerformanceStats) SetDataInBytesAfterReduction(v int64)`

SetDataInBytesAfterReduction sets DataInBytesAfterReduction field to given value.

### HasDataInBytesAfterReduction

`func (o *UsageAndPerformanceStats) HasDataInBytesAfterReduction() bool`

HasDataInBytesAfterReduction returns a boolean if a field has been set.

### SetDataInBytesAfterReductionNil

`func (o *UsageAndPerformanceStats) SetDataInBytesAfterReductionNil(b bool)`

 SetDataInBytesAfterReductionNil sets the value for DataInBytesAfterReduction to be an explicit nil

### UnsetDataInBytesAfterReduction
`func (o *UsageAndPerformanceStats) UnsetDataInBytesAfterReduction()`

UnsetDataInBytesAfterReduction ensures that no value is present for DataInBytesAfterReduction, not even an explicit nil
### GetMinUsablePhysicalCapacityBytes

`func (o *UsageAndPerformanceStats) GetMinUsablePhysicalCapacityBytes() int64`

GetMinUsablePhysicalCapacityBytes returns the MinUsablePhysicalCapacityBytes field if non-nil, zero value otherwise.

### GetMinUsablePhysicalCapacityBytesOk

`func (o *UsageAndPerformanceStats) GetMinUsablePhysicalCapacityBytesOk() (*int64, bool)`

GetMinUsablePhysicalCapacityBytesOk returns a tuple with the MinUsablePhysicalCapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUsablePhysicalCapacityBytes

`func (o *UsageAndPerformanceStats) SetMinUsablePhysicalCapacityBytes(v int64)`

SetMinUsablePhysicalCapacityBytes sets MinUsablePhysicalCapacityBytes field to given value.

### HasMinUsablePhysicalCapacityBytes

`func (o *UsageAndPerformanceStats) HasMinUsablePhysicalCapacityBytes() bool`

HasMinUsablePhysicalCapacityBytes returns a boolean if a field has been set.

### SetMinUsablePhysicalCapacityBytesNil

`func (o *UsageAndPerformanceStats) SetMinUsablePhysicalCapacityBytesNil(b bool)`

 SetMinUsablePhysicalCapacityBytesNil sets the value for MinUsablePhysicalCapacityBytes to be an explicit nil

### UnsetMinUsablePhysicalCapacityBytes
`func (o *UsageAndPerformanceStats) UnsetMinUsablePhysicalCapacityBytes()`

UnsetMinUsablePhysicalCapacityBytes ensures that no value is present for MinUsablePhysicalCapacityBytes, not even an explicit nil
### GetNumBytesRead

`func (o *UsageAndPerformanceStats) GetNumBytesRead() int64`

GetNumBytesRead returns the NumBytesRead field if non-nil, zero value otherwise.

### GetNumBytesReadOk

`func (o *UsageAndPerformanceStats) GetNumBytesReadOk() (*int64, bool)`

GetNumBytesReadOk returns a tuple with the NumBytesRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumBytesRead

`func (o *UsageAndPerformanceStats) SetNumBytesRead(v int64)`

SetNumBytesRead sets NumBytesRead field to given value.

### HasNumBytesRead

`func (o *UsageAndPerformanceStats) HasNumBytesRead() bool`

HasNumBytesRead returns a boolean if a field has been set.

### SetNumBytesReadNil

`func (o *UsageAndPerformanceStats) SetNumBytesReadNil(b bool)`

 SetNumBytesReadNil sets the value for NumBytesRead to be an explicit nil

### UnsetNumBytesRead
`func (o *UsageAndPerformanceStats) UnsetNumBytesRead()`

UnsetNumBytesRead ensures that no value is present for NumBytesRead, not even an explicit nil
### GetNumBytesWritten

`func (o *UsageAndPerformanceStats) GetNumBytesWritten() int64`

GetNumBytesWritten returns the NumBytesWritten field if non-nil, zero value otherwise.

### GetNumBytesWrittenOk

`func (o *UsageAndPerformanceStats) GetNumBytesWrittenOk() (*int64, bool)`

GetNumBytesWrittenOk returns a tuple with the NumBytesWritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumBytesWritten

`func (o *UsageAndPerformanceStats) SetNumBytesWritten(v int64)`

SetNumBytesWritten sets NumBytesWritten field to given value.

### HasNumBytesWritten

`func (o *UsageAndPerformanceStats) HasNumBytesWritten() bool`

HasNumBytesWritten returns a boolean if a field has been set.

### SetNumBytesWrittenNil

`func (o *UsageAndPerformanceStats) SetNumBytesWrittenNil(b bool)`

 SetNumBytesWrittenNil sets the value for NumBytesWritten to be an explicit nil

### UnsetNumBytesWritten
`func (o *UsageAndPerformanceStats) UnsetNumBytesWritten()`

UnsetNumBytesWritten ensures that no value is present for NumBytesWritten, not even an explicit nil
### GetPhysicalCapacityBytes

`func (o *UsageAndPerformanceStats) GetPhysicalCapacityBytes() int64`

GetPhysicalCapacityBytes returns the PhysicalCapacityBytes field if non-nil, zero value otherwise.

### GetPhysicalCapacityBytesOk

`func (o *UsageAndPerformanceStats) GetPhysicalCapacityBytesOk() (*int64, bool)`

GetPhysicalCapacityBytesOk returns a tuple with the PhysicalCapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalCapacityBytes

`func (o *UsageAndPerformanceStats) SetPhysicalCapacityBytes(v int64)`

SetPhysicalCapacityBytes sets PhysicalCapacityBytes field to given value.

### HasPhysicalCapacityBytes

`func (o *UsageAndPerformanceStats) HasPhysicalCapacityBytes() bool`

HasPhysicalCapacityBytes returns a boolean if a field has been set.

### SetPhysicalCapacityBytesNil

`func (o *UsageAndPerformanceStats) SetPhysicalCapacityBytesNil(b bool)`

 SetPhysicalCapacityBytesNil sets the value for PhysicalCapacityBytes to be an explicit nil

### UnsetPhysicalCapacityBytes
`func (o *UsageAndPerformanceStats) UnsetPhysicalCapacityBytes()`

UnsetPhysicalCapacityBytes ensures that no value is present for PhysicalCapacityBytes, not even an explicit nil
### GetReadIos

`func (o *UsageAndPerformanceStats) GetReadIos() int64`

GetReadIos returns the ReadIos field if non-nil, zero value otherwise.

### GetReadIosOk

`func (o *UsageAndPerformanceStats) GetReadIosOk() (*int64, bool)`

GetReadIosOk returns a tuple with the ReadIos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadIos

`func (o *UsageAndPerformanceStats) SetReadIos(v int64)`

SetReadIos sets ReadIos field to given value.

### HasReadIos

`func (o *UsageAndPerformanceStats) HasReadIos() bool`

HasReadIos returns a boolean if a field has been set.

### SetReadIosNil

`func (o *UsageAndPerformanceStats) SetReadIosNil(b bool)`

 SetReadIosNil sets the value for ReadIos to be an explicit nil

### UnsetReadIos
`func (o *UsageAndPerformanceStats) UnsetReadIos()`

UnsetReadIos ensures that no value is present for ReadIos, not even an explicit nil
### GetReadLatencyMsecs

`func (o *UsageAndPerformanceStats) GetReadLatencyMsecs() float64`

GetReadLatencyMsecs returns the ReadLatencyMsecs field if non-nil, zero value otherwise.

### GetReadLatencyMsecsOk

`func (o *UsageAndPerformanceStats) GetReadLatencyMsecsOk() (*float64, bool)`

GetReadLatencyMsecsOk returns a tuple with the ReadLatencyMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadLatencyMsecs

`func (o *UsageAndPerformanceStats) SetReadLatencyMsecs(v float64)`

SetReadLatencyMsecs sets ReadLatencyMsecs field to given value.

### HasReadLatencyMsecs

`func (o *UsageAndPerformanceStats) HasReadLatencyMsecs() bool`

HasReadLatencyMsecs returns a boolean if a field has been set.

### SetReadLatencyMsecsNil

`func (o *UsageAndPerformanceStats) SetReadLatencyMsecsNil(b bool)`

 SetReadLatencyMsecsNil sets the value for ReadLatencyMsecs to be an explicit nil

### UnsetReadLatencyMsecs
`func (o *UsageAndPerformanceStats) UnsetReadLatencyMsecs()`

UnsetReadLatencyMsecs ensures that no value is present for ReadLatencyMsecs, not even an explicit nil
### GetSystemCapacityBytes

`func (o *UsageAndPerformanceStats) GetSystemCapacityBytes() int64`

GetSystemCapacityBytes returns the SystemCapacityBytes field if non-nil, zero value otherwise.

### GetSystemCapacityBytesOk

`func (o *UsageAndPerformanceStats) GetSystemCapacityBytesOk() (*int64, bool)`

GetSystemCapacityBytesOk returns a tuple with the SystemCapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemCapacityBytes

`func (o *UsageAndPerformanceStats) SetSystemCapacityBytes(v int64)`

SetSystemCapacityBytes sets SystemCapacityBytes field to given value.

### HasSystemCapacityBytes

`func (o *UsageAndPerformanceStats) HasSystemCapacityBytes() bool`

HasSystemCapacityBytes returns a boolean if a field has been set.

### SetSystemCapacityBytesNil

`func (o *UsageAndPerformanceStats) SetSystemCapacityBytesNil(b bool)`

 SetSystemCapacityBytesNil sets the value for SystemCapacityBytes to be an explicit nil

### UnsetSystemCapacityBytes
`func (o *UsageAndPerformanceStats) UnsetSystemCapacityBytes()`

UnsetSystemCapacityBytes ensures that no value is present for SystemCapacityBytes, not even an explicit nil
### GetSystemUsageBytes

`func (o *UsageAndPerformanceStats) GetSystemUsageBytes() int64`

GetSystemUsageBytes returns the SystemUsageBytes field if non-nil, zero value otherwise.

### GetSystemUsageBytesOk

`func (o *UsageAndPerformanceStats) GetSystemUsageBytesOk() (*int64, bool)`

GetSystemUsageBytesOk returns a tuple with the SystemUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUsageBytes

`func (o *UsageAndPerformanceStats) SetSystemUsageBytes(v int64)`

SetSystemUsageBytes sets SystemUsageBytes field to given value.

### HasSystemUsageBytes

`func (o *UsageAndPerformanceStats) HasSystemUsageBytes() bool`

HasSystemUsageBytes returns a boolean if a field has been set.

### SetSystemUsageBytesNil

`func (o *UsageAndPerformanceStats) SetSystemUsageBytesNil(b bool)`

 SetSystemUsageBytesNil sets the value for SystemUsageBytes to be an explicit nil

### UnsetSystemUsageBytes
`func (o *UsageAndPerformanceStats) UnsetSystemUsageBytes()`

UnsetSystemUsageBytes ensures that no value is present for SystemUsageBytes, not even an explicit nil
### GetTotalPhysicalRawUsageBytes

`func (o *UsageAndPerformanceStats) GetTotalPhysicalRawUsageBytes() int64`

GetTotalPhysicalRawUsageBytes returns the TotalPhysicalRawUsageBytes field if non-nil, zero value otherwise.

### GetTotalPhysicalRawUsageBytesOk

`func (o *UsageAndPerformanceStats) GetTotalPhysicalRawUsageBytesOk() (*int64, bool)`

GetTotalPhysicalRawUsageBytesOk returns a tuple with the TotalPhysicalRawUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPhysicalRawUsageBytes

`func (o *UsageAndPerformanceStats) SetTotalPhysicalRawUsageBytes(v int64)`

SetTotalPhysicalRawUsageBytes sets TotalPhysicalRawUsageBytes field to given value.

### HasTotalPhysicalRawUsageBytes

`func (o *UsageAndPerformanceStats) HasTotalPhysicalRawUsageBytes() bool`

HasTotalPhysicalRawUsageBytes returns a boolean if a field has been set.

### SetTotalPhysicalRawUsageBytesNil

`func (o *UsageAndPerformanceStats) SetTotalPhysicalRawUsageBytesNil(b bool)`

 SetTotalPhysicalRawUsageBytesNil sets the value for TotalPhysicalRawUsageBytes to be an explicit nil

### UnsetTotalPhysicalRawUsageBytes
`func (o *UsageAndPerformanceStats) UnsetTotalPhysicalRawUsageBytes()`

UnsetTotalPhysicalRawUsageBytes ensures that no value is present for TotalPhysicalRawUsageBytes, not even an explicit nil
### GetTotalPhysicalUsageBytes

`func (o *UsageAndPerformanceStats) GetTotalPhysicalUsageBytes() int64`

GetTotalPhysicalUsageBytes returns the TotalPhysicalUsageBytes field if non-nil, zero value otherwise.

### GetTotalPhysicalUsageBytesOk

`func (o *UsageAndPerformanceStats) GetTotalPhysicalUsageBytesOk() (*int64, bool)`

GetTotalPhysicalUsageBytesOk returns a tuple with the TotalPhysicalUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPhysicalUsageBytes

`func (o *UsageAndPerformanceStats) SetTotalPhysicalUsageBytes(v int64)`

SetTotalPhysicalUsageBytes sets TotalPhysicalUsageBytes field to given value.

### HasTotalPhysicalUsageBytes

`func (o *UsageAndPerformanceStats) HasTotalPhysicalUsageBytes() bool`

HasTotalPhysicalUsageBytes returns a boolean if a field has been set.

### SetTotalPhysicalUsageBytesNil

`func (o *UsageAndPerformanceStats) SetTotalPhysicalUsageBytesNil(b bool)`

 SetTotalPhysicalUsageBytesNil sets the value for TotalPhysicalUsageBytes to be an explicit nil

### UnsetTotalPhysicalUsageBytes
`func (o *UsageAndPerformanceStats) UnsetTotalPhysicalUsageBytes()`

UnsetTotalPhysicalUsageBytes ensures that no value is present for TotalPhysicalUsageBytes, not even an explicit nil
### GetWriteIos

`func (o *UsageAndPerformanceStats) GetWriteIos() int64`

GetWriteIos returns the WriteIos field if non-nil, zero value otherwise.

### GetWriteIosOk

`func (o *UsageAndPerformanceStats) GetWriteIosOk() (*int64, bool)`

GetWriteIosOk returns a tuple with the WriteIos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteIos

`func (o *UsageAndPerformanceStats) SetWriteIos(v int64)`

SetWriteIos sets WriteIos field to given value.

### HasWriteIos

`func (o *UsageAndPerformanceStats) HasWriteIos() bool`

HasWriteIos returns a boolean if a field has been set.

### SetWriteIosNil

`func (o *UsageAndPerformanceStats) SetWriteIosNil(b bool)`

 SetWriteIosNil sets the value for WriteIos to be an explicit nil

### UnsetWriteIos
`func (o *UsageAndPerformanceStats) UnsetWriteIos()`

UnsetWriteIos ensures that no value is present for WriteIos, not even an explicit nil
### GetWriteLatencyMsecs

`func (o *UsageAndPerformanceStats) GetWriteLatencyMsecs() float64`

GetWriteLatencyMsecs returns the WriteLatencyMsecs field if non-nil, zero value otherwise.

### GetWriteLatencyMsecsOk

`func (o *UsageAndPerformanceStats) GetWriteLatencyMsecsOk() (*float64, bool)`

GetWriteLatencyMsecsOk returns a tuple with the WriteLatencyMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteLatencyMsecs

`func (o *UsageAndPerformanceStats) SetWriteLatencyMsecs(v float64)`

SetWriteLatencyMsecs sets WriteLatencyMsecs field to given value.

### HasWriteLatencyMsecs

`func (o *UsageAndPerformanceStats) HasWriteLatencyMsecs() bool`

HasWriteLatencyMsecs returns a boolean if a field has been set.

### SetWriteLatencyMsecsNil

`func (o *UsageAndPerformanceStats) SetWriteLatencyMsecsNil(b bool)`

 SetWriteLatencyMsecsNil sets the value for WriteLatencyMsecs to be an explicit nil

### UnsetWriteLatencyMsecs
`func (o *UsageAndPerformanceStats) UnsetWriteLatencyMsecs()`

UnsetWriteLatencyMsecs ensures that no value is present for WriteLatencyMsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


