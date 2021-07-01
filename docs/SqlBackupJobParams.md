# SqlBackupJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AagBackupPreferenceType** | Pointer to **NullableInt32** | Preference type for backing up databases that are part of an AAG. Only applicable if &#39;use_aag_preferences_from_sql_server&#39; is set to false. | [optional] 
**BackupDatabaseVolumesOnly** | Pointer to **NullableBool** | If set to true, only the volumes associated with databases should be backed up. The user cannot select additional volumes at host level for backup.  If set to false, all the volumes on the host machine will be backed up. In this case, the user can further select the exact set of volumes using host level params.  Note that the volumes associated with selected databases will always be included in the backup. | [optional] 
**BackupSystemDbs** | Pointer to **NullableBool** | Set to true if system databases should be backed up. | [optional] 
**ContinueAfterError** | Pointer to **NullableBool** | Whether backup should continue after encountering a page checksum error. | [optional] 
**EnableChecksum** | Pointer to **NullableBool** | Whether backup checksums are enabled. | [optional] 
**EnableIncrementalBackupAfterRestart** | Pointer to **NullableBool** | If this is set to true, then incremental backup will be performed after the server restarts, otherwise a full-backup will be done. | [optional] 
**FullBackupType** | Pointer to **NullableInt32** | The type of SQL full backup to be used for this job. | [optional] 
**IsCopyOnlyFull** | Pointer to **NullableBool** | Whether full backups should be copy-only. | [optional] 
**IsCopyOnlyLog** | Pointer to **NullableBool** | Whether log backups should be copy-only. | [optional] 
**NumDbsPerBatch** | Pointer to **NullableInt32** | The number of databases to be backed up per batch. This is only applicable for file based sql backup. If this is not specified, we use the value specified in magneto_vss_sql_app_file_batch_size gflag. | [optional] 
**NumStreams** | Pointer to **NullableInt32** | The number of streams to be used in native sql backup command. This is only applicable for native sql backup. If this is not specified, we use the value specified in magneto_sql_num_streams_for_each_db_backup gflag. | [optional] 
**UseAagPreferencesFromSqlServer** | Pointer to **NullableBool** | Set to true if we should use AAG preferences specified at the SQL server host. | [optional] 
**UserDbPreferenceType** | Pointer to **NullableInt32** | Preference type for backing up user databases on the host. | [optional] 
**WithClause** | Pointer to **NullableString** | &#39;with_clause&#39; contains &#39;with clause&#39; to be used in native sql backup command. This is only applicable for native sql backup. Here user can specify multiple backup options. Example: \&quot;WITH BUFFERCOUNT &#x3D; 575, MAXTRANSFERSIZE &#x3D; 2097152\&quot;. If this is not specified, we use the value specified in magneto_sql_native_backup_with_clause gflag. | [optional] 

## Methods

### NewSqlBackupJobParams

`func NewSqlBackupJobParams() *SqlBackupJobParams`

NewSqlBackupJobParams instantiates a new SqlBackupJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlBackupJobParamsWithDefaults

`func NewSqlBackupJobParamsWithDefaults() *SqlBackupJobParams`

NewSqlBackupJobParamsWithDefaults instantiates a new SqlBackupJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAagBackupPreferenceType

`func (o *SqlBackupJobParams) GetAagBackupPreferenceType() int32`

GetAagBackupPreferenceType returns the AagBackupPreferenceType field if non-nil, zero value otherwise.

### GetAagBackupPreferenceTypeOk

`func (o *SqlBackupJobParams) GetAagBackupPreferenceTypeOk() (*int32, bool)`

GetAagBackupPreferenceTypeOk returns a tuple with the AagBackupPreferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAagBackupPreferenceType

`func (o *SqlBackupJobParams) SetAagBackupPreferenceType(v int32)`

SetAagBackupPreferenceType sets AagBackupPreferenceType field to given value.

### HasAagBackupPreferenceType

`func (o *SqlBackupJobParams) HasAagBackupPreferenceType() bool`

HasAagBackupPreferenceType returns a boolean if a field has been set.

### SetAagBackupPreferenceTypeNil

`func (o *SqlBackupJobParams) SetAagBackupPreferenceTypeNil(b bool)`

 SetAagBackupPreferenceTypeNil sets the value for AagBackupPreferenceType to be an explicit nil

### UnsetAagBackupPreferenceType
`func (o *SqlBackupJobParams) UnsetAagBackupPreferenceType()`

UnsetAagBackupPreferenceType ensures that no value is present for AagBackupPreferenceType, not even an explicit nil
### GetBackupDatabaseVolumesOnly

`func (o *SqlBackupJobParams) GetBackupDatabaseVolumesOnly() bool`

GetBackupDatabaseVolumesOnly returns the BackupDatabaseVolumesOnly field if non-nil, zero value otherwise.

### GetBackupDatabaseVolumesOnlyOk

`func (o *SqlBackupJobParams) GetBackupDatabaseVolumesOnlyOk() (*bool, bool)`

GetBackupDatabaseVolumesOnlyOk returns a tuple with the BackupDatabaseVolumesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDatabaseVolumesOnly

`func (o *SqlBackupJobParams) SetBackupDatabaseVolumesOnly(v bool)`

SetBackupDatabaseVolumesOnly sets BackupDatabaseVolumesOnly field to given value.

### HasBackupDatabaseVolumesOnly

`func (o *SqlBackupJobParams) HasBackupDatabaseVolumesOnly() bool`

HasBackupDatabaseVolumesOnly returns a boolean if a field has been set.

### SetBackupDatabaseVolumesOnlyNil

`func (o *SqlBackupJobParams) SetBackupDatabaseVolumesOnlyNil(b bool)`

 SetBackupDatabaseVolumesOnlyNil sets the value for BackupDatabaseVolumesOnly to be an explicit nil

### UnsetBackupDatabaseVolumesOnly
`func (o *SqlBackupJobParams) UnsetBackupDatabaseVolumesOnly()`

UnsetBackupDatabaseVolumesOnly ensures that no value is present for BackupDatabaseVolumesOnly, not even an explicit nil
### GetBackupSystemDbs

`func (o *SqlBackupJobParams) GetBackupSystemDbs() bool`

GetBackupSystemDbs returns the BackupSystemDbs field if non-nil, zero value otherwise.

### GetBackupSystemDbsOk

`func (o *SqlBackupJobParams) GetBackupSystemDbsOk() (*bool, bool)`

GetBackupSystemDbsOk returns a tuple with the BackupSystemDbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSystemDbs

`func (o *SqlBackupJobParams) SetBackupSystemDbs(v bool)`

SetBackupSystemDbs sets BackupSystemDbs field to given value.

### HasBackupSystemDbs

`func (o *SqlBackupJobParams) HasBackupSystemDbs() bool`

HasBackupSystemDbs returns a boolean if a field has been set.

### SetBackupSystemDbsNil

`func (o *SqlBackupJobParams) SetBackupSystemDbsNil(b bool)`

 SetBackupSystemDbsNil sets the value for BackupSystemDbs to be an explicit nil

### UnsetBackupSystemDbs
`func (o *SqlBackupJobParams) UnsetBackupSystemDbs()`

UnsetBackupSystemDbs ensures that no value is present for BackupSystemDbs, not even an explicit nil
### GetContinueAfterError

`func (o *SqlBackupJobParams) GetContinueAfterError() bool`

GetContinueAfterError returns the ContinueAfterError field if non-nil, zero value otherwise.

### GetContinueAfterErrorOk

`func (o *SqlBackupJobParams) GetContinueAfterErrorOk() (*bool, bool)`

GetContinueAfterErrorOk returns a tuple with the ContinueAfterError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueAfterError

`func (o *SqlBackupJobParams) SetContinueAfterError(v bool)`

SetContinueAfterError sets ContinueAfterError field to given value.

### HasContinueAfterError

`func (o *SqlBackupJobParams) HasContinueAfterError() bool`

HasContinueAfterError returns a boolean if a field has been set.

### SetContinueAfterErrorNil

`func (o *SqlBackupJobParams) SetContinueAfterErrorNil(b bool)`

 SetContinueAfterErrorNil sets the value for ContinueAfterError to be an explicit nil

### UnsetContinueAfterError
`func (o *SqlBackupJobParams) UnsetContinueAfterError()`

UnsetContinueAfterError ensures that no value is present for ContinueAfterError, not even an explicit nil
### GetEnableChecksum

`func (o *SqlBackupJobParams) GetEnableChecksum() bool`

GetEnableChecksum returns the EnableChecksum field if non-nil, zero value otherwise.

### GetEnableChecksumOk

`func (o *SqlBackupJobParams) GetEnableChecksumOk() (*bool, bool)`

GetEnableChecksumOk returns a tuple with the EnableChecksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableChecksum

`func (o *SqlBackupJobParams) SetEnableChecksum(v bool)`

SetEnableChecksum sets EnableChecksum field to given value.

### HasEnableChecksum

`func (o *SqlBackupJobParams) HasEnableChecksum() bool`

HasEnableChecksum returns a boolean if a field has been set.

### SetEnableChecksumNil

`func (o *SqlBackupJobParams) SetEnableChecksumNil(b bool)`

 SetEnableChecksumNil sets the value for EnableChecksum to be an explicit nil

### UnsetEnableChecksum
`func (o *SqlBackupJobParams) UnsetEnableChecksum()`

UnsetEnableChecksum ensures that no value is present for EnableChecksum, not even an explicit nil
### GetEnableIncrementalBackupAfterRestart

`func (o *SqlBackupJobParams) GetEnableIncrementalBackupAfterRestart() bool`

GetEnableIncrementalBackupAfterRestart returns the EnableIncrementalBackupAfterRestart field if non-nil, zero value otherwise.

### GetEnableIncrementalBackupAfterRestartOk

`func (o *SqlBackupJobParams) GetEnableIncrementalBackupAfterRestartOk() (*bool, bool)`

GetEnableIncrementalBackupAfterRestartOk returns a tuple with the EnableIncrementalBackupAfterRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIncrementalBackupAfterRestart

`func (o *SqlBackupJobParams) SetEnableIncrementalBackupAfterRestart(v bool)`

SetEnableIncrementalBackupAfterRestart sets EnableIncrementalBackupAfterRestart field to given value.

### HasEnableIncrementalBackupAfterRestart

`func (o *SqlBackupJobParams) HasEnableIncrementalBackupAfterRestart() bool`

HasEnableIncrementalBackupAfterRestart returns a boolean if a field has been set.

### SetEnableIncrementalBackupAfterRestartNil

`func (o *SqlBackupJobParams) SetEnableIncrementalBackupAfterRestartNil(b bool)`

 SetEnableIncrementalBackupAfterRestartNil sets the value for EnableIncrementalBackupAfterRestart to be an explicit nil

### UnsetEnableIncrementalBackupAfterRestart
`func (o *SqlBackupJobParams) UnsetEnableIncrementalBackupAfterRestart()`

UnsetEnableIncrementalBackupAfterRestart ensures that no value is present for EnableIncrementalBackupAfterRestart, not even an explicit nil
### GetFullBackupType

`func (o *SqlBackupJobParams) GetFullBackupType() int32`

GetFullBackupType returns the FullBackupType field if non-nil, zero value otherwise.

### GetFullBackupTypeOk

`func (o *SqlBackupJobParams) GetFullBackupTypeOk() (*int32, bool)`

GetFullBackupTypeOk returns a tuple with the FullBackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullBackupType

`func (o *SqlBackupJobParams) SetFullBackupType(v int32)`

SetFullBackupType sets FullBackupType field to given value.

### HasFullBackupType

`func (o *SqlBackupJobParams) HasFullBackupType() bool`

HasFullBackupType returns a boolean if a field has been set.

### SetFullBackupTypeNil

`func (o *SqlBackupJobParams) SetFullBackupTypeNil(b bool)`

 SetFullBackupTypeNil sets the value for FullBackupType to be an explicit nil

### UnsetFullBackupType
`func (o *SqlBackupJobParams) UnsetFullBackupType()`

UnsetFullBackupType ensures that no value is present for FullBackupType, not even an explicit nil
### GetIsCopyOnlyFull

`func (o *SqlBackupJobParams) GetIsCopyOnlyFull() bool`

GetIsCopyOnlyFull returns the IsCopyOnlyFull field if non-nil, zero value otherwise.

### GetIsCopyOnlyFullOk

`func (o *SqlBackupJobParams) GetIsCopyOnlyFullOk() (*bool, bool)`

GetIsCopyOnlyFullOk returns a tuple with the IsCopyOnlyFull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCopyOnlyFull

`func (o *SqlBackupJobParams) SetIsCopyOnlyFull(v bool)`

SetIsCopyOnlyFull sets IsCopyOnlyFull field to given value.

### HasIsCopyOnlyFull

`func (o *SqlBackupJobParams) HasIsCopyOnlyFull() bool`

HasIsCopyOnlyFull returns a boolean if a field has been set.

### SetIsCopyOnlyFullNil

`func (o *SqlBackupJobParams) SetIsCopyOnlyFullNil(b bool)`

 SetIsCopyOnlyFullNil sets the value for IsCopyOnlyFull to be an explicit nil

### UnsetIsCopyOnlyFull
`func (o *SqlBackupJobParams) UnsetIsCopyOnlyFull()`

UnsetIsCopyOnlyFull ensures that no value is present for IsCopyOnlyFull, not even an explicit nil
### GetIsCopyOnlyLog

`func (o *SqlBackupJobParams) GetIsCopyOnlyLog() bool`

GetIsCopyOnlyLog returns the IsCopyOnlyLog field if non-nil, zero value otherwise.

### GetIsCopyOnlyLogOk

`func (o *SqlBackupJobParams) GetIsCopyOnlyLogOk() (*bool, bool)`

GetIsCopyOnlyLogOk returns a tuple with the IsCopyOnlyLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCopyOnlyLog

`func (o *SqlBackupJobParams) SetIsCopyOnlyLog(v bool)`

SetIsCopyOnlyLog sets IsCopyOnlyLog field to given value.

### HasIsCopyOnlyLog

`func (o *SqlBackupJobParams) HasIsCopyOnlyLog() bool`

HasIsCopyOnlyLog returns a boolean if a field has been set.

### SetIsCopyOnlyLogNil

`func (o *SqlBackupJobParams) SetIsCopyOnlyLogNil(b bool)`

 SetIsCopyOnlyLogNil sets the value for IsCopyOnlyLog to be an explicit nil

### UnsetIsCopyOnlyLog
`func (o *SqlBackupJobParams) UnsetIsCopyOnlyLog()`

UnsetIsCopyOnlyLog ensures that no value is present for IsCopyOnlyLog, not even an explicit nil
### GetNumDbsPerBatch

`func (o *SqlBackupJobParams) GetNumDbsPerBatch() int32`

GetNumDbsPerBatch returns the NumDbsPerBatch field if non-nil, zero value otherwise.

### GetNumDbsPerBatchOk

`func (o *SqlBackupJobParams) GetNumDbsPerBatchOk() (*int32, bool)`

GetNumDbsPerBatchOk returns a tuple with the NumDbsPerBatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDbsPerBatch

`func (o *SqlBackupJobParams) SetNumDbsPerBatch(v int32)`

SetNumDbsPerBatch sets NumDbsPerBatch field to given value.

### HasNumDbsPerBatch

`func (o *SqlBackupJobParams) HasNumDbsPerBatch() bool`

HasNumDbsPerBatch returns a boolean if a field has been set.

### SetNumDbsPerBatchNil

`func (o *SqlBackupJobParams) SetNumDbsPerBatchNil(b bool)`

 SetNumDbsPerBatchNil sets the value for NumDbsPerBatch to be an explicit nil

### UnsetNumDbsPerBatch
`func (o *SqlBackupJobParams) UnsetNumDbsPerBatch()`

UnsetNumDbsPerBatch ensures that no value is present for NumDbsPerBatch, not even an explicit nil
### GetNumStreams

`func (o *SqlBackupJobParams) GetNumStreams() int32`

GetNumStreams returns the NumStreams field if non-nil, zero value otherwise.

### GetNumStreamsOk

`func (o *SqlBackupJobParams) GetNumStreamsOk() (*int32, bool)`

GetNumStreamsOk returns a tuple with the NumStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStreams

`func (o *SqlBackupJobParams) SetNumStreams(v int32)`

SetNumStreams sets NumStreams field to given value.

### HasNumStreams

`func (o *SqlBackupJobParams) HasNumStreams() bool`

HasNumStreams returns a boolean if a field has been set.

### SetNumStreamsNil

`func (o *SqlBackupJobParams) SetNumStreamsNil(b bool)`

 SetNumStreamsNil sets the value for NumStreams to be an explicit nil

### UnsetNumStreams
`func (o *SqlBackupJobParams) UnsetNumStreams()`

UnsetNumStreams ensures that no value is present for NumStreams, not even an explicit nil
### GetUseAagPreferencesFromSqlServer

`func (o *SqlBackupJobParams) GetUseAagPreferencesFromSqlServer() bool`

GetUseAagPreferencesFromSqlServer returns the UseAagPreferencesFromSqlServer field if non-nil, zero value otherwise.

### GetUseAagPreferencesFromSqlServerOk

`func (o *SqlBackupJobParams) GetUseAagPreferencesFromSqlServerOk() (*bool, bool)`

GetUseAagPreferencesFromSqlServerOk returns a tuple with the UseAagPreferencesFromSqlServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAagPreferencesFromSqlServer

`func (o *SqlBackupJobParams) SetUseAagPreferencesFromSqlServer(v bool)`

SetUseAagPreferencesFromSqlServer sets UseAagPreferencesFromSqlServer field to given value.

### HasUseAagPreferencesFromSqlServer

`func (o *SqlBackupJobParams) HasUseAagPreferencesFromSqlServer() bool`

HasUseAagPreferencesFromSqlServer returns a boolean if a field has been set.

### SetUseAagPreferencesFromSqlServerNil

`func (o *SqlBackupJobParams) SetUseAagPreferencesFromSqlServerNil(b bool)`

 SetUseAagPreferencesFromSqlServerNil sets the value for UseAagPreferencesFromSqlServer to be an explicit nil

### UnsetUseAagPreferencesFromSqlServer
`func (o *SqlBackupJobParams) UnsetUseAagPreferencesFromSqlServer()`

UnsetUseAagPreferencesFromSqlServer ensures that no value is present for UseAagPreferencesFromSqlServer, not even an explicit nil
### GetUserDbPreferenceType

`func (o *SqlBackupJobParams) GetUserDbPreferenceType() int32`

GetUserDbPreferenceType returns the UserDbPreferenceType field if non-nil, zero value otherwise.

### GetUserDbPreferenceTypeOk

`func (o *SqlBackupJobParams) GetUserDbPreferenceTypeOk() (*int32, bool)`

GetUserDbPreferenceTypeOk returns a tuple with the UserDbPreferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDbPreferenceType

`func (o *SqlBackupJobParams) SetUserDbPreferenceType(v int32)`

SetUserDbPreferenceType sets UserDbPreferenceType field to given value.

### HasUserDbPreferenceType

`func (o *SqlBackupJobParams) HasUserDbPreferenceType() bool`

HasUserDbPreferenceType returns a boolean if a field has been set.

### SetUserDbPreferenceTypeNil

`func (o *SqlBackupJobParams) SetUserDbPreferenceTypeNil(b bool)`

 SetUserDbPreferenceTypeNil sets the value for UserDbPreferenceType to be an explicit nil

### UnsetUserDbPreferenceType
`func (o *SqlBackupJobParams) UnsetUserDbPreferenceType()`

UnsetUserDbPreferenceType ensures that no value is present for UserDbPreferenceType, not even an explicit nil
### GetWithClause

`func (o *SqlBackupJobParams) GetWithClause() string`

GetWithClause returns the WithClause field if non-nil, zero value otherwise.

### GetWithClauseOk

`func (o *SqlBackupJobParams) GetWithClauseOk() (*string, bool)`

GetWithClauseOk returns a tuple with the WithClause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithClause

`func (o *SqlBackupJobParams) SetWithClause(v string)`

SetWithClause sets WithClause field to given value.

### HasWithClause

`func (o *SqlBackupJobParams) HasWithClause() bool`

HasWithClause returns a boolean if a field has been set.

### SetWithClauseNil

`func (o *SqlBackupJobParams) SetWithClauseNil(b bool)`

 SetWithClauseNil sets the value for WithClause to be an explicit nil

### UnsetWithClause
`func (o *SqlBackupJobParams) UnsetWithClause()`

UnsetWithClause ensures that no value is present for WithClause, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


