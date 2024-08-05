# RecoverOracleAppOriginalSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptCompleteRecovery** | Pointer to **NullableBool** | Whether or not this is a complete recovery attempt. | [optional] 
**RollForwardLogPathVec** | Pointer to **[]string** | List of archive logs to apply on Database after overwrite restore. | [optional] 
**RollForwardTimeMsecs** | Pointer to **NullableInt64** | UTC time in msecs till which we have to roll-forward the database. | [optional] 
**StopActivePassive** | Pointer to **NullableBool** | Specifies whether allowed to automatically stop active passive resource. | [optional] 
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
### GetRollForwardTimeMsecs

`func (o *RecoverOracleAppOriginalSourceConfig) GetRollForwardTimeMsecs() int64`

GetRollForwardTimeMsecs returns the RollForwardTimeMsecs field if non-nil, zero value otherwise.

### GetRollForwardTimeMsecsOk

`func (o *RecoverOracleAppOriginalSourceConfig) GetRollForwardTimeMsecsOk() (*int64, bool)`

GetRollForwardTimeMsecsOk returns a tuple with the RollForwardTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollForwardTimeMsecs

`func (o *RecoverOracleAppOriginalSourceConfig) SetRollForwardTimeMsecs(v int64)`

SetRollForwardTimeMsecs sets RollForwardTimeMsecs field to given value.

### HasRollForwardTimeMsecs

`func (o *RecoverOracleAppOriginalSourceConfig) HasRollForwardTimeMsecs() bool`

HasRollForwardTimeMsecs returns a boolean if a field has been set.

### SetRollForwardTimeMsecsNil

`func (o *RecoverOracleAppOriginalSourceConfig) SetRollForwardTimeMsecsNil(b bool)`

 SetRollForwardTimeMsecsNil sets the value for RollForwardTimeMsecs to be an explicit nil

### UnsetRollForwardTimeMsecs
`func (o *RecoverOracleAppOriginalSourceConfig) UnsetRollForwardTimeMsecs()`

UnsetRollForwardTimeMsecs ensures that no value is present for RollForwardTimeMsecs, not even an explicit nil
### GetStopActivePassive

`func (o *RecoverOracleAppOriginalSourceConfig) GetStopActivePassive() bool`

GetStopActivePassive returns the StopActivePassive field if non-nil, zero value otherwise.

### GetStopActivePassiveOk

`func (o *RecoverOracleAppOriginalSourceConfig) GetStopActivePassiveOk() (*bool, bool)`

GetStopActivePassiveOk returns a tuple with the StopActivePassive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopActivePassive

`func (o *RecoverOracleAppOriginalSourceConfig) SetStopActivePassive(v bool)`

SetStopActivePassive sets StopActivePassive field to given value.

### HasStopActivePassive

`func (o *RecoverOracleAppOriginalSourceConfig) HasStopActivePassive() bool`

HasStopActivePassive returns a boolean if a field has been set.

### SetStopActivePassiveNil

`func (o *RecoverOracleAppOriginalSourceConfig) SetStopActivePassiveNil(b bool)`

 SetStopActivePassiveNil sets the value for StopActivePassive to be an explicit nil

### UnsetStopActivePassive
`func (o *RecoverOracleAppOriginalSourceConfig) UnsetStopActivePassive()`

UnsetStopActivePassive ensures that no value is present for StopActivePassive, not even an explicit nil
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
### GetGranularRestoreInfo

`func (o *RecoverOracleAppOriginalSourceConfig) GetGranularRestoreInfo() CommonOracleAppSourceConfigGranularRestoreInfo`

GetGranularRestoreInfo returns the GranularRestoreInfo field if non-nil, zero value otherwise.

### GetGranularRestoreInfoOk

`func (o *RecoverOracleAppOriginalSourceConfig) GetGranularRestoreInfoOk() (*CommonOracleAppSourceConfigGranularRestoreInfo, bool)`

GetGranularRestoreInfoOk returns a tuple with the GranularRestoreInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularRestoreInfo

`func (o *RecoverOracleAppOriginalSourceConfig) SetGranularRestoreInfo(v CommonOracleAppSourceConfigGranularRestoreInfo)`

SetGranularRestoreInfo sets GranularRestoreInfo field to given value.

### HasGranularRestoreInfo

`func (o *RecoverOracleAppOriginalSourceConfig) HasGranularRestoreInfo() bool`

HasGranularRestoreInfo returns a boolean if a field has been set.

### GetOracleArchiveLogInfo

`func (o *RecoverOracleAppOriginalSourceConfig) GetOracleArchiveLogInfo() CommonOracleAppSourceConfigOracleArchiveLogInfo`

GetOracleArchiveLogInfo returns the OracleArchiveLogInfo field if non-nil, zero value otherwise.

### GetOracleArchiveLogInfoOk

`func (o *RecoverOracleAppOriginalSourceConfig) GetOracleArchiveLogInfoOk() (*CommonOracleAppSourceConfigOracleArchiveLogInfo, bool)`

GetOracleArchiveLogInfoOk returns a tuple with the OracleArchiveLogInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleArchiveLogInfo

`func (o *RecoverOracleAppOriginalSourceConfig) SetOracleArchiveLogInfo(v CommonOracleAppSourceConfigOracleArchiveLogInfo)`

SetOracleArchiveLogInfo sets OracleArchiveLogInfo field to given value.

### HasOracleArchiveLogInfo

`func (o *RecoverOracleAppOriginalSourceConfig) HasOracleArchiveLogInfo() bool`

HasOracleArchiveLogInfo returns a boolean if a field has been set.

### GetOracleRecoveryValidationInfo

`func (o *RecoverOracleAppOriginalSourceConfig) GetOracleRecoveryValidationInfo() CommonOracleAppSourceConfigOracleRecoveryValidationInfo`

GetOracleRecoveryValidationInfo returns the OracleRecoveryValidationInfo field if non-nil, zero value otherwise.

### GetOracleRecoveryValidationInfoOk

`func (o *RecoverOracleAppOriginalSourceConfig) GetOracleRecoveryValidationInfoOk() (*CommonOracleAppSourceConfigOracleRecoveryValidationInfo, bool)`

GetOracleRecoveryValidationInfoOk returns a tuple with the OracleRecoveryValidationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleRecoveryValidationInfo

`func (o *RecoverOracleAppOriginalSourceConfig) SetOracleRecoveryValidationInfo(v CommonOracleAppSourceConfigOracleRecoveryValidationInfo)`

SetOracleRecoveryValidationInfo sets OracleRecoveryValidationInfo field to given value.

### HasOracleRecoveryValidationInfo

`func (o *RecoverOracleAppOriginalSourceConfig) HasOracleRecoveryValidationInfo() bool`

HasOracleRecoveryValidationInfo returns a boolean if a field has been set.

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
### GetRestoreSpfileOrPfileInfo

`func (o *RecoverOracleAppOriginalSourceConfig) GetRestoreSpfileOrPfileInfo() CommonOracleAppSourceConfigRestoreSpfileOrPfileInfo`

GetRestoreSpfileOrPfileInfo returns the RestoreSpfileOrPfileInfo field if non-nil, zero value otherwise.

### GetRestoreSpfileOrPfileInfoOk

`func (o *RecoverOracleAppOriginalSourceConfig) GetRestoreSpfileOrPfileInfoOk() (*CommonOracleAppSourceConfigRestoreSpfileOrPfileInfo, bool)`

GetRestoreSpfileOrPfileInfoOk returns a tuple with the RestoreSpfileOrPfileInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreSpfileOrPfileInfo

`func (o *RecoverOracleAppOriginalSourceConfig) SetRestoreSpfileOrPfileInfo(v CommonOracleAppSourceConfigRestoreSpfileOrPfileInfo)`

SetRestoreSpfileOrPfileInfo sets RestoreSpfileOrPfileInfo field to given value.

### HasRestoreSpfileOrPfileInfo

`func (o *RecoverOracleAppOriginalSourceConfig) HasRestoreSpfileOrPfileInfo() bool`

HasRestoreSpfileOrPfileInfo returns a boolean if a field has been set.

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
### GetUseScnForRestore

`func (o *RecoverOracleAppOriginalSourceConfig) GetUseScnForRestore() bool`

GetUseScnForRestore returns the UseScnForRestore field if non-nil, zero value otherwise.

### GetUseScnForRestoreOk

`func (o *RecoverOracleAppOriginalSourceConfig) GetUseScnForRestoreOk() (*bool, bool)`

GetUseScnForRestoreOk returns a tuple with the UseScnForRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseScnForRestore

`func (o *RecoverOracleAppOriginalSourceConfig) SetUseScnForRestore(v bool)`

SetUseScnForRestore sets UseScnForRestore field to given value.

### HasUseScnForRestore

`func (o *RecoverOracleAppOriginalSourceConfig) HasUseScnForRestore() bool`

HasUseScnForRestore returns a boolean if a field has been set.

### SetUseScnForRestoreNil

`func (o *RecoverOracleAppOriginalSourceConfig) SetUseScnForRestoreNil(b bool)`

 SetUseScnForRestoreNil sets the value for UseScnForRestore to be an explicit nil

### UnsetUseScnForRestore
`func (o *RecoverOracleAppOriginalSourceConfig) UnsetUseScnForRestore()`

UnsetUseScnForRestore ensures that no value is present for UseScnForRestore, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


