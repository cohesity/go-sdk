# TimeSeriesStatsForMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataPoints** | Pointer to [**[]DataPoint**](DataPoint.md) | Specifies a list of data points. | [optional] 
**MetricName** | Pointer to **NullableString** | Specifies the metric name. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the data points. | [optional] 

## Methods

### NewTimeSeriesStatsForMetric

`func NewTimeSeriesStatsForMetric() *TimeSeriesStatsForMetric`

NewTimeSeriesStatsForMetric instantiates a new TimeSeriesStatsForMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSeriesStatsForMetricWithDefaults

`func NewTimeSeriesStatsForMetricWithDefaults() *TimeSeriesStatsForMetric`

NewTimeSeriesStatsForMetricWithDefaults instantiates a new TimeSeriesStatsForMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataPoints

`func (o *TimeSeriesStatsForMetric) GetDataPoints() []DataPoint`

GetDataPoints returns the DataPoints field if non-nil, zero value otherwise.

### GetDataPointsOk

`func (o *TimeSeriesStatsForMetric) GetDataPointsOk() (*[]DataPoint, bool)`

GetDataPointsOk returns a tuple with the DataPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPoints

`func (o *TimeSeriesStatsForMetric) SetDataPoints(v []DataPoint)`

SetDataPoints sets DataPoints field to given value.

### HasDataPoints

`func (o *TimeSeriesStatsForMetric) HasDataPoints() bool`

HasDataPoints returns a boolean if a field has been set.

### SetDataPointsNil

`func (o *TimeSeriesStatsForMetric) SetDataPointsNil(b bool)`

 SetDataPointsNil sets the value for DataPoints to be an explicit nil

### UnsetDataPoints
`func (o *TimeSeriesStatsForMetric) UnsetDataPoints()`

UnsetDataPoints ensures that no value is present for DataPoints, not even an explicit nil
### GetMetricName

`func (o *TimeSeriesStatsForMetric) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *TimeSeriesStatsForMetric) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *TimeSeriesStatsForMetric) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *TimeSeriesStatsForMetric) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### SetMetricNameNil

`func (o *TimeSeriesStatsForMetric) SetMetricNameNil(b bool)`

 SetMetricNameNil sets the value for MetricName to be an explicit nil

### UnsetMetricName
`func (o *TimeSeriesStatsForMetric) UnsetMetricName()`

UnsetMetricName ensures that no value is present for MetricName, not even an explicit nil
### GetType

`func (o *TimeSeriesStatsForMetric) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimeSeriesStatsForMetric) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimeSeriesStatsForMetric) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TimeSeriesStatsForMetric) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *TimeSeriesStatsForMetric) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *TimeSeriesStatsForMetric) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


