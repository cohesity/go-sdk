# HeliosStorageArraySnapshotSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaySchedule** | Pointer to [**DaySchedule**](DaySchedule.md) |  | [optional] 
**HourSchedule** | Pointer to [**HourSchedule**](HourSchedule.md) |  | [optional] 
**MinuteSchedule** | Pointer to [**MinuteSchedule**](MinuteSchedule.md) |  | [optional] 
**MonthSchedule** | Pointer to [**MonthSchedule**](MonthSchedule.md) |  | [optional] 
**Unit** | **NullableString** | Specifies how often to start new Protection Group Runs of a Protection Group. &lt;br&gt;&#39;Minutes&#39; specifies that Protection Group run starts periodically after certain number of minutes specified in &#39;frequency&#39; field. &lt;br&gt;&#39;Hours&#39; specifies that Protection Group run starts periodically after certain number of hours specified in &#39;frequency&#39; field. | 
**WeekSchedule** | Pointer to [**WeekSchedule**](WeekSchedule.md) |  | [optional] 
**YearSchedule** | Pointer to [**YearSchedule**](YearSchedule.md) |  | [optional] 

## Methods

### NewHeliosStorageArraySnapshotSchedule

`func NewHeliosStorageArraySnapshotSchedule(unit NullableString, ) *HeliosStorageArraySnapshotSchedule`

NewHeliosStorageArraySnapshotSchedule instantiates a new HeliosStorageArraySnapshotSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosStorageArraySnapshotScheduleWithDefaults

`func NewHeliosStorageArraySnapshotScheduleWithDefaults() *HeliosStorageArraySnapshotSchedule`

NewHeliosStorageArraySnapshotScheduleWithDefaults instantiates a new HeliosStorageArraySnapshotSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaySchedule

`func (o *HeliosStorageArraySnapshotSchedule) GetDaySchedule() DaySchedule`

GetDaySchedule returns the DaySchedule field if non-nil, zero value otherwise.

### GetDayScheduleOk

`func (o *HeliosStorageArraySnapshotSchedule) GetDayScheduleOk() (*DaySchedule, bool)`

GetDayScheduleOk returns a tuple with the DaySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaySchedule

`func (o *HeliosStorageArraySnapshotSchedule) SetDaySchedule(v DaySchedule)`

SetDaySchedule sets DaySchedule field to given value.

### HasDaySchedule

`func (o *HeliosStorageArraySnapshotSchedule) HasDaySchedule() bool`

HasDaySchedule returns a boolean if a field has been set.

### GetHourSchedule

`func (o *HeliosStorageArraySnapshotSchedule) GetHourSchedule() HourSchedule`

GetHourSchedule returns the HourSchedule field if non-nil, zero value otherwise.

### GetHourScheduleOk

`func (o *HeliosStorageArraySnapshotSchedule) GetHourScheduleOk() (*HourSchedule, bool)`

GetHourScheduleOk returns a tuple with the HourSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourSchedule

`func (o *HeliosStorageArraySnapshotSchedule) SetHourSchedule(v HourSchedule)`

SetHourSchedule sets HourSchedule field to given value.

### HasHourSchedule

`func (o *HeliosStorageArraySnapshotSchedule) HasHourSchedule() bool`

HasHourSchedule returns a boolean if a field has been set.

### GetMinuteSchedule

`func (o *HeliosStorageArraySnapshotSchedule) GetMinuteSchedule() MinuteSchedule`

GetMinuteSchedule returns the MinuteSchedule field if non-nil, zero value otherwise.

### GetMinuteScheduleOk

`func (o *HeliosStorageArraySnapshotSchedule) GetMinuteScheduleOk() (*MinuteSchedule, bool)`

GetMinuteScheduleOk returns a tuple with the MinuteSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinuteSchedule

`func (o *HeliosStorageArraySnapshotSchedule) SetMinuteSchedule(v MinuteSchedule)`

SetMinuteSchedule sets MinuteSchedule field to given value.

### HasMinuteSchedule

`func (o *HeliosStorageArraySnapshotSchedule) HasMinuteSchedule() bool`

HasMinuteSchedule returns a boolean if a field has been set.

### GetMonthSchedule

`func (o *HeliosStorageArraySnapshotSchedule) GetMonthSchedule() MonthSchedule`

GetMonthSchedule returns the MonthSchedule field if non-nil, zero value otherwise.

### GetMonthScheduleOk

`func (o *HeliosStorageArraySnapshotSchedule) GetMonthScheduleOk() (*MonthSchedule, bool)`

GetMonthScheduleOk returns a tuple with the MonthSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthSchedule

`func (o *HeliosStorageArraySnapshotSchedule) SetMonthSchedule(v MonthSchedule)`

SetMonthSchedule sets MonthSchedule field to given value.

### HasMonthSchedule

`func (o *HeliosStorageArraySnapshotSchedule) HasMonthSchedule() bool`

HasMonthSchedule returns a boolean if a field has been set.

### GetUnit

`func (o *HeliosStorageArraySnapshotSchedule) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *HeliosStorageArraySnapshotSchedule) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *HeliosStorageArraySnapshotSchedule) SetUnit(v string)`

SetUnit sets Unit field to given value.


### SetUnitNil

`func (o *HeliosStorageArraySnapshotSchedule) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *HeliosStorageArraySnapshotSchedule) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetWeekSchedule

`func (o *HeliosStorageArraySnapshotSchedule) GetWeekSchedule() WeekSchedule`

GetWeekSchedule returns the WeekSchedule field if non-nil, zero value otherwise.

### GetWeekScheduleOk

`func (o *HeliosStorageArraySnapshotSchedule) GetWeekScheduleOk() (*WeekSchedule, bool)`

GetWeekScheduleOk returns a tuple with the WeekSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekSchedule

`func (o *HeliosStorageArraySnapshotSchedule) SetWeekSchedule(v WeekSchedule)`

SetWeekSchedule sets WeekSchedule field to given value.

### HasWeekSchedule

`func (o *HeliosStorageArraySnapshotSchedule) HasWeekSchedule() bool`

HasWeekSchedule returns a boolean if a field has been set.

### GetYearSchedule

`func (o *HeliosStorageArraySnapshotSchedule) GetYearSchedule() YearSchedule`

GetYearSchedule returns the YearSchedule field if non-nil, zero value otherwise.

### GetYearScheduleOk

`func (o *HeliosStorageArraySnapshotSchedule) GetYearScheduleOk() (*YearSchedule, bool)`

GetYearScheduleOk returns a tuple with the YearSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearSchedule

`func (o *HeliosStorageArraySnapshotSchedule) SetYearSchedule(v YearSchedule)`

SetYearSchedule sets YearSchedule field to given value.

### HasYearSchedule

`func (o *HeliosStorageArraySnapshotSchedule) HasYearSchedule() bool`

HasYearSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


