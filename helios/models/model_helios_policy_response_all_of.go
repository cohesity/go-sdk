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

// HeliosPolicyResponseAllOf struct for HeliosPolicyResponseAllOf
type HeliosPolicyResponseAllOf struct {
	// Specifies a unique policy id assigned by the Helios.
	Id NullableString `json:"id,omitempty"`
	// In case of global policy response, specifies the number of policies linked to this global policy on the cluster.
	NumLinkedPolicies NullableInt64 `json:"numLinkedPolicies,omitempty"`
	// Specifies the number of object protections using the protection policy.
	NumObjectProtections NullableInt64 `json:"numObjectProtections,omitempty"`
}

// NewHeliosPolicyResponseAllOf instantiates a new HeliosPolicyResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeliosPolicyResponseAllOf() *HeliosPolicyResponseAllOf {
	this := HeliosPolicyResponseAllOf{}
	return &this
}

// NewHeliosPolicyResponseAllOfWithDefaults instantiates a new HeliosPolicyResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeliosPolicyResponseAllOfWithDefaults() *HeliosPolicyResponseAllOf {
	this := HeliosPolicyResponseAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosPolicyResponseAllOf) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosPolicyResponseAllOf) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *HeliosPolicyResponseAllOf) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *HeliosPolicyResponseAllOf) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *HeliosPolicyResponseAllOf) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *HeliosPolicyResponseAllOf) UnsetId() {
	o.Id.Unset()
}

// GetNumLinkedPolicies returns the NumLinkedPolicies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosPolicyResponseAllOf) GetNumLinkedPolicies() int64 {
	if o == nil || o.NumLinkedPolicies.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NumLinkedPolicies.Get()
}

// GetNumLinkedPoliciesOk returns a tuple with the NumLinkedPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosPolicyResponseAllOf) GetNumLinkedPoliciesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NumLinkedPolicies.Get(), o.NumLinkedPolicies.IsSet()
}

// HasNumLinkedPolicies returns a boolean if a field has been set.
func (o *HeliosPolicyResponseAllOf) HasNumLinkedPolicies() bool {
	if o != nil && o.NumLinkedPolicies.IsSet() {
		return true
	}

	return false
}

// SetNumLinkedPolicies gets a reference to the given NullableInt64 and assigns it to the NumLinkedPolicies field.
func (o *HeliosPolicyResponseAllOf) SetNumLinkedPolicies(v int64) {
	o.NumLinkedPolicies.Set(&v)
}
// SetNumLinkedPoliciesNil sets the value for NumLinkedPolicies to be an explicit nil
func (o *HeliosPolicyResponseAllOf) SetNumLinkedPoliciesNil() {
	o.NumLinkedPolicies.Set(nil)
}

// UnsetNumLinkedPolicies ensures that no value is present for NumLinkedPolicies, not even an explicit nil
func (o *HeliosPolicyResponseAllOf) UnsetNumLinkedPolicies() {
	o.NumLinkedPolicies.Unset()
}

// GetNumObjectProtections returns the NumObjectProtections field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosPolicyResponseAllOf) GetNumObjectProtections() int64 {
	if o == nil || o.NumObjectProtections.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NumObjectProtections.Get()
}

// GetNumObjectProtectionsOk returns a tuple with the NumObjectProtections field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosPolicyResponseAllOf) GetNumObjectProtectionsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NumObjectProtections.Get(), o.NumObjectProtections.IsSet()
}

// HasNumObjectProtections returns a boolean if a field has been set.
func (o *HeliosPolicyResponseAllOf) HasNumObjectProtections() bool {
	if o != nil && o.NumObjectProtections.IsSet() {
		return true
	}

	return false
}

// SetNumObjectProtections gets a reference to the given NullableInt64 and assigns it to the NumObjectProtections field.
func (o *HeliosPolicyResponseAllOf) SetNumObjectProtections(v int64) {
	o.NumObjectProtections.Set(&v)
}
// SetNumObjectProtectionsNil sets the value for NumObjectProtections to be an explicit nil
func (o *HeliosPolicyResponseAllOf) SetNumObjectProtectionsNil() {
	o.NumObjectProtections.Set(nil)
}

// UnsetNumObjectProtections ensures that no value is present for NumObjectProtections, not even an explicit nil
func (o *HeliosPolicyResponseAllOf) UnsetNumObjectProtections() {
	o.NumObjectProtections.Unset()
}

func (o HeliosPolicyResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.NumLinkedPolicies.IsSet() {
		toSerialize["numLinkedPolicies"] = o.NumLinkedPolicies.Get()
	}
	if o.NumObjectProtections.IsSet() {
		toSerialize["numObjectProtections"] = o.NumObjectProtections.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableHeliosPolicyResponseAllOf struct {
	value *HeliosPolicyResponseAllOf
	isSet bool
}

func (v NullableHeliosPolicyResponseAllOf) Get() *HeliosPolicyResponseAllOf {
	return v.value
}

func (v *NullableHeliosPolicyResponseAllOf) Set(val *HeliosPolicyResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHeliosPolicyResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHeliosPolicyResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeliosPolicyResponseAllOf(val *HeliosPolicyResponseAllOf) *NullableHeliosPolicyResponseAllOf {
	return &NullableHeliosPolicyResponseAllOf{value: val, isSet: true}
}

func (v NullableHeliosPolicyResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeliosPolicyResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


