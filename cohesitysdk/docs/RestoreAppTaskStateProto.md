# RestoreAppTaskStateProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppRestoreProgressMonitorSubtaskPath** | Pointer to **NullableString** | The Pulse task path to the application restore task sub tree. If the application restore has to wait on other tasks (for example, a SQL db restore may wait for a tail log backup or a VM restore), then this would represent a sub-tree of &#39;progress_monitor_task_path&#39; in PerformRestoreTaskStateProto. | [optional] 
**ChildRestoreAppParamsVec** | Pointer to [**[]RestoreAppParams**](RestoreAppParams.md) | This has list of the restore app params for all the child restore tasks. This is populated iff the &#39;is_parent_task&#39; is set to true. | [optional] 
**LastFinishedLogBackupStartTimeUsecs** | Pointer to **NullableInt64** | The start time of the last finished log backup run. For SQL application, this is set iff we need to take a tail log backup. | [optional] 
**RestoreAppParams** | Pointer to [**RestoreAppParams**](RestoreAppParams.md) |  | [optional] 

## Methods

### NewRestoreAppTaskStateProto

`func NewRestoreAppTaskStateProto() *RestoreAppTaskStateProto`

NewRestoreAppTaskStateProto instantiates a new RestoreAppTaskStateProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreAppTaskStateProtoWithDefaults

`func NewRestoreAppTaskStateProtoWithDefaults() *RestoreAppTaskStateProto`

NewRestoreAppTaskStateProtoWithDefaults instantiates a new RestoreAppTaskStateProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppRestoreProgressMonitorSubtaskPath

`func (o *RestoreAppTaskStateProto) GetAppRestoreProgressMonitorSubtaskPath() string`

GetAppRestoreProgressMonitorSubtaskPath returns the AppRestoreProgressMonitorSubtaskPath field if non-nil, zero value otherwise.

### GetAppRestoreProgressMonitorSubtaskPathOk

`func (o *RestoreAppTaskStateProto) GetAppRestoreProgressMonitorSubtaskPathOk() (*string, bool)`

GetAppRestoreProgressMonitorSubtaskPathOk returns a tuple with the AppRestoreProgressMonitorSubtaskPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRestoreProgressMonitorSubtaskPath

`func (o *RestoreAppTaskStateProto) SetAppRestoreProgressMonitorSubtaskPath(v string)`

SetAppRestoreProgressMonitorSubtaskPath sets AppRestoreProgressMonitorSubtaskPath field to given value.

### HasAppRestoreProgressMonitorSubtaskPath

`func (o *RestoreAppTaskStateProto) HasAppRestoreProgressMonitorSubtaskPath() bool`

HasAppRestoreProgressMonitorSubtaskPath returns a boolean if a field has been set.

### SetAppRestoreProgressMonitorSubtaskPathNil

`func (o *RestoreAppTaskStateProto) SetAppRestoreProgressMonitorSubtaskPathNil(b bool)`

 SetAppRestoreProgressMonitorSubtaskPathNil sets the value for AppRestoreProgressMonitorSubtaskPath to be an explicit nil

### UnsetAppRestoreProgressMonitorSubtaskPath
`func (o *RestoreAppTaskStateProto) UnsetAppRestoreProgressMonitorSubtaskPath()`

UnsetAppRestoreProgressMonitorSubtaskPath ensures that no value is present for AppRestoreProgressMonitorSubtaskPath, not even an explicit nil
### GetChildRestoreAppParamsVec

`func (o *RestoreAppTaskStateProto) GetChildRestoreAppParamsVec() []RestoreAppParams`

GetChildRestoreAppParamsVec returns the ChildRestoreAppParamsVec field if non-nil, zero value otherwise.

### GetChildRestoreAppParamsVecOk

`func (o *RestoreAppTaskStateProto) GetChildRestoreAppParamsVecOk() (*[]RestoreAppParams, bool)`

GetChildRestoreAppParamsVecOk returns a tuple with the ChildRestoreAppParamsVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildRestoreAppParamsVec

`func (o *RestoreAppTaskStateProto) SetChildRestoreAppParamsVec(v []RestoreAppParams)`

SetChildRestoreAppParamsVec sets ChildRestoreAppParamsVec field to given value.

### HasChildRestoreAppParamsVec

`func (o *RestoreAppTaskStateProto) HasChildRestoreAppParamsVec() bool`

HasChildRestoreAppParamsVec returns a boolean if a field has been set.

### SetChildRestoreAppParamsVecNil

`func (o *RestoreAppTaskStateProto) SetChildRestoreAppParamsVecNil(b bool)`

 SetChildRestoreAppParamsVecNil sets the value for ChildRestoreAppParamsVec to be an explicit nil

### UnsetChildRestoreAppParamsVec
`func (o *RestoreAppTaskStateProto) UnsetChildRestoreAppParamsVec()`

UnsetChildRestoreAppParamsVec ensures that no value is present for ChildRestoreAppParamsVec, not even an explicit nil
### GetLastFinishedLogBackupStartTimeUsecs

`func (o *RestoreAppTaskStateProto) GetLastFinishedLogBackupStartTimeUsecs() int64`

GetLastFinishedLogBackupStartTimeUsecs returns the LastFinishedLogBackupStartTimeUsecs field if non-nil, zero value otherwise.

### GetLastFinishedLogBackupStartTimeUsecsOk

`func (o *RestoreAppTaskStateProto) GetLastFinishedLogBackupStartTimeUsecsOk() (*int64, bool)`

GetLastFinishedLogBackupStartTimeUsecsOk returns a tuple with the LastFinishedLogBackupStartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFinishedLogBackupStartTimeUsecs

`func (o *RestoreAppTaskStateProto) SetLastFinishedLogBackupStartTimeUsecs(v int64)`

SetLastFinishedLogBackupStartTimeUsecs sets LastFinishedLogBackupStartTimeUsecs field to given value.

### HasLastFinishedLogBackupStartTimeUsecs

`func (o *RestoreAppTaskStateProto) HasLastFinishedLogBackupStartTimeUsecs() bool`

HasLastFinishedLogBackupStartTimeUsecs returns a boolean if a field has been set.

### SetLastFinishedLogBackupStartTimeUsecsNil

`func (o *RestoreAppTaskStateProto) SetLastFinishedLogBackupStartTimeUsecsNil(b bool)`

 SetLastFinishedLogBackupStartTimeUsecsNil sets the value for LastFinishedLogBackupStartTimeUsecs to be an explicit nil

### UnsetLastFinishedLogBackupStartTimeUsecs
`func (o *RestoreAppTaskStateProto) UnsetLastFinishedLogBackupStartTimeUsecs()`

UnsetLastFinishedLogBackupStartTimeUsecs ensures that no value is present for LastFinishedLogBackupStartTimeUsecs, not even an explicit nil
### GetRestoreAppParams

`func (o *RestoreAppTaskStateProto) GetRestoreAppParams() RestoreAppParams`

GetRestoreAppParams returns the RestoreAppParams field if non-nil, zero value otherwise.

### GetRestoreAppParamsOk

`func (o *RestoreAppTaskStateProto) GetRestoreAppParamsOk() (*RestoreAppParams, bool)`

GetRestoreAppParamsOk returns a tuple with the RestoreAppParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAppParams

`func (o *RestoreAppTaskStateProto) SetRestoreAppParams(v RestoreAppParams)`

SetRestoreAppParams sets RestoreAppParams field to given value.

### HasRestoreAppParams

`func (o *RestoreAppTaskStateProto) HasRestoreAppParams() bool`

HasRestoreAppParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


