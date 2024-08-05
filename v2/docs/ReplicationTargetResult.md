# ReplicationTargetResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **NullableInt64** | Specifies the id of the cluster. | [optional] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | Specifies the incarnation id of the cluster. | [optional] 
**ClusterName** | Pointer to **NullableString** | Specifies the name of the cluster. | [optional] [readonly] 
**AwsTargetConfig** | Pointer to [**AWSTargetConfig**](AWSTargetConfig.md) |  | [optional] 
**AzureTargetConfig** | Pointer to [**AzureTargetConfig**](AzureTargetConfig.md) |  | [optional] 
**DataLockConstraints** | Pointer to [**DataLockConstraints**](DataLockConstraints.md) |  | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of replication in Unix epoch Timestamp(in microseconds) for a target. | [optional] 
**EntriesChanged** | Pointer to **NullableInt64** | Specifies the number of metadata actions completed during the protection run. | [optional] 
**ExpiryTimeUsecs** | Pointer to **NullableInt64** | Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds) for an object. | [optional] 
**IsInBound** | Pointer to **NullableBool** | Specifies the direction of the replication. If the snapshot is replicated to this cluster, then isInBound is true. If the snapshot is replicated from this cluster to another cluster, then isInBound is false. | [optional] 
**IsManuallyDeleted** | Pointer to **NullableBool** | Specifies whether the snapshot is deleted manually. | [optional] 
**Message** | Pointer to **NullableString** | Message about the replication run. | [optional] 
**MultiObjectReplication** | Pointer to **NullableBool** | Specifies whether view based replication was used. In this case, the view containing all objects is replicated as a whole instead of replicating on a per object basis. | [optional] 
**OnLegalHold** | Pointer to **NullableBool** | Specifies the legal hold status for a replication target. | [optional] 
**PercentageCompleted** | Pointer to **NullableInt32** | Specifies the progress in percentage. | [optional] 
**QueuedTimeUsecs** | Pointer to **NullableInt64** | Specifies the time when the replication is queued for schedule in Unix epoch Timestamp(in microseconds) for a target. | [optional] 
**ReplicationTaskId** | Pointer to **NullableString** | Task UID for a replication protection run. This is for tasks that are replicated from another cluster. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of replication in Unix epoch Timestamp(in microseconds) for a target. | [optional] 
**Stats** | Pointer to [**ReplicationDataStats**](ReplicationDataStats.md) |  | [optional] 
**Status** | Pointer to **NullableString** | Status of the replication for a target. &#39;Running&#39; indicates that the run is still running. &#39;Canceled&#39; indicates that the run has been canceled. &#39;Canceling&#39; indicates that the run is in the process of being canceled. &#39;Paused&#39; indicates that the ongoing run has been paused. &#39;Failed&#39; indicates that the run has failed. &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening. &#39;Succeeded&#39; indicates that the run has finished successfully. &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages. &#39;Skipped&#39; indicates that the run was skipped. | [optional] 

## Methods

### NewReplicationTargetResult

`func NewReplicationTargetResult() *ReplicationTargetResult`

NewReplicationTargetResult instantiates a new ReplicationTargetResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationTargetResultWithDefaults

`func NewReplicationTargetResultWithDefaults() *ReplicationTargetResult`

NewReplicationTargetResultWithDefaults instantiates a new ReplicationTargetResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ReplicationTargetResult) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ReplicationTargetResult) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ReplicationTargetResult) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ReplicationTargetResult) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *ReplicationTargetResult) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *ReplicationTargetResult) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *ReplicationTargetResult) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *ReplicationTargetResult) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *ReplicationTargetResult) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *ReplicationTargetResult) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *ReplicationTargetResult) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *ReplicationTargetResult) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetClusterName

`func (o *ReplicationTargetResult) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *ReplicationTargetResult) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *ReplicationTargetResult) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *ReplicationTargetResult) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### SetClusterNameNil

`func (o *ReplicationTargetResult) SetClusterNameNil(b bool)`

 SetClusterNameNil sets the value for ClusterName to be an explicit nil

### UnsetClusterName
`func (o *ReplicationTargetResult) UnsetClusterName()`

UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil
### GetAwsTargetConfig

`func (o *ReplicationTargetResult) GetAwsTargetConfig() AWSTargetConfig`

GetAwsTargetConfig returns the AwsTargetConfig field if non-nil, zero value otherwise.

### GetAwsTargetConfigOk

`func (o *ReplicationTargetResult) GetAwsTargetConfigOk() (*AWSTargetConfig, bool)`

GetAwsTargetConfigOk returns a tuple with the AwsTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsTargetConfig

`func (o *ReplicationTargetResult) SetAwsTargetConfig(v AWSTargetConfig)`

SetAwsTargetConfig sets AwsTargetConfig field to given value.

### HasAwsTargetConfig

`func (o *ReplicationTargetResult) HasAwsTargetConfig() bool`

HasAwsTargetConfig returns a boolean if a field has been set.

### GetAzureTargetConfig

`func (o *ReplicationTargetResult) GetAzureTargetConfig() AzureTargetConfig`

GetAzureTargetConfig returns the AzureTargetConfig field if non-nil, zero value otherwise.

### GetAzureTargetConfigOk

`func (o *ReplicationTargetResult) GetAzureTargetConfigOk() (*AzureTargetConfig, bool)`

GetAzureTargetConfigOk returns a tuple with the AzureTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTargetConfig

`func (o *ReplicationTargetResult) SetAzureTargetConfig(v AzureTargetConfig)`

SetAzureTargetConfig sets AzureTargetConfig field to given value.

### HasAzureTargetConfig

`func (o *ReplicationTargetResult) HasAzureTargetConfig() bool`

HasAzureTargetConfig returns a boolean if a field has been set.

### GetDataLockConstraints

`func (o *ReplicationTargetResult) GetDataLockConstraints() DataLockConstraints`

GetDataLockConstraints returns the DataLockConstraints field if non-nil, zero value otherwise.

### GetDataLockConstraintsOk

`func (o *ReplicationTargetResult) GetDataLockConstraintsOk() (*DataLockConstraints, bool)`

GetDataLockConstraintsOk returns a tuple with the DataLockConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLockConstraints

`func (o *ReplicationTargetResult) SetDataLockConstraints(v DataLockConstraints)`

SetDataLockConstraints sets DataLockConstraints field to given value.

### HasDataLockConstraints

`func (o *ReplicationTargetResult) HasDataLockConstraints() bool`

HasDataLockConstraints returns a boolean if a field has been set.

### GetEndTimeUsecs

`func (o *ReplicationTargetResult) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *ReplicationTargetResult) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *ReplicationTargetResult) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *ReplicationTargetResult) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *ReplicationTargetResult) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *ReplicationTargetResult) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetEntriesChanged

`func (o *ReplicationTargetResult) GetEntriesChanged() int64`

GetEntriesChanged returns the EntriesChanged field if non-nil, zero value otherwise.

### GetEntriesChangedOk

`func (o *ReplicationTargetResult) GetEntriesChangedOk() (*int64, bool)`

GetEntriesChangedOk returns a tuple with the EntriesChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntriesChanged

`func (o *ReplicationTargetResult) SetEntriesChanged(v int64)`

SetEntriesChanged sets EntriesChanged field to given value.

### HasEntriesChanged

`func (o *ReplicationTargetResult) HasEntriesChanged() bool`

HasEntriesChanged returns a boolean if a field has been set.

### SetEntriesChangedNil

`func (o *ReplicationTargetResult) SetEntriesChangedNil(b bool)`

 SetEntriesChangedNil sets the value for EntriesChanged to be an explicit nil

### UnsetEntriesChanged
`func (o *ReplicationTargetResult) UnsetEntriesChanged()`

UnsetEntriesChanged ensures that no value is present for EntriesChanged, not even an explicit nil
### GetExpiryTimeUsecs

`func (o *ReplicationTargetResult) GetExpiryTimeUsecs() int64`

GetExpiryTimeUsecs returns the ExpiryTimeUsecs field if non-nil, zero value otherwise.

### GetExpiryTimeUsecsOk

`func (o *ReplicationTargetResult) GetExpiryTimeUsecsOk() (*int64, bool)`

GetExpiryTimeUsecsOk returns a tuple with the ExpiryTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeUsecs

`func (o *ReplicationTargetResult) SetExpiryTimeUsecs(v int64)`

SetExpiryTimeUsecs sets ExpiryTimeUsecs field to given value.

### HasExpiryTimeUsecs

`func (o *ReplicationTargetResult) HasExpiryTimeUsecs() bool`

HasExpiryTimeUsecs returns a boolean if a field has been set.

### SetExpiryTimeUsecsNil

`func (o *ReplicationTargetResult) SetExpiryTimeUsecsNil(b bool)`

 SetExpiryTimeUsecsNil sets the value for ExpiryTimeUsecs to be an explicit nil

### UnsetExpiryTimeUsecs
`func (o *ReplicationTargetResult) UnsetExpiryTimeUsecs()`

UnsetExpiryTimeUsecs ensures that no value is present for ExpiryTimeUsecs, not even an explicit nil
### GetIsInBound

`func (o *ReplicationTargetResult) GetIsInBound() bool`

GetIsInBound returns the IsInBound field if non-nil, zero value otherwise.

### GetIsInBoundOk

`func (o *ReplicationTargetResult) GetIsInBoundOk() (*bool, bool)`

GetIsInBoundOk returns a tuple with the IsInBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInBound

`func (o *ReplicationTargetResult) SetIsInBound(v bool)`

SetIsInBound sets IsInBound field to given value.

### HasIsInBound

`func (o *ReplicationTargetResult) HasIsInBound() bool`

HasIsInBound returns a boolean if a field has been set.

### SetIsInBoundNil

`func (o *ReplicationTargetResult) SetIsInBoundNil(b bool)`

 SetIsInBoundNil sets the value for IsInBound to be an explicit nil

### UnsetIsInBound
`func (o *ReplicationTargetResult) UnsetIsInBound()`

UnsetIsInBound ensures that no value is present for IsInBound, not even an explicit nil
### GetIsManuallyDeleted

`func (o *ReplicationTargetResult) GetIsManuallyDeleted() bool`

GetIsManuallyDeleted returns the IsManuallyDeleted field if non-nil, zero value otherwise.

### GetIsManuallyDeletedOk

`func (o *ReplicationTargetResult) GetIsManuallyDeletedOk() (*bool, bool)`

GetIsManuallyDeletedOk returns a tuple with the IsManuallyDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManuallyDeleted

`func (o *ReplicationTargetResult) SetIsManuallyDeleted(v bool)`

SetIsManuallyDeleted sets IsManuallyDeleted field to given value.

### HasIsManuallyDeleted

`func (o *ReplicationTargetResult) HasIsManuallyDeleted() bool`

HasIsManuallyDeleted returns a boolean if a field has been set.

### SetIsManuallyDeletedNil

`func (o *ReplicationTargetResult) SetIsManuallyDeletedNil(b bool)`

 SetIsManuallyDeletedNil sets the value for IsManuallyDeleted to be an explicit nil

### UnsetIsManuallyDeleted
`func (o *ReplicationTargetResult) UnsetIsManuallyDeleted()`

UnsetIsManuallyDeleted ensures that no value is present for IsManuallyDeleted, not even an explicit nil
### GetMessage

`func (o *ReplicationTargetResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ReplicationTargetResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ReplicationTargetResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ReplicationTargetResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ReplicationTargetResult) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ReplicationTargetResult) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetMultiObjectReplication

`func (o *ReplicationTargetResult) GetMultiObjectReplication() bool`

GetMultiObjectReplication returns the MultiObjectReplication field if non-nil, zero value otherwise.

### GetMultiObjectReplicationOk

`func (o *ReplicationTargetResult) GetMultiObjectReplicationOk() (*bool, bool)`

GetMultiObjectReplicationOk returns a tuple with the MultiObjectReplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiObjectReplication

`func (o *ReplicationTargetResult) SetMultiObjectReplication(v bool)`

SetMultiObjectReplication sets MultiObjectReplication field to given value.

### HasMultiObjectReplication

`func (o *ReplicationTargetResult) HasMultiObjectReplication() bool`

HasMultiObjectReplication returns a boolean if a field has been set.

### SetMultiObjectReplicationNil

`func (o *ReplicationTargetResult) SetMultiObjectReplicationNil(b bool)`

 SetMultiObjectReplicationNil sets the value for MultiObjectReplication to be an explicit nil

### UnsetMultiObjectReplication
`func (o *ReplicationTargetResult) UnsetMultiObjectReplication()`

UnsetMultiObjectReplication ensures that no value is present for MultiObjectReplication, not even an explicit nil
### GetOnLegalHold

`func (o *ReplicationTargetResult) GetOnLegalHold() bool`

GetOnLegalHold returns the OnLegalHold field if non-nil, zero value otherwise.

### GetOnLegalHoldOk

`func (o *ReplicationTargetResult) GetOnLegalHoldOk() (*bool, bool)`

GetOnLegalHoldOk returns a tuple with the OnLegalHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnLegalHold

`func (o *ReplicationTargetResult) SetOnLegalHold(v bool)`

SetOnLegalHold sets OnLegalHold field to given value.

### HasOnLegalHold

`func (o *ReplicationTargetResult) HasOnLegalHold() bool`

HasOnLegalHold returns a boolean if a field has been set.

### SetOnLegalHoldNil

`func (o *ReplicationTargetResult) SetOnLegalHoldNil(b bool)`

 SetOnLegalHoldNil sets the value for OnLegalHold to be an explicit nil

### UnsetOnLegalHold
`func (o *ReplicationTargetResult) UnsetOnLegalHold()`

UnsetOnLegalHold ensures that no value is present for OnLegalHold, not even an explicit nil
### GetPercentageCompleted

`func (o *ReplicationTargetResult) GetPercentageCompleted() int32`

GetPercentageCompleted returns the PercentageCompleted field if non-nil, zero value otherwise.

### GetPercentageCompletedOk

`func (o *ReplicationTargetResult) GetPercentageCompletedOk() (*int32, bool)`

GetPercentageCompletedOk returns a tuple with the PercentageCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageCompleted

`func (o *ReplicationTargetResult) SetPercentageCompleted(v int32)`

SetPercentageCompleted sets PercentageCompleted field to given value.

### HasPercentageCompleted

`func (o *ReplicationTargetResult) HasPercentageCompleted() bool`

HasPercentageCompleted returns a boolean if a field has been set.

### SetPercentageCompletedNil

`func (o *ReplicationTargetResult) SetPercentageCompletedNil(b bool)`

 SetPercentageCompletedNil sets the value for PercentageCompleted to be an explicit nil

### UnsetPercentageCompleted
`func (o *ReplicationTargetResult) UnsetPercentageCompleted()`

UnsetPercentageCompleted ensures that no value is present for PercentageCompleted, not even an explicit nil
### GetQueuedTimeUsecs

`func (o *ReplicationTargetResult) GetQueuedTimeUsecs() int64`

GetQueuedTimeUsecs returns the QueuedTimeUsecs field if non-nil, zero value otherwise.

### GetQueuedTimeUsecsOk

`func (o *ReplicationTargetResult) GetQueuedTimeUsecsOk() (*int64, bool)`

GetQueuedTimeUsecsOk returns a tuple with the QueuedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedTimeUsecs

`func (o *ReplicationTargetResult) SetQueuedTimeUsecs(v int64)`

SetQueuedTimeUsecs sets QueuedTimeUsecs field to given value.

### HasQueuedTimeUsecs

`func (o *ReplicationTargetResult) HasQueuedTimeUsecs() bool`

HasQueuedTimeUsecs returns a boolean if a field has been set.

### SetQueuedTimeUsecsNil

`func (o *ReplicationTargetResult) SetQueuedTimeUsecsNil(b bool)`

 SetQueuedTimeUsecsNil sets the value for QueuedTimeUsecs to be an explicit nil

### UnsetQueuedTimeUsecs
`func (o *ReplicationTargetResult) UnsetQueuedTimeUsecs()`

UnsetQueuedTimeUsecs ensures that no value is present for QueuedTimeUsecs, not even an explicit nil
### GetReplicationTaskId

`func (o *ReplicationTargetResult) GetReplicationTaskId() string`

GetReplicationTaskId returns the ReplicationTaskId field if non-nil, zero value otherwise.

### GetReplicationTaskIdOk

`func (o *ReplicationTargetResult) GetReplicationTaskIdOk() (*string, bool)`

GetReplicationTaskIdOk returns a tuple with the ReplicationTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationTaskId

`func (o *ReplicationTargetResult) SetReplicationTaskId(v string)`

SetReplicationTaskId sets ReplicationTaskId field to given value.

### HasReplicationTaskId

`func (o *ReplicationTargetResult) HasReplicationTaskId() bool`

HasReplicationTaskId returns a boolean if a field has been set.

### SetReplicationTaskIdNil

`func (o *ReplicationTargetResult) SetReplicationTaskIdNil(b bool)`

 SetReplicationTaskIdNil sets the value for ReplicationTaskId to be an explicit nil

### UnsetReplicationTaskId
`func (o *ReplicationTargetResult) UnsetReplicationTaskId()`

UnsetReplicationTaskId ensures that no value is present for ReplicationTaskId, not even an explicit nil
### GetStartTimeUsecs

`func (o *ReplicationTargetResult) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *ReplicationTargetResult) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *ReplicationTargetResult) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *ReplicationTargetResult) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *ReplicationTargetResult) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *ReplicationTargetResult) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetStats

`func (o *ReplicationTargetResult) GetStats() ReplicationDataStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ReplicationTargetResult) GetStatsOk() (*ReplicationDataStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ReplicationTargetResult) SetStats(v ReplicationDataStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ReplicationTargetResult) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStatus

`func (o *ReplicationTargetResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReplicationTargetResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReplicationTargetResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReplicationTargetResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ReplicationTargetResult) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ReplicationTargetResult) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


