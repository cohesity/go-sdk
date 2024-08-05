# ClientStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | Pointer to **NullableString** | Specifies the stats metric. | [optional] 
**Value** | Pointer to **NullableInt64** | Specifies the stats value. | [optional] 
**ValueInLastHours** | Pointer to [**[]ClientStatsInLastHours**](ClientStatsInLastHours.md) | Specifies the stats value in last hours. | [optional] 

## Methods

### NewClientStats

`func NewClientStats() *ClientStats`

NewClientStats instantiates a new ClientStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientStatsWithDefaults

`func NewClientStatsWithDefaults() *ClientStats`

NewClientStatsWithDefaults instantiates a new ClientStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *ClientStats) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *ClientStats) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *ClientStats) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *ClientStats) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### SetMetricNil

`func (o *ClientStats) SetMetricNil(b bool)`

 SetMetricNil sets the value for Metric to be an explicit nil

### UnsetMetric
`func (o *ClientStats) UnsetMetric()`

UnsetMetric ensures that no value is present for Metric, not even an explicit nil
### GetValue

`func (o *ClientStats) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ClientStats) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ClientStats) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *ClientStats) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ClientStats) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ClientStats) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetValueInLastHours

`func (o *ClientStats) GetValueInLastHours() []ClientStatsInLastHours`

GetValueInLastHours returns the ValueInLastHours field if non-nil, zero value otherwise.

### GetValueInLastHoursOk

`func (o *ClientStats) GetValueInLastHoursOk() (*[]ClientStatsInLastHours, bool)`

GetValueInLastHoursOk returns a tuple with the ValueInLastHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueInLastHours

`func (o *ClientStats) SetValueInLastHours(v []ClientStatsInLastHours)`

SetValueInLastHours sets ValueInLastHours field to given value.

### HasValueInLastHours

`func (o *ClientStats) HasValueInLastHours() bool`

HasValueInLastHours returns a boolean if a field has been set.

### SetValueInLastHoursNil

`func (o *ClientStats) SetValueInLastHoursNil(b bool)`

 SetValueInLastHoursNil sets the value for ValueInLastHours to be an explicit nil

### UnsetValueInLastHours
`func (o *ClientStats) UnsetValueInLastHours()`

UnsetValueInLastHours ensures that no value is present for ValueInLastHours, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


