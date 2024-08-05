# RecoverOracleAppNewSourceConfigRecoverDatabaseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BctFilePath** | Pointer to **NullableString** | Specifies BCT file path. | [optional] 
**DatabaseName** | Pointer to **NullableString** | Specifies a new name for the restored database. If this field is not specified, then the original database will be overwritten after recovery. | [optional] 
**DbConfigFilePath** | Pointer to **NullableString** | Specifies the config file path on selected host which configures the restored database. | [optional] 
**DbFilesDestination** | Pointer to **NullableString** | Specifies the location to restore database files. | [optional] 
**DisasterRecoveryOptions** | Pointer to [**DisasterRecoveryOptions**](DisasterRecoveryOptions.md) |  | [optional] 
**EnableArchiveLogMode** | Pointer to **NullableBool** | Specifies archive log mode for oracle restore. | [optional] 
**IsMultiStageRestore** | Pointer to **NullableBool** | Specifies whether this task is a multistage restore task. If set, we migrate the DB after clone completes. | [optional] 
**NewNameClause** | Pointer to **NullableString** | Specifies newname clause for db files which allows user to have full control on how their database files can be renamed during the oracle alternate restore workflow. | [optional] 
**NoFilenameCheck** | Pointer to **NullableBool** | Specifies whether to validate filenames or not in Oracle alternate restore workflow. | [optional] 
**NumTempfiles** | Pointer to **NullableInt32** | Specifies no. of tempfiles to be used for the recovered database. | [optional] 
**OracleBaseFolder** | Pointer to **NullableString** | Specifies the oracle base folder at selected host. | [optional] 
**OracleHomeFolder** | Pointer to **NullableString** | Specifies the oracle home folder at selected host. | [optional] 
**OracleUpdateRestoreOptions** | Pointer to [**NullableMigrateCloneParams**](MigrateCloneParams.md) | Specifies the parameters that are needed for updating oracle restore options. | [optional] 
**PfileParameterMap** | Pointer to [**[]KeyValuePair**](KeyValuePair.md) | Specifies a key value pair for pfile parameters. | [optional] 
**RedoLogConfig** | Pointer to [**RedoLogGroupConfig**](RedoLogGroupConfig.md) | Specifies redo log config. | [optional] 
**RestoreToRac** | Pointer to **NullableBool** | Whether or not to restore to a RAC database. | [optional] 
**SkipCloneNid** | Pointer to **NullableBool** | Whether or not to skip the nid step in Oracle Clone workflow. Applicable to both smart and old clone workflow. | [optional] 
**DbChannels** | Pointer to [**[]OracleDbChannel**](OracleDbChannel.md) | Specifies the Oracle database node channels info. If not specified, the default values assigned by the server are applied to all the databases. | [optional] 
**GranularRestoreInfo** | Pointer to [**CommonOracleAppSourceConfigGranularRestoreInfo**](CommonOracleAppSourceConfigGranularRestoreInfo.md) |  | [optional] 
**OracleArchiveLogInfo** | Pointer to [**CommonOracleAppSourceConfigOracleArchiveLogInfo**](CommonOracleAppSourceConfigOracleArchiveLogInfo.md) |  | [optional] 
**OracleRecoveryValidationInfo** | Pointer to [**CommonOracleAppSourceConfigOracleRecoveryValidationInfo**](CommonOracleAppSourceConfigOracleRecoveryValidationInfo.md) |  | [optional] 
**RecoveryMode** | Pointer to **NullableBool** | Specifies if database should be left in recovery mode. | [optional] 
**RestoreSpfileOrPfileInfo** | Pointer to [**CommonOracleAppSourceConfigRestoreSpfileOrPfileInfo**](CommonOracleAppSourceConfigRestoreSpfileOrPfileInfo.md) |  | [optional] 
**RestoreTimeUsecs** | Pointer to **NullableInt64** | Specifies the time in the past to which the Oracle db needs to be restored. This allows for granular recovery of Oracle databases. If this is not set, the Oracle db will be restored from the full/incremental snapshot. | [optional] 
**ShellEvironmentVars** | Pointer to [**[]ShellKeyValuePair**](ShellKeyValuePair.md) | Specifies key value pairs of shell variables which defines the restore shell environment. | [optional] 
**UseScnForRestore** | Pointer to **NullableBool** | Specifies whether database recovery performed should use scn value or not. | [optional] 

## Methods

### NewRecoverOracleAppNewSourceConfigRecoverDatabaseParams

`func NewRecoverOracleAppNewSourceConfigRecoverDatabaseParams() *RecoverOracleAppNewSourceConfigRecoverDatabaseParams`

NewRecoverOracleAppNewSourceConfigRecoverDatabaseParams instantiates a new RecoverOracleAppNewSourceConfigRecoverDatabaseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverOracleAppNewSourceConfigRecoverDatabaseParamsWithDefaults

`func NewRecoverOracleAppNewSourceConfigRecoverDatabaseParamsWithDefaults() *RecoverOracleAppNewSourceConfigRecoverDatabaseParams`

NewRecoverOracleAppNewSourceConfigRecoverDatabaseParamsWithDefaults instantiates a new RecoverOracleAppNewSourceConfigRecoverDatabaseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBctFilePath

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetBctFilePath() string`

GetBctFilePath returns the BctFilePath field if non-nil, zero value otherwise.

### GetBctFilePathOk

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetBctFilePathOk() (*string, bool)`

GetBctFilePathOk returns a tuple with the BctFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBctFilePath

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetBctFilePath(v string)`

SetBctFilePath sets BctFilePath field to given value.

### HasBctFilePath

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) HasBctFilePath() bool`

HasBctFilePath returns a boolean if a field has been set.

### SetBctFilePathNil

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetBctFilePathNil(b bool)`

 SetBctFilePathNil sets the value for BctFilePath to be an explicit nil

### UnsetBctFilePath
`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) UnsetBctFilePath()`

UnsetBctFilePath ensures that no value is present for BctFilePath, not even an explicit nil
### GetDatabaseName

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### SetDatabaseNameNil

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetDatabaseNameNil(b bool)`

 SetDatabaseNameNil sets the value for DatabaseName to be an explicit nil

### UnsetDatabaseName
`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) UnsetDatabaseName()`

UnsetDatabaseName ensures that no value is present for DatabaseName, not even an explicit nil
### GetDbConfigFilePath

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetDbConfigFilePath() string`

GetDbConfigFilePath returns the DbConfigFilePath field if non-nil, zero value otherwise.

### GetDbConfigFilePathOk

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetDbConfigFilePathOk() (*string, bool)`

GetDbConfigFilePathOk returns a tuple with the DbConfigFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbConfigFilePath

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetDbConfigFilePath(v string)`

SetDbConfigFilePath sets DbConfigFilePath field to given value.

### HasDbConfigFilePath

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) HasDbConfigFilePath() bool`

HasDbConfigFilePath returns a boolean if a field has been set.

### SetDbConfigFilePathNil

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetDbConfigFilePathNil(b bool)`

 SetDbConfigFilePathNil sets the value for DbConfigFilePath to be an explicit nil

### UnsetDbConfigFilePath
`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) UnsetDbConfigFilePath()`

UnsetDbConfigFilePath ensures that no value is present for DbConfigFilePath, not even an explicit nil
### GetDbFilesDestination

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetDbFilesDestination() string`

GetDbFilesDestination returns the DbFilesDestination field if non-nil, zero value otherwise.

### GetDbFilesDestinationOk

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetDbFilesDestinationOk() (*string, bool)`

GetDbFilesDestinationOk returns a tuple with the DbFilesDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbFilesDestination

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetDbFilesDestination(v string)`

SetDbFilesDestination sets DbFilesDestination field to given value.

### HasDbFilesDestination

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) HasDbFilesDestination() bool`

HasDbFilesDestination returns a boolean if a field has been set.

### SetDbFilesDestinationNil

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetDbFilesDestinationNil(b bool)`

 SetDbFilesDestinationNil sets the value for DbFilesDestination to be an explicit nil

### UnsetDbFilesDestination
`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) UnsetDbFilesDestination()`

UnsetDbFilesDestination ensures that no value is present for DbFilesDestination, not even an explicit nil
### GetDisasterRecoveryOptions

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetDisasterRecoveryOptions() DisasterRecoveryOptions`

GetDisasterRecoveryOptions returns the DisasterRecoveryOptions field if non-nil, zero value otherwise.

### GetDisasterRecoveryOptionsOk

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetDisasterRecoveryOptionsOk() (*DisasterRecoveryOptions, bool)`

GetDisasterRecoveryOptionsOk returns a tuple with the DisasterRecoveryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisasterRecoveryOptions

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetDisasterRecoveryOptions(v DisasterRecoveryOptions)`

SetDisasterRecoveryOptions sets DisasterRecoveryOptions field to given value.

### HasDisasterRecoveryOptions

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) HasDisasterRecoveryOptions() bool`

HasDisasterRecoveryOptions returns a boolean if a field has been set.

### GetEnableArchiveLogMode

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetEnableArchiveLogMode() bool`

GetEnableArchiveLogMode returns the EnableArchiveLogMode field if non-nil, zero value otherwise.

### GetEnableArchiveLogModeOk

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetEnableArchiveLogModeOk() (*bool, bool)`

GetEnableArchiveLogModeOk returns a tuple with the EnableArchiveLogMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableArchiveLogMode

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetEnableArchiveLogMode(v bool)`

SetEnableArchiveLogMode sets EnableArchiveLogMode field to given value.

### HasEnableArchiveLogMode

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) HasEnableArchiveLogMode() bool`

HasEnableArchiveLogMode returns a boolean if a field has been set.

### SetEnableArchiveLogModeNil

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetEnableArchiveLogModeNil(b bool)`

 SetEnableArchiveLogModeNil sets the value for EnableArchiveLogMode to be an explicit nil

### UnsetEnableArchiveLogMode
`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) UnsetEnableArchiveLogMode()`

UnsetEnableArchiveLogMode ensures that no value is present for EnableArchiveLogMode, not even an explicit nil
### GetIsMultiStageRestore

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetIsMultiStageRestore() bool`

GetIsMultiStageRestore returns the IsMultiStageRestore field if non-nil, zero value otherwise.

### GetIsMultiStageRestoreOk

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetIsMultiStageRestoreOk() (*bool, bool)`

GetIsMultiStageRestoreOk returns a tuple with the IsMultiStageRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiStageRestore

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetIsMultiStageRestore(v bool)`

SetIsMultiStageRestore sets IsMultiStageRestore field to given value.

### HasIsMultiStageRestore

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) HasIsMultiStageRestore() bool`

HasIsMultiStageRestore returns a boolean if a field has been set.

### SetIsMultiStageRestoreNil

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetIsMultiStageRestoreNil(b bool)`

 SetIsMultiStageRestoreNil sets the value for IsMultiStageRestore to be an explicit nil

### UnsetIsMultiStageRestore
`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) UnsetIsMultiStageRestore()`

UnsetIsMultiStageRestore ensures that no value is present for IsMultiStageRestore, not even an explicit nil
### GetNewNameClause

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetNewNameClause() string`

GetNewNameClause returns the NewNameClause field if non-nil, zero value otherwise.

### GetNewNameClauseOk

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetNewNameClauseOk() (*string, bool)`

GetNewNameClauseOk returns a tuple with the NewNameClause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewNameClause

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetNewNameClause(v string)`

SetNewNameClause sets NewNameClause field to given value.

### HasNewNameClause

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) HasNewNameClause() bool`

HasNewNameClause returns a boolean if a field has been set.

### SetNewNameClauseNil

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetNewNameClauseNil(b bool)`

 SetNewNameClauseNil sets the value for NewNameClause to be an explicit nil

### UnsetNewNameClause
`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) UnsetNewNameClause()`

UnsetNewNameClause ensures that no value is present for NewNameClause, not even an explicit nil
### GetNoFilenameCheck

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetNoFilenameCheck() bool`

GetNoFilenameCheck returns the NoFilenameCheck field if non-nil, zero value otherwise.

### GetNoFilenameCheckOk

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetNoFilenameCheckOk() (*bool, bool)`

GetNoFilenameCheckOk returns a tuple with the NoFilenameCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoFilenameCheck

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetNoFilenameCheck(v bool)`

SetNoFilenameCheck sets NoFilenameCheck field to given value.

### HasNoFilenameCheck

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) HasNoFilenameCheck() bool`

HasNoFilenameCheck returns a boolean if a field has been set.

### SetNoFilenameCheckNil

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetNoFilenameCheckNil(b bool)`

 SetNoFilenameCheckNil sets the value for NoFilenameCheck to be an explicit nil

### UnsetNoFilenameCheck
`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) UnsetNoFilenameCheck()`

UnsetNoFilenameCheck ensures that no value is present for NoFilenameCheck, not even an explicit nil
### GetNumTempfiles

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetNumTempfiles() int32`

GetNumTempfiles returns the NumTempfiles field if non-nil, zero value otherwise.

### GetNumTempfilesOk

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetNumTempfilesOk() (*int32, bool)`

GetNumTempfilesOk returns a tuple with the NumTempfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTempfiles

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetNumTempfiles(v int32)`

SetNumTempfiles sets NumTempfiles field to given value.

### HasNumTempfiles

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) HasNumTempfiles() bool`

HasNumTempfiles returns a boolean if a field has been set.

### SetNumTempfilesNil

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetNumTempfilesNil(b bool)`

 SetNumTempfilesNil sets the value for NumTempfiles to be an explicit nil

### UnsetNumTempfiles
`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) UnsetNumTempfiles()`

UnsetNumTempfiles ensures that no value is present for NumTempfiles, not even an explicit nil
### GetOracleBaseFolder

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetOracleBaseFolder() string`

GetOracleBaseFolder returns the OracleBaseFolder field if non-nil, zero value otherwise.

### GetOracleBaseFolderOk

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetOracleBaseFolderOk() (*string, bool)`

GetOracleBaseFolderOk returns a tuple with the OracleBaseFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleBaseFolder

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetOracleBaseFolder(v string)`

SetOracleBaseFolder sets OracleBaseFolder field to given value.

### HasOracleBaseFolder

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) HasOracleBaseFolder() bool`

HasOracleBaseFolder returns a boolean if a field has been set.

### SetOracleBaseFolderNil

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetOracleBaseFolderNil(b bool)`

 SetOracleBaseFolderNil sets the value for OracleBaseFolder to be an explicit nil

### UnsetOracleBaseFolder
`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) UnsetOracleBaseFolder()`

UnsetOracleBaseFolder ensures that no value is present for OracleBaseFolder, not even an explicit nil
### GetOracleHomeFolder

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetOracleHomeFolder() string`

GetOracleHomeFolder returns the OracleHomeFolder field if non-nil, zero value otherwise.

### GetOracleHomeFolderOk

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetOracleHomeFolderOk() (*string, bool)`

GetOracleHomeFolderOk returns a tuple with the OracleHomeFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleHomeFolder

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetOracleHomeFolder(v string)`

SetOracleHomeFolder sets OracleHomeFolder field to given value.

### HasOracleHomeFolder

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) HasOracleHomeFolder() bool`

HasOracleHomeFolder returns a boolean if a field has been set.

### SetOracleHomeFolderNil

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetOracleHomeFolderNil(b bool)`

 SetOracleHomeFolderNil sets the value for OracleHomeFolder to be an explicit nil

### UnsetOracleHomeFolder
`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) UnsetOracleHomeFolder()`

UnsetOracleHomeFolder ensures that no value is present for OracleHomeFolder, not even an explicit nil
### GetOracleUpdateRestoreOptions

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetOracleUpdateRestoreOptions() MigrateCloneParams`

GetOracleUpdateRestoreOptions returns the OracleUpdateRestoreOptions field if non-nil, zero value otherwise.

### GetOracleUpdateRestoreOptionsOk

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetOracleUpdateRestoreOptionsOk() (*MigrateCloneParams, bool)`

GetOracleUpdateRestoreOptionsOk returns a tuple with the OracleUpdateRestoreOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleUpdateRestoreOptions

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetOracleUpdateRestoreOptions(v MigrateCloneParams)`

SetOracleUpdateRestoreOptions sets OracleUpdateRestoreOptions field to given value.

### HasOracleUpdateRestoreOptions

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) HasOracleUpdateRestoreOptions() bool`

HasOracleUpdateRestoreOptions returns a boolean if a field has been set.

### SetOracleUpdateRestoreOptionsNil

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetOracleUpdateRestoreOptionsNil(b bool)`

 SetOracleUpdateRestoreOptionsNil sets the value for OracleUpdateRestoreOptions to be an explicit nil

### UnsetOracleUpdateRestoreOptions
`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) UnsetOracleUpdateRestoreOptions()`

UnsetOracleUpdateRestoreOptions ensures that no value is present for OracleUpdateRestoreOptions, not even an explicit nil
### GetPfileParameterMap

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetPfileParameterMap() []KeyValuePair`

GetPfileParameterMap returns the PfileParameterMap field if non-nil, zero value otherwise.

### GetPfileParameterMapOk

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetPfileParameterMapOk() (*[]KeyValuePair, bool)`

GetPfileParameterMapOk returns a tuple with the PfileParameterMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfileParameterMap

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetPfileParameterMap(v []KeyValuePair)`

SetPfileParameterMap sets PfileParameterMap field to given value.

### HasPfileParameterMap

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) HasPfileParameterMap() bool`

HasPfileParameterMap returns a boolean if a field has been set.

### SetPfileParameterMapNil

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetPfileParameterMapNil(b bool)`

 SetPfileParameterMapNil sets the value for PfileParameterMap to be an explicit nil

### UnsetPfileParameterMap
`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) UnsetPfileParameterMap()`

UnsetPfileParameterMap ensures that no value is present for PfileParameterMap, not even an explicit nil
### GetRedoLogConfig

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetRedoLogConfig() RedoLogGroupConfig`

GetRedoLogConfig returns the RedoLogConfig field if non-nil, zero value otherwise.

### GetRedoLogConfigOk

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetRedoLogConfigOk() (*RedoLogGroupConfig, bool)`

GetRedoLogConfigOk returns a tuple with the RedoLogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedoLogConfig

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetRedoLogConfig(v RedoLogGroupConfig)`

SetRedoLogConfig sets RedoLogConfig field to given value.

### HasRedoLogConfig

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) HasRedoLogConfig() bool`

HasRedoLogConfig returns a boolean if a field has been set.

### GetRestoreToRac

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetRestoreToRac() bool`

GetRestoreToRac returns the RestoreToRac field if non-nil, zero value otherwise.

### GetRestoreToRacOk

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetRestoreToRacOk() (*bool, bool)`

GetRestoreToRacOk returns a tuple with the RestoreToRac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreToRac

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetRestoreToRac(v bool)`

SetRestoreToRac sets RestoreToRac field to given value.

### HasRestoreToRac

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) HasRestoreToRac() bool`

HasRestoreToRac returns a boolean if a field has been set.

### SetRestoreToRacNil

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetRestoreToRacNil(b bool)`

 SetRestoreToRacNil sets the value for RestoreToRac to be an explicit nil

### UnsetRestoreToRac
`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) UnsetRestoreToRac()`

UnsetRestoreToRac ensures that no value is present for RestoreToRac, not even an explicit nil
### GetSkipCloneNid

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetSkipCloneNid() bool`

GetSkipCloneNid returns the SkipCloneNid field if non-nil, zero value otherwise.

### GetSkipCloneNidOk

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetSkipCloneNidOk() (*bool, bool)`

GetSkipCloneNidOk returns a tuple with the SkipCloneNid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCloneNid

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetSkipCloneNid(v bool)`

SetSkipCloneNid sets SkipCloneNid field to given value.

### HasSkipCloneNid

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) HasSkipCloneNid() bool`

HasSkipCloneNid returns a boolean if a field has been set.

### SetSkipCloneNidNil

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetSkipCloneNidNil(b bool)`

 SetSkipCloneNidNil sets the value for SkipCloneNid to be an explicit nil

### UnsetSkipCloneNid
`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) UnsetSkipCloneNid()`

UnsetSkipCloneNid ensures that no value is present for SkipCloneNid, not even an explicit nil
### GetDbChannels

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetDbChannels() []OracleDbChannel`

GetDbChannels returns the DbChannels field if non-nil, zero value otherwise.

### GetDbChannelsOk

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetDbChannelsOk() (*[]OracleDbChannel, bool)`

GetDbChannelsOk returns a tuple with the DbChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbChannels

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetDbChannels(v []OracleDbChannel)`

SetDbChannels sets DbChannels field to given value.

### HasDbChannels

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) HasDbChannels() bool`

HasDbChannels returns a boolean if a field has been set.

### SetDbChannelsNil

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetDbChannelsNil(b bool)`

 SetDbChannelsNil sets the value for DbChannels to be an explicit nil

### UnsetDbChannels
`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) UnsetDbChannels()`

UnsetDbChannels ensures that no value is present for DbChannels, not even an explicit nil
### GetGranularRestoreInfo

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetGranularRestoreInfo() CommonOracleAppSourceConfigGranularRestoreInfo`

GetGranularRestoreInfo returns the GranularRestoreInfo field if non-nil, zero value otherwise.

### GetGranularRestoreInfoOk

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetGranularRestoreInfoOk() (*CommonOracleAppSourceConfigGranularRestoreInfo, bool)`

GetGranularRestoreInfoOk returns a tuple with the GranularRestoreInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularRestoreInfo

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetGranularRestoreInfo(v CommonOracleAppSourceConfigGranularRestoreInfo)`

SetGranularRestoreInfo sets GranularRestoreInfo field to given value.

### HasGranularRestoreInfo

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) HasGranularRestoreInfo() bool`

HasGranularRestoreInfo returns a boolean if a field has been set.

### GetOracleArchiveLogInfo

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetOracleArchiveLogInfo() CommonOracleAppSourceConfigOracleArchiveLogInfo`

GetOracleArchiveLogInfo returns the OracleArchiveLogInfo field if non-nil, zero value otherwise.

### GetOracleArchiveLogInfoOk

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetOracleArchiveLogInfoOk() (*CommonOracleAppSourceConfigOracleArchiveLogInfo, bool)`

GetOracleArchiveLogInfoOk returns a tuple with the OracleArchiveLogInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleArchiveLogInfo

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetOracleArchiveLogInfo(v CommonOracleAppSourceConfigOracleArchiveLogInfo)`

SetOracleArchiveLogInfo sets OracleArchiveLogInfo field to given value.

### HasOracleArchiveLogInfo

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) HasOracleArchiveLogInfo() bool`

HasOracleArchiveLogInfo returns a boolean if a field has been set.

### GetOracleRecoveryValidationInfo

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetOracleRecoveryValidationInfo() CommonOracleAppSourceConfigOracleRecoveryValidationInfo`

GetOracleRecoveryValidationInfo returns the OracleRecoveryValidationInfo field if non-nil, zero value otherwise.

### GetOracleRecoveryValidationInfoOk

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetOracleRecoveryValidationInfoOk() (*CommonOracleAppSourceConfigOracleRecoveryValidationInfo, bool)`

GetOracleRecoveryValidationInfoOk returns a tuple with the OracleRecoveryValidationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleRecoveryValidationInfo

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetOracleRecoveryValidationInfo(v CommonOracleAppSourceConfigOracleRecoveryValidationInfo)`

SetOracleRecoveryValidationInfo sets OracleRecoveryValidationInfo field to given value.

### HasOracleRecoveryValidationInfo

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) HasOracleRecoveryValidationInfo() bool`

HasOracleRecoveryValidationInfo returns a boolean if a field has been set.

### GetRecoveryMode

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetRecoveryMode() bool`

GetRecoveryMode returns the RecoveryMode field if non-nil, zero value otherwise.

### GetRecoveryModeOk

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetRecoveryModeOk() (*bool, bool)`

GetRecoveryModeOk returns a tuple with the RecoveryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryMode

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetRecoveryMode(v bool)`

SetRecoveryMode sets RecoveryMode field to given value.

### HasRecoveryMode

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) HasRecoveryMode() bool`

HasRecoveryMode returns a boolean if a field has been set.

### SetRecoveryModeNil

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetRecoveryModeNil(b bool)`

 SetRecoveryModeNil sets the value for RecoveryMode to be an explicit nil

### UnsetRecoveryMode
`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) UnsetRecoveryMode()`

UnsetRecoveryMode ensures that no value is present for RecoveryMode, not even an explicit nil
### GetRestoreSpfileOrPfileInfo

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetRestoreSpfileOrPfileInfo() CommonOracleAppSourceConfigRestoreSpfileOrPfileInfo`

GetRestoreSpfileOrPfileInfo returns the RestoreSpfileOrPfileInfo field if non-nil, zero value otherwise.

### GetRestoreSpfileOrPfileInfoOk

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetRestoreSpfileOrPfileInfoOk() (*CommonOracleAppSourceConfigRestoreSpfileOrPfileInfo, bool)`

GetRestoreSpfileOrPfileInfoOk returns a tuple with the RestoreSpfileOrPfileInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreSpfileOrPfileInfo

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetRestoreSpfileOrPfileInfo(v CommonOracleAppSourceConfigRestoreSpfileOrPfileInfo)`

SetRestoreSpfileOrPfileInfo sets RestoreSpfileOrPfileInfo field to given value.

### HasRestoreSpfileOrPfileInfo

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) HasRestoreSpfileOrPfileInfo() bool`

HasRestoreSpfileOrPfileInfo returns a boolean if a field has been set.

### GetRestoreTimeUsecs

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetRestoreTimeUsecs() int64`

GetRestoreTimeUsecs returns the RestoreTimeUsecs field if non-nil, zero value otherwise.

### GetRestoreTimeUsecsOk

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetRestoreTimeUsecsOk() (*int64, bool)`

GetRestoreTimeUsecsOk returns a tuple with the RestoreTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTimeUsecs

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetRestoreTimeUsecs(v int64)`

SetRestoreTimeUsecs sets RestoreTimeUsecs field to given value.

### HasRestoreTimeUsecs

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) HasRestoreTimeUsecs() bool`

HasRestoreTimeUsecs returns a boolean if a field has been set.

### SetRestoreTimeUsecsNil

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetRestoreTimeUsecsNil(b bool)`

 SetRestoreTimeUsecsNil sets the value for RestoreTimeUsecs to be an explicit nil

### UnsetRestoreTimeUsecs
`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) UnsetRestoreTimeUsecs()`

UnsetRestoreTimeUsecs ensures that no value is present for RestoreTimeUsecs, not even an explicit nil
### GetShellEvironmentVars

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetShellEvironmentVars() []ShellKeyValuePair`

GetShellEvironmentVars returns the ShellEvironmentVars field if non-nil, zero value otherwise.

### GetShellEvironmentVarsOk

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetShellEvironmentVarsOk() (*[]ShellKeyValuePair, bool)`

GetShellEvironmentVarsOk returns a tuple with the ShellEvironmentVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShellEvironmentVars

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetShellEvironmentVars(v []ShellKeyValuePair)`

SetShellEvironmentVars sets ShellEvironmentVars field to given value.

### HasShellEvironmentVars

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) HasShellEvironmentVars() bool`

HasShellEvironmentVars returns a boolean if a field has been set.

### SetShellEvironmentVarsNil

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetShellEvironmentVarsNil(b bool)`

 SetShellEvironmentVarsNil sets the value for ShellEvironmentVars to be an explicit nil

### UnsetShellEvironmentVars
`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) UnsetShellEvironmentVars()`

UnsetShellEvironmentVars ensures that no value is present for ShellEvironmentVars, not even an explicit nil
### GetUseScnForRestore

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetUseScnForRestore() bool`

GetUseScnForRestore returns the UseScnForRestore field if non-nil, zero value otherwise.

### GetUseScnForRestoreOk

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) GetUseScnForRestoreOk() (*bool, bool)`

GetUseScnForRestoreOk returns a tuple with the UseScnForRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseScnForRestore

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetUseScnForRestore(v bool)`

SetUseScnForRestore sets UseScnForRestore field to given value.

### HasUseScnForRestore

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) HasUseScnForRestore() bool`

HasUseScnForRestore returns a boolean if a field has been set.

### SetUseScnForRestoreNil

`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) SetUseScnForRestoreNil(b bool)`

 SetUseScnForRestoreNil sets the value for UseScnForRestore to be an explicit nil

### UnsetUseScnForRestore
`func (o *RecoverOracleAppNewSourceConfigRecoverDatabaseParams) UnsetUseScnForRestore()`

UnsetUseScnForRestore ensures that no value is present for UseScnForRestore, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


