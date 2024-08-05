# ViewStatsDataUsageStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudDataWrittenBytes** | Pointer to **NullableInt64** | Specifies the total data written on cloud tiers, as computed by the Cohesity Cluster. | [optional] 
**CloudDataWrittenBytesTimestampUsec** | Pointer to **NullableInt64** | Specifies Timestamp of CloudDataWrittenBytes. | [optional] 
**CloudTotalPhysicalUsageBytes** | Pointer to **NullableInt64** | Specifies the total cloud capacity, as computed by the Cohesity Cluster, after the size of the data has been reduced by change-block tracking, compression and deduplication. | [optional] 
**CloudTotalPhysicalUsageBytesTimestampUsec** | Pointer to **NullableInt64** | Specifies Timestamp of CloudTotalPhysicalUsageBytes. | [optional] 
**DataInBytes** | Pointer to **NullableInt64** | Specifies the data read from the protected objects by the Cohesity Cluster before any data reduction using deduplication and compression. | [optional] 
**DataInBytesAfterDedup** | Pointer to **NullableInt64** | Specifies the size of the data has been reduced by change-block tracking and deduplication but before compression or data is replicated to other nodes as per RF or Erasure Coding policy. | [optional] 
**DataInBytesAfterDedupTimestampUsec** | Pointer to **NullableInt64** | Specifies Timestamp of DataInBytesAfterDedup. | [optional] 
**DataInBytesTimestampUsec** | Pointer to **NullableInt64** | Specifies Timestamp of DataInBytes. | [optional] 
**DataProtectLogicalUsageBytes** | Pointer to **NullableInt64** | Specifies the logical data used by Data Protect on Cohesity cluster. | [optional] 
**DataProtectLogicalUsageBytesTimestampUsec** | Pointer to **NullableInt64** | Specifies Timestamp of DataProtectLogicalUsageBytes. | [optional] 
**DataProtectPhysicalUsageBytes** | Pointer to **NullableInt64** | Specifies the physical data used by Data Protect on Cohesity cluster. | [optional] 
**DataProtectPhysicalUsageBytesTimestampUsec** | Pointer to **NullableInt64** | Specifies Timestamp of DataProtectPhysicalUsageBytes. | [optional] 
**DataWrittenBytes** | Pointer to **NullableInt64** | Specifies the data written after it has been reduced by deduplication and compression. This does not include resiliency impact. | [optional] 
**DataWrittenBytesTimestampUsec** | Pointer to **NullableInt64** | Specifies Timestamp of DataWrittenBytes. | [optional] 
**FileServicesLogicalUsageBytes** | Pointer to **NullableInt64** | Specifies the logical data used by File services on Cohesity cluster. | [optional] 
**FileServicesLogicalUsageBytesTimestampUsec** | Pointer to **NullableInt64** | Specifies Timestamp of FileServicesLogicalUsageBytes. | [optional] 
**FileServicesPhysicalUsageBytes** | Pointer to **NullableInt64** | Specifies the physical data used by File services on Cohesity cluster. | [optional] 
**FileServicesPhysicalUsageBytesTimestampUsec** | Pointer to **NullableInt64** | Specifies Timestamp of FileServicesPhysicalUsageBytes. | [optional] 
**LocalDataWrittenBytes** | Pointer to **NullableInt64** | Specifies the total data written on local tiers, as computed by the Cohesity Cluster, after the size of the data has been reduced by change-block tracking, deduplication and compression. This does not include resiliency impact. | [optional] 
**LocalDataWrittenBytesTimestampUsec** | Pointer to **NullableInt64** | Specifies Timestamp of LocalDataWrittenBytes. | [optional] 
**LocalTierResiliencyImpactBytes** | Pointer to **NullableInt64** | Specifies the size of the data has been replicated to other nodes as per RF or Erasure Coding policy. | [optional] 
**LocalTierResiliencyImpactBytesTimestampUsec** | Pointer to **NullableInt64** | Specifies Timestamp of LocalTierResiliencyImpactBytes. | [optional] 
**LocalTotalPhysicalUsageBytes** | Pointer to **NullableInt64** | Specifies the total local capacity, as computed by the Cohesity Cluster, after the size of the data has been reduced by change-block tracking, compression and deduplication. | [optional] 
**LocalTotalPhysicalUsageBytesTimestampUsec** | Pointer to **NullableInt64** | Specifies Timestamp of LocalTotalPhysicalUsageBytes. | [optional] 
**NumDirectories** | Pointer to **NullableInt64** | Specifies the number of directories. | [optional] 
**NumFiles** | Pointer to **NullableInt64** | Specifies the number of files. | [optional] 
**OutdatedLogicalUsageBytes** | Pointer to **NullableInt64** | Specifies the logical usage as computed by the Cohesity Cluster. This field is computed on a same frequency as &#39;StorageConsumedBytes&#39;, and it may not be the latest value. It is used to compute reduction ratio. | [optional] 
**OutdatedLogicalUsageBytesTimestampUsec** | Pointer to **NullableInt64** | Specifies Timestamp of OutdatedLogicalUsageBytes. | [optional] 
**StorageConsumedBytes** | Pointer to **NullableInt64** | Specifies the total capacity, as computed by the Cohesity Cluster, after the size of the data has been reduced by change-block tracking, compression and deduplication. This includes resiliency impact. | [optional] 
**StorageConsumedBytesTimestampUsec** | Pointer to **NullableInt64** | Specifies Timestamp of StorageConsumedBytes. | [optional] 
**TotalLogicalUsageBytes** | Pointer to **NullableInt64** | Provides the combined data residing on protected objects. The size of data before reduction by deduplication and compression. | [optional] 
**TotalLogicalUsageBytesTimestampUsec** | Pointer to **NullableInt64** | Specifies Timestamp of TotalLogicalUsageBytes. | [optional] 
**UniquePhysicalDataBytes** | Pointer to **NullableInt64** | Specifies the unique physical data usage in bytes. | [optional] 

## Methods

### NewViewStatsDataUsageStats

`func NewViewStatsDataUsageStats() *ViewStatsDataUsageStats`

NewViewStatsDataUsageStats instantiates a new ViewStatsDataUsageStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewStatsDataUsageStatsWithDefaults

`func NewViewStatsDataUsageStatsWithDefaults() *ViewStatsDataUsageStats`

NewViewStatsDataUsageStatsWithDefaults instantiates a new ViewStatsDataUsageStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudDataWrittenBytes

`func (o *ViewStatsDataUsageStats) GetCloudDataWrittenBytes() int64`

GetCloudDataWrittenBytes returns the CloudDataWrittenBytes field if non-nil, zero value otherwise.

### GetCloudDataWrittenBytesOk

`func (o *ViewStatsDataUsageStats) GetCloudDataWrittenBytesOk() (*int64, bool)`

GetCloudDataWrittenBytesOk returns a tuple with the CloudDataWrittenBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDataWrittenBytes

`func (o *ViewStatsDataUsageStats) SetCloudDataWrittenBytes(v int64)`

SetCloudDataWrittenBytes sets CloudDataWrittenBytes field to given value.

### HasCloudDataWrittenBytes

`func (o *ViewStatsDataUsageStats) HasCloudDataWrittenBytes() bool`

HasCloudDataWrittenBytes returns a boolean if a field has been set.

### SetCloudDataWrittenBytesNil

`func (o *ViewStatsDataUsageStats) SetCloudDataWrittenBytesNil(b bool)`

 SetCloudDataWrittenBytesNil sets the value for CloudDataWrittenBytes to be an explicit nil

### UnsetCloudDataWrittenBytes
`func (o *ViewStatsDataUsageStats) UnsetCloudDataWrittenBytes()`

UnsetCloudDataWrittenBytes ensures that no value is present for CloudDataWrittenBytes, not even an explicit nil
### GetCloudDataWrittenBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) GetCloudDataWrittenBytesTimestampUsec() int64`

GetCloudDataWrittenBytesTimestampUsec returns the CloudDataWrittenBytesTimestampUsec field if non-nil, zero value otherwise.

### GetCloudDataWrittenBytesTimestampUsecOk

`func (o *ViewStatsDataUsageStats) GetCloudDataWrittenBytesTimestampUsecOk() (*int64, bool)`

GetCloudDataWrittenBytesTimestampUsecOk returns a tuple with the CloudDataWrittenBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDataWrittenBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) SetCloudDataWrittenBytesTimestampUsec(v int64)`

SetCloudDataWrittenBytesTimestampUsec sets CloudDataWrittenBytesTimestampUsec field to given value.

### HasCloudDataWrittenBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) HasCloudDataWrittenBytesTimestampUsec() bool`

HasCloudDataWrittenBytesTimestampUsec returns a boolean if a field has been set.

### SetCloudDataWrittenBytesTimestampUsecNil

`func (o *ViewStatsDataUsageStats) SetCloudDataWrittenBytesTimestampUsecNil(b bool)`

 SetCloudDataWrittenBytesTimestampUsecNil sets the value for CloudDataWrittenBytesTimestampUsec to be an explicit nil

### UnsetCloudDataWrittenBytesTimestampUsec
`func (o *ViewStatsDataUsageStats) UnsetCloudDataWrittenBytesTimestampUsec()`

UnsetCloudDataWrittenBytesTimestampUsec ensures that no value is present for CloudDataWrittenBytesTimestampUsec, not even an explicit nil
### GetCloudTotalPhysicalUsageBytes

`func (o *ViewStatsDataUsageStats) GetCloudTotalPhysicalUsageBytes() int64`

GetCloudTotalPhysicalUsageBytes returns the CloudTotalPhysicalUsageBytes field if non-nil, zero value otherwise.

### GetCloudTotalPhysicalUsageBytesOk

`func (o *ViewStatsDataUsageStats) GetCloudTotalPhysicalUsageBytesOk() (*int64, bool)`

GetCloudTotalPhysicalUsageBytesOk returns a tuple with the CloudTotalPhysicalUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudTotalPhysicalUsageBytes

`func (o *ViewStatsDataUsageStats) SetCloudTotalPhysicalUsageBytes(v int64)`

SetCloudTotalPhysicalUsageBytes sets CloudTotalPhysicalUsageBytes field to given value.

### HasCloudTotalPhysicalUsageBytes

`func (o *ViewStatsDataUsageStats) HasCloudTotalPhysicalUsageBytes() bool`

HasCloudTotalPhysicalUsageBytes returns a boolean if a field has been set.

### SetCloudTotalPhysicalUsageBytesNil

`func (o *ViewStatsDataUsageStats) SetCloudTotalPhysicalUsageBytesNil(b bool)`

 SetCloudTotalPhysicalUsageBytesNil sets the value for CloudTotalPhysicalUsageBytes to be an explicit nil

### UnsetCloudTotalPhysicalUsageBytes
`func (o *ViewStatsDataUsageStats) UnsetCloudTotalPhysicalUsageBytes()`

UnsetCloudTotalPhysicalUsageBytes ensures that no value is present for CloudTotalPhysicalUsageBytes, not even an explicit nil
### GetCloudTotalPhysicalUsageBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) GetCloudTotalPhysicalUsageBytesTimestampUsec() int64`

GetCloudTotalPhysicalUsageBytesTimestampUsec returns the CloudTotalPhysicalUsageBytesTimestampUsec field if non-nil, zero value otherwise.

### GetCloudTotalPhysicalUsageBytesTimestampUsecOk

`func (o *ViewStatsDataUsageStats) GetCloudTotalPhysicalUsageBytesTimestampUsecOk() (*int64, bool)`

GetCloudTotalPhysicalUsageBytesTimestampUsecOk returns a tuple with the CloudTotalPhysicalUsageBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudTotalPhysicalUsageBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) SetCloudTotalPhysicalUsageBytesTimestampUsec(v int64)`

SetCloudTotalPhysicalUsageBytesTimestampUsec sets CloudTotalPhysicalUsageBytesTimestampUsec field to given value.

### HasCloudTotalPhysicalUsageBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) HasCloudTotalPhysicalUsageBytesTimestampUsec() bool`

HasCloudTotalPhysicalUsageBytesTimestampUsec returns a boolean if a field has been set.

### SetCloudTotalPhysicalUsageBytesTimestampUsecNil

`func (o *ViewStatsDataUsageStats) SetCloudTotalPhysicalUsageBytesTimestampUsecNil(b bool)`

 SetCloudTotalPhysicalUsageBytesTimestampUsecNil sets the value for CloudTotalPhysicalUsageBytesTimestampUsec to be an explicit nil

### UnsetCloudTotalPhysicalUsageBytesTimestampUsec
`func (o *ViewStatsDataUsageStats) UnsetCloudTotalPhysicalUsageBytesTimestampUsec()`

UnsetCloudTotalPhysicalUsageBytesTimestampUsec ensures that no value is present for CloudTotalPhysicalUsageBytesTimestampUsec, not even an explicit nil
### GetDataInBytes

`func (o *ViewStatsDataUsageStats) GetDataInBytes() int64`

GetDataInBytes returns the DataInBytes field if non-nil, zero value otherwise.

### GetDataInBytesOk

`func (o *ViewStatsDataUsageStats) GetDataInBytesOk() (*int64, bool)`

GetDataInBytesOk returns a tuple with the DataInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataInBytes

`func (o *ViewStatsDataUsageStats) SetDataInBytes(v int64)`

SetDataInBytes sets DataInBytes field to given value.

### HasDataInBytes

`func (o *ViewStatsDataUsageStats) HasDataInBytes() bool`

HasDataInBytes returns a boolean if a field has been set.

### SetDataInBytesNil

`func (o *ViewStatsDataUsageStats) SetDataInBytesNil(b bool)`

 SetDataInBytesNil sets the value for DataInBytes to be an explicit nil

### UnsetDataInBytes
`func (o *ViewStatsDataUsageStats) UnsetDataInBytes()`

UnsetDataInBytes ensures that no value is present for DataInBytes, not even an explicit nil
### GetDataInBytesAfterDedup

`func (o *ViewStatsDataUsageStats) GetDataInBytesAfterDedup() int64`

GetDataInBytesAfterDedup returns the DataInBytesAfterDedup field if non-nil, zero value otherwise.

### GetDataInBytesAfterDedupOk

`func (o *ViewStatsDataUsageStats) GetDataInBytesAfterDedupOk() (*int64, bool)`

GetDataInBytesAfterDedupOk returns a tuple with the DataInBytesAfterDedup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataInBytesAfterDedup

`func (o *ViewStatsDataUsageStats) SetDataInBytesAfterDedup(v int64)`

SetDataInBytesAfterDedup sets DataInBytesAfterDedup field to given value.

### HasDataInBytesAfterDedup

`func (o *ViewStatsDataUsageStats) HasDataInBytesAfterDedup() bool`

HasDataInBytesAfterDedup returns a boolean if a field has been set.

### SetDataInBytesAfterDedupNil

`func (o *ViewStatsDataUsageStats) SetDataInBytesAfterDedupNil(b bool)`

 SetDataInBytesAfterDedupNil sets the value for DataInBytesAfterDedup to be an explicit nil

### UnsetDataInBytesAfterDedup
`func (o *ViewStatsDataUsageStats) UnsetDataInBytesAfterDedup()`

UnsetDataInBytesAfterDedup ensures that no value is present for DataInBytesAfterDedup, not even an explicit nil
### GetDataInBytesAfterDedupTimestampUsec

`func (o *ViewStatsDataUsageStats) GetDataInBytesAfterDedupTimestampUsec() int64`

GetDataInBytesAfterDedupTimestampUsec returns the DataInBytesAfterDedupTimestampUsec field if non-nil, zero value otherwise.

### GetDataInBytesAfterDedupTimestampUsecOk

`func (o *ViewStatsDataUsageStats) GetDataInBytesAfterDedupTimestampUsecOk() (*int64, bool)`

GetDataInBytesAfterDedupTimestampUsecOk returns a tuple with the DataInBytesAfterDedupTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataInBytesAfterDedupTimestampUsec

`func (o *ViewStatsDataUsageStats) SetDataInBytesAfterDedupTimestampUsec(v int64)`

SetDataInBytesAfterDedupTimestampUsec sets DataInBytesAfterDedupTimestampUsec field to given value.

### HasDataInBytesAfterDedupTimestampUsec

`func (o *ViewStatsDataUsageStats) HasDataInBytesAfterDedupTimestampUsec() bool`

HasDataInBytesAfterDedupTimestampUsec returns a boolean if a field has been set.

### SetDataInBytesAfterDedupTimestampUsecNil

`func (o *ViewStatsDataUsageStats) SetDataInBytesAfterDedupTimestampUsecNil(b bool)`

 SetDataInBytesAfterDedupTimestampUsecNil sets the value for DataInBytesAfterDedupTimestampUsec to be an explicit nil

### UnsetDataInBytesAfterDedupTimestampUsec
`func (o *ViewStatsDataUsageStats) UnsetDataInBytesAfterDedupTimestampUsec()`

UnsetDataInBytesAfterDedupTimestampUsec ensures that no value is present for DataInBytesAfterDedupTimestampUsec, not even an explicit nil
### GetDataInBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) GetDataInBytesTimestampUsec() int64`

GetDataInBytesTimestampUsec returns the DataInBytesTimestampUsec field if non-nil, zero value otherwise.

### GetDataInBytesTimestampUsecOk

`func (o *ViewStatsDataUsageStats) GetDataInBytesTimestampUsecOk() (*int64, bool)`

GetDataInBytesTimestampUsecOk returns a tuple with the DataInBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataInBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) SetDataInBytesTimestampUsec(v int64)`

SetDataInBytesTimestampUsec sets DataInBytesTimestampUsec field to given value.

### HasDataInBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) HasDataInBytesTimestampUsec() bool`

HasDataInBytesTimestampUsec returns a boolean if a field has been set.

### SetDataInBytesTimestampUsecNil

`func (o *ViewStatsDataUsageStats) SetDataInBytesTimestampUsecNil(b bool)`

 SetDataInBytesTimestampUsecNil sets the value for DataInBytesTimestampUsec to be an explicit nil

### UnsetDataInBytesTimestampUsec
`func (o *ViewStatsDataUsageStats) UnsetDataInBytesTimestampUsec()`

UnsetDataInBytesTimestampUsec ensures that no value is present for DataInBytesTimestampUsec, not even an explicit nil
### GetDataProtectLogicalUsageBytes

`func (o *ViewStatsDataUsageStats) GetDataProtectLogicalUsageBytes() int64`

GetDataProtectLogicalUsageBytes returns the DataProtectLogicalUsageBytes field if non-nil, zero value otherwise.

### GetDataProtectLogicalUsageBytesOk

`func (o *ViewStatsDataUsageStats) GetDataProtectLogicalUsageBytesOk() (*int64, bool)`

GetDataProtectLogicalUsageBytesOk returns a tuple with the DataProtectLogicalUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProtectLogicalUsageBytes

`func (o *ViewStatsDataUsageStats) SetDataProtectLogicalUsageBytes(v int64)`

SetDataProtectLogicalUsageBytes sets DataProtectLogicalUsageBytes field to given value.

### HasDataProtectLogicalUsageBytes

`func (o *ViewStatsDataUsageStats) HasDataProtectLogicalUsageBytes() bool`

HasDataProtectLogicalUsageBytes returns a boolean if a field has been set.

### SetDataProtectLogicalUsageBytesNil

`func (o *ViewStatsDataUsageStats) SetDataProtectLogicalUsageBytesNil(b bool)`

 SetDataProtectLogicalUsageBytesNil sets the value for DataProtectLogicalUsageBytes to be an explicit nil

### UnsetDataProtectLogicalUsageBytes
`func (o *ViewStatsDataUsageStats) UnsetDataProtectLogicalUsageBytes()`

UnsetDataProtectLogicalUsageBytes ensures that no value is present for DataProtectLogicalUsageBytes, not even an explicit nil
### GetDataProtectLogicalUsageBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) GetDataProtectLogicalUsageBytesTimestampUsec() int64`

GetDataProtectLogicalUsageBytesTimestampUsec returns the DataProtectLogicalUsageBytesTimestampUsec field if non-nil, zero value otherwise.

### GetDataProtectLogicalUsageBytesTimestampUsecOk

`func (o *ViewStatsDataUsageStats) GetDataProtectLogicalUsageBytesTimestampUsecOk() (*int64, bool)`

GetDataProtectLogicalUsageBytesTimestampUsecOk returns a tuple with the DataProtectLogicalUsageBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProtectLogicalUsageBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) SetDataProtectLogicalUsageBytesTimestampUsec(v int64)`

SetDataProtectLogicalUsageBytesTimestampUsec sets DataProtectLogicalUsageBytesTimestampUsec field to given value.

### HasDataProtectLogicalUsageBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) HasDataProtectLogicalUsageBytesTimestampUsec() bool`

HasDataProtectLogicalUsageBytesTimestampUsec returns a boolean if a field has been set.

### SetDataProtectLogicalUsageBytesTimestampUsecNil

`func (o *ViewStatsDataUsageStats) SetDataProtectLogicalUsageBytesTimestampUsecNil(b bool)`

 SetDataProtectLogicalUsageBytesTimestampUsecNil sets the value for DataProtectLogicalUsageBytesTimestampUsec to be an explicit nil

### UnsetDataProtectLogicalUsageBytesTimestampUsec
`func (o *ViewStatsDataUsageStats) UnsetDataProtectLogicalUsageBytesTimestampUsec()`

UnsetDataProtectLogicalUsageBytesTimestampUsec ensures that no value is present for DataProtectLogicalUsageBytesTimestampUsec, not even an explicit nil
### GetDataProtectPhysicalUsageBytes

`func (o *ViewStatsDataUsageStats) GetDataProtectPhysicalUsageBytes() int64`

GetDataProtectPhysicalUsageBytes returns the DataProtectPhysicalUsageBytes field if non-nil, zero value otherwise.

### GetDataProtectPhysicalUsageBytesOk

`func (o *ViewStatsDataUsageStats) GetDataProtectPhysicalUsageBytesOk() (*int64, bool)`

GetDataProtectPhysicalUsageBytesOk returns a tuple with the DataProtectPhysicalUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProtectPhysicalUsageBytes

`func (o *ViewStatsDataUsageStats) SetDataProtectPhysicalUsageBytes(v int64)`

SetDataProtectPhysicalUsageBytes sets DataProtectPhysicalUsageBytes field to given value.

### HasDataProtectPhysicalUsageBytes

`func (o *ViewStatsDataUsageStats) HasDataProtectPhysicalUsageBytes() bool`

HasDataProtectPhysicalUsageBytes returns a boolean if a field has been set.

### SetDataProtectPhysicalUsageBytesNil

`func (o *ViewStatsDataUsageStats) SetDataProtectPhysicalUsageBytesNil(b bool)`

 SetDataProtectPhysicalUsageBytesNil sets the value for DataProtectPhysicalUsageBytes to be an explicit nil

### UnsetDataProtectPhysicalUsageBytes
`func (o *ViewStatsDataUsageStats) UnsetDataProtectPhysicalUsageBytes()`

UnsetDataProtectPhysicalUsageBytes ensures that no value is present for DataProtectPhysicalUsageBytes, not even an explicit nil
### GetDataProtectPhysicalUsageBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) GetDataProtectPhysicalUsageBytesTimestampUsec() int64`

GetDataProtectPhysicalUsageBytesTimestampUsec returns the DataProtectPhysicalUsageBytesTimestampUsec field if non-nil, zero value otherwise.

### GetDataProtectPhysicalUsageBytesTimestampUsecOk

`func (o *ViewStatsDataUsageStats) GetDataProtectPhysicalUsageBytesTimestampUsecOk() (*int64, bool)`

GetDataProtectPhysicalUsageBytesTimestampUsecOk returns a tuple with the DataProtectPhysicalUsageBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProtectPhysicalUsageBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) SetDataProtectPhysicalUsageBytesTimestampUsec(v int64)`

SetDataProtectPhysicalUsageBytesTimestampUsec sets DataProtectPhysicalUsageBytesTimestampUsec field to given value.

### HasDataProtectPhysicalUsageBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) HasDataProtectPhysicalUsageBytesTimestampUsec() bool`

HasDataProtectPhysicalUsageBytesTimestampUsec returns a boolean if a field has been set.

### SetDataProtectPhysicalUsageBytesTimestampUsecNil

`func (o *ViewStatsDataUsageStats) SetDataProtectPhysicalUsageBytesTimestampUsecNil(b bool)`

 SetDataProtectPhysicalUsageBytesTimestampUsecNil sets the value for DataProtectPhysicalUsageBytesTimestampUsec to be an explicit nil

### UnsetDataProtectPhysicalUsageBytesTimestampUsec
`func (o *ViewStatsDataUsageStats) UnsetDataProtectPhysicalUsageBytesTimestampUsec()`

UnsetDataProtectPhysicalUsageBytesTimestampUsec ensures that no value is present for DataProtectPhysicalUsageBytesTimestampUsec, not even an explicit nil
### GetDataWrittenBytes

`func (o *ViewStatsDataUsageStats) GetDataWrittenBytes() int64`

GetDataWrittenBytes returns the DataWrittenBytes field if non-nil, zero value otherwise.

### GetDataWrittenBytesOk

`func (o *ViewStatsDataUsageStats) GetDataWrittenBytesOk() (*int64, bool)`

GetDataWrittenBytesOk returns a tuple with the DataWrittenBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWrittenBytes

`func (o *ViewStatsDataUsageStats) SetDataWrittenBytes(v int64)`

SetDataWrittenBytes sets DataWrittenBytes field to given value.

### HasDataWrittenBytes

`func (o *ViewStatsDataUsageStats) HasDataWrittenBytes() bool`

HasDataWrittenBytes returns a boolean if a field has been set.

### SetDataWrittenBytesNil

`func (o *ViewStatsDataUsageStats) SetDataWrittenBytesNil(b bool)`

 SetDataWrittenBytesNil sets the value for DataWrittenBytes to be an explicit nil

### UnsetDataWrittenBytes
`func (o *ViewStatsDataUsageStats) UnsetDataWrittenBytes()`

UnsetDataWrittenBytes ensures that no value is present for DataWrittenBytes, not even an explicit nil
### GetDataWrittenBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) GetDataWrittenBytesTimestampUsec() int64`

GetDataWrittenBytesTimestampUsec returns the DataWrittenBytesTimestampUsec field if non-nil, zero value otherwise.

### GetDataWrittenBytesTimestampUsecOk

`func (o *ViewStatsDataUsageStats) GetDataWrittenBytesTimestampUsecOk() (*int64, bool)`

GetDataWrittenBytesTimestampUsecOk returns a tuple with the DataWrittenBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWrittenBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) SetDataWrittenBytesTimestampUsec(v int64)`

SetDataWrittenBytesTimestampUsec sets DataWrittenBytesTimestampUsec field to given value.

### HasDataWrittenBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) HasDataWrittenBytesTimestampUsec() bool`

HasDataWrittenBytesTimestampUsec returns a boolean if a field has been set.

### SetDataWrittenBytesTimestampUsecNil

`func (o *ViewStatsDataUsageStats) SetDataWrittenBytesTimestampUsecNil(b bool)`

 SetDataWrittenBytesTimestampUsecNil sets the value for DataWrittenBytesTimestampUsec to be an explicit nil

### UnsetDataWrittenBytesTimestampUsec
`func (o *ViewStatsDataUsageStats) UnsetDataWrittenBytesTimestampUsec()`

UnsetDataWrittenBytesTimestampUsec ensures that no value is present for DataWrittenBytesTimestampUsec, not even an explicit nil
### GetFileServicesLogicalUsageBytes

`func (o *ViewStatsDataUsageStats) GetFileServicesLogicalUsageBytes() int64`

GetFileServicesLogicalUsageBytes returns the FileServicesLogicalUsageBytes field if non-nil, zero value otherwise.

### GetFileServicesLogicalUsageBytesOk

`func (o *ViewStatsDataUsageStats) GetFileServicesLogicalUsageBytesOk() (*int64, bool)`

GetFileServicesLogicalUsageBytesOk returns a tuple with the FileServicesLogicalUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileServicesLogicalUsageBytes

`func (o *ViewStatsDataUsageStats) SetFileServicesLogicalUsageBytes(v int64)`

SetFileServicesLogicalUsageBytes sets FileServicesLogicalUsageBytes field to given value.

### HasFileServicesLogicalUsageBytes

`func (o *ViewStatsDataUsageStats) HasFileServicesLogicalUsageBytes() bool`

HasFileServicesLogicalUsageBytes returns a boolean if a field has been set.

### SetFileServicesLogicalUsageBytesNil

`func (o *ViewStatsDataUsageStats) SetFileServicesLogicalUsageBytesNil(b bool)`

 SetFileServicesLogicalUsageBytesNil sets the value for FileServicesLogicalUsageBytes to be an explicit nil

### UnsetFileServicesLogicalUsageBytes
`func (o *ViewStatsDataUsageStats) UnsetFileServicesLogicalUsageBytes()`

UnsetFileServicesLogicalUsageBytes ensures that no value is present for FileServicesLogicalUsageBytes, not even an explicit nil
### GetFileServicesLogicalUsageBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) GetFileServicesLogicalUsageBytesTimestampUsec() int64`

GetFileServicesLogicalUsageBytesTimestampUsec returns the FileServicesLogicalUsageBytesTimestampUsec field if non-nil, zero value otherwise.

### GetFileServicesLogicalUsageBytesTimestampUsecOk

`func (o *ViewStatsDataUsageStats) GetFileServicesLogicalUsageBytesTimestampUsecOk() (*int64, bool)`

GetFileServicesLogicalUsageBytesTimestampUsecOk returns a tuple with the FileServicesLogicalUsageBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileServicesLogicalUsageBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) SetFileServicesLogicalUsageBytesTimestampUsec(v int64)`

SetFileServicesLogicalUsageBytesTimestampUsec sets FileServicesLogicalUsageBytesTimestampUsec field to given value.

### HasFileServicesLogicalUsageBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) HasFileServicesLogicalUsageBytesTimestampUsec() bool`

HasFileServicesLogicalUsageBytesTimestampUsec returns a boolean if a field has been set.

### SetFileServicesLogicalUsageBytesTimestampUsecNil

`func (o *ViewStatsDataUsageStats) SetFileServicesLogicalUsageBytesTimestampUsecNil(b bool)`

 SetFileServicesLogicalUsageBytesTimestampUsecNil sets the value for FileServicesLogicalUsageBytesTimestampUsec to be an explicit nil

### UnsetFileServicesLogicalUsageBytesTimestampUsec
`func (o *ViewStatsDataUsageStats) UnsetFileServicesLogicalUsageBytesTimestampUsec()`

UnsetFileServicesLogicalUsageBytesTimestampUsec ensures that no value is present for FileServicesLogicalUsageBytesTimestampUsec, not even an explicit nil
### GetFileServicesPhysicalUsageBytes

`func (o *ViewStatsDataUsageStats) GetFileServicesPhysicalUsageBytes() int64`

GetFileServicesPhysicalUsageBytes returns the FileServicesPhysicalUsageBytes field if non-nil, zero value otherwise.

### GetFileServicesPhysicalUsageBytesOk

`func (o *ViewStatsDataUsageStats) GetFileServicesPhysicalUsageBytesOk() (*int64, bool)`

GetFileServicesPhysicalUsageBytesOk returns a tuple with the FileServicesPhysicalUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileServicesPhysicalUsageBytes

`func (o *ViewStatsDataUsageStats) SetFileServicesPhysicalUsageBytes(v int64)`

SetFileServicesPhysicalUsageBytes sets FileServicesPhysicalUsageBytes field to given value.

### HasFileServicesPhysicalUsageBytes

`func (o *ViewStatsDataUsageStats) HasFileServicesPhysicalUsageBytes() bool`

HasFileServicesPhysicalUsageBytes returns a boolean if a field has been set.

### SetFileServicesPhysicalUsageBytesNil

`func (o *ViewStatsDataUsageStats) SetFileServicesPhysicalUsageBytesNil(b bool)`

 SetFileServicesPhysicalUsageBytesNil sets the value for FileServicesPhysicalUsageBytes to be an explicit nil

### UnsetFileServicesPhysicalUsageBytes
`func (o *ViewStatsDataUsageStats) UnsetFileServicesPhysicalUsageBytes()`

UnsetFileServicesPhysicalUsageBytes ensures that no value is present for FileServicesPhysicalUsageBytes, not even an explicit nil
### GetFileServicesPhysicalUsageBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) GetFileServicesPhysicalUsageBytesTimestampUsec() int64`

GetFileServicesPhysicalUsageBytesTimestampUsec returns the FileServicesPhysicalUsageBytesTimestampUsec field if non-nil, zero value otherwise.

### GetFileServicesPhysicalUsageBytesTimestampUsecOk

`func (o *ViewStatsDataUsageStats) GetFileServicesPhysicalUsageBytesTimestampUsecOk() (*int64, bool)`

GetFileServicesPhysicalUsageBytesTimestampUsecOk returns a tuple with the FileServicesPhysicalUsageBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileServicesPhysicalUsageBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) SetFileServicesPhysicalUsageBytesTimestampUsec(v int64)`

SetFileServicesPhysicalUsageBytesTimestampUsec sets FileServicesPhysicalUsageBytesTimestampUsec field to given value.

### HasFileServicesPhysicalUsageBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) HasFileServicesPhysicalUsageBytesTimestampUsec() bool`

HasFileServicesPhysicalUsageBytesTimestampUsec returns a boolean if a field has been set.

### SetFileServicesPhysicalUsageBytesTimestampUsecNil

`func (o *ViewStatsDataUsageStats) SetFileServicesPhysicalUsageBytesTimestampUsecNil(b bool)`

 SetFileServicesPhysicalUsageBytesTimestampUsecNil sets the value for FileServicesPhysicalUsageBytesTimestampUsec to be an explicit nil

### UnsetFileServicesPhysicalUsageBytesTimestampUsec
`func (o *ViewStatsDataUsageStats) UnsetFileServicesPhysicalUsageBytesTimestampUsec()`

UnsetFileServicesPhysicalUsageBytesTimestampUsec ensures that no value is present for FileServicesPhysicalUsageBytesTimestampUsec, not even an explicit nil
### GetLocalDataWrittenBytes

`func (o *ViewStatsDataUsageStats) GetLocalDataWrittenBytes() int64`

GetLocalDataWrittenBytes returns the LocalDataWrittenBytes field if non-nil, zero value otherwise.

### GetLocalDataWrittenBytesOk

`func (o *ViewStatsDataUsageStats) GetLocalDataWrittenBytesOk() (*int64, bool)`

GetLocalDataWrittenBytesOk returns a tuple with the LocalDataWrittenBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDataWrittenBytes

`func (o *ViewStatsDataUsageStats) SetLocalDataWrittenBytes(v int64)`

SetLocalDataWrittenBytes sets LocalDataWrittenBytes field to given value.

### HasLocalDataWrittenBytes

`func (o *ViewStatsDataUsageStats) HasLocalDataWrittenBytes() bool`

HasLocalDataWrittenBytes returns a boolean if a field has been set.

### SetLocalDataWrittenBytesNil

`func (o *ViewStatsDataUsageStats) SetLocalDataWrittenBytesNil(b bool)`

 SetLocalDataWrittenBytesNil sets the value for LocalDataWrittenBytes to be an explicit nil

### UnsetLocalDataWrittenBytes
`func (o *ViewStatsDataUsageStats) UnsetLocalDataWrittenBytes()`

UnsetLocalDataWrittenBytes ensures that no value is present for LocalDataWrittenBytes, not even an explicit nil
### GetLocalDataWrittenBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) GetLocalDataWrittenBytesTimestampUsec() int64`

GetLocalDataWrittenBytesTimestampUsec returns the LocalDataWrittenBytesTimestampUsec field if non-nil, zero value otherwise.

### GetLocalDataWrittenBytesTimestampUsecOk

`func (o *ViewStatsDataUsageStats) GetLocalDataWrittenBytesTimestampUsecOk() (*int64, bool)`

GetLocalDataWrittenBytesTimestampUsecOk returns a tuple with the LocalDataWrittenBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDataWrittenBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) SetLocalDataWrittenBytesTimestampUsec(v int64)`

SetLocalDataWrittenBytesTimestampUsec sets LocalDataWrittenBytesTimestampUsec field to given value.

### HasLocalDataWrittenBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) HasLocalDataWrittenBytesTimestampUsec() bool`

HasLocalDataWrittenBytesTimestampUsec returns a boolean if a field has been set.

### SetLocalDataWrittenBytesTimestampUsecNil

`func (o *ViewStatsDataUsageStats) SetLocalDataWrittenBytesTimestampUsecNil(b bool)`

 SetLocalDataWrittenBytesTimestampUsecNil sets the value for LocalDataWrittenBytesTimestampUsec to be an explicit nil

### UnsetLocalDataWrittenBytesTimestampUsec
`func (o *ViewStatsDataUsageStats) UnsetLocalDataWrittenBytesTimestampUsec()`

UnsetLocalDataWrittenBytesTimestampUsec ensures that no value is present for LocalDataWrittenBytesTimestampUsec, not even an explicit nil
### GetLocalTierResiliencyImpactBytes

`func (o *ViewStatsDataUsageStats) GetLocalTierResiliencyImpactBytes() int64`

GetLocalTierResiliencyImpactBytes returns the LocalTierResiliencyImpactBytes field if non-nil, zero value otherwise.

### GetLocalTierResiliencyImpactBytesOk

`func (o *ViewStatsDataUsageStats) GetLocalTierResiliencyImpactBytesOk() (*int64, bool)`

GetLocalTierResiliencyImpactBytesOk returns a tuple with the LocalTierResiliencyImpactBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTierResiliencyImpactBytes

`func (o *ViewStatsDataUsageStats) SetLocalTierResiliencyImpactBytes(v int64)`

SetLocalTierResiliencyImpactBytes sets LocalTierResiliencyImpactBytes field to given value.

### HasLocalTierResiliencyImpactBytes

`func (o *ViewStatsDataUsageStats) HasLocalTierResiliencyImpactBytes() bool`

HasLocalTierResiliencyImpactBytes returns a boolean if a field has been set.

### SetLocalTierResiliencyImpactBytesNil

`func (o *ViewStatsDataUsageStats) SetLocalTierResiliencyImpactBytesNil(b bool)`

 SetLocalTierResiliencyImpactBytesNil sets the value for LocalTierResiliencyImpactBytes to be an explicit nil

### UnsetLocalTierResiliencyImpactBytes
`func (o *ViewStatsDataUsageStats) UnsetLocalTierResiliencyImpactBytes()`

UnsetLocalTierResiliencyImpactBytes ensures that no value is present for LocalTierResiliencyImpactBytes, not even an explicit nil
### GetLocalTierResiliencyImpactBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) GetLocalTierResiliencyImpactBytesTimestampUsec() int64`

GetLocalTierResiliencyImpactBytesTimestampUsec returns the LocalTierResiliencyImpactBytesTimestampUsec field if non-nil, zero value otherwise.

### GetLocalTierResiliencyImpactBytesTimestampUsecOk

`func (o *ViewStatsDataUsageStats) GetLocalTierResiliencyImpactBytesTimestampUsecOk() (*int64, bool)`

GetLocalTierResiliencyImpactBytesTimestampUsecOk returns a tuple with the LocalTierResiliencyImpactBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTierResiliencyImpactBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) SetLocalTierResiliencyImpactBytesTimestampUsec(v int64)`

SetLocalTierResiliencyImpactBytesTimestampUsec sets LocalTierResiliencyImpactBytesTimestampUsec field to given value.

### HasLocalTierResiliencyImpactBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) HasLocalTierResiliencyImpactBytesTimestampUsec() bool`

HasLocalTierResiliencyImpactBytesTimestampUsec returns a boolean if a field has been set.

### SetLocalTierResiliencyImpactBytesTimestampUsecNil

`func (o *ViewStatsDataUsageStats) SetLocalTierResiliencyImpactBytesTimestampUsecNil(b bool)`

 SetLocalTierResiliencyImpactBytesTimestampUsecNil sets the value for LocalTierResiliencyImpactBytesTimestampUsec to be an explicit nil

### UnsetLocalTierResiliencyImpactBytesTimestampUsec
`func (o *ViewStatsDataUsageStats) UnsetLocalTierResiliencyImpactBytesTimestampUsec()`

UnsetLocalTierResiliencyImpactBytesTimestampUsec ensures that no value is present for LocalTierResiliencyImpactBytesTimestampUsec, not even an explicit nil
### GetLocalTotalPhysicalUsageBytes

`func (o *ViewStatsDataUsageStats) GetLocalTotalPhysicalUsageBytes() int64`

GetLocalTotalPhysicalUsageBytes returns the LocalTotalPhysicalUsageBytes field if non-nil, zero value otherwise.

### GetLocalTotalPhysicalUsageBytesOk

`func (o *ViewStatsDataUsageStats) GetLocalTotalPhysicalUsageBytesOk() (*int64, bool)`

GetLocalTotalPhysicalUsageBytesOk returns a tuple with the LocalTotalPhysicalUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTotalPhysicalUsageBytes

`func (o *ViewStatsDataUsageStats) SetLocalTotalPhysicalUsageBytes(v int64)`

SetLocalTotalPhysicalUsageBytes sets LocalTotalPhysicalUsageBytes field to given value.

### HasLocalTotalPhysicalUsageBytes

`func (o *ViewStatsDataUsageStats) HasLocalTotalPhysicalUsageBytes() bool`

HasLocalTotalPhysicalUsageBytes returns a boolean if a field has been set.

### SetLocalTotalPhysicalUsageBytesNil

`func (o *ViewStatsDataUsageStats) SetLocalTotalPhysicalUsageBytesNil(b bool)`

 SetLocalTotalPhysicalUsageBytesNil sets the value for LocalTotalPhysicalUsageBytes to be an explicit nil

### UnsetLocalTotalPhysicalUsageBytes
`func (o *ViewStatsDataUsageStats) UnsetLocalTotalPhysicalUsageBytes()`

UnsetLocalTotalPhysicalUsageBytes ensures that no value is present for LocalTotalPhysicalUsageBytes, not even an explicit nil
### GetLocalTotalPhysicalUsageBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) GetLocalTotalPhysicalUsageBytesTimestampUsec() int64`

GetLocalTotalPhysicalUsageBytesTimestampUsec returns the LocalTotalPhysicalUsageBytesTimestampUsec field if non-nil, zero value otherwise.

### GetLocalTotalPhysicalUsageBytesTimestampUsecOk

`func (o *ViewStatsDataUsageStats) GetLocalTotalPhysicalUsageBytesTimestampUsecOk() (*int64, bool)`

GetLocalTotalPhysicalUsageBytesTimestampUsecOk returns a tuple with the LocalTotalPhysicalUsageBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTotalPhysicalUsageBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) SetLocalTotalPhysicalUsageBytesTimestampUsec(v int64)`

SetLocalTotalPhysicalUsageBytesTimestampUsec sets LocalTotalPhysicalUsageBytesTimestampUsec field to given value.

### HasLocalTotalPhysicalUsageBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) HasLocalTotalPhysicalUsageBytesTimestampUsec() bool`

HasLocalTotalPhysicalUsageBytesTimestampUsec returns a boolean if a field has been set.

### SetLocalTotalPhysicalUsageBytesTimestampUsecNil

`func (o *ViewStatsDataUsageStats) SetLocalTotalPhysicalUsageBytesTimestampUsecNil(b bool)`

 SetLocalTotalPhysicalUsageBytesTimestampUsecNil sets the value for LocalTotalPhysicalUsageBytesTimestampUsec to be an explicit nil

### UnsetLocalTotalPhysicalUsageBytesTimestampUsec
`func (o *ViewStatsDataUsageStats) UnsetLocalTotalPhysicalUsageBytesTimestampUsec()`

UnsetLocalTotalPhysicalUsageBytesTimestampUsec ensures that no value is present for LocalTotalPhysicalUsageBytesTimestampUsec, not even an explicit nil
### GetNumDirectories

`func (o *ViewStatsDataUsageStats) GetNumDirectories() int64`

GetNumDirectories returns the NumDirectories field if non-nil, zero value otherwise.

### GetNumDirectoriesOk

`func (o *ViewStatsDataUsageStats) GetNumDirectoriesOk() (*int64, bool)`

GetNumDirectoriesOk returns a tuple with the NumDirectories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDirectories

`func (o *ViewStatsDataUsageStats) SetNumDirectories(v int64)`

SetNumDirectories sets NumDirectories field to given value.

### HasNumDirectories

`func (o *ViewStatsDataUsageStats) HasNumDirectories() bool`

HasNumDirectories returns a boolean if a field has been set.

### SetNumDirectoriesNil

`func (o *ViewStatsDataUsageStats) SetNumDirectoriesNil(b bool)`

 SetNumDirectoriesNil sets the value for NumDirectories to be an explicit nil

### UnsetNumDirectories
`func (o *ViewStatsDataUsageStats) UnsetNumDirectories()`

UnsetNumDirectories ensures that no value is present for NumDirectories, not even an explicit nil
### GetNumFiles

`func (o *ViewStatsDataUsageStats) GetNumFiles() int64`

GetNumFiles returns the NumFiles field if non-nil, zero value otherwise.

### GetNumFilesOk

`func (o *ViewStatsDataUsageStats) GetNumFilesOk() (*int64, bool)`

GetNumFilesOk returns a tuple with the NumFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFiles

`func (o *ViewStatsDataUsageStats) SetNumFiles(v int64)`

SetNumFiles sets NumFiles field to given value.

### HasNumFiles

`func (o *ViewStatsDataUsageStats) HasNumFiles() bool`

HasNumFiles returns a boolean if a field has been set.

### SetNumFilesNil

`func (o *ViewStatsDataUsageStats) SetNumFilesNil(b bool)`

 SetNumFilesNil sets the value for NumFiles to be an explicit nil

### UnsetNumFiles
`func (o *ViewStatsDataUsageStats) UnsetNumFiles()`

UnsetNumFiles ensures that no value is present for NumFiles, not even an explicit nil
### GetOutdatedLogicalUsageBytes

`func (o *ViewStatsDataUsageStats) GetOutdatedLogicalUsageBytes() int64`

GetOutdatedLogicalUsageBytes returns the OutdatedLogicalUsageBytes field if non-nil, zero value otherwise.

### GetOutdatedLogicalUsageBytesOk

`func (o *ViewStatsDataUsageStats) GetOutdatedLogicalUsageBytesOk() (*int64, bool)`

GetOutdatedLogicalUsageBytesOk returns a tuple with the OutdatedLogicalUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutdatedLogicalUsageBytes

`func (o *ViewStatsDataUsageStats) SetOutdatedLogicalUsageBytes(v int64)`

SetOutdatedLogicalUsageBytes sets OutdatedLogicalUsageBytes field to given value.

### HasOutdatedLogicalUsageBytes

`func (o *ViewStatsDataUsageStats) HasOutdatedLogicalUsageBytes() bool`

HasOutdatedLogicalUsageBytes returns a boolean if a field has been set.

### SetOutdatedLogicalUsageBytesNil

`func (o *ViewStatsDataUsageStats) SetOutdatedLogicalUsageBytesNil(b bool)`

 SetOutdatedLogicalUsageBytesNil sets the value for OutdatedLogicalUsageBytes to be an explicit nil

### UnsetOutdatedLogicalUsageBytes
`func (o *ViewStatsDataUsageStats) UnsetOutdatedLogicalUsageBytes()`

UnsetOutdatedLogicalUsageBytes ensures that no value is present for OutdatedLogicalUsageBytes, not even an explicit nil
### GetOutdatedLogicalUsageBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) GetOutdatedLogicalUsageBytesTimestampUsec() int64`

GetOutdatedLogicalUsageBytesTimestampUsec returns the OutdatedLogicalUsageBytesTimestampUsec field if non-nil, zero value otherwise.

### GetOutdatedLogicalUsageBytesTimestampUsecOk

`func (o *ViewStatsDataUsageStats) GetOutdatedLogicalUsageBytesTimestampUsecOk() (*int64, bool)`

GetOutdatedLogicalUsageBytesTimestampUsecOk returns a tuple with the OutdatedLogicalUsageBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutdatedLogicalUsageBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) SetOutdatedLogicalUsageBytesTimestampUsec(v int64)`

SetOutdatedLogicalUsageBytesTimestampUsec sets OutdatedLogicalUsageBytesTimestampUsec field to given value.

### HasOutdatedLogicalUsageBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) HasOutdatedLogicalUsageBytesTimestampUsec() bool`

HasOutdatedLogicalUsageBytesTimestampUsec returns a boolean if a field has been set.

### SetOutdatedLogicalUsageBytesTimestampUsecNil

`func (o *ViewStatsDataUsageStats) SetOutdatedLogicalUsageBytesTimestampUsecNil(b bool)`

 SetOutdatedLogicalUsageBytesTimestampUsecNil sets the value for OutdatedLogicalUsageBytesTimestampUsec to be an explicit nil

### UnsetOutdatedLogicalUsageBytesTimestampUsec
`func (o *ViewStatsDataUsageStats) UnsetOutdatedLogicalUsageBytesTimestampUsec()`

UnsetOutdatedLogicalUsageBytesTimestampUsec ensures that no value is present for OutdatedLogicalUsageBytesTimestampUsec, not even an explicit nil
### GetStorageConsumedBytes

`func (o *ViewStatsDataUsageStats) GetStorageConsumedBytes() int64`

GetStorageConsumedBytes returns the StorageConsumedBytes field if non-nil, zero value otherwise.

### GetStorageConsumedBytesOk

`func (o *ViewStatsDataUsageStats) GetStorageConsumedBytesOk() (*int64, bool)`

GetStorageConsumedBytesOk returns a tuple with the StorageConsumedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConsumedBytes

`func (o *ViewStatsDataUsageStats) SetStorageConsumedBytes(v int64)`

SetStorageConsumedBytes sets StorageConsumedBytes field to given value.

### HasStorageConsumedBytes

`func (o *ViewStatsDataUsageStats) HasStorageConsumedBytes() bool`

HasStorageConsumedBytes returns a boolean if a field has been set.

### SetStorageConsumedBytesNil

`func (o *ViewStatsDataUsageStats) SetStorageConsumedBytesNil(b bool)`

 SetStorageConsumedBytesNil sets the value for StorageConsumedBytes to be an explicit nil

### UnsetStorageConsumedBytes
`func (o *ViewStatsDataUsageStats) UnsetStorageConsumedBytes()`

UnsetStorageConsumedBytes ensures that no value is present for StorageConsumedBytes, not even an explicit nil
### GetStorageConsumedBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) GetStorageConsumedBytesTimestampUsec() int64`

GetStorageConsumedBytesTimestampUsec returns the StorageConsumedBytesTimestampUsec field if non-nil, zero value otherwise.

### GetStorageConsumedBytesTimestampUsecOk

`func (o *ViewStatsDataUsageStats) GetStorageConsumedBytesTimestampUsecOk() (*int64, bool)`

GetStorageConsumedBytesTimestampUsecOk returns a tuple with the StorageConsumedBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConsumedBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) SetStorageConsumedBytesTimestampUsec(v int64)`

SetStorageConsumedBytesTimestampUsec sets StorageConsumedBytesTimestampUsec field to given value.

### HasStorageConsumedBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) HasStorageConsumedBytesTimestampUsec() bool`

HasStorageConsumedBytesTimestampUsec returns a boolean if a field has been set.

### SetStorageConsumedBytesTimestampUsecNil

`func (o *ViewStatsDataUsageStats) SetStorageConsumedBytesTimestampUsecNil(b bool)`

 SetStorageConsumedBytesTimestampUsecNil sets the value for StorageConsumedBytesTimestampUsec to be an explicit nil

### UnsetStorageConsumedBytesTimestampUsec
`func (o *ViewStatsDataUsageStats) UnsetStorageConsumedBytesTimestampUsec()`

UnsetStorageConsumedBytesTimestampUsec ensures that no value is present for StorageConsumedBytesTimestampUsec, not even an explicit nil
### GetTotalLogicalUsageBytes

`func (o *ViewStatsDataUsageStats) GetTotalLogicalUsageBytes() int64`

GetTotalLogicalUsageBytes returns the TotalLogicalUsageBytes field if non-nil, zero value otherwise.

### GetTotalLogicalUsageBytesOk

`func (o *ViewStatsDataUsageStats) GetTotalLogicalUsageBytesOk() (*int64, bool)`

GetTotalLogicalUsageBytesOk returns a tuple with the TotalLogicalUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLogicalUsageBytes

`func (o *ViewStatsDataUsageStats) SetTotalLogicalUsageBytes(v int64)`

SetTotalLogicalUsageBytes sets TotalLogicalUsageBytes field to given value.

### HasTotalLogicalUsageBytes

`func (o *ViewStatsDataUsageStats) HasTotalLogicalUsageBytes() bool`

HasTotalLogicalUsageBytes returns a boolean if a field has been set.

### SetTotalLogicalUsageBytesNil

`func (o *ViewStatsDataUsageStats) SetTotalLogicalUsageBytesNil(b bool)`

 SetTotalLogicalUsageBytesNil sets the value for TotalLogicalUsageBytes to be an explicit nil

### UnsetTotalLogicalUsageBytes
`func (o *ViewStatsDataUsageStats) UnsetTotalLogicalUsageBytes()`

UnsetTotalLogicalUsageBytes ensures that no value is present for TotalLogicalUsageBytes, not even an explicit nil
### GetTotalLogicalUsageBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) GetTotalLogicalUsageBytesTimestampUsec() int64`

GetTotalLogicalUsageBytesTimestampUsec returns the TotalLogicalUsageBytesTimestampUsec field if non-nil, zero value otherwise.

### GetTotalLogicalUsageBytesTimestampUsecOk

`func (o *ViewStatsDataUsageStats) GetTotalLogicalUsageBytesTimestampUsecOk() (*int64, bool)`

GetTotalLogicalUsageBytesTimestampUsecOk returns a tuple with the TotalLogicalUsageBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLogicalUsageBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) SetTotalLogicalUsageBytesTimestampUsec(v int64)`

SetTotalLogicalUsageBytesTimestampUsec sets TotalLogicalUsageBytesTimestampUsec field to given value.

### HasTotalLogicalUsageBytesTimestampUsec

`func (o *ViewStatsDataUsageStats) HasTotalLogicalUsageBytesTimestampUsec() bool`

HasTotalLogicalUsageBytesTimestampUsec returns a boolean if a field has been set.

### SetTotalLogicalUsageBytesTimestampUsecNil

`func (o *ViewStatsDataUsageStats) SetTotalLogicalUsageBytesTimestampUsecNil(b bool)`

 SetTotalLogicalUsageBytesTimestampUsecNil sets the value for TotalLogicalUsageBytesTimestampUsec to be an explicit nil

### UnsetTotalLogicalUsageBytesTimestampUsec
`func (o *ViewStatsDataUsageStats) UnsetTotalLogicalUsageBytesTimestampUsec()`

UnsetTotalLogicalUsageBytesTimestampUsec ensures that no value is present for TotalLogicalUsageBytesTimestampUsec, not even an explicit nil
### GetUniquePhysicalDataBytes

`func (o *ViewStatsDataUsageStats) GetUniquePhysicalDataBytes() int64`

GetUniquePhysicalDataBytes returns the UniquePhysicalDataBytes field if non-nil, zero value otherwise.

### GetUniquePhysicalDataBytesOk

`func (o *ViewStatsDataUsageStats) GetUniquePhysicalDataBytesOk() (*int64, bool)`

GetUniquePhysicalDataBytesOk returns a tuple with the UniquePhysicalDataBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniquePhysicalDataBytes

`func (o *ViewStatsDataUsageStats) SetUniquePhysicalDataBytes(v int64)`

SetUniquePhysicalDataBytes sets UniquePhysicalDataBytes field to given value.

### HasUniquePhysicalDataBytes

`func (o *ViewStatsDataUsageStats) HasUniquePhysicalDataBytes() bool`

HasUniquePhysicalDataBytes returns a boolean if a field has been set.

### SetUniquePhysicalDataBytesNil

`func (o *ViewStatsDataUsageStats) SetUniquePhysicalDataBytesNil(b bool)`

 SetUniquePhysicalDataBytesNil sets the value for UniquePhysicalDataBytes to be an explicit nil

### UnsetUniquePhysicalDataBytes
`func (o *ViewStatsDataUsageStats) UnsetUniquePhysicalDataBytes()`

UnsetUniquePhysicalDataBytes ensures that no value is present for UniquePhysicalDataBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


