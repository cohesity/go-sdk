# SnmpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentPort** | Pointer to **NullableInt32** | AgentPort is the TCP port SNMP agent is using. | [optional] 
**Operation** | Pointer to **NullableString** | Operation is the operation of configuring SNMP services. | [optional] 
**ReadUser** | Pointer to [**SnmpUser**](SnmpUser.md) |  | [optional] 
**Server** | Pointer to **NullableString** | Server is the IP address of Network Management System. | [optional] 
**SystemInfo** | Pointer to [**SnmpSysInfo**](SnmpSysInfo.md) |  | [optional] 
**TrapPort** | Pointer to **NullableInt32** | TrapPort is the TCP port SNMP agent is using. | [optional] 
**TrapUser** | Pointer to [**SnmpUser**](SnmpUser.md) |  | [optional] 
**Version** | Pointer to **NullableString** | SnmpVersion is the SNMP version to talk with SNMP agent. It is SNMP V2 or SNMP V3. | [optional] 
**Vip** | Pointer to **NullableString** | Vip is the IP address SNMP agent and SNMP Trap Daemon will use. It should be one of the VIPs assigned to the cluster. | [optional] 
**WriteUser** | Pointer to [**SnmpUser**](SnmpUser.md) |  | [optional] 

## Methods

### NewSnmpConfig

`func NewSnmpConfig() *SnmpConfig`

NewSnmpConfig instantiates a new SnmpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpConfigWithDefaults

`func NewSnmpConfigWithDefaults() *SnmpConfig`

NewSnmpConfigWithDefaults instantiates a new SnmpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentPort

`func (o *SnmpConfig) GetAgentPort() int32`

GetAgentPort returns the AgentPort field if non-nil, zero value otherwise.

### GetAgentPortOk

`func (o *SnmpConfig) GetAgentPortOk() (*int32, bool)`

GetAgentPortOk returns a tuple with the AgentPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPort

`func (o *SnmpConfig) SetAgentPort(v int32)`

SetAgentPort sets AgentPort field to given value.

### HasAgentPort

`func (o *SnmpConfig) HasAgentPort() bool`

HasAgentPort returns a boolean if a field has been set.

### SetAgentPortNil

`func (o *SnmpConfig) SetAgentPortNil(b bool)`

 SetAgentPortNil sets the value for AgentPort to be an explicit nil

### UnsetAgentPort
`func (o *SnmpConfig) UnsetAgentPort()`

UnsetAgentPort ensures that no value is present for AgentPort, not even an explicit nil
### GetOperation

`func (o *SnmpConfig) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *SnmpConfig) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *SnmpConfig) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *SnmpConfig) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### SetOperationNil

`func (o *SnmpConfig) SetOperationNil(b bool)`

 SetOperationNil sets the value for Operation to be an explicit nil

### UnsetOperation
`func (o *SnmpConfig) UnsetOperation()`

UnsetOperation ensures that no value is present for Operation, not even an explicit nil
### GetReadUser

`func (o *SnmpConfig) GetReadUser() SnmpUser`

GetReadUser returns the ReadUser field if non-nil, zero value otherwise.

### GetReadUserOk

`func (o *SnmpConfig) GetReadUserOk() (*SnmpUser, bool)`

GetReadUserOk returns a tuple with the ReadUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadUser

`func (o *SnmpConfig) SetReadUser(v SnmpUser)`

SetReadUser sets ReadUser field to given value.

### HasReadUser

`func (o *SnmpConfig) HasReadUser() bool`

HasReadUser returns a boolean if a field has been set.

### GetServer

`func (o *SnmpConfig) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *SnmpConfig) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *SnmpConfig) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *SnmpConfig) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *SnmpConfig) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *SnmpConfig) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil
### GetSystemInfo

`func (o *SnmpConfig) GetSystemInfo() SnmpSysInfo`

GetSystemInfo returns the SystemInfo field if non-nil, zero value otherwise.

### GetSystemInfoOk

`func (o *SnmpConfig) GetSystemInfoOk() (*SnmpSysInfo, bool)`

GetSystemInfoOk returns a tuple with the SystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemInfo

`func (o *SnmpConfig) SetSystemInfo(v SnmpSysInfo)`

SetSystemInfo sets SystemInfo field to given value.

### HasSystemInfo

`func (o *SnmpConfig) HasSystemInfo() bool`

HasSystemInfo returns a boolean if a field has been set.

### GetTrapPort

`func (o *SnmpConfig) GetTrapPort() int32`

GetTrapPort returns the TrapPort field if non-nil, zero value otherwise.

### GetTrapPortOk

`func (o *SnmpConfig) GetTrapPortOk() (*int32, bool)`

GetTrapPortOk returns a tuple with the TrapPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrapPort

`func (o *SnmpConfig) SetTrapPort(v int32)`

SetTrapPort sets TrapPort field to given value.

### HasTrapPort

`func (o *SnmpConfig) HasTrapPort() bool`

HasTrapPort returns a boolean if a field has been set.

### SetTrapPortNil

`func (o *SnmpConfig) SetTrapPortNil(b bool)`

 SetTrapPortNil sets the value for TrapPort to be an explicit nil

### UnsetTrapPort
`func (o *SnmpConfig) UnsetTrapPort()`

UnsetTrapPort ensures that no value is present for TrapPort, not even an explicit nil
### GetTrapUser

`func (o *SnmpConfig) GetTrapUser() SnmpUser`

GetTrapUser returns the TrapUser field if non-nil, zero value otherwise.

### GetTrapUserOk

`func (o *SnmpConfig) GetTrapUserOk() (*SnmpUser, bool)`

GetTrapUserOk returns a tuple with the TrapUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrapUser

`func (o *SnmpConfig) SetTrapUser(v SnmpUser)`

SetTrapUser sets TrapUser field to given value.

### HasTrapUser

`func (o *SnmpConfig) HasTrapUser() bool`

HasTrapUser returns a boolean if a field has been set.

### GetVersion

`func (o *SnmpConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SnmpConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SnmpConfig) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SnmpConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *SnmpConfig) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *SnmpConfig) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetVip

`func (o *SnmpConfig) GetVip() string`

GetVip returns the Vip field if non-nil, zero value otherwise.

### GetVipOk

`func (o *SnmpConfig) GetVipOk() (*string, bool)`

GetVipOk returns a tuple with the Vip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVip

`func (o *SnmpConfig) SetVip(v string)`

SetVip sets Vip field to given value.

### HasVip

`func (o *SnmpConfig) HasVip() bool`

HasVip returns a boolean if a field has been set.

### SetVipNil

`func (o *SnmpConfig) SetVipNil(b bool)`

 SetVipNil sets the value for Vip to be an explicit nil

### UnsetVip
`func (o *SnmpConfig) UnsetVip()`

UnsetVip ensures that no value is present for Vip, not even an explicit nil
### GetWriteUser

`func (o *SnmpConfig) GetWriteUser() SnmpUser`

GetWriteUser returns the WriteUser field if non-nil, zero value otherwise.

### GetWriteUserOk

`func (o *SnmpConfig) GetWriteUserOk() (*SnmpUser, bool)`

GetWriteUserOk returns a tuple with the WriteUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteUser

`func (o *SnmpConfig) SetWriteUser(v SnmpUser)`

SetWriteUser sets WriteUser field to given value.

### HasWriteUser

`func (o *SnmpConfig) HasWriteUser() bool`

HasWriteUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


