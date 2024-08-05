# NasEnvJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableAuditLogging** | Pointer to **NullableBool** | Specifies whether to audit log the file tiering activity. | [optional] [default to false]
**FilePath** | Pointer to [**FileFilteringPolicy**](FileFilteringPolicy.md) |  | [optional] 
**FileSize** | Pointer to [**FileSizePolicy**](FileSizePolicy.md) |  | [optional] 
**IncludeAllFiles** | Pointer to **NullableBool** | If set, all files in the view will be uptiered regardless of file_select_policy, num_file_access, hot_file_window, file_size constraints. | [optional] [default to false]
**Target** | Pointer to [**NullableDataTieringTarget**](DataTieringTarget.md) |  | [optional] 
**UptieringFileAge** | Pointer to [**UptieringFileAgePolicy**](UptieringFileAgePolicy.md) |  | [optional] 
**AutoOrphanDataCleanup** | Pointer to **NullableBool** | Specifies whether to remove the orphan data from the target if the symlink is removed from the source. | [optional] [default to true]
**DowntieringFileAge** | Pointer to [**DowntieringFileAgePolicy**](DowntieringFileAgePolicy.md) |  | [optional] 
**SkipBackSymlink** | Pointer to **NullableBool** | Specifies whether to create a symlink for the migrated data from source to target. | [optional] [default to true]
**BackupExistingSnapshot** | Pointer to **NullableBool** | Specifies that snapshot label is not set for Data-Protect Netapp Volumes backup. If field is set to true, existing oldest snapshot is used for backup and subsequent incremental will be selected in ascending order of snapshot create time on the source. If snapshot label is set, this field is set to false. | [optional] 
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether or not the Protection Group should continue regardless of whether or not an error was encountered during protection group run. | [optional] 
**EnableFasterIncrementalBackups** | Pointer to **NullableBool** | Specifies whether this job will enable faster incremental backups using change list or similar APIs | [optional] 
**EncryptionEnabled** | Pointer to **NullableBool** | Specifies whether the protection group should use encryption while backup or not. | [optional] 
**FileLockConfig** | Pointer to [**FileLevelDataLockConfig**](FileLevelDataLockConfig.md) |  | [optional] 
**FilePathFilters** | Pointer to [**FileFilteringPolicy**](FileFilteringPolicy.md) |  | [optional] 
**FilterIpConfig** | Pointer to [**FilterIpConfig**](FilterIpConfig.md) |  | [optional] 
**ModifySourcePermissions** | Pointer to **NullableBool** | Specifies if the NAS source permissions should be modified internally to allow backups. | [optional] 
**NasProtocol** | Pointer to **NullableString** | Specifies the preferred protocol to use if this device supports multiple protocols. | [optional] 
**NfsVersionPreference** | Pointer to **NullableString** | Specifies the preference of NFS version to be backed up if a volume supports multiple versions of NFS. | [optional] 
**SnapshotLabel** | Pointer to [**SnapshotLabel**](SnapshotLabel.md) |  | [optional] 
**ThrottlingConfig** | Pointer to [**NasThrottlingConfig**](NasThrottlingConfig.md) |  | [optional] 
**UseChangelist** | Pointer to **NullableBool** | Specify whether to use the Isilon Changelist API to directly discover changed files/directories for faster incremental backup. Cohesity will keep an extra snapshot which will be deleted by the next successful backup. | [optional] 

## Methods

### NewNasEnvJobParams

`func NewNasEnvJobParams() *NasEnvJobParams`

NewNasEnvJobParams instantiates a new NasEnvJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNasEnvJobParamsWithDefaults

`func NewNasEnvJobParamsWithDefaults() *NasEnvJobParams`

NewNasEnvJobParamsWithDefaults instantiates a new NasEnvJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableAuditLogging

`func (o *NasEnvJobParams) GetEnableAuditLogging() bool`

GetEnableAuditLogging returns the EnableAuditLogging field if non-nil, zero value otherwise.

### GetEnableAuditLoggingOk

`func (o *NasEnvJobParams) GetEnableAuditLoggingOk() (*bool, bool)`

GetEnableAuditLoggingOk returns a tuple with the EnableAuditLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAuditLogging

`func (o *NasEnvJobParams) SetEnableAuditLogging(v bool)`

SetEnableAuditLogging sets EnableAuditLogging field to given value.

### HasEnableAuditLogging

`func (o *NasEnvJobParams) HasEnableAuditLogging() bool`

HasEnableAuditLogging returns a boolean if a field has been set.

### SetEnableAuditLoggingNil

`func (o *NasEnvJobParams) SetEnableAuditLoggingNil(b bool)`

 SetEnableAuditLoggingNil sets the value for EnableAuditLogging to be an explicit nil

### UnsetEnableAuditLogging
`func (o *NasEnvJobParams) UnsetEnableAuditLogging()`

UnsetEnableAuditLogging ensures that no value is present for EnableAuditLogging, not even an explicit nil
### GetFilePath

`func (o *NasEnvJobParams) GetFilePath() FileFilteringPolicy`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *NasEnvJobParams) GetFilePathOk() (*FileFilteringPolicy, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *NasEnvJobParams) SetFilePath(v FileFilteringPolicy)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *NasEnvJobParams) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetFileSize

`func (o *NasEnvJobParams) GetFileSize() FileSizePolicy`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *NasEnvJobParams) GetFileSizeOk() (*FileSizePolicy, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *NasEnvJobParams) SetFileSize(v FileSizePolicy)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *NasEnvJobParams) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetIncludeAllFiles

`func (o *NasEnvJobParams) GetIncludeAllFiles() bool`

GetIncludeAllFiles returns the IncludeAllFiles field if non-nil, zero value otherwise.

### GetIncludeAllFilesOk

`func (o *NasEnvJobParams) GetIncludeAllFilesOk() (*bool, bool)`

GetIncludeAllFilesOk returns a tuple with the IncludeAllFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllFiles

`func (o *NasEnvJobParams) SetIncludeAllFiles(v bool)`

SetIncludeAllFiles sets IncludeAllFiles field to given value.

### HasIncludeAllFiles

`func (o *NasEnvJobParams) HasIncludeAllFiles() bool`

HasIncludeAllFiles returns a boolean if a field has been set.

### SetIncludeAllFilesNil

`func (o *NasEnvJobParams) SetIncludeAllFilesNil(b bool)`

 SetIncludeAllFilesNil sets the value for IncludeAllFiles to be an explicit nil

### UnsetIncludeAllFiles
`func (o *NasEnvJobParams) UnsetIncludeAllFiles()`

UnsetIncludeAllFiles ensures that no value is present for IncludeAllFiles, not even an explicit nil
### GetTarget

`func (o *NasEnvJobParams) GetTarget() DataTieringTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *NasEnvJobParams) GetTargetOk() (*DataTieringTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *NasEnvJobParams) SetTarget(v DataTieringTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *NasEnvJobParams) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *NasEnvJobParams) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *NasEnvJobParams) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetUptieringFileAge

`func (o *NasEnvJobParams) GetUptieringFileAge() UptieringFileAgePolicy`

GetUptieringFileAge returns the UptieringFileAge field if non-nil, zero value otherwise.

### GetUptieringFileAgeOk

`func (o *NasEnvJobParams) GetUptieringFileAgeOk() (*UptieringFileAgePolicy, bool)`

GetUptieringFileAgeOk returns a tuple with the UptieringFileAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptieringFileAge

`func (o *NasEnvJobParams) SetUptieringFileAge(v UptieringFileAgePolicy)`

SetUptieringFileAge sets UptieringFileAge field to given value.

### HasUptieringFileAge

`func (o *NasEnvJobParams) HasUptieringFileAge() bool`

HasUptieringFileAge returns a boolean if a field has been set.

### GetAutoOrphanDataCleanup

`func (o *NasEnvJobParams) GetAutoOrphanDataCleanup() bool`

GetAutoOrphanDataCleanup returns the AutoOrphanDataCleanup field if non-nil, zero value otherwise.

### GetAutoOrphanDataCleanupOk

`func (o *NasEnvJobParams) GetAutoOrphanDataCleanupOk() (*bool, bool)`

GetAutoOrphanDataCleanupOk returns a tuple with the AutoOrphanDataCleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoOrphanDataCleanup

`func (o *NasEnvJobParams) SetAutoOrphanDataCleanup(v bool)`

SetAutoOrphanDataCleanup sets AutoOrphanDataCleanup field to given value.

### HasAutoOrphanDataCleanup

`func (o *NasEnvJobParams) HasAutoOrphanDataCleanup() bool`

HasAutoOrphanDataCleanup returns a boolean if a field has been set.

### SetAutoOrphanDataCleanupNil

`func (o *NasEnvJobParams) SetAutoOrphanDataCleanupNil(b bool)`

 SetAutoOrphanDataCleanupNil sets the value for AutoOrphanDataCleanup to be an explicit nil

### UnsetAutoOrphanDataCleanup
`func (o *NasEnvJobParams) UnsetAutoOrphanDataCleanup()`

UnsetAutoOrphanDataCleanup ensures that no value is present for AutoOrphanDataCleanup, not even an explicit nil
### GetDowntieringFileAge

`func (o *NasEnvJobParams) GetDowntieringFileAge() DowntieringFileAgePolicy`

GetDowntieringFileAge returns the DowntieringFileAge field if non-nil, zero value otherwise.

### GetDowntieringFileAgeOk

`func (o *NasEnvJobParams) GetDowntieringFileAgeOk() (*DowntieringFileAgePolicy, bool)`

GetDowntieringFileAgeOk returns a tuple with the DowntieringFileAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowntieringFileAge

`func (o *NasEnvJobParams) SetDowntieringFileAge(v DowntieringFileAgePolicy)`

SetDowntieringFileAge sets DowntieringFileAge field to given value.

### HasDowntieringFileAge

`func (o *NasEnvJobParams) HasDowntieringFileAge() bool`

HasDowntieringFileAge returns a boolean if a field has been set.

### GetSkipBackSymlink

`func (o *NasEnvJobParams) GetSkipBackSymlink() bool`

GetSkipBackSymlink returns the SkipBackSymlink field if non-nil, zero value otherwise.

### GetSkipBackSymlinkOk

`func (o *NasEnvJobParams) GetSkipBackSymlinkOk() (*bool, bool)`

GetSkipBackSymlinkOk returns a tuple with the SkipBackSymlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipBackSymlink

`func (o *NasEnvJobParams) SetSkipBackSymlink(v bool)`

SetSkipBackSymlink sets SkipBackSymlink field to given value.

### HasSkipBackSymlink

`func (o *NasEnvJobParams) HasSkipBackSymlink() bool`

HasSkipBackSymlink returns a boolean if a field has been set.

### SetSkipBackSymlinkNil

`func (o *NasEnvJobParams) SetSkipBackSymlinkNil(b bool)`

 SetSkipBackSymlinkNil sets the value for SkipBackSymlink to be an explicit nil

### UnsetSkipBackSymlink
`func (o *NasEnvJobParams) UnsetSkipBackSymlink()`

UnsetSkipBackSymlink ensures that no value is present for SkipBackSymlink, not even an explicit nil
### GetBackupExistingSnapshot

`func (o *NasEnvJobParams) GetBackupExistingSnapshot() bool`

GetBackupExistingSnapshot returns the BackupExistingSnapshot field if non-nil, zero value otherwise.

### GetBackupExistingSnapshotOk

`func (o *NasEnvJobParams) GetBackupExistingSnapshotOk() (*bool, bool)`

GetBackupExistingSnapshotOk returns a tuple with the BackupExistingSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupExistingSnapshot

`func (o *NasEnvJobParams) SetBackupExistingSnapshot(v bool)`

SetBackupExistingSnapshot sets BackupExistingSnapshot field to given value.

### HasBackupExistingSnapshot

`func (o *NasEnvJobParams) HasBackupExistingSnapshot() bool`

HasBackupExistingSnapshot returns a boolean if a field has been set.

### SetBackupExistingSnapshotNil

`func (o *NasEnvJobParams) SetBackupExistingSnapshotNil(b bool)`

 SetBackupExistingSnapshotNil sets the value for BackupExistingSnapshot to be an explicit nil

### UnsetBackupExistingSnapshot
`func (o *NasEnvJobParams) UnsetBackupExistingSnapshot()`

UnsetBackupExistingSnapshot ensures that no value is present for BackupExistingSnapshot, not even an explicit nil
### GetContinueOnError

`func (o *NasEnvJobParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *NasEnvJobParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *NasEnvJobParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *NasEnvJobParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *NasEnvJobParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *NasEnvJobParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetEnableFasterIncrementalBackups

`func (o *NasEnvJobParams) GetEnableFasterIncrementalBackups() bool`

GetEnableFasterIncrementalBackups returns the EnableFasterIncrementalBackups field if non-nil, zero value otherwise.

### GetEnableFasterIncrementalBackupsOk

`func (o *NasEnvJobParams) GetEnableFasterIncrementalBackupsOk() (*bool, bool)`

GetEnableFasterIncrementalBackupsOk returns a tuple with the EnableFasterIncrementalBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFasterIncrementalBackups

`func (o *NasEnvJobParams) SetEnableFasterIncrementalBackups(v bool)`

SetEnableFasterIncrementalBackups sets EnableFasterIncrementalBackups field to given value.

### HasEnableFasterIncrementalBackups

`func (o *NasEnvJobParams) HasEnableFasterIncrementalBackups() bool`

HasEnableFasterIncrementalBackups returns a boolean if a field has been set.

### SetEnableFasterIncrementalBackupsNil

`func (o *NasEnvJobParams) SetEnableFasterIncrementalBackupsNil(b bool)`

 SetEnableFasterIncrementalBackupsNil sets the value for EnableFasterIncrementalBackups to be an explicit nil

### UnsetEnableFasterIncrementalBackups
`func (o *NasEnvJobParams) UnsetEnableFasterIncrementalBackups()`

UnsetEnableFasterIncrementalBackups ensures that no value is present for EnableFasterIncrementalBackups, not even an explicit nil
### GetEncryptionEnabled

`func (o *NasEnvJobParams) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *NasEnvJobParams) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *NasEnvJobParams) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *NasEnvJobParams) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### SetEncryptionEnabledNil

`func (o *NasEnvJobParams) SetEncryptionEnabledNil(b bool)`

 SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil

### UnsetEncryptionEnabled
`func (o *NasEnvJobParams) UnsetEncryptionEnabled()`

UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
### GetFileLockConfig

`func (o *NasEnvJobParams) GetFileLockConfig() FileLevelDataLockConfig`

GetFileLockConfig returns the FileLockConfig field if non-nil, zero value otherwise.

### GetFileLockConfigOk

`func (o *NasEnvJobParams) GetFileLockConfigOk() (*FileLevelDataLockConfig, bool)`

GetFileLockConfigOk returns a tuple with the FileLockConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLockConfig

`func (o *NasEnvJobParams) SetFileLockConfig(v FileLevelDataLockConfig)`

SetFileLockConfig sets FileLockConfig field to given value.

### HasFileLockConfig

`func (o *NasEnvJobParams) HasFileLockConfig() bool`

HasFileLockConfig returns a boolean if a field has been set.

### GetFilePathFilters

`func (o *NasEnvJobParams) GetFilePathFilters() FileFilteringPolicy`

GetFilePathFilters returns the FilePathFilters field if non-nil, zero value otherwise.

### GetFilePathFiltersOk

`func (o *NasEnvJobParams) GetFilePathFiltersOk() (*FileFilteringPolicy, bool)`

GetFilePathFiltersOk returns a tuple with the FilePathFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePathFilters

`func (o *NasEnvJobParams) SetFilePathFilters(v FileFilteringPolicy)`

SetFilePathFilters sets FilePathFilters field to given value.

### HasFilePathFilters

`func (o *NasEnvJobParams) HasFilePathFilters() bool`

HasFilePathFilters returns a boolean if a field has been set.

### GetFilterIpConfig

`func (o *NasEnvJobParams) GetFilterIpConfig() FilterIpConfig`

GetFilterIpConfig returns the FilterIpConfig field if non-nil, zero value otherwise.

### GetFilterIpConfigOk

`func (o *NasEnvJobParams) GetFilterIpConfigOk() (*FilterIpConfig, bool)`

GetFilterIpConfigOk returns a tuple with the FilterIpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterIpConfig

`func (o *NasEnvJobParams) SetFilterIpConfig(v FilterIpConfig)`

SetFilterIpConfig sets FilterIpConfig field to given value.

### HasFilterIpConfig

`func (o *NasEnvJobParams) HasFilterIpConfig() bool`

HasFilterIpConfig returns a boolean if a field has been set.

### GetModifySourcePermissions

`func (o *NasEnvJobParams) GetModifySourcePermissions() bool`

GetModifySourcePermissions returns the ModifySourcePermissions field if non-nil, zero value otherwise.

### GetModifySourcePermissionsOk

`func (o *NasEnvJobParams) GetModifySourcePermissionsOk() (*bool, bool)`

GetModifySourcePermissionsOk returns a tuple with the ModifySourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifySourcePermissions

`func (o *NasEnvJobParams) SetModifySourcePermissions(v bool)`

SetModifySourcePermissions sets ModifySourcePermissions field to given value.

### HasModifySourcePermissions

`func (o *NasEnvJobParams) HasModifySourcePermissions() bool`

HasModifySourcePermissions returns a boolean if a field has been set.

### SetModifySourcePermissionsNil

`func (o *NasEnvJobParams) SetModifySourcePermissionsNil(b bool)`

 SetModifySourcePermissionsNil sets the value for ModifySourcePermissions to be an explicit nil

### UnsetModifySourcePermissions
`func (o *NasEnvJobParams) UnsetModifySourcePermissions()`

UnsetModifySourcePermissions ensures that no value is present for ModifySourcePermissions, not even an explicit nil
### GetNasProtocol

`func (o *NasEnvJobParams) GetNasProtocol() string`

GetNasProtocol returns the NasProtocol field if non-nil, zero value otherwise.

### GetNasProtocolOk

`func (o *NasEnvJobParams) GetNasProtocolOk() (*string, bool)`

GetNasProtocolOk returns a tuple with the NasProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasProtocol

`func (o *NasEnvJobParams) SetNasProtocol(v string)`

SetNasProtocol sets NasProtocol field to given value.

### HasNasProtocol

`func (o *NasEnvJobParams) HasNasProtocol() bool`

HasNasProtocol returns a boolean if a field has been set.

### SetNasProtocolNil

`func (o *NasEnvJobParams) SetNasProtocolNil(b bool)`

 SetNasProtocolNil sets the value for NasProtocol to be an explicit nil

### UnsetNasProtocol
`func (o *NasEnvJobParams) UnsetNasProtocol()`

UnsetNasProtocol ensures that no value is present for NasProtocol, not even an explicit nil
### GetNfsVersionPreference

`func (o *NasEnvJobParams) GetNfsVersionPreference() string`

GetNfsVersionPreference returns the NfsVersionPreference field if non-nil, zero value otherwise.

### GetNfsVersionPreferenceOk

`func (o *NasEnvJobParams) GetNfsVersionPreferenceOk() (*string, bool)`

GetNfsVersionPreferenceOk returns a tuple with the NfsVersionPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsVersionPreference

`func (o *NasEnvJobParams) SetNfsVersionPreference(v string)`

SetNfsVersionPreference sets NfsVersionPreference field to given value.

### HasNfsVersionPreference

`func (o *NasEnvJobParams) HasNfsVersionPreference() bool`

HasNfsVersionPreference returns a boolean if a field has been set.

### SetNfsVersionPreferenceNil

`func (o *NasEnvJobParams) SetNfsVersionPreferenceNil(b bool)`

 SetNfsVersionPreferenceNil sets the value for NfsVersionPreference to be an explicit nil

### UnsetNfsVersionPreference
`func (o *NasEnvJobParams) UnsetNfsVersionPreference()`

UnsetNfsVersionPreference ensures that no value is present for NfsVersionPreference, not even an explicit nil
### GetSnapshotLabel

`func (o *NasEnvJobParams) GetSnapshotLabel() SnapshotLabel`

GetSnapshotLabel returns the SnapshotLabel field if non-nil, zero value otherwise.

### GetSnapshotLabelOk

`func (o *NasEnvJobParams) GetSnapshotLabelOk() (*SnapshotLabel, bool)`

GetSnapshotLabelOk returns a tuple with the SnapshotLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotLabel

`func (o *NasEnvJobParams) SetSnapshotLabel(v SnapshotLabel)`

SetSnapshotLabel sets SnapshotLabel field to given value.

### HasSnapshotLabel

`func (o *NasEnvJobParams) HasSnapshotLabel() bool`

HasSnapshotLabel returns a boolean if a field has been set.

### GetThrottlingConfig

`func (o *NasEnvJobParams) GetThrottlingConfig() NasThrottlingConfig`

GetThrottlingConfig returns the ThrottlingConfig field if non-nil, zero value otherwise.

### GetThrottlingConfigOk

`func (o *NasEnvJobParams) GetThrottlingConfigOk() (*NasThrottlingConfig, bool)`

GetThrottlingConfigOk returns a tuple with the ThrottlingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingConfig

`func (o *NasEnvJobParams) SetThrottlingConfig(v NasThrottlingConfig)`

SetThrottlingConfig sets ThrottlingConfig field to given value.

### HasThrottlingConfig

`func (o *NasEnvJobParams) HasThrottlingConfig() bool`

HasThrottlingConfig returns a boolean if a field has been set.

### GetUseChangelist

`func (o *NasEnvJobParams) GetUseChangelist() bool`

GetUseChangelist returns the UseChangelist field if non-nil, zero value otherwise.

### GetUseChangelistOk

`func (o *NasEnvJobParams) GetUseChangelistOk() (*bool, bool)`

GetUseChangelistOk returns a tuple with the UseChangelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseChangelist

`func (o *NasEnvJobParams) SetUseChangelist(v bool)`

SetUseChangelist sets UseChangelist field to given value.

### HasUseChangelist

`func (o *NasEnvJobParams) HasUseChangelist() bool`

HasUseChangelist returns a boolean if a field has been set.

### SetUseChangelistNil

`func (o *NasEnvJobParams) SetUseChangelistNil(b bool)`

 SetUseChangelistNil sets the value for UseChangelist to be an explicit nil

### UnsetUseChangelist
`func (o *NasEnvJobParams) UnsetUseChangelist()`

UnsetUseChangelist ensures that no value is present for UseChangelist, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


