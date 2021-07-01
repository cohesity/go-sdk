# VaultRunStatsSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureTimeSeries** | Pointer to [**[]VaultRunInfo**](VaultRunInfo.md) | Specifies the time series for the failed runs that ended in the given time frame. | [optional] 
**NumFailedRuns** | Pointer to **NullableInt64** | Specifies the number of runs that ended in failure during the given time frame. | [optional] 
**NumInProgressRuns** | Pointer to **NullableInt64** | Specifies the number of runs that were currently in progress at the time that the API call was made. | [optional] 
**NumQueuedRuns** | Pointer to **NullableInt64** | Specifies the number of runs that were currently queued at the time that the API call was made. | [optional] 
**NumSuccessfulRuns** | Pointer to **NullableInt64** | Specifies the number of runs that ended in success during the given time frame. | [optional] 
**SuccessTimeSeries** | Pointer to [**[]VaultRunInfo**](VaultRunInfo.md) | Specifies the time series for the successful runs that ended in the given time frame. | [optional] 

## Methods

### NewVaultRunStatsSummary

`func NewVaultRunStatsSummary() *VaultRunStatsSummary`

NewVaultRunStatsSummary instantiates a new VaultRunStatsSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultRunStatsSummaryWithDefaults

`func NewVaultRunStatsSummaryWithDefaults() *VaultRunStatsSummary`

NewVaultRunStatsSummaryWithDefaults instantiates a new VaultRunStatsSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureTimeSeries

`func (o *VaultRunStatsSummary) GetFailureTimeSeries() []VaultRunInfo`

GetFailureTimeSeries returns the FailureTimeSeries field if non-nil, zero value otherwise.

### GetFailureTimeSeriesOk

`func (o *VaultRunStatsSummary) GetFailureTimeSeriesOk() (*[]VaultRunInfo, bool)`

GetFailureTimeSeriesOk returns a tuple with the FailureTimeSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureTimeSeries

`func (o *VaultRunStatsSummary) SetFailureTimeSeries(v []VaultRunInfo)`

SetFailureTimeSeries sets FailureTimeSeries field to given value.

### HasFailureTimeSeries

`func (o *VaultRunStatsSummary) HasFailureTimeSeries() bool`

HasFailureTimeSeries returns a boolean if a field has been set.

### GetNumFailedRuns

`func (o *VaultRunStatsSummary) GetNumFailedRuns() int64`

GetNumFailedRuns returns the NumFailedRuns field if non-nil, zero value otherwise.

### GetNumFailedRunsOk

`func (o *VaultRunStatsSummary) GetNumFailedRunsOk() (*int64, bool)`

GetNumFailedRunsOk returns a tuple with the NumFailedRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFailedRuns

`func (o *VaultRunStatsSummary) SetNumFailedRuns(v int64)`

SetNumFailedRuns sets NumFailedRuns field to given value.

### HasNumFailedRuns

`func (o *VaultRunStatsSummary) HasNumFailedRuns() bool`

HasNumFailedRuns returns a boolean if a field has been set.

### SetNumFailedRunsNil

`func (o *VaultRunStatsSummary) SetNumFailedRunsNil(b bool)`

 SetNumFailedRunsNil sets the value for NumFailedRuns to be an explicit nil

### UnsetNumFailedRuns
`func (o *VaultRunStatsSummary) UnsetNumFailedRuns()`

UnsetNumFailedRuns ensures that no value is present for NumFailedRuns, not even an explicit nil
### GetNumInProgressRuns

`func (o *VaultRunStatsSummary) GetNumInProgressRuns() int64`

GetNumInProgressRuns returns the NumInProgressRuns field if non-nil, zero value otherwise.

### GetNumInProgressRunsOk

`func (o *VaultRunStatsSummary) GetNumInProgressRunsOk() (*int64, bool)`

GetNumInProgressRunsOk returns a tuple with the NumInProgressRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInProgressRuns

`func (o *VaultRunStatsSummary) SetNumInProgressRuns(v int64)`

SetNumInProgressRuns sets NumInProgressRuns field to given value.

### HasNumInProgressRuns

`func (o *VaultRunStatsSummary) HasNumInProgressRuns() bool`

HasNumInProgressRuns returns a boolean if a field has been set.

### SetNumInProgressRunsNil

`func (o *VaultRunStatsSummary) SetNumInProgressRunsNil(b bool)`

 SetNumInProgressRunsNil sets the value for NumInProgressRuns to be an explicit nil

### UnsetNumInProgressRuns
`func (o *VaultRunStatsSummary) UnsetNumInProgressRuns()`

UnsetNumInProgressRuns ensures that no value is present for NumInProgressRuns, not even an explicit nil
### GetNumQueuedRuns

`func (o *VaultRunStatsSummary) GetNumQueuedRuns() int64`

GetNumQueuedRuns returns the NumQueuedRuns field if non-nil, zero value otherwise.

### GetNumQueuedRunsOk

`func (o *VaultRunStatsSummary) GetNumQueuedRunsOk() (*int64, bool)`

GetNumQueuedRunsOk returns a tuple with the NumQueuedRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumQueuedRuns

`func (o *VaultRunStatsSummary) SetNumQueuedRuns(v int64)`

SetNumQueuedRuns sets NumQueuedRuns field to given value.

### HasNumQueuedRuns

`func (o *VaultRunStatsSummary) HasNumQueuedRuns() bool`

HasNumQueuedRuns returns a boolean if a field has been set.

### SetNumQueuedRunsNil

`func (o *VaultRunStatsSummary) SetNumQueuedRunsNil(b bool)`

 SetNumQueuedRunsNil sets the value for NumQueuedRuns to be an explicit nil

### UnsetNumQueuedRuns
`func (o *VaultRunStatsSummary) UnsetNumQueuedRuns()`

UnsetNumQueuedRuns ensures that no value is present for NumQueuedRuns, not even an explicit nil
### GetNumSuccessfulRuns

`func (o *VaultRunStatsSummary) GetNumSuccessfulRuns() int64`

GetNumSuccessfulRuns returns the NumSuccessfulRuns field if non-nil, zero value otherwise.

### GetNumSuccessfulRunsOk

`func (o *VaultRunStatsSummary) GetNumSuccessfulRunsOk() (*int64, bool)`

GetNumSuccessfulRunsOk returns a tuple with the NumSuccessfulRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSuccessfulRuns

`func (o *VaultRunStatsSummary) SetNumSuccessfulRuns(v int64)`

SetNumSuccessfulRuns sets NumSuccessfulRuns field to given value.

### HasNumSuccessfulRuns

`func (o *VaultRunStatsSummary) HasNumSuccessfulRuns() bool`

HasNumSuccessfulRuns returns a boolean if a field has been set.

### SetNumSuccessfulRunsNil

`func (o *VaultRunStatsSummary) SetNumSuccessfulRunsNil(b bool)`

 SetNumSuccessfulRunsNil sets the value for NumSuccessfulRuns to be an explicit nil

### UnsetNumSuccessfulRuns
`func (o *VaultRunStatsSummary) UnsetNumSuccessfulRuns()`

UnsetNumSuccessfulRuns ensures that no value is present for NumSuccessfulRuns, not even an explicit nil
### GetSuccessTimeSeries

`func (o *VaultRunStatsSummary) GetSuccessTimeSeries() []VaultRunInfo`

GetSuccessTimeSeries returns the SuccessTimeSeries field if non-nil, zero value otherwise.

### GetSuccessTimeSeriesOk

`func (o *VaultRunStatsSummary) GetSuccessTimeSeriesOk() (*[]VaultRunInfo, bool)`

GetSuccessTimeSeriesOk returns a tuple with the SuccessTimeSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessTimeSeries

`func (o *VaultRunStatsSummary) SetSuccessTimeSeries(v []VaultRunInfo)`

SetSuccessTimeSeries sets SuccessTimeSeries field to given value.

### HasSuccessTimeSeries

`func (o *VaultRunStatsSummary) HasSuccessTimeSeries() bool`

HasSuccessTimeSeries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


