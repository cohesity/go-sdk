# EntitySchemaProtoTimeSeriesDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricDescriptiveName** | Pointer to **NullableString** | Specifies a descriptive name for the metric that is displayed in the Advanced Diagnostics of the Cohesity Dashboard. For example for the &#39;kUnmorphedUsageBytes&#39; metric, the descriptive name is \&quot;Total Logical Space Used\&quot;. | [optional] 
**MetricName** | Pointer to **NullableString** | Specifies the name of the metric such as &#39;kUnmorphedUsageBytes&#39;. It should be unique in an entity schema. | [optional] 
**MetricUnit** | Pointer to [**EntitySchemaProtoTimeSeriesDescriptorMetricUnit**](EntitySchemaProtoTimeSeriesDescriptorMetricUnit.md) |  | [optional] 
**RawMetricPublishIntervalHintSecs** | Pointer to **NullableInt32** | Specifies a suggestion for the interval to collect raw data points. | [optional] 
**TimeToLiveSecs** | Pointer to **NullableInt64** | Specifies how long the data point will be stored. Note: In statsv2, as timeseries data of an entity is stored per scribe row with metrics as columns, it is good to have time_to_live_secs per schema(defined below) For existing schemas, we will consider highest time_to_live_secs of all metrics as expiration time for all metrics defined in schema. | [optional] 
**ValueType** | Pointer to **NullableInt32** | Specifies the value type for this metric. A metric of type &#39;string\&quot; is not supported, instead use &#39;bytes&#39;. Note that an aggregate metric of type &#39;bytes&#39; is not supported. 0 specifies a value type of Int64. 1 specifies a value type of Double. 2 specifies a value type of String. 3 specifies a value type of Bytes. | [optional] 

## Methods

### NewEntitySchemaProtoTimeSeriesDescriptor

`func NewEntitySchemaProtoTimeSeriesDescriptor() *EntitySchemaProtoTimeSeriesDescriptor`

NewEntitySchemaProtoTimeSeriesDescriptor instantiates a new EntitySchemaProtoTimeSeriesDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitySchemaProtoTimeSeriesDescriptorWithDefaults

`func NewEntitySchemaProtoTimeSeriesDescriptorWithDefaults() *EntitySchemaProtoTimeSeriesDescriptor`

NewEntitySchemaProtoTimeSeriesDescriptorWithDefaults instantiates a new EntitySchemaProtoTimeSeriesDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricDescriptiveName

`func (o *EntitySchemaProtoTimeSeriesDescriptor) GetMetricDescriptiveName() string`

GetMetricDescriptiveName returns the MetricDescriptiveName field if non-nil, zero value otherwise.

### GetMetricDescriptiveNameOk

`func (o *EntitySchemaProtoTimeSeriesDescriptor) GetMetricDescriptiveNameOk() (*string, bool)`

GetMetricDescriptiveNameOk returns a tuple with the MetricDescriptiveName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDescriptiveName

`func (o *EntitySchemaProtoTimeSeriesDescriptor) SetMetricDescriptiveName(v string)`

SetMetricDescriptiveName sets MetricDescriptiveName field to given value.

### HasMetricDescriptiveName

`func (o *EntitySchemaProtoTimeSeriesDescriptor) HasMetricDescriptiveName() bool`

HasMetricDescriptiveName returns a boolean if a field has been set.

### SetMetricDescriptiveNameNil

`func (o *EntitySchemaProtoTimeSeriesDescriptor) SetMetricDescriptiveNameNil(b bool)`

 SetMetricDescriptiveNameNil sets the value for MetricDescriptiveName to be an explicit nil

### UnsetMetricDescriptiveName
`func (o *EntitySchemaProtoTimeSeriesDescriptor) UnsetMetricDescriptiveName()`

UnsetMetricDescriptiveName ensures that no value is present for MetricDescriptiveName, not even an explicit nil
### GetMetricName

`func (o *EntitySchemaProtoTimeSeriesDescriptor) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *EntitySchemaProtoTimeSeriesDescriptor) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *EntitySchemaProtoTimeSeriesDescriptor) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *EntitySchemaProtoTimeSeriesDescriptor) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### SetMetricNameNil

`func (o *EntitySchemaProtoTimeSeriesDescriptor) SetMetricNameNil(b bool)`

 SetMetricNameNil sets the value for MetricName to be an explicit nil

### UnsetMetricName
`func (o *EntitySchemaProtoTimeSeriesDescriptor) UnsetMetricName()`

UnsetMetricName ensures that no value is present for MetricName, not even an explicit nil
### GetMetricUnit

`func (o *EntitySchemaProtoTimeSeriesDescriptor) GetMetricUnit() EntitySchemaProtoTimeSeriesDescriptorMetricUnit`

GetMetricUnit returns the MetricUnit field if non-nil, zero value otherwise.

### GetMetricUnitOk

`func (o *EntitySchemaProtoTimeSeriesDescriptor) GetMetricUnitOk() (*EntitySchemaProtoTimeSeriesDescriptorMetricUnit, bool)`

GetMetricUnitOk returns a tuple with the MetricUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricUnit

`func (o *EntitySchemaProtoTimeSeriesDescriptor) SetMetricUnit(v EntitySchemaProtoTimeSeriesDescriptorMetricUnit)`

SetMetricUnit sets MetricUnit field to given value.

### HasMetricUnit

`func (o *EntitySchemaProtoTimeSeriesDescriptor) HasMetricUnit() bool`

HasMetricUnit returns a boolean if a field has been set.

### GetRawMetricPublishIntervalHintSecs

`func (o *EntitySchemaProtoTimeSeriesDescriptor) GetRawMetricPublishIntervalHintSecs() int32`

GetRawMetricPublishIntervalHintSecs returns the RawMetricPublishIntervalHintSecs field if non-nil, zero value otherwise.

### GetRawMetricPublishIntervalHintSecsOk

`func (o *EntitySchemaProtoTimeSeriesDescriptor) GetRawMetricPublishIntervalHintSecsOk() (*int32, bool)`

GetRawMetricPublishIntervalHintSecsOk returns a tuple with the RawMetricPublishIntervalHintSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawMetricPublishIntervalHintSecs

`func (o *EntitySchemaProtoTimeSeriesDescriptor) SetRawMetricPublishIntervalHintSecs(v int32)`

SetRawMetricPublishIntervalHintSecs sets RawMetricPublishIntervalHintSecs field to given value.

### HasRawMetricPublishIntervalHintSecs

`func (o *EntitySchemaProtoTimeSeriesDescriptor) HasRawMetricPublishIntervalHintSecs() bool`

HasRawMetricPublishIntervalHintSecs returns a boolean if a field has been set.

### SetRawMetricPublishIntervalHintSecsNil

`func (o *EntitySchemaProtoTimeSeriesDescriptor) SetRawMetricPublishIntervalHintSecsNil(b bool)`

 SetRawMetricPublishIntervalHintSecsNil sets the value for RawMetricPublishIntervalHintSecs to be an explicit nil

### UnsetRawMetricPublishIntervalHintSecs
`func (o *EntitySchemaProtoTimeSeriesDescriptor) UnsetRawMetricPublishIntervalHintSecs()`

UnsetRawMetricPublishIntervalHintSecs ensures that no value is present for RawMetricPublishIntervalHintSecs, not even an explicit nil
### GetTimeToLiveSecs

`func (o *EntitySchemaProtoTimeSeriesDescriptor) GetTimeToLiveSecs() int64`

GetTimeToLiveSecs returns the TimeToLiveSecs field if non-nil, zero value otherwise.

### GetTimeToLiveSecsOk

`func (o *EntitySchemaProtoTimeSeriesDescriptor) GetTimeToLiveSecsOk() (*int64, bool)`

GetTimeToLiveSecsOk returns a tuple with the TimeToLiveSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToLiveSecs

`func (o *EntitySchemaProtoTimeSeriesDescriptor) SetTimeToLiveSecs(v int64)`

SetTimeToLiveSecs sets TimeToLiveSecs field to given value.

### HasTimeToLiveSecs

`func (o *EntitySchemaProtoTimeSeriesDescriptor) HasTimeToLiveSecs() bool`

HasTimeToLiveSecs returns a boolean if a field has been set.

### SetTimeToLiveSecsNil

`func (o *EntitySchemaProtoTimeSeriesDescriptor) SetTimeToLiveSecsNil(b bool)`

 SetTimeToLiveSecsNil sets the value for TimeToLiveSecs to be an explicit nil

### UnsetTimeToLiveSecs
`func (o *EntitySchemaProtoTimeSeriesDescriptor) UnsetTimeToLiveSecs()`

UnsetTimeToLiveSecs ensures that no value is present for TimeToLiveSecs, not even an explicit nil
### GetValueType

`func (o *EntitySchemaProtoTimeSeriesDescriptor) GetValueType() int32`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *EntitySchemaProtoTimeSeriesDescriptor) GetValueTypeOk() (*int32, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *EntitySchemaProtoTimeSeriesDescriptor) SetValueType(v int32)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *EntitySchemaProtoTimeSeriesDescriptor) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### SetValueTypeNil

`func (o *EntitySchemaProtoTimeSeriesDescriptor) SetValueTypeNil(b bool)`

 SetValueTypeNil sets the value for ValueType to be an explicit nil

### UnsetValueType
`func (o *EntitySchemaProtoTimeSeriesDescriptor) UnsetValueType()`

UnsetValueType ensures that no value is present for ValueType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


