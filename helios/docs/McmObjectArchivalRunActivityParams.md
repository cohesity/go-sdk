# McmObjectArchivalRunActivityParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunId** | Pointer to **NullableString** | Specifies the ID of the Protection Run. | [optional] 
**RunType** | Pointer to **NullableString** | Specifies the type of the Protection Run. | [optional] 
**ProtectionGroupId** | Pointer to **NullableString** | Specifies the Protection Group Id. | [optional] 
**ProtectionGroupName** | Pointer to **NullableString** | Specifies the Protection Group name. | [optional] 
**SnapshotId** | Pointer to **NullableString** | Specifies the id of the object snapshot that is created as a part of this Run. This field is only populated for runs which are successful. | [optional] 
**PolicyId** | Pointer to **NullableString** | Specifies the Protection Policy Id. | [optional] 
**PolicyName** | Pointer to **NullableString** | Specifies the Protection Policy Name. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of Run in Unix epoch Timestamp(in microseconds). | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of Run in Unix epoch Timestamp(in microseconds). | [optional] 
**Status** | Pointer to **NullableString** | Status of the Run. &#39;Running&#39; indicates that the run is still running. &#39;Canceled&#39; indicates that the run has been canceled. &#39;Canceling&#39; indicates that the run is in the process of being canceled. &#39;Failed&#39; indicates that the run has failed. &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening. &#39;Succeeded&#39; indicates that the run has finished successfully. &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages. | [optional] 
**ProgressTaskId** | Pointer to **NullableString** | Progress monitor task id for the Run. | [optional] 
**ArchivalTargetId** | Pointer to **NullableInt64** | Specifies the id of archival target. | [optional] 
**ArchivalTargetName** | Pointer to **NullableString** | Specifies the name of archival target. | [optional] 
**ProtectionEnvironmentType** | Pointer to **NullableString** | Specifies the type of protection environment. | [optional] 

## Methods

### NewMcmObjectArchivalRunActivityParams

`func NewMcmObjectArchivalRunActivityParams() *McmObjectArchivalRunActivityParams`

NewMcmObjectArchivalRunActivityParams instantiates a new McmObjectArchivalRunActivityParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmObjectArchivalRunActivityParamsWithDefaults

`func NewMcmObjectArchivalRunActivityParamsWithDefaults() *McmObjectArchivalRunActivityParams`

NewMcmObjectArchivalRunActivityParamsWithDefaults instantiates a new McmObjectArchivalRunActivityParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunId

`func (o *McmObjectArchivalRunActivityParams) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *McmObjectArchivalRunActivityParams) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *McmObjectArchivalRunActivityParams) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *McmObjectArchivalRunActivityParams) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### SetRunIdNil

`func (o *McmObjectArchivalRunActivityParams) SetRunIdNil(b bool)`

 SetRunIdNil sets the value for RunId to be an explicit nil

### UnsetRunId
`func (o *McmObjectArchivalRunActivityParams) UnsetRunId()`

UnsetRunId ensures that no value is present for RunId, not even an explicit nil
### GetRunType

`func (o *McmObjectArchivalRunActivityParams) GetRunType() string`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *McmObjectArchivalRunActivityParams) GetRunTypeOk() (*string, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *McmObjectArchivalRunActivityParams) SetRunType(v string)`

SetRunType sets RunType field to given value.

### HasRunType

`func (o *McmObjectArchivalRunActivityParams) HasRunType() bool`

HasRunType returns a boolean if a field has been set.

### SetRunTypeNil

`func (o *McmObjectArchivalRunActivityParams) SetRunTypeNil(b bool)`

 SetRunTypeNil sets the value for RunType to be an explicit nil

### UnsetRunType
`func (o *McmObjectArchivalRunActivityParams) UnsetRunType()`

UnsetRunType ensures that no value is present for RunType, not even an explicit nil
### GetProtectionGroupId

`func (o *McmObjectArchivalRunActivityParams) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *McmObjectArchivalRunActivityParams) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *McmObjectArchivalRunActivityParams) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *McmObjectArchivalRunActivityParams) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *McmObjectArchivalRunActivityParams) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *McmObjectArchivalRunActivityParams) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetProtectionGroupName

`func (o *McmObjectArchivalRunActivityParams) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *McmObjectArchivalRunActivityParams) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *McmObjectArchivalRunActivityParams) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *McmObjectArchivalRunActivityParams) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### SetProtectionGroupNameNil

`func (o *McmObjectArchivalRunActivityParams) SetProtectionGroupNameNil(b bool)`

 SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil

### UnsetProtectionGroupName
`func (o *McmObjectArchivalRunActivityParams) UnsetProtectionGroupName()`

UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
### GetSnapshotId

`func (o *McmObjectArchivalRunActivityParams) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *McmObjectArchivalRunActivityParams) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *McmObjectArchivalRunActivityParams) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *McmObjectArchivalRunActivityParams) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### SetSnapshotIdNil

`func (o *McmObjectArchivalRunActivityParams) SetSnapshotIdNil(b bool)`

 SetSnapshotIdNil sets the value for SnapshotId to be an explicit nil

### UnsetSnapshotId
`func (o *McmObjectArchivalRunActivityParams) UnsetSnapshotId()`

UnsetSnapshotId ensures that no value is present for SnapshotId, not even an explicit nil
### GetPolicyId

`func (o *McmObjectArchivalRunActivityParams) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *McmObjectArchivalRunActivityParams) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *McmObjectArchivalRunActivityParams) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *McmObjectArchivalRunActivityParams) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *McmObjectArchivalRunActivityParams) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *McmObjectArchivalRunActivityParams) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetPolicyName

`func (o *McmObjectArchivalRunActivityParams) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *McmObjectArchivalRunActivityParams) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *McmObjectArchivalRunActivityParams) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *McmObjectArchivalRunActivityParams) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *McmObjectArchivalRunActivityParams) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *McmObjectArchivalRunActivityParams) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetStartTimeUsecs

`func (o *McmObjectArchivalRunActivityParams) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *McmObjectArchivalRunActivityParams) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *McmObjectArchivalRunActivityParams) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *McmObjectArchivalRunActivityParams) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *McmObjectArchivalRunActivityParams) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *McmObjectArchivalRunActivityParams) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetEndTimeUsecs

`func (o *McmObjectArchivalRunActivityParams) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *McmObjectArchivalRunActivityParams) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *McmObjectArchivalRunActivityParams) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *McmObjectArchivalRunActivityParams) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *McmObjectArchivalRunActivityParams) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *McmObjectArchivalRunActivityParams) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetStatus

`func (o *McmObjectArchivalRunActivityParams) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *McmObjectArchivalRunActivityParams) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *McmObjectArchivalRunActivityParams) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *McmObjectArchivalRunActivityParams) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *McmObjectArchivalRunActivityParams) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *McmObjectArchivalRunActivityParams) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetProgressTaskId

`func (o *McmObjectArchivalRunActivityParams) GetProgressTaskId() string`

GetProgressTaskId returns the ProgressTaskId field if non-nil, zero value otherwise.

### GetProgressTaskIdOk

`func (o *McmObjectArchivalRunActivityParams) GetProgressTaskIdOk() (*string, bool)`

GetProgressTaskIdOk returns a tuple with the ProgressTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressTaskId

`func (o *McmObjectArchivalRunActivityParams) SetProgressTaskId(v string)`

SetProgressTaskId sets ProgressTaskId field to given value.

### HasProgressTaskId

`func (o *McmObjectArchivalRunActivityParams) HasProgressTaskId() bool`

HasProgressTaskId returns a boolean if a field has been set.

### SetProgressTaskIdNil

`func (o *McmObjectArchivalRunActivityParams) SetProgressTaskIdNil(b bool)`

 SetProgressTaskIdNil sets the value for ProgressTaskId to be an explicit nil

### UnsetProgressTaskId
`func (o *McmObjectArchivalRunActivityParams) UnsetProgressTaskId()`

UnsetProgressTaskId ensures that no value is present for ProgressTaskId, not even an explicit nil
### GetArchivalTargetId

`func (o *McmObjectArchivalRunActivityParams) GetArchivalTargetId() int64`

GetArchivalTargetId returns the ArchivalTargetId field if non-nil, zero value otherwise.

### GetArchivalTargetIdOk

`func (o *McmObjectArchivalRunActivityParams) GetArchivalTargetIdOk() (*int64, bool)`

GetArchivalTargetIdOk returns a tuple with the ArchivalTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalTargetId

`func (o *McmObjectArchivalRunActivityParams) SetArchivalTargetId(v int64)`

SetArchivalTargetId sets ArchivalTargetId field to given value.

### HasArchivalTargetId

`func (o *McmObjectArchivalRunActivityParams) HasArchivalTargetId() bool`

HasArchivalTargetId returns a boolean if a field has been set.

### SetArchivalTargetIdNil

`func (o *McmObjectArchivalRunActivityParams) SetArchivalTargetIdNil(b bool)`

 SetArchivalTargetIdNil sets the value for ArchivalTargetId to be an explicit nil

### UnsetArchivalTargetId
`func (o *McmObjectArchivalRunActivityParams) UnsetArchivalTargetId()`

UnsetArchivalTargetId ensures that no value is present for ArchivalTargetId, not even an explicit nil
### GetArchivalTargetName

`func (o *McmObjectArchivalRunActivityParams) GetArchivalTargetName() string`

GetArchivalTargetName returns the ArchivalTargetName field if non-nil, zero value otherwise.

### GetArchivalTargetNameOk

`func (o *McmObjectArchivalRunActivityParams) GetArchivalTargetNameOk() (*string, bool)`

GetArchivalTargetNameOk returns a tuple with the ArchivalTargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalTargetName

`func (o *McmObjectArchivalRunActivityParams) SetArchivalTargetName(v string)`

SetArchivalTargetName sets ArchivalTargetName field to given value.

### HasArchivalTargetName

`func (o *McmObjectArchivalRunActivityParams) HasArchivalTargetName() bool`

HasArchivalTargetName returns a boolean if a field has been set.

### SetArchivalTargetNameNil

`func (o *McmObjectArchivalRunActivityParams) SetArchivalTargetNameNil(b bool)`

 SetArchivalTargetNameNil sets the value for ArchivalTargetName to be an explicit nil

### UnsetArchivalTargetName
`func (o *McmObjectArchivalRunActivityParams) UnsetArchivalTargetName()`

UnsetArchivalTargetName ensures that no value is present for ArchivalTargetName, not even an explicit nil
### GetProtectionEnvironmentType

`func (o *McmObjectArchivalRunActivityParams) GetProtectionEnvironmentType() string`

GetProtectionEnvironmentType returns the ProtectionEnvironmentType field if non-nil, zero value otherwise.

### GetProtectionEnvironmentTypeOk

`func (o *McmObjectArchivalRunActivityParams) GetProtectionEnvironmentTypeOk() (*string, bool)`

GetProtectionEnvironmentTypeOk returns a tuple with the ProtectionEnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionEnvironmentType

`func (o *McmObjectArchivalRunActivityParams) SetProtectionEnvironmentType(v string)`

SetProtectionEnvironmentType sets ProtectionEnvironmentType field to given value.

### HasProtectionEnvironmentType

`func (o *McmObjectArchivalRunActivityParams) HasProtectionEnvironmentType() bool`

HasProtectionEnvironmentType returns a boolean if a field has been set.

### SetProtectionEnvironmentTypeNil

`func (o *McmObjectArchivalRunActivityParams) SetProtectionEnvironmentTypeNil(b bool)`

 SetProtectionEnvironmentTypeNil sets the value for ProtectionEnvironmentType to be an explicit nil

### UnsetProtectionEnvironmentType
`func (o *McmObjectArchivalRunActivityParams) UnsetProtectionEnvironmentType()`

UnsetProtectionEnvironmentType ensures that no value is present for ProtectionEnvironmentType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


