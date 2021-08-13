/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package helios
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/helios/utils"
)

var _ = NullableBool{}

// FlashbladeParams Specifies the information related to Registered Pure Flashblade.
type FlashbladeParams struct {
	RegistrationParams *FlashBladeRegistrationInfo `json:"registrationParams,omitempty"`
	// Specifies uuid of the pure flashblade server.
	Uuid NullableString `json:"uuid,omitempty"`
	// Specifies list of data vips that are assigned to cohesity cluster to create nfs share mountpoints.
	AssignedDataVips []string `json:"assignedDataVips,omitempty"`
	// Specifies the capacity in bytes assigned on pure flashblade for remote storage usage on cohesity cluster.
	AssignedCapacityBytes NullableInt64 `json:"assignedCapacityBytes,omitempty"`
	// If true, cohesity cluster uses all available capacity on pure flashblade for remote storage.
	IsDedicatedStorage NullableBool `json:"isDedicatedStorage,omitempty"`
	// Available data vips configured on pure flashblade.
	AvailableDataVips *[]string `json:"availableDataVips,omitempty"`
	// Available capacity on pure flashblade.
	AvailableCapacity NullableInt64 `json:"availableCapacity,omitempty"`
	// Number of new file systems created on pure flashblade when assignedCapacityBytes is updated.
	CreatedFileSystemCount NullableInt64 `json:"createdFileSystemCount,omitempty"`
	// Number of file systems that are updated on pure flashblade when assignedCapacityBytes is updated.
	UpdatedFileSystemCount NullableInt64 `json:"updatedFileSystemCount,omitempty"`
	// Software OS and version running on pure flashblade
	SoftwareOsVersion NullableString `json:"softwareOsVersion,omitempty"`
	// Name of the pure flashblade specified on pure storage.
	Name NullableString `json:"name,omitempty"`
	// Total capacity of pure flashblade.
	TotalCapacity NullableInt64 `json:"totalCapacity,omitempty"`
}

// NewFlashbladeParams instantiates a new FlashbladeParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlashbladeParams() *FlashbladeParams {
	this := FlashbladeParams{}
	return &this
}

// NewFlashbladeParamsWithDefaults instantiates a new FlashbladeParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlashbladeParamsWithDefaults() *FlashbladeParams {
	this := FlashbladeParams{}
	return &this
}

// GetRegistrationParams returns the RegistrationParams field value if set, zero value otherwise.
func (o *FlashbladeParams) GetRegistrationParams() FlashBladeRegistrationInfo {
	if o == nil || o.RegistrationParams == nil {
		var ret FlashBladeRegistrationInfo
		return ret
	}
	return *o.RegistrationParams
}

// GetRegistrationParamsOk returns a tuple with the RegistrationParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlashbladeParams) GetRegistrationParamsOk() (*FlashBladeRegistrationInfo, bool) {
	if o == nil || o.RegistrationParams == nil {
		return nil, false
	}
	return o.RegistrationParams, true
}

// HasRegistrationParams returns a boolean if a field has been set.
func (o *FlashbladeParams) HasRegistrationParams() bool {
	if o != nil && o.RegistrationParams != nil {
		return true
	}

	return false
}

// SetRegistrationParams gets a reference to the given FlashBladeRegistrationInfo and assigns it to the RegistrationParams field.
func (o *FlashbladeParams) SetRegistrationParams(v FlashBladeRegistrationInfo) {
	o.RegistrationParams = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlashbladeParams) GetUuid() string {
	if o == nil || o.Uuid.Get() == nil {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlashbladeParams) GetUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *FlashbladeParams) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *FlashbladeParams) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *FlashbladeParams) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *FlashbladeParams) UnsetUuid() {
	o.Uuid.Unset()
}

// GetAssignedDataVips returns the AssignedDataVips field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlashbladeParams) GetAssignedDataVips() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.AssignedDataVips
}

// GetAssignedDataVipsOk returns a tuple with the AssignedDataVips field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlashbladeParams) GetAssignedDataVipsOk() (*[]string, bool) {
	if o == nil || o.AssignedDataVips == nil {
		return nil, false
	}
	return &o.AssignedDataVips, true
}

// HasAssignedDataVips returns a boolean if a field has been set.
func (o *FlashbladeParams) HasAssignedDataVips() bool {
	if o != nil && o.AssignedDataVips != nil {
		return true
	}

	return false
}

// SetAssignedDataVips gets a reference to the given []string and assigns it to the AssignedDataVips field.
func (o *FlashbladeParams) SetAssignedDataVips(v []string) {
	o.AssignedDataVips = v
}

// GetAssignedCapacityBytes returns the AssignedCapacityBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlashbladeParams) GetAssignedCapacityBytes() int64 {
	if o == nil || o.AssignedCapacityBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AssignedCapacityBytes.Get()
}

// GetAssignedCapacityBytesOk returns a tuple with the AssignedCapacityBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlashbladeParams) GetAssignedCapacityBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AssignedCapacityBytes.Get(), o.AssignedCapacityBytes.IsSet()
}

// HasAssignedCapacityBytes returns a boolean if a field has been set.
func (o *FlashbladeParams) HasAssignedCapacityBytes() bool {
	if o != nil && o.AssignedCapacityBytes.IsSet() {
		return true
	}

	return false
}

// SetAssignedCapacityBytes gets a reference to the given NullableInt64 and assigns it to the AssignedCapacityBytes field.
func (o *FlashbladeParams) SetAssignedCapacityBytes(v int64) {
	o.AssignedCapacityBytes.Set(&v)
}
// SetAssignedCapacityBytesNil sets the value for AssignedCapacityBytes to be an explicit nil
func (o *FlashbladeParams) SetAssignedCapacityBytesNil() {
	o.AssignedCapacityBytes.Set(nil)
}

// UnsetAssignedCapacityBytes ensures that no value is present for AssignedCapacityBytes, not even an explicit nil
func (o *FlashbladeParams) UnsetAssignedCapacityBytes() {
	o.AssignedCapacityBytes.Unset()
}

// GetIsDedicatedStorage returns the IsDedicatedStorage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlashbladeParams) GetIsDedicatedStorage() bool {
	if o == nil || o.IsDedicatedStorage.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsDedicatedStorage.Get()
}

// GetIsDedicatedStorageOk returns a tuple with the IsDedicatedStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlashbladeParams) GetIsDedicatedStorageOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsDedicatedStorage.Get(), o.IsDedicatedStorage.IsSet()
}

// HasIsDedicatedStorage returns a boolean if a field has been set.
func (o *FlashbladeParams) HasIsDedicatedStorage() bool {
	if o != nil && o.IsDedicatedStorage.IsSet() {
		return true
	}

	return false
}

// SetIsDedicatedStorage gets a reference to the given NullableBool and assigns it to the IsDedicatedStorage field.
func (o *FlashbladeParams) SetIsDedicatedStorage(v bool) {
	o.IsDedicatedStorage.Set(&v)
}
// SetIsDedicatedStorageNil sets the value for IsDedicatedStorage to be an explicit nil
func (o *FlashbladeParams) SetIsDedicatedStorageNil() {
	o.IsDedicatedStorage.Set(nil)
}

// UnsetIsDedicatedStorage ensures that no value is present for IsDedicatedStorage, not even an explicit nil
func (o *FlashbladeParams) UnsetIsDedicatedStorage() {
	o.IsDedicatedStorage.Unset()
}

// GetAvailableDataVips returns the AvailableDataVips field value if set, zero value otherwise.
func (o *FlashbladeParams) GetAvailableDataVips() []string {
	if o == nil || o.AvailableDataVips == nil {
		var ret []string
		return ret
	}
	return *o.AvailableDataVips
}

// GetAvailableDataVipsOk returns a tuple with the AvailableDataVips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlashbladeParams) GetAvailableDataVipsOk() (*[]string, bool) {
	if o == nil || o.AvailableDataVips == nil {
		return nil, false
	}
	return o.AvailableDataVips, true
}

// HasAvailableDataVips returns a boolean if a field has been set.
func (o *FlashbladeParams) HasAvailableDataVips() bool {
	if o != nil && o.AvailableDataVips != nil {
		return true
	}

	return false
}

// SetAvailableDataVips gets a reference to the given []string and assigns it to the AvailableDataVips field.
func (o *FlashbladeParams) SetAvailableDataVips(v []string) {
	o.AvailableDataVips = &v
}

// GetAvailableCapacity returns the AvailableCapacity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlashbladeParams) GetAvailableCapacity() int64 {
	if o == nil || o.AvailableCapacity.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AvailableCapacity.Get()
}

// GetAvailableCapacityOk returns a tuple with the AvailableCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlashbladeParams) GetAvailableCapacityOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AvailableCapacity.Get(), o.AvailableCapacity.IsSet()
}

// HasAvailableCapacity returns a boolean if a field has been set.
func (o *FlashbladeParams) HasAvailableCapacity() bool {
	if o != nil && o.AvailableCapacity.IsSet() {
		return true
	}

	return false
}

// SetAvailableCapacity gets a reference to the given NullableInt64 and assigns it to the AvailableCapacity field.
func (o *FlashbladeParams) SetAvailableCapacity(v int64) {
	o.AvailableCapacity.Set(&v)
}
// SetAvailableCapacityNil sets the value for AvailableCapacity to be an explicit nil
func (o *FlashbladeParams) SetAvailableCapacityNil() {
	o.AvailableCapacity.Set(nil)
}

// UnsetAvailableCapacity ensures that no value is present for AvailableCapacity, not even an explicit nil
func (o *FlashbladeParams) UnsetAvailableCapacity() {
	o.AvailableCapacity.Unset()
}

// GetCreatedFileSystemCount returns the CreatedFileSystemCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlashbladeParams) GetCreatedFileSystemCount() int64 {
	if o == nil || o.CreatedFileSystemCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CreatedFileSystemCount.Get()
}

// GetCreatedFileSystemCountOk returns a tuple with the CreatedFileSystemCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlashbladeParams) GetCreatedFileSystemCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedFileSystemCount.Get(), o.CreatedFileSystemCount.IsSet()
}

// HasCreatedFileSystemCount returns a boolean if a field has been set.
func (o *FlashbladeParams) HasCreatedFileSystemCount() bool {
	if o != nil && o.CreatedFileSystemCount.IsSet() {
		return true
	}

	return false
}

// SetCreatedFileSystemCount gets a reference to the given NullableInt64 and assigns it to the CreatedFileSystemCount field.
func (o *FlashbladeParams) SetCreatedFileSystemCount(v int64) {
	o.CreatedFileSystemCount.Set(&v)
}
// SetCreatedFileSystemCountNil sets the value for CreatedFileSystemCount to be an explicit nil
func (o *FlashbladeParams) SetCreatedFileSystemCountNil() {
	o.CreatedFileSystemCount.Set(nil)
}

// UnsetCreatedFileSystemCount ensures that no value is present for CreatedFileSystemCount, not even an explicit nil
func (o *FlashbladeParams) UnsetCreatedFileSystemCount() {
	o.CreatedFileSystemCount.Unset()
}

// GetUpdatedFileSystemCount returns the UpdatedFileSystemCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlashbladeParams) GetUpdatedFileSystemCount() int64 {
	if o == nil || o.UpdatedFileSystemCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedFileSystemCount.Get()
}

// GetUpdatedFileSystemCountOk returns a tuple with the UpdatedFileSystemCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlashbladeParams) GetUpdatedFileSystemCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedFileSystemCount.Get(), o.UpdatedFileSystemCount.IsSet()
}

// HasUpdatedFileSystemCount returns a boolean if a field has been set.
func (o *FlashbladeParams) HasUpdatedFileSystemCount() bool {
	if o != nil && o.UpdatedFileSystemCount.IsSet() {
		return true
	}

	return false
}

// SetUpdatedFileSystemCount gets a reference to the given NullableInt64 and assigns it to the UpdatedFileSystemCount field.
func (o *FlashbladeParams) SetUpdatedFileSystemCount(v int64) {
	o.UpdatedFileSystemCount.Set(&v)
}
// SetUpdatedFileSystemCountNil sets the value for UpdatedFileSystemCount to be an explicit nil
func (o *FlashbladeParams) SetUpdatedFileSystemCountNil() {
	o.UpdatedFileSystemCount.Set(nil)
}

// UnsetUpdatedFileSystemCount ensures that no value is present for UpdatedFileSystemCount, not even an explicit nil
func (o *FlashbladeParams) UnsetUpdatedFileSystemCount() {
	o.UpdatedFileSystemCount.Unset()
}

// GetSoftwareOsVersion returns the SoftwareOsVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlashbladeParams) GetSoftwareOsVersion() string {
	if o == nil || o.SoftwareOsVersion.Get() == nil {
		var ret string
		return ret
	}
	return *o.SoftwareOsVersion.Get()
}

// GetSoftwareOsVersionOk returns a tuple with the SoftwareOsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlashbladeParams) GetSoftwareOsVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SoftwareOsVersion.Get(), o.SoftwareOsVersion.IsSet()
}

// HasSoftwareOsVersion returns a boolean if a field has been set.
func (o *FlashbladeParams) HasSoftwareOsVersion() bool {
	if o != nil && o.SoftwareOsVersion.IsSet() {
		return true
	}

	return false
}

// SetSoftwareOsVersion gets a reference to the given NullableString and assigns it to the SoftwareOsVersion field.
func (o *FlashbladeParams) SetSoftwareOsVersion(v string) {
	o.SoftwareOsVersion.Set(&v)
}
// SetSoftwareOsVersionNil sets the value for SoftwareOsVersion to be an explicit nil
func (o *FlashbladeParams) SetSoftwareOsVersionNil() {
	o.SoftwareOsVersion.Set(nil)
}

// UnsetSoftwareOsVersion ensures that no value is present for SoftwareOsVersion, not even an explicit nil
func (o *FlashbladeParams) UnsetSoftwareOsVersion() {
	o.SoftwareOsVersion.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlashbladeParams) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlashbladeParams) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *FlashbladeParams) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *FlashbladeParams) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *FlashbladeParams) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *FlashbladeParams) UnsetName() {
	o.Name.Unset()
}

// GetTotalCapacity returns the TotalCapacity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlashbladeParams) GetTotalCapacity() int64 {
	if o == nil || o.TotalCapacity.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TotalCapacity.Get()
}

// GetTotalCapacityOk returns a tuple with the TotalCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlashbladeParams) GetTotalCapacityOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TotalCapacity.Get(), o.TotalCapacity.IsSet()
}

// HasTotalCapacity returns a boolean if a field has been set.
func (o *FlashbladeParams) HasTotalCapacity() bool {
	if o != nil && o.TotalCapacity.IsSet() {
		return true
	}

	return false
}

// SetTotalCapacity gets a reference to the given NullableInt64 and assigns it to the TotalCapacity field.
func (o *FlashbladeParams) SetTotalCapacity(v int64) {
	o.TotalCapacity.Set(&v)
}
// SetTotalCapacityNil sets the value for TotalCapacity to be an explicit nil
func (o *FlashbladeParams) SetTotalCapacityNil() {
	o.TotalCapacity.Set(nil)
}

// UnsetTotalCapacity ensures that no value is present for TotalCapacity, not even an explicit nil
func (o *FlashbladeParams) UnsetTotalCapacity() {
	o.TotalCapacity.Unset()
}

func (o FlashbladeParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RegistrationParams != nil {
		toSerialize["registrationParams"] = o.RegistrationParams
	}
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	if o.AssignedDataVips != nil {
		toSerialize["assignedDataVips"] = o.AssignedDataVips
	}
	if o.AssignedCapacityBytes.IsSet() {
		toSerialize["assignedCapacityBytes"] = o.AssignedCapacityBytes.Get()
	}
	if o.IsDedicatedStorage.IsSet() {
		toSerialize["isDedicatedStorage"] = o.IsDedicatedStorage.Get()
	}
	if o.AvailableDataVips != nil {
		toSerialize["availableDataVips"] = o.AvailableDataVips
	}
	if o.AvailableCapacity.IsSet() {
		toSerialize["availableCapacity"] = o.AvailableCapacity.Get()
	}
	if o.CreatedFileSystemCount.IsSet() {
		toSerialize["createdFileSystemCount"] = o.CreatedFileSystemCount.Get()
	}
	if o.UpdatedFileSystemCount.IsSet() {
		toSerialize["updatedFileSystemCount"] = o.UpdatedFileSystemCount.Get()
	}
	if o.SoftwareOsVersion.IsSet() {
		toSerialize["softwareOsVersion"] = o.SoftwareOsVersion.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.TotalCapacity.IsSet() {
		toSerialize["totalCapacity"] = o.TotalCapacity.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableFlashbladeParams struct {
	value *FlashbladeParams
	isSet bool
}

func (v NullableFlashbladeParams) Get() *FlashbladeParams {
	return v.value
}

func (v *NullableFlashbladeParams) Set(val *FlashbladeParams) {
	v.value = val
	v.isSet = true
}

func (v NullableFlashbladeParams) IsSet() bool {
	return v.isSet
}

func (v *NullableFlashbladeParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlashbladeParams(val *FlashbladeParams) *NullableFlashbladeParams {
	return &NullableFlashbladeParams{value: val, isSet: true}
}

func (v NullableFlashbladeParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlashbladeParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


