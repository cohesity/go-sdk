# DataTieringAnalysisInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagsInfo** | Pointer to [**[]DataTieringTagObject**](DataTieringTagObject.md) | Array of Tag objects. | [optional] 
**Progress** | Pointer to [**ProgressSummary**](ProgressSummary.md) |  | [optional] 
**Status** | Pointer to **NullableString** | Status of the analysis run. &#39;Running&#39; indicates that the run is still running. &#39;Canceled&#39; indicates that the run has been canceled. &#39;Canceling&#39; indicates that the run is in the process of being  canceled. &#39;Failed&#39; indicates that the run has failed. &#39;Missed&#39; indicates that the run was unable to take place at the  scheduled time because the previous run was still happening. &#39;Succeeded&#39; indicates that the run has finished successfully. &#39;SucceededWithWarning&#39; indicates that the run finished  successfully, but there were some warning messages. &#39;OnHold&#39; indicates that the run has On hold. | [optional] 
**Message** | Pointer to **NullableString** | message from the last analysis run stating the error message in case of error in analysis run or warning message if the run finished successfully, but there were some warning messages. | [optional] 

## Methods

### NewDataTieringAnalysisInfo

`func NewDataTieringAnalysisInfo() *DataTieringAnalysisInfo`

NewDataTieringAnalysisInfo instantiates a new DataTieringAnalysisInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTieringAnalysisInfoWithDefaults

`func NewDataTieringAnalysisInfoWithDefaults() *DataTieringAnalysisInfo`

NewDataTieringAnalysisInfoWithDefaults instantiates a new DataTieringAnalysisInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagsInfo

`func (o *DataTieringAnalysisInfo) GetTagsInfo() []DataTieringTagObject`

GetTagsInfo returns the TagsInfo field if non-nil, zero value otherwise.

### GetTagsInfoOk

`func (o *DataTieringAnalysisInfo) GetTagsInfoOk() (*[]DataTieringTagObject, bool)`

GetTagsInfoOk returns a tuple with the TagsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsInfo

`func (o *DataTieringAnalysisInfo) SetTagsInfo(v []DataTieringTagObject)`

SetTagsInfo sets TagsInfo field to given value.

### HasTagsInfo

`func (o *DataTieringAnalysisInfo) HasTagsInfo() bool`

HasTagsInfo returns a boolean if a field has been set.

### SetTagsInfoNil

`func (o *DataTieringAnalysisInfo) SetTagsInfoNil(b bool)`

 SetTagsInfoNil sets the value for TagsInfo to be an explicit nil

### UnsetTagsInfo
`func (o *DataTieringAnalysisInfo) UnsetTagsInfo()`

UnsetTagsInfo ensures that no value is present for TagsInfo, not even an explicit nil
### GetProgress

`func (o *DataTieringAnalysisInfo) GetProgress() ProgressSummary`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *DataTieringAnalysisInfo) GetProgressOk() (*ProgressSummary, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *DataTieringAnalysisInfo) SetProgress(v ProgressSummary)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *DataTieringAnalysisInfo) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetStatus

`func (o *DataTieringAnalysisInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DataTieringAnalysisInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DataTieringAnalysisInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DataTieringAnalysisInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *DataTieringAnalysisInfo) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *DataTieringAnalysisInfo) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetMessage

`func (o *DataTieringAnalysisInfo) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DataTieringAnalysisInfo) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DataTieringAnalysisInfo) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DataTieringAnalysisInfo) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *DataTieringAnalysisInfo) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *DataTieringAnalysisInfo) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


