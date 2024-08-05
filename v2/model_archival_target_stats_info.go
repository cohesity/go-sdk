/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// checks if the ArchivalTargetStatsInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArchivalTargetStatsInfo{}

// ArchivalTargetStatsInfo Specifies the stats of an archival run target.
type ArchivalTargetStatsInfo struct {
	// Specifies the archival task id. This is a protection group UID which only applies when archival type is 'Tape'.
	ArchivalTaskId NullableString `json:"archivalTaskId,omitempty"`
	// Specifies the ownership context for the target.
	OwnershipContext NullableString `json:"ownershipContext,omitempty"`
	// Specifies the archival target ID.
	TargetId NullableInt64 `json:"targetId,omitempty"`
	// Specifies the archival target name.
	TargetName NullableString `json:"targetName,omitempty"`
	// Specifies the archival target type.
	TargetType NullableString `json:"targetType,omitempty"`
	TierSettings *ArchivalTargetTierInfo `json:"tierSettings,omitempty"`
	// Specifies the usage type for the target.
	UsageType NullableString `json:"usageType,omitempty"`
	BackupGenericStats *BackupGenericStats `json:"backupGenericStats,omitempty"`
	NasStats *BackupNasStats `json:"nasStats,omitempty"`
	// Specifies stats for objects.
	Objects []ObjectStatsInfo `json:"objects,omitempty"`
}

// NewArchivalTargetStatsInfo instantiates a new ArchivalTargetStatsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArchivalTargetStatsInfo() *ArchivalTargetStatsInfo {
	this := ArchivalTargetStatsInfo{}
	return &this
}

// NewArchivalTargetStatsInfoWithDefaults instantiates a new ArchivalTargetStatsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArchivalTargetStatsInfoWithDefaults() *ArchivalTargetStatsInfo {
	this := ArchivalTargetStatsInfo{}
	return &this
}

// GetArchivalTaskId returns the ArchivalTaskId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArchivalTargetStatsInfo) GetArchivalTaskId() string {
	if o == nil || IsNil(o.ArchivalTaskId.Get()) {
		var ret string
		return ret
	}
	return *o.ArchivalTaskId.Get()
}

// GetArchivalTaskIdOk returns a tuple with the ArchivalTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchivalTargetStatsInfo) GetArchivalTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArchivalTaskId.Get(), o.ArchivalTaskId.IsSet()
}

// HasArchivalTaskId returns a boolean if a field has been set.
func (o *ArchivalTargetStatsInfo) HasArchivalTaskId() bool {
	if o != nil && o.ArchivalTaskId.IsSet() {
		return true
	}

	return false
}

// SetArchivalTaskId gets a reference to the given NullableString and assigns it to the ArchivalTaskId field.
func (o *ArchivalTargetStatsInfo) SetArchivalTaskId(v string) {
	o.ArchivalTaskId.Set(&v)
}
// SetArchivalTaskIdNil sets the value for ArchivalTaskId to be an explicit nil
func (o *ArchivalTargetStatsInfo) SetArchivalTaskIdNil() {
	o.ArchivalTaskId.Set(nil)
}

// UnsetArchivalTaskId ensures that no value is present for ArchivalTaskId, not even an explicit nil
func (o *ArchivalTargetStatsInfo) UnsetArchivalTaskId() {
	o.ArchivalTaskId.Unset()
}

// GetOwnershipContext returns the OwnershipContext field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArchivalTargetStatsInfo) GetOwnershipContext() string {
	if o == nil || IsNil(o.OwnershipContext.Get()) {
		var ret string
		return ret
	}
	return *o.OwnershipContext.Get()
}

// GetOwnershipContextOk returns a tuple with the OwnershipContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchivalTargetStatsInfo) GetOwnershipContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnershipContext.Get(), o.OwnershipContext.IsSet()
}

// HasOwnershipContext returns a boolean if a field has been set.
func (o *ArchivalTargetStatsInfo) HasOwnershipContext() bool {
	if o != nil && o.OwnershipContext.IsSet() {
		return true
	}

	return false
}

// SetOwnershipContext gets a reference to the given NullableString and assigns it to the OwnershipContext field.
func (o *ArchivalTargetStatsInfo) SetOwnershipContext(v string) {
	o.OwnershipContext.Set(&v)
}
// SetOwnershipContextNil sets the value for OwnershipContext to be an explicit nil
func (o *ArchivalTargetStatsInfo) SetOwnershipContextNil() {
	o.OwnershipContext.Set(nil)
}

// UnsetOwnershipContext ensures that no value is present for OwnershipContext, not even an explicit nil
func (o *ArchivalTargetStatsInfo) UnsetOwnershipContext() {
	o.OwnershipContext.Unset()
}

// GetTargetId returns the TargetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArchivalTargetStatsInfo) GetTargetId() int64 {
	if o == nil || IsNil(o.TargetId.Get()) {
		var ret int64
		return ret
	}
	return *o.TargetId.Get()
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchivalTargetStatsInfo) GetTargetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetId.Get(), o.TargetId.IsSet()
}

// HasTargetId returns a boolean if a field has been set.
func (o *ArchivalTargetStatsInfo) HasTargetId() bool {
	if o != nil && o.TargetId.IsSet() {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given NullableInt64 and assigns it to the TargetId field.
func (o *ArchivalTargetStatsInfo) SetTargetId(v int64) {
	o.TargetId.Set(&v)
}
// SetTargetIdNil sets the value for TargetId to be an explicit nil
func (o *ArchivalTargetStatsInfo) SetTargetIdNil() {
	o.TargetId.Set(nil)
}

// UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
func (o *ArchivalTargetStatsInfo) UnsetTargetId() {
	o.TargetId.Unset()
}

// GetTargetName returns the TargetName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArchivalTargetStatsInfo) GetTargetName() string {
	if o == nil || IsNil(o.TargetName.Get()) {
		var ret string
		return ret
	}
	return *o.TargetName.Get()
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchivalTargetStatsInfo) GetTargetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetName.Get(), o.TargetName.IsSet()
}

// HasTargetName returns a boolean if a field has been set.
func (o *ArchivalTargetStatsInfo) HasTargetName() bool {
	if o != nil && o.TargetName.IsSet() {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given NullableString and assigns it to the TargetName field.
func (o *ArchivalTargetStatsInfo) SetTargetName(v string) {
	o.TargetName.Set(&v)
}
// SetTargetNameNil sets the value for TargetName to be an explicit nil
func (o *ArchivalTargetStatsInfo) SetTargetNameNil() {
	o.TargetName.Set(nil)
}

// UnsetTargetName ensures that no value is present for TargetName, not even an explicit nil
func (o *ArchivalTargetStatsInfo) UnsetTargetName() {
	o.TargetName.Unset()
}

// GetTargetType returns the TargetType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArchivalTargetStatsInfo) GetTargetType() string {
	if o == nil || IsNil(o.TargetType.Get()) {
		var ret string
		return ret
	}
	return *o.TargetType.Get()
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchivalTargetStatsInfo) GetTargetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetType.Get(), o.TargetType.IsSet()
}

// HasTargetType returns a boolean if a field has been set.
func (o *ArchivalTargetStatsInfo) HasTargetType() bool {
	if o != nil && o.TargetType.IsSet() {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given NullableString and assigns it to the TargetType field.
func (o *ArchivalTargetStatsInfo) SetTargetType(v string) {
	o.TargetType.Set(&v)
}
// SetTargetTypeNil sets the value for TargetType to be an explicit nil
func (o *ArchivalTargetStatsInfo) SetTargetTypeNil() {
	o.TargetType.Set(nil)
}

// UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
func (o *ArchivalTargetStatsInfo) UnsetTargetType() {
	o.TargetType.Unset()
}

// GetTierSettings returns the TierSettings field value if set, zero value otherwise.
func (o *ArchivalTargetStatsInfo) GetTierSettings() ArchivalTargetTierInfo {
	if o == nil || IsNil(o.TierSettings) {
		var ret ArchivalTargetTierInfo
		return ret
	}
	return *o.TierSettings
}

// GetTierSettingsOk returns a tuple with the TierSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchivalTargetStatsInfo) GetTierSettingsOk() (*ArchivalTargetTierInfo, bool) {
	if o == nil || IsNil(o.TierSettings) {
		return nil, false
	}
	return o.TierSettings, true
}

// HasTierSettings returns a boolean if a field has been set.
func (o *ArchivalTargetStatsInfo) HasTierSettings() bool {
	if o != nil && !IsNil(o.TierSettings) {
		return true
	}

	return false
}

// SetTierSettings gets a reference to the given ArchivalTargetTierInfo and assigns it to the TierSettings field.
func (o *ArchivalTargetStatsInfo) SetTierSettings(v ArchivalTargetTierInfo) {
	o.TierSettings = &v
}

// GetUsageType returns the UsageType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArchivalTargetStatsInfo) GetUsageType() string {
	if o == nil || IsNil(o.UsageType.Get()) {
		var ret string
		return ret
	}
	return *o.UsageType.Get()
}

// GetUsageTypeOk returns a tuple with the UsageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchivalTargetStatsInfo) GetUsageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsageType.Get(), o.UsageType.IsSet()
}

// HasUsageType returns a boolean if a field has been set.
func (o *ArchivalTargetStatsInfo) HasUsageType() bool {
	if o != nil && o.UsageType.IsSet() {
		return true
	}

	return false
}

// SetUsageType gets a reference to the given NullableString and assigns it to the UsageType field.
func (o *ArchivalTargetStatsInfo) SetUsageType(v string) {
	o.UsageType.Set(&v)
}
// SetUsageTypeNil sets the value for UsageType to be an explicit nil
func (o *ArchivalTargetStatsInfo) SetUsageTypeNil() {
	o.UsageType.Set(nil)
}

// UnsetUsageType ensures that no value is present for UsageType, not even an explicit nil
func (o *ArchivalTargetStatsInfo) UnsetUsageType() {
	o.UsageType.Unset()
}

// GetBackupGenericStats returns the BackupGenericStats field value if set, zero value otherwise.
func (o *ArchivalTargetStatsInfo) GetBackupGenericStats() BackupGenericStats {
	if o == nil || IsNil(o.BackupGenericStats) {
		var ret BackupGenericStats
		return ret
	}
	return *o.BackupGenericStats
}

// GetBackupGenericStatsOk returns a tuple with the BackupGenericStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchivalTargetStatsInfo) GetBackupGenericStatsOk() (*BackupGenericStats, bool) {
	if o == nil || IsNil(o.BackupGenericStats) {
		return nil, false
	}
	return o.BackupGenericStats, true
}

// HasBackupGenericStats returns a boolean if a field has been set.
func (o *ArchivalTargetStatsInfo) HasBackupGenericStats() bool {
	if o != nil && !IsNil(o.BackupGenericStats) {
		return true
	}

	return false
}

// SetBackupGenericStats gets a reference to the given BackupGenericStats and assigns it to the BackupGenericStats field.
func (o *ArchivalTargetStatsInfo) SetBackupGenericStats(v BackupGenericStats) {
	o.BackupGenericStats = &v
}

// GetNasStats returns the NasStats field value if set, zero value otherwise.
func (o *ArchivalTargetStatsInfo) GetNasStats() BackupNasStats {
	if o == nil || IsNil(o.NasStats) {
		var ret BackupNasStats
		return ret
	}
	return *o.NasStats
}

// GetNasStatsOk returns a tuple with the NasStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchivalTargetStatsInfo) GetNasStatsOk() (*BackupNasStats, bool) {
	if o == nil || IsNil(o.NasStats) {
		return nil, false
	}
	return o.NasStats, true
}

// HasNasStats returns a boolean if a field has been set.
func (o *ArchivalTargetStatsInfo) HasNasStats() bool {
	if o != nil && !IsNil(o.NasStats) {
		return true
	}

	return false
}

// SetNasStats gets a reference to the given BackupNasStats and assigns it to the NasStats field.
func (o *ArchivalTargetStatsInfo) SetNasStats(v BackupNasStats) {
	o.NasStats = &v
}

// GetObjects returns the Objects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArchivalTargetStatsInfo) GetObjects() []ObjectStatsInfo {
	if o == nil {
		var ret []ObjectStatsInfo
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArchivalTargetStatsInfo) GetObjectsOk() ([]ObjectStatsInfo, bool) {
	if o == nil || IsNil(o.Objects) {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *ArchivalTargetStatsInfo) HasObjects() bool {
	if o != nil && !IsNil(o.Objects) {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []ObjectStatsInfo and assigns it to the Objects field.
func (o *ArchivalTargetStatsInfo) SetObjects(v []ObjectStatsInfo) {
	o.Objects = v
}

func (o ArchivalTargetStatsInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArchivalTargetStatsInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ArchivalTaskId.IsSet() {
		toSerialize["archivalTaskId"] = o.ArchivalTaskId.Get()
	}
	if o.OwnershipContext.IsSet() {
		toSerialize["ownershipContext"] = o.OwnershipContext.Get()
	}
	if o.TargetId.IsSet() {
		toSerialize["targetId"] = o.TargetId.Get()
	}
	if o.TargetName.IsSet() {
		toSerialize["targetName"] = o.TargetName.Get()
	}
	if o.TargetType.IsSet() {
		toSerialize["targetType"] = o.TargetType.Get()
	}
	if !IsNil(o.TierSettings) {
		toSerialize["tierSettings"] = o.TierSettings
	}
	if o.UsageType.IsSet() {
		toSerialize["usageType"] = o.UsageType.Get()
	}
	if !IsNil(o.BackupGenericStats) {
		toSerialize["backupGenericStats"] = o.BackupGenericStats
	}
	if !IsNil(o.NasStats) {
		toSerialize["nasStats"] = o.NasStats
	}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	return toSerialize, nil
}

type NullableArchivalTargetStatsInfo struct {
	value *ArchivalTargetStatsInfo
	isSet bool
}

func (v NullableArchivalTargetStatsInfo) Get() *ArchivalTargetStatsInfo {
	return v.value
}

func (v *NullableArchivalTargetStatsInfo) Set(val *ArchivalTargetStatsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableArchivalTargetStatsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableArchivalTargetStatsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArchivalTargetStatsInfo(val *ArchivalTargetStatsInfo) *NullableArchivalTargetStatsInfo {
	return &NullableArchivalTargetStatsInfo{value: val, isSet: true}
}

func (v NullableArchivalTargetStatsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArchivalTargetStatsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


