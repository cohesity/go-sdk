# DataTieringObjectAnalysisInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **NullableString** | A message about the error if encountered while performing data tiering analysis. | [optional] 
**Stats** | Pointer to [**[]DataTieringShareStats**](DataTieringShareStats.md) | Specifies the source share analysis stats. | [optional] 
**Status** | Pointer to **NullableString** | Status of the analysis run. &#39;Running&#39; indicates that the run is still running. &#39;Canceled&#39; indicates that the run has been canceled. &#39;Canceling&#39; indicates that the run is in the process of being  canceled. &#39;Failed&#39; indicates that the run has failed. &#39;Missed&#39; indicates that the run was unable to take place at the  scheduled time because the previous run was still happening. &#39;Succeeded&#39; indicates that the run has finished successfully. &#39;SucceededWithWarning&#39; indicates that the run finished  successfully, but there were some warning messages. &#39;OnHold&#39; indicates that the run has On hold. &#39;Skipped&#39; indicates that the run was skipped. | [optional] 
**TagsInfo** | Pointer to [**[]DataTieringTagObject**](DataTieringTagObject.md) | Array of Tag objects. | [optional] 

## Methods

### NewDataTieringObjectAnalysisInfo

`func NewDataTieringObjectAnalysisInfo() *DataTieringObjectAnalysisInfo`

NewDataTieringObjectAnalysisInfo instantiates a new DataTieringObjectAnalysisInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTieringObjectAnalysisInfoWithDefaults

`func NewDataTieringObjectAnalysisInfoWithDefaults() *DataTieringObjectAnalysisInfo`

NewDataTieringObjectAnalysisInfoWithDefaults instantiates a new DataTieringObjectAnalysisInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *DataTieringObjectAnalysisInfo) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DataTieringObjectAnalysisInfo) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DataTieringObjectAnalysisInfo) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DataTieringObjectAnalysisInfo) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *DataTieringObjectAnalysisInfo) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *DataTieringObjectAnalysisInfo) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetStats

`func (o *DataTieringObjectAnalysisInfo) GetStats() []DataTieringShareStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *DataTieringObjectAnalysisInfo) GetStatsOk() (*[]DataTieringShareStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *DataTieringObjectAnalysisInfo) SetStats(v []DataTieringShareStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *DataTieringObjectAnalysisInfo) HasStats() bool`

HasStats returns a boolean if a field has been set.

### SetStatsNil

`func (o *DataTieringObjectAnalysisInfo) SetStatsNil(b bool)`

 SetStatsNil sets the value for Stats to be an explicit nil

### UnsetStats
`func (o *DataTieringObjectAnalysisInfo) UnsetStats()`

UnsetStats ensures that no value is present for Stats, not even an explicit nil
### GetStatus

`func (o *DataTieringObjectAnalysisInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DataTieringObjectAnalysisInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DataTieringObjectAnalysisInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DataTieringObjectAnalysisInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *DataTieringObjectAnalysisInfo) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *DataTieringObjectAnalysisInfo) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTagsInfo

`func (o *DataTieringObjectAnalysisInfo) GetTagsInfo() []DataTieringTagObject`

GetTagsInfo returns the TagsInfo field if non-nil, zero value otherwise.

### GetTagsInfoOk

`func (o *DataTieringObjectAnalysisInfo) GetTagsInfoOk() (*[]DataTieringTagObject, bool)`

GetTagsInfoOk returns a tuple with the TagsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsInfo

`func (o *DataTieringObjectAnalysisInfo) SetTagsInfo(v []DataTieringTagObject)`

SetTagsInfo sets TagsInfo field to given value.

### HasTagsInfo

`func (o *DataTieringObjectAnalysisInfo) HasTagsInfo() bool`

HasTagsInfo returns a boolean if a field has been set.

### SetTagsInfoNil

`func (o *DataTieringObjectAnalysisInfo) SetTagsInfoNil(b bool)`

 SetTagsInfoNil sets the value for TagsInfo to be an explicit nil

### UnsetTagsInfo
`func (o *DataTieringObjectAnalysisInfo) UnsetTagsInfo()`

UnsetTagsInfo ensures that no value is present for TagsInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


