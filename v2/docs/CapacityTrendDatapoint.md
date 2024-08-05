# CapacityTrendDatapoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataPointStats** | Pointer to [**[]CapacityTrendDatapointStats**](CapacityTrendDatapointStats.md) | Specifies an array of tag and its corresponding statistic. | [optional] 
**DataPointTimestamp** | Pointer to **int64** | Specifies the timestamp of this data point. | [optional] 
**ErrorMessages** | Pointer to **[]string** | Specifies error messages, if any error is encountered while fetching the datapoint stats. | [optional] 

## Methods

### NewCapacityTrendDatapoint

`func NewCapacityTrendDatapoint() *CapacityTrendDatapoint`

NewCapacityTrendDatapoint instantiates a new CapacityTrendDatapoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapacityTrendDatapointWithDefaults

`func NewCapacityTrendDatapointWithDefaults() *CapacityTrendDatapoint`

NewCapacityTrendDatapointWithDefaults instantiates a new CapacityTrendDatapoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataPointStats

`func (o *CapacityTrendDatapoint) GetDataPointStats() []CapacityTrendDatapointStats`

GetDataPointStats returns the DataPointStats field if non-nil, zero value otherwise.

### GetDataPointStatsOk

`func (o *CapacityTrendDatapoint) GetDataPointStatsOk() (*[]CapacityTrendDatapointStats, bool)`

GetDataPointStatsOk returns a tuple with the DataPointStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPointStats

`func (o *CapacityTrendDatapoint) SetDataPointStats(v []CapacityTrendDatapointStats)`

SetDataPointStats sets DataPointStats field to given value.

### HasDataPointStats

`func (o *CapacityTrendDatapoint) HasDataPointStats() bool`

HasDataPointStats returns a boolean if a field has been set.

### SetDataPointStatsNil

`func (o *CapacityTrendDatapoint) SetDataPointStatsNil(b bool)`

 SetDataPointStatsNil sets the value for DataPointStats to be an explicit nil

### UnsetDataPointStats
`func (o *CapacityTrendDatapoint) UnsetDataPointStats()`

UnsetDataPointStats ensures that no value is present for DataPointStats, not even an explicit nil
### GetDataPointTimestamp

`func (o *CapacityTrendDatapoint) GetDataPointTimestamp() int64`

GetDataPointTimestamp returns the DataPointTimestamp field if non-nil, zero value otherwise.

### GetDataPointTimestampOk

`func (o *CapacityTrendDatapoint) GetDataPointTimestampOk() (*int64, bool)`

GetDataPointTimestampOk returns a tuple with the DataPointTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPointTimestamp

`func (o *CapacityTrendDatapoint) SetDataPointTimestamp(v int64)`

SetDataPointTimestamp sets DataPointTimestamp field to given value.

### HasDataPointTimestamp

`func (o *CapacityTrendDatapoint) HasDataPointTimestamp() bool`

HasDataPointTimestamp returns a boolean if a field has been set.

### GetErrorMessages

`func (o *CapacityTrendDatapoint) GetErrorMessages() []string`

GetErrorMessages returns the ErrorMessages field if non-nil, zero value otherwise.

### GetErrorMessagesOk

`func (o *CapacityTrendDatapoint) GetErrorMessagesOk() (*[]string, bool)`

GetErrorMessagesOk returns a tuple with the ErrorMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessages

`func (o *CapacityTrendDatapoint) SetErrorMessages(v []string)`

SetErrorMessages sets ErrorMessages field to given value.

### HasErrorMessages

`func (o *CapacityTrendDatapoint) HasErrorMessages() bool`

HasErrorMessages returns a boolean if a field has been set.

### SetErrorMessagesNil

`func (o *CapacityTrendDatapoint) SetErrorMessagesNil(b bool)`

 SetErrorMessagesNil sets the value for ErrorMessages to be an explicit nil

### UnsetErrorMessages
`func (o *CapacityTrendDatapoint) UnsetErrorMessages()`

UnsetErrorMessages ensures that no value is present for ErrorMessages, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


