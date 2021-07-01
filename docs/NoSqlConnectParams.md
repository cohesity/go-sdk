# NoSqlConnectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CassandraAdditionalParams** | Pointer to [**CassandraAdditionalParams**](CassandraAdditionalParams.md) |  | [optional] 
**CassandraConnectParams** | Pointer to [**CassandraConnectParams**](CassandraConnectParams.md) |  | [optional] 
**CouchbaseConnectParams** | Pointer to [**CouchbaseConnectParams**](CouchbaseConnectParams.md) |  | [optional] 
**HbaseConnectParams** | Pointer to [**HBaseConnectParams**](HBaseConnectParams.md) |  | [optional] 
**HdfsConnectParams** | Pointer to [**HdfsConnectParams**](HdfsConnectParams.md) |  | [optional] 
**HiveConnectParams** | Pointer to [**HiveConnectParams**](HiveConnectParams.md) |  | [optional] 
**MongodbAdditionalParams** | Pointer to [**MongoDBAdditionalParams**](MongoDBAdditionalParams.md) |  | [optional] 
**MongodbConnectParams** | Pointer to [**MongoDBConnectParams**](MongoDBConnectParams.md) |  | [optional] 

## Methods

### NewNoSqlConnectParams

`func NewNoSqlConnectParams() *NoSqlConnectParams`

NewNoSqlConnectParams instantiates a new NoSqlConnectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoSqlConnectParamsWithDefaults

`func NewNoSqlConnectParamsWithDefaults() *NoSqlConnectParams`

NewNoSqlConnectParamsWithDefaults instantiates a new NoSqlConnectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCassandraAdditionalParams

`func (o *NoSqlConnectParams) GetCassandraAdditionalParams() CassandraAdditionalParams`

GetCassandraAdditionalParams returns the CassandraAdditionalParams field if non-nil, zero value otherwise.

### GetCassandraAdditionalParamsOk

`func (o *NoSqlConnectParams) GetCassandraAdditionalParamsOk() (*CassandraAdditionalParams, bool)`

GetCassandraAdditionalParamsOk returns a tuple with the CassandraAdditionalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraAdditionalParams

`func (o *NoSqlConnectParams) SetCassandraAdditionalParams(v CassandraAdditionalParams)`

SetCassandraAdditionalParams sets CassandraAdditionalParams field to given value.

### HasCassandraAdditionalParams

`func (o *NoSqlConnectParams) HasCassandraAdditionalParams() bool`

HasCassandraAdditionalParams returns a boolean if a field has been set.

### GetCassandraConnectParams

`func (o *NoSqlConnectParams) GetCassandraConnectParams() CassandraConnectParams`

GetCassandraConnectParams returns the CassandraConnectParams field if non-nil, zero value otherwise.

### GetCassandraConnectParamsOk

`func (o *NoSqlConnectParams) GetCassandraConnectParamsOk() (*CassandraConnectParams, bool)`

GetCassandraConnectParamsOk returns a tuple with the CassandraConnectParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraConnectParams

`func (o *NoSqlConnectParams) SetCassandraConnectParams(v CassandraConnectParams)`

SetCassandraConnectParams sets CassandraConnectParams field to given value.

### HasCassandraConnectParams

`func (o *NoSqlConnectParams) HasCassandraConnectParams() bool`

HasCassandraConnectParams returns a boolean if a field has been set.

### GetCouchbaseConnectParams

`func (o *NoSqlConnectParams) GetCouchbaseConnectParams() CouchbaseConnectParams`

GetCouchbaseConnectParams returns the CouchbaseConnectParams field if non-nil, zero value otherwise.

### GetCouchbaseConnectParamsOk

`func (o *NoSqlConnectParams) GetCouchbaseConnectParamsOk() (*CouchbaseConnectParams, bool)`

GetCouchbaseConnectParamsOk returns a tuple with the CouchbaseConnectParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseConnectParams

`func (o *NoSqlConnectParams) SetCouchbaseConnectParams(v CouchbaseConnectParams)`

SetCouchbaseConnectParams sets CouchbaseConnectParams field to given value.

### HasCouchbaseConnectParams

`func (o *NoSqlConnectParams) HasCouchbaseConnectParams() bool`

HasCouchbaseConnectParams returns a boolean if a field has been set.

### GetHbaseConnectParams

`func (o *NoSqlConnectParams) GetHbaseConnectParams() HBaseConnectParams`

GetHbaseConnectParams returns the HbaseConnectParams field if non-nil, zero value otherwise.

### GetHbaseConnectParamsOk

`func (o *NoSqlConnectParams) GetHbaseConnectParamsOk() (*HBaseConnectParams, bool)`

GetHbaseConnectParamsOk returns a tuple with the HbaseConnectParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseConnectParams

`func (o *NoSqlConnectParams) SetHbaseConnectParams(v HBaseConnectParams)`

SetHbaseConnectParams sets HbaseConnectParams field to given value.

### HasHbaseConnectParams

`func (o *NoSqlConnectParams) HasHbaseConnectParams() bool`

HasHbaseConnectParams returns a boolean if a field has been set.

### GetHdfsConnectParams

`func (o *NoSqlConnectParams) GetHdfsConnectParams() HdfsConnectParams`

GetHdfsConnectParams returns the HdfsConnectParams field if non-nil, zero value otherwise.

### GetHdfsConnectParamsOk

`func (o *NoSqlConnectParams) GetHdfsConnectParamsOk() (*HdfsConnectParams, bool)`

GetHdfsConnectParamsOk returns a tuple with the HdfsConnectParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsConnectParams

`func (o *NoSqlConnectParams) SetHdfsConnectParams(v HdfsConnectParams)`

SetHdfsConnectParams sets HdfsConnectParams field to given value.

### HasHdfsConnectParams

`func (o *NoSqlConnectParams) HasHdfsConnectParams() bool`

HasHdfsConnectParams returns a boolean if a field has been set.

### GetHiveConnectParams

`func (o *NoSqlConnectParams) GetHiveConnectParams() HiveConnectParams`

GetHiveConnectParams returns the HiveConnectParams field if non-nil, zero value otherwise.

### GetHiveConnectParamsOk

`func (o *NoSqlConnectParams) GetHiveConnectParamsOk() (*HiveConnectParams, bool)`

GetHiveConnectParamsOk returns a tuple with the HiveConnectParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveConnectParams

`func (o *NoSqlConnectParams) SetHiveConnectParams(v HiveConnectParams)`

SetHiveConnectParams sets HiveConnectParams field to given value.

### HasHiveConnectParams

`func (o *NoSqlConnectParams) HasHiveConnectParams() bool`

HasHiveConnectParams returns a boolean if a field has been set.

### GetMongodbAdditionalParams

`func (o *NoSqlConnectParams) GetMongodbAdditionalParams() MongoDBAdditionalParams`

GetMongodbAdditionalParams returns the MongodbAdditionalParams field if non-nil, zero value otherwise.

### GetMongodbAdditionalParamsOk

`func (o *NoSqlConnectParams) GetMongodbAdditionalParamsOk() (*MongoDBAdditionalParams, bool)`

GetMongodbAdditionalParamsOk returns a tuple with the MongodbAdditionalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbAdditionalParams

`func (o *NoSqlConnectParams) SetMongodbAdditionalParams(v MongoDBAdditionalParams)`

SetMongodbAdditionalParams sets MongodbAdditionalParams field to given value.

### HasMongodbAdditionalParams

`func (o *NoSqlConnectParams) HasMongodbAdditionalParams() bool`

HasMongodbAdditionalParams returns a boolean if a field has been set.

### GetMongodbConnectParams

`func (o *NoSqlConnectParams) GetMongodbConnectParams() MongoDBConnectParams`

GetMongodbConnectParams returns the MongodbConnectParams field if non-nil, zero value otherwise.

### GetMongodbConnectParamsOk

`func (o *NoSqlConnectParams) GetMongodbConnectParamsOk() (*MongoDBConnectParams, bool)`

GetMongodbConnectParamsOk returns a tuple with the MongodbConnectParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbConnectParams

`func (o *NoSqlConnectParams) SetMongodbConnectParams(v MongoDBConnectParams)`

SetMongodbConnectParams sets MongodbConnectParams field to given value.

### HasMongodbConnectParams

`func (o *NoSqlConnectParams) HasMongodbConnectParams() bool`

HasMongodbConnectParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


