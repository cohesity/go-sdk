# TaskLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddTaskId** | Pointer to **NullableString** | Specifies the id of adding task. | [optional] 
**EditTaskId** | Pointer to **[]string** | Specifies the id of editing task. | [optional] 

## Methods

### NewTaskLogs

`func NewTaskLogs() *TaskLogs`

NewTaskLogs instantiates a new TaskLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskLogsWithDefaults

`func NewTaskLogsWithDefaults() *TaskLogs`

NewTaskLogsWithDefaults instantiates a new TaskLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddTaskId

`func (o *TaskLogs) GetAddTaskId() string`

GetAddTaskId returns the AddTaskId field if non-nil, zero value otherwise.

### GetAddTaskIdOk

`func (o *TaskLogs) GetAddTaskIdOk() (*string, bool)`

GetAddTaskIdOk returns a tuple with the AddTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTaskId

`func (o *TaskLogs) SetAddTaskId(v string)`

SetAddTaskId sets AddTaskId field to given value.

### HasAddTaskId

`func (o *TaskLogs) HasAddTaskId() bool`

HasAddTaskId returns a boolean if a field has been set.

### SetAddTaskIdNil

`func (o *TaskLogs) SetAddTaskIdNil(b bool)`

 SetAddTaskIdNil sets the value for AddTaskId to be an explicit nil

### UnsetAddTaskId
`func (o *TaskLogs) UnsetAddTaskId()`

UnsetAddTaskId ensures that no value is present for AddTaskId, not even an explicit nil
### GetEditTaskId

`func (o *TaskLogs) GetEditTaskId() []string`

GetEditTaskId returns the EditTaskId field if non-nil, zero value otherwise.

### GetEditTaskIdOk

`func (o *TaskLogs) GetEditTaskIdOk() (*[]string, bool)`

GetEditTaskIdOk returns a tuple with the EditTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditTaskId

`func (o *TaskLogs) SetEditTaskId(v []string)`

SetEditTaskId sets EditTaskId field to given value.

### HasEditTaskId

`func (o *TaskLogs) HasEditTaskId() bool`

HasEditTaskId returns a boolean if a field has been set.

### SetEditTaskIdNil

`func (o *TaskLogs) SetEditTaskIdNil(b bool)`

 SetEditTaskIdNil sets the value for EditTaskId to be an explicit nil

### UnsetEditTaskId
`func (o *TaskLogs) UnsetEditTaskId()`

UnsetEditTaskId ensures that no value is present for EditTaskId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


