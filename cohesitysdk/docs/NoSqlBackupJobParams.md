# NoSqlBackupJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandwidthBytesPerSecond** | Pointer to **NullableInt64** | Net bandwidth bytes per second. | [optional] 
**CassandraBackupJobParams** | Pointer to [**CassandraBackupJobParams**](CassandraBackupJobParams.md) |  | [optional] 
**CompactionJobIntervalSecs** | Pointer to **NullableInt64** | Frequency at which compaction jobs should run in seconds. Will be only applicable for Cassandra, Mongo and Couchbase environment. | [optional] 
**Concurrency** | Pointer to **NullableInt32** | Max number of mappers. | [optional] 
**CouchbaseBackupJobParams** | Pointer to **map[string]interface{}** | Contains any additional couchbase environment specific backup params at the job level. | [optional] 
**GcJobIntervalSecs** | Pointer to **NullableInt64** | Frequency at which garbage collection jobs should run in seconds. | [optional] 
**GcRetentionPeriodDays** | Pointer to **NullableInt32** | Retention period for logs of this job in days. | [optional] 
**HbaseBackupJobParams** | Pointer to [**HBaseBackupJobParams**](HBaseBackupJobParams.md) |  | [optional] 
**HdfsBackupJobParams** | Pointer to [**HdfsBackupJobParams**](HdfsBackupJobParams.md) |  | [optional] 
**HiveBackupJobParams** | Pointer to [**HiveBackupJobParams**](HiveBackupJobParams.md) |  | [optional] 
**LastCompactionRunTimeUsecs** | Pointer to **NullableInt64** | The last time (in usecs) when the compaction ran for this jobs. | [optional] 
**LastGcRunTimeUsecs** | Pointer to **NullableInt64** | The last time (in usecs) when the gc ran for this jobs. | [optional] 
**MongodbBackupJobParams** | Pointer to [**MongoDBBackupJobParams**](MongoDBBackupJobParams.md) |  | [optional] 

## Methods

### NewNoSqlBackupJobParams

`func NewNoSqlBackupJobParams() *NoSqlBackupJobParams`

NewNoSqlBackupJobParams instantiates a new NoSqlBackupJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoSqlBackupJobParamsWithDefaults

`func NewNoSqlBackupJobParamsWithDefaults() *NoSqlBackupJobParams`

NewNoSqlBackupJobParamsWithDefaults instantiates a new NoSqlBackupJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandwidthBytesPerSecond

`func (o *NoSqlBackupJobParams) GetBandwidthBytesPerSecond() int64`

GetBandwidthBytesPerSecond returns the BandwidthBytesPerSecond field if non-nil, zero value otherwise.

### GetBandwidthBytesPerSecondOk

`func (o *NoSqlBackupJobParams) GetBandwidthBytesPerSecondOk() (*int64, bool)`

GetBandwidthBytesPerSecondOk returns a tuple with the BandwidthBytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthBytesPerSecond

`func (o *NoSqlBackupJobParams) SetBandwidthBytesPerSecond(v int64)`

SetBandwidthBytesPerSecond sets BandwidthBytesPerSecond field to given value.

### HasBandwidthBytesPerSecond

`func (o *NoSqlBackupJobParams) HasBandwidthBytesPerSecond() bool`

HasBandwidthBytesPerSecond returns a boolean if a field has been set.

### SetBandwidthBytesPerSecondNil

`func (o *NoSqlBackupJobParams) SetBandwidthBytesPerSecondNil(b bool)`

 SetBandwidthBytesPerSecondNil sets the value for BandwidthBytesPerSecond to be an explicit nil

### UnsetBandwidthBytesPerSecond
`func (o *NoSqlBackupJobParams) UnsetBandwidthBytesPerSecond()`

UnsetBandwidthBytesPerSecond ensures that no value is present for BandwidthBytesPerSecond, not even an explicit nil
### GetCassandraBackupJobParams

`func (o *NoSqlBackupJobParams) GetCassandraBackupJobParams() CassandraBackupJobParams`

GetCassandraBackupJobParams returns the CassandraBackupJobParams field if non-nil, zero value otherwise.

### GetCassandraBackupJobParamsOk

`func (o *NoSqlBackupJobParams) GetCassandraBackupJobParamsOk() (*CassandraBackupJobParams, bool)`

GetCassandraBackupJobParamsOk returns a tuple with the CassandraBackupJobParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraBackupJobParams

`func (o *NoSqlBackupJobParams) SetCassandraBackupJobParams(v CassandraBackupJobParams)`

SetCassandraBackupJobParams sets CassandraBackupJobParams field to given value.

### HasCassandraBackupJobParams

`func (o *NoSqlBackupJobParams) HasCassandraBackupJobParams() bool`

HasCassandraBackupJobParams returns a boolean if a field has been set.

### GetCompactionJobIntervalSecs

`func (o *NoSqlBackupJobParams) GetCompactionJobIntervalSecs() int64`

GetCompactionJobIntervalSecs returns the CompactionJobIntervalSecs field if non-nil, zero value otherwise.

### GetCompactionJobIntervalSecsOk

`func (o *NoSqlBackupJobParams) GetCompactionJobIntervalSecsOk() (*int64, bool)`

GetCompactionJobIntervalSecsOk returns a tuple with the CompactionJobIntervalSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactionJobIntervalSecs

`func (o *NoSqlBackupJobParams) SetCompactionJobIntervalSecs(v int64)`

SetCompactionJobIntervalSecs sets CompactionJobIntervalSecs field to given value.

### HasCompactionJobIntervalSecs

`func (o *NoSqlBackupJobParams) HasCompactionJobIntervalSecs() bool`

HasCompactionJobIntervalSecs returns a boolean if a field has been set.

### SetCompactionJobIntervalSecsNil

`func (o *NoSqlBackupJobParams) SetCompactionJobIntervalSecsNil(b bool)`

 SetCompactionJobIntervalSecsNil sets the value for CompactionJobIntervalSecs to be an explicit nil

### UnsetCompactionJobIntervalSecs
`func (o *NoSqlBackupJobParams) UnsetCompactionJobIntervalSecs()`

UnsetCompactionJobIntervalSecs ensures that no value is present for CompactionJobIntervalSecs, not even an explicit nil
### GetConcurrency

`func (o *NoSqlBackupJobParams) GetConcurrency() int32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *NoSqlBackupJobParams) GetConcurrencyOk() (*int32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *NoSqlBackupJobParams) SetConcurrency(v int32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *NoSqlBackupJobParams) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### SetConcurrencyNil

`func (o *NoSqlBackupJobParams) SetConcurrencyNil(b bool)`

 SetConcurrencyNil sets the value for Concurrency to be an explicit nil

### UnsetConcurrency
`func (o *NoSqlBackupJobParams) UnsetConcurrency()`

UnsetConcurrency ensures that no value is present for Concurrency, not even an explicit nil
### GetCouchbaseBackupJobParams

`func (o *NoSqlBackupJobParams) GetCouchbaseBackupJobParams() map[string]interface{}`

GetCouchbaseBackupJobParams returns the CouchbaseBackupJobParams field if non-nil, zero value otherwise.

### GetCouchbaseBackupJobParamsOk

`func (o *NoSqlBackupJobParams) GetCouchbaseBackupJobParamsOk() (*map[string]interface{}, bool)`

GetCouchbaseBackupJobParamsOk returns a tuple with the CouchbaseBackupJobParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseBackupJobParams

`func (o *NoSqlBackupJobParams) SetCouchbaseBackupJobParams(v map[string]interface{})`

SetCouchbaseBackupJobParams sets CouchbaseBackupJobParams field to given value.

### HasCouchbaseBackupJobParams

`func (o *NoSqlBackupJobParams) HasCouchbaseBackupJobParams() bool`

HasCouchbaseBackupJobParams returns a boolean if a field has been set.

### GetGcJobIntervalSecs

`func (o *NoSqlBackupJobParams) GetGcJobIntervalSecs() int64`

GetGcJobIntervalSecs returns the GcJobIntervalSecs field if non-nil, zero value otherwise.

### GetGcJobIntervalSecsOk

`func (o *NoSqlBackupJobParams) GetGcJobIntervalSecsOk() (*int64, bool)`

GetGcJobIntervalSecsOk returns a tuple with the GcJobIntervalSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcJobIntervalSecs

`func (o *NoSqlBackupJobParams) SetGcJobIntervalSecs(v int64)`

SetGcJobIntervalSecs sets GcJobIntervalSecs field to given value.

### HasGcJobIntervalSecs

`func (o *NoSqlBackupJobParams) HasGcJobIntervalSecs() bool`

HasGcJobIntervalSecs returns a boolean if a field has been set.

### SetGcJobIntervalSecsNil

`func (o *NoSqlBackupJobParams) SetGcJobIntervalSecsNil(b bool)`

 SetGcJobIntervalSecsNil sets the value for GcJobIntervalSecs to be an explicit nil

### UnsetGcJobIntervalSecs
`func (o *NoSqlBackupJobParams) UnsetGcJobIntervalSecs()`

UnsetGcJobIntervalSecs ensures that no value is present for GcJobIntervalSecs, not even an explicit nil
### GetGcRetentionPeriodDays

`func (o *NoSqlBackupJobParams) GetGcRetentionPeriodDays() int32`

GetGcRetentionPeriodDays returns the GcRetentionPeriodDays field if non-nil, zero value otherwise.

### GetGcRetentionPeriodDaysOk

`func (o *NoSqlBackupJobParams) GetGcRetentionPeriodDaysOk() (*int32, bool)`

GetGcRetentionPeriodDaysOk returns a tuple with the GcRetentionPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcRetentionPeriodDays

`func (o *NoSqlBackupJobParams) SetGcRetentionPeriodDays(v int32)`

SetGcRetentionPeriodDays sets GcRetentionPeriodDays field to given value.

### HasGcRetentionPeriodDays

`func (o *NoSqlBackupJobParams) HasGcRetentionPeriodDays() bool`

HasGcRetentionPeriodDays returns a boolean if a field has been set.

### SetGcRetentionPeriodDaysNil

`func (o *NoSqlBackupJobParams) SetGcRetentionPeriodDaysNil(b bool)`

 SetGcRetentionPeriodDaysNil sets the value for GcRetentionPeriodDays to be an explicit nil

### UnsetGcRetentionPeriodDays
`func (o *NoSqlBackupJobParams) UnsetGcRetentionPeriodDays()`

UnsetGcRetentionPeriodDays ensures that no value is present for GcRetentionPeriodDays, not even an explicit nil
### GetHbaseBackupJobParams

`func (o *NoSqlBackupJobParams) GetHbaseBackupJobParams() HBaseBackupJobParams`

GetHbaseBackupJobParams returns the HbaseBackupJobParams field if non-nil, zero value otherwise.

### GetHbaseBackupJobParamsOk

`func (o *NoSqlBackupJobParams) GetHbaseBackupJobParamsOk() (*HBaseBackupJobParams, bool)`

GetHbaseBackupJobParamsOk returns a tuple with the HbaseBackupJobParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseBackupJobParams

`func (o *NoSqlBackupJobParams) SetHbaseBackupJobParams(v HBaseBackupJobParams)`

SetHbaseBackupJobParams sets HbaseBackupJobParams field to given value.

### HasHbaseBackupJobParams

`func (o *NoSqlBackupJobParams) HasHbaseBackupJobParams() bool`

HasHbaseBackupJobParams returns a boolean if a field has been set.

### GetHdfsBackupJobParams

`func (o *NoSqlBackupJobParams) GetHdfsBackupJobParams() HdfsBackupJobParams`

GetHdfsBackupJobParams returns the HdfsBackupJobParams field if non-nil, zero value otherwise.

### GetHdfsBackupJobParamsOk

`func (o *NoSqlBackupJobParams) GetHdfsBackupJobParamsOk() (*HdfsBackupJobParams, bool)`

GetHdfsBackupJobParamsOk returns a tuple with the HdfsBackupJobParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsBackupJobParams

`func (o *NoSqlBackupJobParams) SetHdfsBackupJobParams(v HdfsBackupJobParams)`

SetHdfsBackupJobParams sets HdfsBackupJobParams field to given value.

### HasHdfsBackupJobParams

`func (o *NoSqlBackupJobParams) HasHdfsBackupJobParams() bool`

HasHdfsBackupJobParams returns a boolean if a field has been set.

### GetHiveBackupJobParams

`func (o *NoSqlBackupJobParams) GetHiveBackupJobParams() HiveBackupJobParams`

GetHiveBackupJobParams returns the HiveBackupJobParams field if non-nil, zero value otherwise.

### GetHiveBackupJobParamsOk

`func (o *NoSqlBackupJobParams) GetHiveBackupJobParamsOk() (*HiveBackupJobParams, bool)`

GetHiveBackupJobParamsOk returns a tuple with the HiveBackupJobParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveBackupJobParams

`func (o *NoSqlBackupJobParams) SetHiveBackupJobParams(v HiveBackupJobParams)`

SetHiveBackupJobParams sets HiveBackupJobParams field to given value.

### HasHiveBackupJobParams

`func (o *NoSqlBackupJobParams) HasHiveBackupJobParams() bool`

HasHiveBackupJobParams returns a boolean if a field has been set.

### GetLastCompactionRunTimeUsecs

`func (o *NoSqlBackupJobParams) GetLastCompactionRunTimeUsecs() int64`

GetLastCompactionRunTimeUsecs returns the LastCompactionRunTimeUsecs field if non-nil, zero value otherwise.

### GetLastCompactionRunTimeUsecsOk

`func (o *NoSqlBackupJobParams) GetLastCompactionRunTimeUsecsOk() (*int64, bool)`

GetLastCompactionRunTimeUsecsOk returns a tuple with the LastCompactionRunTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCompactionRunTimeUsecs

`func (o *NoSqlBackupJobParams) SetLastCompactionRunTimeUsecs(v int64)`

SetLastCompactionRunTimeUsecs sets LastCompactionRunTimeUsecs field to given value.

### HasLastCompactionRunTimeUsecs

`func (o *NoSqlBackupJobParams) HasLastCompactionRunTimeUsecs() bool`

HasLastCompactionRunTimeUsecs returns a boolean if a field has been set.

### SetLastCompactionRunTimeUsecsNil

`func (o *NoSqlBackupJobParams) SetLastCompactionRunTimeUsecsNil(b bool)`

 SetLastCompactionRunTimeUsecsNil sets the value for LastCompactionRunTimeUsecs to be an explicit nil

### UnsetLastCompactionRunTimeUsecs
`func (o *NoSqlBackupJobParams) UnsetLastCompactionRunTimeUsecs()`

UnsetLastCompactionRunTimeUsecs ensures that no value is present for LastCompactionRunTimeUsecs, not even an explicit nil
### GetLastGcRunTimeUsecs

`func (o *NoSqlBackupJobParams) GetLastGcRunTimeUsecs() int64`

GetLastGcRunTimeUsecs returns the LastGcRunTimeUsecs field if non-nil, zero value otherwise.

### GetLastGcRunTimeUsecsOk

`func (o *NoSqlBackupJobParams) GetLastGcRunTimeUsecsOk() (*int64, bool)`

GetLastGcRunTimeUsecsOk returns a tuple with the LastGcRunTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastGcRunTimeUsecs

`func (o *NoSqlBackupJobParams) SetLastGcRunTimeUsecs(v int64)`

SetLastGcRunTimeUsecs sets LastGcRunTimeUsecs field to given value.

### HasLastGcRunTimeUsecs

`func (o *NoSqlBackupJobParams) HasLastGcRunTimeUsecs() bool`

HasLastGcRunTimeUsecs returns a boolean if a field has been set.

### SetLastGcRunTimeUsecsNil

`func (o *NoSqlBackupJobParams) SetLastGcRunTimeUsecsNil(b bool)`

 SetLastGcRunTimeUsecsNil sets the value for LastGcRunTimeUsecs to be an explicit nil

### UnsetLastGcRunTimeUsecs
`func (o *NoSqlBackupJobParams) UnsetLastGcRunTimeUsecs()`

UnsetLastGcRunTimeUsecs ensures that no value is present for LastGcRunTimeUsecs, not even an explicit nil
### GetMongodbBackupJobParams

`func (o *NoSqlBackupJobParams) GetMongodbBackupJobParams() MongoDBBackupJobParams`

GetMongodbBackupJobParams returns the MongodbBackupJobParams field if non-nil, zero value otherwise.

### GetMongodbBackupJobParamsOk

`func (o *NoSqlBackupJobParams) GetMongodbBackupJobParamsOk() (*MongoDBBackupJobParams, bool)`

GetMongodbBackupJobParamsOk returns a tuple with the MongodbBackupJobParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbBackupJobParams

`func (o *NoSqlBackupJobParams) SetMongodbBackupJobParams(v MongoDBBackupJobParams)`

SetMongodbBackupJobParams sets MongodbBackupJobParams field to given value.

### HasMongodbBackupJobParams

`func (o *NoSqlBackupJobParams) HasMongodbBackupJobParams() bool`

HasMongodbBackupJobParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


