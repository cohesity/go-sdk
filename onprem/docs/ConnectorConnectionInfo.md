# ConnectorConnectionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsActive** | Pointer to **NullableBool** | Specifies whether the connector is currently connected to cohesity server. | [optional] 
**Message** | Pointer to **NullableString** | Specifies possible error message when the connector is not able to connect. | [optional] 
**LastConnectedTimestampMsecs** | Pointer to **NullableInt64** | Specifies last timestamp for which connection status was known. | [optional] 

## Methods

### NewConnectorConnectionInfo

`func NewConnectorConnectionInfo() *ConnectorConnectionInfo`

NewConnectorConnectionInfo instantiates a new ConnectorConnectionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorConnectionInfoWithDefaults

`func NewConnectorConnectionInfoWithDefaults() *ConnectorConnectionInfo`

NewConnectorConnectionInfoWithDefaults instantiates a new ConnectorConnectionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsActive

`func (o *ConnectorConnectionInfo) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ConnectorConnectionInfo) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ConnectorConnectionInfo) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ConnectorConnectionInfo) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *ConnectorConnectionInfo) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *ConnectorConnectionInfo) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetMessage

`func (o *ConnectorConnectionInfo) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConnectorConnectionInfo) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConnectorConnectionInfo) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ConnectorConnectionInfo) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ConnectorConnectionInfo) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ConnectorConnectionInfo) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetLastConnectedTimestampMsecs

`func (o *ConnectorConnectionInfo) GetLastConnectedTimestampMsecs() int64`

GetLastConnectedTimestampMsecs returns the LastConnectedTimestampMsecs field if non-nil, zero value otherwise.

### GetLastConnectedTimestampMsecsOk

`func (o *ConnectorConnectionInfo) GetLastConnectedTimestampMsecsOk() (*int64, bool)`

GetLastConnectedTimestampMsecsOk returns a tuple with the LastConnectedTimestampMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectedTimestampMsecs

`func (o *ConnectorConnectionInfo) SetLastConnectedTimestampMsecs(v int64)`

SetLastConnectedTimestampMsecs sets LastConnectedTimestampMsecs field to given value.

### HasLastConnectedTimestampMsecs

`func (o *ConnectorConnectionInfo) HasLastConnectedTimestampMsecs() bool`

HasLastConnectedTimestampMsecs returns a boolean if a field has been set.

### SetLastConnectedTimestampMsecsNil

`func (o *ConnectorConnectionInfo) SetLastConnectedTimestampMsecsNil(b bool)`

 SetLastConnectedTimestampMsecsNil sets the value for LastConnectedTimestampMsecs to be an explicit nil

### UnsetLastConnectedTimestampMsecs
`func (o *ConnectorConnectionInfo) UnsetLastConnectedTimestampMsecs()`

UnsetLastConnectedTimestampMsecs ensures that no value is present for LastConnectedTimestampMsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


