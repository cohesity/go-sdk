# CloudSpinTargetResultAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of Cloud Spin in Unix epoch Timestamp(in microseconds) for a target. | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of Cloud Spin in Unix epoch Timestamp(in microseconds) for a target. | [optional] 
**Status** | Pointer to **NullableString** | Status of the Cloud Spin for a target. &#39;Running&#39; indicates that the run is still running. &#39;Canceled&#39; indicates that the run has been canceled. &#39;Canceling&#39; indicates that the run is in the process of being canceled. &#39;Paused&#39; indicates that the ongoing run has been paused. &#39;Failed&#39; indicates that the run has failed. &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening. &#39;Succeeded&#39; indicates that the run has finished successfully. &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages. | [optional] 
**Message** | Pointer to **NullableString** | Message about the Cloud Spin run. | [optional] 
**Stats** | Pointer to [**CloudSpinDataStats**](CloudSpinDataStats.md) |  | [optional] 
**IsManuallyDeleted** | Pointer to **NullableBool** | Specifies whether the snapshot is deleted manually. | [optional] 
**ExpiryTimeUsecs** | Pointer to **NullableInt64** | Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds) for an object. | [optional] 
**CloudspinTaskId** | Pointer to **NullableString** | Task ID for a CloudSpin protection run. | [optional] 
**ProgressTaskId** | Pointer to **NullableString** | Progress monitor task id for Cloud Spin run. | [optional] 
**DataLockConstraints** | Pointer to [**DataLockConstraints**](DataLockConstraints.md) |  | [optional] 
**OnLegalHold** | Pointer to **NullableBool** | Specifies the legal hold status for a cloud spin target. | [optional] 

## Methods

### NewCloudSpinTargetResultAllOf

`func NewCloudSpinTargetResultAllOf() *CloudSpinTargetResultAllOf`

NewCloudSpinTargetResultAllOf instantiates a new CloudSpinTargetResultAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSpinTargetResultAllOfWithDefaults

`func NewCloudSpinTargetResultAllOfWithDefaults() *CloudSpinTargetResultAllOf`

NewCloudSpinTargetResultAllOfWithDefaults instantiates a new CloudSpinTargetResultAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTimeUsecs

`func (o *CloudSpinTargetResultAllOf) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *CloudSpinTargetResultAllOf) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *CloudSpinTargetResultAllOf) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *CloudSpinTargetResultAllOf) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *CloudSpinTargetResultAllOf) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *CloudSpinTargetResultAllOf) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetEndTimeUsecs

`func (o *CloudSpinTargetResultAllOf) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *CloudSpinTargetResultAllOf) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *CloudSpinTargetResultAllOf) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *CloudSpinTargetResultAllOf) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *CloudSpinTargetResultAllOf) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *CloudSpinTargetResultAllOf) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetStatus

`func (o *CloudSpinTargetResultAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudSpinTargetResultAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudSpinTargetResultAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudSpinTargetResultAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CloudSpinTargetResultAllOf) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CloudSpinTargetResultAllOf) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetMessage

`func (o *CloudSpinTargetResultAllOf) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CloudSpinTargetResultAllOf) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CloudSpinTargetResultAllOf) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CloudSpinTargetResultAllOf) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CloudSpinTargetResultAllOf) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CloudSpinTargetResultAllOf) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetStats

`func (o *CloudSpinTargetResultAllOf) GetStats() CloudSpinDataStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *CloudSpinTargetResultAllOf) GetStatsOk() (*CloudSpinDataStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *CloudSpinTargetResultAllOf) SetStats(v CloudSpinDataStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *CloudSpinTargetResultAllOf) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetIsManuallyDeleted

`func (o *CloudSpinTargetResultAllOf) GetIsManuallyDeleted() bool`

GetIsManuallyDeleted returns the IsManuallyDeleted field if non-nil, zero value otherwise.

### GetIsManuallyDeletedOk

`func (o *CloudSpinTargetResultAllOf) GetIsManuallyDeletedOk() (*bool, bool)`

GetIsManuallyDeletedOk returns a tuple with the IsManuallyDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManuallyDeleted

`func (o *CloudSpinTargetResultAllOf) SetIsManuallyDeleted(v bool)`

SetIsManuallyDeleted sets IsManuallyDeleted field to given value.

### HasIsManuallyDeleted

`func (o *CloudSpinTargetResultAllOf) HasIsManuallyDeleted() bool`

HasIsManuallyDeleted returns a boolean if a field has been set.

### SetIsManuallyDeletedNil

`func (o *CloudSpinTargetResultAllOf) SetIsManuallyDeletedNil(b bool)`

 SetIsManuallyDeletedNil sets the value for IsManuallyDeleted to be an explicit nil

### UnsetIsManuallyDeleted
`func (o *CloudSpinTargetResultAllOf) UnsetIsManuallyDeleted()`

UnsetIsManuallyDeleted ensures that no value is present for IsManuallyDeleted, not even an explicit nil
### GetExpiryTimeUsecs

`func (o *CloudSpinTargetResultAllOf) GetExpiryTimeUsecs() int64`

GetExpiryTimeUsecs returns the ExpiryTimeUsecs field if non-nil, zero value otherwise.

### GetExpiryTimeUsecsOk

`func (o *CloudSpinTargetResultAllOf) GetExpiryTimeUsecsOk() (*int64, bool)`

GetExpiryTimeUsecsOk returns a tuple with the ExpiryTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeUsecs

`func (o *CloudSpinTargetResultAllOf) SetExpiryTimeUsecs(v int64)`

SetExpiryTimeUsecs sets ExpiryTimeUsecs field to given value.

### HasExpiryTimeUsecs

`func (o *CloudSpinTargetResultAllOf) HasExpiryTimeUsecs() bool`

HasExpiryTimeUsecs returns a boolean if a field has been set.

### SetExpiryTimeUsecsNil

`func (o *CloudSpinTargetResultAllOf) SetExpiryTimeUsecsNil(b bool)`

 SetExpiryTimeUsecsNil sets the value for ExpiryTimeUsecs to be an explicit nil

### UnsetExpiryTimeUsecs
`func (o *CloudSpinTargetResultAllOf) UnsetExpiryTimeUsecs()`

UnsetExpiryTimeUsecs ensures that no value is present for ExpiryTimeUsecs, not even an explicit nil
### GetCloudspinTaskId

`func (o *CloudSpinTargetResultAllOf) GetCloudspinTaskId() string`

GetCloudspinTaskId returns the CloudspinTaskId field if non-nil, zero value otherwise.

### GetCloudspinTaskIdOk

`func (o *CloudSpinTargetResultAllOf) GetCloudspinTaskIdOk() (*string, bool)`

GetCloudspinTaskIdOk returns a tuple with the CloudspinTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudspinTaskId

`func (o *CloudSpinTargetResultAllOf) SetCloudspinTaskId(v string)`

SetCloudspinTaskId sets CloudspinTaskId field to given value.

### HasCloudspinTaskId

`func (o *CloudSpinTargetResultAllOf) HasCloudspinTaskId() bool`

HasCloudspinTaskId returns a boolean if a field has been set.

### SetCloudspinTaskIdNil

`func (o *CloudSpinTargetResultAllOf) SetCloudspinTaskIdNil(b bool)`

 SetCloudspinTaskIdNil sets the value for CloudspinTaskId to be an explicit nil

### UnsetCloudspinTaskId
`func (o *CloudSpinTargetResultAllOf) UnsetCloudspinTaskId()`

UnsetCloudspinTaskId ensures that no value is present for CloudspinTaskId, not even an explicit nil
### GetProgressTaskId

`func (o *CloudSpinTargetResultAllOf) GetProgressTaskId() string`

GetProgressTaskId returns the ProgressTaskId field if non-nil, zero value otherwise.

### GetProgressTaskIdOk

`func (o *CloudSpinTargetResultAllOf) GetProgressTaskIdOk() (*string, bool)`

GetProgressTaskIdOk returns a tuple with the ProgressTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressTaskId

`func (o *CloudSpinTargetResultAllOf) SetProgressTaskId(v string)`

SetProgressTaskId sets ProgressTaskId field to given value.

### HasProgressTaskId

`func (o *CloudSpinTargetResultAllOf) HasProgressTaskId() bool`

HasProgressTaskId returns a boolean if a field has been set.

### SetProgressTaskIdNil

`func (o *CloudSpinTargetResultAllOf) SetProgressTaskIdNil(b bool)`

 SetProgressTaskIdNil sets the value for ProgressTaskId to be an explicit nil

### UnsetProgressTaskId
`func (o *CloudSpinTargetResultAllOf) UnsetProgressTaskId()`

UnsetProgressTaskId ensures that no value is present for ProgressTaskId, not even an explicit nil
### GetDataLockConstraints

`func (o *CloudSpinTargetResultAllOf) GetDataLockConstraints() DataLockConstraints`

GetDataLockConstraints returns the DataLockConstraints field if non-nil, zero value otherwise.

### GetDataLockConstraintsOk

`func (o *CloudSpinTargetResultAllOf) GetDataLockConstraintsOk() (*DataLockConstraints, bool)`

GetDataLockConstraintsOk returns a tuple with the DataLockConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLockConstraints

`func (o *CloudSpinTargetResultAllOf) SetDataLockConstraints(v DataLockConstraints)`

SetDataLockConstraints sets DataLockConstraints field to given value.

### HasDataLockConstraints

`func (o *CloudSpinTargetResultAllOf) HasDataLockConstraints() bool`

HasDataLockConstraints returns a boolean if a field has been set.

### GetOnLegalHold

`func (o *CloudSpinTargetResultAllOf) GetOnLegalHold() bool`

GetOnLegalHold returns the OnLegalHold field if non-nil, zero value otherwise.

### GetOnLegalHoldOk

`func (o *CloudSpinTargetResultAllOf) GetOnLegalHoldOk() (*bool, bool)`

GetOnLegalHoldOk returns a tuple with the OnLegalHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnLegalHold

`func (o *CloudSpinTargetResultAllOf) SetOnLegalHold(v bool)`

SetOnLegalHold sets OnLegalHold field to given value.

### HasOnLegalHold

`func (o *CloudSpinTargetResultAllOf) HasOnLegalHold() bool`

HasOnLegalHold returns a boolean if a field has been set.

### SetOnLegalHoldNil

`func (o *CloudSpinTargetResultAllOf) SetOnLegalHoldNil(b bool)`

 SetOnLegalHoldNil sets the value for OnLegalHold to be an explicit nil

### UnsetOnLegalHold
`func (o *CloudSpinTargetResultAllOf) UnsetOnLegalHold()`

UnsetOnLegalHold ensures that no value is present for OnLegalHold, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


