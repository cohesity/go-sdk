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

// checks if the SmbConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmbConfig{}

// SmbConfig Specifies the SMB config settings for this View.
type SmbConfig struct {
	// Specifies whether fast durable handle is enabled. If enabled, view open handle will be kept in memory, which results in a higher performance. But the handles cannot be recovered if node or service crashes.
	EnableFastDurableHandle NullableBool `json:"enableFastDurableHandle,omitempty"`
	// Specifies if access-based enumeration should be enabled. If 'true', only files and folders that the user has permissions to access are visible on the SMB share for that user.
	EnableSmbAccessBasedEnumeration NullableBool `json:"enableSmbAccessBasedEnumeration,omitempty"`
	// Specifies the SMB encryption for the View. If set, it enables the SMB encryption for the View. Encryption is supported only by SMB 3.x dialects. Dialects that do not support would still access data in unencrypted format.
	EnableSmbEncryption NullableBool `json:"enableSmbEncryption,omitempty"`
	// Specifies whether SMB opportunistic lock is enabled.
	EnableSmbOplock NullableBool `json:"enableSmbOplock,omitempty"`
	// If set, it enables discovery of view for SMB.
	EnableSmbViewDiscovery NullableBool `json:"enableSmbViewDiscovery,omitempty"`
	// Specifies the SMB encryption for all the sessions for the View. If set, encryption is enforced for all the sessions for the View. When enabled all future and existing unencrypted sessions are disallowed.
	EnforceSmbEncryption NullableBool `json:"enforceSmbEncryption,omitempty"`
	SharePermissions *SmbConfigSharePermissions `json:"sharePermissions,omitempty"`
	SmbPermissionsInfo *SmbConfigSmbPermissionsInfo `json:"smbPermissionsInfo,omitempty"`
}

// NewSmbConfig instantiates a new SmbConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmbConfig() *SmbConfig {
	this := SmbConfig{}
	return &this
}

// NewSmbConfigWithDefaults instantiates a new SmbConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmbConfigWithDefaults() *SmbConfig {
	this := SmbConfig{}
	return &this
}

// GetEnableFastDurableHandle returns the EnableFastDurableHandle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmbConfig) GetEnableFastDurableHandle() bool {
	if o == nil || IsNil(o.EnableFastDurableHandle.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableFastDurableHandle.Get()
}

// GetEnableFastDurableHandleOk returns a tuple with the EnableFastDurableHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmbConfig) GetEnableFastDurableHandleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableFastDurableHandle.Get(), o.EnableFastDurableHandle.IsSet()
}

// HasEnableFastDurableHandle returns a boolean if a field has been set.
func (o *SmbConfig) HasEnableFastDurableHandle() bool {
	if o != nil && o.EnableFastDurableHandle.IsSet() {
		return true
	}

	return false
}

// SetEnableFastDurableHandle gets a reference to the given NullableBool and assigns it to the EnableFastDurableHandle field.
func (o *SmbConfig) SetEnableFastDurableHandle(v bool) {
	o.EnableFastDurableHandle.Set(&v)
}
// SetEnableFastDurableHandleNil sets the value for EnableFastDurableHandle to be an explicit nil
func (o *SmbConfig) SetEnableFastDurableHandleNil() {
	o.EnableFastDurableHandle.Set(nil)
}

// UnsetEnableFastDurableHandle ensures that no value is present for EnableFastDurableHandle, not even an explicit nil
func (o *SmbConfig) UnsetEnableFastDurableHandle() {
	o.EnableFastDurableHandle.Unset()
}

// GetEnableSmbAccessBasedEnumeration returns the EnableSmbAccessBasedEnumeration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmbConfig) GetEnableSmbAccessBasedEnumeration() bool {
	if o == nil || IsNil(o.EnableSmbAccessBasedEnumeration.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableSmbAccessBasedEnumeration.Get()
}

// GetEnableSmbAccessBasedEnumerationOk returns a tuple with the EnableSmbAccessBasedEnumeration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmbConfig) GetEnableSmbAccessBasedEnumerationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableSmbAccessBasedEnumeration.Get(), o.EnableSmbAccessBasedEnumeration.IsSet()
}

// HasEnableSmbAccessBasedEnumeration returns a boolean if a field has been set.
func (o *SmbConfig) HasEnableSmbAccessBasedEnumeration() bool {
	if o != nil && o.EnableSmbAccessBasedEnumeration.IsSet() {
		return true
	}

	return false
}

// SetEnableSmbAccessBasedEnumeration gets a reference to the given NullableBool and assigns it to the EnableSmbAccessBasedEnumeration field.
func (o *SmbConfig) SetEnableSmbAccessBasedEnumeration(v bool) {
	o.EnableSmbAccessBasedEnumeration.Set(&v)
}
// SetEnableSmbAccessBasedEnumerationNil sets the value for EnableSmbAccessBasedEnumeration to be an explicit nil
func (o *SmbConfig) SetEnableSmbAccessBasedEnumerationNil() {
	o.EnableSmbAccessBasedEnumeration.Set(nil)
}

// UnsetEnableSmbAccessBasedEnumeration ensures that no value is present for EnableSmbAccessBasedEnumeration, not even an explicit nil
func (o *SmbConfig) UnsetEnableSmbAccessBasedEnumeration() {
	o.EnableSmbAccessBasedEnumeration.Unset()
}

// GetEnableSmbEncryption returns the EnableSmbEncryption field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmbConfig) GetEnableSmbEncryption() bool {
	if o == nil || IsNil(o.EnableSmbEncryption.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableSmbEncryption.Get()
}

// GetEnableSmbEncryptionOk returns a tuple with the EnableSmbEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmbConfig) GetEnableSmbEncryptionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableSmbEncryption.Get(), o.EnableSmbEncryption.IsSet()
}

// HasEnableSmbEncryption returns a boolean if a field has been set.
func (o *SmbConfig) HasEnableSmbEncryption() bool {
	if o != nil && o.EnableSmbEncryption.IsSet() {
		return true
	}

	return false
}

// SetEnableSmbEncryption gets a reference to the given NullableBool and assigns it to the EnableSmbEncryption field.
func (o *SmbConfig) SetEnableSmbEncryption(v bool) {
	o.EnableSmbEncryption.Set(&v)
}
// SetEnableSmbEncryptionNil sets the value for EnableSmbEncryption to be an explicit nil
func (o *SmbConfig) SetEnableSmbEncryptionNil() {
	o.EnableSmbEncryption.Set(nil)
}

// UnsetEnableSmbEncryption ensures that no value is present for EnableSmbEncryption, not even an explicit nil
func (o *SmbConfig) UnsetEnableSmbEncryption() {
	o.EnableSmbEncryption.Unset()
}

// GetEnableSmbOplock returns the EnableSmbOplock field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmbConfig) GetEnableSmbOplock() bool {
	if o == nil || IsNil(o.EnableSmbOplock.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableSmbOplock.Get()
}

// GetEnableSmbOplockOk returns a tuple with the EnableSmbOplock field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmbConfig) GetEnableSmbOplockOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableSmbOplock.Get(), o.EnableSmbOplock.IsSet()
}

// HasEnableSmbOplock returns a boolean if a field has been set.
func (o *SmbConfig) HasEnableSmbOplock() bool {
	if o != nil && o.EnableSmbOplock.IsSet() {
		return true
	}

	return false
}

// SetEnableSmbOplock gets a reference to the given NullableBool and assigns it to the EnableSmbOplock field.
func (o *SmbConfig) SetEnableSmbOplock(v bool) {
	o.EnableSmbOplock.Set(&v)
}
// SetEnableSmbOplockNil sets the value for EnableSmbOplock to be an explicit nil
func (o *SmbConfig) SetEnableSmbOplockNil() {
	o.EnableSmbOplock.Set(nil)
}

// UnsetEnableSmbOplock ensures that no value is present for EnableSmbOplock, not even an explicit nil
func (o *SmbConfig) UnsetEnableSmbOplock() {
	o.EnableSmbOplock.Unset()
}

// GetEnableSmbViewDiscovery returns the EnableSmbViewDiscovery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmbConfig) GetEnableSmbViewDiscovery() bool {
	if o == nil || IsNil(o.EnableSmbViewDiscovery.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableSmbViewDiscovery.Get()
}

// GetEnableSmbViewDiscoveryOk returns a tuple with the EnableSmbViewDiscovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmbConfig) GetEnableSmbViewDiscoveryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableSmbViewDiscovery.Get(), o.EnableSmbViewDiscovery.IsSet()
}

// HasEnableSmbViewDiscovery returns a boolean if a field has been set.
func (o *SmbConfig) HasEnableSmbViewDiscovery() bool {
	if o != nil && o.EnableSmbViewDiscovery.IsSet() {
		return true
	}

	return false
}

// SetEnableSmbViewDiscovery gets a reference to the given NullableBool and assigns it to the EnableSmbViewDiscovery field.
func (o *SmbConfig) SetEnableSmbViewDiscovery(v bool) {
	o.EnableSmbViewDiscovery.Set(&v)
}
// SetEnableSmbViewDiscoveryNil sets the value for EnableSmbViewDiscovery to be an explicit nil
func (o *SmbConfig) SetEnableSmbViewDiscoveryNil() {
	o.EnableSmbViewDiscovery.Set(nil)
}

// UnsetEnableSmbViewDiscovery ensures that no value is present for EnableSmbViewDiscovery, not even an explicit nil
func (o *SmbConfig) UnsetEnableSmbViewDiscovery() {
	o.EnableSmbViewDiscovery.Unset()
}

// GetEnforceSmbEncryption returns the EnforceSmbEncryption field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmbConfig) GetEnforceSmbEncryption() bool {
	if o == nil || IsNil(o.EnforceSmbEncryption.Get()) {
		var ret bool
		return ret
	}
	return *o.EnforceSmbEncryption.Get()
}

// GetEnforceSmbEncryptionOk returns a tuple with the EnforceSmbEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmbConfig) GetEnforceSmbEncryptionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnforceSmbEncryption.Get(), o.EnforceSmbEncryption.IsSet()
}

// HasEnforceSmbEncryption returns a boolean if a field has been set.
func (o *SmbConfig) HasEnforceSmbEncryption() bool {
	if o != nil && o.EnforceSmbEncryption.IsSet() {
		return true
	}

	return false
}

// SetEnforceSmbEncryption gets a reference to the given NullableBool and assigns it to the EnforceSmbEncryption field.
func (o *SmbConfig) SetEnforceSmbEncryption(v bool) {
	o.EnforceSmbEncryption.Set(&v)
}
// SetEnforceSmbEncryptionNil sets the value for EnforceSmbEncryption to be an explicit nil
func (o *SmbConfig) SetEnforceSmbEncryptionNil() {
	o.EnforceSmbEncryption.Set(nil)
}

// UnsetEnforceSmbEncryption ensures that no value is present for EnforceSmbEncryption, not even an explicit nil
func (o *SmbConfig) UnsetEnforceSmbEncryption() {
	o.EnforceSmbEncryption.Unset()
}

// GetSharePermissions returns the SharePermissions field value if set, zero value otherwise.
func (o *SmbConfig) GetSharePermissions() SmbConfigSharePermissions {
	if o == nil || IsNil(o.SharePermissions) {
		var ret SmbConfigSharePermissions
		return ret
	}
	return *o.SharePermissions
}

// GetSharePermissionsOk returns a tuple with the SharePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmbConfig) GetSharePermissionsOk() (*SmbConfigSharePermissions, bool) {
	if o == nil || IsNil(o.SharePermissions) {
		return nil, false
	}
	return o.SharePermissions, true
}

// HasSharePermissions returns a boolean if a field has been set.
func (o *SmbConfig) HasSharePermissions() bool {
	if o != nil && !IsNil(o.SharePermissions) {
		return true
	}

	return false
}

// SetSharePermissions gets a reference to the given SmbConfigSharePermissions and assigns it to the SharePermissions field.
func (o *SmbConfig) SetSharePermissions(v SmbConfigSharePermissions) {
	o.SharePermissions = &v
}

// GetSmbPermissionsInfo returns the SmbPermissionsInfo field value if set, zero value otherwise.
func (o *SmbConfig) GetSmbPermissionsInfo() SmbConfigSmbPermissionsInfo {
	if o == nil || IsNil(o.SmbPermissionsInfo) {
		var ret SmbConfigSmbPermissionsInfo
		return ret
	}
	return *o.SmbPermissionsInfo
}

// GetSmbPermissionsInfoOk returns a tuple with the SmbPermissionsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmbConfig) GetSmbPermissionsInfoOk() (*SmbConfigSmbPermissionsInfo, bool) {
	if o == nil || IsNil(o.SmbPermissionsInfo) {
		return nil, false
	}
	return o.SmbPermissionsInfo, true
}

// HasSmbPermissionsInfo returns a boolean if a field has been set.
func (o *SmbConfig) HasSmbPermissionsInfo() bool {
	if o != nil && !IsNil(o.SmbPermissionsInfo) {
		return true
	}

	return false
}

// SetSmbPermissionsInfo gets a reference to the given SmbConfigSmbPermissionsInfo and assigns it to the SmbPermissionsInfo field.
func (o *SmbConfig) SetSmbPermissionsInfo(v SmbConfigSmbPermissionsInfo) {
	o.SmbPermissionsInfo = &v
}

func (o SmbConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmbConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.EnableFastDurableHandle.IsSet() {
		toSerialize["enableFastDurableHandle"] = o.EnableFastDurableHandle.Get()
	}
	if o.EnableSmbAccessBasedEnumeration.IsSet() {
		toSerialize["enableSmbAccessBasedEnumeration"] = o.EnableSmbAccessBasedEnumeration.Get()
	}
	if o.EnableSmbEncryption.IsSet() {
		toSerialize["enableSmbEncryption"] = o.EnableSmbEncryption.Get()
	}
	if o.EnableSmbOplock.IsSet() {
		toSerialize["enableSmbOplock"] = o.EnableSmbOplock.Get()
	}
	if o.EnableSmbViewDiscovery.IsSet() {
		toSerialize["enableSmbViewDiscovery"] = o.EnableSmbViewDiscovery.Get()
	}
	if o.EnforceSmbEncryption.IsSet() {
		toSerialize["enforceSmbEncryption"] = o.EnforceSmbEncryption.Get()
	}
	if !IsNil(o.SharePermissions) {
		toSerialize["sharePermissions"] = o.SharePermissions
	}
	if !IsNil(o.SmbPermissionsInfo) {
		toSerialize["smbPermissionsInfo"] = o.SmbPermissionsInfo
	}
	return toSerialize, nil
}

type NullableSmbConfig struct {
	value *SmbConfig
	isSet bool
}

func (v NullableSmbConfig) Get() *SmbConfig {
	return v.value
}

func (v *NullableSmbConfig) Set(val *SmbConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSmbConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSmbConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmbConfig(val *SmbConfig) *NullableSmbConfig {
	return &NullableSmbConfig{value: val, isSet: true}
}

func (v NullableSmbConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmbConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


