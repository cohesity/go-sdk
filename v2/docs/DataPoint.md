# DataPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DoubleValue** | Pointer to **NullableFloat64** | Specifies the data point value in double format. | [optional] 
**Int64Value** | Pointer to **NullableInt64** | Specifies the data point value in int64 format. | [optional] 
**StringValue** | Pointer to **NullableString** | Specifies the data point value in string format. | [optional] 
**TimestampMsecs** | Pointer to **NullableInt64** | Specifies the timestamp of the data point. | [optional] 

## Methods

### NewDataPoint

`func NewDataPoint() *DataPoint`

NewDataPoint instantiates a new DataPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataPointWithDefaults

`func NewDataPointWithDefaults() *DataPoint`

NewDataPointWithDefaults instantiates a new DataPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoubleValue

`func (o *DataPoint) GetDoubleValue() float64`

GetDoubleValue returns the DoubleValue field if non-nil, zero value otherwise.

### GetDoubleValueOk

`func (o *DataPoint) GetDoubleValueOk() (*float64, bool)`

GetDoubleValueOk returns a tuple with the DoubleValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleValue

`func (o *DataPoint) SetDoubleValue(v float64)`

SetDoubleValue sets DoubleValue field to given value.

### HasDoubleValue

`func (o *DataPoint) HasDoubleValue() bool`

HasDoubleValue returns a boolean if a field has been set.

### SetDoubleValueNil

`func (o *DataPoint) SetDoubleValueNil(b bool)`

 SetDoubleValueNil sets the value for DoubleValue to be an explicit nil

### UnsetDoubleValue
`func (o *DataPoint) UnsetDoubleValue()`

UnsetDoubleValue ensures that no value is present for DoubleValue, not even an explicit nil
### GetInt64Value

`func (o *DataPoint) GetInt64Value() int64`

GetInt64Value returns the Int64Value field if non-nil, zero value otherwise.

### GetInt64ValueOk

`func (o *DataPoint) GetInt64ValueOk() (*int64, bool)`

GetInt64ValueOk returns a tuple with the Int64Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInt64Value

`func (o *DataPoint) SetInt64Value(v int64)`

SetInt64Value sets Int64Value field to given value.

### HasInt64Value

`func (o *DataPoint) HasInt64Value() bool`

HasInt64Value returns a boolean if a field has been set.

### SetInt64ValueNil

`func (o *DataPoint) SetInt64ValueNil(b bool)`

 SetInt64ValueNil sets the value for Int64Value to be an explicit nil

### UnsetInt64Value
`func (o *DataPoint) UnsetInt64Value()`

UnsetInt64Value ensures that no value is present for Int64Value, not even an explicit nil
### GetStringValue

`func (o *DataPoint) GetStringValue() string`

GetStringValue returns the StringValue field if non-nil, zero value otherwise.

### GetStringValueOk

`func (o *DataPoint) GetStringValueOk() (*string, bool)`

GetStringValueOk returns a tuple with the StringValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValue

`func (o *DataPoint) SetStringValue(v string)`

SetStringValue sets StringValue field to given value.

### HasStringValue

`func (o *DataPoint) HasStringValue() bool`

HasStringValue returns a boolean if a field has been set.

### SetStringValueNil

`func (o *DataPoint) SetStringValueNil(b bool)`

 SetStringValueNil sets the value for StringValue to be an explicit nil

### UnsetStringValue
`func (o *DataPoint) UnsetStringValue()`

UnsetStringValue ensures that no value is present for StringValue, not even an explicit nil
### GetTimestampMsecs

`func (o *DataPoint) GetTimestampMsecs() int64`

GetTimestampMsecs returns the TimestampMsecs field if non-nil, zero value otherwise.

### GetTimestampMsecsOk

`func (o *DataPoint) GetTimestampMsecsOk() (*int64, bool)`

GetTimestampMsecsOk returns a tuple with the TimestampMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampMsecs

`func (o *DataPoint) SetTimestampMsecs(v int64)`

SetTimestampMsecs sets TimestampMsecs field to given value.

### HasTimestampMsecs

`func (o *DataPoint) HasTimestampMsecs() bool`

HasTimestampMsecs returns a boolean if a field has been set.

### SetTimestampMsecsNil

`func (o *DataPoint) SetTimestampMsecsNil(b bool)`

 SetTimestampMsecsNil sets the value for TimestampMsecs to be an explicit nil

### UnsetTimestampMsecs
`func (o *DataPoint) UnsetTimestampMsecs()`

UnsetTimestampMsecs ensures that no value is present for TimestampMsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


