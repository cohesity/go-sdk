# CommonRecoverSqlAppTargetParamsNewSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataFileDirectoryLocation** | **NullableString** | Specifies the directory where to put the database data files. Missing directory will be automatically created. | 
**DatabaseName** | Pointer to **NullableString** | Specifies a new name for the restored database. If this field is not specified, then the original database will be overwritten after recovery. | [optional] 
**Host** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the source id of target host where databases will be recovered. This source id can be a physical host or virtual machine. | 
**InstanceName** | **NullableString** | Specifies an instance name of the Sql Server that should be used for restoring databases to. | 
**LogFileDirectoryLocation** | **NullableString** | Specifies the directory where to put the database log files. Missing directory will be automatically created. | 
**KeepCdc** | Pointer to **NullableBool** | Specifies whether to keep CDC (Change Data Capture) on recovered databases or not. If not passed, this is assumed to be true. If withNoRecovery is passed as true, then this field must not be set to true. Passing this field as true in this scenario will be a invalid request. | [optional] 
**MultiStageRestoreOptions** | Pointer to [**MultiStageRestoreOptions**](MultiStageRestoreOptions.md) |  | [optional] 
**NativeLogRecoveryWithClause** | Pointer to **NullableString** | Specifies the WITH clause to be used in native sql log restore command. This is only applicable for native log restore. | [optional] 
**NativeRecoveryWithClause** | Pointer to **NullableString** | &#39;with_clause&#39; contains &#39;with clause&#39; to be used in native sql restore command. This is only applicable for database restore of native sql backup. Here user can specify multiple restore options. Example: &#39;WITH BUFFERCOUNT &#x3D; 575, MAXTRANSFERSIZE &#x3D; 2097152&#39;. | [optional] 
**OverwritingPolicy** | Pointer to **NullableString** | Specifies a policy to be used while recovering existing databases. | [optional] 
**ReplayEntireLastLog** | Pointer to **NullableBool** | Specifies the option to set replay last log bit while creating the sql restore task and doing restore to latest point-in-time. If this is set to true, we will replay the entire last log without STOPAT. | [optional] 
**RestoreTimeUsecs** | Pointer to **NullableInt64** | Specifies the time in the past to which the Sql database needs to be restored. This allows for granular recovery of Sql databases. If this is not set, the Sql database will be restored from the full/incremental snapshot. | [optional] 
**SecondaryDataFilesDirList** | Pointer to [**[]FilenamePatternToDirectory**](FilenamePatternToDirectory.md) | Specifies the secondary data filename pattern and corresponding direcories of the DB. Secondary data files are optional and are user defined. The recommended file extention for secondary files is \&quot;.ndf\&quot;. If this option is specified and the destination folders do not exist they will be automatically created. | [optional] 
**WithNoRecovery** | Pointer to **NullableBool** | Specifies the flag to bring DBs online or not after successful recovery. If this is passed as true, then it means DBs won&#39;t be brought online. | [optional] 

## Methods

### NewCommonRecoverSqlAppTargetParamsNewSourceConfig

`func NewCommonRecoverSqlAppTargetParamsNewSourceConfig(dataFileDirectoryLocation NullableString, host NullableRecoveryObjectIdentifier, instanceName NullableString, logFileDirectoryLocation NullableString, ) *CommonRecoverSqlAppTargetParamsNewSourceConfig`

NewCommonRecoverSqlAppTargetParamsNewSourceConfig instantiates a new CommonRecoverSqlAppTargetParamsNewSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonRecoverSqlAppTargetParamsNewSourceConfigWithDefaults

`func NewCommonRecoverSqlAppTargetParamsNewSourceConfigWithDefaults() *CommonRecoverSqlAppTargetParamsNewSourceConfig`

NewCommonRecoverSqlAppTargetParamsNewSourceConfigWithDefaults instantiates a new CommonRecoverSqlAppTargetParamsNewSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataFileDirectoryLocation

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetDataFileDirectoryLocation() string`

GetDataFileDirectoryLocation returns the DataFileDirectoryLocation field if non-nil, zero value otherwise.

### GetDataFileDirectoryLocationOk

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetDataFileDirectoryLocationOk() (*string, bool)`

GetDataFileDirectoryLocationOk returns a tuple with the DataFileDirectoryLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFileDirectoryLocation

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetDataFileDirectoryLocation(v string)`

SetDataFileDirectoryLocation sets DataFileDirectoryLocation field to given value.


### SetDataFileDirectoryLocationNil

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetDataFileDirectoryLocationNil(b bool)`

 SetDataFileDirectoryLocationNil sets the value for DataFileDirectoryLocation to be an explicit nil

### UnsetDataFileDirectoryLocation
`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) UnsetDataFileDirectoryLocation()`

UnsetDataFileDirectoryLocation ensures that no value is present for DataFileDirectoryLocation, not even an explicit nil
### GetDatabaseName

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### SetDatabaseNameNil

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetDatabaseNameNil(b bool)`

 SetDatabaseNameNil sets the value for DatabaseName to be an explicit nil

### UnsetDatabaseName
`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) UnsetDatabaseName()`

UnsetDatabaseName ensures that no value is present for DatabaseName, not even an explicit nil
### GetHost

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetHost() RecoveryObjectIdentifier`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetHostOk() (*RecoveryObjectIdentifier, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetHost(v RecoveryObjectIdentifier)`

SetHost sets Host field to given value.


### SetHostNil

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetInstanceName

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.


### SetInstanceNameNil

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetInstanceNameNil(b bool)`

 SetInstanceNameNil sets the value for InstanceName to be an explicit nil

### UnsetInstanceName
`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) UnsetInstanceName()`

UnsetInstanceName ensures that no value is present for InstanceName, not even an explicit nil
### GetLogFileDirectoryLocation

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetLogFileDirectoryLocation() string`

GetLogFileDirectoryLocation returns the LogFileDirectoryLocation field if non-nil, zero value otherwise.

### GetLogFileDirectoryLocationOk

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetLogFileDirectoryLocationOk() (*string, bool)`

GetLogFileDirectoryLocationOk returns a tuple with the LogFileDirectoryLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFileDirectoryLocation

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetLogFileDirectoryLocation(v string)`

SetLogFileDirectoryLocation sets LogFileDirectoryLocation field to given value.


### SetLogFileDirectoryLocationNil

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetLogFileDirectoryLocationNil(b bool)`

 SetLogFileDirectoryLocationNil sets the value for LogFileDirectoryLocation to be an explicit nil

### UnsetLogFileDirectoryLocation
`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) UnsetLogFileDirectoryLocation()`

UnsetLogFileDirectoryLocation ensures that no value is present for LogFileDirectoryLocation, not even an explicit nil
### GetKeepCdc

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetKeepCdc() bool`

GetKeepCdc returns the KeepCdc field if non-nil, zero value otherwise.

### GetKeepCdcOk

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetKeepCdcOk() (*bool, bool)`

GetKeepCdcOk returns a tuple with the KeepCdc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepCdc

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetKeepCdc(v bool)`

SetKeepCdc sets KeepCdc field to given value.

### HasKeepCdc

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) HasKeepCdc() bool`

HasKeepCdc returns a boolean if a field has been set.

### SetKeepCdcNil

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetKeepCdcNil(b bool)`

 SetKeepCdcNil sets the value for KeepCdc to be an explicit nil

### UnsetKeepCdc
`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) UnsetKeepCdc()`

UnsetKeepCdc ensures that no value is present for KeepCdc, not even an explicit nil
### GetMultiStageRestoreOptions

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetMultiStageRestoreOptions() MultiStageRestoreOptions`

GetMultiStageRestoreOptions returns the MultiStageRestoreOptions field if non-nil, zero value otherwise.

### GetMultiStageRestoreOptionsOk

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetMultiStageRestoreOptionsOk() (*MultiStageRestoreOptions, bool)`

GetMultiStageRestoreOptionsOk returns a tuple with the MultiStageRestoreOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiStageRestoreOptions

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetMultiStageRestoreOptions(v MultiStageRestoreOptions)`

SetMultiStageRestoreOptions sets MultiStageRestoreOptions field to given value.

### HasMultiStageRestoreOptions

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) HasMultiStageRestoreOptions() bool`

HasMultiStageRestoreOptions returns a boolean if a field has been set.

### GetNativeLogRecoveryWithClause

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetNativeLogRecoveryWithClause() string`

GetNativeLogRecoveryWithClause returns the NativeLogRecoveryWithClause field if non-nil, zero value otherwise.

### GetNativeLogRecoveryWithClauseOk

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetNativeLogRecoveryWithClauseOk() (*string, bool)`

GetNativeLogRecoveryWithClauseOk returns a tuple with the NativeLogRecoveryWithClause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeLogRecoveryWithClause

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetNativeLogRecoveryWithClause(v string)`

SetNativeLogRecoveryWithClause sets NativeLogRecoveryWithClause field to given value.

### HasNativeLogRecoveryWithClause

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) HasNativeLogRecoveryWithClause() bool`

HasNativeLogRecoveryWithClause returns a boolean if a field has been set.

### SetNativeLogRecoveryWithClauseNil

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetNativeLogRecoveryWithClauseNil(b bool)`

 SetNativeLogRecoveryWithClauseNil sets the value for NativeLogRecoveryWithClause to be an explicit nil

### UnsetNativeLogRecoveryWithClause
`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) UnsetNativeLogRecoveryWithClause()`

UnsetNativeLogRecoveryWithClause ensures that no value is present for NativeLogRecoveryWithClause, not even an explicit nil
### GetNativeRecoveryWithClause

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetNativeRecoveryWithClause() string`

GetNativeRecoveryWithClause returns the NativeRecoveryWithClause field if non-nil, zero value otherwise.

### GetNativeRecoveryWithClauseOk

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetNativeRecoveryWithClauseOk() (*string, bool)`

GetNativeRecoveryWithClauseOk returns a tuple with the NativeRecoveryWithClause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeRecoveryWithClause

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetNativeRecoveryWithClause(v string)`

SetNativeRecoveryWithClause sets NativeRecoveryWithClause field to given value.

### HasNativeRecoveryWithClause

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) HasNativeRecoveryWithClause() bool`

HasNativeRecoveryWithClause returns a boolean if a field has been set.

### SetNativeRecoveryWithClauseNil

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetNativeRecoveryWithClauseNil(b bool)`

 SetNativeRecoveryWithClauseNil sets the value for NativeRecoveryWithClause to be an explicit nil

### UnsetNativeRecoveryWithClause
`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) UnsetNativeRecoveryWithClause()`

UnsetNativeRecoveryWithClause ensures that no value is present for NativeRecoveryWithClause, not even an explicit nil
### GetOverwritingPolicy

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetOverwritingPolicy() string`

GetOverwritingPolicy returns the OverwritingPolicy field if non-nil, zero value otherwise.

### GetOverwritingPolicyOk

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetOverwritingPolicyOk() (*string, bool)`

GetOverwritingPolicyOk returns a tuple with the OverwritingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwritingPolicy

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetOverwritingPolicy(v string)`

SetOverwritingPolicy sets OverwritingPolicy field to given value.

### HasOverwritingPolicy

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) HasOverwritingPolicy() bool`

HasOverwritingPolicy returns a boolean if a field has been set.

### SetOverwritingPolicyNil

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetOverwritingPolicyNil(b bool)`

 SetOverwritingPolicyNil sets the value for OverwritingPolicy to be an explicit nil

### UnsetOverwritingPolicy
`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) UnsetOverwritingPolicy()`

UnsetOverwritingPolicy ensures that no value is present for OverwritingPolicy, not even an explicit nil
### GetReplayEntireLastLog

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetReplayEntireLastLog() bool`

GetReplayEntireLastLog returns the ReplayEntireLastLog field if non-nil, zero value otherwise.

### GetReplayEntireLastLogOk

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetReplayEntireLastLogOk() (*bool, bool)`

GetReplayEntireLastLogOk returns a tuple with the ReplayEntireLastLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayEntireLastLog

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetReplayEntireLastLog(v bool)`

SetReplayEntireLastLog sets ReplayEntireLastLog field to given value.

### HasReplayEntireLastLog

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) HasReplayEntireLastLog() bool`

HasReplayEntireLastLog returns a boolean if a field has been set.

### SetReplayEntireLastLogNil

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetReplayEntireLastLogNil(b bool)`

 SetReplayEntireLastLogNil sets the value for ReplayEntireLastLog to be an explicit nil

### UnsetReplayEntireLastLog
`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) UnsetReplayEntireLastLog()`

UnsetReplayEntireLastLog ensures that no value is present for ReplayEntireLastLog, not even an explicit nil
### GetRestoreTimeUsecs

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetRestoreTimeUsecs() int64`

GetRestoreTimeUsecs returns the RestoreTimeUsecs field if non-nil, zero value otherwise.

### GetRestoreTimeUsecsOk

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetRestoreTimeUsecsOk() (*int64, bool)`

GetRestoreTimeUsecsOk returns a tuple with the RestoreTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTimeUsecs

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetRestoreTimeUsecs(v int64)`

SetRestoreTimeUsecs sets RestoreTimeUsecs field to given value.

### HasRestoreTimeUsecs

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) HasRestoreTimeUsecs() bool`

HasRestoreTimeUsecs returns a boolean if a field has been set.

### SetRestoreTimeUsecsNil

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetRestoreTimeUsecsNil(b bool)`

 SetRestoreTimeUsecsNil sets the value for RestoreTimeUsecs to be an explicit nil

### UnsetRestoreTimeUsecs
`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) UnsetRestoreTimeUsecs()`

UnsetRestoreTimeUsecs ensures that no value is present for RestoreTimeUsecs, not even an explicit nil
### GetSecondaryDataFilesDirList

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetSecondaryDataFilesDirList() []FilenamePatternToDirectory`

GetSecondaryDataFilesDirList returns the SecondaryDataFilesDirList field if non-nil, zero value otherwise.

### GetSecondaryDataFilesDirListOk

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetSecondaryDataFilesDirListOk() (*[]FilenamePatternToDirectory, bool)`

GetSecondaryDataFilesDirListOk returns a tuple with the SecondaryDataFilesDirList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDataFilesDirList

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetSecondaryDataFilesDirList(v []FilenamePatternToDirectory)`

SetSecondaryDataFilesDirList sets SecondaryDataFilesDirList field to given value.

### HasSecondaryDataFilesDirList

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) HasSecondaryDataFilesDirList() bool`

HasSecondaryDataFilesDirList returns a boolean if a field has been set.

### SetSecondaryDataFilesDirListNil

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetSecondaryDataFilesDirListNil(b bool)`

 SetSecondaryDataFilesDirListNil sets the value for SecondaryDataFilesDirList to be an explicit nil

### UnsetSecondaryDataFilesDirList
`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) UnsetSecondaryDataFilesDirList()`

UnsetSecondaryDataFilesDirList ensures that no value is present for SecondaryDataFilesDirList, not even an explicit nil
### GetWithNoRecovery

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetWithNoRecovery() bool`

GetWithNoRecovery returns the WithNoRecovery field if non-nil, zero value otherwise.

### GetWithNoRecoveryOk

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetWithNoRecoveryOk() (*bool, bool)`

GetWithNoRecoveryOk returns a tuple with the WithNoRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithNoRecovery

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetWithNoRecovery(v bool)`

SetWithNoRecovery sets WithNoRecovery field to given value.

### HasWithNoRecovery

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) HasWithNoRecovery() bool`

HasWithNoRecovery returns a boolean if a field has been set.

### SetWithNoRecoveryNil

`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetWithNoRecoveryNil(b bool)`

 SetWithNoRecoveryNil sets the value for WithNoRecovery to be an explicit nil

### UnsetWithNoRecovery
`func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) UnsetWithNoRecovery()`

UnsetWithNoRecovery ensures that no value is present for WithNoRecovery, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


