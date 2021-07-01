# PostgresNodeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultPassword** | Pointer to **NullableString** | Specifies the default password to access the postgres database. | [optional] 
**DefaultUsername** | Pointer to **NullableString** | Specifies the default username to access the postgres database. | [optional] 
**NodeId** | Pointer to **NullableInt64** | Specifies the id of the node where postgres database is running. | [optional] 
**NodeIp** | Pointer to **NullableString** | Specifies the ip of the node where postgres database is running. | [optional] 
**Port** | Pointer to **NullableInt32** | Specifies the information where postgres database is running. | [optional] 

## Methods

### NewPostgresNodeInfo

`func NewPostgresNodeInfo() *PostgresNodeInfo`

NewPostgresNodeInfo instantiates a new PostgresNodeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresNodeInfoWithDefaults

`func NewPostgresNodeInfoWithDefaults() *PostgresNodeInfo`

NewPostgresNodeInfoWithDefaults instantiates a new PostgresNodeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultPassword

`func (o *PostgresNodeInfo) GetDefaultPassword() string`

GetDefaultPassword returns the DefaultPassword field if non-nil, zero value otherwise.

### GetDefaultPasswordOk

`func (o *PostgresNodeInfo) GetDefaultPasswordOk() (*string, bool)`

GetDefaultPasswordOk returns a tuple with the DefaultPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPassword

`func (o *PostgresNodeInfo) SetDefaultPassword(v string)`

SetDefaultPassword sets DefaultPassword field to given value.

### HasDefaultPassword

`func (o *PostgresNodeInfo) HasDefaultPassword() bool`

HasDefaultPassword returns a boolean if a field has been set.

### SetDefaultPasswordNil

`func (o *PostgresNodeInfo) SetDefaultPasswordNil(b bool)`

 SetDefaultPasswordNil sets the value for DefaultPassword to be an explicit nil

### UnsetDefaultPassword
`func (o *PostgresNodeInfo) UnsetDefaultPassword()`

UnsetDefaultPassword ensures that no value is present for DefaultPassword, not even an explicit nil
### GetDefaultUsername

`func (o *PostgresNodeInfo) GetDefaultUsername() string`

GetDefaultUsername returns the DefaultUsername field if non-nil, zero value otherwise.

### GetDefaultUsernameOk

`func (o *PostgresNodeInfo) GetDefaultUsernameOk() (*string, bool)`

GetDefaultUsernameOk returns a tuple with the DefaultUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUsername

`func (o *PostgresNodeInfo) SetDefaultUsername(v string)`

SetDefaultUsername sets DefaultUsername field to given value.

### HasDefaultUsername

`func (o *PostgresNodeInfo) HasDefaultUsername() bool`

HasDefaultUsername returns a boolean if a field has been set.

### SetDefaultUsernameNil

`func (o *PostgresNodeInfo) SetDefaultUsernameNil(b bool)`

 SetDefaultUsernameNil sets the value for DefaultUsername to be an explicit nil

### UnsetDefaultUsername
`func (o *PostgresNodeInfo) UnsetDefaultUsername()`

UnsetDefaultUsername ensures that no value is present for DefaultUsername, not even an explicit nil
### GetNodeId

`func (o *PostgresNodeInfo) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *PostgresNodeInfo) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *PostgresNodeInfo) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *PostgresNodeInfo) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### SetNodeIdNil

`func (o *PostgresNodeInfo) SetNodeIdNil(b bool)`

 SetNodeIdNil sets the value for NodeId to be an explicit nil

### UnsetNodeId
`func (o *PostgresNodeInfo) UnsetNodeId()`

UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
### GetNodeIp

`func (o *PostgresNodeInfo) GetNodeIp() string`

GetNodeIp returns the NodeIp field if non-nil, zero value otherwise.

### GetNodeIpOk

`func (o *PostgresNodeInfo) GetNodeIpOk() (*string, bool)`

GetNodeIpOk returns a tuple with the NodeIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIp

`func (o *PostgresNodeInfo) SetNodeIp(v string)`

SetNodeIp sets NodeIp field to given value.

### HasNodeIp

`func (o *PostgresNodeInfo) HasNodeIp() bool`

HasNodeIp returns a boolean if a field has been set.

### SetNodeIpNil

`func (o *PostgresNodeInfo) SetNodeIpNil(b bool)`

 SetNodeIpNil sets the value for NodeIp to be an explicit nil

### UnsetNodeIp
`func (o *PostgresNodeInfo) UnsetNodeIp()`

UnsetNodeIp ensures that no value is present for NodeIp, not even an explicit nil
### GetPort

`func (o *PostgresNodeInfo) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PostgresNodeInfo) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PostgresNodeInfo) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PostgresNodeInfo) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *PostgresNodeInfo) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *PostgresNodeInfo) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


