# RigelClaimLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **NullableString** | Specifies the message of this event. | [optional] 
**TimeStamp** | Pointer to **NullableInt64** | Specifies the time stamp in microseconds of the event. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the severity of the event. | [optional] 

## Methods

### NewRigelClaimLog

`func NewRigelClaimLog() *RigelClaimLog`

NewRigelClaimLog instantiates a new RigelClaimLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRigelClaimLogWithDefaults

`func NewRigelClaimLogWithDefaults() *RigelClaimLog`

NewRigelClaimLogWithDefaults instantiates a new RigelClaimLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *RigelClaimLog) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RigelClaimLog) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RigelClaimLog) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RigelClaimLog) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *RigelClaimLog) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *RigelClaimLog) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetTimeStamp

`func (o *RigelClaimLog) GetTimeStamp() int64`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *RigelClaimLog) GetTimeStampOk() (*int64, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *RigelClaimLog) SetTimeStamp(v int64)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *RigelClaimLog) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.

### SetTimeStampNil

`func (o *RigelClaimLog) SetTimeStampNil(b bool)`

 SetTimeStampNil sets the value for TimeStamp to be an explicit nil

### UnsetTimeStamp
`func (o *RigelClaimLog) UnsetTimeStamp()`

UnsetTimeStamp ensures that no value is present for TimeStamp, not even an explicit nil
### GetType

`func (o *RigelClaimLog) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RigelClaimLog) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RigelClaimLog) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RigelClaimLog) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *RigelClaimLog) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *RigelClaimLog) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


