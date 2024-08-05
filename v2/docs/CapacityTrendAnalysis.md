# CapacityTrendAnalysis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataPointVec** | Pointer to [**[]CapacityTrendDatapoint**](CapacityTrendDatapoint.md) | Vector of data points. Each data point contains statistics on analysis and downtiered data. | [optional] 
**LastUpdatedTimeUsecs** | Pointer to **NullableInt64** | Last updated time of capacity trend. | [optional] 

## Methods

### NewCapacityTrendAnalysis

`func NewCapacityTrendAnalysis() *CapacityTrendAnalysis`

NewCapacityTrendAnalysis instantiates a new CapacityTrendAnalysis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapacityTrendAnalysisWithDefaults

`func NewCapacityTrendAnalysisWithDefaults() *CapacityTrendAnalysis`

NewCapacityTrendAnalysisWithDefaults instantiates a new CapacityTrendAnalysis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataPointVec

`func (o *CapacityTrendAnalysis) GetDataPointVec() []CapacityTrendDatapoint`

GetDataPointVec returns the DataPointVec field if non-nil, zero value otherwise.

### GetDataPointVecOk

`func (o *CapacityTrendAnalysis) GetDataPointVecOk() (*[]CapacityTrendDatapoint, bool)`

GetDataPointVecOk returns a tuple with the DataPointVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPointVec

`func (o *CapacityTrendAnalysis) SetDataPointVec(v []CapacityTrendDatapoint)`

SetDataPointVec sets DataPointVec field to given value.

### HasDataPointVec

`func (o *CapacityTrendAnalysis) HasDataPointVec() bool`

HasDataPointVec returns a boolean if a field has been set.

### SetDataPointVecNil

`func (o *CapacityTrendAnalysis) SetDataPointVecNil(b bool)`

 SetDataPointVecNil sets the value for DataPointVec to be an explicit nil

### UnsetDataPointVec
`func (o *CapacityTrendAnalysis) UnsetDataPointVec()`

UnsetDataPointVec ensures that no value is present for DataPointVec, not even an explicit nil
### GetLastUpdatedTimeUsecs

`func (o *CapacityTrendAnalysis) GetLastUpdatedTimeUsecs() int64`

GetLastUpdatedTimeUsecs returns the LastUpdatedTimeUsecs field if non-nil, zero value otherwise.

### GetLastUpdatedTimeUsecsOk

`func (o *CapacityTrendAnalysis) GetLastUpdatedTimeUsecsOk() (*int64, bool)`

GetLastUpdatedTimeUsecsOk returns a tuple with the LastUpdatedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimeUsecs

`func (o *CapacityTrendAnalysis) SetLastUpdatedTimeUsecs(v int64)`

SetLastUpdatedTimeUsecs sets LastUpdatedTimeUsecs field to given value.

### HasLastUpdatedTimeUsecs

`func (o *CapacityTrendAnalysis) HasLastUpdatedTimeUsecs() bool`

HasLastUpdatedTimeUsecs returns a boolean if a field has been set.

### SetLastUpdatedTimeUsecsNil

`func (o *CapacityTrendAnalysis) SetLastUpdatedTimeUsecsNil(b bool)`

 SetLastUpdatedTimeUsecsNil sets the value for LastUpdatedTimeUsecs to be an explicit nil

### UnsetLastUpdatedTimeUsecs
`func (o *CapacityTrendAnalysis) UnsetLastUpdatedTimeUsecs()`

UnsetLastUpdatedTimeUsecs ensures that no value is present for LastUpdatedTimeUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


