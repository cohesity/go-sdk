# Sample

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FloatValue** | Pointer to **NullableFloat64** | Specifies the value of the data sample if the type is float64. This field is nil if the type of the data is not a float value. | [optional] 
**TimestampMsecs** | Pointer to **NullableInt64** | Specifies the timestamp when the data sample occured. | [optional] 
**Value** | Pointer to **NullableInt64** | Specifies the value of the data sample if the type is int64. This field is nil if the type of the data is not an int value. | [optional] 

## Methods

### NewSample

`func NewSample() *Sample`

NewSample instantiates a new Sample object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSampleWithDefaults

`func NewSampleWithDefaults() *Sample`

NewSampleWithDefaults instantiates a new Sample object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFloatValue

`func (o *Sample) GetFloatValue() float64`

GetFloatValue returns the FloatValue field if non-nil, zero value otherwise.

### GetFloatValueOk

`func (o *Sample) GetFloatValueOk() (*float64, bool)`

GetFloatValueOk returns a tuple with the FloatValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatValue

`func (o *Sample) SetFloatValue(v float64)`

SetFloatValue sets FloatValue field to given value.

### HasFloatValue

`func (o *Sample) HasFloatValue() bool`

HasFloatValue returns a boolean if a field has been set.

### SetFloatValueNil

`func (o *Sample) SetFloatValueNil(b bool)`

 SetFloatValueNil sets the value for FloatValue to be an explicit nil

### UnsetFloatValue
`func (o *Sample) UnsetFloatValue()`

UnsetFloatValue ensures that no value is present for FloatValue, not even an explicit nil
### GetTimestampMsecs

`func (o *Sample) GetTimestampMsecs() int64`

GetTimestampMsecs returns the TimestampMsecs field if non-nil, zero value otherwise.

### GetTimestampMsecsOk

`func (o *Sample) GetTimestampMsecsOk() (*int64, bool)`

GetTimestampMsecsOk returns a tuple with the TimestampMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampMsecs

`func (o *Sample) SetTimestampMsecs(v int64)`

SetTimestampMsecs sets TimestampMsecs field to given value.

### HasTimestampMsecs

`func (o *Sample) HasTimestampMsecs() bool`

HasTimestampMsecs returns a boolean if a field has been set.

### SetTimestampMsecsNil

`func (o *Sample) SetTimestampMsecsNil(b bool)`

 SetTimestampMsecsNil sets the value for TimestampMsecs to be an explicit nil

### UnsetTimestampMsecs
`func (o *Sample) UnsetTimestampMsecs()`

UnsetTimestampMsecs ensures that no value is present for TimestampMsecs, not even an explicit nil
### GetValue

`func (o *Sample) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Sample) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Sample) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *Sample) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *Sample) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Sample) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


