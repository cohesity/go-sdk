# OperationEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **NullableString** | Specifies the message describing the event. | [optional] 
**Severity** | Pointer to **NullableString** | Specifies the severity of an event. | [optional] 
**Timestamp** | Pointer to **int64** | Specifies the unix epoch timestamp in microseconds when this event took place. | [optional] 

## Methods

### NewOperationEvents

`func NewOperationEvents() *OperationEvents`

NewOperationEvents instantiates a new OperationEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationEventsWithDefaults

`func NewOperationEventsWithDefaults() *OperationEvents`

NewOperationEventsWithDefaults instantiates a new OperationEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *OperationEvents) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OperationEvents) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OperationEvents) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OperationEvents) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *OperationEvents) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *OperationEvents) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetSeverity

`func (o *OperationEvents) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *OperationEvents) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *OperationEvents) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *OperationEvents) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *OperationEvents) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *OperationEvents) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
### GetTimestamp

`func (o *OperationEvents) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *OperationEvents) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *OperationEvents) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *OperationEvents) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


