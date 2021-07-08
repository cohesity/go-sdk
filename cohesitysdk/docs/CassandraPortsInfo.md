# CassandraPortsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JmxPort** | Pointer to **NullableInt32** | Specifies the Cassandra JMX port. | [optional] 
**NativeTransportPort** | Pointer to **NullableInt32** | Specifies the port for the CQL native transport. | [optional] 
**RpcPort** | Pointer to **NullableInt32** | Specifies the Remote Procedure Call (RPC) port for general mechanism for client-server applications. | [optional] 
**SslStoragePort** | Pointer to **NullableInt32** | Specifies the SSL port for encrypted communication. | [optional] 
**StoragePort** | Pointer to **NullableInt32** | Specifies the TCP port for data. Internally used by Cassandra bulk loader. | [optional] 

## Methods

### NewCassandraPortsInfo

`func NewCassandraPortsInfo() *CassandraPortsInfo`

NewCassandraPortsInfo instantiates a new CassandraPortsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCassandraPortsInfoWithDefaults

`func NewCassandraPortsInfoWithDefaults() *CassandraPortsInfo`

NewCassandraPortsInfoWithDefaults instantiates a new CassandraPortsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJmxPort

`func (o *CassandraPortsInfo) GetJmxPort() int32`

GetJmxPort returns the JmxPort field if non-nil, zero value otherwise.

### GetJmxPortOk

`func (o *CassandraPortsInfo) GetJmxPortOk() (*int32, bool)`

GetJmxPortOk returns a tuple with the JmxPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJmxPort

`func (o *CassandraPortsInfo) SetJmxPort(v int32)`

SetJmxPort sets JmxPort field to given value.

### HasJmxPort

`func (o *CassandraPortsInfo) HasJmxPort() bool`

HasJmxPort returns a boolean if a field has been set.

### SetJmxPortNil

`func (o *CassandraPortsInfo) SetJmxPortNil(b bool)`

 SetJmxPortNil sets the value for JmxPort to be an explicit nil

### UnsetJmxPort
`func (o *CassandraPortsInfo) UnsetJmxPort()`

UnsetJmxPort ensures that no value is present for JmxPort, not even an explicit nil
### GetNativeTransportPort

`func (o *CassandraPortsInfo) GetNativeTransportPort() int32`

GetNativeTransportPort returns the NativeTransportPort field if non-nil, zero value otherwise.

### GetNativeTransportPortOk

`func (o *CassandraPortsInfo) GetNativeTransportPortOk() (*int32, bool)`

GetNativeTransportPortOk returns a tuple with the NativeTransportPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeTransportPort

`func (o *CassandraPortsInfo) SetNativeTransportPort(v int32)`

SetNativeTransportPort sets NativeTransportPort field to given value.

### HasNativeTransportPort

`func (o *CassandraPortsInfo) HasNativeTransportPort() bool`

HasNativeTransportPort returns a boolean if a field has been set.

### SetNativeTransportPortNil

`func (o *CassandraPortsInfo) SetNativeTransportPortNil(b bool)`

 SetNativeTransportPortNil sets the value for NativeTransportPort to be an explicit nil

### UnsetNativeTransportPort
`func (o *CassandraPortsInfo) UnsetNativeTransportPort()`

UnsetNativeTransportPort ensures that no value is present for NativeTransportPort, not even an explicit nil
### GetRpcPort

`func (o *CassandraPortsInfo) GetRpcPort() int32`

GetRpcPort returns the RpcPort field if non-nil, zero value otherwise.

### GetRpcPortOk

`func (o *CassandraPortsInfo) GetRpcPortOk() (*int32, bool)`

GetRpcPortOk returns a tuple with the RpcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcPort

`func (o *CassandraPortsInfo) SetRpcPort(v int32)`

SetRpcPort sets RpcPort field to given value.

### HasRpcPort

`func (o *CassandraPortsInfo) HasRpcPort() bool`

HasRpcPort returns a boolean if a field has been set.

### SetRpcPortNil

`func (o *CassandraPortsInfo) SetRpcPortNil(b bool)`

 SetRpcPortNil sets the value for RpcPort to be an explicit nil

### UnsetRpcPort
`func (o *CassandraPortsInfo) UnsetRpcPort()`

UnsetRpcPort ensures that no value is present for RpcPort, not even an explicit nil
### GetSslStoragePort

`func (o *CassandraPortsInfo) GetSslStoragePort() int32`

GetSslStoragePort returns the SslStoragePort field if non-nil, zero value otherwise.

### GetSslStoragePortOk

`func (o *CassandraPortsInfo) GetSslStoragePortOk() (*int32, bool)`

GetSslStoragePortOk returns a tuple with the SslStoragePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslStoragePort

`func (o *CassandraPortsInfo) SetSslStoragePort(v int32)`

SetSslStoragePort sets SslStoragePort field to given value.

### HasSslStoragePort

`func (o *CassandraPortsInfo) HasSslStoragePort() bool`

HasSslStoragePort returns a boolean if a field has been set.

### SetSslStoragePortNil

`func (o *CassandraPortsInfo) SetSslStoragePortNil(b bool)`

 SetSslStoragePortNil sets the value for SslStoragePort to be an explicit nil

### UnsetSslStoragePort
`func (o *CassandraPortsInfo) UnsetSslStoragePort()`

UnsetSslStoragePort ensures that no value is present for SslStoragePort, not even an explicit nil
### GetStoragePort

`func (o *CassandraPortsInfo) GetStoragePort() int32`

GetStoragePort returns the StoragePort field if non-nil, zero value otherwise.

### GetStoragePortOk

`func (o *CassandraPortsInfo) GetStoragePortOk() (*int32, bool)`

GetStoragePortOk returns a tuple with the StoragePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePort

`func (o *CassandraPortsInfo) SetStoragePort(v int32)`

SetStoragePort sets StoragePort field to given value.

### HasStoragePort

`func (o *CassandraPortsInfo) HasStoragePort() bool`

HasStoragePort returns a boolean if a field has been set.

### SetStoragePortNil

`func (o *CassandraPortsInfo) SetStoragePortNil(b bool)`

 SetStoragePortNil sets the value for StoragePort to be an explicit nil

### UnsetStoragePort
`func (o *CassandraPortsInfo) UnsetStoragePort()`

UnsetStoragePort ensures that no value is present for StoragePort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


