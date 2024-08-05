# DaySchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frequency** | **NullableInt64** | Specifies a factor to multiply the unit by, to determine the backup schedule. &lt;br&gt; Example: If &#39;frequency&#39; set to 2 and the unit is &#39;Hours&#39;, then Snapshots are backed up every 2 hours. &lt;br&gt; This field is only applicable if unit is &#39;Minutes&#39;, &#39;Hours&#39; or &#39;Days&#39;. | 

## Methods

### NewDaySchedule

`func NewDaySchedule(frequency NullableInt64, ) *DaySchedule`

NewDaySchedule instantiates a new DaySchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDayScheduleWithDefaults

`func NewDayScheduleWithDefaults() *DaySchedule`

NewDayScheduleWithDefaults instantiates a new DaySchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequency

`func (o *DaySchedule) GetFrequency() int64`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *DaySchedule) GetFrequencyOk() (*int64, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *DaySchedule) SetFrequency(v int64)`

SetFrequency sets Frequency field to given value.


### SetFrequencyNil

`func (o *DaySchedule) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *DaySchedule) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


