# NetappProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupExistingSnapshot** | Pointer to **NullableBool** | Specifies that snapshot label is not set for Data-Protect Netapp Volumes backup. If field is set to true, existing oldest snapshot is used for backup and subsequent incremental will be selected in ascending order of snapshot create time on the source. If snapshot label is set, this field is set to false. | [optional] 
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether or not the Protection Group should continue regardless of whether or not an error was encountered during protection group run. | [optional] 
**ContinuousSnapshots** | Pointer to [**ContinuousSnapshotParams**](ContinuousSnapshotParams.md) |  | [optional] 
**DirectCloudArchive** | Pointer to **NullableBool** | Specifies whether or not to store the snapshots in this run directly in an Archive Target instead of on the Cluster. If this is set to true, the associated policy must have exactly one Archive Target associated with it and the policy must be set up to archive after every run. Also, a Storage Domain cannot be specified. Default behavior is &#39;false&#39;. | [optional] 
**EncryptionEnabled** | Pointer to **NullableBool** | Specifies whether the protection group should use encryption while backup or not. | [optional] 
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the objects to be excluded in the Protection Group. | [optional] 
**FileFilters** | Pointer to [**FileFilteringPolicy**](FileFilteringPolicy.md) |  | [optional] 
**FileLockConfig** | Pointer to [**FileLevelDataLockConfig**](FileLevelDataLockConfig.md) |  | [optional] 
**FilterIpConfig** | Pointer to [**FilterIpConfig**](FilterIpConfig.md) |  | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**ModifySourcePermissions** | Pointer to **NullableBool** | Specifies if the Netapp source permissions should be modified internally to allow backups. | [optional] 
**NativeFormat** | Pointer to **NullableBool** | Specifies whether or not to enable native format for direct archive job. This field is set to true if native format should be used for archiving. | [optional] 
**NfsVersionPreference** | Pointer to **NullableString** | Specifies the preference of NFS version to be backed up if a volume supports multiple versions of NFS. | [optional] 
**Objects** | [**[]NetappProtectionGroupObjectParams**](NetappProtectionGroupObjectParams.md) | Specifies the objects to be included in the Protection Group. | 
**PrePostScript** | Pointer to [**HostBasedBackupScriptParams**](HostBasedBackupScriptParams.md) |  | [optional] 
**Protocol** | Pointer to **NullableString** | Specifies the preferred protocol to use if this device supports multiple protocols. | [optional] 
**SnapMirrorConfig** | Pointer to [**SnapMirrorConfig**](SnapMirrorConfig.md) |  | [optional] 
**SnapshotLabel** | Pointer to [**SnapshotLabel**](SnapshotLabel.md) |  | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies the id of the parent of the objects. | [optional] [readonly] 
**SourceName** | Pointer to **NullableString** | Specifies the name of the parent of the objects. | [optional] [readonly] 
**ThrottlingConfig** | Pointer to [**NasThrottlingConfig**](NasThrottlingConfig.md) |  | [optional] 

## Methods

### NewNetappProtectionGroupParams

`func NewNetappProtectionGroupParams(objects []NetappProtectionGroupObjectParams, ) *NetappProtectionGroupParams`

NewNetappProtectionGroupParams instantiates a new NetappProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetappProtectionGroupParamsWithDefaults

`func NewNetappProtectionGroupParamsWithDefaults() *NetappProtectionGroupParams`

NewNetappProtectionGroupParamsWithDefaults instantiates a new NetappProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupExistingSnapshot

`func (o *NetappProtectionGroupParams) GetBackupExistingSnapshot() bool`

GetBackupExistingSnapshot returns the BackupExistingSnapshot field if non-nil, zero value otherwise.

### GetBackupExistingSnapshotOk

`func (o *NetappProtectionGroupParams) GetBackupExistingSnapshotOk() (*bool, bool)`

GetBackupExistingSnapshotOk returns a tuple with the BackupExistingSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupExistingSnapshot

`func (o *NetappProtectionGroupParams) SetBackupExistingSnapshot(v bool)`

SetBackupExistingSnapshot sets BackupExistingSnapshot field to given value.

### HasBackupExistingSnapshot

`func (o *NetappProtectionGroupParams) HasBackupExistingSnapshot() bool`

HasBackupExistingSnapshot returns a boolean if a field has been set.

### SetBackupExistingSnapshotNil

`func (o *NetappProtectionGroupParams) SetBackupExistingSnapshotNil(b bool)`

 SetBackupExistingSnapshotNil sets the value for BackupExistingSnapshot to be an explicit nil

### UnsetBackupExistingSnapshot
`func (o *NetappProtectionGroupParams) UnsetBackupExistingSnapshot()`

UnsetBackupExistingSnapshot ensures that no value is present for BackupExistingSnapshot, not even an explicit nil
### GetContinueOnError

`func (o *NetappProtectionGroupParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *NetappProtectionGroupParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *NetappProtectionGroupParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *NetappProtectionGroupParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *NetappProtectionGroupParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *NetappProtectionGroupParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetContinuousSnapshots

`func (o *NetappProtectionGroupParams) GetContinuousSnapshots() ContinuousSnapshotParams`

GetContinuousSnapshots returns the ContinuousSnapshots field if non-nil, zero value otherwise.

### GetContinuousSnapshotsOk

`func (o *NetappProtectionGroupParams) GetContinuousSnapshotsOk() (*ContinuousSnapshotParams, bool)`

GetContinuousSnapshotsOk returns a tuple with the ContinuousSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuousSnapshots

`func (o *NetappProtectionGroupParams) SetContinuousSnapshots(v ContinuousSnapshotParams)`

SetContinuousSnapshots sets ContinuousSnapshots field to given value.

### HasContinuousSnapshots

`func (o *NetappProtectionGroupParams) HasContinuousSnapshots() bool`

HasContinuousSnapshots returns a boolean if a field has been set.

### GetDirectCloudArchive

`func (o *NetappProtectionGroupParams) GetDirectCloudArchive() bool`

GetDirectCloudArchive returns the DirectCloudArchive field if non-nil, zero value otherwise.

### GetDirectCloudArchiveOk

`func (o *NetappProtectionGroupParams) GetDirectCloudArchiveOk() (*bool, bool)`

GetDirectCloudArchiveOk returns a tuple with the DirectCloudArchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectCloudArchive

`func (o *NetappProtectionGroupParams) SetDirectCloudArchive(v bool)`

SetDirectCloudArchive sets DirectCloudArchive field to given value.

### HasDirectCloudArchive

`func (o *NetappProtectionGroupParams) HasDirectCloudArchive() bool`

HasDirectCloudArchive returns a boolean if a field has been set.

### SetDirectCloudArchiveNil

`func (o *NetappProtectionGroupParams) SetDirectCloudArchiveNil(b bool)`

 SetDirectCloudArchiveNil sets the value for DirectCloudArchive to be an explicit nil

### UnsetDirectCloudArchive
`func (o *NetappProtectionGroupParams) UnsetDirectCloudArchive()`

UnsetDirectCloudArchive ensures that no value is present for DirectCloudArchive, not even an explicit nil
### GetEncryptionEnabled

`func (o *NetappProtectionGroupParams) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *NetappProtectionGroupParams) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *NetappProtectionGroupParams) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *NetappProtectionGroupParams) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### SetEncryptionEnabledNil

`func (o *NetappProtectionGroupParams) SetEncryptionEnabledNil(b bool)`

 SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil

### UnsetEncryptionEnabled
`func (o *NetappProtectionGroupParams) UnsetEncryptionEnabled()`

UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
### GetExcludeObjectIds

`func (o *NetappProtectionGroupParams) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *NetappProtectionGroupParams) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *NetappProtectionGroupParams) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *NetappProtectionGroupParams) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### SetExcludeObjectIdsNil

`func (o *NetappProtectionGroupParams) SetExcludeObjectIdsNil(b bool)`

 SetExcludeObjectIdsNil sets the value for ExcludeObjectIds to be an explicit nil

### UnsetExcludeObjectIds
`func (o *NetappProtectionGroupParams) UnsetExcludeObjectIds()`

UnsetExcludeObjectIds ensures that no value is present for ExcludeObjectIds, not even an explicit nil
### GetFileFilters

`func (o *NetappProtectionGroupParams) GetFileFilters() FileFilteringPolicy`

GetFileFilters returns the FileFilters field if non-nil, zero value otherwise.

### GetFileFiltersOk

`func (o *NetappProtectionGroupParams) GetFileFiltersOk() (*FileFilteringPolicy, bool)`

GetFileFiltersOk returns a tuple with the FileFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileFilters

`func (o *NetappProtectionGroupParams) SetFileFilters(v FileFilteringPolicy)`

SetFileFilters sets FileFilters field to given value.

### HasFileFilters

`func (o *NetappProtectionGroupParams) HasFileFilters() bool`

HasFileFilters returns a boolean if a field has been set.

### GetFileLockConfig

`func (o *NetappProtectionGroupParams) GetFileLockConfig() FileLevelDataLockConfig`

GetFileLockConfig returns the FileLockConfig field if non-nil, zero value otherwise.

### GetFileLockConfigOk

`func (o *NetappProtectionGroupParams) GetFileLockConfigOk() (*FileLevelDataLockConfig, bool)`

GetFileLockConfigOk returns a tuple with the FileLockConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLockConfig

`func (o *NetappProtectionGroupParams) SetFileLockConfig(v FileLevelDataLockConfig)`

SetFileLockConfig sets FileLockConfig field to given value.

### HasFileLockConfig

`func (o *NetappProtectionGroupParams) HasFileLockConfig() bool`

HasFileLockConfig returns a boolean if a field has been set.

### GetFilterIpConfig

`func (o *NetappProtectionGroupParams) GetFilterIpConfig() FilterIpConfig`

GetFilterIpConfig returns the FilterIpConfig field if non-nil, zero value otherwise.

### GetFilterIpConfigOk

`func (o *NetappProtectionGroupParams) GetFilterIpConfigOk() (*FilterIpConfig, bool)`

GetFilterIpConfigOk returns a tuple with the FilterIpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterIpConfig

`func (o *NetappProtectionGroupParams) SetFilterIpConfig(v FilterIpConfig)`

SetFilterIpConfig sets FilterIpConfig field to given value.

### HasFilterIpConfig

`func (o *NetappProtectionGroupParams) HasFilterIpConfig() bool`

HasFilterIpConfig returns a boolean if a field has been set.

### GetIndexingPolicy

`func (o *NetappProtectionGroupParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *NetappProtectionGroupParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *NetappProtectionGroupParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *NetappProtectionGroupParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetModifySourcePermissions

`func (o *NetappProtectionGroupParams) GetModifySourcePermissions() bool`

GetModifySourcePermissions returns the ModifySourcePermissions field if non-nil, zero value otherwise.

### GetModifySourcePermissionsOk

`func (o *NetappProtectionGroupParams) GetModifySourcePermissionsOk() (*bool, bool)`

GetModifySourcePermissionsOk returns a tuple with the ModifySourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifySourcePermissions

`func (o *NetappProtectionGroupParams) SetModifySourcePermissions(v bool)`

SetModifySourcePermissions sets ModifySourcePermissions field to given value.

### HasModifySourcePermissions

`func (o *NetappProtectionGroupParams) HasModifySourcePermissions() bool`

HasModifySourcePermissions returns a boolean if a field has been set.

### SetModifySourcePermissionsNil

`func (o *NetappProtectionGroupParams) SetModifySourcePermissionsNil(b bool)`

 SetModifySourcePermissionsNil sets the value for ModifySourcePermissions to be an explicit nil

### UnsetModifySourcePermissions
`func (o *NetappProtectionGroupParams) UnsetModifySourcePermissions()`

UnsetModifySourcePermissions ensures that no value is present for ModifySourcePermissions, not even an explicit nil
### GetNativeFormat

`func (o *NetappProtectionGroupParams) GetNativeFormat() bool`

GetNativeFormat returns the NativeFormat field if non-nil, zero value otherwise.

### GetNativeFormatOk

`func (o *NetappProtectionGroupParams) GetNativeFormatOk() (*bool, bool)`

GetNativeFormatOk returns a tuple with the NativeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeFormat

`func (o *NetappProtectionGroupParams) SetNativeFormat(v bool)`

SetNativeFormat sets NativeFormat field to given value.

### HasNativeFormat

`func (o *NetappProtectionGroupParams) HasNativeFormat() bool`

HasNativeFormat returns a boolean if a field has been set.

### SetNativeFormatNil

`func (o *NetappProtectionGroupParams) SetNativeFormatNil(b bool)`

 SetNativeFormatNil sets the value for NativeFormat to be an explicit nil

### UnsetNativeFormat
`func (o *NetappProtectionGroupParams) UnsetNativeFormat()`

UnsetNativeFormat ensures that no value is present for NativeFormat, not even an explicit nil
### GetNfsVersionPreference

`func (o *NetappProtectionGroupParams) GetNfsVersionPreference() string`

GetNfsVersionPreference returns the NfsVersionPreference field if non-nil, zero value otherwise.

### GetNfsVersionPreferenceOk

`func (o *NetappProtectionGroupParams) GetNfsVersionPreferenceOk() (*string, bool)`

GetNfsVersionPreferenceOk returns a tuple with the NfsVersionPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsVersionPreference

`func (o *NetappProtectionGroupParams) SetNfsVersionPreference(v string)`

SetNfsVersionPreference sets NfsVersionPreference field to given value.

### HasNfsVersionPreference

`func (o *NetappProtectionGroupParams) HasNfsVersionPreference() bool`

HasNfsVersionPreference returns a boolean if a field has been set.

### SetNfsVersionPreferenceNil

`func (o *NetappProtectionGroupParams) SetNfsVersionPreferenceNil(b bool)`

 SetNfsVersionPreferenceNil sets the value for NfsVersionPreference to be an explicit nil

### UnsetNfsVersionPreference
`func (o *NetappProtectionGroupParams) UnsetNfsVersionPreference()`

UnsetNfsVersionPreference ensures that no value is present for NfsVersionPreference, not even an explicit nil
### GetObjects

`func (o *NetappProtectionGroupParams) GetObjects() []NetappProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *NetappProtectionGroupParams) GetObjectsOk() (*[]NetappProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *NetappProtectionGroupParams) SetObjects(v []NetappProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.


### GetPrePostScript

`func (o *NetappProtectionGroupParams) GetPrePostScript() HostBasedBackupScriptParams`

GetPrePostScript returns the PrePostScript field if non-nil, zero value otherwise.

### GetPrePostScriptOk

`func (o *NetappProtectionGroupParams) GetPrePostScriptOk() (*HostBasedBackupScriptParams, bool)`

GetPrePostScriptOk returns a tuple with the PrePostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePostScript

`func (o *NetappProtectionGroupParams) SetPrePostScript(v HostBasedBackupScriptParams)`

SetPrePostScript sets PrePostScript field to given value.

### HasPrePostScript

`func (o *NetappProtectionGroupParams) HasPrePostScript() bool`

HasPrePostScript returns a boolean if a field has been set.

### GetProtocol

`func (o *NetappProtectionGroupParams) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NetappProtectionGroupParams) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NetappProtectionGroupParams) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *NetappProtectionGroupParams) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *NetappProtectionGroupParams) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *NetappProtectionGroupParams) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetSnapMirrorConfig

`func (o *NetappProtectionGroupParams) GetSnapMirrorConfig() SnapMirrorConfig`

GetSnapMirrorConfig returns the SnapMirrorConfig field if non-nil, zero value otherwise.

### GetSnapMirrorConfigOk

`func (o *NetappProtectionGroupParams) GetSnapMirrorConfigOk() (*SnapMirrorConfig, bool)`

GetSnapMirrorConfigOk returns a tuple with the SnapMirrorConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapMirrorConfig

`func (o *NetappProtectionGroupParams) SetSnapMirrorConfig(v SnapMirrorConfig)`

SetSnapMirrorConfig sets SnapMirrorConfig field to given value.

### HasSnapMirrorConfig

`func (o *NetappProtectionGroupParams) HasSnapMirrorConfig() bool`

HasSnapMirrorConfig returns a boolean if a field has been set.

### GetSnapshotLabel

`func (o *NetappProtectionGroupParams) GetSnapshotLabel() SnapshotLabel`

GetSnapshotLabel returns the SnapshotLabel field if non-nil, zero value otherwise.

### GetSnapshotLabelOk

`func (o *NetappProtectionGroupParams) GetSnapshotLabelOk() (*SnapshotLabel, bool)`

GetSnapshotLabelOk returns a tuple with the SnapshotLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotLabel

`func (o *NetappProtectionGroupParams) SetSnapshotLabel(v SnapshotLabel)`

SetSnapshotLabel sets SnapshotLabel field to given value.

### HasSnapshotLabel

`func (o *NetappProtectionGroupParams) HasSnapshotLabel() bool`

HasSnapshotLabel returns a boolean if a field has been set.

### GetSourceId

`func (o *NetappProtectionGroupParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *NetappProtectionGroupParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *NetappProtectionGroupParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *NetappProtectionGroupParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *NetappProtectionGroupParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *NetappProtectionGroupParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *NetappProtectionGroupParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *NetappProtectionGroupParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *NetappProtectionGroupParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *NetappProtectionGroupParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *NetappProtectionGroupParams) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *NetappProtectionGroupParams) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetThrottlingConfig

`func (o *NetappProtectionGroupParams) GetThrottlingConfig() NasThrottlingConfig`

GetThrottlingConfig returns the ThrottlingConfig field if non-nil, zero value otherwise.

### GetThrottlingConfigOk

`func (o *NetappProtectionGroupParams) GetThrottlingConfigOk() (*NasThrottlingConfig, bool)`

GetThrottlingConfigOk returns a tuple with the ThrottlingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingConfig

`func (o *NetappProtectionGroupParams) SetThrottlingConfig(v NasThrottlingConfig)`

SetThrottlingConfig sets ThrottlingConfig field to given value.

### HasThrottlingConfig

`func (o *NetappProtectionGroupParams) HasThrottlingConfig() bool`

HasThrottlingConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


