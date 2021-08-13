# AgentUpgradeTaskState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies the name of the task. | [optional] 
**Description** | Pointer to **NullableString** | Specifies the description of the task. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the ID of the task. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the time, as a Unix epoch timestamp in microseconds, when the task started execution. | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the time when the upgrade task completed execution as a Unix epoch Timestamp (in microseconds). | [optional] 
**Status** | Pointer to **NullableString** | Specifies the status of the task.&lt;br&gt; &#39;Scheduled&#39; indicates that the upgrade task is yet to start.&lt;br&gt; &#39;Running&#39; indicates that the upgrade task has started execution.&lt;br&gt; &#39;Succeeded&#39; indicates that the upgrade task completed without an error.&lt;br&gt; &#39;Failed&#39; indicates that upgrade has failed for all agents. &#39;PartiallyFailed&#39; indicates that upgrade has failed for some agents. | [optional] 
**AgentIDs** | Pointer to **[]int64** | Specifies the agents upgraded in the task. | [optional] 
**ClusterVersion** | Pointer to **NullableString** | Specifies the version to which agents are upgraded. | [optional] 
**Type** | Pointer to **NullableString** | Specifes the type of task.&lt;br&gt; &#39;Auto&#39; indicates an auto agent upgrade task which is started after a cluster upgrade.&lt;br&gt; &#39;Manual&#39; indicates a schedule based agent upgrade task.&lt;br&gt; &#39;Retry&#39; indicates an agent upgrade task which was retried. | [optional] 
**ScheduleTimeUsecs** | Pointer to **NullableInt64** | Specifies the time when the task should start execution as a Unix epoch Timestamp (in microseconds). If no schedule is specified, the task will start immediately. | [optional] 
**ScheduleEndTimeUsecs** | Pointer to **NullableInt64** | Specifies the time before which the upgrade task should start execution as a Unix epoch Timestamp (in microseconds). If this is not specified the task will start anytime after scheduleTimeUsecs. | [optional] 
**RetriedTaskID** | Pointer to **NullableInt64** | Specifies ID of a task which was retried if type is &#39;Retry&#39;. | [optional] 
**IsRetryable** | Pointer to **NullableBool** | Specifies if a task can be retried. | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 
**Agents** | Pointer to [**[]AgentUpgradeInfoObject**](AgentUpgradeInfoObject.md) | Specifies the upgrade information for each agent. | [optional] 

## Methods

### NewAgentUpgradeTaskState

`func NewAgentUpgradeTaskState() *AgentUpgradeTaskState`

NewAgentUpgradeTaskState instantiates a new AgentUpgradeTaskState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentUpgradeTaskStateWithDefaults

`func NewAgentUpgradeTaskStateWithDefaults() *AgentUpgradeTaskState`

NewAgentUpgradeTaskStateWithDefaults instantiates a new AgentUpgradeTaskState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AgentUpgradeTaskState) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentUpgradeTaskState) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentUpgradeTaskState) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentUpgradeTaskState) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AgentUpgradeTaskState) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AgentUpgradeTaskState) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *AgentUpgradeTaskState) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentUpgradeTaskState) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentUpgradeTaskState) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentUpgradeTaskState) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AgentUpgradeTaskState) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AgentUpgradeTaskState) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *AgentUpgradeTaskState) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentUpgradeTaskState) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentUpgradeTaskState) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AgentUpgradeTaskState) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AgentUpgradeTaskState) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AgentUpgradeTaskState) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetStartTimeUsecs

`func (o *AgentUpgradeTaskState) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *AgentUpgradeTaskState) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *AgentUpgradeTaskState) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *AgentUpgradeTaskState) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *AgentUpgradeTaskState) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *AgentUpgradeTaskState) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetEndTimeUsecs

`func (o *AgentUpgradeTaskState) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *AgentUpgradeTaskState) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *AgentUpgradeTaskState) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *AgentUpgradeTaskState) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *AgentUpgradeTaskState) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *AgentUpgradeTaskState) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetStatus

`func (o *AgentUpgradeTaskState) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentUpgradeTaskState) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentUpgradeTaskState) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentUpgradeTaskState) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AgentUpgradeTaskState) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AgentUpgradeTaskState) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetAgentIDs

`func (o *AgentUpgradeTaskState) GetAgentIDs() []int64`

GetAgentIDs returns the AgentIDs field if non-nil, zero value otherwise.

### GetAgentIDsOk

`func (o *AgentUpgradeTaskState) GetAgentIDsOk() (*[]int64, bool)`

GetAgentIDsOk returns a tuple with the AgentIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentIDs

`func (o *AgentUpgradeTaskState) SetAgentIDs(v []int64)`

SetAgentIDs sets AgentIDs field to given value.

### HasAgentIDs

`func (o *AgentUpgradeTaskState) HasAgentIDs() bool`

HasAgentIDs returns a boolean if a field has been set.

### SetAgentIDsNil

`func (o *AgentUpgradeTaskState) SetAgentIDsNil(b bool)`

 SetAgentIDsNil sets the value for AgentIDs to be an explicit nil

### UnsetAgentIDs
`func (o *AgentUpgradeTaskState) UnsetAgentIDs()`

UnsetAgentIDs ensures that no value is present for AgentIDs, not even an explicit nil
### GetClusterVersion

`func (o *AgentUpgradeTaskState) GetClusterVersion() string`

GetClusterVersion returns the ClusterVersion field if non-nil, zero value otherwise.

### GetClusterVersionOk

`func (o *AgentUpgradeTaskState) GetClusterVersionOk() (*string, bool)`

GetClusterVersionOk returns a tuple with the ClusterVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterVersion

`func (o *AgentUpgradeTaskState) SetClusterVersion(v string)`

SetClusterVersion sets ClusterVersion field to given value.

### HasClusterVersion

`func (o *AgentUpgradeTaskState) HasClusterVersion() bool`

HasClusterVersion returns a boolean if a field has been set.

### SetClusterVersionNil

`func (o *AgentUpgradeTaskState) SetClusterVersionNil(b bool)`

 SetClusterVersionNil sets the value for ClusterVersion to be an explicit nil

### UnsetClusterVersion
`func (o *AgentUpgradeTaskState) UnsetClusterVersion()`

UnsetClusterVersion ensures that no value is present for ClusterVersion, not even an explicit nil
### GetType

`func (o *AgentUpgradeTaskState) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgentUpgradeTaskState) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgentUpgradeTaskState) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AgentUpgradeTaskState) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *AgentUpgradeTaskState) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AgentUpgradeTaskState) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetScheduleTimeUsecs

`func (o *AgentUpgradeTaskState) GetScheduleTimeUsecs() int64`

GetScheduleTimeUsecs returns the ScheduleTimeUsecs field if non-nil, zero value otherwise.

### GetScheduleTimeUsecsOk

`func (o *AgentUpgradeTaskState) GetScheduleTimeUsecsOk() (*int64, bool)`

GetScheduleTimeUsecsOk returns a tuple with the ScheduleTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTimeUsecs

`func (o *AgentUpgradeTaskState) SetScheduleTimeUsecs(v int64)`

SetScheduleTimeUsecs sets ScheduleTimeUsecs field to given value.

### HasScheduleTimeUsecs

`func (o *AgentUpgradeTaskState) HasScheduleTimeUsecs() bool`

HasScheduleTimeUsecs returns a boolean if a field has been set.

### SetScheduleTimeUsecsNil

`func (o *AgentUpgradeTaskState) SetScheduleTimeUsecsNil(b bool)`

 SetScheduleTimeUsecsNil sets the value for ScheduleTimeUsecs to be an explicit nil

### UnsetScheduleTimeUsecs
`func (o *AgentUpgradeTaskState) UnsetScheduleTimeUsecs()`

UnsetScheduleTimeUsecs ensures that no value is present for ScheduleTimeUsecs, not even an explicit nil
### GetScheduleEndTimeUsecs

`func (o *AgentUpgradeTaskState) GetScheduleEndTimeUsecs() int64`

GetScheduleEndTimeUsecs returns the ScheduleEndTimeUsecs field if non-nil, zero value otherwise.

### GetScheduleEndTimeUsecsOk

`func (o *AgentUpgradeTaskState) GetScheduleEndTimeUsecsOk() (*int64, bool)`

GetScheduleEndTimeUsecsOk returns a tuple with the ScheduleEndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleEndTimeUsecs

`func (o *AgentUpgradeTaskState) SetScheduleEndTimeUsecs(v int64)`

SetScheduleEndTimeUsecs sets ScheduleEndTimeUsecs field to given value.

### HasScheduleEndTimeUsecs

`func (o *AgentUpgradeTaskState) HasScheduleEndTimeUsecs() bool`

HasScheduleEndTimeUsecs returns a boolean if a field has been set.

### SetScheduleEndTimeUsecsNil

`func (o *AgentUpgradeTaskState) SetScheduleEndTimeUsecsNil(b bool)`

 SetScheduleEndTimeUsecsNil sets the value for ScheduleEndTimeUsecs to be an explicit nil

### UnsetScheduleEndTimeUsecs
`func (o *AgentUpgradeTaskState) UnsetScheduleEndTimeUsecs()`

UnsetScheduleEndTimeUsecs ensures that no value is present for ScheduleEndTimeUsecs, not even an explicit nil
### GetRetriedTaskID

`func (o *AgentUpgradeTaskState) GetRetriedTaskID() int64`

GetRetriedTaskID returns the RetriedTaskID field if non-nil, zero value otherwise.

### GetRetriedTaskIDOk

`func (o *AgentUpgradeTaskState) GetRetriedTaskIDOk() (*int64, bool)`

GetRetriedTaskIDOk returns a tuple with the RetriedTaskID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetriedTaskID

`func (o *AgentUpgradeTaskState) SetRetriedTaskID(v int64)`

SetRetriedTaskID sets RetriedTaskID field to given value.

### HasRetriedTaskID

`func (o *AgentUpgradeTaskState) HasRetriedTaskID() bool`

HasRetriedTaskID returns a boolean if a field has been set.

### SetRetriedTaskIDNil

`func (o *AgentUpgradeTaskState) SetRetriedTaskIDNil(b bool)`

 SetRetriedTaskIDNil sets the value for RetriedTaskID to be an explicit nil

### UnsetRetriedTaskID
`func (o *AgentUpgradeTaskState) UnsetRetriedTaskID()`

UnsetRetriedTaskID ensures that no value is present for RetriedTaskID, not even an explicit nil
### GetIsRetryable

`func (o *AgentUpgradeTaskState) GetIsRetryable() bool`

GetIsRetryable returns the IsRetryable field if non-nil, zero value otherwise.

### GetIsRetryableOk

`func (o *AgentUpgradeTaskState) GetIsRetryableOk() (*bool, bool)`

GetIsRetryableOk returns a tuple with the IsRetryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRetryable

`func (o *AgentUpgradeTaskState) SetIsRetryable(v bool)`

SetIsRetryable sets IsRetryable field to given value.

### HasIsRetryable

`func (o *AgentUpgradeTaskState) HasIsRetryable() bool`

HasIsRetryable returns a boolean if a field has been set.

### SetIsRetryableNil

`func (o *AgentUpgradeTaskState) SetIsRetryableNil(b bool)`

 SetIsRetryableNil sets the value for IsRetryable to be an explicit nil

### UnsetIsRetryable
`func (o *AgentUpgradeTaskState) UnsetIsRetryable()`

UnsetIsRetryable ensures that no value is present for IsRetryable, not even an explicit nil
### GetError

`func (o *AgentUpgradeTaskState) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AgentUpgradeTaskState) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AgentUpgradeTaskState) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *AgentUpgradeTaskState) HasError() bool`

HasError returns a boolean if a field has been set.

### GetAgents

`func (o *AgentUpgradeTaskState) GetAgents() []AgentUpgradeInfoObject`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *AgentUpgradeTaskState) GetAgentsOk() (*[]AgentUpgradeInfoObject, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *AgentUpgradeTaskState) SetAgents(v []AgentUpgradeInfoObject)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *AgentUpgradeTaskState) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### SetAgentsNil

`func (o *AgentUpgradeTaskState) SetAgentsNil(b bool)`

 SetAgentsNil sets the value for Agents to be an explicit nil

### UnsetAgents
`func (o *AgentUpgradeTaskState) UnsetAgents()`

UnsetAgents ensures that no value is present for Agents, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


