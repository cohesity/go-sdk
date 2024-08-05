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

// checks if the MfaConfigInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MfaConfigInfo{}

// MfaConfigInfo Holds the MFA configuration to be returned or stored.
type MfaConfigInfo struct {
	// Specifies the list of mechanism to receive the OTP code.
	AuthenticationTypes []string `json:"authenticationTypes,omitempty"`
	// Specifies whether MFA is enabled on a cluster level.
	Enabled *bool `json:"enabled,omitempty"`
	// Specifies whether user MFA setting needs to be retained.
	RetainUserMfaSettings NullableBool `json:"retainUserMfaSettings,omitempty"`
}

// NewMfaConfigInfo instantiates a new MfaConfigInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMfaConfigInfo() *MfaConfigInfo {
	this := MfaConfigInfo{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewMfaConfigInfoWithDefaults instantiates a new MfaConfigInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMfaConfigInfoWithDefaults() *MfaConfigInfo {
	this := MfaConfigInfo{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetAuthenticationTypes returns the AuthenticationTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MfaConfigInfo) GetAuthenticationTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AuthenticationTypes
}

// GetAuthenticationTypesOk returns a tuple with the AuthenticationTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MfaConfigInfo) GetAuthenticationTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.AuthenticationTypes) {
		return nil, false
	}
	return o.AuthenticationTypes, true
}

// HasAuthenticationTypes returns a boolean if a field has been set.
func (o *MfaConfigInfo) HasAuthenticationTypes() bool {
	if o != nil && !IsNil(o.AuthenticationTypes) {
		return true
	}

	return false
}

// SetAuthenticationTypes gets a reference to the given []string and assigns it to the AuthenticationTypes field.
func (o *MfaConfigInfo) SetAuthenticationTypes(v []string) {
	o.AuthenticationTypes = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *MfaConfigInfo) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MfaConfigInfo) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *MfaConfigInfo) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *MfaConfigInfo) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRetainUserMfaSettings returns the RetainUserMfaSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MfaConfigInfo) GetRetainUserMfaSettings() bool {
	if o == nil || IsNil(o.RetainUserMfaSettings.Get()) {
		var ret bool
		return ret
	}
	return *o.RetainUserMfaSettings.Get()
}

// GetRetainUserMfaSettingsOk returns a tuple with the RetainUserMfaSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MfaConfigInfo) GetRetainUserMfaSettingsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetainUserMfaSettings.Get(), o.RetainUserMfaSettings.IsSet()
}

// HasRetainUserMfaSettings returns a boolean if a field has been set.
func (o *MfaConfigInfo) HasRetainUserMfaSettings() bool {
	if o != nil && o.RetainUserMfaSettings.IsSet() {
		return true
	}

	return false
}

// SetRetainUserMfaSettings gets a reference to the given NullableBool and assigns it to the RetainUserMfaSettings field.
func (o *MfaConfigInfo) SetRetainUserMfaSettings(v bool) {
	o.RetainUserMfaSettings.Set(&v)
}
// SetRetainUserMfaSettingsNil sets the value for RetainUserMfaSettings to be an explicit nil
func (o *MfaConfigInfo) SetRetainUserMfaSettingsNil() {
	o.RetainUserMfaSettings.Set(nil)
}

// UnsetRetainUserMfaSettings ensures that no value is present for RetainUserMfaSettings, not even an explicit nil
func (o *MfaConfigInfo) UnsetRetainUserMfaSettings() {
	o.RetainUserMfaSettings.Unset()
}

func (o MfaConfigInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MfaConfigInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthenticationTypes != nil {
		toSerialize["authenticationTypes"] = o.AuthenticationTypes
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if o.RetainUserMfaSettings.IsSet() {
		toSerialize["retainUserMfaSettings"] = o.RetainUserMfaSettings.Get()
	}
	return toSerialize, nil
}

type NullableMfaConfigInfo struct {
	value *MfaConfigInfo
	isSet bool
}

func (v NullableMfaConfigInfo) Get() *MfaConfigInfo {
	return v.value
}

func (v *NullableMfaConfigInfo) Set(val *MfaConfigInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMfaConfigInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMfaConfigInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMfaConfigInfo(val *MfaConfigInfo) *NullableMfaConfigInfo {
	return &NullableMfaConfigInfo{value: val, isSet: true}
}

func (v NullableMfaConfigInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMfaConfigInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


