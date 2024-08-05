# CreateUpgradeTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentIDs** | Pointer to **[]int64** | Specifies agent IDs to be upgraded in the task. | [optional] 
**Description** | Pointer to **NullableString** | Specifies the description of the task. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the task. | [optional] 
**RetryTaskId** | Pointer to **NullableInt64** | Specifies the task that needs to be retried. | [optional] 
**ScheduleEndTimeUsecs** | Pointer to **NullableInt64** | Specifies the time before which the upgrade task should start execution as a Unix epoch Timestamp (in microseconds). If this is not specified the task will start anytime after scheduleTimeUsecs. | [optional] 
**ScheduleTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the task specified by the user as a Unix epoch Timestamp (in microseconds). If no schedule is specified, the agents will be upgraded immediately. | [optional] 

## Methods

### NewCreateUpgradeTaskRequest

`func NewCreateUpgradeTaskRequest() *CreateUpgradeTaskRequest`

NewCreateUpgradeTaskRequest instantiates a new CreateUpgradeTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpgradeTaskRequestWithDefaults

`func NewCreateUpgradeTaskRequestWithDefaults() *CreateUpgradeTaskRequest`

NewCreateUpgradeTaskRequestWithDefaults instantiates a new CreateUpgradeTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentIDs

`func (o *CreateUpgradeTaskRequest) GetAgentIDs() []int64`

GetAgentIDs returns the AgentIDs field if non-nil, zero value otherwise.

### GetAgentIDsOk

`func (o *CreateUpgradeTaskRequest) GetAgentIDsOk() (*[]int64, bool)`

GetAgentIDsOk returns a tuple with the AgentIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentIDs

`func (o *CreateUpgradeTaskRequest) SetAgentIDs(v []int64)`

SetAgentIDs sets AgentIDs field to given value.

### HasAgentIDs

`func (o *CreateUpgradeTaskRequest) HasAgentIDs() bool`

HasAgentIDs returns a boolean if a field has been set.

### SetAgentIDsNil

`func (o *CreateUpgradeTaskRequest) SetAgentIDsNil(b bool)`

 SetAgentIDsNil sets the value for AgentIDs to be an explicit nil

### UnsetAgentIDs
`func (o *CreateUpgradeTaskRequest) UnsetAgentIDs()`

UnsetAgentIDs ensures that no value is present for AgentIDs, not even an explicit nil
### GetDescription

`func (o *CreateUpgradeTaskRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateUpgradeTaskRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateUpgradeTaskRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateUpgradeTaskRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateUpgradeTaskRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateUpgradeTaskRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *CreateUpgradeTaskRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUpgradeTaskRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUpgradeTaskRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateUpgradeTaskRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateUpgradeTaskRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateUpgradeTaskRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRetryTaskId

`func (o *CreateUpgradeTaskRequest) GetRetryTaskId() int64`

GetRetryTaskId returns the RetryTaskId field if non-nil, zero value otherwise.

### GetRetryTaskIdOk

`func (o *CreateUpgradeTaskRequest) GetRetryTaskIdOk() (*int64, bool)`

GetRetryTaskIdOk returns a tuple with the RetryTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryTaskId

`func (o *CreateUpgradeTaskRequest) SetRetryTaskId(v int64)`

SetRetryTaskId sets RetryTaskId field to given value.

### HasRetryTaskId

`func (o *CreateUpgradeTaskRequest) HasRetryTaskId() bool`

HasRetryTaskId returns a boolean if a field has been set.

### SetRetryTaskIdNil

`func (o *CreateUpgradeTaskRequest) SetRetryTaskIdNil(b bool)`

 SetRetryTaskIdNil sets the value for RetryTaskId to be an explicit nil

### UnsetRetryTaskId
`func (o *CreateUpgradeTaskRequest) UnsetRetryTaskId()`

UnsetRetryTaskId ensures that no value is present for RetryTaskId, not even an explicit nil
### GetScheduleEndTimeUsecs

`func (o *CreateUpgradeTaskRequest) GetScheduleEndTimeUsecs() int64`

GetScheduleEndTimeUsecs returns the ScheduleEndTimeUsecs field if non-nil, zero value otherwise.

### GetScheduleEndTimeUsecsOk

`func (o *CreateUpgradeTaskRequest) GetScheduleEndTimeUsecsOk() (*int64, bool)`

GetScheduleEndTimeUsecsOk returns a tuple with the ScheduleEndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleEndTimeUsecs

`func (o *CreateUpgradeTaskRequest) SetScheduleEndTimeUsecs(v int64)`

SetScheduleEndTimeUsecs sets ScheduleEndTimeUsecs field to given value.

### HasScheduleEndTimeUsecs

`func (o *CreateUpgradeTaskRequest) HasScheduleEndTimeUsecs() bool`

HasScheduleEndTimeUsecs returns a boolean if a field has been set.

### SetScheduleEndTimeUsecsNil

`func (o *CreateUpgradeTaskRequest) SetScheduleEndTimeUsecsNil(b bool)`

 SetScheduleEndTimeUsecsNil sets the value for ScheduleEndTimeUsecs to be an explicit nil

### UnsetScheduleEndTimeUsecs
`func (o *CreateUpgradeTaskRequest) UnsetScheduleEndTimeUsecs()`

UnsetScheduleEndTimeUsecs ensures that no value is present for ScheduleEndTimeUsecs, not even an explicit nil
### GetScheduleTimeUsecs

`func (o *CreateUpgradeTaskRequest) GetScheduleTimeUsecs() int64`

GetScheduleTimeUsecs returns the ScheduleTimeUsecs field if non-nil, zero value otherwise.

### GetScheduleTimeUsecsOk

`func (o *CreateUpgradeTaskRequest) GetScheduleTimeUsecsOk() (*int64, bool)`

GetScheduleTimeUsecsOk returns a tuple with the ScheduleTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTimeUsecs

`func (o *CreateUpgradeTaskRequest) SetScheduleTimeUsecs(v int64)`

SetScheduleTimeUsecs sets ScheduleTimeUsecs field to given value.

### HasScheduleTimeUsecs

`func (o *CreateUpgradeTaskRequest) HasScheduleTimeUsecs() bool`

HasScheduleTimeUsecs returns a boolean if a field has been set.

### SetScheduleTimeUsecsNil

`func (o *CreateUpgradeTaskRequest) SetScheduleTimeUsecsNil(b bool)`

 SetScheduleTimeUsecsNil sets the value for ScheduleTimeUsecs to be an explicit nil

### UnsetScheduleTimeUsecs
`func (o *CreateUpgradeTaskRequest) UnsetScheduleTimeUsecs()`

UnsetScheduleTimeUsecs ensures that no value is present for ScheduleTimeUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


