# RestoreFilesTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies if the Restore Task should continue even if the copy operation of some files and folders fails. If true, the Cohesity Cluster ignores intermittent errors and recovers as many files and folders as possible. If false, the Restore Task stops recovering when a copy operation fails. | [optional] 
**FileRecoveryMethod** | Pointer to **NullableString** | Specifies the type of method to be used to perform file recovery. &#39;kAutoDeploy&#39; indicates that file restore operation wiil be performed using an ephemeral agent. &#39;kUseExistingAgent&#39; indicates that file restore operation wiil be performed using an persistent agent. &#39;kUseHypervisorAPIs&#39; indicates that file restore operation wiil be performed using an hypervisor API&#39;s. | [optional] 
**Filenames** | Pointer to **[]string** | Array of Files or Folders.  Specifies the files and folders to recover from the snapshot. | [optional] 
**IsFileBasedVolumeRestore** | Pointer to **NullableBool** | Specifies whether this is a file based volume restore. | [optional] 
**MountDisksOnVm** | Pointer to **NullableBool** | Sepcifies whether this will attach disks or mount disks on the VM side OR use Storage Proxy RPCs to stream data | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the Restore Task. This field must be set and must be a unique name. | [optional] 
**NewBaseDirectory** | Pointer to **NullableString** | Specifies an optional root folder where to recover the selected files and folders. By default, files and folders are restored to their original path. | [optional] 
**Overwrite** | Pointer to **NullableBool** | If true, any existing files and folders on the operating system are overwritten by the recovered files or folders. This is the default. If false, existing files and folders are not overwritten. | [optional] 
**Password** | Pointer to **NullableString** | Specifies password of the username to access the target source. | [optional] 
**PreserveAttributes** | Pointer to **NullableBool** | If true, the Restore Tasks preserves the original file and folder attributes. This is the default. | [optional] 
**RestoredFileInfoList** | Pointer to [**[]RestoredFileInfoList**](RestoredFileInfoList.md) | Specifies information regarding files and directories. | [optional] 
**SourceObjectInfo** | Pointer to [**NullableRestoreObjectDetails**](RestoreObjectDetails.md) | Specifies information about the source object (such as a VM) that contains the files and folders to recover. In addition, it contains information about the Protection Job and Job Run that captured the snapshot to recover from. To specify a particular snapshot, you must specify a jobRunId and a startTimeUsecs. If jobRunId and startTimeUsecs are not specified, the last Job Run of the specified Job is used. | [optional] 
**TargetHostType** | Pointer to **NullableString** | Specifies the target host types to be restored. &#39;kLinux&#39; indicates the Linux operating system. &#39;kWindows&#39; indicates the Microsoft Windows operating system. &#39;kAix&#39; indicates the IBM AIX operating system. &#39;kSolaris&#39; indicates the Oracle Solaris operating system. &#39;kSapHana&#39; indicates the Sap Hana database system developed by SAP SE. &#39;kOther&#39; indicates the other types of operating system. | [optional] 
**TargetParentSourceId** | Pointer to **NullableInt64** | Specifies the registered source (such as a vCenter Server) that contains the target protection source (such as a VM) where the files and folders are recovered to. This field is not required for a Physical Server. | [optional] 
**TargetSourceId** | Pointer to **NullableInt64** | Specifies the id of the target protection source (such as a VM) where the files and folders are recovered to. | [optional] 
**UseExistingAgent** | Pointer to **NullableBool** | Specifies whether this will use an existing agent on the target vm to do restore. Following field is deprecated and shall not be used. Please refer to the FileRecoveryMethod field for more information. | [optional] 
**Username** | Pointer to **NullableString** | Specifies username to access the target source. | [optional] 

## Methods

### NewRestoreFilesTaskRequest

`func NewRestoreFilesTaskRequest() *RestoreFilesTaskRequest`

NewRestoreFilesTaskRequest instantiates a new RestoreFilesTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreFilesTaskRequestWithDefaults

`func NewRestoreFilesTaskRequestWithDefaults() *RestoreFilesTaskRequest`

NewRestoreFilesTaskRequestWithDefaults instantiates a new RestoreFilesTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *RestoreFilesTaskRequest) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RestoreFilesTaskRequest) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RestoreFilesTaskRequest) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RestoreFilesTaskRequest) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RestoreFilesTaskRequest) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RestoreFilesTaskRequest) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetFileRecoveryMethod

`func (o *RestoreFilesTaskRequest) GetFileRecoveryMethod() string`

GetFileRecoveryMethod returns the FileRecoveryMethod field if non-nil, zero value otherwise.

### GetFileRecoveryMethodOk

`func (o *RestoreFilesTaskRequest) GetFileRecoveryMethodOk() (*string, bool)`

GetFileRecoveryMethodOk returns a tuple with the FileRecoveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileRecoveryMethod

`func (o *RestoreFilesTaskRequest) SetFileRecoveryMethod(v string)`

SetFileRecoveryMethod sets FileRecoveryMethod field to given value.

### HasFileRecoveryMethod

`func (o *RestoreFilesTaskRequest) HasFileRecoveryMethod() bool`

HasFileRecoveryMethod returns a boolean if a field has been set.

### SetFileRecoveryMethodNil

`func (o *RestoreFilesTaskRequest) SetFileRecoveryMethodNil(b bool)`

 SetFileRecoveryMethodNil sets the value for FileRecoveryMethod to be an explicit nil

### UnsetFileRecoveryMethod
`func (o *RestoreFilesTaskRequest) UnsetFileRecoveryMethod()`

UnsetFileRecoveryMethod ensures that no value is present for FileRecoveryMethod, not even an explicit nil
### GetFilenames

`func (o *RestoreFilesTaskRequest) GetFilenames() []string`

GetFilenames returns the Filenames field if non-nil, zero value otherwise.

### GetFilenamesOk

`func (o *RestoreFilesTaskRequest) GetFilenamesOk() (*[]string, bool)`

GetFilenamesOk returns a tuple with the Filenames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilenames

`func (o *RestoreFilesTaskRequest) SetFilenames(v []string)`

SetFilenames sets Filenames field to given value.

### HasFilenames

`func (o *RestoreFilesTaskRequest) HasFilenames() bool`

HasFilenames returns a boolean if a field has been set.

### SetFilenamesNil

`func (o *RestoreFilesTaskRequest) SetFilenamesNil(b bool)`

 SetFilenamesNil sets the value for Filenames to be an explicit nil

### UnsetFilenames
`func (o *RestoreFilesTaskRequest) UnsetFilenames()`

UnsetFilenames ensures that no value is present for Filenames, not even an explicit nil
### GetIsFileBasedVolumeRestore

`func (o *RestoreFilesTaskRequest) GetIsFileBasedVolumeRestore() bool`

GetIsFileBasedVolumeRestore returns the IsFileBasedVolumeRestore field if non-nil, zero value otherwise.

### GetIsFileBasedVolumeRestoreOk

`func (o *RestoreFilesTaskRequest) GetIsFileBasedVolumeRestoreOk() (*bool, bool)`

GetIsFileBasedVolumeRestoreOk returns a tuple with the IsFileBasedVolumeRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFileBasedVolumeRestore

`func (o *RestoreFilesTaskRequest) SetIsFileBasedVolumeRestore(v bool)`

SetIsFileBasedVolumeRestore sets IsFileBasedVolumeRestore field to given value.

### HasIsFileBasedVolumeRestore

`func (o *RestoreFilesTaskRequest) HasIsFileBasedVolumeRestore() bool`

HasIsFileBasedVolumeRestore returns a boolean if a field has been set.

### SetIsFileBasedVolumeRestoreNil

`func (o *RestoreFilesTaskRequest) SetIsFileBasedVolumeRestoreNil(b bool)`

 SetIsFileBasedVolumeRestoreNil sets the value for IsFileBasedVolumeRestore to be an explicit nil

### UnsetIsFileBasedVolumeRestore
`func (o *RestoreFilesTaskRequest) UnsetIsFileBasedVolumeRestore()`

UnsetIsFileBasedVolumeRestore ensures that no value is present for IsFileBasedVolumeRestore, not even an explicit nil
### GetMountDisksOnVm

`func (o *RestoreFilesTaskRequest) GetMountDisksOnVm() bool`

GetMountDisksOnVm returns the MountDisksOnVm field if non-nil, zero value otherwise.

### GetMountDisksOnVmOk

`func (o *RestoreFilesTaskRequest) GetMountDisksOnVmOk() (*bool, bool)`

GetMountDisksOnVmOk returns a tuple with the MountDisksOnVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountDisksOnVm

`func (o *RestoreFilesTaskRequest) SetMountDisksOnVm(v bool)`

SetMountDisksOnVm sets MountDisksOnVm field to given value.

### HasMountDisksOnVm

`func (o *RestoreFilesTaskRequest) HasMountDisksOnVm() bool`

HasMountDisksOnVm returns a boolean if a field has been set.

### SetMountDisksOnVmNil

`func (o *RestoreFilesTaskRequest) SetMountDisksOnVmNil(b bool)`

 SetMountDisksOnVmNil sets the value for MountDisksOnVm to be an explicit nil

### UnsetMountDisksOnVm
`func (o *RestoreFilesTaskRequest) UnsetMountDisksOnVm()`

UnsetMountDisksOnVm ensures that no value is present for MountDisksOnVm, not even an explicit nil
### GetName

`func (o *RestoreFilesTaskRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestoreFilesTaskRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestoreFilesTaskRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RestoreFilesTaskRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RestoreFilesTaskRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RestoreFilesTaskRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNewBaseDirectory

`func (o *RestoreFilesTaskRequest) GetNewBaseDirectory() string`

GetNewBaseDirectory returns the NewBaseDirectory field if non-nil, zero value otherwise.

### GetNewBaseDirectoryOk

`func (o *RestoreFilesTaskRequest) GetNewBaseDirectoryOk() (*string, bool)`

GetNewBaseDirectoryOk returns a tuple with the NewBaseDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewBaseDirectory

`func (o *RestoreFilesTaskRequest) SetNewBaseDirectory(v string)`

SetNewBaseDirectory sets NewBaseDirectory field to given value.

### HasNewBaseDirectory

`func (o *RestoreFilesTaskRequest) HasNewBaseDirectory() bool`

HasNewBaseDirectory returns a boolean if a field has been set.

### SetNewBaseDirectoryNil

`func (o *RestoreFilesTaskRequest) SetNewBaseDirectoryNil(b bool)`

 SetNewBaseDirectoryNil sets the value for NewBaseDirectory to be an explicit nil

### UnsetNewBaseDirectory
`func (o *RestoreFilesTaskRequest) UnsetNewBaseDirectory()`

UnsetNewBaseDirectory ensures that no value is present for NewBaseDirectory, not even an explicit nil
### GetOverwrite

`func (o *RestoreFilesTaskRequest) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *RestoreFilesTaskRequest) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *RestoreFilesTaskRequest) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *RestoreFilesTaskRequest) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### SetOverwriteNil

`func (o *RestoreFilesTaskRequest) SetOverwriteNil(b bool)`

 SetOverwriteNil sets the value for Overwrite to be an explicit nil

### UnsetOverwrite
`func (o *RestoreFilesTaskRequest) UnsetOverwrite()`

UnsetOverwrite ensures that no value is present for Overwrite, not even an explicit nil
### GetPassword

`func (o *RestoreFilesTaskRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RestoreFilesTaskRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RestoreFilesTaskRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RestoreFilesTaskRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *RestoreFilesTaskRequest) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *RestoreFilesTaskRequest) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPreserveAttributes

`func (o *RestoreFilesTaskRequest) GetPreserveAttributes() bool`

GetPreserveAttributes returns the PreserveAttributes field if non-nil, zero value otherwise.

### GetPreserveAttributesOk

`func (o *RestoreFilesTaskRequest) GetPreserveAttributesOk() (*bool, bool)`

GetPreserveAttributesOk returns a tuple with the PreserveAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveAttributes

`func (o *RestoreFilesTaskRequest) SetPreserveAttributes(v bool)`

SetPreserveAttributes sets PreserveAttributes field to given value.

### HasPreserveAttributes

`func (o *RestoreFilesTaskRequest) HasPreserveAttributes() bool`

HasPreserveAttributes returns a boolean if a field has been set.

### SetPreserveAttributesNil

`func (o *RestoreFilesTaskRequest) SetPreserveAttributesNil(b bool)`

 SetPreserveAttributesNil sets the value for PreserveAttributes to be an explicit nil

### UnsetPreserveAttributes
`func (o *RestoreFilesTaskRequest) UnsetPreserveAttributes()`

UnsetPreserveAttributes ensures that no value is present for PreserveAttributes, not even an explicit nil
### GetRestoredFileInfoList

`func (o *RestoreFilesTaskRequest) GetRestoredFileInfoList() []RestoredFileInfoList`

GetRestoredFileInfoList returns the RestoredFileInfoList field if non-nil, zero value otherwise.

### GetRestoredFileInfoListOk

`func (o *RestoreFilesTaskRequest) GetRestoredFileInfoListOk() (*[]RestoredFileInfoList, bool)`

GetRestoredFileInfoListOk returns a tuple with the RestoredFileInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoredFileInfoList

`func (o *RestoreFilesTaskRequest) SetRestoredFileInfoList(v []RestoredFileInfoList)`

SetRestoredFileInfoList sets RestoredFileInfoList field to given value.

### HasRestoredFileInfoList

`func (o *RestoreFilesTaskRequest) HasRestoredFileInfoList() bool`

HasRestoredFileInfoList returns a boolean if a field has been set.

### SetRestoredFileInfoListNil

`func (o *RestoreFilesTaskRequest) SetRestoredFileInfoListNil(b bool)`

 SetRestoredFileInfoListNil sets the value for RestoredFileInfoList to be an explicit nil

### UnsetRestoredFileInfoList
`func (o *RestoreFilesTaskRequest) UnsetRestoredFileInfoList()`

UnsetRestoredFileInfoList ensures that no value is present for RestoredFileInfoList, not even an explicit nil
### GetSourceObjectInfo

`func (o *RestoreFilesTaskRequest) GetSourceObjectInfo() RestoreObjectDetails`

GetSourceObjectInfo returns the SourceObjectInfo field if non-nil, zero value otherwise.

### GetSourceObjectInfoOk

`func (o *RestoreFilesTaskRequest) GetSourceObjectInfoOk() (*RestoreObjectDetails, bool)`

GetSourceObjectInfoOk returns a tuple with the SourceObjectInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceObjectInfo

`func (o *RestoreFilesTaskRequest) SetSourceObjectInfo(v RestoreObjectDetails)`

SetSourceObjectInfo sets SourceObjectInfo field to given value.

### HasSourceObjectInfo

`func (o *RestoreFilesTaskRequest) HasSourceObjectInfo() bool`

HasSourceObjectInfo returns a boolean if a field has been set.

### SetSourceObjectInfoNil

`func (o *RestoreFilesTaskRequest) SetSourceObjectInfoNil(b bool)`

 SetSourceObjectInfoNil sets the value for SourceObjectInfo to be an explicit nil

### UnsetSourceObjectInfo
`func (o *RestoreFilesTaskRequest) UnsetSourceObjectInfo()`

UnsetSourceObjectInfo ensures that no value is present for SourceObjectInfo, not even an explicit nil
### GetTargetHostType

`func (o *RestoreFilesTaskRequest) GetTargetHostType() string`

GetTargetHostType returns the TargetHostType field if non-nil, zero value otherwise.

### GetTargetHostTypeOk

`func (o *RestoreFilesTaskRequest) GetTargetHostTypeOk() (*string, bool)`

GetTargetHostTypeOk returns a tuple with the TargetHostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetHostType

`func (o *RestoreFilesTaskRequest) SetTargetHostType(v string)`

SetTargetHostType sets TargetHostType field to given value.

### HasTargetHostType

`func (o *RestoreFilesTaskRequest) HasTargetHostType() bool`

HasTargetHostType returns a boolean if a field has been set.

### SetTargetHostTypeNil

`func (o *RestoreFilesTaskRequest) SetTargetHostTypeNil(b bool)`

 SetTargetHostTypeNil sets the value for TargetHostType to be an explicit nil

### UnsetTargetHostType
`func (o *RestoreFilesTaskRequest) UnsetTargetHostType()`

UnsetTargetHostType ensures that no value is present for TargetHostType, not even an explicit nil
### GetTargetParentSourceId

`func (o *RestoreFilesTaskRequest) GetTargetParentSourceId() int64`

GetTargetParentSourceId returns the TargetParentSourceId field if non-nil, zero value otherwise.

### GetTargetParentSourceIdOk

`func (o *RestoreFilesTaskRequest) GetTargetParentSourceIdOk() (*int64, bool)`

GetTargetParentSourceIdOk returns a tuple with the TargetParentSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetParentSourceId

`func (o *RestoreFilesTaskRequest) SetTargetParentSourceId(v int64)`

SetTargetParentSourceId sets TargetParentSourceId field to given value.

### HasTargetParentSourceId

`func (o *RestoreFilesTaskRequest) HasTargetParentSourceId() bool`

HasTargetParentSourceId returns a boolean if a field has been set.

### SetTargetParentSourceIdNil

`func (o *RestoreFilesTaskRequest) SetTargetParentSourceIdNil(b bool)`

 SetTargetParentSourceIdNil sets the value for TargetParentSourceId to be an explicit nil

### UnsetTargetParentSourceId
`func (o *RestoreFilesTaskRequest) UnsetTargetParentSourceId()`

UnsetTargetParentSourceId ensures that no value is present for TargetParentSourceId, not even an explicit nil
### GetTargetSourceId

`func (o *RestoreFilesTaskRequest) GetTargetSourceId() int64`

GetTargetSourceId returns the TargetSourceId field if non-nil, zero value otherwise.

### GetTargetSourceIdOk

`func (o *RestoreFilesTaskRequest) GetTargetSourceIdOk() (*int64, bool)`

GetTargetSourceIdOk returns a tuple with the TargetSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSourceId

`func (o *RestoreFilesTaskRequest) SetTargetSourceId(v int64)`

SetTargetSourceId sets TargetSourceId field to given value.

### HasTargetSourceId

`func (o *RestoreFilesTaskRequest) HasTargetSourceId() bool`

HasTargetSourceId returns a boolean if a field has been set.

### SetTargetSourceIdNil

`func (o *RestoreFilesTaskRequest) SetTargetSourceIdNil(b bool)`

 SetTargetSourceIdNil sets the value for TargetSourceId to be an explicit nil

### UnsetTargetSourceId
`func (o *RestoreFilesTaskRequest) UnsetTargetSourceId()`

UnsetTargetSourceId ensures that no value is present for TargetSourceId, not even an explicit nil
### GetUseExistingAgent

`func (o *RestoreFilesTaskRequest) GetUseExistingAgent() bool`

GetUseExistingAgent returns the UseExistingAgent field if non-nil, zero value otherwise.

### GetUseExistingAgentOk

`func (o *RestoreFilesTaskRequest) GetUseExistingAgentOk() (*bool, bool)`

GetUseExistingAgentOk returns a tuple with the UseExistingAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseExistingAgent

`func (o *RestoreFilesTaskRequest) SetUseExistingAgent(v bool)`

SetUseExistingAgent sets UseExistingAgent field to given value.

### HasUseExistingAgent

`func (o *RestoreFilesTaskRequest) HasUseExistingAgent() bool`

HasUseExistingAgent returns a boolean if a field has been set.

### SetUseExistingAgentNil

`func (o *RestoreFilesTaskRequest) SetUseExistingAgentNil(b bool)`

 SetUseExistingAgentNil sets the value for UseExistingAgent to be an explicit nil

### UnsetUseExistingAgent
`func (o *RestoreFilesTaskRequest) UnsetUseExistingAgent()`

UnsetUseExistingAgent ensures that no value is present for UseExistingAgent, not even an explicit nil
### GetUsername

`func (o *RestoreFilesTaskRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RestoreFilesTaskRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RestoreFilesTaskRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RestoreFilesTaskRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *RestoreFilesTaskRequest) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *RestoreFilesTaskRequest) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


