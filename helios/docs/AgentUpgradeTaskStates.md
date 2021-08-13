# AgentUpgradeTaskStates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tasks** | Pointer to [**[]AgentUpgradeTaskState**](AgentUpgradeTaskState.md) | Specifies the list of agent upgrade tasks. | [optional] 

## Methods

### NewAgentUpgradeTaskStates

`func NewAgentUpgradeTaskStates() *AgentUpgradeTaskStates`

NewAgentUpgradeTaskStates instantiates a new AgentUpgradeTaskStates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentUpgradeTaskStatesWithDefaults

`func NewAgentUpgradeTaskStatesWithDefaults() *AgentUpgradeTaskStates`

NewAgentUpgradeTaskStatesWithDefaults instantiates a new AgentUpgradeTaskStates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTasks

`func (o *AgentUpgradeTaskStates) GetTasks() []AgentUpgradeTaskState`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *AgentUpgradeTaskStates) GetTasksOk() (*[]AgentUpgradeTaskState, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *AgentUpgradeTaskStates) SetTasks(v []AgentUpgradeTaskState)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *AgentUpgradeTaskStates) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### SetTasksNil

`func (o *AgentUpgradeTaskStates) SetTasksNil(b bool)`

 SetTasksNil sets the value for Tasks to be an explicit nil

### UnsetTasks
`func (o *AgentUpgradeTaskStates) UnsetTasks()`

UnsetTasks ensures that no value is present for Tasks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


