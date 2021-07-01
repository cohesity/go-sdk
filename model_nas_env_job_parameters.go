/*
 * Cohesity REST API
 *
 * This API list provides operations for interfacing with the Cohesity Cluster.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cohesitysdk

import (
	"encoding/json"
)

// NasEnvJobParameters Specifies job parameters applicable for all 'kGenericNas' Environment type Protection Sources in a Protection Job.
type NasEnvJobParameters struct {
	// Specifies if the backup should continue on with other Protection Sources even if the backup operation of some Protection Source fails. If true, the Cohesity Cluster ignores the errors and continues with remaining Protection Sources in the job. If false, the backup operation stops when an error occurs. This does not apply to non-snapshot based generic NAS backup jobs. If not set, default value is true.
	ContinueOnError NullableBool `json:"continueOnError,omitempty"`
	DataMigrationJobParameters *DataMigrationJobParameters `json:"dataMigrationJobParameters,omitempty"`
	DataUptierJobParameters *DataUptierJobParameters `json:"dataUptierJobParameters,omitempty"`
	// Specifies whether this job will enable faster incremental backups using change list or similar APIs
	EnableFasterIncrementalBackups NullableBool `json:"enableFasterIncrementalBackups,omitempty"`
	// Specifies if the protection job should use encryption while backing up
	EncryptionEnabled NullableBool `json:"encryptionEnabled,omitempty"`
	FileLockConfig *FileLevelDataLockConfig `json:"fileLockConfig,omitempty"`
	FilePathFilters *FilePathFilter `json:"filePathFilters,omitempty"`
	// Specifies the preferred protocol to use for backup. This does not apply to generic NAS and will be ignored. Specifies the protocol used by a NAS server. 'kNfs3' indicates NFS v3 protocol. 'kCifs1' indicates CIFS v1.0 protocol.
	NasProtocol NullableString `json:"nasProtocol,omitempty"`
	SnapshotLabel *SnapshotLabel `json:"snapshotLabel,omitempty"`
}

// NewNasEnvJobParameters instantiates a new NasEnvJobParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNasEnvJobParameters() *NasEnvJobParameters {
	this := NasEnvJobParameters{}
	return &this
}

// NewNasEnvJobParametersWithDefaults instantiates a new NasEnvJobParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNasEnvJobParametersWithDefaults() *NasEnvJobParameters {
	this := NasEnvJobParameters{}
	return &this
}

// GetContinueOnError returns the ContinueOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NasEnvJobParameters) GetContinueOnError() bool {
	if o == nil || o.ContinueOnError.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ContinueOnError.Get()
}

// GetContinueOnErrorOk returns a tuple with the ContinueOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NasEnvJobParameters) GetContinueOnErrorOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContinueOnError.Get(), o.ContinueOnError.IsSet()
}

// HasContinueOnError returns a boolean if a field has been set.
func (o *NasEnvJobParameters) HasContinueOnError() bool {
	if o != nil && o.ContinueOnError.IsSet() {
		return true
	}

	return false
}

// SetContinueOnError gets a reference to the given NullableBool and assigns it to the ContinueOnError field.
func (o *NasEnvJobParameters) SetContinueOnError(v bool) {
	o.ContinueOnError.Set(&v)
}
// SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil
func (o *NasEnvJobParameters) SetContinueOnErrorNil() {
	o.ContinueOnError.Set(nil)
}

// UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
func (o *NasEnvJobParameters) UnsetContinueOnError() {
	o.ContinueOnError.Unset()
}

// GetDataMigrationJobParameters returns the DataMigrationJobParameters field value if set, zero value otherwise.
func (o *NasEnvJobParameters) GetDataMigrationJobParameters() DataMigrationJobParameters {
	if o == nil || o.DataMigrationJobParameters == nil {
		var ret DataMigrationJobParameters
		return ret
	}
	return *o.DataMigrationJobParameters
}

// GetDataMigrationJobParametersOk returns a tuple with the DataMigrationJobParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NasEnvJobParameters) GetDataMigrationJobParametersOk() (*DataMigrationJobParameters, bool) {
	if o == nil || o.DataMigrationJobParameters == nil {
		return nil, false
	}
	return o.DataMigrationJobParameters, true
}

// HasDataMigrationJobParameters returns a boolean if a field has been set.
func (o *NasEnvJobParameters) HasDataMigrationJobParameters() bool {
	if o != nil && o.DataMigrationJobParameters != nil {
		return true
	}

	return false
}

// SetDataMigrationJobParameters gets a reference to the given DataMigrationJobParameters and assigns it to the DataMigrationJobParameters field.
func (o *NasEnvJobParameters) SetDataMigrationJobParameters(v DataMigrationJobParameters) {
	o.DataMigrationJobParameters = &v
}

// GetDataUptierJobParameters returns the DataUptierJobParameters field value if set, zero value otherwise.
func (o *NasEnvJobParameters) GetDataUptierJobParameters() DataUptierJobParameters {
	if o == nil || o.DataUptierJobParameters == nil {
		var ret DataUptierJobParameters
		return ret
	}
	return *o.DataUptierJobParameters
}

// GetDataUptierJobParametersOk returns a tuple with the DataUptierJobParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NasEnvJobParameters) GetDataUptierJobParametersOk() (*DataUptierJobParameters, bool) {
	if o == nil || o.DataUptierJobParameters == nil {
		return nil, false
	}
	return o.DataUptierJobParameters, true
}

// HasDataUptierJobParameters returns a boolean if a field has been set.
func (o *NasEnvJobParameters) HasDataUptierJobParameters() bool {
	if o != nil && o.DataUptierJobParameters != nil {
		return true
	}

	return false
}

// SetDataUptierJobParameters gets a reference to the given DataUptierJobParameters and assigns it to the DataUptierJobParameters field.
func (o *NasEnvJobParameters) SetDataUptierJobParameters(v DataUptierJobParameters) {
	o.DataUptierJobParameters = &v
}

// GetEnableFasterIncrementalBackups returns the EnableFasterIncrementalBackups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NasEnvJobParameters) GetEnableFasterIncrementalBackups() bool {
	if o == nil || o.EnableFasterIncrementalBackups.Get() == nil {
		var ret bool
		return ret
	}
	return *o.EnableFasterIncrementalBackups.Get()
}

// GetEnableFasterIncrementalBackupsOk returns a tuple with the EnableFasterIncrementalBackups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NasEnvJobParameters) GetEnableFasterIncrementalBackupsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EnableFasterIncrementalBackups.Get(), o.EnableFasterIncrementalBackups.IsSet()
}

// HasEnableFasterIncrementalBackups returns a boolean if a field has been set.
func (o *NasEnvJobParameters) HasEnableFasterIncrementalBackups() bool {
	if o != nil && o.EnableFasterIncrementalBackups.IsSet() {
		return true
	}

	return false
}

// SetEnableFasterIncrementalBackups gets a reference to the given NullableBool and assigns it to the EnableFasterIncrementalBackups field.
func (o *NasEnvJobParameters) SetEnableFasterIncrementalBackups(v bool) {
	o.EnableFasterIncrementalBackups.Set(&v)
}
// SetEnableFasterIncrementalBackupsNil sets the value for EnableFasterIncrementalBackups to be an explicit nil
func (o *NasEnvJobParameters) SetEnableFasterIncrementalBackupsNil() {
	o.EnableFasterIncrementalBackups.Set(nil)
}

// UnsetEnableFasterIncrementalBackups ensures that no value is present for EnableFasterIncrementalBackups, not even an explicit nil
func (o *NasEnvJobParameters) UnsetEnableFasterIncrementalBackups() {
	o.EnableFasterIncrementalBackups.Unset()
}

// GetEncryptionEnabled returns the EncryptionEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NasEnvJobParameters) GetEncryptionEnabled() bool {
	if o == nil || o.EncryptionEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.EncryptionEnabled.Get()
}

// GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NasEnvJobParameters) GetEncryptionEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EncryptionEnabled.Get(), o.EncryptionEnabled.IsSet()
}

// HasEncryptionEnabled returns a boolean if a field has been set.
func (o *NasEnvJobParameters) HasEncryptionEnabled() bool {
	if o != nil && o.EncryptionEnabled.IsSet() {
		return true
	}

	return false
}

// SetEncryptionEnabled gets a reference to the given NullableBool and assigns it to the EncryptionEnabled field.
func (o *NasEnvJobParameters) SetEncryptionEnabled(v bool) {
	o.EncryptionEnabled.Set(&v)
}
// SetEncryptionEnabledNil sets the value for EncryptionEnabled to be an explicit nil
func (o *NasEnvJobParameters) SetEncryptionEnabledNil() {
	o.EncryptionEnabled.Set(nil)
}

// UnsetEncryptionEnabled ensures that no value is present for EncryptionEnabled, not even an explicit nil
func (o *NasEnvJobParameters) UnsetEncryptionEnabled() {
	o.EncryptionEnabled.Unset()
}

// GetFileLockConfig returns the FileLockConfig field value if set, zero value otherwise.
func (o *NasEnvJobParameters) GetFileLockConfig() FileLevelDataLockConfig {
	if o == nil || o.FileLockConfig == nil {
		var ret FileLevelDataLockConfig
		return ret
	}
	return *o.FileLockConfig
}

// GetFileLockConfigOk returns a tuple with the FileLockConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NasEnvJobParameters) GetFileLockConfigOk() (*FileLevelDataLockConfig, bool) {
	if o == nil || o.FileLockConfig == nil {
		return nil, false
	}
	return o.FileLockConfig, true
}

// HasFileLockConfig returns a boolean if a field has been set.
func (o *NasEnvJobParameters) HasFileLockConfig() bool {
	if o != nil && o.FileLockConfig != nil {
		return true
	}

	return false
}

// SetFileLockConfig gets a reference to the given FileLevelDataLockConfig and assigns it to the FileLockConfig field.
func (o *NasEnvJobParameters) SetFileLockConfig(v FileLevelDataLockConfig) {
	o.FileLockConfig = &v
}

// GetFilePathFilters returns the FilePathFilters field value if set, zero value otherwise.
func (o *NasEnvJobParameters) GetFilePathFilters() FilePathFilter {
	if o == nil || o.FilePathFilters == nil {
		var ret FilePathFilter
		return ret
	}
	return *o.FilePathFilters
}

// GetFilePathFiltersOk returns a tuple with the FilePathFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NasEnvJobParameters) GetFilePathFiltersOk() (*FilePathFilter, bool) {
	if o == nil || o.FilePathFilters == nil {
		return nil, false
	}
	return o.FilePathFilters, true
}

// HasFilePathFilters returns a boolean if a field has been set.
func (o *NasEnvJobParameters) HasFilePathFilters() bool {
	if o != nil && o.FilePathFilters != nil {
		return true
	}

	return false
}

// SetFilePathFilters gets a reference to the given FilePathFilter and assigns it to the FilePathFilters field.
func (o *NasEnvJobParameters) SetFilePathFilters(v FilePathFilter) {
	o.FilePathFilters = &v
}

// GetNasProtocol returns the NasProtocol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NasEnvJobParameters) GetNasProtocol() string {
	if o == nil || o.NasProtocol.Get() == nil {
		var ret string
		return ret
	}
	return *o.NasProtocol.Get()
}

// GetNasProtocolOk returns a tuple with the NasProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NasEnvJobParameters) GetNasProtocolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NasProtocol.Get(), o.NasProtocol.IsSet()
}

// HasNasProtocol returns a boolean if a field has been set.
func (o *NasEnvJobParameters) HasNasProtocol() bool {
	if o != nil && o.NasProtocol.IsSet() {
		return true
	}

	return false
}

// SetNasProtocol gets a reference to the given NullableString and assigns it to the NasProtocol field.
func (o *NasEnvJobParameters) SetNasProtocol(v string) {
	o.NasProtocol.Set(&v)
}
// SetNasProtocolNil sets the value for NasProtocol to be an explicit nil
func (o *NasEnvJobParameters) SetNasProtocolNil() {
	o.NasProtocol.Set(nil)
}

// UnsetNasProtocol ensures that no value is present for NasProtocol, not even an explicit nil
func (o *NasEnvJobParameters) UnsetNasProtocol() {
	o.NasProtocol.Unset()
}

// GetSnapshotLabel returns the SnapshotLabel field value if set, zero value otherwise.
func (o *NasEnvJobParameters) GetSnapshotLabel() SnapshotLabel {
	if o == nil || o.SnapshotLabel == nil {
		var ret SnapshotLabel
		return ret
	}
	return *o.SnapshotLabel
}

// GetSnapshotLabelOk returns a tuple with the SnapshotLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NasEnvJobParameters) GetSnapshotLabelOk() (*SnapshotLabel, bool) {
	if o == nil || o.SnapshotLabel == nil {
		return nil, false
	}
	return o.SnapshotLabel, true
}

// HasSnapshotLabel returns a boolean if a field has been set.
func (o *NasEnvJobParameters) HasSnapshotLabel() bool {
	if o != nil && o.SnapshotLabel != nil {
		return true
	}

	return false
}

// SetSnapshotLabel gets a reference to the given SnapshotLabel and assigns it to the SnapshotLabel field.
func (o *NasEnvJobParameters) SetSnapshotLabel(v SnapshotLabel) {
	o.SnapshotLabel = &v
}

func (o NasEnvJobParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContinueOnError.IsSet() {
		toSerialize["continueOnError"] = o.ContinueOnError.Get()
	}
	if o.DataMigrationJobParameters != nil {
		toSerialize["dataMigrationJobParameters"] = o.DataMigrationJobParameters
	}
	if o.DataUptierJobParameters != nil {
		toSerialize["dataUptierJobParameters"] = o.DataUptierJobParameters
	}
	if o.EnableFasterIncrementalBackups.IsSet() {
		toSerialize["enableFasterIncrementalBackups"] = o.EnableFasterIncrementalBackups.Get()
	}
	if o.EncryptionEnabled.IsSet() {
		toSerialize["encryptionEnabled"] = o.EncryptionEnabled.Get()
	}
	if o.FileLockConfig != nil {
		toSerialize["fileLockConfig"] = o.FileLockConfig
	}
	if o.FilePathFilters != nil {
		toSerialize["filePathFilters"] = o.FilePathFilters
	}
	if o.NasProtocol.IsSet() {
		toSerialize["nasProtocol"] = o.NasProtocol.Get()
	}
	if o.SnapshotLabel != nil {
		toSerialize["snapshotLabel"] = o.SnapshotLabel
	}
	return json.Marshal(toSerialize)
}

type NullableNasEnvJobParameters struct {
	value *NasEnvJobParameters
	isSet bool
}

func (v NullableNasEnvJobParameters) Get() *NasEnvJobParameters {
	return v.value
}

func (v *NullableNasEnvJobParameters) Set(val *NasEnvJobParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableNasEnvJobParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableNasEnvJobParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNasEnvJobParameters(val *NasEnvJobParameters) *NullableNasEnvJobParameters {
	return &NullableNasEnvJobParameters{value: val, isSet: true}
}

func (v NullableNasEnvJobParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNasEnvJobParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


