# SourceConnectionResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | **NullableString** | Specifies the environment type of the Protection Source. | 
**ConnectionId** | Pointer to **NullableInt64** | Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user. | [optional] 
**CassandraConnectionResponseParams** | Pointer to [**CassandraSourceConfigParams**](CassandraSourceConfigParams.md) |  | [optional] 
**HiveConnectionResponseParams** | Pointer to [**HiveAdditionalParams**](HiveAdditionalParams.md) |  | [optional] 
**HbaseConnectionResponseParams** | Pointer to [**HBaseAdditionalParams**](HBaseAdditionalParams.md) |  | [optional] 
**HdfsConnectionResponseParams** | Pointer to [**HdfsAdditionalParams**](HdfsAdditionalParams.md) |  | [optional] 
**MssqlConnectionResponseParams** | Pointer to [**MssqlConnectionResponseParams**](MssqlConnectionResponseParams.md) |  | [optional] 
**VmwareConnectionResponseParams** | Pointer to [**VmwareAdditionalParams**](VmwareAdditionalParams.md) |  | [optional] 

## Methods

### NewSourceConnectionResponseParams

`func NewSourceConnectionResponseParams(environment NullableString, ) *SourceConnectionResponseParams`

NewSourceConnectionResponseParams instantiates a new SourceConnectionResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceConnectionResponseParamsWithDefaults

`func NewSourceConnectionResponseParamsWithDefaults() *SourceConnectionResponseParams`

NewSourceConnectionResponseParamsWithDefaults instantiates a new SourceConnectionResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *SourceConnectionResponseParams) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SourceConnectionResponseParams) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SourceConnectionResponseParams) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### SetEnvironmentNil

`func (o *SourceConnectionResponseParams) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *SourceConnectionResponseParams) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetConnectionId

`func (o *SourceConnectionResponseParams) GetConnectionId() int64`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *SourceConnectionResponseParams) GetConnectionIdOk() (*int64, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *SourceConnectionResponseParams) SetConnectionId(v int64)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *SourceConnectionResponseParams) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### SetConnectionIdNil

`func (o *SourceConnectionResponseParams) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *SourceConnectionResponseParams) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetCassandraConnectionResponseParams

`func (o *SourceConnectionResponseParams) GetCassandraConnectionResponseParams() CassandraSourceConfigParams`

GetCassandraConnectionResponseParams returns the CassandraConnectionResponseParams field if non-nil, zero value otherwise.

### GetCassandraConnectionResponseParamsOk

`func (o *SourceConnectionResponseParams) GetCassandraConnectionResponseParamsOk() (*CassandraSourceConfigParams, bool)`

GetCassandraConnectionResponseParamsOk returns a tuple with the CassandraConnectionResponseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraConnectionResponseParams

`func (o *SourceConnectionResponseParams) SetCassandraConnectionResponseParams(v CassandraSourceConfigParams)`

SetCassandraConnectionResponseParams sets CassandraConnectionResponseParams field to given value.

### HasCassandraConnectionResponseParams

`func (o *SourceConnectionResponseParams) HasCassandraConnectionResponseParams() bool`

HasCassandraConnectionResponseParams returns a boolean if a field has been set.

### GetHiveConnectionResponseParams

`func (o *SourceConnectionResponseParams) GetHiveConnectionResponseParams() HiveAdditionalParams`

GetHiveConnectionResponseParams returns the HiveConnectionResponseParams field if non-nil, zero value otherwise.

### GetHiveConnectionResponseParamsOk

`func (o *SourceConnectionResponseParams) GetHiveConnectionResponseParamsOk() (*HiveAdditionalParams, bool)`

GetHiveConnectionResponseParamsOk returns a tuple with the HiveConnectionResponseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveConnectionResponseParams

`func (o *SourceConnectionResponseParams) SetHiveConnectionResponseParams(v HiveAdditionalParams)`

SetHiveConnectionResponseParams sets HiveConnectionResponseParams field to given value.

### HasHiveConnectionResponseParams

`func (o *SourceConnectionResponseParams) HasHiveConnectionResponseParams() bool`

HasHiveConnectionResponseParams returns a boolean if a field has been set.

### GetHbaseConnectionResponseParams

`func (o *SourceConnectionResponseParams) GetHbaseConnectionResponseParams() HBaseAdditionalParams`

GetHbaseConnectionResponseParams returns the HbaseConnectionResponseParams field if non-nil, zero value otherwise.

### GetHbaseConnectionResponseParamsOk

`func (o *SourceConnectionResponseParams) GetHbaseConnectionResponseParamsOk() (*HBaseAdditionalParams, bool)`

GetHbaseConnectionResponseParamsOk returns a tuple with the HbaseConnectionResponseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseConnectionResponseParams

`func (o *SourceConnectionResponseParams) SetHbaseConnectionResponseParams(v HBaseAdditionalParams)`

SetHbaseConnectionResponseParams sets HbaseConnectionResponseParams field to given value.

### HasHbaseConnectionResponseParams

`func (o *SourceConnectionResponseParams) HasHbaseConnectionResponseParams() bool`

HasHbaseConnectionResponseParams returns a boolean if a field has been set.

### GetHdfsConnectionResponseParams

`func (o *SourceConnectionResponseParams) GetHdfsConnectionResponseParams() HdfsAdditionalParams`

GetHdfsConnectionResponseParams returns the HdfsConnectionResponseParams field if non-nil, zero value otherwise.

### GetHdfsConnectionResponseParamsOk

`func (o *SourceConnectionResponseParams) GetHdfsConnectionResponseParamsOk() (*HdfsAdditionalParams, bool)`

GetHdfsConnectionResponseParamsOk returns a tuple with the HdfsConnectionResponseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsConnectionResponseParams

`func (o *SourceConnectionResponseParams) SetHdfsConnectionResponseParams(v HdfsAdditionalParams)`

SetHdfsConnectionResponseParams sets HdfsConnectionResponseParams field to given value.

### HasHdfsConnectionResponseParams

`func (o *SourceConnectionResponseParams) HasHdfsConnectionResponseParams() bool`

HasHdfsConnectionResponseParams returns a boolean if a field has been set.

### GetMssqlConnectionResponseParams

`func (o *SourceConnectionResponseParams) GetMssqlConnectionResponseParams() MssqlConnectionResponseParams`

GetMssqlConnectionResponseParams returns the MssqlConnectionResponseParams field if non-nil, zero value otherwise.

### GetMssqlConnectionResponseParamsOk

`func (o *SourceConnectionResponseParams) GetMssqlConnectionResponseParamsOk() (*MssqlConnectionResponseParams, bool)`

GetMssqlConnectionResponseParamsOk returns a tuple with the MssqlConnectionResponseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssqlConnectionResponseParams

`func (o *SourceConnectionResponseParams) SetMssqlConnectionResponseParams(v MssqlConnectionResponseParams)`

SetMssqlConnectionResponseParams sets MssqlConnectionResponseParams field to given value.

### HasMssqlConnectionResponseParams

`func (o *SourceConnectionResponseParams) HasMssqlConnectionResponseParams() bool`

HasMssqlConnectionResponseParams returns a boolean if a field has been set.

### GetVmwareConnectionResponseParams

`func (o *SourceConnectionResponseParams) GetVmwareConnectionResponseParams() VmwareAdditionalParams`

GetVmwareConnectionResponseParams returns the VmwareConnectionResponseParams field if non-nil, zero value otherwise.

### GetVmwareConnectionResponseParamsOk

`func (o *SourceConnectionResponseParams) GetVmwareConnectionResponseParamsOk() (*VmwareAdditionalParams, bool)`

GetVmwareConnectionResponseParamsOk returns a tuple with the VmwareConnectionResponseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareConnectionResponseParams

`func (o *SourceConnectionResponseParams) SetVmwareConnectionResponseParams(v VmwareAdditionalParams)`

SetVmwareConnectionResponseParams sets VmwareConnectionResponseParams field to given value.

### HasVmwareConnectionResponseParams

`func (o *SourceConnectionResponseParams) HasVmwareConnectionResponseParams() bool`

HasVmwareConnectionResponseParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


