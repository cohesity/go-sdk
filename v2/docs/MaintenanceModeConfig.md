# MaintenanceModeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationTimeIntervals** | Pointer to [**[]TimeRangeUsecs**](TimeRangeUsecs.md) | Specifies the absolute intervals where the maintenance schedule is valid, i.e. maintenance_shedule is considered only for these time ranges. (For example, if there is one time range with [now_usecs, now_usecs + 10 days], the action will be done during the maintenance_schedule for the next 10 days.)The start time must be specified. The end time can be -1 which would denote an indefinite maintenance mode. | [optional] 
**MaintenanceSchedule** | Pointer to [**Schedule**](Schedule.md) |  | [optional] 
**UserMessage** | Pointer to **string** | User provided message associated with this maintenance mode. | [optional] 
**WorkflowInterventionSpecList** | Pointer to [**[]WorkflowInterventionSpec**](WorkflowInterventionSpec.md) | Specifies the type of intervention for different workflows when the source goes into maintenance mode. | [optional] 

## Methods

### NewMaintenanceModeConfig

`func NewMaintenanceModeConfig() *MaintenanceModeConfig`

NewMaintenanceModeConfig instantiates a new MaintenanceModeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceModeConfigWithDefaults

`func NewMaintenanceModeConfigWithDefaults() *MaintenanceModeConfig`

NewMaintenanceModeConfigWithDefaults instantiates a new MaintenanceModeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationTimeIntervals

`func (o *MaintenanceModeConfig) GetActivationTimeIntervals() []TimeRangeUsecs`

GetActivationTimeIntervals returns the ActivationTimeIntervals field if non-nil, zero value otherwise.

### GetActivationTimeIntervalsOk

`func (o *MaintenanceModeConfig) GetActivationTimeIntervalsOk() (*[]TimeRangeUsecs, bool)`

GetActivationTimeIntervalsOk returns a tuple with the ActivationTimeIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationTimeIntervals

`func (o *MaintenanceModeConfig) SetActivationTimeIntervals(v []TimeRangeUsecs)`

SetActivationTimeIntervals sets ActivationTimeIntervals field to given value.

### HasActivationTimeIntervals

`func (o *MaintenanceModeConfig) HasActivationTimeIntervals() bool`

HasActivationTimeIntervals returns a boolean if a field has been set.

### GetMaintenanceSchedule

`func (o *MaintenanceModeConfig) GetMaintenanceSchedule() Schedule`

GetMaintenanceSchedule returns the MaintenanceSchedule field if non-nil, zero value otherwise.

### GetMaintenanceScheduleOk

`func (o *MaintenanceModeConfig) GetMaintenanceScheduleOk() (*Schedule, bool)`

GetMaintenanceScheduleOk returns a tuple with the MaintenanceSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceSchedule

`func (o *MaintenanceModeConfig) SetMaintenanceSchedule(v Schedule)`

SetMaintenanceSchedule sets MaintenanceSchedule field to given value.

### HasMaintenanceSchedule

`func (o *MaintenanceModeConfig) HasMaintenanceSchedule() bool`

HasMaintenanceSchedule returns a boolean if a field has been set.

### GetUserMessage

`func (o *MaintenanceModeConfig) GetUserMessage() string`

GetUserMessage returns the UserMessage field if non-nil, zero value otherwise.

### GetUserMessageOk

`func (o *MaintenanceModeConfig) GetUserMessageOk() (*string, bool)`

GetUserMessageOk returns a tuple with the UserMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMessage

`func (o *MaintenanceModeConfig) SetUserMessage(v string)`

SetUserMessage sets UserMessage field to given value.

### HasUserMessage

`func (o *MaintenanceModeConfig) HasUserMessage() bool`

HasUserMessage returns a boolean if a field has been set.

### GetWorkflowInterventionSpecList

`func (o *MaintenanceModeConfig) GetWorkflowInterventionSpecList() []WorkflowInterventionSpec`

GetWorkflowInterventionSpecList returns the WorkflowInterventionSpecList field if non-nil, zero value otherwise.

### GetWorkflowInterventionSpecListOk

`func (o *MaintenanceModeConfig) GetWorkflowInterventionSpecListOk() (*[]WorkflowInterventionSpec, bool)`

GetWorkflowInterventionSpecListOk returns a tuple with the WorkflowInterventionSpecList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInterventionSpecList

`func (o *MaintenanceModeConfig) SetWorkflowInterventionSpecList(v []WorkflowInterventionSpec)`

SetWorkflowInterventionSpecList sets WorkflowInterventionSpecList field to given value.

### HasWorkflowInterventionSpecList

`func (o *MaintenanceModeConfig) HasWorkflowInterventionSpecList() bool`

HasWorkflowInterventionSpecList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


