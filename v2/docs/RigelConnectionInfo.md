# RigelConnectionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsActive** | Pointer to **NullableBool** | Specifies if the connection is active. | [optional] 
**LastConnectedTimestampMsecs** | Pointer to **NullableInt64** | Specifies last timestamp for which connection status was known. | [optional] 
**Message** | Pointer to **NullableString** | Specifies possible error message when rigel is not able to connect. | [optional] 

## Methods

### NewRigelConnectionInfo

`func NewRigelConnectionInfo() *RigelConnectionInfo`

NewRigelConnectionInfo instantiates a new RigelConnectionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRigelConnectionInfoWithDefaults

`func NewRigelConnectionInfoWithDefaults() *RigelConnectionInfo`

NewRigelConnectionInfoWithDefaults instantiates a new RigelConnectionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsActive

`func (o *RigelConnectionInfo) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *RigelConnectionInfo) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *RigelConnectionInfo) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *RigelConnectionInfo) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *RigelConnectionInfo) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *RigelConnectionInfo) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetLastConnectedTimestampMsecs

`func (o *RigelConnectionInfo) GetLastConnectedTimestampMsecs() int64`

GetLastConnectedTimestampMsecs returns the LastConnectedTimestampMsecs field if non-nil, zero value otherwise.

### GetLastConnectedTimestampMsecsOk

`func (o *RigelConnectionInfo) GetLastConnectedTimestampMsecsOk() (*int64, bool)`

GetLastConnectedTimestampMsecsOk returns a tuple with the LastConnectedTimestampMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectedTimestampMsecs

`func (o *RigelConnectionInfo) SetLastConnectedTimestampMsecs(v int64)`

SetLastConnectedTimestampMsecs sets LastConnectedTimestampMsecs field to given value.

### HasLastConnectedTimestampMsecs

`func (o *RigelConnectionInfo) HasLastConnectedTimestampMsecs() bool`

HasLastConnectedTimestampMsecs returns a boolean if a field has been set.

### SetLastConnectedTimestampMsecsNil

`func (o *RigelConnectionInfo) SetLastConnectedTimestampMsecsNil(b bool)`

 SetLastConnectedTimestampMsecsNil sets the value for LastConnectedTimestampMsecs to be an explicit nil

### UnsetLastConnectedTimestampMsecs
`func (o *RigelConnectionInfo) UnsetLastConnectedTimestampMsecs()`

UnsetLastConnectedTimestampMsecs ensures that no value is present for LastConnectedTimestampMsecs, not even an explicit nil
### GetMessage

`func (o *RigelConnectionInfo) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RigelConnectionInfo) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RigelConnectionInfo) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RigelConnectionInfo) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *RigelConnectionInfo) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *RigelConnectionInfo) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


