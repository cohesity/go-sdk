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

// VmwareObjectProtectionResponseParams Specifies the parameters which are specific to VMware object protection.
type VmwareObjectProtectionResponseParams struct {
	// Specifies a list of disks to exclude from being protected. This is only applicable to VM objects.
	ExcludeDisks *[]DiskInfo `json:"excludeDisks,omitempty"`
	// Specifies whether or not to truncate MS Exchange logs while taking an app consistent snapshot of this object. This is only applicable to objects which have a registered MS Exchange app.
	TruncateExchangeLogs NullableBool `json:"truncateExchangeLogs,omitempty"`
	// Specifies the list of IDs of the objects to not be protected in this backup. This field only applies if provided object id is non leaf entity such as Tag or a folder. This can be used to ignore specific objects under a parent object which has been included for protection.
	ExcludeObjectIds *[]int64 `json:"excludeObjectIds,omitempty"`
	CdpInfo *VmwareCdpObject `json:"cdpInfo,omitempty"`
	StandbyInfo *VmwareStandbyObject `json:"standbyInfo,omitempty"`
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

// NewVmwareObjectProtectionResponseParams instantiates a new VmwareObjectProtectionResponseParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmwareObjectProtectionResponseParams() *VmwareObjectProtectionResponseParams {
	this := VmwareObjectProtectionResponseParams{}
	return &this
}

// NewVmwareObjectProtectionResponseParamsWithDefaults instantiates a new VmwareObjectProtectionResponseParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmwareObjectProtectionResponseParamsWithDefaults() *VmwareObjectProtectionResponseParams {
	this := VmwareObjectProtectionResponseParams{}
	return &this
}

// GetExcludeDisks returns the ExcludeDisks field value if set, zero value otherwise.
func (o *VmwareObjectProtectionResponseParams) GetExcludeDisks() []DiskInfo {
	if o == nil || o.ExcludeDisks == nil {
		var ret []DiskInfo
		return ret
	}
	return *o.ExcludeDisks
}

// GetExcludeDisksOk returns a tuple with the ExcludeDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareObjectProtectionResponseParams) GetExcludeDisksOk() (*[]DiskInfo, bool) {
	if o == nil || o.ExcludeDisks == nil {
		return nil, false
	}
	return o.ExcludeDisks, true
}

// HasExcludeDisks returns a boolean if a field has been set.
func (o *VmwareObjectProtectionResponseParams) HasExcludeDisks() bool {
	if o != nil && o.ExcludeDisks != nil {
		return true
	}

	return false
}

// SetExcludeDisks gets a reference to the given []DiskInfo and assigns it to the ExcludeDisks field.
func (o *VmwareObjectProtectionResponseParams) SetExcludeDisks(v []DiskInfo) {
	o.ExcludeDisks = &v
}

// GetTruncateExchangeLogs returns the TruncateExchangeLogs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareObjectProtectionResponseParams) GetTruncateExchangeLogs() bool {
	if o == nil || o.TruncateExchangeLogs.Get() == nil {
		var ret bool
		return ret
	}
	return *o.TruncateExchangeLogs.Get()
}

// GetTruncateExchangeLogsOk returns a tuple with the TruncateExchangeLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareObjectProtectionResponseParams) GetTruncateExchangeLogsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TruncateExchangeLogs.Get(), o.TruncateExchangeLogs.IsSet()
}

// HasTruncateExchangeLogs returns a boolean if a field has been set.
func (o *VmwareObjectProtectionResponseParams) HasTruncateExchangeLogs() bool {
	if o != nil && o.TruncateExchangeLogs.IsSet() {
		return true
	}

	return false
}

// SetTruncateExchangeLogs gets a reference to the given NullableBool and assigns it to the TruncateExchangeLogs field.
func (o *VmwareObjectProtectionResponseParams) SetTruncateExchangeLogs(v bool) {
	o.TruncateExchangeLogs.Set(&v)
}
// SetTruncateExchangeLogsNil sets the value for TruncateExchangeLogs to be an explicit nil
func (o *VmwareObjectProtectionResponseParams) SetTruncateExchangeLogsNil() {
	o.TruncateExchangeLogs.Set(nil)
}

// UnsetTruncateExchangeLogs ensures that no value is present for TruncateExchangeLogs, not even an explicit nil
func (o *VmwareObjectProtectionResponseParams) UnsetTruncateExchangeLogs() {
	o.TruncateExchangeLogs.Unset()
}

// GetExcludeObjectIds returns the ExcludeObjectIds field value if set, zero value otherwise.
func (o *VmwareObjectProtectionResponseParams) GetExcludeObjectIds() []int64 {
	if o == nil || o.ExcludeObjectIds == nil {
		var ret []int64
		return ret
	}
	return *o.ExcludeObjectIds
}

// GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareObjectProtectionResponseParams) GetExcludeObjectIdsOk() (*[]int64, bool) {
	if o == nil || o.ExcludeObjectIds == nil {
		return nil, false
	}
	return o.ExcludeObjectIds, true
}

// HasExcludeObjectIds returns a boolean if a field has been set.
func (o *VmwareObjectProtectionResponseParams) HasExcludeObjectIds() bool {
	if o != nil && o.ExcludeObjectIds != nil {
		return true
	}

	return false
}

// SetExcludeObjectIds gets a reference to the given []int64 and assigns it to the ExcludeObjectIds field.
func (o *VmwareObjectProtectionResponseParams) SetExcludeObjectIds(v []int64) {
	o.ExcludeObjectIds = &v
}

// GetCdpInfo returns the CdpInfo field value if set, zero value otherwise.
func (o *VmwareObjectProtectionResponseParams) GetCdpInfo() VmwareCdpObject {
	if o == nil || o.CdpInfo == nil {
		var ret VmwareCdpObject
		return ret
	}
	return *o.CdpInfo
}

// GetCdpInfoOk returns a tuple with the CdpInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareObjectProtectionResponseParams) GetCdpInfoOk() (*VmwareCdpObject, bool) {
	if o == nil || o.CdpInfo == nil {
		return nil, false
	}
	return o.CdpInfo, true
}

// HasCdpInfo returns a boolean if a field has been set.
func (o *VmwareObjectProtectionResponseParams) HasCdpInfo() bool {
	if o != nil && o.CdpInfo != nil {
		return true
	}

	return false
}

// SetCdpInfo gets a reference to the given VmwareCdpObject and assigns it to the CdpInfo field.
func (o *VmwareObjectProtectionResponseParams) SetCdpInfo(v VmwareCdpObject) {
	o.CdpInfo = &v
}

// GetStandbyInfo returns the StandbyInfo field value if set, zero value otherwise.
func (o *VmwareObjectProtectionResponseParams) GetStandbyInfo() VmwareStandbyObject {
	if o == nil || o.StandbyInfo == nil {
		var ret VmwareStandbyObject
		return ret
	}
	return *o.StandbyInfo
}

// GetStandbyInfoOk returns a tuple with the StandbyInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareObjectProtectionResponseParams) GetStandbyInfoOk() (*VmwareStandbyObject, bool) {
	if o == nil || o.StandbyInfo == nil {
		return nil, false
	}
	return o.StandbyInfo, true
}

// HasStandbyInfo returns a boolean if a field has been set.
func (o *VmwareObjectProtectionResponseParams) HasStandbyInfo() bool {
	if o != nil && o.StandbyInfo != nil {
		return true
	}

	return false
}

// SetStandbyInfo gets a reference to the given VmwareStandbyObject and assigns it to the StandbyInfo field.
func (o *VmwareObjectProtectionResponseParams) SetStandbyInfo(v VmwareStandbyObject) {
	o.StandbyInfo = &v
}

// GetAppConsistentSnapshot returns the AppConsistentSnapshot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareObjectProtectionResponseParams) GetAppConsistentSnapshot() bool {
	if o == nil || o.AppConsistentSnapshot.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AppConsistentSnapshot.Get()
}

// GetAppConsistentSnapshotOk returns a tuple with the AppConsistentSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareObjectProtectionResponseParams) GetAppConsistentSnapshotOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AppConsistentSnapshot.Get(), o.AppConsistentSnapshot.IsSet()
}

// HasAppConsistentSnapshot returns a boolean if a field has been set.
func (o *VmwareObjectProtectionResponseParams) HasAppConsistentSnapshot() bool {
	if o != nil && o.AppConsistentSnapshot.IsSet() {
		return true
	}

	return false
}

// SetAppConsistentSnapshot gets a reference to the given NullableBool and assigns it to the AppConsistentSnapshot field.
func (o *VmwareObjectProtectionResponseParams) SetAppConsistentSnapshot(v bool) {
	o.AppConsistentSnapshot.Set(&v)
}
// SetAppConsistentSnapshotNil sets the value for AppConsistentSnapshot to be an explicit nil
func (o *VmwareObjectProtectionResponseParams) SetAppConsistentSnapshotNil() {
	o.AppConsistentSnapshot.Set(nil)
}

// UnsetAppConsistentSnapshot ensures that no value is present for AppConsistentSnapshot, not even an explicit nil
func (o *VmwareObjectProtectionResponseParams) UnsetAppConsistentSnapshot() {
	o.AppConsistentSnapshot.Unset()
}

// GetFallbackToCrashConsistentSnapshot returns the FallbackToCrashConsistentSnapshot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareObjectProtectionResponseParams) GetFallbackToCrashConsistentSnapshot() bool {
	if o == nil || o.FallbackToCrashConsistentSnapshot.Get() == nil {
		var ret bool
		return ret
	}
	return *o.FallbackToCrashConsistentSnapshot.Get()
}

// GetFallbackToCrashConsistentSnapshotOk returns a tuple with the FallbackToCrashConsistentSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareObjectProtectionResponseParams) GetFallbackToCrashConsistentSnapshotOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FallbackToCrashConsistentSnapshot.Get(), o.FallbackToCrashConsistentSnapshot.IsSet()
}

// HasFallbackToCrashConsistentSnapshot returns a boolean if a field has been set.
func (o *VmwareObjectProtectionResponseParams) HasFallbackToCrashConsistentSnapshot() bool {
	if o != nil && o.FallbackToCrashConsistentSnapshot.IsSet() {
		return true
	}

	return false
}

// SetFallbackToCrashConsistentSnapshot gets a reference to the given NullableBool and assigns it to the FallbackToCrashConsistentSnapshot field.
func (o *VmwareObjectProtectionResponseParams) SetFallbackToCrashConsistentSnapshot(v bool) {
	o.FallbackToCrashConsistentSnapshot.Set(&v)
}
// SetFallbackToCrashConsistentSnapshotNil sets the value for FallbackToCrashConsistentSnapshot to be an explicit nil
func (o *VmwareObjectProtectionResponseParams) SetFallbackToCrashConsistentSnapshotNil() {
	o.FallbackToCrashConsistentSnapshot.Set(nil)
}

// UnsetFallbackToCrashConsistentSnapshot ensures that no value is present for FallbackToCrashConsistentSnapshot, not even an explicit nil
func (o *VmwareObjectProtectionResponseParams) UnsetFallbackToCrashConsistentSnapshot() {
	o.FallbackToCrashConsistentSnapshot.Unset()
}

// GetSkipPhysicalRDMDisks returns the SkipPhysicalRDMDisks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareObjectProtectionResponseParams) GetSkipPhysicalRDMDisks() bool {
	if o == nil || o.SkipPhysicalRDMDisks.Get() == nil {
		var ret bool
		return ret
	}
	return *o.SkipPhysicalRDMDisks.Get()
}

// GetSkipPhysicalRDMDisksOk returns a tuple with the SkipPhysicalRDMDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareObjectProtectionResponseParams) GetSkipPhysicalRDMDisksOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SkipPhysicalRDMDisks.Get(), o.SkipPhysicalRDMDisks.IsSet()
}

// HasSkipPhysicalRDMDisks returns a boolean if a field has been set.
func (o *VmwareObjectProtectionResponseParams) HasSkipPhysicalRDMDisks() bool {
	if o != nil && o.SkipPhysicalRDMDisks.IsSet() {
		return true
	}

	return false
}

// SetSkipPhysicalRDMDisks gets a reference to the given NullableBool and assigns it to the SkipPhysicalRDMDisks field.
func (o *VmwareObjectProtectionResponseParams) SetSkipPhysicalRDMDisks(v bool) {
	o.SkipPhysicalRDMDisks.Set(&v)
}
// SetSkipPhysicalRDMDisksNil sets the value for SkipPhysicalRDMDisks to be an explicit nil
func (o *VmwareObjectProtectionResponseParams) SetSkipPhysicalRDMDisksNil() {
	o.SkipPhysicalRDMDisks.Set(nil)
}

// UnsetSkipPhysicalRDMDisks ensures that no value is present for SkipPhysicalRDMDisks, not even an explicit nil
func (o *VmwareObjectProtectionResponseParams) UnsetSkipPhysicalRDMDisks() {
	o.SkipPhysicalRDMDisks.Unset()
}

// GetIndexingPolicy returns the IndexingPolicy field value if set, zero value otherwise.
func (o *VmwareObjectProtectionResponseParams) GetIndexingPolicy() IndexingPolicy {
	if o == nil || o.IndexingPolicy == nil {
		var ret IndexingPolicy
		return ret
	}
	return *o.IndexingPolicy
}

// GetIndexingPolicyOk returns a tuple with the IndexingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareObjectProtectionResponseParams) GetIndexingPolicyOk() (*IndexingPolicy, bool) {
	if o == nil || o.IndexingPolicy == nil {
		return nil, false
	}
	return o.IndexingPolicy, true
}

// HasIndexingPolicy returns a boolean if a field has been set.
func (o *VmwareObjectProtectionResponseParams) HasIndexingPolicy() bool {
	if o != nil && o.IndexingPolicy != nil {
		return true
	}

	return false
}

// SetIndexingPolicy gets a reference to the given IndexingPolicy and assigns it to the IndexingPolicy field.
func (o *VmwareObjectProtectionResponseParams) SetIndexingPolicy(v IndexingPolicy) {
	o.IndexingPolicy = &v
}

// GetPrePostScript returns the PrePostScript field value if set, zero value otherwise.
func (o *VmwareObjectProtectionResponseParams) GetPrePostScript() PrePostScriptParams {
	if o == nil || o.PrePostScript == nil {
		var ret PrePostScriptParams
		return ret
	}
	return *o.PrePostScript
}

// GetPrePostScriptOk returns a tuple with the PrePostScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareObjectProtectionResponseParams) GetPrePostScriptOk() (*PrePostScriptParams, bool) {
	if o == nil || o.PrePostScript == nil {
		return nil, false
	}
	return o.PrePostScript, true
}

// HasPrePostScript returns a boolean if a field has been set.
func (o *VmwareObjectProtectionResponseParams) HasPrePostScript() bool {
	if o != nil && o.PrePostScript != nil {
		return true
	}

	return false
}

// SetPrePostScript gets a reference to the given PrePostScriptParams and assigns it to the PrePostScript field.
func (o *VmwareObjectProtectionResponseParams) SetPrePostScript(v PrePostScriptParams) {
	o.PrePostScript = &v
}

// GetLeverageSanTransport returns the LeverageSanTransport field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareObjectProtectionResponseParams) GetLeverageSanTransport() bool {
	if o == nil || o.LeverageSanTransport.Get() == nil {
		var ret bool
		return ret
	}
	return *o.LeverageSanTransport.Get()
}

// GetLeverageSanTransportOk returns a tuple with the LeverageSanTransport field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareObjectProtectionResponseParams) GetLeverageSanTransportOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LeverageSanTransport.Get(), o.LeverageSanTransport.IsSet()
}

// HasLeverageSanTransport returns a boolean if a field has been set.
func (o *VmwareObjectProtectionResponseParams) HasLeverageSanTransport() bool {
	if o != nil && o.LeverageSanTransport.IsSet() {
		return true
	}

	return false
}

// SetLeverageSanTransport gets a reference to the given NullableBool and assigns it to the LeverageSanTransport field.
func (o *VmwareObjectProtectionResponseParams) SetLeverageSanTransport(v bool) {
	o.LeverageSanTransport.Set(&v)
}
// SetLeverageSanTransportNil sets the value for LeverageSanTransport to be an explicit nil
func (o *VmwareObjectProtectionResponseParams) SetLeverageSanTransportNil() {
	o.LeverageSanTransport.Set(nil)
}

// UnsetLeverageSanTransport ensures that no value is present for LeverageSanTransport, not even an explicit nil
func (o *VmwareObjectProtectionResponseParams) UnsetLeverageSanTransport() {
	o.LeverageSanTransport.Unset()
}

// GetEnableNBDSSLFallback returns the EnableNBDSSLFallback field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareObjectProtectionResponseParams) GetEnableNBDSSLFallback() bool {
	if o == nil || o.EnableNBDSSLFallback.Get() == nil {
		var ret bool
		return ret
	}
	return *o.EnableNBDSSLFallback.Get()
}

// GetEnableNBDSSLFallbackOk returns a tuple with the EnableNBDSSLFallback field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareObjectProtectionResponseParams) GetEnableNBDSSLFallbackOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EnableNBDSSLFallback.Get(), o.EnableNBDSSLFallback.IsSet()
}

// HasEnableNBDSSLFallback returns a boolean if a field has been set.
func (o *VmwareObjectProtectionResponseParams) HasEnableNBDSSLFallback() bool {
	if o != nil && o.EnableNBDSSLFallback.IsSet() {
		return true
	}

	return false
}

// SetEnableNBDSSLFallback gets a reference to the given NullableBool and assigns it to the EnableNBDSSLFallback field.
func (o *VmwareObjectProtectionResponseParams) SetEnableNBDSSLFallback(v bool) {
	o.EnableNBDSSLFallback.Set(&v)
}
// SetEnableNBDSSLFallbackNil sets the value for EnableNBDSSLFallback to be an explicit nil
func (o *VmwareObjectProtectionResponseParams) SetEnableNBDSSLFallbackNil() {
	o.EnableNBDSSLFallback.Set(nil)
}

// UnsetEnableNBDSSLFallback ensures that no value is present for EnableNBDSSLFallback, not even an explicit nil
func (o *VmwareObjectProtectionResponseParams) UnsetEnableNBDSSLFallback() {
	o.EnableNBDSSLFallback.Unset()
}

func (o VmwareObjectProtectionResponseParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExcludeDisks != nil {
		toSerialize["excludeDisks"] = o.ExcludeDisks
	}
	if o.TruncateExchangeLogs.IsSet() {
		toSerialize["truncateExchangeLogs"] = o.TruncateExchangeLogs.Get()
	}
	if o.ExcludeObjectIds != nil {
		toSerialize["excludeObjectIds"] = o.ExcludeObjectIds
	}
	if o.CdpInfo != nil {
		toSerialize["cdpInfo"] = o.CdpInfo
	}
	if o.StandbyInfo != nil {
		toSerialize["standbyInfo"] = o.StandbyInfo
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

type NullableVmwareObjectProtectionResponseParams struct {
	value *VmwareObjectProtectionResponseParams
	isSet bool
}

func (v NullableVmwareObjectProtectionResponseParams) Get() *VmwareObjectProtectionResponseParams {
	return v.value
}

func (v *NullableVmwareObjectProtectionResponseParams) Set(val *VmwareObjectProtectionResponseParams) {
	v.value = val
	v.isSet = true
}

func (v NullableVmwareObjectProtectionResponseParams) IsSet() bool {
	return v.isSet
}

func (v *NullableVmwareObjectProtectionResponseParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmwareObjectProtectionResponseParams(val *VmwareObjectProtectionResponseParams) *NullableVmwareObjectProtectionResponseParams {
	return &NullableVmwareObjectProtectionResponseParams{value: val, isSet: true}
}

func (v NullableVmwareObjectProtectionResponseParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmwareObjectProtectionResponseParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o VmwareObjectProtectionResponseParams) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}