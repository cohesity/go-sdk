# LogSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | **NullableString** | Specifies how often to start new Protection Group Runs of a Protection Group. &lt;br&gt;&#39;Minutes&#39; specifies that Protection Group run starts periodically after certain number of minutes specified in &#39;frequency&#39; field. &lt;br&gt;&#39;Hours&#39; specifies that Protection Group run starts periodically after certain number of hours specified in &#39;frequency&#39; field. | 
**MinuteSchedule** | Pointer to [**MinuteSchedule**](MinuteSchedule.md) |  | [optional] 
**HourSchedule** | Pointer to [**HourSchedule**](HourSchedule.md) |  | [optional] 

## Methods

### NewLogSchedule

`func NewLogSchedule(unit NullableString, ) *LogSchedule`

NewLogSchedule instantiates a new LogSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogScheduleWithDefaults

`func NewLogScheduleWithDefaults() *LogSchedule`

NewLogScheduleWithDefaults instantiates a new LogSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *LogSchedule) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *LogSchedule) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *LogSchedule) SetUnit(v string)`

SetUnit sets Unit field to given value.


### SetUnitNil

`func (o *LogSchedule) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *LogSchedule) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetMinuteSchedule

`func (o *LogSchedule) GetMinuteSchedule() MinuteSchedule`

GetMinuteSchedule returns the MinuteSchedule field if non-nil, zero value otherwise.

### GetMinuteScheduleOk

`func (o *LogSchedule) GetMinuteScheduleOk() (*MinuteSchedule, bool)`

GetMinuteScheduleOk returns a tuple with the MinuteSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinuteSchedule

`func (o *LogSchedule) SetMinuteSchedule(v MinuteSchedule)`

SetMinuteSchedule sets MinuteSchedule field to given value.

### HasMinuteSchedule

`func (o *LogSchedule) HasMinuteSchedule() bool`

HasMinuteSchedule returns a boolean if a field has been set.

### GetHourSchedule

`func (o *LogSchedule) GetHourSchedule() HourSchedule`

GetHourSchedule returns the HourSchedule field if non-nil, zero value otherwise.

### GetHourScheduleOk

`func (o *LogSchedule) GetHourScheduleOk() (*HourSchedule, bool)`

GetHourScheduleOk returns a tuple with the HourSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourSchedule

`func (o *LogSchedule) SetHourSchedule(v HourSchedule)`

SetHourSchedule sets HourSchedule field to given value.

### HasHourSchedule

`func (o *LogSchedule) HasHourSchedule() bool`

HasHourSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


