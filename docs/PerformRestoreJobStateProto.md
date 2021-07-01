# PerformRestoreJobStateProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdmittedTimeUsecs** | Pointer to **NullableInt64** | The time at which the restore job was admitted to run on a Magneto master. This field will be set only after the status changes to &#39;kAdmitted&#39;. Using this field, amount of time spent in the waiting/queued state and the amount of time taken taken to actually run the job can be determined. wait time &#x3D; admitted_time_usecs - start_time_usecs run time &#x3D; end_time_usecs - admitted_time_usecs | [optional] 
**CancellationRequested** | Pointer to **NullableBool** | Whether this restore job has a pending cancellation request. | [optional] 
**ContinueRestoreOnError** | Pointer to **NullableBool** | Whether to continue with the restore operation if restore of any object fails. | [optional] 
**DeployVmsToCloudTaskState** | Pointer to [**DeployVMsToCloudTaskStateProto**](DeployVMsToCloudTaskStateProto.md) |  | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | If the restore job has finished, this field contains the end time for the job. | [optional] 
**Error** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 
**Name** | Pointer to **NullableString** | The name of the restore job. | [optional] 
**NosqlConnectParams** | Pointer to [**NoSqlConnectParams**](NoSqlConnectParams.md) |  | [optional] 
**NosqlRecoverJobParams** | Pointer to [**NoSqlRecoverJobParams**](NoSqlRecoverJobParams.md) |  | [optional] 
**Objects** | Pointer to [**[]RestoreObject**](RestoreObject.md) | Information on the exact set of objects being restored (along with snapshots they are being recovered from). Even if the user wanted to restore an entire job form the latest snapshot, this will have individual objects and the exact snapshot they are bein restored from. If specified, this can only have leaf-level entities. | [optional] 
**ParentSourceConnectionParams** | Pointer to [**ConnectorParams**](ConnectorParams.md) |  | [optional] 
**PhysicalFlrParallelRestore** | Pointer to **NullableBool** | If enabled, magneto physical file restore will be enabled via job framework | [optional] 
**PowerStateConfig** | Pointer to [**PowerStateConfigProto**](PowerStateConfigProto.md) |  | [optional] 
**PreserveTags** | Pointer to **NullableBool** | Whether to preserve tags for the clone op. This field is currently used by HyperV and VMWare. | [optional] 
**ProgressMonitorTaskPath** | Pointer to **NullableString** | Root path of a Pulse task tracking the progress of the restore job. | [optional] 
**RenameRestoredObjectParam** | Pointer to [**RenameObjectParamProto**](RenameObjectParamProto.md) |  | [optional] 
**RenameRestoredVappParam** | Pointer to [**RenameObjectParamProto**](RenameObjectParamProto.md) |  | [optional] 
**RestoreAcropolisVmsParams** | Pointer to [**RestoreAcropolisVMsParams**](RestoreAcropolisVMsParams.md) |  | [optional] 
**RestoreJobId** | Pointer to **NullableInt64** | A globally unique id for this restore job. | [optional] 
**RestoreKubernetesNamespacesParams** | Pointer to [**RestoreKubernetesNamespacesParams**](RestoreKubernetesNamespacesParams.md) |  | [optional] 
**RestoreKvmVmsParams** | Pointer to [**RestoreKVMVMsParams**](RestoreKVMVMsParams.md) |  | [optional] 
**RestoreParentSource** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**RestorePublicFoldersParams** | Pointer to [**RestoreO365PublicFoldersParams**](RestoreO365PublicFoldersParams.md) |  | [optional] 
**RestoreSiteParams** | Pointer to [**RestoreSiteParams**](RestoreSiteParams.md) |  | [optional] 
**RestoreTaskStateProtoTmpl** | Pointer to [**PerformRestoreTaskStateProto**](PerformRestoreTaskStateProto.md) |  | [optional] 
**RestoreTaskVec** | Pointer to [**[]PerformRestoreJobStateProtoRestoreTask**](PerformRestoreJobStateProtoRestoreTask.md) | Even if the user wanted to restore an entire job from the latest snapshot, this will have info of all the individual objects. | [optional] 
**RestoreVmwareVmParams** | Pointer to [**RestoreVMwareVMParams**](RestoreVMwareVMParams.md) |  | [optional] 
**RestoredObjectsNetworkConfig** | Pointer to [**RestoredObjectNetworkConfigProto**](RestoredObjectNetworkConfigProto.md) |  | [optional] 
**RestoredToDifferentSource** | Pointer to **NullableBool** | Whether restore is being performed to a different parent source. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | The start time for this restore job. | [optional] 
**Status** | Pointer to **NullableInt32** | Status of the restore job. | [optional] 
**Type** | Pointer to **NullableInt32** | The type of restore being performed. | [optional] 
**User** | Pointer to **NullableString** | The user who requested this restore job. | [optional] 
**UserInfo** | Pointer to [**UserInformation**](UserInformation.md) |  | [optional] 
**VcdConfig** | Pointer to [**RestoredObjectVCDConfigProto**](RestoredObjectVCDConfigProto.md) |  | [optional] 
**ViewBoxId** | Pointer to **NullableInt64** | The view box id to which the restore job belongs to. | [optional] 
**Warnings** | Pointer to [**[]ErrorProto**](ErrorProto.md) | Populate warnings on the job if any. The warning messages are propagated from the child restore tasks upon completion of the task. | [optional] 

## Methods

### NewPerformRestoreJobStateProto

`func NewPerformRestoreJobStateProto() *PerformRestoreJobStateProto`

NewPerformRestoreJobStateProto instantiates a new PerformRestoreJobStateProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformRestoreJobStateProtoWithDefaults

`func NewPerformRestoreJobStateProtoWithDefaults() *PerformRestoreJobStateProto`

NewPerformRestoreJobStateProtoWithDefaults instantiates a new PerformRestoreJobStateProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmittedTimeUsecs

`func (o *PerformRestoreJobStateProto) GetAdmittedTimeUsecs() int64`

GetAdmittedTimeUsecs returns the AdmittedTimeUsecs field if non-nil, zero value otherwise.

### GetAdmittedTimeUsecsOk

`func (o *PerformRestoreJobStateProto) GetAdmittedTimeUsecsOk() (*int64, bool)`

GetAdmittedTimeUsecsOk returns a tuple with the AdmittedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmittedTimeUsecs

`func (o *PerformRestoreJobStateProto) SetAdmittedTimeUsecs(v int64)`

SetAdmittedTimeUsecs sets AdmittedTimeUsecs field to given value.

### HasAdmittedTimeUsecs

`func (o *PerformRestoreJobStateProto) HasAdmittedTimeUsecs() bool`

HasAdmittedTimeUsecs returns a boolean if a field has been set.

### SetAdmittedTimeUsecsNil

`func (o *PerformRestoreJobStateProto) SetAdmittedTimeUsecsNil(b bool)`

 SetAdmittedTimeUsecsNil sets the value for AdmittedTimeUsecs to be an explicit nil

### UnsetAdmittedTimeUsecs
`func (o *PerformRestoreJobStateProto) UnsetAdmittedTimeUsecs()`

UnsetAdmittedTimeUsecs ensures that no value is present for AdmittedTimeUsecs, not even an explicit nil
### GetCancellationRequested

`func (o *PerformRestoreJobStateProto) GetCancellationRequested() bool`

GetCancellationRequested returns the CancellationRequested field if non-nil, zero value otherwise.

### GetCancellationRequestedOk

`func (o *PerformRestoreJobStateProto) GetCancellationRequestedOk() (*bool, bool)`

GetCancellationRequestedOk returns a tuple with the CancellationRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationRequested

`func (o *PerformRestoreJobStateProto) SetCancellationRequested(v bool)`

SetCancellationRequested sets CancellationRequested field to given value.

### HasCancellationRequested

`func (o *PerformRestoreJobStateProto) HasCancellationRequested() bool`

HasCancellationRequested returns a boolean if a field has been set.

### SetCancellationRequestedNil

`func (o *PerformRestoreJobStateProto) SetCancellationRequestedNil(b bool)`

 SetCancellationRequestedNil sets the value for CancellationRequested to be an explicit nil

### UnsetCancellationRequested
`func (o *PerformRestoreJobStateProto) UnsetCancellationRequested()`

UnsetCancellationRequested ensures that no value is present for CancellationRequested, not even an explicit nil
### GetContinueRestoreOnError

`func (o *PerformRestoreJobStateProto) GetContinueRestoreOnError() bool`

GetContinueRestoreOnError returns the ContinueRestoreOnError field if non-nil, zero value otherwise.

### GetContinueRestoreOnErrorOk

`func (o *PerformRestoreJobStateProto) GetContinueRestoreOnErrorOk() (*bool, bool)`

GetContinueRestoreOnErrorOk returns a tuple with the ContinueRestoreOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueRestoreOnError

`func (o *PerformRestoreJobStateProto) SetContinueRestoreOnError(v bool)`

SetContinueRestoreOnError sets ContinueRestoreOnError field to given value.

### HasContinueRestoreOnError

`func (o *PerformRestoreJobStateProto) HasContinueRestoreOnError() bool`

HasContinueRestoreOnError returns a boolean if a field has been set.

### SetContinueRestoreOnErrorNil

`func (o *PerformRestoreJobStateProto) SetContinueRestoreOnErrorNil(b bool)`

 SetContinueRestoreOnErrorNil sets the value for ContinueRestoreOnError to be an explicit nil

### UnsetContinueRestoreOnError
`func (o *PerformRestoreJobStateProto) UnsetContinueRestoreOnError()`

UnsetContinueRestoreOnError ensures that no value is present for ContinueRestoreOnError, not even an explicit nil
### GetDeployVmsToCloudTaskState

`func (o *PerformRestoreJobStateProto) GetDeployVmsToCloudTaskState() DeployVMsToCloudTaskStateProto`

GetDeployVmsToCloudTaskState returns the DeployVmsToCloudTaskState field if non-nil, zero value otherwise.

### GetDeployVmsToCloudTaskStateOk

`func (o *PerformRestoreJobStateProto) GetDeployVmsToCloudTaskStateOk() (*DeployVMsToCloudTaskStateProto, bool)`

GetDeployVmsToCloudTaskStateOk returns a tuple with the DeployVmsToCloudTaskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployVmsToCloudTaskState

`func (o *PerformRestoreJobStateProto) SetDeployVmsToCloudTaskState(v DeployVMsToCloudTaskStateProto)`

SetDeployVmsToCloudTaskState sets DeployVmsToCloudTaskState field to given value.

### HasDeployVmsToCloudTaskState

`func (o *PerformRestoreJobStateProto) HasDeployVmsToCloudTaskState() bool`

HasDeployVmsToCloudTaskState returns a boolean if a field has been set.

### GetEndTimeUsecs

`func (o *PerformRestoreJobStateProto) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *PerformRestoreJobStateProto) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *PerformRestoreJobStateProto) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *PerformRestoreJobStateProto) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *PerformRestoreJobStateProto) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *PerformRestoreJobStateProto) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetError

`func (o *PerformRestoreJobStateProto) GetError() ErrorProto`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PerformRestoreJobStateProto) GetErrorOk() (*ErrorProto, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PerformRestoreJobStateProto) SetError(v ErrorProto)`

SetError sets Error field to given value.

### HasError

`func (o *PerformRestoreJobStateProto) HasError() bool`

HasError returns a boolean if a field has been set.

### GetName

`func (o *PerformRestoreJobStateProto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PerformRestoreJobStateProto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PerformRestoreJobStateProto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PerformRestoreJobStateProto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PerformRestoreJobStateProto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PerformRestoreJobStateProto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNosqlConnectParams

`func (o *PerformRestoreJobStateProto) GetNosqlConnectParams() NoSqlConnectParams`

GetNosqlConnectParams returns the NosqlConnectParams field if non-nil, zero value otherwise.

### GetNosqlConnectParamsOk

`func (o *PerformRestoreJobStateProto) GetNosqlConnectParamsOk() (*NoSqlConnectParams, bool)`

GetNosqlConnectParamsOk returns a tuple with the NosqlConnectParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNosqlConnectParams

`func (o *PerformRestoreJobStateProto) SetNosqlConnectParams(v NoSqlConnectParams)`

SetNosqlConnectParams sets NosqlConnectParams field to given value.

### HasNosqlConnectParams

`func (o *PerformRestoreJobStateProto) HasNosqlConnectParams() bool`

HasNosqlConnectParams returns a boolean if a field has been set.

### GetNosqlRecoverJobParams

`func (o *PerformRestoreJobStateProto) GetNosqlRecoverJobParams() NoSqlRecoverJobParams`

GetNosqlRecoverJobParams returns the NosqlRecoverJobParams field if non-nil, zero value otherwise.

### GetNosqlRecoverJobParamsOk

`func (o *PerformRestoreJobStateProto) GetNosqlRecoverJobParamsOk() (*NoSqlRecoverJobParams, bool)`

GetNosqlRecoverJobParamsOk returns a tuple with the NosqlRecoverJobParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNosqlRecoverJobParams

`func (o *PerformRestoreJobStateProto) SetNosqlRecoverJobParams(v NoSqlRecoverJobParams)`

SetNosqlRecoverJobParams sets NosqlRecoverJobParams field to given value.

### HasNosqlRecoverJobParams

`func (o *PerformRestoreJobStateProto) HasNosqlRecoverJobParams() bool`

HasNosqlRecoverJobParams returns a boolean if a field has been set.

### GetObjects

`func (o *PerformRestoreJobStateProto) GetObjects() []RestoreObject`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *PerformRestoreJobStateProto) GetObjectsOk() (*[]RestoreObject, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *PerformRestoreJobStateProto) SetObjects(v []RestoreObject)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *PerformRestoreJobStateProto) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *PerformRestoreJobStateProto) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *PerformRestoreJobStateProto) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetParentSourceConnectionParams

`func (o *PerformRestoreJobStateProto) GetParentSourceConnectionParams() ConnectorParams`

GetParentSourceConnectionParams returns the ParentSourceConnectionParams field if non-nil, zero value otherwise.

### GetParentSourceConnectionParamsOk

`func (o *PerformRestoreJobStateProto) GetParentSourceConnectionParamsOk() (*ConnectorParams, bool)`

GetParentSourceConnectionParamsOk returns a tuple with the ParentSourceConnectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSourceConnectionParams

`func (o *PerformRestoreJobStateProto) SetParentSourceConnectionParams(v ConnectorParams)`

SetParentSourceConnectionParams sets ParentSourceConnectionParams field to given value.

### HasParentSourceConnectionParams

`func (o *PerformRestoreJobStateProto) HasParentSourceConnectionParams() bool`

HasParentSourceConnectionParams returns a boolean if a field has been set.

### GetPhysicalFlrParallelRestore

`func (o *PerformRestoreJobStateProto) GetPhysicalFlrParallelRestore() bool`

GetPhysicalFlrParallelRestore returns the PhysicalFlrParallelRestore field if non-nil, zero value otherwise.

### GetPhysicalFlrParallelRestoreOk

`func (o *PerformRestoreJobStateProto) GetPhysicalFlrParallelRestoreOk() (*bool, bool)`

GetPhysicalFlrParallelRestoreOk returns a tuple with the PhysicalFlrParallelRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalFlrParallelRestore

`func (o *PerformRestoreJobStateProto) SetPhysicalFlrParallelRestore(v bool)`

SetPhysicalFlrParallelRestore sets PhysicalFlrParallelRestore field to given value.

### HasPhysicalFlrParallelRestore

`func (o *PerformRestoreJobStateProto) HasPhysicalFlrParallelRestore() bool`

HasPhysicalFlrParallelRestore returns a boolean if a field has been set.

### SetPhysicalFlrParallelRestoreNil

`func (o *PerformRestoreJobStateProto) SetPhysicalFlrParallelRestoreNil(b bool)`

 SetPhysicalFlrParallelRestoreNil sets the value for PhysicalFlrParallelRestore to be an explicit nil

### UnsetPhysicalFlrParallelRestore
`func (o *PerformRestoreJobStateProto) UnsetPhysicalFlrParallelRestore()`

UnsetPhysicalFlrParallelRestore ensures that no value is present for PhysicalFlrParallelRestore, not even an explicit nil
### GetPowerStateConfig

`func (o *PerformRestoreJobStateProto) GetPowerStateConfig() PowerStateConfigProto`

GetPowerStateConfig returns the PowerStateConfig field if non-nil, zero value otherwise.

### GetPowerStateConfigOk

`func (o *PerformRestoreJobStateProto) GetPowerStateConfigOk() (*PowerStateConfigProto, bool)`

GetPowerStateConfigOk returns a tuple with the PowerStateConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerStateConfig

`func (o *PerformRestoreJobStateProto) SetPowerStateConfig(v PowerStateConfigProto)`

SetPowerStateConfig sets PowerStateConfig field to given value.

### HasPowerStateConfig

`func (o *PerformRestoreJobStateProto) HasPowerStateConfig() bool`

HasPowerStateConfig returns a boolean if a field has been set.

### GetPreserveTags

`func (o *PerformRestoreJobStateProto) GetPreserveTags() bool`

GetPreserveTags returns the PreserveTags field if non-nil, zero value otherwise.

### GetPreserveTagsOk

`func (o *PerformRestoreJobStateProto) GetPreserveTagsOk() (*bool, bool)`

GetPreserveTagsOk returns a tuple with the PreserveTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveTags

`func (o *PerformRestoreJobStateProto) SetPreserveTags(v bool)`

SetPreserveTags sets PreserveTags field to given value.

### HasPreserveTags

`func (o *PerformRestoreJobStateProto) HasPreserveTags() bool`

HasPreserveTags returns a boolean if a field has been set.

### SetPreserveTagsNil

`func (o *PerformRestoreJobStateProto) SetPreserveTagsNil(b bool)`

 SetPreserveTagsNil sets the value for PreserveTags to be an explicit nil

### UnsetPreserveTags
`func (o *PerformRestoreJobStateProto) UnsetPreserveTags()`

UnsetPreserveTags ensures that no value is present for PreserveTags, not even an explicit nil
### GetProgressMonitorTaskPath

`func (o *PerformRestoreJobStateProto) GetProgressMonitorTaskPath() string`

GetProgressMonitorTaskPath returns the ProgressMonitorTaskPath field if non-nil, zero value otherwise.

### GetProgressMonitorTaskPathOk

`func (o *PerformRestoreJobStateProto) GetProgressMonitorTaskPathOk() (*string, bool)`

GetProgressMonitorTaskPathOk returns a tuple with the ProgressMonitorTaskPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressMonitorTaskPath

`func (o *PerformRestoreJobStateProto) SetProgressMonitorTaskPath(v string)`

SetProgressMonitorTaskPath sets ProgressMonitorTaskPath field to given value.

### HasProgressMonitorTaskPath

`func (o *PerformRestoreJobStateProto) HasProgressMonitorTaskPath() bool`

HasProgressMonitorTaskPath returns a boolean if a field has been set.

### SetProgressMonitorTaskPathNil

`func (o *PerformRestoreJobStateProto) SetProgressMonitorTaskPathNil(b bool)`

 SetProgressMonitorTaskPathNil sets the value for ProgressMonitorTaskPath to be an explicit nil

### UnsetProgressMonitorTaskPath
`func (o *PerformRestoreJobStateProto) UnsetProgressMonitorTaskPath()`

UnsetProgressMonitorTaskPath ensures that no value is present for ProgressMonitorTaskPath, not even an explicit nil
### GetRenameRestoredObjectParam

`func (o *PerformRestoreJobStateProto) GetRenameRestoredObjectParam() RenameObjectParamProto`

GetRenameRestoredObjectParam returns the RenameRestoredObjectParam field if non-nil, zero value otherwise.

### GetRenameRestoredObjectParamOk

`func (o *PerformRestoreJobStateProto) GetRenameRestoredObjectParamOk() (*RenameObjectParamProto, bool)`

GetRenameRestoredObjectParamOk returns a tuple with the RenameRestoredObjectParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRestoredObjectParam

`func (o *PerformRestoreJobStateProto) SetRenameRestoredObjectParam(v RenameObjectParamProto)`

SetRenameRestoredObjectParam sets RenameRestoredObjectParam field to given value.

### HasRenameRestoredObjectParam

`func (o *PerformRestoreJobStateProto) HasRenameRestoredObjectParam() bool`

HasRenameRestoredObjectParam returns a boolean if a field has been set.

### GetRenameRestoredVappParam

`func (o *PerformRestoreJobStateProto) GetRenameRestoredVappParam() RenameObjectParamProto`

GetRenameRestoredVappParam returns the RenameRestoredVappParam field if non-nil, zero value otherwise.

### GetRenameRestoredVappParamOk

`func (o *PerformRestoreJobStateProto) GetRenameRestoredVappParamOk() (*RenameObjectParamProto, bool)`

GetRenameRestoredVappParamOk returns a tuple with the RenameRestoredVappParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRestoredVappParam

`func (o *PerformRestoreJobStateProto) SetRenameRestoredVappParam(v RenameObjectParamProto)`

SetRenameRestoredVappParam sets RenameRestoredVappParam field to given value.

### HasRenameRestoredVappParam

`func (o *PerformRestoreJobStateProto) HasRenameRestoredVappParam() bool`

HasRenameRestoredVappParam returns a boolean if a field has been set.

### GetRestoreAcropolisVmsParams

`func (o *PerformRestoreJobStateProto) GetRestoreAcropolisVmsParams() RestoreAcropolisVMsParams`

GetRestoreAcropolisVmsParams returns the RestoreAcropolisVmsParams field if non-nil, zero value otherwise.

### GetRestoreAcropolisVmsParamsOk

`func (o *PerformRestoreJobStateProto) GetRestoreAcropolisVmsParamsOk() (*RestoreAcropolisVMsParams, bool)`

GetRestoreAcropolisVmsParamsOk returns a tuple with the RestoreAcropolisVmsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAcropolisVmsParams

`func (o *PerformRestoreJobStateProto) SetRestoreAcropolisVmsParams(v RestoreAcropolisVMsParams)`

SetRestoreAcropolisVmsParams sets RestoreAcropolisVmsParams field to given value.

### HasRestoreAcropolisVmsParams

`func (o *PerformRestoreJobStateProto) HasRestoreAcropolisVmsParams() bool`

HasRestoreAcropolisVmsParams returns a boolean if a field has been set.

### GetRestoreJobId

`func (o *PerformRestoreJobStateProto) GetRestoreJobId() int64`

GetRestoreJobId returns the RestoreJobId field if non-nil, zero value otherwise.

### GetRestoreJobIdOk

`func (o *PerformRestoreJobStateProto) GetRestoreJobIdOk() (*int64, bool)`

GetRestoreJobIdOk returns a tuple with the RestoreJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreJobId

`func (o *PerformRestoreJobStateProto) SetRestoreJobId(v int64)`

SetRestoreJobId sets RestoreJobId field to given value.

### HasRestoreJobId

`func (o *PerformRestoreJobStateProto) HasRestoreJobId() bool`

HasRestoreJobId returns a boolean if a field has been set.

### SetRestoreJobIdNil

`func (o *PerformRestoreJobStateProto) SetRestoreJobIdNil(b bool)`

 SetRestoreJobIdNil sets the value for RestoreJobId to be an explicit nil

### UnsetRestoreJobId
`func (o *PerformRestoreJobStateProto) UnsetRestoreJobId()`

UnsetRestoreJobId ensures that no value is present for RestoreJobId, not even an explicit nil
### GetRestoreKubernetesNamespacesParams

`func (o *PerformRestoreJobStateProto) GetRestoreKubernetesNamespacesParams() RestoreKubernetesNamespacesParams`

GetRestoreKubernetesNamespacesParams returns the RestoreKubernetesNamespacesParams field if non-nil, zero value otherwise.

### GetRestoreKubernetesNamespacesParamsOk

`func (o *PerformRestoreJobStateProto) GetRestoreKubernetesNamespacesParamsOk() (*RestoreKubernetesNamespacesParams, bool)`

GetRestoreKubernetesNamespacesParamsOk returns a tuple with the RestoreKubernetesNamespacesParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreKubernetesNamespacesParams

`func (o *PerformRestoreJobStateProto) SetRestoreKubernetesNamespacesParams(v RestoreKubernetesNamespacesParams)`

SetRestoreKubernetesNamespacesParams sets RestoreKubernetesNamespacesParams field to given value.

### HasRestoreKubernetesNamespacesParams

`func (o *PerformRestoreJobStateProto) HasRestoreKubernetesNamespacesParams() bool`

HasRestoreKubernetesNamespacesParams returns a boolean if a field has been set.

### GetRestoreKvmVmsParams

`func (o *PerformRestoreJobStateProto) GetRestoreKvmVmsParams() RestoreKVMVMsParams`

GetRestoreKvmVmsParams returns the RestoreKvmVmsParams field if non-nil, zero value otherwise.

### GetRestoreKvmVmsParamsOk

`func (o *PerformRestoreJobStateProto) GetRestoreKvmVmsParamsOk() (*RestoreKVMVMsParams, bool)`

GetRestoreKvmVmsParamsOk returns a tuple with the RestoreKvmVmsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreKvmVmsParams

`func (o *PerformRestoreJobStateProto) SetRestoreKvmVmsParams(v RestoreKVMVMsParams)`

SetRestoreKvmVmsParams sets RestoreKvmVmsParams field to given value.

### HasRestoreKvmVmsParams

`func (o *PerformRestoreJobStateProto) HasRestoreKvmVmsParams() bool`

HasRestoreKvmVmsParams returns a boolean if a field has been set.

### GetRestoreParentSource

`func (o *PerformRestoreJobStateProto) GetRestoreParentSource() EntityProto`

GetRestoreParentSource returns the RestoreParentSource field if non-nil, zero value otherwise.

### GetRestoreParentSourceOk

`func (o *PerformRestoreJobStateProto) GetRestoreParentSourceOk() (*EntityProto, bool)`

GetRestoreParentSourceOk returns a tuple with the RestoreParentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreParentSource

`func (o *PerformRestoreJobStateProto) SetRestoreParentSource(v EntityProto)`

SetRestoreParentSource sets RestoreParentSource field to given value.

### HasRestoreParentSource

`func (o *PerformRestoreJobStateProto) HasRestoreParentSource() bool`

HasRestoreParentSource returns a boolean if a field has been set.

### GetRestorePublicFoldersParams

`func (o *PerformRestoreJobStateProto) GetRestorePublicFoldersParams() RestoreO365PublicFoldersParams`

GetRestorePublicFoldersParams returns the RestorePublicFoldersParams field if non-nil, zero value otherwise.

### GetRestorePublicFoldersParamsOk

`func (o *PerformRestoreJobStateProto) GetRestorePublicFoldersParamsOk() (*RestoreO365PublicFoldersParams, bool)`

GetRestorePublicFoldersParamsOk returns a tuple with the RestorePublicFoldersParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestorePublicFoldersParams

`func (o *PerformRestoreJobStateProto) SetRestorePublicFoldersParams(v RestoreO365PublicFoldersParams)`

SetRestorePublicFoldersParams sets RestorePublicFoldersParams field to given value.

### HasRestorePublicFoldersParams

`func (o *PerformRestoreJobStateProto) HasRestorePublicFoldersParams() bool`

HasRestorePublicFoldersParams returns a boolean if a field has been set.

### GetRestoreSiteParams

`func (o *PerformRestoreJobStateProto) GetRestoreSiteParams() RestoreSiteParams`

GetRestoreSiteParams returns the RestoreSiteParams field if non-nil, zero value otherwise.

### GetRestoreSiteParamsOk

`func (o *PerformRestoreJobStateProto) GetRestoreSiteParamsOk() (*RestoreSiteParams, bool)`

GetRestoreSiteParamsOk returns a tuple with the RestoreSiteParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreSiteParams

`func (o *PerformRestoreJobStateProto) SetRestoreSiteParams(v RestoreSiteParams)`

SetRestoreSiteParams sets RestoreSiteParams field to given value.

### HasRestoreSiteParams

`func (o *PerformRestoreJobStateProto) HasRestoreSiteParams() bool`

HasRestoreSiteParams returns a boolean if a field has been set.

### GetRestoreTaskStateProtoTmpl

`func (o *PerformRestoreJobStateProto) GetRestoreTaskStateProtoTmpl() PerformRestoreTaskStateProto`

GetRestoreTaskStateProtoTmpl returns the RestoreTaskStateProtoTmpl field if non-nil, zero value otherwise.

### GetRestoreTaskStateProtoTmplOk

`func (o *PerformRestoreJobStateProto) GetRestoreTaskStateProtoTmplOk() (*PerformRestoreTaskStateProto, bool)`

GetRestoreTaskStateProtoTmplOk returns a tuple with the RestoreTaskStateProtoTmpl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTaskStateProtoTmpl

`func (o *PerformRestoreJobStateProto) SetRestoreTaskStateProtoTmpl(v PerformRestoreTaskStateProto)`

SetRestoreTaskStateProtoTmpl sets RestoreTaskStateProtoTmpl field to given value.

### HasRestoreTaskStateProtoTmpl

`func (o *PerformRestoreJobStateProto) HasRestoreTaskStateProtoTmpl() bool`

HasRestoreTaskStateProtoTmpl returns a boolean if a field has been set.

### GetRestoreTaskVec

`func (o *PerformRestoreJobStateProto) GetRestoreTaskVec() []PerformRestoreJobStateProtoRestoreTask`

GetRestoreTaskVec returns the RestoreTaskVec field if non-nil, zero value otherwise.

### GetRestoreTaskVecOk

`func (o *PerformRestoreJobStateProto) GetRestoreTaskVecOk() (*[]PerformRestoreJobStateProtoRestoreTask, bool)`

GetRestoreTaskVecOk returns a tuple with the RestoreTaskVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTaskVec

`func (o *PerformRestoreJobStateProto) SetRestoreTaskVec(v []PerformRestoreJobStateProtoRestoreTask)`

SetRestoreTaskVec sets RestoreTaskVec field to given value.

### HasRestoreTaskVec

`func (o *PerformRestoreJobStateProto) HasRestoreTaskVec() bool`

HasRestoreTaskVec returns a boolean if a field has been set.

### SetRestoreTaskVecNil

`func (o *PerformRestoreJobStateProto) SetRestoreTaskVecNil(b bool)`

 SetRestoreTaskVecNil sets the value for RestoreTaskVec to be an explicit nil

### UnsetRestoreTaskVec
`func (o *PerformRestoreJobStateProto) UnsetRestoreTaskVec()`

UnsetRestoreTaskVec ensures that no value is present for RestoreTaskVec, not even an explicit nil
### GetRestoreVmwareVmParams

`func (o *PerformRestoreJobStateProto) GetRestoreVmwareVmParams() RestoreVMwareVMParams`

GetRestoreVmwareVmParams returns the RestoreVmwareVmParams field if non-nil, zero value otherwise.

### GetRestoreVmwareVmParamsOk

`func (o *PerformRestoreJobStateProto) GetRestoreVmwareVmParamsOk() (*RestoreVMwareVMParams, bool)`

GetRestoreVmwareVmParamsOk returns a tuple with the RestoreVmwareVmParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreVmwareVmParams

`func (o *PerformRestoreJobStateProto) SetRestoreVmwareVmParams(v RestoreVMwareVMParams)`

SetRestoreVmwareVmParams sets RestoreVmwareVmParams field to given value.

### HasRestoreVmwareVmParams

`func (o *PerformRestoreJobStateProto) HasRestoreVmwareVmParams() bool`

HasRestoreVmwareVmParams returns a boolean if a field has been set.

### GetRestoredObjectsNetworkConfig

`func (o *PerformRestoreJobStateProto) GetRestoredObjectsNetworkConfig() RestoredObjectNetworkConfigProto`

GetRestoredObjectsNetworkConfig returns the RestoredObjectsNetworkConfig field if non-nil, zero value otherwise.

### GetRestoredObjectsNetworkConfigOk

`func (o *PerformRestoreJobStateProto) GetRestoredObjectsNetworkConfigOk() (*RestoredObjectNetworkConfigProto, bool)`

GetRestoredObjectsNetworkConfigOk returns a tuple with the RestoredObjectsNetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoredObjectsNetworkConfig

`func (o *PerformRestoreJobStateProto) SetRestoredObjectsNetworkConfig(v RestoredObjectNetworkConfigProto)`

SetRestoredObjectsNetworkConfig sets RestoredObjectsNetworkConfig field to given value.

### HasRestoredObjectsNetworkConfig

`func (o *PerformRestoreJobStateProto) HasRestoredObjectsNetworkConfig() bool`

HasRestoredObjectsNetworkConfig returns a boolean if a field has been set.

### GetRestoredToDifferentSource

`func (o *PerformRestoreJobStateProto) GetRestoredToDifferentSource() bool`

GetRestoredToDifferentSource returns the RestoredToDifferentSource field if non-nil, zero value otherwise.

### GetRestoredToDifferentSourceOk

`func (o *PerformRestoreJobStateProto) GetRestoredToDifferentSourceOk() (*bool, bool)`

GetRestoredToDifferentSourceOk returns a tuple with the RestoredToDifferentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoredToDifferentSource

`func (o *PerformRestoreJobStateProto) SetRestoredToDifferentSource(v bool)`

SetRestoredToDifferentSource sets RestoredToDifferentSource field to given value.

### HasRestoredToDifferentSource

`func (o *PerformRestoreJobStateProto) HasRestoredToDifferentSource() bool`

HasRestoredToDifferentSource returns a boolean if a field has been set.

### SetRestoredToDifferentSourceNil

`func (o *PerformRestoreJobStateProto) SetRestoredToDifferentSourceNil(b bool)`

 SetRestoredToDifferentSourceNil sets the value for RestoredToDifferentSource to be an explicit nil

### UnsetRestoredToDifferentSource
`func (o *PerformRestoreJobStateProto) UnsetRestoredToDifferentSource()`

UnsetRestoredToDifferentSource ensures that no value is present for RestoredToDifferentSource, not even an explicit nil
### GetStartTimeUsecs

`func (o *PerformRestoreJobStateProto) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *PerformRestoreJobStateProto) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *PerformRestoreJobStateProto) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *PerformRestoreJobStateProto) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *PerformRestoreJobStateProto) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *PerformRestoreJobStateProto) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetStatus

`func (o *PerformRestoreJobStateProto) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PerformRestoreJobStateProto) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PerformRestoreJobStateProto) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PerformRestoreJobStateProto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PerformRestoreJobStateProto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PerformRestoreJobStateProto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetType

`func (o *PerformRestoreJobStateProto) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PerformRestoreJobStateProto) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PerformRestoreJobStateProto) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *PerformRestoreJobStateProto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *PerformRestoreJobStateProto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PerformRestoreJobStateProto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUser

`func (o *PerformRestoreJobStateProto) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PerformRestoreJobStateProto) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PerformRestoreJobStateProto) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *PerformRestoreJobStateProto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *PerformRestoreJobStateProto) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *PerformRestoreJobStateProto) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetUserInfo

`func (o *PerformRestoreJobStateProto) GetUserInfo() UserInformation`

GetUserInfo returns the UserInfo field if non-nil, zero value otherwise.

### GetUserInfoOk

`func (o *PerformRestoreJobStateProto) GetUserInfoOk() (*UserInformation, bool)`

GetUserInfoOk returns a tuple with the UserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfo

`func (o *PerformRestoreJobStateProto) SetUserInfo(v UserInformation)`

SetUserInfo sets UserInfo field to given value.

### HasUserInfo

`func (o *PerformRestoreJobStateProto) HasUserInfo() bool`

HasUserInfo returns a boolean if a field has been set.

### GetVcdConfig

`func (o *PerformRestoreJobStateProto) GetVcdConfig() RestoredObjectVCDConfigProto`

GetVcdConfig returns the VcdConfig field if non-nil, zero value otherwise.

### GetVcdConfigOk

`func (o *PerformRestoreJobStateProto) GetVcdConfigOk() (*RestoredObjectVCDConfigProto, bool)`

GetVcdConfigOk returns a tuple with the VcdConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcdConfig

`func (o *PerformRestoreJobStateProto) SetVcdConfig(v RestoredObjectVCDConfigProto)`

SetVcdConfig sets VcdConfig field to given value.

### HasVcdConfig

`func (o *PerformRestoreJobStateProto) HasVcdConfig() bool`

HasVcdConfig returns a boolean if a field has been set.

### GetViewBoxId

`func (o *PerformRestoreJobStateProto) GetViewBoxId() int64`

GetViewBoxId returns the ViewBoxId field if non-nil, zero value otherwise.

### GetViewBoxIdOk

`func (o *PerformRestoreJobStateProto) GetViewBoxIdOk() (*int64, bool)`

GetViewBoxIdOk returns a tuple with the ViewBoxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewBoxId

`func (o *PerformRestoreJobStateProto) SetViewBoxId(v int64)`

SetViewBoxId sets ViewBoxId field to given value.

### HasViewBoxId

`func (o *PerformRestoreJobStateProto) HasViewBoxId() bool`

HasViewBoxId returns a boolean if a field has been set.

### SetViewBoxIdNil

`func (o *PerformRestoreJobStateProto) SetViewBoxIdNil(b bool)`

 SetViewBoxIdNil sets the value for ViewBoxId to be an explicit nil

### UnsetViewBoxId
`func (o *PerformRestoreJobStateProto) UnsetViewBoxId()`

UnsetViewBoxId ensures that no value is present for ViewBoxId, not even an explicit nil
### GetWarnings

`func (o *PerformRestoreJobStateProto) GetWarnings() []ErrorProto`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *PerformRestoreJobStateProto) GetWarningsOk() (*[]ErrorProto, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *PerformRestoreJobStateProto) SetWarnings(v []ErrorProto)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *PerformRestoreJobStateProto) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *PerformRestoreJobStateProto) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *PerformRestoreJobStateProto) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


