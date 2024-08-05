# HeliosMinuteSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frequency** | **NullableInt64** | Specifies a factor to multiply the unit by, to determine the backup schedule. &lt;br&gt; Example: If &#39;frequency&#39; set to 2 and the unit is &#39;Hours&#39;, then Snapshots are backed up every 2 hours. If selected unit is &#39;Weeks&#39; or &#39;Months&#39; then frequency will only be applied if policy type is DMaas. | 

## Methods

### NewHeliosMinuteSchedule

`func NewHeliosMinuteSchedule(frequency NullableInt64, ) *HeliosMinuteSchedule`

NewHeliosMinuteSchedule instantiates a new HeliosMinuteSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosMinuteScheduleWithDefaults

`func NewHeliosMinuteScheduleWithDefaults() *HeliosMinuteSchedule`

NewHeliosMinuteScheduleWithDefaults instantiates a new HeliosMinuteSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequency

`func (o *HeliosMinuteSchedule) GetFrequency() int64`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *HeliosMinuteSchedule) GetFrequencyOk() (*int64, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *HeliosMinuteSchedule) SetFrequency(v int64)`

SetFrequency sets Frequency field to given value.


### SetFrequencyNil

`func (o *HeliosMinuteSchedule) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *HeliosMinuteSchedule) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


