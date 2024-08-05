/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CommonRecoverSqlAppTargetParamsNewSourceConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonRecoverSqlAppTargetParamsNewSourceConfig{}

// CommonRecoverSqlAppTargetParamsNewSourceConfig Specifies the destination Source configuration parameters where the databases will be recovered. This is mandatory if recoverToNewSource is set to true.
type CommonRecoverSqlAppTargetParamsNewSourceConfig struct {
	// Specifies the directory where to put the database data files. Missing directory will be automatically created.
	DataFileDirectoryLocation NullableString `json:"dataFileDirectoryLocation"`
	// Specifies a new name for the restored database. If this field is not specified, then the original database will be overwritten after recovery.
	DatabaseName NullableString `json:"databaseName,omitempty"`
	// Specifies the source id of target host where databases will be recovered. This source id can be a physical host or virtual machine.
	Host NullableRecoveryObjectIdentifier `json:"host"`
	// Specifies an instance name of the Sql Server that should be used for restoring databases to.
	InstanceName NullableString `json:"instanceName"`
	// Specifies the directory where to put the database log files. Missing directory will be automatically created.
	LogFileDirectoryLocation NullableString `json:"logFileDirectoryLocation"`
	// Specifies whether to keep CDC (Change Data Capture) on recovered databases or not. If not passed, this is assumed to be true. If withNoRecovery is passed as true, then this field must not be set to true. Passing this field as true in this scenario will be a invalid request.
	KeepCdc NullableBool `json:"keepCdc,omitempty"`
	MultiStageRestoreOptions *MultiStageRestoreOptions `json:"multiStageRestoreOptions,omitempty"`
	// Specifies the WITH clause to be used in native sql log restore command. This is only applicable for native log restore.
	NativeLogRecoveryWithClause NullableString `json:"nativeLogRecoveryWithClause,omitempty"`
	// 'with_clause' contains 'with clause' to be used in native sql restore command. This is only applicable for database restore of native sql backup. Here user can specify multiple restore options. Example: 'WITH BUFFERCOUNT = 575, MAXTRANSFERSIZE = 2097152'.
	NativeRecoveryWithClause NullableString `json:"nativeRecoveryWithClause,omitempty"`
	// Specifies a policy to be used while recovering existing databases.
	OverwritingPolicy NullableString `json:"overwritingPolicy,omitempty"`
	// Specifies the option to set replay last log bit while creating the sql restore task and doing restore to latest point-in-time. If this is set to true, we will replay the entire last log without STOPAT.
	ReplayEntireLastLog NullableBool `json:"replayEntireLastLog,omitempty"`
	// Specifies the time in the past to which the Sql database needs to be restored. This allows for granular recovery of Sql databases. If this is not set, the Sql database will be restored from the full/incremental snapshot.
	RestoreTimeUsecs NullableInt64 `json:"restoreTimeUsecs,omitempty"`
	// Specifies the secondary data filename pattern and corresponding direcories of the DB. Secondary data files are optional and are user defined. The recommended file extention for secondary files is \".ndf\". If this option is specified and the destination folders do not exist they will be automatically created.
	SecondaryDataFilesDirList []FilenamePatternToDirectory `json:"secondaryDataFilesDirList,omitempty"`
	// Specifies the flag to bring DBs online or not after successful recovery. If this is passed as true, then it means DBs won't be brought online.
	WithNoRecovery NullableBool `json:"withNoRecovery,omitempty"`
}

type _CommonRecoverSqlAppTargetParamsNewSourceConfig CommonRecoverSqlAppTargetParamsNewSourceConfig

// NewCommonRecoverSqlAppTargetParamsNewSourceConfig instantiates a new CommonRecoverSqlAppTargetParamsNewSourceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonRecoverSqlAppTargetParamsNewSourceConfig(dataFileDirectoryLocation NullableString, host NullableRecoveryObjectIdentifier, instanceName NullableString, logFileDirectoryLocation NullableString) *CommonRecoverSqlAppTargetParamsNewSourceConfig {
	this := CommonRecoverSqlAppTargetParamsNewSourceConfig{}
	this.DataFileDirectoryLocation = dataFileDirectoryLocation
	this.Host = host
	this.InstanceName = instanceName
	this.LogFileDirectoryLocation = logFileDirectoryLocation
	return &this
}

// NewCommonRecoverSqlAppTargetParamsNewSourceConfigWithDefaults instantiates a new CommonRecoverSqlAppTargetParamsNewSourceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonRecoverSqlAppTargetParamsNewSourceConfigWithDefaults() *CommonRecoverSqlAppTargetParamsNewSourceConfig {
	this := CommonRecoverSqlAppTargetParamsNewSourceConfig{}
	return &this
}

// GetDataFileDirectoryLocation returns the DataFileDirectoryLocation field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetDataFileDirectoryLocation() string {
	if o == nil || o.DataFileDirectoryLocation.Get() == nil {
		var ret string
		return ret
	}

	return *o.DataFileDirectoryLocation.Get()
}

// GetDataFileDirectoryLocationOk returns a tuple with the DataFileDirectoryLocation field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetDataFileDirectoryLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataFileDirectoryLocation.Get(), o.DataFileDirectoryLocation.IsSet()
}

// SetDataFileDirectoryLocation sets field value
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetDataFileDirectoryLocation(v string) {
	o.DataFileDirectoryLocation.Set(&v)
}

// GetDatabaseName returns the DatabaseName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetDatabaseName() string {
	if o == nil || IsNil(o.DatabaseName.Get()) {
		var ret string
		return ret
	}
	return *o.DatabaseName.Get()
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetDatabaseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatabaseName.Get(), o.DatabaseName.IsSet()
}

// HasDatabaseName returns a boolean if a field has been set.
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) HasDatabaseName() bool {
	if o != nil && o.DatabaseName.IsSet() {
		return true
	}

	return false
}

// SetDatabaseName gets a reference to the given NullableString and assigns it to the DatabaseName field.
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetDatabaseName(v string) {
	o.DatabaseName.Set(&v)
}
// SetDatabaseNameNil sets the value for DatabaseName to be an explicit nil
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetDatabaseNameNil() {
	o.DatabaseName.Set(nil)
}

// UnsetDatabaseName ensures that no value is present for DatabaseName, not even an explicit nil
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) UnsetDatabaseName() {
	o.DatabaseName.Unset()
}

// GetHost returns the Host field value
// If the value is explicit nil, the zero value for RecoveryObjectIdentifier will be returned
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetHost() RecoveryObjectIdentifier {
	if o == nil || o.Host.Get() == nil {
		var ret RecoveryObjectIdentifier
		return ret
	}

	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetHostOk() (*RecoveryObjectIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// SetHost sets field value
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetHost(v RecoveryObjectIdentifier) {
	o.Host.Set(&v)
}

// GetInstanceName returns the InstanceName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetInstanceName() string {
	if o == nil || o.InstanceName.Get() == nil {
		var ret string
		return ret
	}

	return *o.InstanceName.Get()
}

// GetInstanceNameOk returns a tuple with the InstanceName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetInstanceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstanceName.Get(), o.InstanceName.IsSet()
}

// SetInstanceName sets field value
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetInstanceName(v string) {
	o.InstanceName.Set(&v)
}

// GetLogFileDirectoryLocation returns the LogFileDirectoryLocation field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetLogFileDirectoryLocation() string {
	if o == nil || o.LogFileDirectoryLocation.Get() == nil {
		var ret string
		return ret
	}

	return *o.LogFileDirectoryLocation.Get()
}

// GetLogFileDirectoryLocationOk returns a tuple with the LogFileDirectoryLocation field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetLogFileDirectoryLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogFileDirectoryLocation.Get(), o.LogFileDirectoryLocation.IsSet()
}

// SetLogFileDirectoryLocation sets field value
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetLogFileDirectoryLocation(v string) {
	o.LogFileDirectoryLocation.Set(&v)
}

// GetKeepCdc returns the KeepCdc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetKeepCdc() bool {
	if o == nil || IsNil(o.KeepCdc.Get()) {
		var ret bool
		return ret
	}
	return *o.KeepCdc.Get()
}

// GetKeepCdcOk returns a tuple with the KeepCdc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetKeepCdcOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeepCdc.Get(), o.KeepCdc.IsSet()
}

// HasKeepCdc returns a boolean if a field has been set.
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) HasKeepCdc() bool {
	if o != nil && o.KeepCdc.IsSet() {
		return true
	}

	return false
}

// SetKeepCdc gets a reference to the given NullableBool and assigns it to the KeepCdc field.
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetKeepCdc(v bool) {
	o.KeepCdc.Set(&v)
}
// SetKeepCdcNil sets the value for KeepCdc to be an explicit nil
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetKeepCdcNil() {
	o.KeepCdc.Set(nil)
}

// UnsetKeepCdc ensures that no value is present for KeepCdc, not even an explicit nil
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) UnsetKeepCdc() {
	o.KeepCdc.Unset()
}

// GetMultiStageRestoreOptions returns the MultiStageRestoreOptions field value if set, zero value otherwise.
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetMultiStageRestoreOptions() MultiStageRestoreOptions {
	if o == nil || IsNil(o.MultiStageRestoreOptions) {
		var ret MultiStageRestoreOptions
		return ret
	}
	return *o.MultiStageRestoreOptions
}

// GetMultiStageRestoreOptionsOk returns a tuple with the MultiStageRestoreOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetMultiStageRestoreOptionsOk() (*MultiStageRestoreOptions, bool) {
	if o == nil || IsNil(o.MultiStageRestoreOptions) {
		return nil, false
	}
	return o.MultiStageRestoreOptions, true
}

// HasMultiStageRestoreOptions returns a boolean if a field has been set.
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) HasMultiStageRestoreOptions() bool {
	if o != nil && !IsNil(o.MultiStageRestoreOptions) {
		return true
	}

	return false
}

// SetMultiStageRestoreOptions gets a reference to the given MultiStageRestoreOptions and assigns it to the MultiStageRestoreOptions field.
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetMultiStageRestoreOptions(v MultiStageRestoreOptions) {
	o.MultiStageRestoreOptions = &v
}

// GetNativeLogRecoveryWithClause returns the NativeLogRecoveryWithClause field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetNativeLogRecoveryWithClause() string {
	if o == nil || IsNil(o.NativeLogRecoveryWithClause.Get()) {
		var ret string
		return ret
	}
	return *o.NativeLogRecoveryWithClause.Get()
}

// GetNativeLogRecoveryWithClauseOk returns a tuple with the NativeLogRecoveryWithClause field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetNativeLogRecoveryWithClauseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NativeLogRecoveryWithClause.Get(), o.NativeLogRecoveryWithClause.IsSet()
}

// HasNativeLogRecoveryWithClause returns a boolean if a field has been set.
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) HasNativeLogRecoveryWithClause() bool {
	if o != nil && o.NativeLogRecoveryWithClause.IsSet() {
		return true
	}

	return false
}

// SetNativeLogRecoveryWithClause gets a reference to the given NullableString and assigns it to the NativeLogRecoveryWithClause field.
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetNativeLogRecoveryWithClause(v string) {
	o.NativeLogRecoveryWithClause.Set(&v)
}
// SetNativeLogRecoveryWithClauseNil sets the value for NativeLogRecoveryWithClause to be an explicit nil
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetNativeLogRecoveryWithClauseNil() {
	o.NativeLogRecoveryWithClause.Set(nil)
}

// UnsetNativeLogRecoveryWithClause ensures that no value is present for NativeLogRecoveryWithClause, not even an explicit nil
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) UnsetNativeLogRecoveryWithClause() {
	o.NativeLogRecoveryWithClause.Unset()
}

// GetNativeRecoveryWithClause returns the NativeRecoveryWithClause field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetNativeRecoveryWithClause() string {
	if o == nil || IsNil(o.NativeRecoveryWithClause.Get()) {
		var ret string
		return ret
	}
	return *o.NativeRecoveryWithClause.Get()
}

// GetNativeRecoveryWithClauseOk returns a tuple with the NativeRecoveryWithClause field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetNativeRecoveryWithClauseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NativeRecoveryWithClause.Get(), o.NativeRecoveryWithClause.IsSet()
}

// HasNativeRecoveryWithClause returns a boolean if a field has been set.
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) HasNativeRecoveryWithClause() bool {
	if o != nil && o.NativeRecoveryWithClause.IsSet() {
		return true
	}

	return false
}

// SetNativeRecoveryWithClause gets a reference to the given NullableString and assigns it to the NativeRecoveryWithClause field.
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetNativeRecoveryWithClause(v string) {
	o.NativeRecoveryWithClause.Set(&v)
}
// SetNativeRecoveryWithClauseNil sets the value for NativeRecoveryWithClause to be an explicit nil
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetNativeRecoveryWithClauseNil() {
	o.NativeRecoveryWithClause.Set(nil)
}

// UnsetNativeRecoveryWithClause ensures that no value is present for NativeRecoveryWithClause, not even an explicit nil
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) UnsetNativeRecoveryWithClause() {
	o.NativeRecoveryWithClause.Unset()
}

// GetOverwritingPolicy returns the OverwritingPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetOverwritingPolicy() string {
	if o == nil || IsNil(o.OverwritingPolicy.Get()) {
		var ret string
		return ret
	}
	return *o.OverwritingPolicy.Get()
}

// GetOverwritingPolicyOk returns a tuple with the OverwritingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetOverwritingPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OverwritingPolicy.Get(), o.OverwritingPolicy.IsSet()
}

// HasOverwritingPolicy returns a boolean if a field has been set.
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) HasOverwritingPolicy() bool {
	if o != nil && o.OverwritingPolicy.IsSet() {
		return true
	}

	return false
}

// SetOverwritingPolicy gets a reference to the given NullableString and assigns it to the OverwritingPolicy field.
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetOverwritingPolicy(v string) {
	o.OverwritingPolicy.Set(&v)
}
// SetOverwritingPolicyNil sets the value for OverwritingPolicy to be an explicit nil
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetOverwritingPolicyNil() {
	o.OverwritingPolicy.Set(nil)
}

// UnsetOverwritingPolicy ensures that no value is present for OverwritingPolicy, not even an explicit nil
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) UnsetOverwritingPolicy() {
	o.OverwritingPolicy.Unset()
}

// GetReplayEntireLastLog returns the ReplayEntireLastLog field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetReplayEntireLastLog() bool {
	if o == nil || IsNil(o.ReplayEntireLastLog.Get()) {
		var ret bool
		return ret
	}
	return *o.ReplayEntireLastLog.Get()
}

// GetReplayEntireLastLogOk returns a tuple with the ReplayEntireLastLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetReplayEntireLastLogOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReplayEntireLastLog.Get(), o.ReplayEntireLastLog.IsSet()
}

// HasReplayEntireLastLog returns a boolean if a field has been set.
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) HasReplayEntireLastLog() bool {
	if o != nil && o.ReplayEntireLastLog.IsSet() {
		return true
	}

	return false
}

// SetReplayEntireLastLog gets a reference to the given NullableBool and assigns it to the ReplayEntireLastLog field.
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetReplayEntireLastLog(v bool) {
	o.ReplayEntireLastLog.Set(&v)
}
// SetReplayEntireLastLogNil sets the value for ReplayEntireLastLog to be an explicit nil
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetReplayEntireLastLogNil() {
	o.ReplayEntireLastLog.Set(nil)
}

// UnsetReplayEntireLastLog ensures that no value is present for ReplayEntireLastLog, not even an explicit nil
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) UnsetReplayEntireLastLog() {
	o.ReplayEntireLastLog.Unset()
}

// GetRestoreTimeUsecs returns the RestoreTimeUsecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetRestoreTimeUsecs() int64 {
	if o == nil || IsNil(o.RestoreTimeUsecs.Get()) {
		var ret int64
		return ret
	}
	return *o.RestoreTimeUsecs.Get()
}

// GetRestoreTimeUsecsOk returns a tuple with the RestoreTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetRestoreTimeUsecsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RestoreTimeUsecs.Get(), o.RestoreTimeUsecs.IsSet()
}

// HasRestoreTimeUsecs returns a boolean if a field has been set.
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) HasRestoreTimeUsecs() bool {
	if o != nil && o.RestoreTimeUsecs.IsSet() {
		return true
	}

	return false
}

// SetRestoreTimeUsecs gets a reference to the given NullableInt64 and assigns it to the RestoreTimeUsecs field.
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetRestoreTimeUsecs(v int64) {
	o.RestoreTimeUsecs.Set(&v)
}
// SetRestoreTimeUsecsNil sets the value for RestoreTimeUsecs to be an explicit nil
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetRestoreTimeUsecsNil() {
	o.RestoreTimeUsecs.Set(nil)
}

// UnsetRestoreTimeUsecs ensures that no value is present for RestoreTimeUsecs, not even an explicit nil
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) UnsetRestoreTimeUsecs() {
	o.RestoreTimeUsecs.Unset()
}

// GetSecondaryDataFilesDirList returns the SecondaryDataFilesDirList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetSecondaryDataFilesDirList() []FilenamePatternToDirectory {
	if o == nil {
		var ret []FilenamePatternToDirectory
		return ret
	}
	return o.SecondaryDataFilesDirList
}

// GetSecondaryDataFilesDirListOk returns a tuple with the SecondaryDataFilesDirList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetSecondaryDataFilesDirListOk() ([]FilenamePatternToDirectory, bool) {
	if o == nil || IsNil(o.SecondaryDataFilesDirList) {
		return nil, false
	}
	return o.SecondaryDataFilesDirList, true
}

// HasSecondaryDataFilesDirList returns a boolean if a field has been set.
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) HasSecondaryDataFilesDirList() bool {
	if o != nil && !IsNil(o.SecondaryDataFilesDirList) {
		return true
	}

	return false
}

// SetSecondaryDataFilesDirList gets a reference to the given []FilenamePatternToDirectory and assigns it to the SecondaryDataFilesDirList field.
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetSecondaryDataFilesDirList(v []FilenamePatternToDirectory) {
	o.SecondaryDataFilesDirList = v
}

// GetWithNoRecovery returns the WithNoRecovery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetWithNoRecovery() bool {
	if o == nil || IsNil(o.WithNoRecovery.Get()) {
		var ret bool
		return ret
	}
	return *o.WithNoRecovery.Get()
}

// GetWithNoRecoveryOk returns a tuple with the WithNoRecovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) GetWithNoRecoveryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WithNoRecovery.Get(), o.WithNoRecovery.IsSet()
}

// HasWithNoRecovery returns a boolean if a field has been set.
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) HasWithNoRecovery() bool {
	if o != nil && o.WithNoRecovery.IsSet() {
		return true
	}

	return false
}

// SetWithNoRecovery gets a reference to the given NullableBool and assigns it to the WithNoRecovery field.
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetWithNoRecovery(v bool) {
	o.WithNoRecovery.Set(&v)
}
// SetWithNoRecoveryNil sets the value for WithNoRecovery to be an explicit nil
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) SetWithNoRecoveryNil() {
	o.WithNoRecovery.Set(nil)
}

// UnsetWithNoRecovery ensures that no value is present for WithNoRecovery, not even an explicit nil
func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) UnsetWithNoRecovery() {
	o.WithNoRecovery.Unset()
}

func (o CommonRecoverSqlAppTargetParamsNewSourceConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonRecoverSqlAppTargetParamsNewSourceConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dataFileDirectoryLocation"] = o.DataFileDirectoryLocation.Get()
	if o.DatabaseName.IsSet() {
		toSerialize["databaseName"] = o.DatabaseName.Get()
	}
	toSerialize["host"] = o.Host.Get()
	toSerialize["instanceName"] = o.InstanceName.Get()
	toSerialize["logFileDirectoryLocation"] = o.LogFileDirectoryLocation.Get()
	if o.KeepCdc.IsSet() {
		toSerialize["keepCdc"] = o.KeepCdc.Get()
	}
	if !IsNil(o.MultiStageRestoreOptions) {
		toSerialize["multiStageRestoreOptions"] = o.MultiStageRestoreOptions
	}
	if o.NativeLogRecoveryWithClause.IsSet() {
		toSerialize["nativeLogRecoveryWithClause"] = o.NativeLogRecoveryWithClause.Get()
	}
	if o.NativeRecoveryWithClause.IsSet() {
		toSerialize["nativeRecoveryWithClause"] = o.NativeRecoveryWithClause.Get()
	}
	if o.OverwritingPolicy.IsSet() {
		toSerialize["overwritingPolicy"] = o.OverwritingPolicy.Get()
	}
	if o.ReplayEntireLastLog.IsSet() {
		toSerialize["replayEntireLastLog"] = o.ReplayEntireLastLog.Get()
	}
	if o.RestoreTimeUsecs.IsSet() {
		toSerialize["restoreTimeUsecs"] = o.RestoreTimeUsecs.Get()
	}
	if o.SecondaryDataFilesDirList != nil {
		toSerialize["secondaryDataFilesDirList"] = o.SecondaryDataFilesDirList
	}
	if o.WithNoRecovery.IsSet() {
		toSerialize["withNoRecovery"] = o.WithNoRecovery.Get()
	}
	return toSerialize, nil
}

func (o *CommonRecoverSqlAppTargetParamsNewSourceConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dataFileDirectoryLocation",
		"host",
		"instanceName",
		"logFileDirectoryLocation",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCommonRecoverSqlAppTargetParamsNewSourceConfig := _CommonRecoverSqlAppTargetParamsNewSourceConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommonRecoverSqlAppTargetParamsNewSourceConfig)

	if err != nil {
		return err
	}

	*o = CommonRecoverSqlAppTargetParamsNewSourceConfig(varCommonRecoverSqlAppTargetParamsNewSourceConfig)

	return err
}

type NullableCommonRecoverSqlAppTargetParamsNewSourceConfig struct {
	value *CommonRecoverSqlAppTargetParamsNewSourceConfig
	isSet bool
}

func (v NullableCommonRecoverSqlAppTargetParamsNewSourceConfig) Get() *CommonRecoverSqlAppTargetParamsNewSourceConfig {
	return v.value
}

func (v *NullableCommonRecoverSqlAppTargetParamsNewSourceConfig) Set(val *CommonRecoverSqlAppTargetParamsNewSourceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonRecoverSqlAppTargetParamsNewSourceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonRecoverSqlAppTargetParamsNewSourceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonRecoverSqlAppTargetParamsNewSourceConfig(val *CommonRecoverSqlAppTargetParamsNewSourceConfig) *NullableCommonRecoverSqlAppTargetParamsNewSourceConfig {
	return &NullableCommonRecoverSqlAppTargetParamsNewSourceConfig{value: val, isSet: true}
}

func (v NullableCommonRecoverSqlAppTargetParamsNewSourceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonRecoverSqlAppTargetParamsNewSourceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


