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

// HeliosReplicationConfigAllOf struct for HeliosReplicationConfigAllOf
type HeliosReplicationConfigAllOf struct {
	// Specifies the type of target to which replication need to be performed.
	TargetType NullableString `json:"targetType,omitempty"`
}

// NewHeliosReplicationConfigAllOf instantiates a new HeliosReplicationConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeliosReplicationConfigAllOf() *HeliosReplicationConfigAllOf {
	this := HeliosReplicationConfigAllOf{}
	return &this
}

// NewHeliosReplicationConfigAllOfWithDefaults instantiates a new HeliosReplicationConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeliosReplicationConfigAllOfWithDefaults() *HeliosReplicationConfigAllOf {
	this := HeliosReplicationConfigAllOf{}
	return &this
}

// GetTargetType returns the TargetType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosReplicationConfigAllOf) GetTargetType() string {
	if o == nil || o.TargetType.Get() == nil {
		var ret string
		return ret
	}
	return *o.TargetType.Get()
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosReplicationConfigAllOf) GetTargetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TargetType.Get(), o.TargetType.IsSet()
}

// HasTargetType returns a boolean if a field has been set.
func (o *HeliosReplicationConfigAllOf) HasTargetType() bool {
	if o != nil && o.TargetType.IsSet() {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given NullableString and assigns it to the TargetType field.
func (o *HeliosReplicationConfigAllOf) SetTargetType(v string) {
	o.TargetType.Set(&v)
}
// SetTargetTypeNil sets the value for TargetType to be an explicit nil
func (o *HeliosReplicationConfigAllOf) SetTargetTypeNil() {
	o.TargetType.Set(nil)
}

// UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
func (o *HeliosReplicationConfigAllOf) UnsetTargetType() {
	o.TargetType.Unset()
}

func (o HeliosReplicationConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TargetType.IsSet() {
		toSerialize["targetType"] = o.TargetType.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableHeliosReplicationConfigAllOf struct {
	value *HeliosReplicationConfigAllOf
	isSet bool
}

func (v NullableHeliosReplicationConfigAllOf) Get() *HeliosReplicationConfigAllOf {
	return v.value
}

func (v *NullableHeliosReplicationConfigAllOf) Set(val *HeliosReplicationConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHeliosReplicationConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHeliosReplicationConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeliosReplicationConfigAllOf(val *HeliosReplicationConfigAllOf) *NullableHeliosReplicationConfigAllOf {
	return &NullableHeliosReplicationConfigAllOf{value: val, isSet: true}
}

func (v NullableHeliosReplicationConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeliosReplicationConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


