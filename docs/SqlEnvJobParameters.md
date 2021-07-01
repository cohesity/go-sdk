# SqlEnvJobParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AagPreference** | Pointer to **NullableString** | Specifies the preference for backing up databases that are part of an AAG. Only applicable if &#39;aagPreferenceFromSqlServer&#39; is set to false or not given. kPrimaryReplicaOnly implies backups should always occur on the primary replica. kSecondaryReplicaOnly implies backups should always occur on the secondary replica. kPreferSecondaryReplica implies secondary replica is preferred for backups. kAnyReplica implies no preference of about whether backups are performed on the primary replica or on a secondary replica. If no secondary replica is available, then performing backups on the primary replica is acceptable. | [optional] 
**AagPreferenceFromSqlServer** | Pointer to **NullableBool** | If true, AAG preferences are taken from the SQL server host. If this is set to false or not given, preferences can be overridden by aagBackupPreference. | [optional] 
**BackupSystemDatabases** | Pointer to **NullableBool** | If true, system databases are backed up. If this is set to false, system databases are not backed up. If this field is not specified, default value is true. | [optional] 
**BackupType** | Pointer to **NullableString** | Specifies the type of the &#39;kFull&#39; backup job. Specifies whether it is Volume level backup or individual files level backup. kSqlVSSVolume implies volume based VSS full backup. kSqlVSSFile implies file based VSS full backup. | [optional] 
**BackupVolumesOnly** | Pointer to **NullableBool** | If set to true, only the volumes associated with databases should be backed up. The user cannot select additional volumes at host level for backup.  If set to false, all the volumes on the host machine will be backed up. In this case, the user can further select the exact set of volumes using host level params.  Note that the volumes associated with selected databases will always be included in the backup. | [optional] 
**IncrementalSnapshotUponRestart** | Pointer to **NullableBool** | If true, the backup of type kSqlVssVolume will be incremental after restart | [optional] 
**IsCopyOnlyFull** | Pointer to **NullableBool** | If true, the backup is a full backup with the copy-only option specified. | [optional] 
**NumStreams** | Pointer to **NullableInt32** | Number of streams to be used in native sql backup. | [optional] 
**UserDatabasePreference** | Pointer to **NullableString** | Specifies the preference for backing up user databases on the host. kBackupAllDatabases implies to backup all databases. kBackupAllExceptAAGDatabases implies not to backup AAG databases. kBackupOnlyAAGDatabases implies to backup only AAG databases. | [optional] 
**WithClause** | Pointer to **NullableString** | With clause is used for setting clauese in native sql backup. | [optional] 

## Methods

### NewSqlEnvJobParameters

`func NewSqlEnvJobParameters() *SqlEnvJobParameters`

NewSqlEnvJobParameters instantiates a new SqlEnvJobParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlEnvJobParametersWithDefaults

`func NewSqlEnvJobParametersWithDefaults() *SqlEnvJobParameters`

NewSqlEnvJobParametersWithDefaults instantiates a new SqlEnvJobParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAagPreference

`func (o *SqlEnvJobParameters) GetAagPreference() string`

GetAagPreference returns the AagPreference field if non-nil, zero value otherwise.

### GetAagPreferenceOk

`func (o *SqlEnvJobParameters) GetAagPreferenceOk() (*string, bool)`

GetAagPreferenceOk returns a tuple with the AagPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAagPreference

`func (o *SqlEnvJobParameters) SetAagPreference(v string)`

SetAagPreference sets AagPreference field to given value.

### HasAagPreference

`func (o *SqlEnvJobParameters) HasAagPreference() bool`

HasAagPreference returns a boolean if a field has been set.

### SetAagPreferenceNil

`func (o *SqlEnvJobParameters) SetAagPreferenceNil(b bool)`

 SetAagPreferenceNil sets the value for AagPreference to be an explicit nil

### UnsetAagPreference
`func (o *SqlEnvJobParameters) UnsetAagPreference()`

UnsetAagPreference ensures that no value is present for AagPreference, not even an explicit nil
### GetAagPreferenceFromSqlServer

`func (o *SqlEnvJobParameters) GetAagPreferenceFromSqlServer() bool`

GetAagPreferenceFromSqlServer returns the AagPreferenceFromSqlServer field if non-nil, zero value otherwise.

### GetAagPreferenceFromSqlServerOk

`func (o *SqlEnvJobParameters) GetAagPreferenceFromSqlServerOk() (*bool, bool)`

GetAagPreferenceFromSqlServerOk returns a tuple with the AagPreferenceFromSqlServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAagPreferenceFromSqlServer

`func (o *SqlEnvJobParameters) SetAagPreferenceFromSqlServer(v bool)`

SetAagPreferenceFromSqlServer sets AagPreferenceFromSqlServer field to given value.

### HasAagPreferenceFromSqlServer

`func (o *SqlEnvJobParameters) HasAagPreferenceFromSqlServer() bool`

HasAagPreferenceFromSqlServer returns a boolean if a field has been set.

### SetAagPreferenceFromSqlServerNil

`func (o *SqlEnvJobParameters) SetAagPreferenceFromSqlServerNil(b bool)`

 SetAagPreferenceFromSqlServerNil sets the value for AagPreferenceFromSqlServer to be an explicit nil

### UnsetAagPreferenceFromSqlServer
`func (o *SqlEnvJobParameters) UnsetAagPreferenceFromSqlServer()`

UnsetAagPreferenceFromSqlServer ensures that no value is present for AagPreferenceFromSqlServer, not even an explicit nil
### GetBackupSystemDatabases

`func (o *SqlEnvJobParameters) GetBackupSystemDatabases() bool`

GetBackupSystemDatabases returns the BackupSystemDatabases field if non-nil, zero value otherwise.

### GetBackupSystemDatabasesOk

`func (o *SqlEnvJobParameters) GetBackupSystemDatabasesOk() (*bool, bool)`

GetBackupSystemDatabasesOk returns a tuple with the BackupSystemDatabases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSystemDatabases

`func (o *SqlEnvJobParameters) SetBackupSystemDatabases(v bool)`

SetBackupSystemDatabases sets BackupSystemDatabases field to given value.

### HasBackupSystemDatabases

`func (o *SqlEnvJobParameters) HasBackupSystemDatabases() bool`

HasBackupSystemDatabases returns a boolean if a field has been set.

### SetBackupSystemDatabasesNil

`func (o *SqlEnvJobParameters) SetBackupSystemDatabasesNil(b bool)`

 SetBackupSystemDatabasesNil sets the value for BackupSystemDatabases to be an explicit nil

### UnsetBackupSystemDatabases
`func (o *SqlEnvJobParameters) UnsetBackupSystemDatabases()`

UnsetBackupSystemDatabases ensures that no value is present for BackupSystemDatabases, not even an explicit nil
### GetBackupType

`func (o *SqlEnvJobParameters) GetBackupType() string`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *SqlEnvJobParameters) GetBackupTypeOk() (*string, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *SqlEnvJobParameters) SetBackupType(v string)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *SqlEnvJobParameters) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.

### SetBackupTypeNil

`func (o *SqlEnvJobParameters) SetBackupTypeNil(b bool)`

 SetBackupTypeNil sets the value for BackupType to be an explicit nil

### UnsetBackupType
`func (o *SqlEnvJobParameters) UnsetBackupType()`

UnsetBackupType ensures that no value is present for BackupType, not even an explicit nil
### GetBackupVolumesOnly

`func (o *SqlEnvJobParameters) GetBackupVolumesOnly() bool`

GetBackupVolumesOnly returns the BackupVolumesOnly field if non-nil, zero value otherwise.

### GetBackupVolumesOnlyOk

`func (o *SqlEnvJobParameters) GetBackupVolumesOnlyOk() (*bool, bool)`

GetBackupVolumesOnlyOk returns a tuple with the BackupVolumesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupVolumesOnly

`func (o *SqlEnvJobParameters) SetBackupVolumesOnly(v bool)`

SetBackupVolumesOnly sets BackupVolumesOnly field to given value.

### HasBackupVolumesOnly

`func (o *SqlEnvJobParameters) HasBackupVolumesOnly() bool`

HasBackupVolumesOnly returns a boolean if a field has been set.

### SetBackupVolumesOnlyNil

`func (o *SqlEnvJobParameters) SetBackupVolumesOnlyNil(b bool)`

 SetBackupVolumesOnlyNil sets the value for BackupVolumesOnly to be an explicit nil

### UnsetBackupVolumesOnly
`func (o *SqlEnvJobParameters) UnsetBackupVolumesOnly()`

UnsetBackupVolumesOnly ensures that no value is present for BackupVolumesOnly, not even an explicit nil
### GetIncrementalSnapshotUponRestart

`func (o *SqlEnvJobParameters) GetIncrementalSnapshotUponRestart() bool`

GetIncrementalSnapshotUponRestart returns the IncrementalSnapshotUponRestart field if non-nil, zero value otherwise.

### GetIncrementalSnapshotUponRestartOk

`func (o *SqlEnvJobParameters) GetIncrementalSnapshotUponRestartOk() (*bool, bool)`

GetIncrementalSnapshotUponRestartOk returns a tuple with the IncrementalSnapshotUponRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalSnapshotUponRestart

`func (o *SqlEnvJobParameters) SetIncrementalSnapshotUponRestart(v bool)`

SetIncrementalSnapshotUponRestart sets IncrementalSnapshotUponRestart field to given value.

### HasIncrementalSnapshotUponRestart

`func (o *SqlEnvJobParameters) HasIncrementalSnapshotUponRestart() bool`

HasIncrementalSnapshotUponRestart returns a boolean if a field has been set.

### SetIncrementalSnapshotUponRestartNil

`func (o *SqlEnvJobParameters) SetIncrementalSnapshotUponRestartNil(b bool)`

 SetIncrementalSnapshotUponRestartNil sets the value for IncrementalSnapshotUponRestart to be an explicit nil

### UnsetIncrementalSnapshotUponRestart
`func (o *SqlEnvJobParameters) UnsetIncrementalSnapshotUponRestart()`

UnsetIncrementalSnapshotUponRestart ensures that no value is present for IncrementalSnapshotUponRestart, not even an explicit nil
### GetIsCopyOnlyFull

`func (o *SqlEnvJobParameters) GetIsCopyOnlyFull() bool`

GetIsCopyOnlyFull returns the IsCopyOnlyFull field if non-nil, zero value otherwise.

### GetIsCopyOnlyFullOk

`func (o *SqlEnvJobParameters) GetIsCopyOnlyFullOk() (*bool, bool)`

GetIsCopyOnlyFullOk returns a tuple with the IsCopyOnlyFull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCopyOnlyFull

`func (o *SqlEnvJobParameters) SetIsCopyOnlyFull(v bool)`

SetIsCopyOnlyFull sets IsCopyOnlyFull field to given value.

### HasIsCopyOnlyFull

`func (o *SqlEnvJobParameters) HasIsCopyOnlyFull() bool`

HasIsCopyOnlyFull returns a boolean if a field has been set.

### SetIsCopyOnlyFullNil

`func (o *SqlEnvJobParameters) SetIsCopyOnlyFullNil(b bool)`

 SetIsCopyOnlyFullNil sets the value for IsCopyOnlyFull to be an explicit nil

### UnsetIsCopyOnlyFull
`func (o *SqlEnvJobParameters) UnsetIsCopyOnlyFull()`

UnsetIsCopyOnlyFull ensures that no value is present for IsCopyOnlyFull, not even an explicit nil
### GetNumStreams

`func (o *SqlEnvJobParameters) GetNumStreams() int32`

GetNumStreams returns the NumStreams field if non-nil, zero value otherwise.

### GetNumStreamsOk

`func (o *SqlEnvJobParameters) GetNumStreamsOk() (*int32, bool)`

GetNumStreamsOk returns a tuple with the NumStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStreams

`func (o *SqlEnvJobParameters) SetNumStreams(v int32)`

SetNumStreams sets NumStreams field to given value.

### HasNumStreams

`func (o *SqlEnvJobParameters) HasNumStreams() bool`

HasNumStreams returns a boolean if a field has been set.

### SetNumStreamsNil

`func (o *SqlEnvJobParameters) SetNumStreamsNil(b bool)`

 SetNumStreamsNil sets the value for NumStreams to be an explicit nil

### UnsetNumStreams
`func (o *SqlEnvJobParameters) UnsetNumStreams()`

UnsetNumStreams ensures that no value is present for NumStreams, not even an explicit nil
### GetUserDatabasePreference

`func (o *SqlEnvJobParameters) GetUserDatabasePreference() string`

GetUserDatabasePreference returns the UserDatabasePreference field if non-nil, zero value otherwise.

### GetUserDatabasePreferenceOk

`func (o *SqlEnvJobParameters) GetUserDatabasePreferenceOk() (*string, bool)`

GetUserDatabasePreferenceOk returns a tuple with the UserDatabasePreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDatabasePreference

`func (o *SqlEnvJobParameters) SetUserDatabasePreference(v string)`

SetUserDatabasePreference sets UserDatabasePreference field to given value.

### HasUserDatabasePreference

`func (o *SqlEnvJobParameters) HasUserDatabasePreference() bool`

HasUserDatabasePreference returns a boolean if a field has been set.

### SetUserDatabasePreferenceNil

`func (o *SqlEnvJobParameters) SetUserDatabasePreferenceNil(b bool)`

 SetUserDatabasePreferenceNil sets the value for UserDatabasePreference to be an explicit nil

### UnsetUserDatabasePreference
`func (o *SqlEnvJobParameters) UnsetUserDatabasePreference()`

UnsetUserDatabasePreference ensures that no value is present for UserDatabasePreference, not even an explicit nil
### GetWithClause

`func (o *SqlEnvJobParameters) GetWithClause() string`

GetWithClause returns the WithClause field if non-nil, zero value otherwise.

### GetWithClauseOk

`func (o *SqlEnvJobParameters) GetWithClauseOk() (*string, bool)`

GetWithClauseOk returns a tuple with the WithClause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithClause

`func (o *SqlEnvJobParameters) SetWithClause(v string)`

SetWithClause sets WithClause field to given value.

### HasWithClause

`func (o *SqlEnvJobParameters) HasWithClause() bool`

HasWithClause returns a boolean if a field has been set.

### SetWithClauseNil

`func (o *SqlEnvJobParameters) SetWithClauseNil(b bool)`

 SetWithClauseNil sets the value for WithClause to be an explicit nil

### UnsetWithClause
`func (o *SqlEnvJobParameters) UnsetWithClause()`

UnsetWithClause ensures that no value is present for WithClause, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


