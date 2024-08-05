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

// checks if the SecurityConfigPasswordReuse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityConfigPasswordReuse{}

// SecurityConfigPasswordReuse Specifies security config for password reuse.
type SecurityConfigPasswordReuse struct {
	// Specifies the number of characters in the new password that needs to be different from the old password (only applicable when changing the user's own password).
	NumDifferentChars NullableInt32 `json:"numDifferentChars,omitempty"`
	// Specifies the minimum number of old passwords that are not allowed for changing the password.
	NumDisallowedOldPasswords NullableInt32 `json:"numDisallowedOldPasswords,omitempty"`
}

// NewSecurityConfigPasswordReuse instantiates a new SecurityConfigPasswordReuse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityConfigPasswordReuse() *SecurityConfigPasswordReuse {
	this := SecurityConfigPasswordReuse{}
	return &this
}

// NewSecurityConfigPasswordReuseWithDefaults instantiates a new SecurityConfigPasswordReuse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityConfigPasswordReuseWithDefaults() *SecurityConfigPasswordReuse {
	this := SecurityConfigPasswordReuse{}
	return &this
}

// GetNumDifferentChars returns the NumDifferentChars field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityConfigPasswordReuse) GetNumDifferentChars() int32 {
	if o == nil || IsNil(o.NumDifferentChars.Get()) {
		var ret int32
		return ret
	}
	return *o.NumDifferentChars.Get()
}

// GetNumDifferentCharsOk returns a tuple with the NumDifferentChars field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityConfigPasswordReuse) GetNumDifferentCharsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumDifferentChars.Get(), o.NumDifferentChars.IsSet()
}

// HasNumDifferentChars returns a boolean if a field has been set.
func (o *SecurityConfigPasswordReuse) HasNumDifferentChars() bool {
	if o != nil && o.NumDifferentChars.IsSet() {
		return true
	}

	return false
}

// SetNumDifferentChars gets a reference to the given NullableInt32 and assigns it to the NumDifferentChars field.
func (o *SecurityConfigPasswordReuse) SetNumDifferentChars(v int32) {
	o.NumDifferentChars.Set(&v)
}
// SetNumDifferentCharsNil sets the value for NumDifferentChars to be an explicit nil
func (o *SecurityConfigPasswordReuse) SetNumDifferentCharsNil() {
	o.NumDifferentChars.Set(nil)
}

// UnsetNumDifferentChars ensures that no value is present for NumDifferentChars, not even an explicit nil
func (o *SecurityConfigPasswordReuse) UnsetNumDifferentChars() {
	o.NumDifferentChars.Unset()
}

// GetNumDisallowedOldPasswords returns the NumDisallowedOldPasswords field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityConfigPasswordReuse) GetNumDisallowedOldPasswords() int32 {
	if o == nil || IsNil(o.NumDisallowedOldPasswords.Get()) {
		var ret int32
		return ret
	}
	return *o.NumDisallowedOldPasswords.Get()
}

// GetNumDisallowedOldPasswordsOk returns a tuple with the NumDisallowedOldPasswords field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityConfigPasswordReuse) GetNumDisallowedOldPasswordsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumDisallowedOldPasswords.Get(), o.NumDisallowedOldPasswords.IsSet()
}

// HasNumDisallowedOldPasswords returns a boolean if a field has been set.
func (o *SecurityConfigPasswordReuse) HasNumDisallowedOldPasswords() bool {
	if o != nil && o.NumDisallowedOldPasswords.IsSet() {
		return true
	}

	return false
}

// SetNumDisallowedOldPasswords gets a reference to the given NullableInt32 and assigns it to the NumDisallowedOldPasswords field.
func (o *SecurityConfigPasswordReuse) SetNumDisallowedOldPasswords(v int32) {
	o.NumDisallowedOldPasswords.Set(&v)
}
// SetNumDisallowedOldPasswordsNil sets the value for NumDisallowedOldPasswords to be an explicit nil
func (o *SecurityConfigPasswordReuse) SetNumDisallowedOldPasswordsNil() {
	o.NumDisallowedOldPasswords.Set(nil)
}

// UnsetNumDisallowedOldPasswords ensures that no value is present for NumDisallowedOldPasswords, not even an explicit nil
func (o *SecurityConfigPasswordReuse) UnsetNumDisallowedOldPasswords() {
	o.NumDisallowedOldPasswords.Unset()
}

func (o SecurityConfigPasswordReuse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityConfigPasswordReuse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.NumDifferentChars.IsSet() {
		toSerialize["numDifferentChars"] = o.NumDifferentChars.Get()
	}
	if o.NumDisallowedOldPasswords.IsSet() {
		toSerialize["numDisallowedOldPasswords"] = o.NumDisallowedOldPasswords.Get()
	}
	return toSerialize, nil
}

type NullableSecurityConfigPasswordReuse struct {
	value *SecurityConfigPasswordReuse
	isSet bool
}

func (v NullableSecurityConfigPasswordReuse) Get() *SecurityConfigPasswordReuse {
	return v.value
}

func (v *NullableSecurityConfigPasswordReuse) Set(val *SecurityConfigPasswordReuse) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityConfigPasswordReuse) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityConfigPasswordReuse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityConfigPasswordReuse(val *SecurityConfigPasswordReuse) *NullableSecurityConfigPasswordReuse {
	return &NullableSecurityConfigPasswordReuse{value: val, isSet: true}
}

func (v NullableSecurityConfigPasswordReuse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityConfigPasswordReuse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


