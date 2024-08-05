# ProgressTaskEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **NullableString** | Specifies the log message describing the current event. | [optional] 
**OccuredAtUsecs** | Pointer to **NullableInt64** | Specifies the time of the event occurance in Unix epoch Timestamp(in microseconds). | [optional] 

## Methods

### NewProgressTaskEvent

`func NewProgressTaskEvent() *ProgressTaskEvent`

NewProgressTaskEvent instantiates a new ProgressTaskEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgressTaskEventWithDefaults

`func NewProgressTaskEventWithDefaults() *ProgressTaskEvent`

NewProgressTaskEventWithDefaults instantiates a new ProgressTaskEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ProgressTaskEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ProgressTaskEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ProgressTaskEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ProgressTaskEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ProgressTaskEvent) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ProgressTaskEvent) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetOccuredAtUsecs

`func (o *ProgressTaskEvent) GetOccuredAtUsecs() int64`

GetOccuredAtUsecs returns the OccuredAtUsecs field if non-nil, zero value otherwise.

### GetOccuredAtUsecsOk

`func (o *ProgressTaskEvent) GetOccuredAtUsecsOk() (*int64, bool)`

GetOccuredAtUsecsOk returns a tuple with the OccuredAtUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccuredAtUsecs

`func (o *ProgressTaskEvent) SetOccuredAtUsecs(v int64)`

SetOccuredAtUsecs sets OccuredAtUsecs field to given value.

### HasOccuredAtUsecs

`func (o *ProgressTaskEvent) HasOccuredAtUsecs() bool`

HasOccuredAtUsecs returns a boolean if a field has been set.

### SetOccuredAtUsecsNil

`func (o *ProgressTaskEvent) SetOccuredAtUsecsNil(b bool)`

 SetOccuredAtUsecsNil sets the value for OccuredAtUsecs to be an explicit nil

### UnsetOccuredAtUsecs
`func (o *ProgressTaskEvent) UnsetOccuredAtUsecs()`

UnsetOccuredAtUsecs ensures that no value is present for OccuredAtUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


