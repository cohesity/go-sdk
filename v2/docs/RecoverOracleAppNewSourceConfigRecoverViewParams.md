# RecoverOracleAppNewSourceConfigRecoverViewParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViewMountPath** | Pointer to **NullableString** | Specifies the directory where cohesity view for app recovery will be mounted. | [optional] 
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

### NewRecoverOracleAppNewSourceConfigRecoverViewParams

`func NewRecoverOracleAppNewSourceConfigRecoverViewParams() *RecoverOracleAppNewSourceConfigRecoverViewParams`

NewRecoverOracleAppNewSourceConfigRecoverViewParams instantiates a new RecoverOracleAppNewSourceConfigRecoverViewParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverOracleAppNewSourceConfigRecoverViewParamsWithDefaults

`func NewRecoverOracleAppNewSourceConfigRecoverViewParamsWithDefaults() *RecoverOracleAppNewSourceConfigRecoverViewParams`

NewRecoverOracleAppNewSourceConfigRecoverViewParamsWithDefaults instantiates a new RecoverOracleAppNewSourceConfigRecoverViewParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViewMountPath

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) GetViewMountPath() string`

GetViewMountPath returns the ViewMountPath field if non-nil, zero value otherwise.

### GetViewMountPathOk

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) GetViewMountPathOk() (*string, bool)`

GetViewMountPathOk returns a tuple with the ViewMountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewMountPath

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) SetViewMountPath(v string)`

SetViewMountPath sets ViewMountPath field to given value.

### HasViewMountPath

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) HasViewMountPath() bool`

HasViewMountPath returns a boolean if a field has been set.

### SetViewMountPathNil

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) SetViewMountPathNil(b bool)`

 SetViewMountPathNil sets the value for ViewMountPath to be an explicit nil

### UnsetViewMountPath
`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) UnsetViewMountPath()`

UnsetViewMountPath ensures that no value is present for ViewMountPath, not even an explicit nil
### GetDbChannels

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) GetDbChannels() []OracleDbChannel`

GetDbChannels returns the DbChannels field if non-nil, zero value otherwise.

### GetDbChannelsOk

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) GetDbChannelsOk() (*[]OracleDbChannel, bool)`

GetDbChannelsOk returns a tuple with the DbChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbChannels

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) SetDbChannels(v []OracleDbChannel)`

SetDbChannels sets DbChannels field to given value.

### HasDbChannels

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) HasDbChannels() bool`

HasDbChannels returns a boolean if a field has been set.

### SetDbChannelsNil

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) SetDbChannelsNil(b bool)`

 SetDbChannelsNil sets the value for DbChannels to be an explicit nil

### UnsetDbChannels
`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) UnsetDbChannels()`

UnsetDbChannels ensures that no value is present for DbChannels, not even an explicit nil
### GetGranularRestoreInfo

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) GetGranularRestoreInfo() CommonOracleAppSourceConfigGranularRestoreInfo`

GetGranularRestoreInfo returns the GranularRestoreInfo field if non-nil, zero value otherwise.

### GetGranularRestoreInfoOk

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) GetGranularRestoreInfoOk() (*CommonOracleAppSourceConfigGranularRestoreInfo, bool)`

GetGranularRestoreInfoOk returns a tuple with the GranularRestoreInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularRestoreInfo

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) SetGranularRestoreInfo(v CommonOracleAppSourceConfigGranularRestoreInfo)`

SetGranularRestoreInfo sets GranularRestoreInfo field to given value.

### HasGranularRestoreInfo

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) HasGranularRestoreInfo() bool`

HasGranularRestoreInfo returns a boolean if a field has been set.

### GetOracleArchiveLogInfo

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) GetOracleArchiveLogInfo() CommonOracleAppSourceConfigOracleArchiveLogInfo`

GetOracleArchiveLogInfo returns the OracleArchiveLogInfo field if non-nil, zero value otherwise.

### GetOracleArchiveLogInfoOk

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) GetOracleArchiveLogInfoOk() (*CommonOracleAppSourceConfigOracleArchiveLogInfo, bool)`

GetOracleArchiveLogInfoOk returns a tuple with the OracleArchiveLogInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleArchiveLogInfo

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) SetOracleArchiveLogInfo(v CommonOracleAppSourceConfigOracleArchiveLogInfo)`

SetOracleArchiveLogInfo sets OracleArchiveLogInfo field to given value.

### HasOracleArchiveLogInfo

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) HasOracleArchiveLogInfo() bool`

HasOracleArchiveLogInfo returns a boolean if a field has been set.

### GetOracleRecoveryValidationInfo

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) GetOracleRecoveryValidationInfo() CommonOracleAppSourceConfigOracleRecoveryValidationInfo`

GetOracleRecoveryValidationInfo returns the OracleRecoveryValidationInfo field if non-nil, zero value otherwise.

### GetOracleRecoveryValidationInfoOk

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) GetOracleRecoveryValidationInfoOk() (*CommonOracleAppSourceConfigOracleRecoveryValidationInfo, bool)`

GetOracleRecoveryValidationInfoOk returns a tuple with the OracleRecoveryValidationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleRecoveryValidationInfo

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) SetOracleRecoveryValidationInfo(v CommonOracleAppSourceConfigOracleRecoveryValidationInfo)`

SetOracleRecoveryValidationInfo sets OracleRecoveryValidationInfo field to given value.

### HasOracleRecoveryValidationInfo

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) HasOracleRecoveryValidationInfo() bool`

HasOracleRecoveryValidationInfo returns a boolean if a field has been set.

### GetRecoveryMode

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) GetRecoveryMode() bool`

GetRecoveryMode returns the RecoveryMode field if non-nil, zero value otherwise.

### GetRecoveryModeOk

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) GetRecoveryModeOk() (*bool, bool)`

GetRecoveryModeOk returns a tuple with the RecoveryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryMode

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) SetRecoveryMode(v bool)`

SetRecoveryMode sets RecoveryMode field to given value.

### HasRecoveryMode

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) HasRecoveryMode() bool`

HasRecoveryMode returns a boolean if a field has been set.

### SetRecoveryModeNil

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) SetRecoveryModeNil(b bool)`

 SetRecoveryModeNil sets the value for RecoveryMode to be an explicit nil

### UnsetRecoveryMode
`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) UnsetRecoveryMode()`

UnsetRecoveryMode ensures that no value is present for RecoveryMode, not even an explicit nil
### GetRestoreSpfileOrPfileInfo

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) GetRestoreSpfileOrPfileInfo() CommonOracleAppSourceConfigRestoreSpfileOrPfileInfo`

GetRestoreSpfileOrPfileInfo returns the RestoreSpfileOrPfileInfo field if non-nil, zero value otherwise.

### GetRestoreSpfileOrPfileInfoOk

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) GetRestoreSpfileOrPfileInfoOk() (*CommonOracleAppSourceConfigRestoreSpfileOrPfileInfo, bool)`

GetRestoreSpfileOrPfileInfoOk returns a tuple with the RestoreSpfileOrPfileInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreSpfileOrPfileInfo

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) SetRestoreSpfileOrPfileInfo(v CommonOracleAppSourceConfigRestoreSpfileOrPfileInfo)`

SetRestoreSpfileOrPfileInfo sets RestoreSpfileOrPfileInfo field to given value.

### HasRestoreSpfileOrPfileInfo

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) HasRestoreSpfileOrPfileInfo() bool`

HasRestoreSpfileOrPfileInfo returns a boolean if a field has been set.

### GetRestoreTimeUsecs

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) GetRestoreTimeUsecs() int64`

GetRestoreTimeUsecs returns the RestoreTimeUsecs field if non-nil, zero value otherwise.

### GetRestoreTimeUsecsOk

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) GetRestoreTimeUsecsOk() (*int64, bool)`

GetRestoreTimeUsecsOk returns a tuple with the RestoreTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTimeUsecs

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) SetRestoreTimeUsecs(v int64)`

SetRestoreTimeUsecs sets RestoreTimeUsecs field to given value.

### HasRestoreTimeUsecs

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) HasRestoreTimeUsecs() bool`

HasRestoreTimeUsecs returns a boolean if a field has been set.

### SetRestoreTimeUsecsNil

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) SetRestoreTimeUsecsNil(b bool)`

 SetRestoreTimeUsecsNil sets the value for RestoreTimeUsecs to be an explicit nil

### UnsetRestoreTimeUsecs
`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) UnsetRestoreTimeUsecs()`

UnsetRestoreTimeUsecs ensures that no value is present for RestoreTimeUsecs, not even an explicit nil
### GetShellEvironmentVars

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) GetShellEvironmentVars() []ShellKeyValuePair`

GetShellEvironmentVars returns the ShellEvironmentVars field if non-nil, zero value otherwise.

### GetShellEvironmentVarsOk

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) GetShellEvironmentVarsOk() (*[]ShellKeyValuePair, bool)`

GetShellEvironmentVarsOk returns a tuple with the ShellEvironmentVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShellEvironmentVars

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) SetShellEvironmentVars(v []ShellKeyValuePair)`

SetShellEvironmentVars sets ShellEvironmentVars field to given value.

### HasShellEvironmentVars

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) HasShellEvironmentVars() bool`

HasShellEvironmentVars returns a boolean if a field has been set.

### SetShellEvironmentVarsNil

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) SetShellEvironmentVarsNil(b bool)`

 SetShellEvironmentVarsNil sets the value for ShellEvironmentVars to be an explicit nil

### UnsetShellEvironmentVars
`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) UnsetShellEvironmentVars()`

UnsetShellEvironmentVars ensures that no value is present for ShellEvironmentVars, not even an explicit nil
### GetUseScnForRestore

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) GetUseScnForRestore() bool`

GetUseScnForRestore returns the UseScnForRestore field if non-nil, zero value otherwise.

### GetUseScnForRestoreOk

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) GetUseScnForRestoreOk() (*bool, bool)`

GetUseScnForRestoreOk returns a tuple with the UseScnForRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseScnForRestore

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) SetUseScnForRestore(v bool)`

SetUseScnForRestore sets UseScnForRestore field to given value.

### HasUseScnForRestore

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) HasUseScnForRestore() bool`

HasUseScnForRestore returns a boolean if a field has been set.

### SetUseScnForRestoreNil

`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) SetUseScnForRestoreNil(b bool)`

 SetUseScnForRestoreNil sets the value for UseScnForRestore to be an explicit nil

### UnsetUseScnForRestore
`func (o *RecoverOracleAppNewSourceConfigRecoverViewParams) UnsetUseScnForRestore()`

UnsetUseScnForRestore ensures that no value is present for UseScnForRestore, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


