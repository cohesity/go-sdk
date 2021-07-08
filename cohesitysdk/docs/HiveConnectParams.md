# HiveConnectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HdfsEntityId** | Pointer to **NullableInt64** | Specifies the entity id of the HDFS source for this Hive | [optional] 
**KerberosPrincipal** | Pointer to **NullableString** | Specifies the kerberos principal. | [optional] 
**Metastore** | Pointer to **NullableString** | Specifies the Hive metastore host. | [optional] 
**ThriftPort** | Pointer to **NullableInt32** | Specifies the Hive metastore thrift Port | [optional] 

## Methods

### NewHiveConnectParams

`func NewHiveConnectParams() *HiveConnectParams`

NewHiveConnectParams instantiates a new HiveConnectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHiveConnectParamsWithDefaults

`func NewHiveConnectParamsWithDefaults() *HiveConnectParams`

NewHiveConnectParamsWithDefaults instantiates a new HiveConnectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHdfsEntityId

`func (o *HiveConnectParams) GetHdfsEntityId() int64`

GetHdfsEntityId returns the HdfsEntityId field if non-nil, zero value otherwise.

### GetHdfsEntityIdOk

`func (o *HiveConnectParams) GetHdfsEntityIdOk() (*int64, bool)`

GetHdfsEntityIdOk returns a tuple with the HdfsEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsEntityId

`func (o *HiveConnectParams) SetHdfsEntityId(v int64)`

SetHdfsEntityId sets HdfsEntityId field to given value.

### HasHdfsEntityId

`func (o *HiveConnectParams) HasHdfsEntityId() bool`

HasHdfsEntityId returns a boolean if a field has been set.

### SetHdfsEntityIdNil

`func (o *HiveConnectParams) SetHdfsEntityIdNil(b bool)`

 SetHdfsEntityIdNil sets the value for HdfsEntityId to be an explicit nil

### UnsetHdfsEntityId
`func (o *HiveConnectParams) UnsetHdfsEntityId()`

UnsetHdfsEntityId ensures that no value is present for HdfsEntityId, not even an explicit nil
### GetKerberosPrincipal

`func (o *HiveConnectParams) GetKerberosPrincipal() string`

GetKerberosPrincipal returns the KerberosPrincipal field if non-nil, zero value otherwise.

### GetKerberosPrincipalOk

`func (o *HiveConnectParams) GetKerberosPrincipalOk() (*string, bool)`

GetKerberosPrincipalOk returns a tuple with the KerberosPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosPrincipal

`func (o *HiveConnectParams) SetKerberosPrincipal(v string)`

SetKerberosPrincipal sets KerberosPrincipal field to given value.

### HasKerberosPrincipal

`func (o *HiveConnectParams) HasKerberosPrincipal() bool`

HasKerberosPrincipal returns a boolean if a field has been set.

### SetKerberosPrincipalNil

`func (o *HiveConnectParams) SetKerberosPrincipalNil(b bool)`

 SetKerberosPrincipalNil sets the value for KerberosPrincipal to be an explicit nil

### UnsetKerberosPrincipal
`func (o *HiveConnectParams) UnsetKerberosPrincipal()`

UnsetKerberosPrincipal ensures that no value is present for KerberosPrincipal, not even an explicit nil
### GetMetastore

`func (o *HiveConnectParams) GetMetastore() string`

GetMetastore returns the Metastore field if non-nil, zero value otherwise.

### GetMetastoreOk

`func (o *HiveConnectParams) GetMetastoreOk() (*string, bool)`

GetMetastoreOk returns a tuple with the Metastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetastore

`func (o *HiveConnectParams) SetMetastore(v string)`

SetMetastore sets Metastore field to given value.

### HasMetastore

`func (o *HiveConnectParams) HasMetastore() bool`

HasMetastore returns a boolean if a field has been set.

### SetMetastoreNil

`func (o *HiveConnectParams) SetMetastoreNil(b bool)`

 SetMetastoreNil sets the value for Metastore to be an explicit nil

### UnsetMetastore
`func (o *HiveConnectParams) UnsetMetastore()`

UnsetMetastore ensures that no value is present for Metastore, not even an explicit nil
### GetThriftPort

`func (o *HiveConnectParams) GetThriftPort() int32`

GetThriftPort returns the ThriftPort field if non-nil, zero value otherwise.

### GetThriftPortOk

`func (o *HiveConnectParams) GetThriftPortOk() (*int32, bool)`

GetThriftPortOk returns a tuple with the ThriftPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThriftPort

`func (o *HiveConnectParams) SetThriftPort(v int32)`

SetThriftPort sets ThriftPort field to given value.

### HasThriftPort

`func (o *HiveConnectParams) HasThriftPort() bool`

HasThriftPort returns a boolean if a field has been set.

### SetThriftPortNil

`func (o *HiveConnectParams) SetThriftPortNil(b bool)`

 SetThriftPortNil sets the value for ThriftPort to be an explicit nil

### UnsetThriftPort
`func (o *HiveConnectParams) UnsetThriftPort()`

UnsetThriftPort ensures that no value is present for ThriftPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


