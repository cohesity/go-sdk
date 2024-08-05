# Schedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeriodicTimeWindows** | Pointer to [**[]TimeWindow**](TimeWindow.md) |  Specifies the time range within the days of the week. | [optional] 
**ScheduleType** | Pointer to **NullableString** | Specifies the type of schedule for this ScheduleProto. | [optional] 
**TimeRanges** | Pointer to [**[]TimeRangeUsecs**](TimeRangeUsecs.md) |  Specifies the time ranges in usecs. | [optional] 
**Timezone** | Pointer to **string** | Specifies the timezone of the user of this ScheduleProto. The timezones have unique names of the form &#39;Area/Location&#39;. | [optional] 

## Methods

### NewSchedule

`func NewSchedule() *Schedule`

NewSchedule instantiates a new Schedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleWithDefaults

`func NewScheduleWithDefaults() *Schedule`

NewScheduleWithDefaults instantiates a new Schedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriodicTimeWindows

`func (o *Schedule) GetPeriodicTimeWindows() []TimeWindow`

GetPeriodicTimeWindows returns the PeriodicTimeWindows field if non-nil, zero value otherwise.

### GetPeriodicTimeWindowsOk

`func (o *Schedule) GetPeriodicTimeWindowsOk() (*[]TimeWindow, bool)`

GetPeriodicTimeWindowsOk returns a tuple with the PeriodicTimeWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicTimeWindows

`func (o *Schedule) SetPeriodicTimeWindows(v []TimeWindow)`

SetPeriodicTimeWindows sets PeriodicTimeWindows field to given value.

### HasPeriodicTimeWindows

`func (o *Schedule) HasPeriodicTimeWindows() bool`

HasPeriodicTimeWindows returns a boolean if a field has been set.

### GetScheduleType

`func (o *Schedule) GetScheduleType() string`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *Schedule) GetScheduleTypeOk() (*string, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *Schedule) SetScheduleType(v string)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *Schedule) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### SetScheduleTypeNil

`func (o *Schedule) SetScheduleTypeNil(b bool)`

 SetScheduleTypeNil sets the value for ScheduleType to be an explicit nil

### UnsetScheduleType
`func (o *Schedule) UnsetScheduleType()`

UnsetScheduleType ensures that no value is present for ScheduleType, not even an explicit nil
### GetTimeRanges

`func (o *Schedule) GetTimeRanges() []TimeRangeUsecs`

GetTimeRanges returns the TimeRanges field if non-nil, zero value otherwise.

### GetTimeRangesOk

`func (o *Schedule) GetTimeRangesOk() (*[]TimeRangeUsecs, bool)`

GetTimeRangesOk returns a tuple with the TimeRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRanges

`func (o *Schedule) SetTimeRanges(v []TimeRangeUsecs)`

SetTimeRanges sets TimeRanges field to given value.

### HasTimeRanges

`func (o *Schedule) HasTimeRanges() bool`

HasTimeRanges returns a boolean if a field has been set.

### GetTimezone

`func (o *Schedule) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Schedule) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Schedule) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *Schedule) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


