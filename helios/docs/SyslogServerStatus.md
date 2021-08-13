# SyslogServerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | The id of the syslog server. | [optional] 
**IsReachable** | Pointer to **NullableBool** | Specify if the syslog server is reachable or not. | [optional] 
**Message** | Pointer to **NullableString** | Description for current status. | [optional] 

## Methods

### NewSyslogServerStatus

`func NewSyslogServerStatus() *SyslogServerStatus`

NewSyslogServerStatus instantiates a new SyslogServerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogServerStatusWithDefaults

`func NewSyslogServerStatusWithDefaults() *SyslogServerStatus`

NewSyslogServerStatusWithDefaults instantiates a new SyslogServerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SyslogServerStatus) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyslogServerStatus) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyslogServerStatus) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SyslogServerStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SyslogServerStatus) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SyslogServerStatus) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsReachable

`func (o *SyslogServerStatus) GetIsReachable() bool`

GetIsReachable returns the IsReachable field if non-nil, zero value otherwise.

### GetIsReachableOk

`func (o *SyslogServerStatus) GetIsReachableOk() (*bool, bool)`

GetIsReachableOk returns a tuple with the IsReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReachable

`func (o *SyslogServerStatus) SetIsReachable(v bool)`

SetIsReachable sets IsReachable field to given value.

### HasIsReachable

`func (o *SyslogServerStatus) HasIsReachable() bool`

HasIsReachable returns a boolean if a field has been set.

### SetIsReachableNil

`func (o *SyslogServerStatus) SetIsReachableNil(b bool)`

 SetIsReachableNil sets the value for IsReachable to be an explicit nil

### UnsetIsReachable
`func (o *SyslogServerStatus) UnsetIsReachable()`

UnsetIsReachable ensures that no value is present for IsReachable, not even an explicit nil
### GetMessage

`func (o *SyslogServerStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SyslogServerStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SyslogServerStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SyslogServerStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *SyslogServerStatus) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *SyslogServerStatus) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


