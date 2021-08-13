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

// HeliosMfa Specifies whether to enable/disable MFA for account.
type HeliosMfa struct {
	// Specifies whether MFA needs to be enabled / disabled.
	Mfa NullableBool `json:"mfa,omitempty"`
}

// NewHeliosMfa instantiates a new HeliosMfa object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeliosMfa() *HeliosMfa {
	this := HeliosMfa{}
	return &this
}

// NewHeliosMfaWithDefaults instantiates a new HeliosMfa object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeliosMfaWithDefaults() *HeliosMfa {
	this := HeliosMfa{}
	return &this
}

// GetMfa returns the Mfa field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosMfa) GetMfa() bool {
	if o == nil || o.Mfa.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Mfa.Get()
}

// GetMfaOk returns a tuple with the Mfa field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosMfa) GetMfaOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Mfa.Get(), o.Mfa.IsSet()
}

// HasMfa returns a boolean if a field has been set.
func (o *HeliosMfa) HasMfa() bool {
	if o != nil && o.Mfa.IsSet() {
		return true
	}

	return false
}

// SetMfa gets a reference to the given NullableBool and assigns it to the Mfa field.
func (o *HeliosMfa) SetMfa(v bool) {
	o.Mfa.Set(&v)
}
// SetMfaNil sets the value for Mfa to be an explicit nil
func (o *HeliosMfa) SetMfaNil() {
	o.Mfa.Set(nil)
}

// UnsetMfa ensures that no value is present for Mfa, not even an explicit nil
func (o *HeliosMfa) UnsetMfa() {
	o.Mfa.Unset()
}

func (o HeliosMfa) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Mfa.IsSet() {
		toSerialize["mfa"] = o.Mfa.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableHeliosMfa struct {
	value *HeliosMfa
	isSet bool
}

func (v NullableHeliosMfa) Get() *HeliosMfa {
	return v.value
}

func (v *NullableHeliosMfa) Set(val *HeliosMfa) {
	v.value = val
	v.isSet = true
}

func (v NullableHeliosMfa) IsSet() bool {
	return v.isSet
}

func (v *NullableHeliosMfa) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeliosMfa(val *HeliosMfa) *NullableHeliosMfa {
	return &NullableHeliosMfa{value: val, isSet: true}
}

func (v NullableHeliosMfa) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeliosMfa) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


