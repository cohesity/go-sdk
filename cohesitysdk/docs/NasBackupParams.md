# NasBackupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupExistingSnapshot** | Pointer to **NullableBool** | This bool parameter will be set only for DP volumes when customer doesn&#39;t select the full_backup_snapshot_label and incremental_backup_snapshot_label. When set to true, backend will be using existing oldest snapshot for the first backup. Each incremental will be selected in ascending of snapshot create time on the source. | [optional] 
**BlacklistedIpAddrs** | Pointer to **[]string** | A list of IP addresses that should not be used. | [optional] 
**ContinueOnError** | Pointer to **NullableBool** | Whether the backup job should continue on errors for snapshot based backups. For non-snapshot-based generic NAS backup jobs, Magneto always continues on errors. | [optional] 
**EncryptionEnabled** | Pointer to **NullableBool** | Whether this backup job should use encryption. | [optional] 
**FilteringPolicy** | Pointer to [**FilteringPolicyProto**](FilteringPolicyProto.md) |  | [optional] 
**FldConfig** | Pointer to [**ViewIdMappingProtoFileLevelDataLockConfig**](ViewIdMappingProtoFileLevelDataLockConfig.md) |  | [optional] 
**FullBackupSnapshotLabel** | Pointer to **NullableString** | Only used when we backup using discovered snapshots. This prefix is to figure out which discovered snapshot we need to use for full backup. | [optional] 
**IncrementalBackupSnapshotLabel** | Pointer to **NullableString** | Only used when we backup using discovered snapshots. This prefix is to figure out which discovered snapshot we need to use for incremental backup. | [optional] 
**IsSourceInitiatedBackup** | Pointer to **NullableBool** | Source initiated backup when the source sends pushes the data like for example snapmirror based backup for netapp. | [optional] 
**MixedModePreference** | Pointer to **NullableInt32** | If the target entity is a mixed mode volume, which NAS protocol type the user prefer to backup. This does not apply to generic NAS and will be ignored. | [optional] 
**S3Viewbackupproperties** | Pointer to [**S3ViewBackupProperties**](S3ViewBackupProperties.md) |  | [optional] 
**SnapshotChangeEnabled** | Pointer to **NullableBool** | Whether this backup job should utilize changelist like API when available for faster incremental backups. | [optional] 

## Methods

### NewNasBackupParams

`func NewNasBackupParams() *NasBackupParams`

NewNasBackupParams instantiates a new NasBackupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNasBackupParamsWithDefaults

`func NewNasBackupParamsWithDefaults() *NasBackupParams`

NewNasBackupParamsWithDefaults instantiates a new NasBackupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupExistingSnapshot

`func (o *NasBackupParams) GetBackupExistingSnapshot() bool`

GetBackupExistingSnapshot returns the BackupExistingSnapshot field if non-nil, zero value otherwise.

### GetBackupExistingSnapshotOk

`func (o *NasBackupParams) GetBackupExistingSnapshotOk() (*bool, bool)`

GetBackupExistingSnapshotOk returns a tuple with the BackupExistingSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupExistingSnapshot

`func (o *NasBackupParams) SetBackupExistingSnapshot(v bool)`

SetBackupExistingSnapshot sets BackupExistingSnapshot field to given value.

### HasBackupExistingSnapshot

`func (o *NasBackupParams) HasBackupExistingSnapshot() bool`

HasBackupExistingSnapshot returns a boolean if a field has been set.

### SetBackupExistingSnapshotNil

`func (o *NasBackupParams) SetBackupExistingSnapshotNil(b bool)`

 SetBackupExistingSnapshotNil sets the value for BackupExistingSnapshot to be an explicit nil

### UnsetBackupExistingSnapshot
`func (o *NasBackupParams) UnsetBackupExistingSnapshot()`

UnsetBackupExistingSnapshot ensures that no value is present for BackupExistingSnapshot, not even an explicit nil
### GetBlacklistedIpAddrs

`func (o *NasBackupParams) GetBlacklistedIpAddrs() []string`

GetBlacklistedIpAddrs returns the BlacklistedIpAddrs field if non-nil, zero value otherwise.

### GetBlacklistedIpAddrsOk

`func (o *NasBackupParams) GetBlacklistedIpAddrsOk() (*[]string, bool)`

GetBlacklistedIpAddrsOk returns a tuple with the BlacklistedIpAddrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistedIpAddrs

`func (o *NasBackupParams) SetBlacklistedIpAddrs(v []string)`

SetBlacklistedIpAddrs sets BlacklistedIpAddrs field to given value.

### HasBlacklistedIpAddrs

`func (o *NasBackupParams) HasBlacklistedIpAddrs() bool`

HasBlacklistedIpAddrs returns a boolean if a field has been set.

### SetBlacklistedIpAddrsNil

`func (o *NasBackupParams) SetBlacklistedIpAddrsNil(b bool)`

 SetBlacklistedIpAddrsNil sets the value for BlacklistedIpAddrs to be an explicit nil

### UnsetBlacklistedIpAddrs
`func (o *NasBackupParams) UnsetBlacklistedIpAddrs()`

UnsetBlacklistedIpAddrs ensures that no value is present for BlacklistedIpAddrs, not even an explicit nil
### GetContinueOnError

`func (o *NasBackupParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *NasBackupParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *NasBackupParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *NasBackupParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *NasBackupParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *NasBackupParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetEncryptionEnabled

`func (o *NasBackupParams) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *NasBackupParams) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *NasBackupParams) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *NasBackupParams) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### SetEncryptionEnabledNil

`func (o *NasBackupParams) SetEncryptionEnabledNil(b bool)`

 SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil

### UnsetEncryptionEnabled
`func (o *NasBackupParams) UnsetEncryptionEnabled()`

UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
### GetFilteringPolicy

`func (o *NasBackupParams) GetFilteringPolicy() FilteringPolicyProto`

GetFilteringPolicy returns the FilteringPolicy field if non-nil, zero value otherwise.

### GetFilteringPolicyOk

`func (o *NasBackupParams) GetFilteringPolicyOk() (*FilteringPolicyProto, bool)`

GetFilteringPolicyOk returns a tuple with the FilteringPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteringPolicy

`func (o *NasBackupParams) SetFilteringPolicy(v FilteringPolicyProto)`

SetFilteringPolicy sets FilteringPolicy field to given value.

### HasFilteringPolicy

`func (o *NasBackupParams) HasFilteringPolicy() bool`

HasFilteringPolicy returns a boolean if a field has been set.

### GetFldConfig

`func (o *NasBackupParams) GetFldConfig() ViewIdMappingProtoFileLevelDataLockConfig`

GetFldConfig returns the FldConfig field if non-nil, zero value otherwise.

### GetFldConfigOk

`func (o *NasBackupParams) GetFldConfigOk() (*ViewIdMappingProtoFileLevelDataLockConfig, bool)`

GetFldConfigOk returns a tuple with the FldConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFldConfig

`func (o *NasBackupParams) SetFldConfig(v ViewIdMappingProtoFileLevelDataLockConfig)`

SetFldConfig sets FldConfig field to given value.

### HasFldConfig

`func (o *NasBackupParams) HasFldConfig() bool`

HasFldConfig returns a boolean if a field has been set.

### GetFullBackupSnapshotLabel

`func (o *NasBackupParams) GetFullBackupSnapshotLabel() string`

GetFullBackupSnapshotLabel returns the FullBackupSnapshotLabel field if non-nil, zero value otherwise.

### GetFullBackupSnapshotLabelOk

`func (o *NasBackupParams) GetFullBackupSnapshotLabelOk() (*string, bool)`

GetFullBackupSnapshotLabelOk returns a tuple with the FullBackupSnapshotLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullBackupSnapshotLabel

`func (o *NasBackupParams) SetFullBackupSnapshotLabel(v string)`

SetFullBackupSnapshotLabel sets FullBackupSnapshotLabel field to given value.

### HasFullBackupSnapshotLabel

`func (o *NasBackupParams) HasFullBackupSnapshotLabel() bool`

HasFullBackupSnapshotLabel returns a boolean if a field has been set.

### SetFullBackupSnapshotLabelNil

`func (o *NasBackupParams) SetFullBackupSnapshotLabelNil(b bool)`

 SetFullBackupSnapshotLabelNil sets the value for FullBackupSnapshotLabel to be an explicit nil

### UnsetFullBackupSnapshotLabel
`func (o *NasBackupParams) UnsetFullBackupSnapshotLabel()`

UnsetFullBackupSnapshotLabel ensures that no value is present for FullBackupSnapshotLabel, not even an explicit nil
### GetIncrementalBackupSnapshotLabel

`func (o *NasBackupParams) GetIncrementalBackupSnapshotLabel() string`

GetIncrementalBackupSnapshotLabel returns the IncrementalBackupSnapshotLabel field if non-nil, zero value otherwise.

### GetIncrementalBackupSnapshotLabelOk

`func (o *NasBackupParams) GetIncrementalBackupSnapshotLabelOk() (*string, bool)`

GetIncrementalBackupSnapshotLabelOk returns a tuple with the IncrementalBackupSnapshotLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalBackupSnapshotLabel

`func (o *NasBackupParams) SetIncrementalBackupSnapshotLabel(v string)`

SetIncrementalBackupSnapshotLabel sets IncrementalBackupSnapshotLabel field to given value.

### HasIncrementalBackupSnapshotLabel

`func (o *NasBackupParams) HasIncrementalBackupSnapshotLabel() bool`

HasIncrementalBackupSnapshotLabel returns a boolean if a field has been set.

### SetIncrementalBackupSnapshotLabelNil

`func (o *NasBackupParams) SetIncrementalBackupSnapshotLabelNil(b bool)`

 SetIncrementalBackupSnapshotLabelNil sets the value for IncrementalBackupSnapshotLabel to be an explicit nil

### UnsetIncrementalBackupSnapshotLabel
`func (o *NasBackupParams) UnsetIncrementalBackupSnapshotLabel()`

UnsetIncrementalBackupSnapshotLabel ensures that no value is present for IncrementalBackupSnapshotLabel, not even an explicit nil
### GetIsSourceInitiatedBackup

`func (o *NasBackupParams) GetIsSourceInitiatedBackup() bool`

GetIsSourceInitiatedBackup returns the IsSourceInitiatedBackup field if non-nil, zero value otherwise.

### GetIsSourceInitiatedBackupOk

`func (o *NasBackupParams) GetIsSourceInitiatedBackupOk() (*bool, bool)`

GetIsSourceInitiatedBackupOk returns a tuple with the IsSourceInitiatedBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSourceInitiatedBackup

`func (o *NasBackupParams) SetIsSourceInitiatedBackup(v bool)`

SetIsSourceInitiatedBackup sets IsSourceInitiatedBackup field to given value.

### HasIsSourceInitiatedBackup

`func (o *NasBackupParams) HasIsSourceInitiatedBackup() bool`

HasIsSourceInitiatedBackup returns a boolean if a field has been set.

### SetIsSourceInitiatedBackupNil

`func (o *NasBackupParams) SetIsSourceInitiatedBackupNil(b bool)`

 SetIsSourceInitiatedBackupNil sets the value for IsSourceInitiatedBackup to be an explicit nil

### UnsetIsSourceInitiatedBackup
`func (o *NasBackupParams) UnsetIsSourceInitiatedBackup()`

UnsetIsSourceInitiatedBackup ensures that no value is present for IsSourceInitiatedBackup, not even an explicit nil
### GetMixedModePreference

`func (o *NasBackupParams) GetMixedModePreference() int32`

GetMixedModePreference returns the MixedModePreference field if non-nil, zero value otherwise.

### GetMixedModePreferenceOk

`func (o *NasBackupParams) GetMixedModePreferenceOk() (*int32, bool)`

GetMixedModePreferenceOk returns a tuple with the MixedModePreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMixedModePreference

`func (o *NasBackupParams) SetMixedModePreference(v int32)`

SetMixedModePreference sets MixedModePreference field to given value.

### HasMixedModePreference

`func (o *NasBackupParams) HasMixedModePreference() bool`

HasMixedModePreference returns a boolean if a field has been set.

### SetMixedModePreferenceNil

`func (o *NasBackupParams) SetMixedModePreferenceNil(b bool)`

 SetMixedModePreferenceNil sets the value for MixedModePreference to be an explicit nil

### UnsetMixedModePreference
`func (o *NasBackupParams) UnsetMixedModePreference()`

UnsetMixedModePreference ensures that no value is present for MixedModePreference, not even an explicit nil
### GetS3Viewbackupproperties

`func (o *NasBackupParams) GetS3Viewbackupproperties() S3ViewBackupProperties`

GetS3Viewbackupproperties returns the S3Viewbackupproperties field if non-nil, zero value otherwise.

### GetS3ViewbackuppropertiesOk

`func (o *NasBackupParams) GetS3ViewbackuppropertiesOk() (*S3ViewBackupProperties, bool)`

GetS3ViewbackuppropertiesOk returns a tuple with the S3Viewbackupproperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Viewbackupproperties

`func (o *NasBackupParams) SetS3Viewbackupproperties(v S3ViewBackupProperties)`

SetS3Viewbackupproperties sets S3Viewbackupproperties field to given value.

### HasS3Viewbackupproperties

`func (o *NasBackupParams) HasS3Viewbackupproperties() bool`

HasS3Viewbackupproperties returns a boolean if a field has been set.

### GetSnapshotChangeEnabled

`func (o *NasBackupParams) GetSnapshotChangeEnabled() bool`

GetSnapshotChangeEnabled returns the SnapshotChangeEnabled field if non-nil, zero value otherwise.

### GetSnapshotChangeEnabledOk

`func (o *NasBackupParams) GetSnapshotChangeEnabledOk() (*bool, bool)`

GetSnapshotChangeEnabledOk returns a tuple with the SnapshotChangeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotChangeEnabled

`func (o *NasBackupParams) SetSnapshotChangeEnabled(v bool)`

SetSnapshotChangeEnabled sets SnapshotChangeEnabled field to given value.

### HasSnapshotChangeEnabled

`func (o *NasBackupParams) HasSnapshotChangeEnabled() bool`

HasSnapshotChangeEnabled returns a boolean if a field has been set.

### SetSnapshotChangeEnabledNil

`func (o *NasBackupParams) SetSnapshotChangeEnabledNil(b bool)`

 SetSnapshotChangeEnabledNil sets the value for SnapshotChangeEnabled to be an explicit nil

### UnsetSnapshotChangeEnabled
`func (o *NasBackupParams) UnsetSnapshotChangeEnabled()`

UnsetSnapshotChangeEnabled ensures that no value is present for SnapshotChangeEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


