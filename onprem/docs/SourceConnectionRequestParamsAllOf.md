# SourceConnectionRequestParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CassandraConnectionParams** | Pointer to [**CassandraConnectionParams**](CassandraConnectionParams.md) |  | [optional] 
**HiveConnectionParams** | Pointer to [**HadoopConnectionParams**](HadoopConnectionParams.md) |  | [optional] 
**HbaseConnectionParams** | Pointer to [**HadoopConnectionParams**](HadoopConnectionParams.md) |  | [optional] 
**HdfsConnectionParams** | Pointer to [**HadoopConnectionParams**](HadoopConnectionParams.md) |  | [optional] 
**MssqlConnectionParams** | Pointer to [**MssqlConnectionParams**](MssqlConnectionParams.md) |  | [optional] 
**OracleConnectionParams** | Pointer to [**OracleConnectionParams**](OracleConnectionParams.md) |  | [optional] 
**VmwareConnectionParams** | Pointer to [**VmwareConnectionParams**](VmwareConnectionParams.md) |  | [optional] 

## Methods

### NewSourceConnectionRequestParamsAllOf

`func NewSourceConnectionRequestParamsAllOf() *SourceConnectionRequestParamsAllOf`

NewSourceConnectionRequestParamsAllOf instantiates a new SourceConnectionRequestParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceConnectionRequestParamsAllOfWithDefaults

`func NewSourceConnectionRequestParamsAllOfWithDefaults() *SourceConnectionRequestParamsAllOf`

NewSourceConnectionRequestParamsAllOfWithDefaults instantiates a new SourceConnectionRequestParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCassandraConnectionParams

`func (o *SourceConnectionRequestParamsAllOf) GetCassandraConnectionParams() CassandraConnectionParams`

GetCassandraConnectionParams returns the CassandraConnectionParams field if non-nil, zero value otherwise.

### GetCassandraConnectionParamsOk

`func (o *SourceConnectionRequestParamsAllOf) GetCassandraConnectionParamsOk() (*CassandraConnectionParams, bool)`

GetCassandraConnectionParamsOk returns a tuple with the CassandraConnectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraConnectionParams

`func (o *SourceConnectionRequestParamsAllOf) SetCassandraConnectionParams(v CassandraConnectionParams)`

SetCassandraConnectionParams sets CassandraConnectionParams field to given value.

### HasCassandraConnectionParams

`func (o *SourceConnectionRequestParamsAllOf) HasCassandraConnectionParams() bool`

HasCassandraConnectionParams returns a boolean if a field has been set.

### GetHiveConnectionParams

`func (o *SourceConnectionRequestParamsAllOf) GetHiveConnectionParams() HadoopConnectionParams`

GetHiveConnectionParams returns the HiveConnectionParams field if non-nil, zero value otherwise.

### GetHiveConnectionParamsOk

`func (o *SourceConnectionRequestParamsAllOf) GetHiveConnectionParamsOk() (*HadoopConnectionParams, bool)`

GetHiveConnectionParamsOk returns a tuple with the HiveConnectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveConnectionParams

`func (o *SourceConnectionRequestParamsAllOf) SetHiveConnectionParams(v HadoopConnectionParams)`

SetHiveConnectionParams sets HiveConnectionParams field to given value.

### HasHiveConnectionParams

`func (o *SourceConnectionRequestParamsAllOf) HasHiveConnectionParams() bool`

HasHiveConnectionParams returns a boolean if a field has been set.

### GetHbaseConnectionParams

`func (o *SourceConnectionRequestParamsAllOf) GetHbaseConnectionParams() HadoopConnectionParams`

GetHbaseConnectionParams returns the HbaseConnectionParams field if non-nil, zero value otherwise.

### GetHbaseConnectionParamsOk

`func (o *SourceConnectionRequestParamsAllOf) GetHbaseConnectionParamsOk() (*HadoopConnectionParams, bool)`

GetHbaseConnectionParamsOk returns a tuple with the HbaseConnectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseConnectionParams

`func (o *SourceConnectionRequestParamsAllOf) SetHbaseConnectionParams(v HadoopConnectionParams)`

SetHbaseConnectionParams sets HbaseConnectionParams field to given value.

### HasHbaseConnectionParams

`func (o *SourceConnectionRequestParamsAllOf) HasHbaseConnectionParams() bool`

HasHbaseConnectionParams returns a boolean if a field has been set.

### GetHdfsConnectionParams

`func (o *SourceConnectionRequestParamsAllOf) GetHdfsConnectionParams() HadoopConnectionParams`

GetHdfsConnectionParams returns the HdfsConnectionParams field if non-nil, zero value otherwise.

### GetHdfsConnectionParamsOk

`func (o *SourceConnectionRequestParamsAllOf) GetHdfsConnectionParamsOk() (*HadoopConnectionParams, bool)`

GetHdfsConnectionParamsOk returns a tuple with the HdfsConnectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsConnectionParams

`func (o *SourceConnectionRequestParamsAllOf) SetHdfsConnectionParams(v HadoopConnectionParams)`

SetHdfsConnectionParams sets HdfsConnectionParams field to given value.

### HasHdfsConnectionParams

`func (o *SourceConnectionRequestParamsAllOf) HasHdfsConnectionParams() bool`

HasHdfsConnectionParams returns a boolean if a field has been set.

### GetMssqlConnectionParams

`func (o *SourceConnectionRequestParamsAllOf) GetMssqlConnectionParams() MssqlConnectionParams`

GetMssqlConnectionParams returns the MssqlConnectionParams field if non-nil, zero value otherwise.

### GetMssqlConnectionParamsOk

`func (o *SourceConnectionRequestParamsAllOf) GetMssqlConnectionParamsOk() (*MssqlConnectionParams, bool)`

GetMssqlConnectionParamsOk returns a tuple with the MssqlConnectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlConnectionParams

`func (o *SourceConnectionRequestParamsAllOf) SetMssqlConnectionParams(v MssqlConnectionParams)`

SetMssqlConnectionParams sets MssqlConnectionParams field to given value.

### HasMssqlConnectionParams

`func (o *SourceConnectionRequestParamsAllOf) HasMssqlConnectionParams() bool`

HasMssqlConnectionParams returns a boolean if a field has been set.

### GetOracleConnectionParams

`func (o *SourceConnectionRequestParamsAllOf) GetOracleConnectionParams() OracleConnectionParams`

GetOracleConnectionParams returns the OracleConnectionParams field if non-nil, zero value otherwise.

### GetOracleConnectionParamsOk

`func (o *SourceConnectionRequestParamsAllOf) GetOracleConnectionParamsOk() (*OracleConnectionParams, bool)`

GetOracleConnectionParamsOk returns a tuple with the OracleConnectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleConnectionParams

`func (o *SourceConnectionRequestParamsAllOf) SetOracleConnectionParams(v OracleConnectionParams)`

SetOracleConnectionParams sets OracleConnectionParams field to given value.

### HasOracleConnectionParams

`func (o *SourceConnectionRequestParamsAllOf) HasOracleConnectionParams() bool`

HasOracleConnectionParams returns a boolean if a field has been set.

### GetVmwareConnectionParams

`func (o *SourceConnectionRequestParamsAllOf) GetVmwareConnectionParams() VmwareConnectionParams`

GetVmwareConnectionParams returns the VmwareConnectionParams field if non-nil, zero value otherwise.

### GetVmwareConnectionParamsOk

`func (o *SourceConnectionRequestParamsAllOf) GetVmwareConnectionParamsOk() (*VmwareConnectionParams, bool)`

GetVmwareConnectionParamsOk returns a tuple with the VmwareConnectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareConnectionParams

`func (o *SourceConnectionRequestParamsAllOf) SetVmwareConnectionParams(v VmwareConnectionParams)`

SetVmwareConnectionParams sets VmwareConnectionParams field to given value.

### HasVmwareConnectionParams

`func (o *SourceConnectionRequestParamsAllOf) HasVmwareConnectionParams() bool`

HasVmwareConnectionParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


