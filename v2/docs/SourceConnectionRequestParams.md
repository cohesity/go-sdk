# SourceConnectionRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | Pointer to **NullableInt64** | Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user. | [optional] 
**Environment** | **NullableString** | Specifies the environment type of the Protection Source. | 
**CassandraConnectionParams** | Pointer to [**CassandraConnectionParams**](CassandraConnectionParams.md) |  | [optional] 
**HbaseConnectionParams** | Pointer to [**HadoopConnectionParams**](HadoopConnectionParams.md) |  | [optional] 
**HdfsConnectionParams** | Pointer to [**HadoopConnectionParams**](HadoopConnectionParams.md) |  | [optional] 
**HiveConnectionParams** | Pointer to [**HadoopConnectionParams**](HadoopConnectionParams.md) |  | [optional] 
**MssqlConnectionParams** | Pointer to [**MssqlConnectionParams**](MssqlConnectionParams.md) |  | [optional] 
**OracleConnectionParams** | Pointer to [**OracleConnectionParams**](OracleConnectionParams.md) |  | [optional] 
**VmwareConnectionParams** | Pointer to [**VmwareConnectionParams**](VmwareConnectionParams.md) |  | [optional] 

## Methods

### NewSourceConnectionRequestParams

`func NewSourceConnectionRequestParams(environment NullableString, ) *SourceConnectionRequestParams`

NewSourceConnectionRequestParams instantiates a new SourceConnectionRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceConnectionRequestParamsWithDefaults

`func NewSourceConnectionRequestParamsWithDefaults() *SourceConnectionRequestParams`

NewSourceConnectionRequestParamsWithDefaults instantiates a new SourceConnectionRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *SourceConnectionRequestParams) GetConnectionId() int64`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *SourceConnectionRequestParams) GetConnectionIdOk() (*int64, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *SourceConnectionRequestParams) SetConnectionId(v int64)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *SourceConnectionRequestParams) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### SetConnectionIdNil

`func (o *SourceConnectionRequestParams) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *SourceConnectionRequestParams) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetEnvironment

`func (o *SourceConnectionRequestParams) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SourceConnectionRequestParams) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SourceConnectionRequestParams) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### SetEnvironmentNil

`func (o *SourceConnectionRequestParams) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *SourceConnectionRequestParams) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetCassandraConnectionParams

`func (o *SourceConnectionRequestParams) GetCassandraConnectionParams() CassandraConnectionParams`

GetCassandraConnectionParams returns the CassandraConnectionParams field if non-nil, zero value otherwise.

### GetCassandraConnectionParamsOk

`func (o *SourceConnectionRequestParams) GetCassandraConnectionParamsOk() (*CassandraConnectionParams, bool)`

GetCassandraConnectionParamsOk returns a tuple with the CassandraConnectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraConnectionParams

`func (o *SourceConnectionRequestParams) SetCassandraConnectionParams(v CassandraConnectionParams)`

SetCassandraConnectionParams sets CassandraConnectionParams field to given value.

### HasCassandraConnectionParams

`func (o *SourceConnectionRequestParams) HasCassandraConnectionParams() bool`

HasCassandraConnectionParams returns a boolean if a field has been set.

### GetHbaseConnectionParams

`func (o *SourceConnectionRequestParams) GetHbaseConnectionParams() HadoopConnectionParams`

GetHbaseConnectionParams returns the HbaseConnectionParams field if non-nil, zero value otherwise.

### GetHbaseConnectionParamsOk

`func (o *SourceConnectionRequestParams) GetHbaseConnectionParamsOk() (*HadoopConnectionParams, bool)`

GetHbaseConnectionParamsOk returns a tuple with the HbaseConnectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseConnectionParams

`func (o *SourceConnectionRequestParams) SetHbaseConnectionParams(v HadoopConnectionParams)`

SetHbaseConnectionParams sets HbaseConnectionParams field to given value.

### HasHbaseConnectionParams

`func (o *SourceConnectionRequestParams) HasHbaseConnectionParams() bool`

HasHbaseConnectionParams returns a boolean if a field has been set.

### GetHdfsConnectionParams

`func (o *SourceConnectionRequestParams) GetHdfsConnectionParams() HadoopConnectionParams`

GetHdfsConnectionParams returns the HdfsConnectionParams field if non-nil, zero value otherwise.

### GetHdfsConnectionParamsOk

`func (o *SourceConnectionRequestParams) GetHdfsConnectionParamsOk() (*HadoopConnectionParams, bool)`

GetHdfsConnectionParamsOk returns a tuple with the HdfsConnectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsConnectionParams

`func (o *SourceConnectionRequestParams) SetHdfsConnectionParams(v HadoopConnectionParams)`

SetHdfsConnectionParams sets HdfsConnectionParams field to given value.

### HasHdfsConnectionParams

`func (o *SourceConnectionRequestParams) HasHdfsConnectionParams() bool`

HasHdfsConnectionParams returns a boolean if a field has been set.

### GetHiveConnectionParams

`func (o *SourceConnectionRequestParams) GetHiveConnectionParams() HadoopConnectionParams`

GetHiveConnectionParams returns the HiveConnectionParams field if non-nil, zero value otherwise.

### GetHiveConnectionParamsOk

`func (o *SourceConnectionRequestParams) GetHiveConnectionParamsOk() (*HadoopConnectionParams, bool)`

GetHiveConnectionParamsOk returns a tuple with the HiveConnectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveConnectionParams

`func (o *SourceConnectionRequestParams) SetHiveConnectionParams(v HadoopConnectionParams)`

SetHiveConnectionParams sets HiveConnectionParams field to given value.

### HasHiveConnectionParams

`func (o *SourceConnectionRequestParams) HasHiveConnectionParams() bool`

HasHiveConnectionParams returns a boolean if a field has been set.

### GetMssqlConnectionParams

`func (o *SourceConnectionRequestParams) GetMssqlConnectionParams() MssqlConnectionParams`

GetMssqlConnectionParams returns the MssqlConnectionParams field if non-nil, zero value otherwise.

### GetMssqlConnectionParamsOk

`func (o *SourceConnectionRequestParams) GetMssqlConnectionParamsOk() (*MssqlConnectionParams, bool)`

GetMssqlConnectionParamsOk returns a tuple with the MssqlConnectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlConnectionParams

`func (o *SourceConnectionRequestParams) SetMssqlConnectionParams(v MssqlConnectionParams)`

SetMssqlConnectionParams sets MssqlConnectionParams field to given value.

### HasMssqlConnectionParams

`func (o *SourceConnectionRequestParams) HasMssqlConnectionParams() bool`

HasMssqlConnectionParams returns a boolean if a field has been set.

### GetOracleConnectionParams

`func (o *SourceConnectionRequestParams) GetOracleConnectionParams() OracleConnectionParams`

GetOracleConnectionParams returns the OracleConnectionParams field if non-nil, zero value otherwise.

### GetOracleConnectionParamsOk

`func (o *SourceConnectionRequestParams) GetOracleConnectionParamsOk() (*OracleConnectionParams, bool)`

GetOracleConnectionParamsOk returns a tuple with the OracleConnectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleConnectionParams

`func (o *SourceConnectionRequestParams) SetOracleConnectionParams(v OracleConnectionParams)`

SetOracleConnectionParams sets OracleConnectionParams field to given value.

### HasOracleConnectionParams

`func (o *SourceConnectionRequestParams) HasOracleConnectionParams() bool`

HasOracleConnectionParams returns a boolean if a field has been set.

### GetVmwareConnectionParams

`func (o *SourceConnectionRequestParams) GetVmwareConnectionParams() VmwareConnectionParams`

GetVmwareConnectionParams returns the VmwareConnectionParams field if non-nil, zero value otherwise.

### GetVmwareConnectionParamsOk

`func (o *SourceConnectionRequestParams) GetVmwareConnectionParamsOk() (*VmwareConnectionParams, bool)`

GetVmwareConnectionParamsOk returns a tuple with the VmwareConnectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareConnectionParams

`func (o *SourceConnectionRequestParams) SetVmwareConnectionParams(v VmwareConnectionParams)`

SetVmwareConnectionParams sets VmwareConnectionParams field to given value.

### HasVmwareConnectionParams

`func (o *SourceConnectionRequestParams) HasVmwareConnectionParams() bool`

HasVmwareConnectionParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


