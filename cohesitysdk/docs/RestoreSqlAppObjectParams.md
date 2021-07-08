# RestoreSqlAppObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaptureTailLogs** | Pointer to **NullableBool** | Set to true if tail logs are to be captured before the restore operation. This is only applicable if we are restoring the SQL database to its original source, and the database is not being renamed. | [optional] 
**ContinueAfterError** | Pointer to **NullableBool** | Whether restore should continue after encountering a page checksum error. | [optional] 
**DataFileDestination** | Pointer to **NullableString** | Which directory to put the database data files. Missing directory will be automatically created. Cannot be empty if not restoring to the original SQL instance. | [optional] 
**DbRestoreOverwritePolicy** | Pointer to **NullableInt32** | Policy to overwrite an existing DB during a restore operation. | [optional] 
**EnableChecksum** | Pointer to **NullableBool** | Whether restore checksums are enabled. | [optional] 
**InstanceName** | Pointer to **NullableString** | The name of the SQL instance that we restore database to. If target_host is not empty, this also cannot be empty. | [optional] 
**IsAutoSyncEnabled** | Pointer to **NullableBool** | The following field is set if auto_sync for multi-stage SQL restore task is enabled. This field is valid only if is_multi_state_restore is set to true. | [optional] 
**IsMultiStageRestore** | Pointer to **NullableBool** | The following field is set if we are creating a multi-stage SQL restore task needed for features such as Hot-Standby. | [optional] 
**KeepCdc** | Pointer to **NullableBool** | Set to true to keep cdc on restored database. | [optional] 
**LogFileDestination** | Pointer to **NullableString** | Which directory to put the database log files. Missing directory will be automatically created. Cannot be empty if not restoring to the original SQL instance. | [optional] 
**MultiStageRestoreOptions** | Pointer to [**SqlUpdateRestoreTaskOptions**](SqlUpdateRestoreTaskOptions.md) |  | [optional] 
**NewDatabaseName** | Pointer to **NullableString** | The new name of the database, if it is going to be renamed. app_entity in RestoreAppObject has to be non-empty for the renaming, otherwise it does not make sense to rename all databases in the owner. | [optional] 
**RestoreTimeSecs** | Pointer to **NullableInt64** | The time to which the SQL database needs to be restored. This allows for granular recovery of SQL databases. If this is not set, the SQL database will be recovered to the full/incremental snapshot (specified in the owner&#39;s restore object in AppOwnerRestoreInfo). | [optional] 
**ResumeRestore** | Pointer to **NullableBool** | Resume restore if sql instance/database exist in restore/recovering state. The database might be in restore/recovering state if previous restore failed or previous  restore was attempted  with norecovery option. | [optional] 
**SecondaryDataFileDestination** | Pointer to **NullableString** | Which directory to put the secondary data files of the database. Secondary data files are optional and are user defined. The recommended file name extension for these is \&quot;.ndf\&quot;.  If this option is specified, the directory will be automatically created if its missing. | [optional] 
**SecondaryDataFileDestinationVec** | Pointer to [**[]FilesToDirectoryMapping**](FilesToDirectoryMapping.md) | Specify the secondary data files and corresponding direcories of the DB. Secondary data files are optional and are user defined. The recommended file extension for secondary files is \&quot;.ndf\&quot;.  If this option is specified and the destination folders do not exist they will be automatically created. | [optional] 
**WithClause** | Pointer to **NullableString** | &#39;with_clause&#39; contains &#39;with clause&#39; to be used in native sql restore command. This is only applicable for db restore of native sql backup. Here user can specify multiple restore options. Example: \&quot;WITH BUFFERCOUNT &#x3D; 575, MAXTRANSFERSIZE &#x3D; 2097152\&quot;. If this is not specified, we use the value specified in magneto_sql_native_restore_with_clause gflag. | [optional] 
**WithNoRecovery** | Pointer to **NullableBool** | Set to true if we want to recover the database in \&quot;NO_RECOVERY\&quot; mode which does not bring it online after restore. | [optional] 

## Methods

### NewRestoreSqlAppObjectParams

`func NewRestoreSqlAppObjectParams() *RestoreSqlAppObjectParams`

NewRestoreSqlAppObjectParams instantiates a new RestoreSqlAppObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreSqlAppObjectParamsWithDefaults

`func NewRestoreSqlAppObjectParamsWithDefaults() *RestoreSqlAppObjectParams`

NewRestoreSqlAppObjectParamsWithDefaults instantiates a new RestoreSqlAppObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaptureTailLogs

`func (o *RestoreSqlAppObjectParams) GetCaptureTailLogs() bool`

GetCaptureTailLogs returns the CaptureTailLogs field if non-nil, zero value otherwise.

### GetCaptureTailLogsOk

`func (o *RestoreSqlAppObjectParams) GetCaptureTailLogsOk() (*bool, bool)`

GetCaptureTailLogsOk returns a tuple with the CaptureTailLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureTailLogs

`func (o *RestoreSqlAppObjectParams) SetCaptureTailLogs(v bool)`

SetCaptureTailLogs sets CaptureTailLogs field to given value.

### HasCaptureTailLogs

`func (o *RestoreSqlAppObjectParams) HasCaptureTailLogs() bool`

HasCaptureTailLogs returns a boolean if a field has been set.

### SetCaptureTailLogsNil

`func (o *RestoreSqlAppObjectParams) SetCaptureTailLogsNil(b bool)`

 SetCaptureTailLogsNil sets the value for CaptureTailLogs to be an explicit nil

### UnsetCaptureTailLogs
`func (o *RestoreSqlAppObjectParams) UnsetCaptureTailLogs()`

UnsetCaptureTailLogs ensures that no value is present for CaptureTailLogs, not even an explicit nil
### GetContinueAfterError

`func (o *RestoreSqlAppObjectParams) GetContinueAfterError() bool`

GetContinueAfterError returns the ContinueAfterError field if non-nil, zero value otherwise.

### GetContinueAfterErrorOk

`func (o *RestoreSqlAppObjectParams) GetContinueAfterErrorOk() (*bool, bool)`

GetContinueAfterErrorOk returns a tuple with the ContinueAfterError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueAfterError

`func (o *RestoreSqlAppObjectParams) SetContinueAfterError(v bool)`

SetContinueAfterError sets ContinueAfterError field to given value.

### HasContinueAfterError

`func (o *RestoreSqlAppObjectParams) HasContinueAfterError() bool`

HasContinueAfterError returns a boolean if a field has been set.

### SetContinueAfterErrorNil

`func (o *RestoreSqlAppObjectParams) SetContinueAfterErrorNil(b bool)`

 SetContinueAfterErrorNil sets the value for ContinueAfterError to be an explicit nil

### UnsetContinueAfterError
`func (o *RestoreSqlAppObjectParams) UnsetContinueAfterError()`

UnsetContinueAfterError ensures that no value is present for ContinueAfterError, not even an explicit nil
### GetDataFileDestination

`func (o *RestoreSqlAppObjectParams) GetDataFileDestination() string`

GetDataFileDestination returns the DataFileDestination field if non-nil, zero value otherwise.

### GetDataFileDestinationOk

`func (o *RestoreSqlAppObjectParams) GetDataFileDestinationOk() (*string, bool)`

GetDataFileDestinationOk returns a tuple with the DataFileDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFileDestination

`func (o *RestoreSqlAppObjectParams) SetDataFileDestination(v string)`

SetDataFileDestination sets DataFileDestination field to given value.

### HasDataFileDestination

`func (o *RestoreSqlAppObjectParams) HasDataFileDestination() bool`

HasDataFileDestination returns a boolean if a field has been set.

### SetDataFileDestinationNil

`func (o *RestoreSqlAppObjectParams) SetDataFileDestinationNil(b bool)`

 SetDataFileDestinationNil sets the value for DataFileDestination to be an explicit nil

### UnsetDataFileDestination
`func (o *RestoreSqlAppObjectParams) UnsetDataFileDestination()`

UnsetDataFileDestination ensures that no value is present for DataFileDestination, not even an explicit nil
### GetDbRestoreOverwritePolicy

`func (o *RestoreSqlAppObjectParams) GetDbRestoreOverwritePolicy() int32`

GetDbRestoreOverwritePolicy returns the DbRestoreOverwritePolicy field if non-nil, zero value otherwise.

### GetDbRestoreOverwritePolicyOk

`func (o *RestoreSqlAppObjectParams) GetDbRestoreOverwritePolicyOk() (*int32, bool)`

GetDbRestoreOverwritePolicyOk returns a tuple with the DbRestoreOverwritePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbRestoreOverwritePolicy

`func (o *RestoreSqlAppObjectParams) SetDbRestoreOverwritePolicy(v int32)`

SetDbRestoreOverwritePolicy sets DbRestoreOverwritePolicy field to given value.

### HasDbRestoreOverwritePolicy

`func (o *RestoreSqlAppObjectParams) HasDbRestoreOverwritePolicy() bool`

HasDbRestoreOverwritePolicy returns a boolean if a field has been set.

### SetDbRestoreOverwritePolicyNil

`func (o *RestoreSqlAppObjectParams) SetDbRestoreOverwritePolicyNil(b bool)`

 SetDbRestoreOverwritePolicyNil sets the value for DbRestoreOverwritePolicy to be an explicit nil

### UnsetDbRestoreOverwritePolicy
`func (o *RestoreSqlAppObjectParams) UnsetDbRestoreOverwritePolicy()`

UnsetDbRestoreOverwritePolicy ensures that no value is present for DbRestoreOverwritePolicy, not even an explicit nil
### GetEnableChecksum

`func (o *RestoreSqlAppObjectParams) GetEnableChecksum() bool`

GetEnableChecksum returns the EnableChecksum field if non-nil, zero value otherwise.

### GetEnableChecksumOk

`func (o *RestoreSqlAppObjectParams) GetEnableChecksumOk() (*bool, bool)`

GetEnableChecksumOk returns a tuple with the EnableChecksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableChecksum

`func (o *RestoreSqlAppObjectParams) SetEnableChecksum(v bool)`

SetEnableChecksum sets EnableChecksum field to given value.

### HasEnableChecksum

`func (o *RestoreSqlAppObjectParams) HasEnableChecksum() bool`

HasEnableChecksum returns a boolean if a field has been set.

### SetEnableChecksumNil

`func (o *RestoreSqlAppObjectParams) SetEnableChecksumNil(b bool)`

 SetEnableChecksumNil sets the value for EnableChecksum to be an explicit nil

### UnsetEnableChecksum
`func (o *RestoreSqlAppObjectParams) UnsetEnableChecksum()`

UnsetEnableChecksum ensures that no value is present for EnableChecksum, not even an explicit nil
### GetInstanceName

`func (o *RestoreSqlAppObjectParams) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *RestoreSqlAppObjectParams) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *RestoreSqlAppObjectParams) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *RestoreSqlAppObjectParams) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### SetInstanceNameNil

`func (o *RestoreSqlAppObjectParams) SetInstanceNameNil(b bool)`

 SetInstanceNameNil sets the value for InstanceName to be an explicit nil

### UnsetInstanceName
`func (o *RestoreSqlAppObjectParams) UnsetInstanceName()`

UnsetInstanceName ensures that no value is present for InstanceName, not even an explicit nil
### GetIsAutoSyncEnabled

`func (o *RestoreSqlAppObjectParams) GetIsAutoSyncEnabled() bool`

GetIsAutoSyncEnabled returns the IsAutoSyncEnabled field if non-nil, zero value otherwise.

### GetIsAutoSyncEnabledOk

`func (o *RestoreSqlAppObjectParams) GetIsAutoSyncEnabledOk() (*bool, bool)`

GetIsAutoSyncEnabledOk returns a tuple with the IsAutoSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoSyncEnabled

`func (o *RestoreSqlAppObjectParams) SetIsAutoSyncEnabled(v bool)`

SetIsAutoSyncEnabled sets IsAutoSyncEnabled field to given value.

### HasIsAutoSyncEnabled

`func (o *RestoreSqlAppObjectParams) HasIsAutoSyncEnabled() bool`

HasIsAutoSyncEnabled returns a boolean if a field has been set.

### SetIsAutoSyncEnabledNil

`func (o *RestoreSqlAppObjectParams) SetIsAutoSyncEnabledNil(b bool)`

 SetIsAutoSyncEnabledNil sets the value for IsAutoSyncEnabled to be an explicit nil

### UnsetIsAutoSyncEnabled
`func (o *RestoreSqlAppObjectParams) UnsetIsAutoSyncEnabled()`

UnsetIsAutoSyncEnabled ensures that no value is present for IsAutoSyncEnabled, not even an explicit nil
### GetIsMultiStageRestore

`func (o *RestoreSqlAppObjectParams) GetIsMultiStageRestore() bool`

GetIsMultiStageRestore returns the IsMultiStageRestore field if non-nil, zero value otherwise.

### GetIsMultiStageRestoreOk

`func (o *RestoreSqlAppObjectParams) GetIsMultiStageRestoreOk() (*bool, bool)`

GetIsMultiStageRestoreOk returns a tuple with the IsMultiStageRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiStageRestore

`func (o *RestoreSqlAppObjectParams) SetIsMultiStageRestore(v bool)`

SetIsMultiStageRestore sets IsMultiStageRestore field to given value.

### HasIsMultiStageRestore

`func (o *RestoreSqlAppObjectParams) HasIsMultiStageRestore() bool`

HasIsMultiStageRestore returns a boolean if a field has been set.

### SetIsMultiStageRestoreNil

`func (o *RestoreSqlAppObjectParams) SetIsMultiStageRestoreNil(b bool)`

 SetIsMultiStageRestoreNil sets the value for IsMultiStageRestore to be an explicit nil

### UnsetIsMultiStageRestore
`func (o *RestoreSqlAppObjectParams) UnsetIsMultiStageRestore()`

UnsetIsMultiStageRestore ensures that no value is present for IsMultiStageRestore, not even an explicit nil
### GetKeepCdc

`func (o *RestoreSqlAppObjectParams) GetKeepCdc() bool`

GetKeepCdc returns the KeepCdc field if non-nil, zero value otherwise.

### GetKeepCdcOk

`func (o *RestoreSqlAppObjectParams) GetKeepCdcOk() (*bool, bool)`

GetKeepCdcOk returns a tuple with the KeepCdc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepCdc

`func (o *RestoreSqlAppObjectParams) SetKeepCdc(v bool)`

SetKeepCdc sets KeepCdc field to given value.

### HasKeepCdc

`func (o *RestoreSqlAppObjectParams) HasKeepCdc() bool`

HasKeepCdc returns a boolean if a field has been set.

### SetKeepCdcNil

`func (o *RestoreSqlAppObjectParams) SetKeepCdcNil(b bool)`

 SetKeepCdcNil sets the value for KeepCdc to be an explicit nil

### UnsetKeepCdc
`func (o *RestoreSqlAppObjectParams) UnsetKeepCdc()`

UnsetKeepCdc ensures that no value is present for KeepCdc, not even an explicit nil
### GetLogFileDestination

`func (o *RestoreSqlAppObjectParams) GetLogFileDestination() string`

GetLogFileDestination returns the LogFileDestination field if non-nil, zero value otherwise.

### GetLogFileDestinationOk

`func (o *RestoreSqlAppObjectParams) GetLogFileDestinationOk() (*string, bool)`

GetLogFileDestinationOk returns a tuple with the LogFileDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFileDestination

`func (o *RestoreSqlAppObjectParams) SetLogFileDestination(v string)`

SetLogFileDestination sets LogFileDestination field to given value.

### HasLogFileDestination

`func (o *RestoreSqlAppObjectParams) HasLogFileDestination() bool`

HasLogFileDestination returns a boolean if a field has been set.

### SetLogFileDestinationNil

`func (o *RestoreSqlAppObjectParams) SetLogFileDestinationNil(b bool)`

 SetLogFileDestinationNil sets the value for LogFileDestination to be an explicit nil

### UnsetLogFileDestination
`func (o *RestoreSqlAppObjectParams) UnsetLogFileDestination()`

UnsetLogFileDestination ensures that no value is present for LogFileDestination, not even an explicit nil
### GetMultiStageRestoreOptions

`func (o *RestoreSqlAppObjectParams) GetMultiStageRestoreOptions() SqlUpdateRestoreTaskOptions`

GetMultiStageRestoreOptions returns the MultiStageRestoreOptions field if non-nil, zero value otherwise.

### GetMultiStageRestoreOptionsOk

`func (o *RestoreSqlAppObjectParams) GetMultiStageRestoreOptionsOk() (*SqlUpdateRestoreTaskOptions, bool)`

GetMultiStageRestoreOptionsOk returns a tuple with the MultiStageRestoreOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiStageRestoreOptions

`func (o *RestoreSqlAppObjectParams) SetMultiStageRestoreOptions(v SqlUpdateRestoreTaskOptions)`

SetMultiStageRestoreOptions sets MultiStageRestoreOptions field to given value.

### HasMultiStageRestoreOptions

`func (o *RestoreSqlAppObjectParams) HasMultiStageRestoreOptions() bool`

HasMultiStageRestoreOptions returns a boolean if a field has been set.

### GetNewDatabaseName

`func (o *RestoreSqlAppObjectParams) GetNewDatabaseName() string`

GetNewDatabaseName returns the NewDatabaseName field if non-nil, zero value otherwise.

### GetNewDatabaseNameOk

`func (o *RestoreSqlAppObjectParams) GetNewDatabaseNameOk() (*string, bool)`

GetNewDatabaseNameOk returns a tuple with the NewDatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDatabaseName

`func (o *RestoreSqlAppObjectParams) SetNewDatabaseName(v string)`

SetNewDatabaseName sets NewDatabaseName field to given value.

### HasNewDatabaseName

`func (o *RestoreSqlAppObjectParams) HasNewDatabaseName() bool`

HasNewDatabaseName returns a boolean if a field has been set.

### SetNewDatabaseNameNil

`func (o *RestoreSqlAppObjectParams) SetNewDatabaseNameNil(b bool)`

 SetNewDatabaseNameNil sets the value for NewDatabaseName to be an explicit nil

### UnsetNewDatabaseName
`func (o *RestoreSqlAppObjectParams) UnsetNewDatabaseName()`

UnsetNewDatabaseName ensures that no value is present for NewDatabaseName, not even an explicit nil
### GetRestoreTimeSecs

`func (o *RestoreSqlAppObjectParams) GetRestoreTimeSecs() int64`

GetRestoreTimeSecs returns the RestoreTimeSecs field if non-nil, zero value otherwise.

### GetRestoreTimeSecsOk

`func (o *RestoreSqlAppObjectParams) GetRestoreTimeSecsOk() (*int64, bool)`

GetRestoreTimeSecsOk returns a tuple with the RestoreTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTimeSecs

`func (o *RestoreSqlAppObjectParams) SetRestoreTimeSecs(v int64)`

SetRestoreTimeSecs sets RestoreTimeSecs field to given value.

### HasRestoreTimeSecs

`func (o *RestoreSqlAppObjectParams) HasRestoreTimeSecs() bool`

HasRestoreTimeSecs returns a boolean if a field has been set.

### SetRestoreTimeSecsNil

`func (o *RestoreSqlAppObjectParams) SetRestoreTimeSecsNil(b bool)`

 SetRestoreTimeSecsNil sets the value for RestoreTimeSecs to be an explicit nil

### UnsetRestoreTimeSecs
`func (o *RestoreSqlAppObjectParams) UnsetRestoreTimeSecs()`

UnsetRestoreTimeSecs ensures that no value is present for RestoreTimeSecs, not even an explicit nil
### GetResumeRestore

`func (o *RestoreSqlAppObjectParams) GetResumeRestore() bool`

GetResumeRestore returns the ResumeRestore field if non-nil, zero value otherwise.

### GetResumeRestoreOk

`func (o *RestoreSqlAppObjectParams) GetResumeRestoreOk() (*bool, bool)`

GetResumeRestoreOk returns a tuple with the ResumeRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeRestore

`func (o *RestoreSqlAppObjectParams) SetResumeRestore(v bool)`

SetResumeRestore sets ResumeRestore field to given value.

### HasResumeRestore

`func (o *RestoreSqlAppObjectParams) HasResumeRestore() bool`

HasResumeRestore returns a boolean if a field has been set.

### SetResumeRestoreNil

`func (o *RestoreSqlAppObjectParams) SetResumeRestoreNil(b bool)`

 SetResumeRestoreNil sets the value for ResumeRestore to be an explicit nil

### UnsetResumeRestore
`func (o *RestoreSqlAppObjectParams) UnsetResumeRestore()`

UnsetResumeRestore ensures that no value is present for ResumeRestore, not even an explicit nil
### GetSecondaryDataFileDestination

`func (o *RestoreSqlAppObjectParams) GetSecondaryDataFileDestination() string`

GetSecondaryDataFileDestination returns the SecondaryDataFileDestination field if non-nil, zero value otherwise.

### GetSecondaryDataFileDestinationOk

`func (o *RestoreSqlAppObjectParams) GetSecondaryDataFileDestinationOk() (*string, bool)`

GetSecondaryDataFileDestinationOk returns a tuple with the SecondaryDataFileDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDataFileDestination

`func (o *RestoreSqlAppObjectParams) SetSecondaryDataFileDestination(v string)`

SetSecondaryDataFileDestination sets SecondaryDataFileDestination field to given value.

### HasSecondaryDataFileDestination

`func (o *RestoreSqlAppObjectParams) HasSecondaryDataFileDestination() bool`

HasSecondaryDataFileDestination returns a boolean if a field has been set.

### SetSecondaryDataFileDestinationNil

`func (o *RestoreSqlAppObjectParams) SetSecondaryDataFileDestinationNil(b bool)`

 SetSecondaryDataFileDestinationNil sets the value for SecondaryDataFileDestination to be an explicit nil

### UnsetSecondaryDataFileDestination
`func (o *RestoreSqlAppObjectParams) UnsetSecondaryDataFileDestination()`

UnsetSecondaryDataFileDestination ensures that no value is present for SecondaryDataFileDestination, not even an explicit nil
### GetSecondaryDataFileDestinationVec

`func (o *RestoreSqlAppObjectParams) GetSecondaryDataFileDestinationVec() []FilesToDirectoryMapping`

GetSecondaryDataFileDestinationVec returns the SecondaryDataFileDestinationVec field if non-nil, zero value otherwise.

### GetSecondaryDataFileDestinationVecOk

`func (o *RestoreSqlAppObjectParams) GetSecondaryDataFileDestinationVecOk() (*[]FilesToDirectoryMapping, bool)`

GetSecondaryDataFileDestinationVecOk returns a tuple with the SecondaryDataFileDestinationVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDataFileDestinationVec

`func (o *RestoreSqlAppObjectParams) SetSecondaryDataFileDestinationVec(v []FilesToDirectoryMapping)`

SetSecondaryDataFileDestinationVec sets SecondaryDataFileDestinationVec field to given value.

### HasSecondaryDataFileDestinationVec

`func (o *RestoreSqlAppObjectParams) HasSecondaryDataFileDestinationVec() bool`

HasSecondaryDataFileDestinationVec returns a boolean if a field has been set.

### SetSecondaryDataFileDestinationVecNil

`func (o *RestoreSqlAppObjectParams) SetSecondaryDataFileDestinationVecNil(b bool)`

 SetSecondaryDataFileDestinationVecNil sets the value for SecondaryDataFileDestinationVec to be an explicit nil

### UnsetSecondaryDataFileDestinationVec
`func (o *RestoreSqlAppObjectParams) UnsetSecondaryDataFileDestinationVec()`

UnsetSecondaryDataFileDestinationVec ensures that no value is present for SecondaryDataFileDestinationVec, not even an explicit nil
### GetWithClause

`func (o *RestoreSqlAppObjectParams) GetWithClause() string`

GetWithClause returns the WithClause field if non-nil, zero value otherwise.

### GetWithClauseOk

`func (o *RestoreSqlAppObjectParams) GetWithClauseOk() (*string, bool)`

GetWithClauseOk returns a tuple with the WithClause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithClause

`func (o *RestoreSqlAppObjectParams) SetWithClause(v string)`

SetWithClause sets WithClause field to given value.

### HasWithClause

`func (o *RestoreSqlAppObjectParams) HasWithClause() bool`

HasWithClause returns a boolean if a field has been set.

### SetWithClauseNil

`func (o *RestoreSqlAppObjectParams) SetWithClauseNil(b bool)`

 SetWithClauseNil sets the value for WithClause to be an explicit nil

### UnsetWithClause
`func (o *RestoreSqlAppObjectParams) UnsetWithClause()`

UnsetWithClause ensures that no value is present for WithClause, not even an explicit nil
### GetWithNoRecovery

`func (o *RestoreSqlAppObjectParams) GetWithNoRecovery() bool`

GetWithNoRecovery returns the WithNoRecovery field if non-nil, zero value otherwise.

### GetWithNoRecoveryOk

`func (o *RestoreSqlAppObjectParams) GetWithNoRecoveryOk() (*bool, bool)`

GetWithNoRecoveryOk returns a tuple with the WithNoRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithNoRecovery

`func (o *RestoreSqlAppObjectParams) SetWithNoRecovery(v bool)`

SetWithNoRecovery sets WithNoRecovery field to given value.

### HasWithNoRecovery

`func (o *RestoreSqlAppObjectParams) HasWithNoRecovery() bool`

HasWithNoRecovery returns a boolean if a field has been set.

### SetWithNoRecoveryNil

`func (o *RestoreSqlAppObjectParams) SetWithNoRecoveryNil(b bool)`

 SetWithNoRecoveryNil sets the value for WithNoRecovery to be an explicit nil

### UnsetWithNoRecovery
`func (o *RestoreSqlAppObjectParams) UnsetWithNoRecovery()`

UnsetWithNoRecovery ensures that no value is present for WithNoRecovery, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


