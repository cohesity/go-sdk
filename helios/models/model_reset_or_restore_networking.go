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

// ResetOrRestoreNetworking Update cluster or node reset state information
type ResetOrRestoreNetworking struct {
	// Cancel reset cluster or node state operation.
	Operation NullableString `json:"operation,omitempty"`
}

// NewResetOrRestoreNetworking instantiates a new ResetOrRestoreNetworking object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResetOrRestoreNetworking() *ResetOrRestoreNetworking {
	this := ResetOrRestoreNetworking{}
	return &this
}

// NewResetOrRestoreNetworkingWithDefaults instantiates a new ResetOrRestoreNetworking object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResetOrRestoreNetworkingWithDefaults() *ResetOrRestoreNetworking {
	this := ResetOrRestoreNetworking{}
	return &this
}

// GetOperation returns the Operation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResetOrRestoreNetworking) GetOperation() string {
	if o == nil || o.Operation.Get() == nil {
		var ret string
		return ret
	}
	return *o.Operation.Get()
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResetOrRestoreNetworking) GetOperationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Operation.Get(), o.Operation.IsSet()
}

// HasOperation returns a boolean if a field has been set.
func (o *ResetOrRestoreNetworking) HasOperation() bool {
	if o != nil && o.Operation.IsSet() {
		return true
	}

	return false
}

// SetOperation gets a reference to the given NullableString and assigns it to the Operation field.
func (o *ResetOrRestoreNetworking) SetOperation(v string) {
	o.Operation.Set(&v)
}
// SetOperationNil sets the value for Operation to be an explicit nil
func (o *ResetOrRestoreNetworking) SetOperationNil() {
	o.Operation.Set(nil)
}

// UnsetOperation ensures that no value is present for Operation, not even an explicit nil
func (o *ResetOrRestoreNetworking) UnsetOperation() {
	o.Operation.Unset()
}

func (o ResetOrRestoreNetworking) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Operation.IsSet() {
		toSerialize["operation"] = o.Operation.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableResetOrRestoreNetworking struct {
	value *ResetOrRestoreNetworking
	isSet bool
}

func (v NullableResetOrRestoreNetworking) Get() *ResetOrRestoreNetworking {
	return v.value
}

func (v *NullableResetOrRestoreNetworking) Set(val *ResetOrRestoreNetworking) {
	v.value = val
	v.isSet = true
}

func (v NullableResetOrRestoreNetworking) IsSet() bool {
	return v.isSet
}

func (v *NullableResetOrRestoreNetworking) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResetOrRestoreNetworking(val *ResetOrRestoreNetworking) *NullableResetOrRestoreNetworking {
	return &NullableResetOrRestoreNetworking{value: val, isSet: true}
}

func (v NullableResetOrRestoreNetworking) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResetOrRestoreNetworking) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


