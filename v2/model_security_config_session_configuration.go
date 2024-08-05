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

// checks if the SecurityConfigSessionConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityConfigSessionConfiguration{}

// SecurityConfigSessionConfiguration Specifies configuration for user sessions.
type SecurityConfigSessionConfiguration struct {
	// Specifies absolute session expiration time in seconds.
	AbsoluteTimeout NullableInt64 `json:"absoluteTimeout,omitempty"`
	// Specifies inactivity session expiration time in seconds.
	InactivityTimeout NullableInt64 `json:"inactivityTimeout,omitempty"`
	// Specifies if limitations on number of active sessions is enabled or not.
	LimitSessions NullableBool `json:"limitSessions,omitempty"`
	// Specifies maximum number of active sessions allowed per user. This applies only when limitSessions is enabled.
	SessionLimitPerUser NullableInt64 `json:"sessionLimitPerUser,omitempty"`
	// Specifies maximum number of active sessions allowed system wide. This applies only when limitSessions is enabled.
	SessionLimitSystemWide NullableInt64 `json:"sessionLimitSystemWide,omitempty"`
}

// NewSecurityConfigSessionConfiguration instantiates a new SecurityConfigSessionConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityConfigSessionConfiguration() *SecurityConfigSessionConfiguration {
	this := SecurityConfigSessionConfiguration{}
	return &this
}

// NewSecurityConfigSessionConfigurationWithDefaults instantiates a new SecurityConfigSessionConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityConfigSessionConfigurationWithDefaults() *SecurityConfigSessionConfiguration {
	this := SecurityConfigSessionConfiguration{}
	return &this
}

// GetAbsoluteTimeout returns the AbsoluteTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityConfigSessionConfiguration) GetAbsoluteTimeout() int64 {
	if o == nil || IsNil(o.AbsoluteTimeout.Get()) {
		var ret int64
		return ret
	}
	return *o.AbsoluteTimeout.Get()
}

// GetAbsoluteTimeoutOk returns a tuple with the AbsoluteTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityConfigSessionConfiguration) GetAbsoluteTimeoutOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AbsoluteTimeout.Get(), o.AbsoluteTimeout.IsSet()
}

// HasAbsoluteTimeout returns a boolean if a field has been set.
func (o *SecurityConfigSessionConfiguration) HasAbsoluteTimeout() bool {
	if o != nil && o.AbsoluteTimeout.IsSet() {
		return true
	}

	return false
}

// SetAbsoluteTimeout gets a reference to the given NullableInt64 and assigns it to the AbsoluteTimeout field.
func (o *SecurityConfigSessionConfiguration) SetAbsoluteTimeout(v int64) {
	o.AbsoluteTimeout.Set(&v)
}
// SetAbsoluteTimeoutNil sets the value for AbsoluteTimeout to be an explicit nil
func (o *SecurityConfigSessionConfiguration) SetAbsoluteTimeoutNil() {
	o.AbsoluteTimeout.Set(nil)
}

// UnsetAbsoluteTimeout ensures that no value is present for AbsoluteTimeout, not even an explicit nil
func (o *SecurityConfigSessionConfiguration) UnsetAbsoluteTimeout() {
	o.AbsoluteTimeout.Unset()
}

// GetInactivityTimeout returns the InactivityTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityConfigSessionConfiguration) GetInactivityTimeout() int64 {
	if o == nil || IsNil(o.InactivityTimeout.Get()) {
		var ret int64
		return ret
	}
	return *o.InactivityTimeout.Get()
}

// GetInactivityTimeoutOk returns a tuple with the InactivityTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityConfigSessionConfiguration) GetInactivityTimeoutOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.InactivityTimeout.Get(), o.InactivityTimeout.IsSet()
}

// HasInactivityTimeout returns a boolean if a field has been set.
func (o *SecurityConfigSessionConfiguration) HasInactivityTimeout() bool {
	if o != nil && o.InactivityTimeout.IsSet() {
		return true
	}

	return false
}

// SetInactivityTimeout gets a reference to the given NullableInt64 and assigns it to the InactivityTimeout field.
func (o *SecurityConfigSessionConfiguration) SetInactivityTimeout(v int64) {
	o.InactivityTimeout.Set(&v)
}
// SetInactivityTimeoutNil sets the value for InactivityTimeout to be an explicit nil
func (o *SecurityConfigSessionConfiguration) SetInactivityTimeoutNil() {
	o.InactivityTimeout.Set(nil)
}

// UnsetInactivityTimeout ensures that no value is present for InactivityTimeout, not even an explicit nil
func (o *SecurityConfigSessionConfiguration) UnsetInactivityTimeout() {
	o.InactivityTimeout.Unset()
}

// GetLimitSessions returns the LimitSessions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityConfigSessionConfiguration) GetLimitSessions() bool {
	if o == nil || IsNil(o.LimitSessions.Get()) {
		var ret bool
		return ret
	}
	return *o.LimitSessions.Get()
}

// GetLimitSessionsOk returns a tuple with the LimitSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityConfigSessionConfiguration) GetLimitSessionsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.LimitSessions.Get(), o.LimitSessions.IsSet()
}

// HasLimitSessions returns a boolean if a field has been set.
func (o *SecurityConfigSessionConfiguration) HasLimitSessions() bool {
	if o != nil && o.LimitSessions.IsSet() {
		return true
	}

	return false
}

// SetLimitSessions gets a reference to the given NullableBool and assigns it to the LimitSessions field.
func (o *SecurityConfigSessionConfiguration) SetLimitSessions(v bool) {
	o.LimitSessions.Set(&v)
}
// SetLimitSessionsNil sets the value for LimitSessions to be an explicit nil
func (o *SecurityConfigSessionConfiguration) SetLimitSessionsNil() {
	o.LimitSessions.Set(nil)
}

// UnsetLimitSessions ensures that no value is present for LimitSessions, not even an explicit nil
func (o *SecurityConfigSessionConfiguration) UnsetLimitSessions() {
	o.LimitSessions.Unset()
}

// GetSessionLimitPerUser returns the SessionLimitPerUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityConfigSessionConfiguration) GetSessionLimitPerUser() int64 {
	if o == nil || IsNil(o.SessionLimitPerUser.Get()) {
		var ret int64
		return ret
	}
	return *o.SessionLimitPerUser.Get()
}

// GetSessionLimitPerUserOk returns a tuple with the SessionLimitPerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityConfigSessionConfiguration) GetSessionLimitPerUserOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SessionLimitPerUser.Get(), o.SessionLimitPerUser.IsSet()
}

// HasSessionLimitPerUser returns a boolean if a field has been set.
func (o *SecurityConfigSessionConfiguration) HasSessionLimitPerUser() bool {
	if o != nil && o.SessionLimitPerUser.IsSet() {
		return true
	}

	return false
}

// SetSessionLimitPerUser gets a reference to the given NullableInt64 and assigns it to the SessionLimitPerUser field.
func (o *SecurityConfigSessionConfiguration) SetSessionLimitPerUser(v int64) {
	o.SessionLimitPerUser.Set(&v)
}
// SetSessionLimitPerUserNil sets the value for SessionLimitPerUser to be an explicit nil
func (o *SecurityConfigSessionConfiguration) SetSessionLimitPerUserNil() {
	o.SessionLimitPerUser.Set(nil)
}

// UnsetSessionLimitPerUser ensures that no value is present for SessionLimitPerUser, not even an explicit nil
func (o *SecurityConfigSessionConfiguration) UnsetSessionLimitPerUser() {
	o.SessionLimitPerUser.Unset()
}

// GetSessionLimitSystemWide returns the SessionLimitSystemWide field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityConfigSessionConfiguration) GetSessionLimitSystemWide() int64 {
	if o == nil || IsNil(o.SessionLimitSystemWide.Get()) {
		var ret int64
		return ret
	}
	return *o.SessionLimitSystemWide.Get()
}

// GetSessionLimitSystemWideOk returns a tuple with the SessionLimitSystemWide field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityConfigSessionConfiguration) GetSessionLimitSystemWideOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SessionLimitSystemWide.Get(), o.SessionLimitSystemWide.IsSet()
}

// HasSessionLimitSystemWide returns a boolean if a field has been set.
func (o *SecurityConfigSessionConfiguration) HasSessionLimitSystemWide() bool {
	if o != nil && o.SessionLimitSystemWide.IsSet() {
		return true
	}

	return false
}

// SetSessionLimitSystemWide gets a reference to the given NullableInt64 and assigns it to the SessionLimitSystemWide field.
func (o *SecurityConfigSessionConfiguration) SetSessionLimitSystemWide(v int64) {
	o.SessionLimitSystemWide.Set(&v)
}
// SetSessionLimitSystemWideNil sets the value for SessionLimitSystemWide to be an explicit nil
func (o *SecurityConfigSessionConfiguration) SetSessionLimitSystemWideNil() {
	o.SessionLimitSystemWide.Set(nil)
}

// UnsetSessionLimitSystemWide ensures that no value is present for SessionLimitSystemWide, not even an explicit nil
func (o *SecurityConfigSessionConfiguration) UnsetSessionLimitSystemWide() {
	o.SessionLimitSystemWide.Unset()
}

func (o SecurityConfigSessionConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityConfigSessionConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AbsoluteTimeout.IsSet() {
		toSerialize["absoluteTimeout"] = o.AbsoluteTimeout.Get()
	}
	if o.InactivityTimeout.IsSet() {
		toSerialize["inactivityTimeout"] = o.InactivityTimeout.Get()
	}
	if o.LimitSessions.IsSet() {
		toSerialize["limitSessions"] = o.LimitSessions.Get()
	}
	if o.SessionLimitPerUser.IsSet() {
		toSerialize["sessionLimitPerUser"] = o.SessionLimitPerUser.Get()
	}
	if o.SessionLimitSystemWide.IsSet() {
		toSerialize["sessionLimitSystemWide"] = o.SessionLimitSystemWide.Get()
	}
	return toSerialize, nil
}

type NullableSecurityConfigSessionConfiguration struct {
	value *SecurityConfigSessionConfiguration
	isSet bool
}

func (v NullableSecurityConfigSessionConfiguration) Get() *SecurityConfigSessionConfiguration {
	return v.value
}

func (v *NullableSecurityConfigSessionConfiguration) Set(val *SecurityConfigSessionConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityConfigSessionConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityConfigSessionConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityConfigSessionConfiguration(val *SecurityConfigSessionConfiguration) *NullableSecurityConfigSessionConfiguration {
	return &NullableSecurityConfigSessionConfiguration{value: val, isSet: true}
}

func (v NullableSecurityConfigSessionConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityConfigSessionConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


