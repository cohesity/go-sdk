# MetricValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricName** | Pointer to **NullableString** | Specifies the metric name. | [optional] 
**TimestampMsecs** | Pointer to **NullableInt64** | Specifies the creation time of a data point as a Unix epoch Timestamp (in milliseconds). | [optional] 
**Value** | Pointer to [**Value**](Value.md) |  | [optional] 

## Methods

### NewMetricValue

`func NewMetricValue() *MetricValue`

NewMetricValue instantiates a new MetricValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricValueWithDefaults

`func NewMetricValueWithDefaults() *MetricValue`

NewMetricValueWithDefaults instantiates a new MetricValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricName

`func (o *MetricValue) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *MetricValue) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *MetricValue) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *MetricValue) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### SetMetricNameNil

`func (o *MetricValue) SetMetricNameNil(b bool)`

 SetMetricNameNil sets the value for MetricName to be an explicit nil

### UnsetMetricName
`func (o *MetricValue) UnsetMetricName()`

UnsetMetricName ensures that no value is present for MetricName, not even an explicit nil
### GetTimestampMsecs

`func (o *MetricValue) GetTimestampMsecs() int64`

GetTimestampMsecs returns the TimestampMsecs field if non-nil, zero value otherwise.

### GetTimestampMsecsOk

`func (o *MetricValue) GetTimestampMsecsOk() (*int64, bool)`

GetTimestampMsecsOk returns a tuple with the TimestampMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampMsecs

`func (o *MetricValue) SetTimestampMsecs(v int64)`

SetTimestampMsecs sets TimestampMsecs field to given value.

### HasTimestampMsecs

`func (o *MetricValue) HasTimestampMsecs() bool`

HasTimestampMsecs returns a boolean if a field has been set.

### SetTimestampMsecsNil

`func (o *MetricValue) SetTimestampMsecsNil(b bool)`

 SetTimestampMsecsNil sets the value for TimestampMsecs to be an explicit nil

### UnsetTimestampMsecs
`func (o *MetricValue) UnsetTimestampMsecs()`

UnsetTimestampMsecs ensures that no value is present for TimestampMsecs, not even an explicit nil
### GetValue

`func (o *MetricValue) GetValue() Value`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MetricValue) GetValueOk() (*Value, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MetricValue) SetValue(v Value)`

SetValue sets Value field to given value.

### HasValue

`func (o *MetricValue) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


