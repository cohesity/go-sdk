# FileDistributionMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricName** | Pointer to **NullableString** | Specifies the name of the metric. | [optional] 
**Value** | Pointer to **NullableInt64** | Specifies the value of specified metric name. | [optional] 

## Methods

### NewFileDistributionMetrics

`func NewFileDistributionMetrics() *FileDistributionMetrics`

NewFileDistributionMetrics instantiates a new FileDistributionMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileDistributionMetricsWithDefaults

`func NewFileDistributionMetricsWithDefaults() *FileDistributionMetrics`

NewFileDistributionMetricsWithDefaults instantiates a new FileDistributionMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricName

`func (o *FileDistributionMetrics) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *FileDistributionMetrics) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *FileDistributionMetrics) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *FileDistributionMetrics) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### SetMetricNameNil

`func (o *FileDistributionMetrics) SetMetricNameNil(b bool)`

 SetMetricNameNil sets the value for MetricName to be an explicit nil

### UnsetMetricName
`func (o *FileDistributionMetrics) UnsetMetricName()`

UnsetMetricName ensures that no value is present for MetricName, not even an explicit nil
### GetValue

`func (o *FileDistributionMetrics) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FileDistributionMetrics) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FileDistributionMetrics) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *FileDistributionMetrics) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *FileDistributionMetrics) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *FileDistributionMetrics) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


