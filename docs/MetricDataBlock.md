# MetricDataBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataPointVec** | Pointer to [**[]MetricDataPoint**](MetricDataPoint.md) | Array of Data Points.  Specifies a list of metric data points for a time series. | [optional] 
**MetricName** | Pointer to **NullableString** | Specifies the name of a metric such as &#39;kDiskAwaitTimeMsecs&#39;. | [optional] 
**Type** | Pointer to **NullableInt32** | Specifies the data type of the data points. 0 specifies a data point of type Int64. 1 specifies a data point of type Double. 2 specifies a data point of type String. 3 specifies a data point of type Bytes. | [optional] 

## Methods

### NewMetricDataBlock

`func NewMetricDataBlock() *MetricDataBlock`

NewMetricDataBlock instantiates a new MetricDataBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDataBlockWithDefaults

`func NewMetricDataBlockWithDefaults() *MetricDataBlock`

NewMetricDataBlockWithDefaults instantiates a new MetricDataBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataPointVec

`func (o *MetricDataBlock) GetDataPointVec() []MetricDataPoint`

GetDataPointVec returns the DataPointVec field if non-nil, zero value otherwise.

### GetDataPointVecOk

`func (o *MetricDataBlock) GetDataPointVecOk() (*[]MetricDataPoint, bool)`

GetDataPointVecOk returns a tuple with the DataPointVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPointVec

`func (o *MetricDataBlock) SetDataPointVec(v []MetricDataPoint)`

SetDataPointVec sets DataPointVec field to given value.

### HasDataPointVec

`func (o *MetricDataBlock) HasDataPointVec() bool`

HasDataPointVec returns a boolean if a field has been set.

### SetDataPointVecNil

`func (o *MetricDataBlock) SetDataPointVecNil(b bool)`

 SetDataPointVecNil sets the value for DataPointVec to be an explicit nil

### UnsetDataPointVec
`func (o *MetricDataBlock) UnsetDataPointVec()`

UnsetDataPointVec ensures that no value is present for DataPointVec, not even an explicit nil
### GetMetricName

`func (o *MetricDataBlock) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *MetricDataBlock) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *MetricDataBlock) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *MetricDataBlock) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### SetMetricNameNil

`func (o *MetricDataBlock) SetMetricNameNil(b bool)`

 SetMetricNameNil sets the value for MetricName to be an explicit nil

### UnsetMetricName
`func (o *MetricDataBlock) UnsetMetricName()`

UnsetMetricName ensures that no value is present for MetricName, not even an explicit nil
### GetType

`func (o *MetricDataBlock) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricDataBlock) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricDataBlock) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *MetricDataBlock) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MetricDataBlock) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MetricDataBlock) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


