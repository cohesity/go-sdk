# HeliosExtendedRetentionSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frequency** | Pointer to **NullableInt32** | Specifies a factor to multiply the unit by, to determine the retention schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is retained. | [optional] 
**Unit** | Pointer to **NullableString** | Specifies the unit interval for retention of Snapshots. &lt;br&gt;&#39;Runs&#39; means that the Snapshot copy retained after the number of Protection Group Runs equals the number specified in the frequency. &lt;br&gt;&#39;Hours&#39; means that the Snapshot copy retained hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. &lt;br&gt;&#39;Days&#39; means that the Snapshot copy gets retained daily at the frequency set in the frequency. &lt;br&gt;&#39;Weeks&#39; means that the Snapshot copy is retained weekly at the frequency set in the frequency. &lt;br&gt;&#39;Months&#39; means that the Snapshot copy is retained monthly at the frequency set in the Frequency. &lt;br&gt;&#39;Years&#39; means that the Snapshot copy is retained yearly at the frequency set in the Frequency. | [optional] 

## Methods

### NewHeliosExtendedRetentionSchedule

`func NewHeliosExtendedRetentionSchedule() *HeliosExtendedRetentionSchedule`

NewHeliosExtendedRetentionSchedule instantiates a new HeliosExtendedRetentionSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosExtendedRetentionScheduleWithDefaults

`func NewHeliosExtendedRetentionScheduleWithDefaults() *HeliosExtendedRetentionSchedule`

NewHeliosExtendedRetentionScheduleWithDefaults instantiates a new HeliosExtendedRetentionSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequency

`func (o *HeliosExtendedRetentionSchedule) GetFrequency() int32`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *HeliosExtendedRetentionSchedule) GetFrequencyOk() (*int32, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *HeliosExtendedRetentionSchedule) SetFrequency(v int32)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *HeliosExtendedRetentionSchedule) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *HeliosExtendedRetentionSchedule) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *HeliosExtendedRetentionSchedule) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetUnit

`func (o *HeliosExtendedRetentionSchedule) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *HeliosExtendedRetentionSchedule) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *HeliosExtendedRetentionSchedule) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *HeliosExtendedRetentionSchedule) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *HeliosExtendedRetentionSchedule) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *HeliosExtendedRetentionSchedule) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


