# RpoSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntervalUnit** | Pointer to **NullableString** | Specifies an RPO policy interval unit which will be used along with the multiplier to calculate the interval for the RPO policy execution. this can be kHours, kDays, KWeeks, kMonths RPOIntervalUnit.  Specifies an RPO Schedule interval unit. kMinutes specifies that the rpo interval unit is hours. kHours specifies that the rpo interval unit is hours. kDays specifies that the rpo interval unit is days. kWeeks specifies that the rpo interval unit is weeks. kMonths specifies that the rpo interval unit is months. | [optional] 
**Multiplier** | Pointer to **NullableInt64** | Specifies the multiplier value to be used with the  RPO interval unit value. | [optional] 

## Methods

### NewRpoSchedule

`func NewRpoSchedule() *RpoSchedule`

NewRpoSchedule instantiates a new RpoSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpoScheduleWithDefaults

`func NewRpoScheduleWithDefaults() *RpoSchedule`

NewRpoScheduleWithDefaults instantiates a new RpoSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntervalUnit

`func (o *RpoSchedule) GetIntervalUnit() string`

GetIntervalUnit returns the IntervalUnit field if non-nil, zero value otherwise.

### GetIntervalUnitOk

`func (o *RpoSchedule) GetIntervalUnitOk() (*string, bool)`

GetIntervalUnitOk returns a tuple with the IntervalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalUnit

`func (o *RpoSchedule) SetIntervalUnit(v string)`

SetIntervalUnit sets IntervalUnit field to given value.

### HasIntervalUnit

`func (o *RpoSchedule) HasIntervalUnit() bool`

HasIntervalUnit returns a boolean if a field has been set.

### SetIntervalUnitNil

`func (o *RpoSchedule) SetIntervalUnitNil(b bool)`

 SetIntervalUnitNil sets the value for IntervalUnit to be an explicit nil

### UnsetIntervalUnit
`func (o *RpoSchedule) UnsetIntervalUnit()`

UnsetIntervalUnit ensures that no value is present for IntervalUnit, not even an explicit nil
### GetMultiplier

`func (o *RpoSchedule) GetMultiplier() int64`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *RpoSchedule) GetMultiplierOk() (*int64, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *RpoSchedule) SetMultiplier(v int64)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *RpoSchedule) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.

### SetMultiplierNil

`func (o *RpoSchedule) SetMultiplierNil(b bool)`

 SetMultiplierNil sets the value for Multiplier to be an explicit nil

### UnsetMultiplier
`func (o *RpoSchedule) UnsetMultiplier()`

UnsetMultiplier ensures that no value is present for Multiplier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


