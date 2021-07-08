# OracleDBConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditLogDest** | Pointer to **NullableString** | Audit log destination. | [optional] 
**BctFilePath** | Pointer to **NullableString** | BCT file path. | [optional] 
**ControlFilePathVec** | Pointer to **[]string** | List of paths where the control file needs to be multiplexed. | [optional] 
**DbConfigFilePath** | Pointer to **NullableString** | Path to the file on oracle host which decides the configuration of restored DB. | [optional] 
**DiagDest** | Pointer to **NullableString** | Diag destination. | [optional] 
**EnableArchiveLogMode** | Pointer to **NullableBool** | If set to false, archive log mode is disabled. | [optional] 
**FraDest** | Pointer to **NullableString** | FRA path. | [optional] 
**FraSizeMb** | Pointer to **NullableInt32** | FRA Size in MBs. | [optional] 
**NumTempfiles** | Pointer to **NullableInt32** | How many tempfiles to use for the recovered database. | [optional] 
**PfileParameterMap** | Pointer to [**[]OracleDBConfigPfileParameterMapEntry**](OracleDBConfigPfileParameterMapEntry.md) | Map of pfile parameters to its values. | [optional] 
**RedoLogConf** | Pointer to [**OracleDBConfigRedoLogGroupConf**](OracleDBConfigRedoLogGroupConf.md) |  | [optional] 
**SgaTargetSize** | Pointer to **NullableString** | SGA_TARGET_SIZE size [ Default value same as Source DB ]. | [optional] 
**SharedPoolSize** | Pointer to **NullableString** | Shared pool size [ Default value same as Source DB ]. | [optional] 

## Methods

### NewOracleDBConfig

`func NewOracleDBConfig() *OracleDBConfig`

NewOracleDBConfig instantiates a new OracleDBConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleDBConfigWithDefaults

`func NewOracleDBConfigWithDefaults() *OracleDBConfig`

NewOracleDBConfigWithDefaults instantiates a new OracleDBConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditLogDest

`func (o *OracleDBConfig) GetAuditLogDest() string`

GetAuditLogDest returns the AuditLogDest field if non-nil, zero value otherwise.

### GetAuditLogDestOk

`func (o *OracleDBConfig) GetAuditLogDestOk() (*string, bool)`

GetAuditLogDestOk returns a tuple with the AuditLogDest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogDest

`func (o *OracleDBConfig) SetAuditLogDest(v string)`

SetAuditLogDest sets AuditLogDest field to given value.

### HasAuditLogDest

`func (o *OracleDBConfig) HasAuditLogDest() bool`

HasAuditLogDest returns a boolean if a field has been set.

### SetAuditLogDestNil

`func (o *OracleDBConfig) SetAuditLogDestNil(b bool)`

 SetAuditLogDestNil sets the value for AuditLogDest to be an explicit nil

### UnsetAuditLogDest
`func (o *OracleDBConfig) UnsetAuditLogDest()`

UnsetAuditLogDest ensures that no value is present for AuditLogDest, not even an explicit nil
### GetBctFilePath

`func (o *OracleDBConfig) GetBctFilePath() string`

GetBctFilePath returns the BctFilePath field if non-nil, zero value otherwise.

### GetBctFilePathOk

`func (o *OracleDBConfig) GetBctFilePathOk() (*string, bool)`

GetBctFilePathOk returns a tuple with the BctFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBctFilePath

`func (o *OracleDBConfig) SetBctFilePath(v string)`

SetBctFilePath sets BctFilePath field to given value.

### HasBctFilePath

`func (o *OracleDBConfig) HasBctFilePath() bool`

HasBctFilePath returns a boolean if a field has been set.

### SetBctFilePathNil

`func (o *OracleDBConfig) SetBctFilePathNil(b bool)`

 SetBctFilePathNil sets the value for BctFilePath to be an explicit nil

### UnsetBctFilePath
`func (o *OracleDBConfig) UnsetBctFilePath()`

UnsetBctFilePath ensures that no value is present for BctFilePath, not even an explicit nil
### GetControlFilePathVec

`func (o *OracleDBConfig) GetControlFilePathVec() []string`

GetControlFilePathVec returns the ControlFilePathVec field if non-nil, zero value otherwise.

### GetControlFilePathVecOk

`func (o *OracleDBConfig) GetControlFilePathVecOk() (*[]string, bool)`

GetControlFilePathVecOk returns a tuple with the ControlFilePathVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlFilePathVec

`func (o *OracleDBConfig) SetControlFilePathVec(v []string)`

SetControlFilePathVec sets ControlFilePathVec field to given value.

### HasControlFilePathVec

`func (o *OracleDBConfig) HasControlFilePathVec() bool`

HasControlFilePathVec returns a boolean if a field has been set.

### SetControlFilePathVecNil

`func (o *OracleDBConfig) SetControlFilePathVecNil(b bool)`

 SetControlFilePathVecNil sets the value for ControlFilePathVec to be an explicit nil

### UnsetControlFilePathVec
`func (o *OracleDBConfig) UnsetControlFilePathVec()`

UnsetControlFilePathVec ensures that no value is present for ControlFilePathVec, not even an explicit nil
### GetDbConfigFilePath

`func (o *OracleDBConfig) GetDbConfigFilePath() string`

GetDbConfigFilePath returns the DbConfigFilePath field if non-nil, zero value otherwise.

### GetDbConfigFilePathOk

`func (o *OracleDBConfig) GetDbConfigFilePathOk() (*string, bool)`

GetDbConfigFilePathOk returns a tuple with the DbConfigFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbConfigFilePath

`func (o *OracleDBConfig) SetDbConfigFilePath(v string)`

SetDbConfigFilePath sets DbConfigFilePath field to given value.

### HasDbConfigFilePath

`func (o *OracleDBConfig) HasDbConfigFilePath() bool`

HasDbConfigFilePath returns a boolean if a field has been set.

### SetDbConfigFilePathNil

`func (o *OracleDBConfig) SetDbConfigFilePathNil(b bool)`

 SetDbConfigFilePathNil sets the value for DbConfigFilePath to be an explicit nil

### UnsetDbConfigFilePath
`func (o *OracleDBConfig) UnsetDbConfigFilePath()`

UnsetDbConfigFilePath ensures that no value is present for DbConfigFilePath, not even an explicit nil
### GetDiagDest

`func (o *OracleDBConfig) GetDiagDest() string`

GetDiagDest returns the DiagDest field if non-nil, zero value otherwise.

### GetDiagDestOk

`func (o *OracleDBConfig) GetDiagDestOk() (*string, bool)`

GetDiagDestOk returns a tuple with the DiagDest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagDest

`func (o *OracleDBConfig) SetDiagDest(v string)`

SetDiagDest sets DiagDest field to given value.

### HasDiagDest

`func (o *OracleDBConfig) HasDiagDest() bool`

HasDiagDest returns a boolean if a field has been set.

### SetDiagDestNil

`func (o *OracleDBConfig) SetDiagDestNil(b bool)`

 SetDiagDestNil sets the value for DiagDest to be an explicit nil

### UnsetDiagDest
`func (o *OracleDBConfig) UnsetDiagDest()`

UnsetDiagDest ensures that no value is present for DiagDest, not even an explicit nil
### GetEnableArchiveLogMode

`func (o *OracleDBConfig) GetEnableArchiveLogMode() bool`

GetEnableArchiveLogMode returns the EnableArchiveLogMode field if non-nil, zero value otherwise.

### GetEnableArchiveLogModeOk

`func (o *OracleDBConfig) GetEnableArchiveLogModeOk() (*bool, bool)`

GetEnableArchiveLogModeOk returns a tuple with the EnableArchiveLogMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableArchiveLogMode

`func (o *OracleDBConfig) SetEnableArchiveLogMode(v bool)`

SetEnableArchiveLogMode sets EnableArchiveLogMode field to given value.

### HasEnableArchiveLogMode

`func (o *OracleDBConfig) HasEnableArchiveLogMode() bool`

HasEnableArchiveLogMode returns a boolean if a field has been set.

### SetEnableArchiveLogModeNil

`func (o *OracleDBConfig) SetEnableArchiveLogModeNil(b bool)`

 SetEnableArchiveLogModeNil sets the value for EnableArchiveLogMode to be an explicit nil

### UnsetEnableArchiveLogMode
`func (o *OracleDBConfig) UnsetEnableArchiveLogMode()`

UnsetEnableArchiveLogMode ensures that no value is present for EnableArchiveLogMode, not even an explicit nil
### GetFraDest

`func (o *OracleDBConfig) GetFraDest() string`

GetFraDest returns the FraDest field if non-nil, zero value otherwise.

### GetFraDestOk

`func (o *OracleDBConfig) GetFraDestOk() (*string, bool)`

GetFraDestOk returns a tuple with the FraDest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraDest

`func (o *OracleDBConfig) SetFraDest(v string)`

SetFraDest sets FraDest field to given value.

### HasFraDest

`func (o *OracleDBConfig) HasFraDest() bool`

HasFraDest returns a boolean if a field has been set.

### SetFraDestNil

`func (o *OracleDBConfig) SetFraDestNil(b bool)`

 SetFraDestNil sets the value for FraDest to be an explicit nil

### UnsetFraDest
`func (o *OracleDBConfig) UnsetFraDest()`

UnsetFraDest ensures that no value is present for FraDest, not even an explicit nil
### GetFraSizeMb

`func (o *OracleDBConfig) GetFraSizeMb() int32`

GetFraSizeMb returns the FraSizeMb field if non-nil, zero value otherwise.

### GetFraSizeMbOk

`func (o *OracleDBConfig) GetFraSizeMbOk() (*int32, bool)`

GetFraSizeMbOk returns a tuple with the FraSizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraSizeMb

`func (o *OracleDBConfig) SetFraSizeMb(v int32)`

SetFraSizeMb sets FraSizeMb field to given value.

### HasFraSizeMb

`func (o *OracleDBConfig) HasFraSizeMb() bool`

HasFraSizeMb returns a boolean if a field has been set.

### SetFraSizeMbNil

`func (o *OracleDBConfig) SetFraSizeMbNil(b bool)`

 SetFraSizeMbNil sets the value for FraSizeMb to be an explicit nil

### UnsetFraSizeMb
`func (o *OracleDBConfig) UnsetFraSizeMb()`

UnsetFraSizeMb ensures that no value is present for FraSizeMb, not even an explicit nil
### GetNumTempfiles

`func (o *OracleDBConfig) GetNumTempfiles() int32`

GetNumTempfiles returns the NumTempfiles field if non-nil, zero value otherwise.

### GetNumTempfilesOk

`func (o *OracleDBConfig) GetNumTempfilesOk() (*int32, bool)`

GetNumTempfilesOk returns a tuple with the NumTempfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTempfiles

`func (o *OracleDBConfig) SetNumTempfiles(v int32)`

SetNumTempfiles sets NumTempfiles field to given value.

### HasNumTempfiles

`func (o *OracleDBConfig) HasNumTempfiles() bool`

HasNumTempfiles returns a boolean if a field has been set.

### SetNumTempfilesNil

`func (o *OracleDBConfig) SetNumTempfilesNil(b bool)`

 SetNumTempfilesNil sets the value for NumTempfiles to be an explicit nil

### UnsetNumTempfiles
`func (o *OracleDBConfig) UnsetNumTempfiles()`

UnsetNumTempfiles ensures that no value is present for NumTempfiles, not even an explicit nil
### GetPfileParameterMap

`func (o *OracleDBConfig) GetPfileParameterMap() []OracleDBConfigPfileParameterMapEntry`

GetPfileParameterMap returns the PfileParameterMap field if non-nil, zero value otherwise.

### GetPfileParameterMapOk

`func (o *OracleDBConfig) GetPfileParameterMapOk() (*[]OracleDBConfigPfileParameterMapEntry, bool)`

GetPfileParameterMapOk returns a tuple with the PfileParameterMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfileParameterMap

`func (o *OracleDBConfig) SetPfileParameterMap(v []OracleDBConfigPfileParameterMapEntry)`

SetPfileParameterMap sets PfileParameterMap field to given value.

### HasPfileParameterMap

`func (o *OracleDBConfig) HasPfileParameterMap() bool`

HasPfileParameterMap returns a boolean if a field has been set.

### SetPfileParameterMapNil

`func (o *OracleDBConfig) SetPfileParameterMapNil(b bool)`

 SetPfileParameterMapNil sets the value for PfileParameterMap to be an explicit nil

### UnsetPfileParameterMap
`func (o *OracleDBConfig) UnsetPfileParameterMap()`

UnsetPfileParameterMap ensures that no value is present for PfileParameterMap, not even an explicit nil
### GetRedoLogConf

`func (o *OracleDBConfig) GetRedoLogConf() OracleDBConfigRedoLogGroupConf`

GetRedoLogConf returns the RedoLogConf field if non-nil, zero value otherwise.

### GetRedoLogConfOk

`func (o *OracleDBConfig) GetRedoLogConfOk() (*OracleDBConfigRedoLogGroupConf, bool)`

GetRedoLogConfOk returns a tuple with the RedoLogConf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedoLogConf

`func (o *OracleDBConfig) SetRedoLogConf(v OracleDBConfigRedoLogGroupConf)`

SetRedoLogConf sets RedoLogConf field to given value.

### HasRedoLogConf

`func (o *OracleDBConfig) HasRedoLogConf() bool`

HasRedoLogConf returns a boolean if a field has been set.

### GetSgaTargetSize

`func (o *OracleDBConfig) GetSgaTargetSize() string`

GetSgaTargetSize returns the SgaTargetSize field if non-nil, zero value otherwise.

### GetSgaTargetSizeOk

`func (o *OracleDBConfig) GetSgaTargetSizeOk() (*string, bool)`

GetSgaTargetSizeOk returns a tuple with the SgaTargetSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgaTargetSize

`func (o *OracleDBConfig) SetSgaTargetSize(v string)`

SetSgaTargetSize sets SgaTargetSize field to given value.

### HasSgaTargetSize

`func (o *OracleDBConfig) HasSgaTargetSize() bool`

HasSgaTargetSize returns a boolean if a field has been set.

### SetSgaTargetSizeNil

`func (o *OracleDBConfig) SetSgaTargetSizeNil(b bool)`

 SetSgaTargetSizeNil sets the value for SgaTargetSize to be an explicit nil

### UnsetSgaTargetSize
`func (o *OracleDBConfig) UnsetSgaTargetSize()`

UnsetSgaTargetSize ensures that no value is present for SgaTargetSize, not even an explicit nil
### GetSharedPoolSize

`func (o *OracleDBConfig) GetSharedPoolSize() string`

GetSharedPoolSize returns the SharedPoolSize field if non-nil, zero value otherwise.

### GetSharedPoolSizeOk

`func (o *OracleDBConfig) GetSharedPoolSizeOk() (*string, bool)`

GetSharedPoolSizeOk returns a tuple with the SharedPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedPoolSize

`func (o *OracleDBConfig) SetSharedPoolSize(v string)`

SetSharedPoolSize sets SharedPoolSize field to given value.

### HasSharedPoolSize

`func (o *OracleDBConfig) HasSharedPoolSize() bool`

HasSharedPoolSize returns a boolean if a field has been set.

### SetSharedPoolSizeNil

`func (o *OracleDBConfig) SetSharedPoolSizeNil(b bool)`

 SetSharedPoolSizeNil sets the value for SharedPoolSize to be an explicit nil

### UnsetSharedPoolSize
`func (o *OracleDBConfig) UnsetSharedPoolSize()`

UnsetSharedPoolSize ensures that no value is present for SharedPoolSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


