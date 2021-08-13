# HeliosLogSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | Pointer to **NullableString** | Specifies how often to start new Protection Group Runs of a Protection Group. &lt;br&gt;&#39;Minutes&#39; specifies that Protection Group run starts periodically after certain number of minutes specified in &#39;frequency&#39; field. &lt;br&gt;&#39;Hours&#39; specifies that Protection Group run starts periodically after certain number of hours specified in &#39;frequency&#39; field. | [optional] 
**MinuteSchedule** | Pointer to [**HeliosMinuteSchedule**](HeliosMinuteSchedule.md) |  | [optional] 
**HourSchedule** | Pointer to [**HeliosHourSchedule**](HeliosHourSchedule.md) |  | [optional] 

## Methods

### NewHeliosLogSchedule

`func NewHeliosLogSchedule() *HeliosLogSchedule`

NewHeliosLogSchedule instantiates a new HeliosLogSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosLogScheduleWithDefaults

`func NewHeliosLogScheduleWithDefaults() *HeliosLogSchedule`

NewHeliosLogScheduleWithDefaults instantiates a new HeliosLogSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *HeliosLogSchedule) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *HeliosLogSchedule) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *HeliosLogSchedule) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *HeliosLogSchedule) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *HeliosLogSchedule) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *HeliosLogSchedule) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetMinuteSchedule

`func (o *HeliosLogSchedule) GetMinuteSchedule() HeliosMinuteSchedule`

GetMinuteSchedule returns the MinuteSchedule field if non-nil, zero value otherwise.

### GetMinuteScheduleOk

`func (o *HeliosLogSchedule) GetMinuteScheduleOk() (*HeliosMinuteSchedule, bool)`

GetMinuteScheduleOk returns a tuple with the MinuteSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinuteSchedule

`func (o *HeliosLogSchedule) SetMinuteSchedule(v HeliosMinuteSchedule)`

SetMinuteSchedule sets MinuteSchedule field to given value.

### HasMinuteSchedule

`func (o *HeliosLogSchedule) HasMinuteSchedule() bool`

HasMinuteSchedule returns a boolean if a field has been set.

### GetHourSchedule

`func (o *HeliosLogSchedule) GetHourSchedule() HeliosHourSchedule`

GetHourSchedule returns the HourSchedule field if non-nil, zero value otherwise.

### GetHourScheduleOk

`func (o *HeliosLogSchedule) GetHourScheduleOk() (*HeliosHourSchedule, bool)`

GetHourScheduleOk returns a tuple with the HourSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourSchedule

`func (o *HeliosLogSchedule) SetHourSchedule(v HeliosHourSchedule)`

SetHourSchedule sets HourSchedule field to given value.

### HasHourSchedule

`func (o *HeliosLogSchedule) HasHourSchedule() bool`

HasHourSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


