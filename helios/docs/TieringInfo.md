# TieringInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Progress** | Pointer to [**ProgressSummary**](ProgressSummary.md) |  | [optional] 
**Status** | Pointer to **NullableString** | Status of the analysis run. &#39;Running&#39; indicates that the run is still running. &#39;Canceled&#39; indicates that the run has been canceled. &#39;Canceling&#39; indicates that the run is in the process of being  canceled. &#39;Failed&#39; indicates that the run has failed. &#39;Missed&#39; indicates that the run was unable to take place at the  scheduled time because the previous run was still happening. &#39;Succeeded&#39; indicates that the run has finished successfully. &#39;SucceededWithWarning&#39; indicates that the run finished  successfully, but there were some warning messages. &#39;OnHold&#39; indicates that the run has On hold. | [optional] 
**Stats** | Pointer to [**DataTieringTaskStats**](DataTieringTaskStats.md) |  | [optional] 

## Methods

### NewTieringInfo

`func NewTieringInfo() *TieringInfo`

NewTieringInfo instantiates a new TieringInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTieringInfoWithDefaults

`func NewTieringInfoWithDefaults() *TieringInfo`

NewTieringInfoWithDefaults instantiates a new TieringInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgress

`func (o *TieringInfo) GetProgress() ProgressSummary`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *TieringInfo) GetProgressOk() (*ProgressSummary, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *TieringInfo) SetProgress(v ProgressSummary)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *TieringInfo) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetStatus

`func (o *TieringInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TieringInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TieringInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TieringInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TieringInfo) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TieringInfo) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStats

`func (o *TieringInfo) GetStats() DataTieringTaskStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *TieringInfo) GetStatsOk() (*DataTieringTaskStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *TieringInfo) SetStats(v DataTieringTaskStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *TieringInfo) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


