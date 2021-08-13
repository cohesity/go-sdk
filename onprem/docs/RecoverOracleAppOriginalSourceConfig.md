# RecoverOracleAppOriginalSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RollForwardLogPathVec** | Pointer to **[]string** | List of archive logs to apply on Database after overwrite restore. | [optional] 
**AttemptCompleteRecovery** | Pointer to **NullableBool** | Whether or not this is a complete recovery attempt. | [optional] 
**RestoreTimeUsecs** | Pointer to **NullableInt64** | Specifies the time in the past to which the Oracle db needs to be restored. This allows for granular recovery of Oracle databases. If this is not set, the Oracle db will be restored from the full/incremental snapshot. | [optional] 
**DbChannels** | Pointer to [**[]OracleDbChannel**](OracleDbChannel.md) | Specifies the Oracle database node channels info. If not specified, the default values assigned by the server are applied to all the databases. | [optional] 
**RecoveryMode** | Pointer to **NullableBool** | Specifies if database should be left in recovery mode. | [optional] 
**ShellEvironmentVars** | Pointer to [**[]ShellKeyValuePair**](ShellKeyValuePair.md) | Specifies key value pairs of shell variables which defines the restore shell environment. | [optional] 
**GranularRestoreInfo** | Pointer to [**RecoverOracleGranularRestoreInfo**](RecoverOracleGranularRestoreInfo.md) | Specifies information about list of objects (PDBs) to restore. | [optional] 
**OracleArchiveLogInfo** | Pointer to [**OracleArchiveLogInfo**](OracleArchiveLogInfo.md) | Specifies Range in Time, Scn or Sequence to restore archive logs of a DB. | [optional] 

## Methods

### NewRecoverOracleAppOriginalSourceConfig

`func NewRecoverOracleAppOriginalSourceConfig() *RecoverOracleAppOriginalSourceConfig`

NewRecoverOracleAppOriginalSourceConfig instantiates a new RecoverOracleAppOriginalSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverOracleAppOriginalSourceConfigWithDefaults

`func NewRecoverOracleAppOriginalSourceConfigWithDefaults() *RecoverOracleAppOriginalSourceConfig`

NewRecoverOracleAppOriginalSourceConfigWithDefaults instantiates a new RecoverOracleAppOriginalSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRollForwardLogPathVec

`func (o *RecoverOracleAppOriginalSourceConfig) GetRollForwardLogPathVec() []string`

GetRollForwardLogPathVec returns the RollForwardLogPathVec field if non-nil, zero value otherwise.

### GetRollForwardLogPathVecOk

`func (o *RecoverOracleAppOriginalSourceConfig) GetRollForwardLogPathVecOk() (*[]string, bool)`

GetRollForwardLogPathVecOk returns a tuple with the RollForwardLogPathVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollForwardLogPathVec

`func (o *RecoverOracleAppOriginalSourceConfig) SetRollForwardLogPathVec(v []string)`

SetRollForwardLogPathVec sets RollForwardLogPathVec field to given value.

### HasRollForwardLogPathVec

`func (o *RecoverOracleAppOriginalSourceConfig) HasRollForwardLogPathVec() bool`

HasRollForwardLogPathVec returns a boolean if a field has been set.

### SetRollForwardLogPathVecNil

`func (o *RecoverOracleAppOriginalSourceConfig) SetRollForwardLogPathVecNil(b bool)`

 SetRollForwardLogPathVecNil sets the value for RollForwardLogPathVec to be an explicit nil

### UnsetRollForwardLogPathVec
`func (o *RecoverOracleAppOriginalSourceConfig) UnsetRollForwardLogPathVec()`

UnsetRollForwardLogPathVec ensures that no value is present for RollForwardLogPathVec, not even an explicit nil
### GetAttemptCompleteRecovery

`func (o *RecoverOracleAppOriginalSourceConfig) GetAttemptCompleteRecovery() bool`

GetAttemptCompleteRecovery returns the AttemptCompleteRecovery field if non-nil, zero value otherwise.

### GetAttemptCompleteRecoveryOk

`func (o *RecoverOracleAppOriginalSourceConfig) GetAttemptCompleteRecoveryOk() (*bool, bool)`

GetAttemptCompleteRecoveryOk returns a tuple with the AttemptCompleteRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptCompleteRecovery

`func (o *RecoverOracleAppOriginalSourceConfig) SetAttemptCompleteRecovery(v bool)`

SetAttemptCompleteRecovery sets AttemptCompleteRecovery field to given value.

### HasAttemptCompleteRecovery

`func (o *RecoverOracleAppOriginalSourceConfig) HasAttemptCompleteRecovery() bool`

HasAttemptCompleteRecovery returns a boolean if a field has been set.

### SetAttemptCompleteRecoveryNil

`func (o *RecoverOracleAppOriginalSourceConfig) SetAttemptCompleteRecoveryNil(b bool)`

 SetAttemptCompleteRecoveryNil sets the value for AttemptCompleteRecovery to be an explicit nil

### UnsetAttemptCompleteRecovery
`func (o *RecoverOracleAppOriginalSourceConfig) UnsetAttemptCompleteRecovery()`

UnsetAttemptCompleteRecovery ensures that no value is present for AttemptCompleteRecovery, not even an explicit nil
### GetRestoreTimeUsecs

`func (o *RecoverOracleAppOriginalSourceConfig) GetRestoreTimeUsecs() int64`

GetRestoreTimeUsecs returns the RestoreTimeUsecs field if non-nil, zero value otherwise.

### GetRestoreTimeUsecsOk

`func (o *RecoverOracleAppOriginalSourceConfig) GetRestoreTimeUsecsOk() (*int64, bool)`

GetRestoreTimeUsecsOk returns a tuple with the RestoreTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTimeUsecs

`func (o *RecoverOracleAppOriginalSourceConfig) SetRestoreTimeUsecs(v int64)`

SetRestoreTimeUsecs sets RestoreTimeUsecs field to given value.

### HasRestoreTimeUsecs

`func (o *RecoverOracleAppOriginalSourceConfig) HasRestoreTimeUsecs() bool`

HasRestoreTimeUsecs returns a boolean if a field has been set.

### SetRestoreTimeUsecsNil

`func (o *RecoverOracleAppOriginalSourceConfig) SetRestoreTimeUsecsNil(b bool)`

 SetRestoreTimeUsecsNil sets the value for RestoreTimeUsecs to be an explicit nil

### UnsetRestoreTimeUsecs
`func (o *RecoverOracleAppOriginalSourceConfig) UnsetRestoreTimeUsecs()`

UnsetRestoreTimeUsecs ensures that no value is present for RestoreTimeUsecs, not even an explicit nil
### GetDbChannels

`func (o *RecoverOracleAppOriginalSourceConfig) GetDbChannels() []OracleDbChannel`

GetDbChannels returns the DbChannels field if non-nil, zero value otherwise.

### GetDbChannelsOk

`func (o *RecoverOracleAppOriginalSourceConfig) GetDbChannelsOk() (*[]OracleDbChannel, bool)`

GetDbChannelsOk returns a tuple with the DbChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbChannels

`func (o *RecoverOracleAppOriginalSourceConfig) SetDbChannels(v []OracleDbChannel)`

SetDbChannels sets DbChannels field to given value.

### HasDbChannels

`func (o *RecoverOracleAppOriginalSourceConfig) HasDbChannels() bool`

HasDbChannels returns a boolean if a field has been set.

### SetDbChannelsNil

`func (o *RecoverOracleAppOriginalSourceConfig) SetDbChannelsNil(b bool)`

 SetDbChannelsNil sets the value for DbChannels to be an explicit nil

### UnsetDbChannels
`func (o *RecoverOracleAppOriginalSourceConfig) UnsetDbChannels()`

UnsetDbChannels ensures that no value is present for DbChannels, not even an explicit nil
### GetRecoveryMode

`func (o *RecoverOracleAppOriginalSourceConfig) GetRecoveryMode() bool`

GetRecoveryMode returns the RecoveryMode field if non-nil, zero value otherwise.

### GetRecoveryModeOk

`func (o *RecoverOracleAppOriginalSourceConfig) GetRecoveryModeOk() (*bool, bool)`

GetRecoveryModeOk returns a tuple with the RecoveryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryMode

`func (o *RecoverOracleAppOriginalSourceConfig) SetRecoveryMode(v bool)`

SetRecoveryMode sets RecoveryMode field to given value.

### HasRecoveryMode

`func (o *RecoverOracleAppOriginalSourceConfig) HasRecoveryMode() bool`

HasRecoveryMode returns a boolean if a field has been set.

### SetRecoveryModeNil

`func (o *RecoverOracleAppOriginalSourceConfig) SetRecoveryModeNil(b bool)`

 SetRecoveryModeNil sets the value for RecoveryMode to be an explicit nil

### UnsetRecoveryMode
`func (o *RecoverOracleAppOriginalSourceConfig) UnsetRecoveryMode()`

UnsetRecoveryMode ensures that no value is present for RecoveryMode, not even an explicit nil
### GetShellEvironmentVars

`func (o *RecoverOracleAppOriginalSourceConfig) GetShellEvironmentVars() []ShellKeyValuePair`

GetShellEvironmentVars returns the ShellEvironmentVars field if non-nil, zero value otherwise.

### GetShellEvironmentVarsOk

`func (o *RecoverOracleAppOriginalSourceConfig) GetShellEvironmentVarsOk() (*[]ShellKeyValuePair, bool)`

GetShellEvironmentVarsOk returns a tuple with the ShellEvironmentVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShellEvironmentVars

`func (o *RecoverOracleAppOriginalSourceConfig) SetShellEvironmentVars(v []ShellKeyValuePair)`

SetShellEvironmentVars sets ShellEvironmentVars field to given value.

### HasShellEvironmentVars

`func (o *RecoverOracleAppOriginalSourceConfig) HasShellEvironmentVars() bool`

HasShellEvironmentVars returns a boolean if a field has been set.

### SetShellEvironmentVarsNil

`func (o *RecoverOracleAppOriginalSourceConfig) SetShellEvironmentVarsNil(b bool)`

 SetShellEvironmentVarsNil sets the value for ShellEvironmentVars to be an explicit nil

### UnsetShellEvironmentVars
`func (o *RecoverOracleAppOriginalSourceConfig) UnsetShellEvironmentVars()`

UnsetShellEvironmentVars ensures that no value is present for ShellEvironmentVars, not even an explicit nil
### GetGranularRestoreInfo

`func (o *RecoverOracleAppOriginalSourceConfig) GetGranularRestoreInfo() RecoverOracleGranularRestoreInfo`

GetGranularRestoreInfo returns the GranularRestoreInfo field if non-nil, zero value otherwise.

### GetGranularRestoreInfoOk

`func (o *RecoverOracleAppOriginalSourceConfig) GetGranularRestoreInfoOk() (*RecoverOracleGranularRestoreInfo, bool)`

GetGranularRestoreInfoOk returns a tuple with the GranularRestoreInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularRestoreInfo

`func (o *RecoverOracleAppOriginalSourceConfig) SetGranularRestoreInfo(v RecoverOracleGranularRestoreInfo)`

SetGranularRestoreInfo sets GranularRestoreInfo field to given value.

### HasGranularRestoreInfo

`func (o *RecoverOracleAppOriginalSourceConfig) HasGranularRestoreInfo() bool`

HasGranularRestoreInfo returns a boolean if a field has been set.

### GetOracleArchiveLogInfo

`func (o *RecoverOracleAppOriginalSourceConfig) GetOracleArchiveLogInfo() OracleArchiveLogInfo`

GetOracleArchiveLogInfo returns the OracleArchiveLogInfo field if non-nil, zero value otherwise.

### GetOracleArchiveLogInfoOk

`func (o *RecoverOracleAppOriginalSourceConfig) GetOracleArchiveLogInfoOk() (*OracleArchiveLogInfo, bool)`

GetOracleArchiveLogInfoOk returns a tuple with the OracleArchiveLogInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleArchiveLogInfo

`func (o *RecoverOracleAppOriginalSourceConfig) SetOracleArchiveLogInfo(v OracleArchiveLogInfo)`

SetOracleArchiveLogInfo sets OracleArchiveLogInfo field to given value.

### HasOracleArchiveLogInfo

`func (o *RecoverOracleAppOriginalSourceConfig) HasOracleArchiveLogInfo() bool`

HasOracleArchiveLogInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


