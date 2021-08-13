# RecoverOracleNewTargetViewConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViewMountPath** | Pointer to **NullableString** | Specifies the directory where cohesity view for app recovery will be mounted. | [optional] 
**RestoreTimeUsecs** | Pointer to **NullableInt64** | Specifies the time in the past to which the Oracle db needs to be restored. This allows for granular recovery of Oracle databases. If this is not set, the Oracle db will be restored from the full/incremental snapshot. | [optional] 
**DbChannels** | Pointer to [**[]OracleDbChannel**](OracleDbChannel.md) | Specifies the Oracle database node channels info. If not specified, the default values assigned by the server are applied to all the databases. | [optional] 
**RecoveryMode** | Pointer to **NullableBool** | Specifies if database should be left in recovery mode. | [optional] 
**ShellEvironmentVars** | Pointer to [**[]ShellKeyValuePair**](ShellKeyValuePair.md) | Specifies key value pairs of shell variables which defines the restore shell environment. | [optional] 
**GranularRestoreInfo** | Pointer to [**RecoverOracleGranularRestoreInfo**](RecoverOracleGranularRestoreInfo.md) | Specifies information about list of objects (PDBs) to restore. | [optional] 
**OracleArchiveLogInfo** | Pointer to [**OracleArchiveLogInfo**](OracleArchiveLogInfo.md) | Specifies Range in Time, Scn or Sequence to restore archive logs of a DB. | [optional] 

## Methods

### NewRecoverOracleNewTargetViewConfig

`func NewRecoverOracleNewTargetViewConfig() *RecoverOracleNewTargetViewConfig`

NewRecoverOracleNewTargetViewConfig instantiates a new RecoverOracleNewTargetViewConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverOracleNewTargetViewConfigWithDefaults

`func NewRecoverOracleNewTargetViewConfigWithDefaults() *RecoverOracleNewTargetViewConfig`

NewRecoverOracleNewTargetViewConfigWithDefaults instantiates a new RecoverOracleNewTargetViewConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViewMountPath

`func (o *RecoverOracleNewTargetViewConfig) GetViewMountPath() string`

GetViewMountPath returns the ViewMountPath field if non-nil, zero value otherwise.

### GetViewMountPathOk

`func (o *RecoverOracleNewTargetViewConfig) GetViewMountPathOk() (*string, bool)`

GetViewMountPathOk returns a tuple with the ViewMountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewMountPath

`func (o *RecoverOracleNewTargetViewConfig) SetViewMountPath(v string)`

SetViewMountPath sets ViewMountPath field to given value.

### HasViewMountPath

`func (o *RecoverOracleNewTargetViewConfig) HasViewMountPath() bool`

HasViewMountPath returns a boolean if a field has been set.

### SetViewMountPathNil

`func (o *RecoverOracleNewTargetViewConfig) SetViewMountPathNil(b bool)`

 SetViewMountPathNil sets the value for ViewMountPath to be an explicit nil

### UnsetViewMountPath
`func (o *RecoverOracleNewTargetViewConfig) UnsetViewMountPath()`

UnsetViewMountPath ensures that no value is present for ViewMountPath, not even an explicit nil
### GetRestoreTimeUsecs

`func (o *RecoverOracleNewTargetViewConfig) GetRestoreTimeUsecs() int64`

GetRestoreTimeUsecs returns the RestoreTimeUsecs field if non-nil, zero value otherwise.

### GetRestoreTimeUsecsOk

`func (o *RecoverOracleNewTargetViewConfig) GetRestoreTimeUsecsOk() (*int64, bool)`

GetRestoreTimeUsecsOk returns a tuple with the RestoreTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTimeUsecs

`func (o *RecoverOracleNewTargetViewConfig) SetRestoreTimeUsecs(v int64)`

SetRestoreTimeUsecs sets RestoreTimeUsecs field to given value.

### HasRestoreTimeUsecs

`func (o *RecoverOracleNewTargetViewConfig) HasRestoreTimeUsecs() bool`

HasRestoreTimeUsecs returns a boolean if a field has been set.

### SetRestoreTimeUsecsNil

`func (o *RecoverOracleNewTargetViewConfig) SetRestoreTimeUsecsNil(b bool)`

 SetRestoreTimeUsecsNil sets the value for RestoreTimeUsecs to be an explicit nil

### UnsetRestoreTimeUsecs
`func (o *RecoverOracleNewTargetViewConfig) UnsetRestoreTimeUsecs()`

UnsetRestoreTimeUsecs ensures that no value is present for RestoreTimeUsecs, not even an explicit nil
### GetDbChannels

`func (o *RecoverOracleNewTargetViewConfig) GetDbChannels() []OracleDbChannel`

GetDbChannels returns the DbChannels field if non-nil, zero value otherwise.

### GetDbChannelsOk

`func (o *RecoverOracleNewTargetViewConfig) GetDbChannelsOk() (*[]OracleDbChannel, bool)`

GetDbChannelsOk returns a tuple with the DbChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbChannels

`func (o *RecoverOracleNewTargetViewConfig) SetDbChannels(v []OracleDbChannel)`

SetDbChannels sets DbChannels field to given value.

### HasDbChannels

`func (o *RecoverOracleNewTargetViewConfig) HasDbChannels() bool`

HasDbChannels returns a boolean if a field has been set.

### SetDbChannelsNil

`func (o *RecoverOracleNewTargetViewConfig) SetDbChannelsNil(b bool)`

 SetDbChannelsNil sets the value for DbChannels to be an explicit nil

### UnsetDbChannels
`func (o *RecoverOracleNewTargetViewConfig) UnsetDbChannels()`

UnsetDbChannels ensures that no value is present for DbChannels, not even an explicit nil
### GetRecoveryMode

`func (o *RecoverOracleNewTargetViewConfig) GetRecoveryMode() bool`

GetRecoveryMode returns the RecoveryMode field if non-nil, zero value otherwise.

### GetRecoveryModeOk

`func (o *RecoverOracleNewTargetViewConfig) GetRecoveryModeOk() (*bool, bool)`

GetRecoveryModeOk returns a tuple with the RecoveryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryMode

`func (o *RecoverOracleNewTargetViewConfig) SetRecoveryMode(v bool)`

SetRecoveryMode sets RecoveryMode field to given value.

### HasRecoveryMode

`func (o *RecoverOracleNewTargetViewConfig) HasRecoveryMode() bool`

HasRecoveryMode returns a boolean if a field has been set.

### SetRecoveryModeNil

`func (o *RecoverOracleNewTargetViewConfig) SetRecoveryModeNil(b bool)`

 SetRecoveryModeNil sets the value for RecoveryMode to be an explicit nil

### UnsetRecoveryMode
`func (o *RecoverOracleNewTargetViewConfig) UnsetRecoveryMode()`

UnsetRecoveryMode ensures that no value is present for RecoveryMode, not even an explicit nil
### GetShellEvironmentVars

`func (o *RecoverOracleNewTargetViewConfig) GetShellEvironmentVars() []ShellKeyValuePair`

GetShellEvironmentVars returns the ShellEvironmentVars field if non-nil, zero value otherwise.

### GetShellEvironmentVarsOk

`func (o *RecoverOracleNewTargetViewConfig) GetShellEvironmentVarsOk() (*[]ShellKeyValuePair, bool)`

GetShellEvironmentVarsOk returns a tuple with the ShellEvironmentVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShellEvironmentVars

`func (o *RecoverOracleNewTargetViewConfig) SetShellEvironmentVars(v []ShellKeyValuePair)`

SetShellEvironmentVars sets ShellEvironmentVars field to given value.

### HasShellEvironmentVars

`func (o *RecoverOracleNewTargetViewConfig) HasShellEvironmentVars() bool`

HasShellEvironmentVars returns a boolean if a field has been set.

### SetShellEvironmentVarsNil

`func (o *RecoverOracleNewTargetViewConfig) SetShellEvironmentVarsNil(b bool)`

 SetShellEvironmentVarsNil sets the value for ShellEvironmentVars to be an explicit nil

### UnsetShellEvironmentVars
`func (o *RecoverOracleNewTargetViewConfig) UnsetShellEvironmentVars()`

UnsetShellEvironmentVars ensures that no value is present for ShellEvironmentVars, not even an explicit nil
### GetGranularRestoreInfo

`func (o *RecoverOracleNewTargetViewConfig) GetGranularRestoreInfo() RecoverOracleGranularRestoreInfo`

GetGranularRestoreInfo returns the GranularRestoreInfo field if non-nil, zero value otherwise.

### GetGranularRestoreInfoOk

`func (o *RecoverOracleNewTargetViewConfig) GetGranularRestoreInfoOk() (*RecoverOracleGranularRestoreInfo, bool)`

GetGranularRestoreInfoOk returns a tuple with the GranularRestoreInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularRestoreInfo

`func (o *RecoverOracleNewTargetViewConfig) SetGranularRestoreInfo(v RecoverOracleGranularRestoreInfo)`

SetGranularRestoreInfo sets GranularRestoreInfo field to given value.

### HasGranularRestoreInfo

`func (o *RecoverOracleNewTargetViewConfig) HasGranularRestoreInfo() bool`

HasGranularRestoreInfo returns a boolean if a field has been set.

### GetOracleArchiveLogInfo

`func (o *RecoverOracleNewTargetViewConfig) GetOracleArchiveLogInfo() OracleArchiveLogInfo`

GetOracleArchiveLogInfo returns the OracleArchiveLogInfo field if non-nil, zero value otherwise.

### GetOracleArchiveLogInfoOk

`func (o *RecoverOracleNewTargetViewConfig) GetOracleArchiveLogInfoOk() (*OracleArchiveLogInfo, bool)`

GetOracleArchiveLogInfoOk returns a tuple with the OracleArchiveLogInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleArchiveLogInfo

`func (o *RecoverOracleNewTargetViewConfig) SetOracleArchiveLogInfo(v OracleArchiveLogInfo)`

SetOracleArchiveLogInfo sets OracleArchiveLogInfo field to given value.

### HasOracleArchiveLogInfo

`func (o *RecoverOracleNewTargetViewConfig) HasOracleArchiveLogInfo() bool`

HasOracleArchiveLogInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


