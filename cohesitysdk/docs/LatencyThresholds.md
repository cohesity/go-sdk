# LatencyThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveTaskMsecs** | Pointer to **NullableInt64** | If the latency of a datastore is above this value, existing backup tasks using the datastore are throttled. | [optional] 
**NewTaskMsecs** | Pointer to **NullableInt64** | If the latency of a datastore is above this value, then new backup tasks using the datastore will not be started. | [optional] 

## Methods

### NewLatencyThresholds

`func NewLatencyThresholds() *LatencyThresholds`

NewLatencyThresholds instantiates a new LatencyThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLatencyThresholdsWithDefaults

`func NewLatencyThresholdsWithDefaults() *LatencyThresholds`

NewLatencyThresholdsWithDefaults instantiates a new LatencyThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveTaskMsecs

`func (o *LatencyThresholds) GetActiveTaskMsecs() int64`

GetActiveTaskMsecs returns the ActiveTaskMsecs field if non-nil, zero value otherwise.

### GetActiveTaskMsecsOk

`func (o *LatencyThresholds) GetActiveTaskMsecsOk() (*int64, bool)`

GetActiveTaskMsecsOk returns a tuple with the ActiveTaskMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTaskMsecs

`func (o *LatencyThresholds) SetActiveTaskMsecs(v int64)`

SetActiveTaskMsecs sets ActiveTaskMsecs field to given value.

### HasActiveTaskMsecs

`func (o *LatencyThresholds) HasActiveTaskMsecs() bool`

HasActiveTaskMsecs returns a boolean if a field has been set.

### SetActiveTaskMsecsNil

`func (o *LatencyThresholds) SetActiveTaskMsecsNil(b bool)`

 SetActiveTaskMsecsNil sets the value for ActiveTaskMsecs to be an explicit nil

### UnsetActiveTaskMsecs
`func (o *LatencyThresholds) UnsetActiveTaskMsecs()`

UnsetActiveTaskMsecs ensures that no value is present for ActiveTaskMsecs, not even an explicit nil
### GetNewTaskMsecs

`func (o *LatencyThresholds) GetNewTaskMsecs() int64`

GetNewTaskMsecs returns the NewTaskMsecs field if non-nil, zero value otherwise.

### GetNewTaskMsecsOk

`func (o *LatencyThresholds) GetNewTaskMsecsOk() (*int64, bool)`

GetNewTaskMsecsOk returns a tuple with the NewTaskMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTaskMsecs

`func (o *LatencyThresholds) SetNewTaskMsecs(v int64)`

SetNewTaskMsecs sets NewTaskMsecs field to given value.

### HasNewTaskMsecs

`func (o *LatencyThresholds) HasNewTaskMsecs() bool`

HasNewTaskMsecs returns a boolean if a field has been set.

### SetNewTaskMsecsNil

`func (o *LatencyThresholds) SetNewTaskMsecsNil(b bool)`

 SetNewTaskMsecsNil sets the value for NewTaskMsecs to be an explicit nil

### UnsetNewTaskMsecs
`func (o *LatencyThresholds) UnsetNewTaskMsecs()`

UnsetNewTaskMsecs ensures that no value is present for NewTaskMsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


