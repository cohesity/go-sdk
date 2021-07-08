# PerformRestoreJobStateProtoRestoreTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to [**RestoreObject**](RestoreObject.md) |  | [optional] 
**ObjectProgressMonitorTaskPath** | Pointer to **NullableString** | The relative task path of the progress monitor for the restore of the above &#39;object&#39;. Please note that this field will be set only after progress monitor is created for this restore job. | [optional] 
**TaskId** | Pointer to **NullableInt64** | Id of the task tracking the restore of the above &#39;object&#39;. | [optional] 

## Methods

### NewPerformRestoreJobStateProtoRestoreTask

`func NewPerformRestoreJobStateProtoRestoreTask() *PerformRestoreJobStateProtoRestoreTask`

NewPerformRestoreJobStateProtoRestoreTask instantiates a new PerformRestoreJobStateProtoRestoreTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformRestoreJobStateProtoRestoreTaskWithDefaults

`func NewPerformRestoreJobStateProtoRestoreTaskWithDefaults() *PerformRestoreJobStateProtoRestoreTask`

NewPerformRestoreJobStateProtoRestoreTaskWithDefaults instantiates a new PerformRestoreJobStateProtoRestoreTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *PerformRestoreJobStateProtoRestoreTask) GetObject() RestoreObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PerformRestoreJobStateProtoRestoreTask) GetObjectOk() (*RestoreObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PerformRestoreJobStateProtoRestoreTask) SetObject(v RestoreObject)`

SetObject sets Object field to given value.

### HasObject

`func (o *PerformRestoreJobStateProtoRestoreTask) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetObjectProgressMonitorTaskPath

`func (o *PerformRestoreJobStateProtoRestoreTask) GetObjectProgressMonitorTaskPath() string`

GetObjectProgressMonitorTaskPath returns the ObjectProgressMonitorTaskPath field if non-nil, zero value otherwise.

### GetObjectProgressMonitorTaskPathOk

`func (o *PerformRestoreJobStateProtoRestoreTask) GetObjectProgressMonitorTaskPathOk() (*string, bool)`

GetObjectProgressMonitorTaskPathOk returns a tuple with the ObjectProgressMonitorTaskPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectProgressMonitorTaskPath

`func (o *PerformRestoreJobStateProtoRestoreTask) SetObjectProgressMonitorTaskPath(v string)`

SetObjectProgressMonitorTaskPath sets ObjectProgressMonitorTaskPath field to given value.

### HasObjectProgressMonitorTaskPath

`func (o *PerformRestoreJobStateProtoRestoreTask) HasObjectProgressMonitorTaskPath() bool`

HasObjectProgressMonitorTaskPath returns a boolean if a field has been set.

### SetObjectProgressMonitorTaskPathNil

`func (o *PerformRestoreJobStateProtoRestoreTask) SetObjectProgressMonitorTaskPathNil(b bool)`

 SetObjectProgressMonitorTaskPathNil sets the value for ObjectProgressMonitorTaskPath to be an explicit nil

### UnsetObjectProgressMonitorTaskPath
`func (o *PerformRestoreJobStateProtoRestoreTask) UnsetObjectProgressMonitorTaskPath()`

UnsetObjectProgressMonitorTaskPath ensures that no value is present for ObjectProgressMonitorTaskPath, not even an explicit nil
### GetTaskId

`func (o *PerformRestoreJobStateProtoRestoreTask) GetTaskId() int64`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *PerformRestoreJobStateProtoRestoreTask) GetTaskIdOk() (*int64, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *PerformRestoreJobStateProtoRestoreTask) SetTaskId(v int64)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *PerformRestoreJobStateProtoRestoreTask) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *PerformRestoreJobStateProtoRestoreTask) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *PerformRestoreJobStateProtoRestoreTask) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


