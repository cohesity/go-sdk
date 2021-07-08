# IopsTile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxReadIops** | Pointer to **NullableInt64** | Maximum Read IOs per second in last 24 hours. | [optional] 
**MaxWriteIops** | Pointer to **NullableInt64** | Maximum number of Write IOs per second in last 24 hours. | [optional] 
**ReadIopsSamples** | Pointer to [**[]Sample**](Sample.md) | Read IOs per second samples taken for the past 24 hours at 10 minutes interval given in descending order of time. | [optional] 
**WriteIopsSamples** | Pointer to [**[]Sample**](Sample.md) | Write IOs per second samples taken for the past 24 hours at 10 minutes interval given in descending order of time. | [optional] 

## Methods

### NewIopsTile

`func NewIopsTile() *IopsTile`

NewIopsTile instantiates a new IopsTile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIopsTileWithDefaults

`func NewIopsTileWithDefaults() *IopsTile`

NewIopsTileWithDefaults instantiates a new IopsTile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxReadIops

`func (o *IopsTile) GetMaxReadIops() int64`

GetMaxReadIops returns the MaxReadIops field if non-nil, zero value otherwise.

### GetMaxReadIopsOk

`func (o *IopsTile) GetMaxReadIopsOk() (*int64, bool)`

GetMaxReadIopsOk returns a tuple with the MaxReadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReadIops

`func (o *IopsTile) SetMaxReadIops(v int64)`

SetMaxReadIops sets MaxReadIops field to given value.

### HasMaxReadIops

`func (o *IopsTile) HasMaxReadIops() bool`

HasMaxReadIops returns a boolean if a field has been set.

### SetMaxReadIopsNil

`func (o *IopsTile) SetMaxReadIopsNil(b bool)`

 SetMaxReadIopsNil sets the value for MaxReadIops to be an explicit nil

### UnsetMaxReadIops
`func (o *IopsTile) UnsetMaxReadIops()`

UnsetMaxReadIops ensures that no value is present for MaxReadIops, not even an explicit nil
### GetMaxWriteIops

`func (o *IopsTile) GetMaxWriteIops() int64`

GetMaxWriteIops returns the MaxWriteIops field if non-nil, zero value otherwise.

### GetMaxWriteIopsOk

`func (o *IopsTile) GetMaxWriteIopsOk() (*int64, bool)`

GetMaxWriteIopsOk returns a tuple with the MaxWriteIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWriteIops

`func (o *IopsTile) SetMaxWriteIops(v int64)`

SetMaxWriteIops sets MaxWriteIops field to given value.

### HasMaxWriteIops

`func (o *IopsTile) HasMaxWriteIops() bool`

HasMaxWriteIops returns a boolean if a field has been set.

### SetMaxWriteIopsNil

`func (o *IopsTile) SetMaxWriteIopsNil(b bool)`

 SetMaxWriteIopsNil sets the value for MaxWriteIops to be an explicit nil

### UnsetMaxWriteIops
`func (o *IopsTile) UnsetMaxWriteIops()`

UnsetMaxWriteIops ensures that no value is present for MaxWriteIops, not even an explicit nil
### GetReadIopsSamples

`func (o *IopsTile) GetReadIopsSamples() []Sample`

GetReadIopsSamples returns the ReadIopsSamples field if non-nil, zero value otherwise.

### GetReadIopsSamplesOk

`func (o *IopsTile) GetReadIopsSamplesOk() (*[]Sample, bool)`

GetReadIopsSamplesOk returns a tuple with the ReadIopsSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadIopsSamples

`func (o *IopsTile) SetReadIopsSamples(v []Sample)`

SetReadIopsSamples sets ReadIopsSamples field to given value.

### HasReadIopsSamples

`func (o *IopsTile) HasReadIopsSamples() bool`

HasReadIopsSamples returns a boolean if a field has been set.

### SetReadIopsSamplesNil

`func (o *IopsTile) SetReadIopsSamplesNil(b bool)`

 SetReadIopsSamplesNil sets the value for ReadIopsSamples to be an explicit nil

### UnsetReadIopsSamples
`func (o *IopsTile) UnsetReadIopsSamples()`

UnsetReadIopsSamples ensures that no value is present for ReadIopsSamples, not even an explicit nil
### GetWriteIopsSamples

`func (o *IopsTile) GetWriteIopsSamples() []Sample`

GetWriteIopsSamples returns the WriteIopsSamples field if non-nil, zero value otherwise.

### GetWriteIopsSamplesOk

`func (o *IopsTile) GetWriteIopsSamplesOk() (*[]Sample, bool)`

GetWriteIopsSamplesOk returns a tuple with the WriteIopsSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteIopsSamples

`func (o *IopsTile) SetWriteIopsSamples(v []Sample)`

SetWriteIopsSamples sets WriteIopsSamples field to given value.

### HasWriteIopsSamples

`func (o *IopsTile) HasWriteIopsSamples() bool`

HasWriteIopsSamples returns a boolean if a field has been set.

### SetWriteIopsSamplesNil

`func (o *IopsTile) SetWriteIopsSamplesNil(b bool)`

 SetWriteIopsSamplesNil sets the value for WriteIopsSamples to be an explicit nil

### UnsetWriteIopsSamples
`func (o *IopsTile) UnsetWriteIopsSamples()`

UnsetWriteIopsSamples ensures that no value is present for WriteIopsSamples, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


