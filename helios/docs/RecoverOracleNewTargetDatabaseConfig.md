# RecoverOracleNewTargetDatabaseConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseName** | Pointer to **NullableString** | Specifies a new name for the restored database. If this field is not specified, then the original database will be overwritten after recovery. | [optional] 
**OracleBaseFolder** | Pointer to **NullableString** | Specifies the oracle base folder at selected host. | [optional] 
**OracleHomeFolder** | Pointer to **NullableString** | Specifies the oracle home folder at selected host. | [optional] 
**DbFilesDestination** | Pointer to **NullableString** | Specifies the location to restore database files. | [optional] 
**DbConfigFilePath** | Pointer to **NullableString** | Specifies the config file path on selected host which configures the restored database. | [optional] 
**EnableArchiveLogMode** | Pointer to **NullableBool** | Specifies archive log mode for oracle restore. | [optional] 
**PfileParameterMap** | Pointer to [**[]KeyValuePair**](KeyValuePair.md) | Specifies a key value pair for pfile parameters. | [optional] 
**BctFilePath** | Pointer to **NullableString** | Specifies BCT file path. | [optional] 
**NumTempfiles** | Pointer to **NullableInt32** | Specifies no. of tempfiles to be used for the recovered database. | [optional] 
**RedoLogConfig** | Pointer to [**RedoLogGroupConfig**](RedoLogGroupConfig.md) | Specifies redo log config. | [optional] 
**IsMultiStageRestore** | Pointer to **NullableBool** | Specifies whether this task is a multistage restore task. If set, we migrate the DB after clone completes. | [optional] 
**OracleUpdateRestoreOptions** | Pointer to [**NullableMigrateCloneParams**](MigrateCloneParams.md) | Specifies the parameters that are needed for updating oracle restore options. | [optional] 
**SkipCloneNid** | Pointer to **NullableBool** | Whether or not to skip the nid step in Oracle Clone workflow. Applicable to both smart and old clone workflow. | [optional] 
**NoFilenameCheck** | Pointer to **NullableBool** | Specifies whether to validate filenames or not in Oracle alternate restore workflow. | [optional] 
**NewNameClause** | Pointer to **NullableString** | Specifies newname clause for db files which allows user to have full control on how their database files can be renamed during the oracle alternate restore workflow. | [optional] 
**RestoreTimeUsecs** | Pointer to **NullableInt64** | Specifies the time in the past to which the Oracle db needs to be restored. This allows for granular recovery of Oracle databases. If this is not set, the Oracle db will be restored from the full/incremental snapshot. | [optional] 
**DbChannels** | Pointer to [**[]OracleDbChannel**](OracleDbChannel.md) | Specifies the Oracle database node channels info. If not specified, the default values assigned by the server are applied to all the databases. | [optional] 
**RecoveryMode** | Pointer to **NullableBool** | Specifies if database should be left in recovery mode. | [optional] 
**ShellEvironmentVars** | Pointer to [**[]ShellKeyValuePair**](ShellKeyValuePair.md) | Specifies key value pairs of shell variables which defines the restore shell environment. | [optional] 
**GranularRestoreInfo** | Pointer to [**RecoverOracleGranularRestoreInfo**](RecoverOracleGranularRestoreInfo.md) | Specifies information about list of objects (PDBs) to restore. | [optional] 
**OracleArchiveLogInfo** | Pointer to [**OracleArchiveLogInfo**](OracleArchiveLogInfo.md) | Specifies Range in Time, Scn or Sequence to restore archive logs of a DB. | [optional] 

## Methods

### NewRecoverOracleNewTargetDatabaseConfig

`func NewRecoverOracleNewTargetDatabaseConfig() *RecoverOracleNewTargetDatabaseConfig`

NewRecoverOracleNewTargetDatabaseConfig instantiates a new RecoverOracleNewTargetDatabaseConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverOracleNewTargetDatabaseConfigWithDefaults

`func NewRecoverOracleNewTargetDatabaseConfigWithDefaults() *RecoverOracleNewTargetDatabaseConfig`

NewRecoverOracleNewTargetDatabaseConfigWithDefaults instantiates a new RecoverOracleNewTargetDatabaseConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseName

`func (o *RecoverOracleNewTargetDatabaseConfig) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *RecoverOracleNewTargetDatabaseConfig) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *RecoverOracleNewTargetDatabaseConfig) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *RecoverOracleNewTargetDatabaseConfig) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### SetDatabaseNameNil

`func (o *RecoverOracleNewTargetDatabaseConfig) SetDatabaseNameNil(b bool)`

 SetDatabaseNameNil sets the value for DatabaseName to be an explicit nil

### UnsetDatabaseName
`func (o *RecoverOracleNewTargetDatabaseConfig) UnsetDatabaseName()`

UnsetDatabaseName ensures that no value is present for DatabaseName, not even an explicit nil
### GetOracleBaseFolder

`func (o *RecoverOracleNewTargetDatabaseConfig) GetOracleBaseFolder() string`

GetOracleBaseFolder returns the OracleBaseFolder field if non-nil, zero value otherwise.

### GetOracleBaseFolderOk

`func (o *RecoverOracleNewTargetDatabaseConfig) GetOracleBaseFolderOk() (*string, bool)`

GetOracleBaseFolderOk returns a tuple with the OracleBaseFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleBaseFolder

`func (o *RecoverOracleNewTargetDatabaseConfig) SetOracleBaseFolder(v string)`

SetOracleBaseFolder sets OracleBaseFolder field to given value.

### HasOracleBaseFolder

`func (o *RecoverOracleNewTargetDatabaseConfig) HasOracleBaseFolder() bool`

HasOracleBaseFolder returns a boolean if a field has been set.

### SetOracleBaseFolderNil

`func (o *RecoverOracleNewTargetDatabaseConfig) SetOracleBaseFolderNil(b bool)`

 SetOracleBaseFolderNil sets the value for OracleBaseFolder to be an explicit nil

### UnsetOracleBaseFolder
`func (o *RecoverOracleNewTargetDatabaseConfig) UnsetOracleBaseFolder()`

UnsetOracleBaseFolder ensures that no value is present for OracleBaseFolder, not even an explicit nil
### GetOracleHomeFolder

`func (o *RecoverOracleNewTargetDatabaseConfig) GetOracleHomeFolder() string`

GetOracleHomeFolder returns the OracleHomeFolder field if non-nil, zero value otherwise.

### GetOracleHomeFolderOk

`func (o *RecoverOracleNewTargetDatabaseConfig) GetOracleHomeFolderOk() (*string, bool)`

GetOracleHomeFolderOk returns a tuple with the OracleHomeFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleHomeFolder

`func (o *RecoverOracleNewTargetDatabaseConfig) SetOracleHomeFolder(v string)`

SetOracleHomeFolder sets OracleHomeFolder field to given value.

### HasOracleHomeFolder

`func (o *RecoverOracleNewTargetDatabaseConfig) HasOracleHomeFolder() bool`

HasOracleHomeFolder returns a boolean if a field has been set.

### SetOracleHomeFolderNil

`func (o *RecoverOracleNewTargetDatabaseConfig) SetOracleHomeFolderNil(b bool)`

 SetOracleHomeFolderNil sets the value for OracleHomeFolder to be an explicit nil

### UnsetOracleHomeFolder
`func (o *RecoverOracleNewTargetDatabaseConfig) UnsetOracleHomeFolder()`

UnsetOracleHomeFolder ensures that no value is present for OracleHomeFolder, not even an explicit nil
### GetDbFilesDestination

`func (o *RecoverOracleNewTargetDatabaseConfig) GetDbFilesDestination() string`

GetDbFilesDestination returns the DbFilesDestination field if non-nil, zero value otherwise.

### GetDbFilesDestinationOk

`func (o *RecoverOracleNewTargetDatabaseConfig) GetDbFilesDestinationOk() (*string, bool)`

GetDbFilesDestinationOk returns a tuple with the DbFilesDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbFilesDestination

`func (o *RecoverOracleNewTargetDatabaseConfig) SetDbFilesDestination(v string)`

SetDbFilesDestination sets DbFilesDestination field to given value.

### HasDbFilesDestination

`func (o *RecoverOracleNewTargetDatabaseConfig) HasDbFilesDestination() bool`

HasDbFilesDestination returns a boolean if a field has been set.

### SetDbFilesDestinationNil

`func (o *RecoverOracleNewTargetDatabaseConfig) SetDbFilesDestinationNil(b bool)`

 SetDbFilesDestinationNil sets the value for DbFilesDestination to be an explicit nil

### UnsetDbFilesDestination
`func (o *RecoverOracleNewTargetDatabaseConfig) UnsetDbFilesDestination()`

UnsetDbFilesDestination ensures that no value is present for DbFilesDestination, not even an explicit nil
### GetDbConfigFilePath

`func (o *RecoverOracleNewTargetDatabaseConfig) GetDbConfigFilePath() string`

GetDbConfigFilePath returns the DbConfigFilePath field if non-nil, zero value otherwise.

### GetDbConfigFilePathOk

`func (o *RecoverOracleNewTargetDatabaseConfig) GetDbConfigFilePathOk() (*string, bool)`

GetDbConfigFilePathOk returns a tuple with the DbConfigFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbConfigFilePath

`func (o *RecoverOracleNewTargetDatabaseConfig) SetDbConfigFilePath(v string)`

SetDbConfigFilePath sets DbConfigFilePath field to given value.

### HasDbConfigFilePath

`func (o *RecoverOracleNewTargetDatabaseConfig) HasDbConfigFilePath() bool`

HasDbConfigFilePath returns a boolean if a field has been set.

### SetDbConfigFilePathNil

`func (o *RecoverOracleNewTargetDatabaseConfig) SetDbConfigFilePathNil(b bool)`

 SetDbConfigFilePathNil sets the value for DbConfigFilePath to be an explicit nil

### UnsetDbConfigFilePath
`func (o *RecoverOracleNewTargetDatabaseConfig) UnsetDbConfigFilePath()`

UnsetDbConfigFilePath ensures that no value is present for DbConfigFilePath, not even an explicit nil
### GetEnableArchiveLogMode

`func (o *RecoverOracleNewTargetDatabaseConfig) GetEnableArchiveLogMode() bool`

GetEnableArchiveLogMode returns the EnableArchiveLogMode field if non-nil, zero value otherwise.

### GetEnableArchiveLogModeOk

`func (o *RecoverOracleNewTargetDatabaseConfig) GetEnableArchiveLogModeOk() (*bool, bool)`

GetEnableArchiveLogModeOk returns a tuple with the EnableArchiveLogMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableArchiveLogMode

`func (o *RecoverOracleNewTargetDatabaseConfig) SetEnableArchiveLogMode(v bool)`

SetEnableArchiveLogMode sets EnableArchiveLogMode field to given value.

### HasEnableArchiveLogMode

`func (o *RecoverOracleNewTargetDatabaseConfig) HasEnableArchiveLogMode() bool`

HasEnableArchiveLogMode returns a boolean if a field has been set.

### SetEnableArchiveLogModeNil

`func (o *RecoverOracleNewTargetDatabaseConfig) SetEnableArchiveLogModeNil(b bool)`

 SetEnableArchiveLogModeNil sets the value for EnableArchiveLogMode to be an explicit nil

### UnsetEnableArchiveLogMode
`func (o *RecoverOracleNewTargetDatabaseConfig) UnsetEnableArchiveLogMode()`

UnsetEnableArchiveLogMode ensures that no value is present for EnableArchiveLogMode, not even an explicit nil
### GetPfileParameterMap

`func (o *RecoverOracleNewTargetDatabaseConfig) GetPfileParameterMap() []KeyValuePair`

GetPfileParameterMap returns the PfileParameterMap field if non-nil, zero value otherwise.

### GetPfileParameterMapOk

`func (o *RecoverOracleNewTargetDatabaseConfig) GetPfileParameterMapOk() (*[]KeyValuePair, bool)`

GetPfileParameterMapOk returns a tuple with the PfileParameterMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfileParameterMap

`func (o *RecoverOracleNewTargetDatabaseConfig) SetPfileParameterMap(v []KeyValuePair)`

SetPfileParameterMap sets PfileParameterMap field to given value.

### HasPfileParameterMap

`func (o *RecoverOracleNewTargetDatabaseConfig) HasPfileParameterMap() bool`

HasPfileParameterMap returns a boolean if a field has been set.

### SetPfileParameterMapNil

`func (o *RecoverOracleNewTargetDatabaseConfig) SetPfileParameterMapNil(b bool)`

 SetPfileParameterMapNil sets the value for PfileParameterMap to be an explicit nil

### UnsetPfileParameterMap
`func (o *RecoverOracleNewTargetDatabaseConfig) UnsetPfileParameterMap()`

UnsetPfileParameterMap ensures that no value is present for PfileParameterMap, not even an explicit nil
### GetBctFilePath

`func (o *RecoverOracleNewTargetDatabaseConfig) GetBctFilePath() string`

GetBctFilePath returns the BctFilePath field if non-nil, zero value otherwise.

### GetBctFilePathOk

`func (o *RecoverOracleNewTargetDatabaseConfig) GetBctFilePathOk() (*string, bool)`

GetBctFilePathOk returns a tuple with the BctFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBctFilePath

`func (o *RecoverOracleNewTargetDatabaseConfig) SetBctFilePath(v string)`

SetBctFilePath sets BctFilePath field to given value.

### HasBctFilePath

`func (o *RecoverOracleNewTargetDatabaseConfig) HasBctFilePath() bool`

HasBctFilePath returns a boolean if a field has been set.

### SetBctFilePathNil

`func (o *RecoverOracleNewTargetDatabaseConfig) SetBctFilePathNil(b bool)`

 SetBctFilePathNil sets the value for BctFilePath to be an explicit nil

### UnsetBctFilePath
`func (o *RecoverOracleNewTargetDatabaseConfig) UnsetBctFilePath()`

UnsetBctFilePath ensures that no value is present for BctFilePath, not even an explicit nil
### GetNumTempfiles

`func (o *RecoverOracleNewTargetDatabaseConfig) GetNumTempfiles() int32`

GetNumTempfiles returns the NumTempfiles field if non-nil, zero value otherwise.

### GetNumTempfilesOk

`func (o *RecoverOracleNewTargetDatabaseConfig) GetNumTempfilesOk() (*int32, bool)`

GetNumTempfilesOk returns a tuple with the NumTempfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTempfiles

`func (o *RecoverOracleNewTargetDatabaseConfig) SetNumTempfiles(v int32)`

SetNumTempfiles sets NumTempfiles field to given value.

### HasNumTempfiles

`func (o *RecoverOracleNewTargetDatabaseConfig) HasNumTempfiles() bool`

HasNumTempfiles returns a boolean if a field has been set.

### SetNumTempfilesNil

`func (o *RecoverOracleNewTargetDatabaseConfig) SetNumTempfilesNil(b bool)`

 SetNumTempfilesNil sets the value for NumTempfiles to be an explicit nil

### UnsetNumTempfiles
`func (o *RecoverOracleNewTargetDatabaseConfig) UnsetNumTempfiles()`

UnsetNumTempfiles ensures that no value is present for NumTempfiles, not even an explicit nil
### GetRedoLogConfig

`func (o *RecoverOracleNewTargetDatabaseConfig) GetRedoLogConfig() RedoLogGroupConfig`

GetRedoLogConfig returns the RedoLogConfig field if non-nil, zero value otherwise.

### GetRedoLogConfigOk

`func (o *RecoverOracleNewTargetDatabaseConfig) GetRedoLogConfigOk() (*RedoLogGroupConfig, bool)`

GetRedoLogConfigOk returns a tuple with the RedoLogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedoLogConfig

`func (o *RecoverOracleNewTargetDatabaseConfig) SetRedoLogConfig(v RedoLogGroupConfig)`

SetRedoLogConfig sets RedoLogConfig field to given value.

### HasRedoLogConfig

`func (o *RecoverOracleNewTargetDatabaseConfig) HasRedoLogConfig() bool`

HasRedoLogConfig returns a boolean if a field has been set.

### GetIsMultiStageRestore

`func (o *RecoverOracleNewTargetDatabaseConfig) GetIsMultiStageRestore() bool`

GetIsMultiStageRestore returns the IsMultiStageRestore field if non-nil, zero value otherwise.

### GetIsMultiStageRestoreOk

`func (o *RecoverOracleNewTargetDatabaseConfig) GetIsMultiStageRestoreOk() (*bool, bool)`

GetIsMultiStageRestoreOk returns a tuple with the IsMultiStageRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiStageRestore

`func (o *RecoverOracleNewTargetDatabaseConfig) SetIsMultiStageRestore(v bool)`

SetIsMultiStageRestore sets IsMultiStageRestore field to given value.

### HasIsMultiStageRestore

`func (o *RecoverOracleNewTargetDatabaseConfig) HasIsMultiStageRestore() bool`

HasIsMultiStageRestore returns a boolean if a field has been set.

### SetIsMultiStageRestoreNil

`func (o *RecoverOracleNewTargetDatabaseConfig) SetIsMultiStageRestoreNil(b bool)`

 SetIsMultiStageRestoreNil sets the value for IsMultiStageRestore to be an explicit nil

### UnsetIsMultiStageRestore
`func (o *RecoverOracleNewTargetDatabaseConfig) UnsetIsMultiStageRestore()`

UnsetIsMultiStageRestore ensures that no value is present for IsMultiStageRestore, not even an explicit nil
### GetOracleUpdateRestoreOptions

`func (o *RecoverOracleNewTargetDatabaseConfig) GetOracleUpdateRestoreOptions() MigrateCloneParams`

GetOracleUpdateRestoreOptions returns the OracleUpdateRestoreOptions field if non-nil, zero value otherwise.

### GetOracleUpdateRestoreOptionsOk

`func (o *RecoverOracleNewTargetDatabaseConfig) GetOracleUpdateRestoreOptionsOk() (*MigrateCloneParams, bool)`

GetOracleUpdateRestoreOptionsOk returns a tuple with the OracleUpdateRestoreOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleUpdateRestoreOptions

`func (o *RecoverOracleNewTargetDatabaseConfig) SetOracleUpdateRestoreOptions(v MigrateCloneParams)`

SetOracleUpdateRestoreOptions sets OracleUpdateRestoreOptions field to given value.

### HasOracleUpdateRestoreOptions

`func (o *RecoverOracleNewTargetDatabaseConfig) HasOracleUpdateRestoreOptions() bool`

HasOracleUpdateRestoreOptions returns a boolean if a field has been set.

### SetOracleUpdateRestoreOptionsNil

`func (o *RecoverOracleNewTargetDatabaseConfig) SetOracleUpdateRestoreOptionsNil(b bool)`

 SetOracleUpdateRestoreOptionsNil sets the value for OracleUpdateRestoreOptions to be an explicit nil

### UnsetOracleUpdateRestoreOptions
`func (o *RecoverOracleNewTargetDatabaseConfig) UnsetOracleUpdateRestoreOptions()`

UnsetOracleUpdateRestoreOptions ensures that no value is present for OracleUpdateRestoreOptions, not even an explicit nil
### GetSkipCloneNid

`func (o *RecoverOracleNewTargetDatabaseConfig) GetSkipCloneNid() bool`

GetSkipCloneNid returns the SkipCloneNid field if non-nil, zero value otherwise.

### GetSkipCloneNidOk

`func (o *RecoverOracleNewTargetDatabaseConfig) GetSkipCloneNidOk() (*bool, bool)`

GetSkipCloneNidOk returns a tuple with the SkipCloneNid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCloneNid

`func (o *RecoverOracleNewTargetDatabaseConfig) SetSkipCloneNid(v bool)`

SetSkipCloneNid sets SkipCloneNid field to given value.

### HasSkipCloneNid

`func (o *RecoverOracleNewTargetDatabaseConfig) HasSkipCloneNid() bool`

HasSkipCloneNid returns a boolean if a field has been set.

### SetSkipCloneNidNil

`func (o *RecoverOracleNewTargetDatabaseConfig) SetSkipCloneNidNil(b bool)`

 SetSkipCloneNidNil sets the value for SkipCloneNid to be an explicit nil

### UnsetSkipCloneNid
`func (o *RecoverOracleNewTargetDatabaseConfig) UnsetSkipCloneNid()`

UnsetSkipCloneNid ensures that no value is present for SkipCloneNid, not even an explicit nil
### GetNoFilenameCheck

`func (o *RecoverOracleNewTargetDatabaseConfig) GetNoFilenameCheck() bool`

GetNoFilenameCheck returns the NoFilenameCheck field if non-nil, zero value otherwise.

### GetNoFilenameCheckOk

`func (o *RecoverOracleNewTargetDatabaseConfig) GetNoFilenameCheckOk() (*bool, bool)`

GetNoFilenameCheckOk returns a tuple with the NoFilenameCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoFilenameCheck

`func (o *RecoverOracleNewTargetDatabaseConfig) SetNoFilenameCheck(v bool)`

SetNoFilenameCheck sets NoFilenameCheck field to given value.

### HasNoFilenameCheck

`func (o *RecoverOracleNewTargetDatabaseConfig) HasNoFilenameCheck() bool`

HasNoFilenameCheck returns a boolean if a field has been set.

### SetNoFilenameCheckNil

`func (o *RecoverOracleNewTargetDatabaseConfig) SetNoFilenameCheckNil(b bool)`

 SetNoFilenameCheckNil sets the value for NoFilenameCheck to be an explicit nil

### UnsetNoFilenameCheck
`func (o *RecoverOracleNewTargetDatabaseConfig) UnsetNoFilenameCheck()`

UnsetNoFilenameCheck ensures that no value is present for NoFilenameCheck, not even an explicit nil
### GetNewNameClause

`func (o *RecoverOracleNewTargetDatabaseConfig) GetNewNameClause() string`

GetNewNameClause returns the NewNameClause field if non-nil, zero value otherwise.

### GetNewNameClauseOk

`func (o *RecoverOracleNewTargetDatabaseConfig) GetNewNameClauseOk() (*string, bool)`

GetNewNameClauseOk returns a tuple with the NewNameClause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewNameClause

`func (o *RecoverOracleNewTargetDatabaseConfig) SetNewNameClause(v string)`

SetNewNameClause sets NewNameClause field to given value.

### HasNewNameClause

`func (o *RecoverOracleNewTargetDatabaseConfig) HasNewNameClause() bool`

HasNewNameClause returns a boolean if a field has been set.

### SetNewNameClauseNil

`func (o *RecoverOracleNewTargetDatabaseConfig) SetNewNameClauseNil(b bool)`

 SetNewNameClauseNil sets the value for NewNameClause to be an explicit nil

### UnsetNewNameClause
`func (o *RecoverOracleNewTargetDatabaseConfig) UnsetNewNameClause()`

UnsetNewNameClause ensures that no value is present for NewNameClause, not even an explicit nil
### GetRestoreTimeUsecs

`func (o *RecoverOracleNewTargetDatabaseConfig) GetRestoreTimeUsecs() int64`

GetRestoreTimeUsecs returns the RestoreTimeUsecs field if non-nil, zero value otherwise.

### GetRestoreTimeUsecsOk

`func (o *RecoverOracleNewTargetDatabaseConfig) GetRestoreTimeUsecsOk() (*int64, bool)`

GetRestoreTimeUsecsOk returns a tuple with the RestoreTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTimeUsecs

`func (o *RecoverOracleNewTargetDatabaseConfig) SetRestoreTimeUsecs(v int64)`

SetRestoreTimeUsecs sets RestoreTimeUsecs field to given value.

### HasRestoreTimeUsecs

`func (o *RecoverOracleNewTargetDatabaseConfig) HasRestoreTimeUsecs() bool`

HasRestoreTimeUsecs returns a boolean if a field has been set.

### SetRestoreTimeUsecsNil

`func (o *RecoverOracleNewTargetDatabaseConfig) SetRestoreTimeUsecsNil(b bool)`

 SetRestoreTimeUsecsNil sets the value for RestoreTimeUsecs to be an explicit nil

### UnsetRestoreTimeUsecs
`func (o *RecoverOracleNewTargetDatabaseConfig) UnsetRestoreTimeUsecs()`

UnsetRestoreTimeUsecs ensures that no value is present for RestoreTimeUsecs, not even an explicit nil
### GetDbChannels

`func (o *RecoverOracleNewTargetDatabaseConfig) GetDbChannels() []OracleDbChannel`

GetDbChannels returns the DbChannels field if non-nil, zero value otherwise.

### GetDbChannelsOk

`func (o *RecoverOracleNewTargetDatabaseConfig) GetDbChannelsOk() (*[]OracleDbChannel, bool)`

GetDbChannelsOk returns a tuple with the DbChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbChannels

`func (o *RecoverOracleNewTargetDatabaseConfig) SetDbChannels(v []OracleDbChannel)`

SetDbChannels sets DbChannels field to given value.

### HasDbChannels

`func (o *RecoverOracleNewTargetDatabaseConfig) HasDbChannels() bool`

HasDbChannels returns a boolean if a field has been set.

### SetDbChannelsNil

`func (o *RecoverOracleNewTargetDatabaseConfig) SetDbChannelsNil(b bool)`

 SetDbChannelsNil sets the value for DbChannels to be an explicit nil

### UnsetDbChannels
`func (o *RecoverOracleNewTargetDatabaseConfig) UnsetDbChannels()`

UnsetDbChannels ensures that no value is present for DbChannels, not even an explicit nil
### GetRecoveryMode

`func (o *RecoverOracleNewTargetDatabaseConfig) GetRecoveryMode() bool`

GetRecoveryMode returns the RecoveryMode field if non-nil, zero value otherwise.

### GetRecoveryModeOk

`func (o *RecoverOracleNewTargetDatabaseConfig) GetRecoveryModeOk() (*bool, bool)`

GetRecoveryModeOk returns a tuple with the RecoveryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryMode

`func (o *RecoverOracleNewTargetDatabaseConfig) SetRecoveryMode(v bool)`

SetRecoveryMode sets RecoveryMode field to given value.

### HasRecoveryMode

`func (o *RecoverOracleNewTargetDatabaseConfig) HasRecoveryMode() bool`

HasRecoveryMode returns a boolean if a field has been set.

### SetRecoveryModeNil

`func (o *RecoverOracleNewTargetDatabaseConfig) SetRecoveryModeNil(b bool)`

 SetRecoveryModeNil sets the value for RecoveryMode to be an explicit nil

### UnsetRecoveryMode
`func (o *RecoverOracleNewTargetDatabaseConfig) UnsetRecoveryMode()`

UnsetRecoveryMode ensures that no value is present for RecoveryMode, not even an explicit nil
### GetShellEvironmentVars

`func (o *RecoverOracleNewTargetDatabaseConfig) GetShellEvironmentVars() []ShellKeyValuePair`

GetShellEvironmentVars returns the ShellEvironmentVars field if non-nil, zero value otherwise.

### GetShellEvironmentVarsOk

`func (o *RecoverOracleNewTargetDatabaseConfig) GetShellEvironmentVarsOk() (*[]ShellKeyValuePair, bool)`

GetShellEvironmentVarsOk returns a tuple with the ShellEvironmentVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShellEvironmentVars

`func (o *RecoverOracleNewTargetDatabaseConfig) SetShellEvironmentVars(v []ShellKeyValuePair)`

SetShellEvironmentVars sets ShellEvironmentVars field to given value.

### HasShellEvironmentVars

`func (o *RecoverOracleNewTargetDatabaseConfig) HasShellEvironmentVars() bool`

HasShellEvironmentVars returns a boolean if a field has been set.

### SetShellEvironmentVarsNil

`func (o *RecoverOracleNewTargetDatabaseConfig) SetShellEvironmentVarsNil(b bool)`

 SetShellEvironmentVarsNil sets the value for ShellEvironmentVars to be an explicit nil

### UnsetShellEvironmentVars
`func (o *RecoverOracleNewTargetDatabaseConfig) UnsetShellEvironmentVars()`

UnsetShellEvironmentVars ensures that no value is present for ShellEvironmentVars, not even an explicit nil
### GetGranularRestoreInfo

`func (o *RecoverOracleNewTargetDatabaseConfig) GetGranularRestoreInfo() RecoverOracleGranularRestoreInfo`

GetGranularRestoreInfo returns the GranularRestoreInfo field if non-nil, zero value otherwise.

### GetGranularRestoreInfoOk

`func (o *RecoverOracleNewTargetDatabaseConfig) GetGranularRestoreInfoOk() (*RecoverOracleGranularRestoreInfo, bool)`

GetGranularRestoreInfoOk returns a tuple with the GranularRestoreInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularRestoreInfo

`func (o *RecoverOracleNewTargetDatabaseConfig) SetGranularRestoreInfo(v RecoverOracleGranularRestoreInfo)`

SetGranularRestoreInfo sets GranularRestoreInfo field to given value.

### HasGranularRestoreInfo

`func (o *RecoverOracleNewTargetDatabaseConfig) HasGranularRestoreInfo() bool`

HasGranularRestoreInfo returns a boolean if a field has been set.

### GetOracleArchiveLogInfo

`func (o *RecoverOracleNewTargetDatabaseConfig) GetOracleArchiveLogInfo() OracleArchiveLogInfo`

GetOracleArchiveLogInfo returns the OracleArchiveLogInfo field if non-nil, zero value otherwise.

### GetOracleArchiveLogInfoOk

`func (o *RecoverOracleNewTargetDatabaseConfig) GetOracleArchiveLogInfoOk() (*OracleArchiveLogInfo, bool)`

GetOracleArchiveLogInfoOk returns a tuple with the OracleArchiveLogInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleArchiveLogInfo

`func (o *RecoverOracleNewTargetDatabaseConfig) SetOracleArchiveLogInfo(v OracleArchiveLogInfo)`

SetOracleArchiveLogInfo sets OracleArchiveLogInfo field to given value.

### HasOracleArchiveLogInfo

`func (o *RecoverOracleNewTargetDatabaseConfig) HasOracleArchiveLogInfo() bool`

HasOracleArchiveLogInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


