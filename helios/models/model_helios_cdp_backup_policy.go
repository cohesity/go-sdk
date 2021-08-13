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

// HeliosCdpBackupPolicy Specifies CDP (Continious Data Protection) backup settings for a Protection Group.
type HeliosCdpBackupPolicy struct {
	Retention *HeliosCdpRetention `json:"retention,omitempty"`
}

// NewHeliosCdpBackupPolicy instantiates a new HeliosCdpBackupPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeliosCdpBackupPolicy() *HeliosCdpBackupPolicy {
	this := HeliosCdpBackupPolicy{}
	return &this
}

// NewHeliosCdpBackupPolicyWithDefaults instantiates a new HeliosCdpBackupPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeliosCdpBackupPolicyWithDefaults() *HeliosCdpBackupPolicy {
	this := HeliosCdpBackupPolicy{}
	return &this
}

// GetRetention returns the Retention field value if set, zero value otherwise.
func (o *HeliosCdpBackupPolicy) GetRetention() HeliosCdpRetention {
	if o == nil || o.Retention == nil {
		var ret HeliosCdpRetention
		return ret
	}
	return *o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeliosCdpBackupPolicy) GetRetentionOk() (*HeliosCdpRetention, bool) {
	if o == nil || o.Retention == nil {
		return nil, false
	}
	return o.Retention, true
}

// HasRetention returns a boolean if a field has been set.
func (o *HeliosCdpBackupPolicy) HasRetention() bool {
	if o != nil && o.Retention != nil {
		return true
	}

	return false
}

// SetRetention gets a reference to the given HeliosCdpRetention and assigns it to the Retention field.
func (o *HeliosCdpBackupPolicy) SetRetention(v HeliosCdpRetention) {
	o.Retention = &v
}

func (o HeliosCdpBackupPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Retention != nil {
		toSerialize["retention"] = o.Retention
	}
	return json.Marshal(toSerialize)
}

type NullableHeliosCdpBackupPolicy struct {
	value *HeliosCdpBackupPolicy
	isSet bool
}

func (v NullableHeliosCdpBackupPolicy) Get() *HeliosCdpBackupPolicy {
	return v.value
}

func (v *NullableHeliosCdpBackupPolicy) Set(val *HeliosCdpBackupPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableHeliosCdpBackupPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableHeliosCdpBackupPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeliosCdpBackupPolicy(val *HeliosCdpBackupPolicy) *NullableHeliosCdpBackupPolicy {
	return &NullableHeliosCdpBackupPolicy{value: val, isSet: true}
}

func (v NullableHeliosCdpBackupPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeliosCdpBackupPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


