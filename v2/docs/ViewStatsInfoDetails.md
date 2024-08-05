# ViewStatsInfoDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | Pointer to **NullableString** | Specifies the stats metric. | [optional] 
**ValueInLastHours** | Pointer to [**[]ViewStatsInLastHours**](ViewStatsInLastHours.md) | Specifies the stats value in last hours. | [optional] 

## Methods

### NewViewStatsInfoDetails

`func NewViewStatsInfoDetails() *ViewStatsInfoDetails`

NewViewStatsInfoDetails instantiates a new ViewStatsInfoDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewStatsInfoDetailsWithDefaults

`func NewViewStatsInfoDetailsWithDefaults() *ViewStatsInfoDetails`

NewViewStatsInfoDetailsWithDefaults instantiates a new ViewStatsInfoDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *ViewStatsInfoDetails) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *ViewStatsInfoDetails) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *ViewStatsInfoDetails) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *ViewStatsInfoDetails) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### SetMetricNil

`func (o *ViewStatsInfoDetails) SetMetricNil(b bool)`

 SetMetricNil sets the value for Metric to be an explicit nil

### UnsetMetric
`func (o *ViewStatsInfoDetails) UnsetMetric()`

UnsetMetric ensures that no value is present for Metric, not even an explicit nil
### GetValueInLastHours

`func (o *ViewStatsInfoDetails) GetValueInLastHours() []ViewStatsInLastHours`

GetValueInLastHours returns the ValueInLastHours field if non-nil, zero value otherwise.

### GetValueInLastHoursOk

`func (o *ViewStatsInfoDetails) GetValueInLastHoursOk() (*[]ViewStatsInLastHours, bool)`

GetValueInLastHoursOk returns a tuple with the ValueInLastHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueInLastHours

`func (o *ViewStatsInfoDetails) SetValueInLastHours(v []ViewStatsInLastHours)`

SetValueInLastHours sets ValueInLastHours field to given value.

### HasValueInLastHours

`func (o *ViewStatsInfoDetails) HasValueInLastHours() bool`

HasValueInLastHours returns a boolean if a field has been set.

### SetValueInLastHoursNil

`func (o *ViewStatsInfoDetails) SetValueInLastHoursNil(b bool)`

 SetValueInLastHoursNil sets the value for ValueInLastHours to be an explicit nil

### UnsetValueInLastHours
`func (o *ViewStatsInfoDetails) UnsetValueInLastHours()`

UnsetValueInLastHours ensures that no value is present for ValueInLastHours, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


