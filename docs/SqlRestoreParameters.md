# SqlRestoreParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaptureTailLogs** | Pointer to **NullableBool** | Set this to true if tail logs are to be captured before the restore operation. This is only applicable if we are restoring the SQL database to its hosting Protection Source, and the database is not being renamed. | [optional] 
**IsAutoSyncEnabled** | Pointer to **NullableBool** | This field determines if Auto Sync enabled/disabled for SQL Multi-stage Restore task | [optional] 
**KeepCdc** | Pointer to **NullableBool** | This field prevents \&quot;change data capture\&quot; settings from being reomved when a database or log backup is restored on another server and database is recovered. | [optional] 
**KeepOffline** | Pointer to **NullableBool** | Set this to true if we want to restore the database and do not want to bring it online after restore.  This is only applicable if we are restoring the database back to its original location. | [optional] 
**NewDatabaseName** | Pointer to **NullableString** | Specifies optionally a new name for the restored database. | [optional] 
**NewInstanceName** | Pointer to **NullableString** | Specifies an instance name of the SQL Server that should be restored. SQL application has many instances. Each instance has a unique name. One of the instances that should be restored must be set in this field. | [optional] 
**RestoreTimeSecs** | Pointer to **NullableInt64** | Specifies the time in the past to which the SQL database needs to be restored. This allows for granular recovery of SQL databases. If this is not set, the SQL database will be restored from the full/incremental snapshot. | [optional] 
**TargetDataFilesDirectory** | Pointer to **NullableString** | Specifies the directory where to put the database data files. Missing directory will be automatically created. This field must be set if restoring to a different target host. | [optional] 
**TargetLogFilesDirectory** | Pointer to **NullableString** | Specifies the directory where to put the database log files. Missing directory will be automatically created. This field must be set if restoring to a different target host. | [optional] 
**TargetSecondaryDataFilesDirectoryList** | Pointer to [**[]FilenamePatternToDirectory**](FilenamePatternToDirectory.md) | Specifies the secondary data filename pattern and corresponding directories of the DB. Secondary data files are optional and are user defined. The recommended file extension for secondary files is \&quot;.ndf\&quot;.  If this option is specified and the destination folders do not exist they will be automatically created. | [optional] 
**WithClause** | Pointer to **NullableString** | WithClause allows you to specify clauses to be used in native sql restore task. | [optional] 

## Methods

### NewSqlRestoreParameters

`func NewSqlRestoreParameters() *SqlRestoreParameters`

NewSqlRestoreParameters instantiates a new SqlRestoreParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlRestoreParametersWithDefaults

`func NewSqlRestoreParametersWithDefaults() *SqlRestoreParameters`

NewSqlRestoreParametersWithDefaults instantiates a new SqlRestoreParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaptureTailLogs

`func (o *SqlRestoreParameters) GetCaptureTailLogs() bool`

GetCaptureTailLogs returns the CaptureTailLogs field if non-nil, zero value otherwise.

### GetCaptureTailLogsOk

`func (o *SqlRestoreParameters) GetCaptureTailLogsOk() (*bool, bool)`

GetCaptureTailLogsOk returns a tuple with the CaptureTailLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureTailLogs

`func (o *SqlRestoreParameters) SetCaptureTailLogs(v bool)`

SetCaptureTailLogs sets CaptureTailLogs field to given value.

### HasCaptureTailLogs

`func (o *SqlRestoreParameters) HasCaptureTailLogs() bool`

HasCaptureTailLogs returns a boolean if a field has been set.

### SetCaptureTailLogsNil

`func (o *SqlRestoreParameters) SetCaptureTailLogsNil(b bool)`

 SetCaptureTailLogsNil sets the value for CaptureTailLogs to be an explicit nil

### UnsetCaptureTailLogs
`func (o *SqlRestoreParameters) UnsetCaptureTailLogs()`

UnsetCaptureTailLogs ensures that no value is present for CaptureTailLogs, not even an explicit nil
### GetIsAutoSyncEnabled

`func (o *SqlRestoreParameters) GetIsAutoSyncEnabled() bool`

GetIsAutoSyncEnabled returns the IsAutoSyncEnabled field if non-nil, zero value otherwise.

### GetIsAutoSyncEnabledOk

`func (o *SqlRestoreParameters) GetIsAutoSyncEnabledOk() (*bool, bool)`

GetIsAutoSyncEnabledOk returns a tuple with the IsAutoSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoSyncEnabled

`func (o *SqlRestoreParameters) SetIsAutoSyncEnabled(v bool)`

SetIsAutoSyncEnabled sets IsAutoSyncEnabled field to given value.

### HasIsAutoSyncEnabled

`func (o *SqlRestoreParameters) HasIsAutoSyncEnabled() bool`

HasIsAutoSyncEnabled returns a boolean if a field has been set.

### SetIsAutoSyncEnabledNil

`func (o *SqlRestoreParameters) SetIsAutoSyncEnabledNil(b bool)`

 SetIsAutoSyncEnabledNil sets the value for IsAutoSyncEnabled to be an explicit nil

### UnsetIsAutoSyncEnabled
`func (o *SqlRestoreParameters) UnsetIsAutoSyncEnabled()`

UnsetIsAutoSyncEnabled ensures that no value is present for IsAutoSyncEnabled, not even an explicit nil
### GetKeepCdc

`func (o *SqlRestoreParameters) GetKeepCdc() bool`

GetKeepCdc returns the KeepCdc field if non-nil, zero value otherwise.

### GetKeepCdcOk

`func (o *SqlRestoreParameters) GetKeepCdcOk() (*bool, bool)`

GetKeepCdcOk returns a tuple with the KeepCdc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepCdc

`func (o *SqlRestoreParameters) SetKeepCdc(v bool)`

SetKeepCdc sets KeepCdc field to given value.

### HasKeepCdc

`func (o *SqlRestoreParameters) HasKeepCdc() bool`

HasKeepCdc returns a boolean if a field has been set.

### SetKeepCdcNil

`func (o *SqlRestoreParameters) SetKeepCdcNil(b bool)`

 SetKeepCdcNil sets the value for KeepCdc to be an explicit nil

### UnsetKeepCdc
`func (o *SqlRestoreParameters) UnsetKeepCdc()`

UnsetKeepCdc ensures that no value is present for KeepCdc, not even an explicit nil
### GetKeepOffline

`func (o *SqlRestoreParameters) GetKeepOffline() bool`

GetKeepOffline returns the KeepOffline field if non-nil, zero value otherwise.

### GetKeepOfflineOk

`func (o *SqlRestoreParameters) GetKeepOfflineOk() (*bool, bool)`

GetKeepOfflineOk returns a tuple with the KeepOffline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepOffline

`func (o *SqlRestoreParameters) SetKeepOffline(v bool)`

SetKeepOffline sets KeepOffline field to given value.

### HasKeepOffline

`func (o *SqlRestoreParameters) HasKeepOffline() bool`

HasKeepOffline returns a boolean if a field has been set.

### SetKeepOfflineNil

`func (o *SqlRestoreParameters) SetKeepOfflineNil(b bool)`

 SetKeepOfflineNil sets the value for KeepOffline to be an explicit nil

### UnsetKeepOffline
`func (o *SqlRestoreParameters) UnsetKeepOffline()`

UnsetKeepOffline ensures that no value is present for KeepOffline, not even an explicit nil
### GetNewDatabaseName

`func (o *SqlRestoreParameters) GetNewDatabaseName() string`

GetNewDatabaseName returns the NewDatabaseName field if non-nil, zero value otherwise.

### GetNewDatabaseNameOk

`func (o *SqlRestoreParameters) GetNewDatabaseNameOk() (*string, bool)`

GetNewDatabaseNameOk returns a tuple with the NewDatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDatabaseName

`func (o *SqlRestoreParameters) SetNewDatabaseName(v string)`

SetNewDatabaseName sets NewDatabaseName field to given value.

### HasNewDatabaseName

`func (o *SqlRestoreParameters) HasNewDatabaseName() bool`

HasNewDatabaseName returns a boolean if a field has been set.

### SetNewDatabaseNameNil

`func (o *SqlRestoreParameters) SetNewDatabaseNameNil(b bool)`

 SetNewDatabaseNameNil sets the value for NewDatabaseName to be an explicit nil

### UnsetNewDatabaseName
`func (o *SqlRestoreParameters) UnsetNewDatabaseName()`

UnsetNewDatabaseName ensures that no value is present for NewDatabaseName, not even an explicit nil
### GetNewInstanceName

`func (o *SqlRestoreParameters) GetNewInstanceName() string`

GetNewInstanceName returns the NewInstanceName field if non-nil, zero value otherwise.

### GetNewInstanceNameOk

`func (o *SqlRestoreParameters) GetNewInstanceNameOk() (*string, bool)`

GetNewInstanceNameOk returns a tuple with the NewInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewInstanceName

`func (o *SqlRestoreParameters) SetNewInstanceName(v string)`

SetNewInstanceName sets NewInstanceName field to given value.

### HasNewInstanceName

`func (o *SqlRestoreParameters) HasNewInstanceName() bool`

HasNewInstanceName returns a boolean if a field has been set.

### SetNewInstanceNameNil

`func (o *SqlRestoreParameters) SetNewInstanceNameNil(b bool)`

 SetNewInstanceNameNil sets the value for NewInstanceName to be an explicit nil

### UnsetNewInstanceName
`func (o *SqlRestoreParameters) UnsetNewInstanceName()`

UnsetNewInstanceName ensures that no value is present for NewInstanceName, not even an explicit nil
### GetRestoreTimeSecs

`func (o *SqlRestoreParameters) GetRestoreTimeSecs() int64`

GetRestoreTimeSecs returns the RestoreTimeSecs field if non-nil, zero value otherwise.

### GetRestoreTimeSecsOk

`func (o *SqlRestoreParameters) GetRestoreTimeSecsOk() (*int64, bool)`

GetRestoreTimeSecsOk returns a tuple with the RestoreTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTimeSecs

`func (o *SqlRestoreParameters) SetRestoreTimeSecs(v int64)`

SetRestoreTimeSecs sets RestoreTimeSecs field to given value.

### HasRestoreTimeSecs

`func (o *SqlRestoreParameters) HasRestoreTimeSecs() bool`

HasRestoreTimeSecs returns a boolean if a field has been set.

### SetRestoreTimeSecsNil

`func (o *SqlRestoreParameters) SetRestoreTimeSecsNil(b bool)`

 SetRestoreTimeSecsNil sets the value for RestoreTimeSecs to be an explicit nil

### UnsetRestoreTimeSecs
`func (o *SqlRestoreParameters) UnsetRestoreTimeSecs()`

UnsetRestoreTimeSecs ensures that no value is present for RestoreTimeSecs, not even an explicit nil
### GetTargetDataFilesDirectory

`func (o *SqlRestoreParameters) GetTargetDataFilesDirectory() string`

GetTargetDataFilesDirectory returns the TargetDataFilesDirectory field if non-nil, zero value otherwise.

### GetTargetDataFilesDirectoryOk

`func (o *SqlRestoreParameters) GetTargetDataFilesDirectoryOk() (*string, bool)`

GetTargetDataFilesDirectoryOk returns a tuple with the TargetDataFilesDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDataFilesDirectory

`func (o *SqlRestoreParameters) SetTargetDataFilesDirectory(v string)`

SetTargetDataFilesDirectory sets TargetDataFilesDirectory field to given value.

### HasTargetDataFilesDirectory

`func (o *SqlRestoreParameters) HasTargetDataFilesDirectory() bool`

HasTargetDataFilesDirectory returns a boolean if a field has been set.

### SetTargetDataFilesDirectoryNil

`func (o *SqlRestoreParameters) SetTargetDataFilesDirectoryNil(b bool)`

 SetTargetDataFilesDirectoryNil sets the value for TargetDataFilesDirectory to be an explicit nil

### UnsetTargetDataFilesDirectory
`func (o *SqlRestoreParameters) UnsetTargetDataFilesDirectory()`

UnsetTargetDataFilesDirectory ensures that no value is present for TargetDataFilesDirectory, not even an explicit nil
### GetTargetLogFilesDirectory

`func (o *SqlRestoreParameters) GetTargetLogFilesDirectory() string`

GetTargetLogFilesDirectory returns the TargetLogFilesDirectory field if non-nil, zero value otherwise.

### GetTargetLogFilesDirectoryOk

`func (o *SqlRestoreParameters) GetTargetLogFilesDirectoryOk() (*string, bool)`

GetTargetLogFilesDirectoryOk returns a tuple with the TargetLogFilesDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLogFilesDirectory

`func (o *SqlRestoreParameters) SetTargetLogFilesDirectory(v string)`

SetTargetLogFilesDirectory sets TargetLogFilesDirectory field to given value.

### HasTargetLogFilesDirectory

`func (o *SqlRestoreParameters) HasTargetLogFilesDirectory() bool`

HasTargetLogFilesDirectory returns a boolean if a field has been set.

### SetTargetLogFilesDirectoryNil

`func (o *SqlRestoreParameters) SetTargetLogFilesDirectoryNil(b bool)`

 SetTargetLogFilesDirectoryNil sets the value for TargetLogFilesDirectory to be an explicit nil

### UnsetTargetLogFilesDirectory
`func (o *SqlRestoreParameters) UnsetTargetLogFilesDirectory()`

UnsetTargetLogFilesDirectory ensures that no value is present for TargetLogFilesDirectory, not even an explicit nil
### GetTargetSecondaryDataFilesDirectoryList

`func (o *SqlRestoreParameters) GetTargetSecondaryDataFilesDirectoryList() []FilenamePatternToDirectory`

GetTargetSecondaryDataFilesDirectoryList returns the TargetSecondaryDataFilesDirectoryList field if non-nil, zero value otherwise.

### GetTargetSecondaryDataFilesDirectoryListOk

`func (o *SqlRestoreParameters) GetTargetSecondaryDataFilesDirectoryListOk() (*[]FilenamePatternToDirectory, bool)`

GetTargetSecondaryDataFilesDirectoryListOk returns a tuple with the TargetSecondaryDataFilesDirectoryList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSecondaryDataFilesDirectoryList

`func (o *SqlRestoreParameters) SetTargetSecondaryDataFilesDirectoryList(v []FilenamePatternToDirectory)`

SetTargetSecondaryDataFilesDirectoryList sets TargetSecondaryDataFilesDirectoryList field to given value.

### HasTargetSecondaryDataFilesDirectoryList

`func (o *SqlRestoreParameters) HasTargetSecondaryDataFilesDirectoryList() bool`

HasTargetSecondaryDataFilesDirectoryList returns a boolean if a field has been set.

### SetTargetSecondaryDataFilesDirectoryListNil

`func (o *SqlRestoreParameters) SetTargetSecondaryDataFilesDirectoryListNil(b bool)`

 SetTargetSecondaryDataFilesDirectoryListNil sets the value for TargetSecondaryDataFilesDirectoryList to be an explicit nil

### UnsetTargetSecondaryDataFilesDirectoryList
`func (o *SqlRestoreParameters) UnsetTargetSecondaryDataFilesDirectoryList()`

UnsetTargetSecondaryDataFilesDirectoryList ensures that no value is present for TargetSecondaryDataFilesDirectoryList, not even an explicit nil
### GetWithClause

`func (o *SqlRestoreParameters) GetWithClause() string`

GetWithClause returns the WithClause field if non-nil, zero value otherwise.

### GetWithClauseOk

`func (o *SqlRestoreParameters) GetWithClauseOk() (*string, bool)`

GetWithClauseOk returns a tuple with the WithClause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithClause

`func (o *SqlRestoreParameters) SetWithClause(v string)`

SetWithClause sets WithClause field to given value.

### HasWithClause

`func (o *SqlRestoreParameters) HasWithClause() bool`

HasWithClause returns a boolean if a field has been set.

### SetWithClauseNil

`func (o *SqlRestoreParameters) SetWithClauseNil(b bool)`

 SetWithClauseNil sets the value for WithClause to be an explicit nil

### UnsetWithClause
`func (o *SqlRestoreParameters) UnsetWithClause()`

UnsetWithClause ensures that no value is present for WithClause, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


