# HeliosBmrSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonthSchedule** | Pointer to [**HeliosMonthSchedule**](HeliosMonthSchedule.md) |  | [optional] 
**Unit** | Pointer to **NullableString** | Specifies how often to start new runs of a Protection Group. &lt;br&gt;&#39;Weeks&#39; specifies that new Protection Group runs start weekly on certain days specified using &#39;dayOfWeek&#39; field. &lt;br&gt;&#39;Months&#39; specifies that new Protection Group runs start monthly on certain day of specific week. | [optional] 
**WeekSchedule** | Pointer to [**HeliosWeekSchedule**](HeliosWeekSchedule.md) |  | [optional] 

## Methods

### NewHeliosBmrSchedule

`func NewHeliosBmrSchedule() *HeliosBmrSchedule`

NewHeliosBmrSchedule instantiates a new HeliosBmrSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosBmrScheduleWithDefaults

`func NewHeliosBmrScheduleWithDefaults() *HeliosBmrSchedule`

NewHeliosBmrScheduleWithDefaults instantiates a new HeliosBmrSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonthSchedule

`func (o *HeliosBmrSchedule) GetMonthSchedule() HeliosMonthSchedule`

GetMonthSchedule returns the MonthSchedule field if non-nil, zero value otherwise.

### GetMonthScheduleOk

`func (o *HeliosBmrSchedule) GetMonthScheduleOk() (*HeliosMonthSchedule, bool)`

GetMonthScheduleOk returns a tuple with the MonthSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthSchedule

`func (o *HeliosBmrSchedule) SetMonthSchedule(v HeliosMonthSchedule)`

SetMonthSchedule sets MonthSchedule field to given value.

### HasMonthSchedule

`func (o *HeliosBmrSchedule) HasMonthSchedule() bool`

HasMonthSchedule returns a boolean if a field has been set.

### GetUnit

`func (o *HeliosBmrSchedule) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *HeliosBmrSchedule) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *HeliosBmrSchedule) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *HeliosBmrSchedule) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *HeliosBmrSchedule) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *HeliosBmrSchedule) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetWeekSchedule

`func (o *HeliosBmrSchedule) GetWeekSchedule() HeliosWeekSchedule`

GetWeekSchedule returns the WeekSchedule field if non-nil, zero value otherwise.

### GetWeekScheduleOk

`func (o *HeliosBmrSchedule) GetWeekScheduleOk() (*HeliosWeekSchedule, bool)`

GetWeekScheduleOk returns a tuple with the WeekSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekSchedule

`func (o *HeliosBmrSchedule) SetWeekSchedule(v HeliosWeekSchedule)`

SetWeekSchedule sets WeekSchedule field to given value.

### HasWeekSchedule

`func (o *HeliosBmrSchedule) HasWeekSchedule() bool`

HasWeekSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


