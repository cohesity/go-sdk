# RestoreTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcropolisParameters** | Pointer to [**AcropolisRestoreParameters**](AcropolisRestoreParameters.md) |  | [optional] 
**ApplicationParameters** | Pointer to [**ApplicationRestoreParameters**](ApplicationRestoreParameters.md) |  | [optional] 
**ArchiveTaskUid** | Pointer to [**NullableUniversalId**](UniversalId.md) | Specifies the uid of the Restore Task that retrieves objects from an archive. This field is only populated when objects must be retrieved from an archive before being restored. This field is deprecated. deprecated:true | [optional] 
**ArchiveTaskUids** | Pointer to [**[]UniversalId**](UniversalId.md) | Specifies the uids of the Restore Task that retrieves objects from an archive. This field is only populated when objects must be retrieved from an archive before being restored. overrideDescription:true | [optional] 
**CanTearDown** | Pointer to **NullableBool** | Specifies whether it is possible to tear down the objects created by the recovery. | [optional] 
**CloneViewParameters** | Pointer to [**NullableUpdateViewParam**](UpdateViewParam.md) | Specifies the View settings used when cloning a View. | [optional] 
**ContinueOnError** | Pointer to **NullableBool** | Specifies if the Restore Task should continue when some operations on some objects fail. If true, the Cohesity Cluster ignores intermittent errors and restores as many objects as possible. | [optional] 
**DatastoreId** | Pointer to **NullableInt64** | Specifies the datastore where the object&#39;s files are recovered to. This field is populated when objects are recovered to a different resource pool or to a different parent source. This field is not populated when objects are recovered to their original datastore locations in the original parent source. | [optional] 
**DeployVmsToCloud** | Pointer to [**DeployVmsToCloud**](DeployVmsToCloud.md) |  | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of the Restore Task as a Unix epoch Timestamp (in microseconds). This field is only populated if the Restore Task completes. | [optional] 
**Error** | Pointer to [**NullableRequestError**](RequestError.md) | Specifies the error reported by the Restore Task (if any) after the Task has finished. | [optional] 
**FullViewName** | Pointer to **NullableString** | Specifies the full name of a View. | [optional] 
**HypervParameters** | Pointer to [**HypervRestoreParameters**](HypervRestoreParameters.md) |  | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the id of the Restore Task assigned by Cohesity Cluster. | [optional] 
**MountVolumesState** | Pointer to [**MountVolumesState**](MountVolumesState.md) |  | [optional] 
**Name** | **NullableString** | Specifies the name of the Restore Task. This field must be set and must be a unique name. | 
**NewParentId** | Pointer to **NullableInt64** | Specify a new registered parent Protection Source. If specified the selected objects are cloned or recovered to this new Protection Source. If not specified, objects are cloned or recovered to the original Protection Source that was managing them. | [optional] 
**Objects** | Pointer to [**[]RestoreObjectDetails**](RestoreObjectDetails.md) | Array of Objects.  Specifies a list of Protection Source objects or Protection Job objects (with specified Protection Source objects). | [optional] 
**OneDriveParameters** | Pointer to [**OneDriveRestoreParameters**](OneDriveRestoreParameters.md) |  | [optional] 
**OutlookParameters** | Pointer to [**OutlookRestoreParameters**](OutlookRestoreParameters.md) |  | [optional] 
**PublicFoldersParameters** | Pointer to [**PublicFoldersRestoreParameters**](PublicFoldersRestoreParameters.md) |  | [optional] 
**RestoreObjectState** | Pointer to [**[]RestoreObjectState**](RestoreObjectState.md) | Array of Object States.  Specifies the states of all the objects for the &#39;kRecoverVMs&#39; and &#39;kCloneVMs&#39; Restore Tasks. | [optional] 
**SharePointParameters** | Pointer to [**SharePointRestoreParameters**](SharePointRestoreParameters.md) |  | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time for the Restore Task as a Unix epoch Timestamp (in microseconds). | [optional] 
**Status** | Pointer to **NullableString** | Specifies the overall status of the Restore Task. &#39;kReadyToSchedule&#39; indicates the Restore Task is waiting to be scheduled. &#39;kProgressMonitorCreated&#39; indicates the progress monitor for the Restore Task has been created. &#39;kRetrievedFromArchive&#39; indicates that the objects to restore have been retrieved from the specified archive. A Task will only ever transition to this state if a retrieval is necessary. &#39;kAdmitted&#39; indicates the task has been admitted. After a task has been admitted, its status does not move back to &#39;kReadyToSchedule&#39; state even if it is rescheduled. &#39;kInProgress&#39; indicates that the Restore Task is in progress. &#39;kFinishingProgressMonitor&#39; indicates that the Restore Task is finishing its progress monitoring. &#39;kFinished&#39; indicates that the Restore Task has finished. The status indicating success or failure is found in the error code that is stored with the Restore Task. &#39;kInternalViewCreated&#39; indicates that internal view for the task has been created. &#39;kZipFileRequested&#39; indicates that request has been sent to create zip files for the files to be downloaded. This state is only going to be present for kDownloadFiles Task. &#39;kCancelled&#39; indicates that task or jb has been cancelled. | [optional] 
**TargetViewCreated** | Pointer to **NullableBool** | Is true if a new View was created by a &#39;kCloneVMs&#39; Restore Task. This field is only set for a &#39;kCloneVMs&#39; Restore Task. | [optional] 
**TearDownErrorMessage** | Pointer to **NullableString** | Specifies the error message about the tear down operation if it fails. | [optional] 
**TearDownStatus** | Pointer to **NullableString** | Specifies the status of the tear down operation. This is only set when the field &#39;CanTearDown&#39; is set to true. &#39;kReadyToSchedule&#39; indicates that the task is waiting to be scheduled. &#39;kAdmitted&#39; indicates that the task has been admitted. &#39;kFinished&#39; indicates that the task is finished with or without error. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of Restore Task.  &#39;kRecoverVMs&#39; specifies a Restore Task that recovers VMs. &#39;kCloneVMs&#39; specifies a Restore Task that clones VMs. &#39;kCloneView&#39; specifies a Restore Task that clones a View. &#39;kMountVolumes&#39; specifies a Restore Task that mounts volumes. &#39;kRestoreFiles&#39; specifies a Restore Task that recovers files and folders. &#39;kRecoverApp&#39; specifies a Restore Task that recovers app. &#39;kCloneApp&#39; specifies a Restore Task that clone app. &#39;kRecoverSanVolume&#39; specifies a Restore Task that recovers SAN volumes. &#39;kConvertAndDeployVMs&#39; specifies a Restore Task that converts and deploy VMs to a target environment. &#39;kMountFileVolume&#39; specifies a Restore Task that mounts a file volume. &#39;kSystem&#39; specifies a Restore Task that recovers a system. &#39;kRecoverVolumes&#39; specifies a Restore Task that recovers volumes via the physical agent. &#39;kDeployVolumes&#39; specifies a Restore Task that deploys volumes to a target environment. &#39;kDownloadFiles&#39; specifies a Restore Task that downloads the requested files and folders in zip format. &#39;kRecoverEmails&#39; specifies a Restore Task that recovers the mailbox/email items. &#39;kRecoverDisks&#39; specifies a Restore Task that recovers the virtual disks. &#39;kRecoverNamespaces&#39; specifies a Restore Task that recovers Kubernetes namespaces. &#39;kCloneVMsToView&#39; specifies a Restore Task that clones VMs into a View. | [optional] 
**Username** | Pointer to **NullableString** | Specifies the Cohesity user who requested this Restore Task. | [optional] 
**ViewBoxId** | Pointer to **NullableInt64** | Specifies the id of the Domain (View Box) where the View is stored. | [optional] 
**VirtualDiskRestoreState** | Pointer to [**VirtualDiskRecoverTaskState**](VirtualDiskRecoverTaskState.md) |  | [optional] 
**VlanParameters** | Pointer to [**VlanParameters**](VlanParameters.md) |  | [optional] 
**VmwareParameters** | Pointer to [**VmwareRestoreParameters**](VmwareRestoreParameters.md) |  | [optional] 

## Methods

### NewRestoreTask

`func NewRestoreTask(name NullableString, ) *RestoreTask`

NewRestoreTask instantiates a new RestoreTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreTaskWithDefaults

`func NewRestoreTaskWithDefaults() *RestoreTask`

NewRestoreTaskWithDefaults instantiates a new RestoreTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcropolisParameters

`func (o *RestoreTask) GetAcropolisParameters() AcropolisRestoreParameters`

GetAcropolisParameters returns the AcropolisParameters field if non-nil, zero value otherwise.

### GetAcropolisParametersOk

`func (o *RestoreTask) GetAcropolisParametersOk() (*AcropolisRestoreParameters, bool)`

GetAcropolisParametersOk returns a tuple with the AcropolisParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcropolisParameters

`func (o *RestoreTask) SetAcropolisParameters(v AcropolisRestoreParameters)`

SetAcropolisParameters sets AcropolisParameters field to given value.

### HasAcropolisParameters

`func (o *RestoreTask) HasAcropolisParameters() bool`

HasAcropolisParameters returns a boolean if a field has been set.

### GetApplicationParameters

`func (o *RestoreTask) GetApplicationParameters() ApplicationRestoreParameters`

GetApplicationParameters returns the ApplicationParameters field if non-nil, zero value otherwise.

### GetApplicationParametersOk

`func (o *RestoreTask) GetApplicationParametersOk() (*ApplicationRestoreParameters, bool)`

GetApplicationParametersOk returns a tuple with the ApplicationParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationParameters

`func (o *RestoreTask) SetApplicationParameters(v ApplicationRestoreParameters)`

SetApplicationParameters sets ApplicationParameters field to given value.

### HasApplicationParameters

`func (o *RestoreTask) HasApplicationParameters() bool`

HasApplicationParameters returns a boolean if a field has been set.

### GetArchiveTaskUid

`func (o *RestoreTask) GetArchiveTaskUid() UniversalId`

GetArchiveTaskUid returns the ArchiveTaskUid field if non-nil, zero value otherwise.

### GetArchiveTaskUidOk

`func (o *RestoreTask) GetArchiveTaskUidOk() (*UniversalId, bool)`

GetArchiveTaskUidOk returns a tuple with the ArchiveTaskUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveTaskUid

`func (o *RestoreTask) SetArchiveTaskUid(v UniversalId)`

SetArchiveTaskUid sets ArchiveTaskUid field to given value.

### HasArchiveTaskUid

`func (o *RestoreTask) HasArchiveTaskUid() bool`

HasArchiveTaskUid returns a boolean if a field has been set.

### SetArchiveTaskUidNil

`func (o *RestoreTask) SetArchiveTaskUidNil(b bool)`

 SetArchiveTaskUidNil sets the value for ArchiveTaskUid to be an explicit nil

### UnsetArchiveTaskUid
`func (o *RestoreTask) UnsetArchiveTaskUid()`

UnsetArchiveTaskUid ensures that no value is present for ArchiveTaskUid, not even an explicit nil
### GetArchiveTaskUids

`func (o *RestoreTask) GetArchiveTaskUids() []UniversalId`

GetArchiveTaskUids returns the ArchiveTaskUids field if non-nil, zero value otherwise.

### GetArchiveTaskUidsOk

`func (o *RestoreTask) GetArchiveTaskUidsOk() (*[]UniversalId, bool)`

GetArchiveTaskUidsOk returns a tuple with the ArchiveTaskUids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveTaskUids

`func (o *RestoreTask) SetArchiveTaskUids(v []UniversalId)`

SetArchiveTaskUids sets ArchiveTaskUids field to given value.

### HasArchiveTaskUids

`func (o *RestoreTask) HasArchiveTaskUids() bool`

HasArchiveTaskUids returns a boolean if a field has been set.

### SetArchiveTaskUidsNil

`func (o *RestoreTask) SetArchiveTaskUidsNil(b bool)`

 SetArchiveTaskUidsNil sets the value for ArchiveTaskUids to be an explicit nil

### UnsetArchiveTaskUids
`func (o *RestoreTask) UnsetArchiveTaskUids()`

UnsetArchiveTaskUids ensures that no value is present for ArchiveTaskUids, not even an explicit nil
### GetCanTearDown

`func (o *RestoreTask) GetCanTearDown() bool`

GetCanTearDown returns the CanTearDown field if non-nil, zero value otherwise.

### GetCanTearDownOk

`func (o *RestoreTask) GetCanTearDownOk() (*bool, bool)`

GetCanTearDownOk returns a tuple with the CanTearDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTearDown

`func (o *RestoreTask) SetCanTearDown(v bool)`

SetCanTearDown sets CanTearDown field to given value.

### HasCanTearDown

`func (o *RestoreTask) HasCanTearDown() bool`

HasCanTearDown returns a boolean if a field has been set.

### SetCanTearDownNil

`func (o *RestoreTask) SetCanTearDownNil(b bool)`

 SetCanTearDownNil sets the value for CanTearDown to be an explicit nil

### UnsetCanTearDown
`func (o *RestoreTask) UnsetCanTearDown()`

UnsetCanTearDown ensures that no value is present for CanTearDown, not even an explicit nil
### GetCloneViewParameters

`func (o *RestoreTask) GetCloneViewParameters() UpdateViewParam`

GetCloneViewParameters returns the CloneViewParameters field if non-nil, zero value otherwise.

### GetCloneViewParametersOk

`func (o *RestoreTask) GetCloneViewParametersOk() (*UpdateViewParam, bool)`

GetCloneViewParametersOk returns a tuple with the CloneViewParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneViewParameters

`func (o *RestoreTask) SetCloneViewParameters(v UpdateViewParam)`

SetCloneViewParameters sets CloneViewParameters field to given value.

### HasCloneViewParameters

`func (o *RestoreTask) HasCloneViewParameters() bool`

HasCloneViewParameters returns a boolean if a field has been set.

### SetCloneViewParametersNil

`func (o *RestoreTask) SetCloneViewParametersNil(b bool)`

 SetCloneViewParametersNil sets the value for CloneViewParameters to be an explicit nil

### UnsetCloneViewParameters
`func (o *RestoreTask) UnsetCloneViewParameters()`

UnsetCloneViewParameters ensures that no value is present for CloneViewParameters, not even an explicit nil
### GetContinueOnError

`func (o *RestoreTask) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RestoreTask) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RestoreTask) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RestoreTask) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RestoreTask) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RestoreTask) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetDatastoreId

`func (o *RestoreTask) GetDatastoreId() int64`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *RestoreTask) GetDatastoreIdOk() (*int64, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *RestoreTask) SetDatastoreId(v int64)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *RestoreTask) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.

### SetDatastoreIdNil

`func (o *RestoreTask) SetDatastoreIdNil(b bool)`

 SetDatastoreIdNil sets the value for DatastoreId to be an explicit nil

### UnsetDatastoreId
`func (o *RestoreTask) UnsetDatastoreId()`

UnsetDatastoreId ensures that no value is present for DatastoreId, not even an explicit nil
### GetDeployVmsToCloud

`func (o *RestoreTask) GetDeployVmsToCloud() DeployVmsToCloud`

GetDeployVmsToCloud returns the DeployVmsToCloud field if non-nil, zero value otherwise.

### GetDeployVmsToCloudOk

`func (o *RestoreTask) GetDeployVmsToCloudOk() (*DeployVmsToCloud, bool)`

GetDeployVmsToCloudOk returns a tuple with the DeployVmsToCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployVmsToCloud

`func (o *RestoreTask) SetDeployVmsToCloud(v DeployVmsToCloud)`

SetDeployVmsToCloud sets DeployVmsToCloud field to given value.

### HasDeployVmsToCloud

`func (o *RestoreTask) HasDeployVmsToCloud() bool`

HasDeployVmsToCloud returns a boolean if a field has been set.

### GetEndTimeUsecs

`func (o *RestoreTask) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *RestoreTask) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *RestoreTask) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *RestoreTask) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *RestoreTask) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *RestoreTask) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetError

`func (o *RestoreTask) GetError() RequestError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RestoreTask) GetErrorOk() (*RequestError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RestoreTask) SetError(v RequestError)`

SetError sets Error field to given value.

### HasError

`func (o *RestoreTask) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *RestoreTask) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *RestoreTask) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetFullViewName

`func (o *RestoreTask) GetFullViewName() string`

GetFullViewName returns the FullViewName field if non-nil, zero value otherwise.

### GetFullViewNameOk

`func (o *RestoreTask) GetFullViewNameOk() (*string, bool)`

GetFullViewNameOk returns a tuple with the FullViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullViewName

`func (o *RestoreTask) SetFullViewName(v string)`

SetFullViewName sets FullViewName field to given value.

### HasFullViewName

`func (o *RestoreTask) HasFullViewName() bool`

HasFullViewName returns a boolean if a field has been set.

### SetFullViewNameNil

`func (o *RestoreTask) SetFullViewNameNil(b bool)`

 SetFullViewNameNil sets the value for FullViewName to be an explicit nil

### UnsetFullViewName
`func (o *RestoreTask) UnsetFullViewName()`

UnsetFullViewName ensures that no value is present for FullViewName, not even an explicit nil
### GetHypervParameters

`func (o *RestoreTask) GetHypervParameters() HypervRestoreParameters`

GetHypervParameters returns the HypervParameters field if non-nil, zero value otherwise.

### GetHypervParametersOk

`func (o *RestoreTask) GetHypervParametersOk() (*HypervRestoreParameters, bool)`

GetHypervParametersOk returns a tuple with the HypervParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervParameters

`func (o *RestoreTask) SetHypervParameters(v HypervRestoreParameters)`

SetHypervParameters sets HypervParameters field to given value.

### HasHypervParameters

`func (o *RestoreTask) HasHypervParameters() bool`

HasHypervParameters returns a boolean if a field has been set.

### GetId

`func (o *RestoreTask) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestoreTask) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestoreTask) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RestoreTask) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *RestoreTask) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RestoreTask) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetMountVolumesState

`func (o *RestoreTask) GetMountVolumesState() MountVolumesState`

GetMountVolumesState returns the MountVolumesState field if non-nil, zero value otherwise.

### GetMountVolumesStateOk

`func (o *RestoreTask) GetMountVolumesStateOk() (*MountVolumesState, bool)`

GetMountVolumesStateOk returns a tuple with the MountVolumesState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountVolumesState

`func (o *RestoreTask) SetMountVolumesState(v MountVolumesState)`

SetMountVolumesState sets MountVolumesState field to given value.

### HasMountVolumesState

`func (o *RestoreTask) HasMountVolumesState() bool`

HasMountVolumesState returns a boolean if a field has been set.

### GetName

`func (o *RestoreTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestoreTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestoreTask) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *RestoreTask) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RestoreTask) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNewParentId

`func (o *RestoreTask) GetNewParentId() int64`

GetNewParentId returns the NewParentId field if non-nil, zero value otherwise.

### GetNewParentIdOk

`func (o *RestoreTask) GetNewParentIdOk() (*int64, bool)`

GetNewParentIdOk returns a tuple with the NewParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewParentId

`func (o *RestoreTask) SetNewParentId(v int64)`

SetNewParentId sets NewParentId field to given value.

### HasNewParentId

`func (o *RestoreTask) HasNewParentId() bool`

HasNewParentId returns a boolean if a field has been set.

### SetNewParentIdNil

`func (o *RestoreTask) SetNewParentIdNil(b bool)`

 SetNewParentIdNil sets the value for NewParentId to be an explicit nil

### UnsetNewParentId
`func (o *RestoreTask) UnsetNewParentId()`

UnsetNewParentId ensures that no value is present for NewParentId, not even an explicit nil
### GetObjects

`func (o *RestoreTask) GetObjects() []RestoreObjectDetails`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RestoreTask) GetObjectsOk() (*[]RestoreObjectDetails, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RestoreTask) SetObjects(v []RestoreObjectDetails)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *RestoreTask) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *RestoreTask) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RestoreTask) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetOneDriveParameters

`func (o *RestoreTask) GetOneDriveParameters() OneDriveRestoreParameters`

GetOneDriveParameters returns the OneDriveParameters field if non-nil, zero value otherwise.

### GetOneDriveParametersOk

`func (o *RestoreTask) GetOneDriveParametersOk() (*OneDriveRestoreParameters, bool)`

GetOneDriveParametersOk returns a tuple with the OneDriveParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneDriveParameters

`func (o *RestoreTask) SetOneDriveParameters(v OneDriveRestoreParameters)`

SetOneDriveParameters sets OneDriveParameters field to given value.

### HasOneDriveParameters

`func (o *RestoreTask) HasOneDriveParameters() bool`

HasOneDriveParameters returns a boolean if a field has been set.

### GetOutlookParameters

`func (o *RestoreTask) GetOutlookParameters() OutlookRestoreParameters`

GetOutlookParameters returns the OutlookParameters field if non-nil, zero value otherwise.

### GetOutlookParametersOk

`func (o *RestoreTask) GetOutlookParametersOk() (*OutlookRestoreParameters, bool)`

GetOutlookParametersOk returns a tuple with the OutlookParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutlookParameters

`func (o *RestoreTask) SetOutlookParameters(v OutlookRestoreParameters)`

SetOutlookParameters sets OutlookParameters field to given value.

### HasOutlookParameters

`func (o *RestoreTask) HasOutlookParameters() bool`

HasOutlookParameters returns a boolean if a field has been set.

### GetPublicFoldersParameters

`func (o *RestoreTask) GetPublicFoldersParameters() PublicFoldersRestoreParameters`

GetPublicFoldersParameters returns the PublicFoldersParameters field if non-nil, zero value otherwise.

### GetPublicFoldersParametersOk

`func (o *RestoreTask) GetPublicFoldersParametersOk() (*PublicFoldersRestoreParameters, bool)`

GetPublicFoldersParametersOk returns a tuple with the PublicFoldersParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicFoldersParameters

`func (o *RestoreTask) SetPublicFoldersParameters(v PublicFoldersRestoreParameters)`

SetPublicFoldersParameters sets PublicFoldersParameters field to given value.

### HasPublicFoldersParameters

`func (o *RestoreTask) HasPublicFoldersParameters() bool`

HasPublicFoldersParameters returns a boolean if a field has been set.

### GetRestoreObjectState

`func (o *RestoreTask) GetRestoreObjectState() []RestoreObjectState`

GetRestoreObjectState returns the RestoreObjectState field if non-nil, zero value otherwise.

### GetRestoreObjectStateOk

`func (o *RestoreTask) GetRestoreObjectStateOk() (*[]RestoreObjectState, bool)`

GetRestoreObjectStateOk returns a tuple with the RestoreObjectState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreObjectState

`func (o *RestoreTask) SetRestoreObjectState(v []RestoreObjectState)`

SetRestoreObjectState sets RestoreObjectState field to given value.

### HasRestoreObjectState

`func (o *RestoreTask) HasRestoreObjectState() bool`

HasRestoreObjectState returns a boolean if a field has been set.

### SetRestoreObjectStateNil

`func (o *RestoreTask) SetRestoreObjectStateNil(b bool)`

 SetRestoreObjectStateNil sets the value for RestoreObjectState to be an explicit nil

### UnsetRestoreObjectState
`func (o *RestoreTask) UnsetRestoreObjectState()`

UnsetRestoreObjectState ensures that no value is present for RestoreObjectState, not even an explicit nil
### GetSharePointParameters

`func (o *RestoreTask) GetSharePointParameters() SharePointRestoreParameters`

GetSharePointParameters returns the SharePointParameters field if non-nil, zero value otherwise.

### GetSharePointParametersOk

`func (o *RestoreTask) GetSharePointParametersOk() (*SharePointRestoreParameters, bool)`

GetSharePointParametersOk returns a tuple with the SharePointParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePointParameters

`func (o *RestoreTask) SetSharePointParameters(v SharePointRestoreParameters)`

SetSharePointParameters sets SharePointParameters field to given value.

### HasSharePointParameters

`func (o *RestoreTask) HasSharePointParameters() bool`

HasSharePointParameters returns a boolean if a field has been set.

### GetStartTimeUsecs

`func (o *RestoreTask) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *RestoreTask) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *RestoreTask) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *RestoreTask) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *RestoreTask) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *RestoreTask) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetStatus

`func (o *RestoreTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RestoreTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RestoreTask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RestoreTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *RestoreTask) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *RestoreTask) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTargetViewCreated

`func (o *RestoreTask) GetTargetViewCreated() bool`

GetTargetViewCreated returns the TargetViewCreated field if non-nil, zero value otherwise.

### GetTargetViewCreatedOk

`func (o *RestoreTask) GetTargetViewCreatedOk() (*bool, bool)`

GetTargetViewCreatedOk returns a tuple with the TargetViewCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetViewCreated

`func (o *RestoreTask) SetTargetViewCreated(v bool)`

SetTargetViewCreated sets TargetViewCreated field to given value.

### HasTargetViewCreated

`func (o *RestoreTask) HasTargetViewCreated() bool`

HasTargetViewCreated returns a boolean if a field has been set.

### SetTargetViewCreatedNil

`func (o *RestoreTask) SetTargetViewCreatedNil(b bool)`

 SetTargetViewCreatedNil sets the value for TargetViewCreated to be an explicit nil

### UnsetTargetViewCreated
`func (o *RestoreTask) UnsetTargetViewCreated()`

UnsetTargetViewCreated ensures that no value is present for TargetViewCreated, not even an explicit nil
### GetTearDownErrorMessage

`func (o *RestoreTask) GetTearDownErrorMessage() string`

GetTearDownErrorMessage returns the TearDownErrorMessage field if non-nil, zero value otherwise.

### GetTearDownErrorMessageOk

`func (o *RestoreTask) GetTearDownErrorMessageOk() (*string, bool)`

GetTearDownErrorMessageOk returns a tuple with the TearDownErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTearDownErrorMessage

`func (o *RestoreTask) SetTearDownErrorMessage(v string)`

SetTearDownErrorMessage sets TearDownErrorMessage field to given value.

### HasTearDownErrorMessage

`func (o *RestoreTask) HasTearDownErrorMessage() bool`

HasTearDownErrorMessage returns a boolean if a field has been set.

### SetTearDownErrorMessageNil

`func (o *RestoreTask) SetTearDownErrorMessageNil(b bool)`

 SetTearDownErrorMessageNil sets the value for TearDownErrorMessage to be an explicit nil

### UnsetTearDownErrorMessage
`func (o *RestoreTask) UnsetTearDownErrorMessage()`

UnsetTearDownErrorMessage ensures that no value is present for TearDownErrorMessage, not even an explicit nil
### GetTearDownStatus

`func (o *RestoreTask) GetTearDownStatus() string`

GetTearDownStatus returns the TearDownStatus field if non-nil, zero value otherwise.

### GetTearDownStatusOk

`func (o *RestoreTask) GetTearDownStatusOk() (*string, bool)`

GetTearDownStatusOk returns a tuple with the TearDownStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTearDownStatus

`func (o *RestoreTask) SetTearDownStatus(v string)`

SetTearDownStatus sets TearDownStatus field to given value.

### HasTearDownStatus

`func (o *RestoreTask) HasTearDownStatus() bool`

HasTearDownStatus returns a boolean if a field has been set.

### SetTearDownStatusNil

`func (o *RestoreTask) SetTearDownStatusNil(b bool)`

 SetTearDownStatusNil sets the value for TearDownStatus to be an explicit nil

### UnsetTearDownStatus
`func (o *RestoreTask) UnsetTearDownStatus()`

UnsetTearDownStatus ensures that no value is present for TearDownStatus, not even an explicit nil
### GetType

`func (o *RestoreTask) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestoreTask) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestoreTask) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RestoreTask) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *RestoreTask) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *RestoreTask) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUsername

`func (o *RestoreTask) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RestoreTask) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RestoreTask) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RestoreTask) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *RestoreTask) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *RestoreTask) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetViewBoxId

`func (o *RestoreTask) GetViewBoxId() int64`

GetViewBoxId returns the ViewBoxId field if non-nil, zero value otherwise.

### GetViewBoxIdOk

`func (o *RestoreTask) GetViewBoxIdOk() (*int64, bool)`

GetViewBoxIdOk returns a tuple with the ViewBoxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewBoxId

`func (o *RestoreTask) SetViewBoxId(v int64)`

SetViewBoxId sets ViewBoxId field to given value.

### HasViewBoxId

`func (o *RestoreTask) HasViewBoxId() bool`

HasViewBoxId returns a boolean if a field has been set.

### SetViewBoxIdNil

`func (o *RestoreTask) SetViewBoxIdNil(b bool)`

 SetViewBoxIdNil sets the value for ViewBoxId to be an explicit nil

### UnsetViewBoxId
`func (o *RestoreTask) UnsetViewBoxId()`

UnsetViewBoxId ensures that no value is present for ViewBoxId, not even an explicit nil
### GetVirtualDiskRestoreState

`func (o *RestoreTask) GetVirtualDiskRestoreState() VirtualDiskRecoverTaskState`

GetVirtualDiskRestoreState returns the VirtualDiskRestoreState field if non-nil, zero value otherwise.

### GetVirtualDiskRestoreStateOk

`func (o *RestoreTask) GetVirtualDiskRestoreStateOk() (*VirtualDiskRecoverTaskState, bool)`

GetVirtualDiskRestoreStateOk returns a tuple with the VirtualDiskRestoreState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDiskRestoreState

`func (o *RestoreTask) SetVirtualDiskRestoreState(v VirtualDiskRecoverTaskState)`

SetVirtualDiskRestoreState sets VirtualDiskRestoreState field to given value.

### HasVirtualDiskRestoreState

`func (o *RestoreTask) HasVirtualDiskRestoreState() bool`

HasVirtualDiskRestoreState returns a boolean if a field has been set.

### GetVlanParameters

`func (o *RestoreTask) GetVlanParameters() VlanParameters`

GetVlanParameters returns the VlanParameters field if non-nil, zero value otherwise.

### GetVlanParametersOk

`func (o *RestoreTask) GetVlanParametersOk() (*VlanParameters, bool)`

GetVlanParametersOk returns a tuple with the VlanParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanParameters

`func (o *RestoreTask) SetVlanParameters(v VlanParameters)`

SetVlanParameters sets VlanParameters field to given value.

### HasVlanParameters

`func (o *RestoreTask) HasVlanParameters() bool`

HasVlanParameters returns a boolean if a field has been set.

### GetVmwareParameters

`func (o *RestoreTask) GetVmwareParameters() VmwareRestoreParameters`

GetVmwareParameters returns the VmwareParameters field if non-nil, zero value otherwise.

### GetVmwareParametersOk

`func (o *RestoreTask) GetVmwareParametersOk() (*VmwareRestoreParameters, bool)`

GetVmwareParametersOk returns a tuple with the VmwareParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParameters

`func (o *RestoreTask) SetVmwareParameters(v VmwareRestoreParameters)`

SetVmwareParameters sets VmwareParameters field to given value.

### HasVmwareParameters

`func (o *RestoreTask) HasVmwareParameters() bool`

HasVmwareParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


