# ThroughputTile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxReadThroughput** | Pointer to **NullableInt64** | Maxium Read throughput in last 24 hours. | [optional] 
**MaxWriteThroughput** | Pointer to **NullableInt64** | Maximum Write throughput in last 24 hours. | [optional] 
**ReadThroughputSamples** | Pointer to [**[]Sample**](Sample.md) | Read throughput samples taken for the past 24 hours at 10 minutes interval given in descending order of time. | [optional] 
**WriteThroughputSamples** | Pointer to [**[]Sample**](Sample.md) | Write throughput samples taken for the past 24 hours at 10 minutes interval given in descending order of time. | [optional] 

## Methods

### NewThroughputTile

`func NewThroughputTile() *ThroughputTile`

NewThroughputTile instantiates a new ThroughputTile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThroughputTileWithDefaults

`func NewThroughputTileWithDefaults() *ThroughputTile`

NewThroughputTileWithDefaults instantiates a new ThroughputTile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxReadThroughput

`func (o *ThroughputTile) GetMaxReadThroughput() int64`

GetMaxReadThroughput returns the MaxReadThroughput field if non-nil, zero value otherwise.

### GetMaxReadThroughputOk

`func (o *ThroughputTile) GetMaxReadThroughputOk() (*int64, bool)`

GetMaxReadThroughputOk returns a tuple with the MaxReadThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReadThroughput

`func (o *ThroughputTile) SetMaxReadThroughput(v int64)`

SetMaxReadThroughput sets MaxReadThroughput field to given value.

### HasMaxReadThroughput

`func (o *ThroughputTile) HasMaxReadThroughput() bool`

HasMaxReadThroughput returns a boolean if a field has been set.

### SetMaxReadThroughputNil

`func (o *ThroughputTile) SetMaxReadThroughputNil(b bool)`

 SetMaxReadThroughputNil sets the value for MaxReadThroughput to be an explicit nil

### UnsetMaxReadThroughput
`func (o *ThroughputTile) UnsetMaxReadThroughput()`

UnsetMaxReadThroughput ensures that no value is present for MaxReadThroughput, not even an explicit nil
### GetMaxWriteThroughput

`func (o *ThroughputTile) GetMaxWriteThroughput() int64`

GetMaxWriteThroughput returns the MaxWriteThroughput field if non-nil, zero value otherwise.

### GetMaxWriteThroughputOk

`func (o *ThroughputTile) GetMaxWriteThroughputOk() (*int64, bool)`

GetMaxWriteThroughputOk returns a tuple with the MaxWriteThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWriteThroughput

`func (o *ThroughputTile) SetMaxWriteThroughput(v int64)`

SetMaxWriteThroughput sets MaxWriteThroughput field to given value.

### HasMaxWriteThroughput

`func (o *ThroughputTile) HasMaxWriteThroughput() bool`

HasMaxWriteThroughput returns a boolean if a field has been set.

### SetMaxWriteThroughputNil

`func (o *ThroughputTile) SetMaxWriteThroughputNil(b bool)`

 SetMaxWriteThroughputNil sets the value for MaxWriteThroughput to be an explicit nil

### UnsetMaxWriteThroughput
`func (o *ThroughputTile) UnsetMaxWriteThroughput()`

UnsetMaxWriteThroughput ensures that no value is present for MaxWriteThroughput, not even an explicit nil
### GetReadThroughputSamples

`func (o *ThroughputTile) GetReadThroughputSamples() []Sample`

GetReadThroughputSamples returns the ReadThroughputSamples field if non-nil, zero value otherwise.

### GetReadThroughputSamplesOk

`func (o *ThroughputTile) GetReadThroughputSamplesOk() (*[]Sample, bool)`

GetReadThroughputSamplesOk returns a tuple with the ReadThroughputSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadThroughputSamples

`func (o *ThroughputTile) SetReadThroughputSamples(v []Sample)`

SetReadThroughputSamples sets ReadThroughputSamples field to given value.

### HasReadThroughputSamples

`func (o *ThroughputTile) HasReadThroughputSamples() bool`

HasReadThroughputSamples returns a boolean if a field has been set.

### SetReadThroughputSamplesNil

`func (o *ThroughputTile) SetReadThroughputSamplesNil(b bool)`

 SetReadThroughputSamplesNil sets the value for ReadThroughputSamples to be an explicit nil

### UnsetReadThroughputSamples
`func (o *ThroughputTile) UnsetReadThroughputSamples()`

UnsetReadThroughputSamples ensures that no value is present for ReadThroughputSamples, not even an explicit nil
### GetWriteThroughputSamples

`func (o *ThroughputTile) GetWriteThroughputSamples() []Sample`

GetWriteThroughputSamples returns the WriteThroughputSamples field if non-nil, zero value otherwise.

### GetWriteThroughputSamplesOk

`func (o *ThroughputTile) GetWriteThroughputSamplesOk() (*[]Sample, bool)`

GetWriteThroughputSamplesOk returns a tuple with the WriteThroughputSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteThroughputSamples

`func (o *ThroughputTile) SetWriteThroughputSamples(v []Sample)`

SetWriteThroughputSamples sets WriteThroughputSamples field to given value.

### HasWriteThroughputSamples

`func (o *ThroughputTile) HasWriteThroughputSamples() bool`

HasWriteThroughputSamples returns a boolean if a field has been set.

### SetWriteThroughputSamplesNil

`func (o *ThroughputTile) SetWriteThroughputSamplesNil(b bool)`

 SetWriteThroughputSamplesNil sets the value for WriteThroughputSamples to be an explicit nil

### UnsetWriteThroughputSamples
`func (o *ThroughputTile) UnsetWriteThroughputSamples()`

UnsetWriteThroughputSamples ensures that no value is present for WriteThroughputSamples, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


