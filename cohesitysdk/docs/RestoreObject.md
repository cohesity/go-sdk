# RestoreObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivalTarget** | Pointer to [**ArchivalTarget**](ArchivalTarget.md) |  | [optional] 
**AttemptNum** | Pointer to **NullableInt32** | The attempt number of the job run to restore from. | [optional] 
**CloudDeployTarget** | Pointer to [**CloudDeployTarget**](CloudDeployTarget.md) |  | [optional] 
**CloudReplicationTarget** | Pointer to [**CloudDeployTarget**](CloudDeployTarget.md) |  | [optional] 
**Entity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**JobId** | Pointer to **NullableInt64** | The job id from which to restore. This is used while communicating with yoda. | [optional] 
**JobInstanceId** | Pointer to **NullableInt64** | Id identifying a specific run to restore from. If this is not specified, and we need to restore from a run, the latest run is used. NOTE: This must be specified for RestoreFiles, RecoverDisks and GetVirtualDisks APIs. | [optional] 
**JobUid** | Pointer to [**UniversalIdProto**](UniversalIdProto.md) |  | [optional] 
**NosqlRecoverParams** | Pointer to [**NoSqlRecoverParams**](NoSqlRecoverParams.md) |  | [optional] 
**ParentSource** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**PointInTimeRestoreTimeUsecs** | Pointer to **NullableInt64** | The time to which the object needs to be restored. If this is not set, then the object will be restored to the full/incremental snapshot. This is applicable only if the object is protected using CDP. | [optional] 
**RestoreAcropolisVmParam** | Pointer to [**RestoreAcropolisVMParam**](RestoreAcropolisVMParam.md) |  | [optional] 
**SnapshotRelativeDirPath** | Pointer to **NullableString** | The relative path to the directory containing the entity&#39;s snapshot. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | The start time of the specific job run. Iff &#39;job_instance_id&#39; is set, this field must be set. In-memory indices on the Magneto master are laid-out by the start time of the job, and this is how the master pulls up a specific run. NOTE: This must be specified for RestoreFiles, RecoverDisks and GetVirtualDisks APIs | [optional] 
**ViewName** | Pointer to **NullableString** | The name of the view where the object&#39;s snapshot is located. | [optional] 
**VmHadIndependentDisks** | Pointer to **NullableBool** | This is applicable only to VMs and is set to true when the VM being recovered or cloned contained independent disks when it was backed up. | [optional] 

## Methods

### NewRestoreObject

`func NewRestoreObject() *RestoreObject`

NewRestoreObject instantiates a new RestoreObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreObjectWithDefaults

`func NewRestoreObjectWithDefaults() *RestoreObject`

NewRestoreObjectWithDefaults instantiates a new RestoreObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivalTarget

`func (o *RestoreObject) GetArchivalTarget() ArchivalTarget`

GetArchivalTarget returns the ArchivalTarget field if non-nil, zero value otherwise.

### GetArchivalTargetOk

`func (o *RestoreObject) GetArchivalTargetOk() (*ArchivalTarget, bool)`

GetArchivalTargetOk returns a tuple with the ArchivalTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalTarget

`func (o *RestoreObject) SetArchivalTarget(v ArchivalTarget)`

SetArchivalTarget sets ArchivalTarget field to given value.

### HasArchivalTarget

`func (o *RestoreObject) HasArchivalTarget() bool`

HasArchivalTarget returns a boolean if a field has been set.

### GetAttemptNum

`func (o *RestoreObject) GetAttemptNum() int32`

GetAttemptNum returns the AttemptNum field if non-nil, zero value otherwise.

### GetAttemptNumOk

`func (o *RestoreObject) GetAttemptNumOk() (*int32, bool)`

GetAttemptNumOk returns a tuple with the AttemptNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptNum

`func (o *RestoreObject) SetAttemptNum(v int32)`

SetAttemptNum sets AttemptNum field to given value.

### HasAttemptNum

`func (o *RestoreObject) HasAttemptNum() bool`

HasAttemptNum returns a boolean if a field has been set.

### SetAttemptNumNil

`func (o *RestoreObject) SetAttemptNumNil(b bool)`

 SetAttemptNumNil sets the value for AttemptNum to be an explicit nil

### UnsetAttemptNum
`func (o *RestoreObject) UnsetAttemptNum()`

UnsetAttemptNum ensures that no value is present for AttemptNum, not even an explicit nil
### GetCloudDeployTarget

`func (o *RestoreObject) GetCloudDeployTarget() CloudDeployTarget`

GetCloudDeployTarget returns the CloudDeployTarget field if non-nil, zero value otherwise.

### GetCloudDeployTargetOk

`func (o *RestoreObject) GetCloudDeployTargetOk() (*CloudDeployTarget, bool)`

GetCloudDeployTargetOk returns a tuple with the CloudDeployTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDeployTarget

`func (o *RestoreObject) SetCloudDeployTarget(v CloudDeployTarget)`

SetCloudDeployTarget sets CloudDeployTarget field to given value.

### HasCloudDeployTarget

`func (o *RestoreObject) HasCloudDeployTarget() bool`

HasCloudDeployTarget returns a boolean if a field has been set.

### GetCloudReplicationTarget

`func (o *RestoreObject) GetCloudReplicationTarget() CloudDeployTarget`

GetCloudReplicationTarget returns the CloudReplicationTarget field if non-nil, zero value otherwise.

### GetCloudReplicationTargetOk

`func (o *RestoreObject) GetCloudReplicationTargetOk() (*CloudDeployTarget, bool)`

GetCloudReplicationTargetOk returns a tuple with the CloudReplicationTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudReplicationTarget

`func (o *RestoreObject) SetCloudReplicationTarget(v CloudDeployTarget)`

SetCloudReplicationTarget sets CloudReplicationTarget field to given value.

### HasCloudReplicationTarget

`func (o *RestoreObject) HasCloudReplicationTarget() bool`

HasCloudReplicationTarget returns a boolean if a field has been set.

### GetEntity

`func (o *RestoreObject) GetEntity() EntityProto`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *RestoreObject) GetEntityOk() (*EntityProto, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *RestoreObject) SetEntity(v EntityProto)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *RestoreObject) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetJobId

`func (o *RestoreObject) GetJobId() int64`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *RestoreObject) GetJobIdOk() (*int64, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *RestoreObject) SetJobId(v int64)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *RestoreObject) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *RestoreObject) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *RestoreObject) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil
### GetJobInstanceId

`func (o *RestoreObject) GetJobInstanceId() int64`

GetJobInstanceId returns the JobInstanceId field if non-nil, zero value otherwise.

### GetJobInstanceIdOk

`func (o *RestoreObject) GetJobInstanceIdOk() (*int64, bool)`

GetJobInstanceIdOk returns a tuple with the JobInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobInstanceId

`func (o *RestoreObject) SetJobInstanceId(v int64)`

SetJobInstanceId sets JobInstanceId field to given value.

### HasJobInstanceId

`func (o *RestoreObject) HasJobInstanceId() bool`

HasJobInstanceId returns a boolean if a field has been set.

### SetJobInstanceIdNil

`func (o *RestoreObject) SetJobInstanceIdNil(b bool)`

 SetJobInstanceIdNil sets the value for JobInstanceId to be an explicit nil

### UnsetJobInstanceId
`func (o *RestoreObject) UnsetJobInstanceId()`

UnsetJobInstanceId ensures that no value is present for JobInstanceId, not even an explicit nil
### GetJobUid

`func (o *RestoreObject) GetJobUid() UniversalIdProto`

GetJobUid returns the JobUid field if non-nil, zero value otherwise.

### GetJobUidOk

`func (o *RestoreObject) GetJobUidOk() (*UniversalIdProto, bool)`

GetJobUidOk returns a tuple with the JobUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobUid

`func (o *RestoreObject) SetJobUid(v UniversalIdProto)`

SetJobUid sets JobUid field to given value.

### HasJobUid

`func (o *RestoreObject) HasJobUid() bool`

HasJobUid returns a boolean if a field has been set.

### GetNosqlRecoverParams

`func (o *RestoreObject) GetNosqlRecoverParams() NoSqlRecoverParams`

GetNosqlRecoverParams returns the NosqlRecoverParams field if non-nil, zero value otherwise.

### GetNosqlRecoverParamsOk

`func (o *RestoreObject) GetNosqlRecoverParamsOk() (*NoSqlRecoverParams, bool)`

GetNosqlRecoverParamsOk returns a tuple with the NosqlRecoverParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNosqlRecoverParams

`func (o *RestoreObject) SetNosqlRecoverParams(v NoSqlRecoverParams)`

SetNosqlRecoverParams sets NosqlRecoverParams field to given value.

### HasNosqlRecoverParams

`func (o *RestoreObject) HasNosqlRecoverParams() bool`

HasNosqlRecoverParams returns a boolean if a field has been set.

### GetParentSource

`func (o *RestoreObject) GetParentSource() EntityProto`

GetParentSource returns the ParentSource field if non-nil, zero value otherwise.

### GetParentSourceOk

`func (o *RestoreObject) GetParentSourceOk() (*EntityProto, bool)`

GetParentSourceOk returns a tuple with the ParentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSource

`func (o *RestoreObject) SetParentSource(v EntityProto)`

SetParentSource sets ParentSource field to given value.

### HasParentSource

`func (o *RestoreObject) HasParentSource() bool`

HasParentSource returns a boolean if a field has been set.

### GetPointInTimeRestoreTimeUsecs

`func (o *RestoreObject) GetPointInTimeRestoreTimeUsecs() int64`

GetPointInTimeRestoreTimeUsecs returns the PointInTimeRestoreTimeUsecs field if non-nil, zero value otherwise.

### GetPointInTimeRestoreTimeUsecsOk

`func (o *RestoreObject) GetPointInTimeRestoreTimeUsecsOk() (*int64, bool)`

GetPointInTimeRestoreTimeUsecsOk returns a tuple with the PointInTimeRestoreTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTimeRestoreTimeUsecs

`func (o *RestoreObject) SetPointInTimeRestoreTimeUsecs(v int64)`

SetPointInTimeRestoreTimeUsecs sets PointInTimeRestoreTimeUsecs field to given value.

### HasPointInTimeRestoreTimeUsecs

`func (o *RestoreObject) HasPointInTimeRestoreTimeUsecs() bool`

HasPointInTimeRestoreTimeUsecs returns a boolean if a field has been set.

### SetPointInTimeRestoreTimeUsecsNil

`func (o *RestoreObject) SetPointInTimeRestoreTimeUsecsNil(b bool)`

 SetPointInTimeRestoreTimeUsecsNil sets the value for PointInTimeRestoreTimeUsecs to be an explicit nil

### UnsetPointInTimeRestoreTimeUsecs
`func (o *RestoreObject) UnsetPointInTimeRestoreTimeUsecs()`

UnsetPointInTimeRestoreTimeUsecs ensures that no value is present for PointInTimeRestoreTimeUsecs, not even an explicit nil
### GetRestoreAcropolisVmParam

`func (o *RestoreObject) GetRestoreAcropolisVmParam() RestoreAcropolisVMParam`

GetRestoreAcropolisVmParam returns the RestoreAcropolisVmParam field if non-nil, zero value otherwise.

### GetRestoreAcropolisVmParamOk

`func (o *RestoreObject) GetRestoreAcropolisVmParamOk() (*RestoreAcropolisVMParam, bool)`

GetRestoreAcropolisVmParamOk returns a tuple with the RestoreAcropolisVmParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAcropolisVmParam

`func (o *RestoreObject) SetRestoreAcropolisVmParam(v RestoreAcropolisVMParam)`

SetRestoreAcropolisVmParam sets RestoreAcropolisVmParam field to given value.

### HasRestoreAcropolisVmParam

`func (o *RestoreObject) HasRestoreAcropolisVmParam() bool`

HasRestoreAcropolisVmParam returns a boolean if a field has been set.

### GetSnapshotRelativeDirPath

`func (o *RestoreObject) GetSnapshotRelativeDirPath() string`

GetSnapshotRelativeDirPath returns the SnapshotRelativeDirPath field if non-nil, zero value otherwise.

### GetSnapshotRelativeDirPathOk

`func (o *RestoreObject) GetSnapshotRelativeDirPathOk() (*string, bool)`

GetSnapshotRelativeDirPathOk returns a tuple with the SnapshotRelativeDirPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotRelativeDirPath

`func (o *RestoreObject) SetSnapshotRelativeDirPath(v string)`

SetSnapshotRelativeDirPath sets SnapshotRelativeDirPath field to given value.

### HasSnapshotRelativeDirPath

`func (o *RestoreObject) HasSnapshotRelativeDirPath() bool`

HasSnapshotRelativeDirPath returns a boolean if a field has been set.

### SetSnapshotRelativeDirPathNil

`func (o *RestoreObject) SetSnapshotRelativeDirPathNil(b bool)`

 SetSnapshotRelativeDirPathNil sets the value for SnapshotRelativeDirPath to be an explicit nil

### UnsetSnapshotRelativeDirPath
`func (o *RestoreObject) UnsetSnapshotRelativeDirPath()`

UnsetSnapshotRelativeDirPath ensures that no value is present for SnapshotRelativeDirPath, not even an explicit nil
### GetStartTimeUsecs

`func (o *RestoreObject) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *RestoreObject) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *RestoreObject) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *RestoreObject) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *RestoreObject) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *RestoreObject) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetViewName

`func (o *RestoreObject) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *RestoreObject) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *RestoreObject) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *RestoreObject) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *RestoreObject) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *RestoreObject) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil
### GetVmHadIndependentDisks

`func (o *RestoreObject) GetVmHadIndependentDisks() bool`

GetVmHadIndependentDisks returns the VmHadIndependentDisks field if non-nil, zero value otherwise.

### GetVmHadIndependentDisksOk

`func (o *RestoreObject) GetVmHadIndependentDisksOk() (*bool, bool)`

GetVmHadIndependentDisksOk returns a tuple with the VmHadIndependentDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmHadIndependentDisks

`func (o *RestoreObject) SetVmHadIndependentDisks(v bool)`

SetVmHadIndependentDisks sets VmHadIndependentDisks field to given value.

### HasVmHadIndependentDisks

`func (o *RestoreObject) HasVmHadIndependentDisks() bool`

HasVmHadIndependentDisks returns a boolean if a field has been set.

### SetVmHadIndependentDisksNil

`func (o *RestoreObject) SetVmHadIndependentDisksNil(b bool)`

 SetVmHadIndependentDisksNil sets the value for VmHadIndependentDisks to be an explicit nil

### UnsetVmHadIndependentDisks
`func (o *RestoreObject) UnsetVmHadIndependentDisks()`

UnsetVmHadIndependentDisks ensures that no value is present for VmHadIndependentDisks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


