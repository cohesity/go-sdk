# McmObjectBackupRunActivityParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunId** | Pointer to **NullableString** | Specifies the ID of the Protection Run. | [optional] 
**ProtectionGroupId** | Pointer to **NullableString** | Specifies the Protection Group Id. | [optional] 
**ProtectionGroupName** | Pointer to **NullableString** | Specifies the Protection Group name. | [optional] 
**PolicyId** | Pointer to **NullableString** | Specifies the Protection Policy Id. | [optional] 
**PolicyName** | Pointer to **NullableString** | Specifies the Protection Policy Name. | [optional] 
**SnapshotId** | Pointer to **NullableString** | Specifies the id of the object snapshot that is created as a part of this Run. This field is only populated for runs which are successful. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of Run in Unix epoch Timestamp(in microseconds). | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of Run in Unix epoch Timestamp(in microseconds). | [optional] 
**Status** | Pointer to **NullableString** | Status of the Run. &#39;Running&#39; indicates that the run is still running. &#39;Canceled&#39; indicates that the run has been canceled. &#39;Canceling&#39; indicates that the run is in the process of being canceled. &#39;Failed&#39; indicates that the run has failed. &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening. &#39;Succeeded&#39; indicates that the run has finished successfully. &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages. | [optional] 
**ProgressTaskId** | Pointer to **NullableString** | Progress monitor task id for the Run. | [optional] 
**ProtectionEnvironmentType** | Pointer to **NullableString** | Specifies the type of protection environment. | [optional] 

## Methods

### NewMcmObjectBackupRunActivityParams

`func NewMcmObjectBackupRunActivityParams() *McmObjectBackupRunActivityParams`

NewMcmObjectBackupRunActivityParams instantiates a new McmObjectBackupRunActivityParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmObjectBackupRunActivityParamsWithDefaults

`func NewMcmObjectBackupRunActivityParamsWithDefaults() *McmObjectBackupRunActivityParams`

NewMcmObjectBackupRunActivityParamsWithDefaults instantiates a new McmObjectBackupRunActivityParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunId

`func (o *McmObjectBackupRunActivityParams) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *McmObjectBackupRunActivityParams) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *McmObjectBackupRunActivityParams) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *McmObjectBackupRunActivityParams) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### SetRunIdNil

`func (o *McmObjectBackupRunActivityParams) SetRunIdNil(b bool)`

 SetRunIdNil sets the value for RunId to be an explicit nil

### UnsetRunId
`func (o *McmObjectBackupRunActivityParams) UnsetRunId()`

UnsetRunId ensures that no value is present for RunId, not even an explicit nil
### GetProtectionGroupId

`func (o *McmObjectBackupRunActivityParams) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *McmObjectBackupRunActivityParams) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *McmObjectBackupRunActivityParams) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *McmObjectBackupRunActivityParams) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *McmObjectBackupRunActivityParams) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *McmObjectBackupRunActivityParams) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetProtectionGroupName

`func (o *McmObjectBackupRunActivityParams) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *McmObjectBackupRunActivityParams) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *McmObjectBackupRunActivityParams) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *McmObjectBackupRunActivityParams) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### SetProtectionGroupNameNil

`func (o *McmObjectBackupRunActivityParams) SetProtectionGroupNameNil(b bool)`

 SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil

### UnsetProtectionGroupName
`func (o *McmObjectBackupRunActivityParams) UnsetProtectionGroupName()`

UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
### GetPolicyId

`func (o *McmObjectBackupRunActivityParams) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *McmObjectBackupRunActivityParams) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *McmObjectBackupRunActivityParams) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *McmObjectBackupRunActivityParams) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *McmObjectBackupRunActivityParams) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *McmObjectBackupRunActivityParams) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetPolicyName

`func (o *McmObjectBackupRunActivityParams) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *McmObjectBackupRunActivityParams) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *McmObjectBackupRunActivityParams) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *McmObjectBackupRunActivityParams) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *McmObjectBackupRunActivityParams) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *McmObjectBackupRunActivityParams) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetSnapshotId

`func (o *McmObjectBackupRunActivityParams) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *McmObjectBackupRunActivityParams) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *McmObjectBackupRunActivityParams) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *McmObjectBackupRunActivityParams) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### SetSnapshotIdNil

`func (o *McmObjectBackupRunActivityParams) SetSnapshotIdNil(b bool)`

 SetSnapshotIdNil sets the value for SnapshotId to be an explicit nil

### UnsetSnapshotId
`func (o *McmObjectBackupRunActivityParams) UnsetSnapshotId()`

UnsetSnapshotId ensures that no value is present for SnapshotId, not even an explicit nil
### GetStartTimeUsecs

`func (o *McmObjectBackupRunActivityParams) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *McmObjectBackupRunActivityParams) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *McmObjectBackupRunActivityParams) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *McmObjectBackupRunActivityParams) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *McmObjectBackupRunActivityParams) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *McmObjectBackupRunActivityParams) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetEndTimeUsecs

`func (o *McmObjectBackupRunActivityParams) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *McmObjectBackupRunActivityParams) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *McmObjectBackupRunActivityParams) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *McmObjectBackupRunActivityParams) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *McmObjectBackupRunActivityParams) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *McmObjectBackupRunActivityParams) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetStatus

`func (o *McmObjectBackupRunActivityParams) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *McmObjectBackupRunActivityParams) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *McmObjectBackupRunActivityParams) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *McmObjectBackupRunActivityParams) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *McmObjectBackupRunActivityParams) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *McmObjectBackupRunActivityParams) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetProgressTaskId

`func (o *McmObjectBackupRunActivityParams) GetProgressTaskId() string`

GetProgressTaskId returns the ProgressTaskId field if non-nil, zero value otherwise.

### GetProgressTaskIdOk

`func (o *McmObjectBackupRunActivityParams) GetProgressTaskIdOk() (*string, bool)`

GetProgressTaskIdOk returns a tuple with the ProgressTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressTaskId

`func (o *McmObjectBackupRunActivityParams) SetProgressTaskId(v string)`

SetProgressTaskId sets ProgressTaskId field to given value.

### HasProgressTaskId

`func (o *McmObjectBackupRunActivityParams) HasProgressTaskId() bool`

HasProgressTaskId returns a boolean if a field has been set.

### SetProgressTaskIdNil

`func (o *McmObjectBackupRunActivityParams) SetProgressTaskIdNil(b bool)`

 SetProgressTaskIdNil sets the value for ProgressTaskId to be an explicit nil

### UnsetProgressTaskId
`func (o *McmObjectBackupRunActivityParams) UnsetProgressTaskId()`

UnsetProgressTaskId ensures that no value is present for ProgressTaskId, not even an explicit nil
### GetProtectionEnvironmentType

`func (o *McmObjectBackupRunActivityParams) GetProtectionEnvironmentType() string`

GetProtectionEnvironmentType returns the ProtectionEnvironmentType field if non-nil, zero value otherwise.

### GetProtectionEnvironmentTypeOk

`func (o *McmObjectBackupRunActivityParams) GetProtectionEnvironmentTypeOk() (*string, bool)`

GetProtectionEnvironmentTypeOk returns a tuple with the ProtectionEnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionEnvironmentType

`func (o *McmObjectBackupRunActivityParams) SetProtectionEnvironmentType(v string)`

SetProtectionEnvironmentType sets ProtectionEnvironmentType field to given value.

### HasProtectionEnvironmentType

`func (o *McmObjectBackupRunActivityParams) HasProtectionEnvironmentType() bool`

HasProtectionEnvironmentType returns a boolean if a field has been set.

### SetProtectionEnvironmentTypeNil

`func (o *McmObjectBackupRunActivityParams) SetProtectionEnvironmentTypeNil(b bool)`

 SetProtectionEnvironmentTypeNil sets the value for ProtectionEnvironmentType to be an explicit nil

### UnsetProtectionEnvironmentType
`func (o *McmObjectBackupRunActivityParams) UnsetProtectionEnvironmentType()`

UnsetProtectionEnvironmentType ensures that no value is present for ProtectionEnvironmentType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


