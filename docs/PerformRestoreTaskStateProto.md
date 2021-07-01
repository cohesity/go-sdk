# PerformRestoreTaskStateProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionExecutorTargetType** | Pointer to **NullableInt32** | Denotes the target for action executor(Bridge / BridgeProxy) on which task on slave should execute actions. | [optional] 
**Base** | Pointer to [**RestoreTaskStateBaseProto**](RestoreTaskStateBaseProto.md) |  | [optional] 
**CanTeardown** | Pointer to **NullableBool** | This is set if the clone operation has created any objects on the primary environment and teardown operation is possible. UI will disable the teardown button only if this is not set or set to false. NOTE: This won&#39;t be reset if the teardown operation subsequently completes as teardown state is managed separately. | [optional] 
**CdpRestoreProgressMonitorTaskPath** | Pointer to **NullableString** | The path of the progress monitor for the task that is responsible for creating the CDP hydrated view. | [optional] 
**CdpRestoreTask** | Pointer to [**PerformRestoreTaskStateProto**](PerformRestoreTaskStateProto.md) |  | [optional] 
**CdpRestoreTaskId** | Pointer to **NullableInt64** | The id of the task that will create the CDP hydrated view. | [optional] 
**CdpRestoreViewName** | Pointer to **NullableString** | The temporary view where the hydrated disks of the CDP restores are kept. | [optional] 
**ChildCloneTaskId** | Pointer to **NullableInt64** | The id of the child clone task triggered by refresh op. | [optional] 
**ChildDestroyTaskId** | Pointer to **NullableInt64** | The following fields are used by clone refresh op. These will be present only in case of refresh task op.  The id of the child destroy clone task triggered by refresh op. | [optional] 
**CloneAppViewInfo** | Pointer to [**CloneAppViewInfoProto**](CloneAppViewInfoProto.md) |  | [optional] 
**CloudDeployInfo** | Pointer to [**CloudDeployInfoProto**](CloudDeployInfoProto.md) |  | [optional] 
**ContinueRestoreOnError** | Pointer to **NullableBool** | Whether to continue with the restore operation if restore of any object fails. | [optional] 
**CreateView** | Pointer to **NullableBool** | True iff the target view needs to be created. | [optional] 
**DatastoreEntityVec** | Pointer to [**[]EntityProto**](EntityProto.md) | Please refer to comments for the field CreateRestoreTaskArg.datastore_entity_vec for more details. | [optional] 
**DeployVmsToCloudTaskState** | Pointer to [**DeployVMsToCloudTaskStateProto**](DeployVMsToCloudTaskStateProto.md) |  | [optional] 
**FolderEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**FullViewName** | Pointer to **NullableString** | The full view name (internal or external). This is composed of an optional Cohesity specific prefix and the user provided view name. | [optional] 
**IncludeVmConfig** | Pointer to **NullableBool** | This is set to true if the vm-config.xml need to be copied in the target view/folder. | [optional] 
**MountVolumesTaskState** | Pointer to [**MountVolumesTaskStateProto**](MountVolumesTaskStateProto.md) |  | [optional] 
**NosqlConnectParams** | Pointer to [**NoSqlConnectParams**](NoSqlConnectParams.md) |  | [optional] 
**NosqlRecoverJobParams** | Pointer to [**NoSqlRecoverJobParams**](NoSqlRecoverJobParams.md) |  | [optional] 
**ObjectNameDEPRECATED** | Pointer to **NullableString** | An optional name to give to the restored object. | [optional] 
**Objects** | Pointer to [**[]RestoreObject**](RestoreObject.md) | Information on the exact set of objects being restored (along with snapshots they are being recovered from). Even if the user wanted to restore an entire job from the latest snapshot, this will have individual objects and the exact snapshot they are being restored from. If specified, this can only have leaf-level entities. | [optional] 
**ObjectsProgressMonitorTaskPaths** | Pointer to **[]string** | Vector containing the relative task path of progress monitors of the objects in the above field &#39;objects&#39; to be restored. There is one to one correspondence between elements in &#39;objects&#39; and &#39;objects_progress_monitor_task_paths&#39;.  Please note that this field will be set only after progress monitor is created for this restore task. | [optional] 
**ParentRestoreJobId** | Pointer to **NullableInt64** | If this a child restore task, this field will contain the id of the parent restore job that spawned this task.  List of env and action type for which this field is applicable are: Acropolis: kRecoverVMs. | [optional] 
**ParentRestoreTaskId** | Pointer to **NullableInt64** | The id of the parent restore task if this is a restore sub-task.  List of environments that use this field: kSQL : Used for multi-stage SQL restore that supports a hot-standy.  This will also be used by refresh op to mark the new clone as internal sub-task. | [optional] 
**PathPrefixDEPRECATED** | Pointer to **NullableString** |  | [optional] 
**PhysicalFlrParallelRestore** | Pointer to **NullableBool** | If enabled, magneto physical file restore will be enabled via job framework | [optional] 
**PhysicalFlrUseNewLockingMethod** | Pointer to **NullableBool** | If enabled, magneto physical file restore will be enabled via job framework | [optional] 
**PowerStateConfig** | Pointer to [**PowerStateConfigProto**](PowerStateConfigProto.md) |  | [optional] 
**PreserveTags** | Pointer to **NullableBool** | This field is currently used by HyperV and VMWare. | [optional] 
**ProgressMonitorTaskPath** | Pointer to **NullableString** | Root path of a Pulse task tracking the progress of the restore task. | [optional] 
**RecoverDisksTaskState** | Pointer to [**RecoverDisksTaskStateProto**](RecoverDisksTaskStateProto.md) |  | [optional] 
**RecoverVolumesTaskState** | Pointer to [**RecoverVolumesTaskStateProto**](RecoverVolumesTaskStateProto.md) |  | [optional] 
**RelatedRestoreTaskId** | Pointer to **NullableInt64** | The task id of a related restore task. For example, a SQL restore operation may involve restoring a VM first and then restoring SQL databases after that. So the corresponding VM restore and SQL database restore tasks will be related, and they will each have their &#39;related_restore_task_id&#39; set to the id of the other task. | [optional] 
**RenameRestoredObjectParam** | Pointer to [**RenameObjectParamProto**](RenameObjectParamProto.md) |  | [optional] 
**RenameRestoredVappParam** | Pointer to [**RenameObjectParamProto**](RenameObjectParamProto.md) |  | [optional] 
**ResourcePoolEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**RestoreAcropolisVmsParams** | Pointer to [**RestoreAcropolisVMsParams**](RestoreAcropolisVMsParams.md) |  | [optional] 
**RestoreAppTaskState** | Pointer to [**RestoreAppTaskStateProto**](RestoreAppTaskStateProto.md) |  | [optional] 
**RestoreFilesTaskState** | Pointer to [**RestoreFilesTaskStateProto**](RestoreFilesTaskStateProto.md) |  | [optional] 
**RestoreHypervVmParams** | Pointer to [**RestoreHyperVVMParams**](RestoreHyperVVMParams.md) |  | [optional] 
**RestoreInfo** | Pointer to [**RestoreInfoProto**](RestoreInfoProto.md) |  | [optional] 
**RestoreKubernetesNamespacesParams** | Pointer to [**RestoreKubernetesNamespacesParams**](RestoreKubernetesNamespacesParams.md) |  | [optional] 
**RestoreKvmVmsParams** | Pointer to [**RestoreKVMVMsParams**](RestoreKVMVMsParams.md) |  | [optional] 
**RestoreOneDriveParams** | Pointer to [**RestoreOneDriveParams**](RestoreOneDriveParams.md) |  | [optional] 
**RestoreOutlookParams** | Pointer to [**RestoreOutlookParams**](RestoreOutlookParams.md) |  | [optional] 
**RestoreParentSource** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**RestorePublicFoldersParams** | Pointer to [**RestoreO365PublicFoldersParams**](RestoreO365PublicFoldersParams.md) |  | [optional] 
**RestoreSiteParams** | Pointer to [**RestoreSiteParams**](RestoreSiteParams.md) |  | [optional] 
**RestoreSubTaskVec** | Pointer to **[]int64** | Inside Magneto, these are represented as regular restore tasks with their own PerformRestoreTaskStateProto. Each restore sub-task will have its parent_restore_task_id field set.  List of environments that use this field: kSQL : Used for multi-stage SQL restore that supports a hot-standy. | [optional] 
**RestoreTaskPurged** | Pointer to **NullableBool** | Whether the restore task is purged. During WAL recovery, purged restore tasks are ignored. | [optional] 
**RestoreViewDatastoreEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**RestoreVmwareVmParams** | Pointer to [**RestoreVMwareVMParams**](RestoreVMwareVMParams.md) |  | [optional] 
**RestoredObjectsNetworkConfig** | Pointer to [**RestoredObjectNetworkConfigProto**](RestoredObjectNetworkConfigProto.md) |  | [optional] 
**RestoredToDifferentSource** | Pointer to **NullableBool** | Whether restore is being performed to a different parent source. | [optional] 
**RetrieveArchiveProgressMonitorTaskPath** | Pointer to **NullableString** | The path of the progress monitor for the task that is responsible for retrieving the objects from the archive. | [optional] 
**RetrieveArchiveStubViewName** | Pointer to **NullableString** | The stub view created by Icebox corresponding to the archive. The stub view is used to selectively retrieve files and folders. | [optional] 
**RetrieveArchiveTaskUidVec** | Pointer to [**[]UniversalIdProto**](UniversalIdProto.md) | The uid of the tasks that will retrieve the objects from the archive. Typically we only retrieve one snapshot for an enity but for point in time restores for SQL/Oracle database, we may need to retrieve multiple snapshots typically one full, and few logs. Hence we may need multiple uids to start retrieval task. | [optional] 
**RetrieveArchiveTaskVec** | Pointer to [**[]RetrieveArchiveTaskStateProto**](RetrieveArchiveTaskStateProto.md) | Proto that contains all the information about the retrieve archive task. Typically we only retrieve one snapshot for an enity but for point in time restores for SQL/Oracle database, we may need to retrieve multiple snapshots typically one full, and few logs. As we may start the multiple retrieval tasks, we need vector of RetrieveArchiveTaskStateProto for storing information of retrieved archive tasks. | [optional] 
**RetrieveArchiveViewName** | Pointer to **NullableString** | The temporary view where the entities that have been retrieved from an archive have been placed in by Icebox. | [optional] 
**SelectedDatastoreIdx** | Pointer to **NullableInt64** | In case of restore job with multi-vm multi-datastore this field denotes the specific datastore index in datastore_entity_vec to be selected for the task. Not going for specific datastore allocation in datastore_entity_vec so that we have required information in case of extensibility for task level retries with different datastore | [optional] 
**StubViewRelativeDirName** | Pointer to **NullableString** | Relative directory inside the stub view that corresponds with the archive. | [optional] 
**VaultRestoreParams** | Pointer to [**VaultParamsRestoreParams**](VaultParamsRestoreParams.md) |  | [optional] 
**VcdConfig** | Pointer to [**RestoredObjectVCDConfigProto**](RestoredObjectVCDConfigProto.md) |  | [optional] 
**VcdStorageProfileDatastoreMorefVec** | Pointer to **[]string** | This field is applicable for VCD type recovery. It defines the compatible datastores for recovery to alternate location. This field is inferred using the storage profile in restore_vmware_vm_params below. | [optional] 
**ViewBoxId** | Pointer to **NullableInt64** | The view box id to which &#39;view_name&#39; belongs to. In case the restore task corresponds to a backup job, this view box corresponds to the view box of the backup job. | [optional] 
**ViewNameDEPRECATED** | Pointer to **NullableString** | The view name as provided by the user for this restore operation. | [optional] 
**ViewParams** | Pointer to [**ViewParams**](ViewParams.md) |  | [optional] 
**VolumeInfoVec** | Pointer to [**[]VolumeInfo**](VolumeInfo.md) | Information regarding volumes that are required for the restore task. This is populated for restore files and mount virtual disk ops. | [optional] 

## Methods

### NewPerformRestoreTaskStateProto

`func NewPerformRestoreTaskStateProto() *PerformRestoreTaskStateProto`

NewPerformRestoreTaskStateProto instantiates a new PerformRestoreTaskStateProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformRestoreTaskStateProtoWithDefaults

`func NewPerformRestoreTaskStateProtoWithDefaults() *PerformRestoreTaskStateProto`

NewPerformRestoreTaskStateProtoWithDefaults instantiates a new PerformRestoreTaskStateProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionExecutorTargetType

`func (o *PerformRestoreTaskStateProto) GetActionExecutorTargetType() int32`

GetActionExecutorTargetType returns the ActionExecutorTargetType field if non-nil, zero value otherwise.

### GetActionExecutorTargetTypeOk

`func (o *PerformRestoreTaskStateProto) GetActionExecutorTargetTypeOk() (*int32, bool)`

GetActionExecutorTargetTypeOk returns a tuple with the ActionExecutorTargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionExecutorTargetType

`func (o *PerformRestoreTaskStateProto) SetActionExecutorTargetType(v int32)`

SetActionExecutorTargetType sets ActionExecutorTargetType field to given value.

### HasActionExecutorTargetType

`func (o *PerformRestoreTaskStateProto) HasActionExecutorTargetType() bool`

HasActionExecutorTargetType returns a boolean if a field has been set.

### SetActionExecutorTargetTypeNil

`func (o *PerformRestoreTaskStateProto) SetActionExecutorTargetTypeNil(b bool)`

 SetActionExecutorTargetTypeNil sets the value for ActionExecutorTargetType to be an explicit nil

### UnsetActionExecutorTargetType
`func (o *PerformRestoreTaskStateProto) UnsetActionExecutorTargetType()`

UnsetActionExecutorTargetType ensures that no value is present for ActionExecutorTargetType, not even an explicit nil
### GetBase

`func (o *PerformRestoreTaskStateProto) GetBase() RestoreTaskStateBaseProto`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *PerformRestoreTaskStateProto) GetBaseOk() (*RestoreTaskStateBaseProto, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *PerformRestoreTaskStateProto) SetBase(v RestoreTaskStateBaseProto)`

SetBase sets Base field to given value.

### HasBase

`func (o *PerformRestoreTaskStateProto) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetCanTeardown

`func (o *PerformRestoreTaskStateProto) GetCanTeardown() bool`

GetCanTeardown returns the CanTeardown field if non-nil, zero value otherwise.

### GetCanTeardownOk

`func (o *PerformRestoreTaskStateProto) GetCanTeardownOk() (*bool, bool)`

GetCanTeardownOk returns a tuple with the CanTeardown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTeardown

`func (o *PerformRestoreTaskStateProto) SetCanTeardown(v bool)`

SetCanTeardown sets CanTeardown field to given value.

### HasCanTeardown

`func (o *PerformRestoreTaskStateProto) HasCanTeardown() bool`

HasCanTeardown returns a boolean if a field has been set.

### SetCanTeardownNil

`func (o *PerformRestoreTaskStateProto) SetCanTeardownNil(b bool)`

 SetCanTeardownNil sets the value for CanTeardown to be an explicit nil

### UnsetCanTeardown
`func (o *PerformRestoreTaskStateProto) UnsetCanTeardown()`

UnsetCanTeardown ensures that no value is present for CanTeardown, not even an explicit nil
### GetCdpRestoreProgressMonitorTaskPath

`func (o *PerformRestoreTaskStateProto) GetCdpRestoreProgressMonitorTaskPath() string`

GetCdpRestoreProgressMonitorTaskPath returns the CdpRestoreProgressMonitorTaskPath field if non-nil, zero value otherwise.

### GetCdpRestoreProgressMonitorTaskPathOk

`func (o *PerformRestoreTaskStateProto) GetCdpRestoreProgressMonitorTaskPathOk() (*string, bool)`

GetCdpRestoreProgressMonitorTaskPathOk returns a tuple with the CdpRestoreProgressMonitorTaskPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdpRestoreProgressMonitorTaskPath

`func (o *PerformRestoreTaskStateProto) SetCdpRestoreProgressMonitorTaskPath(v string)`

SetCdpRestoreProgressMonitorTaskPath sets CdpRestoreProgressMonitorTaskPath field to given value.

### HasCdpRestoreProgressMonitorTaskPath

`func (o *PerformRestoreTaskStateProto) HasCdpRestoreProgressMonitorTaskPath() bool`

HasCdpRestoreProgressMonitorTaskPath returns a boolean if a field has been set.

### SetCdpRestoreProgressMonitorTaskPathNil

`func (o *PerformRestoreTaskStateProto) SetCdpRestoreProgressMonitorTaskPathNil(b bool)`

 SetCdpRestoreProgressMonitorTaskPathNil sets the value for CdpRestoreProgressMonitorTaskPath to be an explicit nil

### UnsetCdpRestoreProgressMonitorTaskPath
`func (o *PerformRestoreTaskStateProto) UnsetCdpRestoreProgressMonitorTaskPath()`

UnsetCdpRestoreProgressMonitorTaskPath ensures that no value is present for CdpRestoreProgressMonitorTaskPath, not even an explicit nil
### GetCdpRestoreTask

`func (o *PerformRestoreTaskStateProto) GetCdpRestoreTask() PerformRestoreTaskStateProto`

GetCdpRestoreTask returns the CdpRestoreTask field if non-nil, zero value otherwise.

### GetCdpRestoreTaskOk

`func (o *PerformRestoreTaskStateProto) GetCdpRestoreTaskOk() (*PerformRestoreTaskStateProto, bool)`

GetCdpRestoreTaskOk returns a tuple with the CdpRestoreTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdpRestoreTask

`func (o *PerformRestoreTaskStateProto) SetCdpRestoreTask(v PerformRestoreTaskStateProto)`

SetCdpRestoreTask sets CdpRestoreTask field to given value.

### HasCdpRestoreTask

`func (o *PerformRestoreTaskStateProto) HasCdpRestoreTask() bool`

HasCdpRestoreTask returns a boolean if a field has been set.

### GetCdpRestoreTaskId

`func (o *PerformRestoreTaskStateProto) GetCdpRestoreTaskId() int64`

GetCdpRestoreTaskId returns the CdpRestoreTaskId field if non-nil, zero value otherwise.

### GetCdpRestoreTaskIdOk

`func (o *PerformRestoreTaskStateProto) GetCdpRestoreTaskIdOk() (*int64, bool)`

GetCdpRestoreTaskIdOk returns a tuple with the CdpRestoreTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdpRestoreTaskId

`func (o *PerformRestoreTaskStateProto) SetCdpRestoreTaskId(v int64)`

SetCdpRestoreTaskId sets CdpRestoreTaskId field to given value.

### HasCdpRestoreTaskId

`func (o *PerformRestoreTaskStateProto) HasCdpRestoreTaskId() bool`

HasCdpRestoreTaskId returns a boolean if a field has been set.

### SetCdpRestoreTaskIdNil

`func (o *PerformRestoreTaskStateProto) SetCdpRestoreTaskIdNil(b bool)`

 SetCdpRestoreTaskIdNil sets the value for CdpRestoreTaskId to be an explicit nil

### UnsetCdpRestoreTaskId
`func (o *PerformRestoreTaskStateProto) UnsetCdpRestoreTaskId()`

UnsetCdpRestoreTaskId ensures that no value is present for CdpRestoreTaskId, not even an explicit nil
### GetCdpRestoreViewName

`func (o *PerformRestoreTaskStateProto) GetCdpRestoreViewName() string`

GetCdpRestoreViewName returns the CdpRestoreViewName field if non-nil, zero value otherwise.

### GetCdpRestoreViewNameOk

`func (o *PerformRestoreTaskStateProto) GetCdpRestoreViewNameOk() (*string, bool)`

GetCdpRestoreViewNameOk returns a tuple with the CdpRestoreViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdpRestoreViewName

`func (o *PerformRestoreTaskStateProto) SetCdpRestoreViewName(v string)`

SetCdpRestoreViewName sets CdpRestoreViewName field to given value.

### HasCdpRestoreViewName

`func (o *PerformRestoreTaskStateProto) HasCdpRestoreViewName() bool`

HasCdpRestoreViewName returns a boolean if a field has been set.

### SetCdpRestoreViewNameNil

`func (o *PerformRestoreTaskStateProto) SetCdpRestoreViewNameNil(b bool)`

 SetCdpRestoreViewNameNil sets the value for CdpRestoreViewName to be an explicit nil

### UnsetCdpRestoreViewName
`func (o *PerformRestoreTaskStateProto) UnsetCdpRestoreViewName()`

UnsetCdpRestoreViewName ensures that no value is present for CdpRestoreViewName, not even an explicit nil
### GetChildCloneTaskId

`func (o *PerformRestoreTaskStateProto) GetChildCloneTaskId() int64`

GetChildCloneTaskId returns the ChildCloneTaskId field if non-nil, zero value otherwise.

### GetChildCloneTaskIdOk

`func (o *PerformRestoreTaskStateProto) GetChildCloneTaskIdOk() (*int64, bool)`

GetChildCloneTaskIdOk returns a tuple with the ChildCloneTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildCloneTaskId

`func (o *PerformRestoreTaskStateProto) SetChildCloneTaskId(v int64)`

SetChildCloneTaskId sets ChildCloneTaskId field to given value.

### HasChildCloneTaskId

`func (o *PerformRestoreTaskStateProto) HasChildCloneTaskId() bool`

HasChildCloneTaskId returns a boolean if a field has been set.

### SetChildCloneTaskIdNil

`func (o *PerformRestoreTaskStateProto) SetChildCloneTaskIdNil(b bool)`

 SetChildCloneTaskIdNil sets the value for ChildCloneTaskId to be an explicit nil

### UnsetChildCloneTaskId
`func (o *PerformRestoreTaskStateProto) UnsetChildCloneTaskId()`

UnsetChildCloneTaskId ensures that no value is present for ChildCloneTaskId, not even an explicit nil
### GetChildDestroyTaskId

`func (o *PerformRestoreTaskStateProto) GetChildDestroyTaskId() int64`

GetChildDestroyTaskId returns the ChildDestroyTaskId field if non-nil, zero value otherwise.

### GetChildDestroyTaskIdOk

`func (o *PerformRestoreTaskStateProto) GetChildDestroyTaskIdOk() (*int64, bool)`

GetChildDestroyTaskIdOk returns a tuple with the ChildDestroyTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildDestroyTaskId

`func (o *PerformRestoreTaskStateProto) SetChildDestroyTaskId(v int64)`

SetChildDestroyTaskId sets ChildDestroyTaskId field to given value.

### HasChildDestroyTaskId

`func (o *PerformRestoreTaskStateProto) HasChildDestroyTaskId() bool`

HasChildDestroyTaskId returns a boolean if a field has been set.

### SetChildDestroyTaskIdNil

`func (o *PerformRestoreTaskStateProto) SetChildDestroyTaskIdNil(b bool)`

 SetChildDestroyTaskIdNil sets the value for ChildDestroyTaskId to be an explicit nil

### UnsetChildDestroyTaskId
`func (o *PerformRestoreTaskStateProto) UnsetChildDestroyTaskId()`

UnsetChildDestroyTaskId ensures that no value is present for ChildDestroyTaskId, not even an explicit nil
### GetCloneAppViewInfo

`func (o *PerformRestoreTaskStateProto) GetCloneAppViewInfo() CloneAppViewInfoProto`

GetCloneAppViewInfo returns the CloneAppViewInfo field if non-nil, zero value otherwise.

### GetCloneAppViewInfoOk

`func (o *PerformRestoreTaskStateProto) GetCloneAppViewInfoOk() (*CloneAppViewInfoProto, bool)`

GetCloneAppViewInfoOk returns a tuple with the CloneAppViewInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneAppViewInfo

`func (o *PerformRestoreTaskStateProto) SetCloneAppViewInfo(v CloneAppViewInfoProto)`

SetCloneAppViewInfo sets CloneAppViewInfo field to given value.

### HasCloneAppViewInfo

`func (o *PerformRestoreTaskStateProto) HasCloneAppViewInfo() bool`

HasCloneAppViewInfo returns a boolean if a field has been set.

### GetCloudDeployInfo

`func (o *PerformRestoreTaskStateProto) GetCloudDeployInfo() CloudDeployInfoProto`

GetCloudDeployInfo returns the CloudDeployInfo field if non-nil, zero value otherwise.

### GetCloudDeployInfoOk

`func (o *PerformRestoreTaskStateProto) GetCloudDeployInfoOk() (*CloudDeployInfoProto, bool)`

GetCloudDeployInfoOk returns a tuple with the CloudDeployInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDeployInfo

`func (o *PerformRestoreTaskStateProto) SetCloudDeployInfo(v CloudDeployInfoProto)`

SetCloudDeployInfo sets CloudDeployInfo field to given value.

### HasCloudDeployInfo

`func (o *PerformRestoreTaskStateProto) HasCloudDeployInfo() bool`

HasCloudDeployInfo returns a boolean if a field has been set.

### GetContinueRestoreOnError

`func (o *PerformRestoreTaskStateProto) GetContinueRestoreOnError() bool`

GetContinueRestoreOnError returns the ContinueRestoreOnError field if non-nil, zero value otherwise.

### GetContinueRestoreOnErrorOk

`func (o *PerformRestoreTaskStateProto) GetContinueRestoreOnErrorOk() (*bool, bool)`

GetContinueRestoreOnErrorOk returns a tuple with the ContinueRestoreOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueRestoreOnError

`func (o *PerformRestoreTaskStateProto) SetContinueRestoreOnError(v bool)`

SetContinueRestoreOnError sets ContinueRestoreOnError field to given value.

### HasContinueRestoreOnError

`func (o *PerformRestoreTaskStateProto) HasContinueRestoreOnError() bool`

HasContinueRestoreOnError returns a boolean if a field has been set.

### SetContinueRestoreOnErrorNil

`func (o *PerformRestoreTaskStateProto) SetContinueRestoreOnErrorNil(b bool)`

 SetContinueRestoreOnErrorNil sets the value for ContinueRestoreOnError to be an explicit nil

### UnsetContinueRestoreOnError
`func (o *PerformRestoreTaskStateProto) UnsetContinueRestoreOnError()`

UnsetContinueRestoreOnError ensures that no value is present for ContinueRestoreOnError, not even an explicit nil
### GetCreateView

`func (o *PerformRestoreTaskStateProto) GetCreateView() bool`

GetCreateView returns the CreateView field if non-nil, zero value otherwise.

### GetCreateViewOk

`func (o *PerformRestoreTaskStateProto) GetCreateViewOk() (*bool, bool)`

GetCreateViewOk returns a tuple with the CreateView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateView

`func (o *PerformRestoreTaskStateProto) SetCreateView(v bool)`

SetCreateView sets CreateView field to given value.

### HasCreateView

`func (o *PerformRestoreTaskStateProto) HasCreateView() bool`

HasCreateView returns a boolean if a field has been set.

### SetCreateViewNil

`func (o *PerformRestoreTaskStateProto) SetCreateViewNil(b bool)`

 SetCreateViewNil sets the value for CreateView to be an explicit nil

### UnsetCreateView
`func (o *PerformRestoreTaskStateProto) UnsetCreateView()`

UnsetCreateView ensures that no value is present for CreateView, not even an explicit nil
### GetDatastoreEntityVec

`func (o *PerformRestoreTaskStateProto) GetDatastoreEntityVec() []EntityProto`

GetDatastoreEntityVec returns the DatastoreEntityVec field if non-nil, zero value otherwise.

### GetDatastoreEntityVecOk

`func (o *PerformRestoreTaskStateProto) GetDatastoreEntityVecOk() (*[]EntityProto, bool)`

GetDatastoreEntityVecOk returns a tuple with the DatastoreEntityVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreEntityVec

`func (o *PerformRestoreTaskStateProto) SetDatastoreEntityVec(v []EntityProto)`

SetDatastoreEntityVec sets DatastoreEntityVec field to given value.

### HasDatastoreEntityVec

`func (o *PerformRestoreTaskStateProto) HasDatastoreEntityVec() bool`

HasDatastoreEntityVec returns a boolean if a field has been set.

### SetDatastoreEntityVecNil

`func (o *PerformRestoreTaskStateProto) SetDatastoreEntityVecNil(b bool)`

 SetDatastoreEntityVecNil sets the value for DatastoreEntityVec to be an explicit nil

### UnsetDatastoreEntityVec
`func (o *PerformRestoreTaskStateProto) UnsetDatastoreEntityVec()`

UnsetDatastoreEntityVec ensures that no value is present for DatastoreEntityVec, not even an explicit nil
### GetDeployVmsToCloudTaskState

`func (o *PerformRestoreTaskStateProto) GetDeployVmsToCloudTaskState() DeployVMsToCloudTaskStateProto`

GetDeployVmsToCloudTaskState returns the DeployVmsToCloudTaskState field if non-nil, zero value otherwise.

### GetDeployVmsToCloudTaskStateOk

`func (o *PerformRestoreTaskStateProto) GetDeployVmsToCloudTaskStateOk() (*DeployVMsToCloudTaskStateProto, bool)`

GetDeployVmsToCloudTaskStateOk returns a tuple with the DeployVmsToCloudTaskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployVmsToCloudTaskState

`func (o *PerformRestoreTaskStateProto) SetDeployVmsToCloudTaskState(v DeployVMsToCloudTaskStateProto)`

SetDeployVmsToCloudTaskState sets DeployVmsToCloudTaskState field to given value.

### HasDeployVmsToCloudTaskState

`func (o *PerformRestoreTaskStateProto) HasDeployVmsToCloudTaskState() bool`

HasDeployVmsToCloudTaskState returns a boolean if a field has been set.

### GetFolderEntity

`func (o *PerformRestoreTaskStateProto) GetFolderEntity() EntityProto`

GetFolderEntity returns the FolderEntity field if non-nil, zero value otherwise.

### GetFolderEntityOk

`func (o *PerformRestoreTaskStateProto) GetFolderEntityOk() (*EntityProto, bool)`

GetFolderEntityOk returns a tuple with the FolderEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderEntity

`func (o *PerformRestoreTaskStateProto) SetFolderEntity(v EntityProto)`

SetFolderEntity sets FolderEntity field to given value.

### HasFolderEntity

`func (o *PerformRestoreTaskStateProto) HasFolderEntity() bool`

HasFolderEntity returns a boolean if a field has been set.

### GetFullViewName

`func (o *PerformRestoreTaskStateProto) GetFullViewName() string`

GetFullViewName returns the FullViewName field if non-nil, zero value otherwise.

### GetFullViewNameOk

`func (o *PerformRestoreTaskStateProto) GetFullViewNameOk() (*string, bool)`

GetFullViewNameOk returns a tuple with the FullViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullViewName

`func (o *PerformRestoreTaskStateProto) SetFullViewName(v string)`

SetFullViewName sets FullViewName field to given value.

### HasFullViewName

`func (o *PerformRestoreTaskStateProto) HasFullViewName() bool`

HasFullViewName returns a boolean if a field has been set.

### SetFullViewNameNil

`func (o *PerformRestoreTaskStateProto) SetFullViewNameNil(b bool)`

 SetFullViewNameNil sets the value for FullViewName to be an explicit nil

### UnsetFullViewName
`func (o *PerformRestoreTaskStateProto) UnsetFullViewName()`

UnsetFullViewName ensures that no value is present for FullViewName, not even an explicit nil
### GetIncludeVmConfig

`func (o *PerformRestoreTaskStateProto) GetIncludeVmConfig() bool`

GetIncludeVmConfig returns the IncludeVmConfig field if non-nil, zero value otherwise.

### GetIncludeVmConfigOk

`func (o *PerformRestoreTaskStateProto) GetIncludeVmConfigOk() (*bool, bool)`

GetIncludeVmConfigOk returns a tuple with the IncludeVmConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeVmConfig

`func (o *PerformRestoreTaskStateProto) SetIncludeVmConfig(v bool)`

SetIncludeVmConfig sets IncludeVmConfig field to given value.

### HasIncludeVmConfig

`func (o *PerformRestoreTaskStateProto) HasIncludeVmConfig() bool`

HasIncludeVmConfig returns a boolean if a field has been set.

### SetIncludeVmConfigNil

`func (o *PerformRestoreTaskStateProto) SetIncludeVmConfigNil(b bool)`

 SetIncludeVmConfigNil sets the value for IncludeVmConfig to be an explicit nil

### UnsetIncludeVmConfig
`func (o *PerformRestoreTaskStateProto) UnsetIncludeVmConfig()`

UnsetIncludeVmConfig ensures that no value is present for IncludeVmConfig, not even an explicit nil
### GetMountVolumesTaskState

`func (o *PerformRestoreTaskStateProto) GetMountVolumesTaskState() MountVolumesTaskStateProto`

GetMountVolumesTaskState returns the MountVolumesTaskState field if non-nil, zero value otherwise.

### GetMountVolumesTaskStateOk

`func (o *PerformRestoreTaskStateProto) GetMountVolumesTaskStateOk() (*MountVolumesTaskStateProto, bool)`

GetMountVolumesTaskStateOk returns a tuple with the MountVolumesTaskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountVolumesTaskState

`func (o *PerformRestoreTaskStateProto) SetMountVolumesTaskState(v MountVolumesTaskStateProto)`

SetMountVolumesTaskState sets MountVolumesTaskState field to given value.

### HasMountVolumesTaskState

`func (o *PerformRestoreTaskStateProto) HasMountVolumesTaskState() bool`

HasMountVolumesTaskState returns a boolean if a field has been set.

### GetNosqlConnectParams

`func (o *PerformRestoreTaskStateProto) GetNosqlConnectParams() NoSqlConnectParams`

GetNosqlConnectParams returns the NosqlConnectParams field if non-nil, zero value otherwise.

### GetNosqlConnectParamsOk

`func (o *PerformRestoreTaskStateProto) GetNosqlConnectParamsOk() (*NoSqlConnectParams, bool)`

GetNosqlConnectParamsOk returns a tuple with the NosqlConnectParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNosqlConnectParams

`func (o *PerformRestoreTaskStateProto) SetNosqlConnectParams(v NoSqlConnectParams)`

SetNosqlConnectParams sets NosqlConnectParams field to given value.

### HasNosqlConnectParams

`func (o *PerformRestoreTaskStateProto) HasNosqlConnectParams() bool`

HasNosqlConnectParams returns a boolean if a field has been set.

### GetNosqlRecoverJobParams

`func (o *PerformRestoreTaskStateProto) GetNosqlRecoverJobParams() NoSqlRecoverJobParams`

GetNosqlRecoverJobParams returns the NosqlRecoverJobParams field if non-nil, zero value otherwise.

### GetNosqlRecoverJobParamsOk

`func (o *PerformRestoreTaskStateProto) GetNosqlRecoverJobParamsOk() (*NoSqlRecoverJobParams, bool)`

GetNosqlRecoverJobParamsOk returns a tuple with the NosqlRecoverJobParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNosqlRecoverJobParams

`func (o *PerformRestoreTaskStateProto) SetNosqlRecoverJobParams(v NoSqlRecoverJobParams)`

SetNosqlRecoverJobParams sets NosqlRecoverJobParams field to given value.

### HasNosqlRecoverJobParams

`func (o *PerformRestoreTaskStateProto) HasNosqlRecoverJobParams() bool`

HasNosqlRecoverJobParams returns a boolean if a field has been set.

### GetObjectNameDEPRECATED

`func (o *PerformRestoreTaskStateProto) GetObjectNameDEPRECATED() string`

GetObjectNameDEPRECATED returns the ObjectNameDEPRECATED field if non-nil, zero value otherwise.

### GetObjectNameDEPRECATEDOk

`func (o *PerformRestoreTaskStateProto) GetObjectNameDEPRECATEDOk() (*string, bool)`

GetObjectNameDEPRECATEDOk returns a tuple with the ObjectNameDEPRECATED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectNameDEPRECATED

`func (o *PerformRestoreTaskStateProto) SetObjectNameDEPRECATED(v string)`

SetObjectNameDEPRECATED sets ObjectNameDEPRECATED field to given value.

### HasObjectNameDEPRECATED

`func (o *PerformRestoreTaskStateProto) HasObjectNameDEPRECATED() bool`

HasObjectNameDEPRECATED returns a boolean if a field has been set.

### SetObjectNameDEPRECATEDNil

`func (o *PerformRestoreTaskStateProto) SetObjectNameDEPRECATEDNil(b bool)`

 SetObjectNameDEPRECATEDNil sets the value for ObjectNameDEPRECATED to be an explicit nil

### UnsetObjectNameDEPRECATED
`func (o *PerformRestoreTaskStateProto) UnsetObjectNameDEPRECATED()`

UnsetObjectNameDEPRECATED ensures that no value is present for ObjectNameDEPRECATED, not even an explicit nil
### GetObjects

`func (o *PerformRestoreTaskStateProto) GetObjects() []RestoreObject`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *PerformRestoreTaskStateProto) GetObjectsOk() (*[]RestoreObject, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *PerformRestoreTaskStateProto) SetObjects(v []RestoreObject)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *PerformRestoreTaskStateProto) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *PerformRestoreTaskStateProto) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *PerformRestoreTaskStateProto) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetObjectsProgressMonitorTaskPaths

`func (o *PerformRestoreTaskStateProto) GetObjectsProgressMonitorTaskPaths() []string`

GetObjectsProgressMonitorTaskPaths returns the ObjectsProgressMonitorTaskPaths field if non-nil, zero value otherwise.

### GetObjectsProgressMonitorTaskPathsOk

`func (o *PerformRestoreTaskStateProto) GetObjectsProgressMonitorTaskPathsOk() (*[]string, bool)`

GetObjectsProgressMonitorTaskPathsOk returns a tuple with the ObjectsProgressMonitorTaskPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectsProgressMonitorTaskPaths

`func (o *PerformRestoreTaskStateProto) SetObjectsProgressMonitorTaskPaths(v []string)`

SetObjectsProgressMonitorTaskPaths sets ObjectsProgressMonitorTaskPaths field to given value.

### HasObjectsProgressMonitorTaskPaths

`func (o *PerformRestoreTaskStateProto) HasObjectsProgressMonitorTaskPaths() bool`

HasObjectsProgressMonitorTaskPaths returns a boolean if a field has been set.

### SetObjectsProgressMonitorTaskPathsNil

`func (o *PerformRestoreTaskStateProto) SetObjectsProgressMonitorTaskPathsNil(b bool)`

 SetObjectsProgressMonitorTaskPathsNil sets the value for ObjectsProgressMonitorTaskPaths to be an explicit nil

### UnsetObjectsProgressMonitorTaskPaths
`func (o *PerformRestoreTaskStateProto) UnsetObjectsProgressMonitorTaskPaths()`

UnsetObjectsProgressMonitorTaskPaths ensures that no value is present for ObjectsProgressMonitorTaskPaths, not even an explicit nil
### GetParentRestoreJobId

`func (o *PerformRestoreTaskStateProto) GetParentRestoreJobId() int64`

GetParentRestoreJobId returns the ParentRestoreJobId field if non-nil, zero value otherwise.

### GetParentRestoreJobIdOk

`func (o *PerformRestoreTaskStateProto) GetParentRestoreJobIdOk() (*int64, bool)`

GetParentRestoreJobIdOk returns a tuple with the ParentRestoreJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRestoreJobId

`func (o *PerformRestoreTaskStateProto) SetParentRestoreJobId(v int64)`

SetParentRestoreJobId sets ParentRestoreJobId field to given value.

### HasParentRestoreJobId

`func (o *PerformRestoreTaskStateProto) HasParentRestoreJobId() bool`

HasParentRestoreJobId returns a boolean if a field has been set.

### SetParentRestoreJobIdNil

`func (o *PerformRestoreTaskStateProto) SetParentRestoreJobIdNil(b bool)`

 SetParentRestoreJobIdNil sets the value for ParentRestoreJobId to be an explicit nil

### UnsetParentRestoreJobId
`func (o *PerformRestoreTaskStateProto) UnsetParentRestoreJobId()`

UnsetParentRestoreJobId ensures that no value is present for ParentRestoreJobId, not even an explicit nil
### GetParentRestoreTaskId

`func (o *PerformRestoreTaskStateProto) GetParentRestoreTaskId() int64`

GetParentRestoreTaskId returns the ParentRestoreTaskId field if non-nil, zero value otherwise.

### GetParentRestoreTaskIdOk

`func (o *PerformRestoreTaskStateProto) GetParentRestoreTaskIdOk() (*int64, bool)`

GetParentRestoreTaskIdOk returns a tuple with the ParentRestoreTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRestoreTaskId

`func (o *PerformRestoreTaskStateProto) SetParentRestoreTaskId(v int64)`

SetParentRestoreTaskId sets ParentRestoreTaskId field to given value.

### HasParentRestoreTaskId

`func (o *PerformRestoreTaskStateProto) HasParentRestoreTaskId() bool`

HasParentRestoreTaskId returns a boolean if a field has been set.

### SetParentRestoreTaskIdNil

`func (o *PerformRestoreTaskStateProto) SetParentRestoreTaskIdNil(b bool)`

 SetParentRestoreTaskIdNil sets the value for ParentRestoreTaskId to be an explicit nil

### UnsetParentRestoreTaskId
`func (o *PerformRestoreTaskStateProto) UnsetParentRestoreTaskId()`

UnsetParentRestoreTaskId ensures that no value is present for ParentRestoreTaskId, not even an explicit nil
### GetPathPrefixDEPRECATED

`func (o *PerformRestoreTaskStateProto) GetPathPrefixDEPRECATED() string`

GetPathPrefixDEPRECATED returns the PathPrefixDEPRECATED field if non-nil, zero value otherwise.

### GetPathPrefixDEPRECATEDOk

`func (o *PerformRestoreTaskStateProto) GetPathPrefixDEPRECATEDOk() (*string, bool)`

GetPathPrefixDEPRECATEDOk returns a tuple with the PathPrefixDEPRECATED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPrefixDEPRECATED

`func (o *PerformRestoreTaskStateProto) SetPathPrefixDEPRECATED(v string)`

SetPathPrefixDEPRECATED sets PathPrefixDEPRECATED field to given value.

### HasPathPrefixDEPRECATED

`func (o *PerformRestoreTaskStateProto) HasPathPrefixDEPRECATED() bool`

HasPathPrefixDEPRECATED returns a boolean if a field has been set.

### SetPathPrefixDEPRECATEDNil

`func (o *PerformRestoreTaskStateProto) SetPathPrefixDEPRECATEDNil(b bool)`

 SetPathPrefixDEPRECATEDNil sets the value for PathPrefixDEPRECATED to be an explicit nil

### UnsetPathPrefixDEPRECATED
`func (o *PerformRestoreTaskStateProto) UnsetPathPrefixDEPRECATED()`

UnsetPathPrefixDEPRECATED ensures that no value is present for PathPrefixDEPRECATED, not even an explicit nil
### GetPhysicalFlrParallelRestore

`func (o *PerformRestoreTaskStateProto) GetPhysicalFlrParallelRestore() bool`

GetPhysicalFlrParallelRestore returns the PhysicalFlrParallelRestore field if non-nil, zero value otherwise.

### GetPhysicalFlrParallelRestoreOk

`func (o *PerformRestoreTaskStateProto) GetPhysicalFlrParallelRestoreOk() (*bool, bool)`

GetPhysicalFlrParallelRestoreOk returns a tuple with the PhysicalFlrParallelRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalFlrParallelRestore

`func (o *PerformRestoreTaskStateProto) SetPhysicalFlrParallelRestore(v bool)`

SetPhysicalFlrParallelRestore sets PhysicalFlrParallelRestore field to given value.

### HasPhysicalFlrParallelRestore

`func (o *PerformRestoreTaskStateProto) HasPhysicalFlrParallelRestore() bool`

HasPhysicalFlrParallelRestore returns a boolean if a field has been set.

### SetPhysicalFlrParallelRestoreNil

`func (o *PerformRestoreTaskStateProto) SetPhysicalFlrParallelRestoreNil(b bool)`

 SetPhysicalFlrParallelRestoreNil sets the value for PhysicalFlrParallelRestore to be an explicit nil

### UnsetPhysicalFlrParallelRestore
`func (o *PerformRestoreTaskStateProto) UnsetPhysicalFlrParallelRestore()`

UnsetPhysicalFlrParallelRestore ensures that no value is present for PhysicalFlrParallelRestore, not even an explicit nil
### GetPhysicalFlrUseNewLockingMethod

`func (o *PerformRestoreTaskStateProto) GetPhysicalFlrUseNewLockingMethod() bool`

GetPhysicalFlrUseNewLockingMethod returns the PhysicalFlrUseNewLockingMethod field if non-nil, zero value otherwise.

### GetPhysicalFlrUseNewLockingMethodOk

`func (o *PerformRestoreTaskStateProto) GetPhysicalFlrUseNewLockingMethodOk() (*bool, bool)`

GetPhysicalFlrUseNewLockingMethodOk returns a tuple with the PhysicalFlrUseNewLockingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalFlrUseNewLockingMethod

`func (o *PerformRestoreTaskStateProto) SetPhysicalFlrUseNewLockingMethod(v bool)`

SetPhysicalFlrUseNewLockingMethod sets PhysicalFlrUseNewLockingMethod field to given value.

### HasPhysicalFlrUseNewLockingMethod

`func (o *PerformRestoreTaskStateProto) HasPhysicalFlrUseNewLockingMethod() bool`

HasPhysicalFlrUseNewLockingMethod returns a boolean if a field has been set.

### SetPhysicalFlrUseNewLockingMethodNil

`func (o *PerformRestoreTaskStateProto) SetPhysicalFlrUseNewLockingMethodNil(b bool)`

 SetPhysicalFlrUseNewLockingMethodNil sets the value for PhysicalFlrUseNewLockingMethod to be an explicit nil

### UnsetPhysicalFlrUseNewLockingMethod
`func (o *PerformRestoreTaskStateProto) UnsetPhysicalFlrUseNewLockingMethod()`

UnsetPhysicalFlrUseNewLockingMethod ensures that no value is present for PhysicalFlrUseNewLockingMethod, not even an explicit nil
### GetPowerStateConfig

`func (o *PerformRestoreTaskStateProto) GetPowerStateConfig() PowerStateConfigProto`

GetPowerStateConfig returns the PowerStateConfig field if non-nil, zero value otherwise.

### GetPowerStateConfigOk

`func (o *PerformRestoreTaskStateProto) GetPowerStateConfigOk() (*PowerStateConfigProto, bool)`

GetPowerStateConfigOk returns a tuple with the PowerStateConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerStateConfig

`func (o *PerformRestoreTaskStateProto) SetPowerStateConfig(v PowerStateConfigProto)`

SetPowerStateConfig sets PowerStateConfig field to given value.

### HasPowerStateConfig

`func (o *PerformRestoreTaskStateProto) HasPowerStateConfig() bool`

HasPowerStateConfig returns a boolean if a field has been set.

### GetPreserveTags

`func (o *PerformRestoreTaskStateProto) GetPreserveTags() bool`

GetPreserveTags returns the PreserveTags field if non-nil, zero value otherwise.

### GetPreserveTagsOk

`func (o *PerformRestoreTaskStateProto) GetPreserveTagsOk() (*bool, bool)`

GetPreserveTagsOk returns a tuple with the PreserveTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveTags

`func (o *PerformRestoreTaskStateProto) SetPreserveTags(v bool)`

SetPreserveTags sets PreserveTags field to given value.

### HasPreserveTags

`func (o *PerformRestoreTaskStateProto) HasPreserveTags() bool`

HasPreserveTags returns a boolean if a field has been set.

### SetPreserveTagsNil

`func (o *PerformRestoreTaskStateProto) SetPreserveTagsNil(b bool)`

 SetPreserveTagsNil sets the value for PreserveTags to be an explicit nil

### UnsetPreserveTags
`func (o *PerformRestoreTaskStateProto) UnsetPreserveTags()`

UnsetPreserveTags ensures that no value is present for PreserveTags, not even an explicit nil
### GetProgressMonitorTaskPath

`func (o *PerformRestoreTaskStateProto) GetProgressMonitorTaskPath() string`

GetProgressMonitorTaskPath returns the ProgressMonitorTaskPath field if non-nil, zero value otherwise.

### GetProgressMonitorTaskPathOk

`func (o *PerformRestoreTaskStateProto) GetProgressMonitorTaskPathOk() (*string, bool)`

GetProgressMonitorTaskPathOk returns a tuple with the ProgressMonitorTaskPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressMonitorTaskPath

`func (o *PerformRestoreTaskStateProto) SetProgressMonitorTaskPath(v string)`

SetProgressMonitorTaskPath sets ProgressMonitorTaskPath field to given value.

### HasProgressMonitorTaskPath

`func (o *PerformRestoreTaskStateProto) HasProgressMonitorTaskPath() bool`

HasProgressMonitorTaskPath returns a boolean if a field has been set.

### SetProgressMonitorTaskPathNil

`func (o *PerformRestoreTaskStateProto) SetProgressMonitorTaskPathNil(b bool)`

 SetProgressMonitorTaskPathNil sets the value for ProgressMonitorTaskPath to be an explicit nil

### UnsetProgressMonitorTaskPath
`func (o *PerformRestoreTaskStateProto) UnsetProgressMonitorTaskPath()`

UnsetProgressMonitorTaskPath ensures that no value is present for ProgressMonitorTaskPath, not even an explicit nil
### GetRecoverDisksTaskState

`func (o *PerformRestoreTaskStateProto) GetRecoverDisksTaskState() RecoverDisksTaskStateProto`

GetRecoverDisksTaskState returns the RecoverDisksTaskState field if non-nil, zero value otherwise.

### GetRecoverDisksTaskStateOk

`func (o *PerformRestoreTaskStateProto) GetRecoverDisksTaskStateOk() (*RecoverDisksTaskStateProto, bool)`

GetRecoverDisksTaskStateOk returns a tuple with the RecoverDisksTaskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverDisksTaskState

`func (o *PerformRestoreTaskStateProto) SetRecoverDisksTaskState(v RecoverDisksTaskStateProto)`

SetRecoverDisksTaskState sets RecoverDisksTaskState field to given value.

### HasRecoverDisksTaskState

`func (o *PerformRestoreTaskStateProto) HasRecoverDisksTaskState() bool`

HasRecoverDisksTaskState returns a boolean if a field has been set.

### GetRecoverVolumesTaskState

`func (o *PerformRestoreTaskStateProto) GetRecoverVolumesTaskState() RecoverVolumesTaskStateProto`

GetRecoverVolumesTaskState returns the RecoverVolumesTaskState field if non-nil, zero value otherwise.

### GetRecoverVolumesTaskStateOk

`func (o *PerformRestoreTaskStateProto) GetRecoverVolumesTaskStateOk() (*RecoverVolumesTaskStateProto, bool)`

GetRecoverVolumesTaskStateOk returns a tuple with the RecoverVolumesTaskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverVolumesTaskState

`func (o *PerformRestoreTaskStateProto) SetRecoverVolumesTaskState(v RecoverVolumesTaskStateProto)`

SetRecoverVolumesTaskState sets RecoverVolumesTaskState field to given value.

### HasRecoverVolumesTaskState

`func (o *PerformRestoreTaskStateProto) HasRecoverVolumesTaskState() bool`

HasRecoverVolumesTaskState returns a boolean if a field has been set.

### GetRelatedRestoreTaskId

`func (o *PerformRestoreTaskStateProto) GetRelatedRestoreTaskId() int64`

GetRelatedRestoreTaskId returns the RelatedRestoreTaskId field if non-nil, zero value otherwise.

### GetRelatedRestoreTaskIdOk

`func (o *PerformRestoreTaskStateProto) GetRelatedRestoreTaskIdOk() (*int64, bool)`

GetRelatedRestoreTaskIdOk returns a tuple with the RelatedRestoreTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedRestoreTaskId

`func (o *PerformRestoreTaskStateProto) SetRelatedRestoreTaskId(v int64)`

SetRelatedRestoreTaskId sets RelatedRestoreTaskId field to given value.

### HasRelatedRestoreTaskId

`func (o *PerformRestoreTaskStateProto) HasRelatedRestoreTaskId() bool`

HasRelatedRestoreTaskId returns a boolean if a field has been set.

### SetRelatedRestoreTaskIdNil

`func (o *PerformRestoreTaskStateProto) SetRelatedRestoreTaskIdNil(b bool)`

 SetRelatedRestoreTaskIdNil sets the value for RelatedRestoreTaskId to be an explicit nil

### UnsetRelatedRestoreTaskId
`func (o *PerformRestoreTaskStateProto) UnsetRelatedRestoreTaskId()`

UnsetRelatedRestoreTaskId ensures that no value is present for RelatedRestoreTaskId, not even an explicit nil
### GetRenameRestoredObjectParam

`func (o *PerformRestoreTaskStateProto) GetRenameRestoredObjectParam() RenameObjectParamProto`

GetRenameRestoredObjectParam returns the RenameRestoredObjectParam field if non-nil, zero value otherwise.

### GetRenameRestoredObjectParamOk

`func (o *PerformRestoreTaskStateProto) GetRenameRestoredObjectParamOk() (*RenameObjectParamProto, bool)`

GetRenameRestoredObjectParamOk returns a tuple with the RenameRestoredObjectParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRestoredObjectParam

`func (o *PerformRestoreTaskStateProto) SetRenameRestoredObjectParam(v RenameObjectParamProto)`

SetRenameRestoredObjectParam sets RenameRestoredObjectParam field to given value.

### HasRenameRestoredObjectParam

`func (o *PerformRestoreTaskStateProto) HasRenameRestoredObjectParam() bool`

HasRenameRestoredObjectParam returns a boolean if a field has been set.

### GetRenameRestoredVappParam

`func (o *PerformRestoreTaskStateProto) GetRenameRestoredVappParam() RenameObjectParamProto`

GetRenameRestoredVappParam returns the RenameRestoredVappParam field if non-nil, zero value otherwise.

### GetRenameRestoredVappParamOk

`func (o *PerformRestoreTaskStateProto) GetRenameRestoredVappParamOk() (*RenameObjectParamProto, bool)`

GetRenameRestoredVappParamOk returns a tuple with the RenameRestoredVappParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRestoredVappParam

`func (o *PerformRestoreTaskStateProto) SetRenameRestoredVappParam(v RenameObjectParamProto)`

SetRenameRestoredVappParam sets RenameRestoredVappParam field to given value.

### HasRenameRestoredVappParam

`func (o *PerformRestoreTaskStateProto) HasRenameRestoredVappParam() bool`

HasRenameRestoredVappParam returns a boolean if a field has been set.

### GetResourcePoolEntity

`func (o *PerformRestoreTaskStateProto) GetResourcePoolEntity() EntityProto`

GetResourcePoolEntity returns the ResourcePoolEntity field if non-nil, zero value otherwise.

### GetResourcePoolEntityOk

`func (o *PerformRestoreTaskStateProto) GetResourcePoolEntityOk() (*EntityProto, bool)`

GetResourcePoolEntityOk returns a tuple with the ResourcePoolEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolEntity

`func (o *PerformRestoreTaskStateProto) SetResourcePoolEntity(v EntityProto)`

SetResourcePoolEntity sets ResourcePoolEntity field to given value.

### HasResourcePoolEntity

`func (o *PerformRestoreTaskStateProto) HasResourcePoolEntity() bool`

HasResourcePoolEntity returns a boolean if a field has been set.

### GetRestoreAcropolisVmsParams

`func (o *PerformRestoreTaskStateProto) GetRestoreAcropolisVmsParams() RestoreAcropolisVMsParams`

GetRestoreAcropolisVmsParams returns the RestoreAcropolisVmsParams field if non-nil, zero value otherwise.

### GetRestoreAcropolisVmsParamsOk

`func (o *PerformRestoreTaskStateProto) GetRestoreAcropolisVmsParamsOk() (*RestoreAcropolisVMsParams, bool)`

GetRestoreAcropolisVmsParamsOk returns a tuple with the RestoreAcropolisVmsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAcropolisVmsParams

`func (o *PerformRestoreTaskStateProto) SetRestoreAcropolisVmsParams(v RestoreAcropolisVMsParams)`

SetRestoreAcropolisVmsParams sets RestoreAcropolisVmsParams field to given value.

### HasRestoreAcropolisVmsParams

`func (o *PerformRestoreTaskStateProto) HasRestoreAcropolisVmsParams() bool`

HasRestoreAcropolisVmsParams returns a boolean if a field has been set.

### GetRestoreAppTaskState

`func (o *PerformRestoreTaskStateProto) GetRestoreAppTaskState() RestoreAppTaskStateProto`

GetRestoreAppTaskState returns the RestoreAppTaskState field if non-nil, zero value otherwise.

### GetRestoreAppTaskStateOk

`func (o *PerformRestoreTaskStateProto) GetRestoreAppTaskStateOk() (*RestoreAppTaskStateProto, bool)`

GetRestoreAppTaskStateOk returns a tuple with the RestoreAppTaskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAppTaskState

`func (o *PerformRestoreTaskStateProto) SetRestoreAppTaskState(v RestoreAppTaskStateProto)`

SetRestoreAppTaskState sets RestoreAppTaskState field to given value.

### HasRestoreAppTaskState

`func (o *PerformRestoreTaskStateProto) HasRestoreAppTaskState() bool`

HasRestoreAppTaskState returns a boolean if a field has been set.

### GetRestoreFilesTaskState

`func (o *PerformRestoreTaskStateProto) GetRestoreFilesTaskState() RestoreFilesTaskStateProto`

GetRestoreFilesTaskState returns the RestoreFilesTaskState field if non-nil, zero value otherwise.

### GetRestoreFilesTaskStateOk

`func (o *PerformRestoreTaskStateProto) GetRestoreFilesTaskStateOk() (*RestoreFilesTaskStateProto, bool)`

GetRestoreFilesTaskStateOk returns a tuple with the RestoreFilesTaskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreFilesTaskState

`func (o *PerformRestoreTaskStateProto) SetRestoreFilesTaskState(v RestoreFilesTaskStateProto)`

SetRestoreFilesTaskState sets RestoreFilesTaskState field to given value.

### HasRestoreFilesTaskState

`func (o *PerformRestoreTaskStateProto) HasRestoreFilesTaskState() bool`

HasRestoreFilesTaskState returns a boolean if a field has been set.

### GetRestoreHypervVmParams

`func (o *PerformRestoreTaskStateProto) GetRestoreHypervVmParams() RestoreHyperVVMParams`

GetRestoreHypervVmParams returns the RestoreHypervVmParams field if non-nil, zero value otherwise.

### GetRestoreHypervVmParamsOk

`func (o *PerformRestoreTaskStateProto) GetRestoreHypervVmParamsOk() (*RestoreHyperVVMParams, bool)`

GetRestoreHypervVmParamsOk returns a tuple with the RestoreHypervVmParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreHypervVmParams

`func (o *PerformRestoreTaskStateProto) SetRestoreHypervVmParams(v RestoreHyperVVMParams)`

SetRestoreHypervVmParams sets RestoreHypervVmParams field to given value.

### HasRestoreHypervVmParams

`func (o *PerformRestoreTaskStateProto) HasRestoreHypervVmParams() bool`

HasRestoreHypervVmParams returns a boolean if a field has been set.

### GetRestoreInfo

`func (o *PerformRestoreTaskStateProto) GetRestoreInfo() RestoreInfoProto`

GetRestoreInfo returns the RestoreInfo field if non-nil, zero value otherwise.

### GetRestoreInfoOk

`func (o *PerformRestoreTaskStateProto) GetRestoreInfoOk() (*RestoreInfoProto, bool)`

GetRestoreInfoOk returns a tuple with the RestoreInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreInfo

`func (o *PerformRestoreTaskStateProto) SetRestoreInfo(v RestoreInfoProto)`

SetRestoreInfo sets RestoreInfo field to given value.

### HasRestoreInfo

`func (o *PerformRestoreTaskStateProto) HasRestoreInfo() bool`

HasRestoreInfo returns a boolean if a field has been set.

### GetRestoreKubernetesNamespacesParams

`func (o *PerformRestoreTaskStateProto) GetRestoreKubernetesNamespacesParams() RestoreKubernetesNamespacesParams`

GetRestoreKubernetesNamespacesParams returns the RestoreKubernetesNamespacesParams field if non-nil, zero value otherwise.

### GetRestoreKubernetesNamespacesParamsOk

`func (o *PerformRestoreTaskStateProto) GetRestoreKubernetesNamespacesParamsOk() (*RestoreKubernetesNamespacesParams, bool)`

GetRestoreKubernetesNamespacesParamsOk returns a tuple with the RestoreKubernetesNamespacesParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreKubernetesNamespacesParams

`func (o *PerformRestoreTaskStateProto) SetRestoreKubernetesNamespacesParams(v RestoreKubernetesNamespacesParams)`

SetRestoreKubernetesNamespacesParams sets RestoreKubernetesNamespacesParams field to given value.

### HasRestoreKubernetesNamespacesParams

`func (o *PerformRestoreTaskStateProto) HasRestoreKubernetesNamespacesParams() bool`

HasRestoreKubernetesNamespacesParams returns a boolean if a field has been set.

### GetRestoreKvmVmsParams

`func (o *PerformRestoreTaskStateProto) GetRestoreKvmVmsParams() RestoreKVMVMsParams`

GetRestoreKvmVmsParams returns the RestoreKvmVmsParams field if non-nil, zero value otherwise.

### GetRestoreKvmVmsParamsOk

`func (o *PerformRestoreTaskStateProto) GetRestoreKvmVmsParamsOk() (*RestoreKVMVMsParams, bool)`

GetRestoreKvmVmsParamsOk returns a tuple with the RestoreKvmVmsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreKvmVmsParams

`func (o *PerformRestoreTaskStateProto) SetRestoreKvmVmsParams(v RestoreKVMVMsParams)`

SetRestoreKvmVmsParams sets RestoreKvmVmsParams field to given value.

### HasRestoreKvmVmsParams

`func (o *PerformRestoreTaskStateProto) HasRestoreKvmVmsParams() bool`

HasRestoreKvmVmsParams returns a boolean if a field has been set.

### GetRestoreOneDriveParams

`func (o *PerformRestoreTaskStateProto) GetRestoreOneDriveParams() RestoreOneDriveParams`

GetRestoreOneDriveParams returns the RestoreOneDriveParams field if non-nil, zero value otherwise.

### GetRestoreOneDriveParamsOk

`func (o *PerformRestoreTaskStateProto) GetRestoreOneDriveParamsOk() (*RestoreOneDriveParams, bool)`

GetRestoreOneDriveParamsOk returns a tuple with the RestoreOneDriveParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreOneDriveParams

`func (o *PerformRestoreTaskStateProto) SetRestoreOneDriveParams(v RestoreOneDriveParams)`

SetRestoreOneDriveParams sets RestoreOneDriveParams field to given value.

### HasRestoreOneDriveParams

`func (o *PerformRestoreTaskStateProto) HasRestoreOneDriveParams() bool`

HasRestoreOneDriveParams returns a boolean if a field has been set.

### GetRestoreOutlookParams

`func (o *PerformRestoreTaskStateProto) GetRestoreOutlookParams() RestoreOutlookParams`

GetRestoreOutlookParams returns the RestoreOutlookParams field if non-nil, zero value otherwise.

### GetRestoreOutlookParamsOk

`func (o *PerformRestoreTaskStateProto) GetRestoreOutlookParamsOk() (*RestoreOutlookParams, bool)`

GetRestoreOutlookParamsOk returns a tuple with the RestoreOutlookParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreOutlookParams

`func (o *PerformRestoreTaskStateProto) SetRestoreOutlookParams(v RestoreOutlookParams)`

SetRestoreOutlookParams sets RestoreOutlookParams field to given value.

### HasRestoreOutlookParams

`func (o *PerformRestoreTaskStateProto) HasRestoreOutlookParams() bool`

HasRestoreOutlookParams returns a boolean if a field has been set.

### GetRestoreParentSource

`func (o *PerformRestoreTaskStateProto) GetRestoreParentSource() EntityProto`

GetRestoreParentSource returns the RestoreParentSource field if non-nil, zero value otherwise.

### GetRestoreParentSourceOk

`func (o *PerformRestoreTaskStateProto) GetRestoreParentSourceOk() (*EntityProto, bool)`

GetRestoreParentSourceOk returns a tuple with the RestoreParentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreParentSource

`func (o *PerformRestoreTaskStateProto) SetRestoreParentSource(v EntityProto)`

SetRestoreParentSource sets RestoreParentSource field to given value.

### HasRestoreParentSource

`func (o *PerformRestoreTaskStateProto) HasRestoreParentSource() bool`

HasRestoreParentSource returns a boolean if a field has been set.

### GetRestorePublicFoldersParams

`func (o *PerformRestoreTaskStateProto) GetRestorePublicFoldersParams() RestoreO365PublicFoldersParams`

GetRestorePublicFoldersParams returns the RestorePublicFoldersParams field if non-nil, zero value otherwise.

### GetRestorePublicFoldersParamsOk

`func (o *PerformRestoreTaskStateProto) GetRestorePublicFoldersParamsOk() (*RestoreO365PublicFoldersParams, bool)`

GetRestorePublicFoldersParamsOk returns a tuple with the RestorePublicFoldersParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestorePublicFoldersParams

`func (o *PerformRestoreTaskStateProto) SetRestorePublicFoldersParams(v RestoreO365PublicFoldersParams)`

SetRestorePublicFoldersParams sets RestorePublicFoldersParams field to given value.

### HasRestorePublicFoldersParams

`func (o *PerformRestoreTaskStateProto) HasRestorePublicFoldersParams() bool`

HasRestorePublicFoldersParams returns a boolean if a field has been set.

### GetRestoreSiteParams

`func (o *PerformRestoreTaskStateProto) GetRestoreSiteParams() RestoreSiteParams`

GetRestoreSiteParams returns the RestoreSiteParams field if non-nil, zero value otherwise.

### GetRestoreSiteParamsOk

`func (o *PerformRestoreTaskStateProto) GetRestoreSiteParamsOk() (*RestoreSiteParams, bool)`

GetRestoreSiteParamsOk returns a tuple with the RestoreSiteParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreSiteParams

`func (o *PerformRestoreTaskStateProto) SetRestoreSiteParams(v RestoreSiteParams)`

SetRestoreSiteParams sets RestoreSiteParams field to given value.

### HasRestoreSiteParams

`func (o *PerformRestoreTaskStateProto) HasRestoreSiteParams() bool`

HasRestoreSiteParams returns a boolean if a field has been set.

### GetRestoreSubTaskVec

`func (o *PerformRestoreTaskStateProto) GetRestoreSubTaskVec() []int64`

GetRestoreSubTaskVec returns the RestoreSubTaskVec field if non-nil, zero value otherwise.

### GetRestoreSubTaskVecOk

`func (o *PerformRestoreTaskStateProto) GetRestoreSubTaskVecOk() (*[]int64, bool)`

GetRestoreSubTaskVecOk returns a tuple with the RestoreSubTaskVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreSubTaskVec

`func (o *PerformRestoreTaskStateProto) SetRestoreSubTaskVec(v []int64)`

SetRestoreSubTaskVec sets RestoreSubTaskVec field to given value.

### HasRestoreSubTaskVec

`func (o *PerformRestoreTaskStateProto) HasRestoreSubTaskVec() bool`

HasRestoreSubTaskVec returns a boolean if a field has been set.

### SetRestoreSubTaskVecNil

`func (o *PerformRestoreTaskStateProto) SetRestoreSubTaskVecNil(b bool)`

 SetRestoreSubTaskVecNil sets the value for RestoreSubTaskVec to be an explicit nil

### UnsetRestoreSubTaskVec
`func (o *PerformRestoreTaskStateProto) UnsetRestoreSubTaskVec()`

UnsetRestoreSubTaskVec ensures that no value is present for RestoreSubTaskVec, not even an explicit nil
### GetRestoreTaskPurged

`func (o *PerformRestoreTaskStateProto) GetRestoreTaskPurged() bool`

GetRestoreTaskPurged returns the RestoreTaskPurged field if non-nil, zero value otherwise.

### GetRestoreTaskPurgedOk

`func (o *PerformRestoreTaskStateProto) GetRestoreTaskPurgedOk() (*bool, bool)`

GetRestoreTaskPurgedOk returns a tuple with the RestoreTaskPurged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTaskPurged

`func (o *PerformRestoreTaskStateProto) SetRestoreTaskPurged(v bool)`

SetRestoreTaskPurged sets RestoreTaskPurged field to given value.

### HasRestoreTaskPurged

`func (o *PerformRestoreTaskStateProto) HasRestoreTaskPurged() bool`

HasRestoreTaskPurged returns a boolean if a field has been set.

### SetRestoreTaskPurgedNil

`func (o *PerformRestoreTaskStateProto) SetRestoreTaskPurgedNil(b bool)`

 SetRestoreTaskPurgedNil sets the value for RestoreTaskPurged to be an explicit nil

### UnsetRestoreTaskPurged
`func (o *PerformRestoreTaskStateProto) UnsetRestoreTaskPurged()`

UnsetRestoreTaskPurged ensures that no value is present for RestoreTaskPurged, not even an explicit nil
### GetRestoreViewDatastoreEntity

`func (o *PerformRestoreTaskStateProto) GetRestoreViewDatastoreEntity() EntityProto`

GetRestoreViewDatastoreEntity returns the RestoreViewDatastoreEntity field if non-nil, zero value otherwise.

### GetRestoreViewDatastoreEntityOk

`func (o *PerformRestoreTaskStateProto) GetRestoreViewDatastoreEntityOk() (*EntityProto, bool)`

GetRestoreViewDatastoreEntityOk returns a tuple with the RestoreViewDatastoreEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreViewDatastoreEntity

`func (o *PerformRestoreTaskStateProto) SetRestoreViewDatastoreEntity(v EntityProto)`

SetRestoreViewDatastoreEntity sets RestoreViewDatastoreEntity field to given value.

### HasRestoreViewDatastoreEntity

`func (o *PerformRestoreTaskStateProto) HasRestoreViewDatastoreEntity() bool`

HasRestoreViewDatastoreEntity returns a boolean if a field has been set.

### GetRestoreVmwareVmParams

`func (o *PerformRestoreTaskStateProto) GetRestoreVmwareVmParams() RestoreVMwareVMParams`

GetRestoreVmwareVmParams returns the RestoreVmwareVmParams field if non-nil, zero value otherwise.

### GetRestoreVmwareVmParamsOk

`func (o *PerformRestoreTaskStateProto) GetRestoreVmwareVmParamsOk() (*RestoreVMwareVMParams, bool)`

GetRestoreVmwareVmParamsOk returns a tuple with the RestoreVmwareVmParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreVmwareVmParams

`func (o *PerformRestoreTaskStateProto) SetRestoreVmwareVmParams(v RestoreVMwareVMParams)`

SetRestoreVmwareVmParams sets RestoreVmwareVmParams field to given value.

### HasRestoreVmwareVmParams

`func (o *PerformRestoreTaskStateProto) HasRestoreVmwareVmParams() bool`

HasRestoreVmwareVmParams returns a boolean if a field has been set.

### GetRestoredObjectsNetworkConfig

`func (o *PerformRestoreTaskStateProto) GetRestoredObjectsNetworkConfig() RestoredObjectNetworkConfigProto`

GetRestoredObjectsNetworkConfig returns the RestoredObjectsNetworkConfig field if non-nil, zero value otherwise.

### GetRestoredObjectsNetworkConfigOk

`func (o *PerformRestoreTaskStateProto) GetRestoredObjectsNetworkConfigOk() (*RestoredObjectNetworkConfigProto, bool)`

GetRestoredObjectsNetworkConfigOk returns a tuple with the RestoredObjectsNetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoredObjectsNetworkConfig

`func (o *PerformRestoreTaskStateProto) SetRestoredObjectsNetworkConfig(v RestoredObjectNetworkConfigProto)`

SetRestoredObjectsNetworkConfig sets RestoredObjectsNetworkConfig field to given value.

### HasRestoredObjectsNetworkConfig

`func (o *PerformRestoreTaskStateProto) HasRestoredObjectsNetworkConfig() bool`

HasRestoredObjectsNetworkConfig returns a boolean if a field has been set.

### GetRestoredToDifferentSource

`func (o *PerformRestoreTaskStateProto) GetRestoredToDifferentSource() bool`

GetRestoredToDifferentSource returns the RestoredToDifferentSource field if non-nil, zero value otherwise.

### GetRestoredToDifferentSourceOk

`func (o *PerformRestoreTaskStateProto) GetRestoredToDifferentSourceOk() (*bool, bool)`

GetRestoredToDifferentSourceOk returns a tuple with the RestoredToDifferentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoredToDifferentSource

`func (o *PerformRestoreTaskStateProto) SetRestoredToDifferentSource(v bool)`

SetRestoredToDifferentSource sets RestoredToDifferentSource field to given value.

### HasRestoredToDifferentSource

`func (o *PerformRestoreTaskStateProto) HasRestoredToDifferentSource() bool`

HasRestoredToDifferentSource returns a boolean if a field has been set.

### SetRestoredToDifferentSourceNil

`func (o *PerformRestoreTaskStateProto) SetRestoredToDifferentSourceNil(b bool)`

 SetRestoredToDifferentSourceNil sets the value for RestoredToDifferentSource to be an explicit nil

### UnsetRestoredToDifferentSource
`func (o *PerformRestoreTaskStateProto) UnsetRestoredToDifferentSource()`

UnsetRestoredToDifferentSource ensures that no value is present for RestoredToDifferentSource, not even an explicit nil
### GetRetrieveArchiveProgressMonitorTaskPath

`func (o *PerformRestoreTaskStateProto) GetRetrieveArchiveProgressMonitorTaskPath() string`

GetRetrieveArchiveProgressMonitorTaskPath returns the RetrieveArchiveProgressMonitorTaskPath field if non-nil, zero value otherwise.

### GetRetrieveArchiveProgressMonitorTaskPathOk

`func (o *PerformRestoreTaskStateProto) GetRetrieveArchiveProgressMonitorTaskPathOk() (*string, bool)`

GetRetrieveArchiveProgressMonitorTaskPathOk returns a tuple with the RetrieveArchiveProgressMonitorTaskPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrieveArchiveProgressMonitorTaskPath

`func (o *PerformRestoreTaskStateProto) SetRetrieveArchiveProgressMonitorTaskPath(v string)`

SetRetrieveArchiveProgressMonitorTaskPath sets RetrieveArchiveProgressMonitorTaskPath field to given value.

### HasRetrieveArchiveProgressMonitorTaskPath

`func (o *PerformRestoreTaskStateProto) HasRetrieveArchiveProgressMonitorTaskPath() bool`

HasRetrieveArchiveProgressMonitorTaskPath returns a boolean if a field has been set.

### SetRetrieveArchiveProgressMonitorTaskPathNil

`func (o *PerformRestoreTaskStateProto) SetRetrieveArchiveProgressMonitorTaskPathNil(b bool)`

 SetRetrieveArchiveProgressMonitorTaskPathNil sets the value for RetrieveArchiveProgressMonitorTaskPath to be an explicit nil

### UnsetRetrieveArchiveProgressMonitorTaskPath
`func (o *PerformRestoreTaskStateProto) UnsetRetrieveArchiveProgressMonitorTaskPath()`

UnsetRetrieveArchiveProgressMonitorTaskPath ensures that no value is present for RetrieveArchiveProgressMonitorTaskPath, not even an explicit nil
### GetRetrieveArchiveStubViewName

`func (o *PerformRestoreTaskStateProto) GetRetrieveArchiveStubViewName() string`

GetRetrieveArchiveStubViewName returns the RetrieveArchiveStubViewName field if non-nil, zero value otherwise.

### GetRetrieveArchiveStubViewNameOk

`func (o *PerformRestoreTaskStateProto) GetRetrieveArchiveStubViewNameOk() (*string, bool)`

GetRetrieveArchiveStubViewNameOk returns a tuple with the RetrieveArchiveStubViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrieveArchiveStubViewName

`func (o *PerformRestoreTaskStateProto) SetRetrieveArchiveStubViewName(v string)`

SetRetrieveArchiveStubViewName sets RetrieveArchiveStubViewName field to given value.

### HasRetrieveArchiveStubViewName

`func (o *PerformRestoreTaskStateProto) HasRetrieveArchiveStubViewName() bool`

HasRetrieveArchiveStubViewName returns a boolean if a field has been set.

### SetRetrieveArchiveStubViewNameNil

`func (o *PerformRestoreTaskStateProto) SetRetrieveArchiveStubViewNameNil(b bool)`

 SetRetrieveArchiveStubViewNameNil sets the value for RetrieveArchiveStubViewName to be an explicit nil

### UnsetRetrieveArchiveStubViewName
`func (o *PerformRestoreTaskStateProto) UnsetRetrieveArchiveStubViewName()`

UnsetRetrieveArchiveStubViewName ensures that no value is present for RetrieveArchiveStubViewName, not even an explicit nil
### GetRetrieveArchiveTaskUidVec

`func (o *PerformRestoreTaskStateProto) GetRetrieveArchiveTaskUidVec() []UniversalIdProto`

GetRetrieveArchiveTaskUidVec returns the RetrieveArchiveTaskUidVec field if non-nil, zero value otherwise.

### GetRetrieveArchiveTaskUidVecOk

`func (o *PerformRestoreTaskStateProto) GetRetrieveArchiveTaskUidVecOk() (*[]UniversalIdProto, bool)`

GetRetrieveArchiveTaskUidVecOk returns a tuple with the RetrieveArchiveTaskUidVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrieveArchiveTaskUidVec

`func (o *PerformRestoreTaskStateProto) SetRetrieveArchiveTaskUidVec(v []UniversalIdProto)`

SetRetrieveArchiveTaskUidVec sets RetrieveArchiveTaskUidVec field to given value.

### HasRetrieveArchiveTaskUidVec

`func (o *PerformRestoreTaskStateProto) HasRetrieveArchiveTaskUidVec() bool`

HasRetrieveArchiveTaskUidVec returns a boolean if a field has been set.

### SetRetrieveArchiveTaskUidVecNil

`func (o *PerformRestoreTaskStateProto) SetRetrieveArchiveTaskUidVecNil(b bool)`

 SetRetrieveArchiveTaskUidVecNil sets the value for RetrieveArchiveTaskUidVec to be an explicit nil

### UnsetRetrieveArchiveTaskUidVec
`func (o *PerformRestoreTaskStateProto) UnsetRetrieveArchiveTaskUidVec()`

UnsetRetrieveArchiveTaskUidVec ensures that no value is present for RetrieveArchiveTaskUidVec, not even an explicit nil
### GetRetrieveArchiveTaskVec

`func (o *PerformRestoreTaskStateProto) GetRetrieveArchiveTaskVec() []RetrieveArchiveTaskStateProto`

GetRetrieveArchiveTaskVec returns the RetrieveArchiveTaskVec field if non-nil, zero value otherwise.

### GetRetrieveArchiveTaskVecOk

`func (o *PerformRestoreTaskStateProto) GetRetrieveArchiveTaskVecOk() (*[]RetrieveArchiveTaskStateProto, bool)`

GetRetrieveArchiveTaskVecOk returns a tuple with the RetrieveArchiveTaskVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrieveArchiveTaskVec

`func (o *PerformRestoreTaskStateProto) SetRetrieveArchiveTaskVec(v []RetrieveArchiveTaskStateProto)`

SetRetrieveArchiveTaskVec sets RetrieveArchiveTaskVec field to given value.

### HasRetrieveArchiveTaskVec

`func (o *PerformRestoreTaskStateProto) HasRetrieveArchiveTaskVec() bool`

HasRetrieveArchiveTaskVec returns a boolean if a field has been set.

### SetRetrieveArchiveTaskVecNil

`func (o *PerformRestoreTaskStateProto) SetRetrieveArchiveTaskVecNil(b bool)`

 SetRetrieveArchiveTaskVecNil sets the value for RetrieveArchiveTaskVec to be an explicit nil

### UnsetRetrieveArchiveTaskVec
`func (o *PerformRestoreTaskStateProto) UnsetRetrieveArchiveTaskVec()`

UnsetRetrieveArchiveTaskVec ensures that no value is present for RetrieveArchiveTaskVec, not even an explicit nil
### GetRetrieveArchiveViewName

`func (o *PerformRestoreTaskStateProto) GetRetrieveArchiveViewName() string`

GetRetrieveArchiveViewName returns the RetrieveArchiveViewName field if non-nil, zero value otherwise.

### GetRetrieveArchiveViewNameOk

`func (o *PerformRestoreTaskStateProto) GetRetrieveArchiveViewNameOk() (*string, bool)`

GetRetrieveArchiveViewNameOk returns a tuple with the RetrieveArchiveViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrieveArchiveViewName

`func (o *PerformRestoreTaskStateProto) SetRetrieveArchiveViewName(v string)`

SetRetrieveArchiveViewName sets RetrieveArchiveViewName field to given value.

### HasRetrieveArchiveViewName

`func (o *PerformRestoreTaskStateProto) HasRetrieveArchiveViewName() bool`

HasRetrieveArchiveViewName returns a boolean if a field has been set.

### SetRetrieveArchiveViewNameNil

`func (o *PerformRestoreTaskStateProto) SetRetrieveArchiveViewNameNil(b bool)`

 SetRetrieveArchiveViewNameNil sets the value for RetrieveArchiveViewName to be an explicit nil

### UnsetRetrieveArchiveViewName
`func (o *PerformRestoreTaskStateProto) UnsetRetrieveArchiveViewName()`

UnsetRetrieveArchiveViewName ensures that no value is present for RetrieveArchiveViewName, not even an explicit nil
### GetSelectedDatastoreIdx

`func (o *PerformRestoreTaskStateProto) GetSelectedDatastoreIdx() int64`

GetSelectedDatastoreIdx returns the SelectedDatastoreIdx field if non-nil, zero value otherwise.

### GetSelectedDatastoreIdxOk

`func (o *PerformRestoreTaskStateProto) GetSelectedDatastoreIdxOk() (*int64, bool)`

GetSelectedDatastoreIdxOk returns a tuple with the SelectedDatastoreIdx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedDatastoreIdx

`func (o *PerformRestoreTaskStateProto) SetSelectedDatastoreIdx(v int64)`

SetSelectedDatastoreIdx sets SelectedDatastoreIdx field to given value.

### HasSelectedDatastoreIdx

`func (o *PerformRestoreTaskStateProto) HasSelectedDatastoreIdx() bool`

HasSelectedDatastoreIdx returns a boolean if a field has been set.

### SetSelectedDatastoreIdxNil

`func (o *PerformRestoreTaskStateProto) SetSelectedDatastoreIdxNil(b bool)`

 SetSelectedDatastoreIdxNil sets the value for SelectedDatastoreIdx to be an explicit nil

### UnsetSelectedDatastoreIdx
`func (o *PerformRestoreTaskStateProto) UnsetSelectedDatastoreIdx()`

UnsetSelectedDatastoreIdx ensures that no value is present for SelectedDatastoreIdx, not even an explicit nil
### GetStubViewRelativeDirName

`func (o *PerformRestoreTaskStateProto) GetStubViewRelativeDirName() string`

GetStubViewRelativeDirName returns the StubViewRelativeDirName field if non-nil, zero value otherwise.

### GetStubViewRelativeDirNameOk

`func (o *PerformRestoreTaskStateProto) GetStubViewRelativeDirNameOk() (*string, bool)`

GetStubViewRelativeDirNameOk returns a tuple with the StubViewRelativeDirName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStubViewRelativeDirName

`func (o *PerformRestoreTaskStateProto) SetStubViewRelativeDirName(v string)`

SetStubViewRelativeDirName sets StubViewRelativeDirName field to given value.

### HasStubViewRelativeDirName

`func (o *PerformRestoreTaskStateProto) HasStubViewRelativeDirName() bool`

HasStubViewRelativeDirName returns a boolean if a field has been set.

### SetStubViewRelativeDirNameNil

`func (o *PerformRestoreTaskStateProto) SetStubViewRelativeDirNameNil(b bool)`

 SetStubViewRelativeDirNameNil sets the value for StubViewRelativeDirName to be an explicit nil

### UnsetStubViewRelativeDirName
`func (o *PerformRestoreTaskStateProto) UnsetStubViewRelativeDirName()`

UnsetStubViewRelativeDirName ensures that no value is present for StubViewRelativeDirName, not even an explicit nil
### GetVaultRestoreParams

`func (o *PerformRestoreTaskStateProto) GetVaultRestoreParams() VaultParamsRestoreParams`

GetVaultRestoreParams returns the VaultRestoreParams field if non-nil, zero value otherwise.

### GetVaultRestoreParamsOk

`func (o *PerformRestoreTaskStateProto) GetVaultRestoreParamsOk() (*VaultParamsRestoreParams, bool)`

GetVaultRestoreParamsOk returns a tuple with the VaultRestoreParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultRestoreParams

`func (o *PerformRestoreTaskStateProto) SetVaultRestoreParams(v VaultParamsRestoreParams)`

SetVaultRestoreParams sets VaultRestoreParams field to given value.

### HasVaultRestoreParams

`func (o *PerformRestoreTaskStateProto) HasVaultRestoreParams() bool`

HasVaultRestoreParams returns a boolean if a field has been set.

### GetVcdConfig

`func (o *PerformRestoreTaskStateProto) GetVcdConfig() RestoredObjectVCDConfigProto`

GetVcdConfig returns the VcdConfig field if non-nil, zero value otherwise.

### GetVcdConfigOk

`func (o *PerformRestoreTaskStateProto) GetVcdConfigOk() (*RestoredObjectVCDConfigProto, bool)`

GetVcdConfigOk returns a tuple with the VcdConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcdConfig

`func (o *PerformRestoreTaskStateProto) SetVcdConfig(v RestoredObjectVCDConfigProto)`

SetVcdConfig sets VcdConfig field to given value.

### HasVcdConfig

`func (o *PerformRestoreTaskStateProto) HasVcdConfig() bool`

HasVcdConfig returns a boolean if a field has been set.

### GetVcdStorageProfileDatastoreMorefVec

`func (o *PerformRestoreTaskStateProto) GetVcdStorageProfileDatastoreMorefVec() []string`

GetVcdStorageProfileDatastoreMorefVec returns the VcdStorageProfileDatastoreMorefVec field if non-nil, zero value otherwise.

### GetVcdStorageProfileDatastoreMorefVecOk

`func (o *PerformRestoreTaskStateProto) GetVcdStorageProfileDatastoreMorefVecOk() (*[]string, bool)`

GetVcdStorageProfileDatastoreMorefVecOk returns a tuple with the VcdStorageProfileDatastoreMorefVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcdStorageProfileDatastoreMorefVec

`func (o *PerformRestoreTaskStateProto) SetVcdStorageProfileDatastoreMorefVec(v []string)`

SetVcdStorageProfileDatastoreMorefVec sets VcdStorageProfileDatastoreMorefVec field to given value.

### HasVcdStorageProfileDatastoreMorefVec

`func (o *PerformRestoreTaskStateProto) HasVcdStorageProfileDatastoreMorefVec() bool`

HasVcdStorageProfileDatastoreMorefVec returns a boolean if a field has been set.

### SetVcdStorageProfileDatastoreMorefVecNil

`func (o *PerformRestoreTaskStateProto) SetVcdStorageProfileDatastoreMorefVecNil(b bool)`

 SetVcdStorageProfileDatastoreMorefVecNil sets the value for VcdStorageProfileDatastoreMorefVec to be an explicit nil

### UnsetVcdStorageProfileDatastoreMorefVec
`func (o *PerformRestoreTaskStateProto) UnsetVcdStorageProfileDatastoreMorefVec()`

UnsetVcdStorageProfileDatastoreMorefVec ensures that no value is present for VcdStorageProfileDatastoreMorefVec, not even an explicit nil
### GetViewBoxId

`func (o *PerformRestoreTaskStateProto) GetViewBoxId() int64`

GetViewBoxId returns the ViewBoxId field if non-nil, zero value otherwise.

### GetViewBoxIdOk

`func (o *PerformRestoreTaskStateProto) GetViewBoxIdOk() (*int64, bool)`

GetViewBoxIdOk returns a tuple with the ViewBoxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewBoxId

`func (o *PerformRestoreTaskStateProto) SetViewBoxId(v int64)`

SetViewBoxId sets ViewBoxId field to given value.

### HasViewBoxId

`func (o *PerformRestoreTaskStateProto) HasViewBoxId() bool`

HasViewBoxId returns a boolean if a field has been set.

### SetViewBoxIdNil

`func (o *PerformRestoreTaskStateProto) SetViewBoxIdNil(b bool)`

 SetViewBoxIdNil sets the value for ViewBoxId to be an explicit nil

### UnsetViewBoxId
`func (o *PerformRestoreTaskStateProto) UnsetViewBoxId()`

UnsetViewBoxId ensures that no value is present for ViewBoxId, not even an explicit nil
### GetViewNameDEPRECATED

`func (o *PerformRestoreTaskStateProto) GetViewNameDEPRECATED() string`

GetViewNameDEPRECATED returns the ViewNameDEPRECATED field if non-nil, zero value otherwise.

### GetViewNameDEPRECATEDOk

`func (o *PerformRestoreTaskStateProto) GetViewNameDEPRECATEDOk() (*string, bool)`

GetViewNameDEPRECATEDOk returns a tuple with the ViewNameDEPRECATED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewNameDEPRECATED

`func (o *PerformRestoreTaskStateProto) SetViewNameDEPRECATED(v string)`

SetViewNameDEPRECATED sets ViewNameDEPRECATED field to given value.

### HasViewNameDEPRECATED

`func (o *PerformRestoreTaskStateProto) HasViewNameDEPRECATED() bool`

HasViewNameDEPRECATED returns a boolean if a field has been set.

### SetViewNameDEPRECATEDNil

`func (o *PerformRestoreTaskStateProto) SetViewNameDEPRECATEDNil(b bool)`

 SetViewNameDEPRECATEDNil sets the value for ViewNameDEPRECATED to be an explicit nil

### UnsetViewNameDEPRECATED
`func (o *PerformRestoreTaskStateProto) UnsetViewNameDEPRECATED()`

UnsetViewNameDEPRECATED ensures that no value is present for ViewNameDEPRECATED, not even an explicit nil
### GetViewParams

`func (o *PerformRestoreTaskStateProto) GetViewParams() ViewParams`

GetViewParams returns the ViewParams field if non-nil, zero value otherwise.

### GetViewParamsOk

`func (o *PerformRestoreTaskStateProto) GetViewParamsOk() (*ViewParams, bool)`

GetViewParamsOk returns a tuple with the ViewParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewParams

`func (o *PerformRestoreTaskStateProto) SetViewParams(v ViewParams)`

SetViewParams sets ViewParams field to given value.

### HasViewParams

`func (o *PerformRestoreTaskStateProto) HasViewParams() bool`

HasViewParams returns a boolean if a field has been set.

### GetVolumeInfoVec

`func (o *PerformRestoreTaskStateProto) GetVolumeInfoVec() []VolumeInfo`

GetVolumeInfoVec returns the VolumeInfoVec field if non-nil, zero value otherwise.

### GetVolumeInfoVecOk

`func (o *PerformRestoreTaskStateProto) GetVolumeInfoVecOk() (*[]VolumeInfo, bool)`

GetVolumeInfoVecOk returns a tuple with the VolumeInfoVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeInfoVec

`func (o *PerformRestoreTaskStateProto) SetVolumeInfoVec(v []VolumeInfo)`

SetVolumeInfoVec sets VolumeInfoVec field to given value.

### HasVolumeInfoVec

`func (o *PerformRestoreTaskStateProto) HasVolumeInfoVec() bool`

HasVolumeInfoVec returns a boolean if a field has been set.

### SetVolumeInfoVecNil

`func (o *PerformRestoreTaskStateProto) SetVolumeInfoVecNil(b bool)`

 SetVolumeInfoVecNil sets the value for VolumeInfoVec to be an explicit nil

### UnsetVolumeInfoVec
`func (o *PerformRestoreTaskStateProto) UnsetVolumeInfoVec()`

UnsetVolumeInfoVec ensures that no value is present for VolumeInfoVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


