# ViewTimeSeriesStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimestampMsecs** | Pointer to **NullableInt64** | Specifies the timestamp in milliseconds. | [optional] 
**Value** | Pointer to **NullableInt64** | Specifies the value of the specified metric at the timestamp. | [optional] 

## Methods

### NewViewTimeSeriesStats

`func NewViewTimeSeriesStats() *ViewTimeSeriesStats`

NewViewTimeSeriesStats instantiates a new ViewTimeSeriesStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewTimeSeriesStatsWithDefaults

`func NewViewTimeSeriesStatsWithDefaults() *ViewTimeSeriesStats`

NewViewTimeSeriesStatsWithDefaults instantiates a new ViewTimeSeriesStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestampMsecs

`func (o *ViewTimeSeriesStats) GetTimestampMsecs() int64`

GetTimestampMsecs returns the TimestampMsecs field if non-nil, zero value otherwise.

### GetTimestampMsecsOk

`func (o *ViewTimeSeriesStats) GetTimestampMsecsOk() (*int64, bool)`

GetTimestampMsecsOk returns a tuple with the TimestampMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampMsecs

`func (o *ViewTimeSeriesStats) SetTimestampMsecs(v int64)`

SetTimestampMsecs sets TimestampMsecs field to given value.

### HasTimestampMsecs

`func (o *ViewTimeSeriesStats) HasTimestampMsecs() bool`

HasTimestampMsecs returns a boolean if a field has been set.

### SetTimestampMsecsNil

`func (o *ViewTimeSeriesStats) SetTimestampMsecsNil(b bool)`

 SetTimestampMsecsNil sets the value for TimestampMsecs to be an explicit nil

### UnsetTimestampMsecs
`func (o *ViewTimeSeriesStats) UnsetTimestampMsecs()`

UnsetTimestampMsecs ensures that no value is present for TimestampMsecs, not even an explicit nil
### GetValue

`func (o *ViewTimeSeriesStats) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ViewTimeSeriesStats) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ViewTimeSeriesStats) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *ViewTimeSeriesStats) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ViewTimeSeriesStats) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ViewTimeSeriesStats) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


