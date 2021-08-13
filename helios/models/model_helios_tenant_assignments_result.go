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

// HeliosTenantAssignmentsResult All properties like entities, sources, policies etc assigned to the tenant.
type HeliosTenantAssignmentsResult struct {
	// The tenant id.
	TenantId NullableString `json:"tenantId,omitempty"`
	Assignments NullableHeliosBaseTenantAssignmentsResult `json:"assignments,omitempty"`
}

// NewHeliosTenantAssignmentsResult instantiates a new HeliosTenantAssignmentsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeliosTenantAssignmentsResult() *HeliosTenantAssignmentsResult {
	this := HeliosTenantAssignmentsResult{}
	return &this
}

// NewHeliosTenantAssignmentsResultWithDefaults instantiates a new HeliosTenantAssignmentsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeliosTenantAssignmentsResultWithDefaults() *HeliosTenantAssignmentsResult {
	this := HeliosTenantAssignmentsResult{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosTenantAssignmentsResult) GetTenantId() string {
	if o == nil || o.TenantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosTenantAssignmentsResult) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *HeliosTenantAssignmentsResult) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *HeliosTenantAssignmentsResult) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *HeliosTenantAssignmentsResult) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *HeliosTenantAssignmentsResult) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetAssignments returns the Assignments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosTenantAssignmentsResult) GetAssignments() HeliosBaseTenantAssignmentsResult {
	if o == nil || o.Assignments.Get() == nil {
		var ret HeliosBaseTenantAssignmentsResult
		return ret
	}
	return *o.Assignments.Get()
}

// GetAssignmentsOk returns a tuple with the Assignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosTenantAssignmentsResult) GetAssignmentsOk() (*HeliosBaseTenantAssignmentsResult, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Assignments.Get(), o.Assignments.IsSet()
}

// HasAssignments returns a boolean if a field has been set.
func (o *HeliosTenantAssignmentsResult) HasAssignments() bool {
	if o != nil && o.Assignments.IsSet() {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given NullableHeliosBaseTenantAssignmentsResult and assigns it to the Assignments field.
func (o *HeliosTenantAssignmentsResult) SetAssignments(v HeliosBaseTenantAssignmentsResult) {
	o.Assignments.Set(&v)
}
// SetAssignmentsNil sets the value for Assignments to be an explicit nil
func (o *HeliosTenantAssignmentsResult) SetAssignmentsNil() {
	o.Assignments.Set(nil)
}

// UnsetAssignments ensures that no value is present for Assignments, not even an explicit nil
func (o *HeliosTenantAssignmentsResult) UnsetAssignments() {
	o.Assignments.Unset()
}

func (o HeliosTenantAssignmentsResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.Assignments.IsSet() {
		toSerialize["assignments"] = o.Assignments.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableHeliosTenantAssignmentsResult struct {
	value *HeliosTenantAssignmentsResult
	isSet bool
}

func (v NullableHeliosTenantAssignmentsResult) Get() *HeliosTenantAssignmentsResult {
	return v.value
}

func (v *NullableHeliosTenantAssignmentsResult) Set(val *HeliosTenantAssignmentsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableHeliosTenantAssignmentsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableHeliosTenantAssignmentsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeliosTenantAssignmentsResult(val *HeliosTenantAssignmentsResult) *NullableHeliosTenantAssignmentsResult {
	return &NullableHeliosTenantAssignmentsResult{value: val, isSet: true}
}

func (v NullableHeliosTenantAssignmentsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeliosTenantAssignmentsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


