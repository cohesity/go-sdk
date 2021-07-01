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

// DataMigrationJobParameters Specifies parameters applicable for data migration jobs in NAS environment.
type DataMigrationJobParameters struct {
	// Identifies the cold files in the NAS source. Files that haven't been accessed/modified in the last cold_file_window are migrated.
	ColdFileWindow NullableInt64 `json:"coldFileWindow,omitempty"`
	// Delete migrated data if no symlink at source is pointing to it.
	DeleteOrphanData NullableBool `json:"deleteOrphanData,omitempty"`
	FilePathFilter *FilePathFilter `json:"filePathFilter,omitempty"`
	// Specifies policy to select a file to migrate based on its creation, last access or modification time. eg. A file can be selected to migrate if it has not been accessed/modified in the ColdFileWindow. enum: kOlderThan, kLastAccessed, kLastModified. Specifies policy for file selection in data migration jobs based on time. 'kOlderThan': Migrate the files that are older than cold file window. 'kLastAccessed': Migrate the files that are not accessed in cold file window. 'kLastModified': Migrate the files that have not been modified in cold file window.
	FileSelectionPolicy NullableString `json:"fileSelectionPolicy,omitempty"`
	// Gives the size criteria to be used for selecting the files to be migrated in bytes. The cold files that are equal and greater than this size are migrated.
	FileSizeBytes NullableInt64 `json:"fileSizeBytes,omitempty"`
	// Specifies policy to select a file to migrate based on its size. eg. A file can be selected to migrate if its size is greater than or smaller than the FileSizeBytes. enum: kGreaterThan, kSmallerThan. Specifies policy for file selection in data migration jobs based on file size. 'kGreaterThan': Migrate the files whose size are greater than specified file size. 'kSmallerThan': Migrate the files whose size are smaller than specified file size.
	FileSizePolicy NullableString `json:"fileSizePolicy,omitempty"`
	// Specifies if data is to be migrated without stub.
	MigrateWithoutStub NullableBool `json:"migrateWithoutStub,omitempty"`
	// Mount path where the target view must be mounted on all NFS clients for accessing the migrated data.
	NfsMountPath NullableString `json:"nfsMountPath,omitempty"`
	// The target view name to which the data will be migrated.
	TargetViewName NullableString `json:"targetViewName,omitempty"`
}

// NewDataMigrationJobParameters instantiates a new DataMigrationJobParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataMigrationJobParameters() *DataMigrationJobParameters {
	this := DataMigrationJobParameters{}
	return &this
}

// NewDataMigrationJobParametersWithDefaults instantiates a new DataMigrationJobParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataMigrationJobParametersWithDefaults() *DataMigrationJobParameters {
	this := DataMigrationJobParameters{}
	return &this
}

// GetColdFileWindow returns the ColdFileWindow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataMigrationJobParameters) GetColdFileWindow() int64 {
	if o == nil || o.ColdFileWindow.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ColdFileWindow.Get()
}

// GetColdFileWindowOk returns a tuple with the ColdFileWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataMigrationJobParameters) GetColdFileWindowOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ColdFileWindow.Get(), o.ColdFileWindow.IsSet()
}

// HasColdFileWindow returns a boolean if a field has been set.
func (o *DataMigrationJobParameters) HasColdFileWindow() bool {
	if o != nil && o.ColdFileWindow.IsSet() {
		return true
	}

	return false
}

// SetColdFileWindow gets a reference to the given NullableInt64 and assigns it to the ColdFileWindow field.
func (o *DataMigrationJobParameters) SetColdFileWindow(v int64) {
	o.ColdFileWindow.Set(&v)
}
// SetColdFileWindowNil sets the value for ColdFileWindow to be an explicit nil
func (o *DataMigrationJobParameters) SetColdFileWindowNil() {
	o.ColdFileWindow.Set(nil)
}

// UnsetColdFileWindow ensures that no value is present for ColdFileWindow, not even an explicit nil
func (o *DataMigrationJobParameters) UnsetColdFileWindow() {
	o.ColdFileWindow.Unset()
}

// GetDeleteOrphanData returns the DeleteOrphanData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataMigrationJobParameters) GetDeleteOrphanData() bool {
	if o == nil || o.DeleteOrphanData.Get() == nil {
		var ret bool
		return ret
	}
	return *o.DeleteOrphanData.Get()
}

// GetDeleteOrphanDataOk returns a tuple with the DeleteOrphanData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataMigrationJobParameters) GetDeleteOrphanDataOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeleteOrphanData.Get(), o.DeleteOrphanData.IsSet()
}

// HasDeleteOrphanData returns a boolean if a field has been set.
func (o *DataMigrationJobParameters) HasDeleteOrphanData() bool {
	if o != nil && o.DeleteOrphanData.IsSet() {
		return true
	}

	return false
}

// SetDeleteOrphanData gets a reference to the given NullableBool and assigns it to the DeleteOrphanData field.
func (o *DataMigrationJobParameters) SetDeleteOrphanData(v bool) {
	o.DeleteOrphanData.Set(&v)
}
// SetDeleteOrphanDataNil sets the value for DeleteOrphanData to be an explicit nil
func (o *DataMigrationJobParameters) SetDeleteOrphanDataNil() {
	o.DeleteOrphanData.Set(nil)
}

// UnsetDeleteOrphanData ensures that no value is present for DeleteOrphanData, not even an explicit nil
func (o *DataMigrationJobParameters) UnsetDeleteOrphanData() {
	o.DeleteOrphanData.Unset()
}

// GetFilePathFilter returns the FilePathFilter field value if set, zero value otherwise.
func (o *DataMigrationJobParameters) GetFilePathFilter() FilePathFilter {
	if o == nil || o.FilePathFilter == nil {
		var ret FilePathFilter
		return ret
	}
	return *o.FilePathFilter
}

// GetFilePathFilterOk returns a tuple with the FilePathFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataMigrationJobParameters) GetFilePathFilterOk() (*FilePathFilter, bool) {
	if o == nil || o.FilePathFilter == nil {
		return nil, false
	}
	return o.FilePathFilter, true
}

// HasFilePathFilter returns a boolean if a field has been set.
func (o *DataMigrationJobParameters) HasFilePathFilter() bool {
	if o != nil && o.FilePathFilter != nil {
		return true
	}

	return false
}

// SetFilePathFilter gets a reference to the given FilePathFilter and assigns it to the FilePathFilter field.
func (o *DataMigrationJobParameters) SetFilePathFilter(v FilePathFilter) {
	o.FilePathFilter = &v
}

// GetFileSelectionPolicy returns the FileSelectionPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataMigrationJobParameters) GetFileSelectionPolicy() string {
	if o == nil || o.FileSelectionPolicy.Get() == nil {
		var ret string
		return ret
	}
	return *o.FileSelectionPolicy.Get()
}

// GetFileSelectionPolicyOk returns a tuple with the FileSelectionPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataMigrationJobParameters) GetFileSelectionPolicyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FileSelectionPolicy.Get(), o.FileSelectionPolicy.IsSet()
}

// HasFileSelectionPolicy returns a boolean if a field has been set.
func (o *DataMigrationJobParameters) HasFileSelectionPolicy() bool {
	if o != nil && o.FileSelectionPolicy.IsSet() {
		return true
	}

	return false
}

// SetFileSelectionPolicy gets a reference to the given NullableString and assigns it to the FileSelectionPolicy field.
func (o *DataMigrationJobParameters) SetFileSelectionPolicy(v string) {
	o.FileSelectionPolicy.Set(&v)
}
// SetFileSelectionPolicyNil sets the value for FileSelectionPolicy to be an explicit nil
func (o *DataMigrationJobParameters) SetFileSelectionPolicyNil() {
	o.FileSelectionPolicy.Set(nil)
}

// UnsetFileSelectionPolicy ensures that no value is present for FileSelectionPolicy, not even an explicit nil
func (o *DataMigrationJobParameters) UnsetFileSelectionPolicy() {
	o.FileSelectionPolicy.Unset()
}

// GetFileSizeBytes returns the FileSizeBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataMigrationJobParameters) GetFileSizeBytes() int64 {
	if o == nil || o.FileSizeBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.FileSizeBytes.Get()
}

// GetFileSizeBytesOk returns a tuple with the FileSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataMigrationJobParameters) GetFileSizeBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FileSizeBytes.Get(), o.FileSizeBytes.IsSet()
}

// HasFileSizeBytes returns a boolean if a field has been set.
func (o *DataMigrationJobParameters) HasFileSizeBytes() bool {
	if o != nil && o.FileSizeBytes.IsSet() {
		return true
	}

	return false
}

// SetFileSizeBytes gets a reference to the given NullableInt64 and assigns it to the FileSizeBytes field.
func (o *DataMigrationJobParameters) SetFileSizeBytes(v int64) {
	o.FileSizeBytes.Set(&v)
}
// SetFileSizeBytesNil sets the value for FileSizeBytes to be an explicit nil
func (o *DataMigrationJobParameters) SetFileSizeBytesNil() {
	o.FileSizeBytes.Set(nil)
}

// UnsetFileSizeBytes ensures that no value is present for FileSizeBytes, not even an explicit nil
func (o *DataMigrationJobParameters) UnsetFileSizeBytes() {
	o.FileSizeBytes.Unset()
}

// GetFileSizePolicy returns the FileSizePolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataMigrationJobParameters) GetFileSizePolicy() string {
	if o == nil || o.FileSizePolicy.Get() == nil {
		var ret string
		return ret
	}
	return *o.FileSizePolicy.Get()
}

// GetFileSizePolicyOk returns a tuple with the FileSizePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataMigrationJobParameters) GetFileSizePolicyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FileSizePolicy.Get(), o.FileSizePolicy.IsSet()
}

// HasFileSizePolicy returns a boolean if a field has been set.
func (o *DataMigrationJobParameters) HasFileSizePolicy() bool {
	if o != nil && o.FileSizePolicy.IsSet() {
		return true
	}

	return false
}

// SetFileSizePolicy gets a reference to the given NullableString and assigns it to the FileSizePolicy field.
func (o *DataMigrationJobParameters) SetFileSizePolicy(v string) {
	o.FileSizePolicy.Set(&v)
}
// SetFileSizePolicyNil sets the value for FileSizePolicy to be an explicit nil
func (o *DataMigrationJobParameters) SetFileSizePolicyNil() {
	o.FileSizePolicy.Set(nil)
}

// UnsetFileSizePolicy ensures that no value is present for FileSizePolicy, not even an explicit nil
func (o *DataMigrationJobParameters) UnsetFileSizePolicy() {
	o.FileSizePolicy.Unset()
}

// GetMigrateWithoutStub returns the MigrateWithoutStub field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataMigrationJobParameters) GetMigrateWithoutStub() bool {
	if o == nil || o.MigrateWithoutStub.Get() == nil {
		var ret bool
		return ret
	}
	return *o.MigrateWithoutStub.Get()
}

// GetMigrateWithoutStubOk returns a tuple with the MigrateWithoutStub field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataMigrationJobParameters) GetMigrateWithoutStubOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MigrateWithoutStub.Get(), o.MigrateWithoutStub.IsSet()
}

// HasMigrateWithoutStub returns a boolean if a field has been set.
func (o *DataMigrationJobParameters) HasMigrateWithoutStub() bool {
	if o != nil && o.MigrateWithoutStub.IsSet() {
		return true
	}

	return false
}

// SetMigrateWithoutStub gets a reference to the given NullableBool and assigns it to the MigrateWithoutStub field.
func (o *DataMigrationJobParameters) SetMigrateWithoutStub(v bool) {
	o.MigrateWithoutStub.Set(&v)
}
// SetMigrateWithoutStubNil sets the value for MigrateWithoutStub to be an explicit nil
func (o *DataMigrationJobParameters) SetMigrateWithoutStubNil() {
	o.MigrateWithoutStub.Set(nil)
}

// UnsetMigrateWithoutStub ensures that no value is present for MigrateWithoutStub, not even an explicit nil
func (o *DataMigrationJobParameters) UnsetMigrateWithoutStub() {
	o.MigrateWithoutStub.Unset()
}

// GetNfsMountPath returns the NfsMountPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataMigrationJobParameters) GetNfsMountPath() string {
	if o == nil || o.NfsMountPath.Get() == nil {
		var ret string
		return ret
	}
	return *o.NfsMountPath.Get()
}

// GetNfsMountPathOk returns a tuple with the NfsMountPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataMigrationJobParameters) GetNfsMountPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NfsMountPath.Get(), o.NfsMountPath.IsSet()
}

// HasNfsMountPath returns a boolean if a field has been set.
func (o *DataMigrationJobParameters) HasNfsMountPath() bool {
	if o != nil && o.NfsMountPath.IsSet() {
		return true
	}

	return false
}

// SetNfsMountPath gets a reference to the given NullableString and assigns it to the NfsMountPath field.
func (o *DataMigrationJobParameters) SetNfsMountPath(v string) {
	o.NfsMountPath.Set(&v)
}
// SetNfsMountPathNil sets the value for NfsMountPath to be an explicit nil
func (o *DataMigrationJobParameters) SetNfsMountPathNil() {
	o.NfsMountPath.Set(nil)
}

// UnsetNfsMountPath ensures that no value is present for NfsMountPath, not even an explicit nil
func (o *DataMigrationJobParameters) UnsetNfsMountPath() {
	o.NfsMountPath.Unset()
}

// GetTargetViewName returns the TargetViewName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataMigrationJobParameters) GetTargetViewName() string {
	if o == nil || o.TargetViewName.Get() == nil {
		var ret string
		return ret
	}
	return *o.TargetViewName.Get()
}

// GetTargetViewNameOk returns a tuple with the TargetViewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataMigrationJobParameters) GetTargetViewNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TargetViewName.Get(), o.TargetViewName.IsSet()
}

// HasTargetViewName returns a boolean if a field has been set.
func (o *DataMigrationJobParameters) HasTargetViewName() bool {
	if o != nil && o.TargetViewName.IsSet() {
		return true
	}

	return false
}

// SetTargetViewName gets a reference to the given NullableString and assigns it to the TargetViewName field.
func (o *DataMigrationJobParameters) SetTargetViewName(v string) {
	o.TargetViewName.Set(&v)
}
// SetTargetViewNameNil sets the value for TargetViewName to be an explicit nil
func (o *DataMigrationJobParameters) SetTargetViewNameNil() {
	o.TargetViewName.Set(nil)
}

// UnsetTargetViewName ensures that no value is present for TargetViewName, not even an explicit nil
func (o *DataMigrationJobParameters) UnsetTargetViewName() {
	o.TargetViewName.Unset()
}

func (o DataMigrationJobParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ColdFileWindow.IsSet() {
		toSerialize["coldFileWindow"] = o.ColdFileWindow.Get()
	}
	if o.DeleteOrphanData.IsSet() {
		toSerialize["deleteOrphanData"] = o.DeleteOrphanData.Get()
	}
	if o.FilePathFilter != nil {
		toSerialize["filePathFilter"] = o.FilePathFilter
	}
	if o.FileSelectionPolicy.IsSet() {
		toSerialize["fileSelectionPolicy"] = o.FileSelectionPolicy.Get()
	}
	if o.FileSizeBytes.IsSet() {
		toSerialize["fileSizeBytes"] = o.FileSizeBytes.Get()
	}
	if o.FileSizePolicy.IsSet() {
		toSerialize["fileSizePolicy"] = o.FileSizePolicy.Get()
	}
	if o.MigrateWithoutStub.IsSet() {
		toSerialize["migrateWithoutStub"] = o.MigrateWithoutStub.Get()
	}
	if o.NfsMountPath.IsSet() {
		toSerialize["nfsMountPath"] = o.NfsMountPath.Get()
	}
	if o.TargetViewName.IsSet() {
		toSerialize["targetViewName"] = o.TargetViewName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDataMigrationJobParameters struct {
	value *DataMigrationJobParameters
	isSet bool
}

func (v NullableDataMigrationJobParameters) Get() *DataMigrationJobParameters {
	return v.value
}

func (v *NullableDataMigrationJobParameters) Set(val *DataMigrationJobParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableDataMigrationJobParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableDataMigrationJobParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataMigrationJobParameters(val *DataMigrationJobParameters) *NullableDataMigrationJobParameters {
	return &NullableDataMigrationJobParameters{value: val, isSet: true}
}

func (v NullableDataMigrationJobParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataMigrationJobParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


