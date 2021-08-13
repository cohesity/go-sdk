/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package onprem
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/onprem/utils"
	"fmt"
)

var _ = NullableBool{}

// VmwareObjectProtectionRequestParams Specifies the parameters which are specific to VMware object protection.
type VmwareObjectProtectionRequestParams struct {
	// Specifies the objects to include in the backup.
	Objects []VmwareObjectProtectionRequest `json:"objects"`
	// Specifies a list of disks to exclude from the backup.
	GlobalExcludeDisks []DiskInfo `json:"globalExcludeDisks,omitempty"`
	// Specifies whether or not to quiesce apps and the file system in order to take app consistent snapshots.
	AppConsistentSnapshot NullableBool `json:"appConsistentSnapshot,omitempty"`
	// Specifies whether or not to fallback to a crash consistent snapshot in the event that an app consistent snapshot fails. This parameter defaults to true and only changes the behavior of the operation if 'appConsistentSnapshot' is set to 'true'.
	FallbackToCrashConsistentSnapshot NullableBool `json:"fallbackToCrashConsistentSnapshot,omitempty"`
	// Specifies whether or not to skip backing up physical RDM disks. Physical RDM disks cannot be backed up, so if you attempt to backup a VM with physical RDM disks and this value is set to 'false', then those VM backups will fail.
	SkipPhysicalRDMDisks NullableBool `json:"skipPhysicalRDMDisks,omitempty"`
	IndexingPolicy *IndexingPolicy `json:"indexingPolicy,omitempty"`
	PrePostScript *PrePostScriptParams `json:"prePostScript,omitempty"`
	// If this field is set to true, then the backup for the objects will be performed using dedicated storage area network (SAN) instead of LAN or managment network.
	LeverageSanTransport NullableBool `json:"leverageSanTransport,omitempty"`
	// If this field is set to true and SAN transport backup fails, then backup will fallback to use NBDSSL transport. This field only applies if 'leverageSanTransport' is set to true.
	EnableNBDSSLFallback NullableBool `json:"enableNBDSSLFallback,omitempty"`
}

// NewVmwareObjectProtectionRequestParams instantiates a new VmwareObjectProtectionRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmwareObjectProtectionRequestParams(objects []VmwareObjectProtectionRequest) *VmwareObjectProtectionRequestParams {
	this := VmwareObjectProtectionRequestParams{}
	this.Objects = objects
	return &this
}

// NewVmwareObjectProtectionRequestParamsWithDefaults instantiates a new VmwareObjectProtectionRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmwareObjectProtectionRequestParamsWithDefaults() *VmwareObjectProtectionRequestParams {
	this := VmwareObjectProtectionRequestParams{}
	return &this
}

// GetObjects returns the Objects field value
func (o *VmwareObjectProtectionRequestParams) GetObjects() []VmwareObjectProtectionRequest {
	if o == nil {
		var ret []VmwareObjectProtectionRequest
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
func (o *VmwareObjectProtectionRequestParams) GetObjectsOk() (*[]VmwareObjectProtectionRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Objects, true
}

// SetObjects sets field value
func (o *VmwareObjectProtectionRequestParams) SetObjects(v []VmwareObjectProtectionRequest) {
	o.Objects = v
}

// GetGlobalExcludeDisks returns the GlobalExcludeDisks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareObjectProtectionRequestParams) GetGlobalExcludeDisks() []DiskInfo {
	if o == nil  {
		var ret []DiskInfo
		return ret
	}
	return o.GlobalExcludeDisks
}

// GetGlobalExcludeDisksOk returns a tuple with the GlobalExcludeDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareObjectProtectionRequestParams) GetGlobalExcludeDisksOk() (*[]DiskInfo, bool) {
	if o == nil || o.GlobalExcludeDisks == nil {
		return nil, false
	}
	return &o.GlobalExcludeDisks, true
}

// HasGlobalExcludeDisks returns a boolean if a field has been set.
func (o *VmwareObjectProtectionRequestParams) HasGlobalExcludeDisks() bool {
	if o != nil && o.GlobalExcludeDisks != nil {
		return true
	}

	return false
}

// SetGlobalExcludeDisks gets a reference to the given []DiskInfo and assigns it to the GlobalExcludeDisks field.
func (o *VmwareObjectProtectionRequestParams) SetGlobalExcludeDisks(v []DiskInfo) {
	o.GlobalExcludeDisks = v
}

// GetAppConsistentSnapshot returns the AppConsistentSnapshot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareObjectProtectionRequestParams) GetAppConsistentSnapshot() bool {
	if o == nil || o.AppConsistentSnapshot.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AppConsistentSnapshot.Get()
}

// GetAppConsistentSnapshotOk returns a tuple with the AppConsistentSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareObjectProtectionRequestParams) GetAppConsistentSnapshotOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AppConsistentSnapshot.Get(), o.AppConsistentSnapshot.IsSet()
}

// HasAppConsistentSnapshot returns a boolean if a field has been set.
func (o *VmwareObjectProtectionRequestParams) HasAppConsistentSnapshot() bool {
	if o != nil && o.AppConsistentSnapshot.IsSet() {
		return true
	}

	return false
}

// SetAppConsistentSnapshot gets a reference to the given NullableBool and assigns it to the AppConsistentSnapshot field.
func (o *VmwareObjectProtectionRequestParams) SetAppConsistentSnapshot(v bool) {
	o.AppConsistentSnapshot.Set(&v)
}
// SetAppConsistentSnapshotNil sets the value for AppConsistentSnapshot to be an explicit nil
func (o *VmwareObjectProtectionRequestParams) SetAppConsistentSnapshotNil() {
	o.AppConsistentSnapshot.Set(nil)
}

// UnsetAppConsistentSnapshot ensures that no value is present for AppConsistentSnapshot, not even an explicit nil
func (o *VmwareObjectProtectionRequestParams) UnsetAppConsistentSnapshot() {
	o.AppConsistentSnapshot.Unset()
}

// GetFallbackToCrashConsistentSnapshot returns the FallbackToCrashConsistentSnapshot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareObjectProtectionRequestParams) GetFallbackToCrashConsistentSnapshot() bool {
	if o == nil || o.FallbackToCrashConsistentSnapshot.Get() == nil {
		var ret bool
		return ret
	}
	return *o.FallbackToCrashConsistentSnapshot.Get()
}

// GetFallbackToCrashConsistentSnapshotOk returns a tuple with the FallbackToCrashConsistentSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareObjectProtectionRequestParams) GetFallbackToCrashConsistentSnapshotOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FallbackToCrashConsistentSnapshot.Get(), o.FallbackToCrashConsistentSnapshot.IsSet()
}

// HasFallbackToCrashConsistentSnapshot returns a boolean if a field has been set.
func (o *VmwareObjectProtectionRequestParams) HasFallbackToCrashConsistentSnapshot() bool {
	if o != nil && o.FallbackToCrashConsistentSnapshot.IsSet() {
		return true
	}

	return false
}

// SetFallbackToCrashConsistentSnapshot gets a reference to the given NullableBool and assigns it to the FallbackToCrashConsistentSnapshot field.
func (o *VmwareObjectProtectionRequestParams) SetFallbackToCrashConsistentSnapshot(v bool) {
	o.FallbackToCrashConsistentSnapshot.Set(&v)
}
// SetFallbackToCrashConsistentSnapshotNil sets the value for FallbackToCrashConsistentSnapshot to be an explicit nil
func (o *VmwareObjectProtectionRequestParams) SetFallbackToCrashConsistentSnapshotNil() {
	o.FallbackToCrashConsistentSnapshot.Set(nil)
}

// UnsetFallbackToCrashConsistentSnapshot ensures that no value is present for FallbackToCrashConsistentSnapshot, not even an explicit nil
func (o *VmwareObjectProtectionRequestParams) UnsetFallbackToCrashConsistentSnapshot() {
	o.FallbackToCrashConsistentSnapshot.Unset()
}

// GetSkipPhysicalRDMDisks returns the SkipPhysicalRDMDisks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareObjectProtectionRequestParams) GetSkipPhysicalRDMDisks() bool {
	if o == nil || o.SkipPhysicalRDMDisks.Get() == nil {
		var ret bool
		return ret
	}
	return *o.SkipPhysicalRDMDisks.Get()
}

// GetSkipPhysicalRDMDisksOk returns a tuple with the SkipPhysicalRDMDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareObjectProtectionRequestParams) GetSkipPhysicalRDMDisksOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SkipPhysicalRDMDisks.Get(), o.SkipPhysicalRDMDisks.IsSet()
}

// HasSkipPhysicalRDMDisks returns a boolean if a field has been set.
func (o *VmwareObjectProtectionRequestParams) HasSkipPhysicalRDMDisks() bool {
	if o != nil && o.SkipPhysicalRDMDisks.IsSet() {
		return true
	}

	return false
}

// SetSkipPhysicalRDMDisks gets a reference to the given NullableBool and assigns it to the SkipPhysicalRDMDisks field.
func (o *VmwareObjectProtectionRequestParams) SetSkipPhysicalRDMDisks(v bool) {
	o.SkipPhysicalRDMDisks.Set(&v)
}
// SetSkipPhysicalRDMDisksNil sets the value for SkipPhysicalRDMDisks to be an explicit nil
func (o *VmwareObjectProtectionRequestParams) SetSkipPhysicalRDMDisksNil() {
	o.SkipPhysicalRDMDisks.Set(nil)
}

// UnsetSkipPhysicalRDMDisks ensures that no value is present for SkipPhysicalRDMDisks, not even an explicit nil
func (o *VmwareObjectProtectionRequestParams) UnsetSkipPhysicalRDMDisks() {
	o.SkipPhysicalRDMDisks.Unset()
}

// GetIndexingPolicy returns the IndexingPolicy field value if set, zero value otherwise.
func (o *VmwareObjectProtectionRequestParams) GetIndexingPolicy() IndexingPolicy {
	if o == nil || o.IndexingPolicy == nil {
		var ret IndexingPolicy
		return ret
	}
	return *o.IndexingPolicy
}

// GetIndexingPolicyOk returns a tuple with the IndexingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareObjectProtectionRequestParams) GetIndexingPolicyOk() (*IndexingPolicy, bool) {
	if o == nil || o.IndexingPolicy == nil {
		return nil, false
	}
	return o.IndexingPolicy, true
}

// HasIndexingPolicy returns a boolean if a field has been set.
func (o *VmwareObjectProtectionRequestParams) HasIndexingPolicy() bool {
	if o != nil && o.IndexingPolicy != nil {
		return true
	}

	return false
}

// SetIndexingPolicy gets a reference to the given IndexingPolicy and assigns it to the IndexingPolicy field.
func (o *VmwareObjectProtectionRequestParams) SetIndexingPolicy(v IndexingPolicy) {
	o.IndexingPolicy = &v
}

// GetPrePostScript returns the PrePostScript field value if set, zero value otherwise.
func (o *VmwareObjectProtectionRequestParams) GetPrePostScript() PrePostScriptParams {
	if o == nil || o.PrePostScript == nil {
		var ret PrePostScriptParams
		return ret
	}
	return *o.PrePostScript
}

// GetPrePostScriptOk returns a tuple with the PrePostScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareObjectProtectionRequestParams) GetPrePostScriptOk() (*PrePostScriptParams, bool) {
	if o == nil || o.PrePostScript == nil {
		return nil, false
	}
	return o.PrePostScript, true
}

// HasPrePostScript returns a boolean if a field has been set.
func (o *VmwareObjectProtectionRequestParams) HasPrePostScript() bool {
	if o != nil && o.PrePostScript != nil {
		return true
	}

	return false
}

// SetPrePostScript gets a reference to the given PrePostScriptParams and assigns it to the PrePostScript field.
func (o *VmwareObjectProtectionRequestParams) SetPrePostScript(v PrePostScriptParams) {
	o.PrePostScript = &v
}

// GetLeverageSanTransport returns the LeverageSanTransport field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareObjectProtectionRequestParams) GetLeverageSanTransport() bool {
	if o == nil || o.LeverageSanTransport.Get() == nil {
		var ret bool
		return ret
	}
	return *o.LeverageSanTransport.Get()
}

// GetLeverageSanTransportOk returns a tuple with the LeverageSanTransport field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareObjectProtectionRequestParams) GetLeverageSanTransportOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LeverageSanTransport.Get(), o.LeverageSanTransport.IsSet()
}

// HasLeverageSanTransport returns a boolean if a field has been set.
func (o *VmwareObjectProtectionRequestParams) HasLeverageSanTransport() bool {
	if o != nil && o.LeverageSanTransport.IsSet() {
		return true
	}

	return false
}

// SetLeverageSanTransport gets a reference to the given NullableBool and assigns it to the LeverageSanTransport field.
func (o *VmwareObjectProtectionRequestParams) SetLeverageSanTransport(v bool) {
	o.LeverageSanTransport.Set(&v)
}
// SetLeverageSanTransportNil sets the value for LeverageSanTransport to be an explicit nil
func (o *VmwareObjectProtectionRequestParams) SetLeverageSanTransportNil() {
	o.LeverageSanTransport.Set(nil)
}

// UnsetLeverageSanTransport ensures that no value is present for LeverageSanTransport, not even an explicit nil
func (o *VmwareObjectProtectionRequestParams) UnsetLeverageSanTransport() {
	o.LeverageSanTransport.Unset()
}

// GetEnableNBDSSLFallback returns the EnableNBDSSLFallback field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareObjectProtectionRequestParams) GetEnableNBDSSLFallback() bool {
	if o == nil || o.EnableNBDSSLFallback.Get() == nil {
		var ret bool
		return ret
	}
	return *o.EnableNBDSSLFallback.Get()
}

// GetEnableNBDSSLFallbackOk returns a tuple with the EnableNBDSSLFallback field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareObjectProtectionRequestParams) GetEnableNBDSSLFallbackOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EnableNBDSSLFallback.Get(), o.EnableNBDSSLFallback.IsSet()
}

// HasEnableNBDSSLFallback returns a boolean if a field has been set.
func (o *VmwareObjectProtectionRequestParams) HasEnableNBDSSLFallback() bool {
	if o != nil && o.EnableNBDSSLFallback.IsSet() {
		return true
	}

	return false
}

// SetEnableNBDSSLFallback gets a reference to the given NullableBool and assigns it to the EnableNBDSSLFallback field.
func (o *VmwareObjectProtectionRequestParams) SetEnableNBDSSLFallback(v bool) {
	o.EnableNBDSSLFallback.Set(&v)
}
// SetEnableNBDSSLFallbackNil sets the value for EnableNBDSSLFallback to be an explicit nil
func (o *VmwareObjectProtectionRequestParams) SetEnableNBDSSLFallbackNil() {
	o.EnableNBDSSLFallback.Set(nil)
}

// UnsetEnableNBDSSLFallback ensures that no value is present for EnableNBDSSLFallback, not even an explicit nil
func (o *VmwareObjectProtectionRequestParams) UnsetEnableNBDSSLFallback() {
	o.EnableNBDSSLFallback.Unset()
}

func (o VmwareObjectProtectionRequestParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["objects"] = o.Objects
	}
	if o.GlobalExcludeDisks != nil {
		toSerialize["globalExcludeDisks"] = o.GlobalExcludeDisks
	}
	if o.AppConsistentSnapshot.IsSet() {
		toSerialize["appConsistentSnapshot"] = o.AppConsistentSnapshot.Get()
	}
	if o.FallbackToCrashConsistentSnapshot.IsSet() {
		toSerialize["fallbackToCrashConsistentSnapshot"] = o.FallbackToCrashConsistentSnapshot.Get()
	}
	if o.SkipPhysicalRDMDisks.IsSet() {
		toSerialize["skipPhysicalRDMDisks"] = o.SkipPhysicalRDMDisks.Get()
	}
	if o.IndexingPolicy != nil {
		toSerialize["indexingPolicy"] = o.IndexingPolicy
	}
	if o.PrePostScript != nil {
		toSerialize["prePostScript"] = o.PrePostScript
	}
	if o.LeverageSanTransport.IsSet() {
		toSerialize["leverageSanTransport"] = o.LeverageSanTransport.Get()
	}
	if o.EnableNBDSSLFallback.IsSet() {
		toSerialize["enableNBDSSLFallback"] = o.EnableNBDSSLFallback.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableVmwareObjectProtectionRequestParams struct {
	value *VmwareObjectProtectionRequestParams
	isSet bool
}

func (v NullableVmwareObjectProtectionRequestParams) Get() *VmwareObjectProtectionRequestParams {
	return v.value
}

func (v *NullableVmwareObjectProtectionRequestParams) Set(val *VmwareObjectProtectionRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableVmwareObjectProtectionRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableVmwareObjectProtectionRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmwareObjectProtectionRequestParams(val *VmwareObjectProtectionRequestParams) *NullableVmwareObjectProtectionRequestParams {
	return &NullableVmwareObjectProtectionRequestParams{value: val, isSet: true}
}

func (v NullableVmwareObjectProtectionRequestParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmwareObjectProtectionRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o VmwareObjectProtectionRequestParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}