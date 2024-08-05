# HeliosTargetSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frequency** | Pointer to **NullableInt32** | Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied. | [optional] 
**Unit** | Pointer to **NullableString** | Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. &lt;br&gt;&#39;Runs&#39; means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. &lt;br&gt;&#39;Hours&#39; means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. &lt;br&gt;&#39;Days&#39; means that the Snapshot copy occurs daily at the frequency set in the frequency. &lt;br&gt;&#39;Weeks&#39; means that the Snapshot copy occurs weekly at the frequency set in the frequency. &lt;br&gt;&#39;Months&#39; means that the Snapshot copy occurs monthly at the frequency set in the Frequency. &lt;br&gt;&#39;Years&#39; means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency. | [optional] 

## Methods

### NewHeliosTargetSchedule

`func NewHeliosTargetSchedule() *HeliosTargetSchedule`

NewHeliosTargetSchedule instantiates a new HeliosTargetSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosTargetScheduleWithDefaults

`func NewHeliosTargetScheduleWithDefaults() *HeliosTargetSchedule`

NewHeliosTargetScheduleWithDefaults instantiates a new HeliosTargetSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequency

`func (o *HeliosTargetSchedule) GetFrequency() int32`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *HeliosTargetSchedule) GetFrequencyOk() (*int32, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *HeliosTargetSchedule) SetFrequency(v int32)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *HeliosTargetSchedule) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *HeliosTargetSchedule) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *HeliosTargetSchedule) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetUnit

`func (o *HeliosTargetSchedule) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *HeliosTargetSchedule) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *HeliosTargetSchedule) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *HeliosTargetSchedule) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *HeliosTargetSchedule) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *HeliosTargetSchedule) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


