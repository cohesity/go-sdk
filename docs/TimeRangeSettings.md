# TimeRangeSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time specified as a Unix epoch Timestamp (in microseconds). | [optional] 
**JobUid** | Pointer to [**UniversalId**](UniversalId.md) |  | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time specified as a Unix epoch Timestamp (in microseconds). | [optional] 

## Methods

### NewTimeRangeSettings

`func NewTimeRangeSettings() *TimeRangeSettings`

NewTimeRangeSettings instantiates a new TimeRangeSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeRangeSettingsWithDefaults

`func NewTimeRangeSettingsWithDefaults() *TimeRangeSettings`

NewTimeRangeSettingsWithDefaults instantiates a new TimeRangeSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTimeUsecs

`func (o *TimeRangeSettings) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *TimeRangeSettings) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *TimeRangeSettings) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *TimeRangeSettings) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *TimeRangeSettings) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *TimeRangeSettings) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetJobUid

`func (o *TimeRangeSettings) GetJobUid() UniversalId`

GetJobUid returns the JobUid field if non-nil, zero value otherwise.

### GetJobUidOk

`func (o *TimeRangeSettings) GetJobUidOk() (*UniversalId, bool)`

GetJobUidOk returns a tuple with the JobUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobUid

`func (o *TimeRangeSettings) SetJobUid(v UniversalId)`

SetJobUid sets JobUid field to given value.

### HasJobUid

`func (o *TimeRangeSettings) HasJobUid() bool`

HasJobUid returns a boolean if a field has been set.

### GetStartTimeUsecs

`func (o *TimeRangeSettings) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *TimeRangeSettings) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *TimeRangeSettings) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *TimeRangeSettings) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *TimeRangeSettings) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *TimeRangeSettings) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


