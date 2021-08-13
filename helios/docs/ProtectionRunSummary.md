# ProtectionRunSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specifies the ID of the Protection Group run. | [optional] 
**ProtectionGroupId** | Pointer to **NullableString** | ProtectionGroupId to which this run belongs. | [optional] 
**ProtectionGroupName** | Pointer to **NullableString** | Name of the Protection Group to which this run belongs. | [optional] 
**Environment** | Pointer to **NullableString** | Specifies the environment type of the Protection Group. | [optional] 
**IsSlaViolated** | Pointer to **NullableBool** | Indicated if SLA has been violated for this run. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of backup run in Unix epoch Timestamp(in microseconds). | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of backup run in Unix epoch Timestamp(in microseconds). | [optional] 
**Status** | Pointer to **NullableString** | Status of the backup run. &#39;Running&#39; indicates that the run is still running. &#39;Canceled&#39; indicates that the run has been canceled. &#39;Canceling&#39; indicates that the run is in the process of being canceled. &#39;Paused&#39; indicates that the ongoing run has been paused. &#39;Failed&#39; indicates that the run has failed. &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening. &#39;Succeeded&#39; indicates that the run has finished successfully. &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages. | [optional] 
**IsFullRun** | Pointer to **NullableBool** | Specifies if the protection run is a full run. | [optional] 
**TotalObjectsCount** | Pointer to **NullableInt64** | Specifies the total number of objects protected in this run. | [optional] 
**SuccessObjectsCount** | Pointer to **NullableInt64** | Specifies the number of objects which are successfully protected in this run. | [optional] 
**LogicalSizeBytes** | Pointer to **NullableInt64** | Specifies total logical size of object(s) in bytes. | [optional] 
**BytesWritten** | Pointer to **NullableInt64** | Specifies total size of data in bytes written after taking backup. | [optional] 

## Methods

### NewProtectionRunSummary

`func NewProtectionRunSummary() *ProtectionRunSummary`

NewProtectionRunSummary instantiates a new ProtectionRunSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionRunSummaryWithDefaults

`func NewProtectionRunSummaryWithDefaults() *ProtectionRunSummary`

NewProtectionRunSummaryWithDefaults instantiates a new ProtectionRunSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProtectionRunSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProtectionRunSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProtectionRunSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProtectionRunSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProtectionRunSummary) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProtectionRunSummary) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetProtectionGroupId

`func (o *ProtectionRunSummary) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *ProtectionRunSummary) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *ProtectionRunSummary) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *ProtectionRunSummary) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *ProtectionRunSummary) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *ProtectionRunSummary) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetProtectionGroupName

`func (o *ProtectionRunSummary) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *ProtectionRunSummary) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *ProtectionRunSummary) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *ProtectionRunSummary) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### SetProtectionGroupNameNil

`func (o *ProtectionRunSummary) SetProtectionGroupNameNil(b bool)`

 SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil

### UnsetProtectionGroupName
`func (o *ProtectionRunSummary) UnsetProtectionGroupName()`

UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
### GetEnvironment

`func (o *ProtectionRunSummary) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ProtectionRunSummary) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ProtectionRunSummary) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ProtectionRunSummary) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *ProtectionRunSummary) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ProtectionRunSummary) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetIsSlaViolated

`func (o *ProtectionRunSummary) GetIsSlaViolated() bool`

GetIsSlaViolated returns the IsSlaViolated field if non-nil, zero value otherwise.

### GetIsSlaViolatedOk

`func (o *ProtectionRunSummary) GetIsSlaViolatedOk() (*bool, bool)`

GetIsSlaViolatedOk returns a tuple with the IsSlaViolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSlaViolated

`func (o *ProtectionRunSummary) SetIsSlaViolated(v bool)`

SetIsSlaViolated sets IsSlaViolated field to given value.

### HasIsSlaViolated

`func (o *ProtectionRunSummary) HasIsSlaViolated() bool`

HasIsSlaViolated returns a boolean if a field has been set.

### SetIsSlaViolatedNil

`func (o *ProtectionRunSummary) SetIsSlaViolatedNil(b bool)`

 SetIsSlaViolatedNil sets the value for IsSlaViolated to be an explicit nil

### UnsetIsSlaViolated
`func (o *ProtectionRunSummary) UnsetIsSlaViolated()`

UnsetIsSlaViolated ensures that no value is present for IsSlaViolated, not even an explicit nil
### GetStartTimeUsecs

`func (o *ProtectionRunSummary) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *ProtectionRunSummary) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *ProtectionRunSummary) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *ProtectionRunSummary) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *ProtectionRunSummary) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *ProtectionRunSummary) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetEndTimeUsecs

`func (o *ProtectionRunSummary) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *ProtectionRunSummary) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *ProtectionRunSummary) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *ProtectionRunSummary) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *ProtectionRunSummary) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *ProtectionRunSummary) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetStatus

`func (o *ProtectionRunSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProtectionRunSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProtectionRunSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProtectionRunSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ProtectionRunSummary) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ProtectionRunSummary) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetIsFullRun

`func (o *ProtectionRunSummary) GetIsFullRun() bool`

GetIsFullRun returns the IsFullRun field if non-nil, zero value otherwise.

### GetIsFullRunOk

`func (o *ProtectionRunSummary) GetIsFullRunOk() (*bool, bool)`

GetIsFullRunOk returns a tuple with the IsFullRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullRun

`func (o *ProtectionRunSummary) SetIsFullRun(v bool)`

SetIsFullRun sets IsFullRun field to given value.

### HasIsFullRun

`func (o *ProtectionRunSummary) HasIsFullRun() bool`

HasIsFullRun returns a boolean if a field has been set.

### SetIsFullRunNil

`func (o *ProtectionRunSummary) SetIsFullRunNil(b bool)`

 SetIsFullRunNil sets the value for IsFullRun to be an explicit nil

### UnsetIsFullRun
`func (o *ProtectionRunSummary) UnsetIsFullRun()`

UnsetIsFullRun ensures that no value is present for IsFullRun, not even an explicit nil
### GetTotalObjectsCount

`func (o *ProtectionRunSummary) GetTotalObjectsCount() int64`

GetTotalObjectsCount returns the TotalObjectsCount field if non-nil, zero value otherwise.

### GetTotalObjectsCountOk

`func (o *ProtectionRunSummary) GetTotalObjectsCountOk() (*int64, bool)`

GetTotalObjectsCountOk returns a tuple with the TotalObjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalObjectsCount

`func (o *ProtectionRunSummary) SetTotalObjectsCount(v int64)`

SetTotalObjectsCount sets TotalObjectsCount field to given value.

### HasTotalObjectsCount

`func (o *ProtectionRunSummary) HasTotalObjectsCount() bool`

HasTotalObjectsCount returns a boolean if a field has been set.

### SetTotalObjectsCountNil

`func (o *ProtectionRunSummary) SetTotalObjectsCountNil(b bool)`

 SetTotalObjectsCountNil sets the value for TotalObjectsCount to be an explicit nil

### UnsetTotalObjectsCount
`func (o *ProtectionRunSummary) UnsetTotalObjectsCount()`

UnsetTotalObjectsCount ensures that no value is present for TotalObjectsCount, not even an explicit nil
### GetSuccessObjectsCount

`func (o *ProtectionRunSummary) GetSuccessObjectsCount() int64`

GetSuccessObjectsCount returns the SuccessObjectsCount field if non-nil, zero value otherwise.

### GetSuccessObjectsCountOk

`func (o *ProtectionRunSummary) GetSuccessObjectsCountOk() (*int64, bool)`

GetSuccessObjectsCountOk returns a tuple with the SuccessObjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessObjectsCount

`func (o *ProtectionRunSummary) SetSuccessObjectsCount(v int64)`

SetSuccessObjectsCount sets SuccessObjectsCount field to given value.

### HasSuccessObjectsCount

`func (o *ProtectionRunSummary) HasSuccessObjectsCount() bool`

HasSuccessObjectsCount returns a boolean if a field has been set.

### SetSuccessObjectsCountNil

`func (o *ProtectionRunSummary) SetSuccessObjectsCountNil(b bool)`

 SetSuccessObjectsCountNil sets the value for SuccessObjectsCount to be an explicit nil

### UnsetSuccessObjectsCount
`func (o *ProtectionRunSummary) UnsetSuccessObjectsCount()`

UnsetSuccessObjectsCount ensures that no value is present for SuccessObjectsCount, not even an explicit nil
### GetLogicalSizeBytes

`func (o *ProtectionRunSummary) GetLogicalSizeBytes() int64`

GetLogicalSizeBytes returns the LogicalSizeBytes field if non-nil, zero value otherwise.

### GetLogicalSizeBytesOk

`func (o *ProtectionRunSummary) GetLogicalSizeBytesOk() (*int64, bool)`

GetLogicalSizeBytesOk returns a tuple with the LogicalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSizeBytes

`func (o *ProtectionRunSummary) SetLogicalSizeBytes(v int64)`

SetLogicalSizeBytes sets LogicalSizeBytes field to given value.

### HasLogicalSizeBytes

`func (o *ProtectionRunSummary) HasLogicalSizeBytes() bool`

HasLogicalSizeBytes returns a boolean if a field has been set.

### SetLogicalSizeBytesNil

`func (o *ProtectionRunSummary) SetLogicalSizeBytesNil(b bool)`

 SetLogicalSizeBytesNil sets the value for LogicalSizeBytes to be an explicit nil

### UnsetLogicalSizeBytes
`func (o *ProtectionRunSummary) UnsetLogicalSizeBytes()`

UnsetLogicalSizeBytes ensures that no value is present for LogicalSizeBytes, not even an explicit nil
### GetBytesWritten

`func (o *ProtectionRunSummary) GetBytesWritten() int64`

GetBytesWritten returns the BytesWritten field if non-nil, zero value otherwise.

### GetBytesWrittenOk

`func (o *ProtectionRunSummary) GetBytesWrittenOk() (*int64, bool)`

GetBytesWrittenOk returns a tuple with the BytesWritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesWritten

`func (o *ProtectionRunSummary) SetBytesWritten(v int64)`

SetBytesWritten sets BytesWritten field to given value.

### HasBytesWritten

`func (o *ProtectionRunSummary) HasBytesWritten() bool`

HasBytesWritten returns a boolean if a field has been set.

### SetBytesWrittenNil

`func (o *ProtectionRunSummary) SetBytesWrittenNil(b bool)`

 SetBytesWrittenNil sets the value for BytesWritten to be an explicit nil

### UnsetBytesWritten
`func (o *ProtectionRunSummary) UnsetBytesWritten()`

UnsetBytesWritten ensures that no value is present for BytesWritten, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


