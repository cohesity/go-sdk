# RestoreFilesParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlacklistedIpAddrs** | Pointer to **[]string** | A list of target IP addresses that should not be used. | [optional] 
**DestinationEpUuid** | Pointer to **NullableString** | Destination endpoint UUID for source s3 objectstore. | [optional] 
**DirectoryNameSecurityStyleMap** | Pointer to [**[]RestoreFilesParamsDirectoryNameSecurityStyleMapEntry**](RestoreFilesParamsDirectoryNameSecurityStyleMapEntry.md) | Directory name security style map contains mapping of the directory name to security style it supports.  This is needed to restore the same permission for the given directory for Qtrees. | [optional] 
**IsArchiveFlr** | Pointer to **NullableBool** | Whether this is a file restore operation from an archive. | [optional] 
**IsFileVolumeRestore** | Pointer to **NullableBool** | Whether this is a file based volume restore. | [optional] 
**IsMountBasedFlr** | Pointer to **NullableBool** | Whether this is a mount based file restore operation | [optional] 
**IsSourceInitiatedRestore** | Pointer to **NullableBool** | Whether this is a source initiated restore. | [optional] 
**MountDisksOnVm** | Pointer to **NullableBool** | Whether this will attach disks or mount disks on the VM side OR use Storage Proxy RPCs to stream data. | [optional] 
**NasBackupParams** | Pointer to [**NasBackupParams**](NasBackupParams.md) |  | [optional] 
**NasProtocolTypeVec** | Pointer to **[]int32** | The NAS protocol type(s) of this restore task. Currently we only support a single restore type. This field is only populated for NAS restore tasks. | [optional] 
**ObjectstoreConfigName** | Pointer to **NullableString** | Object store config name for source initiated backup. | [optional] 
**PhysicalFlrParallelRestore** | Pointer to **NullableBool** | If enabled, magneto physical file restore will be enabled via job framework | [optional] 
**ProxyEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**ProxyEntityParentSource** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**RestoreFilesPreferences** | Pointer to [**RestoreFilesPreferences**](RestoreFilesPreferences.md) |  | [optional] 
**RestoreMethod** | Pointer to **NullableInt32** | Determines the type of method to be used to perform FLR. | [optional] 
**RestoredFileInfoVec** | Pointer to [**[]RestoredFileInfo**](RestoredFileInfo.md) | Information regarding files and directories. | [optional] 
**S3Viewbackupproperties** | Pointer to [**S3ViewBackupProperties**](S3ViewBackupProperties.md) |  | [optional] 
**SourceSnapshotName** | Pointer to **NullableString** | Snapshot name need by source to start the restore. | [optional] 
**TargetEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**TargetEntityCredentials** | Pointer to [**Credentials**](Credentials.md) |  | [optional] 
**TargetEntityParentSource** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**TargetHostEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**TargetHostType** | Pointer to **NullableInt32** | The host environment type. This is set in VMware environment to indicate the OS type of the target entity. NOTE: This is expected to be set since magneto does not know the host type for VMware entities. | [optional] 
**UptierParams** | Pointer to [**FileUptieringParams**](FileUptieringParams.md) |  | [optional] 
**UseExistingAgent** | Pointer to **NullableBool** | Whether this will use an existing agent on the target VM to do the restore. This field is deprecated. restore_method should be used for populating use of existing agent. | [optional] 
**ViewId** | Pointer to **NullableInt64** | View ID | [optional] 
**ViewName** | Pointer to **NullableString** | Name of the S3 view | [optional] 
**VpcConnectorEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 

## Methods

### NewRestoreFilesParams

`func NewRestoreFilesParams() *RestoreFilesParams`

NewRestoreFilesParams instantiates a new RestoreFilesParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreFilesParamsWithDefaults

`func NewRestoreFilesParamsWithDefaults() *RestoreFilesParams`

NewRestoreFilesParamsWithDefaults instantiates a new RestoreFilesParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlacklistedIpAddrs

`func (o *RestoreFilesParams) GetBlacklistedIpAddrs() []string`

GetBlacklistedIpAddrs returns the BlacklistedIpAddrs field if non-nil, zero value otherwise.

### GetBlacklistedIpAddrsOk

`func (o *RestoreFilesParams) GetBlacklistedIpAddrsOk() (*[]string, bool)`

GetBlacklistedIpAddrsOk returns a tuple with the BlacklistedIpAddrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistedIpAddrs

`func (o *RestoreFilesParams) SetBlacklistedIpAddrs(v []string)`

SetBlacklistedIpAddrs sets BlacklistedIpAddrs field to given value.

### HasBlacklistedIpAddrs

`func (o *RestoreFilesParams) HasBlacklistedIpAddrs() bool`

HasBlacklistedIpAddrs returns a boolean if a field has been set.

### SetBlacklistedIpAddrsNil

`func (o *RestoreFilesParams) SetBlacklistedIpAddrsNil(b bool)`

 SetBlacklistedIpAddrsNil sets the value for BlacklistedIpAddrs to be an explicit nil

### UnsetBlacklistedIpAddrs
`func (o *RestoreFilesParams) UnsetBlacklistedIpAddrs()`

UnsetBlacklistedIpAddrs ensures that no value is present for BlacklistedIpAddrs, not even an explicit nil
### GetDestinationEpUuid

`func (o *RestoreFilesParams) GetDestinationEpUuid() string`

GetDestinationEpUuid returns the DestinationEpUuid field if non-nil, zero value otherwise.

### GetDestinationEpUuidOk

`func (o *RestoreFilesParams) GetDestinationEpUuidOk() (*string, bool)`

GetDestinationEpUuidOk returns a tuple with the DestinationEpUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationEpUuid

`func (o *RestoreFilesParams) SetDestinationEpUuid(v string)`

SetDestinationEpUuid sets DestinationEpUuid field to given value.

### HasDestinationEpUuid

`func (o *RestoreFilesParams) HasDestinationEpUuid() bool`

HasDestinationEpUuid returns a boolean if a field has been set.

### SetDestinationEpUuidNil

`func (o *RestoreFilesParams) SetDestinationEpUuidNil(b bool)`

 SetDestinationEpUuidNil sets the value for DestinationEpUuid to be an explicit nil

### UnsetDestinationEpUuid
`func (o *RestoreFilesParams) UnsetDestinationEpUuid()`

UnsetDestinationEpUuid ensures that no value is present for DestinationEpUuid, not even an explicit nil
### GetDirectoryNameSecurityStyleMap

`func (o *RestoreFilesParams) GetDirectoryNameSecurityStyleMap() []RestoreFilesParamsDirectoryNameSecurityStyleMapEntry`

GetDirectoryNameSecurityStyleMap returns the DirectoryNameSecurityStyleMap field if non-nil, zero value otherwise.

### GetDirectoryNameSecurityStyleMapOk

`func (o *RestoreFilesParams) GetDirectoryNameSecurityStyleMapOk() (*[]RestoreFilesParamsDirectoryNameSecurityStyleMapEntry, bool)`

GetDirectoryNameSecurityStyleMapOk returns a tuple with the DirectoryNameSecurityStyleMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryNameSecurityStyleMap

`func (o *RestoreFilesParams) SetDirectoryNameSecurityStyleMap(v []RestoreFilesParamsDirectoryNameSecurityStyleMapEntry)`

SetDirectoryNameSecurityStyleMap sets DirectoryNameSecurityStyleMap field to given value.

### HasDirectoryNameSecurityStyleMap

`func (o *RestoreFilesParams) HasDirectoryNameSecurityStyleMap() bool`

HasDirectoryNameSecurityStyleMap returns a boolean if a field has been set.

### SetDirectoryNameSecurityStyleMapNil

`func (o *RestoreFilesParams) SetDirectoryNameSecurityStyleMapNil(b bool)`

 SetDirectoryNameSecurityStyleMapNil sets the value for DirectoryNameSecurityStyleMap to be an explicit nil

### UnsetDirectoryNameSecurityStyleMap
`func (o *RestoreFilesParams) UnsetDirectoryNameSecurityStyleMap()`

UnsetDirectoryNameSecurityStyleMap ensures that no value is present for DirectoryNameSecurityStyleMap, not even an explicit nil
### GetIsArchiveFlr

`func (o *RestoreFilesParams) GetIsArchiveFlr() bool`

GetIsArchiveFlr returns the IsArchiveFlr field if non-nil, zero value otherwise.

### GetIsArchiveFlrOk

`func (o *RestoreFilesParams) GetIsArchiveFlrOk() (*bool, bool)`

GetIsArchiveFlrOk returns a tuple with the IsArchiveFlr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchiveFlr

`func (o *RestoreFilesParams) SetIsArchiveFlr(v bool)`

SetIsArchiveFlr sets IsArchiveFlr field to given value.

### HasIsArchiveFlr

`func (o *RestoreFilesParams) HasIsArchiveFlr() bool`

HasIsArchiveFlr returns a boolean if a field has been set.

### SetIsArchiveFlrNil

`func (o *RestoreFilesParams) SetIsArchiveFlrNil(b bool)`

 SetIsArchiveFlrNil sets the value for IsArchiveFlr to be an explicit nil

### UnsetIsArchiveFlr
`func (o *RestoreFilesParams) UnsetIsArchiveFlr()`

UnsetIsArchiveFlr ensures that no value is present for IsArchiveFlr, not even an explicit nil
### GetIsFileVolumeRestore

`func (o *RestoreFilesParams) GetIsFileVolumeRestore() bool`

GetIsFileVolumeRestore returns the IsFileVolumeRestore field if non-nil, zero value otherwise.

### GetIsFileVolumeRestoreOk

`func (o *RestoreFilesParams) GetIsFileVolumeRestoreOk() (*bool, bool)`

GetIsFileVolumeRestoreOk returns a tuple with the IsFileVolumeRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFileVolumeRestore

`func (o *RestoreFilesParams) SetIsFileVolumeRestore(v bool)`

SetIsFileVolumeRestore sets IsFileVolumeRestore field to given value.

### HasIsFileVolumeRestore

`func (o *RestoreFilesParams) HasIsFileVolumeRestore() bool`

HasIsFileVolumeRestore returns a boolean if a field has been set.

### SetIsFileVolumeRestoreNil

`func (o *RestoreFilesParams) SetIsFileVolumeRestoreNil(b bool)`

 SetIsFileVolumeRestoreNil sets the value for IsFileVolumeRestore to be an explicit nil

### UnsetIsFileVolumeRestore
`func (o *RestoreFilesParams) UnsetIsFileVolumeRestore()`

UnsetIsFileVolumeRestore ensures that no value is present for IsFileVolumeRestore, not even an explicit nil
### GetIsMountBasedFlr

`func (o *RestoreFilesParams) GetIsMountBasedFlr() bool`

GetIsMountBasedFlr returns the IsMountBasedFlr field if non-nil, zero value otherwise.

### GetIsMountBasedFlrOk

`func (o *RestoreFilesParams) GetIsMountBasedFlrOk() (*bool, bool)`

GetIsMountBasedFlrOk returns a tuple with the IsMountBasedFlr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMountBasedFlr

`func (o *RestoreFilesParams) SetIsMountBasedFlr(v bool)`

SetIsMountBasedFlr sets IsMountBasedFlr field to given value.

### HasIsMountBasedFlr

`func (o *RestoreFilesParams) HasIsMountBasedFlr() bool`

HasIsMountBasedFlr returns a boolean if a field has been set.

### SetIsMountBasedFlrNil

`func (o *RestoreFilesParams) SetIsMountBasedFlrNil(b bool)`

 SetIsMountBasedFlrNil sets the value for IsMountBasedFlr to be an explicit nil

### UnsetIsMountBasedFlr
`func (o *RestoreFilesParams) UnsetIsMountBasedFlr()`

UnsetIsMountBasedFlr ensures that no value is present for IsMountBasedFlr, not even an explicit nil
### GetIsSourceInitiatedRestore

`func (o *RestoreFilesParams) GetIsSourceInitiatedRestore() bool`

GetIsSourceInitiatedRestore returns the IsSourceInitiatedRestore field if non-nil, zero value otherwise.

### GetIsSourceInitiatedRestoreOk

`func (o *RestoreFilesParams) GetIsSourceInitiatedRestoreOk() (*bool, bool)`

GetIsSourceInitiatedRestoreOk returns a tuple with the IsSourceInitiatedRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSourceInitiatedRestore

`func (o *RestoreFilesParams) SetIsSourceInitiatedRestore(v bool)`

SetIsSourceInitiatedRestore sets IsSourceInitiatedRestore field to given value.

### HasIsSourceInitiatedRestore

`func (o *RestoreFilesParams) HasIsSourceInitiatedRestore() bool`

HasIsSourceInitiatedRestore returns a boolean if a field has been set.

### SetIsSourceInitiatedRestoreNil

`func (o *RestoreFilesParams) SetIsSourceInitiatedRestoreNil(b bool)`

 SetIsSourceInitiatedRestoreNil sets the value for IsSourceInitiatedRestore to be an explicit nil

### UnsetIsSourceInitiatedRestore
`func (o *RestoreFilesParams) UnsetIsSourceInitiatedRestore()`

UnsetIsSourceInitiatedRestore ensures that no value is present for IsSourceInitiatedRestore, not even an explicit nil
### GetMountDisksOnVm

`func (o *RestoreFilesParams) GetMountDisksOnVm() bool`

GetMountDisksOnVm returns the MountDisksOnVm field if non-nil, zero value otherwise.

### GetMountDisksOnVmOk

`func (o *RestoreFilesParams) GetMountDisksOnVmOk() (*bool, bool)`

GetMountDisksOnVmOk returns a tuple with the MountDisksOnVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountDisksOnVm

`func (o *RestoreFilesParams) SetMountDisksOnVm(v bool)`

SetMountDisksOnVm sets MountDisksOnVm field to given value.

### HasMountDisksOnVm

`func (o *RestoreFilesParams) HasMountDisksOnVm() bool`

HasMountDisksOnVm returns a boolean if a field has been set.

### SetMountDisksOnVmNil

`func (o *RestoreFilesParams) SetMountDisksOnVmNil(b bool)`

 SetMountDisksOnVmNil sets the value for MountDisksOnVm to be an explicit nil

### UnsetMountDisksOnVm
`func (o *RestoreFilesParams) UnsetMountDisksOnVm()`

UnsetMountDisksOnVm ensures that no value is present for MountDisksOnVm, not even an explicit nil
### GetNasBackupParams

`func (o *RestoreFilesParams) GetNasBackupParams() NasBackupParams`

GetNasBackupParams returns the NasBackupParams field if non-nil, zero value otherwise.

### GetNasBackupParamsOk

`func (o *RestoreFilesParams) GetNasBackupParamsOk() (*NasBackupParams, bool)`

GetNasBackupParamsOk returns a tuple with the NasBackupParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasBackupParams

`func (o *RestoreFilesParams) SetNasBackupParams(v NasBackupParams)`

SetNasBackupParams sets NasBackupParams field to given value.

### HasNasBackupParams

`func (o *RestoreFilesParams) HasNasBackupParams() bool`

HasNasBackupParams returns a boolean if a field has been set.

### GetNasProtocolTypeVec

`func (o *RestoreFilesParams) GetNasProtocolTypeVec() []int32`

GetNasProtocolTypeVec returns the NasProtocolTypeVec field if non-nil, zero value otherwise.

### GetNasProtocolTypeVecOk

`func (o *RestoreFilesParams) GetNasProtocolTypeVecOk() (*[]int32, bool)`

GetNasProtocolTypeVecOk returns a tuple with the NasProtocolTypeVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasProtocolTypeVec

`func (o *RestoreFilesParams) SetNasProtocolTypeVec(v []int32)`

SetNasProtocolTypeVec sets NasProtocolTypeVec field to given value.

### HasNasProtocolTypeVec

`func (o *RestoreFilesParams) HasNasProtocolTypeVec() bool`

HasNasProtocolTypeVec returns a boolean if a field has been set.

### SetNasProtocolTypeVecNil

`func (o *RestoreFilesParams) SetNasProtocolTypeVecNil(b bool)`

 SetNasProtocolTypeVecNil sets the value for NasProtocolTypeVec to be an explicit nil

### UnsetNasProtocolTypeVec
`func (o *RestoreFilesParams) UnsetNasProtocolTypeVec()`

UnsetNasProtocolTypeVec ensures that no value is present for NasProtocolTypeVec, not even an explicit nil
### GetObjectstoreConfigName

`func (o *RestoreFilesParams) GetObjectstoreConfigName() string`

GetObjectstoreConfigName returns the ObjectstoreConfigName field if non-nil, zero value otherwise.

### GetObjectstoreConfigNameOk

`func (o *RestoreFilesParams) GetObjectstoreConfigNameOk() (*string, bool)`

GetObjectstoreConfigNameOk returns a tuple with the ObjectstoreConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectstoreConfigName

`func (o *RestoreFilesParams) SetObjectstoreConfigName(v string)`

SetObjectstoreConfigName sets ObjectstoreConfigName field to given value.

### HasObjectstoreConfigName

`func (o *RestoreFilesParams) HasObjectstoreConfigName() bool`

HasObjectstoreConfigName returns a boolean if a field has been set.

### SetObjectstoreConfigNameNil

`func (o *RestoreFilesParams) SetObjectstoreConfigNameNil(b bool)`

 SetObjectstoreConfigNameNil sets the value for ObjectstoreConfigName to be an explicit nil

### UnsetObjectstoreConfigName
`func (o *RestoreFilesParams) UnsetObjectstoreConfigName()`

UnsetObjectstoreConfigName ensures that no value is present for ObjectstoreConfigName, not even an explicit nil
### GetPhysicalFlrParallelRestore

`func (o *RestoreFilesParams) GetPhysicalFlrParallelRestore() bool`

GetPhysicalFlrParallelRestore returns the PhysicalFlrParallelRestore field if non-nil, zero value otherwise.

### GetPhysicalFlrParallelRestoreOk

`func (o *RestoreFilesParams) GetPhysicalFlrParallelRestoreOk() (*bool, bool)`

GetPhysicalFlrParallelRestoreOk returns a tuple with the PhysicalFlrParallelRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalFlrParallelRestore

`func (o *RestoreFilesParams) SetPhysicalFlrParallelRestore(v bool)`

SetPhysicalFlrParallelRestore sets PhysicalFlrParallelRestore field to given value.

### HasPhysicalFlrParallelRestore

`func (o *RestoreFilesParams) HasPhysicalFlrParallelRestore() bool`

HasPhysicalFlrParallelRestore returns a boolean if a field has been set.

### SetPhysicalFlrParallelRestoreNil

`func (o *RestoreFilesParams) SetPhysicalFlrParallelRestoreNil(b bool)`

 SetPhysicalFlrParallelRestoreNil sets the value for PhysicalFlrParallelRestore to be an explicit nil

### UnsetPhysicalFlrParallelRestore
`func (o *RestoreFilesParams) UnsetPhysicalFlrParallelRestore()`

UnsetPhysicalFlrParallelRestore ensures that no value is present for PhysicalFlrParallelRestore, not even an explicit nil
### GetProxyEntity

`func (o *RestoreFilesParams) GetProxyEntity() EntityProto`

GetProxyEntity returns the ProxyEntity field if non-nil, zero value otherwise.

### GetProxyEntityOk

`func (o *RestoreFilesParams) GetProxyEntityOk() (*EntityProto, bool)`

GetProxyEntityOk returns a tuple with the ProxyEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyEntity

`func (o *RestoreFilesParams) SetProxyEntity(v EntityProto)`

SetProxyEntity sets ProxyEntity field to given value.

### HasProxyEntity

`func (o *RestoreFilesParams) HasProxyEntity() bool`

HasProxyEntity returns a boolean if a field has been set.

### GetProxyEntityParentSource

`func (o *RestoreFilesParams) GetProxyEntityParentSource() EntityProto`

GetProxyEntityParentSource returns the ProxyEntityParentSource field if non-nil, zero value otherwise.

### GetProxyEntityParentSourceOk

`func (o *RestoreFilesParams) GetProxyEntityParentSourceOk() (*EntityProto, bool)`

GetProxyEntityParentSourceOk returns a tuple with the ProxyEntityParentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyEntityParentSource

`func (o *RestoreFilesParams) SetProxyEntityParentSource(v EntityProto)`

SetProxyEntityParentSource sets ProxyEntityParentSource field to given value.

### HasProxyEntityParentSource

`func (o *RestoreFilesParams) HasProxyEntityParentSource() bool`

HasProxyEntityParentSource returns a boolean if a field has been set.

### GetRestoreFilesPreferences

`func (o *RestoreFilesParams) GetRestoreFilesPreferences() RestoreFilesPreferences`

GetRestoreFilesPreferences returns the RestoreFilesPreferences field if non-nil, zero value otherwise.

### GetRestoreFilesPreferencesOk

`func (o *RestoreFilesParams) GetRestoreFilesPreferencesOk() (*RestoreFilesPreferences, bool)`

GetRestoreFilesPreferencesOk returns a tuple with the RestoreFilesPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreFilesPreferences

`func (o *RestoreFilesParams) SetRestoreFilesPreferences(v RestoreFilesPreferences)`

SetRestoreFilesPreferences sets RestoreFilesPreferences field to given value.

### HasRestoreFilesPreferences

`func (o *RestoreFilesParams) HasRestoreFilesPreferences() bool`

HasRestoreFilesPreferences returns a boolean if a field has been set.

### GetRestoreMethod

`func (o *RestoreFilesParams) GetRestoreMethod() int32`

GetRestoreMethod returns the RestoreMethod field if non-nil, zero value otherwise.

### GetRestoreMethodOk

`func (o *RestoreFilesParams) GetRestoreMethodOk() (*int32, bool)`

GetRestoreMethodOk returns a tuple with the RestoreMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreMethod

`func (o *RestoreFilesParams) SetRestoreMethod(v int32)`

SetRestoreMethod sets RestoreMethod field to given value.

### HasRestoreMethod

`func (o *RestoreFilesParams) HasRestoreMethod() bool`

HasRestoreMethod returns a boolean if a field has been set.

### SetRestoreMethodNil

`func (o *RestoreFilesParams) SetRestoreMethodNil(b bool)`

 SetRestoreMethodNil sets the value for RestoreMethod to be an explicit nil

### UnsetRestoreMethod
`func (o *RestoreFilesParams) UnsetRestoreMethod()`

UnsetRestoreMethod ensures that no value is present for RestoreMethod, not even an explicit nil
### GetRestoredFileInfoVec

`func (o *RestoreFilesParams) GetRestoredFileInfoVec() []RestoredFileInfo`

GetRestoredFileInfoVec returns the RestoredFileInfoVec field if non-nil, zero value otherwise.

### GetRestoredFileInfoVecOk

`func (o *RestoreFilesParams) GetRestoredFileInfoVecOk() (*[]RestoredFileInfo, bool)`

GetRestoredFileInfoVecOk returns a tuple with the RestoredFileInfoVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoredFileInfoVec

`func (o *RestoreFilesParams) SetRestoredFileInfoVec(v []RestoredFileInfo)`

SetRestoredFileInfoVec sets RestoredFileInfoVec field to given value.

### HasRestoredFileInfoVec

`func (o *RestoreFilesParams) HasRestoredFileInfoVec() bool`

HasRestoredFileInfoVec returns a boolean if a field has been set.

### SetRestoredFileInfoVecNil

`func (o *RestoreFilesParams) SetRestoredFileInfoVecNil(b bool)`

 SetRestoredFileInfoVecNil sets the value for RestoredFileInfoVec to be an explicit nil

### UnsetRestoredFileInfoVec
`func (o *RestoreFilesParams) UnsetRestoredFileInfoVec()`

UnsetRestoredFileInfoVec ensures that no value is present for RestoredFileInfoVec, not even an explicit nil
### GetS3Viewbackupproperties

`func (o *RestoreFilesParams) GetS3Viewbackupproperties() S3ViewBackupProperties`

GetS3Viewbackupproperties returns the S3Viewbackupproperties field if non-nil, zero value otherwise.

### GetS3ViewbackuppropertiesOk

`func (o *RestoreFilesParams) GetS3ViewbackuppropertiesOk() (*S3ViewBackupProperties, bool)`

GetS3ViewbackuppropertiesOk returns a tuple with the S3Viewbackupproperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Viewbackupproperties

`func (o *RestoreFilesParams) SetS3Viewbackupproperties(v S3ViewBackupProperties)`

SetS3Viewbackupproperties sets S3Viewbackupproperties field to given value.

### HasS3Viewbackupproperties

`func (o *RestoreFilesParams) HasS3Viewbackupproperties() bool`

HasS3Viewbackupproperties returns a boolean if a field has been set.

### GetSourceSnapshotName

`func (o *RestoreFilesParams) GetSourceSnapshotName() string`

GetSourceSnapshotName returns the SourceSnapshotName field if non-nil, zero value otherwise.

### GetSourceSnapshotNameOk

`func (o *RestoreFilesParams) GetSourceSnapshotNameOk() (*string, bool)`

GetSourceSnapshotNameOk returns a tuple with the SourceSnapshotName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSnapshotName

`func (o *RestoreFilesParams) SetSourceSnapshotName(v string)`

SetSourceSnapshotName sets SourceSnapshotName field to given value.

### HasSourceSnapshotName

`func (o *RestoreFilesParams) HasSourceSnapshotName() bool`

HasSourceSnapshotName returns a boolean if a field has been set.

### SetSourceSnapshotNameNil

`func (o *RestoreFilesParams) SetSourceSnapshotNameNil(b bool)`

 SetSourceSnapshotNameNil sets the value for SourceSnapshotName to be an explicit nil

### UnsetSourceSnapshotName
`func (o *RestoreFilesParams) UnsetSourceSnapshotName()`

UnsetSourceSnapshotName ensures that no value is present for SourceSnapshotName, not even an explicit nil
### GetTargetEntity

`func (o *RestoreFilesParams) GetTargetEntity() EntityProto`

GetTargetEntity returns the TargetEntity field if non-nil, zero value otherwise.

### GetTargetEntityOk

`func (o *RestoreFilesParams) GetTargetEntityOk() (*EntityProto, bool)`

GetTargetEntityOk returns a tuple with the TargetEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEntity

`func (o *RestoreFilesParams) SetTargetEntity(v EntityProto)`

SetTargetEntity sets TargetEntity field to given value.

### HasTargetEntity

`func (o *RestoreFilesParams) HasTargetEntity() bool`

HasTargetEntity returns a boolean if a field has been set.

### GetTargetEntityCredentials

`func (o *RestoreFilesParams) GetTargetEntityCredentials() Credentials`

GetTargetEntityCredentials returns the TargetEntityCredentials field if non-nil, zero value otherwise.

### GetTargetEntityCredentialsOk

`func (o *RestoreFilesParams) GetTargetEntityCredentialsOk() (*Credentials, bool)`

GetTargetEntityCredentialsOk returns a tuple with the TargetEntityCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEntityCredentials

`func (o *RestoreFilesParams) SetTargetEntityCredentials(v Credentials)`

SetTargetEntityCredentials sets TargetEntityCredentials field to given value.

### HasTargetEntityCredentials

`func (o *RestoreFilesParams) HasTargetEntityCredentials() bool`

HasTargetEntityCredentials returns a boolean if a field has been set.

### GetTargetEntityParentSource

`func (o *RestoreFilesParams) GetTargetEntityParentSource() EntityProto`

GetTargetEntityParentSource returns the TargetEntityParentSource field if non-nil, zero value otherwise.

### GetTargetEntityParentSourceOk

`func (o *RestoreFilesParams) GetTargetEntityParentSourceOk() (*EntityProto, bool)`

GetTargetEntityParentSourceOk returns a tuple with the TargetEntityParentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEntityParentSource

`func (o *RestoreFilesParams) SetTargetEntityParentSource(v EntityProto)`

SetTargetEntityParentSource sets TargetEntityParentSource field to given value.

### HasTargetEntityParentSource

`func (o *RestoreFilesParams) HasTargetEntityParentSource() bool`

HasTargetEntityParentSource returns a boolean if a field has been set.

### GetTargetHostEntity

`func (o *RestoreFilesParams) GetTargetHostEntity() EntityProto`

GetTargetHostEntity returns the TargetHostEntity field if non-nil, zero value otherwise.

### GetTargetHostEntityOk

`func (o *RestoreFilesParams) GetTargetHostEntityOk() (*EntityProto, bool)`

GetTargetHostEntityOk returns a tuple with the TargetHostEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetHostEntity

`func (o *RestoreFilesParams) SetTargetHostEntity(v EntityProto)`

SetTargetHostEntity sets TargetHostEntity field to given value.

### HasTargetHostEntity

`func (o *RestoreFilesParams) HasTargetHostEntity() bool`

HasTargetHostEntity returns a boolean if a field has been set.

### GetTargetHostType

`func (o *RestoreFilesParams) GetTargetHostType() int32`

GetTargetHostType returns the TargetHostType field if non-nil, zero value otherwise.

### GetTargetHostTypeOk

`func (o *RestoreFilesParams) GetTargetHostTypeOk() (*int32, bool)`

GetTargetHostTypeOk returns a tuple with the TargetHostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetHostType

`func (o *RestoreFilesParams) SetTargetHostType(v int32)`

SetTargetHostType sets TargetHostType field to given value.

### HasTargetHostType

`func (o *RestoreFilesParams) HasTargetHostType() bool`

HasTargetHostType returns a boolean if a field has been set.

### SetTargetHostTypeNil

`func (o *RestoreFilesParams) SetTargetHostTypeNil(b bool)`

 SetTargetHostTypeNil sets the value for TargetHostType to be an explicit nil

### UnsetTargetHostType
`func (o *RestoreFilesParams) UnsetTargetHostType()`

UnsetTargetHostType ensures that no value is present for TargetHostType, not even an explicit nil
### GetUptierParams

`func (o *RestoreFilesParams) GetUptierParams() FileUptieringParams`

GetUptierParams returns the UptierParams field if non-nil, zero value otherwise.

### GetUptierParamsOk

`func (o *RestoreFilesParams) GetUptierParamsOk() (*FileUptieringParams, bool)`

GetUptierParamsOk returns a tuple with the UptierParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptierParams

`func (o *RestoreFilesParams) SetUptierParams(v FileUptieringParams)`

SetUptierParams sets UptierParams field to given value.

### HasUptierParams

`func (o *RestoreFilesParams) HasUptierParams() bool`

HasUptierParams returns a boolean if a field has been set.

### GetUseExistingAgent

`func (o *RestoreFilesParams) GetUseExistingAgent() bool`

GetUseExistingAgent returns the UseExistingAgent field if non-nil, zero value otherwise.

### GetUseExistingAgentOk

`func (o *RestoreFilesParams) GetUseExistingAgentOk() (*bool, bool)`

GetUseExistingAgentOk returns a tuple with the UseExistingAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseExistingAgent

`func (o *RestoreFilesParams) SetUseExistingAgent(v bool)`

SetUseExistingAgent sets UseExistingAgent field to given value.

### HasUseExistingAgent

`func (o *RestoreFilesParams) HasUseExistingAgent() bool`

HasUseExistingAgent returns a boolean if a field has been set.

### SetUseExistingAgentNil

`func (o *RestoreFilesParams) SetUseExistingAgentNil(b bool)`

 SetUseExistingAgentNil sets the value for UseExistingAgent to be an explicit nil

### UnsetUseExistingAgent
`func (o *RestoreFilesParams) UnsetUseExistingAgent()`

UnsetUseExistingAgent ensures that no value is present for UseExistingAgent, not even an explicit nil
### GetViewId

`func (o *RestoreFilesParams) GetViewId() int64`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *RestoreFilesParams) GetViewIdOk() (*int64, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *RestoreFilesParams) SetViewId(v int64)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *RestoreFilesParams) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### SetViewIdNil

`func (o *RestoreFilesParams) SetViewIdNil(b bool)`

 SetViewIdNil sets the value for ViewId to be an explicit nil

### UnsetViewId
`func (o *RestoreFilesParams) UnsetViewId()`

UnsetViewId ensures that no value is present for ViewId, not even an explicit nil
### GetViewName

`func (o *RestoreFilesParams) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *RestoreFilesParams) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *RestoreFilesParams) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *RestoreFilesParams) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *RestoreFilesParams) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *RestoreFilesParams) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil
### GetVpcConnectorEntity

`func (o *RestoreFilesParams) GetVpcConnectorEntity() EntityProto`

GetVpcConnectorEntity returns the VpcConnectorEntity field if non-nil, zero value otherwise.

### GetVpcConnectorEntityOk

`func (o *RestoreFilesParams) GetVpcConnectorEntityOk() (*EntityProto, bool)`

GetVpcConnectorEntityOk returns a tuple with the VpcConnectorEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcConnectorEntity

`func (o *RestoreFilesParams) SetVpcConnectorEntity(v EntityProto)`

SetVpcConnectorEntity sets VpcConnectorEntity field to given value.

### HasVpcConnectorEntity

`func (o *RestoreFilesParams) HasVpcConnectorEntity() bool`

HasVpcConnectorEntity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


