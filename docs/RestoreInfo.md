# RestoreInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivalTarget** | Pointer to [**ArchivalExternalTarget**](ArchivalExternalTarget.md) |  | [optional] 
**AttemptNumber** | Pointer to **NullableInt32** | Specifies the attempt number. | [optional] 
**CloudDeployTarget** | Pointer to [**CloudDeployTargetDetails**](CloudDeployTargetDetails.md) |  | [optional] 
**JobRunId** | Pointer to **NullableInt64** | Specifies the id of the job run. | [optional] 
**JobUid** | Pointer to [**UniversalId**](UniversalId.md) |  | [optional] 
**ParentSource** | Pointer to [**ProtectionSource**](ProtectionSource.md) |  | [optional] 
**RestoreTimeUsecs** | Pointer to **NullableInt64** | This field specifies the time in to which the object needs to be restored. This filed is only applicable when object is being backeup using CDP feature. | [optional] 
**SnapshotRelativeDirPath** | Pointer to **NullableString** | Specifies the relative path of the snapshot directory. | [optional] 
**Source** | Pointer to [**ProtectionSource**](ProtectionSource.md) |  | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time specified as a Unix epoch Timestamp (in microseconds). | [optional] 
**ViewName** | Pointer to **NullableString** | Specifies the name of the view. | [optional] 
**VmHadIndependentDisks** | Pointer to **NullableBool** | Specifies if the VM had independent disks. | [optional] 

## Methods

### NewRestoreInfo

`func NewRestoreInfo() *RestoreInfo`

NewRestoreInfo instantiates a new RestoreInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreInfoWithDefaults

`func NewRestoreInfoWithDefaults() *RestoreInfo`

NewRestoreInfoWithDefaults instantiates a new RestoreInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivalTarget

`func (o *RestoreInfo) GetArchivalTarget() ArchivalExternalTarget`

GetArchivalTarget returns the ArchivalTarget field if non-nil, zero value otherwise.

### GetArchivalTargetOk

`func (o *RestoreInfo) GetArchivalTargetOk() (*ArchivalExternalTarget, bool)`

GetArchivalTargetOk returns a tuple with the ArchivalTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalTarget

`func (o *RestoreInfo) SetArchivalTarget(v ArchivalExternalTarget)`

SetArchivalTarget sets ArchivalTarget field to given value.

### HasArchivalTarget

`func (o *RestoreInfo) HasArchivalTarget() bool`

HasArchivalTarget returns a boolean if a field has been set.

### GetAttemptNumber

`func (o *RestoreInfo) GetAttemptNumber() int32`

GetAttemptNumber returns the AttemptNumber field if non-nil, zero value otherwise.

### GetAttemptNumberOk

`func (o *RestoreInfo) GetAttemptNumberOk() (*int32, bool)`

GetAttemptNumberOk returns a tuple with the AttemptNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptNumber

`func (o *RestoreInfo) SetAttemptNumber(v int32)`

SetAttemptNumber sets AttemptNumber field to given value.

### HasAttemptNumber

`func (o *RestoreInfo) HasAttemptNumber() bool`

HasAttemptNumber returns a boolean if a field has been set.

### SetAttemptNumberNil

`func (o *RestoreInfo) SetAttemptNumberNil(b bool)`

 SetAttemptNumberNil sets the value for AttemptNumber to be an explicit nil

### UnsetAttemptNumber
`func (o *RestoreInfo) UnsetAttemptNumber()`

UnsetAttemptNumber ensures that no value is present for AttemptNumber, not even an explicit nil
### GetCloudDeployTarget

`func (o *RestoreInfo) GetCloudDeployTarget() CloudDeployTargetDetails`

GetCloudDeployTarget returns the CloudDeployTarget field if non-nil, zero value otherwise.

### GetCloudDeployTargetOk

`func (o *RestoreInfo) GetCloudDeployTargetOk() (*CloudDeployTargetDetails, bool)`

GetCloudDeployTargetOk returns a tuple with the CloudDeployTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDeployTarget

`func (o *RestoreInfo) SetCloudDeployTarget(v CloudDeployTargetDetails)`

SetCloudDeployTarget sets CloudDeployTarget field to given value.

### HasCloudDeployTarget

`func (o *RestoreInfo) HasCloudDeployTarget() bool`

HasCloudDeployTarget returns a boolean if a field has been set.

### GetJobRunId

`func (o *RestoreInfo) GetJobRunId() int64`

GetJobRunId returns the JobRunId field if non-nil, zero value otherwise.

### GetJobRunIdOk

`func (o *RestoreInfo) GetJobRunIdOk() (*int64, bool)`

GetJobRunIdOk returns a tuple with the JobRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRunId

`func (o *RestoreInfo) SetJobRunId(v int64)`

SetJobRunId sets JobRunId field to given value.

### HasJobRunId

`func (o *RestoreInfo) HasJobRunId() bool`

HasJobRunId returns a boolean if a field has been set.

### SetJobRunIdNil

`func (o *RestoreInfo) SetJobRunIdNil(b bool)`

 SetJobRunIdNil sets the value for JobRunId to be an explicit nil

### UnsetJobRunId
`func (o *RestoreInfo) UnsetJobRunId()`

UnsetJobRunId ensures that no value is present for JobRunId, not even an explicit nil
### GetJobUid

`func (o *RestoreInfo) GetJobUid() UniversalId`

GetJobUid returns the JobUid field if non-nil, zero value otherwise.

### GetJobUidOk

`func (o *RestoreInfo) GetJobUidOk() (*UniversalId, bool)`

GetJobUidOk returns a tuple with the JobUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobUid

`func (o *RestoreInfo) SetJobUid(v UniversalId)`

SetJobUid sets JobUid field to given value.

### HasJobUid

`func (o *RestoreInfo) HasJobUid() bool`

HasJobUid returns a boolean if a field has been set.

### GetParentSource

`func (o *RestoreInfo) GetParentSource() ProtectionSource`

GetParentSource returns the ParentSource field if non-nil, zero value otherwise.

### GetParentSourceOk

`func (o *RestoreInfo) GetParentSourceOk() (*ProtectionSource, bool)`

GetParentSourceOk returns a tuple with the ParentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSource

`func (o *RestoreInfo) SetParentSource(v ProtectionSource)`

SetParentSource sets ParentSource field to given value.

### HasParentSource

`func (o *RestoreInfo) HasParentSource() bool`

HasParentSource returns a boolean if a field has been set.

### GetRestoreTimeUsecs

`func (o *RestoreInfo) GetRestoreTimeUsecs() int64`

GetRestoreTimeUsecs returns the RestoreTimeUsecs field if non-nil, zero value otherwise.

### GetRestoreTimeUsecsOk

`func (o *RestoreInfo) GetRestoreTimeUsecsOk() (*int64, bool)`

GetRestoreTimeUsecsOk returns a tuple with the RestoreTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTimeUsecs

`func (o *RestoreInfo) SetRestoreTimeUsecs(v int64)`

SetRestoreTimeUsecs sets RestoreTimeUsecs field to given value.

### HasRestoreTimeUsecs

`func (o *RestoreInfo) HasRestoreTimeUsecs() bool`

HasRestoreTimeUsecs returns a boolean if a field has been set.

### SetRestoreTimeUsecsNil

`func (o *RestoreInfo) SetRestoreTimeUsecsNil(b bool)`

 SetRestoreTimeUsecsNil sets the value for RestoreTimeUsecs to be an explicit nil

### UnsetRestoreTimeUsecs
`func (o *RestoreInfo) UnsetRestoreTimeUsecs()`

UnsetRestoreTimeUsecs ensures that no value is present for RestoreTimeUsecs, not even an explicit nil
### GetSnapshotRelativeDirPath

`func (o *RestoreInfo) GetSnapshotRelativeDirPath() string`

GetSnapshotRelativeDirPath returns the SnapshotRelativeDirPath field if non-nil, zero value otherwise.

### GetSnapshotRelativeDirPathOk

`func (o *RestoreInfo) GetSnapshotRelativeDirPathOk() (*string, bool)`

GetSnapshotRelativeDirPathOk returns a tuple with the SnapshotRelativeDirPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotRelativeDirPath

`func (o *RestoreInfo) SetSnapshotRelativeDirPath(v string)`

SetSnapshotRelativeDirPath sets SnapshotRelativeDirPath field to given value.

### HasSnapshotRelativeDirPath

`func (o *RestoreInfo) HasSnapshotRelativeDirPath() bool`

HasSnapshotRelativeDirPath returns a boolean if a field has been set.

### SetSnapshotRelativeDirPathNil

`func (o *RestoreInfo) SetSnapshotRelativeDirPathNil(b bool)`

 SetSnapshotRelativeDirPathNil sets the value for SnapshotRelativeDirPath to be an explicit nil

### UnsetSnapshotRelativeDirPath
`func (o *RestoreInfo) UnsetSnapshotRelativeDirPath()`

UnsetSnapshotRelativeDirPath ensures that no value is present for SnapshotRelativeDirPath, not even an explicit nil
### GetSource

`func (o *RestoreInfo) GetSource() ProtectionSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RestoreInfo) GetSourceOk() (*ProtectionSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RestoreInfo) SetSource(v ProtectionSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *RestoreInfo) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStartTimeUsecs

`func (o *RestoreInfo) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *RestoreInfo) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *RestoreInfo) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *RestoreInfo) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *RestoreInfo) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *RestoreInfo) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetViewName

`func (o *RestoreInfo) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *RestoreInfo) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *RestoreInfo) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *RestoreInfo) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *RestoreInfo) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *RestoreInfo) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil
### GetVmHadIndependentDisks

`func (o *RestoreInfo) GetVmHadIndependentDisks() bool`

GetVmHadIndependentDisks returns the VmHadIndependentDisks field if non-nil, zero value otherwise.

### GetVmHadIndependentDisksOk

`func (o *RestoreInfo) GetVmHadIndependentDisksOk() (*bool, bool)`

GetVmHadIndependentDisksOk returns a tuple with the VmHadIndependentDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmHadIndependentDisks

`func (o *RestoreInfo) SetVmHadIndependentDisks(v bool)`

SetVmHadIndependentDisks sets VmHadIndependentDisks field to given value.

### HasVmHadIndependentDisks

`func (o *RestoreInfo) HasVmHadIndependentDisks() bool`

HasVmHadIndependentDisks returns a boolean if a field has been set.

### SetVmHadIndependentDisksNil

`func (o *RestoreInfo) SetVmHadIndependentDisksNil(b bool)`

 SetVmHadIndependentDisksNil sets the value for VmHadIndependentDisks to be an explicit nil

### UnsetVmHadIndependentDisks
`func (o *RestoreInfo) UnsetVmHadIndependentDisks()`

UnsetVmHadIndependentDisks ensures that no value is present for VmHadIndependentDisks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


