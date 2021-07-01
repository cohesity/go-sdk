# MetricDataPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ValueData**](ValueData.md) |  | [optional] 
**RollupFunction** | Pointer to **NullableInt32** | If this is a rolled up data point, following enum denotes the rollup function used for rolling up. For a raw point this enum is not set. | [optional] 
**TimestampMsecs** | Pointer to **NullableInt64** | Specifies a timestamp when the metric data point was captured. | [optional] 

## Methods

### NewMetricDataPoint

`func NewMetricDataPoint() *MetricDataPoint`

NewMetricDataPoint instantiates a new MetricDataPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDataPointWithDefaults

`func NewMetricDataPointWithDefaults() *MetricDataPoint`

NewMetricDataPointWithDefaults instantiates a new MetricDataPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MetricDataPoint) GetData() ValueData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MetricDataPoint) GetDataOk() (*ValueData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MetricDataPoint) SetData(v ValueData)`

SetData sets Data field to given value.

### HasData

`func (o *MetricDataPoint) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRollupFunction

`func (o *MetricDataPoint) GetRollupFunction() int32`

GetRollupFunction returns the RollupFunction field if non-nil, zero value otherwise.

### GetRollupFunctionOk

`func (o *MetricDataPoint) GetRollupFunctionOk() (*int32, bool)`

GetRollupFunctionOk returns a tuple with the RollupFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollupFunction

`func (o *MetricDataPoint) SetRollupFunction(v int32)`

SetRollupFunction sets RollupFunction field to given value.

### HasRollupFunction

`func (o *MetricDataPoint) HasRollupFunction() bool`

HasRollupFunction returns a boolean if a field has been set.

### SetRollupFunctionNil

`func (o *MetricDataPoint) SetRollupFunctionNil(b bool)`

 SetRollupFunctionNil sets the value for RollupFunction to be an explicit nil

### UnsetRollupFunction
`func (o *MetricDataPoint) UnsetRollupFunction()`

UnsetRollupFunction ensures that no value is present for RollupFunction, not even an explicit nil
### GetTimestampMsecs

`func (o *MetricDataPoint) GetTimestampMsecs() int64`

GetTimestampMsecs returns the TimestampMsecs field if non-nil, zero value otherwise.

### GetTimestampMsecsOk

`func (o *MetricDataPoint) GetTimestampMsecsOk() (*int64, bool)`

GetTimestampMsecsOk returns a tuple with the TimestampMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampMsecs

`func (o *MetricDataPoint) SetTimestampMsecs(v int64)`

SetTimestampMsecs sets TimestampMsecs field to given value.

### HasTimestampMsecs

`func (o *MetricDataPoint) HasTimestampMsecs() bool`

HasTimestampMsecs returns a boolean if a field has been set.

### SetTimestampMsecsNil

`func (o *MetricDataPoint) SetTimestampMsecsNil(b bool)`

 SetTimestampMsecsNil sets the value for TimestampMsecs to be an explicit nil

### UnsetTimestampMsecs
`func (o *MetricDataPoint) UnsetTimestampMsecs()`

UnsetTimestampMsecs ensures that no value is present for TimestampMsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


