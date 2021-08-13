# SyslogServers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyslogServers** | Pointer to [**[]SyslogServer**](SyslogServer.md) | Specifies the list of syslog servers. | [optional] 

## Methods

### NewSyslogServers

`func NewSyslogServers() *SyslogServers`

NewSyslogServers instantiates a new SyslogServers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogServersWithDefaults

`func NewSyslogServersWithDefaults() *SyslogServers`

NewSyslogServersWithDefaults instantiates a new SyslogServers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSyslogServers

`func (o *SyslogServers) GetSyslogServers() []SyslogServer`

GetSyslogServers returns the SyslogServers field if non-nil, zero value otherwise.

### GetSyslogServersOk

`func (o *SyslogServers) GetSyslogServersOk() (*[]SyslogServer, bool)`

GetSyslogServersOk returns a tuple with the SyslogServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogServers

`func (o *SyslogServers) SetSyslogServers(v []SyslogServer)`

SetSyslogServers sets SyslogServers field to given value.

### HasSyslogServers

`func (o *SyslogServers) HasSyslogServers() bool`

HasSyslogServers returns a boolean if a field has been set.

### SetSyslogServersNil

`func (o *SyslogServers) SetSyslogServersNil(b bool)`

 SetSyslogServersNil sets the value for SyslogServers to be an explicit nil

### UnsetSyslogServers
`func (o *SyslogServers) UnsetSyslogServers()`

UnsetSyslogServers ensures that no value is present for SyslogServers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


