# RecoverSqlAppOriginalSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataFileDirectoryLocation** | Pointer to **NullableString** | Specifies the directory where to put the database data files. Missing directory will be automatically created. If you are overwriting the existing database then this field will be ignored. | [optional] 
**LogFileDirectoryLocation** | Pointer to **NullableString** | Specifies the directory where to put the database log files. Missing directory will be automatically created. If you are overwriting the existing database then this field will be ignored. | [optional] 
**CaptureTailLogs** | Pointer to **NullableBool** | Set this to true if tail logs are to be captured before the recovery operation. This is only applicable if database is not being renamed. | [optional] 
**NewDatabaseName** | Pointer to **NullableString** | Specifies a new name for the restored database. If this field is not specified, then the original database will be overwritten after recovery. | [optional] 
**RestoreTimeUsecs** | Pointer to **NullableInt64** | Specifies the time in the past to which the Sql database needs to be restored. This allows for granular recovery of Sql databases. If this is not set, the Sql database will be restored from the full/incremental snapshot. | [optional] 
**SecondaryDataFilesDirList** | Pointer to [**[]FilenamePatternToDirectory**](FilenamePatternToDirectory.md) | Specifies the secondary data filename pattern and corresponding direcories of the DB. Secondary data files are optional and are user defined. The recommended file extention for secondary files is \&quot;.ndf\&quot;. If this option is specified and the destination folders do not exist they will be automatically created. | [optional] 
**WithNoRecovery** | Pointer to **NullableBool** | Specifies the flag to bring DBs online or not after successful recovery. If this is passed as true, then it means DBs won&#39;t be brought online. | [optional] 
**KeepCdc** | Pointer to **NullableBool** | Specifies whether to keep CDC (Change Data Capture) on recovered databases or not. If not passed, this is assumed to be true. If withNoRecovery is passed as true, then this field must not be set to true. Passing this field as true in this scenario will be a invalid request. | [optional] 
**OverwritingPolicy** | Pointer to **NullableString** | Specifies a policy to be used while recovering existing databases. | [optional] 
**MultiStageRestoreOptions** | Pointer to [**MultiStageRestoreOptions**](MultiStageRestoreOptions.md) |  | [optional] 
**NativeRecoveryWithClause** | Pointer to **NullableString** | &#39;with_clause&#39; contains &#39;with clause&#39; to be used in native sql restore command. This is only applicable for database restore of native sql backup. Here user can specify multiple restore options. Example: &#39;WITH BUFFERCOUNT &#x3D; 575, MAXTRANSFERSIZE &#x3D; 2097152&#39;. | [optional] 

## Methods

### NewRecoverSqlAppOriginalSourceConfig

`func NewRecoverSqlAppOriginalSourceConfig() *RecoverSqlAppOriginalSourceConfig`

NewRecoverSqlAppOriginalSourceConfig instantiates a new RecoverSqlAppOriginalSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverSqlAppOriginalSourceConfigWithDefaults

`func NewRecoverSqlAppOriginalSourceConfigWithDefaults() *RecoverSqlAppOriginalSourceConfig`

NewRecoverSqlAppOriginalSourceConfigWithDefaults instantiates a new RecoverSqlAppOriginalSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataFileDirectoryLocation

`func (o *RecoverSqlAppOriginalSourceConfig) GetDataFileDirectoryLocation() string`

GetDataFileDirectoryLocation returns the DataFileDirectoryLocation field if non-nil, zero value otherwise.

### GetDataFileDirectoryLocationOk

`func (o *RecoverSqlAppOriginalSourceConfig) GetDataFileDirectoryLocationOk() (*string, bool)`

GetDataFileDirectoryLocationOk returns a tuple with the DataFileDirectoryLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFileDirectoryLocation

`func (o *RecoverSqlAppOriginalSourceConfig) SetDataFileDirectoryLocation(v string)`

SetDataFileDirectoryLocation sets DataFileDirectoryLocation field to given value.

### HasDataFileDirectoryLocation

`func (o *RecoverSqlAppOriginalSourceConfig) HasDataFileDirectoryLocation() bool`

HasDataFileDirectoryLocation returns a boolean if a field has been set.

### SetDataFileDirectoryLocationNil

`func (o *RecoverSqlAppOriginalSourceConfig) SetDataFileDirectoryLocationNil(b bool)`

 SetDataFileDirectoryLocationNil sets the value for DataFileDirectoryLocation to be an explicit nil

### UnsetDataFileDirectoryLocation
`func (o *RecoverSqlAppOriginalSourceConfig) UnsetDataFileDirectoryLocation()`

UnsetDataFileDirectoryLocation ensures that no value is present for DataFileDirectoryLocation, not even an explicit nil
### GetLogFileDirectoryLocation

`func (o *RecoverSqlAppOriginalSourceConfig) GetLogFileDirectoryLocation() string`

GetLogFileDirectoryLocation returns the LogFileDirectoryLocation field if non-nil, zero value otherwise.

### GetLogFileDirectoryLocationOk

`func (o *RecoverSqlAppOriginalSourceConfig) GetLogFileDirectoryLocationOk() (*string, bool)`

GetLogFileDirectoryLocationOk returns a tuple with the LogFileDirectoryLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFileDirectoryLocation

`func (o *RecoverSqlAppOriginalSourceConfig) SetLogFileDirectoryLocation(v string)`

SetLogFileDirectoryLocation sets LogFileDirectoryLocation field to given value.

### HasLogFileDirectoryLocation

`func (o *RecoverSqlAppOriginalSourceConfig) HasLogFileDirectoryLocation() bool`

HasLogFileDirectoryLocation returns a boolean if a field has been set.

### SetLogFileDirectoryLocationNil

`func (o *RecoverSqlAppOriginalSourceConfig) SetLogFileDirectoryLocationNil(b bool)`

 SetLogFileDirectoryLocationNil sets the value for LogFileDirectoryLocation to be an explicit nil

### UnsetLogFileDirectoryLocation
`func (o *RecoverSqlAppOriginalSourceConfig) UnsetLogFileDirectoryLocation()`

UnsetLogFileDirectoryLocation ensures that no value is present for LogFileDirectoryLocation, not even an explicit nil
### GetCaptureTailLogs

`func (o *RecoverSqlAppOriginalSourceConfig) GetCaptureTailLogs() bool`

GetCaptureTailLogs returns the CaptureTailLogs field if non-nil, zero value otherwise.

### GetCaptureTailLogsOk

`func (o *RecoverSqlAppOriginalSourceConfig) GetCaptureTailLogsOk() (*bool, bool)`

GetCaptureTailLogsOk returns a tuple with the CaptureTailLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureTailLogs

`func (o *RecoverSqlAppOriginalSourceConfig) SetCaptureTailLogs(v bool)`

SetCaptureTailLogs sets CaptureTailLogs field to given value.

### HasCaptureTailLogs

`func (o *RecoverSqlAppOriginalSourceConfig) HasCaptureTailLogs() bool`

HasCaptureTailLogs returns a boolean if a field has been set.

### SetCaptureTailLogsNil

`func (o *RecoverSqlAppOriginalSourceConfig) SetCaptureTailLogsNil(b bool)`

 SetCaptureTailLogsNil sets the value for CaptureTailLogs to be an explicit nil

### UnsetCaptureTailLogs
`func (o *RecoverSqlAppOriginalSourceConfig) UnsetCaptureTailLogs()`

UnsetCaptureTailLogs ensures that no value is present for CaptureTailLogs, not even an explicit nil
### GetNewDatabaseName

`func (o *RecoverSqlAppOriginalSourceConfig) GetNewDatabaseName() string`

GetNewDatabaseName returns the NewDatabaseName field if non-nil, zero value otherwise.

### GetNewDatabaseNameOk

`func (o *RecoverSqlAppOriginalSourceConfig) GetNewDatabaseNameOk() (*string, bool)`

GetNewDatabaseNameOk returns a tuple with the NewDatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDatabaseName

`func (o *RecoverSqlAppOriginalSourceConfig) SetNewDatabaseName(v string)`

SetNewDatabaseName sets NewDatabaseName field to given value.

### HasNewDatabaseName

`func (o *RecoverSqlAppOriginalSourceConfig) HasNewDatabaseName() bool`

HasNewDatabaseName returns a boolean if a field has been set.

### SetNewDatabaseNameNil

`func (o *RecoverSqlAppOriginalSourceConfig) SetNewDatabaseNameNil(b bool)`

 SetNewDatabaseNameNil sets the value for NewDatabaseName to be an explicit nil

### UnsetNewDatabaseName
`func (o *RecoverSqlAppOriginalSourceConfig) UnsetNewDatabaseName()`

UnsetNewDatabaseName ensures that no value is present for NewDatabaseName, not even an explicit nil
### GetRestoreTimeUsecs

`func (o *RecoverSqlAppOriginalSourceConfig) GetRestoreTimeUsecs() int64`

GetRestoreTimeUsecs returns the RestoreTimeUsecs field if non-nil, zero value otherwise.

### GetRestoreTimeUsecsOk

`func (o *RecoverSqlAppOriginalSourceConfig) GetRestoreTimeUsecsOk() (*int64, bool)`

GetRestoreTimeUsecsOk returns a tuple with the RestoreTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTimeUsecs

`func (o *RecoverSqlAppOriginalSourceConfig) SetRestoreTimeUsecs(v int64)`

SetRestoreTimeUsecs sets RestoreTimeUsecs field to given value.

### HasRestoreTimeUsecs

`func (o *RecoverSqlAppOriginalSourceConfig) HasRestoreTimeUsecs() bool`

HasRestoreTimeUsecs returns a boolean if a field has been set.

### SetRestoreTimeUsecsNil

`func (o *RecoverSqlAppOriginalSourceConfig) SetRestoreTimeUsecsNil(b bool)`

 SetRestoreTimeUsecsNil sets the value for RestoreTimeUsecs to be an explicit nil

### UnsetRestoreTimeUsecs
`func (o *RecoverSqlAppOriginalSourceConfig) UnsetRestoreTimeUsecs()`

UnsetRestoreTimeUsecs ensures that no value is present for RestoreTimeUsecs, not even an explicit nil
### GetSecondaryDataFilesDirList

`func (o *RecoverSqlAppOriginalSourceConfig) GetSecondaryDataFilesDirList() []FilenamePatternToDirectory`

GetSecondaryDataFilesDirList returns the SecondaryDataFilesDirList field if non-nil, zero value otherwise.

### GetSecondaryDataFilesDirListOk

`func (o *RecoverSqlAppOriginalSourceConfig) GetSecondaryDataFilesDirListOk() (*[]FilenamePatternToDirectory, bool)`

GetSecondaryDataFilesDirListOk returns a tuple with the SecondaryDataFilesDirList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDataFilesDirList

`func (o *RecoverSqlAppOriginalSourceConfig) SetSecondaryDataFilesDirList(v []FilenamePatternToDirectory)`

SetSecondaryDataFilesDirList sets SecondaryDataFilesDirList field to given value.

### HasSecondaryDataFilesDirList

`func (o *RecoverSqlAppOriginalSourceConfig) HasSecondaryDataFilesDirList() bool`

HasSecondaryDataFilesDirList returns a boolean if a field has been set.

### SetSecondaryDataFilesDirListNil

`func (o *RecoverSqlAppOriginalSourceConfig) SetSecondaryDataFilesDirListNil(b bool)`

 SetSecondaryDataFilesDirListNil sets the value for SecondaryDataFilesDirList to be an explicit nil

### UnsetSecondaryDataFilesDirList
`func (o *RecoverSqlAppOriginalSourceConfig) UnsetSecondaryDataFilesDirList()`

UnsetSecondaryDataFilesDirList ensures that no value is present for SecondaryDataFilesDirList, not even an explicit nil
### GetWithNoRecovery

`func (o *RecoverSqlAppOriginalSourceConfig) GetWithNoRecovery() bool`

GetWithNoRecovery returns the WithNoRecovery field if non-nil, zero value otherwise.

### GetWithNoRecoveryOk

`func (o *RecoverSqlAppOriginalSourceConfig) GetWithNoRecoveryOk() (*bool, bool)`

GetWithNoRecoveryOk returns a tuple with the WithNoRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithNoRecovery

`func (o *RecoverSqlAppOriginalSourceConfig) SetWithNoRecovery(v bool)`

SetWithNoRecovery sets WithNoRecovery field to given value.

### HasWithNoRecovery

`func (o *RecoverSqlAppOriginalSourceConfig) HasWithNoRecovery() bool`

HasWithNoRecovery returns a boolean if a field has been set.

### SetWithNoRecoveryNil

`func (o *RecoverSqlAppOriginalSourceConfig) SetWithNoRecoveryNil(b bool)`

 SetWithNoRecoveryNil sets the value for WithNoRecovery to be an explicit nil

### UnsetWithNoRecovery
`func (o *RecoverSqlAppOriginalSourceConfig) UnsetWithNoRecovery()`

UnsetWithNoRecovery ensures that no value is present for WithNoRecovery, not even an explicit nil
### GetKeepCdc

`func (o *RecoverSqlAppOriginalSourceConfig) GetKeepCdc() bool`

GetKeepCdc returns the KeepCdc field if non-nil, zero value otherwise.

### GetKeepCdcOk

`func (o *RecoverSqlAppOriginalSourceConfig) GetKeepCdcOk() (*bool, bool)`

GetKeepCdcOk returns a tuple with the KeepCdc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepCdc

`func (o *RecoverSqlAppOriginalSourceConfig) SetKeepCdc(v bool)`

SetKeepCdc sets KeepCdc field to given value.

### HasKeepCdc

`func (o *RecoverSqlAppOriginalSourceConfig) HasKeepCdc() bool`

HasKeepCdc returns a boolean if a field has been set.

### SetKeepCdcNil

`func (o *RecoverSqlAppOriginalSourceConfig) SetKeepCdcNil(b bool)`

 SetKeepCdcNil sets the value for KeepCdc to be an explicit nil

### UnsetKeepCdc
`func (o *RecoverSqlAppOriginalSourceConfig) UnsetKeepCdc()`

UnsetKeepCdc ensures that no value is present for KeepCdc, not even an explicit nil
### GetOverwritingPolicy

`func (o *RecoverSqlAppOriginalSourceConfig) GetOverwritingPolicy() string`

GetOverwritingPolicy returns the OverwritingPolicy field if non-nil, zero value otherwise.

### GetOverwritingPolicyOk

`func (o *RecoverSqlAppOriginalSourceConfig) GetOverwritingPolicyOk() (*string, bool)`

GetOverwritingPolicyOk returns a tuple with the OverwritingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwritingPolicy

`func (o *RecoverSqlAppOriginalSourceConfig) SetOverwritingPolicy(v string)`

SetOverwritingPolicy sets OverwritingPolicy field to given value.

### HasOverwritingPolicy

`func (o *RecoverSqlAppOriginalSourceConfig) HasOverwritingPolicy() bool`

HasOverwritingPolicy returns a boolean if a field has been set.

### SetOverwritingPolicyNil

`func (o *RecoverSqlAppOriginalSourceConfig) SetOverwritingPolicyNil(b bool)`

 SetOverwritingPolicyNil sets the value for OverwritingPolicy to be an explicit nil

### UnsetOverwritingPolicy
`func (o *RecoverSqlAppOriginalSourceConfig) UnsetOverwritingPolicy()`

UnsetOverwritingPolicy ensures that no value is present for OverwritingPolicy, not even an explicit nil
### GetMultiStageRestoreOptions

`func (o *RecoverSqlAppOriginalSourceConfig) GetMultiStageRestoreOptions() MultiStageRestoreOptions`

GetMultiStageRestoreOptions returns the MultiStageRestoreOptions field if non-nil, zero value otherwise.

### GetMultiStageRestoreOptionsOk

`func (o *RecoverSqlAppOriginalSourceConfig) GetMultiStageRestoreOptionsOk() (*MultiStageRestoreOptions, bool)`

GetMultiStageRestoreOptionsOk returns a tuple with the MultiStageRestoreOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiStageRestoreOptions

`func (o *RecoverSqlAppOriginalSourceConfig) SetMultiStageRestoreOptions(v MultiStageRestoreOptions)`

SetMultiStageRestoreOptions sets MultiStageRestoreOptions field to given value.

### HasMultiStageRestoreOptions

`func (o *RecoverSqlAppOriginalSourceConfig) HasMultiStageRestoreOptions() bool`

HasMultiStageRestoreOptions returns a boolean if a field has been set.

### GetNativeRecoveryWithClause

`func (o *RecoverSqlAppOriginalSourceConfig) GetNativeRecoveryWithClause() string`

GetNativeRecoveryWithClause returns the NativeRecoveryWithClause field if non-nil, zero value otherwise.

### GetNativeRecoveryWithClauseOk

`func (o *RecoverSqlAppOriginalSourceConfig) GetNativeRecoveryWithClauseOk() (*string, bool)`

GetNativeRecoveryWithClauseOk returns a tuple with the NativeRecoveryWithClause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeRecoveryWithClause

`func (o *RecoverSqlAppOriginalSourceConfig) SetNativeRecoveryWithClause(v string)`

SetNativeRecoveryWithClause sets NativeRecoveryWithClause field to given value.

### HasNativeRecoveryWithClause

`func (o *RecoverSqlAppOriginalSourceConfig) HasNativeRecoveryWithClause() bool`

HasNativeRecoveryWithClause returns a boolean if a field has been set.

### SetNativeRecoveryWithClauseNil

`func (o *RecoverSqlAppOriginalSourceConfig) SetNativeRecoveryWithClauseNil(b bool)`

 SetNativeRecoveryWithClauseNil sets the value for NativeRecoveryWithClause to be an explicit nil

### UnsetNativeRecoveryWithClause
`func (o *RecoverSqlAppOriginalSourceConfig) UnsetNativeRecoveryWithClause()`

UnsetNativeRecoveryWithClause ensures that no value is present for NativeRecoveryWithClause, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


