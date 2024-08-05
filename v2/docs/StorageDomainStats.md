# StorageDomainStats

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

### NewStorageDomainStats

`func NewStorageDomainStats() *StorageDomainStats`

NewStorageDomainStats instantiates a new StorageDomainStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageDomainStatsWithDefaults

`func NewStorageDomainStatsWithDefaults() *StorageDomainStats`

NewStorageDomainStatsWithDefaults instantiates a new StorageDomainStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudDataWrittenBytes

`func (o *StorageDomainStats) GetCloudDataWrittenBytes() int64`

GetCloudDataWrittenBytes returns the CloudDataWrittenBytes field if non-nil, zero value otherwise.

### GetCloudDataWrittenBytesOk

`func (o *StorageDomainStats) GetCloudDataWrittenBytesOk() (*int64, bool)`

GetCloudDataWrittenBytesOk returns a tuple with the CloudDataWrittenBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDataWrittenBytes

`func (o *StorageDomainStats) SetCloudDataWrittenBytes(v int64)`

SetCloudDataWrittenBytes sets CloudDataWrittenBytes field to given value.

### HasCloudDataWrittenBytes

`func (o *StorageDomainStats) HasCloudDataWrittenBytes() bool`

HasCloudDataWrittenBytes returns a boolean if a field has been set.

### SetCloudDataWrittenBytesNil

`func (o *StorageDomainStats) SetCloudDataWrittenBytesNil(b bool)`

 SetCloudDataWrittenBytesNil sets the value for CloudDataWrittenBytes to be an explicit nil

### UnsetCloudDataWrittenBytes
`func (o *StorageDomainStats) UnsetCloudDataWrittenBytes()`

UnsetCloudDataWrittenBytes ensures that no value is present for CloudDataWrittenBytes, not even an explicit nil
### GetCloudDataWrittenBytesTimestampUsec

`func (o *StorageDomainStats) GetCloudDataWrittenBytesTimestampUsec() int64`

GetCloudDataWrittenBytesTimestampUsec returns the CloudDataWrittenBytesTimestampUsec field if non-nil, zero value otherwise.

### GetCloudDataWrittenBytesTimestampUsecOk

`func (o *StorageDomainStats) GetCloudDataWrittenBytesTimestampUsecOk() (*int64, bool)`

GetCloudDataWrittenBytesTimestampUsecOk returns a tuple with the CloudDataWrittenBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDataWrittenBytesTimestampUsec

`func (o *StorageDomainStats) SetCloudDataWrittenBytesTimestampUsec(v int64)`

SetCloudDataWrittenBytesTimestampUsec sets CloudDataWrittenBytesTimestampUsec field to given value.

### HasCloudDataWrittenBytesTimestampUsec

`func (o *StorageDomainStats) HasCloudDataWrittenBytesTimestampUsec() bool`

HasCloudDataWrittenBytesTimestampUsec returns a boolean if a field has been set.

### SetCloudDataWrittenBytesTimestampUsecNil

`func (o *StorageDomainStats) SetCloudDataWrittenBytesTimestampUsecNil(b bool)`

 SetCloudDataWrittenBytesTimestampUsecNil sets the value for CloudDataWrittenBytesTimestampUsec to be an explicit nil

### UnsetCloudDataWrittenBytesTimestampUsec
`func (o *StorageDomainStats) UnsetCloudDataWrittenBytesTimestampUsec()`

UnsetCloudDataWrittenBytesTimestampUsec ensures that no value is present for CloudDataWrittenBytesTimestampUsec, not even an explicit nil
### GetCloudTotalPhysicalUsageBytes

`func (o *StorageDomainStats) GetCloudTotalPhysicalUsageBytes() int64`

GetCloudTotalPhysicalUsageBytes returns the CloudTotalPhysicalUsageBytes field if non-nil, zero value otherwise.

### GetCloudTotalPhysicalUsageBytesOk

`func (o *StorageDomainStats) GetCloudTotalPhysicalUsageBytesOk() (*int64, bool)`

GetCloudTotalPhysicalUsageBytesOk returns a tuple with the CloudTotalPhysicalUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudTotalPhysicalUsageBytes

`func (o *StorageDomainStats) SetCloudTotalPhysicalUsageBytes(v int64)`

SetCloudTotalPhysicalUsageBytes sets CloudTotalPhysicalUsageBytes field to given value.

### HasCloudTotalPhysicalUsageBytes

`func (o *StorageDomainStats) HasCloudTotalPhysicalUsageBytes() bool`

HasCloudTotalPhysicalUsageBytes returns a boolean if a field has been set.

### SetCloudTotalPhysicalUsageBytesNil

`func (o *StorageDomainStats) SetCloudTotalPhysicalUsageBytesNil(b bool)`

 SetCloudTotalPhysicalUsageBytesNil sets the value for CloudTotalPhysicalUsageBytes to be an explicit nil

### UnsetCloudTotalPhysicalUsageBytes
`func (o *StorageDomainStats) UnsetCloudTotalPhysicalUsageBytes()`

UnsetCloudTotalPhysicalUsageBytes ensures that no value is present for CloudTotalPhysicalUsageBytes, not even an explicit nil
### GetCloudTotalPhysicalUsageBytesTimestampUsec

`func (o *StorageDomainStats) GetCloudTotalPhysicalUsageBytesTimestampUsec() int64`

GetCloudTotalPhysicalUsageBytesTimestampUsec returns the CloudTotalPhysicalUsageBytesTimestampUsec field if non-nil, zero value otherwise.

### GetCloudTotalPhysicalUsageBytesTimestampUsecOk

`func (o *StorageDomainStats) GetCloudTotalPhysicalUsageBytesTimestampUsecOk() (*int64, bool)`

GetCloudTotalPhysicalUsageBytesTimestampUsecOk returns a tuple with the CloudTotalPhysicalUsageBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudTotalPhysicalUsageBytesTimestampUsec

`func (o *StorageDomainStats) SetCloudTotalPhysicalUsageBytesTimestampUsec(v int64)`

SetCloudTotalPhysicalUsageBytesTimestampUsec sets CloudTotalPhysicalUsageBytesTimestampUsec field to given value.

### HasCloudTotalPhysicalUsageBytesTimestampUsec

`func (o *StorageDomainStats) HasCloudTotalPhysicalUsageBytesTimestampUsec() bool`

HasCloudTotalPhysicalUsageBytesTimestampUsec returns a boolean if a field has been set.

### SetCloudTotalPhysicalUsageBytesTimestampUsecNil

`func (o *StorageDomainStats) SetCloudTotalPhysicalUsageBytesTimestampUsecNil(b bool)`

 SetCloudTotalPhysicalUsageBytesTimestampUsecNil sets the value for CloudTotalPhysicalUsageBytesTimestampUsec to be an explicit nil

### UnsetCloudTotalPhysicalUsageBytesTimestampUsec
`func (o *StorageDomainStats) UnsetCloudTotalPhysicalUsageBytesTimestampUsec()`

UnsetCloudTotalPhysicalUsageBytesTimestampUsec ensures that no value is present for CloudTotalPhysicalUsageBytesTimestampUsec, not even an explicit nil
### GetDataInBytes

`func (o *StorageDomainStats) GetDataInBytes() int64`

GetDataInBytes returns the DataInBytes field if non-nil, zero value otherwise.

### GetDataInBytesOk

`func (o *StorageDomainStats) GetDataInBytesOk() (*int64, bool)`

GetDataInBytesOk returns a tuple with the DataInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataInBytes

`func (o *StorageDomainStats) SetDataInBytes(v int64)`

SetDataInBytes sets DataInBytes field to given value.

### HasDataInBytes

`func (o *StorageDomainStats) HasDataInBytes() bool`

HasDataInBytes returns a boolean if a field has been set.

### SetDataInBytesNil

`func (o *StorageDomainStats) SetDataInBytesNil(b bool)`

 SetDataInBytesNil sets the value for DataInBytes to be an explicit nil

### UnsetDataInBytes
`func (o *StorageDomainStats) UnsetDataInBytes()`

UnsetDataInBytes ensures that no value is present for DataInBytes, not even an explicit nil
### GetDataInBytesAfterDedup

`func (o *StorageDomainStats) GetDataInBytesAfterDedup() int64`

GetDataInBytesAfterDedup returns the DataInBytesAfterDedup field if non-nil, zero value otherwise.

### GetDataInBytesAfterDedupOk

`func (o *StorageDomainStats) GetDataInBytesAfterDedupOk() (*int64, bool)`

GetDataInBytesAfterDedupOk returns a tuple with the DataInBytesAfterDedup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataInBytesAfterDedup

`func (o *StorageDomainStats) SetDataInBytesAfterDedup(v int64)`

SetDataInBytesAfterDedup sets DataInBytesAfterDedup field to given value.

### HasDataInBytesAfterDedup

`func (o *StorageDomainStats) HasDataInBytesAfterDedup() bool`

HasDataInBytesAfterDedup returns a boolean if a field has been set.

### SetDataInBytesAfterDedupNil

`func (o *StorageDomainStats) SetDataInBytesAfterDedupNil(b bool)`

 SetDataInBytesAfterDedupNil sets the value for DataInBytesAfterDedup to be an explicit nil

### UnsetDataInBytesAfterDedup
`func (o *StorageDomainStats) UnsetDataInBytesAfterDedup()`

UnsetDataInBytesAfterDedup ensures that no value is present for DataInBytesAfterDedup, not even an explicit nil
### GetDataInBytesAfterDedupTimestampUsec

`func (o *StorageDomainStats) GetDataInBytesAfterDedupTimestampUsec() int64`

GetDataInBytesAfterDedupTimestampUsec returns the DataInBytesAfterDedupTimestampUsec field if non-nil, zero value otherwise.

### GetDataInBytesAfterDedupTimestampUsecOk

`func (o *StorageDomainStats) GetDataInBytesAfterDedupTimestampUsecOk() (*int64, bool)`

GetDataInBytesAfterDedupTimestampUsecOk returns a tuple with the DataInBytesAfterDedupTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataInBytesAfterDedupTimestampUsec

`func (o *StorageDomainStats) SetDataInBytesAfterDedupTimestampUsec(v int64)`

SetDataInBytesAfterDedupTimestampUsec sets DataInBytesAfterDedupTimestampUsec field to given value.

### HasDataInBytesAfterDedupTimestampUsec

`func (o *StorageDomainStats) HasDataInBytesAfterDedupTimestampUsec() bool`

HasDataInBytesAfterDedupTimestampUsec returns a boolean if a field has been set.

### SetDataInBytesAfterDedupTimestampUsecNil

`func (o *StorageDomainStats) SetDataInBytesAfterDedupTimestampUsecNil(b bool)`

 SetDataInBytesAfterDedupTimestampUsecNil sets the value for DataInBytesAfterDedupTimestampUsec to be an explicit nil

### UnsetDataInBytesAfterDedupTimestampUsec
`func (o *StorageDomainStats) UnsetDataInBytesAfterDedupTimestampUsec()`

UnsetDataInBytesAfterDedupTimestampUsec ensures that no value is present for DataInBytesAfterDedupTimestampUsec, not even an explicit nil
### GetDataInBytesTimestampUsec

`func (o *StorageDomainStats) GetDataInBytesTimestampUsec() int64`

GetDataInBytesTimestampUsec returns the DataInBytesTimestampUsec field if non-nil, zero value otherwise.

### GetDataInBytesTimestampUsecOk

`func (o *StorageDomainStats) GetDataInBytesTimestampUsecOk() (*int64, bool)`

GetDataInBytesTimestampUsecOk returns a tuple with the DataInBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataInBytesTimestampUsec

`func (o *StorageDomainStats) SetDataInBytesTimestampUsec(v int64)`

SetDataInBytesTimestampUsec sets DataInBytesTimestampUsec field to given value.

### HasDataInBytesTimestampUsec

`func (o *StorageDomainStats) HasDataInBytesTimestampUsec() bool`

HasDataInBytesTimestampUsec returns a boolean if a field has been set.

### SetDataInBytesTimestampUsecNil

`func (o *StorageDomainStats) SetDataInBytesTimestampUsecNil(b bool)`

 SetDataInBytesTimestampUsecNil sets the value for DataInBytesTimestampUsec to be an explicit nil

### UnsetDataInBytesTimestampUsec
`func (o *StorageDomainStats) UnsetDataInBytesTimestampUsec()`

UnsetDataInBytesTimestampUsec ensures that no value is present for DataInBytesTimestampUsec, not even an explicit nil
### GetDataProtectLogicalUsageBytes

`func (o *StorageDomainStats) GetDataProtectLogicalUsageBytes() int64`

GetDataProtectLogicalUsageBytes returns the DataProtectLogicalUsageBytes field if non-nil, zero value otherwise.

### GetDataProtectLogicalUsageBytesOk

`func (o *StorageDomainStats) GetDataProtectLogicalUsageBytesOk() (*int64, bool)`

GetDataProtectLogicalUsageBytesOk returns a tuple with the DataProtectLogicalUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProtectLogicalUsageBytes

`func (o *StorageDomainStats) SetDataProtectLogicalUsageBytes(v int64)`

SetDataProtectLogicalUsageBytes sets DataProtectLogicalUsageBytes field to given value.

### HasDataProtectLogicalUsageBytes

`func (o *StorageDomainStats) HasDataProtectLogicalUsageBytes() bool`

HasDataProtectLogicalUsageBytes returns a boolean if a field has been set.

### SetDataProtectLogicalUsageBytesNil

`func (o *StorageDomainStats) SetDataProtectLogicalUsageBytesNil(b bool)`

 SetDataProtectLogicalUsageBytesNil sets the value for DataProtectLogicalUsageBytes to be an explicit nil

### UnsetDataProtectLogicalUsageBytes
`func (o *StorageDomainStats) UnsetDataProtectLogicalUsageBytes()`

UnsetDataProtectLogicalUsageBytes ensures that no value is present for DataProtectLogicalUsageBytes, not even an explicit nil
### GetDataProtectLogicalUsageBytesTimestampUsec

`func (o *StorageDomainStats) GetDataProtectLogicalUsageBytesTimestampUsec() int64`

GetDataProtectLogicalUsageBytesTimestampUsec returns the DataProtectLogicalUsageBytesTimestampUsec field if non-nil, zero value otherwise.

### GetDataProtectLogicalUsageBytesTimestampUsecOk

`func (o *StorageDomainStats) GetDataProtectLogicalUsageBytesTimestampUsecOk() (*int64, bool)`

GetDataProtectLogicalUsageBytesTimestampUsecOk returns a tuple with the DataProtectLogicalUsageBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProtectLogicalUsageBytesTimestampUsec

`func (o *StorageDomainStats) SetDataProtectLogicalUsageBytesTimestampUsec(v int64)`

SetDataProtectLogicalUsageBytesTimestampUsec sets DataProtectLogicalUsageBytesTimestampUsec field to given value.

### HasDataProtectLogicalUsageBytesTimestampUsec

`func (o *StorageDomainStats) HasDataProtectLogicalUsageBytesTimestampUsec() bool`

HasDataProtectLogicalUsageBytesTimestampUsec returns a boolean if a field has been set.

### SetDataProtectLogicalUsageBytesTimestampUsecNil

`func (o *StorageDomainStats) SetDataProtectLogicalUsageBytesTimestampUsecNil(b bool)`

 SetDataProtectLogicalUsageBytesTimestampUsecNil sets the value for DataProtectLogicalUsageBytesTimestampUsec to be an explicit nil

### UnsetDataProtectLogicalUsageBytesTimestampUsec
`func (o *StorageDomainStats) UnsetDataProtectLogicalUsageBytesTimestampUsec()`

UnsetDataProtectLogicalUsageBytesTimestampUsec ensures that no value is present for DataProtectLogicalUsageBytesTimestampUsec, not even an explicit nil
### GetDataProtectPhysicalUsageBytes

`func (o *StorageDomainStats) GetDataProtectPhysicalUsageBytes() int64`

GetDataProtectPhysicalUsageBytes returns the DataProtectPhysicalUsageBytes field if non-nil, zero value otherwise.

### GetDataProtectPhysicalUsageBytesOk

`func (o *StorageDomainStats) GetDataProtectPhysicalUsageBytesOk() (*int64, bool)`

GetDataProtectPhysicalUsageBytesOk returns a tuple with the DataProtectPhysicalUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProtectPhysicalUsageBytes

`func (o *StorageDomainStats) SetDataProtectPhysicalUsageBytes(v int64)`

SetDataProtectPhysicalUsageBytes sets DataProtectPhysicalUsageBytes field to given value.

### HasDataProtectPhysicalUsageBytes

`func (o *StorageDomainStats) HasDataProtectPhysicalUsageBytes() bool`

HasDataProtectPhysicalUsageBytes returns a boolean if a field has been set.

### SetDataProtectPhysicalUsageBytesNil

`func (o *StorageDomainStats) SetDataProtectPhysicalUsageBytesNil(b bool)`

 SetDataProtectPhysicalUsageBytesNil sets the value for DataProtectPhysicalUsageBytes to be an explicit nil

### UnsetDataProtectPhysicalUsageBytes
`func (o *StorageDomainStats) UnsetDataProtectPhysicalUsageBytes()`

UnsetDataProtectPhysicalUsageBytes ensures that no value is present for DataProtectPhysicalUsageBytes, not even an explicit nil
### GetDataProtectPhysicalUsageBytesTimestampUsec

`func (o *StorageDomainStats) GetDataProtectPhysicalUsageBytesTimestampUsec() int64`

GetDataProtectPhysicalUsageBytesTimestampUsec returns the DataProtectPhysicalUsageBytesTimestampUsec field if non-nil, zero value otherwise.

### GetDataProtectPhysicalUsageBytesTimestampUsecOk

`func (o *StorageDomainStats) GetDataProtectPhysicalUsageBytesTimestampUsecOk() (*int64, bool)`

GetDataProtectPhysicalUsageBytesTimestampUsecOk returns a tuple with the DataProtectPhysicalUsageBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProtectPhysicalUsageBytesTimestampUsec

`func (o *StorageDomainStats) SetDataProtectPhysicalUsageBytesTimestampUsec(v int64)`

SetDataProtectPhysicalUsageBytesTimestampUsec sets DataProtectPhysicalUsageBytesTimestampUsec field to given value.

### HasDataProtectPhysicalUsageBytesTimestampUsec

`func (o *StorageDomainStats) HasDataProtectPhysicalUsageBytesTimestampUsec() bool`

HasDataProtectPhysicalUsageBytesTimestampUsec returns a boolean if a field has been set.

### SetDataProtectPhysicalUsageBytesTimestampUsecNil

`func (o *StorageDomainStats) SetDataProtectPhysicalUsageBytesTimestampUsecNil(b bool)`

 SetDataProtectPhysicalUsageBytesTimestampUsecNil sets the value for DataProtectPhysicalUsageBytesTimestampUsec to be an explicit nil

### UnsetDataProtectPhysicalUsageBytesTimestampUsec
`func (o *StorageDomainStats) UnsetDataProtectPhysicalUsageBytesTimestampUsec()`

UnsetDataProtectPhysicalUsageBytesTimestampUsec ensures that no value is present for DataProtectPhysicalUsageBytesTimestampUsec, not even an explicit nil
### GetDataWrittenBytes

`func (o *StorageDomainStats) GetDataWrittenBytes() int64`

GetDataWrittenBytes returns the DataWrittenBytes field if non-nil, zero value otherwise.

### GetDataWrittenBytesOk

`func (o *StorageDomainStats) GetDataWrittenBytesOk() (*int64, bool)`

GetDataWrittenBytesOk returns a tuple with the DataWrittenBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWrittenBytes

`func (o *StorageDomainStats) SetDataWrittenBytes(v int64)`

SetDataWrittenBytes sets DataWrittenBytes field to given value.

### HasDataWrittenBytes

`func (o *StorageDomainStats) HasDataWrittenBytes() bool`

HasDataWrittenBytes returns a boolean if a field has been set.

### SetDataWrittenBytesNil

`func (o *StorageDomainStats) SetDataWrittenBytesNil(b bool)`

 SetDataWrittenBytesNil sets the value for DataWrittenBytes to be an explicit nil

### UnsetDataWrittenBytes
`func (o *StorageDomainStats) UnsetDataWrittenBytes()`

UnsetDataWrittenBytes ensures that no value is present for DataWrittenBytes, not even an explicit nil
### GetDataWrittenBytesTimestampUsec

`func (o *StorageDomainStats) GetDataWrittenBytesTimestampUsec() int64`

GetDataWrittenBytesTimestampUsec returns the DataWrittenBytesTimestampUsec field if non-nil, zero value otherwise.

### GetDataWrittenBytesTimestampUsecOk

`func (o *StorageDomainStats) GetDataWrittenBytesTimestampUsecOk() (*int64, bool)`

GetDataWrittenBytesTimestampUsecOk returns a tuple with the DataWrittenBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWrittenBytesTimestampUsec

`func (o *StorageDomainStats) SetDataWrittenBytesTimestampUsec(v int64)`

SetDataWrittenBytesTimestampUsec sets DataWrittenBytesTimestampUsec field to given value.

### HasDataWrittenBytesTimestampUsec

`func (o *StorageDomainStats) HasDataWrittenBytesTimestampUsec() bool`

HasDataWrittenBytesTimestampUsec returns a boolean if a field has been set.

### SetDataWrittenBytesTimestampUsecNil

`func (o *StorageDomainStats) SetDataWrittenBytesTimestampUsecNil(b bool)`

 SetDataWrittenBytesTimestampUsecNil sets the value for DataWrittenBytesTimestampUsec to be an explicit nil

### UnsetDataWrittenBytesTimestampUsec
`func (o *StorageDomainStats) UnsetDataWrittenBytesTimestampUsec()`

UnsetDataWrittenBytesTimestampUsec ensures that no value is present for DataWrittenBytesTimestampUsec, not even an explicit nil
### GetFileServicesLogicalUsageBytes

`func (o *StorageDomainStats) GetFileServicesLogicalUsageBytes() int64`

GetFileServicesLogicalUsageBytes returns the FileServicesLogicalUsageBytes field if non-nil, zero value otherwise.

### GetFileServicesLogicalUsageBytesOk

`func (o *StorageDomainStats) GetFileServicesLogicalUsageBytesOk() (*int64, bool)`

GetFileServicesLogicalUsageBytesOk returns a tuple with the FileServicesLogicalUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileServicesLogicalUsageBytes

`func (o *StorageDomainStats) SetFileServicesLogicalUsageBytes(v int64)`

SetFileServicesLogicalUsageBytes sets FileServicesLogicalUsageBytes field to given value.

### HasFileServicesLogicalUsageBytes

`func (o *StorageDomainStats) HasFileServicesLogicalUsageBytes() bool`

HasFileServicesLogicalUsageBytes returns a boolean if a field has been set.

### SetFileServicesLogicalUsageBytesNil

`func (o *StorageDomainStats) SetFileServicesLogicalUsageBytesNil(b bool)`

 SetFileServicesLogicalUsageBytesNil sets the value for FileServicesLogicalUsageBytes to be an explicit nil

### UnsetFileServicesLogicalUsageBytes
`func (o *StorageDomainStats) UnsetFileServicesLogicalUsageBytes()`

UnsetFileServicesLogicalUsageBytes ensures that no value is present for FileServicesLogicalUsageBytes, not even an explicit nil
### GetFileServicesLogicalUsageBytesTimestampUsec

`func (o *StorageDomainStats) GetFileServicesLogicalUsageBytesTimestampUsec() int64`

GetFileServicesLogicalUsageBytesTimestampUsec returns the FileServicesLogicalUsageBytesTimestampUsec field if non-nil, zero value otherwise.

### GetFileServicesLogicalUsageBytesTimestampUsecOk

`func (o *StorageDomainStats) GetFileServicesLogicalUsageBytesTimestampUsecOk() (*int64, bool)`

GetFileServicesLogicalUsageBytesTimestampUsecOk returns a tuple with the FileServicesLogicalUsageBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileServicesLogicalUsageBytesTimestampUsec

`func (o *StorageDomainStats) SetFileServicesLogicalUsageBytesTimestampUsec(v int64)`

SetFileServicesLogicalUsageBytesTimestampUsec sets FileServicesLogicalUsageBytesTimestampUsec field to given value.

### HasFileServicesLogicalUsageBytesTimestampUsec

`func (o *StorageDomainStats) HasFileServicesLogicalUsageBytesTimestampUsec() bool`

HasFileServicesLogicalUsageBytesTimestampUsec returns a boolean if a field has been set.

### SetFileServicesLogicalUsageBytesTimestampUsecNil

`func (o *StorageDomainStats) SetFileServicesLogicalUsageBytesTimestampUsecNil(b bool)`

 SetFileServicesLogicalUsageBytesTimestampUsecNil sets the value for FileServicesLogicalUsageBytesTimestampUsec to be an explicit nil

### UnsetFileServicesLogicalUsageBytesTimestampUsec
`func (o *StorageDomainStats) UnsetFileServicesLogicalUsageBytesTimestampUsec()`

UnsetFileServicesLogicalUsageBytesTimestampUsec ensures that no value is present for FileServicesLogicalUsageBytesTimestampUsec, not even an explicit nil
### GetFileServicesPhysicalUsageBytes

`func (o *StorageDomainStats) GetFileServicesPhysicalUsageBytes() int64`

GetFileServicesPhysicalUsageBytes returns the FileServicesPhysicalUsageBytes field if non-nil, zero value otherwise.

### GetFileServicesPhysicalUsageBytesOk

`func (o *StorageDomainStats) GetFileServicesPhysicalUsageBytesOk() (*int64, bool)`

GetFileServicesPhysicalUsageBytesOk returns a tuple with the FileServicesPhysicalUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileServicesPhysicalUsageBytes

`func (o *StorageDomainStats) SetFileServicesPhysicalUsageBytes(v int64)`

SetFileServicesPhysicalUsageBytes sets FileServicesPhysicalUsageBytes field to given value.

### HasFileServicesPhysicalUsageBytes

`func (o *StorageDomainStats) HasFileServicesPhysicalUsageBytes() bool`

HasFileServicesPhysicalUsageBytes returns a boolean if a field has been set.

### SetFileServicesPhysicalUsageBytesNil

`func (o *StorageDomainStats) SetFileServicesPhysicalUsageBytesNil(b bool)`

 SetFileServicesPhysicalUsageBytesNil sets the value for FileServicesPhysicalUsageBytes to be an explicit nil

### UnsetFileServicesPhysicalUsageBytes
`func (o *StorageDomainStats) UnsetFileServicesPhysicalUsageBytes()`

UnsetFileServicesPhysicalUsageBytes ensures that no value is present for FileServicesPhysicalUsageBytes, not even an explicit nil
### GetFileServicesPhysicalUsageBytesTimestampUsec

`func (o *StorageDomainStats) GetFileServicesPhysicalUsageBytesTimestampUsec() int64`

GetFileServicesPhysicalUsageBytesTimestampUsec returns the FileServicesPhysicalUsageBytesTimestampUsec field if non-nil, zero value otherwise.

### GetFileServicesPhysicalUsageBytesTimestampUsecOk

`func (o *StorageDomainStats) GetFileServicesPhysicalUsageBytesTimestampUsecOk() (*int64, bool)`

GetFileServicesPhysicalUsageBytesTimestampUsecOk returns a tuple with the FileServicesPhysicalUsageBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileServicesPhysicalUsageBytesTimestampUsec

`func (o *StorageDomainStats) SetFileServicesPhysicalUsageBytesTimestampUsec(v int64)`

SetFileServicesPhysicalUsageBytesTimestampUsec sets FileServicesPhysicalUsageBytesTimestampUsec field to given value.

### HasFileServicesPhysicalUsageBytesTimestampUsec

`func (o *StorageDomainStats) HasFileServicesPhysicalUsageBytesTimestampUsec() bool`

HasFileServicesPhysicalUsageBytesTimestampUsec returns a boolean if a field has been set.

### SetFileServicesPhysicalUsageBytesTimestampUsecNil

`func (o *StorageDomainStats) SetFileServicesPhysicalUsageBytesTimestampUsecNil(b bool)`

 SetFileServicesPhysicalUsageBytesTimestampUsecNil sets the value for FileServicesPhysicalUsageBytesTimestampUsec to be an explicit nil

### UnsetFileServicesPhysicalUsageBytesTimestampUsec
`func (o *StorageDomainStats) UnsetFileServicesPhysicalUsageBytesTimestampUsec()`

UnsetFileServicesPhysicalUsageBytesTimestampUsec ensures that no value is present for FileServicesPhysicalUsageBytesTimestampUsec, not even an explicit nil
### GetLocalDataWrittenBytes

`func (o *StorageDomainStats) GetLocalDataWrittenBytes() int64`

GetLocalDataWrittenBytes returns the LocalDataWrittenBytes field if non-nil, zero value otherwise.

### GetLocalDataWrittenBytesOk

`func (o *StorageDomainStats) GetLocalDataWrittenBytesOk() (*int64, bool)`

GetLocalDataWrittenBytesOk returns a tuple with the LocalDataWrittenBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDataWrittenBytes

`func (o *StorageDomainStats) SetLocalDataWrittenBytes(v int64)`

SetLocalDataWrittenBytes sets LocalDataWrittenBytes field to given value.

### HasLocalDataWrittenBytes

`func (o *StorageDomainStats) HasLocalDataWrittenBytes() bool`

HasLocalDataWrittenBytes returns a boolean if a field has been set.

### SetLocalDataWrittenBytesNil

`func (o *StorageDomainStats) SetLocalDataWrittenBytesNil(b bool)`

 SetLocalDataWrittenBytesNil sets the value for LocalDataWrittenBytes to be an explicit nil

### UnsetLocalDataWrittenBytes
`func (o *StorageDomainStats) UnsetLocalDataWrittenBytes()`

UnsetLocalDataWrittenBytes ensures that no value is present for LocalDataWrittenBytes, not even an explicit nil
### GetLocalDataWrittenBytesTimestampUsec

`func (o *StorageDomainStats) GetLocalDataWrittenBytesTimestampUsec() int64`

GetLocalDataWrittenBytesTimestampUsec returns the LocalDataWrittenBytesTimestampUsec field if non-nil, zero value otherwise.

### GetLocalDataWrittenBytesTimestampUsecOk

`func (o *StorageDomainStats) GetLocalDataWrittenBytesTimestampUsecOk() (*int64, bool)`

GetLocalDataWrittenBytesTimestampUsecOk returns a tuple with the LocalDataWrittenBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDataWrittenBytesTimestampUsec

`func (o *StorageDomainStats) SetLocalDataWrittenBytesTimestampUsec(v int64)`

SetLocalDataWrittenBytesTimestampUsec sets LocalDataWrittenBytesTimestampUsec field to given value.

### HasLocalDataWrittenBytesTimestampUsec

`func (o *StorageDomainStats) HasLocalDataWrittenBytesTimestampUsec() bool`

HasLocalDataWrittenBytesTimestampUsec returns a boolean if a field has been set.

### SetLocalDataWrittenBytesTimestampUsecNil

`func (o *StorageDomainStats) SetLocalDataWrittenBytesTimestampUsecNil(b bool)`

 SetLocalDataWrittenBytesTimestampUsecNil sets the value for LocalDataWrittenBytesTimestampUsec to be an explicit nil

### UnsetLocalDataWrittenBytesTimestampUsec
`func (o *StorageDomainStats) UnsetLocalDataWrittenBytesTimestampUsec()`

UnsetLocalDataWrittenBytesTimestampUsec ensures that no value is present for LocalDataWrittenBytesTimestampUsec, not even an explicit nil
### GetLocalTierResiliencyImpactBytes

`func (o *StorageDomainStats) GetLocalTierResiliencyImpactBytes() int64`

GetLocalTierResiliencyImpactBytes returns the LocalTierResiliencyImpactBytes field if non-nil, zero value otherwise.

### GetLocalTierResiliencyImpactBytesOk

`func (o *StorageDomainStats) GetLocalTierResiliencyImpactBytesOk() (*int64, bool)`

GetLocalTierResiliencyImpactBytesOk returns a tuple with the LocalTierResiliencyImpactBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTierResiliencyImpactBytes

`func (o *StorageDomainStats) SetLocalTierResiliencyImpactBytes(v int64)`

SetLocalTierResiliencyImpactBytes sets LocalTierResiliencyImpactBytes field to given value.

### HasLocalTierResiliencyImpactBytes

`func (o *StorageDomainStats) HasLocalTierResiliencyImpactBytes() bool`

HasLocalTierResiliencyImpactBytes returns a boolean if a field has been set.

### SetLocalTierResiliencyImpactBytesNil

`func (o *StorageDomainStats) SetLocalTierResiliencyImpactBytesNil(b bool)`

 SetLocalTierResiliencyImpactBytesNil sets the value for LocalTierResiliencyImpactBytes to be an explicit nil

### UnsetLocalTierResiliencyImpactBytes
`func (o *StorageDomainStats) UnsetLocalTierResiliencyImpactBytes()`

UnsetLocalTierResiliencyImpactBytes ensures that no value is present for LocalTierResiliencyImpactBytes, not even an explicit nil
### GetLocalTierResiliencyImpactBytesTimestampUsec

`func (o *StorageDomainStats) GetLocalTierResiliencyImpactBytesTimestampUsec() int64`

GetLocalTierResiliencyImpactBytesTimestampUsec returns the LocalTierResiliencyImpactBytesTimestampUsec field if non-nil, zero value otherwise.

### GetLocalTierResiliencyImpactBytesTimestampUsecOk

`func (o *StorageDomainStats) GetLocalTierResiliencyImpactBytesTimestampUsecOk() (*int64, bool)`

GetLocalTierResiliencyImpactBytesTimestampUsecOk returns a tuple with the LocalTierResiliencyImpactBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTierResiliencyImpactBytesTimestampUsec

`func (o *StorageDomainStats) SetLocalTierResiliencyImpactBytesTimestampUsec(v int64)`

SetLocalTierResiliencyImpactBytesTimestampUsec sets LocalTierResiliencyImpactBytesTimestampUsec field to given value.

### HasLocalTierResiliencyImpactBytesTimestampUsec

`func (o *StorageDomainStats) HasLocalTierResiliencyImpactBytesTimestampUsec() bool`

HasLocalTierResiliencyImpactBytesTimestampUsec returns a boolean if a field has been set.

### SetLocalTierResiliencyImpactBytesTimestampUsecNil

`func (o *StorageDomainStats) SetLocalTierResiliencyImpactBytesTimestampUsecNil(b bool)`

 SetLocalTierResiliencyImpactBytesTimestampUsecNil sets the value for LocalTierResiliencyImpactBytesTimestampUsec to be an explicit nil

### UnsetLocalTierResiliencyImpactBytesTimestampUsec
`func (o *StorageDomainStats) UnsetLocalTierResiliencyImpactBytesTimestampUsec()`

UnsetLocalTierResiliencyImpactBytesTimestampUsec ensures that no value is present for LocalTierResiliencyImpactBytesTimestampUsec, not even an explicit nil
### GetLocalTotalPhysicalUsageBytes

`func (o *StorageDomainStats) GetLocalTotalPhysicalUsageBytes() int64`

GetLocalTotalPhysicalUsageBytes returns the LocalTotalPhysicalUsageBytes field if non-nil, zero value otherwise.

### GetLocalTotalPhysicalUsageBytesOk

`func (o *StorageDomainStats) GetLocalTotalPhysicalUsageBytesOk() (*int64, bool)`

GetLocalTotalPhysicalUsageBytesOk returns a tuple with the LocalTotalPhysicalUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTotalPhysicalUsageBytes

`func (o *StorageDomainStats) SetLocalTotalPhysicalUsageBytes(v int64)`

SetLocalTotalPhysicalUsageBytes sets LocalTotalPhysicalUsageBytes field to given value.

### HasLocalTotalPhysicalUsageBytes

`func (o *StorageDomainStats) HasLocalTotalPhysicalUsageBytes() bool`

HasLocalTotalPhysicalUsageBytes returns a boolean if a field has been set.

### SetLocalTotalPhysicalUsageBytesNil

`func (o *StorageDomainStats) SetLocalTotalPhysicalUsageBytesNil(b bool)`

 SetLocalTotalPhysicalUsageBytesNil sets the value for LocalTotalPhysicalUsageBytes to be an explicit nil

### UnsetLocalTotalPhysicalUsageBytes
`func (o *StorageDomainStats) UnsetLocalTotalPhysicalUsageBytes()`

UnsetLocalTotalPhysicalUsageBytes ensures that no value is present for LocalTotalPhysicalUsageBytes, not even an explicit nil
### GetLocalTotalPhysicalUsageBytesTimestampUsec

`func (o *StorageDomainStats) GetLocalTotalPhysicalUsageBytesTimestampUsec() int64`

GetLocalTotalPhysicalUsageBytesTimestampUsec returns the LocalTotalPhysicalUsageBytesTimestampUsec field if non-nil, zero value otherwise.

### GetLocalTotalPhysicalUsageBytesTimestampUsecOk

`func (o *StorageDomainStats) GetLocalTotalPhysicalUsageBytesTimestampUsecOk() (*int64, bool)`

GetLocalTotalPhysicalUsageBytesTimestampUsecOk returns a tuple with the LocalTotalPhysicalUsageBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTotalPhysicalUsageBytesTimestampUsec

`func (o *StorageDomainStats) SetLocalTotalPhysicalUsageBytesTimestampUsec(v int64)`

SetLocalTotalPhysicalUsageBytesTimestampUsec sets LocalTotalPhysicalUsageBytesTimestampUsec field to given value.

### HasLocalTotalPhysicalUsageBytesTimestampUsec

`func (o *StorageDomainStats) HasLocalTotalPhysicalUsageBytesTimestampUsec() bool`

HasLocalTotalPhysicalUsageBytesTimestampUsec returns a boolean if a field has been set.

### SetLocalTotalPhysicalUsageBytesTimestampUsecNil

`func (o *StorageDomainStats) SetLocalTotalPhysicalUsageBytesTimestampUsecNil(b bool)`

 SetLocalTotalPhysicalUsageBytesTimestampUsecNil sets the value for LocalTotalPhysicalUsageBytesTimestampUsec to be an explicit nil

### UnsetLocalTotalPhysicalUsageBytesTimestampUsec
`func (o *StorageDomainStats) UnsetLocalTotalPhysicalUsageBytesTimestampUsec()`

UnsetLocalTotalPhysicalUsageBytesTimestampUsec ensures that no value is present for LocalTotalPhysicalUsageBytesTimestampUsec, not even an explicit nil
### GetNumDirectories

`func (o *StorageDomainStats) GetNumDirectories() int64`

GetNumDirectories returns the NumDirectories field if non-nil, zero value otherwise.

### GetNumDirectoriesOk

`func (o *StorageDomainStats) GetNumDirectoriesOk() (*int64, bool)`

GetNumDirectoriesOk returns a tuple with the NumDirectories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDirectories

`func (o *StorageDomainStats) SetNumDirectories(v int64)`

SetNumDirectories sets NumDirectories field to given value.

### HasNumDirectories

`func (o *StorageDomainStats) HasNumDirectories() bool`

HasNumDirectories returns a boolean if a field has been set.

### SetNumDirectoriesNil

`func (o *StorageDomainStats) SetNumDirectoriesNil(b bool)`

 SetNumDirectoriesNil sets the value for NumDirectories to be an explicit nil

### UnsetNumDirectories
`func (o *StorageDomainStats) UnsetNumDirectories()`

UnsetNumDirectories ensures that no value is present for NumDirectories, not even an explicit nil
### GetNumFiles

`func (o *StorageDomainStats) GetNumFiles() int64`

GetNumFiles returns the NumFiles field if non-nil, zero value otherwise.

### GetNumFilesOk

`func (o *StorageDomainStats) GetNumFilesOk() (*int64, bool)`

GetNumFilesOk returns a tuple with the NumFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFiles

`func (o *StorageDomainStats) SetNumFiles(v int64)`

SetNumFiles sets NumFiles field to given value.

### HasNumFiles

`func (o *StorageDomainStats) HasNumFiles() bool`

HasNumFiles returns a boolean if a field has been set.

### SetNumFilesNil

`func (o *StorageDomainStats) SetNumFilesNil(b bool)`

 SetNumFilesNil sets the value for NumFiles to be an explicit nil

### UnsetNumFiles
`func (o *StorageDomainStats) UnsetNumFiles()`

UnsetNumFiles ensures that no value is present for NumFiles, not even an explicit nil
### GetOutdatedLogicalUsageBytes

`func (o *StorageDomainStats) GetOutdatedLogicalUsageBytes() int64`

GetOutdatedLogicalUsageBytes returns the OutdatedLogicalUsageBytes field if non-nil, zero value otherwise.

### GetOutdatedLogicalUsageBytesOk

`func (o *StorageDomainStats) GetOutdatedLogicalUsageBytesOk() (*int64, bool)`

GetOutdatedLogicalUsageBytesOk returns a tuple with the OutdatedLogicalUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutdatedLogicalUsageBytes

`func (o *StorageDomainStats) SetOutdatedLogicalUsageBytes(v int64)`

SetOutdatedLogicalUsageBytes sets OutdatedLogicalUsageBytes field to given value.

### HasOutdatedLogicalUsageBytes

`func (o *StorageDomainStats) HasOutdatedLogicalUsageBytes() bool`

HasOutdatedLogicalUsageBytes returns a boolean if a field has been set.

### SetOutdatedLogicalUsageBytesNil

`func (o *StorageDomainStats) SetOutdatedLogicalUsageBytesNil(b bool)`

 SetOutdatedLogicalUsageBytesNil sets the value for OutdatedLogicalUsageBytes to be an explicit nil

### UnsetOutdatedLogicalUsageBytes
`func (o *StorageDomainStats) UnsetOutdatedLogicalUsageBytes()`

UnsetOutdatedLogicalUsageBytes ensures that no value is present for OutdatedLogicalUsageBytes, not even an explicit nil
### GetOutdatedLogicalUsageBytesTimestampUsec

`func (o *StorageDomainStats) GetOutdatedLogicalUsageBytesTimestampUsec() int64`

GetOutdatedLogicalUsageBytesTimestampUsec returns the OutdatedLogicalUsageBytesTimestampUsec field if non-nil, zero value otherwise.

### GetOutdatedLogicalUsageBytesTimestampUsecOk

`func (o *StorageDomainStats) GetOutdatedLogicalUsageBytesTimestampUsecOk() (*int64, bool)`

GetOutdatedLogicalUsageBytesTimestampUsecOk returns a tuple with the OutdatedLogicalUsageBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutdatedLogicalUsageBytesTimestampUsec

`func (o *StorageDomainStats) SetOutdatedLogicalUsageBytesTimestampUsec(v int64)`

SetOutdatedLogicalUsageBytesTimestampUsec sets OutdatedLogicalUsageBytesTimestampUsec field to given value.

### HasOutdatedLogicalUsageBytesTimestampUsec

`func (o *StorageDomainStats) HasOutdatedLogicalUsageBytesTimestampUsec() bool`

HasOutdatedLogicalUsageBytesTimestampUsec returns a boolean if a field has been set.

### SetOutdatedLogicalUsageBytesTimestampUsecNil

`func (o *StorageDomainStats) SetOutdatedLogicalUsageBytesTimestampUsecNil(b bool)`

 SetOutdatedLogicalUsageBytesTimestampUsecNil sets the value for OutdatedLogicalUsageBytesTimestampUsec to be an explicit nil

### UnsetOutdatedLogicalUsageBytesTimestampUsec
`func (o *StorageDomainStats) UnsetOutdatedLogicalUsageBytesTimestampUsec()`

UnsetOutdatedLogicalUsageBytesTimestampUsec ensures that no value is present for OutdatedLogicalUsageBytesTimestampUsec, not even an explicit nil
### GetStorageConsumedBytes

`func (o *StorageDomainStats) GetStorageConsumedBytes() int64`

GetStorageConsumedBytes returns the StorageConsumedBytes field if non-nil, zero value otherwise.

### GetStorageConsumedBytesOk

`func (o *StorageDomainStats) GetStorageConsumedBytesOk() (*int64, bool)`

GetStorageConsumedBytesOk returns a tuple with the StorageConsumedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConsumedBytes

`func (o *StorageDomainStats) SetStorageConsumedBytes(v int64)`

SetStorageConsumedBytes sets StorageConsumedBytes field to given value.

### HasStorageConsumedBytes

`func (o *StorageDomainStats) HasStorageConsumedBytes() bool`

HasStorageConsumedBytes returns a boolean if a field has been set.

### SetStorageConsumedBytesNil

`func (o *StorageDomainStats) SetStorageConsumedBytesNil(b bool)`

 SetStorageConsumedBytesNil sets the value for StorageConsumedBytes to be an explicit nil

### UnsetStorageConsumedBytes
`func (o *StorageDomainStats) UnsetStorageConsumedBytes()`

UnsetStorageConsumedBytes ensures that no value is present for StorageConsumedBytes, not even an explicit nil
### GetStorageConsumedBytesTimestampUsec

`func (o *StorageDomainStats) GetStorageConsumedBytesTimestampUsec() int64`

GetStorageConsumedBytesTimestampUsec returns the StorageConsumedBytesTimestampUsec field if non-nil, zero value otherwise.

### GetStorageConsumedBytesTimestampUsecOk

`func (o *StorageDomainStats) GetStorageConsumedBytesTimestampUsecOk() (*int64, bool)`

GetStorageConsumedBytesTimestampUsecOk returns a tuple with the StorageConsumedBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConsumedBytesTimestampUsec

`func (o *StorageDomainStats) SetStorageConsumedBytesTimestampUsec(v int64)`

SetStorageConsumedBytesTimestampUsec sets StorageConsumedBytesTimestampUsec field to given value.

### HasStorageConsumedBytesTimestampUsec

`func (o *StorageDomainStats) HasStorageConsumedBytesTimestampUsec() bool`

HasStorageConsumedBytesTimestampUsec returns a boolean if a field has been set.

### SetStorageConsumedBytesTimestampUsecNil

`func (o *StorageDomainStats) SetStorageConsumedBytesTimestampUsecNil(b bool)`

 SetStorageConsumedBytesTimestampUsecNil sets the value for StorageConsumedBytesTimestampUsec to be an explicit nil

### UnsetStorageConsumedBytesTimestampUsec
`func (o *StorageDomainStats) UnsetStorageConsumedBytesTimestampUsec()`

UnsetStorageConsumedBytesTimestampUsec ensures that no value is present for StorageConsumedBytesTimestampUsec, not even an explicit nil
### GetTotalLogicalUsageBytes

`func (o *StorageDomainStats) GetTotalLogicalUsageBytes() int64`

GetTotalLogicalUsageBytes returns the TotalLogicalUsageBytes field if non-nil, zero value otherwise.

### GetTotalLogicalUsageBytesOk

`func (o *StorageDomainStats) GetTotalLogicalUsageBytesOk() (*int64, bool)`

GetTotalLogicalUsageBytesOk returns a tuple with the TotalLogicalUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLogicalUsageBytes

`func (o *StorageDomainStats) SetTotalLogicalUsageBytes(v int64)`

SetTotalLogicalUsageBytes sets TotalLogicalUsageBytes field to given value.

### HasTotalLogicalUsageBytes

`func (o *StorageDomainStats) HasTotalLogicalUsageBytes() bool`

HasTotalLogicalUsageBytes returns a boolean if a field has been set.

### SetTotalLogicalUsageBytesNil

`func (o *StorageDomainStats) SetTotalLogicalUsageBytesNil(b bool)`

 SetTotalLogicalUsageBytesNil sets the value for TotalLogicalUsageBytes to be an explicit nil

### UnsetTotalLogicalUsageBytes
`func (o *StorageDomainStats) UnsetTotalLogicalUsageBytes()`

UnsetTotalLogicalUsageBytes ensures that no value is present for TotalLogicalUsageBytes, not even an explicit nil
### GetTotalLogicalUsageBytesTimestampUsec

`func (o *StorageDomainStats) GetTotalLogicalUsageBytesTimestampUsec() int64`

GetTotalLogicalUsageBytesTimestampUsec returns the TotalLogicalUsageBytesTimestampUsec field if non-nil, zero value otherwise.

### GetTotalLogicalUsageBytesTimestampUsecOk

`func (o *StorageDomainStats) GetTotalLogicalUsageBytesTimestampUsecOk() (*int64, bool)`

GetTotalLogicalUsageBytesTimestampUsecOk returns a tuple with the TotalLogicalUsageBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLogicalUsageBytesTimestampUsec

`func (o *StorageDomainStats) SetTotalLogicalUsageBytesTimestampUsec(v int64)`

SetTotalLogicalUsageBytesTimestampUsec sets TotalLogicalUsageBytesTimestampUsec field to given value.

### HasTotalLogicalUsageBytesTimestampUsec

`func (o *StorageDomainStats) HasTotalLogicalUsageBytesTimestampUsec() bool`

HasTotalLogicalUsageBytesTimestampUsec returns a boolean if a field has been set.

### SetTotalLogicalUsageBytesTimestampUsecNil

`func (o *StorageDomainStats) SetTotalLogicalUsageBytesTimestampUsecNil(b bool)`

 SetTotalLogicalUsageBytesTimestampUsecNil sets the value for TotalLogicalUsageBytesTimestampUsec to be an explicit nil

### UnsetTotalLogicalUsageBytesTimestampUsec
`func (o *StorageDomainStats) UnsetTotalLogicalUsageBytesTimestampUsec()`

UnsetTotalLogicalUsageBytesTimestampUsec ensures that no value is present for TotalLogicalUsageBytesTimestampUsec, not even an explicit nil
### GetUniquePhysicalDataBytes

`func (o *StorageDomainStats) GetUniquePhysicalDataBytes() int64`

GetUniquePhysicalDataBytes returns the UniquePhysicalDataBytes field if non-nil, zero value otherwise.

### GetUniquePhysicalDataBytesOk

`func (o *StorageDomainStats) GetUniquePhysicalDataBytesOk() (*int64, bool)`

GetUniquePhysicalDataBytesOk returns a tuple with the UniquePhysicalDataBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniquePhysicalDataBytes

`func (o *StorageDomainStats) SetUniquePhysicalDataBytes(v int64)`

SetUniquePhysicalDataBytes sets UniquePhysicalDataBytes field to given value.

### HasUniquePhysicalDataBytes

`func (o *StorageDomainStats) HasUniquePhysicalDataBytes() bool`

HasUniquePhysicalDataBytes returns a boolean if a field has been set.

### SetUniquePhysicalDataBytesNil

`func (o *StorageDomainStats) SetUniquePhysicalDataBytesNil(b bool)`

 SetUniquePhysicalDataBytesNil sets the value for UniquePhysicalDataBytes to be an explicit nil

### UnsetUniquePhysicalDataBytes
`func (o *StorageDomainStats) UnsetUniquePhysicalDataBytes()`

UnsetUniquePhysicalDataBytes ensures that no value is present for UniquePhysicalDataBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


