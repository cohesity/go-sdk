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

// VmwareProtectionGroupObjectParams Specifies the input for a protection object in the VMware environment.
type VmwareProtectionGroupObjectParams struct {
	// Specifies the id of the object being protected. This can be a leaf level or non leaf level object.
	Id NullableInt64 `json:"id"`
	// Specifies the name of the virtual machine.
	Name NullableString `json:"name,omitempty"`
	// Specifies whether the vm is part of an Autoprotection. True implies that the vm or its parent directory is autoprotected and will remain part of the autoprotection with additional settings specified here. False implies the object is not part of an Autoprotection and will remain protected and its individual settings here even if a parent directory's Autoprotection setting is altered. Default is false.
	IsAutoprotected NullableBool `json:"isAutoprotected,omitempty"`
	// Specifies the CDP related information for a given object. This field will only be populated when protection group is configured with policy having CDP retnetion settings.
	CdpInfo *VmwareCdpObject `json:"cdpInfo,omitempty"`
	// Specifies the standby related information for a given object. This field will only be populated when standby is configured in backup job settings.
	StandbyInfo *VmwareStandbyObject `json:"standbyInfo,omitempty"`
	// Specifies a list of disks to exclude from being protected. This is only applicable to VM objects.
	ExcludeDisks *[]DiskInfo `json:"excludeDisks,omitempty"`
	// Specifies whether or not to truncate MS Exchange logs while taking an app consistent snapshot of this object. This is only applicable to objects which have a registered MS Exchange app.
	TruncateExchangeLogs NullableBool `json:"truncateExchangeLogs,omitempty"`
}

// NewVmwareProtectionGroupObjectParams instantiates a new VmwareProtectionGroupObjectParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmwareProtectionGroupObjectParams(id NullableInt64) *VmwareProtectionGroupObjectParams {
	this := VmwareProtectionGroupObjectParams{}
	this.Id = id
	return &this
}

// NewVmwareProtectionGroupObjectParamsWithDefaults instantiates a new VmwareProtectionGroupObjectParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmwareProtectionGroupObjectParamsWithDefaults() *VmwareProtectionGroupObjectParams {
	this := VmwareProtectionGroupObjectParams{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *VmwareProtectionGroupObjectParams) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareProtectionGroupObjectParams) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *VmwareProtectionGroupObjectParams) SetId(v int64) {
	o.Id.Set(&v)
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareProtectionGroupObjectParams) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareProtectionGroupObjectParams) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *VmwareProtectionGroupObjectParams) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *VmwareProtectionGroupObjectParams) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *VmwareProtectionGroupObjectParams) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *VmwareProtectionGroupObjectParams) UnsetName() {
	o.Name.Unset()
}

// GetIsAutoprotected returns the IsAutoprotected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareProtectionGroupObjectParams) GetIsAutoprotected() bool {
	if o == nil || o.IsAutoprotected.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsAutoprotected.Get()
}

// GetIsAutoprotectedOk returns a tuple with the IsAutoprotected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareProtectionGroupObjectParams) GetIsAutoprotectedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsAutoprotected.Get(), o.IsAutoprotected.IsSet()
}

// HasIsAutoprotected returns a boolean if a field has been set.
func (o *VmwareProtectionGroupObjectParams) HasIsAutoprotected() bool {
	if o != nil && o.IsAutoprotected.IsSet() {
		return true
	}

	return false
}

// SetIsAutoprotected gets a reference to the given NullableBool and assigns it to the IsAutoprotected field.
func (o *VmwareProtectionGroupObjectParams) SetIsAutoprotected(v bool) {
	o.IsAutoprotected.Set(&v)
}
// SetIsAutoprotectedNil sets the value for IsAutoprotected to be an explicit nil
func (o *VmwareProtectionGroupObjectParams) SetIsAutoprotectedNil() {
	o.IsAutoprotected.Set(nil)
}

// UnsetIsAutoprotected ensures that no value is present for IsAutoprotected, not even an explicit nil
func (o *VmwareProtectionGroupObjectParams) UnsetIsAutoprotected() {
	o.IsAutoprotected.Unset()
}

// GetCdpInfo returns the CdpInfo field value if set, zero value otherwise.
func (o *VmwareProtectionGroupObjectParams) GetCdpInfo() VmwareCdpObject {
	if o == nil || o.CdpInfo == nil {
		var ret VmwareCdpObject
		return ret
	}
	return *o.CdpInfo
}

// GetCdpInfoOk returns a tuple with the CdpInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareProtectionGroupObjectParams) GetCdpInfoOk() (*VmwareCdpObject, bool) {
	if o == nil || o.CdpInfo == nil {
		return nil, false
	}
	return o.CdpInfo, true
}

// HasCdpInfo returns a boolean if a field has been set.
func (o *VmwareProtectionGroupObjectParams) HasCdpInfo() bool {
	if o != nil && o.CdpInfo != nil {
		return true
	}

	return false
}

// SetCdpInfo gets a reference to the given VmwareCdpObject and assigns it to the CdpInfo field.
func (o *VmwareProtectionGroupObjectParams) SetCdpInfo(v VmwareCdpObject) {
	o.CdpInfo = &v
}

// GetStandbyInfo returns the StandbyInfo field value if set, zero value otherwise.
func (o *VmwareProtectionGroupObjectParams) GetStandbyInfo() VmwareStandbyObject {
	if o == nil || o.StandbyInfo == nil {
		var ret VmwareStandbyObject
		return ret
	}
	return *o.StandbyInfo
}

// GetStandbyInfoOk returns a tuple with the StandbyInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareProtectionGroupObjectParams) GetStandbyInfoOk() (*VmwareStandbyObject, bool) {
	if o == nil || o.StandbyInfo == nil {
		return nil, false
	}
	return o.StandbyInfo, true
}

// HasStandbyInfo returns a boolean if a field has been set.
func (o *VmwareProtectionGroupObjectParams) HasStandbyInfo() bool {
	if o != nil && o.StandbyInfo != nil {
		return true
	}

	return false
}

// SetStandbyInfo gets a reference to the given VmwareStandbyObject and assigns it to the StandbyInfo field.
func (o *VmwareProtectionGroupObjectParams) SetStandbyInfo(v VmwareStandbyObject) {
	o.StandbyInfo = &v
}

// GetExcludeDisks returns the ExcludeDisks field value if set, zero value otherwise.
func (o *VmwareProtectionGroupObjectParams) GetExcludeDisks() []DiskInfo {
	if o == nil || o.ExcludeDisks == nil {
		var ret []DiskInfo
		return ret
	}
	return *o.ExcludeDisks
}

// GetExcludeDisksOk returns a tuple with the ExcludeDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareProtectionGroupObjectParams) GetExcludeDisksOk() (*[]DiskInfo, bool) {
	if o == nil || o.ExcludeDisks == nil {
		return nil, false
	}
	return o.ExcludeDisks, true
}

// HasExcludeDisks returns a boolean if a field has been set.
func (o *VmwareProtectionGroupObjectParams) HasExcludeDisks() bool {
	if o != nil && o.ExcludeDisks != nil {
		return true
	}

	return false
}

// SetExcludeDisks gets a reference to the given []DiskInfo and assigns it to the ExcludeDisks field.
func (o *VmwareProtectionGroupObjectParams) SetExcludeDisks(v []DiskInfo) {
	o.ExcludeDisks = &v
}

// GetTruncateExchangeLogs returns the TruncateExchangeLogs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VmwareProtectionGroupObjectParams) GetTruncateExchangeLogs() bool {
	if o == nil || o.TruncateExchangeLogs.Get() == nil {
		var ret bool
		return ret
	}
	return *o.TruncateExchangeLogs.Get()
}

// GetTruncateExchangeLogsOk returns a tuple with the TruncateExchangeLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VmwareProtectionGroupObjectParams) GetTruncateExchangeLogsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TruncateExchangeLogs.Get(), o.TruncateExchangeLogs.IsSet()
}

// HasTruncateExchangeLogs returns a boolean if a field has been set.
func (o *VmwareProtectionGroupObjectParams) HasTruncateExchangeLogs() bool {
	if o != nil && o.TruncateExchangeLogs.IsSet() {
		return true
	}

	return false
}

// SetTruncateExchangeLogs gets a reference to the given NullableBool and assigns it to the TruncateExchangeLogs field.
func (o *VmwareProtectionGroupObjectParams) SetTruncateExchangeLogs(v bool) {
	o.TruncateExchangeLogs.Set(&v)
}
// SetTruncateExchangeLogsNil sets the value for TruncateExchangeLogs to be an explicit nil
func (o *VmwareProtectionGroupObjectParams) SetTruncateExchangeLogsNil() {
	o.TruncateExchangeLogs.Set(nil)
}

// UnsetTruncateExchangeLogs ensures that no value is present for TruncateExchangeLogs, not even an explicit nil
func (o *VmwareProtectionGroupObjectParams) UnsetTruncateExchangeLogs() {
	o.TruncateExchangeLogs.Unset()
}

func (o VmwareProtectionGroupObjectParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.IsAutoprotected.IsSet() {
		toSerialize["isAutoprotected"] = o.IsAutoprotected.Get()
	}
	if o.CdpInfo != nil {
		toSerialize["cdpInfo"] = o.CdpInfo
	}
	if o.StandbyInfo != nil {
		toSerialize["standbyInfo"] = o.StandbyInfo
	}
	if o.ExcludeDisks != nil {
		toSerialize["excludeDisks"] = o.ExcludeDisks
	}
	if o.TruncateExchangeLogs.IsSet() {
		toSerialize["truncateExchangeLogs"] = o.TruncateExchangeLogs.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableVmwareProtectionGroupObjectParams struct {
	value *VmwareProtectionGroupObjectParams
	isSet bool
}

func (v NullableVmwareProtectionGroupObjectParams) Get() *VmwareProtectionGroupObjectParams {
	return v.value
}

func (v *NullableVmwareProtectionGroupObjectParams) Set(val *VmwareProtectionGroupObjectParams) {
	v.value = val
	v.isSet = true
}

func (v NullableVmwareProtectionGroupObjectParams) IsSet() bool {
	return v.isSet
}

func (v *NullableVmwareProtectionGroupObjectParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmwareProtectionGroupObjectParams(val *VmwareProtectionGroupObjectParams) *NullableVmwareProtectionGroupObjectParams {
	return &NullableVmwareProtectionGroupObjectParams{value: val, isSet: true}
}

func (v NullableVmwareProtectionGroupObjectParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmwareProtectionGroupObjectParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


