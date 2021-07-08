# NasEnvJobParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies if the backup should continue on with other Protection Sources even if the backup operation of some Protection Source fails. If true, the Cohesity Cluster ignores the errors and continues with remaining Protection Sources in the job. If false, the backup operation stops when an error occurs. This does not apply to non-snapshot based generic NAS backup jobs. If not set, default value is true. | [optional] 
**DataMigrationJobParameters** | Pointer to [**DataMigrationJobParameters**](DataMigrationJobParameters.md) |  | [optional] 
**DataUptierJobParameters** | Pointer to [**DataUptierJobParameters**](DataUptierJobParameters.md) |  | [optional] 
**EnableFasterIncrementalBackups** | Pointer to **NullableBool** | Specifies whether this job will enable faster incremental backups using change list or similar APIs | [optional] 
**EncryptionEnabled** | Pointer to **NullableBool** | Specifies if the protection job should use encryption while backing up | [optional] 
**FileLockConfig** | Pointer to [**FileLevelDataLockConfig**](FileLevelDataLockConfig.md) |  | [optional] 
**FilePathFilters** | Pointer to [**FilePathFilter**](FilePathFilter.md) |  | [optional] 
**NasProtocol** | Pointer to **NullableString** | Specifies the preferred protocol to use for backup. This does not apply to generic NAS and will be ignored. Specifies the protocol used by a NAS server. &#39;kNfs3&#39; indicates NFS v3 protocol. &#39;kCifs1&#39; indicates CIFS v1.0 protocol. | [optional] 
**SnapshotLabel** | Pointer to [**SnapshotLabel**](SnapshotLabel.md) |  | [optional] 

## Methods

### NewNasEnvJobParameters

`func NewNasEnvJobParameters() *NasEnvJobParameters`

NewNasEnvJobParameters instantiates a new NasEnvJobParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNasEnvJobParametersWithDefaults

`func NewNasEnvJobParametersWithDefaults() *NasEnvJobParameters`

NewNasEnvJobParametersWithDefaults instantiates a new NasEnvJobParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *NasEnvJobParameters) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *NasEnvJobParameters) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *NasEnvJobParameters) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *NasEnvJobParameters) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *NasEnvJobParameters) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *NasEnvJobParameters) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetDataMigrationJobParameters

`func (o *NasEnvJobParameters) GetDataMigrationJobParameters() DataMigrationJobParameters`

GetDataMigrationJobParameters returns the DataMigrationJobParameters field if non-nil, zero value otherwise.

### GetDataMigrationJobParametersOk

`func (o *NasEnvJobParameters) GetDataMigrationJobParametersOk() (*DataMigrationJobParameters, bool)`

GetDataMigrationJobParametersOk returns a tuple with the DataMigrationJobParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataMigrationJobParameters

`func (o *NasEnvJobParameters) SetDataMigrationJobParameters(v DataMigrationJobParameters)`

SetDataMigrationJobParameters sets DataMigrationJobParameters field to given value.

### HasDataMigrationJobParameters

`func (o *NasEnvJobParameters) HasDataMigrationJobParameters() bool`

HasDataMigrationJobParameters returns a boolean if a field has been set.

### GetDataUptierJobParameters

`func (o *NasEnvJobParameters) GetDataUptierJobParameters() DataUptierJobParameters`

GetDataUptierJobParameters returns the DataUptierJobParameters field if non-nil, zero value otherwise.

### GetDataUptierJobParametersOk

`func (o *NasEnvJobParameters) GetDataUptierJobParametersOk() (*DataUptierJobParameters, bool)`

GetDataUptierJobParametersOk returns a tuple with the DataUptierJobParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataUptierJobParameters

`func (o *NasEnvJobParameters) SetDataUptierJobParameters(v DataUptierJobParameters)`

SetDataUptierJobParameters sets DataUptierJobParameters field to given value.

### HasDataUptierJobParameters

`func (o *NasEnvJobParameters) HasDataUptierJobParameters() bool`

HasDataUptierJobParameters returns a boolean if a field has been set.

### GetEnableFasterIncrementalBackups

`func (o *NasEnvJobParameters) GetEnableFasterIncrementalBackups() bool`

GetEnableFasterIncrementalBackups returns the EnableFasterIncrementalBackups field if non-nil, zero value otherwise.

### GetEnableFasterIncrementalBackupsOk

`func (o *NasEnvJobParameters) GetEnableFasterIncrementalBackupsOk() (*bool, bool)`

GetEnableFasterIncrementalBackupsOk returns a tuple with the EnableFasterIncrementalBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFasterIncrementalBackups

`func (o *NasEnvJobParameters) SetEnableFasterIncrementalBackups(v bool)`

SetEnableFasterIncrementalBackups sets EnableFasterIncrementalBackups field to given value.

### HasEnableFasterIncrementalBackups

`func (o *NasEnvJobParameters) HasEnableFasterIncrementalBackups() bool`

HasEnableFasterIncrementalBackups returns a boolean if a field has been set.

### SetEnableFasterIncrementalBackupsNil

`func (o *NasEnvJobParameters) SetEnableFasterIncrementalBackupsNil(b bool)`

 SetEnableFasterIncrementalBackupsNil sets the value for EnableFasterIncrementalBackups to be an explicit nil

### UnsetEnableFasterIncrementalBackups
`func (o *NasEnvJobParameters) UnsetEnableFasterIncrementalBackups()`

UnsetEnableFasterIncrementalBackups ensures that no value is present for EnableFasterIncrementalBackups, not even an explicit nil
### GetEncryptionEnabled

`func (o *NasEnvJobParameters) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *NasEnvJobParameters) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *NasEnvJobParameters) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *NasEnvJobParameters) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### SetEncryptionEnabledNil

`func (o *NasEnvJobParameters) SetEncryptionEnabledNil(b bool)`

 SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil

### UnsetEncryptionEnabled
`func (o *NasEnvJobParameters) UnsetEncryptionEnabled()`

UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
### GetFileLockConfig

`func (o *NasEnvJobParameters) GetFileLockConfig() FileLevelDataLockConfig`

GetFileLockConfig returns the FileLockConfig field if non-nil, zero value otherwise.

### GetFileLockConfigOk

`func (o *NasEnvJobParameters) GetFileLockConfigOk() (*FileLevelDataLockConfig, bool)`

GetFileLockConfigOk returns a tuple with the FileLockConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLockConfig

`func (o *NasEnvJobParameters) SetFileLockConfig(v FileLevelDataLockConfig)`

SetFileLockConfig sets FileLockConfig field to given value.

### HasFileLockConfig

`func (o *NasEnvJobParameters) HasFileLockConfig() bool`

HasFileLockConfig returns a boolean if a field has been set.

### GetFilePathFilters

`func (o *NasEnvJobParameters) GetFilePathFilters() FilePathFilter`

GetFilePathFilters returns the FilePathFilters field if non-nil, zero value otherwise.

### GetFilePathFiltersOk

`func (o *NasEnvJobParameters) GetFilePathFiltersOk() (*FilePathFilter, bool)`

GetFilePathFiltersOk returns a tuple with the FilePathFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePathFilters

`func (o *NasEnvJobParameters) SetFilePathFilters(v FilePathFilter)`

SetFilePathFilters sets FilePathFilters field to given value.

### HasFilePathFilters

`func (o *NasEnvJobParameters) HasFilePathFilters() bool`

HasFilePathFilters returns a boolean if a field has been set.

### GetNasProtocol

`func (o *NasEnvJobParameters) GetNasProtocol() string`

GetNasProtocol returns the NasProtocol field if non-nil, zero value otherwise.

### GetNasProtocolOk

`func (o *NasEnvJobParameters) GetNasProtocolOk() (*string, bool)`

GetNasProtocolOk returns a tuple with the NasProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasProtocol

`func (o *NasEnvJobParameters) SetNasProtocol(v string)`

SetNasProtocol sets NasProtocol field to given value.

### HasNasProtocol

`func (o *NasEnvJobParameters) HasNasProtocol() bool`

HasNasProtocol returns a boolean if a field has been set.

### SetNasProtocolNil

`func (o *NasEnvJobParameters) SetNasProtocolNil(b bool)`

 SetNasProtocolNil sets the value for NasProtocol to be an explicit nil

### UnsetNasProtocol
`func (o *NasEnvJobParameters) UnsetNasProtocol()`

UnsetNasProtocol ensures that no value is present for NasProtocol, not even an explicit nil
### GetSnapshotLabel

`func (o *NasEnvJobParameters) GetSnapshotLabel() SnapshotLabel`

GetSnapshotLabel returns the SnapshotLabel field if non-nil, zero value otherwise.

### GetSnapshotLabelOk

`func (o *NasEnvJobParameters) GetSnapshotLabelOk() (*SnapshotLabel, bool)`

GetSnapshotLabelOk returns a tuple with the SnapshotLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotLabel

`func (o *NasEnvJobParameters) SetSnapshotLabel(v SnapshotLabel)`

SetSnapshotLabel sets SnapshotLabel field to given value.

### HasSnapshotLabel

`func (o *NasEnvJobParameters) HasSnapshotLabel() bool`

HasSnapshotLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


