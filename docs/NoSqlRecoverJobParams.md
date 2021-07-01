# NoSqlRecoverJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandwidthBytesPerSecond** | Pointer to **NullableInt64** | Net bandwidth bytes per second | [optional] 
**CassandraRecoverJobParams** | Pointer to [**CassandraRecoverJobParams**](CassandraRecoverJobParams.md) |  | [optional] 
**Concurrency** | Pointer to **NullableInt32** | Max number of mappers | [optional] 
**CouchbaseRecoverJobParams** | Pointer to [**CouchbaseRecoverJobParams**](CouchbaseRecoverJobParams.md) |  | [optional] 
**HbaseRecoverJobParams** | Pointer to [**HBaseRecoverJobParams**](HBaseRecoverJobParams.md) |  | [optional] 
**HdfsRecoverJobParams** | Pointer to [**HdfsRecoverJobParams**](HdfsRecoverJobParams.md) |  | [optional] 
**HiveRecoverJobParams** | Pointer to [**HiveRecoverJobParams**](HiveRecoverJobParams.md) |  | [optional] 
**MongodbRecoverJobParams** | Pointer to [**MongoDBRecoverJobParams**](MongoDBRecoverJobParams.md) |  | [optional] 
**Overwrite** | Pointer to **NullableBool** | Whether to overwrite or keep the object if the object being recovered already exists in the destination. | [optional] 

## Methods

### NewNoSqlRecoverJobParams

`func NewNoSqlRecoverJobParams() *NoSqlRecoverJobParams`

NewNoSqlRecoverJobParams instantiates a new NoSqlRecoverJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoSqlRecoverJobParamsWithDefaults

`func NewNoSqlRecoverJobParamsWithDefaults() *NoSqlRecoverJobParams`

NewNoSqlRecoverJobParamsWithDefaults instantiates a new NoSqlRecoverJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandwidthBytesPerSecond

`func (o *NoSqlRecoverJobParams) GetBandwidthBytesPerSecond() int64`

GetBandwidthBytesPerSecond returns the BandwidthBytesPerSecond field if non-nil, zero value otherwise.

### GetBandwidthBytesPerSecondOk

`func (o *NoSqlRecoverJobParams) GetBandwidthBytesPerSecondOk() (*int64, bool)`

GetBandwidthBytesPerSecondOk returns a tuple with the BandwidthBytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthBytesPerSecond

`func (o *NoSqlRecoverJobParams) SetBandwidthBytesPerSecond(v int64)`

SetBandwidthBytesPerSecond sets BandwidthBytesPerSecond field to given value.

### HasBandwidthBytesPerSecond

`func (o *NoSqlRecoverJobParams) HasBandwidthBytesPerSecond() bool`

HasBandwidthBytesPerSecond returns a boolean if a field has been set.

### SetBandwidthBytesPerSecondNil

`func (o *NoSqlRecoverJobParams) SetBandwidthBytesPerSecondNil(b bool)`

 SetBandwidthBytesPerSecondNil sets the value for BandwidthBytesPerSecond to be an explicit nil

### UnsetBandwidthBytesPerSecond
`func (o *NoSqlRecoverJobParams) UnsetBandwidthBytesPerSecond()`

UnsetBandwidthBytesPerSecond ensures that no value is present for BandwidthBytesPerSecond, not even an explicit nil
### GetCassandraRecoverJobParams

`func (o *NoSqlRecoverJobParams) GetCassandraRecoverJobParams() CassandraRecoverJobParams`

GetCassandraRecoverJobParams returns the CassandraRecoverJobParams field if non-nil, zero value otherwise.

### GetCassandraRecoverJobParamsOk

`func (o *NoSqlRecoverJobParams) GetCassandraRecoverJobParamsOk() (*CassandraRecoverJobParams, bool)`

GetCassandraRecoverJobParamsOk returns a tuple with the CassandraRecoverJobParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraRecoverJobParams

`func (o *NoSqlRecoverJobParams) SetCassandraRecoverJobParams(v CassandraRecoverJobParams)`

SetCassandraRecoverJobParams sets CassandraRecoverJobParams field to given value.

### HasCassandraRecoverJobParams

`func (o *NoSqlRecoverJobParams) HasCassandraRecoverJobParams() bool`

HasCassandraRecoverJobParams returns a boolean if a field has been set.

### GetConcurrency

`func (o *NoSqlRecoverJobParams) GetConcurrency() int32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *NoSqlRecoverJobParams) GetConcurrencyOk() (*int32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *NoSqlRecoverJobParams) SetConcurrency(v int32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *NoSqlRecoverJobParams) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### SetConcurrencyNil

`func (o *NoSqlRecoverJobParams) SetConcurrencyNil(b bool)`

 SetConcurrencyNil sets the value for Concurrency to be an explicit nil

### UnsetConcurrency
`func (o *NoSqlRecoverJobParams) UnsetConcurrency()`

UnsetConcurrency ensures that no value is present for Concurrency, not even an explicit nil
### GetCouchbaseRecoverJobParams

`func (o *NoSqlRecoverJobParams) GetCouchbaseRecoverJobParams() CouchbaseRecoverJobParams`

GetCouchbaseRecoverJobParams returns the CouchbaseRecoverJobParams field if non-nil, zero value otherwise.

### GetCouchbaseRecoverJobParamsOk

`func (o *NoSqlRecoverJobParams) GetCouchbaseRecoverJobParamsOk() (*CouchbaseRecoverJobParams, bool)`

GetCouchbaseRecoverJobParamsOk returns a tuple with the CouchbaseRecoverJobParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseRecoverJobParams

`func (o *NoSqlRecoverJobParams) SetCouchbaseRecoverJobParams(v CouchbaseRecoverJobParams)`

SetCouchbaseRecoverJobParams sets CouchbaseRecoverJobParams field to given value.

### HasCouchbaseRecoverJobParams

`func (o *NoSqlRecoverJobParams) HasCouchbaseRecoverJobParams() bool`

HasCouchbaseRecoverJobParams returns a boolean if a field has been set.

### GetHbaseRecoverJobParams

`func (o *NoSqlRecoverJobParams) GetHbaseRecoverJobParams() HBaseRecoverJobParams`

GetHbaseRecoverJobParams returns the HbaseRecoverJobParams field if non-nil, zero value otherwise.

### GetHbaseRecoverJobParamsOk

`func (o *NoSqlRecoverJobParams) GetHbaseRecoverJobParamsOk() (*HBaseRecoverJobParams, bool)`

GetHbaseRecoverJobParamsOk returns a tuple with the HbaseRecoverJobParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseRecoverJobParams

`func (o *NoSqlRecoverJobParams) SetHbaseRecoverJobParams(v HBaseRecoverJobParams)`

SetHbaseRecoverJobParams sets HbaseRecoverJobParams field to given value.

### HasHbaseRecoverJobParams

`func (o *NoSqlRecoverJobParams) HasHbaseRecoverJobParams() bool`

HasHbaseRecoverJobParams returns a boolean if a field has been set.

### GetHdfsRecoverJobParams

`func (o *NoSqlRecoverJobParams) GetHdfsRecoverJobParams() HdfsRecoverJobParams`

GetHdfsRecoverJobParams returns the HdfsRecoverJobParams field if non-nil, zero value otherwise.

### GetHdfsRecoverJobParamsOk

`func (o *NoSqlRecoverJobParams) GetHdfsRecoverJobParamsOk() (*HdfsRecoverJobParams, bool)`

GetHdfsRecoverJobParamsOk returns a tuple with the HdfsRecoverJobParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsRecoverJobParams

`func (o *NoSqlRecoverJobParams) SetHdfsRecoverJobParams(v HdfsRecoverJobParams)`

SetHdfsRecoverJobParams sets HdfsRecoverJobParams field to given value.

### HasHdfsRecoverJobParams

`func (o *NoSqlRecoverJobParams) HasHdfsRecoverJobParams() bool`

HasHdfsRecoverJobParams returns a boolean if a field has been set.

### GetHiveRecoverJobParams

`func (o *NoSqlRecoverJobParams) GetHiveRecoverJobParams() HiveRecoverJobParams`

GetHiveRecoverJobParams returns the HiveRecoverJobParams field if non-nil, zero value otherwise.

### GetHiveRecoverJobParamsOk

`func (o *NoSqlRecoverJobParams) GetHiveRecoverJobParamsOk() (*HiveRecoverJobParams, bool)`

GetHiveRecoverJobParamsOk returns a tuple with the HiveRecoverJobParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveRecoverJobParams

`func (o *NoSqlRecoverJobParams) SetHiveRecoverJobParams(v HiveRecoverJobParams)`

SetHiveRecoverJobParams sets HiveRecoverJobParams field to given value.

### HasHiveRecoverJobParams

`func (o *NoSqlRecoverJobParams) HasHiveRecoverJobParams() bool`

HasHiveRecoverJobParams returns a boolean if a field has been set.

### GetMongodbRecoverJobParams

`func (o *NoSqlRecoverJobParams) GetMongodbRecoverJobParams() MongoDBRecoverJobParams`

GetMongodbRecoverJobParams returns the MongodbRecoverJobParams field if non-nil, zero value otherwise.

### GetMongodbRecoverJobParamsOk

`func (o *NoSqlRecoverJobParams) GetMongodbRecoverJobParamsOk() (*MongoDBRecoverJobParams, bool)`

GetMongodbRecoverJobParamsOk returns a tuple with the MongodbRecoverJobParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbRecoverJobParams

`func (o *NoSqlRecoverJobParams) SetMongodbRecoverJobParams(v MongoDBRecoverJobParams)`

SetMongodbRecoverJobParams sets MongodbRecoverJobParams field to given value.

### HasMongodbRecoverJobParams

`func (o *NoSqlRecoverJobParams) HasMongodbRecoverJobParams() bool`

HasMongodbRecoverJobParams returns a boolean if a field has been set.

### GetOverwrite

`func (o *NoSqlRecoverJobParams) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *NoSqlRecoverJobParams) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *NoSqlRecoverJobParams) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *NoSqlRecoverJobParams) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### SetOverwriteNil

`func (o *NoSqlRecoverJobParams) SetOverwriteNil(b bool)`

 SetOverwriteNil sets the value for Overwrite to be an explicit nil

### UnsetOverwrite
`func (o *NoSqlRecoverJobParams) UnsetOverwrite()`

UnsetOverwrite ensures that no value is present for Overwrite, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


