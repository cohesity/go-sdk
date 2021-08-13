# ReplicationTargetResultAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of replication in Unix epoch Timestamp(in microseconds) for a target. | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of replication in Unix epoch Timestamp(in microseconds) for a target. | [optional] 
**QueuedTimeUsecs** | Pointer to **NullableInt64** | Specifies the time when the replication is queued for schedule in Unix epoch Timestamp(in microseconds) for a target. | [optional] 
**Status** | Pointer to **NullableString** | Status of the replication for a target. &#39;Running&#39; indicates that the run is still running. &#39;Canceled&#39; indicates that the run has been canceled. &#39;Canceling&#39; indicates that the run is in the process of being canceled. &#39;Paused&#39; indicates that the ongoing run has been paused. &#39;Failed&#39; indicates that the run has failed. &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening. &#39;Succeeded&#39; indicates that the run has finished successfully. &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages. | [optional] 
**Message** | Pointer to **NullableString** | Message about the replication run. | [optional] 
**PercentageCompleted** | Pointer to **NullableInt32** | Specifies the progress in percentage. | [optional] 
**Stats** | Pointer to [**ReplicationDataStats**](ReplicationDataStats.md) |  | [optional] 
**IsManuallyDeleted** | Pointer to **NullableBool** | Specifies whether the snapshot is deleted manually. | [optional] 
**ExpiryTimeUsecs** | Pointer to **NullableInt64** | Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds) for an object. | [optional] 
**ReplicationTaskId** | Pointer to **NullableString** | Task UID for a replication protection run. This is for tasks that are replicated from another cluster. | [optional] 
**EntriesChanged** | Pointer to **NullableInt32** | Specifies the number of metadata actions completed during the protection run. | [optional] 
**IsInBound** | Pointer to **NullableBool** | Specifies the direction of the replication. If the snapshot is replicated to this cluster, then isInBound is true. If the snapshot is replicated from this cluster to another cluster, then isInBound is false. | [optional] 
**DataLockConstraints** | Pointer to [**DataLockConstraints**](DataLockConstraints.md) |  | [optional] 
**OnLegalHold** | Pointer to **NullableBool** | Specifies the legal hold status for a replication target. | [optional] 

## Methods

### NewReplicationTargetResultAllOf

`func NewReplicationTargetResultAllOf() *ReplicationTargetResultAllOf`

NewReplicationTargetResultAllOf instantiates a new ReplicationTargetResultAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationTargetResultAllOfWithDefaults

`func NewReplicationTargetResultAllOfWithDefaults() *ReplicationTargetResultAllOf`

NewReplicationTargetResultAllOfWithDefaults instantiates a new ReplicationTargetResultAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTimeUsecs

`func (o *ReplicationTargetResultAllOf) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *ReplicationTargetResultAllOf) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *ReplicationTargetResultAllOf) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *ReplicationTargetResultAllOf) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *ReplicationTargetResultAllOf) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *ReplicationTargetResultAllOf) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetEndTimeUsecs

`func (o *ReplicationTargetResultAllOf) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *ReplicationTargetResultAllOf) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *ReplicationTargetResultAllOf) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *ReplicationTargetResultAllOf) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *ReplicationTargetResultAllOf) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *ReplicationTargetResultAllOf) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetQueuedTimeUsecs

`func (o *ReplicationTargetResultAllOf) GetQueuedTimeUsecs() int64`

GetQueuedTimeUsecs returns the QueuedTimeUsecs field if non-nil, zero value otherwise.

### GetQueuedTimeUsecsOk

`func (o *ReplicationTargetResultAllOf) GetQueuedTimeUsecsOk() (*int64, bool)`

GetQueuedTimeUsecsOk returns a tuple with the QueuedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedTimeUsecs

`func (o *ReplicationTargetResultAllOf) SetQueuedTimeUsecs(v int64)`

SetQueuedTimeUsecs sets QueuedTimeUsecs field to given value.

### HasQueuedTimeUsecs

`func (o *ReplicationTargetResultAllOf) HasQueuedTimeUsecs() bool`

HasQueuedTimeUsecs returns a boolean if a field has been set.

### SetQueuedTimeUsecsNil

`func (o *ReplicationTargetResultAllOf) SetQueuedTimeUsecsNil(b bool)`

 SetQueuedTimeUsecsNil sets the value for QueuedTimeUsecs to be an explicit nil

### UnsetQueuedTimeUsecs
`func (o *ReplicationTargetResultAllOf) UnsetQueuedTimeUsecs()`

UnsetQueuedTimeUsecs ensures that no value is present for QueuedTimeUsecs, not even an explicit nil
### GetStatus

`func (o *ReplicationTargetResultAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReplicationTargetResultAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReplicationTargetResultAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReplicationTargetResultAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ReplicationTargetResultAllOf) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ReplicationTargetResultAllOf) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetMessage

`func (o *ReplicationTargetResultAllOf) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ReplicationTargetResultAllOf) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ReplicationTargetResultAllOf) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ReplicationTargetResultAllOf) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ReplicationTargetResultAllOf) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ReplicationTargetResultAllOf) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetPercentageCompleted

`func (o *ReplicationTargetResultAllOf) GetPercentageCompleted() int32`

GetPercentageCompleted returns the PercentageCompleted field if non-nil, zero value otherwise.

### GetPercentageCompletedOk

`func (o *ReplicationTargetResultAllOf) GetPercentageCompletedOk() (*int32, bool)`

GetPercentageCompletedOk returns a tuple with the PercentageCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageCompleted

`func (o *ReplicationTargetResultAllOf) SetPercentageCompleted(v int32)`

SetPercentageCompleted sets PercentageCompleted field to given value.

### HasPercentageCompleted

`func (o *ReplicationTargetResultAllOf) HasPercentageCompleted() bool`

HasPercentageCompleted returns a boolean if a field has been set.

### SetPercentageCompletedNil

`func (o *ReplicationTargetResultAllOf) SetPercentageCompletedNil(b bool)`

 SetPercentageCompletedNil sets the value for PercentageCompleted to be an explicit nil

### UnsetPercentageCompleted
`func (o *ReplicationTargetResultAllOf) UnsetPercentageCompleted()`

UnsetPercentageCompleted ensures that no value is present for PercentageCompleted, not even an explicit nil
### GetStats

`func (o *ReplicationTargetResultAllOf) GetStats() ReplicationDataStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ReplicationTargetResultAllOf) GetStatsOk() (*ReplicationDataStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ReplicationTargetResultAllOf) SetStats(v ReplicationDataStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ReplicationTargetResultAllOf) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetIsManuallyDeleted

`func (o *ReplicationTargetResultAllOf) GetIsManuallyDeleted() bool`

GetIsManuallyDeleted returns the IsManuallyDeleted field if non-nil, zero value otherwise.

### GetIsManuallyDeletedOk

`func (o *ReplicationTargetResultAllOf) GetIsManuallyDeletedOk() (*bool, bool)`

GetIsManuallyDeletedOk returns a tuple with the IsManuallyDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManuallyDeleted

`func (o *ReplicationTargetResultAllOf) SetIsManuallyDeleted(v bool)`

SetIsManuallyDeleted sets IsManuallyDeleted field to given value.

### HasIsManuallyDeleted

`func (o *ReplicationTargetResultAllOf) HasIsManuallyDeleted() bool`

HasIsManuallyDeleted returns a boolean if a field has been set.

### SetIsManuallyDeletedNil

`func (o *ReplicationTargetResultAllOf) SetIsManuallyDeletedNil(b bool)`

 SetIsManuallyDeletedNil sets the value for IsManuallyDeleted to be an explicit nil

### UnsetIsManuallyDeleted
`func (o *ReplicationTargetResultAllOf) UnsetIsManuallyDeleted()`

UnsetIsManuallyDeleted ensures that no value is present for IsManuallyDeleted, not even an explicit nil
### GetExpiryTimeUsecs

`func (o *ReplicationTargetResultAllOf) GetExpiryTimeUsecs() int64`

GetExpiryTimeUsecs returns the ExpiryTimeUsecs field if non-nil, zero value otherwise.

### GetExpiryTimeUsecsOk

`func (o *ReplicationTargetResultAllOf) GetExpiryTimeUsecsOk() (*int64, bool)`

GetExpiryTimeUsecsOk returns a tuple with the ExpiryTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeUsecs

`func (o *ReplicationTargetResultAllOf) SetExpiryTimeUsecs(v int64)`

SetExpiryTimeUsecs sets ExpiryTimeUsecs field to given value.

### HasExpiryTimeUsecs

`func (o *ReplicationTargetResultAllOf) HasExpiryTimeUsecs() bool`

HasExpiryTimeUsecs returns a boolean if a field has been set.

### SetExpiryTimeUsecsNil

`func (o *ReplicationTargetResultAllOf) SetExpiryTimeUsecsNil(b bool)`

 SetExpiryTimeUsecsNil sets the value for ExpiryTimeUsecs to be an explicit nil

### UnsetExpiryTimeUsecs
`func (o *ReplicationTargetResultAllOf) UnsetExpiryTimeUsecs()`

UnsetExpiryTimeUsecs ensures that no value is present for ExpiryTimeUsecs, not even an explicit nil
### GetReplicationTaskId

`func (o *ReplicationTargetResultAllOf) GetReplicationTaskId() string`

GetReplicationTaskId returns the ReplicationTaskId field if non-nil, zero value otherwise.

### GetReplicationTaskIdOk

`func (o *ReplicationTargetResultAllOf) GetReplicationTaskIdOk() (*string, bool)`

GetReplicationTaskIdOk returns a tuple with the ReplicationTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationTaskId

`func (o *ReplicationTargetResultAllOf) SetReplicationTaskId(v string)`

SetReplicationTaskId sets ReplicationTaskId field to given value.

### HasReplicationTaskId

`func (o *ReplicationTargetResultAllOf) HasReplicationTaskId() bool`

HasReplicationTaskId returns a boolean if a field has been set.

### SetReplicationTaskIdNil

`func (o *ReplicationTargetResultAllOf) SetReplicationTaskIdNil(b bool)`

 SetReplicationTaskIdNil sets the value for ReplicationTaskId to be an explicit nil

### UnsetReplicationTaskId
`func (o *ReplicationTargetResultAllOf) UnsetReplicationTaskId()`

UnsetReplicationTaskId ensures that no value is present for ReplicationTaskId, not even an explicit nil
### GetEntriesChanged

`func (o *ReplicationTargetResultAllOf) GetEntriesChanged() int32`

GetEntriesChanged returns the EntriesChanged field if non-nil, zero value otherwise.

### GetEntriesChangedOk

`func (o *ReplicationTargetResultAllOf) GetEntriesChangedOk() (*int32, bool)`

GetEntriesChangedOk returns a tuple with the EntriesChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntriesChanged

`func (o *ReplicationTargetResultAllOf) SetEntriesChanged(v int32)`

SetEntriesChanged sets EntriesChanged field to given value.

### HasEntriesChanged

`func (o *ReplicationTargetResultAllOf) HasEntriesChanged() bool`

HasEntriesChanged returns a boolean if a field has been set.

### SetEntriesChangedNil

`func (o *ReplicationTargetResultAllOf) SetEntriesChangedNil(b bool)`

 SetEntriesChangedNil sets the value for EntriesChanged to be an explicit nil

### UnsetEntriesChanged
`func (o *ReplicationTargetResultAllOf) UnsetEntriesChanged()`

UnsetEntriesChanged ensures that no value is present for EntriesChanged, not even an explicit nil
### GetIsInBound

`func (o *ReplicationTargetResultAllOf) GetIsInBound() bool`

GetIsInBound returns the IsInBound field if non-nil, zero value otherwise.

### GetIsInBoundOk

`func (o *ReplicationTargetResultAllOf) GetIsInBoundOk() (*bool, bool)`

GetIsInBoundOk returns a tuple with the IsInBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInBound

`func (o *ReplicationTargetResultAllOf) SetIsInBound(v bool)`

SetIsInBound sets IsInBound field to given value.

### HasIsInBound

`func (o *ReplicationTargetResultAllOf) HasIsInBound() bool`

HasIsInBound returns a boolean if a field has been set.

### SetIsInBoundNil

`func (o *ReplicationTargetResultAllOf) SetIsInBoundNil(b bool)`

 SetIsInBoundNil sets the value for IsInBound to be an explicit nil

### UnsetIsInBound
`func (o *ReplicationTargetResultAllOf) UnsetIsInBound()`

UnsetIsInBound ensures that no value is present for IsInBound, not even an explicit nil
### GetDataLockConstraints

`func (o *ReplicationTargetResultAllOf) GetDataLockConstraints() DataLockConstraints`

GetDataLockConstraints returns the DataLockConstraints field if non-nil, zero value otherwise.

### GetDataLockConstraintsOk

`func (o *ReplicationTargetResultAllOf) GetDataLockConstraintsOk() (*DataLockConstraints, bool)`

GetDataLockConstraintsOk returns a tuple with the DataLockConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLockConstraints

`func (o *ReplicationTargetResultAllOf) SetDataLockConstraints(v DataLockConstraints)`

SetDataLockConstraints sets DataLockConstraints field to given value.

### HasDataLockConstraints

`func (o *ReplicationTargetResultAllOf) HasDataLockConstraints() bool`

HasDataLockConstraints returns a boolean if a field has been set.

### GetOnLegalHold

`func (o *ReplicationTargetResultAllOf) GetOnLegalHold() bool`

GetOnLegalHold returns the OnLegalHold field if non-nil, zero value otherwise.

### GetOnLegalHoldOk

`func (o *ReplicationTargetResultAllOf) GetOnLegalHoldOk() (*bool, bool)`

GetOnLegalHoldOk returns a tuple with the OnLegalHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnLegalHold

`func (o *ReplicationTargetResultAllOf) SetOnLegalHold(v bool)`

SetOnLegalHold sets OnLegalHold field to given value.

### HasOnLegalHold

`func (o *ReplicationTargetResultAllOf) HasOnLegalHold() bool`

HasOnLegalHold returns a boolean if a field has been set.

### SetOnLegalHoldNil

`func (o *ReplicationTargetResultAllOf) SetOnLegalHoldNil(b bool)`

 SetOnLegalHoldNil sets the value for OnLegalHold to be an explicit nil

### UnsetOnLegalHold
`func (o *ReplicationTargetResultAllOf) UnsetOnLegalHold()`

UnsetOnLegalHold ensures that no value is present for OnLegalHold, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


