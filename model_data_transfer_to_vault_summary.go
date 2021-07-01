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

// DataTransferToVaultSummary Specifies statistics about the transfer of data from this Cohesity Cluster to a Vault.
type DataTransferToVaultSummary struct {
	// Array of Data Transfer Statistics Per Protection Jobs.  Specifies the data transfer summary statistics for each Protection Job that is transferring data from this Cohesity Cluster to this Vault (External Target).
	DataTransferPerProtectionJob []DataTransferToVaultPerProtectionJob `json:"dataTransferPerProtectionJob,omitempty"`
	// Array of Logical Data Transferred Per Day.  Specifies the logical data transferred from this Cohesity Cluster to this Vault during the time period specified using the startTimeMsecs and endTimeMsecs parameters. For each day in the time period, an array element is returned, for example if 7 days are specified, 7 array elements are returned. The logical size is when the data is fully hydrated or expanded.
	LogicalDataTransferredBytesDuringTimeRange []int64 `json:"logicalDataTransferredBytesDuringTimeRange,omitempty"`
	// Specifies the total number of logical bytes that are transferred from this Cohesity Cluster to this Vault. The logical size is when the data is fully hydrated or expanded.
	NumLogicalBytesTransferred NullableInt64 `json:"numLogicalBytesTransferred,omitempty"`
	// Specifies the total number of physical bytes that are transferred from this Cohesity Cluster to this Vault.
	NumPhysicalBytesTransferred NullableInt64 `json:"numPhysicalBytesTransferred,omitempty"`
	// Specifies the number of Protection Jobs that transfer data to this Vault.
	NumProtectionJobs NullableInt64 `json:"numProtectionJobs,omitempty"`
	// Array of Physical Data Transferred Per Day.  Specifies the physical data transferred from this Cohesity Cluster to this Vault during the time period specified using the startTimeMsecs and endTimeMsecs parameters. For each day in the time period, an array element is returned, for example if 7 days are specified, 7 array elements are returned.
	PhysicalDataTransferredBytesDuringTimeRange []int64 `json:"physicalDataTransferredBytesDuringTimeRange,omitempty"`
	// Specifies the storage consumed on the Vault as of last day in the specified time range.
	StorageConsumedBytes NullableInt64 `json:"storageConsumedBytes,omitempty"`
	// The vault Id associated with the vault.
	VaultId NullableInt64 `json:"vaultId,omitempty"`
	// Specifies the name of the Vault (External Target).
	VaultName NullableString `json:"vaultName,omitempty"`
	// Specifies the type of Vault. 'kNearline' indicates a Google Nearline Vault. 'kGlacier' indicates an AWS Glacier Vault. 'kS3' indicates an AWS S3 Vault. 'kAzureStandard' indicates a Microsoft Azure Standard Vault. 'kS3Compatible' indicates an S3 Compatible Vault. (See the online help for supported types.) 'kQStarTape' indicates a QStar Tape Vault. 'kGoogleStandard' indicates a Google Standard Vault. 'kGoogleDRA' indicates a Google DRA Vault. 'kAmazonS3StandardIA' indicates an Amazon S3 Standard-IA Vault. 'kAWSGovCloud' indicates an AWS Gov Cloud Vault. 'kNAS' indicates a NAS Vault. 'kColdline' indicates a Google Coldline Vault. 'kAzureGovCloud' indicates a Microsoft Azure Gov Cloud Vault. 'kAzureArchive' indicates an Azure Archive Vault. 'kAzure' indicates an Azure Vault. 'kGoogle' indicates a Google Vault. 'kAmazon' indicates an Amazon Vault. 'kOracle' indicates an Oracle Vault. 'kOracleTierStandard' indicates an Oracle Tier Standard Vault. 'kOracleTierArchive' indicates an Oracle Tier Archive Vault. 'kAmazonC2S' indicates an Amazon Commercial Cloud Services Vault.
	VaultType NullableString `json:"vaultType,omitempty"`
}

// NewDataTransferToVaultSummary instantiates a new DataTransferToVaultSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataTransferToVaultSummary() *DataTransferToVaultSummary {
	this := DataTransferToVaultSummary{}
	return &this
}

// NewDataTransferToVaultSummaryWithDefaults instantiates a new DataTransferToVaultSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataTransferToVaultSummaryWithDefaults() *DataTransferToVaultSummary {
	this := DataTransferToVaultSummary{}
	return &this
}

// GetDataTransferPerProtectionJob returns the DataTransferPerProtectionJob field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataTransferToVaultSummary) GetDataTransferPerProtectionJob() []DataTransferToVaultPerProtectionJob {
	if o == nil  {
		var ret []DataTransferToVaultPerProtectionJob
		return ret
	}
	return o.DataTransferPerProtectionJob
}

// GetDataTransferPerProtectionJobOk returns a tuple with the DataTransferPerProtectionJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataTransferToVaultSummary) GetDataTransferPerProtectionJobOk() (*[]DataTransferToVaultPerProtectionJob, bool) {
	if o == nil || o.DataTransferPerProtectionJob == nil {
		return nil, false
	}
	return &o.DataTransferPerProtectionJob, true
}

// HasDataTransferPerProtectionJob returns a boolean if a field has been set.
func (o *DataTransferToVaultSummary) HasDataTransferPerProtectionJob() bool {
	if o != nil && o.DataTransferPerProtectionJob != nil {
		return true
	}

	return false
}

// SetDataTransferPerProtectionJob gets a reference to the given []DataTransferToVaultPerProtectionJob and assigns it to the DataTransferPerProtectionJob field.
func (o *DataTransferToVaultSummary) SetDataTransferPerProtectionJob(v []DataTransferToVaultPerProtectionJob) {
	o.DataTransferPerProtectionJob = v
}

// GetLogicalDataTransferredBytesDuringTimeRange returns the LogicalDataTransferredBytesDuringTimeRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataTransferToVaultSummary) GetLogicalDataTransferredBytesDuringTimeRange() []int64 {
	if o == nil  {
		var ret []int64
		return ret
	}
	return o.LogicalDataTransferredBytesDuringTimeRange
}

// GetLogicalDataTransferredBytesDuringTimeRangeOk returns a tuple with the LogicalDataTransferredBytesDuringTimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataTransferToVaultSummary) GetLogicalDataTransferredBytesDuringTimeRangeOk() (*[]int64, bool) {
	if o == nil || o.LogicalDataTransferredBytesDuringTimeRange == nil {
		return nil, false
	}
	return &o.LogicalDataTransferredBytesDuringTimeRange, true
}

// HasLogicalDataTransferredBytesDuringTimeRange returns a boolean if a field has been set.
func (o *DataTransferToVaultSummary) HasLogicalDataTransferredBytesDuringTimeRange() bool {
	if o != nil && o.LogicalDataTransferredBytesDuringTimeRange != nil {
		return true
	}

	return false
}

// SetLogicalDataTransferredBytesDuringTimeRange gets a reference to the given []int64 and assigns it to the LogicalDataTransferredBytesDuringTimeRange field.
func (o *DataTransferToVaultSummary) SetLogicalDataTransferredBytesDuringTimeRange(v []int64) {
	o.LogicalDataTransferredBytesDuringTimeRange = v
}

// GetNumLogicalBytesTransferred returns the NumLogicalBytesTransferred field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataTransferToVaultSummary) GetNumLogicalBytesTransferred() int64 {
	if o == nil || o.NumLogicalBytesTransferred.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NumLogicalBytesTransferred.Get()
}

// GetNumLogicalBytesTransferredOk returns a tuple with the NumLogicalBytesTransferred field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataTransferToVaultSummary) GetNumLogicalBytesTransferredOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NumLogicalBytesTransferred.Get(), o.NumLogicalBytesTransferred.IsSet()
}

// HasNumLogicalBytesTransferred returns a boolean if a field has been set.
func (o *DataTransferToVaultSummary) HasNumLogicalBytesTransferred() bool {
	if o != nil && o.NumLogicalBytesTransferred.IsSet() {
		return true
	}

	return false
}

// SetNumLogicalBytesTransferred gets a reference to the given NullableInt64 and assigns it to the NumLogicalBytesTransferred field.
func (o *DataTransferToVaultSummary) SetNumLogicalBytesTransferred(v int64) {
	o.NumLogicalBytesTransferred.Set(&v)
}
// SetNumLogicalBytesTransferredNil sets the value for NumLogicalBytesTransferred to be an explicit nil
func (o *DataTransferToVaultSummary) SetNumLogicalBytesTransferredNil() {
	o.NumLogicalBytesTransferred.Set(nil)
}

// UnsetNumLogicalBytesTransferred ensures that no value is present for NumLogicalBytesTransferred, not even an explicit nil
func (o *DataTransferToVaultSummary) UnsetNumLogicalBytesTransferred() {
	o.NumLogicalBytesTransferred.Unset()
}

// GetNumPhysicalBytesTransferred returns the NumPhysicalBytesTransferred field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataTransferToVaultSummary) GetNumPhysicalBytesTransferred() int64 {
	if o == nil || o.NumPhysicalBytesTransferred.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NumPhysicalBytesTransferred.Get()
}

// GetNumPhysicalBytesTransferredOk returns a tuple with the NumPhysicalBytesTransferred field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataTransferToVaultSummary) GetNumPhysicalBytesTransferredOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NumPhysicalBytesTransferred.Get(), o.NumPhysicalBytesTransferred.IsSet()
}

// HasNumPhysicalBytesTransferred returns a boolean if a field has been set.
func (o *DataTransferToVaultSummary) HasNumPhysicalBytesTransferred() bool {
	if o != nil && o.NumPhysicalBytesTransferred.IsSet() {
		return true
	}

	return false
}

// SetNumPhysicalBytesTransferred gets a reference to the given NullableInt64 and assigns it to the NumPhysicalBytesTransferred field.
func (o *DataTransferToVaultSummary) SetNumPhysicalBytesTransferred(v int64) {
	o.NumPhysicalBytesTransferred.Set(&v)
}
// SetNumPhysicalBytesTransferredNil sets the value for NumPhysicalBytesTransferred to be an explicit nil
func (o *DataTransferToVaultSummary) SetNumPhysicalBytesTransferredNil() {
	o.NumPhysicalBytesTransferred.Set(nil)
}

// UnsetNumPhysicalBytesTransferred ensures that no value is present for NumPhysicalBytesTransferred, not even an explicit nil
func (o *DataTransferToVaultSummary) UnsetNumPhysicalBytesTransferred() {
	o.NumPhysicalBytesTransferred.Unset()
}

// GetNumProtectionJobs returns the NumProtectionJobs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataTransferToVaultSummary) GetNumProtectionJobs() int64 {
	if o == nil || o.NumProtectionJobs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NumProtectionJobs.Get()
}

// GetNumProtectionJobsOk returns a tuple with the NumProtectionJobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataTransferToVaultSummary) GetNumProtectionJobsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NumProtectionJobs.Get(), o.NumProtectionJobs.IsSet()
}

// HasNumProtectionJobs returns a boolean if a field has been set.
func (o *DataTransferToVaultSummary) HasNumProtectionJobs() bool {
	if o != nil && o.NumProtectionJobs.IsSet() {
		return true
	}

	return false
}

// SetNumProtectionJobs gets a reference to the given NullableInt64 and assigns it to the NumProtectionJobs field.
func (o *DataTransferToVaultSummary) SetNumProtectionJobs(v int64) {
	o.NumProtectionJobs.Set(&v)
}
// SetNumProtectionJobsNil sets the value for NumProtectionJobs to be an explicit nil
func (o *DataTransferToVaultSummary) SetNumProtectionJobsNil() {
	o.NumProtectionJobs.Set(nil)
}

// UnsetNumProtectionJobs ensures that no value is present for NumProtectionJobs, not even an explicit nil
func (o *DataTransferToVaultSummary) UnsetNumProtectionJobs() {
	o.NumProtectionJobs.Unset()
}

// GetPhysicalDataTransferredBytesDuringTimeRange returns the PhysicalDataTransferredBytesDuringTimeRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataTransferToVaultSummary) GetPhysicalDataTransferredBytesDuringTimeRange() []int64 {
	if o == nil  {
		var ret []int64
		return ret
	}
	return o.PhysicalDataTransferredBytesDuringTimeRange
}

// GetPhysicalDataTransferredBytesDuringTimeRangeOk returns a tuple with the PhysicalDataTransferredBytesDuringTimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataTransferToVaultSummary) GetPhysicalDataTransferredBytesDuringTimeRangeOk() (*[]int64, bool) {
	if o == nil || o.PhysicalDataTransferredBytesDuringTimeRange == nil {
		return nil, false
	}
	return &o.PhysicalDataTransferredBytesDuringTimeRange, true
}

// HasPhysicalDataTransferredBytesDuringTimeRange returns a boolean if a field has been set.
func (o *DataTransferToVaultSummary) HasPhysicalDataTransferredBytesDuringTimeRange() bool {
	if o != nil && o.PhysicalDataTransferredBytesDuringTimeRange != nil {
		return true
	}

	return false
}

// SetPhysicalDataTransferredBytesDuringTimeRange gets a reference to the given []int64 and assigns it to the PhysicalDataTransferredBytesDuringTimeRange field.
func (o *DataTransferToVaultSummary) SetPhysicalDataTransferredBytesDuringTimeRange(v []int64) {
	o.PhysicalDataTransferredBytesDuringTimeRange = v
}

// GetStorageConsumedBytes returns the StorageConsumedBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataTransferToVaultSummary) GetStorageConsumedBytes() int64 {
	if o == nil || o.StorageConsumedBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.StorageConsumedBytes.Get()
}

// GetStorageConsumedBytesOk returns a tuple with the StorageConsumedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataTransferToVaultSummary) GetStorageConsumedBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageConsumedBytes.Get(), o.StorageConsumedBytes.IsSet()
}

// HasStorageConsumedBytes returns a boolean if a field has been set.
func (o *DataTransferToVaultSummary) HasStorageConsumedBytes() bool {
	if o != nil && o.StorageConsumedBytes.IsSet() {
		return true
	}

	return false
}

// SetStorageConsumedBytes gets a reference to the given NullableInt64 and assigns it to the StorageConsumedBytes field.
func (o *DataTransferToVaultSummary) SetStorageConsumedBytes(v int64) {
	o.StorageConsumedBytes.Set(&v)
}
// SetStorageConsumedBytesNil sets the value for StorageConsumedBytes to be an explicit nil
func (o *DataTransferToVaultSummary) SetStorageConsumedBytesNil() {
	o.StorageConsumedBytes.Set(nil)
}

// UnsetStorageConsumedBytes ensures that no value is present for StorageConsumedBytes, not even an explicit nil
func (o *DataTransferToVaultSummary) UnsetStorageConsumedBytes() {
	o.StorageConsumedBytes.Unset()
}

// GetVaultId returns the VaultId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataTransferToVaultSummary) GetVaultId() int64 {
	if o == nil || o.VaultId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.VaultId.Get()
}

// GetVaultIdOk returns a tuple with the VaultId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataTransferToVaultSummary) GetVaultIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VaultId.Get(), o.VaultId.IsSet()
}

// HasVaultId returns a boolean if a field has been set.
func (o *DataTransferToVaultSummary) HasVaultId() bool {
	if o != nil && o.VaultId.IsSet() {
		return true
	}

	return false
}

// SetVaultId gets a reference to the given NullableInt64 and assigns it to the VaultId field.
func (o *DataTransferToVaultSummary) SetVaultId(v int64) {
	o.VaultId.Set(&v)
}
// SetVaultIdNil sets the value for VaultId to be an explicit nil
func (o *DataTransferToVaultSummary) SetVaultIdNil() {
	o.VaultId.Set(nil)
}

// UnsetVaultId ensures that no value is present for VaultId, not even an explicit nil
func (o *DataTransferToVaultSummary) UnsetVaultId() {
	o.VaultId.Unset()
}

// GetVaultName returns the VaultName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataTransferToVaultSummary) GetVaultName() string {
	if o == nil || o.VaultName.Get() == nil {
		var ret string
		return ret
	}
	return *o.VaultName.Get()
}

// GetVaultNameOk returns a tuple with the VaultName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataTransferToVaultSummary) GetVaultNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VaultName.Get(), o.VaultName.IsSet()
}

// HasVaultName returns a boolean if a field has been set.
func (o *DataTransferToVaultSummary) HasVaultName() bool {
	if o != nil && o.VaultName.IsSet() {
		return true
	}

	return false
}

// SetVaultName gets a reference to the given NullableString and assigns it to the VaultName field.
func (o *DataTransferToVaultSummary) SetVaultName(v string) {
	o.VaultName.Set(&v)
}
// SetVaultNameNil sets the value for VaultName to be an explicit nil
func (o *DataTransferToVaultSummary) SetVaultNameNil() {
	o.VaultName.Set(nil)
}

// UnsetVaultName ensures that no value is present for VaultName, not even an explicit nil
func (o *DataTransferToVaultSummary) UnsetVaultName() {
	o.VaultName.Unset()
}

// GetVaultType returns the VaultType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataTransferToVaultSummary) GetVaultType() string {
	if o == nil || o.VaultType.Get() == nil {
		var ret string
		return ret
	}
	return *o.VaultType.Get()
}

// GetVaultTypeOk returns a tuple with the VaultType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataTransferToVaultSummary) GetVaultTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VaultType.Get(), o.VaultType.IsSet()
}

// HasVaultType returns a boolean if a field has been set.
func (o *DataTransferToVaultSummary) HasVaultType() bool {
	if o != nil && o.VaultType.IsSet() {
		return true
	}

	return false
}

// SetVaultType gets a reference to the given NullableString and assigns it to the VaultType field.
func (o *DataTransferToVaultSummary) SetVaultType(v string) {
	o.VaultType.Set(&v)
}
// SetVaultTypeNil sets the value for VaultType to be an explicit nil
func (o *DataTransferToVaultSummary) SetVaultTypeNil() {
	o.VaultType.Set(nil)
}

// UnsetVaultType ensures that no value is present for VaultType, not even an explicit nil
func (o *DataTransferToVaultSummary) UnsetVaultType() {
	o.VaultType.Unset()
}

func (o DataTransferToVaultSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DataTransferPerProtectionJob != nil {
		toSerialize["dataTransferPerProtectionJob"] = o.DataTransferPerProtectionJob
	}
	if o.LogicalDataTransferredBytesDuringTimeRange != nil {
		toSerialize["logicalDataTransferredBytesDuringTimeRange"] = o.LogicalDataTransferredBytesDuringTimeRange
	}
	if o.NumLogicalBytesTransferred.IsSet() {
		toSerialize["numLogicalBytesTransferred"] = o.NumLogicalBytesTransferred.Get()
	}
	if o.NumPhysicalBytesTransferred.IsSet() {
		toSerialize["numPhysicalBytesTransferred"] = o.NumPhysicalBytesTransferred.Get()
	}
	if o.NumProtectionJobs.IsSet() {
		toSerialize["numProtectionJobs"] = o.NumProtectionJobs.Get()
	}
	if o.PhysicalDataTransferredBytesDuringTimeRange != nil {
		toSerialize["physicalDataTransferredBytesDuringTimeRange"] = o.PhysicalDataTransferredBytesDuringTimeRange
	}
	if o.StorageConsumedBytes.IsSet() {
		toSerialize["storageConsumedBytes"] = o.StorageConsumedBytes.Get()
	}
	if o.VaultId.IsSet() {
		toSerialize["vaultId"] = o.VaultId.Get()
	}
	if o.VaultName.IsSet() {
		toSerialize["vaultName"] = o.VaultName.Get()
	}
	if o.VaultType.IsSet() {
		toSerialize["vaultType"] = o.VaultType.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDataTransferToVaultSummary struct {
	value *DataTransferToVaultSummary
	isSet bool
}

func (v NullableDataTransferToVaultSummary) Get() *DataTransferToVaultSummary {
	return v.value
}

func (v *NullableDataTransferToVaultSummary) Set(val *DataTransferToVaultSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableDataTransferToVaultSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableDataTransferToVaultSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataTransferToVaultSummary(val *DataTransferToVaultSummary) *NullableDataTransferToVaultSummary {
	return &NullableDataTransferToVaultSummary{value: val, isSet: true}
}

func (v NullableDataTransferToVaultSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataTransferToVaultSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


